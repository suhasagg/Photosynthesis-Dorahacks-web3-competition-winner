use cosmwasm_std::{Addr, Uint128};
use cosmwasm_storage::{bucket, bucket_read, Bucket, ReadonlyBucket};
use cw20::{Cw20ReceiveMsg, Cw20ExecuteMsg};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use cosmwasm_std::{
    BankMsg, Coin, CosmosMsg, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Uint128,
};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Config {
    pub owner: Addr,
    pub reward_pool: Uint128,
    pub total_staked: Uint128,
    pub st_token_supply: Uint128,
    pub reward_rate: Uint128,  // Annual percentage rate (APR)
    pub slashing_rate: Uint128, // Penalty rate for slashing
    pub previous_redemption_rate: Uint128,
    pub total_unbonded: Uint128, 
    pub module_account_balance: Uint128, 
    pub epoch_info: Option<EpochInfo>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakerInfo {
    pub staked_amount: Uint128,
    pub liquid_tokens: Uint128,
    pub reward_debt: Uint128, // To handle pending rewards
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Validator {
    pub address: Addr,
    pub total_staked: Uint128,
}


#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct EpochInfo {
    pub current_epoch: u64,
    pub last_compounded_epoch: u64,
    pub epoch_duration: u64, // Duration of each epoch in blocks or time
}


pub fn config(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<Config> {
    bucket(storage, b"config")
}

pub fn stakers(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<StakerInfo> {
    bucket(storage, b"stakers")
}

pub fn validators(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<Validator> {9
    bucket(storage, b"validators")
}

// Function to initialize the epochs
pub fn initialize_epochs(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    epoch_duration: u64,
) -> StdResult<Response> {
    let epoch_info = EpochInfo {
        current_epoch: 0,
        last_compounded_epoch: 0,
        epoch_duration,
    };

    let mut config = config(deps.storage).load()?;
    config.epoch_info = Some(epoch_info);
    config(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "initialize_epochs")
        .add_attribute("epoch_duration", epoch_duration.to_string()))
}


pub fn update_config(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    total_unbonded: Uint128,
    total_staked: Uint128,
    module_account_balance: Uint128,
    reward_pool: Option<Uint128>, // Optional
    reward_rate: Option<Uint128>, // Optional
    slashing_rate: Option<Uint128>, // Optional
) -> StdResult<Response> {
    // Load the existing config from storage
    let mut config = config(deps.storage).load()?;

    // Ensure that only the owner can update the config
    if info.sender != config.owner {
        return Err(StdError::unauthorized());
    }

    // Update the config fields with the provided parameters
    config.total_unbonded = total_unbonded;
    config.total_staked = total_staked;
    config.module_account_balance = module_account_balance;

    // Optionally update the other fields if provided
    if let Some(reward_pool_value) = reward_pool {
        config.reward_pool = reward_pool_value;
    }
    if let Some(reward_rate_value) = reward_rate {
        config.reward_rate = reward_rate_value;
    }
    if let Some(slashing_rate_value) = slashing_rate {
        config.slashing_rate = slashing_rate_value;
    }

    // Save the updated config back to storage
    config(deps.storage).save(&config)?;

    // Return a response indicating the update was successful
    Ok(Response::new()
        .add_attribute("method", "update_config")
        .add_attribute("total_unbonded", total_unbonded.to_string())
        .add_attribute("total_staked", total_staked.to_string())
        .add_attribute("module_account_balance", module_account_balance.to_string())
        .add_attribute("total_value", total_value.to_string()))
}


pub fn calculate_redemption_rate(deps: DepsMut) -> StdResult<Uint128> {
    let mut config = config(deps.storage).load()?;

    let total_value = config.total_unbonded + config.total_staked + config.module_account_balance;
    if config.st_token_supply.is_zero() {
        return Err(StdError::generic_err("stToken supply cannot be zero"));
    }

    let new_redemption_rate = total_value / config.st_token_supply;

    // Update the previous redemption rate
    config.previous_redemption_rate = new_redemption_rate;
    config(deps.storage).save(&config)?;

    Ok(new_redemption_rate)
}


pub fn stake(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    amount: Uint128,
    validator_addr: Addr,
) -> StdResult<Response> {
    if amount.is_zero() {
        return Err(StdError::generic_err("Amount must be greater than zero"));
    }

    let mut staker_info = stakers(deps.storage).load(info.sender.as_bytes()).unwrap_or_else(|_| StakerInfo {
        staked_amount: Uint128::zero(),
        liquid_tokens: Uint128::zero(),
        reward_debt: Uint128::zero(),
    });

    let mut validator = validators(deps.storage).load(validator_addr.as_bytes()).unwrap_or_else(|_| Validator {
        address: validator_addr.clone(),
        total_staked: Uint128::zero(),
    });

    staker_info.staked_amount += amount;

    // Calculate and update the redemption rate
    let redemption_rate = calculate_redemption_rate(deps.branch())?;
    
    // Calculate the number of liquid tokens to mint based on the redemption rate
    let liquid_tokens_to_mint = amount * redemption_rate;
    staker_info.liquid_tokens += liquid_tokens_to_mint;

    // Update validator's total staked amount
    validator.total_staked += amount;
    validators(deps.storage).save(validator_addr.as_bytes(), &validator)?;

    // Save staker info
    stakers(deps.storage).save(info.sender.as_bytes(), &staker_info)?;

    // Mint liquid tokens
    let mint_msg = CosmosMsg::Bank(BankMsg::Mint {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "liquid_token".to_string(),
            amount: liquid_tokens_to_mint,
        }],
    });

    // Update total staked and stToken supply in config
    let mut config = config(deps.storage).load()?;
    config.total_staked += amount;
    config.st_token_supply += liquid_tokens_to_mint;
    config(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_message(mint_msg)
        .add_attribute("method", "stake")
        .add_attribute("amount", amount.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string())
        .add_attribute("validator", validator.address.to_string()))
}


pub fn distribute_rewards(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    reward_amount: Uint128,
) -> StdResult<Response> {
    let mut config = config(deps.storage).load()?;
    let reward_per_share = reward_amount / config.total_staked;

    config.reward_rate += reward_per_share;
    config.reward_pool += reward_amount;
    config(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "distribute_rewards")
        .add_attribute("reward_amount", reward_amount.to_string()))
}


