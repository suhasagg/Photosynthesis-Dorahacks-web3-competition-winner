package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

const (
	flagLiquidityToken  = "liquidity-token"
	flagLiquidStaking   = "liquid-staking"
	flagMinRewardAmount = "min-reward-amount"
	// add other flags here
)

func GetPhotosynthesisFlags() (cmdFlags []*cobra.Command) {
	cmdFlags = []*cobra.Command{
		flags.String(flagLiquidityToken, "", "address of the liquidity token"),
		flags.Bool(flagLiquidStaking, false, "whether liquid staking is enabled"),
		flags.String(flagMinRewardAmount, "", "minimum reward amount to be liquid staked"),
		// add other flags here
	}
	return cmdFlags
}
