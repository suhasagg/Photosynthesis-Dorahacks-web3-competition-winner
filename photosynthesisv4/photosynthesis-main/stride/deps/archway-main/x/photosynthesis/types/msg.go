package photosynthesis

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// MsgMintNFT represents the message to mint an NFT
type MsgMintNFT struct {
	Sender     sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient  sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Denom      string         `json:"denom" yaml:"denom"`
	TokenID    string         `json:"token_id" yaml:"token_id"`
	TokenURI   string         `json:"token_uri" yaml:"token_uri"`
	Properties []string       `json:"properties" yaml:"properties"`
}

// NewMsgMintNFT creates a new MsgMintNFT instance
func NewMsgMintNFT(sender, recipient sdk.AccAddress, denom, tokenID, tokenURI string, properties []string) MsgMintNFT {
	return MsgMintNFT{
		Sender:     sender,
		Recipient:  recipient,
		Denom:      denom,
		TokenID:    tokenID,
		TokenURI:   tokenURI,
		Properties: properties,
	}
}

// Route returns the message route
func (msg MsgMintNFT) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgMintNFT) Type() string {
	return "mint_nft"
}

// ValidateBasic validates the message basic fields
func (msg MsgMintNFT) ValidateBasic() error {
	if msg.Sender.Empty() {
		return fmt.Errorf("sender address cannot be empty")
	}
	if msg.Recipient.Empty() {
		return fmt.Errorf("recipient address cannot be empty")
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return fmt.Errorf("denom cannot be empty or blank")
	}
	if strings.TrimSpace(msg.TokenID) == "" {
		return fmt.Errorf("token ID cannot be empty or blank")
	}
	if strings.TrimSpace(msg.TokenURI) == "" {
		return fmt.Errorf("token URI cannot be empty or blank")
	}
	if len(msg.Properties) == 0 {
		return fmt.Errorf("properties cannot be empty")
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg MsgMintNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the message signers
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// QueryNFTByIDParams represents the params for the query 'nft_by_id'
type QueryNFTByIDParams struct {
	Denom   string `json:"denom" yaml:"denom"`
	TokenID string `json:"token_id" yaml:"token_id"`
}

// NewQueryNFTByIDParams creates a new QueryNFTByIDParams instance
func NewQueryNFTByIDParams(denom, tokenID string) QueryNFTByIDParams {
	return QueryNFTByIDParams{
		Denom:   denom,
		TokenID: tokenID,
	}
}

// QueryNFTByDenomParams represents the params for the query 'nft_by_denom'
type QueryNFTByDenomParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// MsgLiquidStakeDeposit defines the message for liquid staking Archway rewards.
type MsgLiquidStakeDeposit struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
	Amount          sdk.Coins      `json:"amount"`
}

// NewMsgLiquidStakeDeposit creates a new MsgLiquidStakeDeposit instance.
func NewMsgLiquidStakeDeposit(contractAddress sdk.AccAddress, amount sdk.Coins) MsgLiquidStakeDeposit {
	return MsgLiquidStakeDeposit{
		ContractAddress: contractAddress,
		Amount:          amount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) Type() string {
	return "liquid_stake_deposit"
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) ValidateBasic() error {
	if msg.ContractAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "contract address cannot be empty")
	}

	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount cannot be empty or invalid")
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ContractAddress}
}

// MsgRedeemLiquidTokens defines the message for redeeming liquid tokens.
type MsgRedeemLiquidTokens struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
	Amount          sdk.Coins      `json:"amount"`
}

// NewMsgRedeemLiquidTokens creates a new MsgRedeemLiquidTokens instance.
func NewMsgRedeemLiquidTokens(contractAddress sdk.AccAddress, amount sdk.Coins) MsgRedeemLiquidTokens {
	return MsgRedeemLiquidTokens{
		ContractAddress: contractAddress,
		Amount:          amount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) Type() string {
	return "redeem_liquid_tokens"
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) ValidateBasic() error {
	if msg.ContractAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "contract address cannot be empty")
	}

	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount cannot be empty or invalid")
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ContractAddress}
}

// QueryLiquidTokensParams defines the parameters for the QueryLiquidTokens query.
type QueryLiquidTokensParams struct {
	ContractAddress sdk.AccAddress `json:"contract_address"`
}

type QueryLiquidStakingDepositParams struct {
	SenderAddress string `json:"sender_address"`
	ContractAddr  string `json:"contract_address"`
}

type QueryLiquidStakingDepositResponse struct {
	DepositAmount        sdk.Coins `json:"deposit_amount"`
	LiquidityTokenAmount sdk.Coins `json:"liquidity_token_amount"`
	NextRedemptionTime   int64     `json:"next_redemption_time"`
}

