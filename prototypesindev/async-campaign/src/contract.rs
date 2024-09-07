use cosmwasm_std::{
    entry_point, to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult, Addr,
};
use cosmwasm_storage::{singleton, Singleton};
use serde::{Deserialize, Serialize};
use cosmwasm_std::Empty;
use cosmwasm_storage::{singleton_read};

static KEY_CONFIG: &[u8] = b"config";

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Config {
    pub owner: Addr,
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
    pub title: String,
    pub content: String,
    pub creator: Addr,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Message {
    pub sender: Addr,
    pub content: String,
}

// Define InitMsg for contract instantiation
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct InitMsg {
    pub owner: String, // Address of the owner
}

// Instantiate function to initialize contract state
#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,  // Use InitMsg as an argument
) -> StdResult<Response> {
    println!("Instantiating contract with owner: {}", msg.owner);
    let owner_addr = deps.api.addr_validate(&msg.owner)?; // Validate the owner address

    let config = Config {
        owner: owner_addr.clone(),
        users: vec![],
        campaigns: vec![],
        messages: vec![],
    };
    
    println!("Saving initial contract config with owner: {}", owner_addr);
    singleton(deps.storage, KEY_CONFIG).save(&config)?;
    
    println!("Contract instantiated successfully.");
    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("owner", owner_addr.to_string()))
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub enum ExecuteMsg {
    RegisterUser {},
    CreateCampaign { title: String, content: String },
    SendMessage { content: String },
}

// Defining the QueryMsg enum
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    GetConfig {},
    GetUser { id: Addr },
    GetCampaign { title: String },
}

// Execute function for handling contract messages
#[entry_point]
pub fn execute(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    println!("Executing contract message: {:?}", msg);
    match msg {
        ExecuteMsg::RegisterUser {} => {
            println!("Registering new user: {}", info.sender);
            register_user(deps, info)
        },
        ExecuteMsg::CreateCampaign { title, content } => {
            println!("Creating campaign with title: {} by user: {}", title, info.sender);
            create_campaign(deps, info, title, content)
        },
        ExecuteMsg::SendMessage { content } => {
            println!("Sending message: {} by user: {}", content, info.sender);
            send_message(deps, info, content)
        },
    }
}

// Register a new user
fn register_user(deps: DepsMut, info: MessageInfo) -> StdResult<Response> {
    let mut config: Config = singleton(deps.storage, KEY_CONFIG).load()?;
    if config.users.iter().any(|u| u.id == info.sender) {
        println!("User already registered: {}", info.sender);
        return Err(cosmwasm_std::StdError::generic_err("User already registered"));
    }
    let user = User {
        id: info.sender.clone(),
        campaigns_created: 0,
    };
    config.users.push(user);
    singleton(deps.storage, KEY_CONFIG).save(&config)?;
    
    println!("User successfully registered: {}", info.sender);
    Ok(Response::new().add_attribute("method", "register_user"))
}

// Create a campaign
fn create_campaign(deps: DepsMut, info: MessageInfo, title: String, content: String) -> StdResult<Response> {
    let mut config: Config = singleton(deps.storage, KEY_CONFIG).load()?;
    if !config.users.iter().any(|u| u.id == info.sender) {
        println!("User must be registered to create a campaign: {}", info.sender);
        return Err(cosmwasm_std::StdError::generic_err("User must be registered"));
    }
    let campaign = Campaign {
        title: title.clone(),
        content: content.clone(),
        creator: info.sender.clone(),
    };
    config.campaigns.push(campaign);
    singleton(deps.storage, KEY_CONFIG).save(&config)?;
    
    println!("Campaign successfully created with title: {} by user: {}", title, info.sender);
    Ok(Response::new()
        .add_attribute("method", "create_campaign")
        .add_attribute("title", title))
}

// Send a message
fn send_message(deps: DepsMut, info: MessageInfo, content: String) -> StdResult<Response> {
    let mut config: Config = singleton(deps.storage, KEY_CONFIG).load()?;
    let message = Message {
        sender: info.sender.clone(),
        content: content.clone(),
    };
    config.messages.push(message);
    singleton(deps.storage, KEY_CONFIG).save(&config)?;
    
    println!("Message successfully sent by user: {}", info.sender);
    Ok(Response::new()
        .add_attribute("method", "send_message")
        .add_attribute("content", content))
}

