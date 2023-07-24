package keeper

import (
	"fmt"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
	"github.com/archway-network/archway/x/rewards/types"
	trackingTypes "github.com/archway-network/archway/x/tracking/types"
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
	GetLatestRedemptionRecord(ctx sdk.Context) (types.RedemptionRecord, bool)

	// Get the cumulative liquidity amount
	GetCumulativeLiquidityAmount(ctx sdk.Context) (sdk.Coins, error)

	// Get the total stake of all contracts
	GetTotalStake(ctx sdk.Context) (sdk.Int, error)

	// Get the stake of the given contract address
	GetStake(ctx sdk.Context, contractAddress sdk.AccAddress) (sdk.Int, error)

	// List all contracts
	ListContracts(ctx sdk.Context) ([]types.Contract, error)

	// Send tokens to the given contract address
	SendTokensToContract(ctx sdk.Context, contractAddress sdk.AccAddress, tokens sdk.Int) error

	// Redeem liquid tokens for the given amount
	RedeemLiquidTokens(ctx sdk.Context, amount sdk.Coins) error

	// Distribute redeemed tokens to the Dapps according to their stake
	DistributeRedeemedTokens(ctx sdk.Context, redeemedTokensAmount sdk.Coins) error

	// Delete the latest redemption record
	DeleteRedemptionRecord(ctx sdk.Context, record types.RedemptionRecord) error
}

// Keeper provides module state operations.
type Keeper struct {
	cdc              codec.Codec
	paramStore       paramTypes.Subspace
	state            State
	contractInfoView ContractInfoReaderExpected
	photosynthesis   Photosynthesis
	trackingKeeper   TrackingKeeperExpected
	authKeeper       AuthKeeperExpected
	bankKeeper       BankKeeperExpected
}

