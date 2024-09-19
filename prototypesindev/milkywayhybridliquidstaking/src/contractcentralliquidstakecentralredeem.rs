// contract.rs

use cosmwasm_std::{
    attr, entry_point, to_binary, Addr, BankMsg, Binary, Coin, CosmosMsg, Decimal, Deps, DepsMut,
    Env, MessageInfo, Response, StdError, StdResult, Uint128, Empty
};

use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

// ==================== State Definitions ====================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub owner: Addr, // Owner for access control
    pub current_epoch: u64,
    pub staking_info: HashMap<String, StakeInfo>,
    pub reward_info: HashMap<String, RewardInfo>,
    pub reward_records: HashMap<String, RewardRecord>,
    pub total_liquid_tokens: Uint128,
    pub total_redemption_tokens: Uint128,
    pub contract_addresses: Vec<String>,
    // Internal balances for central pools
    pub central_liquid_pool_balance: Uint128,
    pub central_redemption_pool_balance: Uint128,
    pub central_treasury_staking_pool_balance: Uint128,
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
pub struct InstantiateMsg {
    // No need to include central addresses anymore
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum ExecuteMsg {
    StakeTokens { amount: Uint128 },
    WithdrawRewards {},
    UpdateEpoch {},
    DistributeLiquidTokens {},
    DistributeRedemptionAmounts {},
    // Commands for cron job
    CentralLiquidStake {},
    CentralRedemption {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    TotalStaked { staker: String },
    TotalRewards { staker: String },
    GetState {}, // Added for testing purposes
}

// ==================== Contract Functions ====================

// Instantiate function
#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo, // Use info to get the sender as owner
    _msg: InstantiateMsg,
) -> StdResult<Response> {
    let state = State {
        owner: info.sender.clone(),
        current_epoch: 0,
        staking_info: HashMap::new(),
        reward_info: HashMap::new(),
        reward_records: HashMap::new(),
        total_liquid_tokens: Uint128::zero(),
        total_redemption_tokens: Uint128::zero(),
        contract_addresses: vec![],
        central_liquid_pool_balance: Uint128::zero(),
        central_redemption_pool_balance: Uint128::zero(),
        central_treasury_staking_pool_balance: Uint128::zero(),
    };
    STATE.save(deps.storage, &state)?;

    // Print log of important operation
    println!("Contract instantiated by {}", info.sender);

    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("owner", info.sender.to_string()))
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
        ExecuteMsg::UpdateEpoch {} => increment_epoch(deps, info),
        ExecuteMsg::DistributeLiquidTokens {} => distribute_liquid_tokens(deps, info),
        ExecuteMsg::DistributeRedemptionAmounts {} => {
            distribute_redemption_amounts(deps, info)
        }
        ExecuteMsg::CentralLiquidStake {} => central_liquid_stake(deps, info),
        ExecuteMsg::CentralRedemption {} => central_redemption(deps, info),
    }
}

// Query function
#[entry_point]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::TotalStaked { staker } => query_total_staked(deps, staker),
        QueryMsg::TotalRewards { staker } => query_total_rewards(deps, staker),
        QueryMsg::GetState {} => query_state(deps),
    }
}

