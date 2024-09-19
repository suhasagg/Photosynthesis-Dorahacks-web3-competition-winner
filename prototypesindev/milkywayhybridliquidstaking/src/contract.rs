use cosmwasm_std::{
    attr, entry_point, to_binary, Addr, BankMsg, Binary, Coin, Decimal, Deps, DepsMut, Empty, Env,
    MessageInfo, Response, StdError, StdResult, Uint128,
};
use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

// ==================== State Definitions ====================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub current_epoch: u64,
    pub staking_info: HashMap<String, StakeInfo>,
    pub reward_info: HashMap<String, RewardInfo>,
    pub reward_records: HashMap<String, RewardRecord>,
    pub total_liquid_tokens: Uint128,
    pub total_redemption_tokens: Uint128,
    pub contract_addresses: Vec<String>,
}

pub const STATE: Item<State> = Item::new("state");

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakeInfo {
    pub amount: Uint128,
    pub last_updated_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct RewardInfo {
    pub total_rewards: Uint128,
    pub last_claim_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct RewardRecord {
    pub epoch: u64,
    pub amount: Uint128,
}

// ==================== Message Definitions ====================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum ExecuteMsg {
    StakeTokens { amount: Uint128 },
    WithdrawRewards {},
    UpdateEpoch {},
    DistributeLiquidTokens {},
    DistributeRedemptionAmounts {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    TotalStaked { staker: String },
    TotalRewards { staker: String },
}

// ==================== Contract Functions ====================

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

    // Verify that the staker has sent the correct amount of tokens
    let sent_amount = info
        .funds
        .iter()
        .find(|coin| coin.denom == "token")
        .map(|coin| coin.amount)
        .unwrap_or(Uint128::zero());

    if sent_amount != amount {
        return Err(StdError::generic_err("Incorrect token amount sent"));
    }

    let mut state_data = STATE.load(deps.storage)?;
    let staker = info.sender.to_string();
    let stake_info = state_data
        .staking_info
        .entry(staker.clone())
        .or_insert_with(|| StakeInfo {
            amount: Uint128::zero(),
            last_updated_epoch: state_data.current_epoch,
        });

    stake_info.amount = stake_info
        .amount
        .checked_add(amount)
        .map_err(StdError::overflow)?;
    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_attribute("action", "stake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount", amount.to_string()))
}

// Withdraw rewards function
fn withdraw_rewards(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let staker = info.sender.to_string();
    let reward_info = state_data
        .reward_info
        .entry(staker.clone())
        .or_insert_with(|| RewardInfo {
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
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: rewards_to_withdraw,
            }],
        })
        .add_attribute("action", "withdraw_rewards")
        .add_attribute("staker", staker)
        .add_attribute("amount", rewards_to_withdraw.to_string()))
}


fn distribute_liquid_tokens(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let mut messages = Vec::new();

    // Query the contract's balance
    let contract_balance = deps
        .querier
        .query_balance(env.contract.address.clone(), "token")?;
    let contract_balance_amount = contract_balance.amount;

    // Calculate the total staked amount
    let total_staked: Uint128 = state_data
        .staking_info
        .values()
        .map(|info| info.amount)
        .sum();

    // Calculate the tokens available for distribution 
    let total_liquid_tokens = contract_balance_amount
        .checked_sub(total_staked)
        .map_err(StdError::overflow)?;

    // Update total_liquid_tokens in state
    state_data.total_liquid_tokens = total_liquid_tokens;

    if total_liquid_tokens.is_zero() {
        return Err(StdError::generic_err("No tokens available for distribution"));
    }

    // Log the calculated values
    println!("Contract balance: {}", contract_balance_amount);
    println!("Total staked: {}", total_staked);
    println!("Total liquid tokens to distribute: {}", total_liquid_tokens);

    // Detailed logs for each staker
    for (staker, stake_info) in state_data.staking_info.iter() {
        let stake_amount = stake_info.amount;

        let (ratio, amount) = calculate_distribution_with_ratio(
            &state_data.total_liquid_tokens,
            &stake_amount,
            &total_staked,
        );

        if amount.is_zero() {
            continue;
        }

        // Log detailed information for each staker
        println!("Staker: {}", staker);
        println!("  Total Staked: {}", stake_amount);
        println!("  Staking Ratio: {:.4}", ratio);
        println!("  Liquid Tokens Received: {}", amount);

        let msg = BankMsg::Send {
            to_address: staker.clone(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount,
            }],
        };

        state_data.total_liquid_tokens = state_data
            .total_liquid_tokens
            .checked_sub(amount)
            .map_err(StdError::overflow)?;

        messages.push(msg);
    }

    if messages.is_empty() {
        return Err(StdError::generic_err("No stakers to distribute tokens"));
    }

    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_liquid_tokens"))
}


