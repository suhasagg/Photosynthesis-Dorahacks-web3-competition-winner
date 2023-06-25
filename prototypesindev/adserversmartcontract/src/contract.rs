use cosmwasm_std::{
    attr, to_binary,Binary, CosmosMsg, Deps, DepsMut, Empty, Env, Event, MessageInfo,
    QueryRequest, Response, StdError, Uint128,
};
use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
use cosmwasm_std::{coins, Addr, BlockInfo, ContractInfo, Timestamp, TransactionInfo};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage, ReadonlySingleton, Singleton};
use serde::{Deserialize, Serialize};
use cosmwasm_std::WasmMsg;
use cw20::{Cw20ExecuteMsg, Cw20ReceiveMsg};
use cw_storage_plus::{Item, Map};
use schemars::JsonSchema;
use std::fmt::Debug;
use std::ops::Deref;
use cosmwasm_std::Storage;
use cosmwasm_storage::singleton;
use cosmwasm_std::to_vec;
use cosmwasm_std::{StdResult, from_slice};
use cosmwasm_storage::singleton_read;


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

pub struct AdserverState {
    ads: Vec<Ad>,
}

impl AdserverState {
    fn serve_ad(&mut self, ad_id: &String) -> Option<&mut Ad> {
        self.ads.iter_mut().find(|ad| ad.id == *ad_id)
    }

    fn delete_ad(&mut self, ad_id: &String) -> bool {
        if let Some(index) = self.ads.iter().position(|ad| ad.id == *ad_id) {
            self.ads.remove(index);
            true
        } else {
            false
        }
    }
}



#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct TotalViewsResponse {
    pub total_views: u64,
}




#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub ads: Vec<Ad>,
    pub total_views: u64,
    pub plt_address: String,
}


impl State {
    pub fn save(&self, storage: &mut dyn Storage) -> StdResult<()> {
        println!("Saving state: {:#?}", self);
        
        let mut singleton = singleton(storage, STATE_KEY);
        let data = to_vec(self)?;
        singleton.save(&data)
    }
    
    pub fn load(storage: &dyn Storage) -> StdResult<Self> {
        let singleton = singleton_read(storage, STATE_KEY);
        let data: Vec<u8> = singleton.load()
            .map_err(|_err| StdError::generic_err("Failed to load state"))?;

        let loaded_state = from_slice(&data)
            .map_err(|_err| StdError::generic_err("Failed to deserialize state"))?;

        println!("Loaded state: {:#?}", loaded_state);
        
        Ok(loaded_state)
    }

    fn default() -> Self {
        State {
            ads: vec![],
            total_views: 0,
            plt_address: "".to_string(),
        }
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}


pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<Response> {
    // initialize the state
    let state = State {
        ads: vec![],
        total_views: 0,
        plt_address: "".to_string(),
    };

    // save the state
    state.save(deps.storage)?;

    Ok(Response::new())
}


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
   // storage: &mut dyn Storage,
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    let number = Uint128::from(20u128);
    let storage: &mut dyn Storage = deps.storage;
    match msg {
        HandleMsg::AddAd {
            id,
            image_url,
            target_url,
            reward_address,
        } => add_ad(deps.storage, env, info, id, image_url, target_url, reward_address, number),
        HandleMsg::ServeAd { id } => serve_ad(deps.storage, env, id),
        HandleMsg::DeleteAd { id } => delete_ad(deps.storage, id),
    }
}

pub fn update_ad(storage: &mut dyn Storage, ad: Ad) -> StdResult<()> {
    let mut ad_storage = singleton(storage, ad.id.as_bytes());
    println!("Saving state: {:#?}", ad);
    ad_storage.save(&ad)    
}

