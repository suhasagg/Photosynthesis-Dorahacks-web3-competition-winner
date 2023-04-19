use cosmwasm_std::{
    attr, to_binary, Addr, Binary, CosmosMsg, Deps, DepsMut, Empty, Env, Event, MessageInfo,
    QueryRequest, Response, StdError, StdResult, Uint128,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage, ReadonlySingleton, Singleton};
use serde::{Deserialize, Serialize};

use cw20::{Cw20ExecuteMsg, Cw20ReceiveMsg};
use cw_storage_plus::{Item, Map};
use schemars::JsonSchema;
use serde::export::fmt::Debug;
use std::ops::Deref;

const STATE_KEY: &[u8] = b"state";
const AD_PREFIX: &[u8] = b"ad";
const TOTAL_VIEWS_KEY: &[u8] = b"total_views";

pub const CW20_STAKING_CONTRACT: &str = "staking_contract";
pub const CW20_STAKING_BALANCE: &str = "staking_balance";

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Ad {
    pub id: String,
    pub image_url: String,
    pub target_url: String,
    pub views: u64,
    pub reward_address: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub ads: Vec<Ad>,
    pub total_views: u64,
    pub plt_address: String,
}

impl Default for State {
    fn default() -> Self {
        State {
            ads: vec![],
            total_views: 0,
            plt_address: "".to_string(),
        }
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum HandleMsg {
    AddAd {
        id: String,
        image_url: String,
        target_url: String,
        reward_address: String,
    },
    ServeAd {
        id: String,
    },
    DeleteAd {
        id: String,
    },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    Ad {
        id: String,
    },
    Ads,
    TotalViews,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum EpochMsg {
    DistributeRewards,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct QueryAdResponse {
    pub id: String,
    pub image_url: String,
    pub target_url: String,
    pub views: u64,
    pub reward_address: String,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct QueryAllAdsResponse {
    pub ads: Vec<QueryAdResponse>,
}

pub fn handle(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    match msg {
        HandleMsg::AddAd {
            id,
            image_url,
            target_url,
            reward_address,
        } => add_ad(deps, env, info, id, image_url, target_url, reward_address),
        HandleMsg::ServeAd { id } => serve_ad(deps, env, id),
        HandleMsg::DeleteAd { id } => delete_ad(deps, env, info, id),
    }
}

pub fn add_ad(
deps: DepsMut,
env: Env,
info: MessageInfo,
id: String,
image_url: String,
target_url: String,
reward_address: String,
plt_amount: Uint128,
) -> StdResult<Response> {
let ad = Ad {
id: id.clone(),
image_url,
target_url,
views: 0,
reward_address,
};
let mut state = State::load(deps.storage)?;
state.ads.push(ad.clone());
State::save(&mut deps.storage, &state)?;


let mut attributes = vec![attr("action", "add_ad"), attr("reward_address", reward_address)];
attributes.push(attr("id", id.clone()));
attributes.push(attr("image_url", image_url.clone()));
attributes.push(attr("target_url", target_url.clone()));
let event = Event::new("add_ad").add_attributes(attributes);

// Reward the caller for adding an ad
let reward_recipient = env.message.sender.clone();
let reward_msg = Cw20ExecuteMsg::Transfer {
    recipient: reward_recipient.into(),
    amount: plt_amount,
};
let reward_attrs = vec![    attr("action", "reward"),    attr("recipient", reward_recipient),    attr("amount", plt_amount),];
let reward_event = Event::new("reward").add_attributes(reward_attrs);

// Stake the required amount of PLT tokens
let plt_address = deps.api.addr_validate(&PLT_ADDRESS)?;
let staking_msg = Cw20ExecuteMsg::Send {
    contract: plt_address.into(),
    amount: plt_amount,
    msg: to_binary(&LiquidStakingExecuteMsg::StakeTokens {
        staker_addr: env.contract.address.into(),
    })?,
};

// Emit a staking event
let staking_attrs = vec![    attr("action", "staking"),    attr("contract_address", plt_address),    attr("staker_address", env.contract.address),    attr("amount", plt_amount),];
let staking_event = Event::new("staking").add_attributes(staking_attrs);

Ok(Response::new()
    .add_messages(vec![
        CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: env.contract.address.into(),
            msg: to_binary(&HandleMsg::ServeAd { id: id.clone() })?,
            funds: vec![],
        }),
        CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: plt_address.into(),
            msg: to_binary(&staking_msg)?,
            funds: vec![],
        }),
        CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: PLT_REWARDS_CONTRACT.to_string().into(),
            msg: to_binary(&Cw20ExecuteMsg::Send {
                contract: plt_address.into(),
                amount: plt_amount,
                msg: to_binary(&EpochMsg::DistributeRewards)?,
            })?,
            funds: vec![],
        }),
    ])
    .add_events(vec![event, reward_event, staking_event]))
}

fn delete_ad(storage: &mut dyn Storage, ad_id: String) {
let mut state = State::load(storage).unwrap();
let ad_idx = state.ads.iter().position(|ad| ad.id == ad_id).unwrap();
state.ads.remove(ad_idx);
State::save(storage, &state).unwrap();
}

fn get_all_ad_ids(deps: &Deps) -> StdResult<Vec<String>> {
let state = State::load(deps.storage)?;
Ok(state.ads.iter().map(|ad| ad.id.clone()).collect())
}


fn get_all_ad_ids(deps: &Deps) -> StdResult<Vec<String>> {
let state = State::load(deps.storage)?;
Ok(state.ads.iter().map(|ad| ad.id.clone()).collect())
}

