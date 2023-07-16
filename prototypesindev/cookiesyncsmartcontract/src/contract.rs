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

// Define a Config struct for the configuration parameters of the contract
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Config {
    pub parameter: String, // The actual configuration parameters. Replace with your specific parameters.
    // add other configuration parameters here
}
// Implement methods for the Config struct
impl Config {
     // This method is used to load the configuration from the contract's storage
    pub fn load(storage: &mut dyn Storage) -> StdResult<Params> {
      // Create a readonly singleton that gives access to the data stored in the contract's storage at the key CONFIG_KEY
        let singleton = ReadonlySingleton::new(storage, CONFIG_KEY);
       // Load and return the configuration from the storage
        singleton.load()
    }
         // This method is used to save the configuration to the contract's storage
    pub fn save<S: Storage>(&self, storage: &mut S) -> StdResult<()> {
         // Create a singleton that gives access to the data stored in the contract's storage at the key CONFIG_KEY
        let mut singleton = Singleton::new(storage, CONFIG_KEY);
         // Save the configuration to the storage
        singleton.save(self)
    }
}

// Define a InitMsg struct that represents the initialization message of the contract
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct InitMsg {}

// This function is called once when the contract is instantiated, and it cannot be called again.
// It takes mutable dependencies to the contract's storage and environment,
// information about the message, and the message itself.

pub fn instantiate(
    deps: DepsMut, // Provides access to the features of the Cosmos SDK that the contract depends on.
    _env: Env,     // Provides information about the current state of the blockchain to the contract.
    _info: MessageInfo,  // Information about the incoming message such as sender, funds sent, etc.
    _msg: InitMsg, // The actual message that has been sent to the contract during the instantiation.
) -> StdResult<Response> {  // The result of the function. It can either be an error or a successful response.
    // initialize the state
    // Initialize the state of the contract.
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
 // Save the state to the contract's storage.
    // save the state
    state.save(deps.storage)?;
// Return an OK response. As no events or attributes are added to this response,
    // it will simply signify that the instantiation was successful.
    Ok(Response::new())
}
// This is the main Cookie struct, representing a web cookie in your application.
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Cookie {
    pub id: String,   // A unique identifier for the cookie
    pub domain: String,  // The domain that the cookie is associated with
    pub data: String,  // The actual data contained within the cookie
    pub expiration: u64, // The expiration time of the cookie as a Unix timestamp
}
// Implement methods for the Cookie struct

impl Cookie {
    // This method checks whether the cookie has expired
    pub fn is_expired(&self, env: &Env) -> bool {
          // Compare the expiration time of the cookie with the current blockchain time.
        // If the cookie's expiration time is less than or equal to the current time, it is expired.
        self.expiration <= env.block.time.seconds()
    }
}
// The Params struct holds parameters that can be used to configure your contract.
// Adjust field names and types as needed.
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Params {
    field1: String, // This is a placeholder for an actual field. Replace with a specific parameter.
    field2: i32,    // and adjust their types as needed.   
}
// The CookiePacketData struct represents a "packet" of data that contains a Cookie and its source and destination.
#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct CookiePacketData {
    pub cookie: Cookie, // The CookiePacketData struct represents a "packet" of data that contains a Cookie and its source and destination.
    pub source_pub: String,  // The public identifier of the source of the transfer
    pub dest_pub: String,   // The public identifier of the destination of the transfer
}
// This is an implementation block for the CookiePacketData struct.
impl CookiePacketData {
    // This function is a constructor that takes a cookie, a source public key and a destination public key 
    // and returns a new instance of the struct.
    pub fn new(cookie: Cookie, source_pub: String, dest_pub: String) -> Self {
        Self {
            cookie,
            source_pub,
            dest_pub,
        }
    }

