use cosmwasm_std::{attr, to_binary, BankMsg, Binary, Deps, DepsMut, Env, HandleResponse, InitResponse, MessageInfo, StdError, StdResult};
use cw2::set_contract_version;
use cw_storage_plus::Map;

const CONTRACT_NAME: &str = "crates.io:john_wick_coin";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct HandleMsg {
    #[serde(rename = "stake_tokens")]
    StakeTokens {},
    #[serde(rename = "unstake_tokens")]
    UnstakeTokens {},
    #[serde(rename = "vote")]
    Vote { proposal_id: u64 },   
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryMsg {
}

const STAKES: Map<&[u8], u128> = Map::new("stakes"); 

pub fn init(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<InitResponse> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    Ok(InitResponse::default())
}

pub fn handle(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<HandleResponse> {
    match msg {
        HandleMsg::StakeTokens {} => stake_tokens(deps, info),
        HandleMsg::UnstakeTokens {} => unstake_tokens(deps, info),
        HandleMsg::Vote { proposal_id } => vote(deps, info, proposal_id),
    }
}


fn stake_tokens(
    deps: DepsMut,
    info: MessageInfo,
) -> StdResult<HandleResponse> {
    let amount = info.sent_funds.iter().find(|c| c.denom == "ujwc").map(|c| c.amount.u128()).unwrap_or(0);
    let sender = info.sender.as_bytes();
    let mut stake = STAKES.may_load(deps.storage, sender)?.unwrap_or_default();
    stake += amount;
    STAKES.save(deps.storage, sender, &stake)?;
    
    Ok(HandleResponse {
        messages: vec![],
        attributes: vec![attr("action", "stake_tokens"), attr("staker", info.sender)],
        data: None,
    })
}

fn unstake_tokens(
    deps: DepsMut,
    info: MessageInfo,
) -> StdResult<HandleResponse> {
    let sender = info.sender.as_bytes();
    let mut stake = STAKES.may_load(deps.storage, sender)?.unwrap_or_default();
    if stake == 0 {
        return Err(StdError::generic_err("No tokens to unstake"));
    }
    let amount = stake;
    stake = 0;
    STAKES.save(deps.storage, sender, &stake)?;

    Ok(HandleResponse {
        messages: vec![BankMsg::Send {
            to_address: info.sender.to_string(),
            amount: vec![coin(amount, "ujwc")],
        }.into()],
        attributes: vec![attr("action", "unstake_tokens"), attr("staker", info.sender)],
        data: None,
    })
}

fn vote(
    deps: DepsMut,
    info: MessageInfo,
    proposal_id: u64,
) -> StdResult<HandleResponse> {
    let sender = info.sender.as_bytes();
    let stake = STAKES.may_load(deps.storage, sender)?.unwrap_or_default();
    if stake == 0 {
        return Err(StdError::generic_err("No staked tokens to vote"));
    }
   
    Ok(HandleResponse {
        messages: vec![],
        attributes: vec![attr("action", "vote"), attr("voter", info.sender)],
        data: None,
    })
}
