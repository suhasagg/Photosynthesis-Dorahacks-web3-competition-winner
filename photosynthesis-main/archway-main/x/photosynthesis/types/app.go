package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	ModuleName = "photosynthesis"
)

var (
	RouterKey      = ModuleName
	StoreKey       = ModuleName
	RedemptionCode = []byte{0x1}
)

type PhotosynthesisApp struct {
	cdc      *codec.Codec
	key      *sdk.KVStoreKey
	keeper   Keeper
}

func NewPhotosynthesisApp(cdc *codec.Codec, key *sdk.KVStoreKey, keeper Keeper) *PhotosynthesisApp {
	return &PhotosynthesisApp{
		cdc:      cdc,
		key:      key,
		keeper:   keeper,
	}
}

func (app *PhotosynthesisApp) Name() string {
	return ModuleName
}

func (app *PhotosynthesisApp) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (app *PhotosynthesisApp) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (app *PhotosynthesisApp) ValidateGenesis(data json.RawMessage) error {
	var state GenesisState
	if err := ModuleCdc.UnmarshalJSON(data, &state); err != nil {
		return err
	}
	return state.Validate()
}

func (app *PhotosynthesisApp) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var state GenesisState
	if len(data) == 0 {
		state = DefaultGenesisState()
	} else {
		ModuleCdc.MustUnmarshalJSON(data, &state)
	}
	InitGenesis(ctx, app.keeper, state)
	return []abci.ValidatorUpdate{}
}

func (app *PhotosynthesisApp) ExportGenesis(ctx sdk.Context) json.RawMessage {
	state := ExportGenesis(ctx, app.keeper)
	return ModuleCdc.MustMarshalJSON(state)
}

func (app *PhotosynthesisApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.keeper.BeginBlocker(ctx)
}

func (app *PhotosynthesisApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.keeper.EndBlocker(ctx)
}