 // This function checks if the cookie id, source_pub, and dest_pub are not empty. 
    // If any of them is empty, it returns an error; otherwise, it returns Ok.
    pub fn validate_basic(&self) -> StdResult<()> {
        if self.cookie.id.is_empty() || self.source_pub.is_empty() || self.dest_pub.is_empty() {
            return Err(StdError::generic_err("invalid packet data"));
        }
        Ok(())
    }
    
    // This function attempts to save the current instance of the struct to the provided mutable storage.
    // It first serializes the instance into bytes, then attempts to save these bytes into the singleton.
    // If it's successful, it returns Ok, otherwise it returns the error it encountered.
    pub fn save(&self, storage: &mut dyn Storage) -> StdResult<()> {
        println!("Saving state: {:#?}", self);
        
        let mut singleton = singleton(storage, STATE_KEY);
        let data = to_vec(self)?;
        singleton.save(&data)
    }
    // This static method attempts to load an instance of the struct from the provided storage.
    // It first attempts to load the data from the singleton. If successful, it deserializes the data into an instance of the struct.
    // If any of these operations fails, it returns the error it encountered.
    // Otherwise, it returns the loaded instance of the struct.
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
// This is the init function for the smart contract.
// It is invoked once when the contract is instantiated
pub fn init(
    deps: DepsMut, // This is a mutable reference to the dependencies of the contract. It contains things like the contract's storage and API.
    env: Env, // This provides information about the blockchain environment.
    info: MessageInfo, // This provides information about the message that has been sent to the contract.
) -> Result<Response, StdError> { // This function will return either a Response indicating success or an StdError indicating failure.
     // Create an instance of CookiePacketData with initial values.
    let cookie_packet_data = CookiePacketData::new(
        Cookie {
            id: String::from("default_id"), // Initial value for the id of the Cookie.
            domain: String::from("default_domain"), // Initial value for the domain of the Cookie.
            data: String::from("default_data"), // Initial value for the data of the Cookie.
            expiration: 0, // Initial value for the expiration of the Cookie.
        },
        String::from("new york times"),  // Initial value for the source of the CookiePacketData.
        String::from("cnn") // Initial value for the destination of the CookiePacketData.
    );
     // Save the created CookiePacketData to the contract's storage.
    cookie_packet_data.save(deps.storage)?;
    // Return a Response with no data, events or messages, indicating that the init function has executed successfully.
    Ok(Response::default())
} 
// The HandleMsg enum represents the different types of messages that the contract can handle.
// Each variant of the enum corresponds to a different action that can be taken on the contract.
pub enum HandleMsg {
    BatchSyncSave {
        source_pub: String,
        dest_pub: String,
        cookies: Vec<Cookie>,
    },
   
    // The Sync variant is used when a Cookie needs to be synced from a source to a destination.
    // It includes the Cookie to be synced, and the public identifiers of the source and destination.
    Sync {
        cookie: Cookie,  // The Cookie to be synced
        source_pub: String, // The public identifier of the source.
        dest_pub: String,  // The public identifier of the destination.
    },
      // The CreateCookie variant is used when a new Cookie needs to be created.
    CreateCookie, 
}
// The QueryMsg enum represents the different types of read-only queries that the contract can handle.
// Each variant of the enum corresponds to a different type of query
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]  // This ensures that all enum variants will be serialized as snake_case.
pub enum QueryMsg {
   // The GetCookie variant is used when the client wants to get a specific Cookie by its ID.
    GetCookie { id: String }, // The ID of the Cookie to retrieve.
    // The ListCookies variant is used when the client wants to get a list of all Cookies.
    ListCookies {},
    // The GetParams variant is used when the client wants to retrieve the parameters of the contract.
    GetParams {},
}
// The QueryAnswer enum represents the possible responses that the contract can give to a query.
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum QueryAnswer {
    // The NoCookie variant is returned when the client has asked for a cookie with a specific ID, but no such cookie exists.
    // It includes the ID that was queried.
    NoCookie { id: String }, // The ID of the cookie that was queried but not found.
     // The Cookie variant is returned when the client has asked for a cookie with a specific ID and the cookie was found.
    // It includes the found Cookie.
    Cookie { cookie: Cookie }, // The Cookie that was found.
   // The Cookies variant is returned when the client has asked for a list of all cookies.
    // It includes a vector of all the Cookies.
    Cookies { cookies: Vec<Cookie> }, // A vector containing all the cookies.
    // The Params variant is returned when the client has asked for the parameters of the contract.
    // It includes the parameters of the contract.
    Params { params: Params }, // The parameters of the contract.
}

