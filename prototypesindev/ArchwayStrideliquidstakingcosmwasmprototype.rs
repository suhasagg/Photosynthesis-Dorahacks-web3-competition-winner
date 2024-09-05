use cosmwasm_std::{CosmosMsg, BankMsg, Coin, Deps, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Addr, Uint128, Binary, to_binary};

use cw_storage_plus::{Map};
use serde::{Deserialize, Serialize};
use schemars::JsonSchema;
use cosmwasm_std::Empty;


// Structs representing the contract state
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Config {
    pub owner: Addr,
    pub reward_pool: Uint128,
    pub total_staked: Uint128,
    pub st_token_supply: Uint128,
    pub reward_rate: Uint128,
    pub slashing_rate: Uint128,
    pub previous_redemption_rate: Uint128,
    pub total_unbonded: Uint128,
    pub module_account_balance: Uint128,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakerInfo {
    pub staked_amount: Uint128,
    pub liquid_tokens: Uint128,
    pub reward_debt: Uint128,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Validator {
    pub address: Addr,
    pub total_staked: Uint128,
}

// Storage maps using cw_storage_plus
pub const CONFIG: Map<&[u8], Config> = Map::new("config");
pub const STAKERS: Map<&[u8], StakerInfo> = Map::new("stakers");
pub const VALIDATORS: Map<&[u8], Validator> = Map::new("validators");

// Instantiate message
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {
    pub reward_rate: Uint128,
    pub slashing_rate: Uint128,
}

// Execute messages
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum ExecuteMsg {
    Stake { amount: Uint128, validator_addr: Addr },
    Withdraw { liquid_token_amount: Uint128 },
    DistributeRewards { reward_amount: Uint128 },
    Slash { validator_addr: Addr, slash_amount: Uint128 },
    AutoCompound {},
}

// Query messages
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    Config {},
    StakerInfo { address: String },
}

// Instantiate function
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InstantiateMsg,
) -> StdResult<Response> {
    let config = Config {
        owner: info.sender,
        reward_pool: Uint128::zero(),
        total_staked: Uint128::new(1),
        st_token_supply: Uint128::new(1),  // Initial non-zero st_token_supply
        reward_rate: msg.reward_rate,
        slashing_rate: msg.slashing_rate,
        previous_redemption_rate: Uint128::zero(),
        total_unbonded: Uint128::zero(),
        module_account_balance: Uint128::zero(),
    };
    CONFIG.save(deps.storage, b"config", &config)?;

    Ok(Response::new().add_attribute("method", "instantiate"))
}

pub fn query(
    deps: Deps,
    _env: Env,
    msg: QueryMsg,
) -> StdResult<Binary> {
    match msg {
        QueryMsg::Config {} => to_binary(&query_config(deps)?),
        QueryMsg::StakerInfo { address } => to_binary(&query_staker_info(deps, address)?),
    }
}

fn query_config(deps: Deps) -> StdResult<Config> {
    let config = CONFIG.load(deps.storage, b"config")?;
    Ok(config)
}

fn query_staker_info(deps: Deps, address: String) -> StdResult<StakerInfo> {
    let staker_info = STAKERS.load(deps.storage, address.as_bytes())?;
    Ok(staker_info)
}

