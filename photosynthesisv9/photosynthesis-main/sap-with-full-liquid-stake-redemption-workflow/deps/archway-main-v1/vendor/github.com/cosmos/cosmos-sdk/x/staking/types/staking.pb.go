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
	cmd := exec.Command("/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/build/archwayd",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/photo1",
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

	cmd1 := exec.Command("/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/build/strided",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1",
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
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/build/strided",
		"--home",
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/dockernet/state/stride1",
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
		"/media/swordfish/Hydra3/photo/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/stridev1-archway-photosynthesis-with-interchain-accounts/build/strided",
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
	//records, _, _ := k.rewardKeeper.GetState().RewardsRecord(ctx)GetRewardsRecords(ctx, sdk.AccAddress(contractAddress), nil)
	//recordsLimitMax := k.rewardKeeper.MaxWithdrawRecords(ctx)

	// Get all rewards records for the given address by limit
	//pageReq := &query.PageRequest{Limit: recordsLimitMax}
	_, records := state.RewardsRecord(ctx).Export()
	totalRewards := sdk.NewCoins()
	rewardAddressStr := string(rewardaddr.Bytes())
	for _, record := range records {
		if record.RewardsAddress == rewardAddressStr {
			totalRewards = totalRewards.Add(record.Rewards...)
		}
	}
	return totalRewards
}

func (k PhotosynthesisKeeper) BeginBlocker(ctx sdk.Context) abci.ResponseBeginBlock {
	state := k.rewardKeeper.GetState()
	k.Logger(ctx).Info("Retrieved state from rewardKeeper: %+v", state)

	contractmeta := state.ContractMetadataState(ctx)
	k.Logger(ctx).Info("Retrieved contract metadata state: %+v", contractmeta)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		k.Logger(ctx).Info("Iterating over contract metadata: %+v", meta)

		for _, epochInfo := range k.epochKeeper.AllEpochInfos(ctx) {
			k.Logger(ctx).Info("Checking epoch info: %+v", epochInfo)

			switch epochInfo.Identifier {
			case epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH:
				k.Logger(ctx).Info("Processing LiquidStakeDappRewards epoch: %+v", epochInfo)

				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info("Retrieved EpochInfo: %+v", info)

				if meta.MinimumRewardAmount > 0 {
					k.Logger(ctx).Info("MinimumRewardAmount is greater than 0: %+v", meta.MinimumRewardAmount)

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(meta.LiquidStakeInterval) == 0 {
						k.Logger(ctx).Info("CurrentEpoch %v is not 0 and is a multiple of LiquidStakeInterval %v", info.CurrentEpoch, meta.LiquidStakeInterval)

						if meta.RewardsAddress != "" {
							k.Logger(ctx).Info("RewardsAddress is not empty: %v", meta.RewardsAddress)

							rewardAmount := k.GetCumulativeRewardAmount(ctx, state, sdk.AccAddress(meta.RewardsAddress))
							k.Logger(ctx).Info("Retrieved CumulativeRewardAmount: %v", rewardAmount)

							if rewardAmount.AmountOf("stake").Int64() >= int64(meta.MinimumRewardAmount) {
								k.Logger(ctx).Info("CumulativeRewardAmount is greater than or equal to MinimumRewardAmount")

								record := k.CreateContractLiquidStakeDepositRecordsForEpoch(ctx, state, sdk.AccAddress(meta.RewardsAddress), info.CurrentEpoch)
								k.Logger(ctx).Info("Created ContractLiquidStakeDepositRecordsForEpoch: %+v", record)

								_ = k.EnqueueLiquidStakeRecord(ctx, record)
								k.Logger(ctx).Info("Enqueued LiquidStakeRecord")

								types.EmitLiquidStakeDepositRecordCreatedEvent(ctx, record.String(), record.Amount)
								k.Logger(ctx).Info("EmitLiquidStakeDepositRecordCreatedEvent for record: %v and amount: %v", record.String(), record.Amount)
							}
						}
					}
				}

			case epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH:
				k.Logger(ctx).Info("Processing ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: %+v", epochInfo)

				// Process liquid staking deposits for contracts with enabled liquid staking
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info("Retrieved EpochInfo for epochstypes.ARCH_CENTRAL_LIQUID_STAKE_INTERVAL_EPOCH: %+v", info)

				infoliquidstaking, _ := k.epochKeeper.GetEpochInfo(ctx, epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH)
				k.Logger(ctx).Info("Retrieved EpochInfo for epochstypes.LIQUID_STAKING_DApp_Rewards_EPOCH: %+v", infoliquidstaking)
				if meta.MinimumRewardAmount > 0 {
					k.Logger(ctx).Info("MinimumRewardAmount is greater than 0: %+v", meta.MinimumRewardAmount)

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(1) == 0 {
						k.Logger(ctx).Info("CurrentEpoch %v is not 0 and is a multiple of 1", info.CurrentEpoch)

						// Get total liquid stake deposit records till epoch
						tls, _ := k.GetTotalLiquidStake(ctx, infoliquidstaking.CurrentEpoch)
						k.Logger(ctx).Info("Retrieved TotalLiquidStake: %v", tls)

						ls, _ := k.LiquidStake(ctx, info.CurrentEpoch, tls.Int64())
						k.Logger(ctx).Info("Computed LiquidStake: %v", ls)

						k.DistributeLiquidity(ctx, infoliquidstaking.CurrentEpoch, ls.Int64())
						k.Logger(ctx).Info("Distributed Liquidity for epoch %v and liquid stake %v", infoliquidstaking.CurrentEpoch, ls)
					}
				}

			case epochstypes.REDEMPTION_RATE_QUERY_EPOCH:
				k.Logger(ctx).Info("Processing REDEMPTION_RATE_QUERY_EPOCH: %+v", epochInfo)

				// Process redemption rate query and update redemption rate threshold if necessary
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info("Retrieved EpochInfo for epochstypes.REDEMPTION_RATE_QUERY_EPOCH: %+v", info)

				if info.CurrentEpoch%int64(meta.RedemptionIntervalThreshold) == 0 {
					k.Logger(ctx).Info("CurrentEpoch %v is a multiple of RedemptionIntervalThreshold %v", info.CurrentEpoch, meta.RedemptionIntervalThreshold)

					redemptionRateInterval := meta.RedemptionRateThreshold
					k.Logger(ctx).Info("Using RedemptionRateThreshold: %v", redemptionRateInterval)

					if info.CurrentEpoch != 0 && info.CurrentEpoch%int64(redemptionRateInterval) == 0 {
						k.Logger(ctx).Info("CurrentEpoch %v is not 0 and is a multiple of RedemptionRateThreshold %v", info.CurrentEpoch, redemptionRateInterval)

						redemptionRate, err := k.QueryRedemptionRate(ctx)
						if err != nil {
							k.Logger(ctx).Info("Error in QueryRedemptionRate: %s", err)
							fmt.Errorf("Error in redemption rate query %s", err)
						} else {
							k.Logger(ctx).Info("Successfully queried RedemptionRate: %v", redemptionRate)

							if uint64(redemptionRate) > meta.RedemptionRateThreshold {
								k.Logger(ctx).Info("RedemptionRate %v is greater than RedemptionRateThreshold %v", redemptionRate, meta.RedemptionRateThreshold)

								redemptionInterval := meta.RedemptionIntervalThreshold
								timeSinceLatestRedemption := k.GetTimeSinceLatestRedemption(ctx, epochstypes.REDEMPTION_RATE_QUERY_EPOCH)
								k.Logger(ctx).Info("TimeSinceLatestRedemption: %v", timeSinceLatestRedemption)

								if uint64(timeSinceLatestRedemption) >= redemptionInterval {
									k.Logger(ctx).Info("TimeSinceLatestRedemption %v is greater than or equal to RedemptionIntervalThreshold %v", timeSinceLatestRedemption, redemptionInterval)

									// Redeem liquid tokens and distribute to Dapps
									tls, _ := k.GetTotalLiquidStake(ctx, info.CurrentEpoch)
									k.Logger(ctx).Info("TotalLiquidStake: %v", tls)

									amount, _ := k.RedeemLiquidTokens(ctx, &types.Coin{Amount: tls.Int64()})
									k.Logger(ctx).Info("Amount from RedeemLiquidTokens: %v", amount)

									types.EmitRewardsDistributedEvent(ctx, meta.RewardsAddress, amount, 1)
									k.Logger(ctx).Info("RewardsDistributedEvent emitted for RewardsAddress %v, amount %v, and event number 1", meta.RewardsAddress, amount)
								}
							}
						}
					}
				}

			case epochstypes.REWARDS_WITHDRAWAL_EPOCH:
				k.Logger(ctx).Info("Processing REWARDS_WITHDRAWAL_EPOCH: %+v", epochInfo)

				// Distribute rewards to contracts with enabled rewards withdrawal
				info, _ := k.epochKeeper.GetEpochInfo(ctx, epochInfo.Identifier)
				k.Logger(ctx).Info("Retrieved EpochInfo for epochstypes.REWARDS_WITHDRAWAL_EPOCH: %+v", info)

				totalRewards := sdk.NewCoins()
				if meta.RewardsWithdrawalInterval > 0 && info.CurrentEpoch != 0 && info.CurrentEpoch%int64(meta.RewardsWithdrawalInterval) == 0 {
					k.Logger(ctx).Info("CurrentEpoch %v is not 0 and is a multiple of RewardsWithdrawalInterval %v", info.CurrentEpoch, meta.RewardsWithdrawalInterval)

					_, records := state.RewardsRecord(ctx).Export()
					k.Logger(ctx).Info("Retrieved %v reward records", len(records))

					for _, record := range records {
						totalRewards = totalRewards.Add(record.Rewards...)
						k.Logger(ctx).Info("Accumulated rewards: %v", totalRewards)
					}

					if !totalRewards.IsZero() {
						k.Logger(ctx).Info("Total rewards is not zero. Proceeding with sending the coins.")

						if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, rewardstypes.ContractRewardCollector, sdk.AccAddress(meta.RewardsAddress), totalRewards); err != nil {
							panic(fmt.Errorf("sending rewards (%s) to the rewards address (%s): %w", totalRewards, meta.RewardsAddress, err))
						} else {
							k.Logger(ctx).Info("Successfully sent %v coins to address %v", totalRewards, meta.RewardsAddress)

							rewardstypes.EmitRewardsWithdrawEvent(ctx, sdk.AccAddress(meta.RewardsAddress), totalRewards)
							k.Logger(ctx).Info("Emitting rewards withdrawal event for address %v with total rewards %v", meta.RewardsAddress, totalRewards)
						}
					}

					// Clean up (safe if there were no rewards)
					state.RewardsRecord(ctx).DeleteRewardsRecords(records...)
					k.Logger(ctx).Info("Deleted %v reward records", len(records))
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
                                                                                                                                                                                                                                                                                                                               0x78, 0xfc, 0x9f, 0xde, 0x4a, 0xa8, 0x1d, 0x77,
	0xa2, 0x9c, 0xe8, 0x44, 0x5f, 0x9a, 0xd5, 0x89, 0xa6, 0xb2, 0xf7, 0xc9, 0x2d, 0xe8, 0xdf, 0x59,
	0xc8, 0x1f, 0xe0, 0x00, 0xbb, 0x14, 0x59, 0x53, 0x93, 0xa6, 0x7c, 0x6b, 0x6e, 0x4c, 0xe5, 0x67,
	0x4b, 0x7d, 0xed, 0x78, 0xca, 0xa0, 0xf9, 0xe1, 0x25, 0x83, 0xe6, 0x37, 0x61, 0x85, 0x3f, 0x87,
	0x23, 0x1b, 0xa5, 0xb7, 0x97, 0x9b, 0x1b, 0x31, 0xca, 0xc5, 0x7d, 0xf9, 0x5a, 0x8e, 0x1e, 0x5d,
	0x14, 0x7d, 0x0d, 0x4a, 0x9c, 0x23, 0x6e, 0xcc, 0x5c, 0xfc, 0x5a, 0xfc, 0x2c, 0x4d, 0x6c, 0x1a,
	0x26, 0xb8, 0xf8, 0x7c, 0x4f, 0x2e, 0xd0, 0x3b, 0x80, 0x4e, 0xa3, 0x2f, 0x23, 0x9d, 0xd8, 0x9d,
	0x5c, 0xfe, 0xf3, 0x93, 0xb1, 0xbe, 0x21, 0xe5, 0xa7, 0x79, 0x0c, 0x73, 0x2d, 0x26, 0x86, 0x68,
	0x5f, 0x05, 0xe0, 0x76, 0x75, 0x6c, 0xe2, 0xf9, 0xae, 0x7a, 0xee, 0x5c, 0x9d, 0x8c, 0xf5, 0x35,
	0x89, 0x12, 0xef, 0x19, 0x66, 0x91, 0x2f, 0x5a, 0xfc, 0x77, 0x22, 0xb3, 0x3f, 0xd2, 0x00, 0xc5,
	0x2d, 0xdf, 0x24, 0x74, 0xc0, 0xdf, 0x67, 0x7c, 0x10, 0x4f, 0x4c, 0xcd, 0xda, 0x93, 0x07, 0xf1,
	0x58, 0x3e, 0x1c, 0xc4, 0x13, 0x95, 0xf2, 0xf5, 0xb8, 0x3d, 0x66, 0x55, 0x1c, 0x15, 0x4c, 0x17,
	0x53, 0x92, 0x18, 0xe6, 0x9d, 0x50, 0x7a, 0xaa, 0x1f, 0x66, 0x8c, 0x3f, 0x6a, 0xb0, 0x31, 0x95,
	0x51, 0xd1, 0x61, 0x7f, 0x08, 0x28, 0x48, 0x6c, 0x0a, 0x7f, 0x8d, 0xd4, 0xa1, 0xe7, 0x4e, 0xd0,
	0xb5, 0x60, 0xaa, 0xef, 0x7e, 0x7a, 0x1d, 0x3e, 0x27, 0x7c, 0xfe, 0x3b, 0x0d, 0xd6, 0x93, 0xea,
	0x23, 0x43, 0x6e, 0xc3, 0x52, 0x52, 0xbb, 0x32, 0xe1, 0xd5, 0x67, 0x31, 0x41, 0x9d, 0xfe, 0x82,
	0x3c, 0xfa, 0x6e, 0x5c, 0xae, 0xf2, 0xdb, 0xd9, 0x8d, 0x67, 0xf6, 0x46, 0x78, 0xa6, 0x74, 0xd9,
	0xe6, 0x44, 0x3c, 0xfe, 0xab, 0x41, 0xee, 0xc0, 0xf7, 0xfb, 0xc8, 0x87, 0x35, 0xcf, 0x67, 0x1d,
	0x9e, 0x59, 0xc4, 0xee, 0xa8, 0x47, 0xb7, 0xec, 0x83, 0xbb, 0xf3, 0x39, 0xe9, 0x9f, 0x63, 0x7d,
	0x1a, 0xca, 0x2c, 0x7b, 0x3e, 0x6b, 0x0a, 0xca, 0x91, 0x7c, 0x92, 0xbf, 0x07, 0xcb, 0x17, 0x95,
	0xc9, 0x2e, 0xf9, 0xbd, 0xb9, 0x95, 0x5d, 0x84, 0x99, 0x8c, 0xf5, 0xf5, 0xb8, 0x62, 0x22, 0xb2,
	0x61, 0x2e, 0x75, 0x13, 0xda, 0x77, 0x0a, 0x3c, 0x7e, 0xff, 0x7a, 0xa0, 0x6b, 0x5f, 0xfe, 0xad,
	0x06, 0x10, 0x7f, 0x79, 0x40, 0xaf, 0xc3, 0xcb, 0xcd, 0xef, 0xdc, 0x6e, 0x75, 0x0e, 0x8f, 0x6e,
	0x1e, 0xdd, 0x39, 0xec, 0xdc, 0xb9, 0x7d, 0x78, 0xb0, 0xb7, 0xdb, 0xbe, 0xd5, 0xde, 0x6b, 0xad,
	0x66, 0xaa, 0xe5, 0x7b, 0xf7, 0xeb, 0xa5, 0x3b, 0x1e, 0x1d, 0x10, 0xcb, 0x39, 0x71, 0x88, 0x8d,
	0x5e, 0x83, 0xf5, 0x8b, 0xdc, 0x7c, 0xb5, 0xd7, 0x5a, 0xd5, 0xaa, 0x4b, 0xf7, 0xee, 0xd7, 0x0b,
	0x72, 0x16, 0x23, 0x36, 0xda, 0x84, 0xab, 0xd3, 0x7c, 0xed, 0xdb, 0xdf, 0x5a, 0xcd, 0x56, 0x97,
	0xef, 0xdd, 0xaf, 0x17, 0xa3, 0xa1, 0x0d, 0x19, 0x80, 0x92, 0x9c, 0x0a, 0x6f, 0xa1, 0x0a, 0xf7,
	0xee, 0xd7, 0xf3, 0xd2, 0x81, 0xd5, 0xdc, 0xfb, 0x1f, 0xd5, 0x32, 0xcd, 0x5b, 0x9f, 0x3c, 0xaa,
	0x69, 0x0f, 0x1f, 0xd5, 0xb4, 0xbf, 0x3f, 0xaa, 0x69, 0x1f, 0x3c, 0xae, 0x65, 0x1e, 0x3e, 0xae,
	0x65, 0xfe, 0xfc, 0xb8, 0x96, 0xf9, 0xfe, 0xeb, 0x4f, 0xf4, 0xdd, 0x79, 0xf4, 0x51, 0x5b, 0x78,
	0xb1, 0x9b, 0x17, 0x6d, 0xf8, 0xcd, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x48, 0x4c, 0x86,
	0xf3, 0x16, 0x00, 0x00,
}

func (this *Pool) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return StakingDescription()
}
func StakingDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 9836 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0xbd, 0x6b, 0x70, 0x24, 0xd7,
		0x75, 0x18, 0x8c, 0x9e, 0x07, 0x30, 0x73, 0x30, 0x00, 0x06, 0x17, 0x58, 0x70, 0x76, 0x76, 0x17,
		0x00, 0x9b, 0xaf, 0xe5, 0x92, 0xc4, 0x92, 0x4b, 0xee, 0x92, 0x9c, 0x95, 0x44, 0xcd, 0x60, 0x66,
		0xb1, 0xe0, 0xe2, 0xc5, 0x06, 0xb0, 0xa4, 0x28, 0xf9, 0x9b, 0x6a, 0xcc, 0x5c, 0x0c, 0x9a, 0x98,
		0xe9, 0x6e, 0x76, 0xf7, 0xec, 0x02, 0x2b, 0xe9, 0x2b, 0xea, 0x11, 0x45, 0xa2, 0xcb, 0xb1, 0x14,
		0xa5, 0x62, 0x89, 0xd2, 0x2a, 0x94, 0xe5, 0x44, 0x8e, 0xac, 0xc4, 0x92, 0xa5, 0x28, 0x71, 0x92,
		0xaa, 0x48, 0xa9, 0x72, 0x2c, 0x29, 0x71, 0x4a, 0x4a, 0x5c, 0x89, 0xe3, 0x72, 0x56, 0x0e, 0xc5,
		0x4a, 0x14, 0x45, 0x89, 0xe5, 0xb5, 0x92, 0x38, 0xa5, 0xca, 0xa3, 0xee, 0xab, 0x5f, 0xf3, 0x04,
		0xb8, 0x2b, 0xc9, 0x71, 0x7e, 0x61, 0xee, 0xb9, 0xe7, 0x9c, 0x7b, 0xce, 0xb9, 0xe7, 0x9e, 0x7b,
		0xee, 0xab, 0x01, 0x5f, 0x3c, 0x0f, 0xb3, 0x35, 0xc3, 0xa8, 0xd5, 0xf1, 0x69, 0xd3, 0x32, 0x1c,
		0x63, 0xab, 0xb9, 0x7d, 0xba, 0x8a, 0xed, 0x8a, 0xa5, 0x99, 0x8e, 0x61, 0xcd, 0x51, 0x18, 0x1a,
		0x63, 0x18, 0x73, 0x02, 0x43, 0x5e, 0x86, 0xf1, 0x0b, 0x5a, 0x1d, 0x17, 0x5d, 0xc4, 0x75, 0xec,
		0xa0, 0x27, 0x20, 0xb6, 0xad, 0xd5, 0x71, 0x46, 0x9a, 0x8d, 0x9e, 0x1c, 0x3e, 0x73, 0xf7, 0x5c,
		0x88, 0x68, 0x2e, 0x48, 0xb1, 0x46, 0xc0, 0x0a, 0xa5, 0x90, 0x5f, 0x8f, 0xc1, 0x44, 0x9b, 0x5a,
		0x84, 0x20, 0xa6, 0xab, 0x0d, 0xc2, 0x51, 0x3a, 0x99, 0x54, 0xe8, 0x6f, 0x94, 0x81, 0x21, 0x53,
		0xad, 0xec, 0xaa, 0x35, 0x9c, 0x89, 0x50, 0xb0, 0x28, 0xa2, 0x69, 0x80, 0x2a, 0x36, 0xb1, 0x5e,
		0xc5, 0x7a, 0x65, 0x3f, 0x13, 0x9d, 0x8d, 0x9e, 0x4c, 0x2a, 0x3e, 0x08, 0x7a, 0x00, 0xc6, 0xcd,
		0xe6, 0x56, 0x5d, 0xab, 0x94, 0x7d, 0x68, 0x30, 0x1b, 0x3d, 0x19, 0x57, 0xd2, 0xac, 0xa2, 0xe8,
		0x21, 0xdf, 0x07, 0x63, 0x57, 0xb1, 0xba, 0xeb, 0x47, 0x1d, 0xa6, 0xa8, 0xa3, 0x04, 0xec, 0x43,
		0x9c, 0x87, 0x54, 0x03, 0xdb, 0xb6, 0x5a, 0xc3, 0x65, 0x67, 0xdf, 0xc4, 0x99, 0x18, 0xd5, 0x7e,
		0xb6, 0x45, 0xfb, 0xb0, 0xe6, 0xc3, 0x9c, 0x6a, 0x63, 0xdf, 0xc4, 0x28, 0x0f, 0x49, 0xac, 0x37,
		0x1b, 0x8c, 0x43, 0xbc, 0x83, 0xfd, 0x4a, 0x7a, 0xb3, 0x11, 0xe6, 0x92, 0x20, 0x64, 0x9c, 0xc5,
		0x90, 0x8d, 0xad, 0x2b, 0x5a, 0x05, 0x67, 0x06, 0x29, 0x83, 0xfb, 0x5a, 0x18, 0xac, 0xb3, 0xfa,
		0x30, 0x0f, 0x41, 0x87, 0xe6, 0x21, 0x89, 0xf7, 0x1c, 0xac, 0xdb, 0x9a, 0xa1, 0x67, 0x86, 0x28,
		0x93, 0x7b, 0xda, 0xf4, 0x22, 0xae, 0x57, 0xc3, 0x2c, 0x3c, 0x3a, 0x74, 0x0e, 0x86, 0x0c, 0xd3,
		0xd1, 0x0c, 0xdd, 0xce, 0x24, 0x66, 0xa5, 0x93, 0xc3, 0x67, 0x8e, 0xb7, 0x75, 0x84, 0x55, 0x86,
		0xa3, 0x08, 0x64, 0xb4, 0x08, 0x69, 0xdb, 0x68, 0x5a, 0x15, 0x5c, 0xae, 0x18, 0x55, 0x5c, 0xd6,
		0xf4, 0x6d, 0x23, 0x93, 0xa4, 0x0c, 0x66, 0x5a, 0x15, 0xa1, 0x88, 0xf3, 0x46, 0x15, 0x2f, 0xea,
		0xdb, 0x86, 0x32, 0x6a, 0x07, 0xca, 0x68, 0x0a, 0x06, 0xed, 0x7d, 0xdd, 0x51, 0xf7, 0x32, 0x29,
		0xea, 0x21, 0xbc, 0x24, 0xff, 0xe6, 0x20, 0x8c, 0xf5, 0xe3, 0x62, 0xe7, 0x21, 0xbe, 0x4d, 0xb4,
		0xcc, 0x44, 0x0e, 0x62, 0x03, 0x46, 0x13, 0x34, 0xe2, 0xe0, 0x21, 0x8d, 0x98, 0x87, 0x61, 0x1d,
		0xdb, 0x0e, 0xae, 0x32, 0x8f, 0x88, 0xf6, 0xe9, 0x53, 0xc0, 0x88, 0x5a, 0x5d, 0x2a, 0x76, 0x28,
		0x97, 0x7a, 0x0e, 0xc6, 0x5c, 0x91, 0xca, 0x96, 0xaa, 0xd7, 0x84, 0x6f, 0x9e, 0xee, 0x25, 0xc9,
		0x5c, 0x49, 0xd0, 0x29, 0x84, 0x4c, 0x19, 0xc5, 0x81, 0x32, 0x2a, 0x02, 0x18, 0x3a, 0x36, 0xb6,
		0xcb, 0x55, 0x5c, 0xa9, 0x67, 0x12, 0x1d, 0xac, 0xb4, 0x4a, 0x50, 0x5a, 0xac, 0x64, 0x30, 0x68,
		0xa5, 0x8e, 0x9e, 0xf4, 0x5c, 0x6d, 0xa8, 0x83, 0xa7, 0x2c, 0xb3, 0x41, 0xd6, 0xe2, 0x6d, 0x9b,
		0x30, 0x6a, 0x61, 0xe2, 0xf7, 0xb8, 0xca, 0x35, 0x4b, 0x52, 0x21, 0xe6, 0x7a, 0x6a, 0xa6, 0x70,
		0x32, 0xa6, 0xd8, 0x88, 0xe5, 0x2f, 0xa2, 0xbb, 0xc0, 0x05, 0x94, 0xa9, 0x5b, 0x01, 0x8d, 0x42,
		0x29, 0x01, 0x5c, 0x51, 0x1b, 0x38, 0x7b, 0x0d, 0x46, 0x83, 0xe6, 0x41, 0x93, 0x10, 0xb7, 0x1d,
		0xd5, 0x72, 0xa8, 0x17, 0xc6, 0x15, 0x56, 0x40, 0x69, 0x88, 0x62, 0xbd, 0x4a, 0xa3, 0x5c, 0x5c,
		0x21, 0x3f, 0xd1, 0x5b, 0x3d, 0x85, 0xa3, 0x54, 0xe1, 0x7b, 0x5b, 0x7b, 0x34, 0xc0, 0x39, 0xac,
		0x77, 0xf6, 0x71, 0x18, 0x09, 0x28, 0xd0, 0x6f, 0xd3, 0xf2, 0xbb, 0xe0, 0x48, 0x5b, 0xd6, 0xe8,
		0x39, 0x98, 0x6c, 0xea, 0x9a, 0xee, 0x60, 0xcb, 0xb4, 0x30, 0xf1, 0x58, 0xd6, 0x54, 0xe6, 0x3f,
		0x0c, 0x75, 0xf0, 0xb9, 0x4d, 0x3f, 0x36, 0xe3, 0xa2, 0x4c, 0x34, 0x5b, 0x81, 0xa7, 0x92, 0x89,
		0xef, 0x0d, 0xa5, 0x5f, 0x7a, 0xe9, 0xa5, 0x97, 0x22, 0xf2, 0xd7, 0x06, 0x61, 0xb2, 0xdd, 0x98,
		0x69, 0x3b, 0x7c, 0xa7, 0x60, 0x50, 0x6f, 0x36, 0xb6, 0xb0, 0x45, 0x8d, 0x14, 0x57, 0x78, 0x09,
		0xe5, 0x21, 0x5e, 0x57, 0xb7, 0x70, 0x3d, 0x13, 0x9b, 0x95, 0x4e, 0x8e, 0x9e, 0x79, 0xa0, 0xaf,
		0x51, 0x39, 0xb7, 0x44, 0x48, 0x14, 0x46, 0x89, 0xde, 0x02, 0x31, 0x1e, 0xa2, 0x09, 0x87, 0x53,
		0xfd, 0x71, 0x20, 0x63, 0x49, 0xa1, 0x74, 0xe8, 0x18, 0x24, 0xc9, 0x5f, 0xe6, 0x1b, 0x83, 0x54,
		0xe6, 0x04, 0x01, 0x10, 0xbf, 0x40, 0x59, 0x48, 0xd0, 0x61, 0x52, 0xc5, 0x62, 0x6a, 0x73, 0xcb,
		0xc4, 0xb1, 0xaa, 0x78, 0x5b, 0x6d, 0xd6, 0x9d, 0xf2, 0x15, 0xb5, 0xde, 0xc4, 0xd4, 0xe1, 0x93,
		0x4a, 0x8a, 0x03, 0x2f, 0x13, 0x18, 0x9a, 0x81, 0x61, 0x36, 0xaa, 0x34, 0xbd, 0x8a, 0xf7, 0x68,
		0xf4, 0x8c, 0x2b, 0x6c, 0xa0, 0x2d, 0x12, 0x08, 0x69, 0xfe, 0x05, 0xdb, 0xd0, 0x85, 0x6b, 0xd2,
		0x26, 0x08, 0x80, 0x36, 0xff, 0x78, 0x38, 0x70, 0x9f, 0x68, 0xaf, 0x5e, 0xcb, 0x58, 0xba, 0x0f,
		0xc6, 0x28, 0xc6, 0xa3, 0xbc, 0xeb, 0xd5, 0x7a, 0x66, 0x7c, 0x56, 0x3a, 0x99, 0x50, 0x46, 0x19,
		0x78, 0x95, 0x43, 0xe5, 0xaf, 0x44, 0x20, 0x46, 0x03, 0xcb, 0x18, 0x0c, 0x6f, 0xbc, 0x6d, 0xad,
		0x54, 0x2e, 0xae, 0x6e, 0x16, 0x96, 0x4a, 0x69, 0x09, 0x8d, 0x02, 0x50, 0xc0, 0x85, 0xa5, 0xd5,
		0xfc, 0x46, 0x3a, 0xe2, 0x96, 0x17, 0x57, 0x36, 0xce, 0x3d, 0x96, 0x8e, 0xba, 0x04, 0x9b, 0x0c,
		0x10, 0xf3, 0x23, 0x3c, 0x7a, 0x26, 0x1d, 0x47, 0x69, 0x48, 0x31, 0x06, 0x8b, 0xcf, 0x95, 0x8a,
		0xe7, 0x1e, 0x4b, 0x0f, 0x06, 0x21, 0x8f, 0x9e, 0x49, 0x0f, 0xa1, 0x11, 0x48, 0x52, 0x48, 0x61,
		0x75, 0x75, 0x29, 0x9d, 0x70, 0x79, 0xae, 0x6f, 0x28, 0x8b, 0x2b, 0x0b, 0xe9, 0xa4, 0xcb, 0x73,
		0x41, 0x59, 0xdd, 0x5c, 0x4b, 0x83, 0xcb, 0x61, 0xb9, 0xb4, 0xbe, 0x9e, 0x5f, 0x28, 0xa5, 0x87,
		0x5d, 0x8c, 0xc2, 0xdb, 0x36, 0x4a, 0xeb, 0xe9, 0x54, 0x40, 0xac, 0x47, 0xcf, 0xa4, 0x47, 0xdc,
		0x26, 0x4a, 0x2b, 0x9b, 0xcb, 0xe9, 0x51, 0x34, 0x0e, 0x23, 0xac, 0x09, 0x21, 0xc4, 0x58, 0x08,
		0x74, 0xee, 0xb1, 0x74, 0xda, 0x13, 0x84, 0x71, 0x19, 0x0f, 0x00, 0xce, 0x3d, 0x96, 0x46, 0xf2,
		0x3c, 0xc4, 0xa9, 0x1b, 0x22, 0x04, 0xa3, 0x4b, 0xf9, 0x42, 0x69, 0xa9, 0xbc, 0xba, 0xb6, 0xb1,
		0xb8, 0xba, 0x92, 0x5f, 0x4a, 0x4b, 0x1e, 0x4c, 0x29, 0x3d, 0xb3, 0xb9, 0xa8, 0x94, 0x8a, 0xe9,
		0x88, 0x1f, 0xb6, 0x56, 0xca, 0x6f, 0x94, 0x8a, 0xe9, 0xa8, 0x5c, 0x81, 0xc9, 0x76, 0x01, 0xb5,
		0xed, 0x10, 0xf2, 0xf9, 0x42, 0xa4, 0x83, 0x2f, 0x50, 0x5e, 0x61, 0x5f, 0x90, 0xbf, 0x1b, 0x81,
		0x89, 0x36, 0x93, 0x4a, 0xdb, 0x46, 0x9e, 0x82, 0x38, 0xf3, 0x65, 0x36, 0xcd, 0xde, 0xdf, 0x76,
		0x76, 0xa2, 0x9e, 0xdd, 0x32, 0xd5, 0x52, 0x3a, 0x7f, 0xaa, 0x11, 0xed, 0x90, 0x6a, 0x10, 0x16,
		0x2d, 0x0e, 0xfb, 0x73, 0x2d, 0xc1, 0x9f, 0xcd, 0x8f, 0xe7, 0xfa, 0x99, 0x1f, 0x29, 0xec, 0x60,
		0x93, 0x40, 0xbc, 0xcd, 0x24, 0x70, 0x1e, 0xc6, 0x5b, 0x18, 0xf5, 0x1d, 0x8c, 0xdf, 0x27, 0x41,
		0xa6, 0x93, 0x71, 0x7a, 0x84, 0xc4, 0x48, 0x20, 0x24, 0x9e, 0x0f, 0x5b, 0xf0, 0xce, 0xce, 0x9d,
		0xd0, 0xd2, 0xd7, 0x9f, 0x95, 0x60, 0xaa, 0x7d, 0x4a, 0xd9, 0x56, 0x86, 0xb7, 0xc0, 0x60, 0x03,
		0x3b, 0x3b, 0x86, 0x48, 0xab, 0xee, 0x6d, 0x33, 0x59, 0x93, 0xea, 0x70, 0x67, 0x73, 0x2a, 0xff,
		0x6c, 0x1f, 0xed, 0x94, 0x17, 0x32, 0x69, 0x5a, 0x24, 0xfd, 0x50, 0x04, 0x8e, 0xb4, 0x65, 0xde,
		0x56, 0xd0, 0x13, 0x00, 0x9a, 0x6e, 0x36, 0x1d, 0x96, 0x3a, 0xb1, 0x48, 0x9c, 0xa4, 0x10, 0x1a,
		0xbc, 0x48, 0x94, 0x6d, 0x3a, 0x6e, 0x7d, 0x94, 0xd6, 0x03, 0x03, 0x51, 0x84, 0x27, 0x3c, 0x41,
		0x63, 0x54, 0xd0, 0xe9, 0x0e, 0x9a, 0xb6, 0x38, 0xe6, 0xc3, 0x90, 0xae, 0xd4, 0x35, 0xac, 0x3b,
		0x65, 0xdb, 0xb1, 0xb0, 0xda, 0xd0, 0xf4, 0x1a, 0x9d, 0x6a, 0x12, 0xb9, 0xf8, 0xb6, 0x5a, 0xb7,
		0xb1, 0x32, 0xc6, 0xaa, 0xd7, 0x45, 0x2d, 0xa1, 0xa0, 0x0e, 0x64, 0xf9, 0x28, 0x06, 0x03, 0x14,
		0xac, 0xda, 0xa5, 0x90, 0x3f, 0x92, 0x84, 0x61, 0x5f, 0x02, 0x8e, 0xee, 0x84, 0xd4, 0x0b, 0xea,
		0x15, 0xb5, 0x2c, 0x16, 0x55, 0xcc, 0x12, 0xc3, 0x04, 0xb6, 0xc6, 0x17, 0x56, 0x0f, 0xc3, 0x24,
		0x45, 0x31, 0x9a, 0x0e, 0xb6, 0xca, 0x95, 0xba, 0x6a, 0xdb, 0xd4, 0x68, 0x09, 0x8a, 0x8a, 0x48,
		0xdd, 0x2a, 0xa9, 0x9a, 0x17, 0x35, 0xe8, 0x2c, 0x4c, 0x50, 0x8a, 0x46, 0xb3, 0xee, 0x68, 0x66,
		0x1d, 0x97, 0xc9, 0x32, 0xcf, 0xa6, 0x53, 0x8e, 0x2b, 0xd9, 0x38, 0xc1, 0x58, 0xe6, 0x08, 0x44,
		0x22, 0x1b, 0x15, 0xe1, 0x04, 0x25, 0xab, 0x61, 0x1d, 0x5b, 0xaa, 0x83, 0xcb, 0xf8, 0xc5, 0xa6,
		0x5a, 0xb7, 0xcb, 0xaa, 0x5e, 0x2d, 0xef, 0xa8, 0xf6, 0x4e, 0x66, 0x92, 0x30, 0x28, 0x44, 0x32,
		0x92, 0x72, 0x94, 0x20, 0x2e, 0x70, 0xbc, 0x12, 0x45, 0xcb, 0xeb, 0xd5, 0x8b, 0xaa, 0xbd, 0x83,
		0x72, 0x30, 0x45, 0xb9, 0xd8, 0x8e, 0xa5, 0xe9, 0xb5, 0x72, 0x65, 0x07, 0x57, 0x76, 0xcb, 0x4d,
		0x67, 0xfb, 0x89, 0xcc, 0x31, 0x7f, 0xfb, 0x54, 0xc2, 0x75, 0x8a, 0x33, 0x4f, 0x50, 0x36, 0x9d,
		0xed, 0x27, 0xd0, 0x3a, 0xa4, 0x48, 0x67, 0x34, 0xb4, 0x6b, 0xb8, 0xbc, 0x6d, 0x58, 0x74, 0x0e,
		0x1d, 0x6d, 0x13, 0x9a, 0x7c, 0x16, 0x9c, 0x5b, 0xe5, 0x04, 0xcb, 0x46, 0x15, 0xe7, 0xe2, 0xeb,
		0x6b, 0xa5, 0x52, 0x51, 0x19, 0x16, 0x5c, 0x2e, 0x18, 0x16, 0x71, 0xa8, 0x9a, 0xe1, 0x1a, 0x78,
		0x98, 0x39, 0x54, 0xcd, 0x10, 0xe6, 0x3d, 0x0b, 0x13, 0x95, 0x0a, 0xd3, 0x59, 0xab, 0x94, 0xf9,
		0x62, 0xcc, 0xce, 0xa4, 0x03, 0xc6, 0xaa, 0x54, 0x16, 0x18, 0x02, 0xf7, 0x71, 0x1b, 0x3d, 0x09,
		0x47, 0x3c, 0x63, 0xf9, 0x09, 0xc7, 0x5b, 0xb4, 0x0c, 0x93, 0x9e, 0x85, 0x09, 0x73, 0xbf, 0x95,
		0x10, 0x05, 0x5a, 0x34, 0xf7, 0xc3, 0x64, 0x8f, 0xc3, 0xa4, 0xb9, 0x63, 0xb6, 0xd2, 0x9d, 0xf2,
		0xd3, 0x21, 0x73, 0xc7, 0x0c, 0x13, 0xde, 0x43, 0x57, 0xe6, 0x16, 0xae, 0xa8, 0x0e, 0xae, 0x66,
		0xee, 0xf0, 0xa3, 0xfb, 0x2a, 0xd0, 0x1c, 0xa4, 0x2b, 0x95, 0x32, 0xd6, 0xd5, 0xad, 0x3a, 0x2e,
		0xab, 0x16, 0xd6, 0x55, 0x3b, 0x33, 0x43, 0x91, 0x63, 0x8e, 0xd5, 0xc4, 0xca, 0x68, 0xa5, 0x52,
		0xa2, 0x95, 0x79, 0x5a, 0x87, 0x4e, 0xc1, 0xb8, 0xb1, 0xf5, 0x42, 0x85, 0x79, 0x64, 0xd9, 0xb4,
		0xf0, 0xb6, 0xb6, 0x97, 0xb9, 0x9b, 0x9a, 0x77, 0x8c, 0x54, 0x50, 0x7f, 0x5c, 0xa3, 0x60, 0x74,
		0x3f, 0xa4, 0x2b, 0xf6, 0x8e, 0x6a, 0x99, 0x34, 0x24, 0xdb, 0xa6, 0x5a, 0xc1, 0x99, 0x7b, 0x18,
		0x2a, 0x83, 0xaf, 0x08, 0x30, 0x19, 0x11, 0xf6, 0x55, 0x6d, 0xdb, 0x11, 0x1c, 0xef, 0x63, 0x23,
		0x82, 0xc2, 0x38, 0xb7, 0x93, 0x90, 0x26, 0x96, 0x08, 0x34, 0x7c, 0x92, 0xa2, 0x8d, 0x9a, 0x3b,
		0xa6, 0xbf, 0xdd, 0xbb, 0x60, 0x84, 0x60, 0x7a, 0x8d, 0xde, 0xcf, 0x12, 0x37, 0x73, 0xc7, 0xd7,
		0xe2, 0x63, 0x30, 0x45, 0x90, 0x1a, 0xd8, 0x51, 0xab, 0xaa, 0xa3, 0xfa, 0xb0, 0x1f, 0xa4, 0xd8,
		0xc4, 0xec, 0xcb, 0xbc, 0x32, 0x20, 0xa7, 0xd5, 0xdc, 0xda, 0x77, 0x1d, 0xeb, 0x21, 0x26, 0x27,
		0x81, 0x09, 0xd7, 0xba, 0x6d, 0xc9, 0xb9, 0x9c, 0x83, 0x94, 0xdf, 0xef, 0x51, 0x12, 0x98, 0xe7,
		0xa7, 0x25, 0x92, 0x04, 0xcd, 0xaf, 0x16, 0x49, 0xfa, 0xf2, 0x7c, 0x29, 0x1d, 0x21, 0x69, 0xd4,
		0xd2, 0xe2, 0x46, 0xa9, 0xac, 0x6c, 0xae, 0x6c, 0x2c, 0x2e, 0x97, 0xd2, 0x51, 0x5f, 0x62, 0xff,
		0x74, 0x2c, 0x71, 0x6f, 0xfa, 0x3e, 0xf9, 0xdb, 0x11, 0x18, 0x0d, 0xae, 0xd4, 0xd0, 0x9b, 0xe0,
		0x0e, 0xb1, 0xad, 0x62, 0x63, 0xa7, 0x7c, 0x55, 0xb3, 0xe8, 0x80, 0x6c, 0xa8, 0x6c, 0x72, 0x74,
		0xfd, 0x67, 0x92, 0x63, 0xad, 0x63, 0xe7, 0x59, 0xcd, 0x22, 0xc3, 0xad, 0xa1, 0x3a, 0x68, 0x09,
		0x66, 0x74, 0xa3, 0x6c, 0x3b, 0xaa, 0x5e, 0x55, 0xad, 0x6a, 0xd9, 0xdb, 0xd0, 0x2a, 0xab, 0x95,
		0x0a, 0xb6, 0x6d, 0x83, 0x4d, 0x84, 0x2e, 0x97, 0xe3, 0xba, 0xb1, 0xce, 0x91, 0xbd, 0x19, 0x22,
		0xcf, 0x51, 0x43, 0xee, 0x1b, 0xed, 0xe4, 0xbe, 0xc7, 0x20, 0xd9, 0x50, 0xcd, 0x32, 0xd6, 0x1d,
		0x6b, 0x9f, 0xe6, 0xe7, 0x09, 0x25, 0xd1, 0x50, 0xcd, 0x12, 0x29, 0xff, 0x44, 0x96, 0x49, 0x4f,
		0xc7, 0x12, 0x89, 0x74, 0xf2, 0xe9, 0x58, 0x22, 0x99, 0x06, 0xf9, 0xb5, 0x28, 0xa4, 0xfc, 0xf9,
		0x3a, 0x59, 0xfe, 0x54, 0xe8, 0x8c, 0x25, 0xd1, 0x98, 0x76, 0x57, 0xd7, 0xec, 0x7e, 0x6e, 0x9e,
		0x4c, 0x65, 0xb9, 0x41, 0x96, 0x1c, 0x2b, 0x8c, 0x92, 0xa4, 0x11, 0xc4, 0xd9, 0x30, 0x4b, 0x46,
		0x12, 0x0a, 0x2f, 0xa1, 0x05, 0x18, 0x7c, 0xc1, 0xa6, 0xbc, 0x07, 0x29, 0xef, 0xbb, 0xbb, 0xf3,
		0x7e, 0x7a, 0x9d, 0x32, 0x4f, 0x3e, 0xbd, 0x5e, 0x5e, 0x59, 0x55, 0x96, 0xf3, 0x4b, 0x0a, 0x27,
		0x47, 0x47, 0x21, 0x56, 0x57, 0xaf, 0xed, 0x07, 0x27, 0x3d, 0x0a, 0xea, 0xb7, 0x13, 0x8e, 0x42,
		0xec, 0x2a, 0x56, 0x77, 0x83, 0x53, 0x0d, 0x05, 0xdd, 0xc6, 0xc1, 0x70, 0x1a, 0xe2, 0xd4, 0x5e,
		0x08, 0x80, 0x5b, 0x2c, 0x3d, 0x80, 0x12, 0x10, 0x9b, 0x5f, 0x55, 0xc8, 0x80, 0x48, 0x43, 0x8a,
		0x41, 0xcb, 0x6b, 0x8b, 0xa5, 0xf9, 0x52, 0x3a, 0x22, 0x9f, 0x85, 0x41, 0x66, 0x04, 0x32, 0x58,
		0x5c, 0x33, 0xa4, 0x07, 0x78, 0x91, 0xf3, 0x90, 0x44, 0xed, 0xe6, 0x72, 0xa1, 0xa4, 0xa4, 0x23,
		0xc1, 0xae, 0x8e, 0xa5, 0xe3, 0xb2, 0x0d, 0x29, 0x7f, 0x1e, 0xfe, 0x93, 0x59, 0x8c, 0x7f, 0x55,
		0x82, 0x61, 0x5f, 0x5e, 0x4d, 0x12, 0x22, 0xb5, 0x5e, 0x37, 0xae, 0x96, 0xd5, 0xba, 0xa6, 0xda,
		0xdc, 0x35, 0x80, 0x82, 0xf2, 0x04, 0xd2, 0x6f, 0xd7, 0xfd, 0x84, 0x86, 0x48, 0x3c, 0x3d, 0x28,
		0x7f, 0x4a, 0x82, 0x74, 0x38, 0xb1, 0x0d, 0x89, 0x29, 0xfd, 0x34, 0xc5, 0x94, 0x3f, 0x29, 0xc1,
		0x68, 0x30, 0x9b, 0x0d, 0x89, 0x77, 0xe7, 0x4f, 0x55, 0xbc, 0x3f, 0x8c, 0xc0, 0x48, 0x20, 0x87,
		0xed, 0x57, 0xba, 0x17, 0x61, 0x5c, 0xab, 0xe2, 0x86, 0x69, 0x38, 0x58, 0xaf, 0xec, 0x97, 0xeb,
		0xf8, 0x0a, 0xae, 0x67, 0x64, 0x1a, 0x34, 0x4e, 0x77, 0xcf, 0x92, 0xe7, 0x16, 0x3d, 0xba, 0x25,
		0x42, 0x96, 0x9b, 0x58, 0x2c, 0x96, 0x96, 0xd7, 0x56, 0x37, 0x4a, 0x2b, 0xf3, 0x6f, 0x2b, 0x6f,
		0xae, 0x5c, 0x5a, 0x59, 0x7d, 0x76, 0x45, 0x49, 0x6b, 0x21, 0xb4, 0xdb, 0x38, 0xec, 0xd7, 0x20,
		0x1d, 0x16, 0x0a, 0xdd, 0x01, 0xed, 0xc4, 0x4a, 0x0f, 0xa0, 0x09, 0x18, 0x5b, 0x59, 0x2d, 0xaf,
		0x2f, 0x16, 0x4b, 0xe5, 0xd2, 0x85, 0x0b, 0xa5, 0xf9, 0x8d, 0x75, 0xb6, 0xef, 0xe1, 0x62, 0x6f,
		0x04, 0x06, 0xb8, 0xfc, 0x4a, 0x14, 0x26, 0xda, 0x48, 0x82, 0xf2, 0x7c, 0xc5, 0xc2, 0x16, 0x51,
		0x0f, 0xf5, 0x23, 0xfd, 0x1c, 0xc9, 0x19, 0xd6, 0x54, 0xcb, 0xe1, 0x0b, 0x9c, 0xfb, 0x81, 0x58,
		0x49, 0x77, 0xb4, 0x6d, 0x0d, 0x5b, 0x7c, 0x3f, 0x89, 0x2d, 0x63, 0xc6, 0x3c, 0x38, 0xdb, 0x52,
		0x7a, 0x10, 0x90, 0x69, 0xd8, 0x9a, 0xa3, 0x5d, 0xc1, 0x65, 0x4d, 0x17, 0x9b, 0x4f, 0x64, 0x59,
		0x13, 0x53, 0xd2, 0xa2, 0x66, 0x51, 0x77, 0x5c, 0x6c, 0x1d, 0xd7, 0xd4, 0x10, 0x36, 0x09, 0xe6,
		0x51, 0x25, 0x2d, 0x6a, 0x5c, 0xec, 0x3b, 0x21, 0x55, 0x35, 0x9a, 0x24, 0xd7, 0x63, 0x78, 0x64,
		0xee, 0x90, 0x94, 0x61, 0x06, 0x73, 0x51, 0x78, 0x16, 0xef, 0xed, 0x7a, 0xa5, 0x94, 0x61, 0x06,
		0x63, 0x28, 0xf7, 0xc1, 0x98, 0x5a, 0xab, 0x59, 0x84, 0xb9, 0x60, 0xc4, 0xd6, 0x25, 0xa3, 0x2e,
		0x98, 0x22, 0x66, 0x9f, 0x86, 0x84, 0xb0, 0x03, 0x99, 0xaa, 0x89, 0x25, 0xca, 0x26, 0x5b, 0x6c,
		0x47, 0x4e, 0x26, 0x95, 0x84, 0x2e, 0x2a, 0xef, 0x84, 0x94, 0x66, 0x97, 0xbd, 0x4d, 0xfc, 0xc8,
		0x6c, 0xe4, 0x64, 0x42, 0x19, 0xd6, 0x6c, 0x77, 0x03, 0x54, 0xfe, 0x6c, 0x04, 0x46, 0x83, 0x87,
		0x10, 0xa8, 0x08, 0x89, 0xba, 0x51, 0x51, 0xa9, 0x6b, 0xb1, 0x13, 0xb0, 0x93, 0x3d, 0xce, 0x2d,
		0xe6, 0x96, 0x38, 0xbe, 0xe2, 0x52, 0x66, 0xff, 0xb9, 0x04, 0x09, 0x01, 0x46, 0x53, 0x10, 0x33,
		0x55, 0x67, 0x87, 0xb2, 0x8b, 0x17, 0x22, 0x69, 0x49, 0xa1, 0x65, 0x02, 0xb7, 0x4d, 0x55, 0xa7,
		0x2e, 0xc0, 0xe1, 0xa4, 0x4c, 0xfa, 0xb5, 0x8e, 0xd5, 0x2a, 0x5d, 0xf4, 0x18, 0x8d, 0x06, 0xd6,
		0x1d, 0x5b, 0xf4, 0x2b, 0x87, 0xcf, 0x73, 0x30, 0x7a, 0x00, 0xc6, 0x1d, 0x4b, 0xd5, 0xea, 0x01,
		0xdc, 0x18, 0xc5, 0x4d, 0x8b, 0x0a, 0x17, 0x39, 0x07, 0x47, 0x05, 0xdf, 0x2a, 0x76, 0xd4, 0xca,
		0x0e, 0xae, 0x7a, 0x44, 0x83, 0x74, 0x73, 0xe3, 0x0e, 0x8e, 0x50, 0xe4, 0xf5, 0x82, 0x56, 0xfe,
		0xb6, 0x04, 0xe3, 0x62, 0x99, 0x56, 0x75, 0x8d, 0xb5, 0x0c, 0xa0, 0xea, 0xba, 0xe1, 0xf8, 0xcd,
		0xd5, 0xea, 0xca, 0x2d, 0x74, 0x73, 0x79, 0x97, 0x48, 0xf1, 0x31, 0xc8, 0x36, 0x00, 0xbc, 0x9a,
		0x8e, 0x66, 0x9b, 0x81, 0x61, 0x7e, 0xc2, 0x44, 0x8f, 0x29, 0xd9, 0xc2, 0x1e, 0x18, 0x88, 0xac,
		0xe7, 0xd0, 0x24, 0xc4, 0xb7, 0x70, 0x4d, 0xd3, 0xf9, 0xbe, 0x31, 0x2b, 0x88, 0xed, 0x97, 0x98,
		0xbb, 0xfd, 0x52, 0xf8, 0xff, 0x61, 0xa2, 0x62, 0x34, 0xc2, 0xe2, 0x16, 0xd2, 0xa1, 0xcd, 0x05,
		0xfb, 0xa2, 0xf4, 0xfc, 0x43, 0x1c, 0xa9, 0x66, 0xd4, 0x55, 0xbd, 0x36, 0x67, 0x58, 0x35, 0xef,
		0x98, 0x95, 0x64, 0x3c, 0xb6, 0xef, 0xb0, 0xd5, 0xdc, 0xfa, 0x53, 0x49, 0xfa, 0xe5, 0x48, 0x74,
		0x61, 0xad, 0xf0, 0xb9, 0x48, 0x76, 0x81, 0x11, 0xae, 0x09, 0x63, 0x28, 0x78, 0xbb, 0x8e, 0x2b,
		0x44, 0x41, 0xf8, 0xfe, 0x03, 0x30, 0x59, 0x33, 0x6a, 0x06, 0xe5, 0x74, 0x9a, 0xfc, 0xe2, 0xe7,
		0xb4, 0x49, 0x17, 0x9a, 0xed, 0x79, 0xa8, 0x9b, 0x5b, 0x81, 0x09, 0x8e, 0x5c, 0xa6, 0x07, 0x45,
		0x6c, 0x19, 0x83, 0xba, 0xee, 0xa1, 0x65, 0xbe, 0xf8, 0x3a, 0x9d, 0xbe, 0x95, 0x71, 0x4e, 0x4a,
		0xea, 0xd8, 0x4a, 0x27, 0xa7, 0xc0, 0x91, 0x00, 0x3f, 0x36, 0x48, 0xb1, 0xd5, 0x83, 0xe3, 0x6f,
		0x71, 0x8e, 0x13, 0x3e, 0x8e, 0xeb, 0x9c, 0x34, 0x37, 0x0f, 0x23, 0x07, 0xe1, 0xf5, 0x4f, 0x38,
		0xaf, 0x14, 0xf6, 0x33, 0x59, 0x80, 0x31, 0xca, 0xa4, 0xd2, 0xb4, 0x1d, 0xa3, 0x41, 0x23, 0x60,
		0x77, 0x36, 0xbf, 0xfd, 0x3a, 0x1b, 0x35, 0xa3, 0x84, 0x6c, 0xde, 0xa5, 0xca, 0xe5, 0x80, 0x9e,
		0x8d, 0x55, 0x71, 0xa5, 0xde, 0x83, 0xc3, 0xd7, 0xb9, 0x20, 0x2e, 0x7e, 0xee, 0x32, 0x4c, 0x92,
		0xdf, 0x34, 0x40, 0xf9, 0x25, 0xe9, 0xbd, 0xe1, 0x96, 0xf9, 0xf6, 0xfb, 0xd8, 0xc0, 0x9c, 0x70,
		0x19, 0xf8, 0x64, 0xf2, 0xf5, 0x62, 0x0d, 0x3b, 0x0e, 0xb6, 0xec, 0xb2, 0x5a, 0x6f, 0x27, 0x9e,
		0x6f, 0xc7, 0x22, 0xf3, 0xf1, 0x1f, 0x04, 0x7b, 0x71, 0x81, 0x51, 0xe6, 0xeb, 0xf5, 0xdc, 0x26,
		0xdc, 0xd1, 0xc6, 0x2b, 0xfa, 0xe0, 0xf9, 0x0a, 0xe7, 0x39, 0xd9, 0xe2, 0x19, 0x84, 0xed, 0x1a,
		0x08, 0xb8, 0xdb, 0x97, 0x7d, 0xf0, 0xfc, 0x04, 0xe7, 0x89, 0x38, 0xad, 0xe8, 0x52, 0xc2, 0xf1,
		0x69, 0x18, 0xbf, 0x82, 0xad, 0x2d, 0xc3, 0xe6, 0xbb, 0x44, 0x7d, 0xb0, 0xfb, 0x24, 0x67, 0x37,
		0xc6, 0x09, 0xe9, 0xb6, 0x11, 0xe1, 0xf5, 0x24, 0x24, 0xb6, 0xd5, 0x0a, 0xee, 0x83, 0xc5, 0x75,
		0xce, 0x62, 0x88, 0xe0, 0x13, 0xd2, 0x3c, 0xa4, 0x6a, 0x06, 0x9f, 0xa3, 0x7a, 0x93, 0x7f, 0x8a,
		0x93, 0x0f, 0x0b, 0x1a, 0xce, 0xc2, 0x34, 0xcc, 0x66, 0x9d, 0x4c, 0x60, 0xbd, 0x59, 0xfc, 0x35,
		0xc1, 0x42, 0xd0, 0x70, 0x16, 0x07, 0x30, 0xeb, 0xab, 0x82, 0x85, 0xed, 0xb3, 0xe7, 0x53, 0x30,
		0x6c, 0xe8, 0xf5, 0x7d, 0x43, 0xef, 0x47, 0x88, 0x4f, 0x73, 0x0e, 0xc0, 0x49, 0x08, 0x83, 0xf3,
		0x90, 0xec, 0xb7, 0x23, 0xfe, 0xfa, 0x0f, 0xc4, 0xf0, 0x10, 0x3d, 0xb0, 0x00, 0x63, 0x22, 0x40,
		0x69, 0x86, 0xde, 0x07, 0x8b, 0xbf, 0xc1, 0x59, 0x8c, 0xfa, 0xc8, 0xb8, 0x1a, 0x0e, 0xb6, 0x9d,
		0x1a, 0xee, 0x87, 0xc9, 0x67, 0x85, 0x1a, 0x9c, 0x84, 0x9b, 0x72, 0x0b, 0xeb, 0x95, 0x9d, 0xfe,
		0x38, 0xfc, 0xaa, 0x30, 0xa5, 0xa0, 0x21, 0x2c, 0xe6, 0x61, 0xa4, 0xa1, 0x5a, 0xf6, 0x8e, 0x5a,
		0xef, 0xab, 0x3b, 0xfe, 0x26, 0xe7, 0x91, 0x72, 0x89, 0xb8, 0x45, 0x9a, 0xfa, 0x41, 0xd8, 0x7c,
		0x4e, 0x58, 0xc4, 0x47, 0xc6, 0x87, 0x9e, 0xed, 0xd0, 0x2d, 0xb5, 0x83, 0x70, 0xfb, 0x35, 0x31,
		0xf4, 0x18, 0xed, 0xb2, 0x9f, 0xe3, 0x79, 0x48, 0xda, 0xda, 0xb5, 0xbe, 0xd8, 0x7c, 0x5e, 0xf4,
		0x34, 0x25, 0x20, 0xc4, 0x6f, 0x83, 0xa3, 0x6d, 0xa7, 0x89, 0x3e, 0x98, 0xfd, 0x2d, 0xce, 0x6c,
		0xaa, 0xcd, 0x54, 0xc1, 0x43, 0xc2, 0x41, 0x59, 0xfe, 0x6d, 0x11, 0x12, 0x70, 0x88, 0xd7, 0x1a,
		0x59, 0x35, 0xd8, 0xea, 0xf6, 0xc1, 0xac, 0xf6, 0xeb, 0xc2, 0x6a, 0x8c, 0x36, 0x60, 0xb5, 0x0d,
		0x98, 0xe2, 0x1c, 0x0f, 0xd6, 0xaf, 0x5f, 0x10, 0x81, 0x95, 0x51, 0x6f, 0x06, 0x7b, 0xf7, 0xed,
		0x90, 0x75, 0xcd, 0x29, 0xd2, 0x53, 0xbb, 0xdc, 0x50, 0xcd, 0x3e, 0x38, 0x7f, 0x91, 0x73, 0x16,
		0x11, 0xdf, 0xcd, 0x6f, 0xed, 0x65, 0xd5, 0x24, 0xcc, 0x9f, 0x83, 0x8c, 0x60, 0xde, 0xd4, 0x2d,
		0x5c, 0x31, 0x6a, 0xba, 0x76, 0x0d, 0x57, 0xfb, 0x60, 0xfd, 0x1b, 0xa1, 0xae, 0xda, 0xf4, 0x91,
		0x13, 0xce, 0x8b, 0x90, 0x76, 0x73, 0x95, 0xb2, 0xd6, 0x30, 0x0d, 0xcb, 0xe9, 0xc1, 0xf1, 0x4b,
		0xa2, 0xa7, 0x5c, 0xba, 0x45, 0x4a, 0x96, 0x2b, 0x01, 0x3b, 0x67, 0xee, 0xd7, 0x25, 0xbf, 0xcc,
		0x19, 0x8d, 0x78, 0x54, 0x3c, 0x70, 0x54, 0x8c, 0x86, 0xa9, 0x5a, 0xfd, 0xc4, 0xbf, 0xbf, 0x23,
		0x02, 0x07, 0x27, 0xe1, 0x81, 0x83, 0x64, 0x74, 0x64, 0xb6, 0xef, 0x83, 0xc3, 0x57, 0x44, 0xe0,
		0x10, 0x34, 0x9c, 0x85, 0x48, 0x18, 0xfa, 0x60, 0xf1, 0x77, 0x05, 0x0b, 0x41, 0x43, 0x58, 0x3c,
		0xe3, 0x4d, 0xb4, 0x16, 0xae, 0x69, 0xb6, 0x63, 0xb1, 0xa4, 0xb8, 0x3b, 0xab, 0xbf, 0xf7, 0x83,
		0x60, 0x12, 0xa6, 0xf8, 0x48, 0x49, 0x24, 0xe2, 0x9b, 0xac, 0x74, 0xcd, 0xd4, 0x5b, 0xb0, 0xdf,
		0x14, 0x91, 0xc8, 0x47, 0x46, 0x64, 0xf3, 0x65, 0x88, 0xc4, 0xec, 0x15, 0xb2, 0x52, 0xe8, 0x83,
		0xdd, 0xdf, 0x0f, 0x09, 0xb7, 0x2e, 0x68, 0x09, 0x4f, 0x5f, 0xfe, 0xd3, 0xd4, 0x77, 0xf1, 0x7e,
		0x5f, 0xde, 0xf9, 0x0f, 0x42, 0xf9, 0xcf, 0x26, 0xa3, 0x64, 0x31, 0x64, 0x2c, 0x94, 0x4f, 0xa1,
		0x5e, 0xb7, 0x8a, 0x32, 0xef, 0xf9, 0x11, 0xd7, 0x37, 0x98, 0x4e, 0xe5, 0x96, 0x88, 0x93, 0x07,
		0x93, 0x9e, 0xde, 0xcc, 0xde, 0xf7, 0x23, 0xd7, 0xcf, 0x03, 0x39, 0x4f, 0xee, 0x02, 0x8c, 0x04,
		0x12, 0x9e, 0xde, 0xac, 0xde, 0xcf, 0x59, 0xa5, 0xfc, 0xf9, 0x4e, 0xee, 0x2c, 0xc4, 0x48, 0xf2,
		0xd2, 0x9b, 0xfc, 0x2f, 0x70, 0x72, 0x8a, 0x9e, 0x7b, 0x33, 0x24, 0x44, 0xd2, 0xd2, 0x9b, 0xf4,
		0x03, 0x9c, 0xd4, 0x25, 0x21, 0xe4, 0x22, 0x61, 0xe9, 0x4d, 0xfe, 0x17, 0x05, 0xb9, 0x20, 0x21,
		0xe4, 0xfd, 0x9b, 0xf0, 0xab, 0x3f, 0x1f, 0xe3, 0x93, 0x8e, 0xb0, 0xdd, 0x79, 0x18, 0xe2, 0x99,
		0x4a, 0x6f, 0xea, 0x0f, 0xf1, 0xc6, 0x05, 0x45, 0xee, 0x71, 0x88, 0xf7, 0x69, 0xf0, 0x5f, 0xe0,
		0xa4, 0x0c, 0x3f, 0x37, 0x0f, 0xc3, 0xbe, 0xec, 0xa4, 0x37, 0xf9, 0x5f, 0xe2, 0xe4, 0x7e, 0x2a,
		0x22, 0x3a, 0xcf, 0x4e, 0x7a, 0x33, 0xf8, 0x45, 0x21, 0x3a, 0xa7, 0x20, 0x66, 0x13, 0x89, 0x49,
		0x6f, 0xea, 0x0f, 0x0b, 0xab, 0x0b, 0x92, 0xdc, 0x53, 0x90, 0x74, 0x27, 0x9b, 0xde, 0xf4, 0x1f,
		0xe1, 0xf4, 0x1e, 0x0d, 0xb1, 0x80, 0x6f, 0xb2, 0xeb, 0xcd, 0xe2, 0x2f, 0x0b, 0x0b, 0xf8, 0xa8,
		0xc8, 0x30, 0x0a, 0x27, 0x30, 0xbd, 0x39, 0x7d, 0x54, 0x0c, 0xa3, 0x50, 0xfe, 0x42, 0x7a, 0x93,
		0xc6, 0xfc, 0xde, 0x2c, 0xfe, 0x8a, 0xe8, 0x4d, 0x8a, 0x4f, 0xc4, 0x08, 0x67, 0x04, 0xbd, 0x79,
		0xfc, 0x92, 0x10, 0x23, 0x94, 0x10, 0xe4, 0xd6, 0x00, 0xb5, 0x66, 0x03, 0xbd, 0xf9, 0x7d, 0x8c,
		0xf3, 0x1b, 0x6f, 0x49, 0x06, 0x72, 0xcf, 0xc2, 0x54, 0xfb, 0x4c, 0xa0, 0x37, 0xd7, 0x8f, 0xff,
		0x28, 0xb4, 0x76, 0xf3, 0x27, 0x02, 0xb9, 0x0d, 0x6f, 0x4a, 0xf1, 0x67, 0x01, 0xbd, 0xd9, 0xbe,
		0xf2, 0xa3, 0x60, 0xe0, 0xf6, 0x27, 0x01, 0xb9, 0x3c, 0x80, 0x37, 0x01, 0xf7, 0xe6, 0xf5, 0x49,
		0xce, 0xcb, 0x47, 0x44, 0x86, 0x06, 0x9f, 0x7f, 0x7b, 0xd3, 0x5f, 0x17, 0x43, 0x83, 0x53, 0x90,
		0xa1, 0x21, 0xa6, 0xde, 0xde, 0xd4, 0x9f, 0x12, 0x43, 0x43, 0x90, 0x10, 0xcf, 0xf6, 0xcd, 0x6e,
		0xbd, 0x39, 0x7c, 0x5a, 0x78, 0xb6, 0x8f, 0x2a, 0xb7, 0x02, 0xe3, 0x2d, 0x13, 0x62, 0x6f, 0x56,
		0xbf, 0xcc, 0x59, 0xa5, 0xc3, 0xf3, 0xa1, 0x7f, 0xf2, 0xe2, 0x93, 0x61, 0x6f, 0x6e, 0x9f, 0x09,
		0x4d, 0x5e, 0x7c, 0x2e, 0xcc, 0x9d, 0x87, 0x84, 0xde, 0xac, 0xd7, 0xc9, 0xe0, 0x41, 0xdd, 0x6f,
		0x02, 0x66, 0xfe, 0xe3, 0x8f, 0xb9, 0x75, 0x04, 0x41, 0xee, 0x2c, 0xc4, 0x71, 0x63, 0x0b, 0x57,
		0x7b, 0x51, 0x7e, 0xff, 0xc7, 0x22, 0x60, 0x12, 0xec, 0xdc, 0x53, 0x00, 0x6c, 0x6b, 0x84, 0x1e,
		0x06, 0xf6, 0xa0, 0xfd, 0x4f, 0x3f, 0xe6, 0x57, 0x6f, 0x3c, 0x12, 0x8f, 0x01, 0xbb, 0xc8, 0xd3,
		0x9d, 0xc1, 0x0f, 0x82, 0x0c, 0x68, 0x8f, 0x3c, 0x09, 0x43, 0x2f, 0xd8, 0x86, 0xee, 0xa8, 0xb5,
		0x5e, 0xd4, 0xff, 0x99, 0x53, 0x0b, 0x7c, 0x62, 0xb0, 0x86, 0x61, 0x61, 0x47, 0xad, 0xd9, 0xbd,
		0x68, 0xff, 0x0b, 0xa7, 0x75, 0x09, 0x08, 0x71, 0x45, 0xb5, 0x9d, 0x7e, 0xf4, 0xfe, 0x23, 0x41,
		0x2c, 0x08, 0x88, 0xd0, 0xe4, 0xf7, 0x2e, 0xde, 0xef, 0x45, 0xfb, 0x43, 0x21, 0x34, 0xc7, 0xcf,
		0xbd, 0x19, 0x92, 0xe4, 0x27, 0xbb, 0x4f, 0xd7, 0x83, 0xf8, 0x8f, 0x39, 0xb1, 0x47, 0x41, 0x5a,
		0xb6, 0x9d, 0xaa, 0xa3, 0xf5, 0x36, 0xf6, 0x4d, 0xde, 0xd3, 0x02, 0x3f, 0x97, 0x87, 0x61, 0xdb,
		0xa9, 0x56, 0x9b, 0x3c, 0x3f, 0xed, 0x41, 0xfe, 0x27, 0x3f, 0x76, 0xb7, 0x2c, 0x5c, 0x1a, 0xd2,
		0xdb, 0x57, 0x77, 0x1d, 0xd3, 0xa0, 0x07, 0x1e, 0xbd, 0x38, 0xfc, 0x88, 0x73, 0xf0, 0x91, 0xe4,
		0xe6, 0x21, 0x45, 0x74, 0xb1, 0xb0, 0x89, 0xe9, 0xe9, 0x54, 0x0f, 0x16, 0xff, 0x95, 0x1b, 0x20,
		0x40, 0x54, 0xf8, 0xb9, 0xaf, 0xbf, 0x36, 0x2d, 0x7d, 0xeb, 0xb5, 0x69, 0xe9, 0x0f, 0x5f, 0x9b,
		0x96, 0x3e, 0xfc, 0xdd, 0xe9, 0x81, 0x6f, 0x7d, 0x77, 0x7a, 0xe0, 0xf7, 0xbe, 0x3b, 0x3d, 0xd0,
		0x7e, 0x97, 0x18, 0x16, 0x8c, 0x05, 0x83, 0xed, 0x0f, 0x3f, 0x2f, 0xd7, 0x34, 0x67, 0xa7, 0xb9,
		0x35, 0x57, 0x31, 0x1a, 0x74, 0x1b, 0xd7, 0xdb, 0xad, 0x75, 0x17, 0x39, 0xf0, 0xde, 0x28, 0x1c,
		0xad, 0x18, 0x76, 0xc3, 0xb0, 0xcb, 0x6c, 0xbf, 0x97, 0x15, 0xf8, 0x8e, 0x6f, 0xca, 0x5f, 0xd5,
		0xc7, 0xa6, 0xef, 0x45, 0x18, 0xa5, 0xaa, 0xd3, 0xed, 0x2e, 0xea, 0x6d, 0x3d, 0x03, 0xc4, 0x37,
		0xfe, 0x55, 0x9c, 0x6a, 0x3d, 0xe2, 0x12, 0xd2, 0xd3, 0xfb, 0x0d, 0x98, 0xd4, 0x1a, 0x66, 0x1d,
		0xd3, 0x6d, 0xfe, 0xb2, 0x5b, 0xd7, 0x9b, 0xdf, 0x37, 0x39, 0xbf, 0x09, 0x8f, 0x7c, 0x51, 0x50,
		0xe7, 0x96, 0x60, 0x5c, 0xad, 0x54, 0xb0, 0x19, 0x60, 0xd9, 0xa3, 0x5b, 0x84, 0x80, 0x69, 0x4e,
		0xe9, 0x72, 0x2b, 0x3c, 0xd5, 0xa9, 0x6b, 0x9e, 0xbf, 0xc7, 0x67, 0x79, 0x0b, 0xd7, 0xb0, 0xfe,
		0x90, 0x8e, 0x9d, 0xab, 0x86, 0xb5, 0xcb, 0xcd, 0xfb, 0x10, 0x6b, 0x6a, 0x90, 0xdd, 0x60, 0x86,
		0xf7, 0x47, 0x61, 0x9a, 0x55, 0x9c, 0xde, 0x52, 0x6d, 0x7c, 0xfa, 0xca, 0x23, 0x5b, 0xd8, 0x51,
		0x1f, 0x39, 0x5d, 0x31, 0x34, 0x9d, 0xf7, 0xc4, 0x04, 0xef, 0x17, 0x52, 0x3f, 0xc7, 0xeb, 0xb3,
		0x6d, 0xb7, 0xe9, 0xe5, 0x05, 0x88, 0xcd, 0x1b, 0x9a, 0x8e, 0x26, 0x21, 0x5e, 0xc5, 0xba, 0xd1,
		0xe0, 0x77, 0xee, 0x58, 0x01, 0xdd, 0x05, 0x83, 0x6a, 0xc3, 0x68, 0xea, 0x0e, 0x3b, 0xa1, 0x28,
		0x0c, 0x7f, 0xfd, 0xc6, 0xcc, 0xc0, 0xef, 0xdf, 0x98, 0x89, 0x2e, 0xea, 0x8e, 0xc2, 0xab, 0x72,
		0xb1, 0xef, 0xbd, 0x3a, 0x23, 0xc9, 0x4f, 0xc3, 0x50, 0x11, 0x57, 0x0e, 0xc3, 0xab, 0x88, 0x2b,
		0x21, 0x5e, 0xf7, 0x43, 0x62, 0x51, 0x77, 0xd8, 0xad, 0xc8, 0x13, 0x10, 0xd5, 0x74, 0x76, 0xd1,
		0x26, 0xd4, 0x3e, 0x81, 0x13, 0xd4, 0x22, 0xae, 0xb8, 0xa8, 0x55, 0x5c, 0x09, 0xa3, 0x12, 0xf6,
		0x04, 0x5e, 0x28, 0xfe, 0xde, 0xbf, 0x9b, 0x1e, 0x78, 0xe9, 0xb5, 0xe9, 0x81, 0x8e, 0x3d, 0xe1,
		0x1f, 0x03, 0xdc, 0xc4, 0xbc, 0x0b, 0xec, 0xea, 0x2e, 0x3b, 0x23, 0x71, 0xbb, 0xe1, 0x77, 0x06,
		0x41, 0xe6, 0x38, 0xb6, 0xa3, 0xee, 0x6a, 0x7a, 0xcd, 0xed, 0x09, 0xb5, 0xe9, 0xec, 0x5c, 0xe3,
		0x5d, 0x31, 0xc5, 0xbb, 0x82, 0xe3, 0x74, 0xef, 0x8d, 0x6c, 0xe7, 0xd1, 0x95, 0xed, 0xd1, 0xe7,
		0xf2, 0x3f, 0x8b, 0x02, 0x5a, 0x77, 0xd4, 0x5d, 0x9c, 0x6f, 0x3a, 0x3b, 0x86, 0xa5, 0x5d, 0x63,
		0xb1, 0x0c, 0x03, 0x34, 0xd4, 0xbd, 0xb2, 0x63, 0xec, 0x62, 0xdd, 0xa6, 0xa6, 0x19, 0x3e, 0x73,
		0x74, 0xae, 0x8d, 0x7f, 0xcc, 0x91, 0xae, 0x2b, 0x3c, 0xf0, 0xb9, 0xef, 0xcc, 0xdc, 0xd7, 0xdb,
		0x0a, 0x14, 0x99, 0x24, 0xd7, 0x7b, 0x1b, 0x94, 0x31, 0xba, 0x0c, 0xec, 0x92, 0x45, 0xb9, 0xae,
		0xd9, 0x0e, 0xbf, 0xa7, 0x7d, 0x76, 0xae, 0xbd, 0xee, 0x73, 0xad, 0x62, 0xce, 0x5d, 0x56, 0xeb,
		0x5a, 0x55, 0x75, 0x0c, 0xcb, 0xbe, 0x38, 0xa0, 0x24, 0x29, 0xab, 0x25, 0xcd, 0x76, 0xd0, 0x06,
		0x24, 0xab, 0x58, 0xdf, 0x67, 0x6c, 0xa3, 0x6f, 0x8c, 0x6d, 0x82, 0x70, 0xa2, 0x5c, 0x9f, 0x03,
		0xa4, 0xfa, 0xf1, 0xc4, 0xc3, 0x24, 0x76, 0xbf, 0xb2, 0x03, 0xfb, 0x00, 0x67, 0xfa, 0x8e, 0x62,
		0x5c, 0x0d, 0x83, 0xb2, 0xf7, 0x02, 0x78, 0x6d, 0xa2, 0x0c, 0x0c, 0xa9, 0xd5, 0xaa, 0x85, 0x6d,
		0x9b, 0x1e, 0x00, 0x26, 0x15, 0x51, 0xcc, 0x8d, 0xff, 0x8b, 0x2f, 0x3f, 0x34, 0x12, 0xe0, 0x58,
		0x48, 0x01, 0x5c, 0x71, 0x49, 0x4f, 0x7d, 0x4a, 0x82, 0xf1, 0x96, 0x16, 0x91, 0x0c, 0xd3, 0xf9,
		0xcd, 0x8d, 0x8b, 0xab, 0xca, 0xe2, 0xf3, 0xf9, 0x8d, 0xc5, 0xd5, 0x95, 0x32, 0xbb, 0xf2, 0xbf,
		0xb2, 0xbe, 0x56, 0x9a, 0x5f, 0xbc, 0xb0, 0x58, 0x2a, 0xa6, 0x07, 0xd0, 0x0c, 0x1c, 0x6b, 0x83,
		0x53, 0x2c, 0x2d, 0x95, 0x16, 0xf2, 0x1b, 0xa5, 0xb4, 0x84, 0xee, 0x84, 0x13, 0x6d, 0x99, 0xb8,
		0x28, 0x91, 0x0e, 0x28, 0x4a, 0xc9, 0x45, 0x89, 0x16, 0x2e, 0x74, 0x1c, 0x45, 0x0f, 0x76, 0xf5,
		0x9f, 0x3d, 0x77, 0xb8, 0x04, 0xc7, 0xd3, 0x7b, 0x22, 0x70, 0x34, 0x3c, 0x65, 0xa8, 0xfa, 0x7e,
		0x87, 0x57, 0x9f, 0x1d, 0xa2, 0xd9, 0x45, 0x88, 0xe6, 0xf5, 0x7d, 0x74, 0x94, 0xe5, 0xd3, 0xe5,
		0xa6, 0x55, 0xe7, 0x31, 0x68, 0x88, 0x94, 0x37, 0xad, 0x3a, 0x89, 0x4d, 0xe2, 0xa2, 0xbf, 0x74,
		0x32, 0xc5, 0x6f, 0xef, 0xe7, 0xd2, 0x1f, 0x7b, 0x75, 0x66, 0xe0, 0x0b, 0xaf, 0xce, 0x0c, 0xfc,
		0xf0, 0xd3, 0x33, 0x03, 0x2f, 0xfd, 0xc1, 0xec, 0x40, 0x61, 0x37, 0xac, 0xde, 0x57, 0x7b, 0xce,
		0xa6, 0x89, 0xbc, 0xbe, 0x4f, 0x03, 0xd1, 0x9a, 0xf4, 0x7c, 0x9c, 0x2a, 0x27, 0x0e, 0x50, 0xa7,
		0xc3, 0x07, 0xa8, 0xcf, 0xe2, 0x7a, 0xfd, 0x92, 0x6e, 0x5c, 0xa5, 0xbd, 0xea, 0xd9, 0xe0, 0xa3,
		0x11, 0x98, 0x6e, 0x99, 0x36, 0x79, 0x86, 0xd1, 0xe9, 0xf9, 0x6b, 0x0e, 0x12, 0x45, 0x91, 0xb8,
		0x64, 0x60, 0xc8, 0xc6, 0x15, 0x43, 0xaf, 0xb2, 0x91, 0x1e, 0x55, 0x44, 0x91, 0xa8, 0xad, 0xab,
		0xba, 0x61, 0xf3, 0x3b, 0xf7, 0xac, 0x50, 0xf8, 0x84, 0x74, 0xb0, 0x7c, 0x61, 0x44, 0xb4, 0x24,
		0xd4, 0x7c, 0xa4, 0xe7, 0x91, 0xf2, 0x2e, 0xd1, 0xd2, 0x55, 0x22, 0x70, 0xac, 0xdc, 0xaf, 0x55,
		0x7e, 0x29, 0x02, 0x33, 0x61, 0xab, 0x90, 0xb4, 0xcd, 0x76, 0xd4, 0x86, 0xd9, 0xc9, 0x2c, 0xe7,
		0x21, 0xb9, 0x21, 0x70, 0x0e, 0x6c, 0x97, 0xeb, 0x07, 0xb4, 0xcb, 0xa8, 0xdb, 0x94, 0x30, 0xcc,
		0x99, 0x3e, 0x0d, 0xe3, 0xea, 0x71, 0x28, 0xcb, 0x7c, 0x2e, 0x06, 0x27, 0xe8, 0xa3, 0x2c, 0xab,
		0xa1, 0xe9, 0xce, 0xe9, 0x8a, 0xb5, 0x6f, 0x3a, 0x34, 0x71, 0x33, 0xb6, 0xb9, 0x5d, 0xc6, 0xbd,
		0xea, 0x39, 0x56, 0xdd, 0x61, 0xe4, 0x6c, 0x43, 0x7c, 0x8d, 0xd0, 0x11, 0x8b, 0x38, 0x86, 0xa3,
		0xd6, 0xb9, 0xa5, 0x58, 0x81, 0x40, 0xd9, 0x43, 0xae, 0x08, 0x83, 0x6a, 0xe2, 0x0d, 0x57, 0x1d,
		0xab, 0xdb, 0xec, 0x3e, 0x7c, 0x94, 0x0e, 0xa8, 0x04, 0x01, 0xd0, 0xab, 0xef, 0x93, 0x10, 0x57,
		0x9b, 0xec, 0x2a, 0x47, 0x94, 0x8c, 0x34, 0x5a, 0x90, 0x2f, 0xc1, 0x10, 0x3f, 0x50, 0x46, 0x69,
		0x88, 0xee, 0xe2, 0x7d, 0xda, 0x4e, 0x4a, 0x21, 0x3f, 0xd1, 0x1c, 0xc4, 0xa9, 0xf0, 0x7c, 0x02,
		0xc9, 0xcc, 0xb5, 0x48, 0x3f, 0x47, 0x85, 0x54, 0x18, 0x9a, 0xfc, 0x34, 0x24, 0x8a, 0x46, 0x43,
		0xd3, 0x8d, 0x20, 0xb7, 0x24, 0xe3, 0x46, 0x65, 0x36, 0x9b, 0x3c, 0xdf, 0x50, 0x58, 0x01, 0x4d,
		0xc1, 0x20, 0x7b, 0x1f, 0xc1, 0xaf, 0xa3, 0xf0, 0x92, 0x3c, 0x0f, 0x43, 0x94, 0xf7, 0xaa, 0x89,
		0x10, 0x7f, 0x59, 0xc7, 0x1f, 0x62, 0xd0, 0xd4, 0x94, 0xb3, 0x8f, 0x78, 0xc2, 0x22, 0x88, 0x55,
		0x55, 0x47, 0xe5, 0x7a, 0xd3, 0xdf, 0xf2, 0x5b, 0x20, 0xc1, 0x99, 0xd8, 0xe8, 0x0c, 0x44, 0x0d,
		0xd3, 0xe6, 0x17, 0x4a, 0xb2, 0x9d, 0x54, 0x59, 0x35, 0x0b, 0x31, 0x92, 0xa9, 0x28, 0x04, 0xb9,
		0xa0, 0x74, 0x0c, 0xaa, 0x4f, 0xf8, 0x82, 0xaa, 0xaf, 0xcb, 0x7d, 0x3f, 0x59, 0x97, 0xb6, 0xb8,
		0x83, 0xeb, 0x2c, 0x9f, 0x8e, 0xc0, 0xb4, 0xaf, 0xf6, 0x0a, 0xb6, 0x6c, 0xcd, 0xd0, 0xf9, 0x7c,
		0xce, 0xbc, 0x05, 0xf9, 0x84, 0xe4, 0xf5, 0x1d, 0xdc, 0xe5, 0xcd, 0x10, 0xcd, 0x9b, 0x26, 0xca,
		0x42, 0x82, 0x96, 0x2b, 0x06, 0xf3, 0x97, 0x98, 0xe2, 0x96, 0x49, 0x9d, 0x6d, 0x6c, 0x3b, 0x57,
		0x55, 0xcb, 0x7d, 0x42, 0x28, 0xca, 0xf2, 0x93, 0x90, 0x9c, 0x37, 0x74, 0x1b, 0xeb, 0x76, 0x93,
		0x8e, 0xc1, 0xad, 0xba, 0x51, 0xd9, 0xe5, 0x1c, 0x58, 0x81, 0x18, 0x5c, 0x35, 0x4d, 0x4a, 0x19,
		0x53, 0xc8, 0x4f, 0x96, 0x1b, 0x16, 0xd6, 0x3b, 0x9a, 0xe8, 0xc9, 0x83, 0x9b, 0x88, 0x2b, 0xe9,
		0xda, 0xe8, 0x7f, 0x4a, 0x70, 0xbc, 0x75, 0x40, 0xed, 0xe2, 0x7d, 0xfb, 0xa0, 0xe3, 0xe9, 0x39,
		0x48, 0xae, 0xd1, 0x77, 0xfc, 0x97, 0xf0, 0x3e, 0xca, 0xc2, 0x10, 0xae, 0x9e, 0x39, 0x7b, 0xf6,
		0x91, 0x27, 0x99, 0xb7, 0x5f, 0x1c, 0x50, 0x04, 0x00, 0x4d, 0x43, 0xd2, 0xc6, 0x15, 0xf3, 0xcc,
		0xd9, 0x73, 0xbb, 0x8f, 0x30, 0xf7, 0x22, 0x19, 0x90, 0x0b, 0xca, 0x25, 0x88, 0xd6, 0xdf, 0xfb,
		0xf4, 0x8c, 0x54, 0x88, 0x43, 0xd4, 0x6e, 0x36, 0x6e, 0xab, 0x8f, 0xbc, 0x12, 0x87, 0x59, 0x3f,
		0x25, 0x8d, 0x54, 0x6e, 0x56, 0xc2, 0x6d, 0x90, 0xf6, 0xd9, 0x80, 0x62, 0x74, 0x48, 0x66, 0xbb,
		0x5a, 0x52, 0xfe, 0x0d, 0x09, 0x52, 0x6e, 0xaa, 0xb4, 0x8e, 0x1d, 0x74, 0xde, 0x9f, 0xff, 0xf0,
		0x61, 0x73, 0x6c, 0x2e, 0xdc, 0x96, 0x97, 0xd2, 0x29, 0x3e, 0x74, 0xf4, 0x38, 0x75, 0x44, 0xd3,
		0xb0, 0xf9, 0xb3, 0xb2, 0x1e, 0xa4, 0x2e, 0x32, 0x7a, 0x10, 0x10, 0x8d, 0x70, 0xe5, 0x2b, 0x86,
		0xa3, 0xe9, 0xb5, 0xb2, 0x69, 0x5c, 0xe5, 0x8f, 0x75, 0xa3, 0x4a, 0x9a, 0xd6, 0x5c, 0xa6, 0x15,
		0x6b, 0x04, 0x4e, 0x84, 0x4e, 0xba, 0x5c, 0x82, 0xe9, 0x1d, 0x09, 0x02, 0xa2, 0x88, 0xce, 0xc3,
		0x90, 0xd9, 0xdc, 0x2a, 0x8b, 0x88, 0x31, 0x7c, 0xe6, 0x78, 0xbb, 0xf1, 0x2f, 0xfc, 0x83, 0x47,
		0x80, 0x41, 0xb3, 0xb9, 0x45, 0xbc, 0xe5, 0x4e, 0x48, 0xb5, 0x11, 0x66, 0xf8, 0x8a, 0x27, 0x07,
		0xfd, 0x7c, 0x04, 0xd7, 0xa0, 0x6c, 0x5a, 0x9a, 0x61, 0x69, 0xce, 0x3e, 0xcd, 0x5f, 0xa3, 0x4a,
		0x5a, 0x54, 0xac, 0x71, 0xb8, 0xbc, 0x0b, 0x63, 0xeb, 0x74, 0x7d, 0xeb, 0x49, 0x7e, 0xd6, 0x93,
		0x4f, 0xea, 0x2d, 0x5f, 0x47, 0xc9, 0x22, 0x2d, 0x92, 0x15, 0x9e, 0xe9, 0xe8, 0x9d, 0x8f, 0x1f,
		0xdc, 0x3b, 0x83, 0x19, 0xe2, 0x1f, 0x1d, 0x0d, 0x0c, 0x4e, 0xe6, 0x9c, 0xfe, 0xf0, 0xd5, 0xaf,
		0x63, 0xf6, 0xca, 0x26, 0xb2, 0xdd, 0x27, 0xd5, 0x6c, 0x8f, 0x30, 0x9a, 0xed, 0x39, 0x84, 0xe4,
		0x27, 0x61, 0x64, 0x4d, 0xb5, 0x9c, 0x75, 0xec, 0x5c, 0xc4, 0x6a, 0x15, 0x5b, 0xc1, 0x59, 0x77,
		0x44, 0xcc, 0xba, 0x08, 0x62, 0x74, 0x6a, 0x65, 0xb3, 0x0e, 0xfd, 0x2d, 0xef, 0x40, 0x8c, 0xde,
		0x0c, 0x75, 0x67, 0x64, 0x4e, 0xc1, 0x66, 0x64, 0x12, 0x4b, 0xf7, 0x1d, 0x6c, 0x8b, 0xf4, 0x96,
		0x16, 0xd0, 0x63, 0x62, 0x5e, 0x8d, 0x76, 0x9f, 0x57, 0xb9, 0x23, 0xf2, 0xd9, 0xb5, 0x0e, 0x43,
		0x05, 0x12, 0x8a, 0x17, 0x8b, 0xae, 0x20, 0x92, 0x27, 0x08, 0x5a, 0x86, 0x31, 0x53, 0xb5, 0x1c,
		0xfa, 0x24, 0x66, 0x87, 0x6a, 0xc1, 0x7d, 0x7d, 0xa6, 0x75, 0xe4, 0x05, 0x94, 0xe5, 0xad, 0x8c,
		0x98, 0x7e, 0xa0, 0xfc, 0xef, 0x63, 0x30, 0xc8, 0x8d, 0xf1, 0x66, 0x18, 0xe2, 0x66, 0xe5, 0xde,
		0x79, 0x62, 0xae, 0x75, 0x62, 0x9a, 0x73, 0x27, 0x10, 0xce, 0x4f, 0xd0, 0xa0, 0x7b, 0x21, 0x51,
		0xd9, 0x51, 0x35, 0xbd, 0xac, 0x55, 0xc5, 0x56, 0xc3, 0x6b, 0x37, 0x66, 0x86, 0xe6, 0x09, 0x6c,
		0xb1, 0xa8, 0x0c, 0xd1, 0xca, 0xc5, 0x2a, 0xc9, 0x04, 0x76, 0xb0, 0x56, 0xdb, 0x71, 0xf8, 0x08,
		0xe3, 0x25, 0xf4, 0x04, 0xc4, 0x88, 0x43, 0xf0, 0x07, 0x93, 0xd9, 0x96, 0x0d, 0x1f, 0x37, 0xd9,
		0x2b, 0x24, 0x48, 0xc3, 0x1f, 0xfe, 0xce, 0x8c, 0xa4, 0x50, 0x0a, 0x34, 0x0f, 0x23, 0x75, 0xd5,
		0x76, 0xca, 0x74, 0x06, 0x23, 0xcd, 0xc7, 0xf9, 0x7a, 0xbb, 0xc5, 0x20, 0xdc, 0xb0, 0x5c, 0xf4,
		0x61, 0x42, 0xc5, 0x40, 0x55, 0x74, 0x12, 0xd2, 0x94, 0x49, 0xc5, 0x68, 0x34, 0x34, 0x87, 0xe5,
		0x56, 0x83, 0xd4, 0xee, 0xa3, 0x04, 0x3e, 0x4f, 0xc1, 0x34, 0xc3, 0x3a, 0x06, 0x49, 0xfa, 0x44,
		0x8b, 0xa2, 0xb0, 0xeb, 0xc8, 0x09, 0x02, 0xa0, 0x95, 0xf7, 0xc1, 0x98, 0x17, 0x1f, 0x19, 0x4a,
		0x82, 0x71, 0xf1, 0xc0, 0x14, 0xf1, 0x61, 0x98, 0xd4, 0xf1, 0x1e, 0xbd, 0x20, 0x1d, 0xc0, 0x4e,
		0x52, 0x6c, 0x44, 0xea, 0x2e, 0x07, 0x29, 0xee, 0x81, 0xd1, 0x8a, 0x30, 0x3e, 0xc3, 0x05, 0x8a,
		0x3b, 0xe2, 0x42, 0x29, 0xda, 0x51, 0x48, 0xa8, 0xa6, 0xc9, 0x10, 0x86, 0x79, 0x7c, 0x34, 0x4d,
		0x5a, 0x75, 0x0a, 0xc6, 0xa9, 0x8e, 0x16, 0xb6, 0x9b, 0x75, 0x87, 0x33, 0x49, 0x51, 0x9c, 0x31,
		0x52, 0xa1, 0x30, 0x38, 0xc5, 0xbd, 0x0b, 0x46, 0xf0, 0x15, 0xad, 0x8a, 0xf5, 0x0a, 0x66, 0x78,
		0x23, 0x14, 0x2f, 0x25, 0x80, 0x14, 0xe9, 0x7e, 0x70, 0xe3, 0x5e, 0x59, 0xc4, 0xe4, 0x51, 0xc6,
		0x4f, 0xc0, 0xf3, 0x0c, 0x2c, 0x67, 0x20, 0x56, 0x54, 0x1d, 0x95, 0x24, 0x18, 0xce, 0x1e, 0x9b,
		0x68, 0x52, 0x0a, 0xf9, 0x29, 0x7f, 0x2f, 0x02, 0xb1, 0xcb, 0x86, 0x83, 0xd1, 0xa3, 0xbe, 0x04,
		0x70, 0xb4, 0x9d, 0x3f, 0xaf, 0x6b, 0x35, 0x1d, 0x57, 0x97, 0xed, 0x9a, 0xef, 0x7b, 0x0a, 0x9e,
		0x3b, 0x45, 0x02, 0xee, 0x34, 0x09, 0x71, 0xcb, 0x68, 0xea, 0x55, 0x71, 0x93, 0x97, 0x16, 0x50,
		0x09, 0x12, 0xae, 0x97, 0xc4, 0x7a, 0x79, 0xc9, 0x18, 0xf1, 0x12, 0xe2, 0xc3, 0x1c, 0xa0, 0x0c,
		0x6d, 0x71, 0x67, 0x29, 0x40, 0xd2, 0x0d, 0x5e, 0xdc, 0xdb, 0xfa, 0x73, 0x58, 0x8f, 0x8c, 0x4c,
		0x26, 0x6e, 0xdf, 0xbb, 0xc6, 0x63, 0x1e, 0x97, 0x76, 0x2b, 0xb8, 0xf5, 0x02, 0x6e, 0xc5, 0xbf,
		0xed, 0x30, 0x44, 0xf5, 0xf2, 0xdc, 0x8a, 0x7d, 0xdf, 0xe1, 0x38, 0x24, 0x6d, 0xad, 0xa6, 0xab,
		0x4e, 0xd3, 0xc2, 0xdc, 0xf3, 0x3c, 0x80, 0xfc, 0x55, 0x09, 0x06, 0x99, 0x27, 0xfb, 0xec, 0x26,
		0xb5, 0xb7, 0x5b, 0xa4, 0x93, 0xdd, 0xa2, 0x87, 0xb7, 0x5b, 0x1e, 0xc0, 0x15, 0xc6, 0xe6, 0x4f,
		0xee, 0xdb, 0x64, 0x0c, 0x4c, 0xc4, 0x75, 0xad, 0xc6, 0x07, 0xaa, 0x8f, 0x48, 0xfe, 0xb7, 0x12,
		0x49, 0x62, 0x79, 0x3d, 0xca, 0xc3, 0x88, 0x90, 0xab, 0xbc, 0x5d, 0x57, 0x6b, 0xdc, 0x77, 0x4e,
		0x74, 0x14, 0xee, 0x42, 0x5d, 0xad, 0x29, 0xc3, 0x5c, 0x1e, 0x52, 0x68, 0xdf, 0x0f, 0x91, 0x0e,
		0xfd, 0x10, 0xe8, 0xf8, 0xe8, 0xe1, 0x3a, 0x3e, 0xd0, 0x45, 0xb1, 0x70, 0x17, 0x7d, 0x29, 0x42,
		0x17, 0x33, 0xa6, 0x61, 0xab, 0xf5, 0x9f, 0xc4, 0x88, 0x38, 0x06, 0x49, 0xd3, 0xa8, 0x97, 0x59,
		0x0d, 0xbb, 0xe1, 0x9e, 0x30, 0x8d, 0xba, 0xd2, 0xd2, 0xed, 0xf1, 0x5b, 0x34, 0x5c, 0x06, 0x6f,
		0x81, 0xd5, 0x86, 0xc2, 0x56, 0xb3, 0x20, 0xc5, 0x4c, 0xc1, 0xe7, 0xb2, 0x87, 0x89, 0x0d, 0xe8,
		0xe4, 0x28, 0xb5, 0xce, 0xbd, 0x4c, 0x6c, 0x86, 0xa9, 0x70, 0x3c, 0x42, 0xc1, 0x42, 0x7f, 0xbb,
		0x55, 0xb0, 0xdf, 0x2d, 0x15, 0x8e, 0x27, 0xff, 0x55, 0x09, 0x60, 0x89, 0x58, 0x96, 0xea, 0x4b,
		0x66, 0x21, 0x9b, 0x8a, 0x50, 0x0e, 0xb4, 0x3c, 0xdd, 0xa9, 0xd3, 0x78, 0xfb, 0x29, 0xdb, 0x2f,
		0xf7, 0x3c, 0x8c, 0x78, 0xce, 0x68, 0x63, 0x21, 0xcc, 0x74, 0x97, 0xac, 0x7a, 0x1d, 0x3b, 0x4a,
		0xea, 0x8a, 0xaf, 0x24, 0xff, 0x63, 0x09, 0x92, 0x54, 0xa6, 0x65, 0xec, 0xa8, 0x81, 0x3e, 0x94,
		0x0e, 0xdf, 0x87, 0x27, 0x00, 0x18, 0x1b, 0x5b, 0xbb, 0x86, 0xb9, 0x67, 0x25, 0x29, 0x64, 0x5d,
		0xbb, 0x86, 0xd1, 0x39, 0xd7, 0xe0, 0xd1, 0xee, 0x06, 0x17, 0x59, 0x37, 0x37, 0xfb, 0x1d, 0x30,
		0x44, 0x3f, 0x51, 0xb5, 0x67, 0xf3, 0x44, 0x7a, 0x50, 0x6f, 0x36, 0x36, 0xf6, 0x6c, 0xf9, 0x05,
		0x18, 0xda, 0xd8, 0x63, 0x7b, 0x23, 0xc7, 0x20, 0x69, 0x19, 0x06, 0x9f, 0x93, 0x59, 0x2e, 0x94,
		0x20, 0x00, 0x3a, 0x05, 0x89, 0xfd, 0x80, 0x88, 0xb7, 0x1f, 0xe0, 0x6d, 0x68, 0x44, 0xfb, 0xda,
		0xd0, 0x38, 0xf5, 0xaf, 0x25, 0x18, 0xf6, 0xc5, 0x07, 0xf4, 0x08, 0x1c, 0x29, 0x2c, 0xad, 0xce,
		0x5f, 0x2a, 0x2f, 0x16, 0xcb, 0x17, 0x96, 0xf2, 0x0b, 0xde, 0x1b, 0xae, 0xec, 0xd4, 0xcb, 0xd7,
		0x67, 0x91, 0x0f, 0x77, 0x53, 0xa7, 0x3b, 0x4a, 0xe8, 0x34, 0x4c, 0x06, 0x49, 0xf2, 0x85, 0xf5,
		0xd2, 0xca, 0x46, 0x5a, 0xca, 0x1e, 0x79, 0xf9, 0xfa, 0xec, 0xb8, 0x8f, 0x22, 0xbf, 0x65, 0x63,
		0xdd, 0x69, 0x25, 0x98, 0x5f, 0x5d, 0x5e, 0x5e, 0xdc, 0x48, 0x47, 0x5a, 0x08, 0x78, 0xc0, 0xbe,
		0x1f, 0xc6, 0x83, 0x04, 0x2b, 0x8b, 0x4b, 0xe9, 0x68, 0x16, 0xbd, 0x7c, 0x7d, 0x76, 0xd4, 0x87,
		0xbd, 0xa2, 0xd5, 0xb3, 0x89, 0x0f, 0x7e, 0x66, 0x7a, 0xe0, 0x57, 0x7f, 0x65, 0x5a, 0x22, 0x9a,
		0x8d, 0x04, 0x62, 0x04, 0x7a, 0x10, 0xee, 0x58, 0x5f, 0x5c, 0x58, 0x29, 0x15, 0xcb, 0xcb, 0xeb,
		0x0b, 0x62, 0x0f, 0x5a, 0x68, 0x37, 0xf6, 0xf2, 0xf5, 0xd9, 0x61, 0xae, 0x52, 0x27, 0xec, 0x35,
		0xa5, 0x74, 0x79, 0x75, 0xa3, 0x94, 0x96, 0x18, 0xf6, 0x9a, 0x85, 0xaf, 0x18, 0x0e, 0xfb, 0x86,
		0xdd, 0xc3, 0x70, 0xb4, 0x0d, 0xb6, 0xab, 0xd8, 0xf8, 0xcb, 0xd7, 0x67, 0x47, 0xd6, 0x2c, 0xcc,
		0xc6, 0x0f, 0xa5, 0x98, 0x83, 0x4c, 0x2b, 0xc5, 0xea, 0xda, 0xea, 0x7a, 0x7e, 0x29, 0x3d, 0x9b,
		0x4d, 0xbf, 0x7c, 0x7d, 0x36, 0x25, 0x82, 0x21, 0xdd, 0xe8, 0x77, 0x35, 0xbb, 0x9d, 0x2b, 0x9e,
		0x3f, 0x79, 0x08, 0xee, 0xee, 0x70, 0xc6, 0x24, 0x4e, 0x27, 0x0e, 0x75, 0xca, 0xd4, 0x71, 0x9f,
		0x3d, 0xdb, 0x63, 0xfb, 0xb9, 0xf7, 0xd2, 0xe9, 0xf0, 0x27, 0x58, 0xd9, 0xae, 0x8b, 0x3b, 0xf9,
		0x43, 0x12, 0x8c, 0x5e, 0xd4, 0x6c, 0xc7, 0xb0, 0xb4, 0x8a, 0x5a, 0xa7, 0x2f, 0xb7, 0xce, 0xf5,
		0x1b, 0x5b, 0x43, 0x43, 0xfd, 0x29, 0x18, 0xbc, 0xa2, 0xd6, 0x59, 0x50, 0x8b, 0xd2, 0x0f, 0xcd,
		0x74, 0x38, 0xf2, 0x71, 0x43, 0x9b, 0x60, 0xc0, 0xc8, 0xe4, 0x5f, 0x8f, 0xc0, 0x18, 0x1d, 0x0c,
		0x36, 0xfb, 0x04, 0x19, 0x59, 0x63, 0x15, 0x20, 0x66, 0xa9, 0x0e, 0xdf, 0x34, 0x2c, 0xcc, 0xf1,
		0xd3, 0xc7, 0x7b, 0xfb, 0x38, 0x4b, 0x2b, 0xe2, 0x8a, 0x42, 0x69, 0xd1, 0x3b, 0x20, 0xd1, 0x50,
		0xf7, 0xca, 0x94, 0x0f, 0x5b, 0xb9, 0xe4, 0x0f, 0xc6, 0xe7, 0xe6, 0x8d, 0x99, 0xb1, 0x7d, 0xb5,
		0x51, 0xcf, 0xc9, 0x82, 0x8f, 0xac, 0x0c, 0x35, 0xd4, 0x3d, 0x22, 0x22, 0x32, 0x61, 0x8c, 0x40,
		0x2b, 0x3b, 0xaa, 0x5e, 0xc3, 0xac, 0x11, 0xba, 0x05, 0x5a, 0xb8, 0x78, 0xe0, 0x46, 0xa6, 0xbc,
		0x46, 0x7c, 0xec, 0x64, 0x65, 0xa4, 0xa1, 0xee, 0xcd, 0x53, 0x00, 0x69, 0x31, 0x97, 0xf8, 0xd8,
		0xab, 0x33, 0x03, 0xf4, 0x44, 0xf7, 0xdb, 0x12, 0x80, 0x67, 0x31, 0xf4, 0x0e, 0x48, 0x57, 0xdc,
		0x12, 0xa5, 0x15, 0x67, 0x93, 0xf7, 0x75, 0xea, 0x8b, 0x90, 0xbd, 0xd9, 0xdc, 0xfc, 0xad, 0x1b,
		0x33, 0x92, 0x32, 0x56, 0x09, 0x75, 0xc5, 0xdb, 0x61, 0xb8, 0x69, 0x56, 0x55, 0x07, 0x97, 0xe9,
		0x3a, 0x2e, 0xd2, 0x73, 0x9e, 0x9f, 0x26, 0xbc, 0x6e, 0xde, 0x98, 0x41, 0x4c, 0x2d, 0x1f, 0xb1,
		0x4c, 0x67, 0x7f, 0x60, 0x10, 0x42, 0xe0, 0xd3, 0xe9, 0x1b, 0x12, 0x0c, 0x17, 0x7d, 0x77, 0x2a,
		0x33, 0x30, 0xd4, 0x30, 0x74, 0x6d, 0x97, 0xfb, 0x63, 0x52, 0x11, 0x45, 0x94, 0x85, 0x04, 0x7b,
		0xcc, 0xea, 0xec, 0x8b, 0xad, 0x50, 0x51, 0x26, 0x54, 0x57, 0xf1, 0x96, 0xad, 0x89, 0xde, 0x50,
		0x44, 0x11, 0x5d, 0x80, 0xb4, 0x8d, 0x2b, 0x4d, 0x4b, 0x73, 0xf6, 0xcb, 0x15, 0x43, 0x77, 0xd4,
		0x8a, 0xc3, 0x9e, 0x45, 0x16, 0x8e, 0xdd, 0xbc, 0x31, 0x73, 0x07, 0x93, 0x35, 0x8c, 0x21, 0x2b,
		0x63, 0x02, 0x34, 0xcf, 0x20, 0xa4, 0x85, 0x2a, 0x76, 0x54, 0xad, 0x6e, 0x67, 0xd8, 0xe5, 0x04,
		0x51, 0xf4, 0xe9, 0xf2, 0xf9, 0x21, 0xff, 0xc6, 0xd6, 0x05, 0x48, 0x1b, 0x26, 0xb6, 0x02, 0x89,
		0xa8, 0x14, 0x6e, 0x39, 0x8c, 0x21, 0x2b, 0x63, 0x02, 0x24, 0x92, 0x54, 0x87, 0x74, 0xb3, 0x58,
		0x28, 0x9a, 0xcd, 0x2d, 0x6f, 0x3f, 0x6c, 0xb2, 0xa5, 0x37, 0xf2, 0xfa, 0x7e, 0xe1, 0x51, 0x8f,
		0x7b, 0x98, 0x4e, 0xfe, 0xe6, 0x97, 0x1f, 0x9a, 0xe4, 0xae, 0xe1, 0xed, 0x4f, 0x5d, 0xc2, 0xfb,
		0xa4, 0xfb, 0x39, 0xea, 0x1a, 0xc5, 0x24, 0x69, 0xe7, 0x0b, 0xaa, 0x56, 0x17, 0xcf, 0xfb, 0x15,
		0x5e, 0x42, 0x39, 0x18, 0xb4, 0x1d, 0xd5, 0x69, 0xda, 0xfc, 0xa4, 0x57, 0xee, 0xe4, 0x6a, 0x05,
		0x43, 0xaf, 0xae, 0x53, 0x4c, 0x85, 0x53, 0xa0, 0x0b, 0x30, 0xc8, 0x8f, 0xd0, 0xe3, 0x07, 0x1e,
		0xdf, 0xf4, 0xae, 0x04, 0xa3, 0x26, 0x16, 0xa9, 0xe2, 0x3a, 0xae, 0xb1, 0xb4, 0x6a, 0x47, 0x25,
		0xab, 0x0f, 0xfa, 0xed, 0xbd, 0xc2, 0xe2, 0x81, 0x07, 0x21, 0xb7, 0x54, 0x98, 0x9f, 0xac, 0x8c,
		0xb9, 0xa0, 0x75, 0x0a, 0x41, 0x97, 0x02, 0x97, 0x7f, 0xf9, 0x07, 0x2a, 0xef, 0xea, 0xa4, 0xbe,
		0xcf, 0xa7, 0xc5, 0xfe, 0x84, 0xff, 0xea, 0xf0, 0x05, 0x48, 0x37, 0xf5, 0x2d, 0x43, 0xa7, 0x6f,
		0x70, 0x79, 0x7e, 0x4f, 0xd6, 0x77, 0x51, 0xbf, 0x73, 0x84, 0x31, 0x64, 0x65, 0xcc, 0x05, 0x5d,
		0x64, 0xab, 0x80, 0x2a, 0x8c, 0x7a, 0x58, 0x74, 0xa0, 0x26, 0x7b, 0x0e, 0xd4, 0x3b, 0xf9, 0x40,
		0x3d, 0x12, 0x6e, 0xc5, 0x1b, 0xab, 0x23, 0x2e, 0x90, 0x90, 0xa1, 0x8b, 0x00, 0x5e, 0x78, 0xa0,
		0xfb, 0x14, 0xc3, 0x9d, 0x3b, 0xde, 0x8b, 0x31, 0x62, 0xbd, 0xe7, 0xd1, 0xa2, 0x77, 0xc1, 0x44,
		0x43, 0xd3, 0xcb, 0x36, 0xae, 0x6f, 0x97, 0xb9, 0x81, 0x09, 0x4b, 0xfa, 0x09, 0xa5, 0xc2, 0xd2,
		0xc1, 0xfc, 0xe1, 0xe6, 0x8d, 0x99, 0x2c, 0x0f, 0xa1, 0xad, 0x2c, 0x65, 0x65, 0xbc, 0xa1, 0xe9,
		0xeb, 0xb8, 0xbe, 0x5d, 0x74, 0x61, 0xb9, 0xd4, 0x07, 0x5f, 0x9d, 0x19, 0xe0, 0xc3, 0x75, 0x40,
		0x3e, 0x47, 0xf7, 0xce, 0xf9, 0x30, 0xc3, 0x36, 0x59, 0x93, 0xa8, 0xa2, 0xc0, 0xaf, 0x1a, 0x78,
		0x00, 0x36, 0xcc, 0x5f, 0xfa, 0x83, 0x59, 0x49, 0xfe, 0xbc, 0x04, 0x83, 0xc5, 0xcb, 0x6b, 0xaa,
		0x66, 0xa1, 0x45, 0x18, 0xf7, 0x3c, 0x27, 0x38, 0xc8, 0x8f, 0xdf, 0xbc, 0x31, 0x93, 0x09, 0x3b,
		0x97, 0x3b, 0xca, 0x3d, 0x07, 0x16, 0xc3, 0x7c, 0xb1, 0xd3, 0xc2, 0x35, 0xc0, 0xaa, 0x05, 0x45,
		0x6e, 0x5d, 0xd6, 0x86, 0xd4, 0x2c, 0xc1, 0x10, 0x93, 0xd6, 0x46, 0x39, 0x88, 0x9b, 0xe4, 0x07,
		0x3f, 0x18, 0x98, 0xee, 0xe8, 0xbc, 0x14, 0xdf, 0xdd, 0xc8, 0x24, 0x24, 0xf2, 0x47, 0x22, 0x00,
		0xc5, 0xcb, 0x97, 0x37, 0x2c, 0xcd, 0xac, 0x63, 0xe7, 0x56, 0x6a, 0xbe, 0x01, 0x47, 0x7c, 0xab,
		0x24, 0xab, 0x12, 0xd2, 0x7e, 0xf6, 0xe6, 0x8d, 0x99, 0xe3, 0x61, 0xed, 0x7d, 0x68, 0xb2, 0x32,
		0xe1, 0xad, 0x97, 0xac, 0x4a, 0x5b, 0xae, 0x55, 0xdb, 0x71, 0xb9, 0x46, 0x3b, 0x73, 0xf5, 0xa1,
		0xf9, 0xb9, 0x16, 0x6d, 0xa7, 0xbd, 0x69, 0xd7, 0x61, 0xd8, 0x33, 0x89, 0x8d, 0x8a, 0x90, 0x70,
		0xf8, 0x6f, 0x6e, 0x61, 0xb9, 0xb3, 0x85, 0x05, 0x19, 0xb7, 0xb2, 0x4b, 0x29, 0xff, 0xa9, 0x04,
		0xe0, 0xf9, 0xec, 0xcf, 0xa6, 0x8b, 0x91, 0x50, 0xce, 0x03, 0x6f, 0xf4, 0x50, 0xa9, 0x1a, 0xa7,
		0x0e, 0xd9, 0xf3, 0xe7, 0x23, 0x30, 0xb1, 0x29, 0x22, 0xcf, 0xcf, 0xbc, 0x0d, 0xd6, 0x60, 0x08,
		0xeb, 0x8e, 0xa5, 0x51, 0x23, 0x90, 0xde, 0x7e, 0xb8, 0x53, 0x6f, 0xb7, 0xd1, 0x89, 0x7e, 0x44,
		0x4a, 0x6c, 0xba, 0x73, 0x36, 0x21, 0x6b, 0xfc, 0x62, 0x14, 0x32, 0x9d, 0x28, 0xd1, 0x3c, 0x8c,
		0x55, 0x2c, 0xcc, 0x2e, 0x5e, 0xf9, 0x77, 0xfe, 0x0a, 0x59, 0x2f, 0xb3, 0x0c, 0x21, 0xc8, 0xca,
		0xa8, 0x80, 0xf0, 0xd9, 0xa3, 0x06, 0x24, 0xed, 0x23, 0x6e, 0x47, 0xef, 0x6f, 0xf5, 0x97, 0xe7,
		0xc9, 0x7c, 0xfa, 0x10, 0x8d, 0x04, 0x19, 0xb0, 0xf9, 0x63, 0xd4, 0x83, 0xd2, 0x09, 0xe4, 0x45,
		0x18, 0xd3, 0x74, 0xcd, 0xd1, 0xd4, 0x7a, 0x79, 0x4b, 0xad, 0xab, 0x7a, 0xe5, 0x30, 0x59, 0x33,
		0x0b, 0xf9, 0xbc, 0xd9, 0x10, 0x3b, 0x59, 0x19, 0xe5, 0x90, 0x02, 0x03, 0xa0, 0x8b, 0x30, 0x24,
		0x9a, 0x8a, 0x1d, 0x2a, 0xdb, 0x10, 0xe4, 0xbe, 0x04, 0xef, 0x17, 0xa2, 0x30, 0xae, 0xe0, 0xea,
		0xff, 0xeb, 0x8a, 0x83, 0x75, 0xc5, 0x32, 0x00, 0x1b, 0xee, 0x24, 0xc0, 0x1e, 0xa2, 0x37, 0x48,
		0xc0, 0x48, 0x32, 0x0e, 0x45, 0xdb, 0xf1, 0xf5, 0xc7, 0x8d, 0x08, 0xa4, 0xfc, 0xfd, 0xf1, 0xe7,
		0x74, 0x56, 0x42, 0x8b, 0x5e, 0x24, 0x8a, 0xf1, 0x4f, 0xef, 0x76, 0x88, 0x44, 0x2d, 0xde, 0xdb,
		0x3d, 0x04, 0xfd, 0xb7, 0x08, 0x0c, 0xae, 0xa9, 0x96, 0xda, 0xb0, 0x51, 0xa5, 0x25, 0xd3, 0x14,
		0xdb, 0x8f, 0x2d, 0x1f, 0x58, 0xe7, 0xbb, 0x1d, 0x3d, 0x12, 0xcd, 0x8f, 0xb5, 0x49, 0x34, 0xdf,
		0x0a, 0xa3, 0x64, 0x39, 0xec, 0xbb, 0xc2, 0x40, 0xac, 0x3d, 0x52, 0x38, 0xea, 0x71, 0x09, 0xd6,
		0xb3, 0xd5, 0xf2, 0x65, 0xff, 0x1d, 0x86, 0x61, 0x82, 0xe1, 0x05, 0x66, 0x42, 0x3e, 0xe5, 0x2d,
		0x4b, 0x7d, 0x95, 0xb2, 0x02, 0x0d, 0x75, 0xaf, 0xc4, 0x0a, 0x68, 0x09, 0xd0, 0x8e, 0xbb, 0x33,
		0x52, 0xf6, 0xcc, 0x49, 0xe8, 0x4f, 0xdc, 0xbc, 0x31, 0x73, 0x94, 0xd1, 0xb7, 0xe2, 0xc8, 0xca,
		0xb8, 0x07, 0x14, 0xdc, 0x1e, 0x03, 0x20, 0x7a, 0x95, 0xd9, 0x15, 0x6e, 0xb6, 0xdc, 0x39, 0x72,
		0xf3, 0xc6, 0xcc, 0x38, 0xe3, 0xe2, 0xd5, 0xc9, 0x4a, 0x92, 0x14, 0x8a, 0xe4, 0xb7, 0xcf, 0xb3,
		0x3f, 0x23, 0x01, 0xf2, 0x42, 0xbe, 0x82, 0x6d, 0x93, 0xac, 0xcf, 0x48, 0x22, 0xee, 0xcb, 0x9a,
		0xa5, 0xee, 0x89, 0xb8, 0x47, 0x2f, 0x12, 0x71, 0xdf, 0x48, 0x79, 0xd2, 0x0b, 0x8f, 0x91, 0x5e,
		0xf7, 0x99, 0xb9, 0x8b, 0x84, 0xe3, 0xe1, 0x80, 0xfc, 0x4f, 0x25, 0x38, 0xda, 0xe2, 0x51, 0xae,
		0xb0, 0xff, 0x1f, 0x20, 0xcb, 0x57, 0xc9, 0xbf, 0xa3, 0xc8, 0x84, 0x3e, 0xb0, 0x83, 0x8e, 0x5b,
		0x2d, 0x71, 0xf7, 0xd6, 0x45, 0x78, 0x76, 0x61, 0xfe, 0x1f, 0x49, 0x30, 0xe9, 0x6f, 0xde, 0x55,
		0x64, 0x05, 0x52, 0xfe, 0xd6, 0xb9, 0x0a, 0x77, 0xf7, 0xa3, 0x02, 0x97, 0x3e, 0x40, 0x8f, 0x9e,
		0xf1, 0x86, 0x2b, 0xdb, 0x3b, 0x7b, 0xa4, 0x6f, 0x6b, 0x08, 0x99, 0xc2, 0xc3, 0x36, 0x46, 0xfb,
		0xe3, 0x7f, 0x4b, 0x10, 0x5b, 0x33, 0x8c, 0x3a, 0x32, 0x60, 0x5c, 0x37, 0x9c, 0x32, 0xf1, 0x2c,
		0x5c, 0xf5, 0xdf, 0x5b, 0x4f, 0x16, 0xe6, 0x0f, 0x66, 0xa4, 0xef, 0xdf, 0x98, 0x69, 0x65, 0xa5,
		0x8c, 0xe9, 0x86, 0x53, 0xa0, 0x10, 0x7e, 0x75, 0xfd, 0x5d, 0x30, 0x12, 0x6c, 0x8c, 0x45, 0xc9,
		0x67, 0x0f, 0xdc, 0x58, 0x90, 0xcd, 0xcd, 0x1b, 0x33, 0x93, 0xde, 0x88, 0x71, 0xc1, 0xb2, 0x92,
		0xda, 0xf2, 0xb5, 0xce, 0xae, 0x77, 0xfd, 0xf0, 0xd5, 0x19, 0xe9, 0xd4, 0x57, 0x24, 0x00, 0x6f,
		0xe7, 0x01, 0x3d, 0x08, 0x77, 0x14, 0x56, 0x57, 0x8a, 0xe5, 0xf5, 0x8d, 0xfc, 0xc6, 0xe6, 0x7a,
		0xf0, 0x8e, 0xb7, 0xd8, 0x1e, 0xb7, 0x4d, 0x5c, 0xd1, 0xb6, 0x35, 0x5c, 0x45, 0xf7, 0xc2, 0x64,
		0x10, 0x9b, 0x94, 0x4a, 0xc5, 0xb4, 0x94, 0x4d, 0xbd, 0x7c, 0x7d, 0x36, 0xc1, 0x72, 0x31, 0x5c,
		0x45, 0x27, 0xe1, 0x48, 0x2b, 0xde, 0xe2, 0xca, 0x42, 0x3a, 0x92, 0x1d, 0x79, 0xf9, 0xfa, 0x6c,
		0xd2, 0x4d, 0xda, 0x90, 0x0c, 0xc8, 0x8f, 0xc9, 0xf9, 0x45, 0xb3, 0xf0, 0xf2, 0xf5, 0xd9, 0x41,
		0x66, 0xc0, 0x6c, 0xec, 0x83, 0x9f, 0x99, 0x1e, 0xb8, 0xe5, 0x37, 0xc1, 0xff, 0x78, 0xa8, 0xe3,
		0xae, 0x77, 0x0d, 0xeb, 0xd8, 0xd6, 0xec, 0x43, 0xed, 0x7a, 0xf7, 0xb5, 0x93, 0x2e, 0xff, 0x6e,
		0x1c, 0x52, 0x0b, 0xac, 0x15, 0xd2, 0x11, 0x18, 0xbd, 0x09, 0x06, 0x4d, 0x3a, 0x8d, 0xb8, 0xc7,
		0x68, 0x1d, 0x1c, 0x9e, 0x4d, 0x36, 0xee, 0x5d, 0x2e, 0x36, 0xf5, 0xd8, 0xfc, 0x32, 0x07, 0xbb,
		0x63, 0xe6, 0xdd, 0x9a, 0x4a, 0x1d, 0x68, 0xbf, 0x87, 0xe5, 0x2c, 0x7c, 0x6b, 0x25, 0xcc, 0x4f,
		0x66, 0xf7, 0x42, 0x36, 0x08, 0x84, 0xdd, 0x0e, 0x7b, 0xbf, 0x04, 0x47, 0x28, 0x96, 0x37, 0x11,
		0x53, 0x4c, 0x91, 0xec, 0x9f, 0xea, 0xa4, 0xc2, 0x92, 0x6a, 0x7b, 0x77, 0x3d, 0xd8, 0x7d, 0xae,
		0xbb, 0xf9, 0x44, 0x78, 0xdc, 0xd7, 0x78, 0x98, 0xad, 0xac, 0x4c, 0xd4, 0x5b, 0x28, 0x6d, 0xb4,
		0x10, 0xb8, 0xd0, 0x17, 0x3b, 0xd8, 0x56, 0xbb, 0xff, 0x72, 0xdf, 0xd3, 0x30, 0xec, 0xc5, 0x12,
		0x9b, 0xff, 0xdf, 0x97, 0xfe, 0xe7, 0x0e, 0x3f, 0x31, 0xfa, 0x80, 0x04, 0x47, 0xbc, 0xd9, 0xdc,
		0xcf, 0x96, 0xfd, 0x7f, 0x9c, 0x07, 0x0e, 0xb0, 0x10, 0x0a, 0x1b, 0xa7, 0x2d, 0x5f, 0x59, 0x99,
		0x6c, 0xb6, 0x92, 0x92, 0x25, 0xd8, 0x88, 0x3f, 0xb2, 0xda, 0x19, 0xf1, 0x09, 0xc8, 0xfe, 0x43,
		0x73, 0x90, 0x01, 0xfb, 0x9f, 0x1d, 0xa6, 0x61, 0x39, 0xb8, 0x4a, 0x37, 0xe4, 0x12, 0x8a, 0x5b,
		0x96, 0x57, 0x00, 0xb5, 0x76, 0x6e, 0xf8, 0x02, 0xa3, 0xf7, 0x3e, 0x05, 0x4d, 0x42, 0xdc, 0x7f,
		0xc5, 0x8f, 0x15, 0x72, 0x89, 0x0f, 0xf2, 0xe9, 0xf3, 0x96, 0x8f, 0xf9, 0xef, 0x44, 0xe0, 0x94,
		0xff, 0x78, 0xe8, 0xc5, 0x26, 0xb6, 0xf6, 0xdd, 0x21, 0x6a, 0xaa, 0x35, 0x4d, 0xf7, 0xbf, 0x82,
		0x38, 0xea, 0x9f, 0xf0, 0x29, 0xae, 0xb0, 0x93, 0xfc, 0x41, 0x09, 0x86, 0xd7, 0xd4, 0x1a, 0x56,
		0xf0, 0x8b, 0x4d, 0x6c, 0x3b, 0x6d, 0x6e, 0x99, 0x4f, 0xc1, 0xa0, 0xb1, 0xbd, 0x2d, 0xce, 0xb4,
		0x63, 0x0a, 0x2f, 0x11, 0x9d, 0xeb, 0x5a, 0x43, 0x63, 0xd7, 0xc1, 0x62, 0x0a, 0x2b, 0xa0, 0x19,
		0x18, 0xae, 0x18, 0x4d, 0x9d, 0x0f, 0xb9, 0x4c, 0x4c, 0x7c, 0x6b, 0xa5, 0xa9, 0xb3, 0x21, 0x47,
		0x8c, 0x68, 0xe1, 0x2b, 0xd8, 0xb2, 0xd9, 0xd7, 0x25, 0x13, 0x8a, 0x28, 0xca, 0x4f, 0x41, 0x8a,
		0x49, 0xc2, 0x27, 0xe3, 0xa3, 0x90, 0xa0, 0x37, 0xad, 0x3c, 0x79, 0x86, 0x48, 0xf9, 0x12, 0xbb,
		0xab, 0xce, 0xf8, 0x33, 0x91, 0x58, 0xa1, 0x50, 0xe8, 0x68, 0xe5, 0x93, 0xbd, 0xa3, 0x06, 0xb3,
		0xa1, 0x6b, 0xe1, 0xdf, 0x8a, 0xc3, 0x11, 0x7e, 0x78, 0xa7, 0x9a, 0xda, 0xe9, 0x1d, 0xc7, 0x11,
		0x6f, 0x27, 0x80, 0x67, 0xc1, 0xaa, 0xa9, 0xc9, 0xfb, 0x10, 0xbb, 0xe8, 0x38, 0x26, 0x3a, 0x05,
		0x71, 0xab, 0x59, 0xc7, 0x62, 0x33, 0xc8, 0xdd, 0xae, 0x57, 0x4d, 0x6d, 0x8e, 0x20, 0x28, 0xcd,
		0x3a, 0x56, 0x18, 0x0a, 0x2a, 0xc1, 0xcc, 0x76, 0xb3, 0x5e, 0xdf, 0x2f, 0x57, 0x31, 0xfd, 0x77,
		0x59, 0xee, 0x3f, 0x9c, 0xc0, 0x7b, 0xa6, 0x2a, 0x3e, 0x5b, 0x49, 0x0c, 0x73, 0x9c, 0xa2, 0x15,
		0x29, 0x96, 0xf8, 0x67, 0x13, 0x25, 0x81, 0x23, 0xff, 0x7e, 0x04, 0x12, 0x82, 0x35, 0xbd, 0x3c,
		0x8e, 0xeb, 0xb8, 0xe2, 0x18, 0xe2, 0x30, 0xc5, 0x2d, 0x23, 0x04, 0xd1, 0x1a, 0xef, 0xbc, 0xe4,
		0xc5, 0x01, 0x85, 0x14, 0x08, 0xcc, 0xbd, 0xd2, 0x4f, 0x60, 0x66, 0x93, 0xf4, 0x67, 0xcc, 0x34,
		0xc4, 0xaa, 0xed, 0xe2, 0x80, 0x42, 0x4b, 0x28, 0x03, 0x83, 0x64, 0xd0, 0x38, 0xac, 0xb7, 0x08,
		0x9c, 0x97, 0xd1, 0x14, 0xc4, 0x4d, 0xd5, 0xa9, 0xb0, 0xdb, 0x76, 0xa4, 0x82, 0x15, 0xd1, 0xe3,
		0x30, 0xc8, 0x5e, 0x65, 0x87, 0xff, 0x17, 0x0d, 0x31, 0x06, 0xfb, 0xfc, 0x1d, 0x91, 0x7b, 0x4d,
		0x75, 0x1c, 0x6c, 0xe9, 0x84, 0x21, 0x43, 0x47, 0x08, 0x62, 0x5b, 0x46, 0x75, 0x9f, 0xff, 0x7f,
		0x1c, 0xfa, 0x9b, 0xff, 0x43, 0x0e, 0xea, 0x0f, 0x65, 0x5a, 0xc9, 0xfe, 0x2d, 0x58, 0x4a, 0x00,
		0x0b, 0x04, 0xa9, 0x04, 0x13, 0x6a, 0xb5, 0xaa, 0xb1, 0x7f, 0x55, 0x53, 0xde, 0xd2, 0x68, 0xf0,
		0xb0, 0xe9, 0x3f, 0x7d, 0xeb, 0xd4, 0x17, 0xc8, 0x23, 0x28, 0x70, 0xfc, 0x42, 0x12, 0x86, 0x4c,
		0x26, 0x94, 0x7c, 0x1e, 0xc6, 0x5b, 0x24, 0x25, 0xf2, 0xed, 0x6a, 0x7a, 0x55, 0xbc, 0x73, 0x20,
		0xbf, 0x09, 0x8c, 0x7e, 0xb0, 0x92, 0x1d, 0x53, 0xd1, 0xdf, 0x85, 0xf7, 0x76, 0x7e, 0x0e, 0x33,
		0xea, 0x7b, 0x0e, 0xa3, 0x9a, 0x5a, 0x21, 0x49, 0xf9, 0xf3, 0x47, 0x30, 0xf9, 0xd6, 0x47, 0x30,
		0x35, 0xac, 0x8b, 0x89, 0x99, 0x54, 0xa9, 0xa6, 0x66, 0x53, 0x77, 0xf4, 0x3e, 0xa0, 0x69, 0x9f,
		0xf7, 0xfd, 0xa6, 0x6f, 0x62, 0x62, 0x0b, 0xf9, 0xb5, 0x45, 0xd7, 0x8f, 0xbf, 0x16, 0x81, 0xe3,
		0x3e, 0x3f, 0xf6, 0x21, 0xb7, 0xba, 0x73, 0xb6, 0xbd, 0xc7, 0xf7, 0xf1, 0x36, 0xf9, 0x12, 0xc4,
		0x08, 0x3e, 0xea, 0xf1, 0xef, 0x32, 0x32, 0x5f, 0xf8, 0xe6, 0x3f, 0x94, 0x83, 0x07, 0x5a, 0x81,
		0x5e, 0xa1, 0x4c, 0x0a, 0x1f, 0xe8, 0xdf, 0x7e, 0x69, 0xef, 0xdb, 0xa1, 0xf6, 0xad, 0x33, 0x63,
		0xd8, 0x86, 0xaf, 0x9f, 0xed, 0xf8, 0x76, 0x95, 0x05, 0xd3, 0xee, 0xf9, 0xd5, 0x01, 0x22, 0x75,
		0xa7, 0xa7, 0x01, 0xdd, 0x7a, 0xb0, 0xcf, 0x4c, 0x6d, 0x0f, 0xa6, 0x9e, 0x21, 0x6d, 0x7b, 0x2b,
		0x68, 0x11, 0xf2, 0xa7, 0xdc, 0x83, 0x3e, 0x89, 0xff, 0xcf, 0x3d, 0x71, 0x88, 0x07, 0x9e, 0x7c,
		0x7c, 0xed, 0x78, 0xef, 0x5c, 0xc7, 0xa9, 0x64, 0xce, 0x37, 0x8d, 0x28, 0x3e, 0x4a, 0xf9, 0xd7,
		0x24, 0xb8, 0xa3, 0xa5, 0x69, 0x1e, 0xe3, 0x17, 0xda, 0xbc, 0x62, 0x38, 0x54, 0xd2, 0xb3, 0xd0,
		0x46, 0xd8, 0xfb, 0x7a, 0x0a, 0xcb, 0xa4, 0x08, 0x48, 0xfb, 0x16, 0x38, 0x12, 0x14, 0x56, 0x98,
		0xe9, 0x1e, 0x18, 0x0d, 0x6e, 0x16, 0x73, 0x73, 0x8d, 0x04, 0xb6, 0x8b, 0xe5, 0x72, 0xd8, 0xce,
		0xae, 0xae, 0x25, 0x48, 0xba, 0xa8, 0x3c, 0x3b, 0xee, 0x5b, 0x55, 0x8f, 0x52, 0xfe, 0x88, 0x04,
		0xb3, 0xc1, 0x16, 0x7c, 0x79, 0xd2, 0xc1, 0x84, 0xbd, 0x65, 0x5d, 0xfc, 0x3d, 0x09, 0xee, 0xec,
		0x22, 0x13, 0x37, 0xc0, 0x35, 0x98, 0xf4, 0x6d, 0x12, 0x88, 0x10, 0x2e, 0xba, 0xfd, 0x54, 0xef,
		0x0c, 0xd5, 0x5d, 0x13, 0x1f, 0x23, 0x46, 0xf9, 0xdc, 0x77, 0x66, 0x26, 0x5a, 0xeb, 0x6c, 0x65,
		0xa2, 0x75, 0x61, 0x7f, 0x0b, 0xfd, 0xe3, 0x15, 0x09, 0xee, 0x0f, 0xaa, 0xda, 0x26, 0xd5, 0xfd,
		0x69, 0xf5, 0xc3, 0xbf, 0x91, 0xe0, 0x54, 0x3f, 0xc2, 0xf1, 0x0e, 0xd9, 0x82, 0x09, 0x2f, 0x09,
		0x0f, 0xf7, 0xc7, 0x81, 0x52, 0x7b, 0xe6, 0xa5, 0xc8, 0xe5, 0x76, 0x1b, 0x0c, 0x6f, 0xf2, 0x81,
		0xe5, 0xef, 0x72, 0xd7, 0xc8, 0xc1, 0x8d, 0x5e, 0x61, 0xe4, 0xc0, 0x56, 0x6f, 0x9b, 0xbe, 0x88,
		0xb4, 0xe9, 0x0b, 0x2f, 0x6b, 0x97, 0xaf, 0xf0, 0xb8, 0xd5, 0x66, 0x7b, 0xee, 0xed, 0x30, 0xd1,
		0xc6, 0x95, 0xf9, 0xa8, 0x3e, 0x80, 0x27, 0x2b, 0xa8, 0xd5, 0x59, 0xe5, 0x7d, 0x98, 0xa1, 0xed,
		0xb6, 0x31, 0xf4, 0xed, 0x56, 0xb9, 0xc1, 0x63, 0x4b, 0xdb, 0xa6, 0xb9, 0xee, 0x8b, 0x30, 0xc8,
		0xfa, 0x99, 0xab, 0x7b, 0x08, 0x47, 0xe1, 0x0c, 0xe4, 0x4f, 0x88, 0x58, 0x56, 0x14, 0x62, 0xb7,
		0x1f, 0x43, 0xfd, 0xe8, 0x7a, 0x8b, 0xc6, 0x90, 0xcf, 0x18, 0xdf, 0x16, 0x51, 0xad, 0xbd, 0x74,
		0xdc, 0x1c, 0x95, 0x5b, 0x16, 0xd5, 0x98, 0x6d, 0x6e, 0x6f, 0xf8, 0xfa, 0x15, 0x11, 0xbe, 0x5c,
		0x9d, 0x7a, 0x84, 0xaf, 0x9f, 0x8e, 0xe9, 0xdd, 0x40, 0xd6, 0x43, 0xcc, 0x3f, 0x8b, 0x81, 0xec,
		0x87, 0x12, 0x1c, 0xa5, 0xba, 0xf9, 0xf7, 0x28, 0x0e, 0x6a, 0xf2, 0x07, 0x01, 0xd9, 0x56, 0xa5,
		0xdc, 0x76, 0x74, 0xa7, 0x6d, 0xab, 0x72, 0x39, 0x30, 0xbf, 0x3c, 0x08, 0xa8, 0x1a, 0xd8, 0x89,
		0xa2, 0xd8, 0xec, 0x02, 0x5d, 0xba, 0xea, 0xdb, 0xe8, 0x68, 0xd3, 0x9d, 0xb1, 0x5b, 0xd0, 0x9d,
		0xdf, 0x92, 0x20, 0xdb, 0x4e, 0x65, 0xde, 0x7d, 0x1a, 0x4c, 0x05, 0xce, 0x0f, 0xc2, 0x3d, 0xf8,
		0x60, 0x3f, 0xbb, 0x3c, 0xa1, 0x61, 0x74, 0xc4, 0xc2, 0xb7, 0x3b, 0x0f, 0x98, 0x09, 0x7a, 0x68,
		0x6b, 0x66, 0xfd, 0x53, 0x1b, 0x3e, 0x5f, 0x6e, 0x89, 0xab, 0x7f, 0x26, 0x72, 0xef, 0x3d, 0x98,
		0xee, 0x20, 0xf5, 0xed, 0x9e, 0xf7, 0x76, 0x3a, 0x76, 0xe6, 0xad, 0x4e, 0xdf, 0x1f, 0xe3, 0x23,
		0x21, 0x78, 0x39, 0xdb, 0xb7, 0x16, 0x6b, 0xf7, 0xba, 0x4b, 0x7e, 0x1b, 0x1c, 0x6b, 0x4b, 0xc5,
		0x65, 0xcb, 0x41, 0x6c, 0x47, 0xb3, 0x1d, 0x2e, 0xd6, 0xbd, 0x9d, 0xc4, 0x0a, 0x51, 0x53, 0x1a,
		0x19, 0x41, 0x9a, 0xb2, 0x5e, 0x33, 0x8c, 0x3a, 0x17, 0x43, 0xbe, 0x04, 0xe3, 0x3e, 0x18, 0x6f,
		0xe4, 0x1c, 0xc4, 0x4c, 0x83, 0x7f, 0xb9, 0x60, 0xf8, 0xcc, 0xf1, 0x8e, 0x1b, 0xfb, 0x86, 0x51,
		0xe7, 0x6a, 0x53, 0x7c, 0x79, 0x12, 0x10, 0x63, 0x46, 0xf7, 0xf8, 0x45, 0x13, 0xeb, 0x30, 0x11,
		0x80, 0xf2, 0x46, 0xde, 0xd0, 0xf9, 0xc1, 0x99, 0xef, 0x1f, 0x81, 0x38, 0xe5, 0x8a, 0x3e, 0x2e,
		0x05, 0x3e, 0x2d, 0x34, 0xd7, 0x89, 0x4d, 0xfb, 0x35, 0x71, 0xf6, 0x74, 0xdf, 0xf8, 0x3c, 0x67,
		0x3b, 0xf5, 0xde, 0x7f, 0xf9, 0xfa, 0x47, 0x23, 0x77, 0x23, 0xf9, 0x74, 0x87, 0xd5, 0xb8, 0x6f,
		0xbc, 0x7c, 0x36, 0xf0, 0x2c, 0xfe, 0xa1, 0xfe, 0x9a, 0x12, 0x92, 0xcd, 0xf5, 0x8b, 0xce, 0x05,
		0x3b, 0x4f, 0x05, 0x3b, 0x8b, 0x1e, 0xed, 0x2d, 0xd8, 0xe9, 0x77, 0x06, 0x07, 0xcd, 0xbb, 0xd1,
		0xef, 0x4a, 0x30, 0xd9, 0x6e, 0x49, 0x87, 0x9e, 0xe8, 0x4f, 0x8a, 0xd6, 0x94, 0x22, 0xfb, 0xe4,
		0x21, 0x28, 0xb9, 0x2a, 0x0b, 0x54, 0x95, 0x3c, 0x7a, 0xea, 0x10, 0xaa, 0x9c, 0xf6, 0x6f, 0xfd,
		0xff, 0x0f, 0x09, 0x4e, 0x74, 0x5d, 0x21, 0xa1, 0x7c, 0x7f, 0x52, 0x76, 0xc9, 0x9d, 0xb2, 0x85,
		0x37, 0xc2, 0x82, 0x6b, 0xfc, 0x0c, 0xd5, 0xf8, 0x12, 0x5a, 0x3c, 0x8c, 0xc6, 0x6d, 0xcf, 0x57,
		0xd0, 0x6f, 0x07, 0x2f, 0x1d, 0x76, 0x77, 0xa7, 0x96, 0x85, 0x47, 0x8f, 0x81, 0xd1, 0x9a, 0xd4,
		0xca, 0xcf, 0x51, 0x15, 0x14, 0xb4, 0xf6, 0x06, 0x3b, 0xed, 0xf4, 0x3b, 0x83, 0x81, 0xff, 0xdd,
		0xe8, 0xbf, 0x4b, 0xed, 0xef, 0x10, 0x3e, 0xde, 0x55, 0xc4, 0xce, 0x8b, 0xaa, 0xec, 0x13, 0x07,
		0x27, 0xe4, 0x4a, 0x36, 0xa8, 0x92, 0x35, 0x84, 0x6f, 0xb5, 0x92, 0x6d, 0x3b, 0x11, 0x7d, 0x43,
		0x82, 0xc9, 0x76, 0x6b, 0x92, 0x1e, 0xc3, 0xb2, 0xcb, 0x22, 0xab, 0xc7, 0xb0, 0xec, 0xb6, 0x00,
		0x92, 0xdf, 0x44, 0x95, 0x3f, 0x87, 0x1e, 0xeb, 0xa4, 0x7c, 0xd7, 0x5e, 0x24, 0x63, 0xb1, 0x6b,
		0x92, 0xdf, 0x63, 0x2c, 0xf6, 0xb3, 0x8e, 0xe9, 0x31, 0x16, 0xfb, 0x5a, 0x63, 0xf4, 0x1e, 0x8b,
		0xae, 0x66, 0x7d, 0x76, 0xa3, 0x8d, 0xbe, 0x26, 0xc1, 0x48, 0x20, 0x23, 0x46, 0x8f, 0x74, 0x15,
		0xb4, 0xdd, 0x82, 0x21, 0x7b, 0xe6, 0x20, 0x24, 0x5c, 0x97, 0x45, 0xaa, 0xcb, 0x3c, 0xca, 0x1f,
		0x46, 0x97, 0xe0, 0x31, 0xea, 0xb7, 0x24, 0x98, 0x68, 0x93, 0x65, 0xf6, 0x18, 0x85, 0x9d, 0x93,
		0xe6, 0xec, 0x13, 0x07, 0x27, 0xe4, 0x5a, 0x5d, 0xa0, 0x5a, 0xbd, 0x15, 0xbd, 0xe5, 0x30, 0x5a,
		0xf9, 0xe6, 0xe7, 0x1b, 0xde, 0x95, 0x2c, 0x5f, 0x3b, 0xe8, 0xdc, 0x01, 0x05, 0x13, 0x0a, 0x3d,
		0x7e, 0x60, 0x3a, 0xae, 0xcf, 0xb3, 0x54, 0x9f, 0x67, 0xd0, 0xea, 0x1b, 0xd3, 0xa7, 0x75, 0x5a,
		0xff, 0x52, 0xeb, 0xe3, 0xc0, 0xee, 0x5e, 0xd4, 0x36, 0x59, 0xcd, 0x3e, 0x7a, 0x20, 0x1a, 0xae,
		0xd4, 0x13, 0x54, 0xa9, 0x33, 0xe8, 0xe1, 0x4e, 0x4a, 0xf9, 0xee, 0xdd, 0x69, 0xfa, 0xb6, 0x71,
		0xfa, 0x9d, 0x2c, 0x05, 0x7e, 0x37, 0x7a, 0x8f, 0xb8, 0xf3, 0x74, 0xb2, 0x6b, 0xbb, 0xbe, 0x3c,
		0x36, 0x7b, 0x7f, 0x1f, 0x98, 0x5c, 0xae, 0xbb, 0xa9, 0x5c, 0xd3, 0xe8, 0x78, 0x27, 0xb9, 0x48,
		0x2e, 0x8b, 0x3e, 0x24, 0xb9, 0xd7, 0x24, 0x4f, 0x75, 0xe7, 0xed, 0x4f, 0x76, 0xb3, 0x0f, 0xf4,
		0x85, 0xcb, 0x25, 0xb9, 0x97, 0x4a, 0x32, 0x8b, 0xa6, 0x3b, 0x4a, 0xc2, 0x52, 0xdf, 0x5b, 0x7d,
		0xa9, 0xe0, 0x7f, 0x4d, 0xc1, 0x4c, 0x87, 0x16, 0x9d, 0xbd, 0x1e, 0x67, 0x5c, 0x5d, 0xde, 0xc8,
		0xf6, 0x7c, 0x03, 0x7b, 0xab, 0xbf, 0xed, 0xda, 0xe7, 0x81, 0xd8, 0xef, 0xc4, 0x00, 0x2d, 0xdb,
		0xb5, 0x79, 0x0b, 0xb3, 0xff, 0x33, 0xc9, 0x47, 0x79, 0xe8, 0xf1, 0x97, 0xf4, 0x86, 0x1e, 0x7f,
		0x2d, 0x07, 0x9e, 0x53, 0x45, 0x0e, 0xf6, 0x64, 0xb3, 0xef, 0x37, 0x55, 0xd1, 0x9f, 0xc8, 0x9b,
		0xaa, 0xf6, 0x57, 0xae, 0x63, 0xb7, 0xee, 0x6d, 0x46, 0xfc, 0xb0, 0xef, 0x53, 0xf8, 0x53, 0xc9,
		0xc1, 0x2e, 0x4f, 0x25, 0x33, 0x1d, 0xdf, 0x43, 0x72, 0x6a, 0x74, 0x56, 0x7c, 0xe9, 0x74, 0xa8,
		0xbf, 0x4b, 0xb2, 0xfc, 0x53, 0xa8, 0xde, 0x16, 0xc2, 0x71, 0xc8, 0xb6, 0xba, 0x93, 0x3b, 0xa8,
		0x3f, 0x1a, 0x85, 0xf4, 0xb2, 0x5d, 0x2b, 0x55, 0x35, 0xe7, 0x36, 0xf9, 0xda, 0x53, 0x9d, 0xdf,
		0xbb, 0xa0, 0x9b, 0x37, 0x66, 0x46, 0x99, 0x4d, 0xbb, 0x58, 0xb2, 0x01, 0x63, 0xa1, 0x57, 0xc6,
		0xdc, 0xb3, 0x8a, 0x87, 0x79, 0xec, 0x1c, 0x62, 0x25, 0xd3, 0xe7, 0x09, 0x3e, 0xff, 0x46, 0x7b,
		0xed, 0x9d, 0x99, 0x39, 0xd4, 0xc5, 0xdb, 0xf9, 0x38, 0xd0, 0xeb, 0xb3, 0x2c, 0x64, 0xc2, 0x9d,
		0xe2, 0xf6, 0xd8, 0x6b, 0x12, 0x0c, 0x2f, 0xdb, 0x22, 0x15, 0xc4, 0x3f, 0xa3, 0x4f, 0x93, 0x1e,
		0x77, 0x3f, 0x13, 0x1e, 0xed, 0xcf, 0x6f, 0xfd, 0x9f, 0x0e, 0x1f, 0x90, 0x8f, 0xc0, 0x84, 0x4f,
		0x47, 0x57, 0xf7, 0x6f, 0x46, 0x68, 0x6c, 0x2c, 0xe0, 0x9a, 0xa6, 0xbb, 0x19, 0x24, 0xfe, 0xf3,
		0xfa, 0xe8, 0xc2, 0xb3, 0x71, 0xec, 0x30, 0x36, 0xde, 0xa5, 0x81, 0x21, 0x64, 0x4b, 0x77, 0xc3,
		0x6b, 0xb9, 0xf5, 0x39, 0x90, 0x74, 0x80, 0x2f, 0xed, 0x84, 0x1e, 0xfd, 0xc8, 0xaf, 0x4b, 0x30,
		0xb2, 0x6c, 0xd7, 0x36, 0xf5, 0xea, 0xff, 0xd5, 0x7e, 0xbb, 0x0d, 0x47, 0x02, 0x5a, 0xde, 0x26,
		0x73, 0x9e, 0x79, 0x25, 0x06, 0xd1, 0x65, 0xbb, 0x86, 0x5e, 0x84, 0xb1, 0x70, 0xa2, 0xd0, 0x31,
		0xff, 0x6b, 0x9d, 0x05, 0x3a, 0xaf, 0xd1, 0x3a, 0xcf, 0x18, 0x68, 0x17, 0x46, 0x82, 0xb3, 0xc5,
		0xc9, 0x2e, 0x4c, 0x02, 0x98, 0xd9, 0x87, 0xfb, 0xc5, 0x74, 0x1b, 0x7b, 0x07, 0x24, 0xdc, 0x40,
		0x77, 0x57, 0x17, 0x6a, 0x81, 0xd4, 0x39, 0xa3, 0x6d, 0x13, 0x4e, 0x88, 0xf5, 0xc2, 0xa1, 0xa4,
		0x9b, 0xf5, 0x42, 0xb8, 0x5d, 0xad, 0xd7, 0x69, 0x58, 0x6d, 0x01, 0xf8, 0xc6, 0xc0, 0x3d, 0x5d,
		0x38, 0x78, 0x68, 0xd9, 0x87, 0xfa, 0x42, 0x73, 0x0f, 0x9a, 0x6e, 0x71, 0x02, 0xfe, 0x7f, 0x02,
		0x00, 0x00, 0xff, 0xff, 0x99, 0xdd, 0x46, 0x27, 0xd1, 0x97, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *CommissionRates) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CommissionRates)
	if !ok {
		that2, ok := that.(CommissionRates)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Rate.Equal(that1.Rate) {
		return false
	}
	if !this.MaxRate.Equal(that1.MaxRate) {
		return false
	}
	if !this.MaxChangeRate.Equal(that1.MaxChangeRate) {
		return false
	}
	return true
}
func (this *Commission) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Commission)
	if !ok {
		that2, ok := that.(Commission)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CommissionRates.Equal(&that1.CommissionRates) {
		return false
	}
	if !this.UpdateTime.Equal(that1.UpdateTime) {
		return false
	}
	return true
}
func (this *Description) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Description)
	if !ok {
		that2, ok := that.(Description)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Identity != that1.Identity {
		return false
	}
	if this.Website != that1.Website {
		return false
	}
	if this.SecurityContact != that1.SecurityContact {
		return false
	}
	if this.Details != that1.Details {
		return false
	}
	return true
}
func (this *UnbondingDelegationEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnbondingDelegationEntry)
	if !ok {
		that2, ok := that.(UnbondingDelegationEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CreationHeight != that1.CreationHeight {
		return false
	}
	if !this.CompletionTime.Equal(that1.CompletionTime) {
		return false
	}
	if !this.InitialBalance.Equal(that1.InitialBalance) {
		return false
	}
	if !this.Balance.Equal(that1.Balance) {
		return false
	}
	return true
}
func (this *RedelegationEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RedelegationEntry)
	if !ok {
		that2, ok := that.(RedelegationEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CreationHeight != that1.CreationHeight {
		return false
	}
	if !this.CompletionTime.Equal(that1.CompletionTime) {
		return false
	}
	if !this.InitialBalance.Equal(that1.InitialBalance) {
		return false
	}
	if !this.SharesDst.Equal(that1.SharesDst) {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UnbondingTime != that1.UnbondingTime {
		return false
	}
	if this.MaxValidators != that1.MaxValidators {
		return false
	}
	if this.MaxEntries != that1.MaxEntries {
		return false
	}
	if this.HistoricalEntries != that1.HistoricalEntries {
		return false
	}
	if this.BondDenom != that1.BondDenom {
		return false
	}
	return true
}
func (this *RedelegationEntryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RedelegationEntryResponse)
	if !ok {
		that2, ok := that.(RedelegationEntryResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RedelegationEntry.Equal(&that1.RedelegationEntry) {
		return false
	}
	if !this.Balance.Equal(that1.Balance) {
		return false
	}
	return true
}
func (this *Pool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pool)
	if !ok {
		that2, ok := that.(Pool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.NotBondedTokens.Equal(that1.NotBondedTokens) {
		return false
	}
	if !this.BondedTokens.Equal(that1.BondedTokens) {
		return false
	}
	return true
}
func (m *HistoricalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoricalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HistoricalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Valset) > 0 {
		for iNdEx := len(m.Valset) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valset[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CommissionRates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommissionRates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommissionRates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxChangeRate.Size()
		i -= size
		if _, err := m.MaxChangeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxRate.Size()
		i -= size
		if _, err := m.MaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Commission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintStaking(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommissionRates.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Description) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Description) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Description) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinSelfDelegation.Size()
		i -= size
		if _, err := m.MinSelfDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size, err := m.Commission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UnbondingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UnbondingTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintStaking(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x4a
	if m.UnbondingHeight != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.UnbondingHeight))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.DelegatorShares.Size()
		i -= size
		if _, err := m.DelegatorShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Tokens.Size()
		i -= size
		if _, err := m.Tokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Status != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ConsensusPubkey != nil {
		{
			size, err := m.ConsensusPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStaking(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValAddresses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValAddresses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValAddresses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintStaking(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DVPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DVPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DVPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DVPairs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DVPairs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DVPairs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pairs) > 0 {
		for iNdEx := len(m.Pairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DVVTriplet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DVVTriplet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DVVTriplet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorDstAddress) > 0 {
		i -= len(m.ValidatorDstAddress)
		copy(dAtA[i:], m.ValidatorDstAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorDstAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorSrcAddress) > 0 {
		i -= len(m.ValidatorSrcAddress)
		copy(dAtA[i:], m.ValidatorSrcAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorSrcAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DVVTriplets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DVVTriplets) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DVVTriplets) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Triplets) > 0 {
		for iNdEx := len(m.Triplets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Triplets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnbondingDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnbondingDelegationEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingDelegationEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingDelegationEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InitialBalance.Size()
		i -= size
		if _, err := m.InitialBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n8, err8 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVarintStaking(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0x12
	if m.CreationHeight != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RedelegationEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedelegationEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedelegationEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesDst.Size()
		i -= size
		if _, err := m.SharesDst.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InitialBalance.Size()
		i -= size
		if _, err := m.InitialBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n9, err9 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err9 != nil {
		return 0, err9
	}
	i -= n9
	i = encodeVarintStaking(dAtA, i, uint64(n9))
	i--
	dAtA[i] = 0x12
	if m.CreationHeight != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Redelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Redelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Redelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ValidatorDstAddress) > 0 {
		i -= len(m.ValidatorDstAddress)
		copy(dAtA[i:], m.ValidatorDstAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorDstAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorSrcAddress) > 0 {
		i -= len(m.ValidatorSrcAddress)
		copy(dAtA[i:], m.ValidatorSrcAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValidatorSrcAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BondDenom) > 0 {
		i -= len(m.BondDenom)
		copy(dAtA[i:], m.BondDenom)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.BondDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.HistoricalEntries != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.HistoricalEntries))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxEntries != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.MaxEntries))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxValidators != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.MaxValidators))
		i--
		dAtA[i] = 0x10
	}
	n10, err10 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.UnbondingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingTime):])
	if err10 != nil {
		return 0, err10
	}
	i -= n10
	i = encodeVarintStaking(dAtA, i, uint64(n10))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DelegationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Delegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RedelegationEntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedelegationEntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedelegationEntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.RedelegationEntry.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RedelegationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedelegationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedelegationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Redelegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BondedTokens.Size()
		i -= size
		if _, err := m.BondedTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.NotBondedTokens.Size()
		i -= size
		if _, err := m.NotBondedTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HistoricalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Header.Size()
	n += 1 + l + sovStaking(uint64(l))
	if len(m.Valset) > 0 {
		for _, e := range m.Valset {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *CommissionRates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Rate.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.MaxRate.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.MaxChangeRate.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *Commission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommissionRates.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime)
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *Description) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.ConsensusPubkey != nil {
		l = m.ConsensusPubkey.Size()
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovStaking(uint64(m.Status))
	}
	l = m.Tokens.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.DelegatorShares.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.Description.Size()
	n += 1 + l + sovStaking(uint64(l))
	if m.UnbondingHeight != 0 {
		n += 1 + sovStaking(uint64(m.UnbondingHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UnbondingTime)
	n += 1 + l + sovStaking(uint64(l))
	l = m.Commission.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.MinSelfDelegation.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *ValAddresses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *DVPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *DVPairs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pairs) > 0 {
		for _, e := range m.Pairs {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *DVVTriplet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorSrcAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorDstAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *DVVTriplets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Triplets) > 0 {
		for _, e := range m.Triplets {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = m.Shares.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *UnbondingDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *UnbondingDelegationEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovStaking(uint64(m.CreationHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovStaking(uint64(l))
	l = m.InitialBalance.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *RedelegationEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovStaking(uint64(m.CreationHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovStaking(uint64(l))
	l = m.InitialBalance.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.SharesDst.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *Redelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorSrcAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.ValidatorDstAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingTime)
	n += 1 + l + sovStaking(uint64(l))
	if m.MaxValidators != 0 {
		n += 1 + sovStaking(uint64(m.MaxValidators))
	}
	if m.MaxEntries != 0 {
		n += 1 + sovStaking(uint64(m.MaxEntries))
	}
	if m.HistoricalEntries != 0 {
		n += 1 + sovStaking(uint64(m.HistoricalEntries))
	}
	l = len(m.BondDenom)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *DelegationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Delegation.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *RedelegationEntryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RedelegationEntry.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func (m *RedelegationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Redelegation.Size()
	n += 1 + l + sovStaking(uint64(l))
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NotBondedTokens.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = m.BondedTokens.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ValAddresses) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ValAddresses{`,
		`Addresses:` + fmt.Sprintf("%v", this.Addresses) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStaking(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HistoricalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistoricalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoricalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Valset = append(m.Valset, Validator{})
			if err := m.Valset[len(m.Valset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommissionRates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommissionRates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommissionRates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxChangeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxChangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Commission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Commission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommissionRates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Description) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Description: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Description: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusPubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusPubkey == nil {
				m.ConsensusPubkey = &types1.Any{}
			}
			if err := m.ConsensusPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BondStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegatorShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingHeight", wireType)
			}
			m.UnbondingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UnbondingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValAddresses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValAddresses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValAddresses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DVPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DVPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DVPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DVPairs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DVPairs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DVPairs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pairs = append(m.Pairs, DVPair{})
			if err := m.Pairs[len(m.Pairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DVVTriplet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DVVTriplet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DVVTriplet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSrcAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSrcAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorDstAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorDstAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DVVTriplets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DVVTriplets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DVVTriplets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Triplets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Triplets = append(m.Triplets, DVVTriplet{})
			if err := m.Triplets[len(m.Triplets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnbondingDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnbondingDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, UnbondingDelegationEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnbondingDelegationEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnbondingDelegationEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingDelegationEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RedelegationEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RedelegationEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedelegationEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesDst", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesDst.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Redelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Redelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Redelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSrcAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorSrcAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorDstAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorDstAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, RedelegationEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.UnbondingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValidators", wireType)
			}
			m.MaxValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxValidators |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEntries", wireType)
			}
			m.MaxEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalEntries", wireType)
			}
			m.HistoricalEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoricalEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DelegationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DelegationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Delegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RedelegationEntryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RedelegationEntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedelegationEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedelegationEntry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedelegationEntry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RedelegationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RedelegationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redelegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Redelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, RedelegationEntryResponse{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotBondedTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NotBondedTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondedTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondedTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