// This is the handle function for the smart contract.
// It is invoked whenever a transaction is sent to the contract.
pub fn handle(
    // storage: &mut dyn Storage,
     deps: DepsMut, // This is a mutable reference to the dependencies of the contract. It contains things like the contract's storage and API.
     env: Env, // This provides information about the blockchain environment.
     info: MessageInfo, // This provides information about the message that has been sent to the contract.
     msg: HandleMsg, // This is the message that has been sent to the contract.
 ) -> StdResult<Response> {
     let number = Uint128::from(20u128);
      // Retrieve the contract's storage from the dependencies.
     let storage: &mut dyn Storage = deps.storage;
       // Match on the HandleMsg to determine what action to take.
     match msg {
         HandleMsg::BatchSyncSave { source_pub, dest_pub, cookies } => 
                batch_sync_save(deps, env, source_pub, dest_pub, cookies),           
         // If the message is a Sync message, call the sync function with the appropriate parameters.
         HandleMsg::Sync{cookie, source_pub, dest_pub} => sync(deps, env, source_pub, dest_pub, cookie),
        // If the message is a CreateCookie message, respond with an empty response.
        // Note: It seems like this should be doing more than just returning an empty response. You might want to add code here to create a cookie.
         HandleMsg::CreateCookie => {
            Ok(Response::new()) 
        },
    }
}


// This function is responsible for syncing CookiePacketData and updates the state of the contract.
pub fn sync(
    deps: DepsMut, // Provides access to the features of the Cosmos SDK that the contract depends on.
    env: Env,       // Provides information about the current state of the blockchain to the contract.
    source_pub: String, // The public key of the source that wants to sync.
    dest_pub: String,  // The public key of the destination where the source wants to sync.
    cookie: Cookie,     // The public key of the destination where the source wants to sync.
) -> StdResult<Response> {  // The result of the function. It can either be an error or a successful response.
      // Formats the destination public key, source public key and cookie id into a string.
    let result = format!("{}:{}:{}", &dest_pub, &source_pub, &cookie.id);
    // Sets the formatted string as a key and the serialized cookie as a value in the storage.
    deps.storage.set(result.as_bytes(), &to_binary(&cookie)?);
      // Creates an attribute vector with the action and the public keys of the source and destination, and the cookie id.
    let mut attributes = vec![attr("action", "sync"), attr("from", &source_pub), attr("to", &dest_pub), attr("cookie", &cookie.id)];
    // Creates an event that shows the synchronization action has occurred.
    let event = Event::new("sync").add_attributes(attributes);
    println!("Event created: {:#?}", event);
    // Returns a successful response with the event.
    Ok(Response::new().add_event(event)) 

}
//Cookie batch size can be tuned to optimise gas usage and archway rewards
pub fn batch_sync_save(
    deps: DepsMut,
    env: Env,
    source_pub: String,
    dest_pub: String,
    cookies: Vec<Cookie>,
) -> StdResult<Response> {
let mut events = Vec::new();
// Define a vector to store serialized cookies key-value pairs
let mut kv_pairs: Vec<(Vec<u8>, Binary)> = Vec::new();
let mut cookie_ids: Vec<String> = Vec::new();

for cookie in cookies {
    // Format the key and serialize the cookie
    let result = format!("{}:{}:{}", &dest_pub, &source_pub, &cookie.id);
    let serialized_cookie = match to_binary(&cookie) {
        Ok(cookie) => cookie,
        Err(e) => return Err(e), // If cookie serialization fails, return the error
    };

    // Add key-value pair to the kv_pairs vector
    kv_pairs.push((result.as_bytes().to_vec(), serialized_cookie));
    cookie_ids.push(cookie.id.clone());

    // Create an event for each synchronization
    let mut attributes = vec![attr("action", "sync"), attr("from", &source_pub), attr("to", &dest_pub), attr("cookie", &cookie.id)];
    let event = Event::new("batch-sync").add_attributes(attributes);
    println!("Event created: {:#?}", event); // print the event
    events.push(event);
    
}

// Define a key for batch storage
let batch_storage_key = format!("{}:{}:{}", &dest_pub, &source_pub, cookie_ids.last().unwrap_or(&String::new()));
let serialized_cookie_ids = to_binary(&cookie_ids)?;

// Save the serialized cookie_ids vector in storage under the batch_storage_key
deps.storage.set(batch_storage_key.as_bytes(), &serialized_cookie_ids);

// Return a successful response with all events
Ok(Response::new().add_events(events))

}

