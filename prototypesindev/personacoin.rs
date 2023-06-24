use cosmwasm_std::{
    to_binary, Api, Binary, CanonicalAddr, Env,
    Querier, QueryRequest, StdError, StdResult, Storage, Uint128,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage};
use serde::{Deserialize, Serialize};
use std::convert::TryInto;
use cosmwasm_storage::singleton;
use cosmwasm_std::to_vec;
use cosmwasm_storage::singleton_read;
use cosmwasm_std::from_slice;
use cosmwasm_std::Event;
use cosmwasm_std::DepsMut;
use cosmwasm_std::MessageInfo;
use cosmwasm_std::Response;
use cosmwasm_std::attr;
use cosmwasm_std::Deps;
use cosmwasm_std::MemoryStorage;
use cosmwasm_std::testing::MockQuerier;
use cosmwasm_std::ContractInfo;
use cosmwasm_std::WasmQuery::ContractInfo as OtherContractInfo;
use cosmwasm_std::testing::MockApi;
use cosmwasm_std::BlockInfo;
use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
use cosmwasm_storage::ReadonlySingleton;
use cosmwasm_storage::Singleton;
use cosmwasm_std::coins;
use cosmwasm_std::Order;
use cosmwasm_std::testing::MockStorage;
const CONFIG_KEY: &[u8] = b"config";
const STATE_KEY: &[u8] = b"state";
const COOKIE_KEY: &[u8] = b"cookie";
const PORT_ID: &str = "cookiesync";
const TIMEOUT_HEIGHT: u64 = 100;
const TIMEOUT_TIMESTAMP: u64 = 0;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Config {
    pub parameter: String, // replace with your actual configuration parameters
    // add other configuration parameters here
}

impl Config {
    
    pub fn load(storage: &mut dyn Storage) -> StdResult<Params> {
        let singleton = ReadonlySingleton::new(storage, CONFIG_KEY);
        singleton.load()
    }
        
