use cosmwasm_std::{
    attr, to_binary, Api, Binary, CanonicalAddr, CosmosMsg, Deps, DepsMut, Env, MessageInfo,
    Response, StdError, StdResult, Storage, Uint128, WasmMsg,
};
use cosmwasm_storage::{
    prefixed, prefixed_read, singleton, singleton_read, PrefixedStorage, ReadonlyPrefixedStorage,
    ReadonlySingleton, Singleton,
};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::convert::TryInto;

use cosmwasm_std::entry_point;

pub fn option_bytes_to_u128(value: Option<Vec<u8>>) -> u128 {
    match value {
        Some(bytes) => {
            // Convert the Vec<u8> to a u128 using from_be_bytes
            let u128_value = u128::from_be_bytes(bytes.as_slice().try_into().unwrap());
            u128_value
        },
        None => {
        0
    }
       
    }
}


// Add the missing imports and types

pub fn init(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<Response> {
    // Implement your logic here
    unimplemented!()
}

pub fn handle(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    match msg {
        HandleMsg::AddStake { params } => add_stake(deps, env, info, params),
        HandleMsg::RedeemLiquidTokens { amount } => redeem_liquid_tokens(deps, env, info, Uint128::from(amount).u128()),
    }
}


pub fn query(deps: Deps, env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::RedemptionRate { denom } => query_redemption_rate(deps, env, denom),
    }
}

fn query_redemption_rate(
    deps: Deps,
    _env: Env,
    denom: String,
) -> StdResult<Binary> {
    // For the sake of this example, let's assume a fixed redemption rate
    let redemption_rate = match denom.as_str() {
        "stake" => 1.0,
        "liquid" => 0.5,
        _ => return Err(StdError::generic_err("Unsupported denom")),
    };

    to_binary(&redemption_rate)
}

fn add_stake(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    _params: AddStakeParams,
) -> StdResult<Response> {
    // For this example, let's just store the staked amount in the contract state
    let mut store = PrefixedStorage::new(deps.storage, b"stakes");
    
    let sent_funds = info.funds.clone();


    
   // Check that `sent_funds` method is available on the `MessageInfo` struct
/*
let sent_funds = if let Some(funds) = info.sent_funds() {
    funds
} else {
    return Err(StdError::generic_err("No funds were sent with the message"));
};
*/

  // Declare and initialize current_stake before using it
let current_stake = 0 as u128;
// Specify the type of `new_stake` variable as `u128` explicitly
let new_stake: u128 = current_stake + Uint128::from(sent_funds[0].amount).u128();

// Convert the address to bytes before storing it
let sender_address_raw = "some_sender_address".as_bytes().to_vec();
store.set(&sender_address_raw, &new_stake.to_be_bytes());



   // Return a `Response` object with the `add_attribute` method call
// Return a `Result` object with the `Response` value inside it
// Use the `with_attributes` method instead of the `add_attribute` method
 Ok(Response::default())

   
}

fn redeem_liquid_tokens(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    amount: u128,
) -> Result<Response,StdError> {
    // For this example, let's just remove the redeemed amount from the contract state
    let mut store = PrefixedStorage::new(deps.storage, b"stakes");
    let sender_address_raw = deps.api.addr_canonicalize(info.sender.as_str())?;
    let currentstake = store.get(&sender_address_raw);
    let current_stake = option_bytes_to_u128(currentstake); 
    if current_stake < amount {
        return Err(StdError::generic_err("Not enough staked tokens to redeem"));
    }

   // Convert the `current_stake` and `amount` variables to `Uint128`

let amount = amount as u128;
let new_stake = current_stake.checked_sub(amount);

   // Return a `Result` object with the `Response` value inside it
// Return a `Result` object with the `Response` value inside it
// Use the `with_attributes` method instead of the `add_attribute` method
Ok(Response::default())

}




// Add the missing structs and attributes
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct AddStakeParams {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum HandleMsg {
    AddStake { params: AddStakeParams },
    RedeemLiquidTokens { amount: Uint128 },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    RedemptionRate { denom: String },
}


#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct ContractState {
    pub owner: CanonicalAddr,
    pub total_staked: Uint128,
    pub total_liquid: Uint128,
}

impl ContractState {
    pub fn new(owner: CanonicalAddr) -> Self {
        ContractState {
            owner,
            total_staked: Uint128::zero(),
            total_liquid: Uint128::zero(),
        }
    }
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum ContractError {
    Unauthorized {},
    InsufficientFunds {},
    InvalidParameters {},
}

impl ContractError {
    pub fn unauthorized() -> Self {
        ContractError::Unauthorized {}
    }

    pub fn insufficient_funds() -> Self {
        ContractError::InsufficientFunds {}
    }

    pub fn invalid_parameters() -> Self {
        ContractError::InvalidParameters {}
    }
}

impl std::fmt::Display for ContractError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match *self {
            ContractError::Unauthorized {} => write!(f, "Unauthorized"),
            ContractError::InsufficientFunds {} => write!(f, "Insufficient funds"),
            ContractError::InvalidParameters {} => write!(f, "Invalid parameters"),
        }
    }
}

impl std::error::Error for ContractError {}

#[no_mangle]
pub fn interface_version_8() -> u32 {
    // Implementation of the interface_version_5 function
    1
}
