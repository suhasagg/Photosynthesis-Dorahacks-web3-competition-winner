use cosmwasm_std::{
    attr, entry_point, to_binary, Addr, Binary, Coin, CosmosMsg, Deps, DepsMut, Env,
    MessageInfo, Order, Response, StdError, StdResult, Storage, Timestamp, Uint128,
    WasmMsg,
};
use cosmwasm_storage::{Bucket, BucketReadonly, ReadonlyBucket};
use cw20::{Cw20ExecuteMsg, Cw20ReceiveMsg};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::cmp::Ordering as CmpOrdering;

// Constants for storage keys
const CONFIG_KEY: &[u8] = b"config";
const ADS_KEY: &[u8] = b"ads";
const BIDS_KEY_PREFIX: &[u8] = b"bids";
const REWARDS_KEY: &[u8] = b"rewards";
const EPOCH_DURATION: u64 = 86400; // 1 day in seconds

// ====================================
// Configuration Structure
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Config {
    pub admin: Addr,                  // Admin address
    pub reward_token: Addr,           // CW20 token address for rewards
    pub min_bid_amount: Uint128,      // Minimum bid amount
    pub epoch_duration: u64,          // Duration of each reward epoch in seconds
    pub last_reward_distribution: u64, // Timestamp of the last reward distribution
}

impl Config {
    pub fn save(&self, storage: &mut dyn Storage) -> StdResult<()> {
        cosmwasm_storage::singleton(storage, CONFIG_KEY).save(self)
    }

    pub fn load(storage: &dyn Storage) -> StdResult<Self> {
        cosmwasm_storage::singleton_read(storage, CONFIG_KEY).load()
    }
}

// ====================================
// Ad Structure
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Ad {
    pub id: String,
    pub creator: Addr,
    pub image_url: String,
    pub target_url: String,
    pub total_views: u64,
    pub active: bool,
    pub created_at: u64,
    pub updated_at: u64,
}

impl Ad {
    pub fn save(&self, storage: &mut dyn Storage) -> StdResult<()> {
        let mut ads_bucket: Bucket<Ad> = Bucket::new(storage, ADS_KEY);
        ads_bucket.save(self.id.as_bytes(), self)
    }

    pub fn load(storage: &dyn Storage, id: &str) -> StdResult<Self> {
        let ads_bucket: BucketReadonly<Ad> = Bucket::new(storage, ADS_KEY);
        ads_bucket.load(id.as_bytes())
    }

    pub fn may_load(storage: &dyn Storage, id: &str) -> StdResult<Option<Self>> {
        let ads_bucket: BucketReadonly<Ad> = Bucket::new(storage, ADS_KEY);
        ads_bucket.may_load(id.as_bytes())
    }

    pub fn remove(storage: &mut dyn Storage, id: &str) {
        let mut ads_bucket: Bucket<Ad> = Bucket::new(storage, ADS_KEY);
        ads_bucket.remove(id.as_bytes());
    }
}

// ====================================
// Bid Structure
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Bid {
    pub bidder: Addr,
    pub amount: Uint128,
    pub timestamp: u64,
    pub expires_at: u64,
}

impl Bid {
    pub fn save(storage: &mut dyn Storage, ad_id: &str, bid: &Bid) -> StdResult<()> {
        let mut bids_bucket: Bucket<Bid> = Bucket::multilevel(storage, &[BIDS_KEY_PREFIX, ad_id.as_bytes()]);
        bids_bucket.save(bid.bidder.as_bytes(), bid)
    }

    pub fn load(storage: &dyn Storage, ad_id: &str, bidder: &Addr) -> StdResult<Bid> {
        let bids_bucket: BucketReadonly<Bid> = Bucket::multilevel(storage, &[BIDS_KEY_PREFIX, ad_id.as_bytes()]);
        bids_bucket.load(bidder.as_bytes())
    }

    pub fn remove(storage: &mut dyn Storage, ad_id: &str, bidder: &Addr) {
        let mut bids_bucket: Bucket<Bid> = Bucket::multilevel(storage, &[BIDS_KEY_PREFIX, ad_id.as_bytes()]);
        bids_bucket.remove(bidder.as_bytes());
    }