// This function is the main entry point for message handling of the contract.
// It takes mutable dependencies to the contract's storage and environment,
// information about the message, and the message itself.

pub fn handle_msg( 
    deps: DepsMut,  // Provides access to the features of the Cosmos SDK that the contract depends on.
    env: Env,       // Provides information about the current state of the blockchain to the contract.
    info: MessageInfo, // Information about the incoming message such as sender, funds sent, etc.
    msg: HandleMsg,   // The actual message that has been sent to the contract.
) -> StdResult<Response> { // The result of the function. It can either be an error or a successful response.
  // Match on the specific type of message received.
    match msg {
        // handle the HandleMsg::CreateCookie message
        // When the message is of type CreateCookie,
        // create a new cookie, save it to the contract's storage, and return an event
        HandleMsg::CreateCookie {} => {
           // Define the cookie to be stored
            let cookie = Cookie {
                id: "123".to_string(),
                domain: "example.com".to_string(),
                data: "cookie_data".to_string(),
                expiration: 0,  // You have this as a string, but it should be a u64
            };
            // Save the cookie into the contract's storage using the cookie's id as the key
            deps.storage.set(cookie.id.as_bytes(), &to_binary(&cookie)?);
             // Create an attribute for the event
            let mut attributes = vec![attr("action", "create_cookie")];
              // Create an event to signal the cookie creation
            let event = Event::new("create_cookie").add_attributes(attributes);
            println!("Event created: {:#?}", event);
             // Return the result with the event included in the response
            Ok(Response::new().add_event(event))
        
        },
         // When the message is of type Sync, call the sync function
        HandleMsg::Sync { cookie, source_pub, dest_pub } => sync(deps, env, dest_pub, source_pub, cookie),
        HandleMsg::BatchSyncSave { source_pub, dest_pub, cookies } => 
                batch_sync_save(deps, env, source_pub, dest_pub, cookies),   
    }
}
// This function is the main entry point for query handling of the contract.
// It takes dependencies to the contract's storage and environment, and the query message itself.
pub fn query_msg(
    deps: Deps,    // Provides access to the features of the Cosmos SDK that the contract depends on.
    env: Env,       // Provides information about the current state of the blockchain to the contract.
    msg: QueryMsg,  // The actual query that has been sent to the contract.
) -> StdResult<Binary> { // The result of the function. It can either be an error or a successful response.
     // Match on the specific type of query message received.
    match msg {
        // handle the QueryMsg::GetCookie message
         // When the query message is of type GetCookie, return the specified cookie from the storage,
        // or an error message if it does not exist.
        QueryMsg::GetCookie { id } => {
            let cookie = match deps.storage.get(id.as_bytes()) {
                None => return to_binary(&QueryAnswer::NoCookie { id }),// If there is no cookie, return an error message
                Some(cookie_data) => from_slice(&cookie_data)?,// If there is a cookie, deserialize it
            };
              // Serialize the response and return it
            to_binary(&QueryAnswer::Cookie { cookie })
        }  
        // handle the QueryMsg::ListCookies message
         // When the query message is of type ListCookies, return a list of all cookies from the storage.
        QueryMsg::ListCookies {} => {
              // Iterate over each item in storage, deserializing each item into a Cookie
            let cookies: StdResult<Vec<Cookie>> = deps
            .storage
            .range(None, None, Order::Ascending)
            .map(|item| {
                let (key, value) = item;
                let value_vec: Vec<u8> = value.to_vec();  // convert the slice to a Vec
                Ok(from_slice::<Cookie>(&value_vec)?) // Deserialize into a Cookie
            })
            .collect();
         // Serialize the response and return it
            to_binary(&QueryAnswer::Cookies { cookies: cookies? })
        }
        // handle the QueryMsg::GetParams message
        // When the query message is of type GetParams, return the contract's parameters from the storage.
        QueryMsg::GetParams {} => {
              // Create a mock storage to read the configuration
            let mut storage: Box<dyn Storage> = Box::new(MockStorage::new());
             // Load the configuration from the storage
            let params = Config::load(&mut *storage)?;
             // Serialize the response and return it
            to_binary(&QueryAnswer::Params { params })
        }
    }
}

