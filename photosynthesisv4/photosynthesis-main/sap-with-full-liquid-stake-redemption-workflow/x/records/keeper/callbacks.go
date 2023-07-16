package keeper

import (
	icacallbackstypes "github.com/Stride-Labs/stride/v9/x/icacallbacks/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

const TRANSFER = "transfer"

// ICACallbacks wrapper struct for stakeibc keeper
type ICACallback func(Keeper, sdk.Context, channeltypes.Packet, *icacallbackstypes.AcknowledgementResponse, []byte) error

type ICACallbacks struct {
	k            Keeper
	icacallbacks map[string]ICACallback
}

var _ icacallbackstypes.ICACallbackHandler = ICACallbacks{}

func (k Keeper) ICACallbackHandler() ICACallbacks {
	return ICACallbacks{k, make(map[string]ICACallback)}
}

func (c ICACallbacks) CallICACallback(ctx sdk.Context, id string, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {
	return c.icacallbacks[id](c.k, ctx, packet, ackResponse, args)
}

func (c ICACallbacks) HasICACallback(id string) bool {
	_, found := c.icacallbacks[id]
	return found
}

func (c ICACallbacks) AddICACallback(id string, fn interface{}) icacallbackstypes.ICACallbackHandler {
	c.icacallbacks[id] = fn.(ICACallback)
	return c
}

func (c ICACallbacks) RegisterICACallbacks() icacallbackstypes.ICACallbackHandler {
	a := c.AddICACallback(TRANSFER, ICACallback(TransferCallback))
	return a.(ICACallbacks)
}
