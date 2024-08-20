use cosmwasm_std::{attr, coin, to_binary, BankMsg, Binary, Deps, DepsMut, Env, HandleResponse, InitResponse, MessageInfo, StdError, StdResult, Uint128};
use cw2::set_contract_version;
use cw_storage_plus::{Map, Item};

const CONTRACT_NAME: &str = "crates.io:jumanji_journey_token";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct HandleMsg {
    #[serde(rename = "stake_tokens")]
    StakeTokens {},
    #[serde(rename = "unstake_tokens")]
    UnstakeTokens {},
    #[serde(rename = "create_proposal")]
    CreateProposal { description: String },
    #[serde(rename = "vote")]
    Vote { proposal_id: u64, vote: VoteOption },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum VoteOption {
    Yes,
    No,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct QueryMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Proposal {
    pub id: u64,
    pub description: String,
    pub yes_votes: Uint128,
    pub no_votes: Uint128,
    pub executed: bool,
}

const STAKES: Map<&[u8], u128> = Map::new("stakes");
const PROPOSALS: Map<u64, Proposal> = Map::new("proposals");
const PROPOSAL_COUNT: Item<u64> = Item::new("proposal_count");

pub fn init(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<InitResponse> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    PROPOSAL_COUNT.save(deps.storage, &0)?;
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
        HandleMsg::CreateProposal { description } => create_proposal(deps, info, description),
        HandleMsg::Vote { proposal_id, vote } => vote(deps, info, proposal_id, vote),
    }
}

fn stake_tokens(
    deps: DepsMut,
    info: MessageInfo,
) -> StdResult<HandleResponse> {
    let amount = info.sent_funds.iter().find(|c| c.denom == "ujjt").map(|c| c.amount.u128()).unwrap_or(0);
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
            amount: vec![coin(amount, "ujjt")],
        }.into()],
        attributes: vec![attr("action", "unstake_tokens"), attr("staker", info.sender)],
        data: None,
    })
}

fn create_proposal(
    deps: DepsMut,
    info: MessageInfo,
    description: String,
) -> StdResult<HandleResponse> {
    let mut proposal_count = PROPOSAL_COUNT.load(deps.storage)?;
    proposal_count += 1;

    let proposal = Proposal {
        id: proposal_count,
        description,
        yes_votes: Uint128::zero(),
        no_votes: Uint128::zero(),
        executed: false,
    };

    PROPOSALS.save(deps.storage, proposal_count, &proposal)?;
    PROPOSAL_COUNT.save(deps.storage, &proposal_count)?;

    Ok(HandleResponse {
        messages: vec![],
        attributes: vec![
            attr("action", "create_proposal"),
            attr("proposal_id", proposal_count.to_string()),
            attr("creator", info.sender),
        ],
        data: None,
    })
}

fn vote(
    deps: DepsMut,
    info: MessageInfo,
    proposal_id: u64,
    vote: VoteOption,
) -> StdResult<HandleResponse> {
    let sender = info.sender.as_bytes();
    let stake = STAKES.may_load(deps.storage, sender)?.unwrap_or_default();
    if stake == 0 {
        return Err(StdError::generic_err("No staked tokens to vote"));
    }

    let mut proposal = PROPOSALS.load(deps.storage, proposal_id)?;
    if proposal.executed {
        return Err(StdError::generic_err("Proposal already executed"));
    }

    match vote {
        VoteOption::Yes => proposal.yes_votes += Uint128::from(stake),
        VoteOption::No => proposal.no_votes += Uint128::from(stake),
    }

    PROPOSALS.save(deps.storage, proposal_id, &proposal)?;

    Ok(HandleResponse {
        messages: vec![],
        attributes: vec![
            attr("action", "vote"),
            attr("proposal_id", proposal_id.to_string()),
            attr("voter", info.sender),
            attr("vote", format!("{:?}", vote)),
        ],
        data: None,
    })
}

