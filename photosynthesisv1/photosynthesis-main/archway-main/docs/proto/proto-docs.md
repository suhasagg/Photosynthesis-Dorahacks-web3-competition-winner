<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [archway/photosynthesis/photosynthesis.proto](#archway/photosynthesis/photosynthesis.proto)
    - [Coin](#.Coin)
    - [Contract](#.Contract)
    - [DepositRecord](#.DepositRecord)
    - [DepositRecords](#.DepositRecords)
    - [RedemptionRecord](#.RedemptionRecord)
    - [RedemptionRecords](#.RedemptionRecords)
  
- [archway/photosynthesis/query.proto](#archway/photosynthesis/query.proto)
    - [QueryAirdropStatusParams](#.QueryAirdropStatusParams)
    - [QueryAirdropStatusResponse](#.QueryAirdropStatusResponse)
    - [QueryArchLiquidStakeIntervalRequest](#.QueryArchLiquidStakeIntervalRequest)
    - [QueryArchLiquidStakeIntervalResponse](#.QueryArchLiquidStakeIntervalResponse)
    - [QueryCumulativeLiquidityAmountRequest](#.QueryCumulativeLiquidityAmountRequest)
    - [QueryCumulativeLiquidityAmountResponse](#.QueryCumulativeLiquidityAmountResponse)
    - [QueryLatestRedemptionRecordRequest](#.QueryLatestRedemptionRecordRequest)
    - [QueryLatestRedemptionRecordResponse](#.QueryLatestRedemptionRecordResponse)
    - [QueryLiquidStakingDepositParams](#.QueryLiquidStakingDepositParams)
    - [QueryLiquidStakingDepositResponse](#.QueryLiquidStakingDepositResponse)
    - [QueryLiquidTokensParams](#.QueryLiquidTokensParams)
    - [QueryLiquidityTokenBalanceParams](#.QueryLiquidityTokenBalanceParams)
    - [QueryLiquidityTokenBalanceResponse](#.QueryLiquidityTokenBalanceResponse)
    - [QueryRedemptionIntervalRequest](#.QueryRedemptionIntervalRequest)
    - [QueryRedemptionIntervalResponse](#.QueryRedemptionIntervalResponse)
    - [QueryRedemptionRateParams](#.QueryRedemptionRateParams)
    - [QueryRedemptionRateQueryIntervalRequest](#.QueryRedemptionRateQueryIntervalRequest)
    - [QueryRedemptionRateQueryIntervalResponse](#.QueryRedemptionRateQueryIntervalResponse)
    - [QueryRedemptionRateResponse](#.QueryRedemptionRateResponse)
    - [QueryRedemptionRateThresholdRequest](#.QueryRedemptionRateThresholdRequest)
    - [QueryRedemptionRateThresholdResponse](#.QueryRedemptionRateThresholdResponse)
    - [QueryRewardsWithdrawalIntervalRequest](#.QueryRewardsWithdrawalIntervalRequest)
    - [QueryRewardsWithdrawalIntervalResponse](#.QueryRewardsWithdrawalIntervalResponse)
  
    - [Query](#.Query)
  
- [archway/photosynthesis/tx.proto](#archway/photosynthesis/tx.proto)
    - [MsgLiquidStakeDeposit](#.MsgLiquidStakeDeposit)
    - [MsgMintNFT](#.MsgMintNFT)
    - [MsgRedeemLiquidTokens](#.MsgRedeemLiquidTokens)
    - [MsgSetArchLiquidStakeInterval](#.MsgSetArchLiquidStakeInterval)
    - [MsgSetArchLiquidStakeIntervalResponse](#.MsgSetArchLiquidStakeIntervalResponse)
    - [MsgSetRedemptionInterval](#.MsgSetRedemptionInterval)
    - [MsgSetRedemptionIntervalResponse](#.MsgSetRedemptionIntervalResponse)
    - [MsgSetRedemptionRateQueryInterval](#.MsgSetRedemptionRateQueryInterval)
    - [MsgSetRedemptionRateQueryIntervalResponse](#.MsgSetRedemptionRateQueryIntervalResponse)
    - [MsgSetRedemptionRateThreshold](#.MsgSetRedemptionRateThreshold)
    - [MsgSetRedemptionRateThresholdResponse](#.MsgSetRedemptionRateThresholdResponse)
    - [MsgSetRewardsWithdrawalInterval](#.MsgSetRewardsWithdrawalInterval)
    - [MsgSetRewardsWithdrawalIntervalResponse](#.MsgSetRewardsWithdrawalIntervalResponse)
  
    - [Msg](#.Msg)
  
- [archway/rewards/v1beta1/rewards.proto](#archway/rewards/v1beta1/rewards.proto)
    - [BlockRewards](#archway.rewards.v1beta1.BlockRewards)
    - [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata)
    - [FlatFee](#archway.rewards.v1beta1.FlatFee)
    - [Params](#archway.rewards.v1beta1.Params)
    - [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord)
    - [TxRewards](#archway.rewards.v1beta1.TxRewards)
  
- [archway/rewards/v1beta1/events.proto](#archway/rewards/v1beta1/events.proto)
    - [ContractFlatFeeSetEvent](#archway.rewards.v1beta1.ContractFlatFeeSetEvent)
    - [ContractMetadataSetEvent](#archway.rewards.v1beta1.ContractMetadataSetEvent)
    - [ContractRewardCalculationEvent](#archway.rewards.v1beta1.ContractRewardCalculationEvent)
    - [MinConsensusFeeSetEvent](#archway.rewards.v1beta1.MinConsensusFeeSetEvent)
    - [RewardsWithdrawEvent](#archway.rewards.v1beta1.RewardsWithdrawEvent)
  
- [archway/rewards/v1beta1/genesis.proto](#archway/rewards/v1beta1/genesis.proto)
    - [GenesisState](#archway.rewards.v1beta1.GenesisState)
  
- [archway/rewards/v1beta1/query.proto](#archway/rewards/v1beta1/query.proto)
    - [BlockTracking](#archway.rewards.v1beta1.BlockTracking)
    - [QueryBlockRewardsTrackingRequest](#archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest)
    - [QueryBlockRewardsTrackingResponse](#archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse)
    - [QueryContractMetadataRequest](#archway.rewards.v1beta1.QueryContractMetadataRequest)
    - [QueryContractMetadataResponse](#archway.rewards.v1beta1.QueryContractMetadataResponse)
    - [QueryEstimateTxFeesRequest](#archway.rewards.v1beta1.QueryEstimateTxFeesRequest)
    - [QueryEstimateTxFeesResponse](#archway.rewards.v1beta1.QueryEstimateTxFeesResponse)
    - [QueryFlatFeeRequest](#archway.rewards.v1beta1.QueryFlatFeeRequest)
    - [QueryFlatFeeResponse](#archway.rewards.v1beta1.QueryFlatFeeResponse)
    - [QueryOutstandingRewardsRequest](#archway.rewards.v1beta1.QueryOutstandingRewardsRequest)
    - [QueryOutstandingRewardsResponse](#archway.rewards.v1beta1.QueryOutstandingRewardsResponse)
    - [QueryParamsRequest](#archway.rewards.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#archway.rewards.v1beta1.QueryParamsResponse)
    - [QueryRewardsPoolRequest](#archway.rewards.v1beta1.QueryRewardsPoolRequest)
    - [QueryRewardsPoolResponse](#archway.rewards.v1beta1.QueryRewardsPoolResponse)
    - [QueryRewardsRecordsRequest](#archway.rewards.v1beta1.QueryRewardsRecordsRequest)
    - [QueryRewardsRecordsResponse](#archway.rewards.v1beta1.QueryRewardsRecordsResponse)
  
    - [Query](#archway.rewards.v1beta1.Query)
  
- [archway/rewards/v1beta1/tx.proto](#archway/rewards/v1beta1/tx.proto)
    - [MsgSetContractMetadata](#archway.rewards.v1beta1.MsgSetContractMetadata)
    - [MsgSetContractMetadataResponse](#archway.rewards.v1beta1.MsgSetContractMetadataResponse)
    - [MsgSetFlatFee](#archway.rewards.v1beta1.MsgSetFlatFee)
    - [MsgSetFlatFeeResponse](#archway.rewards.v1beta1.MsgSetFlatFeeResponse)
    - [MsgWithdrawRewards](#archway.rewards.v1beta1.MsgWithdrawRewards)
    - [MsgWithdrawRewards.RecordIDs](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs)
    - [MsgWithdrawRewards.RecordsLimit](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit)
    - [MsgWithdrawRewardsResponse](#archway.rewards.v1beta1.MsgWithdrawRewardsResponse)
  
    - [Msg](#archway.rewards.v1beta1.Msg)
  
- [archway/tracking/v1beta1/tracking.proto](#archway/tracking/v1beta1/tracking.proto)
    - [BlockTracking](#archway.tracking.v1beta1.BlockTracking)
    - [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo)
    - [TxInfo](#archway.tracking.v1beta1.TxInfo)
    - [TxTracking](#archway.tracking.v1beta1.TxTracking)
  
    - [ContractOperation](#archway.tracking.v1beta1.ContractOperation)
  
- [archway/tracking/v1beta1/genesis.proto](#archway/tracking/v1beta1/genesis.proto)
    - [GenesisState](#archway.tracking.v1beta1.GenesisState)
  
- [archway/tracking/v1beta1/query.proto](#archway/tracking/v1beta1/query.proto)
    - [QueryBlockGasTrackingRequest](#archway.tracking.v1beta1.QueryBlockGasTrackingRequest)
    - [QueryBlockGasTrackingResponse](#archway.tracking.v1beta1.QueryBlockGasTrackingResponse)
  
    - [Query](#archway.tracking.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="archway/photosynthesis/photosynthesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/photosynthesis.proto



<a name=".Coin"></a>

### Coin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [int64](#int64) |  |  |






<a name=".Contract"></a>

### Contract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Use string to represent sdk.AccAddress, you'll need to |
| `creator` | [string](#string) |  | handle the conversion in your application logic

Use string to represent sdk.AccAddress, you'll need to |
| `name` | [string](#string) |  | handle the conversion in your application logic |
| `stake` | [int64](#int64) |  | Use string to represent sdk.Int, you'll need to handle |
| `rewards` | [int64](#int64) |  | the conversion in your application logic

Use string to represent sdk.Int, you'll need to handle |
| `activated` | [bool](#bool) |  | the conversion in your application logic |






<a name=".DepositRecord"></a>

### DepositRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `epoch` | [int64](#int64) |  |  |
| `amount` | [int64](#int64) |  | Use string to represent sdk.Int, you'll need to handle |
| `status` | [string](#string) |  | the conversion in your application logic |






<a name=".DepositRecords"></a>

### DepositRecords



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [DepositRecord](#DepositRecord) | repeated |  |






<a name=".RedemptionRecord"></a>

### RedemptionRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `timestamp` | [string](#string) |  |  |
| `liquidity_amount` | [string](#string) |  |  |






<a name=".RedemptionRecords"></a>

### RedemptionRecords



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [RedemptionRecord](#RedemptionRecord) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/photosynthesis/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/query.proto



<a name=".QueryAirdropStatusParams"></a>

### QueryAirdropStatusParams
QueryAirdropStatusParams defines the parameters for the QueryAirdropStatus
query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  |  |






<a name=".QueryAirdropStatusResponse"></a>

### QueryAirdropStatusResponse
QueryAirdropStatusResponse defines the response to the QueryAirdropStatus
query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_amount` | [Coin](#Coin) | repeated |  |
| `vesting_schedule` | [string](#string) |  |  |
| `current_balance` | [Coin](#Coin) | repeated |  |






<a name=".QueryArchLiquidStakeIntervalRequest"></a>

### QueryArchLiquidStakeIntervalRequest







<a name=".QueryArchLiquidStakeIntervalResponse"></a>

### QueryArchLiquidStakeIntervalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `arch_liquid_stake_interval` | [string](#string) |  |  |






<a name=".QueryCumulativeLiquidityAmountRequest"></a>

### QueryCumulativeLiquidityAmountRequest







<a name=".QueryCumulativeLiquidityAmountResponse"></a>

### QueryCumulativeLiquidityAmountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cumulative_liquidity_amount` | [uint64](#uint64) |  |  |






<a name=".QueryLatestRedemptionRecordRequest"></a>

### QueryLatestRedemptionRecordRequest







<a name=".QueryLatestRedemptionRecordResponse"></a>

### QueryLatestRedemptionRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `latest_redemption_record` | [string](#string) |  |  |






<a name=".QueryLiquidStakingDepositParams"></a>

### QueryLiquidStakingDepositParams
QueryLiquidStakingDepositParams defines the parameters for the
QueryLiquidStakingDeposit query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name=".QueryLiquidStakingDepositResponse"></a>

### QueryLiquidStakingDepositResponse
QueryLiquidStakingDepositResponse defines the response to the
QueryLiquidStakingDeposit query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_amount` | [Coin](#Coin) | repeated |  |
| `liquidity_token_amount` | [Coin](#Coin) | repeated |  |
| `next_redemption_time` | [int64](#int64) |  |  |






<a name=".QueryLiquidTokensParams"></a>

### QueryLiquidTokensParams
QueryLiquidTokensParams defines the parameters for the QueryLiquidTokens
query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |






<a name=".QueryLiquidityTokenBalanceParams"></a>

### QueryLiquidityTokenBalanceParams
QueryLiquidityTokenBalanceParams defines the parameters for the
QueryLiquidityTokenBalance query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  |  |






<a name=".QueryLiquidityTokenBalanceResponse"></a>

### QueryLiquidityTokenBalanceResponse
QueryLiquidityTokenBalanceResponse defines the response to the
QueryLiquidityTokenBalance query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balance` | [Coin](#Coin) | repeated |  |






<a name=".QueryRedemptionIntervalRequest"></a>

### QueryRedemptionIntervalRequest







<a name=".QueryRedemptionIntervalResponse"></a>

### QueryRedemptionIntervalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_interval` | [string](#string) |  |  |






<a name=".QueryRedemptionRateParams"></a>

### QueryRedemptionRateParams
QueryRedemptionRateParams defines the parameters for the QueryRedemptionRate
query.






<a name=".QueryRedemptionRateQueryIntervalRequest"></a>

### QueryRedemptionRateQueryIntervalRequest







<a name=".QueryRedemptionRateQueryIntervalResponse"></a>

### QueryRedemptionRateQueryIntervalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_rate_query_interval` | [string](#string) |  |  |






<a name=".QueryRedemptionRateResponse"></a>

### QueryRedemptionRateResponse
QueryRedemptionRateResponse defines the response to the QueryRedemptionRate
query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_rate` | [string](#string) |  |  |






<a name=".QueryRedemptionRateThresholdRequest"></a>

### QueryRedemptionRateThresholdRequest







<a name=".QueryRedemptionRateThresholdResponse"></a>

### QueryRedemptionRateThresholdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_rate_threshold` | [string](#string) |  |  |






<a name=".QueryRewardsWithdrawalIntervalRequest"></a>

### QueryRewardsWithdrawalIntervalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |






<a name=".QueryRewardsWithdrawalIntervalResponse"></a>

### QueryRewardsWithdrawalIntervalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_withdrawal_interval` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name=".Query"></a>

### Query
Define the gRPC service interface for the remote methods.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `QueryLiquidTokens` | [.QueryLiquidTokensParams](#QueryLiquidTokensParams) | [.Coin](#Coin) |  | |
| `QueryLiquidStakingDeposit` | [.QueryLiquidStakingDepositParams](#QueryLiquidStakingDepositParams) | [.QueryLiquidStakingDepositResponse](#QueryLiquidStakingDepositResponse) |  | |
| `QueryLiquidityTokenBalance` | [.QueryLiquidityTokenBalanceParams](#QueryLiquidityTokenBalanceParams) | [.QueryLiquidityTokenBalanceResponse](#QueryLiquidityTokenBalanceResponse) |  | |
| `QueryRedemptionRate` | [.QueryRedemptionRateParams](#QueryRedemptionRateParams) | [.QueryRedemptionRateResponse](#QueryRedemptionRateResponse) |  | |
| `QueryAirdropStatus` | [.QueryAirdropStatusParams](#QueryAirdropStatusParams) | [.QueryAirdropStatusResponse](#QueryAirdropStatusResponse) |  | |
| `QueryArchLiquidStakeInterval` | [.QueryArchLiquidStakeIntervalRequest](#QueryArchLiquidStakeIntervalRequest) | [.QueryArchLiquidStakeIntervalResponse](#QueryArchLiquidStakeIntervalResponse) |  | |
| `QueryRedemptionRateQueryInterval` | [.QueryRedemptionRateQueryIntervalRequest](#QueryRedemptionRateQueryIntervalRequest) | [.QueryRedemptionRateQueryIntervalResponse](#QueryRedemptionRateQueryIntervalResponse) |  | |
| `QueryRedemptionInterval` | [.QueryRedemptionIntervalRequest](#QueryRedemptionIntervalRequest) | [.QueryRedemptionIntervalResponse](#QueryRedemptionIntervalResponse) |  | |
| `QueryRedemptionRateThreshold` | [.QueryRedemptionRateThresholdRequest](#QueryRedemptionRateThresholdRequest) | [.QueryRedemptionRateThresholdResponse](#QueryRedemptionRateThresholdResponse) |  | |
| `QueryRewardsWithdrawalInterval` | [.QueryRewardsWithdrawalIntervalRequest](#QueryRewardsWithdrawalIntervalRequest) | [.QueryRewardsWithdrawalIntervalResponse](#QueryRewardsWithdrawalIntervalResponse) |  | |
| `QueryLatestRedemptionRecord` | [.QueryLatestRedemptionRecordRequest](#QueryLatestRedemptionRecordRequest) | [.QueryLatestRedemptionRecordResponse](#QueryLatestRedemptionRecordResponse) |  | |
| `QueryCumulativeLiquidityAmount` | [.QueryCumulativeLiquidityAmountRequest](#QueryCumulativeLiquidityAmountRequest) | [.QueryCumulativeLiquidityAmountResponse](#QueryCumulativeLiquidityAmountResponse) |  | |

 <!-- end services -->



<a name="archway/photosynthesis/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/tx.proto



<a name=".MsgLiquidStakeDeposit"></a>

### MsgLiquidStakeDeposit
MsgLiquidStakeDeposit defines the message for liquid staking Archway rewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `amount` | [Coin](#Coin) | repeated |  |






<a name=".MsgMintNFT"></a>

### MsgMintNFT
MsgMintNFT represents the message to mint an NFT


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `token_id` | [string](#string) |  |  |
| `token_uri` | [string](#string) |  |  |
| `properties` | [string](#string) | repeated |  |
| `creation_time` | [string](#string) |  |  |
| `last_update_time` | [string](#string) |  |  |






<a name=".MsgRedeemLiquidTokens"></a>

### MsgRedeemLiquidTokens
MsgRedeemLiquidTokens defines the message for redeeming liquid tokens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `amount` | [Coin](#Coin) | repeated |  |






<a name=".MsgSetArchLiquidStakeInterval"></a>

### MsgSetArchLiquidStakeInterval



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `interval` | [uint64](#uint64) |  |  |






<a name=".MsgSetArchLiquidStakeIntervalResponse"></a>

### MsgSetArchLiquidStakeIntervalResponse







<a name=".MsgSetRedemptionInterval"></a>

### MsgSetRedemptionInterval
MsgSetRedemptionInterval defines a message for setting the redemption
interval for liquid tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `interval` | [uint64](#uint64) |  |  |






<a name=".MsgSetRedemptionIntervalResponse"></a>

### MsgSetRedemptionIntervalResponse







<a name=".MsgSetRedemptionRateQueryInterval"></a>

### MsgSetRedemptionRateQueryInterval
MsgSetRedemptionRateQueryInterval defines a message for setting the
redemption rate query interval


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `interval` | [uint64](#uint64) |  |  |






<a name=".MsgSetRedemptionRateQueryIntervalResponse"></a>

### MsgSetRedemptionRateQueryIntervalResponse







<a name=".MsgSetRedemptionRateThreshold"></a>

### MsgSetRedemptionRateThreshold
MsgSetRedemptionRateThreshold defines a message for setting the redemption
rate threshold for liquid tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `threshold` | [string](#string) |  |  |






<a name=".MsgSetRedemptionRateThresholdResponse"></a>

### MsgSetRedemptionRateThresholdResponse







<a name=".MsgSetRewardsWithdrawalInterval"></a>

### MsgSetRewardsWithdrawalInterval
MsgSetRewardsWithdrawalInterval defines a message for setting the rewards
withdrawal interval for the specified contract address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `interval` | [uint64](#uint64) |  |  |






<a name=".MsgSetRewardsWithdrawalIntervalResponse"></a>

### MsgSetRewardsWithdrawalIntervalResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name=".Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetArchLiquidStakeInterval` | [.MsgSetArchLiquidStakeInterval](#MsgSetArchLiquidStakeInterval) | [.MsgSetArchLiquidStakeIntervalResponse](#MsgSetArchLiquidStakeIntervalResponse) |  | |
| `SetRedemptionRateQueryInterval` | [.MsgSetRedemptionRateQueryInterval](#MsgSetRedemptionRateQueryInterval) | [.MsgSetRedemptionRateQueryIntervalResponse](#MsgSetRedemptionRateQueryIntervalResponse) |  | |
| `SetRedemptionInterval` | [.MsgSetRedemptionInterval](#MsgSetRedemptionInterval) | [.MsgSetRedemptionIntervalResponse](#MsgSetRedemptionIntervalResponse) |  | |
| `SetRedemptionRateThreshold` | [.MsgSetRedemptionRateThreshold](#MsgSetRedemptionRateThreshold) | [.MsgSetRedemptionRateThresholdResponse](#MsgSetRedemptionRateThresholdResponse) |  | |
| `SetRewardsWithdrawalInterval` | [.MsgSetRewardsWithdrawalInterval](#MsgSetRewardsWithdrawalInterval) | [.MsgSetRewardsWithdrawalIntervalResponse](#MsgSetRewardsWithdrawalIntervalResponse) |  | |

 <!-- end services -->



<a name="archway/rewards/v1beta1/rewards.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/rewards.proto



<a name="archway.rewards.v1beta1.BlockRewards"></a>

### BlockRewards
BlockRewards defines block related rewards distribution data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height defines the block height. |
| `inflation_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | inflation_rewards is the rewards to be distributed. |
| `max_gas` | [uint64](#uint64) |  | max_gas defines the maximum gas for the block that is used to distribute inflation rewards (consensus parameter). |






<a name="archway.rewards.v1beta1.ContractMetadata"></a>

### ContractMetadata
ContractMetadata defines the contract rewards distribution options for a
particular contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address (bech32 encoded). |
| `owner_address` | [string](#string) |  | owner_address is the contract owner address that can modify contract reward options (bech32 encoded). That could be the contract admin or the contract itself. If owner_address is set to contract address, contract can modify the metadata on its own using WASM bindings. |
| `rewards_address` | [string](#string) |  | rewards_address is an address to distribute rewards to (bech32 encoded). If not set (empty), rewards are not distributed for this contract. |
| `minimum_reward_amount` | [uint64](#uint64) |  |  |
| `liquidity_token_address` | [string](#string) |  |  |
| `liquid_stake_interval` | [uint64](#uint64) |  |  |
| `redemption_interval` | [uint64](#uint64) |  |  |
| `rewards_withdrawal_interval` | [uint64](#uint64) |  |  |
| `redemption_address` | [string](#string) |  |  |
| `redemption_rate_threshold` | [uint64](#uint64) |  |  |
| `redemption_interval_threshold` | [uint64](#uint64) |  |  |
| `maximum_threshold` | [uint64](#uint64) |  |  |
| `archway_reward_funds_transfer_address` | [string](#string) |  |  |
| `liquidity_provider_address` | [string](#string) |  |  |
| `liquidity_provider_commission` | [uint64](#uint64) |  |  |
| `airdrop_duration` | [uint64](#uint64) |  |  |
| `airdrop_recipient_address` | [string](#string) |  |  |
| `airdrop_vesting_period` | [uint64](#uint64) |  |  |






<a name="archway.rewards.v1beta1.FlatFee"></a>

### FlatFee
FlatFee defines the flat fee for a particular contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address (bech32 encoded). |
| `flat_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee defines the minimum flat fee set by the contract_owner |






<a name="archway.rewards.v1beta1.Params"></a>

### Params
Params defines the module parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rewards_ratio` | [string](#string) |  | inflation_rewards_ratio defines the percentage of minted inflation tokens that are used for dApp rewards [0.0, 1.0]. If set to 0.0, no inflation rewards are distributed. |
| `tx_fee_rebate_ratio` | [string](#string) |  | tx_fee_rebate_ratio defines the percentage of tx fees that are used for dApp rewards [0.0, 1.0]. If set to 0.0, no fee rewards are distributed. |
| `max_withdraw_records` | [uint64](#uint64) |  | max_withdraw_records defines the maximum number of RewardsRecord objects used for the withdrawal operation. |






<a name="archway.rewards.v1beta1.RewardsRecord"></a>

### RewardsRecord
RewardsRecord defines a record that is used to distribute rewards later (lazy
distribution). This record is being created by the x/rewards EndBlocker and
pruned after the rewards are distributed. An actual rewards x/bank transfer
might be triggered by a Tx (via CLI for example) or by a contract via WASM
bindings. For a contract to trigger rewards transfer, contract address must
be set as the rewards_address in a corresponding ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id is the unique ID of the record. |
| `rewards_address` | [string](#string) |  | rewards_address is the address to distribute rewards to (bech32 encoded). |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards are the rewards to be transferred later. |
| `calculated_height` | [int64](#int64) |  | calculated_height defines the block height of rewards calculation event. |
| `calculated_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | calculated_time defines the block time of rewards calculation event. |






<a name="archway.rewards.v1beta1.TxRewards"></a>

### TxRewards
TxRewards defines transaction related rewards distribution data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [uint64](#uint64) |  | tx_id is the tracking transaction ID (x/tracking is the data source for this value). |
| `height` | [int64](#int64) |  | height defines the block height. |
| `fee_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee_rewards is the rewards to be distributed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/events.proto



<a name="archway.rewards.v1beta1.ContractFlatFeeSetEvent"></a>

### ContractFlatFeeSetEvent
ContractFlatFeeSetEvent is emitted when the contract flat fee is updated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the bech32 address of the contract for which the flat fee is set |
| `flat_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee defines the amount that has been set as the minimum fee for the contract |






<a name="archway.rewards.v1beta1.ContractMetadataSetEvent"></a>

### ContractMetadataSetEvent
ContractMetadataSetEvent is emitted when the contract metadata is created or
updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address. |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata defines the new contract metadata state. |






<a name="archway.rewards.v1beta1.ContractRewardCalculationEvent"></a>

### ContractRewardCalculationEvent
ContractRewardCalculationEvent is emitted when the contract reward is
calculated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address. |
| `gas_consumed` | [uint64](#uint64) |  | gas_consumed defines the total gas consumption by all WASM operations within one transaction. |
| `inflation_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | inflation_rewards defines the inflation rewards portions of the rewards. |
| `fee_rebate_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee_rebate_rewards defines the fee rebate rewards portions of the rewards. |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata defines the contract metadata (if set). |






<a name="archway.rewards.v1beta1.MinConsensusFeeSetEvent"></a>

### MinConsensusFeeSetEvent
MinConsensusFeeSetEvent is emitted when the minimum consensus fee is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fee` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | fee defines the updated minimum gas unit price. |






<a name="archway.rewards.v1beta1.RewardsWithdrawEvent"></a>

### RewardsWithdrawEvent
RewardsWithdrawEvent is emitted when credited rewards for a specific
rewards_address are distributed. Event could be triggered by a transaction
(via CLI for example) or by a contract via WASM bindings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_address` | [string](#string) |  | rewards_address defines the rewards address rewards are distributed to. |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards defines the total rewards being distributed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/genesis.proto



<a name="archway.rewards.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the initial state of the tracking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.rewards.v1beta1.Params) |  | params defines all the module parameters. |
| `contracts_metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) | repeated | contracts_metadata defines a list of all contracts metadata. |
| `block_rewards` | [BlockRewards](#archway.rewards.v1beta1.BlockRewards) | repeated | block_rewards defines a list of all block rewards objects. |
| `tx_rewards` | [TxRewards](#archway.rewards.v1beta1.TxRewards) | repeated | tx_rewards defines a list of all tx rewards objects. |
| `min_consensus_fee` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | min_consensus_fee defines the minimum gas unit price. |
| `rewards_record_last_id` | [uint64](#uint64) |  | rewards_record_last_id defines the last unique ID for a RewardsRecord objs. |
| `rewards_records` | [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord) | repeated | rewards_records defines a list of all active (undistributed) rewards records. |
| `flat_fees` | [FlatFee](#archway.rewards.v1beta1.FlatFee) | repeated | flat_fees defines a list of contract flat fee. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/query.proto



<a name="archway.rewards.v1beta1.BlockTracking"></a>

### BlockTracking
BlockTracking is the tracking information for a block.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rewards` | [BlockRewards](#archway.rewards.v1beta1.BlockRewards) |  | inflation_rewards defines the inflation rewards for the block. |
| `tx_rewards` | [TxRewards](#archway.rewards.v1beta1.TxRewards) | repeated | tx_rewards defines the transaction rewards for the block. |






<a name="archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest"></a>

### QueryBlockRewardsTrackingRequest
QueryBlockRewardsTrackingRequest is the request for
Query.BlockRewardsTracking.






<a name="archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse"></a>

### QueryBlockRewardsTrackingResponse
QueryBlockRewardsTrackingResponse is the response for
Query.BlockRewardsTracking.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockTracking](#archway.rewards.v1beta1.BlockTracking) |  |  |






<a name="archway.rewards.v1beta1.QueryContractMetadataRequest"></a>

### QueryContractMetadataRequest
QueryContractMetadataRequest is the request for Query.ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryContractMetadataResponse"></a>

### QueryContractMetadataResponse
QueryContractMetadataResponse is the response for Query.ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  |  |






<a name="archway.rewards.v1beta1.QueryEstimateTxFeesRequest"></a>

### QueryEstimateTxFeesRequest
QueryEstimateTxFeesRequest is the request for Query.EstimateTxFees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_limit` | [uint64](#uint64) |  | gas_limit is the transaction gas limit. |
| `contract_address` | [string](#string) |  | contract_address whose flat fee is considered when estimating tx fees. |






<a name="archway.rewards.v1beta1.QueryEstimateTxFeesResponse"></a>

### QueryEstimateTxFeesResponse
QueryEstimateTxFeesResponse is the response for Query.EstimateTxFees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_unit_price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | gas_unit_price defines the minimum transaction fee per gas unit. |
| `estimated_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | estimated_fee is the estimated transaction fee for a given gas limit. |






<a name="archway.rewards.v1beta1.QueryFlatFeeRequest"></a>

### QueryFlatFeeRequest
QueryFlatFeeRequest is the request for Query.FlatFeet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryFlatFeeResponse"></a>

### QueryFlatFeeResponse
QueryFlatFeeResponse is the response for Query.FlatFee


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `flat_fee_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee_amount defines the minimum flat fee set by the contract_owner per contract execution. |






<a name="archway.rewards.v1beta1.QueryOutstandingRewardsRequest"></a>

### QueryOutstandingRewardsRequest
QueryOutstandingRewardsRequest is the request for Query.OutstandingRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the target address to query calculated rewards for (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryOutstandingRewardsResponse"></a>

### QueryOutstandingRewardsResponse
QueryOutstandingRewardsResponse is the response for Query.OutstandingRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | total_rewards is the total rewards credited to the rewards_address. |
| `records_num` | [uint64](#uint64) |  | records_num is the total number of RewardsRecord objects stored for the rewards_address. |






<a name="archway.rewards.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request for Query.Params.






<a name="archway.rewards.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response for Query.Params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.rewards.v1beta1.Params) |  |  |






<a name="archway.rewards.v1beta1.QueryRewardsPoolRequest"></a>

### QueryRewardsPoolRequest
QueryRewardsPoolRequest is the request for Query.RewardsPool.






<a name="archway.rewards.v1beta1.QueryRewardsPoolResponse"></a>

### QueryRewardsPoolResponse
QueryRewardsPoolResponse is the response for Query.RewardsPool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `undistributed_funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | undistributed_funds are undistributed yet tokens (ready for withdrawal). |
| `treasury_funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | treasury_funds are treasury tokens available (no mechanism is available to withdraw ATM). Treasury tokens are collected on a block basis. Those tokens are unused block rewards. |






<a name="archway.rewards.v1beta1.QueryRewardsRecordsRequest"></a>

### QueryRewardsRecordsRequest
QueryRewardsRecordsRequest is the request for Query.RewardsRecords.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the target address to query records for (bech32 encoded). |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is an optional pagination options for the request. |






<a name="archway.rewards.v1beta1.QueryRewardsRecordsResponse"></a>

### QueryRewardsRecordsResponse
QueryRewardsRecordsResponse is the response for Query.RewardsRecords.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord) | repeated | records is the list of rewards records. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination is the pagination details in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.rewards.v1beta1.Query"></a>

### Query
Query service for the tracking module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#archway.rewards.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#archway.rewards.v1beta1.QueryParamsResponse) | Params returns module parameters. | GET|/archway/rewards/v1/params|
| `ContractMetadata` | [QueryContractMetadataRequest](#archway.rewards.v1beta1.QueryContractMetadataRequest) | [QueryContractMetadataResponse](#archway.rewards.v1beta1.QueryContractMetadataResponse) | ContractMetadata returns the contract rewards parameters (metadata). | GET|/archway/rewards/v1/contract_metadata|
| `BlockRewardsTracking` | [QueryBlockRewardsTrackingRequest](#archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest) | [QueryBlockRewardsTrackingResponse](#archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse) | BlockRewardsTracking returns block rewards tracking for the current block. | GET|/archway/rewards/v1/block_rewards_tracking|
| `RewardsPool` | [QueryRewardsPoolRequest](#archway.rewards.v1beta1.QueryRewardsPoolRequest) | [QueryRewardsPoolResponse](#archway.rewards.v1beta1.QueryRewardsPoolResponse) | RewardsPool returns the current undistributed rewards pool funds. | GET|/archway/rewards/v1/rewards_pool|
| `EstimateTxFees` | [QueryEstimateTxFeesRequest](#archway.rewards.v1beta1.QueryEstimateTxFeesRequest) | [QueryEstimateTxFeesResponse](#archway.rewards.v1beta1.QueryEstimateTxFeesResponse) | EstimateTxFees returns the estimated transaction fees for the given transaction gas limit using the minimum consensus fee value for the current block. | GET|/archway/rewards/v1/estimate_tx_fees|
| `RewardsRecords` | [QueryRewardsRecordsRequest](#archway.rewards.v1beta1.QueryRewardsRecordsRequest) | [QueryRewardsRecordsResponse](#archway.rewards.v1beta1.QueryRewardsRecordsResponse) | RewardsRecords returns the paginated list of RewardsRecord objects stored for the provided rewards_address. | GET|/archway/rewards/v1/rewards_records|
| `OutstandingRewards` | [QueryOutstandingRewardsRequest](#archway.rewards.v1beta1.QueryOutstandingRewardsRequest) | [QueryOutstandingRewardsResponse](#archway.rewards.v1beta1.QueryOutstandingRewardsResponse) | OutstandingRewards returns total rewards credited from different contracts for the provided rewards_address. | GET|/archway/rewards/v1/outstanding_rewards|
| `FlatFee` | [QueryFlatFeeRequest](#archway.rewards.v1beta1.QueryFlatFeeRequest) | [QueryFlatFeeResponse](#archway.rewards.v1beta1.QueryFlatFeeResponse) | FlatFee returns the flat fee set by the contract owner for the provided contract_address | GET|/archway/rewards/v1/flat_fee|

 <!-- end services -->



<a name="archway/rewards/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/tx.proto



<a name="archway.rewards.v1beta1.MsgSetContractMetadata"></a>

### MsgSetContractMetadata
MsgSetContractMetadata is the request for Msg.SetContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  | sender_address is the msg sender address (bech32 encoded). |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata is the contract metadata to set / update. If metadata exists, non-empty fields will be updated. |






<a name="archway.rewards.v1beta1.MsgSetContractMetadataResponse"></a>

### MsgSetContractMetadataResponse
MsgSetContractMetadataResponse is the response for Msg.SetContractMetadata.






<a name="archway.rewards.v1beta1.MsgSetFlatFee"></a>

### MsgSetFlatFee
MsgSetFlatFee is the request for Msg.SetFlatFee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  | sender_address is the msg sender address (bech32 encoded). |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |
| `flat_fee_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee_amount defines the minimum flat fee set by the contract_owner |






<a name="archway.rewards.v1beta1.MsgSetFlatFeeResponse"></a>

### MsgSetFlatFeeResponse
MsgSetFlatFeeResponse is the response for Msg.SetFlatFee.






<a name="archway.rewards.v1beta1.MsgWithdrawRewards"></a>

### MsgWithdrawRewards
MsgWithdrawRewards is the request for Msg.WithdrawRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the address to distribute rewards to (bech32 encoded). |
| `records_limit` | [MsgWithdrawRewards.RecordsLimit](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit) |  | records_limit defines the maximum number of RewardsRecord objects to process. If provided limit is 0, the default limit is used. |
| `record_ids` | [MsgWithdrawRewards.RecordIDs](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs) |  | record_ids defines specific RewardsRecord object IDs to process. |






<a name="archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs"></a>

### MsgWithdrawRewards.RecordIDs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ids` | [uint64](#uint64) | repeated |  |






<a name="archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit"></a>

### MsgWithdrawRewards.RecordsLimit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `limit` | [uint64](#uint64) |  |  |






<a name="archway.rewards.v1beta1.MsgWithdrawRewardsResponse"></a>

### MsgWithdrawRewardsResponse
MsgWithdrawRewardsResponse is the response for Msg.WithdrawRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records_num` | [uint64](#uint64) |  | records_num is the number of RewardsRecord objects processed. |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards are the total rewards transferred. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.rewards.v1beta1.Msg"></a>

### Msg
Msg defines the module messaging service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetContractMetadata` | [MsgSetContractMetadata](#archway.rewards.v1beta1.MsgSetContractMetadata) | [MsgSetContractMetadataResponse](#archway.rewards.v1beta1.MsgSetContractMetadataResponse) | SetContractMetadata creates or updates an existing contract metadata. Method is authorized to the contract owner (admin if no metadata exists). | |
| `WithdrawRewards` | [MsgWithdrawRewards](#archway.rewards.v1beta1.MsgWithdrawRewards) | [MsgWithdrawRewardsResponse](#archway.rewards.v1beta1.MsgWithdrawRewardsResponse) | WithdrawRewards performs collected rewards distribution. Rewards might be credited from multiple contracts (rewards_address must be set in the corresponding contract metadata). | |
| `SetFlatFee` | [MsgSetFlatFee](#archway.rewards.v1beta1.MsgSetFlatFee) | [MsgSetFlatFeeResponse](#archway.rewards.v1beta1.MsgSetFlatFeeResponse) | SetFlatFee sets or updates or removes the flat fee to interact with the contract Method is authorized to the contract owner. | |

 <!-- end services -->



<a name="archway/tracking/v1beta1/tracking.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/tracking.proto



<a name="archway.tracking.v1beta1.BlockTracking"></a>

### BlockTracking
BlockTracking is the tracking information for a block.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [TxTracking](#archway.tracking.v1beta1.TxTracking) | repeated | txs defines the list of transactions tracked in the block. |






<a name="archway.tracking.v1beta1.ContractOperationInfo"></a>

### ContractOperationInfo
ContractOperationInfo keeps a single contract operation gas consumption data.
Object is being created by the IngestGasRecord call from the wasmd.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id defines the unique operation ID. |
| `tx_id` | [uint64](#uint64) |  | tx_id defines a transaction ID operation relates to (TxInfo.id). |
| `contract_address` | [string](#string) |  | contract_address defines the contract address operation relates to. |
| `operation_type` | [ContractOperation](#archway.tracking.v1beta1.ContractOperation) |  | operation_type defines the gas consumption type. |
| `vm_gas` | [uint64](#uint64) |  | vm_gas is the gas consumption reported by the WASM VM. Value is adjusted by this module (CalculateUpdatedGas func). |
| `sdk_gas` | [uint64](#uint64) |  | sdk_gas is the gas consumption reported by the SDK gas meter and the WASM GasRegister (cost of Execute/Query/etc). Value is adjusted by this module (CalculateUpdatedGas func). |






<a name="archway.tracking.v1beta1.TxInfo"></a>

### TxInfo
TxInfo keeps a transaction gas tracking data.
Object is being created at the module EndBlocker.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id defines the unique transaction ID. |
| `height` | [int64](#int64) |  | height defines the block height of the transaction. |
| `total_gas` | [uint64](#uint64) |  | total_gas defines total gas consumption by the transaction. It is the sum of gas consumed by all contract operations (VM + SDK gas). |






<a name="archway.tracking.v1beta1.TxTracking"></a>

### TxTracking
TxTracking is the tracking information for a single transaction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [TxInfo](#archway.tracking.v1beta1.TxInfo) |  | info defines the transaction details. |
| `contract_operations` | [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo) | repeated | contract_operations defines the list of contract operations consumed by the transaction. |





 <!-- end messages -->


<a name="archway.tracking.v1beta1.ContractOperation"></a>

### ContractOperation
ContractOperation denotes which operation consumed gas.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTRACT_OPERATION_UNSPECIFIED | 0 | Invalid or unknown operation |
| CONTRACT_OPERATION_INSTANTIATION | 1 | Instantiate operation |
| CONTRACT_OPERATION_EXECUTION | 2 | Execute operation |
| CONTRACT_OPERATION_QUERY | 3 | Query |
| CONTRACT_OPERATION_MIGRATE | 4 | Migrate operation |
| CONTRACT_OPERATION_IBC | 5 | IBC operations |
| CONTRACT_OPERATION_SUDO | 6 | Sudo operation |
| CONTRACT_OPERATION_REPLY | 7 | Reply callback operation |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/tracking/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/genesis.proto



<a name="archway.tracking.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the initial state of the tracking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_info_last_id` | [uint64](#uint64) |  | tx_info_last_id defines the last unique ID for a TxInfo objs. |
| `tx_infos` | [TxInfo](#archway.tracking.v1beta1.TxInfo) | repeated | tx_infos defines a list of all the tracked transactions. |
| `contract_op_info_last_id` | [uint64](#uint64) |  | contract_op_info_last_id defines the last unique ID for ContractOperationInfo objs. |
| `contract_op_infos` | [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo) | repeated | contract_op_infos defines a list of all the tracked contract operations. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/tracking/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/query.proto



<a name="archway.tracking.v1beta1.QueryBlockGasTrackingRequest"></a>

### QueryBlockGasTrackingRequest
QueryBlockGasTrackingRequest is the request for Query.BlockGasTracking.






<a name="archway.tracking.v1beta1.QueryBlockGasTrackingResponse"></a>

### QueryBlockGasTrackingResponse
QueryBlockGasTrackingResponse is the response for Query.BlockGasTracking.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockTracking](#archway.tracking.v1beta1.BlockTracking) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.tracking.v1beta1.Query"></a>

### Query
Query service for the tracking module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `BlockGasTracking` | [QueryBlockGasTrackingRequest](#archway.tracking.v1beta1.QueryBlockGasTrackingRequest) | [QueryBlockGasTrackingResponse](#archway.tracking.v1beta1.QueryBlockGasTrackingResponse) | BlockGasTracking returns block gas tracking for the current block | GET|/archway/tracking/v1/block_gas_tracking|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

