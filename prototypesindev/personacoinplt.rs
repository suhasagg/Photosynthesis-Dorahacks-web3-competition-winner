use cosmwasm_std::{
    log, to_binary, Api, Binary, CanonicalAddr, Env, Extern, HandleResponse, HumanAddr, InitResponse,
    Querier, QueryRequest, StdError, StdResult, Storage, Uint128,
};
use cosmwasm_storage::{PrefixedStorage, ReadonlyPrefixedStorage};
use serde::{Deserialize, Serialize};
use std::convert::TryInto;

const COOKIE_KEY: &[u8] = b"cookie";
const PORT_ID: &str = "cookiesync";
const TIMEOUT_HEIGHT: u64 = 100;
const TIMEOUT_TIMESTAMP: u64 = 0;

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct Cookie {
    pub id: String,
    pub domain: String,
    pub data: String,
    pub expiration: u64,
}

impl Cookie {
    pub fn is_expired(&self, env: &Env) -> bool {
        self.expiration <= env.block.time
    }
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct CookiePacketData {
    pub cookie: Cookie,
    pub source_chain: String,
    pub dest_chain: String,
}

impl CookiePacketData {
    pub fn new(cookie: Cookie, source_chain: String, dest_chain: String) -> Self {
        Self {
            cookie,
            source_chain,
            dest_chain,
        }
    }

    pub fn validate_basic(&self) -> StdResult<()> {
        if self.cookie.id.is_empty() || self.source_chain.is_empty() || self.dest_chain.is_empty() {
            return Err(StdError::generic_err("invalid packet data"));
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct State {
    pub plt_token: CanonicalAddr,
}

impl State {
    pub fn load(storage: &dyn Storage) -> StdResult<Self> {
        let state = ReadonlyPrefixedStorage::new(b"state.", storage);
        match state.get(b"plt_token") {
            Some(data) => Ok(Self {
                plt_token: CanonicalAddr::from(data),
            }),
            None => Err(StdError::generic_err("state not found")),
        }
    }

    pub fn save(&mut self, storage: &mut dyn Storage) -> StdResult<()> {
        let mut state = PrefixedStorage::new(b"state.", storage);
        state.set(b"plt_token", self.plt_token.as_slice());
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub enum HandleMsg {
    Sync(Cookie),
    MintPLT {
        recipient: HumanAddr,
        amount: Uint128,
    },
    BurnPLT {
        owner: HumanAddr,
        amount: Uint128,
    },
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub enum QueryMsg {
    Sync { cookie_id: String },
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct SyncResponse {
    pub cookie: Cookie,
}

pub fn init(
    deps: &mut Extern<impl Api + Storage + Querier>,
    _env: Env,
) -> StdResult<InitResponse> {
    let state = State {
        plt_token: deps
            .api
            .canonical_address(&deps.api.human_address("plt_token")?)?,
    };
    state.save(&mut deps.storage)?;

    Ok(InitResponse::default())
}

pub fn handle(
    deps: &mut Extern<impl Api + Storage + Querier>,
    env: Env,
    msg: HandleMsg,
) -> StdResult



pub fn handle(
    deps: &mut Extern<impl Api + Storage + Querier>,
    env: Env,
    msg: HandleMsg,
) -> StdResult
<HandleResponse> {
match msg {
HandleMsg::Sync(cookie) => sync(deps, env, cookie),
HandleMsg::MintPLT { recipient, amount } => mint_plt(deps.storage, env, recipient, amount),
HandleMsg::BurnPLT { amount } => burn_plt(deps.storage, env, info, amount),
}
}

pub fn sync(
deps: &mut Extern<impl Api + Storage + Querier + IbcChannel>,
env: Env,
cookie: Cookie,
) -> StdResult<HandleResponse> {
let source_port = env.contract.address.port.clone();
let source_channel = env.contract.address.channel_id.clone();
let dest_chain = cookie.domain.clone();
let packet_data = CookiePacketData::new(cookie, env.chain_id.clone(), dest_chain.clone());


let channel = deps.ibc_channel.create_or_join_channel(
    &env,
    HumanAddr::from(PORT_ID),
    HumanAddr::from(dest_chain),
    source_port.clone(),
    source_channel.clone(),
    TIMEOUT_HEIGHT,
    TIMEOUT_TIMESTAMP,
)?;

channel.send_packet(
    &deps.ibc_channel,
    to_binary(&packet_data)?,
    TIMEOUT_HEIGHT,
    TIMEOUT_TIMESTAMP,
)?;

let res = HandleResponse {
    messages: vec![],
    log: vec![
        log("action", "sync"),
        log("from", env.contract.address.to_string()),
        log("to", format!("{}:{}", PORT_ID, dest_chain)),
    ],
    data: None,
};
Ok(res)
}

pub fn mint_plt(
storage: &mut dyn Storage,
env: Env,
recipient: HumanAddr,
amount: Uint128,
) -> StdResult<HandleResponse> {
let mut state = State::load(storage)?;


let plt = deps.api.addr_validate(&state.plt_address)?;
let sender = env.message.sender;

if sender != plt {
    return Err(StdError::generic_err("unauthorized"));
}

state.plt_supply += amount;
store_state(storage, &state)?;

let cw20 = Cw20Coin {
    address: state.plt_address.to_string(),
    amount: amount,
};
let mint_msg = Cw20HandleMsg::Mint {
    recipient: recipient.into(),
    amount: amount.into(),
};
let msg = SubMsg::new(CosmosMsg::Wasm(WasmMsg::Execute {
    contract_addr: state.plt_address.to_string(),
    msg: to_binary(&mint_msg)?,
    funds: vec![],
}));

let res = HandleResponse {
    messages: vec![msg],
    log: vec![
        log("action", "mint_plt"),
        log("from", plt.to_string()),
        log("to", recipient.to_string()),
        log("amount", amount.to_string()),
    ],
    data: None,
};
Ok(res)
}

pub fn burn_plt(
storage: &mut dyn Storage,
env: Env,
info: MessageInfo,
amount: Uint128,
) -> StdResult<HandleResponse> {
let mut state = State::load(storage)?;


let plt = deps.api.addr_validate(&state.plt_address)?;
let sender = info.sender;

if sender != plt {
    return Err(StdError::generic_err("unauthorized"));
}

if state.plt_supply < amount {
    return Err(StdError::generic_err("insufficient supply"));
}

state.plt_supply -= amount;
store_state(storage, &state)?;

let cw20 = Cw20Coin {
    address: state.plt_address.to_string(),
    amount



 let cw20 = Cw20Coin {
    address: state.plt_address.to_string(),
    amount
: Uint128(amount),
};
let mint_msg = Cw20HandleMsg::Mint {
recipient: recipient.clone(),
amount: cw20.amount,
};
let mint = CosmosMsg::Wasm(WasmMsg::Execute {
contract_addr: cw20.address,
msg: to_binary(&mint_msg)?,
funds: vec![],
});
Ok(HandleResponse {
messages: vec![mint],
log: vec![
log("action", "mint_plt"),
log("recipient", recipient),
log("amount", amount.to_string()),
],
data: None,
})