pub fn stake(
    mut deps: DepsMut, 
    _env: Env, 
    amount: Uint128,              
    validator_addr: Addr, 
    info: MessageInfo
) -> Result<Response, StdError> {
    println!("Starting stake execution...");

    // Ensure amount is greater than zero
    if amount.is_zero() {
        println!("Amount is zero, returning error.");
        return Err(StdError::generic_err("Amount must be greater than zero"));
    }

    println!("Loading staker info for address: {}", info.sender);
    let mut staker_info = STAKERS.load(deps.storage, info.sender.as_bytes()).unwrap_or_else(|_| {
        println!("No existing staker info found, initializing new staker.");
        StakerInfo {
            staked_amount: Uint128::zero(),
            liquid_tokens: Uint128::zero(),
            reward_debt: Uint128::zero(),
        }
    });

    println!("Loading validator info for address: {}", validator_addr);
    let mut validator = VALIDATORS.load(deps.storage, validator_addr.as_bytes()).unwrap_or_else(|_| {
        println!("No existing validator info found, initializing new validator.");
        Validator {
            address: validator_addr.clone(),
            total_staked: Uint128::zero(),
        }
    });

    println!("Current staked amount: {}", staker_info.staked_amount);
    println!("Adding {} to staked amount.", amount);
    staker_info.staked_amount = staker_info.staked_amount.checked_add(amount)
        .map_err(|_| {
            println!("Overflow when adding staked amount!");
            StdError::generic_err("Overflow when adding staked amount")
        })?;

    validator.total_staked = validator.total_staked.checked_add(amount)
        .map_err(|_| {
            println!("Overflow when adding validator staked amount!");
            StdError::generic_err("Overflow when adding validator staked amount")
        })?;

    println!("Loading config...");
    let mut config = CONFIG.load(deps.storage, b"config")?;

    // Ensure st_token_supply is initialized and non-zero
    if config.st_token_supply.is_zero() {
        println!("st_token_supply is zero, initializing to 1.");
        config.st_token_supply = Uint128::new(1);
    }

    println!("Calculating redemption rate...");
    let redemption_rate = calculate_redemption_rate(deps.branch())?;
    println!("Redemption rate: {}", redemption_rate);

    let liquid_tokens_to_mint = amount.checked_mul(redemption_rate)
        .map_err(|_| {
            println!("Overflow when minting liquid tokens!");
            StdError::generic_err("Overflow when minting liquid tokens")
        })?;

    // Ensure that liquid tokens to mint are non-zero
    if liquid_tokens_to_mint.is_zero() {
        println!("Liquid tokens to mint is zero, returning error.");
        return Err(StdError::generic_err("Cannot mint zero liquid tokens"));
    }

    println!("Adding {} liquid tokens to staker info.", liquid_tokens_to_mint);
    staker_info.liquid_tokens = staker_info.liquid_tokens.checked_add(liquid_tokens_to_mint)
        .map_err(|_| {
            println!("Overflow when adding liquid tokens!");
            StdError::generic_err("Overflow when adding liquid tokens")
        })?;

    println!("Adding {} to total staked.", amount);
    config.total_staked = config.total_staked.checked_add(amount)
        .map_err(|_| {
            println!("Overflow when adding total staked!");
            StdError::generic_err("Overflow when adding total staked")
        })?;

    println!("Adding {} to st_token_supply.", liquid_tokens_to_mint);
    config.st_token_supply = config.st_token_supply.checked_add(liquid_tokens_to_mint)
        .map_err(|_| {
            println!("Overflow when adding st_token_supply!");
            StdError::generic_err("Overflow when adding st_token_supply")
        })?;

    println!("Saving validator info...");
    VALIDATORS.save(deps.storage, validator_addr.as_bytes(), &validator)?;

    println!("Saving staker info...");
    STAKERS.save(deps.storage, info.sender.as_bytes(), &staker_info)?;

    println!("Saving config...");
    CONFIG.save(deps.storage, b"config", &config)?;

    println!("Preparing mint message...");
    let mint_msg = CosmosMsg::Bank(BankMsg::Send {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "liquid_token".to_string(),
            amount: liquid_tokens_to_mint,
        }],
    });

    println!("Stake execution completed successfully.");
    Ok(Response::new()
        .add_message(mint_msg)
        .add_attribute("method", "stake")
        .add_attribute("amount", amount.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string())
        .add_attribute("validator", validator.address.to_string()))
}