pub fn slash(
    deps: DepsMut,
    _env: Env,
    validator_addr: Addr,
    slash_amount: Uint128,
) -> StdResult<Response> {
    let mut validator = validators(deps.storage).load(validator_addr.as_bytes())?;
    let mut config = config(deps.storage).load()?;

    let penalty = slash_amount * config.slashing_rate / Uint128::from(100u128);
    validator.total_staked -= penalty;

    // Adjust global state
    config.total_staked -= penalty;
    config.reward_pool -= penalty; // Reducing reward pool by the penalty
    validators(deps.storage).save(validator_addr.as_bytes(), &validator)?;
    config(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "slash")
        .add_attribute("validator", validator_addr.to_string())
        .add_attribute("penalty", penalty.to_string()))
}

pub fn withdraw(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    liquid_token_amount: Uint128,
) -> StdResult<Response> {
    let mut staker_info = stakers(deps.storage).load(info.sender.as_bytes())?;
    if staker_info.liquid_tokens < liquid_token_amount {
        return Err(StdError::generic_err("Insufficient liquid tokens"));
    }

    // Calculate the redemption rate
    let redemption_rate = calculate_redemption_rate(deps.branch())?;

    // Calculate the amount of staked assets to return
    let original_staked = liquid_token_amount * redemption_rate;

    staker_info.staked_amount -= original_staked;
    staker_info.liquid_tokens -= liquid_token_amount;
    stakers(deps.storage).save(info.sender.as_bytes(), &staker_info)?;

    // Update total staked and stToken supply in config
    let mut config = config(deps.storage).load()?;
    config.total_staked -= original_staked;
    config.st_token_supply -= liquid_token_amount;
    config(deps.storage).save(&config)?;

    let send_msg = BankMsg::Send {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "staked_token".to_string(),
            amount: original_staked,
        }],
    };

    Ok(Response::new()
        .add_message(send_msg)
        .add_attribute("method", "withdraw")
        .add_attribute("liquid_token_amount", liquid_token_amount.to_string())
        .add_attribute("redeemed_amount", original_staked.to_string()))
}



