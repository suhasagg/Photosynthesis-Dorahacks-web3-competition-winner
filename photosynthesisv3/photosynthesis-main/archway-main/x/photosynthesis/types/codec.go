package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ModuleCdc references the global x/module_name module codec.
// This codec is provided to all the modules the application depends on.
var ModuleCdc = codec.NewLegacyAmino()

// RegisterCodec registers concrete types on the codec.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetArchLiquidStakeInterval{}, "module_name/SetArchLiquidStakeInterval", nil)
	cdc.RegisterConcrete(&MsgSetRedemptionRateQueryInterval{}, "module_name/SetRedemptionRateQueryInterval", nil)
	cdc.RegisterConcrete(&MsgSetRedemptionInterval{}, "module_name/SetRedemptionInterval", nil)
	cdc.RegisterConcrete(&MsgSetRedemptionRateThreshold{}, "module_name/SetRedemptionRateThreshold", nil)
	cdc.RegisterConcrete(&MsgSetRewardsWithdrawalInterval{}, "module_name/SetRewardsWithdrawalInterval", nil)
}

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterInterfaces registers the module's interface types.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetArchLiquidStakeInterval{},
		&MsgSetRedemptionRateQueryInterval{},
		&MsgSetRedemptionInterval{},
		&MsgSetRedemptionRateThreshold{},
		&MsgSetRewardsWithdrawalInterval{},
	)
}

// GetConfig returns a fully-constructed Config.
