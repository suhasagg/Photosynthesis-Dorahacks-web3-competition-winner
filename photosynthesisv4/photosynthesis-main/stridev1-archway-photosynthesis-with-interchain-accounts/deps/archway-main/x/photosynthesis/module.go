package types

import (
	"encoding/json"
	"fmt"

	"github.com/tendermint/tendermint/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Module struct {
	app *PhotosynthesisApp
}

func NewModule(app *PhotosynthesisApp) Module {
	return Module{
		app: app,
	}
}

func (m Module) Name() string {
	return ModuleName
}

func (m Module) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (m Module) Route() string {
	return RouterKey
}

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg  {
		case MsgEnableLiquidStaking:
			return handleEnableLiquidStaking(ctx, k, msg)
		case MsgSetLiquidStakeInterval:
			return handleSetLiquidStakeInterval(ctx, k, msg)
		case MsgSetMinRewardsToLiquidStake:
			return handleSetMinRewardsToLiquidStake(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unrecognized photosynthesis message type")
		}
	}
}

func (m Module) QuerierRoute() string {
	return ModuleName
}

func (m Module) NewQuerierHandler() sdk.Querier {
	return nil
}

func (m Module) InitGenesis(ctx sdk.Context, data json.RawMessage) []types.ValidatorUpdate {
	var state GenesisState
	if err := ModuleCdc.UnmarshalJSON(data, &state); err != nil {
		panic(fmt.Errorf("failed to unmarshal %s genesis state: %w", ModuleName, err))
	}
	InitGenesis(ctx, m.app.keeper, state)
	return []types.ValidatorUpdate{}
}

func handleEnableLiquidStaking(ctx sdk.Context, k Keeper, msg MsgEnableLiquidStaking) (*sdk.Result, error) {
	contractAddr := msg.ContractAddress

	// Get the existing metadata for the contract
	store := ctx.KVStore(k.storeKey)
	metadata := k.getContractMetadata(store, contractAddr)

	// Update the metadata to enable liquid staking
	metadata.EnableLiquidStaking = true
	k.setContractMetadata(store, contractAddr, metadata)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleSetLiquidStakeInterval(ctx sdk.Context, k Keeper, msg MsgSetLiquidStakeInterval) (*sdk.Result, error) {
	contractAddr := msg.ContractAddress
	interval := msg.Interval

	// Get the existing metadata for the contract
	store := ctx.KVStore(k.storeKey)
	metadata := k.getContractMetadata(store, contractAddr)

	// Update the metadata to set the liquid stake interval
	metadata.LiquidStakeInterval = interval
	k.setContractMetadata(store, contractAddr, metadata)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleSetMinRewardsToLiquidStake(ctx sdk.Context,k Keeper, msg MsgSetMinRewardsToLiquidStake) (*sdk.Result, error) {
	contractAddr := msg.ContractAddress
	amount := msg.Amount

	// Get the existing metadata for the contract
	store := ctx.KVStore(k.storeKey)
	metadata := k.getContractMetadata(store, contractAddr)

	// Update the metadata to set the minimum rewards required for liquid staking
	metadata.MinRewardsToLiquidStake = amount
	k.setContractMetadata(store, contractAddr, metadata)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
