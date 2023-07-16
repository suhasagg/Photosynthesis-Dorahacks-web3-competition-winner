package keeper_test

import (
	"context"
	"strings"

	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/stretchr/testify/suite"

	"github.com/Stride-Labs/stride/v9/x/interchainquery/types"
)

const (
	HostChainId = "GAIA"
)

type MsgSubmitQueryResponseTestCase struct {
	validMsg types.MsgSubmitQueryResponse
	goCtx    context.Context
	query    types.Query
}

func (s *KeeperTestSuite) SetupMsgSubmitQueryResponse() MsgSubmitQueryResponseTestCase {
	// set up IBC
	s.CreateTransferChannel(HostChainId)

	// define the query
	goCtx := sdk.WrapSDKContext(s.Ctx)
	h, err := s.App.StakeibcKeeper.GetLightClientHeightSafely(s.Ctx, s.TransferPath.EndpointA.ConnectionID)
	s.Require().NoError(err)
	height := int64(h - 1) // start at the (LC height) - 1  height, which is the height the query executes at!
	result := []byte("result-example")
	proofOps := crypto.ProofOps{}
	fromAddress := s.TestAccs[0].String()
	expectedId := "9792c1d779a3846a8de7ae82f31a74d308b279a521fa9e0d5c4f08917117bf3e"

	_, addr, _ := bech32.DecodeAndConvert(s.TestAccs[0].String())
	data := banktypes.CreateAccountBalancesPrefix(addr)
	// save the query to Stride state, so it can be retrieved in the response
	query := types.Query{
		Id:           expectedId,
		CallbackId:   "withdrawalbalance",
		ChainId:      HostChainId,
		ConnectionId: s.TransferPath.EndpointA.ConnectionID,
		QueryType:    types.BANK_STORE_QUERY_WITH_PROOF,
		Request:      append(data, []byte(HostChainId)...),
		Ttl:          uint64(12545592938) * uint64(1000000000), // set ttl to August 2050, mult by nano conversion factor
	}

	return MsgSubmitQueryResponseTestCase{
		validMsg: types.MsgSubmitQueryResponse{
			ChainId:     HostChainId,
			QueryId:     expectedId,
			Result:      result,
			ProofOps:    &proofOps,
			Height:      height,
			FromAddress: fromAddress,
		},
		goCtx: goCtx,
		query: query,
	}
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_WrongProof() {
	tc := s.SetupMsgSubmitQueryResponse()

	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().ErrorContains(err, "Unable to verify membership proof: proof cannot be empty")
	s.Require().Nil(resp)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_UnknownId() {
	tc := s.SetupMsgSubmitQueryResponse()

	tc.query.Id = tc.query.Id + "INVALID_SUFFIX" // create an invalid query id
	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Equal(&types.MsgSubmitQueryResponseResponse{}, resp)

	// check that the query is STILL in the store, as it should NOT be deleted because the query was not found
	_, found := s.App.InterchainqueryKeeper.GetQuery(s.Ctx, tc.query.Id)
	s.Require().True(found)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_ExceededTtl() {
	tc := s.SetupMsgSubmitQueryResponse()

	// Remove key from the query type so to bypass the VerifyKeyProof function
	tc.query.QueryType = strings.ReplaceAll(tc.query.QueryType, "key", "")

	// set ttl to be expired
	tc.query.Ttl = uint64(1)
	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	// check that the query was deleted (since the query timed out)
	_, found := s.App.InterchainqueryKeeper.GetQuery(s.Ctx, tc.query.Id)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_FindAndInvokeCallback_WrongHostZone() {
	tc := s.SetupMsgSubmitQueryResponse()

	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	// rather than testing by executing the callback in its entirety,
	//   check by invoking it without a registered host zone and catching the appropriate error
	err := s.App.InterchainqueryKeeper.InvokeCallback(s.Ctx, &tc.validMsg, tc.query)
	s.Require().ErrorContains(err, "no registered zone for queried chain ID", "callback was invoked")
}

// To write this test, we need to write data to Gaia, then get the proof for that data and check it using the LC
// As a first pass, to verify proof checking, we will use an example from Stride integration testing
//     //   ...down the line, we may want to write tests here that verify the merkle check using proofs from tendermint's proof_test library, https://github.com/cometbft/cometbft/blob/75d51e18f740c7cbfb7d8b4d49182ee6c7f41982/crypto/merkle/proof_test.go
// func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_VerifyProofSuccess() {
// 	tc := s.SetupMsgSubmitQueryResponse()
// 	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

// 	// set the msgHeight to the light client's height (required to verify retured proofs)
// 	clientHeight := s.HostChain.GetClientState(s.TransferPath.EndpointA.ClientID).GetLatestHeight().GetRevisionHeight()
// 	tc.validMsg.Height = int64(clientHeight)
// proofOps := &crypto.ProofOps{
// 	Ops: []crypto.ProofOp{
// 		{
// 			Type: "ics23:iavl",
// 			Key:  []uint8{0x2, 0x20, 0x35, 0x2, 0x7c, 0x8f, 0xd9, 0x54, 0xf2, 0x7f, 0xd2, 0xab, 0x34, 0x63, 0xfa, 0x69, 0xff, 0x3b, 0x65, 0xc2, 0x74, 0x33, 0x25, 0x28, 0x1b, 0x43, 0xb, 0xe2, 0x38, 0x26, 0xe1, 0xef, 0x6c, 0xbd, 0x75, 0x61, 0x74, 0x6f, 0x6d},
// 			Data: []uint8{0x12, 0xf3, 0x4, 0xa, 0x27, 0x2, 0x20, 0x35, 0x2, 0x7c, 0x8f, 0xd9, 0x54, 0xf2, 0x7f, 0xd2, 0xab, 0x34, 0x63, 0xfa, 0x69, 0xff, 0x3b, 0x65, 0xc2, 0x74, 0x33, 0x25, 0x28, 0x1b, 0x43, 0xb, 0xe2, 0x38, 0x26, 0xe1, 0xef, 0x6c, 0xbd, 0x75, 0x61, 0x74, 0x6f, 0x6d, 0x12, 0x9e, 0x2, 0xa, 0x1b, 0x2, 0x14, 0xf1, 0x82, 0x96, 0x76, 0xdb, 0x57, 0x76, 0x82, 0xe9, 0x44, 0xfc, 0x34, 0x93, 0xd4, 0x51, 0xb6, 0x7f, 0xf3, 0xe2, 0x9f, 0x75, 0x61, 0x74, 0x6f, 0x6d, 0x12, 0x11, 0xa, 0x5, 0x75, 0x61, 0x74, 0x6f, 0x6d, 0x12, 0x8, 0x32, 0x30, 0x38, 0x31, 0x37, 0x31, 0x30, 0x31, 0x1a, 0xc, 0x8, 0x1, 0x18, 0x1, 0x20, 0x1, 0x2a, 0x4, 0x0, 0x2, 0xfe, 0xe, 0x22, 0x2c, 0x8, 0x1, 0x12, 0x5, 0x2, 0x4, 0xfe, 0xe, 0x20, 0x1a, 0x21, 0x20, 0x10, 0x3b, 0x75, 0xc5, 0x32, 0x2a, 0xda, 0x35, 0x9b, 0xce, 0x51, 0x73, 0x19, 0xe2, 0xfd, 0x3f, 0xc6, 0x2f, 0x65, 0x6a, 0x98, 0x40, 0x94, 0xea, 0xf4, 0x43, 0x7, 0xee, 0xa7, 0x8f, 0xe2, 0x89, 0x22, 0x2c, 0x8, 0x1, 0x12, 0x5, 0x4, 0x6, 0xfe, 0xe, 0x20, 0x1a, 0x21, 0x20, 0x49, 0x95, 0xab, 0x6e, 0xaf, 0x5f, 0x7d, 0x9a, 0x6b, 0x63, 0x6f, 0x6e, 0x28, 0xb9, 0x10, 0x74, 0x78, 0xaf, 0x6, 0x5e, 0x43, 0x28, 0xb9, 0xa9, 0xb, 0x1f, 0x9b, 0x68, 0x28, 0x7a, 0x5a, 0x4f, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0x6, 0xa, 0xfe, 0xe, 0x20, 0xb3, 0xac, 0xd0, 0xfa, 0x6a, 0xc4, 0x34, 0xc2, 0xf1, 0xc4, 0x96, 0x58, 0x97, 0xf0, 0x16, 0x67, 0x12, 0x12, 0x2, 0x6, 0x2f, 0x5b, 0x62, 0xe5, 0x21, 0x69, 0xc1, 0xd2, 0xa6, 0x95, 0xd6, 0x1f, 0x20, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0x8, 0x10, 0xfe, 0xe, 0x20, 0xad, 0x66, 0xbf, 0x41, 0x79, 0x42, 0xff, 0xaf, 0x5f, 0x19, 0xf1, 0x71, 0x6e, 0x78, 0xf8, 0xdf, 0x23, 0x70, 0xa7, 0xc1, 0x18, 0x75, 0xfa, 0x74, 0x4, 0xe, 0x98, 0x77, 0x53, 0xd2, 0x27, 0xfd, 0x20, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0xa, 0x1c, 0xfe, 0xe, 0x20, 0x11, 0x72, 0x36, 0x3c, 0x39, 0x41, 0xcf, 0xfb, 0x67, 0x6, 0xdd, 0xf5, 0xbd, 0x54, 0x6, 0xf5, 0xb4, 0x87, 0x54, 0xdd, 0xf6, 0xf9, 0x4c, 0x13, 0x58, 0x1e, 0x71, 0x2c, 0xbe, 0x4e, 0xde, 0xf5, 0x20, 0x1a, 0xa6, 0x2, 0xa, 0x27, 0x2, 0x20, 0x81, 0x35, 0x1f, 0xa5, 0x39, 0x67, 0x33, 0xc6, 0xb1, 0x9f, 0x65, 0xd5, 0xf4, 0xea, 0x11, 0x3f, 0x8d, 0x72, 0x98, 0x1e, 0x54, 0xe2, 0x98, 0x13, 0x53, 0x52, 0x7, 0x4b, 0xef, 0x5c, 0x54, 0x5, 0x75, 0x61, 0x74, 0x6f, 0x6d, 0x12, 0xf, 0xa, 0x5, 0x75, 0x61, 0x74, 0x6f, 0x6d, 0x12, 0x6, 0x32, 0x32, 0x35, 0x39, 0x37, 0x31, 0x1a, 0xc, 0x8, 0x1, 0x18, 0x1, 0x20, 0x1, 0x2a, 0x4, 0x0, 0x2, 0xd4, 0xe, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0x2, 0x4, 0xfe, 0xe, 0x20, 0xdb, 0xc8, 0xb, 0xe4, 0x68, 0x4f, 0x87, 0xf1, 0x1d, 0x9d, 0xf0, 0x98, 0xe3, 0x44, 0x8b, 0x4c, 0x38, 0xe7, 0x8a, 0x92, 0x34, 0xa9, 0xd8, 0x5d, 0xaa, 0x6f, 0x54, 0x9d, 0xa2, 0x57, 0x1b, 0x83, 0x20, 0x22, 0x2c, 0x8, 0x1, 0x12, 0x5, 0x4, 0x6, 0xfe, 0xe, 0x20, 0x1a, 0x21, 0x20, 0x49, 0x95, 0xab, 0x6e, 0xaf, 0x5f, 0x7d, 0x9a, 0x6b, 0x63, 0x6f, 0x6e, 0x28, 0xb9, 0x10, 0x74, 0x78, 0xaf, 0x6, 0x5e, 0x43, 0x28, 0xb9, 0xa9, 0xb, 0x1f, 0x9b, 0x68, 0x28, 0x7a, 0x5a, 0x4f, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0x6, 0xa, 0xfe, 0xe, 0x20, 0xb3, 0xac, 0xd0, 0xfa, 0x6a, 0xc4, 0x34, 0xc2, 0xf1, 0xc4, 0x96, 0x58, 0x97, 0xf0, 0x16, 0x67, 0x12, 0x12, 0x2, 0x6, 0x2f, 0x5b, 0x62, 0xe5, 0x21, 0x69, 0xc1, 0xd2, 0xa6, 0x95, 0xd6, 0x1f, 0x20, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0x8, 0x10, 0xfe, 0xe, 0x20, 0xad, 0x66, 0xbf, 0x41, 0x79, 0x42, 0xff, 0xaf, 0x5f, 0x19, 0xf1, 0x71, 0x6e, 0x78, 0xf8, 0xdf, 0x23, 0x70, 0xa7, 0xc1, 0x18, 0x75, 0xfa, 0x74, 0x4, 0xe, 0x98, 0x77, 0x53, 0xd2, 0x27, 0xfd, 0x20, 0x22, 0x2a, 0x8, 0x1, 0x12, 0x26, 0xa, 0x1c, 0xfe, 0xe, 0x20, 0x11, 0x72, 0x36, 0x3c, 0x39, 0x41, 0xcf, 0xfb, 0x67, 0x6, 0xdd, 0xf5, 0xbd, 0x54, 0x6, 0xf5, 0xb4, 0x87, 0x54, 0xdd, 0xf6, 0xf9, 0x4c, 0x13, 0x58, 0x1e, 0x71, 0x2c, 0xbe, 0x4e, 0xde, 0xf5, 0x20},
// 		},
// 		{
// 			Type: "ics23:simple",
// 			Key:  []uint8{0x62, 0x61, 0x6e, 0x6b},
// 			Data: []uint8{0xa, 0xfe, 0x1, 0xa, 0x4, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x20, 0x2e, 0x83, 0x16, 0x8, 0x2, 0xd7, 0x8f, 0x89, 0xba, 0xe4, 0xd4, 0x2, 0xcc, 0xbb, 0xc2, 0xe9, 0xc1, 0x86, 0xd4, 0x34, 0x5a, 0xbb, 0xaa, 0xba, 0x21, 0x5c, 0xdf, 0x9b, 0x34, 0x56, 0xe, 0x99, 0x1a, 0x9, 0x8, 0x1, 0x18, 0x1, 0x20, 0x1, 0x2a, 0x1, 0x0, 0x22, 0x27, 0x8, 0x1, 0x12, 0x1, 0x1, 0x1a, 0x20, 0x3d, 0x3d, 0xe1, 0xf9, 0x9b, 0x9d, 0x2c, 0xd, 0x80, 0x30, 0xe6, 0xf8, 0x2a, 0x5a, 0xe6, 0x22, 0x24, 0x3e, 0xdb, 0xd4, 0x32, 0x49, 0x7e, 0x29, 0xf2, 0x74, 0xef, 0xba, 0x91, 0xac, 0xd0, 0xc6, 0x22, 0x25, 0x8, 0x1, 0x12, 0x21, 0x1, 0xd2, 0x48, 0xfd, 0x48, 0x62, 0x63, 0xef, 0xcb, 0xa9, 0x4b, 0xfd, 0x8b, 0x59, 0xf1, 0x5f, 0x34, 0x11, 0x11, 0x43, 0x9a, 0xe9, 0x6c, 0x60, 0xe0, 0x4f, 0x2b, 0xab, 0xff, 0xff, 0x3d, 0x90, 0x2a, 0x22, 0x27, 0x8, 0x1, 0x12, 0x1, 0x1, 0x1a, 0x20, 0xae, 0xc4, 0xd4, 0x4a, 0x49, 0x72, 0x6e, 0xd9, 0xc, 0xb1, 0xea, 0x9c, 0x36, 0xd3, 0x61, 0xd9, 0x99, 0x1e, 0xf9, 0xfb, 0xfe, 0x7, 0xd, 0xaa, 0x9f, 0x7f, 0xb3, 0x60, 0x17, 0x21, 0xf6, 0x5, 0x22, 0x27, 0x8, 0x1, 0x12, 0x1, 0x1, 0x1a, 0x20, 0xf6, 0x20, 0x2c, 0x8d, 0xd7, 0xb1, 0x34, 0x4b, 0xc4, 0xb0, 0x4e, 0xc4, 0xeb, 0x4, 0x5b, 0x3, 0x7c, 0xe4, 0x4d, 0xa4, 0xb, 0x14, 0xc3, 0x54, 0x1f, 0x56, 0x3c, 0x19, 0xea, 0xd6, 0xa5, 0xe3, 0x22, 0x27, 0x8, 0x1, 0x12, 0x1, 0x1, 0x1a, 0x20, 0xe1, 0xd4, 0xef, 0x74, 0x6a, 0xa8, 0x28, 0x16, 0x6a, 0xc1, 0xa6, 0xfe, 0x9f, 0x2e, 0xde, 0x2c, 0x46, 0xd6, 0x13, 0x58, 0x73, 0x63, 0x8c, 0x9f, 0x1b, 0xec, 0xb5, 0x97, 0xe7, 0xe0, 0xbe, 0x23},
// 		},
// 	},
// }

// 	tc.validMsg.ProofOps = proofOps
// 	err := s.App.InterchainqueryKeeper.VerifyKeyProof(s.Ctx, &tc.validMsg, tc.query)
// 	s.Require().NoError(err)
// }
