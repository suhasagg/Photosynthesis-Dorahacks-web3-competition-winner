package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	ModuleName = "photosynthesis"
	// StoreKey is the module KV storage prefix key.
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the module.
	QuerierRoute = ModuleName
	// RouterKey is the msg router key for the module.
	RouterKey = ModuleName
)

const (
	CumulativeLiquidityAmountKey = "cumulative_liquidity_amount"
	RedemptionRecordPrefix       = "redemption_"
	ContractPrefix               = "contract_"
	LatestRedemptionTimeKey      = "latest_redemption_time"
)

const (
	RedemptionRateQueryInterval = "redemption_rate_query_interval"
	RedemptionRateThreshold     = "redemption_rate_threshold"
	RedemptionIntervalThreshold = "redemption_interval_threshold"
)

const (
	// KeyArchLiquidStakeInterval is the key for Archway liquid stake interval
	KeyArchLiquidStakeInterval = "ArchLiquidStakeInterval"

	// KeyRedemptionRateQueryInterval is the key for redemption rate query interval
	KeyRedemptionRateQueryInterval = "RedemptionRateQueryInterval"

	// KeyRedemptionInterval is the key for redemption interval for liquid tokens
	KeyRedemptionInterval = "RedemptionInterval"

	// KeyRedemptionRateThreshold is the key for redemption rate threshold for liquid tokens
	KeyRedemptionRateThreshold = "RedemptionRateThreshold"

	// KeyRewardsWithdrawalIntervalPrefix is the key prefix  for rewards withdrawal interval for the specified contract address
	KeyRewardsWithdrawalIntervalPrefix = "RewardsWithdrawalInterval:"
)

func timestampToUint64(timestamp string) uint64 {
	timeFormat := time.RFC3339
	parsedTime, _ := time.Parse(timeFormat, timestamp)
	ctimestamp := timestamppb.New(parsedTime)
	return uint64(ctimestamp.GetSeconds()*1000) + uint64(ctimestamp.GetNanos()/1000000)
}

func GetRedemptionRecordKey(timestamp string) []byte {
	key := "redemption_" + timestamp
	return []byte(key)
}

func GetContractKey(address string) []byte {
	key := "contract_" + address
	return []byte(key)
}

func GetStakeKey(address string) []byte {
	key := "stake_" + address
	return []byte(key)
}

// GetRewardsWithdrawalIntervalKey returns the key for rewards withdrawal interval for the specified contract address
func GetRewardsWithdrawalIntervalKey(contractAddress sdk.AccAddress) []byte {
	return []byte(KeyRewardsWithdrawalIntervalPrefix + contractAddress.String())
}
