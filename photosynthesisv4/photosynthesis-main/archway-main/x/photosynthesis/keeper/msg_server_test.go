package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/archway-network/archway/x/photosynthesis/keeper"
	"github.com/archway-network/archway/x/photosynthesis/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgServerTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	keeper    keeper.PhotosynthesisKeeper
	msgServer *keeper.MsgServer
}

func (s *MsgServerTestSuite) SetupTest() {
	app := simapp.Setup(false)
	s.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	s.msgServer = keeper.NewMsgServer(s.keeper)
}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
}

func (s *MsgServerTestSuite) TestSetArchLiquidStakeInterval() {
	msg := &types.MsgSetArchLiquidStakeInterval{
		Interval: 100,
	}

	res, err := s.msgServer.SetArchLiquidStakeInterval(s.ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	storedInterval, _ := s.keeper.GetArchLiquidStakeInterval(s.ctx)
	s.Require().Equal(uint64(100), storedInterval)
}

func (s *MsgServerTestSuite) TestSetRedemptionRateQueryInterval() {
	msg := &types.MsgSetRedemptionRateQueryInterval{
		Interval: 100,
	}

	res, err := s.msgServer.SetRedemptionRateQueryInterval(s.ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	storedInterval := s.keeper.GetRedemptionRateQueryInterval(s.ctx)
	s.Require().Equal(uint64(100), storedInterval)
}

func (s *MsgServerTestSuite) TestSetRedemptionInterval() {
	msg := &types.MsgSetRedemptionInterval{
		Interval: 100,
	}

	res, err := s.msgServer.SetRedemptionInterval(s.ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	storedInterval, _ := s.keeper.GetRedemptionInterval(s.ctx)
	s.Require().Equal(uint64(100), storedInterval)
}

func (s *MsgServerTestSuite) TestSetRedemptionRateThreshold() {
	msg := &types.MsgSetRedemptionRateThreshold{
		Threshold: "1.5",
	}

	res, err := s.msgServer.SetRedemptionRateThreshold(s.ctx, msg)
	s.Require().NoError(err)
	s.Require().NotNil(res)

	storedThreshold, _ := s.keeper.GetRedemptionRateThreshold(s.ctx)
	s.Require().Equal(float64(1.5), storedThreshold)
}