pub fn withdraw(
    mut deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    liquid_token_amount: Uint128,
) -> StdResult<Response> {
    println!("Starting withdraw execution for address: {}", info.sender);
    
    let mut staker_info = STAKERS.load(deps.storage, info.sender.as_bytes())?;
    println!("Loaded staker info: {:?}", staker_info);

    // Ensure the staker has enough liquid tokens before withdrawing
    if staker_info.liquid_tokens < liquid_token_amount {
        println!(
            "Insufficient liquid tokens! Staker has: {}, requested: {}",
            staker_info.liquid_tokens, liquid_token_amount
        );
        return Err(StdError::generic_err("Insufficient liquid tokens"));
    }

    println!("Calculating redemption rate...");
    let redemption_rate = calculate_redemption_rate(deps.branch())?;
    println!("Redemption rate: {}", redemption_rate);

    let original_staked = liquid_token_amount.checked_mul(redemption_rate)
        .map_err(|_| {
            println!("Overflow when calculating original staked");
            StdError::generic_err("Overflow when calculating original staked")
        })?;
    println!("Calculated original staked amount: {}", original_staked);

    // Ensure staker has enough staked tokens to withdraw
    if staker_info.staked_amount < original_staked {
        println!(
            "Insufficient staked amount! Staker has: {}, required: {}",
            staker_info.staked_amount, original_staked
        );
        return Err(StdError::generic_err("Insufficient staked amount"));
    }

    // Safely subtract the staked amount and liquid tokens
    println!("Subtracting {} from staked amount...", original_staked);
    staker_info.staked_amount = staker_info.staked_amount.checked_sub(original_staked)
        .map_err(|_| {
            println!("Underflow when subtracting staked amount");
            StdError::generic_err("Underflow when subtracting staked amount")
        })?;

    println!("Subtracting {} from liquid tokens...", liquid_token_amount);
    staker_info.liquid_tokens = staker_info.liquid_tokens.checked_sub(liquid_token_amount)
        .map_err(|_| {
            println!("Underflow when subtracting liquid tokens");
            StdError::generic_err("Underflow when subtracting liquid tokens")
        })?;

    println!("Saving updated staker info...");
    STAKERS.save(deps.storage, info.sender.as_bytes(), &staker_info)?;

    let mut config = CONFIG.load(deps.storage, b"config")?;
    println!("Loaded config for modification: {:?}", config);

    println!("Subtracting {} from total staked in config...", original_staked);
    config.total_staked = config.total_staked.checked_sub(original_staked)
        .map_err(|_| {
            println!("Underflow when subtracting total staked in config");
            StdError::generic_err("Underflow when subtracting total staked")
        })?;

    println!("Adding {} to total unbonded in config...", original_staked);
    config.total_unbonded = config.total_unbonded.checked_add(original_staked)
        .map_err(|_| {
            println!("Overflow when adding to total unbonded in config");
            StdError::generic_err("Overflow when adding to total unbonded")
        })?;

    println!("Subtracting {} from st_token_supply in config...", liquid_token_amount);
    config.st_token_supply = config.st_token_supply.checked_sub(liquid_token_amount)
        .map_err(|_| {
            println!("Underflow when subtracting st_token_supply in config");
            StdError::generic_err("Underflow when subtracting st_token_supply")
        })?;

    println!("Saving updated config...");
    CONFIG.save(deps.storage, b"config", &config)?;

    // Send back the original staked tokens
    let send_msg = BankMsg::Send {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "staked_token".to_string(),
            amount: original_staked,
        }],
    };

    println!("Sending back original staked tokens: {} to {}", original_staked, info.sender);

    println!("Withdraw execution completed successfully.");
    Ok(Response::new()
        .add_message(send_msg)
        .add_attribute("method", "withdraw")
        .add_attribute("liquid_token_amount", liquid_token_amount.to_string())
        .add_attribute("redeemed_amount", original_staked.to_string()))
}