    pub fn save<S: Storage>(&self, storage: &mut S) -> StdResult<()> {
        let mut singleton = Singleton::new(storage, CONFIG_KEY);
        singleton.save(self)
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct InitMsg {}


pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<Response> {
    // initialize the state
    let state = CookiePacketData {
        cookie:  Cookie {
            id: String::from("default_id"),
            domain: String::from("default_domain"),
            data: String::from("default_data"),
            expiration: 0,
        },
        source_pub: "new york times".to_string(),
        dest_pub: "cnn".to_string(),
    };

    // save the state
    state.save(deps.storage)?;

    Ok(Response::new())
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Cookie {
    pub id: String,
    pub domain: String,
    pub data: String,
    pub expiration: u64,
}

impl Cookie {
    pub fn is_expired(&self, env: &Env) -> bool {
        self.expiration <= env.block.time.seconds()
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Params {
    field1: String, // Replace 'field1' and 'field2' with your actual field names
    field2: i32,    // and adjust their types as needed.   
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct CookiePacketData {
    pub cookie: Cookie,
    pub source_pub: String,
    pub dest_pub: String,
}

impl CookiePacketData {
    pub fn new(cookie: Cookie, source_pub: String, dest_pub: String) -> Self {
        Self {
            cookie,
            source_pub,
            dest_pub,
        }
    }

    pub fn validate_basic(&self) -> StdResult<()> {
        if self.cookie.id.is_empty() || self.source_pub.is_empty() || self.dest_pub.is_empty() {
            return Err(StdError::generic_err("invalid packet data"));
        }
        Ok(())
    }
    
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
}

pub fn init(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
) -> Result<Response, StdError> {
    let cookie_packet_data = CookiePacketData::new(
        Cookie {
            id: String::from("default_id"),
            domain: String::from("default_domain"),
            data: String::from("default_data"),
            expiration: 0,
        },
        String::from("new york times"), 
        String::from("cnn")
    );
    cookie_packet_data.save(deps.storage)?;
    Ok(Response::default())
} 

pub enum HandleMsg {
    Sync {
        cookie: Cookie,
        source_pub: String,
        dest_pub: String,
    },
    CreateCookie, 
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    GetCookie { id: String },
    ListCookies {},
    GetParams {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum QueryAnswer {
    NoCookie { id: String },
    Cookie { cookie: Cookie },
    Cookies { cookies: Vec<Cookie> },
    Params { params: Params },
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
         HandleMsg::Sync{cookie, source_pub, dest_pub} => sync(deps, env, source_pub, dest_pub, cookie),
         HandleMsg::CreateCookie => {
            Ok(Response::new()) // Do nothing for this case, just return an Ok response
        },
     }
 }
 


pub fn sync(
    deps: DepsMut,
    env: Env,
    source_pub: String, 
    dest_pub: String,
    cookie: Cookie,
) -> StdResult<Response> {
    let result = format!("{}:{}:{}", &dest_pub, &source_pub, &cookie.id);
    deps.storage.set(result.as_bytes(), &to_binary(&cookie)?);
    let mut attributes = vec![attr("action", "sync"), attr("from", &source_pub), attr("to", &dest_pub), attr("cookie", &cookie.id)];
    let event = Event::new("sync").add_attributes(attributes);
    println!("Event created: {:#?}", event);
    Ok(Response::new().add_event(event)) 

}


pub fn handle_msg(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    match msg {
        // handle the HandleMsg::CreateCookie message
        HandleMsg::CreateCookie {} => {
            let cookie = Cookie {
                id: "123".to_string(),
                domain: "example.com".to_string(),
                data: "cookie_data".to_string(),
                expiration: 0,  // You have this as a string, but it should be a u64
            };
            deps.storage.set(cookie.id.as_bytes(), &to_binary(&cookie)?);
            let mut attributes = vec![attr("action", "create_cookie")];
            let event = Event::new("create_cookie").add_attributes(attributes);
            println!("Event created: {:#?}", event);
            Ok(Response::new().add_event(event))
        
        }
        HandleMsg::Sync { cookie, source_pub, dest_pub } => sync(deps, env, dest_pub, source_pub, cookie),
    }
}

pub fn query_msg(
    deps: Deps,
    env: Env,
    msg: QueryMsg,
) -> StdResult<Binary> {
    match msg {
        // handle the QueryMsg::GetCookie message
        QueryMsg::GetCookie { id } => {
            let cookie = match deps.storage.get(id.as_bytes()) {
                None => return to_binary(&QueryAnswer::NoCookie { id }),
                Some(cookie_data) => from_slice(&cookie_data)?,
            };
            to_binary(&QueryAnswer::Cookie { cookie })
        }  
        // handle the QueryMsg::ListCookies message
        QueryMsg::ListCookies {} => {
            let cookies: StdResult<Vec<Cookie>> = deps
            .storage
            .range(None, None, Order::Ascending)
            .map(|item| {
                let (key, value) = item;
                let value_vec: Vec<u8> = value.to_vec();  // convert the slice to a Vec
                Ok(from_slice::<Cookie>(&value_vec)?)
            })
            .collect();
            to_binary(&QueryAnswer::Cookies { cookies: cookies? })
        }
        // handle the QueryMsg::GetParams message
        QueryMsg::GetParams {} => {
            let mut storage: Box<dyn Storage> = Box::new(MockStorage::new());
            let params = Config::load(&mut *storage)?;
            to_binary(&QueryAnswer::Params { params })
        }
    }
}

fn main() {
    // initialize the dependencies
    let mut deps = mock_dependencies();

    let env = mock_env(); 

    let info = mock_info("sender_address", &coins(1000, "earth")); 

    let cookie = Cookie {
        id: "123".to_string(),
        domain: "test".to_string(),
        data: "cookie_data".to_string(),
        expiration: 0,
    };

    let cookie1 = Cookie {
        id: "456".to_string(),
        domain: "test1".to_string(),
        data: "cookie_data1".to_string(),
        expiration: 0,
    };
    let source_pub = "new york times".to_string();
    let dest_pub = "cnn".to_string();
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
            HandleMsg::Sync{
                cookie: cookie.clone(),
                source_pub: source_pub.clone(),
                dest_pub: dest_pub.clone(),
            },
        ) {
            Ok(_response) => println!("Cookie Synced successfully."),
            Err(e) => println!("Failed to add Sync cookies: {}", e),
        }
    }


    {
        match handle(
           // storage_ref,
            deps.as_mut(),
            env.clone(),
            info.clone(),
            HandleMsg::Sync{
                cookie: cookie1.clone(),
                source_pub: source_pub.clone(),
                dest_pub: dest_pub.clone(),
            },
        ) {
            Ok(_response) => println!("Cookie Synced successfully."),
            Err(e) => println!("Failed to add Sync cookies: {}", e),
        }
    }

}


