package ante

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/archway-network/archway/x/photosynthesis/ante"
	"github.com/archway-network/archway/x/photosynthesis/types"
)

func TestPhotosynthesisAnteHandler(t *testing.T) {
	type testCase struct {
		name string
		// Inputs
		txMsgs []sdk.Msg // transaction messages
		// Params
		liquidStakingEnabled bool          // is liquid staking enabled
		minRewardAmount      sdk.Int       // minimum reward amount for liquid staking
		liquidTokenAddress   sdk.AccAddress // liquidity token address
		liquidStakeInterval  uint64        // interval for liquid staking rewards
		redemptionInterval   uint64        // interval for redeeming liquid tokens
		rewardsWithdrawalInterval uint64    // interval for withdrawing rewards from the pool
		redemptionAddress    sdk.AccAddress // address to distribute redeemed tokens
		redemptionRateThreshold sdk.Dec     // rate threshold for triggering token redemption
		redemptionIntervalThreshold uint64  // interval threshold for triggering token redemption
		maxStakingRewards     sdk.Coins     // maximum staking rewards
		archwayRewardsAddress sdk.AccAddress // Archway reward funds transfer address
		liquidityProviderAddress sdk.AccAddress // liquidity provider address for staking rewards
		liquidityProviderCommission sdk.Dec // commission rate for liquidity providers
		airdropDuration      uint64        // duration for claiming airdrop tokens
		airdropRecipientAddress sdk.AccAddress // airdrop recipient address
		airdropVestingPeriod  uint64        // vesting period for airdrop tokens
		autoCompoundInterval uint64        // interval for auto-compounding rewards
		autoCompoundPercentage sdk.Dec     // percentage of rewards for auto-compounding
		delegationIcaAddress  sdk.AccAddress // delegation ICA address for staking rewards
		feeIcaAddress         sdk.AccAddress // fee ICA address for charging liquidity provider fees
		transferTimeout       uint64        // timeout for IBC transfers
		// Output expected
		errExpected bool
	}

	testCases := []testCase{
		{
			name:                        "OK: All parameters set correctly",
			txMsgs:                      []sdk.Msg{testutils.NewMockMsg()},
			liquidStakingEnabled:        true,
			minRewardAmount:             sdk.NewInt(100),
			liquidTokenAddress:          sdk.AccAddress([]byte("liquidity-token")),
			liquidStakeInterval:         100,
			redemptionInterval:          200,
			rewardsWithdrawalInterval:   300,
			redemptionAddress:           sdk.AccAddress([]byte("redemption-address")),
			redemptionRateThreshold:     sdk.NewDecWithPrec(5, 1),
			redemptionIntervalThreshold: 400,
			maxStakingRewards:           sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000))),
			archwayRewardsAddress:       sdk.AccAddress([]byte("archway-rewards-address")),
			liquidityProviderAddress:    sdk.AccAddress([]byte("liquidity-provider-address")),
			airdropVestingPeriod:        1000,
			reinvestmentInterval:        100,
			reinvestmentPercentage:      sdk.NewDecWithPrec(50, 2),
			delegationIcaAddress:        sdk.AccAddress([]byte("delegation-ica-address")),
			feeIcaAddress:               sdk.AccAddress([]byte("fee-ica-address")),
			transferTimeout:             3600,
		},
		{
			name:                "Fail: Invalid minimum reward amount",
			enableLiquidStaking: true,
			minimumRewardAmount: sdk.NewCoin("", sdk.NewInt(1000)),
			errExpected:         true,
		},
	}
		// Run test cases
		for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		// Set up test context
		ctx := sdk.Context{}
		paramsKeeper := params.NewKeeper(nil, nil, rewardsTypes.ModuleName)
		paramsKeeper.SetParams(ctx, rewardsTypes.DefaultParams())
		rewardsKeeper := NewKeeper(paramsKeeper, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		contract := &rewardsTypes.Contract{
		Enabled:               tc.enabled,
		MinimumRewardAmount:   tc.minimumRewardAmount,
		LiquidityTokenAddress: tc.liquidityTokenAddress,
		LiquidStakeInterval:   tc.liquidStakeInterval,
		RedemptionInterval:    tc.redemptionInterval,
		RewardsWithdrawalInterval: tc.rewardsWithdrawalInterval,
		RedemptionAddress:     tc.redemptionAddress,
		RedemptionRateThreshold: tc.redemptionRateThreshold,
		RedemptionIntervalThreshold: tc.redemptionIntervalThreshold,
		MaximumThreshold:      tc.maximumThreshold,
		ArchwayRewardFundsTransferAddress: tc.archwayRewardFundsTransferAddress,
		LiquidityProviderAddress: tc.liquidityProviderAddress,
		LiquidityProviderCommission: tc.liquidityProviderCommission,
		AirdropDuration: tc.airdropDuration,
		AirdropRecipientAddress: tc.airdropRecipientAddress,
		AirdropVestingPeriod: tc.airdropVestingPeriod,
		ReinvestmentInterval: tc.reinvestmentInterval,
		ReinvestmentPercentage: tc.reinvestmentPercentage,
		DelegationIcaAddress: tc.delegationIcaAddress,
		FeeIcaAddress: tc.feeIcaAddress,
		TransferTimeout: tc.transferTimeout,
	}

		// Run validation
		err := ante.ValidatePhotosynthesisContract(contract)
		if tc.errExpected {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
	})
	}
	}