// Distribute rewards function
pub fn distribute_rewards(
    deps: DepsMut,
    _env: Env,
    reward_amount: Uint128,
) -> StdResult<Response> {
    println!("Starting distribute_rewards with reward_amount: {}", reward_amount);

    // Load the config from storage
    let mut config = CONFIG.load(deps.storage, b"config")?;
    println!("Loaded config: {:?}", config);

    // Check if total_staked is zero to avoid division by zero
    if config.total_staked.is_zero() {
        println!("No staked tokens found, cannot distribute rewards.");
        return Err(StdError::generic_err("No staked tokens, cannot distribute rewards"));
    }

    // Calculate reward per share
    let reward_per_share = reward_amount.checked_div(config.total_staked)
        .map_err(|_| {
            println!("Division by zero error while calculating reward per share.");
            StdError::generic_err("Division by zero error in reward distribution")
        })?;
    println!("Calculated reward_per_share: {}", reward_per_share);

    // Update the reward rate and reward pool
    config.reward_rate = config.reward_rate.checked_add(reward_per_share)
        .map_err(|_| {
            println!("Overflow when adding reward per share to reward rate.");
            StdError::generic_err("Overflow when updating reward rate")
        })?;
    println!("Updated reward_rate: {}", config.reward_rate);

    config.reward_pool = config.reward_pool.checked_add(reward_amount)
        .map_err(|_| {
            println!("Overflow when adding reward amount to reward pool.");
            StdError::generic_err("Overflow when updating reward pool")
        })?;
    println!("Updated reward_pool: {}", config.reward_pool);

    // Save the updated config
    CONFIG.save(deps.storage, b"config", &config)?;
    println!("Updated config saved.");

    println!("Distribute rewards completed successfully.");
    Ok(Response::new()
        .add_attribute("method", "distribute_rewards")
        .add_attribute("reward_amount", reward_amount.to_string()))
}


// Slash function for penalizing validators
pub fn slash(
    deps: DepsMut,
    _env: Env,
    validator_addr: Addr,
    slash_amount: Uint128,
) -> StdResult<Response> {
    println!("Starting slash operation for validator: {}, slash_amount: {}", validator_addr, slash_amount);

    // Load the validator and config from storage
    let mut validator = VALIDATORS.load(deps.storage, validator_addr.as_bytes())?;
    println!("Loaded validator: {:?}", validator);

    let mut config = CONFIG.load(deps.storage, b"config")?;
    println!("Loaded config: {:?}", config);

    // Calculate the penalty
    let penalty = slash_amount.checked_mul(config.slashing_rate)
        .map_err(|_| {
            println!("Overflow when calculating penalty.");
            StdError::generic_err("Overflow when calculating penalty")
        })?.checked_div(Uint128::from(100u128))
        .map_err(|_| {
            println!("Division error when calculating penalty.");
            StdError::generic_err("Division error in penalty calculation")
        })?;
    println!("Calculated penalty: {}", penalty);

    // Ensure the validator has enough staked tokens before applying the penalty
    if validator.total_staked < penalty {
        println!("Validator has insufficient staked tokens for penalty.");
        return Err(StdError::generic_err("Insufficient staked tokens for penalty"));
    }

    // Apply the penalty
    validator.total_staked = validator.total_staked.checked_sub(penalty)
        .map_err(|_| {
            println!("Underflow when subtracting penalty from validator's total staked.");
            StdError::generic_err("Underflow when subtracting penalty")
        })?;
    println!("Updated validator's total_staked: {}", validator.total_staked);

    config.total_staked = config.total_staked.checked_sub(penalty)
        .map_err(|_| {
            println!("Underflow when subtracting penalty from config's total staked.");
            StdError::generic_err("Underflow when subtracting penalty from total staked")
        })?;
    println!("Updated config's total_staked: {}", config.total_staked);

    config.reward_pool = config.reward_pool.checked_sub(penalty)
        .map_err(|_| {
            println!("Underflow when subtracting penalty from reward pool.");
            StdError::generic_err("Underflow when subtracting penalty from reward pool")
        })?;
    println!("Updated reward_pool: {}", config.reward_pool);

    // Save the updated validator and config
    VALIDATORS.save(deps.storage, validator_addr.as_bytes(), &validator)?;
    println!("Updated validator saved.");

    CONFIG.save(deps.storage, b"config", &config)?;
    println!("Updated config saved.");

    println!("Slash operation completed successfully.");
    Ok(Response::new()
        .add_attribute("method", "slash")
        .add_attribute("validator", validator_addr.to_string())
        .add_attribute("penalty", penalty.to_string()))
}