fn distribute_redemption_amounts(deps: DepsMut, env: Env) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    let mut messages = Vec::new();

    // Query the contract's balance
    let contract_balance = deps
        .querier
        .query_balance(env.contract.address.clone(), "token")?;
    let contract_balance_amount = contract_balance.amount;

    // Calculate the total staked amount
    let total_staked: Uint128 = state_data
        .staking_info
        .values()
        .map(|info| info.amount)
        .sum();

    // Calculate the tokens available for distribution 
    let total_redemption_tokens = contract_balance_amount
        .checked_sub(total_staked)
        .map_err(StdError::overflow)?;

    // Update total_redemption_tokens in state
    state_data.total_redemption_tokens = total_redemption_tokens;

    if total_redemption_tokens.is_zero() {
        return Err(StdError::generic_err(
            "No tokens available for redemption distribution",
        ));
    }

    // Log the calculated values
    println!("Contract balance: {}", contract_balance_amount);
    println!("Total staked: {}", total_staked);
    println!(
        "Total redemption tokens to distribute: {}",
        total_redemption_tokens
    );

    // Detailed logs for each staker
    for (staker, stake_info) in state_data.staking_info.iter() {
        let stake_amount = stake_info.amount;

        let (ratio, amount) = calculate_distribution_with_ratio(
            &state_data.total_redemption_tokens,
            &stake_amount,
            &total_staked,
        );

        if amount.is_zero() {
            continue;
        }

        // Log detailed information for each staker
        println!("Staker: {}", staker);
        println!("  Total Staked: {}", stake_amount);
        println!("  Staking Ratio: {:.4}", ratio);
        println!("  Redemption Tokens Received: {}", amount);

        let msg = BankMsg::Send {
            to_address: staker.clone(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount,
            }],
        };

        state_data.total_redemption_tokens = state_data
            .total_redemption_tokens
            .checked_sub(amount)
            .map_err(StdError::overflow)?;

        messages.push(msg);
    }

    if messages.is_empty() {
        return Err(StdError::generic_err(
            "No stakers to distribute redemption amounts",
        ));
    }

    STATE.save(deps.storage, &state_data)?;

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_redemption_amounts"))
}

// Calculate distribution with staking ratio
fn calculate_distribution_with_ratio(
    total_tokens: &Uint128,
    stake_amount: &Uint128,
    total_staked: &Uint128,
) -> (Decimal, Uint128) {
    if total_staked.is_zero() {
        (Decimal::zero(), Uint128::zero())
    } else {
        let ratio = Decimal::from_ratio(*stake_amount, *total_staked);
        let amount = ratio * *total_tokens;
        (ratio, amount)
    }
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

// ==================== Unit Tests ====================

#[cfg(test)]
mod tests {
    use super::*;
    use anyhow::Error;
    use cosmwasm_std::{Addr, Coin};
    use cw_multi_test::{App, AppBuilder, Contract, ContractWrapper, Executor};

    fn mock_app() -> App {
        AppBuilder::new().build(|_router, _api, _storage| {})
    }

    fn contract_staking() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_full_liquid_staking_workflow() -> Result<(), Error> {
        let mut app = mock_app();

        // Store the contract code
        let code_id = app.store_code(contract_staking());
        println!("Contract code stored with code_id: {}", code_id);

        // Instantiate the contract
        let contract_addr = app.instantiate_contract(
            code_id,
            Addr::unchecked("owner"),
            &InstantiateMsg {},
            &[],
            "StakingContract",
            None,
        )?;
        println!("Contract instantiated at address: {}", contract_addr);

        // Mint initial tokens to staker1 and staker2
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: "staker1".to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(3000),
            }],
        }))?;
        println!("Minted 3000 tokens to staker1");

        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: "staker2".to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(2000),
            }],
        }))?;
        println!("Minted 2000 tokens to staker2");

        // Staker1 stakes 1000 tokens
        let stake_amount_1 = Uint128::new(1000);
        app.execute_contract(
            Addr::unchecked("staker1"),
            contract_addr.clone(),
            &ExecuteMsg::StakeTokens {
                amount: stake_amount_1,
            },
            &vec![Coin {
                denom: "token".to_string(),
                amount: stake_amount_1,
            }],
        )?;
        println!("Staker1 staked 1000 tokens");

        // Staker2 stakes 500 tokens
        let stake_amount_2 = Uint128::new(500);
        app.execute_contract(
            Addr::unchecked("staker2"),
            contract_addr.clone(),
            &ExecuteMsg::StakeTokens {
                amount: stake_amount_2,
            },
            &vec![Coin {
                denom: "token".to_string(),
                amount: stake_amount_2,
            }],
        )?;
        println!("Staker2 staked 500 tokens");

        // Simulate reward distribution by incrementing the epoch
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::UpdateEpoch {},
            &[],
        )?;
        println!("Epoch updated to simulate reward distribution");

        // Mint rewards to the contract
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: contract_addr.to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(1500),
            }],
        }))?;
        println!(
            "Minted 1500 tokens to contract address: {}",
            contract_addr
        );

        // Distribute liquid tokens
        println!("Distributing liquid tokens...");
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::DistributeLiquidTokens {},
            &[],
        )?;
        println!("Distributed liquid tokens");

        // Query balances after distribution
        let staker1_balance = app.wrap().query_balance("staker1", "token")?;
        let staker2_balance = app.wrap().query_balance("staker2", "token")?;
        println!("Staker1 balance: {}", staker1_balance.amount);
        println!("Staker2 balance: {}", staker2_balance.amount);


        // Mint redemption tokens to the contract
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: contract_addr.to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(1500),
            }],
        }))?;
        println!(
            "Minted 1500 redemption tokens to contract address: {}",
            contract_addr
        );

        // Distribute redemption amounts
        println!("Distributing redemption amounts...");
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::DistributeRedemptionAmounts {},
            &[],
        )?;
        println!("Distributed redemption amounts");

        // Query balances after redemption distribution
        let staker1_balance_after = app.wrap().query_balance("staker1", "token")?;
        let staker2_balance_after = app.wrap().query_balance("staker2", "token")?;
        println!(
            "Staker1 balance after redemption: {}",
            staker1_balance_after.amount
        );
        println!(
            "Staker2 balance after redemption: {}",
            staker2_balance_after.amount
        );


        Ok(())
    }
}
