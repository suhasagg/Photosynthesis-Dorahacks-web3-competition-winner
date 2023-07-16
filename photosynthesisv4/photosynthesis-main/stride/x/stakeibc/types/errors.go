package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

// x/stakeibc module sentinel errors
var (
	ErrInvalidVersion                    = errorsmod.Register(ModuleName, 1501, "invalid version")
	ErrInvalidToken                      = errorsmod.Register(ModuleName, 1502, "invalid token denom")
	ErrInvalidHostZone                   = errorsmod.Register(ModuleName, 1503, "host zone not registered")
	ErrICAStake                          = errorsmod.Register(ModuleName, 1504, "ICA stake failed")
	ErrEpochNotFound                     = errorsmod.Register(ModuleName, 1505, "epoch not found")
	ErrRecordNotFound                    = errorsmod.Register(ModuleName, 1506, "record not found")
	ErrInvalidAmount                     = errorsmod.Register(ModuleName, 1507, "invalid amount")
	ErrValidatorAlreadyExists            = errorsmod.Register(ModuleName, 1508, "validator already exists")
	ErrNoValidatorWeights                = errorsmod.Register(ModuleName, 1509, "no non-zero validator weights")
	ErrValidatorNotFound                 = errorsmod.Register(ModuleName, 1510, "validator not found")
	ErrWeightsNotDifferent               = errorsmod.Register(ModuleName, 1511, "validator weights haven't changed")
	ErrValidatorDelegationChg            = errorsmod.Register(ModuleName, 1512, "can't change delegation on validator")
	ErrAcctNotScopedForFunc              = errorsmod.Register(ModuleName, 1513, "this account can't call this function")
	ErrInsufficientFunds                 = errorsmod.Register(ModuleName, 1514, "balance is insufficient")
	ErrInvalidUserRedemptionRecord       = errorsmod.Register(ModuleName, 1515, "user redemption record error")
	ErrRequiredFieldEmpty                = errorsmod.Register(ModuleName, 1516, "required field is missing")
	ErrInvalidNumValidator               = errorsmod.Register(ModuleName, 1517, "invalid number of validators")
	ErrValidatorNotRemoved               = errorsmod.Register(ModuleName, 1518, "validator not removed")
	ErrHostZoneNotFound                  = errorsmod.Register(ModuleName, 1519, "host zone not found")
	ErrOutsideIcqWindow                  = errorsmod.Register(ModuleName, 1520, "outside time window that accepts icqs")
	ErrParamNotFound                     = errorsmod.Register(ModuleName, 1521, "param not found")
	ErrUnmarshalFailure                  = errorsmod.Register(ModuleName, 1522, "unable to unmarshal data structure")
	ErrMarshalFailure                    = errorsmod.Register(ModuleName, 1523, "unable to marshal data structure")
	ErrInvalidPacketCompletionTime       = errorsmod.Register(ModuleName, 1524, "invalid packet completion time")
	ErrIntCast                           = errorsmod.Register(ModuleName, 1525, "unable to cast to safe cast int")
	ErrFeeAccountNotRegistered           = errorsmod.Register(ModuleName, 1526, "fee account is not registered")
	ErrRedemptionRateOutsideSafetyBounds = errorsmod.Register(ModuleName, 1527, "redemption rate outside safety bounds")
	ErrTxMsgDataInvalid                  = errorsmod.Register(ModuleName, 1528, "TxMsgData invalid")
	ErrFailedToRegisterHostZone          = errorsmod.Register(ModuleName, 1529, "failed to register host zone")
	ErrInvalidInterchainAccountAddress   = errorsmod.Register(ModuleName, 1530, "invalid interchain account address")
	ErrICAAccountNotFound                = errorsmod.Register(ModuleName, 1531, "ICA acccount not found on host zone")
	ErrICATxFailed                       = errorsmod.Register(ModuleName, 1532, "failed to submit ICA transaction")
	ErrICQFailed                         = errorsmod.Register(ModuleName, 1533, "failed to submit ICQ")
	ErrDivisionByZero                    = errorsmod.Register(ModuleName, 1534, "division by zero")
	ErrSlashExceedsSafetyThreshold       = errorsmod.Register(ModuleName, 1535, "slash is greater than safety threshold")
	ErrInvalidEpoch                      = errorsmod.Register(ModuleName, 1536, "invalid epoch tracker")
	ErrHostZoneICAAccountNotFound        = errorsmod.Register(ModuleName, 1537, "host zone's ICA account not found")
	ErrNoValidatorAmts                   = errorsmod.Register(ModuleName, 1538, "could not fetch validator amts")
	ErrMaxNumValidators                  = errorsmod.Register(ModuleName, 1539, "max number of validators reached")
	ErrUndelegationAmount                = errorsmod.Register(ModuleName, 1540, "Undelegation amount is greater than stakedBal")
	ErrRewardCollectorAccountNotFound    = errorsmod.Register(ModuleName, 1541, "Reward Collector account not found")
	ErrHaltedHostZone                    = errorsmod.Register(ModuleName, 1542, "Halted host zone found")
	ErrInsufficientLiquidStake           = errorsmod.Register(ModuleName, 1543, "Liquid staked amount is too small")
)
