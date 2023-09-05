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
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode"
	"path/filepath"
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
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Current Directory: ", dir)
	}
	trimmedPath := trimToLastNDirectoryLevels(dir, 5)
	newPath := filepath.Join(trimmedPath, "dockernet/state/photo1")
	cmd := exec.Command("archwayd",
		"--home",
		newPath,
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
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Current Directory: ", dir)
	}
	trimmedPath = trimToLastNDirectoryLevels(dir, 5)
	newPath = filepath.Join(trimmedPath, "dockernet/state/stride1")



	cmd1 := exec.Command("strided",
		"--home",
		newPath,
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
        newPath,
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
	return nil
}

// RedeemLiquidTokens redeems liquid tokens and distributes them accordingly
func (k PhotosynthesisKeeper) RedeemLiquidTokens(ctx sdk.Context, cumLiquidityAmount *types.Coin) (int64, error) {
	// Get the list of contracts
	//contracts, err := k.ListContracts(ctx)
	//	if err != nil {
	//		return 0, err
	//	}
	var redeemedAmount int64
	// Iterate over the contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Calculate the redeemed amount for each contract
		//meta := k.rewardKeeper.GetContractMetadata(ctx, sdk.AccAddress(meta.ContractAddress))
		redeemedAmount = int64(meta.RedemptionRateThreshold) * cumLiquidityAmount.Amount
		coin := sdk.NewCoins(sdk.NewCoin("", sdk.NewInt(redeemedAmount)))
		// Transfer the redeemed tokens from the module account to the contract address
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(meta.RewardsAddress), coin)
		if err != nil {
			return true
		}

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
		log.Fatalf("Failed to convert byte array to uint64: %v", err)
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
func (k PhotosynthesisKeeper) ProcessRedemptionRateQueries(ctx sdk.Context) error {
	info, _ := k.epochKeeper.GetEpochInfo(ctx, epochstypes.DAY_EPOCH)

	redemptionRateQueryInterval := k.GetRedemptionRateQueryInterval(ctx)
	if info.CurrentEpoch%int64(redemptionRateQueryInterval) != 0 {
		return nil
	}

	redemptionRate, err := k.QueryRedemptionRate(ctx)
	if err != nil {
		return err
	}
	//var redemptionRate float64

	redemptionRateThreshold := k.GetParam(ctx, types.RedemptionRateThreshold)
	if redemptionRate > float64(redemptionRateThreshold) {
		redemptionIntervalThreshold := k.GetParam(ctx, types.RedemptionIntervalThreshold)
		timeSinceLatestRedemption := ctx.BlockTime().Sub(k.GetLatestRedemptionTime(ctx))

		if timeSinceLatestRedemption.Milliseconds() >= redemptionIntervalThreshold {
			cumLiquidityAmount, _ := k.GetCumulativeLiquidityAmount(ctx)
			redeemedamount, err := k.RedeemLiquidTokens(ctx, &cumLiquidityAmount)
			if err != nil {
				return err
			}

			err = k.DistributeRedeemedTokens(ctx, &types.Coin{Amount: redeemedamount})
			if err != nil {
				return err
			}

			err = k.DeleteRedemptionRecord(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DistributeRedeemedTokens distributes redeemed tokens to all contracts based on their stake.
func (k PhotosynthesisKeeper) DistributeRedeemedTokens(ctx sdk.Context, redeemedTokensAmount *types.Coin) error {
	totalStake, err := k.GetTotalStake(ctx)
	if err != nil {
		return err
	}

	contracts, err := k.ListContracts(ctx)
	if err != nil {
		return err
	}

	for _, contract := range contracts {

		tokensToDistribute := (redeemedTokensAmount.Amount * contract.Stake) / totalStake.Amount
		if tokensToDistribute == 0 {
			continue
		}

		err = k.SendTokensToContract(ctx, sdk.AccAddress(contract.Address), sdk.NewInt(tokensToDistribute))
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteRedemptionRecord deletes the latest redemption record from the store.
func (k PhotosynthesisKeeper) DeleteRedemptionRecord(ctx sdk.Context) error {
	record, found := k.GetLatestRedemptionRecord(ctx)
	if !found {
		return nil
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetRedemptionRecordKey(record.Timestamp)

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
func (k PhotosynthesisKeeper) GetStake(ctx sdk.Context, address sdk.AccAddress) (*types.Coin, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(string(address))
	value := store.Get(key)
	if value == nil {
		return &types.Coin{}, fmt.Errorf("stake not found for address %s", address.String())
	}
	var stake types.Coin
	k.cdc.Unmarshal(value, &stake)
	return &stake, nil
}

// SetStake sets the stake of a contract.
func (k PhotosynthesisKeeper) SetStake(ctx sdk.Context, address sdk.AccAddress, stake *types.Coin) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(string(address))
	stakedamount := k.cdc.MustMarshal(stake)
	store.Set(key, stakedamount)
	return nil
}

// GetTotalStake calculates the total stake across all contracts
func (k PhotosynthesisKeeper) GetTotalStake(ctx sdk.Context) (*types.Coin, error) {
	var totalStake int64
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		lstake, err := k.GetStake(ctx, sdk.AccAddress(meta.RewardsAddress))
		if err != nil {
			return true
		}

		totalStake += lstake.Amount
		return false
	})
	return &types.Coin{"", totalStake}, nil
}

// SendTokensToContract sends tokens to a contract address
func (k PhotosynthesisKeeper) SendTokensToContract(ctx sdk.Context, address sdk.AccAddress, amount sdk.Int) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(sdk.NewCoin("", amount)))
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
func (k PhotosynthesisKeeper) GetCumulativeLiquidityAmount(ctx sdk.Context) (types.Coin, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.CumulativeLiquidityAmountKey
	bz := store.Get([]byte(key))
	if bz == nil {
		return types.Coin{}, fmt.Errorf("cumulative liquidity amount not found")
	}
	var coins types.Coin
	k.cdc.MustUnmarshal(bz, &coins)
	return coins, nil
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
	balance := k.bankKeeper.GetAllBalances(ctx, senderAddr)
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

func (k PhotosynthesisKeeper) QueryRedemptionRate(ctx sdk.Context) (float64, error) {
	cmd := exec.Command(
		"strided",
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

	return redemptionRate, nil
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

func (k PhotosynthesisKeeper) GetCumulativeRewardAmount(ctx sdk.Context, state rewardKeeper.State, rewardaddr sdk.AccAddress) sdk.Coins {
	//records, _, _ := k.rewardKeeper.GetState().RewardsRecord(ctx)GetRewardsRec