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

pub fn init(
    _deps: &mut Extern<impl Api + Storage + Querier>,
    _env: Env,
) -> StdResult<InitResponse> {
    Ok(InitResponse::default())
}

pub fn handle(
    deps: &mut Extern<impl Api + Storage + Querier>,
    env: Env,
    msg: HandleMsg,
) -> StdResult<HandleResponse> {
    match msg {
        HandleMsg::Sync(cookie) => sync(deps, env, cookie),
    }
}

pub fn sync(
    deps: &mut Extern<impl Api + Storage + Querier>,
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

pub fn query(
    deps: &Extern<impl Storage>,
    _env: Env,
    msg: QueryMsg,
) -> StdResult<Binary> {
    match msg {
        QueryMsg(Sync { cookie_id }) => {
let cookie = match get_cookie(deps.storage, &cookie_id) {
Some(cookie) => cookie,
None => {
return Err(StdError::generic_err(format!(
"Cookie with id {} not found",
cookie_id
)))
}
};


        to_binary(&SyncResponse { cookie })
    }
}
}

#[cfg(test)]
mod tests {
use super::*;
use cosmwasm_std::testing::{mock_dependencies, mock_env};
use cosmwasm_std::{coins, from_binary};


#[test]
fn proper_initialization() {
    let mut deps = mock_dependencies(&[]);

    let msg = InitMsg {};
    let env = mock_env();
    let res = init(&mut deps, env, msg).unwrap();
    assert_eq!(0, res.messages.len());

    let query_res = query(&deps, mock_env(), QueryMsg::Sync { cookie_id: String::new() }).unwrap();
    let value: SyncResponse = from_binary(&query_res).unwrap();
    assert_eq!(
        value,
        SyncResponse {
            cookie: Cookie {
                id: String::new(),
                domain: String::new(),
                data: String::new(),
                expiration: String::new()
            }
        }
    );
}

#[test]
fn sync_cookie() {
    let mut deps = mock_dependencies(&[]);

    let msg = InitMsg {};
    let env = mock_env();
    let _res = init(&mut deps, env, msg).unwrap();

    let cookie = Cookie {
        id: "123".to_string(),
        domain: "example.com".to_string(),
        data: "cookie_data".to_string(),
        expiration: "2022-01-01".to_string(),
    };

    let sync_msg = SyncMsg { cookie };
    let env = mock_env();
    let info = mock_info("sender", &coins(1000, "earth"));
    let res = handle(&mut deps, env, info, sync_msg).unwrap();
    assert_eq!(0, res.messages.len());

    let query_res = query(&deps, mock_env(), QueryMsg::Sync { cookie_id: "123".to_string() }).unwrap();
    let value: SyncResponse = from_binary(&query_res).unwrap();
    assert_eq!(value.cookie, cookie);
}

#[test]
fn sync_cookie_fails_if_not_enough_funds() {
    let mut deps = mock_dependencies(&[]);

    let msg = InitMsg {};
    let env = mock_env();
    let _res = init(&mut deps, env, msg).unwrap();

    let cookie = Cookie {
        id: "123".to_string(),
        domain: "example.com".to_string(),
        data: "cookie_data".to_string(),
        expiration: "2022-01-01".to_string(),
    };

    let sync_msg = SyncMsg { cookie };
    let env = mock_env();
    let info = mock_info("sender", &coins(99, "earth"));
    let res = handle(&mut deps, env, info, sync_msg);
    assert!(res.is_err());
}

#[test]
fn sync_cookie_fails_if_cookie_already_exists() {
    let mut deps = mock_dependencies(&[]);

    let msg = InitMsg {};
    let env = mock_env();
    let _res = init(&mut deps, env, msg).unwrap();

    let cookie = Cookie {
        id: "123".to_string(),
        domain: "example.com".to_string(),
        data: "cookie_data".to_string(),
        expiration: "2022-01-01".to_string(),
    };

// handle the message types
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
                let (_, value) = item?;
                Ok(from_slice::<Cookie>(&value)?)
            })
            .collect();
        to_binary(&QueryAnswer::Cookies { cookies: cookies? })
    }
    // handle the QueryMsg::GetParams message
    QueryMsg::GetParams {} => {
        let params = Config::load(deps.storage)?;
        to_binary(&QueryAnswer::Params { params })
    }
}

// handle the message types
match msg {
    // handle the HandleMsg::CreateCookie message
    HandleMsg::CreateCookie {} => {
        let cookie = Cookie {
            id: "123".to_string(),
            domain: "example.com".to_string(),
            data: "cookie_data".to_string(),
            expiration: "2022-01-01".to_string(),
        };
        deps.storage.set(cookie.id.as_bytes(), &to_binary(&cookie)?);
        Ok(HandleResponse {
            messages: vec![],
            attributes: vec![attr("action", "create_cookie")],
            data: None,
        })
    }
}

fn (k Keeper) OnRecvPacket(ctx sdk.Context, packet exported.PacketI, data types.CookieSyncPacketData) error {
    // Store the received cookie data on the destination chain
    k.AppendCookie(ctx, data.Cookie)

    // Send a reward transaction
    recipient := data.Cookie.Domain
    reward := archway_sdk::Reward {
        points: 10,
        memo: "Reward for cookie synchronization",
        sender: k.rewardSender,
    }
    reward_tx := archway_sdk::send_reward(recipient, reward)

    // Emit an event
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(
            types.EventTypeReward,
            sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
            sdk.NewAttribute(types.AttributeKeyRewardPoints, fmt.Sprintf("%d", reward.points)),
            sdk.NewAttribute(types.AttributeKeyRewardTxHash, reward_tx.hash),
        ),
    )

    return nil
}

fn (c *CookieSync) Init(ctx sdk.Context, req *types.InitRequest) error {
    c.keeper = *NewKeeper()

    // Set the reward sender address
    c.rewardSender = req.RewardSender

    return nil
}
