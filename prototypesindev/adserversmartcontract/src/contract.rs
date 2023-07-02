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

// Constants used throughout the module
const STATE_KEY: &[u8] = b"state";
const AD_PREFIX: &[u8] = b"ad";
const TOTAL_VIEWS_KEY: &[u8] = b"total_views";

pub const CW20_STAKING_CONTRACT: &str = "staking_contract";
pub const CW20_STAKING_BALANCE: &str = "staking_balance";
// Structure for Ad
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Ad {
    pub id: String, 
    pub image_url: String,
    pub target_url: String,
    pub views: u64,
    pub reward_address: String,
}
// State for Adserver
// Serve ad method for the state
pub struct AdserverState {
    ads: Vec<Ad>,
}
// Serve ad method for the state
impl AdserverState {
    fn serve_ad(&mut self, ad_id: &String) -> Option<&mut Ad> {
        self.ads.iter_mut().find(|ad| ad.id == *ad_id)
    }
 // Delete ad method for the state
    fn delete_ad(&mut self, ad_id: &String) -> bool {
        if let Some(index) = self.ads.iter().position(|ad| ad.id == *ad_id) {
            self.ads.remove(index);
            true
        } else {
            false
        }
    }
}


// Structure for Total Views response
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct TotalViewsResponse {
    pub total_views: u64,
}



// Structure for State of the contract
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub ads: Vec<Ad>,
    pub total_views: u64,
    pub plt_address: String,
}

// Implement methods for state structure
impl State {
    pub fn save(&self, storage: &mut dyn Storage) -> StdResult<()> {
        println!("Saving state: {:#?}", self);
        
        let mut singleton = singleton(storage, STATE_KEY);
        let data = to_vec(self)?;
        singleton.save(&data)
    }
     // load method for State structure
    pub fn load(storage: &dyn Storage) -> StdResult<Self> {
        let singleton = singleton_read(storage, STATE_KEY);
        let data: Vec<u8> = singleton.load()
            .map_err(|_err| StdError::generic_err("Failed to load state"))?;

        let loaded_state = from_slice(&data)
            .map_err(|_err| StdError::generic_err("Failed to deserialize state"))?;

        println!("Loaded state: {:#?}", loaded_state);
        
        Ok(loaded_state)
    }
// Provide a default State when there is no state available in the storage.
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
    BatchServeAds {
        ids: Vec<String>,
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
        HandleMsg::BatchServeAds { ids } => {
            batch_serve_ads(storage, env, ids)
        },
    }
}

pub fn update_ad(storage: &mut dyn Storage, ad: Ad) -> StdResult<()> {
    let mut ad_storage = singleton(storage, ad.id.as_bytes());
    println!("Saving state: {:#?}", ad);
    ad_storage.save(&ad)    
}

// Function to serve an ad, incrementing its view count
  // Mutable reference to the Storage interface, providing contract storage manipulation
  // Current blockchain environment
  // Unique identifier for the ad
