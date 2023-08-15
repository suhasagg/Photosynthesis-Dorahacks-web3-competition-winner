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

	return