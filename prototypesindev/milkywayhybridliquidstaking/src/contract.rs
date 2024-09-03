use cosmwasm_std::{
    attr, BankMsg, Coin, Deps, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Uint128,
    Decimal, Binary, OverflowError,
};
use cosmwasm_std::to_binary;
use cosmwasm_std::Empty;
use cw_storage_plus::Item;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use schemars::JsonSchema;
use cosmwasm_std::entry_point;
use crate::msg::{ExecuteMsg, InstantiateMsg, QueryMsg};
use crate::state::{StakeInfo, RewardInfo, RewardRecord, State, STATE};

// Instantiate function
#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> StdResult<Response> {
    let state = State {
        current_epoch: 0,
        staking_info: HashMap::new(),
        reward_info: HashMap::new(),
        reward_records: HashMap::new(),
        total_liquid_tokens: Uint128::zero(),
        total_redemption_tokens: Uint128::zero(),
        contract_addresses: vec![],
    };
    STATE.save(deps.storage, &state)?;
    Ok(Response::new().add_attribute("method", "instantiate"))
}

// Execute function
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

// Query function
#[entry_point]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::TotalStaked { staker } => query_total_staked(deps, staker),
        QueryMsg::TotalRewards { staker } => query_total_rewards(deps, staker),
    }
}

// Increment epoch function
fn increment_epoch(deps: DepsMut) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    state_data.current_epoch += 1;
    STATE.save(deps.storage, &state_data)?;
    Ok(Response::new().add_attribute("method", "increment_epoch"))
}

// Staking tokens function
fn stake_tokens(deps: DepsMut, info: MessageInfo, amount: Uint128) -> StdResult<Response> {
    if amount.is_zero() {
        return Err(StdError::generic_err("Cannot stake zero tokens"));
    }

    let mut state_data = STATE.load(deps.storage)?;
    let staker = info.sender.to_string();
    let stake_info = state_data.staking_info.entry(staker.clone()).or_insert_with(|| StakeInfo {
        amount: Uint128::zero(),
        last_updated_epoch: state_data.current_epoch,
    });

    stake_info.amount = stake_info
        .amount
        .checked_add(amount)
        .map_err(StdError::overflow)?;
    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_message(BankMsg::Send {
            to_address: "staking_pool".to_string(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        })
        .add_attribute("action", "stake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount", amount.to_string()))
}

// Withdraw rewards function
fn withdraw_rewards(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let staker = info.sender.to_string();
    let reward_info = state_data.reward_info.entry(staker.clone()).or_insert_with(|| RewardInfo {
        total_rewards: Uint128::zero(),
        last_claim_epoch: state_data.current_epoch,
    });

    let rewards_to_withdraw = reward_info.total_rewards;
    reward_info.total_rewards = Uint128::zero(); // Reset rewards after withdrawal
    STATE.save(deps.storage, &state_data)?;

    if rewards_to_withdraw.is_zero() {
        return Err(StdError::generic_err("No rewards available for withdrawal"));
    }

    Ok(Response::new()
        .add_message(BankMsg::Send {
            to_address: staker.clone(),
            amount: vec![Coin { denom: "token".to_string(), amount: rewards_to_withdraw }],
        })
        .add_attribute("action", "withdraw_rewards")
        .add_attribute("staker", staker)
        .add_attribute("amount", rewards_to_withdraw.to_string()))
}


// Distribute liquid tokens function
fn distribute_liquid_tokens(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let mut messages = Vec::new(); // Collect all messages to be dispatched

    for contract in &state_data.contract_addresses {
        let stake_amount = state_data
            .staking_info
            .get(contract)
            .map(|info| info.amount)
            .unwrap_or(Uint128::zero());

        let amount = calculate_distribution(&state_data.total_liquid_tokens, &stake_amount);

        if amount.is_zero() {
            continue; // Skip zero distributions
        }

        let msg = BankMsg::Send {
            to_address: contract.clone(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        };

        state_data.total_liquid_tokens = state_data
            .total_liquid_tokens
            .checked_sub(amount)
            .map_err(StdError::overflow)?;
        
        messages.push(msg);
    }

    if messages.is_empty() {
        return Err(StdError::generic_err("No contracts to distribute tokens"));
    }

    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_liquid_tokens"))
}

// Distribute redemption amounts function
fn distribute_redemption_amounts(deps: DepsMut, _env: Env) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let mut messages = Vec::new(); // Collect all messages to be dispatched

    let total_tokens = state_data
        .total_liquid_tokens
        .checked_add(state_data.total_redemption_tokens)
        .map_err(StdError::overflow)?;

    for contract in &state_data.contract_addresses {
        let stake_amount = state_data
            .staking_info
            .get(contract)
            .map(|info| info.amount)
            .unwrap_or(Uint128::zero());

        let ratio = calculate_ratio(&stake_amount, &total_tokens);
        let amount = calculate_amount(&state_data.total_redemption_tokens, &ratio);

        if amount.is_zero() {
            continue; // Skip zero distributions
        }

        let msg = BankMsg::Send {
            to_address: contract.clone(),
            amount: vec![Coin { denom: "token".to_string(), amount }],
        };

        state_data.total_redemption_tokens = state_data
            .total_redemption_tokens
            .checked_sub(amount)
            .map_err(StdError::overflow)?;

        messages.push(msg);
    }

    if messages.is_empty() {
        return Err(StdError::generic_err("No contracts to distribute redemption amounts"));
    }

    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_redemption_amounts"))
}


