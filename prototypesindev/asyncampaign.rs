use cosmwasm_std::{
    entry_point, to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult, Storage,
    BankMsg, Coin, Addr,
};
use cosmwasm_storage::{singleton, singleton_read, Singleton};
use serde::{Deserialize, Serialize};

static KEY_CONFIG: &[u8] = b"config";

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Config {
    pub users: Vec<User>,   
    pub campaigns: Vec<Campaign>,
    pub messages: Vec<Message>,   
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct User {
    pub id: Addr,
    pub campaigns_created: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Campaign {
    title: String,
    content: String,
    creator: Addr,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Message {
    sender: Addr,
    content: String,
}

#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> StdResult<Response> {
    let config = Config {
        users: vec![],
        campaigns: vec![],
        messages: vec![],
    };
    singleton(&mut deps.storage, KEY_CONFIG).save(&config)?;
    Ok(Response::default())
}

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    match msg {
        ExecuteMsg::RegisterUser {} => register_user(deps, env, info),
        ExecuteMsg::CreateCampaign { title, content } => create_campaign(deps, env, info, title, content),
        ExecuteMsg::SendMessage { content } => send_message(deps, env, info, content),
    }
}

fn register_user(deps: DepsMut, _env: Env, info: MessageInfo) -> StdResult<Response> {
    let mut config: Config = singleton(&mut deps.storage, KEY_CONFIG).load()?;
    if config.users.iter().any(|u| u.id == info.sender) {
        return Err(cosmwasm_std::StdError::generic_err("User already registered"));
    }
    let user = User {
        id: info.sender.clone(),
        campaigns_created: 0,
    };
    config.users.push(user);
    singleton(&mut deps.storage, KEY_CONFIG).save(&config)?;
    Ok(Response::new().add_attribute("method", "register_user"))
}

fn create_campaign(deps: DepsMut, _env: Env, info: MessageInfo, title: String, content: String) -> StdResult<Response> {
    let mut config: Config = singleton(&mut deps.storage, KEY_CONFIG).load()?;
    if !config.users.iter().any(|u| u.id == info.sender) {
        return Err(cosmwasm_std::StdError::generic_err("User must be registered"));
    }
    let campaign = Campaign {
        title,
        content,
        creator: info.sender.clone

    };
    config.campaigns.push(campaign);
    singleton(&mut deps.storage, KEY_CONFIG).save(&config)?;
    Ok(Response::new().add_attribute("method", "create_campaign"))
}

fn send_message(deps: DepsMut, _env: Env, info: MessageInfo, content: String) -> StdResult<Response> {
    let mut config: Config = singleton(&mut deps.storage, KEY_CONFIG).load()?;
    let message = Message {
        sender: info.sender.clone(),
        content,
    };
    config.messages.push(message);
    singleton(&mut deps.storage, KEY_CONFIG).save(&config)?;
    Ok(Response::new().add_attribute("method", "send_message"))
}