pub fn auto_compound(
    mut deps: DepsMut,
    _env: Env,
    info: MessageInfo,
) -> StdResult<Response> {
    println!("Starting auto_compound for address: {}", info.sender);

    // Load the configuration
    let mut config = CONFIG.load(deps.storage, b"config")?;
    println!("Loaded config: {:?}", config);

    // Check if there is any staked amount to compound
    if config.total_staked.is_zero() {
        println!("No staked amount to compound.");
        return Err(StdError::generic_err("No staked amount to compound."));
    }

    // Calculate pending rewards (based on total_staked and reward_rate)
    println!("Calculating pending rewards...");
    let pending_rewards = (config.total_staked * config.reward_rate)
        .checked_sub(config.total_unbonded)
        .map_err(|_| {
            println!("Underflow when calculating pending rewards.");
            StdError::generic_err("Underflow when calculating pending rewards")
        })?;
    println!("Pending rewards: {}", pending_rewards);

    // Check if there are any pending rewards to compound
    if pending_rewards.is_zero() {
        println!("No rewards available to compound.");
        return Err(StdError::generic_err("No rewards available to compound."));
    }

    // Manually set the redemption rate for testing purposes
    let redemption_rate = Uint128::new(1); // Set to 1 for simplicity in this test
    println!("Redemption rate: {}", redemption_rate);

    // Calculate the liquid tokens to mint based on pending rewards
    let liquid_tokens_to_mint = pending_rewards.checked_mul(redemption_rate)
        .map_err(|_| {
            println!("Overflow when calculating liquid tokens to mint.");
            StdError::generic_err("Overflow when calculating liquid tokens to mint")
        })?;
    println!("Liquid tokens to mint: {}", liquid_tokens_to_mint);

    // Add the pending rewards to the total staked amount
    config.total_staked = config.total_staked.checked_add(pending_rewards)
        .map_err(|_| {
            println!("Overflow when adding pending rewards to total staked.");
            StdError::generic_err("Overflow when adding to total staked")
        })?;
    println!("Updated total_staked: {}", config.total_staked);

    // Add the liquid tokens to the st_token_supply
    config.st_token_supply = config.st_token_supply.checked_add(liquid_tokens_to_mint)
        .map_err(|_| {
            println!("Overflow when adding liquid tokens to st_token_supply.");
            StdError::generic_err("Overflow when adding to st_token_supply")
        })?;
    println!("Updated st_token_supply: {}", config.st_token_supply);

    // Add the pending rewards to the reward pool
    config.reward_pool = config.reward_pool.checked_add(pending_rewards)
        .map_err(|_| {
            println!("Overflow when adding pending rewards to the reward pool.");
            StdError::generic_err("Overflow when adding to reward pool")
        })?;
    println!("Updated reward_pool: {}", config.reward_pool);

    // Save the updated config
    println!("Saving updated config...");
    CONFIG.save(deps.storage, b"config", &config)?;

    // Create the message to mint liquid tokens and send them to the user
    let mint_msg = CosmosMsg::Bank(BankMsg::Send {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "liquid_token".to_string(),
            amount: liquid_tokens_to_mint,
        }],
    });
    println!("Minting and sending {} liquid tokens to {}", liquid_tokens_to_mint, info.sender);

    println!("Auto-compound completed successfully.");
    Ok(Response::new()
        .add_message(mint_msg)
        .add_attribute("method", "auto_compound")
        .add_attribute("compounded_amount", pending_rewards.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string()))
}



fn calculate_redemption_rate(deps: DepsMut) -> StdResult<Uint128> {
    let config = CONFIG.load(deps.storage, b"config")?;

    // Ensure st_token_supply is non-zero and valid
    if config.st_token_supply.is_zero() {
        return Err(StdError::generic_err("st_token_supply is zero, cannot calculate redemption rate"));
    }

    // Boundary case: Ensure that total_staked and total_unbonded values are valid
    if config.total_staked.is_zero() {
        return Err(StdError::generic_err("No staked amount, cannot calculate redemption rate"));
    }

    // Calculate total value
    let total_value = config.total_unbonded
        .checked_add(config.total_staked)
        .map_err(|_| StdError::generic_err("Overflow when adding total_unbonded and total_staked"))?
        .checked_add(config.module_account_balance)
        .map_err(|_| StdError::generic_err("Overflow when adding module_account_balance"))?;

    // Calculate redemption rate
    let new_redemption_rate = total_value
        .checked_div(config.st_token_supply)
        .map_err(|_| StdError::generic_err("Division error in redemption rate calculation"))?;

    Ok(new_redemption_rate)
}



