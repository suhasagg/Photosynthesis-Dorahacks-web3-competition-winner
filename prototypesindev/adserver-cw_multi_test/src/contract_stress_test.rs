use cosmwasm_std::{
    attr, to_binary, Addr, Binary, CosmosMsg, Deps, DepsMut, Empty, Env, Event, MessageInfo,
    QueryRequest, Response, StdError, StdResult, Uint128, WasmMsg, Storage, from_binary,
};
use cw_storage_plus::{Item};
use cw20::{Cw20ExecuteMsg};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use log::{info, debug, warn, error}; 

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

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum ExecuteMsg {
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
pub struct TotalViewsResponse {
    pub total_views: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryAdResponse {
    pub id: String,
    pub image_url: String,
    pub target_url: String,
    pub views: u64,
    pub reward_address: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryAllAdsResponse {
    pub ads: Vec<QueryAdResponse>,
}

pub const STATE: Item<State> = Item::new("state");

pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<Response> {
    info!("Initializing contract state.");
    
    let state = State {
        ads: vec![],
        total_views: 0,
        plt_address: "".to_string(),
    };
    STATE.save(deps.storage, &state)?;
    
    info!("Contract initialized successfully.");
    Ok(Response::new())
}

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
            reward_address,
        } => {
            info!("Executing AddAd with id: {}", id);
            add_ad(deps, env, info, id, image_url, target_url, reward_address)
        },
        ExecuteMsg::ServeAd { id } => {
            info!("Executing ServeAd with id: {}", id);
            serve_ad(deps, env, id)
        },
        ExecuteMsg::DeleteAd { id } => {
            info!("Executing DeleteAd with id: {}", id);
            delete_ad(deps, id)
        },
        ExecuteMsg::BatchServeAds { ids } => {
            info!("Executing BatchServeAds with ids: {:?}", ids);
            batch_serve_ads(deps, env, ids)
        },
    }
}

fn add_ad(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    id: String,
    image_url: String,
    target_url: String,
    reward_address: String,
) -> StdResult<Response> {
    info!("Adding new ad with id: {}", id);
    
    let mut state = STATE.load(deps.storage)?;
    let ad = Ad {
        id: id.clone(),
        image_url: image_url.clone(),
        target_url: target_url.clone(),
        views: 0,
        reward_address: reward_address.clone(),
    };

    state.ads.push(ad);
    STATE.save(deps.storage, &state)?;

    info!("Ad {} added successfully.", id);

    let mut attributes = vec![attr("action", "add_ad"), attr("reward_address", reward_address)];
    attributes.push(attr("id", id));
    attributes.push(attr("image_url", image_url));
    attributes.push(attr("target_url", target_url));

    let event = Event::new("add_ad").add_attributes(attributes);
    Ok(Response::new().add_event(event))
}


fn serve_ad(deps: DepsMut, _env: Env, id: String) -> StdResult<Response> {
    let mut state = STATE.load(deps.storage)?;
    let ad = state.ads.iter_mut().find(|ad| ad.id == id);

    if let Some(ad) = ad {
        ad.views += 1;
        let image_url = ad.image_url.clone();
        let target_url = ad.target_url.clone();
        STATE.save(deps.storage, &state)?;

        let mut attributes = vec![attr("action", "serve_ad"), attr("id", id)];
        attributes.push(attr("image_url", image_url));
        attributes.push(attr("target_url", target_url));

        let event = Event::new("serve_ad").add_attributes(attributes);
        Ok(Response::new().add_event(event))
    } else {
        Err(StdError::generic_err("Ad not found"))
    }
}

fn delete_ad(deps: DepsMut, id: String) -> StdResult<Response> {
    info!("Deleting ad with id: {}", id);

    let mut state = STATE.load(deps.storage)?;
    let ad_index = state.ads.iter().position(|ad| ad.id == id);

    if let Some(index) = ad_index {
        state.ads.remove(index);
        info!("Ad {} deleted successfully.", id);
    } else {
        warn!("Failed to delete ad. Ad with id {} not found.", id);
        return Err(StdError::generic_err("Ad not found"));
    }

    STATE.save(deps.storage, &state)?;

    let event = Event::new("delete_ad").add_attribute("action", "delete_ad").add_attribute("id", id);
    Ok(Response::new().add_event(event))
}

fn batch_serve_ads(deps: DepsMut, _env: Env, ids: Vec<String>) -> StdResult<Response> {
    info!("Batch serving ads with ids: {:?}", ids);
    
    let mut state = STATE.load(deps.storage)?;
    let mut events = Vec::new();

    for id in ids {
        let ad = state.ads.iter_mut().find(|ad| ad.id == id);

        if let Some(ad) = ad {
            ad.views += 1;
            info!("Ad {} served in batch. Total views: {}", id, ad.views);

            let event = Event::new("serve_ad")
                .add_attribute("action", "serve_ad")
                .add_attribute("id", id.clone())
                .add_attribute("image_url", ad.image_url.clone())
                .add_attribute("target_url", ad.target_url.clone());
            events.push(event);
        } else {
            warn!("Ad with id {} not found during batch serve.", id);
        }
    }

    STATE.save(deps.storage, &state)?;
    Ok(Response::new().add_events(events))
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Ad { id } => {
            info!("Querying ad with id: {}", id);
            query_ad(deps, id)
        },
        QueryMsg::Ads => {
            info!("Querying all ads.");
            query_all_ads(deps)
        },
        QueryMsg::TotalViews => {
            info!("Querying total views.");
            query_total_views(deps)
        },
    }
}