    pub fn load_all(storage: &dyn Storage, ad_id: &str) -> StdResult<Vec<Bid>> {
        let bids_bucket: BucketReadonly<Bid> = Bucket::multilevel(storage, &[BIDS_KEY_PREFIX, ad_id.as_bytes()]);
        let bids = bids_bucket
            .range(None, None, Order::Ascending)
            .map(|item| item.map(|(_, bid)| bid))
            .collect::<StdResult<Vec<Bid>>>()?;
        Ok(bids)
    }
}

// ====================================
// Reward Structure
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Reward {
    pub recipient: Addr,
    pub amount: Uint128,
}

impl Reward {
    pub fn save(storage: &mut dyn Storage, recipient: &Addr, amount: Uint128) -> StdResult<()> {
        let mut rewards_bucket: Bucket<Uint128> = Bucket::new(storage, REWARDS_KEY);
        let existing = rewards_bucket.may_load(recipient.as_bytes())?.unwrap_or_else(Uint128::zero);
        rewards_bucket.save(recipient.as_bytes(), &(existing + amount))
    }

    pub fn load_all(storage: &dyn Storage) -> StdResult<Vec<Reward>> {
        let rewards_bucket: BucketReadonly<Uint128> = Bucket::new(storage, REWARDS_KEY);
        let rewards = rewards_bucket
            .range(None, None, Order::Ascending)
            .map(|item| {
                item.map(|(key, amount)| Reward {
                    recipient: Addr::unchecked(String::from_utf8(key).unwrap()),
                    amount,
                })
            })
            .collect::<StdResult<Vec<Reward>>>()?;
        Ok(rewards)
    }

    pub fn remove_all(storage: &mut dyn Storage) {
        let mut rewards_bucket: Bucket<Uint128> = Bucket::new(storage, REWARDS_KEY);
        rewards_bucket.clear();
    }
}

// ====================================
// InstantiateMsg
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {
    pub admin: String,
    pub reward_token: String,
    pub min_bid_amount: Uint128,
    pub epoch_duration: u64,
}

// ====================================
// ExecuteMsg
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum ExecuteMsg {
    AddAd {
        id: String,
        image_url: String,
        target_url: String,
    },
    DeleteAd {
        id: String,
    },
    PlaceBid {
        ad_id: String,
        amount: Uint128,
        expires_in: u64, // Duration in seconds
    },
    ServeAd {
        ad_id: String,
    },
    DistributeRewards {},
    UpdateConfig {
        admin: Option<String>,
        reward_token: Option<String>,
        min_bid_amount: Option<Uint128>,
        epoch_duration: Option<u64>,
    },
}

// ====================================
// QueryMsg
// ====================================

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    GetAd {
        id: String,
    },
    GetAllAds {},
    GetBids {
        ad_id: String,
    },
    GetConfig {},
    GetRewards {},
}

// ====================================
// Instantiate Entry Point
// ====================================

#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: InstantiateMsg,
) -> StdResult<Response> {
    let config = Config {
        admin: deps.api.addr_validate(&msg.admin)?,
        reward_token: deps.api.addr_validate(&msg.reward_token)?,
        min_bid_amount: msg.min_bid_amount,
        epoch_duration: msg.epoch_duration,
        last_reward_distribution: env.block.time.seconds(),
    };

    config.save(deps.storage)?;

    Ok(Response::new()
        .add_attribute("action", "instantiate")
        .add_attribute("admin", config.admin)
        .add_attribute("reward_token", config.reward_token))
}

// ====================================
// Execute Entry Point
// ====================================

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    match msg {
        ExecuteMsg::AddAd {
            id,
            image_url,
            target_url,
        } => execute_add_ad(deps, env, info, id, image_url, target_url),
        ExecuteMsg::DeleteAd { id } => execute_delete_ad(deps, info, id),
        ExecuteMsg::PlaceBid {
            ad_id,
            amount,
            expires_in,
        } => execute_place_bid(deps, env, info, ad_id, amount, expires_in),
        ExecuteMsg::ServeAd { ad_id } => execute_serve_ad(deps, env, ad_id),
        ExecuteMsg::DistributeRewards {} => execute_distribute_rewards(deps, env, info),
        ExecuteMsg::UpdateConfig {
            admin,
            reward_token,
            min_bid_amount,
            epoch_duration,
        } => execute_update_config(deps, info, admin, reward_token, min_bid_amount, epoch_duration),
    }
}

// ====================================
// Execute Handlers
// ====================================

