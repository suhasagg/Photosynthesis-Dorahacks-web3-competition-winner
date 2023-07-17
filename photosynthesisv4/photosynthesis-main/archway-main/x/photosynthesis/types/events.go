package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
// LiquidStakeDepositRecordCreatedEvent is emitted when a liquid staking deposit record is created.
type LiquidStakeDepositRecordCreatedEvent struct {
	RecordID      string `json:"record_id"`
	RewardsAmount int64  `json:"rewards_amount"`
}

// RedemptionRateUpdatedEvent is emitted when the redemption rate threshold is updated.
type RedemptionRateUpdatedEvent struct {
	ContractAddress string `json:"contract_address"`
	NewThreshold    int64  `json:"new_threshold"`
}

// RewardsDistributedEvent is emitted when rewards are distributed to contracts.
type RewardsDistributedEvent struct {
	RewardAddress string `json:"reward_address"`
	RewardsAmount int64  `json:"rewards_amount"`
	NumContracts  int    `json:"num_contracts"`
}

*/

// EmitLiquidStakeDepositRecordCreatedEvent emits a LiquidStakeDepositRecordCreatedEvent.
func EmitLiquidStakeDepositRecordCreatedEvent(ctx sdk.Context, recordID string, rewardsAmount int64) {
	event := LiquidStakeDepositRecordCreatedEvent{
		RecordId:      recordID,
		RewardsAmount: rewardsAmount,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		panic(fmt.Errorf("failed to emit LiquidStakeDepositRecordCreatedEvent: %w", err))
	}
}

// EmitRedemptionRateUpdatedEvent emits a RedemptionRateUpdatedEvent.
func EmitRedemptionRateUpdatedEvent(ctx sdk.Context, contractAddress string, newThreshold int64) {
	event := RedemptionRateUpdatedEvent{
		ContractAddress: contractAddress,
		NewThreshold:    newThreshold,
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		panic(fmt.Errorf("failed to emit RedemptionRateUpdatedEvent: %w", err))
	}
}

// EmitRewardsDistributedEvent emits a RewardsDistributedEvent.
func EmitRewardsDistributedEvent(ctx sdk.Context, rewardAddress string, rewardsAmount int64, numContracts int) {
	event := RewardsDistributedEvent{
		RewardAddress: rewardAddress,
		RewardsAmount: rewardsAmount,
		NumContracts:  int32(numContracts),
	}
	if err := ctx.EventManager().EmitTypedEvent(&event); err != nil {
		panic(fmt.Errorf("failed to emit RewardsDistributedEvent: %w", err))
	}
}

/*
func EmitRewardsWithdrawEvent(ctx sdk.Context, rewardAddress sdk.AccAddress, rewards int64) {
	err := ctx.EventManager().EmitTypedEvent(&RewardsWithdrawEvent{
		RewardAddress: rewardAddress.String(),
		Rewards:       rewards,
	})
	if err != nil {
		panic(fmt.Errorf("sending RewardsWithdrawEvent event: %w", err))
	}
}
*/
