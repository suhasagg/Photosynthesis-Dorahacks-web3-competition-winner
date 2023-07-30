package cli

import "github.com/spf13/cobra"

const (
	flagOwnerAddress                      = "owner-address"
	flagRewardsAddress                    = "rewards-address"
	flagRecordsLimit                      = "records-limit"
	flagRecordIDs                         = "record-ids"
	flagAirdropDuration                   = "airdrop-duration"
	flagAirdropRecipientAddress           = "airdrop-recipient-address"
	flagAirdropVestingPeriod              = "airdrop-vesting-period"
	flagArchwayRewardFundsTransferAddress = "archway-reward-funds-transfer-address"
	flagContractAddress                   = "contract-address"
	flagLiquidStakeInterval               = "liquid-stake-interval"
	flagLiquidityProviderAddress          = "liquidity-provider-address"
	flagLiquidityProviderCommission       = "liquidity-provider-commission"
	flagLiquidityTokenAddress             = "liquidity-token-address"
	flagMaximumThreshold                  = "maximum-threshold"
	flagMinimumRewardAmount               = "minimum-reward-amount"
	flagRedemptionAddress                 = "redemption-address"
	flagRedemptionInterval                = "redemption-interval"
	flagRedemptionIntervalThreshold       = "redemption-interval-threshold"
	flagRedemptionRateThreshold           = "redemption-rate-threshold"
	flagRewardsWithdrawalInterval         = "rewards-withdrawal-interval"
)

func addOwnerAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagOwnerAddress, "", "Address of the contract owner (bech 32)")
}

func addRewardsAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagRewardsAddress, "", "Rewards address to distribute contract rewards to (bech 32)")
}

func addRecordsLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRecordsLimit, 0, "Max number of rewards records to use (value can not be higher than the MaxWithdrawRecords module param")
}

func addRecordIDsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSlice(flagRecordIDs, []string{}, "Rewards record IDs to use (number of IDs can not be higher than the MaxWithdrawRecords module param")
}

func addAirdropDurationFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagAirdropDuration, 0, "Duration of the airdrop in seconds or another unit of time")
}

func addAirdropRecipientAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagAirdropRecipientAddress, "", "Address of the recipient of the airdrop")
}

func addAirdropVestingPeriodFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagAirdropVestingPeriod, 0, "Duration of the vesting period for the airdropped tokens")
}

func addArchwayRewardFundsTransferAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagArchwayRewardFundsTransferAddress, "", "Address from which the reward funds will be transferred from")
}

func addContractAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagContractAddress, "", "Address of the contract to which these settings apply")
}

func addLiquidStakeIntervalFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagLiquidStakeInterval, 0, "Duration of the liquidity stake interval in seconds or another unit of time")
}

func addLiquidityProviderAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagLiquidityProviderAddress, "", "Address of the liquidity provider")
}

func addLiquidityProviderCommissionFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagLiquidityProviderCommission, 0, "Commission rate for the liquidity provider")
}

func addLiquidityTokenAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagLiquidityTokenAddress, "", "Address of the liquidity token being used")
}

func addMaximumThresholdFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagMaximumThreshold, 0, "Maximum threshold for rewards")
}

func addMinimumRewardAmountFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagMinimumRewardAmount, 0, "Minimum amount of rewards that can be earned")
}

func addRedemptionAddressFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagRedemptionAddress, "", "Address where redeemed tokens will be sent to")
}

func addRedemptionIntervalFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRedemptionInterval, 0, "Duration of the redemption interval in seconds or another unit of time")
}

func addRedemptionIntervalThresholdFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRedemptionIntervalThreshold, 0, "Threshold for the redemption interval")
}

func addRedemptionRateThresholdFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRedemptionRateThreshold, 0, "Rate threshold for redemption")
}

func addRewardsWithdrawalIntervalFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRewardsWithdrawalInterval, 0, "Duration of the rewards withdrawal interval in seconds or another unit of time")
}
