package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/archway-network/archway/pkg"
	"github.com/archway-network/archway/x/rewards/types"
)

// GetTxCmd builds tx command group for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Transaction commands for the rewards module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		getTxSetContractMetadataCmd(),
		getTxWithdrawRewardsCmd(),
		getTxSetFlatFeeCmd(),
	)

	return cmd
}

func getTxSetContractMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-contract-metadata [contract-address]",
		Args:  cobra.ExactArgs(1),
		Short: "Create / modify contract metadata (contract rewards parameters)",
		Long: fmt.Sprintf(`Create / modify contract metadata (contract rewards parameters).
Use the %q and / or the %q flag to specify which metadata field to set / update.`,
			flagOwnerAddress, flagRewardsAddress,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddr := clientCtx.GetFromAddress()
			//_, err = pkg.ParseAccAddressArg("contract-address", args[0])
			if err != nil {
				return err
			}

			contractadd, err := cmd.Flags().GetString(flagContractAddress)

			ownerAddress, err := cmd.Flags().GetString(flagOwnerAddress)
			if err != nil {
				return err
			}

			rewardsAddress, err := cmd.Flags().GetString(flagRewardsAddress)
			if err != nil {
				return err
			}

			airdropDuration, err := cmd.Flags().GetUint64(flagAirdropDuration)
			if err != nil {
				return err
			}

			airdropRecipientAddress, err := cmd.Flags().GetString(flagAirdropRecipientAddress)
			if err != nil {
				return err
			}

			airdropVestingPeriod, err := cmd.Flags().GetUint64(flagAirdropVestingPeriod)
			if err != nil {
				return err
			}

			archwayRewardFundsTransferAddress, err := cmd.Flags().GetString(flagArchwayRewardFundsTransferAddress)
			if err != nil {
				return err
			}

			liquidStakeInterval, err := cmd.Flags().GetUint64(flagLiquidStakeInterval)
			if err != nil {
				return err
			}

			liquidityProviderAddress, err := cmd.Flags().GetString(flagLiquidityProviderAddress)
			if err != nil {
				return err
			}

			liquidityProviderCommission, err := cmd.Flags().GetUint64(flagLiquidityProviderCommission)
			if err != nil {
				return err
			}

			liquidityTokenAddress, err := cmd.Flags().GetString(flagLiquidityTokenAddress)
			if err != nil {
				return err
			}

			maximumThreshold, err := cmd.Flags().GetUint64(flagMaximumThreshold)
			if err != nil {
				return err
			}

			minimumRewardAmount, err := cmd.Flags().GetUint64(flagMinimumRewardAmount)
			if err != nil {
				return err
			}

			redemptionAddress, err := cmd.Flags().GetString(flagRedemptionAddress)
			if err != nil {
				return err
			}

			redemptionInterval, err := cmd.Flags().GetUint64(flagRedemptionInterval)
			if err != nil {
				return err
			}

			redemptionIntervalThreshold, err := cmd.Flags().GetUint64(flagRedemptionIntervalThreshold)
			if err != nil {
				return err
			}

			redemptionRateThreshold, err := cmd.Flags().GetUint64(flagRedemptionRateThreshold)

			if err != nil {
				return err
			}

			rewardsWithdrawalInterval, err := cmd.Flags().GetUint64(flagRewardsWithdrawalInterval)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetContractMetadata(
				senderAddr.String(),
				types.ContractMetadata{
					OwnerAddress:                      ownerAddress,
					RewardsAddress:                    rewardsAddress,
					AirdropDuration:                   airdropDuration,
					AirdropRecipientAddress:           airdropRecipientAddress,
					AirdropVestingPeriod:              airdropVestingPeriod,
					ArchwayRewardFundsTransferAddress: archwayRewardFundsTransferAddress,
					ContractAddress:                   contractadd,
					LiquidStakeInterval:               liquidStakeInterval,
					LiquidityProviderAddress:          liquidityProviderAddress,
					LiquidityProviderCommission:       liquidityProviderCommission,
					LiquidityTokenAddress:             liquidityTokenAddress,
					MaximumThreshold:                  maximumThreshold,
					MinimumRewardAmount:               minimumRewardAmount,
					RedemptionAddress:                 redemptionAddress,
					RedemptionInterval:                redemptionInterval,
					RedemptionIntervalThreshold:       redemptionIntervalThreshold,
					RedemptionRateThreshold:           redemptionRateThreshold,
					RewardsWithdrawalInterval:         rewardsWithdrawalInterval,
				},
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	addOwnerAddressFlag(cmd)
	addRewardsAddressFlag(cmd)
	addAirdropDurationFlag(cmd)
	addAirdropRecipientAddressFlag(cmd)
	addAirdropVestingPeriodFlag(cmd)
	addArchwayRewardFundsTransferAddressFlag(cmd)
	addContractAddressFlag(cmd)
	addLiquidStakeIntervalFlag(cmd)
	addLiquidityProviderAddressFlag(cmd)
	addLiquidityProviderCommissionFlag(cmd)
	addLiquidityTokenAddressFlag(cmd)
	addMaximumThresholdFlag(cmd)
	addMinimumRewardAmountFlag(cmd)
	addRedemptionAddressFlag(cmd)
	addRedemptionIntervalFlag(cmd)
	addRedemptionIntervalThresholdFlag(cmd)
	addRedemptionRateThresholdFlag(cmd)
	addRewardsWithdrawalIntervalFlag(cmd)

	return cmd
}

func getTxWithdrawRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-rewards",
		Args:  cobra.NoArgs,
		Short: "Withdraw current credited rewards for the transaction sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddr := clientCtx.GetFromAddress()

			recordsLimit, err := pkg.GetUint64Flag(cmd, flagRecordsLimit, true)
			if err != nil {
				return err
			}

			recordIDs, err := pkg.GetUint64SliceFlag(cmd, flagRecordIDs, true)
			if err != nil {
				return err
			}

			if (len(recordIDs) > 0 && recordsLimit > 0) || (len(recordIDs) == 0 && recordsLimit == 0) {
				return fmt.Errorf("one of (%q, %q) flags must be set", flagRecordIDs, flagRecordsLimit)
			}

			var msg sdk.Msg
			if recordsLimit > 0 {
				msg = types.NewMsgWithdrawRewardsByLimit(senderAddr, recordsLimit)
			} else {
				msg = types.NewMsgWithdrawRewardsByIDs(senderAddr, recordIDs)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addRecordsLimitFlag(cmd)
	addRecordIDsFlag(cmd)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getTxSetFlatFeeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-flat-fee [contract-address] [fee-amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Set / modify contract flat fee",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddr := clientCtx.GetFromAddress()

			contractAddress, err := pkg.ParseAccAddressArg("contract-address", args[0])
			if err != nil {
				return err
			}

			deposit, err := pkg.ParseCoinArg("fee-amount", args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgFlatFee(senderAddr, contractAddress, deposit)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