#[entry_point]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    println!("Querying contract with message: {:?}", msg);
    match msg {
        QueryMsg::GetConfig {} => {
            println!("Querying full config");
            to_binary(&singleton_read::<Config>(deps.storage, KEY_CONFIG).load()?)
        },
        QueryMsg::GetUser { id } => {
            println!("Querying user with id: {}", id);
            query_user(deps, id)
        },
        QueryMsg::GetCampaign { title } => {
            println!("Querying campaign with title: {}", title);
            query_campaign(deps, title)
        },
    }
}

fn query_user(deps: Deps, id: Addr) -> StdResult<Binary> {
    let config: Config = singleton_read(deps.storage, KEY_CONFIG).load()?;
    if let Some(user) = config.users.iter().find(|u| u.id == id) {
        println!("User found with id: {}", id);
        to_binary(user)
    } else {
        println!("User not found with id: {}", id);
        Err(cosmwasm_std::StdError::generic_err("User not found"))
    }
}

fn query_campaign(deps: Deps, title: String) -> StdResult<Binary> {
    let config: Config = singleton_read(deps.storage, KEY_CONFIG).load()?;
    if let Some(campaign) = config.campaigns.iter().find(|c| c.title == title) {
        println!("Campaign found with title: {}", title);
        to_binary(campaign)
    } else {
        println!("Campaign not found with title: {}", title);
        Err(cosmwasm_std::StdError::generic_err("Campaign not found"))
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};
    use cosmwasm_std::Addr;

    fn mock_app() -> App {
        App::default()
    }

    fn contract_campaign() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

    #[test]
    fn test_user_registration_and_campaign_creation() {
        let mut app = mock_app();

        let code_id = app.store_code(contract_campaign());

        let init_msg = InitMsg {
            owner: "creator".to_string(),  // InitMsg for initialization
        };

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("creator"),
                &init_msg,  
                &[],
                "CampaignContract",
                None,
            )
            .unwrap();

        // Register a user
        app.execute_contract(
            Addr::unchecked("user1"),
            contract_addr.clone(),
            &ExecuteMsg::RegisterUser {},
            &[],
        )
        .unwrap();

        // Query the state to check if the user is registered
        let config: Config = app
            .wrap()
            .query_wasm_smart(contract_addr.clone(), &QueryMsg::GetConfig {})
            .unwrap();

        assert_eq!(config.users.len(), 1);
        assert_eq!(config.users[0].id, Addr::unchecked("user1"));

        // Create a campaign
        app.execute_contract(
            Addr::unchecked("user1"),
            contract_addr.clone(),
            &ExecuteMsg::CreateCampaign {
                title: "Campaign Title".to_string(),
                content: "Campaign Content".to_string(),
            },
            &[],
        )
        .unwrap();

        // Query the state to check if the campaign is created
        let config: Config = app
            .wrap()
            .query_wasm_smart(contract_addr.clone(), &QueryMsg::GetConfig {})
            .unwrap();

        assert_eq!(config.campaigns.len(), 1);
        assert_eq!(config.campaigns[0].title, "Campaign Title");
        assert_eq!(config.campaigns[0].content, "Campaign Content");
        assert_eq!(config.campaigns[0].creator, Addr::unchecked("user1"));
    }

    #[test]
    fn test_message_sending() {
        let mut app = mock_app();

        let code_id = app.store_code(contract_campaign());

        let init_msg = InitMsg {
            owner: "creator".to_string(),  
        };

        let contract_addr = app
            .instantiate_contract(
                code_id,
                Addr::unchecked("creator"),
                &init_msg,  
                &[],
                "CampaignContract",
                None,
            )
            .unwrap();

        // Register a user
        app.execute_contract(
            Addr::unchecked("user1"),
            contract_addr.clone(),
            &ExecuteMsg::RegisterUser {},
            &[],
        )
        .unwrap();

        // Send a message
        app.execute_contract(
            Addr::unchecked("user1"),
            contract_addr.clone(),
            &ExecuteMsg::SendMessage {
                content: "Hello, World!".to_string(),
            },
            &[],
        )
        .unwrap();

        // Query the state to check if the message was sent
        let config: Config = app
            .wrap()
            .query_wasm_smart(contract_addr, &QueryMsg::GetConfig {})
            .unwrap();

        assert_eq!(config.messages.len(), 1);
        assert_eq!(config.messages[0].content, "Hello, World!");
        assert_eq!(config.messages[0].sender, Addr::unchecked("user1"));
    }
}