// Calculate the ratio of a given stake amount to the total tokens
fn calculate_ratio(stake_amount: &Uint128, total_tokens: &Uint128) -> Decimal {
    if total_tokens.is_zero() {
        Decimal::zero()
    } else {
        Decimal::from_ratio(*stake_amount, *total_tokens)
    }
}

// Calculate the amount to distribute based on the total redemption tokens and the ratio
fn calculate_amount(total_redemption_tokens: &Uint128, ratio: &Decimal) -> Uint128 {
    (*ratio * *total_redemption_tokens)
}

// Calculate distribution
fn calculate_distribution(total_liquid_tokens: &Uint128, _amount: &Uint128) -> Uint128 {
    *total_liquid_tokens // Simplified for demo
}

// Query total staked
fn query_total_staked(deps: Deps, staker: String) -> StdResult<Binary> {
    let state_data = STATE.load(deps.storage)?;
    if let Some(stake_info) = state_data.staking_info.get(&staker) {
        to_binary(&stake_info.amount)
    } else {
        Err(StdError::generic_err("Staker not found"))
    }
}

// Query total rewards
fn query_total_rewards(deps: Deps, staker: String) -> StdResult<Binary> {
    let state_data = STATE.load(deps.storage)?;
    if let Some(reward_info) = state_data.reward_info.get(&staker) {
        to_binary(&reward_info.total_rewards)
    } else {
        Err(StdError::generic_err("Staker not found"))
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{Addr, Coin, Uint128, StdError};
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};
    use std::convert::TryInto;

    fn mock_app() -> App {
        App::default()
    }

    fn contract_staking() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_full_liquid_staking_workflow() -> Result<(), StdError> {
        let mut app = mock_app();

        // Store the contract code
        let code_id = app.store_code(contract_staking());
        println!("Contract code stored with code_id: {}", code_id);

        // Instantiate the contract
        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("owner"),
                &InstantiateMsg {},
                &[],
                "StakingContract",
                None,
            )
            .map_err(|err| StdError::generic_err(format!("Failed to instantiate contract: {:?}", err)))?;
        println!("Contract instantiated at address: {}", contract_addr);

        // Mint initial tokens to staker1 and staker2
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: "staker1".to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(3000),
            }],
        }))
        .map_err(|err| StdError::generic_err(format!("Failed to mint tokens to staker1: {:?}", err)))?;
        println!("Minted 3000 tokens to staker1");

        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: "staker2".to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(2000),
            }],
        }))
        .map_err(|err| StdError::generic_err(format!("Failed to mint tokens to staker2: {:?}", err)))?;
        println!("Minted 2000 tokens to staker2");

      

        // Staker2 stakes 500 tokens
        let stake_amount_2 = Uint128::new(500);
        app.execute_contract(
            Addr::unchecked("staker2"),
            contract_addr.clone(),
            &ExecuteMsg::StakeTokens { amount: stake_amount_2 },
            &vec![Coin {
                denom: "token".to_string(),
                amount: stake_amount_2,
            }],
        )
        .map_err(|err| StdError::generic_err(format!("Failed to stake tokens by staker2: {:?}", err)))?;
        println!("Staker2 staked 500 tokens");

        // Simulate reward distribution by incrementing the epoch
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::UpdateEpoch {},
            &[],
        )
        .map_err(|err| StdError::generic_err(format!("Failed to update epoch: {:?}", err)))?;
        println!("Epoch updated to simulate reward distribution");

        // Mint rewards to the contract
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: contract_addr.to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(1500), // Total rewards to be distributed
            }],
        }))
        .map_err(|err| StdError::generic_err(format!("Failed to mint rewards to contract: {:?}", err)))?;
        println!("Minted 1500 tokens to contract address: {}", contract_addr);

       // Mint rewards to the contract
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: contract_addr.to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(1500), // Total rewards to be distributed
            }],
        }))
        .map_err(|err| StdError::generic_err(format!("Failed to mint rewards to contract: {:?}", err)))?;
        println!("Minted 1500 tokens to contract address: {}", contract_addr);

        Ok(())
    }
}