fn main() {
    // initialize the dependencies
   // Create a mock environment for testing. 
// This includes a storage system and an API that can be used for interactions that would usually be handled by the Cosmos SDK.
    let mut deps = mock_dependencies();
// Creates a mock environment with default parameters.
    let env = mock_env(); 
// Create mock transaction information. In this case, it is as if a transaction was sent by "sender_address" with 1000 "arch" coins.
    let info = mock_info("sender_address", &coins(1000, "arch")); 
// Initialize a couple of cookie objects to be used in the tests.
    let cookie = Cookie {
        id: "123".to_string(),
        domain: "test".to_string(),
        data: "cookie_data".to_string(),
        expiration: 0, // Cookie does not expire
    };

    let cookie1 = Cookie {
        id: "456".to_string(),
        domain: "test1".to_string(),
        data: "cookie_data1".to_string(),
        expiration: 0, // Cookie does not expire
    };

    let cookie2 = Cookie {
        id: "789".to_string(),
        domain: "test2".to_string(),
        data: "cookie_data2".to_string(),
        expiration: 0,
    };
    
    let cookie3 = Cookie {
        id: "012".to_string(),
        domain: "test3".to_string(),
        data: "cookie_data3".to_string(),
        expiration: 0,
    };
    
    let cookie4 = Cookie {
        id: "345".to_string(),
        domain: "test4".to_string(),
        data: "cookie_data4".to_string(),
        expiration: 0,
    };

    let cookies = vec![cookie.clone(), cookie1.clone(), cookie2.clone(), cookie3.clone(), cookie4.clone()];
    // Define source and destination for the cookies.
    let source_pub = "new york times".to_string();
    let dest_pub = "cnn".to_string();
    // Initialize the InitMsg struct for contract initialization.
    let msg = InitMsg {};
// Try to instantiate the contract state with the initial parameters.
// The success or failure of the operation will be printed to the console.
    match instantiate(deps.as_mut(), env.clone(), info.clone(), msg) {
        Ok(_response) => println!("State instantiated successfully."),
        Err(e) => println!("Failed to instantiate state: {}", e),
    }

   {
         match handle(
             deps.as_mut(),
             env.clone(),
             info.clone(),
             HandleMsg::BatchSyncSave {
                  source_pub: source_pub.clone(),
                  dest_pub: dest_pub.clone(),
                  cookies,
            },
        ) {
           Ok(_response) => println!("Cookies batch synced successfully."),
           Err(e) => println!("Failed to batch sync cookies: {}", e),
          }
    }

    // Try to handle the Sync message with the created cookies.
// The success or failure of the operation will be printed to the console.
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