fn execute_add_ad(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    id: String,
    image_url: String,
    target_url: String,
) -> StdResult<Response> {
    // Validate that the ad does not already exist
    if Ad::may_load(deps.storage, &id)?.is_some() {
        return Err(StdError::generic_err("Ad ID already exists"));
    }

    let ad = Ad {
        id: id.clone(),
        creator: info.sender.clone(),
        image_url,
        target_url,
        total_views: 0,
        active: true,
        created_at: env.block.time.seconds(),
        updated_at: env.block.time.seconds(),
    };

    ad.save(deps.storage)?;

    Ok(Response::new()
        .add_attribute("action", "add_ad")
        .add_attribute("ad_id", id)
        .add_attribute("creator", info.sender))
}

fn execute_delete_ad(deps: DepsMut, info: MessageInfo, id: String) -> StdResult<Response> {
    let ad = Ad::load(deps.storage, &id)?;

    // Only admin or ad creator can delete the ad
    let config = Config::load(deps.storage)?;
    if info.sender != ad.creator && info.sender != config.admin {
        return Err(StdError::unauthorized());
    }

    Ad::remove(deps.storage, &id);

    Ok(Response::new()
        .add_attribute("action", "delete_ad")
        .add_attribute("ad_id", id))
}

fn execute_place_bid(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    ad_id: String,
    amount: Uint128,
    expires_in: u64,
) -> StdResult<Response> {
    let config = Config::load(deps.storage)?;

    // Validate that the ad exists and is active
    let ad = Ad::load(deps.storage, &ad_id)?;
    if !ad.active {
        return Err(StdError::generic_err("Ad is not active"));
    }

    // Validate bid amount
    if amount < config.min_bid_amount {
        return Err(StdError::generic_err("Bid amount is too low"));
    }

    // Create and save the bid
    let bid = Bid {
        bidder: info.sender.clone(),
        amount,
        timestamp: env.block.time.seconds(),
        expires_at: env.block.time.seconds() + expires_in,
    };

    Bid::save(deps.storage, &ad_id, &bid)?;

    Ok(Response::new()
        .add_attribute("action", "place_bid")
        .add_attribute("ad_id", ad_id)
        .add_attribute("bidder", info.sender)
        .add_attribute("amount", amount))
}

fn execute_serve_ad(deps: DepsMut, env: Env, ad_id: String) -> StdResult<Response> {
    let mut ad = Ad::load(deps.storage, &ad_id)?;

    // Fetch all valid bids
    let bids = Bid::load_all(deps.storage, &ad_id)?
        .into_iter()
        .filter(|bid| bid.expires_at > env.block.time.seconds())
        .collect::<Vec<Bid>>();

    if bids.is_empty() {
        return Err(StdError::generic_err("No valid bids for this ad"));
    }

    // Select the highest bid
    let winning_bid = bids
        .iter()
        .max_by(|a, b| a.amount.cmp(&b.amount))
        .unwrap()
        .clone();

    // Increment ad views
    ad.total_views += 1;
    ad.updated_at = env.block.time.seconds();
    ad.save(deps.storage)?;

    // Accumulate rewards
    Reward::save(
        deps.storage,
        &ad.creator,
        winning_bid.amount, 
    )?;

    Ok(Response::new()
        .add_attribute("action", "serve_ad")
        .add_attribute("ad_id", ad_id)
        .add_attribute("bidder", winning_bid.bidder)
        .add_attribute("amount", winning_bid.amount))
}

fn execute_distribute_rewards(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> StdResult<Response> {
    let mut config = Config::load(deps.storage)?;

    // Only admin can distribute rewards
    if info.sender != config.admin {
        return Err(StdError::unauthorized());
    }

    // Check if enough time has passed since the last distribution
    if env.block.time.seconds() < config.last_reward_distribution + config.epoch_duration {
        return Err(StdError::generic_err("Epoch duration has not passed yet"));
    }

    let rewards = Reward::load_all(deps.storage)?;

    if rewards.is_empty() {
        return Err(StdError::generic_err("No rewards to distribute"));
    }

    let mut messages: Vec<CosmosMsg> = vec![];

    for reward in rewards {
        let msg = CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: config.reward_token.to_string(),
            msg: to_binary(&Cw20ExecuteMsg::Transfer {
                recipient: reward.recipient.to_string(),
                amount: reward.amount,
            })?,
            funds: vec![],
        });

        messages.push(msg);
    }

    // Clear rewards after distribution
    Reward::remove_all(deps.storage);

    // Update last reward distribution time
    config.last_reward_distribution = env.block.time.seconds();
    config.save(deps.storage)?;

    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("action", "distribute_rewards"))
}

