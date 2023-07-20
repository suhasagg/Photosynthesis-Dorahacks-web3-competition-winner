package cli

import (
	"fmt"
	"github.com/archway-network/archway/x/photosynthesis/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

/*
func GetQueryCmd() *cobra.Command {
	photosynthesisQueryCmd := &cobra.Command{
		Use:                        photosynthesis.Module.Name(),
		Short:                      "Querying commands for the photosynthesis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	photosynthesisQueryCmd.AddCommand(
		GetCmdQueryLiquidityProvider(),
		// add other query commands here
	)

	return photosynthesisQueryCmd
}

// GetCmdQueryLiquidityProvider returns the command to query a liquidity provider
func GetCmdQueryLiquidityProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidity-provider [liquidity-provider-address]",
		Short: "query a liquidity provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			liquidityProviderAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryLiquidityProviderRequest{LiquidityProviderAddress: liquidityProviderAddr.String()}
			res, err := queryClient.LiquidityProvider(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

*/

// GetQueryCmd builds the query command group for the module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the rewards module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add query commands
	cmd.AddCommand(
		GetCmdQueryArchLiquidStakeInterval(),
		GetCmdQueryRedemptionRateQueryInterval(),
		GetCmdQueryRedemptionInterval(),
		GetCmdQueryRedemptionRateThreshold(),
		GetCmdQueryRewardsWithdrawalInterval(),
		GetCmdQueryLatestRedemptionRecord(),
		GetCmdQueryCumulativeLiquidityAmount(),
		GetCmdQueryTotalStake(),
		GetCmdQueryStake(),
		GetCmdListContracts(),
	)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// Sample query command functions
// Replace these functions with the actual query command implementations for your module

func GetCmdQueryArchLiquidStakeInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-arch-liquid-stake-interval",
		Short: "Query Archway liquid stake interval",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Archway liquid stake interval: Implement me!")
			return nil
		},
	}

	return cmd
}

func GetCmdQueryRedemptionRateQueryInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-redemption-rate-query-interval",
		Short: "Query redemption rate query interval",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Redemption rate query interval: Implement me!")
			return nil
		},
	}

	return cmd
}

func GetCmdQueryRedemptionInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-redemption-interval",
		Short: "Query redemption interval for liquid tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Redemption interval for liquid tokens: Implement me!")
			return nil
		},
	}

	return cmd
}

func GetCmdQueryRedemptionRateThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-redemption-rate-threshold",
		Short: "Query the redemption rate threshold for liquid tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Redemption rate threshold for liquid tokens: Implement me!")
			return nil
		},
	}

	return cmd
}

func GetCmdQueryRewardsWithdrawalInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-rewards-withdrawal-interval [contract_address]",
		Short: "Query the rewards withdrawal interval for the specified contract address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Rewards withdrawal interval for contract address %s: Implement me!\n", args[0])
			return nil
		},
	}

	return cmd
}

func GetCmdQueryLatestRedemptionRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-latest-redemption-record",
		Short: "Get the latest redemption record",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implement the query logic here
			return nil
		},
	}

	return cmd
}

func GetCmdQueryCumulativeLiquidityAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-cumulative-liquidity-amount",
		Short: "Get the cumulative liquidity amount",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implement the query logic here
			return nil
		},
	}

	return cmd
}

func GetCmdQueryTotalStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-total-stake",
		Short: "Get the total stake of all contracts",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implement the query logic here
			return nil
		},
	}

	return cmd
}

func GetCmdQueryStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-stake [contract_address]",
		Short: "Get the stake of the given contract address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implement the query logic here
			return nil
		},
	}

	return cmd
}

func GetCmdListContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contracts",
		Short: "List all contracts",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Implement the query logic here
			return nil
		},
	}

	return cmd
}
