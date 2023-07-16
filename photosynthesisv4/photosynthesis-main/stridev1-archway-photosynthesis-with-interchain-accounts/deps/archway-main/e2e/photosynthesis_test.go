package e2e

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/your-organization/photosynthesis/x/photosynthesis"
)

func TestPhotosynthesis(t *testing.T) {
	ctx, _, k, _, _, _ := createTestInput()
	params := k.GetParams(ctx)

	// Test minting an NFT
	sender := testAddr1
	recipient := testAddr2
	denom := "mydenom"
	tokenID := "1"
	tokenURI := "https://nft-uri.com"
	properties := []string{"prop1", "prop2"}

	msg := photosynthesis.NewMsgMintNFT(sender, recipient, denom, tokenID, tokenURI, properties)
	_, err := handleMsgMintNFT(ctx, msg, k)
	require.NoError(t, err)

	// Test querying an NFT by ID
	q := photosynthesis.NewQueryNFTByIDParams(denom, tokenID)
	res, err := k.QueryNFTByID(ctx, q)
	require.NoError(t, err)
	require.Equal(t, res.Denom, denom)
	require.Equal(t, res.TokenID, tokenID)
	require.Equal(t, res.TokenURI, tokenURI)
	require.Equal(t, res.Properties, properties)

	// Test liquid staking deposit
	contractAddress := testAddr1
	amount := sdk.Coins{sdk.NewInt64Coin("stake", 1000000)}

	msg = photosynthesis.NewMsgLiquidStakeDeposit(contractAddress, amount)
	_, err = handleMsgLiquidStakeDeposit(ctx, msg, k)
	require.NoError(t, err)

	// Test redeeming liquid tokens
	msg = photosynthesis.NewMsgRedeemLiquidTokens(contractAddress, amount)
	_, err = handleMsgRedeemLiquidTokens(ctx, msg, k)
	require.NoError(t, err)

	// Test querying liquidity token balance
	q2 := photosynthesis.NewQueryLiquidityTokenBalanceParams(sender)
	res2, err := k.QueryLiquidityTokenBalance(ctx, q2)
	require.NoError(t, err)
	require.Equal(t, res2.LiquidityTokenBal, amount)

	// Test querying liquid staking deposit
	q3 := photosynthesis.NewQueryLiquidStakingDepositParams(sender, contractAddress.String())
	res3, err := k.QueryLiquidStakingDeposit(ctx, q3)
	require.NoError(t, err)
	require.Equal(t, res3.DepositAmount, amount)

	// Test querying redemption rate
	q4 := photosynthesis.NewQueryRedemptionRateParams()
	_, err = k.QueryRedemptionRate(ctx, q4)
	require.NoError(t, err)

	// Test querying airdrop status
	q5 := photosynthesis.NewQueryAirdropStatusParams(sender)
	_, err = k.QueryAirdropStatus(ctx, q5)
	require.NoError(t, err)
}