pub fn distribute_rewards(
deps: DepsMut,
env: Env,
_msg: EpochMsg,
) -> Result<Response<Empty>, ContractError> {
let mut state = State::load(deps.storage)?;


let plt_address = env.contract.address.clone();
let plt_balance = get_plt_balance(deps.as_ref(), plt_address.clone())?;

// Calculate reward amount per view
let reward_amount = plt_balance.checked_div(state.total_views).unwrap_or_default();

// Iterate through all ads and distribute rewards
for mut ad in state.ads.iter_mut() {
    let ad_reward_amount = ad.views.checked_mul(reward_amount).unwrap_or_default();
    ad.views = 0;

    // Send rewards to the reward address specified by the ad creator
    let reward_address = deps.api.addr_validate(&ad.reward_address)?;
    let reward_msg = Cw20ExecuteMsg::Transfer {
        recipient: reward_address.into(),
        amount: ad_reward_amount.into(),
    };
    let reward_attrs = vec![        attr("action", "reward"),        attr("recipient", ad.reward_address.clone()),        attr("amount", ad_reward_amount),    ];
    let reward_event = Event::new("reward").add_attributes(reward_attrs);
    let res = execute(
        deps.as_ref(),
        env.clone(),
        ExecuteMsg::Cw20(reward_msg),
        vec![reward_event],
    )?;

    // Check if the reward transfer failed
    if !res.is_ok() {
        ad.views = ad_reward_amount.checked_div(reward_amount).unwrap_or_default();
    }
}

// Save state
State::save(deps.storage, &state)?;

Ok(Response::new().add_attribute("action", "distribute_rewards"))
}

// Queries the total number of ad views
pub fn query_total_views(deps: Deps) -> StdResult<QueryTotalViewsResponse> {
let state = State::load(deps.storage)?;
Ok(QueryTotalViewsResponse {
total_views: state.total_views,
})
}

// Queries the ad with the given ID
pub fn query_ad(deps: Deps, id: String) -> StdResult<QueryAdResponse> {
let ad = get_ad(deps.as_ref(), id)?;
Ok(QueryAdResponse {
id: ad.id,
image_url: ad.image_url,
target_url: ad.target_url,
views: ad.views,
reward_address: ad.reward_address,
})
}

// Queries all ads
pub fn query_all_ads(deps: Deps) -> StdResult<QueryAllAdsResponse> {
let ad_ids = get_all_ad_ids(deps.as_ref())?;
let mut ads: Vec<QueryAdResponse> = Vec::new();
for id in ad_ids {
let ad = query_ad(deps.as_ref(), id.clone())?;
ads.push(ad);
}
Ok(QueryAllAdsResponse { ads })
}

// Helper functions go here

fn get_all_ad_ids(deps: &Deps) -> StdResult<Vec<String>> {
let state = State::load(deps.storage)?;
Ok(state.ads.iter().map(|ad| ad.id.clone()).collect())
}

fn get_ad(deps: &Deps, id: String) -> StdResult<Ad> {
let state = State::load(deps.storage)?;
for ad in state.ads.iter() {
if ad.id == id {
return Ok(ad.clone());
}
}
Err(ContractError::AdNotFound {})
}

fn delete_ad(storage: &mut dyn Storage, id: String) {
let mut state = State::load(storage).unwrap();
let ad_index = state.ads.iter().position

if ad.id == id {
return Ok(ad.clone());
}
}


pub fn delete_ad(storage: &mut dyn Storage, id: String) {
let mut state = State::load(storage).unwrap();
let ad_index = state.ads.iter().position(|ad| ad.id == id).unwrap();
state.ads.remove(ad_index);
State::save(storage, &state).unwrap();
}

fn query_ad(deps: Deps, id: String) -> StdResult<QueryAdResponse> {
let ad = get_ad(deps.as_ref(), id)?;
Ok(QueryAdResponse {
id: ad.id,
image_url: ad.image_url,
target_url: ad.target_url,
views: ad.views,
reward_address: ad.reward_address,
})
}

fn query_all_ads(deps: Deps) -> StdResult<QueryAllAdsResponse> {
let ad_ids = get_all_ad_ids(deps.as_ref())?;
let ads: Vec<QueryAdResponse> = ad_ids
.iter()
.map(|id| query_ad(deps.as_ref(), id.to_string()).unwrap())
.collect();
Ok(QueryAllAdsResponse { ads })
}

fn query_total_views(deps: Deps) -> StdResult<TotalViewsResponse> {
let state = State::load(deps.storage)?;
Ok(TotalViewsResponse {
total_views: state.total_views,
})
}

fn distribute_rewards(deps: Deps, env: Env) -> StdResult<Response> {
let state = State::load(deps.storage)?;
let mut messages: Vec<CosmosMsg> = vec![];
let mut total_rewards: u64 = 0;
for ad in &state.ads {
if ad.views > 0 {
let reward_amount = (env.message.sent_funds[0].amount.u128() / ad.views) as u64;
total_rewards += reward_amount * ad.views;
messages.push(CosmosMsg::Wasm(WasmMsg::Execute {
contract_addr: ad.reward_address.clone(),
msg: to_binary(&Cw20ExecuteMsg::Transfer {
recipient: env.message.sender.clone().into(),
amount: Uint128(reward_amount),
})
.unwrap(),
send: vec![],
}));
}
}
if total_rewards > 0 {
let event = Event::new("distribute_rewards").add_attribute(
"total_rewards",
total_rewards.to_string(),
);
Ok(Response::new()
.add_messages(messages)
.add_event(event))
} else {
Err(StdError::generic_err("No rewards to distribute"))
}
}
