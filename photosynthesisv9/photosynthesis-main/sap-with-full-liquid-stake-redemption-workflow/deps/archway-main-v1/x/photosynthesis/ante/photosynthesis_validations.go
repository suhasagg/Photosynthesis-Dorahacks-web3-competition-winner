package ante

import (
	"errors"
	rewardsTypes "github.com/archway-network/archway/x/rewards/types"
)

func ValidatePhotosynthesisContract(contract *rewardsTypes.ContractMetadata) error {
	if contract.MinimumRewardAmount == 0 {
		return errors.New("Contract has not enabled liquid staking or has not set a minimum reward amount to be liquid staked")
	}

	if contract.LiquidStakeInterval <= 0 {
		return errors.New("Invalid liquid stake interval")
	}

	if contract.RedemptionInterval <= 0 {
		return errors.New("Invalid redemption interval")
	}

	if contract.RewardsWithdrawalInterval <= 0 {
		return errors.New("Invalid rewards withdrawal interval")
	}

	if contract.RedemptionRateThreshold <= 0 {
		return errors.New("Invalid redemption rate threshold")
	}

	if contract.RedemptionIntervalThreshold <= 0 {
		return errors.New("Invalid redemption interval threshold")
	}

	if contract.MaximumThreshold == 0 {
		return errors.New("Invalid maximum threshold for staking rewards")
	}

	if contract.LiquidityProviderCommission <= 0 {
		return errors.New("Invalid commission rate for liquidity providers")
	}

	if contract.AirdropDuration <= 0 {
		return errors.New("Invalid airdrop duration")
	}

	if contract.AirdropVestingPeriod <= 0 {
		return errors.New("Invalid airdrop vesting period")
	}

	// Call next ante handler in decorator chain
	return nil
}