// Increment epoch function
fn increment_epoch(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    let mut state_data = STATE.load(deps.storage)?;
    // Only owner can increment epoch
    if info.sender != state_data.owner {
        return Err(StdError::generic_err("Unauthorized"));
    }
    state_data.current_epoch += 1;
    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Epoch incremented to {} by {}",
        state_data.current_epoch, info.sender
    );

    Ok(Response::new()
        .add_attribute("method", "increment_epoch")
        .add_attribute("current_epoch", state_data.current_epoch.to_string()))
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

    // Limit mutable borrow scope
    let stake_amount: Uint128;
    {
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
            .map_err(|_| StdError::generic_err("Overflow in stake amount"))?;

        stake_amount = stake_info.amount;
    } // Mutable borrow ends here

    // Increase the central treasury staking pool balance
    state_data.central_treasury_staking_pool_balance = state_data
        .central_treasury_staking_pool_balance
        .checked_add(amount)
        .map_err(|_| StdError::generic_err("Overflow in treasury staking pool balance"))?;

    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Staker {} staked {} tokens. Total staked in pool: {}. Staker total stake: {}.",
        staker,
        amount,
        state_data.central_treasury_staking_pool_balance,
        stake_amount
    );

    Ok(Response::new()
        .add_attribute("action", "stake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount_staked", amount.to_string())
        .add_attribute(
            "total_staked_in_pool",
            state_data.central_treasury_staking_pool_balance.to_string(),
        )
        .add_attribute("staker_total_stake", stake_amount.to_string()))
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

    if rewards_to_withdraw.is_zero() {
        return Err(StdError::generic_err("No rewards available for withdrawal"));
    }

    // Decrease the central liquid pool balance
    state_data.central_liquid_pool_balance = state_data
        .central_liquid_pool_balance
        .checked_sub(rewards_to_withdraw)
        .map_err(|_| StdError::generic_err("Overflow in liquid pool balance"))?;

    STATE.save(deps.storage, &state_data)?;

    // Send tokens from the contract's balance to the staker
    let send_msg = BankMsg::Send {
        to_address: staker.clone(),
        amount: vec![Coin {
            denom: "token".to_string(),
            amount: rewards_to_withdraw,
        }],
    };

    // Print log of important operation
    println!(
        "Staker {} withdrew rewards: {} tokens. Remaining liquid pool balance: {}.",
        staker,
        rewards_to_withdraw,
        state_data.central_liquid_pool_balance
    );

    Ok(Response::new()
        .add_message(send_msg)
        .add_attribute("action", "withdraw_rewards")
        .add_attribute("staker", staker)
        .add_attribute("amount_withdrawn", rewards_to_withdraw.to_string())
        .add_attribute(
            "remaining_liquid_pool_balance",
            state_data.central_liquid_pool_balance.to_string(),
        ))
}

// Distribute liquid tokens function
fn distribute_liquid_tokens(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    // Only owner can distribute
    let mut state_data = STATE.load(deps.storage)?;
    if info.sender != state_data.owner {
        return Err(StdError::generic_err("Unauthorized"));
    }

    let mut messages = Vec::new();

    let total_liquid_tokens = state_data.central_liquid_pool_balance;

    if total_liquid_tokens.is_zero() {
        return Err(StdError::generic_err("No tokens available for distribution"));
    }

    // Calculate the total staked amount
    let total_staked: Uint128 = state_data
        .staking_info
        .values()
        .map(|info| info.amount)
        .sum();

    // Distribute tokens to each staker
    for (staker, stake_info) in state_data.staking_info.iter() {
        let stake_amount = stake_info.amount;

        let (_ratio, amount) = calculate_distribution_with_ratio(
            &total_liquid_tokens,
            &stake_amount,
            &total_staked,
        );

        if amount.is_zero() {
            continue;
        }

        // Decrease the central liquid pool balance
        state_data.central_liquid_pool_balance = state_data
            .central_liquid_pool_balance
            .checked_sub(amount)
            .map_err(|_| StdError::generic_err("Overflow in liquid pool balance"))?;

        // Send tokens from the contract's balance to the staker
        let msg = BankMsg::Send {
            to_address: staker.clone(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount,
            }],
        };

        messages.push(CosmosMsg::Bank(msg));

        // Print log of distribution to each staker
        println!(
            "Distributed {} liquid tokens to {}.",
            amount, staker
        );
    }

    if messages.is_empty() {
        return Err(StdError::generic_err("No stakers to distribute tokens"));
    }

    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Distributed total of {} liquid tokens. Remaining liquid pool balance: {}.",
        total_liquid_tokens,
        state_data.central_liquid_pool_balance
    );

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_liquid_tokens")
        .add_attribute(
            "remaining_liquid_pool_balance",
            state_data.central_liquid_pool_balance.to_string(),
        ))
}

