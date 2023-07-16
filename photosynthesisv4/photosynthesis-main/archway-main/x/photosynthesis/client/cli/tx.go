package cli

import (
	"fmt"
	"github.com/archway-network/archway/x/photosynthesis/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"
)

// GetTxCmd builds tx command group for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Transaction commands for the Photosynthesis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add set-arch-liquid-stake-interval command
	cmd.AddCommand(SetCmdArchLiquidStakeIntervalCmd())
	// Add set-redemption-rate-query-interval command
	cmd.AddCommand(SetCmdRedemptionRateQueryIntervalCmd())
	// Add set-redemption-interval command
	cmd.AddCommand(GetCmdSetRedemptionInterval())
	// Add set-redemption-rate-threshold command
	cmd.AddCommand(GetCmdSetRedemptionRateThreshold())
	// Add set-rewards-withdrawal-interval command
	cmd.AddCommand(GetCmdSetRewardsWithdrawalInterval())

	// Set help template
	cmd.SetHelpTemplate(fmt.Sprintf(`%s
Transaction commands for the rewards module.

Usage:
  %s [command]

Available Commands:
  set-arch-liquid-stake-interval       Set Archway liquid stake interval
  set-redemption-rate-query-interval   Set redemption rate query interval
  set-redemption-interval              Set redemption interval for liquid tokens
  set-redemption-rate-threshold        Set the redemption rate threshold for liquid tokens
  set-rewards-withdrawal-interval      Sets the rewards withdrawal interval for the specified contract address

Use "%s [command] --help" for more information about a command.
`, cmd.Short, cmd.Use, cmd.Use))

	return cmd
}