fn execute_update_config(
    deps: DepsMut,
    info: MessageInfo,
    admin: Option<String>,
    reward_token: Option<String>,
    min_bid_amount: Option<Uint128>,
    epoch_duration: Option<u64>,
) -> StdResult<Response> {
    let mut config = Config::load(deps.storage)?;

    // Only current admin can update config
    if info.sender != config.admin {
        return Err(StdError::unauthorized());
    }

    if let Some(admin) = admin {
        config.admin = deps.api.addr_validate(&admin)?;
    }

    if let Some(reward_token) = reward_token {
        config.reward_token = deps.api.addr_validate(&reward_token)?;
    }

    if let Some(min_bid_amount) = min_bid_amount {
        config.min_bid_amount = min_bid_amount;
    }

    if let Some(epoch_duration) = epoch_duration {
        config.epoch_duration = epoch_duration;
    }

    config.save(deps.storage)?;

    Ok(Response::new().add_attribute("action", "update_config"))
}

// ====================================
// Query Entry Point
// ====================================

#[entry_point]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetAd { id } => to_binary(&query_get_ad(deps, id)?),
        QueryMsg::GetAllAds {} => to_binary(&query_get_all_ads(deps)?),
        QueryMsg::GetBids { ad_id } => to_binary(&query_get_bids(deps, ad_id)?),
        QueryMsg::GetConfig {} => to_binary(&query_get_config(deps)?),
        QueryMsg::GetRewards {} => to_binary(&query_get_rewards(deps)?),
    }
}

// ====================================
// Query Handlers
// ====================================

fn query_get_ad(deps: Deps, id: String) -> StdResult<Ad> {
    let ad = Ad::load(deps.storage, &id)?;
    Ok(ad)
}

fn query_get_all_ads(deps: Deps) -> StdResult<Vec<Ad>> {
    let ads_bucket: BucketReadonly<Ad> = Bucket::new(deps.storage, ADS_KEY);
    let ads = ads_bucket
        .range(None, None, Order::Ascending)
        .map(|item| item.map(|(_, ad)| ad))
        .collect::<StdResult<Vec<Ad>>>()?;
    Ok(ads)
}

fn query_get_bids(deps: Deps, ad_id: String) -> StdResult<Vec<Bid>> {
    let bids = Bid::load_all(deps.storage, &ad_id)?;
    Ok(bids)
}

fn query_get_config(deps: Deps) -> StdResult<Config> {
    let config = Config::load(deps.storage)?;
    Ok(config)
}

fn query_get_rewards(deps: Deps) -> StdResult<Vec<Reward>> {
    let rewards = Reward::load_all(deps.storage)?;
    Ok(rewards)
}

// ====================================
// Tests
// ====================================

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{
        mock_dependencies, mock_env, mock_info, MockApi, MockQuerier, MockStorage,
    };

    #[test]
    fn test_add_and_get_ad() {
        let mut deps = mock_dependencies();
        let env = mock_env();
        let info = mock_info("creator", &[]);

        let instantiate_msg = InstantiateMsg {
            admin: "admin".to_string(),
            reward_token: "reward_token".to_string(),
            min_bid_amount: Uint128::new(100),
            epoch_duration: EPOCH_DURATION,
        };

        let _res = instantiate(deps.as_mut(), env.clone(), info.clone(), instantiate_msg).unwrap();

        let add_ad_msg = ExecuteMsg::AddAd {
            id: "ad1".to_string(),
            image_url: "https://magnite.com/image.png".to_string(),
            target_url: "https://netflix.com".to_string(),
        };

        let _res = execute(deps.as_mut(), env.clone(), info.clone(), add_ad_msg).unwrap();

        let query_msg = QueryMsg::GetAd {
            id: "ad1".to_string(),
        };

        let res = query(deps.as_ref(), env.clone(), query_msg).unwrap();
        let ad: Ad = from_binary(&res).unwrap();

        assert_eq!(ad.id, "ad1");
        assert_eq!(ad.creator, Addr::unchecked("creator"));
        assert_eq!(ad.image_url, "https://magnite.com/image.png");
        assert_eq!(ad.target_url, "https://netflix.com");
    }

}

// ====================================
// Helper Functions
// ====================================

fn from_binary<T: serde::de::DeserializeOwned>(binary: &Binary) -> StdResult<T> {
    Ok(serde_json::from_slice(binary.as_slice())?)
}