// ProcessRedemptionRateQueries queries the redemption rate at specific epochs and checks
// if the redemption rate is above a threshold. If the rate is above the threshold, it
// initiates the redemption process and distributes the redeemed tokens to Dapps according
// to their stake. It also deletes the latest redemption record.
func (k Keeper) ProcessRedemptionRateQueries(ctx sdk.Context, epochInfo types.EpochInfo) error {
	if epochInfo.Identifier != types.REDEMPTION_RATE_QUERY_EPOCH {
		return nil
	}

	redemptionRateQueryInterval := k.GetRedemptionRateQueryInterval(ctx)

	if epochInfo.Number%redemptionRateQueryInterval != 0 {
		return nil
	}

	err := k.QueryRedemptionRate(ctx, epochInfo.Number)
	if err != nil {
		return err
	}

	redemptionRateThreshold := k.GetRedemptionRateThreshold(ctx)
	if redemptionRate.Compare(redemptionRateThreshold) > 0 {
		redemptionIntervalThreshold := k.GetRedemptionIntervalThreshold(ctx)
		timeSinceLatestRedemption := ctx.BlockTime().Sub(k.GetLatestRedemptionTime(ctx))

		if timeSinceLatestRedemption >= redemptionIntervalThreshold {
			cumLiquidityAmount, _ := k.GetCumulativeLiquidityAmount(ctx)
			err = k.RedeemLiquidTokens(ctx, cumLiquidityAmount)
			if err != nil {
				return err
			}

			err = k.DistributeRedeemedTokens(ctx, cumLiquidityAmount)
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
func (k Keeper) DistributeRedeemedTokens(ctx sdk.Context, redeemedTokensAmount sdk.Coins) error {
	totalStake, err := k.GetTotalStake(ctx)
	if err != nil {
		return err
	}

	contracts, err := k.ListContracts(ctx)
	if err != nil {
		return err
	}

	for _, contract := range contracts {
		stake, err := k.GetStake(ctx, contract.Address)
		if err != nil {
			return err
		}

		tokensToDistribute := redeemedTokensAmount.Mul(stake).Quo(totalStake)
		if tokensToDistribute.IsZero() {
			continue
		}

		err = k.SendTokensToContract(ctx, contract.Address, tokensToDistribute)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteRedemptionRecord deletes the latest redemption record from the store.
func (k Keeper) DeleteRedemptionRecord(ctx sdk.Context) error {
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
func (k Keeper) UpdateContract(ctx sdk.Context, contract types.Contract) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetContractKey(contract.Address)
	value := k.cdc.MustMarshalBinaryBare(&contract)
	store.Set(key, value)
	return nil
}

// DeleteContract deletes the contract information from the store.
func (k Keeper) DeleteContract(ctx sdk.Context, address sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetContractKey(address)
	store.Delete(key)
	return nil
}

// GetStake retrieves the stake of a contract.
func (k Keeper) GetStake(ctx sdk.Context, address sdk.AccAddress) (sdk.Int, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(address)
	value := store.Get(key)
	if value == nil {
		return sdk.Int{}, fmt.Errorf("stake not found for address %s", address.String())
	}
	var stake sdk.Int
	err := k.cdc.UnmarshalBinaryBare(value, &stake)
	if err != nil {
		return sdk.Int{}, fmt.Errorf("failed to unmarshal stake: %s", err)
	}
	return stake, nil
}

// SetStake sets the stake of a contract.
func (k Keeper) SetStake(ctx sdk.Context, address sdk.AccAddress, stake sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetStakeKey(address)
	value := k.cdc.MustMarshalBinaryBare(&stake)
	store.Set(key, value)
	return nil
}


// GetTotalStake calculates the total stake across all contracts
func (k Keeper) GetTotalStake(ctx sdk.Context) (sdk.Int, error) {
	var totalStake sdk.Int
	contracts, err := k.ListContracts(ctx)
	if err != nil {
		return sdk.Int{}, err
	}
	for _, contract := range contracts {
		stake, err := k.GetStake(ctx, contract.Address)
		if err != nil {
			return sdk.Int{}, err
		}
		totalStake = totalStake.Add(stake)
	}
	return totalStake, nil
}

// SendTokensToContract sends tokens to a contract address
func (k Keeper) SendTokensToContract(ctx sdk.Context, address sdk.AccAddress, amount sdk.Int) error {
	err := k.bankKeeper.SendCoins(ctx, k.contractOwner, address, sdk.NewCoins(sdk.NewCoin(k.rewardTokenSymbol, amount)))
	if err != nil {
		return fmt.Errorf("failed to send tokens to contract: %s", err)
	}
	return nil
}

// GetRedemptionRateThreshold returns the redemption rate threshold
func (k Keeper) GetRedemptionRateThreshold(ctx sdk.Context) sdk.Dec {
	return k.params.RedemptionRateThreshold
}

// GetRedemptionIntervalThreshold returns the redemption interval threshold
func (k Keeper) GetRedemptionIntervalThreshold(ctx sdk.Context) time.Duration {
	return k.params.RedemptionIntervalThreshold
}

// GetLatestRedemptionRecord gets the latest redemption record
func (k Keeper) GetLatestRedemptionRecord(ctx sdk.Context) (types.RedemptionRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStoreReversePrefixIterator(store, types.RedemptionRecordPrefix)
	defer iter.Close()
	if !iter.Valid() {
		return types.RedemptionRecord{}, false
	}
	var record types.RedemptionRecord
	k.cdc.MustUnmarshalBinaryBare(iter.Value(), &record)
	return record, true
}

// SetLatestRedemptionRecord sets the latest redemption record
func (k Keeper) SetLatestRedemptionRecord(ctx sdk.Context, record types.RedemptionRecord) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRedemptionRecordKey(record.Timestamp)
	value := k.cdc.MustMarshalBinaryBare(&record)
	store.Set(key, value)
	return nil
}

// GetCumulativeLiquidityAmount gets the cumulative liquidity amount
func (k Keeper) GetCumulativeLiquidityAmount(ctx sdk.Context) (sdk.Coins, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.CumulativeLiquidityAmountKey
	bz := store.Get(key)
	if bz == nil {
		return nil, fmt.Errorf("cumulative liquidity amount not found")
	}
	var coins sdk.Coins
	err := k.cdc.UnmarshalBinaryBare(bz, &coins)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cumulative liquidity amount: %s", err)
	}
	return coins, nil
}

// SetCumulativeLiquidityAmount sets the cumulative liquidity amount
func (k Keeper) SetCumulativeLiquidityAmount(ctx sdk.Context, amount sdk.Coins) error {
	store := ctx.KVStore(k.storeKey)
	key := types.CumulativeLiquidityAmountKey
	bz, err := k.cdc.MarshalBinaryBare(amount)
	if err != nil {
		return fmt.Errorf("failed to marshal cumulative liquidity amount: %s", err)
	}
	store.Set(key, bz)
	return nil
}
// RecordRewards records rewards for a contract by adding them to the existing rewards
func (k Keeper) RecordRewards(ctx sdk.Context, contractAddress sdk.AccAddress, amount sdk.Coins) error {
	rewards := k.rewardsKeeper.GetRewards(ctx, contractAddress) // retrieve existing rewards for the contract
	newRewards := rewards.Add(amount...) // add the new rewards to the existing rewards
	k.rewardsKeeper.SetRewards(ctx, contractAddress, newRewards) // set the new rewards for the contract
	return nil
}

// GetRewards returns the rewards for a contract
func (k Keeper) GetRewards(ctx sdk.Context, contractAddress sdk.AccAddress) sdk.Coins {
	return k.rewardsKeeper.GetRewards(ctx, contractAddress) // retrieve the rewards for the contract
}

// DistributeRewards distributes rewards to all contracts that have opted for liquid staking
func (k Keeper) DistributeRewards(ctx sdk.Context, epochNumber uint64, cumulativeRewards sdk.Coins) error {
	params := k.GetParams(ctx)
	for _, dapp := range k.recordsKeeper.GetContractsToLiquidStake(ctx) { // iterate over all contracts that have opted for liquid staking
		rewardAmount := k.GetRewards(ctx, dapp.ContractAddress) // retrieve the rewards for the contract
		if rewardAmount.IsZero() {
			continue
		}

		if rewardAmount.AmountOf(params.RewardDenom).LT(params.MinRewardsToLiquidStake) { // check if the rewards are greater than the minimum required to distribute
			continue
		}

		if epochNumber%dapp.LiquidStakeInterval != 0 { // check if the current epoch is divisible by the contract's liquid stake interval
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


//This function is called by DistributeRewards to actually distribute rewards to a single Dapp. It takes the Dapp object, the amount of rewards to distribute, and the cumulative rewards as inputs. It first checks if the rewards are greater than the minimum required to distribute. If so, it calculates the ratio of rewards to distribute based on the total stake of the Dapp and distributes rewards proportionally to each stakeholder. It then mints the rewards tokens and sends them to the Dapp, updates the cumulative rewards, and emits an event.
func (k Keeper) DistributeRewardsToDapp(ctx sdk.Context, dapp types.Contract, rewards sdk.Coins, cumulativeRewards sdk.Coins) error {
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

//This function returns the current balance of liquidity tokens for a given address.
func (k Keeper) LiquidityTokenBalance(ctx sdk.Context, senderAddr sdk.AccAddress) (sdk.Coins, error) {
	balance := k.bankKeeper.GetCoins(ctx, senderAddr)
	return balance, nil
}

//This function returns the liquid staking deposit object for a given sender and contract address.
func (k Keeper) LiquidStakingDeposit(ctx sdk.Context, senderAddr sdk.AccAddress, contractAddr sdk.AccAddress) (*types.LiquidStakingDeposit, error) {
	deposit, found := k.GetLiquidStakingDeposit(ctx, senderAddr, contractAddr)
	if !found {
		return nil, types.ErrDepositNotFound
	}
	return &deposit, nil
}

func (k Keeper) RedemptionRate(ctx sdk.Context) (sdk.Dec, error) {
	return k.GetRedemptionRate(ctx), nil
}

//This function returns the airdrop status for a given address, including the total amount of the airdrop, the amount currently vested, and the amount remaining to be vested.
func (k Keeper) AirdropStatus(ctx sdk.Context, senderAddr sdk.AccAddress) (*types.AirdropStatus, error) {
	status, found := k.GetAirdropStatus(ctx, senderAddr)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAirdropNotFound, "address %s", senderAddr.String())
	}

	elapsedDuration := ctx.BlockTime().Sub(status.StartTime)
	elapsedMonths := elapsedDuration / (30 * 24 * time.Hour) // assuming 30-day months
	vestingAmount := sdk.NewCoin(status.Amount.Denom, status.Amount.Amount.MulRaw(int64(100 - elapsedMonths*10)).QuoRaw(100))
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

//This function is similar to AirdropStatus, but takes an address as input instead of using the caller's address.
func (k Keeper) QueryAirdropStatus(ctx sdk.Context, addr sdk.AccAddress) (*types.AirdropStatus, error) {
	status, found := k.GetAirdropStatus(ctx, addr)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAirdropNotFound, "address %s", addr.String())
	}

	elapsedDuration := ctx.BlockTime().Sub(status.StartTime)
	elapsedMonths := elapsedDuration / (30 * 24 * time.Hour) // assuming 30-day months
	vestingAmount := sdk.NewCoin(status.Amount.Denom, status.Amount.Amount.MulRaw(int64(100 - elapsedMonths*10)).QuoRaw(100))
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
