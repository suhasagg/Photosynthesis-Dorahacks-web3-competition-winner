package keeper

import (
	"bytes"
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
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"gopkg.in/yaml.v2"
	"log"
	"os/exec"
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
	rewardKeeper     rewardKeeper.Keeper
	epochKeeper      epochsmodulekeeper.Keeper
}

func NewPhotosynthesisKeeper(cdc codec.Codec, paramStore paramTypes.Subspace, storeKey storetypes.StoreKey,
	contractInfoView rewardKeeper.ContractInfoReaderExpected,
	trackingKeeper rewardKeeper.TrackingKeeperExpected,
	ak rewardKeeper.AuthKeeperExpected,
	bk rewardKeeper.BankKeeperExpected,
	rewardKeeper rewardKeeper.Keeper,
	epochKeeper epochsmodulekeeper.Keeper) PhotosynthesisKeeper {

	return PhotosynthesisKeeper{
		cdc:        cdc,
		paramStore: paramStore,
		//state:            NewState(codec.Codec, storeKey),
		storeKey:         storeKey,
		contractInfoView: contractInfoView,
		trackingKeeper:   trackingKeeper,
		authKeeper:       ak,
		bankKeeper:       bk,
		rewardKeeper:     rewardKeeper,
		epochKeeper:      epochKeeper,
	}
}

// struct to hold the YAML data for stride address account balance for different crypto tokens
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

// CreateContractLiquidStakeDepositRecordsForEpoch creates a new deposit record for the given contract and epoch
func (k PhotosynthesisKeeper) CreateContractLiquidStakeDepositRecordsForEpoch(ctx sdk.Context, state rewardKeeper.State, rewardAddress sdk.AccAddress, epoch int64) *types.DepositRecord {
	// Determine the contract's liquid stake deposit amount for the given epoch
	// This will depend on your specific application logic and may involve calculations or querying other modules
	amount := k.GetCumulativeRewardAmount(ctx, state, rewardAddress)

	// Create a new deposit record with the appropriate fields
	depositRecord := types.DepositRecord{
		ContractAddress: string(rewardAddress.Bytes()),
		Epoch:           epoch,
		Amount:          amount.AmountOf("stake").Int64(),
		Status:          "pending",
	}

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

func (k PhotosynthesisKeeper) GetTotalLiquidStake(ctx sdk.Context, epoch int64) (sdk.Int, error) {
	totalLiquidStake := sdk.ZeroInt()
	store := ctx.KVStore(k.storeKey)
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
		var updatedRecords *types.DepositRecords
		updatedRecords = &types.DepositRecords{}
		for _, record := range depositRecords {
			if record.Status == "pending" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
				record.Status = "completed"
				updatedRecords.Records = append(updatedRecords.Records, record)
			}
		}
		if len(updatedRecords.Records) > 0 {
			store.Set([]byte(meta.RewardsAddress), k.cdc.MustMarshal(updatedRecords))
		}
		// Add the contract's liquid stake to the total liquid stake
		totalLiquidStake = totalLiquidStake.Add(contractLiquidStake)
		return false
	})

	return totalLiquidStake, nil
}

func (k PhotosynthesisKeeper) LiquidStake(ctx sdk.Context, epoch int64, tls int64) (sdk.Int, error) {
	ls := strconv.FormatInt(tls, 10)
	cmd := exec.Command("/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/build/archwayd",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/photo1",
		"tx",
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
		log.Fatal(err)
		return sdk.ZeroInt(), err
	}

	fmt.Printf("The output of the command is: \n%s\n", out)

	cmd1 := exec.Command("strided",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1",
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

	// Run the command and capture its output
	out1, err1 := cmd1.Output()
	if err1 != nil {
		log.Fatal(err1)
		return sdk.ZeroInt(), err
	}

	fmt.Printf("Output: \n%s\n", out1)

	cmd2 := exec.Command(
		"strided",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1",
		"q",
		"bank",
		"balances",
		"stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8",
	)

	// Execute the command and capture its output
	out, err = cmd2.Output()
	if err != nil {
		log.Fatal(err)
		return sdk.ZeroInt(), err

	}
	data := Balance{}
	err = yaml.Unmarshal(out, &data)
	if err != nil {
		log.Fatal(err)
		return sdk.ZeroInt(), err
	}

	// Find the balance for stuarch
	for _, balance := range data.Balances {
		if balance.Denom == "stuarch" {
			fmt.Println("The balance for stuarch is:", balance.Amount)
			val, err := strconv.ParseInt(balance.Amount, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return sdk.ZeroInt(), err
			}
			return sdk.NewInt(val), nil
		}
	}

	//var liquidityAmount int64
	//liquidityAmount := 10
	//amount,err :=  k.GetTotalLiquidStake(ctx,epoch)
	// Transfer reward funds from Archway to liquidity provider //
	//TODO STRIDE INTERACTION
	//err1 := k.TransferRewardFunds(ctx, contract.ArchwayRewardFundsTransferAddress, contract.LiquidityProviderAddress, amount)
	//if err != nil {
	//	return err1
	//}
	return sdk.ZeroInt(), nil
	//Distribute liquidity tokens to DApps
	//k.DistributeLiquidity(ctx, epoch, tls)
}

func (k PhotosynthesisKeeper) DistributeLiquidity(ctx sdk.Context, epoch int64, liquidityAmount int64) {
	// Get the total stake amount from all deposit records
	totalStake := sdk.NewInt(0)

	// Calculate the cumulative stake for each contract
	cumulativeStakes := make(map[string]sdk.Int)
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
			if record.Status == "completed" {
				cumulativeStakes[record.ContractAddress] = sdk.ZeroInt()
			}
		}
		for _, record := range depositRecords {
			if record.Status == "completed" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
				cumulativeStakes[record.ContractAddress] = cumulativeStakes[record.ContractAddress].Add(sdk.NewInt(record.Amount))
				totalStake = totalStake.Add(sdk.NewInt(record.Amount))
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

	// Distribute the liquidity tokens to each contract proportionally
	for contractAddr, contractStake := range cumulativeStakes {
		// Calculate the proportion of the stake for the current contract
		stakeProportion := sdk.NewDecFromInt(contractStake).Quo(sdk.NewDecFromInt(totalStake))
		stakeratio, err := stakeProportion.Float64()
		// Calculate the amount of liquidity tokens to distribute for the current contract
		liquidityTokensAmount := stakeratio * float64(liquidityAmount)
		// Distribute the calculated amount of liquidity tokens to the contract's liquidity token address
		contractAddress, err := sdk.AccAddressFromBech32(contractAddr)
		if err != nil {
			panic(err)
		}
		log.Printf("Contract Address: %s, Liquid Token Amount: %d\n", contractAddress, liquidityTokensAmount)
		err = k.bankKeeper.SendCoins(ctx, contractAddr, contractAddr, sdk.NewCoins(sdk.NewCoin("stStake", sdk.NewInt(int64(liquidityTokensAmount)))))
		if err != nil {
			panic(err)
		}

	}

}

// DeleteLiquidStake DepositRecord deletes completed deposit records for a given contract
func (k *PhotosynthesisKeeper) DeleteLiquidStakeDepositRecord(ctx sdk.Context, contractAddress sdk.AccAddress) error {
	store