pub fn serve_ad(storage: &mut dyn Storage, env: Env, id: String) -> StdResult<Response> {
   // Retrieve the ad with the specified id from the storage
    let mut ad = get_ad(storage, id.clone())?;
    // Increment the view count of the ad
    ad.views += 1;
    // Update the ad in the storage with the new view count
    update_ad(storage, ad.clone())?;
    // Prepare a list of attributes for the event of serving an ad
    let mut attributes = vec![attr("action", "serve_ad"), attr("id", id)];
    // Add additional attributes, the image_url and target_url of the ad
    attributes.push(attr("image_url", ad.image_url.clone()));
   
    attributes.push(attr("target_url", ad.target_url.clone()));
   // Create a new event "serve_ad" with the prepared attributes
    let event = Event::new("serve_ad").add_attributes(attributes);
    // Print the created event to the console
    println!("Event created: {:#?}", event);
    // Return a response with the created event
    Ok(Response::new().add_event(event))
}
//Optimised gas usage - tune batch size accordingly for optimium gas and rewards
//Performance Improvement: Batch serving allows for the processing of multiple ads in a single transaction, reducing the overhead and latency associated with sending multiple individual transactions. This can lead to significant performance improvements, especially when serving a large number of ads.
pub fn batch_serve_ads(storage: &mut dyn Storage, env: Env, ids: Vec<String>) -> StdResult<Response> {
    let mut events = Vec::new();

    for id in ids {
        // Retrieve the ad with the specified id from the storage
        let mut ad = get_ad(storage, id.clone())?;
        // Increment the view count of the ad
        ad.views += 1;
        // Update the ad in the storage with the new view count
        update_ad(storage, ad.clone())?;
        // Prepare a list of attributes for the event of serving an ad
        let mut attributes = vec![attr("action", "serve_ad"), attr("id", id)];
        // Add additional attributes, the image_url and target_url of the ad
        attributes.push(attr("image_url", ad.image_url.clone()));
        attributes.push(attr("target_url", ad.target_url.clone()));
        // Create a new event "serve_ad" with the prepared attributes
        let event = Event::new("serve_ad").add_attributes(attributes);
        // Print the created event to the console
        println!("Event created: {:#?}", event);
        // Add the event to the events vector
        events.push(event);
    }

    // Return a response with all the created events
    Ok(Response::new().add_events(events))
}

// This function is responsible for adding advertisements in our ad system. It accepts
// a number of parameters including the ad information, the environment, the message information, 
// and some data related to reward system. 

