package keeper

import (
"fmt"
sdkmath "cosmossdk.io/math"
"github.com/cosmos/cosmos-sdk/codec"
sdk "github.com/cosmos/cosmos-sdk/types"
sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
"github.com/spf13/cast"
"github.com/archway-Labs/archway/v5/utils"
icqtypes "github.com/archway-Labs/archway/v5/x/interchainquery/types"
)


func QueryRedemptionRateCallback(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(query.ChainId, ICQCallbackID_QueryRedemptionRate,
"Starting query redemption rate callback, QueryId: %vs, QueryType: %s, Connection: %s", query.Id, query.QueryType, query.ConnectionId))

// Confirm host exists
chainId := query.ChainId
hostZone, found := k.GetHostZone(ctx, chainId)
if !found {
return sdkerrors.Wrapf(types.ErrHostZoneNotFound, "no registered zone for queried chain ID (%s)", chainId)
}

// Unmarshal the query response args to determine the redemption rate
redemptionRate, err := sdk.NewDecFromJSON(string(args)).ToInt()
if err != nil {
return sdkerrors.Wrap(err, "unable to determine redemption rate from query response")
}

k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_QueryRedemptionRate,
"Query response - Redemption Rate: %v", redemptionRate))

// Update the redemption rate for the host zone
hostZone.RedemptionRate = redemptionRate
k.SetHostZone(ctx, hostZone)

return nil
}
