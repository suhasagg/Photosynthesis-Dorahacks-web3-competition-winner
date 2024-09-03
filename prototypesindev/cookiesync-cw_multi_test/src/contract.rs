use cosmwasm_std::{
    attr, to_binary, Addr, Binary, Deps, DepsMut, Env, Event, MessageInfo, Response, StdError,
    StdResult, Storage, Uint128, Order, Empty,
};
use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};


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
    let state = State {
        cookies: vec![],
        total_cookies: 0,
    };
    STATE.save(deps.storage, &state)?;
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
        } => add_cookie(deps, env, info, id, domain, data, expiration),
        ExecuteMsg::SyncCookie { id } => sync_cookie(deps, env, id),
        ExecuteMsg::DeleteCookie { id } => delete_cookie(deps, id),
        ExecuteMsg::BatchSyncCookies { ids } => batch_sync_cookies(deps, env, ids),
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
    let mut state = STATE.load(deps.storage)?;
    let cookie = Cookie {
        id: id.clone(),
        domain: domain.clone(),
        data: data.clone(),
        expiration,
    };

    state.cookies.push(cookie);
    STATE.save(deps.storage, &state)?;

    let mut attributes = vec![attr("action", "add_cookie")];
    attributes.push(attr("id", id));
    attributes.push(attr("domain", domain));
    attributes.push(attr("data", data));
    attributes.push(attr("expiration", expiration.to_string()));

    let event = Event::new("add_cookie").add_attributes(attributes);
    Ok(Response::new().add_event(event))
}

fn sync_cookie(deps: DepsMut, _env: Env, id: String) -> StdResult<Response> {
    let mut state = STATE.load(deps.storage)?;

    // Find the cookie with the given id
    if let Some(cookie) = state.cookies.iter_mut().find(|cookie| cookie.id == id) {
        let domain = cookie.domain.clone();
        let data = cookie.data.clone();

        // Save the updated state
        STATE.save(deps.storage, &state)?;

        // Create the response with the required attributes
        let mut attributes = vec![attr("action", "sync_cookie"), attr("id", id)];
        attributes.push(attr("domain", domain));
        attributes.push(attr("data", data));

        let event = Event::new("sync_cookie").add_attributes(attributes);
        Ok(Response::new().add_event(event))
    } else {
        Err(StdError::generic_err("Cookie not found"))
    }
}


fn delete_cookie(deps: DepsMut, id: String) -> StdResult<Response> {
    let mut state = STATE.load(deps.storage)?;
    let cookie_index = state.cookies.iter().position(|cookie| cookie.id == id);

    if let Some(index) = cookie_index {
        state.cookies.remove(index);
    } else {
        return Err(StdError::generic_err("Cookie not found"));
    }

    STATE.save(deps.storage, &state)?;

    let event = Event::new("delete_cookie").add_attribute("action", "delete_cookie").add_attribute("id", id);
    Ok(Response::new().add_event(event))
}

fn batch_sync_cookies(deps: DepsMut, _env: Env, ids: Vec<String>) -> StdResult<Response> {
    let mut state = STATE.load(deps.storage)?;
    let mut events = Vec::new();

    for id in ids {
        let cookie = state.cookies.iter_mut().find(|cookie| cookie.id == id);

        if let Some(cookie) = cookie {
            let event = Event::new("sync_cookie")
                .add_attribute("action", "sync_cookie")
                .add_attribute("id", id.clone())
                .add_attribute("domain", cookie.domain.clone())
                .add_attribute("data", cookie.data.clone());
            events.push(event);
        }
    }

    STATE.save(deps.storage, &state)?;
    Ok(Response::new().add_events(events))
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Cookie { id } => query_cookie(deps, id),
        QueryMsg::Cookies => query_all_cookies(deps),
        QueryMsg::TotalCookies => query_total_cookies(deps),
    }
}

fn query_cookie(deps: Deps, id: String) -> StdResult<Binary> {
    let state = STATE.load(deps.storage)?;
    let cookie = state.cookies.iter().find(|&cookie| cookie.id == id);

    if let Some(cookie) = cookie {
        to_binary(&QueryCookieResponse {
            id: cookie.id.clone(),
            domain: cookie.domain.clone(),
            data: cookie.data.clone(),
            expiration: cookie.expiration,
        })
    } else {
        Err(StdError::generic_err("Cookie not found"))
    }
}

fn query_all_cookies(deps: Deps) -> StdResult<Binary> {
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

    to_binary(&QueryAllCookiesResponse { cookies })
}

fn query_total_cookies(deps: Deps) -> StdResult<Binary> {
    let state = STATE.load(deps.storage)?;
    to_binary(&TotalCookiesResponse {
        total_cookies: state.total_cookies,
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};

    fn mock_app() -> App {
        App::default()
    }

    fn contract_cookie_sync() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_cookie_sync() {
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
}
