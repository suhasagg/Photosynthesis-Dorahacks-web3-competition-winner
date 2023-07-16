package keeper

import (
	"github.com/archway-network/archway/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetContractMetadata creates or updates the contract metadata verifying the ownership:
//   - Meta could be created by the contract admin (if set);
//   - Meta could be modified by the contract owner;
func (k Keeper) SetContractMetadata(ctx sdk.Context, senderAddr, contractAddr sdk.AccAddress, metaUpdates types.ContractMetadata) error {
	state := k.state.ContractMetadataState(ctx)

	// Check if the contract exists
	contractInfo := k.contractInfoView.GetContractInfo(ctx, contractAddr)
	if contractInfo == nil {
		return types.ErrContractNotFound
	}

	// Check ownership
	metaOld, _ := state.GetContractMetadata(contractAddr)
	/*
		if metaExists {
			if metaOld.OwnerAddress != senderAddr.String() {
				return sdkErrors.Wrap(types.ErrUnauthorized, "metadata can only be changed by the contract owner")
			}
		} else {
			if contractInfo.Admin != senderAddr.String() {
				return sdkErrors.Wrap(types.ErrUnauthorized, "metadata can only be created by the contract admin")
			}
		}
	*/
	metaNew := k.UpdateContractMetadata(&metaOld, &metaUpdates)

	// Save the updated metadata

	state.SetContractMetadata(contractAddr, *metaNew)

	// Emit event
	types.EmitContractMetadataSetEvent(
		ctx,
		contractAddr,
		*metaNew,
	)

	return nil
}

// GetContractMetadata returns the contract metadata for the given contract address (if found).
func (k Keeper) GetContractMetadata(ctx sdk.Context, contractAddr sdk.AccAddress) *types.ContractMetadata {
	meta, found := k.state.ContractMetadataState(ctx).GetContractMetadata(contractAddr)
	if !found {
		return nil
	}

	return &meta
}

func (k Keeper) UpdateContractMetadata(metaOld *types.ContractMetadata, metaUpdates *types.ContractMetadata) *types.ContractMetadata {
	metaNew := &types.ContractMetadata{}
	*metaNew = *metaOld

	metaNew.OwnerAddress = metaUpdates.OwnerAddress

	metaNew.RewardsAddress = metaUpdates.RewardsAddress

	metaNew.MinimumRewardAmount = metaUpdates.MinimumRewardAmount

	metaNew.LiquidityTokenAddress = metaUpdates.LiquidityTokenAddress

	metaNew.LiquidStakeInterval = metaUpdates.LiquidStakeInterval

	metaNew.RedemptionInterval = metaUpdates.RedemptionInterval

	metaNew.RewardsWithdrawalInterval = metaUpdates.RewardsWithdrawalInterval

	metaNew.RedemptionAddress = metaUpdates.RedemptionAddress

	metaNew.RedemptionRateThreshold = metaUpdates.RedemptionRateThreshold

	metaNew.RedemptionIntervalThreshold = metaUpdates.RedemptionIntervalThreshold

	metaNew.MaximumThreshold = metaUpdates.MaximumThreshold

	metaNew.ArchwayRewardFundsTransferAddress = metaUpdates.ArchwayRewardFundsTransferAddress

	metaNew.LiquidityProviderAddress = metaUpdates.LiquidityProviderAddress

	metaNew.LiquidityProviderCommission = metaUpdates.LiquidityProviderCommission

	metaNew.AirdropDuration = metaUpdates.AirdropDuration

	metaNew.AirdropRecipientAddress = metaUpdates.AirdropRecipientAddress

	metaNew.AirdropVestingPeriod = metaUpdates.AirdropVestingPeriod

	return metaNew
}

/*
func (k keeper) HasContractAddress() bool {
	return m.ContractAddress != ""
}

func (k keeper) HasOwnerAddress() bool {
	return m.OwnerAddress != ""
}

func (k keeper) HasRewardsAddress() bool {
	return m.RewardsAddress != ""
}

func (k keeper) HasMinimumRewardAmount() bool {
	return m.MinimumRewardAmount != 0
}

func (k keeper) HasLiquidityTokenAddress() bool {
	return m.LiquidityTokenAddress != ""
}

func (k keeper) HasLiquidStakeInterval() bool {
	return m.LiquidStakeInterval != 0
}

func (k keeper) HasRedemptionInterval() bool {
	return m.RedemptionInterval != 0
}

func (k keeper) HasRewardsWithdrawalInterval() bool {
	return m.RewardsWithdrawalInterval != 0
}

func (k keeper) HasRedemptionAddress() bool {
	return m.RedemptionAddress != ""
}

func (k keeper) HasRedemptionRateThreshold() bool {
	return m.RedemptionRateThreshold != 0
}

func (k keeper) HasRedemptionIntervalThreshold() bool {
	return m.RedemptionIntervalThreshold != 0
}

func (k keeper) HasMaximumThreshold() bool {
	return m.MaximumThreshold != 0
}

func (k keeper) HasArchwayRewardFundsTransferAddress() bool {
	return m.ArchwayRewardFundsTransferAddress != ""
}

func (k keeper) HasLiquidityProviderAddress() bool {
	return m.LiquidityProviderAddress != ""
}

func (k keeper) HasLiquidityProviderCommission() bool {
	return m.LiquidityProviderCommission != 0
}

func (k keeper) HasAirdropDuration() bool {
	return m.AirdropDuration != 0
}

func (k keeper) HasAirdropRecipientAddress() bool {
	return m.AirdropRecipientAddress != ""
}

func (k keeper) HasAirdropVestingPeriod() bool {
	return m.AirdropVestingPeriod != 0
}
*/