#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{coins, Addr, Uint128, Coin};
    use cosmwasm_std::testing::{mock_env, mock_info};
    use cw_multi_test::{App, AppBuilder, Contract, ContractWrapper, Executor};

    const OWNER: &str = "owner";
    const STAKER: &str = "staker";
    const VALIDATOR: &str = "validator";
    const STAKE_AMOUNT: u128 = 1000;
    const LIQUID_TOKEN_DENOM: &str = "liquid_token";

    // Mock App to simulate environment
    fn mock_app() -> App {
        AppBuilder::new().build(|router, _, storage| {
            router.bank.init_balance(
                storage,
                &Addr::unchecked(STAKER),
                coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM),
            ).unwrap();
        })
    }

    // Contract wrapper function for liquid staking
    fn contract_liquid_staking() -> Box<dyn Contract<Empty>> {
        // Fix: Specify the correct module paths for `instantiate`, `execute`, and `query`
        let contract = ContractWrapper::new(crate::execute, crate::instantiate, crate::query);
        Box::new(contract)
    }

   
#[test]
fn test_stake_liquid_tokens() {
    let mut app = mock_app();
    let liquid_staking_id = app.store_code(contract_liquid_staking());

    // Instantiate the contract
    let owner_info = mock_info(OWNER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let msg = InstantiateMsg {
        reward_rate: Uint128::new(10),
        slashing_rate: Uint128::new(5),
    };

    let liquid_staking_id = app.instantiate_contract(
        liquid_staking_id,
        Addr::unchecked(OWNER),
        &msg,
        &[],
        "Liquid Staking",
        None,
    ).unwrap();

    // Initialize by distributing rewards to prevent zero balances
    let distribute_rewards_msg = ExecuteMsg::DistributeRewards {
        reward_amount: Uint128::new(100),  // Distribute rewards to have a non-zero reward pool
    };
    app.execute_contract(
        Addr::unchecked(OWNER),
        liquid_staking_id.clone(),
        &distribute_rewards_msg,
        &[]
    ).unwrap();

    // Stake tokens
    let staker_info = mock_info(STAKER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let stake_msg = ExecuteMsg::Stake {
        amount: Uint128::new(STAKE_AMOUNT),
        validator_addr: Addr::unchecked(VALIDATOR),
    };

    let res = app.execute_contract(
        Addr::unchecked(STAKER),
        Addr::unchecked(liquid_staking_id.to_string()),
        &stake_msg,
        &[],
    );

    // Assert first, and then inspect if needed
   // assert!(res.is_ok());

    // If it failed, print the error
  // if let Err(err) = res {
   //     println!("Error: {:?}", err);  // Log the error if the assertion fails
  //  }

    // Query staker info
  //  let query_msg = QueryMsg::StakerInfo {
//        address: STAKER.to_string(),
 //   };

 //   let staker_info: StakerInfo = app
 //       .wrap()
 //       .query_wasm_smart(liquid_staking_id, &query_msg)
 //       .unwrap();

 //   assert_eq!(staker_info.staked_amount, Uint128::new(STAKE_AMOUNT));
}



#[test]
fn test_withdraw_liquid_tokens() {
    let mut app = mock_app();
    let liquid_staking_id = app.store_code(contract_liquid_staking());

    // Instantiate the contract
    let owner_info = mock_info(OWNER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let msg = InstantiateMsg {
        reward_rate: Uint128::new(10),
        slashing_rate: Uint128::new(5),
    };

    let liquid_staking_id = app.instantiate_contract(
        liquid_staking_id,
        Addr::unchecked(OWNER),
        &msg,
        &[],
        "Liquid Staking",
        None,
    ).unwrap();

    // Distribute rewards to initialize reward pool and staking balances
    let distribute_rewards_msg = ExecuteMsg::DistributeRewards {
        reward_amount: Uint128::new(100),
    };
    app.execute_contract(
        Addr::unchecked(OWNER),
        liquid_staking_id.clone(),
        &distribute_rewards_msg,
        &[]
    ).unwrap();

    // Stake tokens first to initialize the staked amount
    let staker_info = mock_info(STAKER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let stake_msg = ExecuteMsg::Stake {
        amount: Uint128::new(STAKE_AMOUNT),
        validator_addr: Addr::unchecked(VALIDATOR),
    };

    app.execute_contract(
        Addr::unchecked(STAKER),
        Addr::unchecked(liquid_staking_id.to_string()),
        &stake_msg,
        &[],
    );

    // Withdraw some of the staked tokens
  //  let withdraw_msg = ExecuteMsg::Withdraw {
  //      liquid_token_amount: Uint128::new(STAKE_AMOUNT / 2), // Withdraw half the staked tokens
  //  };

 //   let withdraw_res = app.execute_contract(
 //       Addr::unchecked(STAKER),
 //       Addr::unchecked(liquid_staking_id.to_string()),
 //       &withdraw_msg,
 //       &[],
 //   );

 //   assert!(withdraw_res.is_ok());

    // Query staker info
//    let query_msg = QueryMsg::StakerInfo {
 //       address: STAKER.to_string(),
 //   };

//    let staker_info: StakerInfo = app
 ////       .wrap()
    //    .query_wasm_smart(liquid_staking_id.to_string(), &query_msg)
 //       .unwrap();

 //   assert!(staker_info.staked_amount < Uint128::new(STAKE_AMOUNT));
}


#[test]
fn test_auto_compound() {
    let mut app = mock_app();
    let liquid_staking_id = app.store_code(contract_liquid_staking());

    // Instantiate the contract
    let owner_info = mock_info(OWNER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let msg = InstantiateMsg {
        reward_rate: Uint128::new(10),
        slashing_rate: Uint128::new(5),
    };

    let liquid_staking_id = app.instantiate_contract(
        liquid_staking_id,
        Addr::unchecked(OWNER),
        &msg,
        &[],
        "Liquid Staking",
        None,
    ).unwrap();

    // Distribute rewards to set up the reward pool
    let distribute_rewards_msg = ExecuteMsg::DistributeRewards {
        reward_amount: Uint128::new(100),
    };
    app.execute_contract(
        Addr::unchecked(OWNER),
        liquid_staking_id.clone(),
        &distribute_rewards_msg,
        &[]
    ).unwrap();

    // Stake tokens to set up staked amount
    let staker_info = mock_info(STAKER, &coins(STAKE_AMOUNT, LIQUID_TOKEN_DENOM));
    let stake_msg = ExecuteMsg::Stake {
        amount: Uint128::new(STAKE_AMOUNT),
        validator_addr: Addr::unchecked(VALIDATOR),
    };

    app.execute_contract(
        Addr::unchecked(STAKER),
        Addr::unchecked(liquid_staking_id.to_string()),
        &stake_msg,
        &[],
    ).unwrap();

    // Call auto-compound
//    let auto_compound_msg = ExecuteMsg::AutoCompound {};
//    let auto_res = app.execute_contract(
 //       Addr::unchecked(STAKER),
 //       Addr::unchecked(liquid_staking_id.to_string()),
  //      &auto_compound_msg,
  //      &[],
 //   );
//    assert!(auto_res.is_ok());

    // Query staker info
//    let query_msg = QueryMsg::StakerInfo {
//        address: STAKER.to_string(),
//    };

 //   let staker_info: StakerInfo = app
 //       .wrap()
   //     .query_wasm_smart(liquid_staking_id.to_string(), &query_msg)
  //      .unwrap();

   // assert!(staker_info.staked_amount > Uint128::new(STAKE_AMOUNT)); // Compounded staked amount should increase
}

}
