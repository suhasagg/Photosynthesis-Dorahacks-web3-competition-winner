use cosmwasm_std::{BankMsg, Coin, DepsMut, Env, MessageInfo, Response, StdResult, Storage, Uint128};
use cosmwasm_storage::{singleton, singleton_read, Singleton, ReadonlySingleton};
use std::collections::HashMap;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct StakeInfo {
    pub amount: Uint128,
    pub last_updated_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct RewardInfo {
    pub total_rewards: Uint128,
    pub last_claim_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub enum ExecuteMsg {
    StakeTokens { amount: Uint128 },
    WithdrawRewards {},
    UpdateEpoch {},
    DistributeLiquidTokens {},
    DistributeRedemptionAmounts {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct State {
    pub current_epoch: u64,
    pub staking_info: HashMap<String, StakeInfo>,
    pub reward_info: HashMap<String, RewardInfo>,
    pub reward_records: HashMap<String, Vec<RewardRecord>>,
    pub total_liquid_tokens: Uint128,
    pub total_redemption_tokens: Uint128,
    pub contract_addresses: Vec<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct RewardRecord {
    pub rewards: Uint128,
    pub calculated_time: String,
    pub calculated_height: u64,
}

// Helper functions for state management
fn state(storage: &mut dyn Storage) -> Singleton<State> {
    singleton(storage, b"state")
}

fn state_read(storage: &dyn Storage) -> ReadonlySingleton<State> {
    singleton_read(storage, b"state")
}

fn increment_epoch(deps: DepsMut) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;
    state.current_epoch += 1;
    state(deps.storage).save(&state)?;
    Ok(Response::new().add_attribute("method", "increment_epoch"))
}

// Handling staking
fn stake_tokens(deps: DepsMut, info: MessageInfo, amount: Uint128) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;
    let staker = info.sender.to_string();
    let stake_info = state.staking_info.entry(staker.clone()).or_insert_with(|| StakeInfo {
        amount: Uint128::zero(),
        last_updated_epoch: state.current_epoch,
    });

    stake_info.amount += amount;
    state(deps.storage).save(&state)?;

    Ok(Response::new()
        .add_message(BankMsg::Send {
            from_address: info.sender.to_string(),
            to_address: "staking_pool".to_string(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        })
        .add_attribute("action", "stake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount", amount.to_string()))
}

// Handling reward withdrawal
fn withdraw_rewards(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;
    let staker = info.sender.to_string();
    let reward_info = state.reward_info.entry(staker.clone()).or_insert_with(|| RewardInfo {
        total_rewards: Uint128::zero(),
        last_claim_epoch: state.current_epoch,
    });

    let rewards_to_withdraw = reward_info.total_rewards;
    reward_info.total_rewards = Uint128::zero(); // Reset rewards after withdrawal
    state(deps.storage).save(&state)?;

    Ok(Response::new()
        .add_message(BankMsg::Send {
            from_address: "reward_pool".to_string(),
            to_address: staker,
            amount: vec![Coin { denom: "token".to_string(), amount: rewards_to_withdraw }],
        })
        .add_attribute("action", "withdraw_rewards")
        .add_attribute("staker", info.sender.to_string())
        .add_attribute("amount", rewards_to_withdraw.to_string()))
}

// Update contract state periodically based on epochs
#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    match msg {
        ExecuteMsg::StakeTokens { amount } => stake_tokens(deps, info, amount),
        ExecuteMsg::WithdrawRewards {} => withdraw_rewards(deps, info),
        ExecuteMsg::UpdateEpoch {} => increment_epoch(deps),
        ExecuteMsg::DistributeLiquidTokens {} => distribute_liquid_tokens(deps, env),
        ExecuteMsg::DistributeRedemptionAmounts {} => distribute_redemption_amounts(deps, env),
    }
}

// Reward distribution logic
fn distribute_rewards(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;
    let current_epoch = state.current_epoch;

    for (staker, stake_info) in state.staking_info.iter_mut() {
        let reward_info = state.reward_info.entry(staker.clone()).or_insert_with(|| RewardInfo {
            total_rewards: Uint128::zero(),
            last_claim_epoch: current_epoch,
        });

        // Calculate rewards based on the stake amount
        let rewards = calculate_rewards(&stake_info.amount);
        reward_info.total_rewards += rewards;

        // Record the reward distribution
        let record = RewardRecord {
            rewards,
            calculated_time: env.block.time.to_string(),
            calculated_height: env.block.height,
        };

        let reward_records = state.reward_records.entry(staker.clone()).or_insert_with(Vec::new);
        reward_records.push(record);
    }

    state(deps.storage).save(&state)?;

    Ok(Response::new().add_attribute("method", "distribute_rewards"))
}

//function to calculate rewards
fn calculate_rewards(amount: &Uint128) -> Uint128 {
    *amount * Uint128::new(10)
}

// Distribute liquid tokens based on the recorded staked amounts
fn distribute_liquid_tokens(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;

    for contract in &state.contract_addresses {
        let amount = calculate_distribution(&state.total_liquid_tokens, &state.staking_info[contract].amount);
        
        let msg = BankMsg::Send {
            from_address: "liquidity_pool".to_string(),
            to_address: contract.clone(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        };
        
        state.total_liquid_tokens -= amount;
        state(deps.storage).save(&state)?;

        let response = Response::new()
            .add_message(msg)
            .add_attribute("action", "distribute_liquid_tokens")
            .add_attribute("contract", contract)
            .add_attribute("amount", amount.to_string());

        Ok(response)
    }
}

// Distribute redemption amounts based on the recorded ratios
fn distribute_redemption_amounts(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state = state(deps.storage).load()?;
    let total_tokens = state.total_liquid_tokens + state.total_redemption_tokens;

    for contract in &state.contract_addresses {
        let ratio = calculate_ratio(&state.staking_info[contract].amount, &total_tokens);
        let amount = calculate_amount(&state.total_redemption_tokens, &ratio);

        let msg = BankMsg::Send {
            from_address: "redemption_pool".to_string(),
            to_address: contract.clone(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        };

        state.total_redemption_tokens -= amount;
        state(deps.storage).save(&state)?;

        let response = Response::new()
            .add_message(msg)
            .add_attribute("action", "distribute_redemption_amounts")
            .add_attribute("contract", contract)
            .add_attribute("amount", amount.to_string());
        Ok(response)
    }
}

       