fn query_ad(deps: Deps, id: String) -> StdResult<Binary> {
    info!("Querying specific ad with id: {}", id);
    
    let state = STATE.load(deps.storage)?;
    let ad = state.ads.iter().find(|&ad| ad.id == id);

    if let Some(ad) = ad {
        info!("Ad {} found. Returning data.", id);
        to_binary(&QueryAdResponse {
            id: ad.id.clone(),
            image_url: ad.image_url.clone(),
            target_url: ad.target_url.clone(),
            views: ad.views,
            reward_address: ad.reward_address.clone(),
        })
    } else {
        warn!("Ad with id {} not found.", id);
        Err(StdError::generic_err("Ad not found"))
    }
}

fn query_all_ads(deps: Deps) -> StdResult<Binary> {
    info!("Querying all ads.");
    
    let state = STATE.load(deps.storage)?;
    let ads: Vec<QueryAdResponse> = state
        .ads
        .iter()
        .map(|ad| QueryAdResponse {
            id: ad.id.clone(),
            image_url: ad.image_url.clone(),
            target_url: ad.target_url.clone(),
            views: ad.views,
            reward_address: ad.reward_address.clone(),
        })
        .collect();

    info!("Returning all ads data.");
    to_binary(&QueryAllAdsResponse { ads })
}

fn query_total_views(deps: Deps) -> StdResult<Binary> {
    info!("Querying total views.");
    
    let state = STATE.load(deps.storage)?;
    info!("Total views found: {}", state.total_views);
    to_binary(&TotalViewsResponse {
        total_views: state.total_views,
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};
    use std::time::Instant;

    fn mock_app() -> App {
        App::default()
    }

    fn contract_adserver() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_adserver() {
        let _ = env_logger::builder().is_test(true).try_init(); // Initialize the logger for tests

        let mut app = mock_app();

        let code_id = app.store_code(contract_adserver());

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("owner"),
                &InitMsg {},
                &[],
                "AdServer",
                None,
            )
            .unwrap();

        // Add an ad
        let ad_id = "ad1".to_string();
        let image_url = "https://usatoday.com/image1".to_string();
        let target_url = "https://serotoninsmile.com".to_string();
        let reward_address = "reward1".to_string();

        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::AddAd {
                id: ad_id.clone(),
                image_url: image_url.clone(),
                target_url: target_url.clone(),
                reward_address: reward_address.clone(),
            },
            &[],
        )
        .unwrap();

        // Serve the ad
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::ServeAd { id: ad_id.clone() },
            &[],
        )
        .unwrap();

        // Query the ad to check if it has been served
        let ad: QueryAdResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr.clone(),
                &QueryMsg::Ad { id: ad_id.clone() },
            )
            .unwrap();

        assert_eq!(ad.views, 1);

        // Delete the ad
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::DeleteAd { id: ad_id.clone() },
            &[],
        )
        .unwrap();

        // Query all ads to ensure the ad has been deleted
        let ads: QueryAllAdsResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr,
                &QueryMsg::Ads,
            )
            .unwrap();

        assert!(ads.ads.is_empty());
    }

    #[test]
    fn stress_test_add_and_serve_ads() {
        let _ = env_logger::builder().is_test(true).try_init(); // Initialize the logger for tests

        let mut app = mock_app();
        let code_id = app.store_code(contract_adserver());

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("owner"),
                &InitMsg {},
                &[],
                "AdServer",
                None,
            )
            .unwrap();

        let start = Instant::now();
        let num_ads = 1000;

        // Stress Test: Adding a large number of ads
        for i in 0..num_ads {
            let ad_id = format!("ad{}", i);
            let image_url = format!("https://usatoday.com/image{}", i);
            let target_url = "https://holographicnftswithapurpose.com".to_string();
            let reward_address = format!("reward{}", i);

            app.execute_contract(
                Addr::unchecked("owner"),
                contract_addr.clone(),
                &ExecuteMsg::AddAd {
                    id: ad_id,
                    image_url,
                    target_url,
                    reward_address,
                },
                &[],
            )
            .unwrap();
        }

        let duration = start.elapsed();
        println!("Time taken to add {} ads: {:?}", num_ads, duration);

        // Stress Test: Serving all ads
        let start = Instant::now();
        for i in 0..num_ads {
            let ad_id = format!("ad{}", i);

            app.execute_contract(
                Addr::unchecked("owner"),
                contract_addr.clone(),
                &ExecuteMsg::ServeAd { id: ad_id },
                &[],
            )
            .unwrap();
        }

        let duration = start.elapsed();
        println!("Time taken to serve {} ads: {:?}", num_ads, duration);

        // Query one of the ads to verify it has been served
        let ad: QueryAdResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr.clone(),
                &QueryMsg::Ad {
                    id: "ad0".to_string(),
                },
            )
            .unwrap();

        assert_eq!(ad.views, 1);
    }
}