func NewQueryLiquidStakingDepositParams(senderAddr string, contractAddr string) QueryLiquidStakingDepositParams {
	return QueryLiquidStakingDepositParams{
		SenderAddress: senderAddr,
		ContractAddr:  contractAddr,
	}
}

func (p QueryLiquidStakingDepositParams) String() string {
	return fmt.Sprintf("Sender Address: %s\nContract Address: %s", p.SenderAddress, p.ContractAddr)
}

type QueryLiquidityTokenBalanceParams struct {
	SenderAddress string `json:"sender_address"`
}

type QueryLiquidityTokenBalanceResponse struct {
	Balance sdk.Coins `json:"balance"`
}

func NewQueryLiquidityTokenBalanceParams(senderAddr string) QueryLiquidityTokenBalanceParams {
	return QueryLiquidityTokenBalanceParams{
		SenderAddress: senderAddr,
	}
}

func (p QueryLiquidityTokenBalanceParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}

type QueryRedemptionRateParams struct{}

type QueryRedemptionRateResponse struct {
	RedemptionRate sdk.Dec `json:"redemption_rate"`
}

func (p QueryRedemptionRateParams) String() string {
	return "Redemption Rate Query"
}

type QueryAirdropStatusParams struct {
	SenderAddress string `json:"sender_address"`
}

type QueryAirdropStatusResponse struct {
	TotalAmount     sdk.Coins `json:"total_amount"`
	VestingSchedule string    `json:"vesting_schedule"`
	CurrentBalance  sdk.Coins `json:"current_balance"`
}

func NewQueryAirdropStatusParams(senderAddr string) QueryAirdropStatusParams {
	return QueryAirdropStatusParams{
		SenderAddress: senderAddr,
	}
}

func (p QueryAirdropStatusParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}

type QueryLiquidityTokenBalanceParams struct {
	SenderAddress sdk.AccAddress `json:"sender_address" yaml:"sender_address"`
}

func NewQueryLiquidityTokenBalanceParams(senderAddr sdk.AccAddress) QueryLiquidityTokenBalanceParams {
	return QueryLiquidityTokenBalanceParams{
		SenderAddress: senderAddr,
	}
}

type QueryLiquidityTokenBalanceResponse struct {
	LiquidityTokenBal sdk.Coins `json:"liquidity_token_balance" yaml:"liquidity_token_balance"`
}

func NewQueryLiquidityTokenBalanceResponse(liquidityTokenBal sdk.Coins) QueryLiquidityTokenBalanceResponse {
	return QueryLiquidityTokenBalanceResponse{
		LiquidityTokenBal: liquidityTokenBal,
	}
}

type QueryLiquidStakingDepositParams struct {
	SenderAddress string `json:"sender_address"`
	ContractAddr  string `json:"contract_address"`
}

type QueryLiquidStakingDepositResponse struct {
	DepositAmount        sdk.Coins `json:"deposit_amount"`
	LiquidityTokenAmount sdk.Coins `json:"liquidity_token_amount"`
	NextRedemptionTime   int64     `json:"next_redemption_time"`
}

func NewQueryLiquidStakingDepositParams(senderAddr string, contractAddr string) QueryLiquidStakingDepositParams {
	return QueryLiquidStakingDepositParams{
		SenderAddress: senderAddr,
		ContractAddr:  contractAddr,
	}
}

func (p QueryLiquidStakingDepositParams) String() string {
	return fmt.Sprintf("Sender Address: %s\nContract Address: %s", p.SenderAddress, p.ContractAddr)
}

type QueryLiquidityTokenBalanceParams struct {
	SenderAddress string `json:"sender_address"`
}

type QueryLiquidityTokenBalanceResponse struct {
	Balance sdk.Coins `json:"balance"`
}

func NewQueryLiquidityTokenBalanceParams(senderAddr string) QueryLiquidityTokenBalanceParams {
	return QueryLiquidityTokenBalanceParams{
		SenderAddress: senderAddr,
	}
}

func (p QueryLiquidityTokenBalanceParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}

type QueryRedemptionRateParams struct{}

type QueryRedemptionRateResponse struct {
	RedemptionRate sdk.Dec `json:"redemption_rate"`
}

func (p QueryRedemptionRateParams) String() string {
	return "Redemption Rate Query"
}

type QueryAirdropStatusParams struct {
	SenderAddress string `json:"sender_address"`
}

type QueryAirdropStatusResponse struct {
	TotalAmount     sdk.Coins `json:"total_amount"`
	VestingSchedule string    `json:"vesting_schedule"`
	CurrentBalance  sdk.Coins `json:"current_balance"`
}

func NewQueryAirdropStatusParams(senderAddr string) QueryAirdropStatusParams {
	return QueryAirdropStatusParams{
		SenderAddress: senderAddr,
	}
}

func (p QueryAirdropStatusParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}