// Distribute redemption amounts function
fn distribute_redemption_amounts(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    // Only owner can distribute
    let mut state_data = STATE.load(deps.storage)?;
    if info.sender != state_data.owner {
        return Err(StdError::generic_err("Unauthorized"));
    }

    let mut messages = Vec::new();

    let total_redemption_tokens = state_data.central_redemption_pool_balance;

    if total_redemption_tokens.is_zero() {
        return Err(StdError::generic_err(
            "No tokens available for redemption distribution",
        ));
    }

    // Calculate the total staked amount
    let total_staked: Uint128 = state_data
        .staking_info
        .values()
        .map(|info| info.amount)
        .sum();

    // Distribute tokens to each staker
    for (staker, stake_info) in state_data.staking_info.iter() {
        let stake_amount = stake_info.amount;

        let (_ratio, amount) = calculate_distribution_with_ratio(
            &total_redemption_tokens,
            &stake_amount,
            &total_staked,
        );

        if amount.is_zero() {
            continue;
        }

        // Decrease the central redemption pool balance
        state_data.central_redemption_pool_balance = state_data
            .central_redemption_pool_balance
            .checked_sub(amount)
            .map_err(|_| StdError::generic_err("Overflow in redemption pool balance"))?;

        // Send tokens from the contract's balance to the staker
        let msg = BankMsg::Send {
            to_address: staker.clone(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount,
            }],
        };

        messages.push(CosmosMsg::Bank(msg));

        // Print log of distribution to each staker
        println!(
            "Distributed {} redemption tokens to {}.",
            amount, staker
        );
    }

    if messages.is_empty() {
        return Err(StdError::generic_err(
            "No stakers to distribute redemption amounts",
        ));
    }

    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Distributed total of {} redemption tokens. Remaining redemption pool balance: {}.",
        total_redemption_tokens,
        state_data.central_redemption_pool_balance
    );

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_redemption_amounts")
        .add_attribute(
            "remaining_redemption_pool_balance",
            state_data.central_redemption_pool_balance.to_string(),
        ))
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

// Central Liquid Stake function
fn central_liquid_stake(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    // Only owner can execute
    let mut state_data = STATE.load(deps.storage)?;
    if info.sender != state_data.owner {
        return Err(StdError::generic_err("Unauthorized"));
    }

    // Simulate the amount to be added to the central liquid pool
    let simulated_liquid_tokens = Uint128::new(1000); // Replace with actual logic

    // Decrease the central treasury staking pool balance
    state_data.central_treasury_staking_pool_balance = state_data
        .central_treasury_staking_pool_balance
        .checked_sub(simulated_liquid_tokens)
        .map_err(|_| StdError::generic_err("Insufficient treasury staking pool balance"))?;

    // Increase the central liquid pool balance
    state_data.central_liquid_pool_balance = state_data
        .central_liquid_pool_balance
        .checked_add(simulated_liquid_tokens)
        .map_err(|_| StdError::generic_err("Overflow in liquid pool balance"))?;

    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Central liquid stake executed. Added {} tokens to central liquid pool. New balances - Central Liquid Pool: {}, Central Treasury Staking Pool: {}.",
        simulated_liquid_tokens,
        state_data.central_liquid_pool_balance,
        state_data.central_treasury_staking_pool_balance
    );

    Ok(Response::new()
        .add_attribute("method", "central_liquid_stake")
        .add_attribute("amount_added", simulated_liquid_tokens.to_string())
        .add_attribute(
            "central_liquid_pool_balance",
            state_data.central_liquid_pool_balance.to_string(),
        )
        .add_attribute(
            "central_treasury_staking_pool_balance",
            state_data.central_treasury_staking_pool_balance.to_string(),
        ))
}

