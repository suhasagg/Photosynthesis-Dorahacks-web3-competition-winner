package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

const (
	flagLiquidityToken  = "liquidity-token"
	flagLiquidStaking   = "liquid-staking"
	flagMinRewardAmount = "min-reward-amount"
	// add other flags here
)

// Get common flags
func commonFlags() []string {
	return []string{flags.FlagChainID, flags.FlagFrom, flags.FlagGas, flags.FlagGasAdjustment, flags.FlagGasPrices, flags.FlagBroadcastMode, flags.FlagDryRun, flags.FlagGenerateOnly}
}

// Add common flags to a command
func AddCommonFlags(cmd *cobra.Command) {
	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().Uint64(flags.FlagGas, flags.DefaultGasLimit, "gas limit to set per-transaction; set to value other than 0 to manually set")
	cmd.Flags().Float64(flags.FlagGasAdjustment, flags.DefaultGasAdjustment, "adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored")
	cmd.Flags().String(flags.FlagGasPrices, "", "gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)")
	cmd.Flags().String(flags.FlagBroadcastMode, flags.BroadcastSync, "transaction broadcasting mode (sync|async|block)")
	cmd.Flags().Bool(flags.FlagDryRun, false, "ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it")
	cmd.Flags().Bool(flags.FlagGenerateOnly, false, "build an unsigned transaction and write it to STDOUT")
	cmd.Flags().String(flagLiquidityToken, "", "address of the liquidity token")
	cmd.Flags().Bool(flagLiquidStaking, false, "whether liquid staking is enabled")
	cmd.Flags().String(flagMinRewardAmount, "", "minimum reward amount to be liquid staked")
}
