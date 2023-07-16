use cosmwasm_std::{
    to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct HandleMsg {
    pub user_id: String,
    pub cookie: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct QueryMsg {
    pub user_id: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct QueryResult {
    pub cookie: String,
}

pub fn init(deps: DepsMut, _env: Env, _info: MessageInfo, _msg: InitMsg) -> StdResult<Response> {
    Ok(Response::default())
}

pub fn handle(deps: DepsMut, _env: Env, _info: MessageInfo, msg: HandleMsg) -> StdResult<Response> {
    let mut storage = PrefixedStorage::new(deps.storage, b"cookies");
    storage.set(msg.user_id.as_bytes(), msg.cookie.as_bytes());
    Ok(Response::default())
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    let storage = ReadonlyPrefixedStorage::new(deps.storage, b"cookies");
    let cookie = storage
        .get(msg.user_id.as_bytes())
        .map(|bytes| String::from_utf8(bytes).unwrap())
        .unwrap_or_else(|| String::from("Cookie not found"));

    to_binary(&QueryResult { cookie })
}

#[no_mangle]
pub fn interface_version_8() -> u32 {
    // Implementation of the interface_version_5 function
    1
}

#[no_mangle]
pub extern "C" fn instantiate() {
    // Implementation of instantiate function
}
