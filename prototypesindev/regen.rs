use cosmwasm_std::{
    Addr, BankMsg, Coin, CosmosMsg, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Uint128,
};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use cosmwasm_storage::{bucket, bucket_read, Bucket, ReadonlyBucket};

// Configuration structure
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Config {
    pub owner: Addr,
    pub carbon_credit_denom: String,  // Denomination for carbon credits
    pub staking_token_denom: String,  // Denomination for staking tokens
    pub total_staked: Uint128,        // Total tokens staked
}

// Structure to track carbon credits
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct CarbonCredit {
    pub owner: Addr,
    pub amount: Uint128,
}

// Structure to track ecological assets
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct EcologicalAsset {
    pub id: u64,
    pub owner: Addr,
    pub metadata: String,
    pub value: Uint128,
}

// Structure to track staking information
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakingInfo {
    pub staker_address: Addr,
    pub staked_amount: Uint128,
}

// the storage buckets for configuration, carbon credits, ecological assets, and staking info
pub fn config_write(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<Config> {
    bucket(storage, b"config")
}

pub fn config_read(storage: &dyn cosmwasm_std::Storage) -> ReadonlyBucket<Config> {
    bucket_read(storage, b"config")
}

pub fn carbon_credits_write(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<CarbonCredit> {
    bucket(storage, b"carbon_credits")
}

pub fn carbon_credits_read(storage: &dyn cosmwasm_std::Storage) -> ReadonlyBucket<CarbonCredit> {
    bucket_read(storage, b"carbon_credits")
}

pub fn ecological_assets_write(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<EcologicalAsset> {
    bucket(storage, b"ecological_assets")
}

pub fn ecological_assets_read(storage: &dyn cosmwasm_std::Storage) -> ReadonlyBucket<EcologicalAsset> {
    bucket_read(storage, b"ecological_assets")
}

pub fn staking_info_write(storage: &mut dyn cosmwasm_std::Storage) -> Bucket<StakingInfo> {
    bucket(storage, b"staking_info")
}

pub fn staking_info_read(storage: &dyn cosmwasm_std::Storage) -> ReadonlyBucket<StakingInfo> {
    bucket_read(storage, b"staking_info")
}

// Instantiation function to initialize the contract
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    carbon_credit_denom: String,
    staking_token_denom: String,
) -> StdResult<Response> {
    let config = Config {
        owner: info.sender.clone(),
        carbon_credit_denom,
        staking_token_denom,
        total_staked: Uint128::zero(),
    };

    config_write(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("owner", info.sender))
}

// Function to issue carbon credits
pub fn issue_carbon_credits(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    amount: Uint128,
) -> StdResult<Response> {
    let config = config_read(deps.storage).load()?;

    let mut carbon_credit = carbon_credits_read(deps.storage)
        .load(info.sender.as_bytes())
        .unwrap_or(CarbonCredit {
            owner: info.sender.clone(),
            amount: Uint128::zero(),
        });

    carbon_credit.amount += amount;

    carbon_credits_write(deps.storage).save(info.sender.as_bytes(), &carbon_credit)?;

    Ok(Response::new()
        .add_attribute("method", "issue_carbon_credits")
        .add_attribute("owner", info.sender.to_string())
        .add_attribute("amount", amount.to_string()))
}

// Function to transfer carbon credits
pub fn transfer_carbon_credits(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    recipient: Addr,
    amount: Uint128,
) -> StdResult<Response> {
    let mut carbon_credit = carbon_credits_read(deps.storage).load(info.sender.as_bytes())?;
    
    if carbon_credit.amount < amount {
        return Err(StdError::generic_err("Insufficient carbon credits"));
    }
    
    carbon_credit.amount -= amount;
    carbon_credits_write(deps.storage).save(info.sender.as_bytes(), &carbon_credit)?;

    let mut recipient_credit = carbon_credits_read(deps.storage)
        .load(recipient.as_bytes())
        .unwrap_or(CarbonCredit {
            owner: recipient.clone(),
            amount: Uint128::zero(),
        });
    
    recipient_credit.amount += amount;
    carbon_credits_write(deps.storage).save(recipient.as_bytes(), &recipient_credit)?;

    Ok(Response::new()
        .add_attribute("method", "transfer_carbon_credits")
        .add_attribute("from", info.sender.to_string())
        .add_attribute("to", recipient.to_string())
        .add_attribute("amount", amount.to_string()))
}

// Function to retire (burn) carbon credits
pub fn retire_carbon_credits(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    amount: Uint128,
) -> StdResult<Response> {
    let mut carbon_credit = carbon_credits_read(deps.storage).load(info.sender.as_bytes())?;

    if carbon_credit.amount < amount {
        return Err(StdError::generic_err("Insufficient carbon credits to retire"));
    }

    carbon_credit.amount -= amount;
    carbon_credits_write(deps.storage).save(info.sender.as_bytes(), &carbon_credit)?;

    Ok(Response::new()
        .add_attribute("method", "retire_carbon_credits")
        .add_attribute("owner", info.sender.to_string())
        .add_attribute("amount_retired", amount.to_string()))
}

// Function to create an ecological asset
pub fn create_ecological_asset(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    id: u64,
    metadata: String,
    value: Uint128,
) -> StdResult<Response> {
    let asset = EcologicalAsset {
        id,
        owner: info.sender.clone(),
        metadata,
        value,
    };

    ecological_assets_write(deps.storage).save(&id.to_be_bytes(), &asset)?;

    Ok(Response::new()
        .add_attribute("method", "create_ecological_asset")
        .add_attribute("owner", info.sender.to_string())
        .add_attribute("asset_id", id.to_string())
        .add_attribute("value", value.to_string()))
}

// Function to transfer an ecological asset
pub fn transfer_ecological_asset(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    recipient: Addr,
    id: u64,
) -> StdResult<Response> {
    let mut asset = ecological_assets_read(deps.storage).load(&id.to_be_bytes())?;

    if asset.owner != info.sender {
        return Err(StdError::unauthorized());
    }

    asset.owner = recipient.clone();
    ecological_assets_write(deps.storage).save(&id.to_be_bytes(), &asset)?;

    Ok(Response::new()
        .add_attribute("method", "transfer_ecological_asset")
        .add_attribute("from", info.sender.to_string())
        .add_attribute("to", recipient.to_string())
        .add_attribute("asset_id", id.to_string()))
}

// Function to stake tokens
pub fn stake_tokens(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    amount: Uint128,
) -> StdResult<Response> {
    let mut config = config_read(deps.storage).load()?;
    config.total_staked += amount;

    let mut staking_info = staking_info_read(deps.storage)
        .load(info.sender.as_bytes())
        .unwrap_or(StakingInfo {
            staker_address: info.sender.clone(),
            staked_amount: Uint128::zero(),
        });

    staking_info.staked_amount += amount;

    staking_info_write(deps.storage).save(info.sender.as_bytes(), &staking_info)?;
    config_write(deps.storage).save(&config)?;

    Ok(Response::new()
        .add_attribute("method", "stake_tokens")
        .add_attribute("staker", info.sender.to_string())
        .add_attribute("amount_staked", amount.to_string()))
}

pub fn distribute_staking_rewards(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    total_rewards: Uint128,
) -> StdResult<Response> {
    let config = config_read(deps.storage).load()?;

    // Collect all stakers' information
    let stakers: Vec<StakingInfo> = staking_info_read(deps.storage)
        .range(None, None, cosmwasm_std::Order::Ascending)
        .map(|item| item.unwrap().1) 
        .collect();

    // List to hold the reward distribution messages
    let mut messages: Vec<CosmosMsg> = vec![];

    // Distribute rewards based on the staked amount of each staker
    for staker in stakers {
        let reward = if config.total_staked.is_zero() {
            Uint128::zero()
        } else {
            total_rewards * staker.staked_amount / config.total_staked
        };

        // Only send rewards if the calculated reward is non-zero
        if !reward.is_zero() {
            let reward_msg = BankMsg::Send {
                to_address: staker.staker_address.to_string(),
                amount: vec![Coin {
                    denom: config.staking_token_denom.clone(),
                    amount: reward,
                }],
            };
            messages.push(CosmosMsg::Bank(reward_msg));
        }
    }

    // Create the response including the reward distribution messages
    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("method", "distribute_staking_rewards")
        .add_attribute("total_rewards", total_rewards.to_string()))
}