pub fn auto_compound(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> StdResult<Response> {
    let mut config = config(deps.storage).load()?;
    let epoch_info = config.epoch_info.as_mut().ok_or_else(|| StdError::generic_err("Epoch info not initialized"))?;

    // Check if it's time to run the auto-compounding based on the epoch duration
    if env.block.height < epoch_info.current_epoch + epoch_info.epoch_duration {
        return Err(StdError::generic_err("Not time for auto-compounding yet"));
    }

    // Update the current epoch
     epoch_info.last_compounded_epoch = env.block.height;
     epoch_info.current_epoch += 1;
     config.epoch_info = Some(*epoch_info);
     config(deps.storage).save(&config)?;


    // Perform auto-compounding for the current staker
    let mut staker_info = stakers(deps.storage).load(info.sender.as_bytes())?;
    if staker_info.staked_amount.is_zero() {
        return Err(StdError::generic_err("No staked amount to compound."));
    }

    // Calculate pending rewards
    let pending_rewards = (staker_info.staked_amount * config.reward_rate) - staker_info.reward_debt;

    if pending_rewards.is_zero() {
        return Err(StdError::generic_err("No rewards available to compound."));
    }

    // Calculate the redemption rate
    let redemption_rate = calculate_redemption_rate(deps.branch())?;

    // Calculate the amount of liquid tokens to mint based on the pending rewards
    let liquid_tokens_to_mint = pending_rewards * redemption_rate;

    // Update staker's staked amount and liquid tokens
    staker_info.staked_amount += pending_rewards;
    staker_info.liquid_tokens += liquid_tokens_to_mint;
    staker_info.reward_debt += pending_rewards;

    // Update the validator's total staked amount
    let mut validator_info = validators(deps.storage).load(info.sender.as_bytes())?;
    validator_info.total_staked += pending_rewards;

    // Save the updated staker and validator information
    stakers(deps.storage).save(info.sender.as_bytes(), &staker_info)?;
    validators(deps.storage).save(info.sender.as_bytes(), &validator_info)?;

    // Mint additional liquid tokens for the staker
    let mint_msg = CosmosMsg::Bank(BankMsg::Mint {
        to_address: info.sender.to_string(),
        amount: vec![Coin {
            denom: "liquid_token".to_string(),
            amount: liquid_tokens_to_mint,
        }],
    });

    // Update the protocol's total staked amount and stToken supply
    config.total_staked += pending_rewards;
    config.st_token_supply += liquid_tokens_to_mint;
    config(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_message(mint_msg)
        .add_attribute("method", "auto_compound")
        .add_attribute("compounded_amount", pending_rewards.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string()))
}



pub fn submit_proposal(
    deps: DepsMut,
    info: MessageInfo,
    proposal_id: u64,
    description: String,
) -> StdResult<Response> {
    // TODO: proposal submission logic
    Ok(Response::new()
        .add_attribute("method", "submit_proposal")
        .add_attribute("proposal_id", proposal_id.to_string())
        .add_attribute("description", description))
}

pub fn vote_proposal(
    deps: DepsMut,
    info: MessageInfo,
    proposal_id: u64,
    vote: bool,
) -> StdResult<Response> {
    // TODO: voting logic
    Ok(Response::new()
        .add_attribute("method", "vote_proposal")
        .add_attribute("proposal_id", proposal_id.to_string())
        .add_attribute("vote", vote.to_string()))
}


pub fn handle_ibc_stake(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    ibc_channel: String,
    amount: Uint128,
) -> StdResult<Response> {
    let redemption_rate = calculate_redemption_rate(deps.branch())?;

    let liquid_tokens_to_mint = amount * redemption_rate;

    let ibc_stake_msg = format!(
        "Staking {} tokens from {} via IBC channel {}",
        amount, info.sender, ibc_channel
    );

    // Update stToken supply in config
    let mut config = config(deps.storage).load()?;
    config.st_token_supply += liquid_tokens_to_mint;
    config.deps.save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "handle_ibc_stake")
        .add_attribute("ibc_channel", ibc_channel)
        .add_attribute("staker", info.sender.to_string())
        .add_attribute("amount", amount.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string())
        .add_message(CosmosMsg::Custom(ibc_stake_msg))) // IBC packet sending logic
}


