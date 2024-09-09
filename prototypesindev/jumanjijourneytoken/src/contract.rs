use cosmwasm_std::{coin, to_json_binary, BankMsg, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdError, StdResult, Uint128, Addr, Empty};
use cw_storage_plus::Map;
use cw2::set_contract_version;
use serde::{Deserialize, Serialize};
use schemars::JsonSchema;

const CONTRACT_NAME: &str = "crates.io:jumanji_journey_token";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

// Messages
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InitMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum HandleMsg {
    StakeTokens {},
    UnstakeTokens {},
    CreateProposal { description: String },
    Vote { proposal_id: u64, vote: VoteOption },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum VoteOption {
    Yes,
    No,
}

// Proposal structure
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct Proposal {
    pub id: u64,
    pub description: String,
    pub yes_votes: Uint128,
    pub no_votes: Uint128,
    pub executed: bool,
}

// Query messages
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub enum QueryMsg {
    GetStake { address: String },
    GetProposal { proposal_id: u64 },
}

// Query responses
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakeResponse {
    pub stake: Uint128,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct ProposalResponse {
    pub proposal_id: u64,
    pub description: String,
}

// Storage
const STAKES: Map<&Addr, Uint128> = Map::new("stakes");
const PROPOSALS: Map<u64, Proposal> = Map::new("proposals");

// Initialize the contract
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> StdResult<Response> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    Ok(Response::new())
}

// Handle messages
pub fn execute(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> StdResult<Response> {
    match msg {
        HandleMsg::StakeTokens {} => stake_tokens(deps, info),
        HandleMsg::UnstakeTokens {} => unstake_tokens(deps, info),
        HandleMsg::CreateProposal { description } => create_proposal(deps, info, description),
        HandleMsg::Vote { proposal_id, vote } => process_vote(deps, info, proposal_id, vote),
    }
}


fn stake_tokens(
    deps: DepsMut,
    info: MessageInfo,
) -> StdResult<Response> {
    println!("Staking tokens for sender: {}", info.sender);
    let amount = info.funds.iter().find(|c| c.denom == "ujjt").map(|c| c.amount).unwrap_or(Uint128::zero());

    if amount.is_zero() {
        println!("No ujjt tokens sent for staking.");
        return Err(StdError::generic_err("Must send non-zero amount of ujjt tokens"));
    }

    let staker = info.sender.clone();
    let current_stake = STAKES.may_load(deps.storage, &staker)?.unwrap_or(Uint128::zero());
    let new_stake = current_stake.checked_add(amount).map_err(|_| StdError::generic_err("Overflow during staking"))?;

    println!("Staking {} ujjt tokens for staker: {}", amount, staker);
    STAKES.save(deps.storage, &staker, &new_stake)?;

    Ok(Response::new()
        .add_attribute("action", "stake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount", amount.to_string()))
}

fn unstake_tokens(
    deps: DepsMut,
    info: MessageInfo,
) -> StdResult<Response> {
    println!("Unstaking tokens for sender: {}", info.sender);
    let staker = info.sender.clone();
    let current_stake = STAKES.may_load(deps.storage, &staker)?.unwrap_or(Uint128::zero());

    if current_stake.is_zero() {
        println!("No tokens available for unstaking.");
        return Err(StdError::generic_err("No tokens to unstake"));
    }

    println!("Unstaking {} ujjt tokens for staker: {}", current_stake, staker);
    STAKES.save(deps.storage, &staker, &Uint128::zero())?;

    Ok(Response::new()
        .add_message(BankMsg::Send {
            to_address: staker.to_string(),
            amount: vec![coin(current_stake.u128(), "ujjt")],
        })
        .add_attribute("action", "unstake_tokens")
        .add_attribute("staker", staker)
        .add_attribute("amount", current_stake.to_string()))
}


// Create a proposal
fn create_proposal(
    deps: DepsMut,
    info: MessageInfo,
    description: String,
) -> StdResult<Response> {
    let proposal = Proposal {
        id: 1,
        description,
        yes_votes: Uint128::zero(),
        no_votes: Uint128::zero(),
        executed: false,
    };

    PROPOSALS.save(deps.storage, 1, &proposal)?;

    Ok(Response::new()
        .add_attribute("action", "create_proposal")
        .add_attribute("creator", info.sender))
}

// Process votes
fn process_vote(
    deps: DepsMut,
    info: MessageInfo,
    proposal_id: u64,
    vote: VoteOption,
) -> StdResult<Response> {
    let mut proposal = PROPOSALS.load(deps.storage, proposal_id)?;

    match vote {
        VoteOption::Yes => proposal.yes_votes += Uint128::from(1u128),
        VoteOption::No => proposal.no_votes += Uint128::from(1u128),
    }

    PROPOSALS.save(deps.storage, proposal_id, &proposal)?;

    Ok(Response::new()
        .add_attribute("action", "vote")
        .add_attribute("voter", info.sender))
}

// Query functions
pub fn query(
    deps: Deps,
    _env: Env,
    msg: QueryMsg,
) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetStake { address } => to_json_binary(&query_stake(deps, address)?),
        QueryMsg::GetProposal { proposal_id } => to_json_binary(&query_proposal(deps, proposal_id)?),
    }
}