// Central Redemption function
fn central_redemption(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    // Only owner can execute
    let mut state_data = STATE.load(deps.storage)?;
    if info.sender != state_data.owner {
        return Err(StdError::generic_err("Unauthorized"));
    }

    // Simulate the amount to be added to the central redemption pool
    let simulated_redemption_tokens = Uint128::new(500); // Replace with actual logic

    // Decrease the central treasury staking pool balance
    state_data.central_treasury_staking_pool_balance = state_data
        .central_treasury_staking_pool_balance
        .checked_sub(simulated_redemption_tokens)
        .map_err(|_| StdError::generic_err("Insufficient treasury staking pool balance"))?;

    // Increase the central redemption pool balance
    state_data.central_redemption_pool_balance = state_data
        .central_redemption_pool_balance
        .checked_add(simulated_redemption_tokens)
        .map_err(|_| StdError::generic_err("Overflow in redemption pool balance"))?;

    STATE.save(deps.storage, &state_data)?;

    // Print log of important operation
    println!(
        "Central redemption executed. Added {} tokens to central redemption pool. New balances - Central Redemption Pool: {}, Central Treasury Staking Pool: {}.",
        simulated_redemption_tokens,
        state_data.central_redemption_pool_balance,
        state_data.central_treasury_staking_pool_balance
    );

    Ok(Response::new()
        .add_attribute("method", "central_redemption")
        .add_attribute("amount_added", simulated_redemption_tokens.to_string())
        .add_attribute(
            "central_redemption_pool_balance",
            state_data.central_redemption_pool_balance.to_string(),
        )
        .add_attribute(
            "central_treasury_staking_pool_balance",
            state_data.central_treasury_staking_pool_balance.to_string(),
        ))
}

// Query total staked
fn query_total_staked(deps: Deps, staker: String) -> StdResult<Binary> {
    let state_data = STATE.load(deps.storage)?;
    if let Some(stake_info) = state_data.staking_info.get(&staker) {
        Ok(to_binary(&stake_info.amount)?)
    } else {
        Err(StdError::generic_err("Staker not found"))
    }
}

// Query total rewards
fn query_total_rewards(deps: Deps, staker: String) -> StdResult<Binary> {
    let state_data = STATE.load(deps.storage)?;
    if let Some(reward_info) = state_data.reward_info.get(&staker) {
        Ok(to_binary(&reward_info.total_rewards)?)
    } else {
        Err(StdError::generic_err("Staker not found"))
    }
}

// Query the entire state (for testing purposes)
fn query_state(deps: Deps) -> StdResult<Binary> {
    let state = STATE.load(deps.storage)?;
    Ok(to_binary(&state)?)
}


