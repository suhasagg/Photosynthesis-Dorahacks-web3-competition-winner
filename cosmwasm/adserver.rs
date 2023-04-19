
use cosmwasm_std::{
    to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage};
use serde::{Deserialize, Serialize};
use schemars::JsonSchema;
use std::convert::TryInto;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct HandleMsg {
    pub ad_slot: String,
    pub ad_url: String,
    pub bid: u128,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryMsg {
    pub ad_slot: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryResult {
    pub ad_url: String,
}

pub fn init(deps: DepsMut, _env: Env, _info: MessageInfo, _msg: InitMsg) -> StdResult<Response> {
    Ok(Response::default())
}

pub fn handle(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    let mut storage = PrefixedStorage::new(deps.storage, b"ads");


    let existing_bid: Option<u128> = storage
        .get(msg.ad_slot.as_bytes())
        .map(|bytes| {
            let bid = u128::from_be_bytes(bytes.try_into().unwrap());
            bid
        });

    if existing_bid.is_none() || msg.bid > existing_bid.unwrap() {
        storage.set(msg.ad_slot.as_bytes(), &msg.bid.to_be_bytes());
    }

    Ok(Response::default())
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    let storage = ReadonlyPrefixedStorage::new(deps.storage, b"ads");

    let ad_url = storage
        .get(msg.ad_slot.as_bytes())
        .map(|bytes| String::from_utf8(bytes).unwrap())
        .unwrap_or_else(|| String::from("No ad found"));

    to_binary(&QueryResult { ad_url })
}

#[no_mangle]
pub fn interface_version_8() -> u32 {
    // Implementation of the interface_version_5 function
    1
}