pub fn add_ad(
 // Mutable reference to our Storage interface, allowing us to interact with the contract's storage    
storage: &mut dyn Storage,
 // The environment in which the contract is executed
env: Env,
// Information about the incoming message such as the sender
info: MessageInfo,
// Unique identifier for the ad
id: String,
 // URL to the image used in the ad
image_url: String,
  // URL the ad should direct to when clicked
target_url: String,
  // Address to which rewards should be sent
reward_address: String,
  // Amount of rewards to be sent
plt_amount: Uint128,
) -> StdResult<Response> {
// Create an instance of Ad with the provided details   
let ad = Ad {
id: id.clone(),
image_url: image_url.clone(),
target_url: target_url.clone(),
views: 0,
reward_address: reward_address.clone(),
};
 // Load the current state of our storage
let mut state = State::load(storage)?;
// Add the newly created ad to the list of ads
state.ads.push(ad.clone());
  // Save the updated state back to the storage
State::save(&state, storage)?;
 // Create attributes for the event when an ad is added
let mut attributes = vec![attr("action", "add_ad"), attr("reward_address", reward_address)];
// Add more attributes related to the ad
attributes.push(attr("id", id.clone()));
attributes.push(attr("image_url", image_url));
attributes.push(attr("target_url", target_url));
// Create an event that an ad has been added
let event = Event::new("add_ad").add_attributes(attributes);
 // Prepare for the reward system. Set the address that will receive the reward.
let reward_recipient = "0x123abc...";
 // Set the message that will be used to transfer the reward
let reward_msg = Cw20ExecuteMsg::Transfer {
    recipient: reward_recipient.into(),
    amount: plt_amount,
};
// Create attributes for the reward event
let reward_attrs = vec![    attr("action", "reward"),    attr("recipient", reward_recipient),    attr("amount", plt_amount),];
 // Create a reward event with the prepared attributes
let reward_event = Event::new("reward").add_attributes(reward_attrs);
 // Define the contract address for staking
const PLT_ADDRESS: &'static str = "0x123abc..."; 
//let plt_address = deps.api.addr_validate(&PLT_ADDRESS)?;
let plt_address = "0x123abc..."; 
 // Create attributes for the staking event
let staking_attrs = vec![    attr("action", "staking"),    attr("contract_address", plt_address),    attr("staker_address", env.contract.address.clone()),    attr("amount", plt_amount),];
// Create a staking event with the prepared attributes
let staking_event = Event::new("staking").add_attributes(staking_attrs);
// Define the address for the rewards contract
const PLT_REWARDS_CONTRACT: &'static str = "0x123abc...";
// Print the created events for debug purposes
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

// Function to delete an advertisement from the system
 // Mutable reference to the Storage interface, providing contract storage manipulation,// Unique identifier for the ad
fn delete_ad(storage: &mut dyn Storage, ad_id: String) -> StdResult<Response> {
   // Load the current state from the storage
    let mut state = State::load(storage)?;
    // Find the index of the ad with the provided id in the list of ads
    let ad_idx = state
        .ads
        .iter()
        .position(|ad| ad.id == ad_id)
         // If the ad is not found, return an error
        .ok_or_else(|| StdError::not_found("Ad"))?;
    // Remove the ad from the list of ads in the state
    state.ads.remove(ad_idx);
    // Save the updated state back to the storage
    State::save(&state, storage)?;
// Prepare a list of attributes for the event of deleting an ad
    let mut attributes = vec![
        attr("action", "delete_ad"),
        attr("id", ad_id),
    ];
    // Prepare a list of attributes for the event of deleting an ad
    let event = Event::new("delete_ad").add_attributes(attributes);
     // Print the created event to the console
    println!("Event created: {:#?}", event);
    // Return a response with the created event
    Ok(Response::new().add_event(event))
}

// This function distributes rewards to the creators of the ads based on the views their ads have gotten.
pub fn distribute_rewards(
deps: DepsMut, // Mutable dependencies include the API, storage and the event manager
env: Env, // Provides information about the contract environment
_msg: EpochMsg, // Epoch message
) -> StdResult<Response> {
 // Load the current state of the storage
let mut state = State::load(deps.storage)?;
let storage: &mut dyn Storage = deps.storage;
 

let plt_address = env.contract.address.clone();  // clone the address of the contract
let reward_amount: u64 = 10; // specify the type and value as per your needs
 // Reward amount per view
// Iterate through all ads and distribute rewards
// Iterate through all ads and distribute rewards based on views
for mut ad in state.ads.iter_mut() {
    let ad_reward_amount = ad.views.checked_mul(reward_amount).unwrap_or_default(); // Calculate the reward for the ad
    let ad_reward_amountv1 = ad_reward_amount.to_string();
    ad.views = 0; // Reset views

    // Send rewards to the reward address specified by the ad creator
     // Validate the reward address and send the rewards
    let reward_address = deps.api.addr_validate(&ad.reward_address)?;
    let reward_msg = Cw20ExecuteMsg::Transfer {
        recipient: reward_address.into(), // Set the recipient of the reward
        amount: ad_reward_amount.into(), // Set the amount of the reward
    };
    let reward_attrs = vec![        attr("action", "reward"),        attr("recipient", ad.reward_address.clone()),        attr("amount", ad_reward_amountv1),    ];
    let reward_event = Event::new("reward").add_attributes(reward_attrs);
    println!("Event created: {:#?}", reward_event);
}
// Save the state
State::save(&state, storage)?;
 // Return a successful Response
Ok(Response::new().add_attribute("action", "distribute_rewards"))
}


// Queries the ad with the given ID
pub fn query_ad(storage: &mut dyn Storage, id: String) -> StdResult<QueryAdResponse> {
// Get the ad with the given ID
let ad = get_ad(storage, id)?;
Ok(QueryAdResponse {
id: ad.id,
image_url: ad.image_url,
target_url: ad.target_url,
views: ad.views,
reward_address: ad.reward_address,
})
}

// This function returns all the ads
pub fn query_all_ads(storage: &mut dyn Storage) -> StdResult<QueryAllAdsResponse> {
     // Get all ad ids
    let ad_ids = get_all_ad_ids(storage)?;
   
    let mut ads: Vec<QueryAdResponse> = Vec::new();
    // Iterate over all ad ids and push their data into the ads vector
    for id in ad_ids {
        let ad = query_ad(storage, id.clone())?;
        ads.push(ad);
    }
    Ok(QueryAllAdsResponse { ads })
}



// Helper functions go here
// This helper function gets all ad ids
fn get_all_ad_ids(storage: &mut dyn Storage) -> StdResult<Vec<String>> {
let state = State::load(storage)?;
Ok(state.ads.iter().map(|ad| ad.id.clone()).collect())
}

// This function gets an ad by its id
pub fn get_ad(storage: &mut dyn Storage, id: String) -> Result<Ad,StdError>{
    let state = State::load(storage)?;
    for ad in state.ads.iter() {
        if ad.id == id {
            return Ok(ad.clone());
        }
    }
    Err(StdError::not_found("Ad"))
}

// This function queries the total number of views
fn query_total_views(deps: Deps) -> StdResult<TotalViewsResponse> {
let state = State::load(deps.storage)?;
Ok(TotalViewsResponse {
total_views: state.total_views,
})
}

fn main() {
    // Create mock dependencies
    let mut deps = mock_dependencies();

    // Create a mock environment
    let env = mock_env(); 

    // Create mock info about the sender
    let info = mock_info("sender_address", &coins(1000, "earth")); 

    // Define multiple ads details for testing
    let ad_ids = vec![
        String::from("test_id1"),
        String::from("test_id2"),
        String::from("test_id3"),
    ];
    let image_urls = vec![
        String::from("test_image_url1"),
        String::from("test_image_url2"),
        String::from("test_image_url3"),
    ];
    let target_urls = vec![
        String::from("test_target_url1"),
        String::from("test_target_url2"),
        String::from("test_target_url3"),
    ];
    let reward_addresses = vec![
        String::from("test_reward_address1"),
        String::from("test_reward_address2"),
        String::from("test_reward_address3"),
    ];

    // Initialize a new contract
    let msg = InitMsg {};

    // Try to instantiate the contract
    match instantiate(deps.as_mut(), env.clone(), info.clone(), msg) {
        Ok(_response) => println!("State instantiated successfully."),
        Err(e) => println!("Failed to instantiate state: {}", e),
    }

    let single_ad_id = ad_ids.first().unwrap();  // Get the first ad ID
    match handle(
        deps.as_mut(),
        env.clone(),
        info.clone(),
        HandleMsg::ServeAd {
            id: single_ad_id.clone(),
        },
    ) {
        Ok(_response) => println!("Ad {} served successfully.", single_ad_id),
        Err(e) => println!("Failed to serve ad {}: {}", single_ad_id, e),
    }



    // Add multiple ads
    for (((ad_id, image_url), target_url), reward_address) in ad_ids.iter().zip(&image_urls).zip(&target_urls).zip(&reward_addresses) {
        match handle(
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
            Ok(_response) => println!("Ad {} added successfully.", ad_id),
            Err(e) => println!("Failed to add ad {}: {}", ad_id, e),
        }
    }

    // Serve multiple ads
    match handle(
        deps.as_mut(),
        env.clone(),
        info.clone(),
        HandleMsg::BatchServeAds {
            ids: ad_ids.clone(),
        },
    ) {
        Ok(_response) => println!("Ads served successfully."),
        Err(e) => println!("Failed to serve ads: {}", e),
    }

    // Delete the ads
    for ad_id in ad_ids {
        match handle(
            deps.as_mut(),
            env.clone(),
            info.clone(),
            HandleMsg::DeleteAd {
                id: ad_id.clone(),
            },
        ) {
            Ok(_response) => println!("Ad {} deleted successfully.", ad_id),
            Err(e) => println!("Failed to delete ad {}: {}", ad_id, e),
        }
    }
}
