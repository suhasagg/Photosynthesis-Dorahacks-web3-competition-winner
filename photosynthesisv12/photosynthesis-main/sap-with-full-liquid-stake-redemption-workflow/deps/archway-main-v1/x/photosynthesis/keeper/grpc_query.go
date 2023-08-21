package keeper

import (
	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"strconv"
)

type QueryServer struct {
	keeper PhotosynthesisKeeper
}

func NewQueryServer(keeper PhotosynthesisKeeper) *QueryServer {
	return &QueryServer{keeper: keeper}
}

func (qs *QueryServer) RegisterGRPCServer(_ *grpc.Server) {}

func (qs *QueryServer) RegisterGateway(_ *runtime.ServeMux, _ string) {}

func (qs *QueryServer) QueryArchLiquidStakeInterval(ctx sdk.Context, req *types.QueryArchLiquidStakeIntervalRequest) (*types.QueryArchLiquidStakeIntervalResponse, error) {
	interval, err := qs.keeper.GetArchLiquidStakeInterval(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryArchLiquidStakeIntervalResponse{ArchLiquidStakeInterval: strconv.FormatUint(interval, 10)}, nil
}

func (qs *QueryServer) QueryRedemptionRateQueryInterval(ctx sdk.Context, req *types.QueryRedemptionRateQueryIntervalRequest) (*types.QueryRedemptionRateQueryIntervalResponse, error) {
	interval := qs.keeper.GetRedemptionRateQueryInterval(ctx)

	return &types.QueryRedemptionRateQueryIntervalResponse{RedemptionRateQueryInterval: strconv.FormatUint(interval, 10)}, nil
}

func (qs *QueryServer) QueryRedemptionInterval(ctx sdk.Context, req *types.QueryRedemptionIntervalRequest) (*types.QueryRedemptionIntervalResponse, error) {
	interval, err := qs.keeper.GetRedemptionInterval(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryRedemptionIntervalResponse{RedemptionInterval: strconv.FormatUint(interval, 10)}, nil
}

func (qs *QueryServer) QueryRedemptionRateThreshold(ctx sdk.Context, req *types.QueryRedemptionRateThresholdRequest) (*types.QueryRedemptionRateThresholdResponse, error) {
	threshold, err := qs.keeper.GetRedemptionRateThreshold(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryRedemptionRateThresholdResponse{RedemptionRateThreshold: threshold}, nil
}

func (qs *QueryServer) QueryRewardsWithdrawalInterval(ctx sdk.Context, req *types.QueryRewardsWithdrawalIntervalRequest) (*types.QueryRewardsWithdrawalIntervalResponse, error) {
	interval, err := qs.keeper.GetRewardsWithdrawalInterval(ctx, sdk.AccAddress(req.ContractAddress))
	if err != nil {
		return nil, err
	}

	return &types.QueryRewardsWithdrawalIntervalResponse{RewardsWithdrawalInterval: strconv.FormatUint(interval, 10)}, nil
}

func (qs *QueryServer) QueryLatestRedemptionRecord(ctx sdk.Context, req *types.QueryLatestRedemptionRecordRequest) (*types.QueryLatestRedemptionRecordResponse, error) {
	record, _ := qs.keeper.GetLatestRedemptionRecord(ctx)

	return &types.QueryLatestRedemptionRecordResponse{LatestRedemptionRecord: record.Timestamp}, nil
}

func (qs *QueryServer) QueryCumulativeLiquidityAmount(ctx sdk.Context, req *types.QueryCumulativeLiquidityAmountRequest) (*types.QueryCumulativeLiquidityAmountResponse, error) {
	amount, _ := qs.keeper.GetCumulativeLiquidityAmount(ctx)
	return &types.QueryCumulativeLiquidityAmountResponse{CumulativeLiquidityAmount: uint64(amount.Amount)}, nil
}