// tests.rs

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{Addr, Coin, Uint128};
    use cw_multi_test::{App, AppBuilder, Contract, ContractWrapper, Executor};

    // Helper function to create a mock app
    fn mock_app() -> App {
        AppBuilder::new().build(|_router, _api, _storage| {})
    }

    // Helper function to wrap the contract
    fn contract_staking() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_full_workflow_with_detailed_logs() -> Result<(), anyhow::Error> {
        // Initialize the app
        let mut app = mock_app();
        println!("Initialized mock app.");

        // Store the contract code
        let code_id = app.store_code(contract_staking());
        println!("Stored contract code with code_id: {}", code_id);

        // Instantiate the contract
        let owner = Addr::unchecked("owner");
        let contract_addr = app.instantiate_contract(
            code_id,
            owner.clone(),
            &InstantiateMsg {},
            &[],
            "StakingContract",
            None,
        )?;
        println!("Instantiated contract at address: {}", contract_addr);

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

        // Mint tokens to the contract's balance to simulate initial funds
        app.sudo(cw_multi_test::SudoMsg::Bank(cw_multi_test::BankSudo::Mint {
            to_address: contract_addr.to_string(),
            amount: vec![Coin {
                denom: "token".to_string(),
                amount: Uint128::new(5000),
            }],
        }))?;
        println!("Minted 5000 tokens to contract's balance");

        // Check initial balances
        let staker1_initial_balance = app.wrap().query_balance("staker1", "token")?.amount;
        let staker2_initial_balance = app.wrap().query_balance("staker2", "token")?.amount;
        let contract_initial_balance = app.wrap().query_balance(contract_addr.to_string(), "token")?.amount;

        println!("Initial Balances:");
        println!("Staker1: {}", staker1_initial_balance);
        println!("Staker2: {}", staker2_initial_balance);
        println!("Contract: {}", contract_initial_balance);

        // Staker1 stakes 1000 tokens
        let stake_amount_1 = Uint128::new(1000);
        let res = app.execute_contract(
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
        println!("Staker1 staked 1000 tokens.");

        // Extract and print attributes from response
        println!("StakeTokens Response Attributes for Staker1:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Staker2 stakes 500 tokens
        let stake_amount_2 = Uint128::new(500);
        let res = app.execute_contract(
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
        println!("Staker2 staked 500 tokens.");

        // Extract and print attributes from response
        println!("StakeTokens Response Attributes for Staker2:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Check internal balances after staking
        let state: State = app
            .wrap()
            .query_wasm_smart(&contract_addr, &QueryMsg::GetState {})?;
        println!(
            "Central Treasury Staking Pool Balance after staking: {}",
            state.central_treasury_staking_pool_balance
        );

        // Simulate central liquid staking (only owner can execute)
        let res = app.execute_contract(
            owner.clone(),
            contract_addr.clone(),
            &ExecuteMsg::CentralLiquidStake {},
            &[],
        )?;
        println!("Executed Central Liquid Stake.");

        // Extract and print attributes from response
        println!("CentralLiquidStake Response Attributes:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Distribute liquid tokens (only owner can execute)
        let res = app.execute_contract(
            owner.clone(),
            contract_addr.clone(),
            &ExecuteMsg::DistributeLiquidTokens {},
            &[],
        )?;
        println!("Distributed Liquid Tokens.");

        // Extract and print attributes from response
        println!("DistributeLiquidTokens Response Attributes:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Query balances after distribution
        let staker1_balance = app.wrap().query_balance("staker1", "token")?.amount;
        let staker2_balance = app.wrap().query_balance("staker2", "token")?.amount;
        println!("Balances after Liquid Token Distribution:");
        println!("Staker1: {}", staker1_balance);
        println!("Staker2: {}", staker2_balance);

        // Simulate central redemption (only owner can execute)
        let res = app.execute_contract(
            owner.clone(),
            contract_addr.clone(),
            &ExecuteMsg::CentralRedemption {},
            &[],
        )?;
        println!("Executed Central Redemption.");

        // Extract and print attributes from response
        println!("CentralRedemption Response Attributes:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Distribute redemption amounts (only owner can execute)
        let res = app.execute_contract(
            owner.clone(),
            contract_addr.clone(),
            &ExecuteMsg::DistributeRedemptionAmounts {},
            &[],
        )?;
        println!("Distributed Redemption Amounts.");

        // Extract and print attributes from response
        println!("DistributeRedemptionAmounts Response Attributes:");
        for event in res.events {
            println!("Event: {}", event.ty);
            for attr in event.attributes {
                println!("  {}: {}", attr.key, attr.value);
            }
        }

        // Query balances after redemption distribution
        let staker1_balance_after = app.wrap().query_balance("staker1", "token")?.amount;
        let staker2_balance_after = app.wrap().query_balance("staker2", "token")?.amount;
        println!("Balances after Redemption Distribution:");
        println!("Staker1: {}", staker1_balance_after);
        println!("Staker2: {}", staker2_balance_after);

        // Query final state
        let state: State = app
            .wrap()
            .query_wasm_smart(&contract_addr, &QueryMsg::GetState {})?;
        println!("Final State of the Contract:");
        println!("{:#?}", state);

        // Check contract's balance after all distributions
        let contract_balance = app
            .wrap()
            .query_balance(contract_addr.to_string(), "token")?
            .amount;
        println!("Contract's final balance: {}", contract_balance);

        // Additional detailed logs and calculations

        // Calculate how much each staker received from liquid token distribution
        let staker1_liquid_diff = staker1_balance - (staker1_initial_balance - stake_amount_1);
        let staker2_liquid_diff = staker2_balance - (staker2_initial_balance - stake_amount_2);
        println!("Staker1 received {} tokens from Liquid Distribution.", staker1_liquid_diff);
        println!("Staker2 received {} tokens from Liquid Distribution.", staker2_liquid_diff);

        // Calculate how much each staker received from redemption distribution
        let staker1_redemption_diff = staker1_balance_after - staker1_balance;
        let staker2_redemption_diff = staker2_balance_after - staker2_balance;
        println!("Staker1 received {} tokens from Redemption Distribution.", staker1_redemption_diff);
        println!("Staker2 received {} tokens from Redemption Distribution.", staker2_redemption_diff);

        // Print final balances
        println!("Final Balances:");
        println!("Staker1: {}", staker1_balance_after);
        println!("Staker2: {}", staker2_balance_after);
        println!("Contract: {}", contract_balance);

        Ok(())
    }
}