// Query proposal
fn query_proposal(deps: Deps, proposal_id: u64) -> StdResult<ProposalResponse> {
    let proposal = PROPOSALS.load(deps.storage, proposal_id)?;
    Ok(ProposalResponse {
        proposal_id,
        description: proposal.description,
    })
}

// Query stake
fn query_stake(deps: Deps, address: String) -> StdResult<StakeResponse> {
    let addr = deps.api.addr_validate(&address)?;
    let stake = STAKES.may_load(deps.storage, &addr)?.unwrap_or(Uint128::zero());
    Ok(StakeResponse { stake })
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::{coins, Addr};
    use cw_multi_test::{App, Contract, ContractWrapper, Executor, SudoMsg, BankSudo}; 

    fn mock_app() -> App {
        App::default()
    }

    fn contract_jumanji() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

#[test]
fn test_stake_tokens() {
    println!("Testing stake_tokens function...");
    let mut app = mock_app();
    let code_id = app.store_code(contract_jumanji());
    let contract_addr = app.instantiate_contract(
        code_id,
        Addr::unchecked("creator"),
        &InitMsg {},
        &[],
        "JumanjiJourney",
        None,
    ).unwrap();

    println!("Minting tokens to creator address...");
    app.sudo(SudoMsg::Bank(BankSudo::Mint {
        to_address: "creator".to_string(),
        amount: vec![coin(100, "ujjt")],
    })).unwrap();

    let stake_msg = HandleMsg::StakeTokens {};
    let res = app.execute_contract(
        Addr::unchecked("creator"),
        contract_addr.clone(),
        &stake_msg,
        &coins(100, "ujjt"),
    ).unwrap();

    println!("Checking events for stake_tokens action...");
    let wasm_event = res.events.iter().find(|e| e.ty == "wasm").unwrap();
    let action_attr = wasm_event.attributes.iter().find(|a| a.key == "action").unwrap();
    assert_eq!(action_attr.value, "stake_tokens");
}

#[test]
fn test_unstake_tokens() {
    println!("Testing unstake_tokens function...");
    let mut app = mock_app();
    let code_id = app.store_code(contract_jumanji());
    let contract_addr = app.instantiate_contract(
        code_id,
        Addr::unchecked("creator"),
        &InitMsg {},
        &[],
        "JumanjiJourney",
        None,
    ).unwrap();

    println!("Minting tokens to creator address...");
    app.sudo(SudoMsg::Bank(BankSudo::Mint {
        to_address: "creator".to_string(),
        amount: vec![coin(100, "ujjt")],
    })).unwrap();

    println!("Staking tokens...");
    let stake_msg = HandleMsg::StakeTokens {};
    app.execute_contract(
        Addr::unchecked("creator"),
        contract_addr.clone(),
        &stake_msg,
        &coins(100, "ujjt"),
    ).unwrap();

    println!("Unstaking tokens...");
    let unstake_msg = HandleMsg::UnstakeTokens {};
    let res = app.execute_contract(
        Addr::unchecked("creator"),
        contract_addr.clone(),
        &unstake_msg,
        &[],
    ).unwrap();

    println!("Checking events for unstake_tokens action...");
    let wasm_event = res.events.iter().find(|e| e.ty == "wasm").unwrap();
    let action_attr = wasm_event.attributes.iter().find(|a| a.key == "action").unwrap();
    assert_eq!(action_attr.value, "unstake_tokens");
}

}
