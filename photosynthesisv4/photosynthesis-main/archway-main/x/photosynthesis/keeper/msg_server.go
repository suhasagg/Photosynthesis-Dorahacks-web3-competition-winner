package keeper

import (
	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgServer implements the module gRPC messaging service.
type MsgServer struct {
	keeper PhotosynthesisKeeper
}

// NewMsgServer creates a new gRPC messaging server.
func NewMsgServer(keeper PhotosynthesisKeeper) *MsgServer {
	return &MsgServer{
		keeper: keeper,
	}
}

func (k MsgServer) SetArchLiquidStakeInterval(ctx sdk.Context, msg *types.MsgSetArchLiquidStakeInterval) (*types.MsgSetArchLiquidStakeIntervalResponse, error) {
	k.keeper.SetArchLiquidStakeIntervalStore(ctx, msg.Interval)
	return &types.MsgSetArchLiquidStakeIntervalResponse{}, nil
}

func (k MsgServer) SetRedemptionRateQueryInterval(ctx sdk.Context, msg *types.MsgSetRedemptionRateQueryInterval) (*types.MsgSetRedemptionRateQueryIntervalResponse, error) {
	k.keeper.SetRedemptionRateQueryIntervalStore(ctx, msg.Interval)

	return &types.MsgSetRedemptionRateQueryIntervalResponse{}, nil
}

func (k MsgServer) SetRedemptionInterval(ctx sdk.Context, msg *types.MsgSetRedemptionInterval) (*types.MsgSetRedemptionIntervalResponse, error) {
	k.keeper.SetRedemptionIntervalStore(ctx, msg.Interval)
	return &types.MsgSetRedemptionIntervalResponse{}, nil
}

func (k MsgServer) SetRedemptionRateThreshold(ctx sdk.Context, msg *types.MsgSetRedemptionRateThreshold) (*types.MsgSetRedemptionRateThresholdResponse, error) {
	k.keeper.SetRedemptionRateThresholdStore(ctx, msg.Threshold)

	return &types.MsgSetRedemptionRateThresholdResponse{}, nil
}

func (k MsgServer) SetRewardsWithdrawalInterval(ctx sdk.Context, msg *types.MsgSetRewardsWithdrawalInterval) (*types.MsgSetRewardsWithdrawalIntervalResponse, error) {
	k.keeper.SetRewardsWithdrawalIntervalStore(ctx, sdk.AccAddress(msg.ContractAddress), msg.Interval)

	return &types.MsgSetRewardsWithdrawalIntervalResponse{}, nil
}
