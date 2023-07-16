package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/archway-network/archway/x/rewards/types"
)

// RegisterInvariants registers all module invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "module-account-balance", ModuleAccountBalanceInvariant(k))
}

// invariant to check that total liquid staked amount is less than or equal to total rewards amount
func TotalLiquidStakedInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// get total rewards amount
		totalRewards := k.GetTotalRewardsAmount(ctx)

		// get total liquid staked amount
		totalLiquidStaked := k.GetTotalLiquidStakedAmount(ctx)

		// check if total liquid staked is less than or equal to total rewards
		if totalLiquidStaked.GT(totalRewards) {
			return sdk.FormatInvariant(types.ModuleName, "total liquid staked should be less than or equal to total rewards"), false
		}

		return "", true
	}
}

// invariant to check that the sum of all Dapp liquid stake amounts is less than or equal to their total rewards amount
func DappLiquidStakeInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		ok := true

		// iterate over all Dapps
		k.IterateDapps(ctx, func(dapp types.Dapp) bool {
			// get the total rewards amount for this Dapp
			totalRewards := dapp.Rewards.Amount

			// get the liquid stake amount for this Dapp
			liquidStake := k.GetDappLiquidStakeAmount(ctx, dapp.Address)

			// check if liquid stake is greater than total rewards
			if liquidStake.GT(totalRewards) {
				msg += fmt.Sprintf("Dapp [%s] liquid stake should be less than or equal to its total rewards amount\n", dapp.Address)
				ok = false
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, msg), ok
	}
}

// invariant to check that the total commission charged by Archway is less than or equal to the total liquid staking commission received
func CommissionInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// get total liquid staking commission received
		totalLiquidityCommission := k.GetTotalLiquidityCommission(ctx)

		// get total commission charged by Archway
		totalArchwayCommission := k.GetTotalArchwayCommission(ctx)

		// check if total commission charged is less than or equal to total commission received
		if totalArchwayCommission.GT(totalLiquidityCommission) {
			return sdk.FormatInvariant(types.ModuleName, "total commission charged by Archway should be less than or equal to total liquidity commission received"), false
		}

		return "", true
	}
}

func RewardsWithdrawalInvariant(k Keeper, ctx sdk.Context) (string, bool) {
	epoch := k.GetCurrentEpoch(ctx)
	totalRewards := k.GetRewardsForWithdrawal(ctx, epoch)

	for _, contractAddr := range k.GetAllContractAddresses(ctx) {
		withdrawalAmount := k.WithdrawRewards(ctx, contractAddr, epoch)

		if withdrawalAmount.GT(totalRewards) {
			return fmt.Sprintf("withdrawn rewards (%v) for contract address %v in epoch %v exceeds total rewards available (%v)", withdrawalAmount, contractAddr, epoch, totalRewards), false
		}

		totalRewards = totalRewards.Sub(withdrawalAmount)
	}

	return "", true
}

func LiquidStakeDepositInvariant(k Keeper, ctx sdk.Context) (string, bool) {
	epoch := k.GetCurrentEpoch(ctx)
	totalRewards := k.GetRewardsForLiquidStaking(ctx, epoch)

	var depositedRewards sdk.Int
	records := k.GetLiquidStakeQueueForEpoch(ctx, epoch)
	for _, record := range records {
		depositedRewards = depositedRewards.Add(record.RewardAmount)
	}

	if depositedRewards.GT(totalRewards) {
		return fmt.Sprintf("deposited rewards (%v) for epoch %v in liquid stake queue exceeds total rewards available (%v)", depositedRewards, epoch, totalRewards), false
	}

	return "", true
}

// Invariant checks that contract reward amount is greater than or equal to the minimum threshold for liquid staking
func ContractRewardAmountAboveThresholdInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var invarBroken bool

		// iterate over all the contracts
		k.IterateContracts(ctx, func(contract types.Contract) bool {
			rewardAmount, err := k.GetContractRewardAmount(ctx, contract.Address)
			if err != nil {
				msg = err.Error()
				invarBroken = true
				return true
			}

			// check if reward amount is greater than or equal to the threshold for liquid staking
			if rewardAmount.LT(contract.MinimumRewardThreshold) {
				msg = fmt.Sprintf("contract reward amount (%v) is below the minimum threshold (%v) for liquid staking for contract address %s", rewardAmount, contract.MinimumRewardThreshold, contract.Address.String())
				invarBroken = true
				return true
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "contract reward amount above threshold", msg), invarBroken
	}
}

// Invariant checks that there is no negative balance in Archway account
func ArchwayAccountHasPositiveBalanceInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var invarBroken bool

		// get the Archway account
		archwayAccount := k.GetArchwayAccount(ctx)

		// check if balance is negative
		if archwayAccount.Balance.IsNegative() {
			msg = fmt.Sprintf("archway account has negative balance: %v", archwayAccount.Balance)
			invarBroken = true
			return sdk.FormatInvariant(types.ModuleName, "archway account has positive balance", msg), invarBroken
		}

		return sdk.FormatInvariant(types.ModuleName, "archway account has positive balance", msg), invarBroken
	}
}