pub fn serve_ad(storage: &mut dyn Storage, env: Env, id: String) -> StdResult<Response> {
    let mut ad = get_ad(storage, id.clone())?;
    ad.views += 1;
    update_ad(storage, ad.clone())?;
    let mut attributes = vec![attr("action", "serve_ad"), attr("id", id)];
    attributes.push(attr("image_url", ad.image_url.clone()));
    attributes.push(attr("target_url", ad.target_url.clone()));
    let event = Event::new("serve_ad").add_attributes(attributes);
    println!("Event created: {:#?}", event);
    Ok(Response::new().add_event(event))
}

pub fn add_ad(
storage: &mut dyn Storage,
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
image_url: image_url.clone(),
target_url: target_url.clone(),
views: 0,
reward_address: reward_address.clone(),
};
//let storage: &mut dyn Storage = deps.storage;
let mut state = State::load(storage)?;
state.ads.push(ad.clone());
State::save(&state, storage)?;

let mut attributes = vec![attr("action", "add_ad"), attr("reward_address", reward_address)];
attributes.push(attr("id", id.clone()));
attributes.push(attr("image_url", image_url));
attributes.push(attr("target_url", target_url));
let event = Event::new("add_ad").add_attributes(attributes);
let reward_recipient = "0x123abc...";
let reward_msg = Cw20ExecuteMsg::Transfer {
    recipient: reward_recipient.into(),
    amount: plt_amount,
};
let reward_attrs = vec![    attr("action", "reward"),    attr("recipient", reward_recipient),    attr("amount", plt_amount),];
let reward_event = Event::new("reward").add_attributes(reward_attrs);
const PLT_ADDRESS: &'static str = "0x123abc..."; 
//let plt_address = deps.api.addr_validate(&PLT_ADDRESS)?;
let plt_address = "0x123abc..."; 
// Emit a staking event
let staking_attrs = vec![    attr("action", "staking"),    attr("contract_address", plt_address),    attr("staker_address", env.contract.address.clone()),    attr("amount", plt_amount),];
let staking_event = Event::new("staking").add_attributes(staking_attrs);
const PLT_REWARDS_CONTRACT: &'static str = "0x123abc...";
println!("Event created: {:#?}", event);
println!("Event created: {:#?}", reward_event);
println!("Event created: {:#?}", staking_event);
Ok(Response::new()
    .add_messages(vec![
       // CosmosMsg::Wasm(WasmMsg::Execute {
       //     contract_addr: env.contract.address.into(),
       //    msg: to_binary(&HandleMsg::ServeAd { storage: storage,id: id.clone() })?,
       //     funds: vec![],
      //  }),
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



fn delete_ad(storage: &mut dyn Storage, ad_id: String) -> StdResult<Response> {
   // let storage: &mut dyn Storage = deps.storage;
    let mut state = State::load(storage)?;
    let ad_idx = state
        .ads
        .iter()
        .position(|ad| ad.id == ad_id)
        .ok_or_else(|| StdError::not_found("Ad"))?;
    state.ads.remove(ad_idx);
    State::save(&state, storage)?;

    let mut attributes = vec![
        attr("action", "delete_ad"),
        attr("id", ad_id),
    ];
    let event = Event::new("delete_ad").add_attributes(attributes);
    println!("Event created: {:#?}", event);
    Ok(Response::new().add_event(event))
}


pub fn distribute_rewards(
deps: DepsMut,
env: Env,
_msg: EpochMsg,
) -> StdResult<Response> {
let mut state = State::load(deps.storage)?;
let storage: &mut dyn Storage = deps.storage;
 

let plt_address = env.contract.address.clone();
let reward_amount: u64 = 10; // specify the type and value as per your needs
 
// Iterate through all ads and distribute rewards
for mut ad in state.ads.iter_mut() {
    let ad_reward_amount = ad.views.checked_mul(reward_amount).unwrap_or_default();
    let ad_reward_amountv1 = ad_reward_amount.to_string();
    ad.views = 0;

    // Send rewards to the reward address specified by the ad creator
    let reward_address = deps.api.addr_validate(&ad.reward_address)?;
    let reward_msg = Cw20ExecuteMsg::Transfer {
        recipient: reward_address.into(),
        amount: ad_reward_amount.into(),
    };
    let reward_attrs = vec![        attr("action", "reward"),        attr("recipient", ad.reward_address.clone()),        attr("amount", ad_reward_amountv1),    ];
    let reward_event = Event::new("reward").add_attributes(reward_attrs);
    println!("Event created: {:#?}", reward_event);
}

State::save(&state, storage)?;
Ok(Response::new().add_attribute("action", "distribute_rewards"))
}


// Queries the ad with the given ID
pub fn query_ad(storage: &mut dyn Storage, id: String) -> StdResult<QueryAdResponse> {
let ad = get_ad(storage, id)?;
Ok(QueryAdResponse {
id: ad.id,
image_url: ad.image_url,
target_url: ad.target_url,
views: ad.views,
reward_address: ad.reward_address,
})
}


pub fn query_all_ads(storage: &mut dyn Storage) -> StdResult<QueryAllAdsResponse> {
    let ad_ids = get_all_ad_ids(storage)?;
    let mut ads: Vec<QueryAdResponse> = Vec::new();
    for id in ad_ids {
        let ad = query_ad(storage, id.clone())?;
        ads.push(ad);
    }
    Ok(QueryAllAdsResponse { ads })
}



// Helper functions go here

fn get_all_ad_ids(storage: &mut dyn Storage) -> StdResult<Vec<String>> {
let state = State::load(storage)?;
Ok(state.ads.iter().map(|ad| ad.id.clone()).collect())
}

pub fn get_ad(storage: &mut dyn Storage, id: String) -> Result<Ad,StdError>{
    let state = State::load(storage)?;
    for ad in state.ads.iter() {
        if ad.id == id {
            return Ok(ad.clone());
        }
    }
    Err(StdError::not_found("Ad"))
}


fn query_total_views(deps: Deps) -> StdResult<TotalViewsResponse> {
let state = State::load(deps.storage)?;
Ok(TotalViewsResponse {
total_views: state.total_views,
})
}


fn main() {
    let mut deps = mock_dependencies();

    let env = mock_env(); 

    let info = mock_info("sender_address", &coins(1000, "earth")); 

    let ad_id = String::from("test_id");
    let image_url = String::from("test_image_url");
    let target_url = String::from("test_target_url");
    let reward_address = String::from("test_reward_address");
    let mut storage: Box<dyn Storage> = Box::new(cosmwasm_std::testing::MockStorage::new());
    let storage_ref: &mut dyn Storage = &mut *storage;

    let msg = InitMsg {};

    match instantiate(deps.as_mut(), env.clone(), info.clone(), msg) {
        Ok(_response) => println!("State instantiated successfully."),
        Err(e) => println!("Failed to instantiate state: {}", e),
    }

    
    {
        match handle(
           // storage_ref,
            deps.as_mut(),
            env.clone(),
            info.clone(),
            HandleMsg::AddAd {
                id: ad_id.clone(),
                image_url: image_url.clone(),
                target_url: target_url.clone(),
                reward_address: reward_address.clone(),
            },
        ) {
            Ok(_response) => println!("Ad added successfully."),
            Err(e) => println!("Failed to add ad: {}", e),
        }
    }
    
    {
        match handle(
          //  storage_ref,
            deps.as_mut(),
            env.clone(),
            info.clone(),
            HandleMsg::ServeAd {
                id: ad_id.clone(),
            },
        ) {
            Ok(_response) => println!("Ad served successfully."),
            Err(e) => println!("Failed to serve ad: {}", e),
        }
    }
    
    {
        match handle(
          //  storage_ref,
            deps.as_mut(),
            env,
            info,
            HandleMsg::DeleteAd {
                id: ad_id,
            },
        ) {
            Ok(_response) => println!("Ad deleted successfully."),
            Err(e) => println!("Failed to delete ad: {}", e),
        }
    }
    

}