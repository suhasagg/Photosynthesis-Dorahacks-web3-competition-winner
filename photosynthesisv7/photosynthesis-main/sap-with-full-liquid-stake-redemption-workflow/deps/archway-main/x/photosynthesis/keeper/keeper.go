package keeper

import (
	"bufio"
	"bytes"
	"container/ring"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	epochsmodulekeeper "github.com/Stride-Labs/stride/v4/x/epochs/keeper"
	epochstypes "github.com/Stride-Labs/stride/v4/x/epochs/types"
	"github.com/archway-network/archway/x/photosynthesis/types"
	rewardKeeper "github.com/archway-network/archway/x/rewards/keeper"
	rewardstypes "github.com/archway-network/archway/x/rewards/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	dt "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Photosynthesis interface {
	// Query the redemption rate for the given epoch number
	QueryRedemptionRate(ctx sdk.Context, epochNumber uint64) (sdk.Dec, error)

	// Get the redemption rate query interval
	GetRedemptionRateQueryInterval(ctx sdk.Context) uint64

	// Get the redemption rate threshold
	GetRedemptionRateThreshold(ctx sdk.Context) sdk.Dec

	// Get the redemption interval threshold
	GetRedemptionIntervalThreshold(ctx sdk.Context) time.Duration

	// Get the latest redemption record
	GetLatestRedemptionRecord(ctx sdk.Context) (RedemptionRecord, bool)

	// Get the cumulative liquidity amount
	GetCumulativeLiquidityAmount(ctx sdk.Context) (sdk.Coins, error)

	// Get the total stake of all contracts
	GetTotalStake(ctx sdk.Context) (sdk.Int, error)

	// Get the stake of the given contract address
	GetStake(ctx sdk.Context, contractAddress sdk.AccAddress) (sdk.Int, error)

	// List all contracts
	ListContracts(ctx sdk.Context) ([]Contract, error)

	// Send tokens to the given contract address
	SendTokensToContract(ctx sdk.Context, contractAddress sdk.AccAddress, tokens sdk.Int) error

	// Redeem liquid tokens for the given amount
	RedeemLiquidTokens(ctx sdk.Context, amount sdk.Coins) error

	// Distribute redeemed tokens to the Dapps according to their stake
	DistributeRedeemedTokens(ctx sdk.Context, redeemedTokensAmount sdk.Coins) error

	// Delete the latest redemption record
	DeleteRedemptionRecord(ctx sdk.Context, record RedemptionRecord) error

	GetLatestRedemptionTime(ctx sdk.Context) time.Time

	//ExportGenesis(ctx sdk.Context) *types.GenesisState
	LiquidStake(ctx sdk.Context, epoch string, depositRecords []DepositRecord) error
	DistributeLiquidity(ctx sdk.Context, epoch string, depositRecords []DepositRecord) error
	EnqueueLiquidStakeRecord(ctx sdk.Context, record DepositRecord) error
	//InitGenesis(ctx sdk.Context, state *types.GenesisState)
	BeginBlocker(ctx sdk.Context) abci.ResponseBeginBlock
	EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate
}

// Keeper provides module state operations.
type PhotosynthesisKeeper struct {
	cdc              codec.Codec
	paramStore       paramTypes.Subspace
	storeKey         storetypes.StoreKey
	contractInfoView rewardKeeper.ContractInfoReaderExpected
	trackingKeeper   rewardKeeper.TrackingKeeperExpected
	authKeeper       rewardKeeper.AuthKeeperExpected
	bankKeeper       rewardKeeper.BankKeeperExpected
	bankkeeper       bankkeeper.Keeper
	rewardKeeper     rewardKeeper.Keeper
	epochKeeper      epochsmodulekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	IBCKeeper        ibckeeper.Keeper
}

type BroadcastReq struct {
	Tx   sdk.Tx `json:"tx"`
	Mode string `json:"mode"`
}

func NewPhotosynthesisKeeper(cdc codec.Codec, paramStore paramTypes.Subspace, storeKey storetypes.StoreKey,
	contractInfoView rewardKeeper.ContractInfoReaderExpected,
	trackingKeeper rewardKeeper.TrackingKeeperExpected,
	ak rewardKeeper.AuthKeeperExpected,
	bk rewardKeeper.BankKeeperExpected,
	bankkeeper bankkeeper.Keeper,
	rewardKeeper rewardKeeper.Keeper,
	epochKeeper epochsmodulekeeper.Keeper,
	TransferKeeper ibctransferkeeper.Keeper,
	IBCKeeper ibckeeper.Keeper) PhotosynthesisKeeper {

	return PhotosynthesisKeeper{
		cdc:        cdc,
		paramStore: paramStore,
		//state:            NewState(codec.Codec, storeKey),
		storeKey:         storeKey,
		contractInfoView: contractInfoView,
		trackingKeeper:   trackingKeeper,
		authKeeper:       ak,
		bankKeeper:       bk,
		bankkeeper:       bankkeeper,
		rewardKeeper:     rewardKeeper,
		epochKeeper:      epochKeeper,
		TransferKeeper:   TransferKeeper,
		IBCKeeper:        IBCKeeper,
	}
}

type StreamType byte

const (
	// Stdout represents standard out stream type
	Stdout StreamType = 1
	// Stderr represents standard error streap type
	Stderr StreamType = 2
)

func demultiplexDockerStream(reader io.Reader, stdoutWriter, stderrWriter io.Writer) error {
	buffer := make([]byte, 8)
	for {
		_, err := io.ReadFull(reader, buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		streamType := StreamType(buffer[0])
		length := binary.BigEndian.Uint32(buffer[4:8])

		// Create a limited reader so that only `length` bytes are read for this message.
		limitedReader := &io.LimitedReader{R: reader, N: int64(length)}

		var writer io.Writer
		switch streamType {
		case Stdout:
			writer = stdoutWriter
		case Stderr:
			writer = stderrWriter
		default:
			return fmt.Errorf("invalid stream type: %v", streamType)
		}

		if _, err = io.Copy(writer, limitedReader); err != nil {
			return err
		}
	}
	return nil
}

type DockerPool struct {
	Client *client.Client
}

func NewDockerPool() (*DockerPool, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerPool{Client: cli}, nil
}

func (dp *DockerPool) CreateAndStartExec(ctx context.Context, containerID string, command []string) ([]byte, string) {
	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	config := dt.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		User:         "root",
		Cmd:          command, // Replace with your actual command and args
	}

	// Create an exec instance
	response, err := cli.ContainerExecCreate(ctx, containerID, config)
	if err != nil {
		panic(err)
	}

	// Create output buffers
	var stdout, stderr bytes.Buffer

	// Attach to the exec instance
	hijackedResponse, err := cli.ContainerExecAttach(ctx, response.ID, dt.ExecStartCheck{})
	if err != nil {
		panic(err)
	}
	defer hijackedResponse.Close()

	// Copy the output to our buffers
	outputDone := make(chan error)
	go func() {
		err = demultiplexDockerStream(hijackedResponse.Reader, &stdout, &stderr)
		outputDone <- err
	}()

	// Wait for the command to finish
	select {
	case err := <-outputDone:
		if err != nil {
			panic(err)
		}
	case <-ctx.Done():
		panic(ctx.Err())
	}

	// Inspect the exec instance to get the exit code
	inspectResponse, err := cli.ContainerExecInspect(ctx, response.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Exit code: %d\n", inspectResponse.ExitCode)
	fmt.Printf("Stdout: %s\n", stdout.String())
	fmt.Printf("Stderr: %s\n", stderr.String())
	return stdout.Bytes(), stderr.String()
}

// struct to hold the YAML data for address account balance for different crypto tokens
type Balance struct {
	Balances []struct {
		Amount string `yaml:"amount"`
		Denom  string `yaml:"denom"`
	} `yaml:"balances"`
}

type DepositRecord struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
	Epoch           int64          `json:"epoch"`
	Amount          int64          `json:"amount"`
	Status          string         `json:"status"`
}

// RedemptionRecord defines a redemption record structure
type RedemptionRecord struct {
	Timestamp       int64   `json:"timestamp" yaml:"timestamp"`
	LiquidityAmount sdk.Int `json:"liquidity_amount" yaml:"liquidity_amount"`
}

// NewRedemptionRecord creates a new RedemptionRecord instance
func NewRedemptionRecord(timestamp int64, liquidityAmount sdk.Int) RedemptionRecord {
	return RedemptionRecord{
		Timestamp:       timestamp,
		LiquidityAmount: liquidityAmount,
	}
}

// Contract defines a simple contract structure
type Contract struct {
	Address   sdk.AccAddress `json:"id" yaml:"id"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Name      string         `json:"name" yaml:"name"`
	Stake     int64          `json:"stake" yaml:"stake"`
	Rewards   int64          `json:"rewards" yaml:"rewards"`
	Activated bool           `json:"activated" yaml:"activated"`
}

// NewContract creates a new Contract instance
func NewContract(address sdk.AccAddress, creator sdk.AccAddress, name string, stake int64, rewards int64, activated bool) Contract {
	return Contract{
		Address:   address,
		Creator:   creator,
		Name:      name,
		Stake:     stake,
		Rewards:   rewards,
		Activated: activated,
	}
}

func trimToFirstNDirectoryLevels(path string, n int) string {
	parts := strings.Split(path, string(filepath.Separator))
	if len(parts) <= n {
		return path
	}
	return strings.Join(parts[:len(parts)-n], string(filepath.Separator))
}

func trimToLastNDirectoryLevels(path string, n int) string {
	parts := strings.Split(path, string(filepath.Separator))
	if len(parts) <= n {
		return path
	}
	return strings.Join(parts[len(parts)-n:], string(filepath.Separator))
}

func trimAndRemoveSpecialChars(str string) string {
	// Custom function to remove special characters
	removeSpecialChars := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}

	// Trim the string and remove special characters
	trimmedStr := strings.Map(removeSpecialChars, str)

	return trimmedStr
}

// GetLatestRedemptionTime retrieves the latest redemption time from the store
func (k PhotosynthesisKeeper) GetLatestRedemptionTime(ctx sdk.Context) time.Time {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(types.LatestRedemptionTimeKey)) {
		return time.Time{}
	}

	bz := store.Get([]byte(types.LatestRedemptionTimeKey))
	latestRedemptionTime := int64(binary.BigEndian.Uint64(bz))
	return time.Unix(latestRedemptionTime, 0)
}

// Logger returns a module-specific logger.
func (k PhotosynthesisKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// CreateContractLiquidStakeDepositRecordsForEpoch creates a new deposit record for the given contract and epoch
func (k PhotosynthesisKeeper) CreateContractLiquidStakeDepositRecordsForEpoch(ctx sdk.Context, state rewardKeeper.State, rewardAddress sdk.AccAddress, epoch int64) *types.DepositRecord {
	// Determine the contract's liquid stake deposit amount for the given epoch
	// This will depend on your specific application logic and may involve calculations or querying other modules
	amount := k.GetCumulativeRewardAmount(ctx, epoch, rewardAddress)

	// Create a new deposit record with the appropriate fields
	depositRecord := types.DepositRecord{
		ContractAddress: string(rewardAddress.Bytes()),
		Epoch:           epoch,
		Amount:          amount,
		Status:          "pending",
	}
	k.Logger(ctx).Info(fmt.Sprintf("Reward Address %v, Epoch %v, Liquid stake deposit record created amount %v, record status %v", rewardAddress, epoch, amount, "pending"))
	return &depositRecord
}

// Implement the EnqueueLiquidStakeRecord method
func (k PhotosynthesisKeeper) EnqueueLiquidStakeRecord(ctx sdk.Context, record *types.DepositRecord) error {
	// Implement the logic for enqueuing liquid stake deposit records here
	// For example, you can store the deposit records in a store using contract addresses as keys
	store := ctx.KVStore(k.storeKey)
	contractAddress := record.ContractAddress
	recordsBytes := store.Get([]byte(contractAddress))
	var records types.DepositRecords
	records.Records = make([]*types.DepositRecord, 0)

	if recordsBytes != nil {
		k.cdc.MustUnmarshal(recordsBytes, &records)
	}
	records.Records = append(records.Records, record)

	store.Set([]byte(contractAddress), k.cdc.MustMarshal(&records))
	return nil
}

/*
// calculateContractLiquidStakeAmount calculates the contract's liquid stake deposit amount for the given epoch
func (k PhotosynthesisKeeper) calculateContractLiquidStakeAmountforEpoch(ctx sdk.Context, contractAddress sdk.AccAddress, epoch int64) sdk.Int {
	// Retrieve the total rewards earned by the contract
	/*
	totalRes :=



	// Define the percentage of rewards to be used for liquid staking
	liquidStakingPercentage := sdk.NewDecWithPrec(10, 2) // 10% for example

	// Calculate the liquid stake deposit amount based on the percentage
	liquidStakeAmount := totalRewards.ToDec().Mul(liquidStakingPercentage).TruncateInt()

	return liquidStakeAmount


	return nil
}
*/

func clearFileContents(filename string) error {
	// Open the file in write mode, truncating it if it exists or creating a new file if it doesn't
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write an empty string to clear the file contents
	_, err = file.WriteString("")
	if err != nil {
		return err
	}

	return nil
}

// GetContractLiquidStakeDepositsTillEpoch retrieves all deposit records for a given contract until the specified epoch
func (k *PhotosynthesisKeeper) GetContractLiquidStakeDepositsTillEpoch(ctx sdk.Context, contractAddress sdk.AccAddress, epoch int64) ([]*types.DepositRecord, error) {
	store := ctx.KVStore(k.storeKey)
	if len(contractAddress) == 0 {
		return nil, nil
	}
	recordsBytes := store.Get(contractAddress)
	//var records *types.DepositRecords
	var records types.DepositRecords
	records.Records = make([]*types.DepositRecord, 0)

	if recordsBytes != nil {
		k.cdc.MustUnmarshal(recordsBytes, &records)
	}

	var depositsTillEpoch []*types.DepositRecord
	//	if records != nil {
	for _, record := range records.Records {
		if record.Epoch <= epoch {
			depositsTillEpoch = append(depositsTillEpoch, record)
		}
	}
	//	}

	return depositsTillEpoch, nil
}

func (k PhotosynthesisKeeper) GetTotalEnqueuedRewards(ctx sdk.Context, epoch int64) (sdk.Int, error) {
	totalLiquidStake := sdk.ZeroInt()
	// Iterate through all contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Retrieve deposit records for the contract
		depositRecords, err := k.GetContractLiquidStakeDepositsTillEpoch(ctx, sdk.AccAddress(meta.RewardsAddress), epoch)
		if err != nil {
			return true
		}

		// Sum up the liquid stake for the contract
		contractLiquidStake := sdk.ZeroInt()

		for _, record := range depositRecords {
			if record.Status == "pending" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
			}
			k.Logger(ctx).Info(fmt.Sprintf("Contract Address %v, Epoch %v, Liquid stake deposit record enqueued amount %v, record status %v", meta.ContractAddress, record.Epoch, record.Amount, record.Status))
		}
		// Add the contract's liquid stake to the total liquid stake
		totalLiquidStake = totalLiquidStake.Add(contractLiquidStake)
		return false
	})

	return totalLiquidStake, nil
}

func (k PhotosynthesisKeeper) GetTotalLiquidStake(ctx sdk.Context, epoch int64) (sdk.Int, error) {
	totalLiquidStake := sdk.ZeroInt()
	store := ctx.KVStore(k.storeKey)
	// Iterate through all contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Retrieve deposit records for the contract
		depositRecords, err := k.GetContractLiquidStakeDepositsTillEpoch(ctx, sdk.AccAddress(meta.RewardsAddress), epoch)
		k.Logger(ctx).Info(fmt.Sprintf("Reward Address %v, Epoch %v", meta.RewardsAddress, epoch))

		if err != nil {
			return true
		}

		// Sum up the liquid stake for the contract
		contractLiquidStake := sdk.ZeroInt()
		var updatedRecords *types.DepositRecords
		updatedRecords = &types.DepositRecords{}
		for _, record := range depositRecords {
			if record.Status == "pending" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
				record.Status = "completed"
				updatedRecords.Records = append(updatedRecords.Records, record)
			}
			k.Logger(ctx).Info(fmt.Sprintf("Contract Address %v, Epoch %v, Liquid stake deposit record amount %v, record status %v", meta.ContractAddress, record.Epoch, record.Amount, record.Status))
		}
		if len(updatedRecords.Records) > 0 {
			store.Set([]byte(meta.RewardsAddress), k.cdc.MustMarshal(updatedRecords))
		}
		// Add the contract's liquid stake to the total liquid stake
		totalLiquidStake = totalLiquidStake.Add(contractLiquidStake)
		k.Logger(ctx).Info(fmt.Sprintf("Total liquid stake %v", totalLiquidStake))
		return false
	})

	return totalLiquidStake, nil
}

func (k PhotosynthesisKeeper) LiquidStake(ctx sdk.Context, epoch int64, tls int64) (sdk.Int, error) {
	k.Logger(ctx).Info(fmt.Sprintf("Liquid stake amount: %d \n", tls))  // Logging the message "Liquid stake amount: " followed by the `tls` value
	text := strconv.FormatInt(tls, 10)                                  // Converting `tls` to string type and storing it in `text`
	k.Logger(ctx).Info(fmt.Sprintf("Liquid stake amount: %s \n", text)) // Logging the message "Liquid stake amount: " followed by the `text` value
	epochText := strconv.FormatInt(epoch, 10)
	// Convert the text to a byte slice because WriteFile requires a byte slice
	data := []byte(text + "\n") // Converting `text` to a byte slice and adding a newline character
	epochdata := []byte(epochText + "\n")
	// Open the file in append mode, create it if it does not exist
	file, err := os.OpenFile("/home/photo/logs/liquidstakeparameters", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	_, err = file.Write(epochdata)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
	}

	file1, err := os.OpenFile("/home/photo/logs/distributionepoch", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	defer file1.Close()

	// Write the data to the file
	_, err = file1.Write(epochdata)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
	}

	/*
		goCtx := sdk.WrapSDKContext(ctx)
		msg := ibctypes.NewMsgTransfer("transfer", "channel-0", sdk.NewCoin("uarch", sdk.NewInt(500)), "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", "stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8", clienttypes.NewHeight(0, uint64(ctx.BlockHeight())+100), 0)
		msgTransferResponse, err := k.TransferKeeper.Transfer(goCtx, msg)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error in LiquidStake: %s", err))
		}
		sequence := msgTransferResponse.Sequence

		k.Logger(ctx).Error(fmt.Sprintf("Sequence in LiquidStake: %s", sequence))
	*/
	/*

		ctx1, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		cmd := []string{
			"archwayd",
			"archwayd",
			"tx",
			"--home=/home/photo/.photo",
			"--node=http://localhost:26657",
			"ibc-transfer",
			"transfer",
			"transfer",
			"channel-0",
			"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8",
			"10uarch", // Assuming tls is a string variable holding some value
			"--chain-id=localnet",
			"--from=pval1",
			"--log_level=trace",
			"--keyring-backend=test",
			"--broadcast-mode=sync",
			"--packet-timeout-height=0-0",
			"--output=json",
			"-y",
		}

		dp, err := NewDockerPool()
		if err != nil {

		}

		containers, err := dp.Client.ContainerList(ctx1, dt.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		var targetContainerID string
		targetImage := "stridezone:photo"
		for _, container := range containers {
			if strings.Contains(container.Image, targetImage) {
				targetContainerID = container.ID
				break // Exit the loop after finding the desired container
			}
		}

		if targetContainerID != "" {
			k.Logger(ctx).Info(fmt.Sprintf("Container ID for %s: %s", "photosynthesis", targetContainerID))
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("Container with name %s not found.", "photosynthesis"))
		}

		dp.CreateAndStartExec(ctx1, targetContainerID, cmd)
	*/

	/*
		script := "/media/usbHDD/deps/archway-main/contrib/localnet/opt/add_cron.sh"

		cmd := exec.Command("/bin/sh", script)

		// This is optional, but you might want to forward the output
		// of the command to the standard output of your program
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error starting the command: %s", err))
			//return
		}

		err = cmd.Wait()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				// Command exited with non-zero status
				errorOutput := string(exitError.Stderr) // Extract error output
				k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
				k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
			} else {
				// Other error occurred
				k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			}
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("Command ran successfully"))
		}

		/*
			go func() {
				script := "/media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh"

				cmd := exec.Command("/bin/sh", script)

				// This is optional, but you might want to forward the output
				// of the command to the standard output of your program
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Start()
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf("Error starting the command: %s", err))
					//return
				}

				err = cmd.Wait()
				if err != nil {
					if exitError, ok := err.(*exec.ExitError); ok {
						// Command exited with non-zero status
						errorOutput := string(exitError.Stderr) // Extract error output
						k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
						k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
					} else {
						// Other error occurred
						k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
					}
				} else {
					k.Logger(ctx).Info(fmt.Sprintf("Command ran successfully"))
				}
			}()
	*/
	// replace this with the path to your script
	/*
		script := "/media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh"

		cmd := exec.Command("/bin/sh", script)

		// This is optional, but you might want to forward the output
		// of the command to the standard output of your program
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			k.Logger(ctx).Info(fmt.Sprintf("Command ran successfully"))
		}

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				// Command exited with non-zero status
				errorOutput := string(exitError.Stderr) // Extract error output
				k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
				k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
			} else {
				// Other error occurred
				k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			}

		}
	*/

	//ls := strconv.FormatInt(tls, 10)
	/*
		    dir, err := os.Getwd()
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
				return sdk.ZeroInt(), err
			} else {
				k.Logger(ctx).Info("Current Directory: ", dir)
			}
			//trimmedPath := trimToFirstNDirectoryLevels(dir, 5)
			//newPath := filepath.Join(trimmedPath, "dockernet/state/photo1")
			cmd := exec.Command("archwayd",
				"--home",
				"/home/photo/.photo",
				"tx",
				"--node",
				"tcp://localhost:26657",
				"ibc-transfer",
				"transfer",
				"transfer",
				"channel-0",
				"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8",
				ls+"uarch",
				"--from",
				"pval1",
				"-y",
			)

			// Execute the command
			out, err := cmd.Output()

			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					// Command exited with non-zero status
					errorOutput := string(exitError.Stderr) // Extract error output
					k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
					k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
				} else {
					// Other error occurred
					k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
				}
				return sdk.ZeroInt(), err
			}

			//fmt.Printf("The output of the command is: \n%s\n", out)
			k.Logger(ctx).Info("The output of the command is: \n%s\n", out)
			dir, err = os.Getwd()
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			} else {
				k.Logger(ctx).Info("Current Directory: ", dir)
			}
			//trimmedPath = trimToFirstNDirectoryLevels(dir, 5)
			//newPath = filepath.Join(trimmedPath, "dockernet/state/stride1")

			cmd1 := exec.Command("strided",
				"--home",
				"/home/stride/.stride",
				"--node",
				"http://localhost:26657",
				"tx",
				"stakeibc",
				"liquid-stake",
				ls,
				"uarch",
				"--keyring-backend",
				"test",
				"--from",
				"admin",
				"--chain-id",
				"STRIDE",
				"-y",
			)

	*/
	// Open a log file
	/*
		logFile, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening log file: %v", err)
		}
		defer logFile.Close()

		// Set up a log writer and direct command outputs to it
		logWriter := bufio.NewWriter(logFile)
		defer logWriter.Flush()
		var wg sync.WaitGroup
		// Add 1 to WaitGroup counter
		wg.Add(1)

		// Launch a goroutine for first command
		go func() {
			defer wg.Done() // Decrease counter when goroutine completes
			/*
				cmd := exec.Command("archwayd",
					"--home",
					"/home/photo/.photo",
					"tx",
					"--node",
					"http://localhost:26657",
					"ibc-transfer",
					"transfer",
					"transfer",
					"channel-0",
					"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8",
					"10"+"uarch",
					"--from",
					"pval1",
					"-y",
				)

			ctx := client.Context{}.WithJSONMarshaler(nil).WithCodec(sdk.NewCodec())


		// Set the source and destination chain IDs
			sourceChainID := "localnet"
			destinationChainID := "STRIDE"

			// Set the source and destination channel IDs
			sourceChannelID := "channel-0"
			destinationChannelID := "channel-1"

			// Set the sender and recipient addresses
			sender := sdk.AccAddress([]byte("sender-address"))
			recipient := sdk.AccAddress([]byte("recipient-address"))

			// Set the amount and denom of the transfer
			amount := sdk.NewInt(100) // 1 token
			denom := "uarch"

			// Build the MsgTransfer message
			msg := ibcchannel.NewMsgTransfer(
				sourceChannelID,
				destinationChannelID,
				amount,
				denom,
				sender,
				recipient,
				fmt.Sprintf("%d", time.Now().UnixNano()),
			)

			// Create a new CLI command to send the message
			cmd := &sdk.Tx{
				Msgs: []sdk.Msg{msg},
				Fee:  sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(5000))), // Fee in uatom
				Memo: "IBC Transfer",
			}

			// Sign and broadcast the transaction
			res, err := authclient.SignAndBroadcastTx(ctx, cmd, sender)
			if err != nil {
				fmt.Printf("Failed to send IBC transfer: %s", err)
				return
			}

	*/
	/*
		go func() {
			// Create a transfer client.
			client := transfer.NewClient(ctx)

			// Specify the transfer parameters.
			srcPort := "transfer"
			srcChannel := "channel-0"
			receiver := "stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"
			amount := int64(100)
			denom := "uarch"

			// Send the transfer message.
			err := client.SendTransfer(srcPort, srcChannel, receiver, amount, denom)
			if err != nil {
				log.Fatalf("Failed to send transfer: %v", err)
			}

			// Set the command's output to our log writer
			//	logWriter := log.Writer(os.Stdout)
			//	cmd := exec.Command("some command") // Replace "some command" with your command.
			//	cmd.Stdout = logWriter
			//		cmd.Stderr = logWriter

			// Execute the command
			//	err = cmd.Run()
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					// Command exited with non-zero status
					errorOutput := string(exitError.Stderr) // Extract error output
					log.Printf("Command failed with exit code: %d\n", exitError.ExitCode())
					log.Printf("Error output: %s", errorOutput)
				} else {
					// Other error occurred
					log.Fatalf("Command execution failed: %v", err)
				}
			}

			// Get current working directory
			dir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Failed to get current directory: %v", err)
			} else {
				log.Printf("Current Directory: %s", dir)
			}
		}()
	*/

	/*
		// Wait for the first goroutine to finish
		wg.Wait()

		// Add 1 to WaitGroup counter
		wg.Add(1)

		// Launch another goroutine for the second command
		go func() {
			defer wg.Done() // Decrease counter when goroutine completes

			cmd1 := exec.Command("strided",
				"--home",
				"/home/stride/.stride",
				"--node",
				"http://localhost:26657",
				"tx",
				"stakeibc",
				"liquid-stake",
				ls,
				"uarch",
				"--keyring-backend",
				"test",
				"--from",
				"admin",
				"--chain-id",
				"STRIDE",
				"-y",
			)

			cmd1.Stdout = logWriter
			cmd1.Stderr = logWriter
			// Run the command and capture its output
			err1 := cmd1.Run()

			// Set the command's output to our log writer

			if err1 != nil {
				if exitError, ok := err1.(*exec.ExitError); ok {
					// Command exited with non-zero status
					errorOutput := string(exitError.Stderr) // Extract error output
					k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
					k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
					k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
				} else {
					// Other error occurred
					k.Logger(ctx).Error(fmt.Sprintf("Error:", err1))
				}
				return
			}

			//fmt.Printf("The output of the command is: \n%s\n", out)
			//k.Logger(ctx).Info("The output of the command is: \n%s\n", out)
			//dir, err = os.Getwd()
			//if err != nil {
			//	k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			//} else {
			//	k.Logger(ctx).Info("Current Directory: ", dir)
			//}
		}()

		// Wait for the second goroutine to finish
		wg.Wait()

		// Run the command and capture its output

		//trimmedPath = trimToFirstNDirectoryLevels(dir, 5)
		//newPath = filepath.Join(trimmedPath, "dockernet/state/stride1")
		cmd2 := exec.Command(
			"strided",
			"--node",
			"http://localhost:26657",
			"--home",
			"/home/stride/.stride",
			"q",
			"bank",
			"balances",
			"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8",
		)

		// Execute the command and capture its output
		out, err := cmd2.Output()
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			return sdk.ZeroInt(), err

		}
		data := Balance{}
		err = yaml.Unmarshal(out, &data)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
			return sdk.ZeroInt(), err
		}

		// Find the balance for stuarch
		for _, balance := range data.Balances {
			if balance.Denom == "stuarch" {
				fmt.Println("The balance for stuarch is:", balance.Amount)
				val, err := strconv.ParseInt(balance.Amount, 10, 64)
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf("Error:", err))
					return sdk.ZeroInt(), err
				}
				return sdk.NewInt(val), nil
			}
		}
	*/

	//	var liquidityAmount int64
	//	liquidityAmount := 10
	//	amount, err := k.GetTotalLiquidStake(ctx, epoch)
	// Transfer reward funds from Archway to liquidity provider
	//TODO STRIDE INTERACTION
	//err1 := k.TransferRewardFunds(ctx, contract.ArchwayRewardFundsTransferAddress, contract.LiquidityProviderAddress, amount)
	//if err != nil {
	//	return err1
	//}
	//	return sdk.ZeroInt(), nil
	//Distribute liquidity tokens to DApps
	//k.DistributeLiquidity(ctx, epoch, tls)
	return sdk.ZeroInt(), nil
}

func (k PhotosynthesisKeeper) DistributeLiquidity(ctx sdk.Context, epoch int64, liquidityAmount int64) {
	// Get the total stake amount from all deposit records
	//liquidityamount := k.bankkeeper.GetAllBalances(ctx, sdk.AccAddress("archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"))
	/*
		var wg sync.WaitGroup

		wg.Add(1)
		logWriter := log.Writer(os.Stdout)
		// Launch another goroutine for the second command
		go func() {
			defer wg.Done()



			cmd.Stdout = logWriter
			cmd.Stderr = logWriter
			// Run the command and capture its output
			err1 := cmd.Run()

			// Set the command's output to our log writer

			if err1 != nil {
				if exitError, ok := err1.(*exec.ExitError); ok {
					// Command exited with non-zero status
					errorOutput := string(exitError.Stderr) // Extract error output
					k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
					k.Logger(ctx).Error(fmt.Sprintf("Command failed with exit code: %d\n", exitError.ExitCode()))
					k.Logger(ctx).Error(fmt.Sprintf("Error output:", errorOutput))
				} else {
					// Other error occurred
					k.Logger(ctx).Error(fmt.Sprintf("Error:", err1))
				}
				return
			}
		    }()
		    wg.Wait()
	*/
	/*
		ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cmd := []string{
			"archwayd",
			"--chain-id",
			"localnet",
			"--node",
			"http://localhost:26657",
			"--home",
			"/home/photo/.photo",
			"q",
			"bank",
			"balances",
			"archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m",
		}

		dp, err := NewDockerPool()
		if err != nil {

		}

		containers, err := dp.Client.ContainerList(ctx1, dt.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		var targetContainerID string
		targetImage := "stridezone:photo"
		for _, container := range containers {
			if strings.Contains(container.Image, targetImage) {
				targetContainerID = container.ID
				break // Exit the loop after finding the desired container
			}
		}

		if targetContainerID != "" {
			k.Logger(ctx).Info(fmt.Sprintf("Container ID for %s: %s", "photosynthesis", targetContainerID))
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("Container with name %s not found.", "photosynthesis"))
		}

		out, error := dp.CreateAndStartExec(ctx1, targetContainerID, cmd)
		k.Logger(ctx).Error(fmt.Sprintf(error))
		data := Balance{}
		err = yaml.Unmarshal(out, &data)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error:", err))

		}

		// Find the balance for stuarch
		for _, balance := range data.Balances {
			if balance.Denom == "ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145" {
				fmt.Println("The balance for stuarch is:", balance.Amount)
				val, err := strconv.ParseInt(balance.Amount, 10, 64)
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf("Error:", err))

				}
				liquidityAmount = val
				k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Token Amount: %v", "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", liquidityAmount))

			}
		}
	*/

	//if liquidityamount.AmountOf("stuarch").Int64() == 0 {
	//	//	return
	//	}

	file, err := os.Open("/home/photo/logs/liquidstakelogs") // change filename.txt with your filename
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	defer file.Close()

	amountRegex, _ := regexp.Compile(`amount: "(\d+)"`)
	denomRegex, _ := regexp.Compile(`denom: ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145`)

	var amount string

	// Initialize a ring buffer with 100 strings
	r := ring.New(1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Always insert the new line at the end of the ring
		r.Value = line
		r = r.Next()

		if amountMatch := amountRegex.FindStringSubmatch(line); amountMatch != nil {
			amount = amountMatch[1]
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Convert ring to slice
	lines := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		if r.Value != nil {
			lines[i] = r.Value.(string)
		} else {
			lines[i] = ""
		}
		r = r.Next()
	}

	// Iterate in reverse
	for i := len(lines) - 1; i >= 0; i-- {
		if denomMatch := denomRegex.FindString(lines[i]); denomMatch != "" {
			k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Tokens Obtained Amount: %v", "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", amount))
			liquidityAmount, err = strconv.ParseInt(amount, 10, 64)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
			}
		}
	}

	totalStake := sdk.NewInt(0)
	// Calculate the cumulative stake for each contract
	cumulativeStakes := make(map[string]sdk.Int)
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Retrieve deposit records for the contract

		depositRecords, err := k.GetContractLiquidStakeDepositsTillEpoch(ctx, sdk.AccAddress(meta.RewardsAddress), epoch)
		k.Logger(ctx).Info(fmt.Sprintf("Reward Address %v, Till Epoch considered for distribution %v", meta.RewardsAddress, epoch))
		if err != nil {
			return true
		}
		k.Logger(ctx).Info(fmt.Sprintf("Deposit Record Stake ratio determination for liquidity distribution %v", depositRecords))
		// Sum up the liquid stake for the contract
		contractLiquidStake := sdk.ZeroInt()
		for _, record := range depositRecords {
			if record.Status == "completed" {
				cumulativeStakes[meta.LiquidityProviderAddress] = sdk.ZeroInt()
			}
		}
		for _, record := range depositRecords {
			if record.Status == "completed" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
				cumulativeStakes[meta.LiquidityProviderAddress] = cumulativeStakes[meta.LiquidityProviderAddress].Add(sdk.NewInt(record.Amount))
				k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s,  Liquid Token Distribution Amount : %d\n, Epoch %d\n", meta.LiquidityProviderAddress, record.Amount, record.Epoch))
				totalStake = totalStake.Add(sdk.NewInt(record.Amount))
				k.Logger(ctx).Info(fmt.Sprintf("Total Stake: %v", totalStake))
			}
		}
		err = k.DeleteLiquidStakeDepositRecord(ctx, sdk.AccAddress(meta.RewardsAddress))
		if err != nil {
			return true
		}
		if totalStake.IsZero() {
			return
		}
		return false
	})

	file, err1 := os.OpenFile("/home/photo/logs/liquiditydistributiontodapps", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // change filename.txt with your filename
	if err1 != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err1.Error()))
	}
	defer file.Close()

	// Distribute the liquidity tokens to each contract proportionally
	for contractAddr, contractStake := range cumulativeStakes {
		// Calculate the proportion of the stake for the current contract
		stakeProportion := sdk.NewDecFromInt(contractStake).Quo(sdk.NewDecFromInt(totalStake))
		k.Logger(ctx).Info(fmt.Sprintf("Stake proportion: %s", stakeProportion))
		stakeratio, err := stakeProportion.Float64()
		// Calculate the amount of liquidity tokens to distribute for the current contract
		liquidityTokensAmount := stakeratio * float64(liquidityAmount)
		// Distribute the calculated amount of liquidity tokens to the contract's liquidity token address
		contractAddress, err := sdk.AccAddressFromBech32(contractAddr)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error()))
		}
		//log.Printf("Contract Address: %s, Liquid Token Amount: %d\n", contractAddress, liquidityTokensAmount)
		k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Token Amount: %d\n", contractAddress, liquidityTokensAmount))
		//err = k.bankkeeper.SendCoins(ctx, sdk.AccAddress("archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"), sdk.AccAddress(contractAddr), sdk.NewCoins(sdk.NewCoin("ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145", sdk.NewInt(int64(liquidityTokensAmount)))))
		data := "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m" + "," + contractAddr + "," + "ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145" + "," + strconv.FormatFloat(liquidityTokensAmount, 'f', 2, 64)
		liquiditydistributiontransferlog := []byte(data + "\n")
		_, err = file.Write(liquiditydistributiontransferlog)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
		}

	}

}

// DeleteLiquidStake DepositRecord deletes completed deposit records for a given contract
func (k *PhotosynthesisKeeper) DeleteLiquidStakeDepositRecord(ctx sdk.Context, contractAddress sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	var recordsBytes []byte
	if len(contractAddress) != 0 {
		recordsBytes = store.Get(contractAddress)
	}

	var records types.DepositRecords
	records.Records = make([]*types.DepositRecord, 0)

	if len(recordsBytes) != 0 {
		k.cdc.MustUnmarshal(recordsBytes, &records)
	}

	var updatedRecords *types.DepositRecords
	updatedRecords = &types.DepositRecords{}
	for _, record := range records.Records {
		if record.Status != "completed" {
			updatedRecords.Records = append(updatedRecords.Records, record)
		}
	}

	if len(updatedRecords.Records) > 0 {
		store.Set(contractAddress.Bytes(), k.cdc.MustMarshal(updatedRecords))
	}
	k.Logger(ctx).Info(fmt.Sprintf("Deleted Liquid stake deposit Record"))

	return nil
}

// RedeemLiquidTokens redeems liquid tokens and distributes them accordingly
func (k PhotosynthesisKeeper) RedeemLiquidTokens(ctx sdk.Context, epoch int64, cumLiquidityAmount *types.Coin) (int64, error) {
	// Get the list of contracts
	//contracts, err := k.ListContracts(ctx)
	//	if err != nil {
	//		return 0, err
	//	}
	var redeemedAmount int64
	// Iterate over the contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		k.Logger(ctx).Info(fmt.Sprintf("Redeem Liquid tokens: %d \n", cumLiquidityAmount)) // Logging the message "Liquid stake amount: " followed by the `tls` value
		text := strconv.FormatInt(cumLiquidityAmount.Amount, 10)                           // Converting `tls` to string type and storing it in `text`
		k.Logger(ctx).Info(fmt.Sprintf("Redeem Liquid tokens: %s \n", text))               // Logging the message "Liquid stake amount: " followed by the `text` value

		// Convert the text to a byte slice because WriteFile requires a byte slice
		data := []byte(text + "\n") // Converting `text` to a byte slice and adding a newline character
		epochdata := []byte(string(epoch) + "\n")
		// Open the file in append mode, create it if it does not exist
		file, err := os.OpenFile("/home/photo/logs/enableredeemstake", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
		}
		defer file.Close()

		// Write the data to the file
		_, err = file.Write(data)
		_, err1 := file.Write(epochdata)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
		}
		if err1 != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err1.Error())) // Logging the error message if an error occurred
		}

		file1, err := os.OpenFile("/home/photo/logs/redeemepoch", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
		}
		defer file1.Close()

		// Write the data to the file
		_, err = file1.Write(epochdata)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
		}

		// Calculate the redeemed amount for each contract
		//meta := k.rewardKeeper.GetContractMetadata(ctx, sdk.AccAddress(meta.ContractAddress))
		//redeemedAmount = int64(meta.RedemptionRateThreshold) * cumLiquidityAmount.Amount
		//coin := sdk.NewCoins(sdk.NewCoin("", sdk.NewInt(redeemedAmount)))
		// Transfer the redeemed tokens from the module account to the contract address
		//err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(meta.RewardsAddress), coin)
		//if err != nil {
		//	return true
		//}

		// Update the contract's stake and rewards
		//contract.Stake = contract.Stake + redeemedAmount
		//contract.Rewards = contract.Rewards + redeemedAmount
		//	err = k.SaveContract(ctx, contract)
		//	if err != nil {
		//			return 0, err
		//		}
		return false
	})

	return redeemedAmount, nil
}

// GetRedemptionRateQueryInterval retrieves the redemption rate query interval from the keeper's parameter store
func (k PhotosynthesisKeeper) GetRedemptionRateQueryInterval(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	var uint64Value uint64
	redemptionRateQueryInterval := store.Get([]byte(types.KeyRedemptionRateQueryInterval))
	err := binary.Read(bytes.NewReader(redemptionRateQueryInterval), binary.BigEndian, &uint64Value)

	if err != nil {
		//log.Fatalf("Failed to convert byte array to uint64: %v", err)
		k.Logger(ctx).Error(fmt.Sprintf("Failed to convert byte array to uint64: %v", err))
	}
	return uint64Value
}

// ListContracts retrieves all stored contracts from the store.
func (k PhotosynthesisKeeper) ListContracts(ctx sdk.Context) ([]*types.Contract, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ContractPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ContractPrefix))

	defer iterator.Close()

	var contracts []*types.Contract
	for ; iterator.Valid(); iterator.Next() {
		var contract *types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), contract)
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (k PhotosynthesisKeeper) SaveContract(ctx sdk.Context, contract *types.Contract) error {
	// Convert the contract address to a store key
	store := ctx.KVStore(k.storeKey)
	key := []byte("contract-" + contract.Address)

	// Marshal the contract
	bz := k.cdc.MustMarshal(contract)

	// Save the contract to the store
	store.Set(key, bz)
	return nil
}

// ProcessRedemptionRateQueries queries the redemption rate at specific epochs and checks
// if the redemption rate is above a threshold. If the rate is above the threshold, it
// initiates the redemption process and distributes the redeemed tokens to Dapps according
// to their stake. It also deletes the latest redemption record.
func (k PhotosynthesisKeeper) ProcessRedemptionRateQueries(ctx sdk.Context, epoch int64, redemptionRateQueryInterval uint64) (float64, error) {
	info, _ := k.epochKeeper.GetEpochInfo(ctx, epochstypes.DAY_EPOCH)

	//redemptionRateQueryInterval := k.GetRedemptionRateQueryInterval(ctx)

	if info.CurrentEpoch%int64(redemptionRateQueryInterval) != 0 {
		return 0.0, nil
	}

	redemptionRate, err := k.QueryRedemptionRate(ctx)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error %s", err))
	}
	var redeemedAmount float64

	redemptionRateThreshold := 1
	if redemptionRate >= float64(redemptionRateThreshold) {
		timeSinceLatestRedemption := ctx.BlockTime().Sub(k.GetLatestRedemptionTime(ctx))
		k.Logger(ctx).Info(fmt.Sprintf("Time since last redemption: %v \n", timeSinceLatestRedemption))
		if timeSinceLatestRedemption.Milliseconds() >= int64(redemptionRateQueryInterval) {
			cumLiquidityAmount, _ := k.GetCumulativeLiquidityAmount(ctx, epoch, "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m")
			k.Logger(ctx).Info(fmt.Sprintf("Cumulative Liquidity Amount %d \n", cumLiquidityAmount))
			_, err := k.RedeemLiquidTokens(ctx, epoch, &types.Coin{Amount: int64(0)})
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error in Redeeming Liquid tokens %s", err))
			}

			redeemepoch, err := getLastStakeEpoch("/home/photo/logs/redeemepoch")
			k.Logger(ctx).Info(fmt.Sprintf("Redeemed epoch %d \n", redeemepoch))
			// Convert the string to a float64
			// Check if getLastStakeEpoch returned an error
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error %s", err))
				// Optionally, instead of terminating the program,
				// you can choose to handle the error in a different way
			}

			// Check if redeemepoch is nil or empty
			if redeemepoch == "" {
				// Insert code here to handle situations where redeemepoch is nil
				// For example, you might want to continue with a default value:
				redeemepoch = "0" // or whatever default value is appropriate
			}
			epochValue, err1 := strconv.ParseInt(redeemepoch, 10, 64)
			if err1 != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error in epoch fetch %s", epochValue))
			}
			redeemedAmount, err = k.DistributeRedeemedTokens(ctx, epochValue)

			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error %d %s", redeemedAmount, err))
			}

			err = k.DeleteRedemptionRecord(ctx)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error %s", err))
			}
		}
	}

	return redeemedAmount, nil
}

// DistributeRedeemedTokens distributes redeemed tokens to all contracts based on their stake.
func (k PhotosynthesisKeeper) DistributeRedeemedTokens(ctx sdk.Context, epoch int64) (float64, error) {
	amount := k.bankkeeper.GetAllBalances(ctx, sdk.AccAddress("archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"))
	redeemedAmount := amount.AmountOf("uarch")
	k.Logger(ctx).Info(fmt.Sprintf("Distribute redemed amount: %d \n", redeemedAmount))
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	cumulativeStakes := make(map[string]float64)
	totalStake := 0.0
	// Initialize map
	dappStakes := make(map[string]float64)

	// Open the file
	file, err := os.Open("/home/photo/logs/liquiditydistributionforDapps")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	// Scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into parts
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			amount, err1 := strconv.ParseFloat(parts[2], 64)
			if err1 != nil {
				fmt.Println("Error converting to float", err1)
			}

			if _, exists := dappStakes[parts[0]]; exists {
				dappStakes[parts[0]] = dappStakes[parts[0]] + amount
			} else {
				dappStakes[parts[0]] = amount
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	// Now cumulativeStakes map is filled with data
	fmt.Println(cumulativeStakes)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Retrieve deposit records for the contract
		contractStake, ok := dappStakes[meta.LiquidityTokenAddress]
		if !ok {
			k.Logger(ctx).Error(fmt.Sprintf("Stake not found for: %v", meta.LiquidityTokenAddress))
			return false
		}
		k.Logger(ctx).Info(fmt.Sprintf("Redemption Amount: %v", contractStake))

		combined := meta.RedemptionAddress + ":" + meta.LiquidityTokenAddress

		if contractStake > 0.0 {
			cumulativeStake, ok := cumulativeStakes[combined]
			if !ok {
				k.Logger(ctx).Error(fmt.Sprintf("Cumulative Stake not found for: %v", combined))
				return false
			}

			cumulativeStakes[combined] = cumulativeStake + contractStake
			k.Logger(ctx).Info(fmt.Sprintf("Cumulative Stakes: %v", cumulativeStakes[combined]))
			totalStake = totalStake + contractStake
			k.Logger(ctx).Info(fmt.Sprintf("Total Stake: %v", totalStake))
		}
		return false
	})

	// Distribute the redeemed liquidity tokens to each contract proportionally
	for contractAddr, contractStake := range cumulativeStakes {
		// Calculate the proportion of the stake for the current contract

		split := strings.Split(contractAddr, ":")
		stakeProportion := contractStake / totalStake
		stakeratio := stakeProportion

		k.Logger(ctx).Info(fmt.Sprintf("Stake ratio: %d \n", stakeratio))
		// Calculate the amount of liquidity tokens to distribute for the current contract
		redeemedTokensAmount := stakeratio * float64(redeemedAmount.Int64())
		k.Logger(ctx).Info(fmt.Sprintf("Redeemed Tokens Amount: %v", redeemedTokensAmount))

		// Distribute the calculated amount of liquidity tokens to the contract's liquidity token address
		contractAddress, err := sdk.AccAddressFromBech32(contractAddr)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error in Decoding Address %s", err))
		}
		//log.Printf("Contract Address: %s, Liquid Token Amount: %d\n", contractAddress, liquidityTokensAmount)
		err = k.bankkeeper.SendCoins(ctx, sdk.AccAddress("archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"), sdk.AccAddress(split[0]), sdk.NewCoins(sdk.NewCoin("uarch", sdk.NewInt(int64(redeemedTokensAmount)))))
		k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s,Redeemed Token Amount: %d\n", contractAddress, redeemedTokensAmount))
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error in Transferring redeemed tokens %s", err))
		}
	}
	return float64(redeemedAmount.Int64()), nil
}

// DeleteRedemptionRecord deletes the latest redemption record from the store.
func (k PhotosynthesisKeeper) DeleteRedemptionRecord(ctx sdk.Context) error {
	record, found := k.GetLatestRedemptionRecord(ctx)
	if !found {
		return nil
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetRedemptionRecordKey(record.Timestamp)
	k.Logger(ctx).Info(fmt.Sprintf("Redemption record deleted"))
	store.Delete(key)
	return nil
}

// UpdateContract updates the contract information in the store.
func (k PhotosynthesisKeeper) UpdateContract(ctx sdk.Context, contract *types.Contract) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetContractKey(contract.Address)
	value := k.cdc.MustMarshal(contract)
	store.Set(key, value)
	return nil
}

// DeleteContract deletes the contract information from the store.
func (k PhotosynthesisKeeper) DeleteContract(ctx sdk.Context, address sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetContractKey(string(address))
	store.Delete(key)
	return nil
}

// GetStake retrieves the stake of a contract.
func (k PhotosynthesisKeeper) GetRedemptionAmount(ctx sdk.Context, address sdk.AccAddress, epoch int64) (*types.Coin, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(string(address) + string(epoch))
	value := store.Get(key)
	if value == nil {
		return &types.Coin{}, fmt.Errorf("stake not found for address %s", address.String())
	}
	k.Logger(ctx).Info(fmt.Sprintf("Get Redemption Amount"))
	var stake types.Coin
	k.cdc.Unmarshal(value, &stake)
	return &stake, nil
}

// SetStake sets the stake of a contract.
func (k PhotosynthesisKeeper) SetRedemptionAmount(ctx sdk.Context, address sdk.AccAddress, epoch int64, stake *types.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(string(address) + string(epoch))
	amount, err := k.GetRedemptionAmount(ctx, address, epoch)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Cumulative Liquid Token Amount: %d\n", amount))
	}
	stakedamount := k.cdc.MustMarshal(&types.Coin{"stuarch", stake.Amount + amount.Amount})
	k.Logger(ctx).Info(fmt.Sprintf("Set Redemption Amount"))
	store.Set(key, stakedamount)
	return nil
}

// SetStake sets the stake of a contract.
func (k PhotosynthesisKeeper) DeleteRedemptionAmount(ctx sdk.Context, epoch int64, address sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(string(address) + string(epoch))
	k.Logger(ctx).Info(fmt.Sprintf("Delete Redemption Amount"))
	store.Delete(key)
	return nil
}

// GetTotalStake calculates the total stake across all contracts
func (k PhotosynthesisKeeper) GetTotalRedemptiomAmount(ctx sdk.Context, epoch int64) (*types.Coin, error) {
	var totalStake int64
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		lstake, err := k.GetRedemptionAmount(ctx, sdk.AccAddress(meta.RewardsAddress), epoch)
		if err != nil {
			return true
		}

		totalStake += lstake.Amount
		return false
	})
	k.Logger(ctx).Info(fmt.Sprintf("Total Redemptiom Amount: %s", totalStake))
	return &types.Coin{"stuarch", totalStake}, nil
}

// SendTokensToContract sends tokens to a contract address
func (k PhotosynthesisKeeper) SendTokensToContract(ctx sdk.Context, address sdk.AccAddress, amount sdk.Int) error {
	err := k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(sdk.NewCoin("", amount)))
	if err != nil {
		return fmt.Errorf("failed to send tokens to contract: %s", err)
	}
	return nil
}

/*
// GetRedemptionRateThreshold returns the redemption rate threshold
func (k PhotosynthesisKeeper) SetRedemptionRateThreshold(ctx sdk.Context) sdk.Dec {
	return k.
}

// GetRedemptionIntervalThreshold returns the redemption interval threshold
func (k PhotosynthesisKeeper) SetRedemptionIntervalThreshold(ctx sdk.Context) time.Duration {
	return k.paramStore.Get(ctx,)
}
*/

// GetLatestRedemptionRecord gets the latest redemption record
func (k PhotosynthesisKeeper) GetLatestRedemptionRecord(ctx sdk.Context) (types.RedemptionRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStoreReversePrefixIterator(store, []byte(types.RedemptionRecordPrefix))
	defer iter.Close()
	if !iter.Valid() {
		return types.RedemptionRecord{}, false
	}
	var record types.RedemptionRecord
	k.cdc.MustUnmarshal(iter.Value(), &record)
	return record, true
}

// SetLatestRedemptionRecord sets the latest redemption record
func (k PhotosynthesisKeeper) SetLatestRedemptionRecord(ctx sdk.Context, record types.RedemptionRecord) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRedemptionRecordKey(record.Timestamp)
	value := k.cdc.MustMarshal(&record)
	store.Set(key, value)
	return nil
}

// GetCumulativeLiquidityAmount gets the cumulative liquidity amount
func (k PhotosynthesisKeeper) GetCumulativeLiquidityAmount(ctx sdk.Context, epoch int64, centralliquidityAddr string) (sdk.Int, error) {
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	//cumulativeStakes := make(map[string]sdk.Int)
	file1, err := os.OpenFile("/home/photo/logs/redeemliquidityamountforDapps", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to open file: %s", err.Error()))
	}
	defer file1.Close()
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {

		text := meta.LiquidityProviderAddress + "," + centralliquidityAddr + "," + string(epoch)
		data := []byte(text + "\n")
		_, err = file1.Write(data)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("An error occurred: %s", err.Error())) // Logging the error message if an error occurred
		}

		//data := k.bankkeeper.GetAllBalances(ctx, sdk.AccAddress(meta.LiquidityTokenAddress))
		/*
			ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			cmd := []string{
				"archwayd",
				"--chain-id",
				"localnet",
				"--node",
				"http://localhost:26657",
				"--home",
				"/home/photo/.photo",
				"q",
				"bank",
				"balances",
				meta.LiquidityTokenAddress,
			}

			dp, err := NewDockerPool()
			if err != nil {

			}

			containers, err := dp.Client.ContainerList(ctx1, dt.ContainerListOptions{})
			if err != nil {
				panic(err)
			}

			var targetContainerID string
			targetImage := "stridezone:photo"
			for _, container := range containers {
				if strings.Contains(container.Image, targetImage) {
					targetContainerID = container.ID
					break // Exit the loop after finding the desired container
				}
			}

			if targetContainerID != "" {
				k.Logger(ctx).Info(fmt.Sprintf("Container ID for %s: %s", "photosynthesis", targetContainerID))
			} else {
				k.Logger(ctx).Info(fmt.Sprintf("Container with name %s not found.", "photosynthesis"))
			}

			out, error := dp.CreateAndStartExec(ctx1, targetContainerID, cmd)
			k.Logger(ctx).Error(fmt.Sprintf(error))
			data := Balance{}
			err = yaml.Unmarshal(out, &data)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error:", err))

			}
		*/
		// Find the balance for stuarch
		/*
			for _, balance := range data.Balances {
				if balance.Denom == "ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145" {
					fmt.Println("The balance for stuarch is:", balance.Amount)
					val, err := strconv.ParseInt(balance.Amount, 10, 64)
					if err != nil {
						k.Logger(ctx).Error(fmt.Sprintf("Error:", err))

					}
					liquidityAmount = val
					k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Token Amount: %v", "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", liquidityAmount))

				}
			}
		*/
		//cumulativeliquidityAmount := int64(0)
		//ibcliquidityAmount := int64(0)
		//stuarchliquidityAmount := int64(0)
		// Find the balance for stuarch
		/*
			for _, balance := range data {
				if balance.Denom == "ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145" {
					fmt.Println("The balance for stuarch is:", balance.Amount)
					val := balance.Amount.Int64()
					ibcliquidityAmount += val
					k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Token ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 Amount: %v", "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", ibcliquidityAmount))

				}

				if balance.Denom == "stuarch" {
					fmt.Println("The balance for stuarch is:", balance.Amount)
					val := balance.Amount.Int64()
					stuarchliquidityAmount += val
					k.Logger(ctx).Info(fmt.Sprintf("Contract Address: %s, Liquid Token stuarch Amount: %v", "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", stuarchliquidityAmount))

				}

			}
		*/
		//	cumulativeliquidityAmount = ibcliquidityAmount + stuarchliquidityAmount
		//	k.Logger(ctx).Info(fmt.Sprintf("Cumulative Liquid Amount %d", cumulativeliquidityAmount))
		// Sum up the liquid stake for the contract
		//	err := k.bankkeeper.SendCoins(ctx, sdk.AccAddress(meta.LiquidityTokenAddress), centralliquidityAddr, sdk.NewCoins(sdk.NewCoin("ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145", sdk.NewInt(ibcliquidityAmount))))
		//	if err != nil {
		//		fmt.Errorf("failed to send tokens to contract: %s", err)
		//	}
		//	err1 := k.bankkeeper.SendCoins(ctx, sdk.AccAddress(meta.LiquidityTokenAddress), centralliquidityAddr, sdk.NewCoins(sdk.NewCoin("stuarch", sdk.NewInt(stuarchliquidityAmount))))
		//	if err1 != nil {
		//		fmt.Errorf("failed to send tokens to contract: %s", err)
		//	}
		//	k.SetRedemptionAmount(ctx, sdk.AccAddress(meta.LiquidityTokenAddress), epoch, &types.Coin{Amount: cumulativeliquidityAmount})
		//	cumulativeStakes[meta.LiquidityTokenAddress] = sdk.NewInt(cumulativeliquidityAmount)

		return false
	})

	//var cumulativeLiquidity int64

	//for _, contractStake := range cumulativeStakes {
	// Calculate the proportion of the stake for the current contract
	//cumulativeLiquidity += contractStake.Int64()

	//}
	//	k.Logger(ctx).Info(fmt.Sprintf(": %d\n", cumulativeLiquidity))

	// Get all rewards records for the given address by limit
	//pageReq := &query.PageRequest{Limit: recordsLimitMax}
	//_, records := state.RewardsRecord(ctx).Export()
	//totalRewards := sdk.NewCoins()
	//rewardAddressStr := string(rewardaddr.Bytes())
	//for _, record := range records {
	//	if record.RewardsAddress == rewardAddressStr {
	//		totalRewards = totalRewards.Add(record.Rewards...)
	//	}
	//}
	//return sdk.NewInt(cumulativeLiquidity), nil

	/*
		store := ctx.KVStore(k.storeKey)
		key := types.CumulativeLiquidityAmountKey
		bz := store.Get([]byte(key))
		if bz == nil {
			return types.Coin{}, fmt.Errorf("cumulative liquidity amount not found")
		}
		var coins types.Coin
		k.cdc.MustUnmarshal(bz, &coins)
		return coins, nil
	*/
	return sdk.ZeroInt(), nil
}

// SetCumulativeLiquidityAmount sets the cumulative liquidity amount
func (k PhotosynthesisKeeper) SetCumulativeLiquidityAmount(ctx sdk.Context, amount *types.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := types.CumulativeLiquidityAmountKey
	bz, err := k.cdc.Marshal(amount)
	if err != nil {
		return fmt.Errorf("failed to marshal cumulative liquidity amount: %s", err)
	}
	store.Set([]byte(key), bz)
	return nil
}

/*

// DistributeRewards distributes rewards to all contracts that have opted for liquid staking
func (k PhotosynthesisKeeper) DistributeRewards(ctx sdk.Context, epochNumber uint64, cumulativeRewards sdk.Coins) error {
	params := k.GetParams(ctx)
	contracts,_ := k.ListContracts(ctx)
	for _, contract := range contracts{ // iterate over all contracts that have opted for liquid staking
		meta := contract.Address
		rewardAmount := k.GetRewards(ctx, dapp.ContractAddress) // retrieve the rewards for the contract
		if rewardAmount.IsZero() {
			continue
		}

		if rewardAmount.AmountOf(params.RewardDenom).LT(params.MinRewardsToLiquidStake) { // check if the rewards are greater than the minimum required to distribute
			continue
		}

		if epochNumber%) != 0 { // check if the current epoch is divisible by the contract's liquid stake interval
			continue
		}

		err := k.RecordRewards(ctx, dapp.ContractAddress, sdk.NewCoins()) // record the rewards for the contract
		if err != nil {
			return fmt.Errorf("failed to record rewards for contract %s: %w", dapp.ContractAddress, err)
		}

		err = k.DistributeRewardsToDapp(ctx, dapp, rewardAmount, cumulativeRewards) // distribute the rewards to the contract
		if err != nil {
			return fmt.Errorf("failed to distribute rewards for contract %s: %w", dapp.ContractAddress, err)
		}
	}

	return nil
}
*/

/*
// This function is called by DistributeRewards to actually distribute rewards to a single Dapp. It takes the Dapp object, the amount of rewards to distribute, and the cumulative rewards as inputs. It first checks if the rewards are greater than the minimum required to distribute. If so, it calculates the ratio of rewards to distribute based on the total stake of the Dapp and distributes rewards proportionally to each stakeholder. It then mints the rewards tokens and sends them to the Dapp, updates the cumulative rewards, and emits an event.
func (k PhotosynthesisKeeper) DistributeRewardsToDapp(ctx sdk.Context, dapp types.Contract, rewards sdk.Coins, cumulativeRewards sdk.Coins) error {
	// Check if rewards are greater than the minimum required to distribute
	if rewards.IsAllLTE(dapp.MinimumRewardsToLiquidStake) {
		return nil
	}

	// Calculate the ratio of rewards to distribute
	totalStake := dapp.TotalStake
	rewardRatio := sdk.NewDecFromInt(rewards.AmountOf(dapp.RewardDenom)).Quo(sdk.NewDecFromInt(totalStake))
	if rewardRatio.GT(sdk.OneDec()) {
		rewardRatio = sdk.OneDec()
	}

	// Calculate the rewards for each dapp stake
	rewardCoins := make([]sdk.Coin, len(dapp.StakeHolders))
	for i, holder := range dapp.StakeHolders {
		share := sdk.NewDecFromInt(holder.Amount).Quo(sdk.NewDecFromInt(totalStake))
		amount := rewards.AmountOf(dapp.RewardDenom).Mul(share.TruncateInt())
		rewardCoins[i] = sdk.NewCoin(dapp.RewardDenom, amount)
	}

	// Mint the rewards tokens and send them to the Dapp
	if err := k.bankKeeper.MintCoins(ctx, types.RewardsMintBurnAcc, rewards); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.RewardsMintBurnAcc, dapp.Address, rewards); err != nil {
		return err
	}

	// Update the cumulative rewards for the Dapp
	dapp.CumulativeRewards = cumulativeRewards.Add(rewards)
	k.SetContract(ctx, dapp)

	// Emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRewardsDistributed,
			sdk.NewAttribute(types.AttributeKeyContractAddress, dapp.Address.String()),
			sdk.NewAttribute(types.AttributeKeyRewardCoins, rewards.String()),
			sdk.NewAttribute(types.AttributeKeyCumulativeRewards, cumulativeRewards.String()),
			sdk.NewAttribute(types.AttributeKeyRewardDistribution, fmt.Sprintf("%v", rewardCoins)),
		),
	)

	return nil
}
*/

// This function returns the current balance of liquidity tokens for a given address.
func (k PhotosynthesisKeeper) LiquidityTokenBalance(ctx sdk.Context, senderAddr sdk.AccAddress) (sdk.Coins, error) {
	balance := k.bankkeeper.GetAllBalances(ctx, senderAddr)
	return balance, nil
}

/*
// This function returns the liquid staking deposit object for a given sender and contract address.
func (k PhotosynthesisKeeper) LiquidStakingDeposit(ctx sdk.Context, senderAddr sdk.AccAddress, contractAddr sdk.AccAddress) (*types.LiquidStakingDeposit, error) {
	deposit, found := k.GetLiquidStakingDeposit(ctx, senderAddr, contractAddr)
	if !found {
		return nil, types.ErrDepositNotFound
	}
	return &deposit, nil
}
*/

func getLastStakeEpoch(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return line, nil
}

func getLastRedemptionRate(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lastRedemptionRate string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "last_redemption_rate") {
			// Split the line on ": " and get the last part (without quotes)
			parts := strings.Split(line, ": ")
			if len(parts) > 1 {
				lastRedemptionRate = strings.Trim(parts[1], "\"")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return lastRedemptionRate, nil
}

func (k PhotosynthesisKeeper) QueryRedemptionRate(ctx sdk.Context) (float64, error) {

	lastRedemptionRate, err := getLastRedemptionRate("/home/photo/logs/redemptionrate")
	if err != nil {
		fmt.Println("Error:", err)

	}
	//fmt.Println("Last Redemption Rate:", lastRedemptionRate)
	k.Logger(ctx).Info(fmt.Sprintf("Last redemption rate: %v", lastRedemptionRate))
	/*
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Current Directory: ", dir)
		}
		//trimmedPath := trimToFirstNDirectoryLevels(dir, 5)
		//newPath := filepath.Join(trimmedPath, "dockernet/state/stride1")

		cmd := exec.Command(
			"strided",
			"--home",
			"/home/stride/.stride",
			"--node",
			"http://localhost:26657",
			"q",
			"stakeibc",
			"show-host-zone",
			"PHOTO",
		)

		// Get the output pipe from the first command
		out1, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error obtaining stdout:", err)
			//return
		}

		// Execute second command: grep redemption_rate:
		cmd2 := exec.Command("grep", "redemption_rate:")
		cmd2.Stdin = out1

		// Get the output pipe from the second command
		out2, err := cmd2.StdoutPipe()
		if err != nil {
			fmt.Println("Error obtaining stdout:", err)
			//return
		}

		// Execute third command: tail -n 1
		cmd3 := exec.Command("tail", "-n", "1")
		cmd3.Stdin = out2

		// Execute the command and capture its output
		out, err := cmd3.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			//return
		}

		// Convert byte array to string
		s := string(out)

		// Convert string to float64
		redemptionRate, err := strconv.ParseFloat(s, 64)
		if err != nil {
		}
	*/
	// Convert the string to a float64
	rate, err := strconv.ParseFloat(lastRedemptionRate, 64)

	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in redemption rate query %s", err))
		return 0, err
	}

	return rate, nil
}

/*
// This function returns the airdrop status for a given address, including the total amount of the airdrop, the amount currently vested, and the amount remaining to be vested.
func (k PhotosynthesisKeeper) AirdropStatus(ctx sdk.Context, senderAddr sdk.AccAddress) (*types.AirdropStatus, error) {
	status, found := k.GetAirdropStatus(ctx, senderAddr)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAirdropNotFound, "address %s", senderAddr.String())
	}

	elapsedDuration := ctx.BlockTime().Sub(status.StartTime)
	elapsedMonths := elapsedDuration / (30 * 24 * time.Hour) // assuming 30-day months
	vestingAmount := sdk.NewCoin(status.Amount.Denom, status.Amount.Amount.MulRaw(int64(100-elapsedMonths*10)).QuoRaw(100))
	remainingAmount := status.Amount.Sub(vestingAmount)

	return &types.AirdropStatus{
		Address:         senderAddr.String(),
		TotalAmount:     status.Amount,
		VestingAmount:   vestingAmount,
		RemainingAmount: remainingAmount,
		StartTime:       status.StartTime,
		VestingDuration: status.VestingDuration,
	}, nil
}


// This function is similar to AirdropStatus, but takes an address as input instead of using the caller's address.
func (k PhotosynthesisKeeper) QueryAirdropStatus(ctx sdk.Context, addr sdk.AccAddress) (*types.AirdropStatus, error) {
	status, found := k.GetAirdropStatus(ctx, addr)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAirdropNotFound, "address %s", addr.String())
	}

	elapsedDuration := ctx.BlockTime().Sub(status.StartTime)
	elapsedMonths := elapsedDuration / (30 * 24 * time.Hour) // assuming 30-day months
	vestingAmount := sdk.NewCoin(status.Amount.Denom, status.Amount.Amount.MulRaw(int64(100-elapsedMonths*10)).QuoRaw(100))
	remainingAmount := status.Amount.Sub(vestingAmount)

	return &types.AirdropStatus{
		Address:         addr.String(),
		TotalAmount:     status.Amount,
		VestingAmount:   vestingAmount,
		RemainingAmount: remainingAmount,
		StartTime:       status.StartTime,
		VestingDuration: status.VestingDuration,
	}, nil
}
*/

func (k PhotosynthesisKeeper) GetCumulativeRewardAmount(ctx sdk.Context, epoch int64, rewardaddr sdk.AccAddress) int64 {
	//records, _, _ := k.rewardKeeper.GetState().RewardsRecord(ctx)GetRewardsRecords(ctx, sdk.AccAddress(contractAddress), nil)
	//recordsLimitMax := k.rewardKeeper.MaxWithdrawRecords(ctx)
	rewardamount := k.bankKeeper.GetAllBalances(ctx, rewardaddr)
	rewardAmountEnqueued, err := k.GetTotalEnqueuedRewards(ctx, epoch)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in deriving reward amount: %s", err))
	}

	// Get all rewards records for the given address by limit
	//pageReq := &query.PageRequest{Limit: recordsLimitMax}
	//_, records := state.RewardsRecord(ctx).Export()
	//totalRewards := sdk.NewCoins()
	//rewardAddressStr := string(rewardaddr.Bytes())
	//for _, record := range records {
	//	if record.RewardsAddress == rewardAddressStr {
	//		totalRewards = totalRewards.Add(record.Rewards...)
	//	}
	//}
	return rewardamount.AmountOf("uarch").Int64() - rewardAmountEnqueued.Int64()
}

func (k PhotosynthesisKeeper) BeginBlocker(ctx sdk.Context) abci.ResponseBeginBlock {
	state := k.rewardKeeper.GetState()
	k.Logger(ctx).Info(fmt.Sprintf("Retrieved state from rewardKeeper: %v", state))

	contractmeta := state.ContractMetadataState(ctx)
	k.Logger(ctx).Info(fmt.Sprintf("Retrieved contract metadata state: %v", contractmeta))

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		k.Logger(ctx).Info(fmt.Sprintf("Iterating over contract metadata: %v", meta))

		for _, epochInfo := range k.epochKeeper.AllEpochInfos(ctx) {
			k.Logger(ctx).Info(fmt.Sprintf("Checking epoch info: %v", epochInfo))

			switch epochInfo.Identifier {
			case epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH:
				k.Logger(ctx).Info(fmt.Sprintf("Processing LiquidStakeDappRewards epoch: %v", epochInfo))

				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info(fmt.Sprintf("Retrieved EpochInfo: %v", info))

				if meta.MinimumRewardAmount > 0 {
					k.Logger(ctx).Info(fmt.Sprintf("MinimumRewardAmount is greater than 0: %v", meta.MinimumRewardAmount))

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(meta.LiquidStakeInterval) == 0 {
						k.Logger(ctx).Info(fmt.Sprintf("CurrentEpoch %v is not 0 and is a multiple of LiquidStakeInterval %v", info.CurrentEpoch, meta.LiquidStakeInterval))

						if meta.RewardsAddress != "" {
							k.Logger(ctx).Info(fmt.Sprintf("RewardsAddress is not empty: %v", meta.RewardsAddress))

							rewardAmount := k.GetCumulativeRewardAmount(ctx, info.CurrentEpoch, sdk.AccAddress(meta.RewardsAddress))
							k.Logger(ctx).Info(fmt.Sprintf("Retrieved CumulativeRewardAmount: %v", rewardAmount))

							if rewardAmount >= int64(meta.MinimumRewardAmount) {
								k.Logger(ctx).Info(fmt.Sprintf("CumulativeRewardAmount is greater than or equal to MinimumRewardAmount"))

								record := k.CreateContractLiquidStakeDepositRecordsForEpoch(ctx, state, sdk.AccAddress(meta.RewardsAddress), info.CurrentEpoch)
								k.Logger(ctx).Info(fmt.Sprintf("Created ContractLiquidStakeDepositRecordsForEpoch: %v", record))

								_ = k.EnqueueLiquidStakeRecord(ctx, record)
								k.Logger(ctx).Info(fmt.Sprintf(("Enqueued LiquidStakeRecord")))

								types.EmitLiquidStakeDepositRecordCreatedEvent(ctx, record.String(), record.Amount)
								k.Logger(ctx).Info(fmt.Sprintf("EmitLiquidStakeDepositRecordCreatedEvent for record: %v and amount: %v", record.String(), record.Amount))
							}
						}
					}
				}

			case epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH:
				k.Logger(ctx).Info(fmt.Sprintf("Processing ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: %v", epochInfo))

				// Process liquid staking deposits for contracts with enabled liquid staking
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info(fmt.Sprintf("Retrieved EpochInfo for epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: %v", info))

				infoliquidstaking, _ := k.epochKeeper.GetEpochInfo(ctx, epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH)
				k.Logger(ctx).Info(fmt.Sprintf("Retrieved EpochInfo for epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH: %v", infoliquidstaking))
				if meta.MinimumRewardAmount > 0 {
					k.Logger(ctx).Info(fmt.Sprintf("MinimumRewardAmount is greater than 0: %v", meta.MinimumRewardAmount))

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(1) == 0 {
						k.Logger(ctx).Info(fmt.Sprintf("CurrentEpoch %v is not 0 and is a multiple of 1", info.CurrentEpoch))

						// Get total liquid stake deposit records till epoch
						tls, _ := k.GetTotalLiquidStake(ctx, infoliquidstaking.CurrentEpoch)
						k.Logger(ctx).Info(fmt.Sprintf("Retrieved TotalLiquidStake: %v", tls))
						ls, _ := k.LiquidStake(ctx, info.CurrentEpoch, tls.Int64())
						k.Logger(ctx).Info(fmt.Sprintf("Retrieved TotalLiquidStake: %v", tls.Int64()))
						k.Logger(ctx).Info(fmt.Sprintf("LiquidStakeScheduled: %v", ls.Int64()+1))
						epoch, err := getLastStakeEpoch("/home/photo/logs/distributionepoch")
						// Convert the string to a float64

						// Check if getLastStakeEpoch returned an error
						if err != nil {
							k.Logger(ctx).Error(fmt.Sprintf("Error %s", err))
							// Optionally, instead of terminating the program,
							// you can choose to handle the error in a different way
						}

						// Check if redeemepoch is nil or empty
						if epoch == "" {
							// Insert code here to handle situations where redeemepoch is nil
							// For example, you might want to continue with a default value:
							epoch = "0" // or whatever default value is appropriate
						}

						epochValue, _ := strconv.ParseInt(epoch, 10, 64)
						if err != nil {
							k.Logger(ctx).Error(fmt.Sprintf("Error in epoch fetch %s", epochValue))
						}
						//err1 := clearFileContents("/home/photo/logs/distributionepoch")
						//if err1 != nil {
						//	k.Logger(ctx).Error(fmt.Sprintf("Error in clearing file contents"))
						//}
						k.DistributeLiquidity(ctx, epochValue, tls.Int64())
						k.Logger(ctx).Info(fmt.Sprintf("Distributed Liquidity for epoch %v and liquid stake %v", epochValue, tls.Int64()))
					}
				}

			case epochstypes.REDEMPTION_RATE_QUERY_EPOCH:
				k.Logger(ctx).Info(fmt.Sprintf("Processing REDEMPTION_RATE_QUERY_EPOCH: %v", epochInfo))

				// Process redemption rate query and update redemption rate threshold if necessary
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info(fmt.Sprintf("Retrieved EpochInfo for epochstypes.REDEMPTION_RATE_QUERY_EPOCH: %+v", info))

				if info.CurrentEpoch%int64(meta.RedemptionIntervalThreshold) == 0 {
					k.Logger(ctx).Info(fmt.Sprintf("CurrentEpoch %v is a multiple of RedemptionIntervalThreshold %v", info.CurrentEpoch, meta.RedemptionIntervalThreshold))

					redemptionRateInterval := meta.RedemptionRateThreshold
					k.Logger(ctx).Info(fmt.Sprintf("Using RedemptionRateThreshold: %v", redemptionRateInterval))

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(redemptionRateInterval) == 0 {
						k.Logger(ctx).Info(fmt.Sprintf("CurrentEpoch %v is not 0 and is a multiple of RedemptionRateThreshold %v", info.CurrentEpoch, redemptionRateInterval))

						redemptionRate, err := k.QueryRedemptionRate(ctx)
						if err != nil {
							k.Logger(ctx).Error(fmt.Sprintf("Error in QueryRedemptionRate: %s", err))
							k.Logger(ctx).Error(fmt.Sprintf("Error in redemption rate query %s", err))
						} else {
							k.Logger(ctx).Info(fmt.Sprintf("Successfully queried RedemptionRate: %v", redemptionRate))

							if redemptionRate > float64(meta.RedemptionRateThreshold) {
								k.Logger(ctx).Info(fmt.Sprintf("RedemptionRate %v is greater than RedemptionRateThreshold %v", redemptionRate, meta.RedemptionRateThreshold))

								redemptionInterval := meta.RedemptionIntervalThreshold
								timeSinceLatestRedemption := k.GetTimeSinceLatestRedemption(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH)
								k.Logger(ctx).Info(fmt.Sprintf("TimeSinceLatestRedemption: %v", timeSinceLatestRedemption))

								if uint64(timeSinceLatestRedemption) >= redemptionInterval {
									k.Logger(ctx).Info(fmt.Sprintf("TimeSinceLatestRedemption %v is greater than or equal to RedemptionIntervalThreshold %v", timeSinceLatestRedemption, redemptionInterval))

									// Redeem liquid tokens and distribute to Dapps
									//tls, _ := k.GetTotalLiquidStake(ctx, info.CurrentEpoch)

									//k.Logger(ctx).Info(fmt.Sprintf("TotalLiquidStake: %v", tls))

									//amount, _ := k.RedeemLiquidTokens(ctx, &types.Coin{Amount: tls.Int64()})
									//k.Logger(ctx).Info(fmt.Sprintf("Amount from RedeemLiquidTokens: %v", amount))

									amount, err := k.ProcessRedemptionRateQueries(ctx, info.CurrentEpoch, meta.RedemptionInterval)
									//k.Logger(ctx).Info(fmt.Sprintf("Amount from RedeemLiquidTokens: %v", amount))
									k.Logger(ctx).Error(fmt.Sprintf("Error processing redemption rate queries %s", err))
									types.EmitRewardsDistributedEvent(ctx, meta.RewardsAddress, int64(amount), 1)
									k.Logger(ctx).Info(fmt.Sprintf("RewardsDistributedEvent emitted for RewardsAddress %v, amount %v, and event number 1", meta.RewardsAddress, amount))
								}
							}
						}
					}
				}

			case epochstypes.REWARDS_WITHDRAWAL_EPOCH:
				k.Logger(ctx).Info(fmt.Sprintf("Processing REWARDS_WITHDRAWAL_EPOCH: %+v", epochInfo))
				_, records := state.RewardsRecord(ctx).Export()
				// Distribute rewards to contracts with enabled rewards withdrawal
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info(fmt.Sprintf("Retrieved EpochInfo for epochstypes.REWARDS_WITHDRAWAL_EPOCH: %+v", info))

				totalRewards := sdk.NewCoins()
				if meta.RewardsWithdrawalInterval > 0 && info.CurrentEpoch != 0 && info.CurrentEpoch%int64(meta.RewardsWithdrawalInterval) == 0 {
					k.Logger(ctx).Info(fmt.Sprintf("CurrentEpoch %v is not 0 and is a multiple of RewardsWithdrawalInterval %v", info.CurrentEpoch, meta.RewardsWithdrawalInterval))
					k.Logger(ctx).Info(fmt.Sprintf("Retrieved %v reward records", len(records)))

					for _, record := range records {
						totalRewards = totalRewards.Add(record.Rewards...)
						k.Logger(ctx).Info(fmt.Sprintf("Accumulated rewards: %v", totalRewards))
					}

					if !totalRewards.IsZero() {
						k.Logger(ctx).Info(fmt.Sprintf("Total rewards is not zero. Proceeding with sending the coins."))

						if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, rewardstypes.ContractRewardCollector, sdk.AccAddress(meta.RewardsAddress), totalRewards); err != nil {
							panic(fmt.Errorf("sending rewards (%s) to the rewards address (%s): %w", totalRewards, meta.RewardsAddress, err))
						} else {
							k.Logger(ctx).Info(fmt.Sprintf("Successfully sent %v coins to address %v", totalRewards, meta.RewardsAddress))

							rewardstypes.EmitRewardsWithdrawEvent(ctx, sdk.AccAddress(meta.RewardsAddress), totalRewards)
							k.Logger(ctx).Info(fmt.Sprintf("Emitting rewards withdrawal event for address %v with total rewards %v", meta.RewardsAddress, totalRewards))
						}
					}
				}
				// Clean up (safe if there were no rewards)
				state.RewardsRecord(ctx).DeleteRewardsRecords(records...)
				k.Logger(ctx).Info(fmt.Sprintf("Deleted %v reward records", len(records)))
			}
		}

		return false
	})

	// Return empty response for begin block
	return abci.ResponseBeginBlock{}
}

func (k PhotosynthesisKeeper) EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	// Process liquid stake deposits
	/*
		liquidStakeInterval := k.GetParam(ctx, types.KeyArchLiquidStakeInterval)
		info,_ := k.epochKeeper.GetEpochInfo(ctx, epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH)
		if info.CurrentEpoch%int64(liquidStakeInterval) == 0 {
			depositRecords,_ := k.GetContractLiquidStakeDepositsTillEpoch(ctx, epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH, ctx.BlockHeight())
			if len(depositRecords) > 0 {
				// Transfer Archway reward funds from the Archway to liquidity provider
				err := k.LiquidStake(ctx, info.CurrentEpoch, depositRecords)
				// Distribute liquidity tokens to Dapps
				err := k.DistributeLiquidity(ctx, depositRecords)
				// Remove liquid stake deposit records
				//k.RemoveContractLiquidStakeDepositRecordsForEpoch(ctx, epochstypes.LIQUID_STAKING_DAPP_REWARDS_EPOCH, ctx.BlockHeight())
			}
		}

		// Process redemption rate query
		redemptionRateInterval := k.GetParam(ctx, types.RedemptionRateQueryInterval)
		if info.CurrentEpoch%int64(redemptionRateInterval) == 0 {
			redemptionRate := k.QueryRedemptionRate(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH)
			if redemptionRate > k.GetParam(ctx, types.RedemptionRateThreshold) {
				redemptionInterval := k.GetParam(ctx, types.RedemptionIntervalThreshold)
				timeSinceLatestRedemption := k.GetTimeSinceLatestRedemption(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH)
				if timeSinceLatestRedemption >= redemptionInterval {
					// Redeem liquid tokens and distribute to Dapps
					k.RedeemAndDistribute(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH, redemptionRate)
					// Update latest redemption time
					k.SetLatestRedemptionTime(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH, ctx.BlockTime())
				}
			}
		}
	*/
	/*
		// Process rewards withdrawal
		rewardsWithdrawalInterval := k.GetParam(ctx, types.RewardsWithdrawalInterval)
		if  info.CurrentEpoch%int64(rewardsWithdrawalInterval) == 0 {
			// Distribute rewards to Dapps
			err := k.DistributeRewards(ctx,info.CurrentEpoch, epochstypes.REWARDS_WITHDRAWAL_EPOCH)

		}
	*/
	return []abci.ValidatorUpdate{}
}

const LatestRedemptionTimeStoreKey = "latest_redemption_time"

func (k *PhotosynthesisKeeper) SetLatestRedemptionTime(ctx sdk.Context, redemptionTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LatestRedemptionTimeKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(redemptionTime.Unix()))
	store.Set([]byte(LatestRedemptionTimeStoreKey), bz)
}

func (k *PhotosynthesisKeeper) GetTimeSinceLatestRedemption(ctx sdk.Context, queryType string) int64 {
	latestRedemptionTime := k.GetLatestRedemptionTime(ctx)

	// Assuming you use the current block time as the reference
	// You can change this to any other reference time
	currentTime := ctx.BlockTime()

	// Calculate the time difference in seconds
	timeDifference := currentTime.Sub(latestRedemptionTime).Seconds()

	return int64(timeDifference)
}

func (k *PhotosynthesisKeeper) GetParam(ctx sdk.Context, key string) int64 {
	var value int64
	//store := ctx.KVStore(k.storeKey)
	k.paramStore.Get(ctx, []byte(key), &value)
	return value
}

// SetArchLiquidStakeInterval sets the Archway liquid stake interval
func (k PhotosynthesisKeeper) SetArchLiquidStakeIntervalStore(ctx sdk.Context, interval uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.KeyArchLiquidStakeInterval), sdk.Uint64ToBigEndian(interval))
}

// SetRedemptionRateQueryInterval sets the redemption rate query interval
func (k PhotosynthesisKeeper) SetRedemptionRateQueryIntervalStore(ctx sdk.Context, interval uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.KeyRedemptionRateQueryInterval), sdk.Uint64ToBigEndian(interval))
}

// SetRedemptionInterval sets the redemption interval for liquid tokens
func (k PhotosynthesisKeeper) SetRedemptionIntervalStore(ctx sdk.Context, interval uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.KeyRedemptionInterval), sdk.Uint64ToBigEndian(interval))
}

// SetRedemptionRateThreshold sets the redemption rate threshold for liquid tokens
func (k PhotosynthesisKeeper) SetRedemptionRateThresholdStore(ctx sdk.Context, threshold string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.KeyRedemptionRateThreshold), []byte(threshold))
}

// SetRewardsWithdrawalInterval sets the rewards withdrawal interval for the specified contract address
func (k PhotosynthesisKeeper) SetRewardsWithdrawalIntervalStore(ctx sdk.Context, contractAddress sdk.AccAddress, interval uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRewardsWithdrawalIntervalKey(contractAddress), sdk.Uint64ToBigEndian(interval))
}

// GetArchLiquidStakeInterval gets the Archway liquid stake interval
func (k PhotosynthesisKeeper) GetArchLiquidStakeInterval(ctx sdk.Context) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyArchLiquidStakeInterval))
	if bz == nil {
		return 0, errors.New("Archway liquid stake interval not set")
	}

	return sdk.BigEndianToUint64(bz), nil
}

// GetRedemptionInterval gets the redemption interval for liquid tokens
func (k PhotosynthesisKeeper) GetRedemptionInterval(ctx sdk.Context) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyRedemptionInterval))
	if bz == nil {
		return 0, errors.New("redemption interval not set")
	}

	return sdk.BigEndianToUint64(bz), nil
}

// GetRedemptionRateThreshold gets the redemption rate threshold for liquid tokens
func (k PhotosynthesisKeeper) GetRedemptionRateThreshold(ctx sdk.Context) (string, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyRedemptionRateThreshold))
	if bz == nil {
		return "", errors.New("redemption rate threshold not set")
	}

	return string(bz), nil
}

// GetRewardsWithdrawalInterval gets the rewards withdrawal interval for the specified contract address
func (k PhotosynthesisKeeper) GetRewardsWithdrawalInterval(ctx sdk.Context, contractAddress sdk.AccAddress) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRewardsWithdrawalIntervalKey(contractAddress))
	if bz == nil {
		return 0, errors.New("rewards withdrawal interval not set for contract address")
	}

	return sdk.BigEndianToUint64(bz), nil
}