// Invariant checks that the total amount of Archway rewards transferred to liquidity providers
// equals the sum of the total amount of liquid tokens minted for each Dapp
func TotalArchwayRewardsEqualsTotalLiquidityTokens(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var totalRewards sdk.Int
		var totalLiquidityTokens sdk.Int

		// Iterate over all LiquidStakeDepositRecords
		k.IterateLiquidStakeDepositRecords(ctx, func(deposit types.LiquidStakeDepositRecord) bool {
			totalRewards = totalRewards.Add(deposit.RewardAmount)
			totalLiquidityTokens = totalLiquidityTokens.Add(deposit.LiquidityTokenAmount)
			return false
		})

		// Iterate over all WithdrawalInfo
		withdrawalInfo := k.GetAllWithdrawalInfo(ctx)
		for _, info := range withdrawalInfo {
			totalLiquidityTokens = totalLiquidityTokens.Add(info.LiquidityTokens)
		}

		if !totalRewards.Equal(totalLiquidityTokens) {
			return sdk.FormatInvariant(types.ModuleName, "total archway rewards equal total liquidity tokens"),
				fmt.Sprintf("total archway rewards (%s) != total liquidity tokens (%s)", totalRewards.String(), totalLiquidityTokens.String()), false
		}

		return sdk.FormatInvariant(types.ModuleName, "total archway rewards equal total liquidity tokens"), "", true
	}
}

// Invariant checks that the total amount of liquidity tokens distributed to Dapps is
// equal to the total amount of liquidity tokens minted by liquidity providers
func TotalLiquidityTokensDistributedEqualsTotalMinted(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var totalDistributed sdk.Int
		var totalMinted sdk.Int

		// Iterate over all WithdrawalInfo
		withdrawalInfo := k.GetAllWithdrawalInfo(ctx)
		for _, info := range withdrawalInfo {
			totalDistributed = totalDistributed.Add(info.LiquidityTokens)
		}

		// Iterate over all LiquidityProviderInfo
		lpInfo := k.GetAllLiquidityProviderInfo(ctx)
		for _, info := range lpInfo {
			totalMinted = totalMinted.Add(info.LiquidityTokensMinted)
		}

		if !totalDistributed.Equal(totalMinted) {
			return sdk.FormatInvariant(types.ModuleName, "total liquidity tokens distributed equal total liquidity tokens minted"),
				fmt.Sprintf("total liquidity tokens distributed (%s) != total liquidity tokens minted (%s)", totalDistributed.String(), totalMinted.String()), false
		}

		return sdk.FormatInvariant(types.ModuleName, "total liquidity tokens distributed equal total liquidity tokens minted"), "", true
	}
}

func LiquidityCommissionInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		totalCommission := sdk.ZeroInt()
		dappIterator := k.GetDappIterator(ctx)
		for ; dappIterator.Valid(); dappIterator.Next() {
			dapp := k.MustGetDapp(ctx, dappIterator.Key())
			totalCommission = totalCommission.Add(dapp.LiquidityCommission)
		}

		archwayCommission := k.GetArchwayCommission(ctx)
		if totalCommission.Equal(archwayCommission) {
			return "", true
		}

		return fmt.Sprintf("total liquidity commission: %s does not match sum of individual dapp liquidity commission: %s", archwayCommission.String(), totalCommission.String()), false
	}
}

func RedemptionRateInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		params := k.GetParams(ctx)
		maxRedemptionRate, err := k.GetMaximumRedemptionRate(ctx)
		if err != nil {
			return fmt.Sprintf("error getting maximum redemption rate: %s", err.Error()), false
		}
		if maxRedemptionRate.LT(params.MinimumRedemptionRate) || maxRedemptionRate.GT(params.MaximumRedemptionRate) {
			return fmt.Sprintf("maximum redemption rate (%s) not within range of minimum (%s) and maximum (%s) redemption rates", maxRedemptionRate.String(), params.MinimumRedemptionRate.String(), params.MaximumRedemptionRate.String()), false
		}

		return "", true
	}
}

func (k Keeper) CheckRedemptionRateInvariant(ctx sdk.Context) (string, bool) {
	var (
		totalLiquidTokens sdk.Int
		totalRedemption   sdk.Int
	)

	// Iterate over all DApps and sum up their liquid token balance and redemption amount
	k.IterateDApps(ctx, func(dApp types.DApp) bool {
		// Add the DApp's liquid token balance to the total
		totalLiquidTokens = totalLiquidTokens.Add(dApp.LiquidTokenBalance.Amount)

		// Add the DApp's redemption amount to the total
		totalRedemption = totalRedemption.Add(dApp.RedemptionAmount.Amount)

		return false
	})

	// Check that the total redemption amount is less than or equal to the total liquid token balance
	if totalRedemption.GT(totalLiquidTokens) {
		return fmt.Sprintf("total redemption amount (%s) is greater than total liquid token balance (%s)", totalRedemption.String(), totalLiquidTokens.String()), true
	}

	return "", true
}

func (k Keeper) CheckLiquidTokenBalanceInvariant(ctx sdk.Context) (string, bool) {
	// Iterate over all DApps and check their liquid token balance
	k.IterateDApps(ctx, func(dApp types.DApp) bool {
		// Check that the DApp's liquid token balance is greater than or equal to zero
		if dApp.LiquidTokenBalance.Amount.LT(sdk.ZeroInt()) {
			return true
		}

		return false
	})

	return "", true
}
