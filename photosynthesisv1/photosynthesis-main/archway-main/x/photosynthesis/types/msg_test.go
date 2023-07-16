package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSetArchLiquidStakeInterval(t *testing.T) {
	fromAddress := sdk.AccAddress("your-address-here")
	interval := int64(100)

	// Test creating a new MsgSetArchLiquidStakeInterval
	msg := NewMsgSetArchLiquidStakeInterval(fromAddress, interval)
	require.NotNil(t, msg)

	// Test ValidateBasic
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// TODO: Add more test cases for MsgSetArchLiquidStakeInterval
}

func TestMsgSetRedemptionRateQueryInterval(t *testing.T) {
	fromAddress := sdk.AccAddress("your-address-here")
	interval := int64(100)

	// Test creating a new MsgSetRedemptionRateQueryInterval
	msg := NewMsgSetRedemptionRateQueryInterval(fromAddress, interval)
	require.NotNil(t, msg)

	// Test ValidateBasic
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// TODO: Add more test cases for MsgSetRedemptionRateQueryInterval
}

func TestMsgSetRedemptionInterval(t *testing.T) {
	fromAddress := sdk.AccAddress("your-address-here")
	interval := uint64(100)

	// Test creating a new MsgSetRedemptionInterval
	msg := NewMsgSetRedemptionInterval(fromAddress, interval)
	require.NotNil(t, msg)

	// Test ValidateBasic
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// TODO: Add more test cases for MsgSetRedemptionInterval
}

func TestMsgSetRedemptionRateThreshold(t *testing.T) {
	fromAddress := sdk.AccAddress("your-address-here")
	threshold := sdk.NewDec(100)

	// Test creating a new MsgSetRedemptionRateThreshold
	msg := NewMsgSetRedemptionRateThreshold(fromAddress, threshold)
	require.NotNil(t, msg)

	// Test ValidateBasic
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// TODO: Add more test cases for MsgSetRedemptionRateThreshold
}

func TestMsgSetRewardsWithdrawalInterval(t *testing.T) {
	contractAddress := sdk.AccAddress("your-address-here")
	interval := uint64(100)

	// Test creating a new MsgSetRewardsWithdrawalInterval
	msg := NewMsgSetRewardsWithdrawalInterval(contractAddress, interval)
	require.NotNil(t, msg)

	// Test ValidateBasic
	err := msg.ValidateBasic()
	require.NoError(t, err)

	// TODO: Add more test cases for MsgSetRewardsWithdrawalInterval
}