pub fn handle_ibc_withdraw(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    ibc_channel: String,
    liquid_token_amount: Uint128,
) -> StdResult<Response> {
    let redemption_rate = calculate_redemption_rate(deps.branch())?;
    let original_staked = liquid_token_amount * redemption_rate;

    let mut staker_info = stakers(deps.storage).load(info.sender.as_bytes())?;
    if staker_info.liquid_tokens < liquid_token_amount {
        return Err(StdError::generic_err("Insufficient liquid tokens"));
    }

    staker_info.liquid_tokens -= liquid_token_amount;
    stakers(deps.storage).save(info.sender.as_bytes(), &staker_info)?;

    let ibc_withdraw_msg = format!(
        "Withdrawing {} liquid tokens from {} via IBC channel {}",
        liquid_token_amount, info.sender, ibc_channel
    );

    // Update total staked and stToken supply in config
    let mut config = config(deps.storage).load()?;
    config.total_staked -= original_staked;
    config.st_token_supply -= liquid_token_amount;
    config.deps.save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "handle_ibc_withdraw")
        .add_attribute("ibc_channel", ibc_channel)
        .add_attribute("staker", info.sender.to_string())
        .add_attribute("liquid_token_amount", liquid_token_amount.to_string())
        .add_attribute("redeemed_amount", original_staked.to_string())
        .add_message(CosmosMsg::Custom(ibc_withdraw_msg))) // IBC withdrawal logic
}



pub fn receive_ibc_stake(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    staker: Addr,
    validator: Addr,
    amount: Uint128,
) -> StdResult<Response> {
    // Verify the packet's authenticity and integrity
    // Record the stake in the local state and associate it with the correct validator

    let mut staker_info = stakers(deps.storage).load(staker.as_bytes()).unwrap_or_else(|_| StakerInfo {
        staked_amount: Uint128::zero(),
        liquid_tokens: Uint128::zero(),
        reward_debt: Uint128::zero(),
    });

    let mut validator_info = validators(deps.storage).load(validator.as_bytes()).unwrap_or_else(|_| Validator {
        address: validator.clone(),
        total_staked: Uint128::zero(),
    });

    // Calculate the redemption rate
    let redemption_rate = calculate_redemption_rate(deps.branch())?;

    // Calculate the number of liquid tokens to mint based on the redemption rate
    let liquid_tokens_to_mint = amount * redemption_rate;

    // Update the staked amount and liquid tokens for the staker
    staker_info.staked_amount += amount;
    staker_info.liquid_tokens += liquid_tokens_to_mint;

    // Update the total staked amount for the validator
    validator_info.total_staked += amount;

    // Save the updated staker and validator information
    stakers(deps.storage).save(staker.as_bytes(), &staker_info)?;
    validators(deps.storage).save(validator.as_bytes(), &validator_info)?;

    // Mint liquid tokens to the staker
    let mint_msg = CosmosMsg::Bank(BankMsg::Mint {
        to_address: staker.to_string(),
        amount: vec![Coin {
            denom: "liquid_token".to_string(),
            amount: liquid_tokens_to_mint,
        }],
    });

    // Update the stToken supply in the configuration
    let mut config = config(deps.storage).load()?;
    config.st_token_supply += liquid_tokens_to_mint;
    config.deps.save(&config)?;

    // Return a response indicating successful processing of the IBC stake
    Ok(Response::new()
        .add_message(mint_msg)
        .add_attribute("method", "receive_ibc_stake")
        .add_attribute("staker", staker.to_string())
        .add_attribute("validator", validator.to_string())
        .add_attribute("amount", amount.to_string())
        .add_attribute("liquid_tokens_minted", liquid_tokens_to_mint.to_string()))
}




pub fn receive_ibc_withdraw(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    staker: Addr,
    liquid_token_amount: Uint128,
) -> StdResult<Response> {
    // Verify the packet's authenticity and integrity
    // Calculate the redemption rate
    let redemption_rate = calculate_redemption_rate(deps.branch())?;

    // Calculate the amount of staked assets to return based on the redemption rate
    let original_staked = liquid_token_amount * redemption_rate;

    // Create a BankMsg to send the staked assets to the staker's address
    let send_msg = BankMsg::Send {
        to_address: staker.to_string(),
        amount: vec![Coin {
            denom: "staked_token".to_string(),
            amount: original_staked,
        }],
    };

    // Update the stToken supply in the configuration
    let mut config = config(deps.storage).load()?;
    config.st_token_supply -= liquid_token_amount;
    config.total_staked -= original_staked;
    config.deps.save(&config)?;

    // Return a response indicating successful processing of the IBC withdrawal
    Ok(Response::new()
        .add_message(send_msg)
        .add_attribute("method", "receive_ibc_withdraw")
        .add_attribute("staker", staker.to_string())
        .add_attribute("liquid_token_amount", liquid_token_amount.to_string())
        .add_attribute("redeemed_amount", original_staked.to_string()))
}
