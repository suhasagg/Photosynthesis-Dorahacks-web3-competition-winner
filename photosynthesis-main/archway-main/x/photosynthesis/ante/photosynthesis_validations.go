package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func NewPhotosynthesisAnteHandler(
	ctx sdk.Context,
	contract types.Contract,
	ak ante.AccountKeeper,
	bk types.BankKeeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer,
	amtGasConsumer ante.Antehandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(),                        // initialize context
		ante.NewRejectExtensionOptionsDecorator(),              // reject txs with unsupported extended options
		ante.NewMempoolFeeDecorator(),                          // check that the transaction has enough fees
		ante.NewValidateMemoDecorator(ak),                      // validate memo
		ante.NewConsumeGasForTxSizeDecorator(ak),               // consume gas for the tx size
		ante.NewRejectFeeGranterDecorator(),                    // reject txs with fee granters
		NewPhotosynthesisContractValidationDecorator(contract), // validate the contract parameters
		ante.NewSetPubKeyDecorator(ak),                         // set public key on the signer
		ante.NewValidateSigCountDecorator(ak),                  // validate the number of signatures
		ante.NewDeductFeeDecorator(ak, bk),                     // deduct fees from the account balance
		ante.NewSigGasConsumeDecorator(ak, sigGasConsumer),     // consume gas for signature verification
		ante.NewSigVerificationDecorator(ak, sigGasConsumer),   // verify signatures
		ante.NewIncrementSequenceDecorator(ak),                 // increment account sequence
		ante.NewGasRefundDecorator(ak),                         // refund gas
		ante.NewAnteTimeoutDecorator(amtGasConsumer),           // check if the tx has reached the antehandler timeout
	)
}

func ValidatePhotosynthesisContract(contract types.Contract) error {
	if !contract.EnableLiquidStaking || contract.MinimumRewardsToLiquidStake.IsZero() {
		return sdk.ErrInvalidAddress("Contract has not enabled liquid staking or has not set a minimum reward amount to be liquid staked")
	}

	if !isValidAddress(contract.LiquidityTokenAddress) {
		return sdk.ErrInvalidAddress("Invalid liquidity token address")
	}

	if contract.LiquidStakeInterval <= 0 {
		return sdk.ErrInvalidAddress("Invalid liquid stake interval")
	}

	if contract.RedemptionInterval <= 0 {
		return sdk.ErrInvalidAddress("Invalid redemption interval")
	}

	if contract.RewardsWithdrawalInterval <= 0 {
		return sdk.ErrInvalidAddress("Invalid rewards withdrawal interval")
	}

	if !isValidAddress(contract.RedemptionAddress) {
		return sdk.ErrInvalidAddress("Invalid redemption address")
	}

	if contract.RedemptionRateThreshold <= 0 {
		return sdk.ErrInvalidAddress("Invalid redemption rate threshold")
	}

	if contract.RedemptionIntervalThreshold <= 0 {
		return sdk.ErrInvalidAddress("Invalid redemption interval threshold")
	}

	if contract.MaximumThresholdRewardsToStake.IsZero() {
		return sdk.ErrInvalidAddress("Invalid maximum threshold for staking rewards")
	}

	// Validate that the contract has set a valid maximum threshold for staking rewards.
	if contract.MaximumThresholdRewardsToStake.IsZero() {
		return sdk.ErrInvalidAddress("Invalid maximum threshold for staking rewards") // error if maximum threshold is not set
	}

	// Validate that the contract has provided a valid Archway reward funds transfer address.
	if !isValidAddress(contract.ArchwayRewardFundsTransferAddress) {
		return sdk.ErrInvalidAddress("Invalid Archway reward funds transfer address")
	}

	// Validate that the contract has provided a valid liquidity provider address for staking rewards.
	if !isValidAddress(contract.LiquidityProviderAddress) {
		return sdk.ErrInvalidAddress("Invalid liquidity provider address")
	}

	// Validate that the contract has provided a valid commission rate for liquidity providers.
	if contract.LiquidityProviderCommissionRate <= 0 {
		return sdk.ErrInvalidAddress("Invalid commission rate for liquidity providers")
	}

	// Validate that the contract has set a valid airdrop duration for claiming airdrop tokens.
	if contract.AirdropDuration <= 0 {
		return sdk.ErrInvalidAddress("Invalid airdrop duration")
	}

	// Validate that the contract has provided a valid airdrop recipient address.
	if !isValidAddress(contract.AirdropRecipientAddress) {
		return sdk.ErrInvalidAddress("Invalid airdrop recipient address")
	}

	// Validate airdrop vesting period
	if contract.AirdropVestingPeriod <= 0 {
		return sdk.ErrInvalidAddress("Invalid airdrop vesting period")
	}

	// Validate reinvestment interval
	if contract.ReinvestmentInterval <= 0 {
		return sdk.ErrInvalidAddress("Invalid reinvestment interval")
	}

	// Validate percentage of rewards for reinvestment
	if contract.ReinvestmentPercentage <= 0 || contract.ReinvestmentPercentage > 100 {
		return sdk.ErrInvalidAddress("Invalid percentage of rewards for reinvestment")
	}

	// Validate delegation ICA address
	if !isValidAddress(contract.DelegationICAAddress) {
		return sdk.ErrInvalidAddress("Invalid delegation ICA address")
	}

	// Validate fee ICA address
	if !isValidAddress(contract.FeeICAAddress) {
		return sdk.ErrInvalidAddress("Invalid fee ICA address")
	}

	// Validate IBC transfer timeout
	if contract.TransferTimeout <= 0 {
		return sdk.ErrInvalidAddress("Invalid IBC transfer timeout")
	}

	// call next ante handler in decorator chain
	return next(ctx, tx, simulate)
}


