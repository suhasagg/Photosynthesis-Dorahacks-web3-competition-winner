package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PhotosynthesisApp struct {
	cdc *codec.LegacyAmino
	key *sdk.KVStoreKey
}

func NewPhotosynthesisApp(cdc *codec.LegacyAmino, key *sdk.KVStoreKey) *PhotosynthesisApp {
	return &PhotosynthesisApp{
		cdc: cdc,
		key: key,
	}
}

func (app *PhotosynthesisApp) Name() string {
	return ModuleName
}

func (app *PhotosynthesisApp) RegisterCodec(cdc *codec.Codec) {
	//RegisterCodec(cdc)
}

/*
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
	var state *GenesisState
	if len(data) == 0 {
		state = DefaultGenesisState()
	} else {
		ModuleCdc.MustUnmarshalJSON(data, state)
	}
	app.keeper.InitGenesis(ctx, state)
	return []abci.ValidatorUpdate{}
}


func (app *PhotosynthesisApp) ExportGenesis(ctx sdk.Context) json.RawMessage {
	state := app.keeper.ExportGenesis(ctx)
	return ModuleCdc.MustMarshalJSON(state)
}
*/
/*
func (app *PhotosynthesisApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.keeper.BeginBlocker(ctx)
}
func (app *PhotosynthesisApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return app.keeper.EndBlocker(ctx)
}
*/