/*

// GetTxCmd builds tx command group for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Transaction commands for the rewards module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// add set-contract-metadata command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}}

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setContractMetadataHelp))

	cmd.Flags().AddFlagSet(flagSetContractMetadata)
	cmd.AddCommand(getTxSetContractMetadataCmd())

	// add withdraw-rewards command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [flags]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, withdrawRewardsHelp))

	cmd.Flags().AddFlagSet(flagWithdrawRewards)
	cmd.AddCommand(getTxWithdrawRewardsCmd())

	// add set-flat-fee command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [flags]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setFlatFeeHelp))

	cmd.Flags().AddFlagSet(flagSetFlatFee)
	cmd.AddCommand(getTxSetFlatFeeCmd())

	// add deposit command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [amount] [flags]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, depositHelp))

	cmd.AddCommand(NewDepositCmd())

	// add set-rewards-withdrawal-interval command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [contract_address] [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRewardsWithdrawalIntervalHelp))

	cmd.AddCommand(GetCmdSetRewardsWithdrawalInterval())

	// add set-redemption-interval command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRedemptionIntervalHelp))

	cmd.AddCommand(GetCmdSetRedemptionInterval())

	// add set-redemption-rate-threshold command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [threshold]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRedemptionRateThresholdHelp))

	cmd.AddCommand(setRedemptionRateThresholdCmd())

	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [contract_address] [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRewardsWithdrawalIntervalHelp))

	cmd.AddCommand(GetCmdSetRewardsWithdrawalInterval())

	// add set-redemption-interval command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRedemptionIntervalHelp))

	cmd.AddCommand(GetCmdSetRedemptionInterval())

	// add set-redemption-rate-threshold command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [threshold]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRedemptionRateThresholdHelp))

	cmd.AddCommand(setRedemptionRateThresholdCmd())

	// add set-redemption-rate-query-interval command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setRedemptionRateQueryIntervalHelp))

	cmd.AddCommand(getCmdSetRedemptionRateQueryInterval())

	// add set-arch-liquid-stake-interval command
	cmd.SetHelpTemplate(fmt.Sprintf(`%s

Usage:
  {{.UseLine}} [interval]

Description:
  {{.Short}}

Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Args:
  {{.ArgsUsage}}

%s`, cmd.Long, setArchLiquidStakeIntervalHelp))

	cmd.AddCommand(getTxSetArchLiquidStakeInterval())

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

			contractAddress, err := pkg.ParseAccAddressArg("contract-address", args[0])
			if err != nil {
				return err
			}

			ownerAddress, err := pkg.ParseAccAddressFlag(cmd, flagOwnerAddress, false)
			if err != nil {
				return err
			}

			rewardsAddress, err := pkg.ParseAccAddressFlag(cmd, flagRewardsAddress, false)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetContractMetadata(senderAddr, contractAddress, ownerAddress, rewardsAddress)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	addOwnerAddressFlag(cmd)
	addRewardsAddressFlag(cmd)

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

func addRecordsLimitFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagRecordsLimit, 0, "maximum number of records to withdraw rewards for")
}

func addRecordIDsFlag(cmd *cobra.Command) {
	cmd.Flags().Uint64Slice(flagRecordIDs, nil, "list of record IDs to withdraw rewards for")
}

func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [amount]",
		Short: "Deposit liquidity tokens into the rewards pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			amount, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			depositor := clientCtx.GetFromAddress()

			msg := types.NewMsgDeposit(depositor, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			txBuilder, err := clientCtx.TxConfig.NewTxBuilder()
			if err != nil {
				return err
			}

			if err := txBuilder.SetMsgs(msg); err != nil {
				return err
			}

			if err := txBuilder.SetGasLimit(flags.DefaultGasLimit); err != nil {
				return err
			}

			if err := txBuilder.SetFeeAmount(flags.DefaultGasFee); err != nil {
				return err
			}

			if err := txBuilder.SetMemo(flags.DefaultMemo); err != nil {
				return err
			}

			if err := tx.Sign(clientCtx, depositor.String(), txBuilder); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, txBuilder)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSetRewardsWithdrawalInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-rewards-withdrawal-interval [contract_address] [interval]",
		Short: "Sets the rewards withdrawal interval for the specified contract address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse input arguments
			contractAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			interval, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgSetRewardsWithdrawalInterval(contractAddress, interval)
			err = cliCtx.EnsureAccountExists()
			if err != nil {
				return err
			}
			tx := auth.NewTxBuilder().WithMsgs(msg).WithGasPrices(gasPrices).WithFeeDenoms(feeDenoms).WithMemo(memo).WithTimeoutHeight(timeoutHeight).Build()
			err = tx.Sign(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			err = cliCtx.BroadcastTx(tx)
			if err != nil {
				return err
			}

			// Print result
			fmt.Printf("Rewards withdrawal interval set to %d epochs for contract address %s\n", interval, contractAddress.String())

			return nil
		},
	}
	return cmd
}

func GetCmdSetRedemptionInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-interval [interval]",
		Short: "Sets the redemption interval for liquid tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse input argument
			interval, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgSetRedemptionInterval(interval)
			err = cliCtx.EnsureAccountExists()
			if err != nil {
				return err
			}
			tx := auth.NewTxBuilder().WithMsgs(msg).WithGasPrices(gasPrices).WithFeeDenoms(feeDenoms).WithMemo(memo).WithTimeoutHeight(timeoutHeight).Build()
			err = tx.Sign(cliCtx.GetFromName())
			if err != nil {
				return err
			}
			err = cliCtx.BroadcastTx(tx)
			if err != nil {
				return err
			}

			// Print result
			fmt.Printf("Redemption interval set to %d epochs\n", interval)

			return nil
		},
	}
	return cmd
}

func getCmdSetRedemptionRateThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-rate-threshold [threshold]",
		Short: "Set the redemption rate threshold for liquid tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app := GetPhotosynthesisApp()

			// Parse the threshold argument
			threshold, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return fmt.Errorf("invalid threshold: %s", args[0])
			}

			// Set the redemption rate threshold in the app params
			app.paramKeeper.SetRedemptionRateThreshold(app.ctx, threshold)

			fmt.Printf("Redemption rate threshold set to %s\n", threshold)

			return nil
		},
	}

	return cmd
}

func GetCmdSetRewardsWithdrawalInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-rewards-withdrawal-interval [contract_address] [interval]",
		Short: "Set rewards withdrawal interval for a contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			contractAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return errors.Wrap(err, "invalid contract address")
			}

			interval, err := strconv.Atoi(args[1])
			if err != nil {
				return errors.Wrap(err, "invalid interval")
			}

			msg := photosynthesis.NewMsgSetRewardsWithdrawalInterval(contractAddr, uint64(interval))
			if err := msg.ValidateBasic(); err != nil {
				return errors.Wrap(err, "invalid message")
			}

			txBuilder := clientCtx.TxConfig.NewTxBuilder()
			if err := txBuilder.SetMsgs(msg); err != nil {
				return errors.Wrap(err, "failed to build transaction")
			}

			if err := client.SignTxAndBroadcast(clientCtx, txBuilder); err != nil {
				return errors.Wrap(err, "failed to sign and broadcast transaction")
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSetRedemptionInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-interval [interval]",
		Short: "Set redemption interval for liquid tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			interval, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.Wrap(err, "invalid interval")
			}

			msg := types.NewMsgSetRedemptionInterval(uint64(interval))
			if err := msg.ValidateBasic(); err != nil {
				return errors.Wrap(err, "invalid message")
			}

			txBuilder := clientCtx.TxConfig.NewTxBuilder()
			if err := txBuilder.SetMsgs(msg); err != nil {
				return errors.Wrap(err, "failed to build transaction")
			}

			if err := client.SignTxAndBroadcast(clientCtx, txBuilder); err != nil {
				return errors.Wrap(err, "failed to sign and broadcast transaction")
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func SetCmdRedemptionRateQueryIntervalCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-redemption-rate-query-interval [interval]",
		Short: "Set redemption rate query interval",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			// Parse input arguments
			interval, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to the network
			msg := types.NewMsgSetRedemptionRateQueryInterval(clientCtx.GetFromAddress(), interval)
			return clientutils.GenerateOrBroadcastMsgs(clientCtx, []sdk.Msg{msg})
		},
	}
}

func SetCmdArchLiquidStakeIntervalCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-arch-liquid-stake-interval [interval]",
		Short: "Set Archway liquid stake interval",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			// Parse input arguments
			interval, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to the network
			msg := types.NewMsgSetArchLiquidStakeInterval(clientCtx.GetFromAddress(), interval)
			return clientutils.GenerateOrBroadcastMsgs(clientCtx, []sdk.Msg{msg})
		},
	}
}


*/

func SetCmdArchLiquidStakeIntervalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-arch-liquid-stake-interval [interval]",
		Short: "Set Archway liquid stake interval",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse input arguments
			interval, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to the network
			msg := types.NewMsgSetArchLiquidStakeInterval(clientCtx.GetFromAddress(), interval)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func SetCmdRedemptionRateQueryIntervalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-rate-query-interval [interval]",
		Short: "Set redemption rate query interval",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse input arguments
			interval, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to the network
			msg := types.NewMsgSetRedemptionRateQueryInterval(clientCtx.GetFromAddress(), interval)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSetRedemptionInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-interval [interval]",
		Short: "Set redemption interval for liquid tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			interval, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid interval: %s", args[0])
			}

			msg := types.NewMsgSetRedemptionInterval(clientCtx.GetFromAddress(), uint64(interval))
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("invalid message: %v", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSetRedemptionRateThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redemption-rate-threshold [threshold]",
		Short: "Set the redemption rate threshold for liquid tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse the threshold argument
			threshold, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return fmt.Errorf("invalid threshold: %s", args[0])
			}

			// Build and sign the transaction, then broadcast to the network
			msg := types.NewMsgSetRedemptionRateThreshold(clientCtx.GetFromAddress(), threshold)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSetRewardsWithdrawalInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-rewards-withdrawal-interval [contract_address] [interval]",
		Short: "Sets the rewards withdrawal interval for the specified contract address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse input arguments
			contractAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			interval, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgSetRewardsWithdrawalInterval(contractAddress, interval)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
