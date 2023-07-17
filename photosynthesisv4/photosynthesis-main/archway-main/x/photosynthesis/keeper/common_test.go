package keeper

import (
	"testing"

	e2eTesting "github.com/archway-network/archway/e2e/testing"
	phototypes "github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type PhotosynthesisKeeperTestSuite struct {
	suite.Suite

	chain *e2eTesting.TestChain
}

func (s *PhotosynthesisKeeperTestSuite) SetupTest() {
	s.chain = e2eTesting.NewTestChain(s.T(), 1)
}

func (s *PhotosynthesisKeeperTestSuite) TestQueries() {
	keeper := s.chain.GetApp().PhotosynthesisKeeper
	ctx := s.chain.GetContext()

	// Setup test data
	archLiquidStakeInterval := uint64(100)
	redemptionRateQueryInterval := uint64(200)
	redemptionInterval := uint64(300)
	redemptionRateThreshold := sdk.NewDec(0.5)

	keeper.SetArchLiquidStakeIntervalStore(ctx, archLiquidStakeInterval)
	keeper.SetRedemptionRateQueryIntervalStore(ctx, redemptionRateQueryInterval)
	keeper.SetRedemptionIntervalStore(ctx, redemptionInterval)
	keeper.SetRedemptionRateThresholdStore(ctx, redemptionRateThreshold.String())

	queryServer := NewQueryServer(keeper)

	// Test QueryArchLiquidStakeInterval
	resp1, err := queryServer.QueryArchLiquidStakeInterval(ctx, &phototypes.QueryArchLiquidStakeIntervalRequest{})
	s.Require().NoError(err)
	s.Assert().Equal(archLiquidStakeInterval, resp1.ArchLiquidStakeInterval)

	// Test QueryRedemptionRateQueryInterval
	resp2, err := queryServer.QueryRedemptionRateQueryInterval(ctx, &phototypes.QueryRedemptionRateQueryIntervalRequest{})
	s.Require().NoError(err)
	s.Assert().Equal(redemptionRateQueryInterval, resp2.RedemptionRateQueryInterval)

	// Test QueryRedemptionInterval
	resp3, err := queryServer.QueryRedemptionInterval(ctx, &phototypes.QueryRedemptionIntervalRequest{})
	s.Require().NoError(err)
	s.Assert().Equal(redemptionInterval, resp3.RedemptionInterval)

	// Test QueryRedemptionRateThreshold
	resp4, err := queryServer.QueryRedemptionRateThreshold(ctx, &phototypes.QueryRedemptionRateThresholdRequest{})
	s.Require().NoError(err)
	s.Assert().Equal(redemptionRateThreshold, resp4.RedemptionRateThreshold)
}

func TestPhotosynthesisKeeper(t *testing.T) {
	suite.Run(t, new(PhotosynthesisKeeperTestSuite))
}
