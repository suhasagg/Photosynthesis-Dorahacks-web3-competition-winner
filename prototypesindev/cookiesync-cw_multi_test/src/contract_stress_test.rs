use cosmwasm_std::{
    attr, to_binary, Addr, Binary, Deps, DepsMut, Env, Event, MessageInfo, Response, StdError,
    StdResult, Storage, Empty,
};
use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use log::{info, debug, warn, error};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Cookie {
    pub id: String,
    pub domain: String,
    pub data: String,
    pub expiration: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub cookies: Vec<Cookie>,
    pub total_cookies: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum ExecuteMsg {
    AddCookie {
        id: String,
        domain: String,
        data: String,
        expiration: u64,
    },
    SyncCookie {
        id: String,
    },
    DeleteCookie {
        id: String,
    },
    BatchSyncCookies {
        ids: Vec<String>,
    },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    Cookie {
        id: String,
    },
    Cookies,
    TotalCookies,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct TotalCookiesResponse {
    pub total_cookies: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryCookieResponse {
    pub id: String,
    pub domain: String,
    pub data: String,
    pub expiration: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryAllCookiesResponse {
    pub cookies: Vec<QueryCookieResponse>,
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
        cookies: vec![],
        total_cookies: 0,
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
        ExecuteMsg::AddCookie {
            id,
            domain,
            data,
            expiration,
        } => {
            info!("Executing AddCookie with id: {}", id);
            add_cookie(deps, env, info, id, domain, data, expiration)
        },
        ExecuteMsg::SyncCookie { id } => {
            info!("Executing SyncCookie with id: {}", id);
            sync_cookie(deps, env, id)
        },
        ExecuteMsg::DeleteCookie { id } => {
            info!("Executing DeleteCookie with id: {}", id);
            delete_cookie(deps, id)
        },
        ExecuteMsg::BatchSyncCookies { ids } => {
            info!("Executing BatchSyncCookies with ids: {:?}", ids);
            batch_sync_cookies(deps, env, ids)
        },
    }
}

fn add_cookie(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    id: String,
    domain: String,
    data: String,
    expiration: u64,
) -> StdResult<Response> {
    info!("Adding new cookie with id: {}", id);
    
    let mut state = STATE.load(deps.storage)?;
    let cookie = Cookie {
        id: id.clone(),
        domain: domain.clone(),
        data: data.clone(),
        expiration,
    };

    state.cookies.push(cookie);
    state.total_cookies += 1;
    STATE.save(deps.storage, &state)?;

    info!("Cookie {} added successfully.", id);

    let mut attributes = vec![attr("action", "add_cookie")];
    attributes.push(attr("id", id));
    attributes.push(attr("domain", domain));
    attributes.push(attr("data", data));
    attributes.push(attr("expiration", expiration.to_string()));

    let event = Event::new("add_cookie").add_attributes(attributes);
    Ok(Response::new().add_event(event))
}

fn sync_cookie(deps: DepsMut, _env: Env, id: String) -> StdResult<Response> {
    info!("Syncing cookie with id: {}", id);
    
    let mut state = STATE.load(deps.storage)?;

    // Find the cookie with the given id
    if let Some(cookie) = state.cookies.iter_mut().find(|cookie| cookie.id == id) {
        let domain = cookie.domain.clone();
        let data = cookie.data.clone();

        // Save the updated state
        STATE.save(deps.storage, &state)?;

        // Create the response with the required attributes
        let mut attributes = vec![attr("action", "sync_cookie"), attr("id", id.clone())];
        attributes.push(attr("domain", domain));
        attributes.push(attr("data", data));

        info!("Cookie {} synced successfully.", id);
        let event = Event::new("sync_cookie").add_attributes(attributes);
        Ok(Response::new().add_event(event))
    } else {
        warn!("Cookie with id {} not found for syncing.", id);
        Err(StdError::generic_err("Cookie not found"))
    }
}


fn delete_cookie(deps: DepsMut, id: String) -> StdResult<Response> {
    info!("Deleting cookie with id: {}", id);

    let mut state = STATE.load(deps.storage)?;
    let cookie_index = state.cookies.iter().position(|cookie| cookie.id == id);

    if let Some(index) = cookie_index {
        state.cookies.remove(index);
        state.total_cookies -= 1;
        info!("Cookie {} deleted successfully.", id);
    } else {
        warn!("Failed to delete cookie. Cookie with id {} not found.", id);
        return Err(StdError::generic_err("Cookie not found"));
    }

    STATE.save(deps.storage, &state)?;

    let event = Event::new("delete_cookie").add_attribute("action", "delete_cookie").add_attribute("id", id);
    Ok(Response::new().add_event(event))
}

fn batch_sync_cookies(deps: DepsMut, _env: Env, ids: Vec<String>) -> StdResult<Response> {
    info!("Batch syncing cookies with ids: {:?}", ids);
    
    let mut state = STATE.load(deps.storage)?;
    let mut events = Vec::new();

    for id in ids {
        let cookie = state.cookies.iter_mut().find(|cookie| cookie.id == id);

        if let Some(cookie) = cookie {
            info!("Cookie {} synced in batch.", id);
            let event = Event::new("sync_cookie")
                .add_attribute("action", "sync_cookie")
                .add_attribute("id", id.clone())
                .add_attribute("domain", cookie.domain.clone())
                .add_attribute("data", cookie.data.clone());
            events.push(event);
        } else {
            warn!("Cookie with id {} not found during batch sync.", id);
        }
    }

    STATE.save(deps.storage, &state)?;
    Ok(Response::new().add_events(events))
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Cookie { id } => {
            info!("Querying cookie with id: {}", id);
            query_cookie(deps, id)
        },
        QueryMsg::Cookies => {
            info!("Querying all cookies.");
            query_all_cookies(deps)
        },
        QueryMsg::TotalCookies => {
            info!("Querying total cookies.");
            query_total_cookies(deps)
        },
    }
}

fn query_cookie(deps: Deps, id: String) -> StdResult<Binary> {
    info!("Querying specific cookie with id: {}", id);
    
    let state = STATE.load(deps.storage)?;
    let cookie = state.cookies.iter().find(|&cookie| cookie.id == id);

    if let Some(cookie) = cookie {
        info!("Cookie {} found. Returning data.", id);
        to_binary(&QueryCookieResponse {
            id: cookie.id.clone(),
            domain: cookie.domain.clone(),
            data: cookie.data.clone(),
            expiration: cookie.expiration,
        })
    } else {
        warn!("Cookie with id {} not found.", id);
        Err(StdError::generic_err("Cookie not found"))
    }
}

fn query_all_cookies(deps: Deps) -> StdResult<Binary> {
    info!("Querying all cookies.");
    
    let state = STATE.load(deps.storage)?;
    let cookies: Vec<QueryCookieResponse> = state
        .cookies
        .iter()
        .map(|cookie| QueryCookieResponse {
            id: cookie.id.clone(),
            domain: cookie.domain.clone(),
            data: cookie.data.clone(),
            expiration: cookie.expiration,
        })
        .collect();

    info!("Returning all cookies data.");
    to_binary(&QueryAllCookiesResponse { cookies })
}

fn query_total_cookies(deps: Deps) -> StdResult<Binary> {
    info!("Querying total cookies.");
    
    let state = STATE.load(deps.storage)?;
    info!("Total cookies found: {}", state.total_cookies);
    to_binary(&TotalCookiesResponse {
        total_cookies: state.total_cookies,
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

    fn contract_cookie_sync() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_cookie_sync() {
        let _ = env_logger::builder().is_test(true).try_init(); // Initialize the logger for tests

        let mut app = mock_app();

        let code_id = app.store_code(contract_cookie_sync());

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("owner"),
                &InitMsg {},
                &[],
                "CookieSync",
                None,
            )
            .unwrap();

        // Add a cookie
        let cookie_id = "cookie1".to_string();
        let domain = "example.com".to_string();
        let data = "cookie_data".to_string();
        let expiration = 1609459200; //timestamp

        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::AddCookie {
                id: cookie_id.clone(),
                domain: domain.clone(),
                data: data.clone(),
                expiration,
            },
            &[],
        )
        .unwrap();

        // Sync the cookie
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::SyncCookie { id: cookie_id.clone() },
            &[],
        )
        .unwrap();

        // Query the cookie to check if it has been synced
        let cookie: QueryCookieResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr.clone(),
                &QueryMsg::Cookie { id: cookie_id.clone() },
            )
            .unwrap();

        assert_eq!(cookie.id, cookie_id);
        assert_eq!(cookie.domain, domain);

        // Delete the cookie
        app.execute_contract(
            Addr::unchecked("owner"),
            contract_addr.clone(),
            &ExecuteMsg::DeleteCookie { id: cookie_id.clone() },
            &[],
        )
        .unwrap();

        // Query all cookies to ensure the cookie has been deleted
        let cookies: QueryAllCookiesResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr,
                &QueryMsg::Cookies,
            )
            .unwrap();

        assert!(cookies.cookies.is_empty());
    }

    #[test]
    fn stress_test_add_and_sync_cookies() {
        let _ = env_logger::builder().is_test(true).try_init(); // Initialize the logger for tests

        let mut app = mock_app();
        let code_id = app.store_code(contract_cookie_sync());

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("owner"),
                &InitMsg {},
                &[],
                "CookieSync",
                None,
            )
            .unwrap();

        let start = Instant::now();
        let num_cookies = 1000;

        // Stress Test: Adding a large number of cookies
        for i in 0..num_cookies {
            let cookie_id = format!("cookie{}", i);
            let domain = format!("domain{}", i);
            let data = format!("data{}", i);
            let expiration = 1609459200 + i as u64; // timestamp variation

            app.execute_contract(
                Addr::unchecked("owner"),
                contract_addr.clone(),
                &ExecuteMsg::AddCookie {
                    id: cookie_id.clone(),
                    domain: domain.clone(),
                    data: data.clone(),
                    expiration,
                },
                &[],
            )
            .unwrap();
        }

        let duration = start.elapsed();
        println!("Time taken to add {} cookies: {:?}", num_cookies, duration);

        // Stress Test: Syncing all cookies
        let start = Instant::now();
        for i in 0..num_cookies {
            let cookie_id = format!("cookie{}", i);

            app.execute_contract(
                Addr::unchecked("owner"),
                contract_addr.clone(),
                &ExecuteMsg::SyncCookie { id: cookie_id.clone() },
                &[],
            )
            .unwrap();
        }

        let duration = start.elapsed();
        println!("Time taken to sync {} cookies: {:?}", num_cookies, duration);

        // Query one of the cookies to verify it has been synced
        let cookie: QueryCookieResponse = app
            .wrap()
            .query_wasm_smart(
                contract_addr.clone(),
                &QueryMsg::Cookie {
                    id: "cookie0".to_string(),
                },
            )
            .unwrap();

        assert_eq!(cookie.id, "cookie0");
        assert_eq!(cookie.domain, "domain0");
    }
}

