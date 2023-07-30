package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

/*

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
//func (msg MsgMintNFT) GetSignBytes() []byte {
//	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
//}

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

// NewMsgLiquidStakeDeposit creates a new MsgLiquidStakeDeposit instance.
func NewMsgLiquidStakeDeposit(contractAddress sdk.AccAddress, amount sdk.Coins) MsgLiquidStakeDeposit {
	return MsgLiquidStakeDeposit{
		ContractAddress: string(contractAddress),
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
//func (msg MsgLiquidStakeDeposit) GetSignBytes() []byte {
//	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
//}

// GetSigners implements the sdk.Msg interface.
func (msg MsgLiquidStakeDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ContractAddress}
}

// MsgRedeemLiquidTokens defines the message for redeeming liquid tokens.

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
//func (msg MsgRedeemLiquidTokens) GetSignBytes() []byte {
//	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
//}

// GetSigners implements the sdk.Msg interface.
func (msg MsgRedeemLiquidTokens) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ContractAddress}
}

func (p QueryLiquidStakingDepositParams) String() string {
	return fmt.Sprintf("Sender Address: %s\nContract Address: %s", p.SenderAddress, p.ContractAddr)
}

func (p QueryRedemptionRateParams) String() string {
	return "Redemption Rate Query"
}

func (p QueryAirdropStatusParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}

func NewQueryLiquidityTokenBalanceResponse(liquidityTokenBal sdk.Coins) QueryLiquidityTokenBalanceResponse {
	return QueryLiquidityTokenBalanceResponse{
		liquidityTokenBal,
	}
}

func NewQueryLiquidStakingDepositParams(senderAddr string, contractAddr string) QueryLiquidStakingDepositParams {
	return QueryLiquidStakingDepositParams{
		SenderAddress: senderAddr,
		ContractAddr:  contractAddr,
	}
}

func NewQueryLiquidityTokenBalanceParams(senderAddr string) QueryLiquidityTokenBalanceParams {
	return QueryLiquidityTokenBalanceParams{
		SenderAddress: senderAddr,
	}
}

func (p QueryLiquidityTokenBalanceParams) String() string {
	return fmt.Sprintf("Sender Address: %s", p.SenderAddress)
}

func NewQueryAirdropStatusParams(senderAddr string) QueryAirdropStatusParams {
	return QueryAirdropStatusParams{
		SenderAddress: senderAddr,
	}
}
*/

var _ sdk.Msg = &MsgSetArchLiquidStakeInterval{}
var _ sdk.Msg = &MsgSetRedemptionRateQueryInterval{}
var _ sdk.Msg = &MsgSetRedemptionInterval{}
var _ sdk.Msg = &MsgSetRedemptionRateThreshold{}
var _ sdk.Msg = &MsgSetRewardsWithdrawalInterval{}

// NewMsgSetArchLiquidStakeInterval creates a new MsgSetArchLiquidStakeInterval instance
func NewMsgSetArchLiquidStakeInterval(fromAddress sdk.AccAddress, interval int64) *MsgSetArchLiquidStakeInterval {
	return &MsgSetArchLiquidStakeInterval{
		FromAddress: fromAddress.String(),
		Interval:    uint64(interval),
	}
}

// Route returns the message route
func (msg *MsgSetArchLiquidStakeInterval) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSetArchLiquidStakeInterval) Type() string {
	return "set_arch_liquid_stake_interval"
}

// ValidateBasic validates the message
func (msg *MsgSetArchLiquidStakeInterval) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.FromAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %v", err)
	}
	if msg.Interval < 0 {
		return errors.New("interval cannot be negative:")
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgSetArchLiquidStakeInterval) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the addresses of the message signers
func (msg *MsgSetArchLiquidStakeInterval) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}

// NewMsgSetRedemptionRateQueryInterval creates a new MsgSetRedemptionRateQueryInterval instance
func NewMsgSetRedemptionRateQueryInterval(fromAddress sdk.AccAddress, interval int64) *MsgSetRedemptionRateQueryInterval {
	return &MsgSetRedemptionRateQueryInterval{
		FromAddress: fromAddress.String(),
		Interval:    uint64(interval),
	}
}

// Route returns the message route
func (msg *MsgSetRedemptionRateQueryInterval) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSetRedemptionRateQueryInterval) Type() string {
	return "set_redemption_rate_query_interval"
}

// ValidateBasic validates the message
func (msg *MsgSetRedemptionRateQueryInterval) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.FromAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %v", err)
	}
	if msg.Interval < 0 {
		return errors.New("interval cannot be negative:")
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgSetRedemptionRateQueryInterval) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the addresses of the message signers
func (msg *MsgSetRedemptionRateQueryInterval) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}

// NewMsgSetRedemptionInterval creates a new MsgSetRedemptionInterval instance
func NewMsgSetRedemptionInterval(fromAddress sdk.AccAddress, interval uint64) *MsgSetRedemptionInterval {
	return &MsgSetRedemptionInterval{
		FromAddress: fromAddress.String(),
		Interval:    interval,
	}
}

// Route returns the message route
func (msg *MsgSetRedemptionInterval) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSetRedemptionInterval) Type() string {
	return "set_redemption_interval"
}

// ValidateBasic validates the message
func (msg *MsgSetRedemptionInterval) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.FromAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %v", err)
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgSetRedemptionInterval) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the addresses of the message signers
func (msg *MsgSetRedemptionInterval) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}

// NewMsgSetRedemptionRateThreshold creates a new MsgSetRedemptionRateThreshold instance
func NewMsgSetRedemptionRateThreshold(fromAddress sdk.AccAddress, threshold sdk.Dec) *MsgSetRedemptionRateThreshold {
	return &MsgSetRedemptionRateThreshold{
		FromAddress: fromAddress.String(),
		Threshold:   threshold.String(),
	}
}

// Route returns the message route
func (msg *MsgSetRedemptionRateThreshold) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSetRedemptionRateThreshold) Type() string {
	return "set_redemption_rate_threshold"
}

// ValidateBasic validates the message
func (msg *MsgSetRedemptionRateThreshold) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.FromAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %v", err)
	}
	threshold, err := sdk.NewDecFromStr(msg.Threshold)
	if err != nil {
		return errors.New("invalid threshold:")
	}
	if threshold.IsNegative() {
		return errors.New("threshold cannot be negative;")
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgSetRedemptionRateThreshold) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the addresses of the message signers
func (msg *MsgSetRedemptionRateThreshold) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.FromAddress)
	return []sdk.AccAddress{fromAddress}
}

// NewMsgSetRewardsWithdrawalInterval creates a new MsgSetRewardsWithdrawalInterval instance
func NewMsgSetRewardsWithdrawalInterval(contractAddress sdk.AccAddress, interval uint64) *MsgSetRewardsWithdrawalInterval {
	return &MsgSetRewardsWithdrawalInterval{
		ContractAddress: contractAddress.String(),
		Interval:        interval,
	}
}

// Route returns the message route
func (msg *MsgSetRewardsWithdrawalInterval) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSetRewardsWithdrawalInterval) Type() string {
	return "set_rewards_withdrawal_interval"
}

// ValidateBasic validates the message
func (msg *MsgSetRewardsWithdrawalInterval) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address: %v", err)
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over
func (msg *MsgSetRewardsWithdrawalInterval) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners returns the addresses of the message signers
func (msg *MsgSetRewardsWithdrawalInterval) GetSigners() []sdk.AccAddress {
	contractAddress, _ := sdk.AccAddressFromBech32(msg.ContractAddress)
	return []sdk.AccAddress{contractAddress}
}
