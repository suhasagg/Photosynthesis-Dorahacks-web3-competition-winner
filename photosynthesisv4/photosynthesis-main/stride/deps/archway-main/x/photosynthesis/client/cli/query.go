package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/archway-network/archway/x/photosynthesis/types"
)

func GetQueryCmd() *cobra.Command {
	photosynthesisQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the photosynthesis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	photosynthesisQueryCmd.AddCommand(
		GetCmdQueryLiquidityProvider(),
		GetCmdQueryRewardParams(),
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

// GetCmdQueryRewardParams returns the command to query reward parameters
func GetCmdQueryRewardParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reward-params",
		Short: "Query the current reward parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			resp, err := queryClient.RewardParams(clientCtx.Context(), &types.QueryRewardParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp.Params)
		},
	}

	return cmd
}
