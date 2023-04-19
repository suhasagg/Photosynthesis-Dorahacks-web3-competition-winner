use cosmwasm_std::{
    Addr, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Storage, to_binary,
    from_binary, Uint128, WasmMsg,CodeInfoResponse,ContractInfoResponse,Env
};

use cosmwasm_std::{
    entry_point, instantiate, to_binary, Attribute, Binary, CodeInfoResponse, ContractInfoResponse,
    DepsMut, Env, MessageInfo, Response, StdResult, WasmMsg,
};

use crate::errors::ContractError;
use crate::msg::{ExecuteMsg, InstantiateMsg};

use cosmwasm_bignumber::{Decimal256};
use cw_storage_plus::Map;
use serde::{Deserialize, Serialize};
use schemars::JsonSchema;

use std::convert::TryInto;

use cw_epoch::{Duration, EpochInfo, EpochManager, Schedule};

const PREFIX_REDEMPTION_RECORD: &[u8] = b"redemption_record";
const PREFIX_CONTRACT_KEY: &[u8] = b"contract";
const PREFIX_STAKE: &[u8] = b"stake";
const PREFIX_REDIRECTION: &[u8] = b"redirection";

const REDEMPTION_RATE_QUERY_EPOCH: &str = "redemption_rate_query";
const DEFAULT_REDEMPTION_RATE_QUERY_INTERVAL: u64 = 2;
const DEFAULT_REDEMPTION_RATE_THRESHOLD: u128 = 100;
const DEFAULT_REDEMPTION_INTERVAL_THRESHOLD: u64 = 10;
use cosmwasm_std::{
    entry_point, to_binary, Binary, Env, InitResponse, MessageInfo, StdResult,
};
use cosmwasm_bignumber::Uint256;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {
    pub owner: String,
    pub total_supply: Uint256,
    pub name: String,
    pub symbol: String,
    pub decimals: u8,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub owner: String,
    pub total_supply: Uint256,
    pub name: String,
    pub symbol: String,
    pub decimals: u8,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum HandleMsg {}
/*
#[entry_point]
pub fn init(_ctx: Env, _msg: InitMsg) -> StdResult<InitResponse<State>> {
    let state = State {
        owner: _msg.owner,
        total_supply: _msg.total_supply,
        name: _msg.name,
        symbol: _msg.symbol,
        decimals: _msg.decimals,
    };

    Ok(InitResponse::new(state))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn proper_initialization() {
        let msg = InitMsg {
            owner: "admin".to_string(),
            total_supply: Uint256::new(1000),
            name: "Test Token".to_string(),
            symbol: "TEST".to_string(),
            decimals: 18,
        };
        let env = Env {
            block: None,
            contract: Default::default(),
            message: Default::default(),
            now: Default::default(),
        };
        let res = init(env, msg).unwrap();
        assert_eq!(res.messages.len(), 0);
    }
}
*/

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct Contract {
    pub address: Addr,
    pub stake: Uint128,
    pub rewards: Uint128,
    pub enable_liquid_staking: bool,
    pub liquid_stake_interval: u64,
    pub rewards_withdrawal_interval: u64,
    pub rewards_to_liquid_stake: Uint128,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct ContractList {
    pub list: Vec<Contract>,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct RedemptionRecord {
    pub timestamp: u64,
    pub liquidity_amount: Uint128,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct RedemptionRateQueryParams {}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub enum RedemptionRateQueryMsg {
    QueryRedemptionRate {
        chain_id: String,
        query_params: RedemptionRateQueryParams,
    },
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct ContractState {
    pub contracts: ContractList,
    pub latest_redemption_record: Option<RedemptionRecord>,
    pub redemption_rate_query_interval: u64,
    pub redemption_rate_threshold: Uint128,
    pub redemption_interval_threshold: u64,
}

impl ContractList {
    pub fn update(&mut self, contract: &Contract) {
        let index = self
            .list
            .iter()
            .position(|c| c.address == contract.address)
            .unwrap();

        self.list[index] = contract.clone();
    }

    pub fn delete(&mut self, address: &Addr) {
        let index = self
            .list
            .iter()
            .position(|c| c.address == *address)
            .unwrap();

        self.list.remove(index);
    }
}

fn get_contract_state(storage: &dyn Storage) -> Result<ContractState, ContractError> {
    cw_storage_plus::singleton_read(storage, b"state")
        .load()
        .map_err(|_| ContractError::Std(StdError::generic_err(
            "Failed to load state",
        )))
}


fn save_contract_state(
    storage: &mut dyn Storage,
    state: &ContractState,
) -> Result<(), ContractError> {
    cw_storage_plus::singleton_save(storage, b"state", state)
        .map_err(|_| ContractError::Std(StdError::generic_err("Failed to save state")))
}

// Define the contract error
#[derive(Debug)]
pub enum ContractError {
    Std(StdError),
    ContractNotFound,
    RedemptionIntervalThresholdNotReached,
}

impl From<StdError> for ContractError {
    fn from(error: StdError) -> Self {
        ContractError::Std(error)
    }
}

// Here you can define the remaining functions and entry points for your contract
// For example, `add_stake`, `redeem_liquid_tokens`, and `update_contract`

// Other helper functions can be added as needed to implement the contract's logic
/*
#[entry_point]
fn process_redemption_rate_queries(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> Result<Response, ContractError> {
    let epoch_info = EpochInfo {
        identifier: REDEMPTION_RATE_QUERY_EPOCH,
        number: env.block.height,
    };

    // You will need to implement the 'process_redemption_rate_queries' function
    process_redemption_rate_queries(deps, env, epoch_info)?;

    Ok(Response::new().add_attribute("action", "process_redemption_rate_queries"))
}

#[entry_point]
fn add_stake(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    #[serde(deserialize_with = "from_binary")] params: AddStakeParams,
) -> Result<Response, ContractError> {
    let staker_address_raw = info.sender.clone();
    let contract_address_raw = deps.api.canonical_address(&params.contract_address)?;
    let stake_amount = params.stake_amount;

    // You will need to implement the 'add_stake' function
    add_stake(deps, env, staker_address_raw, contract_address_raw, stake_amount)?;

    Ok(Response::new().add_attribute("action", "add_stake"))
}

#[entry_point]
fn redeem_liquid_tokens(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> Result<Response, ContractError> {
    let redemption_interval_threshold = get_redemption_interval_threshold(deps.storage)?;
    let time_since_latest_redemption = env.block.time - get_latest_redemption_time(deps.storage)?;

    if time_since_latest_redemption < redemption_interval_threshold {
        return Err(ContractError::RedemptionIntervalThresholdNotReached);
    }

    let cum_liquidity_amount = get_cumulative_liquidity_amount(deps.storage)?;

    // You will need to implement the 'redeem_liquid_tokens' and 'distribute_redeemed_tokens' functions
    redeem_liquid_tokens(deps.as_ref(), cum_liquidity_amount)?;
    distribute_redeemed_tokens(deps.as_ref(), cum_liquidity_amount)?;
    delete_redemption_record(deps.storage)?;

    Ok(Response::new().add_attribute("action", "redeem_liquid_tokens"))
}

#[entry_point]
fn update_contract(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    #[serde(deserialize_with = "from_binary")] contract: Contract,
) -> Result<Response, ContractError> {
    // Ensure that the contract exists
    let contract_info = get_contract(deps.storage, &contract.address)?;

    // Update the contract state
    // You will need to implement the 'update_contract' function
    update_contract(deps, &contract)?;

    Ok(Response::new().add_attribute("action", "update_contract"))
}

// Implement other entry points and helper functions as needed to complete the contract's logic
// Entry point for deleting a contract
#[entry_point]
fn delete_contract(
deps: DepsMut,
env: Env,
info: MessageInfo,
#[serde(deserialize_with = "from_binary")] address: CanonicalAddr,
) -> Result<Response, ContractError> {
// Ensure that the contract exists
let contract_info = load_contract(deps.storage, &address)?;


delete_contract(deps.storage, &address)?;

Ok(Response::new().add_attribute("action", "delete_contract"))
}
*/


// Get the epoch period in seconds from storage
fn get_epoch_period_seconds(storage: &dyn Storage) -> Result<u64, ContractError> {
let key = PREFIX_EPOCH_PERIOD_SECONDS_KEY.mixed_radix_key()?;
let value = storage.get(&key).ok_or(ContractError::NotFound)?;


from_slice(&value).map_err(|_| ContractError::Std(StdError::generic_err(
    "Error parsing epoch period seconds",
)))
}

// Get the redemption rate query interval from storage
fn get_redemption_rate_query_interval(storage: &dyn Storage) -> Result<u64, ContractError> {
let key = PREFIX_REDEMPTION_RATE_QUERY_INTERVAL_KEY.mixed_radix_key()?;
let value = storage.get(&key).ok_or(ContractError::NotFound)?;


from_slice(&value).map_err(|_| ContractError::Std(StdError::generic_err(
    "Error parsing redemption rate query interval",
)))
}

// Check if an epoch starts for a given identifier and block
fn is_epoch_start(storage: &dyn Storage, identifier: &str, block: &BlockInfo) -> bool {
let epoch_period_seconds = match get_epoch_period_seconds(storage) {
Ok(v) => v,
Err(_) => return false,
};


let start_time = block.time.seconds();
start_time % epoch_period_seconds == 0 && identifier == REDEMPTION_RATE_QUERY_EPOCH
}

// Query the redemption rate
fn query_redemption_rate(_storage: &dyn Storage) -> Result<Decimal, ContractError> {
// TODO: Implement the logic to query redemption rate from an oracle or other source


Ok(Decimal::one()) // For now, return a default value of 1.0
}

// Get a parameter by key
fn get_param(storage: &dyn Storage, param_key: &str) -> Result<Decimal, ContractError> {
let key = PREFIX_PARAM_KEY.mixed_radix_key(param_key.as_bytes())?;
let value = storage.get(&key).ok_or(ContractError::NotFound)?;

from_slice(&value).map_err(|_| ContractError::Std(StdError::generic_err(
    "Error parsing parameter",
)))
}

// Update the redemption rate threshold
fn update_redemption_rate_threshold(
storage: &mut dyn Storage,
redemption_rate: Decimal,
) -> Result<(), ContractError> {
let key = PREFIX_PARAM_KEY.mixed_radix_key(b"redemption_rate_threshold")?;
storage.set(&key, &to_vec(&redemption_rate)?);


Ok(())
}

// Get all contracts from storage
fn get_all_contracts(storage: &dyn Storage) -> Result<Vec<Contract>, ContractError> {
list_contracts(storage).map(|contract_list| contract_list.list)
}

// Get the cumulative reward amount for a contract
fn get_cumulative_reward_amount(
_storage: &dyn Storage,
_address: &CanonicalAddr,
) -> Result<Uint256, ContractError> {
// TODO: Implement the logic to get the cumulative reward amount for a contract


Ok(Uint256::zero()) // For now, return a default value of 0
}


use cosmwasm_std::{
    StdResult, Storage, BlockInfo, Env, DepsMut,
    Order, from_slice, HumanAddr,
};
use cosmwasm_bignumber::Uint256;
use cosmwasm_storage::{
    PrefixedStorage, ReadonlyPrefixedStorage, ReadonlySingleton,
    Singleton, ReadonlyBucket, Bucket,
};
use crate::types::{Contract, DepositRecord, EpochState, EpochInfo};
use crate::error::ContractError;
use crate::state::{
    get_param, read_balance, add_balance, sub_balance, get_contract_state,
    save_contract_state, get_all_contracts, read_contract,
    get_contract_liquid_stake_deposit_records_for_epoch, enqueue_liquid_stake_record,
    get_cumulative_reward_amount, get_latest_redemption_time, set_latest_redemption_time,
    get_epoch_period_seconds, get_redemption_rate_query_interval, query_redemption_rate,
    update_redemption_rate_threshold, withdraw_rewards, distribute_rewards,
    is_epoch_start, get_contract_liquid_stake_deposit_records_key,
};

const PREFIX_CONTRACT_KEY: &[u8] = b"contract";

pub struct ContractList {
    pub list: Vec<Contract>,
}

// List all contracts in the store
fn list_contracts(storage: &dyn Storage) -> Result<ContractList, ContractError> {
    let contracts: Vec<Contract> = storage
        .range(
            PREFIX_CONTRACT_KEY,
            None,
            None,
            Order::Ascending,
        )
        .filter_map(|item| {
            let (_, value) = item.map_err(ContractError::Std)?;
            from_slice::<Contract>(&value).ok()
        })
        .collect();

    Ok(ContractList { list: contracts })
}

// Schedule the redemption rate queries using epoch-based processing
fn schedule_redemption_rate_queries(
    deps: DepsMut,
    env: Env,
) -> Result<(), ContractError> {
    let epoch_period_seconds = get_epoch_period_seconds(deps.storage)?;
    let redemption_rate_query_interval = get_redemption_rate_query_interval(deps.storage)?;

    let now = env.block.time.seconds();
    let next_epoch_start_time = now - (now % epoch_period_seconds) + epoch_period_seconds
        + redemption_rate_query_interval;
    let next_epoch_id = (next_epoch_start_time / epoch_period_seconds) as u64;

    let mut state = get_contract_state(deps.storage)?;
    let mut next_epoch = state.next_epoch.unwrap_or_else(|| EpochState {
        id: next_epoch_id,
        identifier: REDEMPTION_RATE_QUERY_EPOCH.to_string(),
        start_time: next_epoch_start_time,
        end_time: next_epoch_start_time + epoch_period_seconds,
    });

    while next_epoch.id <= next_epoch_id {
        process_redemption_rate_queries(
            deps.branch(),
            env.clone(),
            EpochInfo {
                identifier: next_epoch.identifier.clone(),
                number: next_epoch.id,
            },
        )?;

        // Update next_epoch
        next_epoch.start_time += epoch_period_seconds;
        next_epoch.end_time += epoch_period_seconds;
        next_epoch.id += 1;
    }

    state.next_epoch = Some(next_epoch);
    save_contract_state(deps.storage, &state)?;

    Ok(())
}

// Photosynthesis module
const LIQUID_STAKING_DAPP_REWARDS_EPOCH: &str = "LIQUID_STAKING_DAPP_REWARDS_EPOCH";
const REDEMPTION_RATE_QUERY_EPOCH: &str = "REDEMPTION_RATE_QUERY_EPOCH";

use cosmwasm_std::{
    StdResult, BlockInfo, Storage, HumanAddr,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage, AppendStoreMut};
use crate::keeper::Keeper;
use crate::types::{self, DepositRecord};

const LIQUID_STAKING_DAPP_REWARDS_EPOCH: &str = "LIQUID_STAKING_DAPP_REWARDS_EPOCH";
const REDEMPTION_RATE_QUERY_EPOCH: &str = "REDEMPTION_RATE_QUERY_EPOCH";
const REWARDS_WITHDRAWAL_EPOCH: &str = "REWARDS_WITHDRAWAL_EPOCH";

pub fn begin_block<S>(block: BlockInfo, storage: &mut S) -> StdResult<()>
where
    S: Storage,
{
    let mut k = Keeper::new(&mut PrefixedStorage::new(storage, b"epoch::"));
    let ctx = k.create_context();

    // Process liquid staking deposits for contracts with enabled liquid staking
    for contract in k.get_all_contracts(&ctx) {
        if contract.enable_liquid_staking {
            let epoch_number = block.height / contract.liquid_stake_interval;
            let reward_amount = k.get_cumulative_reward_amount(&ctx, contract.address)?;
            if reward_amount >= contract.rewards_to_liquid_stake {
                let deposit_records = k.create_contract_liquid_stake_deposit_records_for_epoch(
                    &ctx,
                    &contract.address,
                    epoch_number,
                );
                k.enqueue_liquid_stake_record(&ctx, &contract.address, deposit_records)?;
            }
        }
    }

    // Process redemption rate query and update redemption rate threshold if necessary
    if k.is_epoch_start(&ctx, REDEMPTION_RATE_QUERY_EPOCH, &block) {
        for contract in k.get_all_contracts(&ctx) {
            if let Some(redemption_address) = contract.redemption_address {
                let redemption_rate = k.query_redemption_rate(&ctx, &redemption_address)?;
                if redemption_rate > k.get_param(&ctx, types::RedemptionRateThreshold)? {
                    let redemption_interval = k.get_param(&ctx, types::RedemptionIntervalThreshold)?;
                    let time_since_latest_redemption = k.get_time_since_latest_redemption(
                        &ctx,
                        &contract.address,
                        block.height / redemption_interval,
                        block.time,
                    )?;
                    if time_since_latest_redemption >= redemption_interval {
                        let deposit_records = k.get_contract_liquid_stake_deposit_records_for_epoch(
                            &ctx,
                            &contract.address,
                            block.height / redemption_interval,
                        )?;
                        k.redeem_and_distribute(
                            &ctx,
                            &contract.address,
                            &redemption_address,
                            redemption_rate,
                            deposit_records,
                        )?;
                        k.set_latest_redemption_time(
                            &ctx,
                            &contract.address,
                            block.height / redemption_interval,
                            block.time,
                        )?;
                    }
                }
            }
        }
    }

    // Distribute rewards to contracts with enabled rewards withdrawal
    for contract in k.get_all_contracts(&ctx) {
        if contract.rewards_with


use cosmwasm_std::{
StdResult, BlockInfo, Storage, HumanAddr,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage};
use crate::keeper::Keeper;
use crate::types::{self, DepositRecord, WithdrawalRecord};

pub fn end_block<S>(_block: BlockInfo, storage: &mut S) -> StdResult<()>
where
S: Storage,
{
let mut k = Keeper::new(&mut PrefixedStorage::new(storage, b"epoch::"));
let ctx = k.create_context();

// Process liquid stake deposits
let liquid_stake_interval = k.get_param(&ctx, types::KeyArchLiquidStakeInterval)?;
if (_block.height + 1) % liquid_stake_interval == 0 {
    let epoch_number = _block.height / liquid_stake_interval;
    for contract in k.get_all_contracts(&ctx) {
        if contract.enable_liquid_staking {
            if epoch_number % contract.liquid_stake_interval == 0 {
                let reward_amount = k.get_cumulative_reward_amount(&ctx, contract.address)?;
                if reward_amount >= contract.rewards_to_liquid_stake {
                    let deposit_records = k.create_contract_liquid_stake_deposit_records_for_epoch(
                        &ctx,
                        &contract.address,
                        epoch_number,
                    );
                    k.enqueue_liquid_stake_record(&ctx, &contract.address, deposit_records)?;
                }
            }
        }
    }
}

// Process redemption rate query
let redemption_rate_interval = k.get_param(&ctx, types::RedemptionRateQueryInterval)?;
if (_block.height + 1) % redemption_rate_interval == 0 {
    let epoch_number = _block.height / redemption_rate_interval;
    for contract in k.get_all_contracts(&ctx) {
        if let Some(redemption_address) = contract.redemption_address {
            let redemption_rate = k.query_redemption_rate(&ctx, &redemption_address)?;
            let redemption_rate_threshold = k.get_param(&ctx, types::RedemptionRateThreshold)?;
            if redemption_rate > redemption_rate_threshold {
                let redemption_interval_threshold =
                    k.get_param(&ctx, types::RedemptionIntervalThreshold)?;
                let time_since_latest_redemption = k.get_time_since_latest_redemption(
                    &ctx,
                    &contract.address,
                    epoch_number,
                )?;
                if time_since_latest_redemption >= redemption_interval_threshold {
                    let deposit_records = k.get_contract_liquid_stake_deposit_records_for_epoch(
                        &ctx,
                        &contract.address,
                        epoch_number,
                    )?;
                    k.redeem_and_distribute(
                        &ctx,
                        &contract.address,
                        &redemption_address,
                        redemption_rate,
                        deposit_records,
                    )?;
                    k.set_latest_redemption_time(
                        &ctx,
                        &contract.address,
                        epoch_number,
                        _block.time,
                    )?;
                }
            }
        }
    }
}

// Process rewards withdrawal
let rewards_withdrawal_interval = k.get_param(&ctx, types::RewardsWithdrawalInterval)?;
if (_block.height + 1) % rewards_withdrawal_interval == 0 {
    for contract in k.get_all_contracts(&ctx) {
        if contract.rewards_withdrawal_interval > 0 {
            let epoch_number = _block.height / contract.rewards_withdrawal_interval;
            let rewards = k.withdraw_rewards(&ctx, &contract.address)?;
            if rewards > 0 {
                k.distribute_rewards(&ctx, &contract.address, rewards)?;
            }
        }
    }
}

Ok(())
}

pub fn enqueue_liquid_stake_record<S>(
ctx: &dyn Context<S>,
contract_address: &HumanAddr,
records: Vec<DepositRecord>,
) -> StdResult<()>
where
S: Storage,
{
let mut prefix_storage = PrefixedStorage::new(ctx.storage(), b"epoch::");
let mut store = AppendStoreMut::attach_or_create(&mut prefix_storage, contract_address.as_bytes())?;
let bytes = to_vec(&records)?;
store.push(bytes.as_slice())?;
Ok(())
}

pub fn get_all_contracts<S>(ctx: &dyn Context<S>) -> Vec<ContractInfo>
where
S: Storage,
{
ReadonlyPrefixedStorage::new(ctx.storage(), b"contract")
.range(None, None, Order::Ascending)
.filter_map(|item| {
if let Ok(contract_info) = from_slice::<ContractInfo>(&item.1) {
Some(contract_info)
} else {
None
}
})
.collect()
}

pub fn get_cumulative_reward_amount<S>(
ctx: &dyn Context<S>,
contract_address: HumanAddr,
) -> StdResult<Uint256>
where
S: Storage,
{
let prefix_storage = ReadonlyPrefixedStorage::new(ctx.storage(), b"contract_reward_cumulative");
let cumulative_reward_key = get_contract_reward_cumulative_key(&contract_address);
prefix_storage
.get(&cumulative_reward_key)
.unwrap_or_else(|| Ok(to_vec(&Uint256::zero()).unwrap()))
.and_then(|bytes| from_slice(&bytes))
}

pub fn create_contract_liquid_stake_deposit_records_for_epoch<S>(
ctx: &dyn Context<S>,
contract_address: &HumanAddr,
epoch_number: u64,
) -> Vec<DepositRecord>
where
S: Storage,
{
let mut records = vec![];
let prefix_storage = ReadonlyPrefixedStorage::new(ctx.storage(), b"epoch::");
let store = AppendStore::attach(&prefix_storage, contract_address.as_bytes()).unwrap();
for i in 0..store.len() {
let bytes = store.get(i).unwrap();
if let Ok(deposit_records) = from_slice::<Vec<DepositRecord>>(bytes) {
for record in deposit_records {
if record.epoch == epoch_number && record.status == "pending" {
records.push(record);
}
}
}
}
records
}

pub fn query_redemption_rate<S>(ctx: &dyn Context<S>, redemption_address: &HumanAddr) -> StdResult<Uint256>
where
S: Storage,
{
let query_msg = QueryMsg::RedemptionRate {};
let res: RedemptionRateResponse = query(&ctx, redemption_address, &query_msg)?;
Ok(res.redemption_rate)
}

pub fn withdraw_rewards<S>(ctx: &dyn Context<S>, contract_address: &HumanAddr) -> StdResult<Uint256>
where
S: Storage,
{
let mut prefix_storage = PrefixedStorage::new(ctx.storage(), b"contract_reward_balance");
let reward_balance_key = get_contract_reward_balance_key(&contract_address);
let rewards = prefix_storage
.get(&reward_balance_key)
.unwrap_or_else(|| Ok(to_vec(&Uint256::zero()).unwrap()))
.and_then(|bytes| from_slice(&bytes))?;
prefix_storage.remove(&reward_balance_key);
Ok(rewards)
}

pub fn enqueue_liquid_stake_record<S>(
    ctx: &dyn Context<S>,
    contract_address: &HumanAddr,
    records: Vec<DepositRecord>,
) -> StdResult<()>
where
    S: Storage,
{
    let mut prefix_storage = PrefixedStorage::new(ctx.storage(), b"epoch::");
    let mut store = AppendStoreMut::attach_or_create(&mut prefix_storage, contract_address.as_bytes())?;
    let bytes = to_vec(&records)?;
    store.push(bytes.as_slice())?;
    Ok(())
}

pub fn distribute_rewards<S>(
    ctx: &dyn Context<S>,
    contract_address: &HumanAddr,
    rewards: Uint256,
) -> StdResult<()>
where
    S: Storage,
{
    let mut prefix_storage = PrefixedStorage::new(ctx.storage(), b"contracts::");
    let mut storage = PrefixedStorage::new(&mut prefix_storage, contract_address.as_bytes());
    let mut contract = read_contract(&storage)?;
    let total_stake = contract.get_total_stake();
    if total_stake.is_zero() {
        return Err(StdError::generic_err("No liquidity provider stake"));
    }
    for (address, stake) in &contract.staked_amount {
        let reward_amount = (rewards * stake) / total_stake;
        add_reward(&mut storage, &address, reward_amount)?;
    }
    Ok(())
}

use cosmwasm_std::{DepsMut, Env, MessageInfo, Response, StdError, StdResult};

const REDEMPTION_RATE_QUERY_EPOCH: &str = "redemption_rate_query";

//#[entry_point]
fn process_redemption_rate_queries(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> Result<Response, StdError> {
    let epoch_info = EpochInfo {
        identifier: REDEMPTION_RATE_QUERY_EPOCH.to_string(),
        number: env.block.height,
    };
    
    // Query redemption rate and perform liquidation if needed
    let redemption_rate_query_interval = get_redemption_rate_query_interval(deps.storage)?;
    let latest_redemption_rate_query_height = get_latest_redemption_rate_query_height(deps.storage)?;
    let block_time = env.block.time;
    
    if epoch_info.number < latest_redemption_rate_query_height + redemption_rate_query_interval {
        return Ok(Response::new());
    }
    
    let redemption_rate = query_redemption_rate(deps.storage)?;
    let redemption_rate_threshold = get_redemption_rate_threshold(deps.storage)?;
    if redemption_rate > redemption_rate_threshold {
        let redemption_interval_threshold = get_redemption_interval_threshold(deps.storage)?;
        let latest_redemption_time = get_latest_redemption_time(deps.storage)?;
        
        if block_time - latest_redemption_time >= Duration::from_seconds(redemption_interval_threshold) {
            let cum_liquidity_amount = get_cumulative_liquidity_amount(deps.storage)?;
            let total_stake = get_total_stake(deps.storage)?;
            let redeemed_amount = cum_liquidity_amount.multiply_ratio(total_stake, Decimal256::from_uint256(Uint256::new(1)));
            
            liquidate_and_distribute_rewards(deps.as_ref(), redeemed_amount)?;
            delete_redemption_record(deps.storage)?;
            
            save_latest_redemption_time(deps.storage, block_time)?;
        }
    }
    
    // Save the latest redemption rate query height
    save_latest_redemption_rate_query_height(deps.storage, epoch_info.number)?;
    
    Ok(Response::new())
}

use cosmwasm_std::{StdError, Storage, ReadonlyStorage};
use cosmwasm_bignumber::{Decimal256, Uint256};
use crate::ContractState;
use crate::ContractError;

const PREFIX_REDEMPTION_RATE_QUERY_INTERVAL: &[u8] = b"redemption_rate_query_interval";
const PREFIX_LATEST_REDEMPTION_RATE_QUERY_HEIGHT: &[u8] = b"latest_redemption_rate_query_height";
const PREFIX_REDEMPTION_RATE_THRESHOLD: &[u8] = b"redemption_rate_threshold";
const PREFIX_REDEMPTION_INTERVAL_THRESHOLD: &[u8] = b"redemption_interval_threshold";
const PREFIX_LATEST_REDEMPTION_TIME: &[u8] = b"latest_redemption_time";
const PREFIX_CUMULATIVE_LIQUIDITY_AMOUNT: &[u8] = b"cumulative_liquidity_amount";
const PREFIX_TOTAL_STAKE: &[u8] = b"total_stake";

// Get the redemption rate query interval from storage
fn get_redemption_rate_query_interval(storage: &dyn Storage) -> Result<u64, ContractError> {
    let value = storage.get(PREFIX_REDEMPTION_RATE_QUERY_INTERVAL).ok_or_else(|| {
        ContractError::Std(StdError::not_found("Redemption rate query interval not found"))
    })?;

    Ok(Uint256::from_be_bytes(value).as_u64())
}

// Get the latest redemption rate query height from storage
fn get_latest_redemption_rate_query_height(storage: &dyn Storage) -> Result<u64, ContractError> {
    let value = storage.get(PREFIX_LATEST_REDEMPTION_RATE_QUERY_HEIGHT).ok_or_else(|| {
        ContractError::Std(StdError::not_found("Latest redemption rate query height not found"))
    })?;

    Ok(Uint256::from_be_bytes(value).as_u64())
}

// Query the redemption rate
fn query_redemption_rate(storage: &dyn Storage) -> Result<Decimal256, ContractError> {
    // TODO: Implement the logic to query redemption rate from an oracle or other source
    // For now, we will return a static value
    let redemption_rate = Decimal256::from_ratio(1u64, 100u64);
    Ok(redemption_rate)
}

// Get the redemption rate threshold from storage
fn get_redemption_rate_threshold(storage: &dyn Storage) -> Result<Uint256, ContractError> {
    let value = storage.get(PREFIX_REDEMPTION_RATE_THRESHOLD).unwrap_or_default();

    Ok(Uint256::from_be_bytes(value))
}

// Get the redemption interval threshold from storage
fn get_redemption_interval_threshold(storage: &dyn Storage) -> Result<u64, ContractError> {
    let value = storage.get(PREFIX_REDEMPTION_INTERVAL_THRESHOLD).unwrap_or_default();

    Ok(Uint256::from_be_bytes(value).as_u64())
}

// Get the latest redemption time from storage
fn get_latest_redemption_time(storage: &dyn Storage) -> Result<u64, ContractError> {
    let value = storage.get(PREFIX_LATEST_REDEMPTION_TIME).unwrap_or_default();

    Ok(Uint256::from_be_bytes(value).as_u64())
}

// Get the cumulative liquidity amount from storage
fn get_cumulative_liquidity_amount(storage: &dyn Storage) -> Result<Uint256, ContractError> {
    let value = storage.get(PREFIX_CUMULATIVE_LIQUIDITY_AMOUNT).unwrap_or_default();

    Ok(Uint256::from_be_bytes(value))
}


// Get the total stake from storage
fn get_total_stake(storage: &dyn Storage) -> Result<Uint256, ContractError> {
    let value = storage
.get(PREFIX_TOTAL_STAKE)
.unwrap_or_else(|| to_binary(&Uint256::zero()).unwrap());

Ok(from_binary(&value)?)
}

// Liquidate the rewards and distribute them to the Dapps according to their stake
fn liquidate_and_distribute_rewards(
storage: &mut dyn Storage,
total_stake: Uint256,
rewards_to_distribute: Uint256,
) -> Result<(), ContractError> {
let mut contract_list = get_contract_list(storage)?;


for contract in contract_list.iter_mut() {
    let rewards_to_distribute_for_contract =
        rewards_to_distribute * contract.stake / total_stake;

    if rewards_to_distribute_for_contract > Uint256::zero() {
        // Send the rewards to the Dapp
        let message = WasmMsg::Execute {
            contract_addr: contract.address.to_string(),
            msg: to_binary(&RedeemRewardsMsg {
                rewards: rewards_to_distribute_for_contract,
            })?,
            funds: vec![],
        };

        let messages = vec![message];
        let attrs = vec![            attr("action", "distribute_rewards"),            attr("amount", &rewards_to_distribute_for_contract.to_string()),        ];
        let response = Response::new().add_messages(messages).add_attributes(attrs);

        contract.rewards += rewards_to_distribute_for_contract;
        save_contract(storage, &contract)?;

        return Ok(response);
    }
}

Ok(Response::new())
}

// Delete the latest redemption record from storage
fn delete_redemption_record(storage: &mut dyn Storage) -> Result<(), ContractError> {
storage.remove(PREFIX_LATEST_REDEMPTION_RECORD);

Ok(())
}

// Save the latest redemption time to storage
fn save_latest_redemption_time(
storage: &mut dyn Storage,
latest_redemption_time: u64,
) -> Result<(), ContractError> {
storage.set(
PREFIX_LATEST_REDEMPTION_TIME,
&to_binary(&latest_redemption_time)?,
);

Ok(())
}

// Save the latest redemption rate query height to storage
fn save_latest_redemption_rate_query_height(
storage: &mut dyn Storage,
latest_redemption_rate_query_height: u64,
) -> Result<(), ContractError> {
storage.set(
PREFIX_LATEST_REDEMPTION_RATE_QUERY_HEIGHT,
&to_binary(&latest_redemption_rate_query_height)?,
);


Ok(())
}




