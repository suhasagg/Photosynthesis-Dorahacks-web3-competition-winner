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
    pub reward_rate: Uint128,  // Annual percentage rate (APR)
    pub slashing_rate: Uint128, // Penalty rate for slashing
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

pub fn config(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<Config> {
    bucket(storage, b"config")
}

pub fn stakers(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<StakerInfo> {
    bucket(storage, b"stakers")
}

pub fn validators(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<Validator> {9
    bucket(storage, b"validators")
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


