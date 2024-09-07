use cosmwasm_std::{DepsMut, Env, MessageInfo, Response, StdError, StdResult, Deps, entry_point, Addr, Binary, to_binary};
use serde::{Deserialize, Serialize};
use cw_storage_plus::{Map, Item};
use cosmwasm_std::Empty;
use cosmwasm_std::to_json_binary;

pub const CONFIG: Item<Config> = Item::new("config");

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Metadata {
    pub description: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct User {
    pub id: String,
    pub campaigns_created: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct CampaignNFT {
    pub token_id: String,
    pub owner: String,
    pub metadata: Metadata,
}

pub const NFTS: Map<&str, CampaignNFT> = Map::new("nfts");
pub const OWNERS: Map<&str, String> = Map::new("owners");

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    GetNft { token_id: String },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub enum ExecuteMsg {
    MintNft { token_id: String, metadata: Metadata },
    TransferNft { recipient: String, token_id: String },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct InitMsg {
    pub owner: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct Config {
    pub owner: Addr,
    pub users: Vec<User>,
}

#[entry_point]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetNft { token_id } => {
            // Log the query action
            println!("Querying NFT with token_id: {}", token_id);
            to_binary(&query_nft(deps, token_id)?)
        },
    }
}

#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: InitMsg,
) -> StdResult<Response> {
    // Log the instantiation action
    println!("Instantiating contract with owner: {}", msg.owner);

    let config = Config {
        owner: deps.api.addr_validate(&msg.owner)?,
        users: vec![],
    };
    CONFIG.save(deps.storage, &config)?;

    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("owner", msg.owner))
}

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    match msg {
        ExecuteMsg::MintNft { token_id, metadata } => {
            println!("Minting NFT with token_id: {} by {}", token_id, info.sender);
            mint_nft(deps, env, info, token_id, metadata)
        },
        ExecuteMsg::TransferNft { recipient, token_id } => {
            println!("Transferring NFT with token_id: {} to {}", token_id, recipient);
            transfer_nft(deps, env, info, recipient, token_id)
        },
    }
}

fn mint_nft(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    token_id: String,
    metadata: Metadata,
) -> StdResult<Response> {
    if NFTS.may_load(deps.storage, &token_id)?.is_some() {
        return Err(StdError::generic_err("NFT already exists"));
    }

    // Log NFT creation details
    println!(
        "Saving NFT with token_id: {}, owner: {}, description: {:?}",
        token_id, info.sender, metadata.description
    );

    let owner = info.sender.to_string();
    let campaign_nft = CampaignNFT {
        token_id: token_id.clone(),
        owner: owner.clone(),
        metadata,
    };

    NFTS.save(deps.storage, &token_id, &campaign_nft)?;
    OWNERS.save(deps.storage, &token_id, &owner)?;

    Ok(Response::new()
        .add_attribute("method", "mint_nft")
        .add_attribute("token_id", token_id)
        .add_attribute("owner", owner))
}

fn transfer_nft(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    recipient: String,
    token_id: String,
) -> StdResult<Response> {
    let nft = NFTS.may_load(deps.storage, &token_id)?.ok_or_else(|| StdError::generic_err("NFT not found"))?;

    if nft.owner != info.sender.to_string() {
        return Err(StdError::generic_err("Unauthorized"));
    }

    // Log NFT transfer details
    println!(
        "Transferring NFT with token_id: {} from {} to {}",
        token_id, info.sender, recipient
    );

    let updated_nft = CampaignNFT {
        owner: recipient.clone(),
        ..nft
    };

    NFTS.save(deps.storage, &token_id, &updated_nft)?;
    OWNERS.save(deps.storage, &token_id, &recipient)?;

    Ok(Response::new()
        .add_attribute("method", "transfer_nft")
        .add_attribute("recipient", recipient)
        .add_attribute("token_id", token_id))
}

fn query_nft(
    deps: Deps,
    token_id: String,
) -> StdResult<CampaignNFT> {
    // Log NFT query details
    println!("Fetching NFT details for token_id: {}", token_id);

    let nft = NFTS.may_load(deps.storage, &token_id)?.ok_or_else(|| StdError::generic_err("NFT not found"))?;
    Ok(nft)
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::Addr;
    use cw_multi_test::{App, Contract, ContractWrapper, Executor};

    fn mock_app() -> App {
        App::default()
    }

    fn contract_campaign() -> Box<dyn Contract<Empty>> {
        let contract = ContractWrapper::new(execute, instantiate, query);
        Box::new(contract)
    }

  #[test]
  fn test_nft_minting() {
    let mut app = mock_app();
    let code_id = app.store_code(contract_campaign());
    let contract_addr = app.instantiate_contract(
        code_id,
        Addr::unchecked("creator"),
        &InitMsg { owner: "creator".to_string() },
        &[],
        "CampaignContract",
        None,
    ).unwrap();

    let mint_msg = ExecuteMsg::MintNft {
        token_id: "nft1".to_string(),
        metadata: Metadata { description: Some("Test NFT".to_string()) },
    };
    let res = app.execute_contract(Addr::unchecked("user1"), contract_addr.clone(), &mint_msg, &[]).unwrap();

    let wasm_event = res.events.iter().find(|e| e.ty == "wasm").unwrap();
    let relevant_attributes: Vec<(&str, &str)> = wasm_event
        .attributes
        .iter()
        .filter(|attr| attr.key != "_contract_addr")
        .map(|attr| (attr.key.as_str(), attr.value.as_str()))
        .collect();

   
    assert_eq!(relevant_attributes, vec![
        ("method", "mint_nft"), 
        ("token_id", "nft1"),
        ("owner", "user1")
    ]);

    let nft: CampaignNFT = app.wrap().query_wasm_smart(contract_addr.clone(), &QueryMsg::GetNft { token_id: "nft1".to_string() }).unwrap();
    assert_eq!(nft.token_id, "nft1");
    assert_eq!(nft.metadata.description.unwrap(), "Test NFT");
   }


    #[test]
    fn test_nft_transfer() {
        let mut app = mock_app();
        let code_id = app.store_code(contract_campaign());
        let contract_addr = app.instantiate_contract(
            code_id,
            Addr::unchecked("creator"),
            &InitMsg { owner: "creator".to_string() },
            &[],
            "CampaignContract",
            None,
        ).unwrap();

        let mint_msg = ExecuteMsg::MintNft {
            token_id: "nft1".to_string(),
            metadata: Metadata { description: Some("Test NFT".to_string()) },
        };
        app.execute_contract(Addr::unchecked("user1"), contract_addr.clone(), &mint_msg, &[]).unwrap();

        let transfer_msg = ExecuteMsg::TransferNft {
            recipient: "user2".to_string(),
            token_id: "nft1".to_string(),
        };
        app.execute_contract(Addr::unchecked("user1"), contract_addr.clone(), &transfer_msg, &[]).unwrap();

        let nft: CampaignNFT = app.wrap().query_wasm_smart(contract_addr.clone(), &QueryMsg::GetNft { token_id: "nft1".to_string() }).unwrap();
        assert_eq!(nft.owner, "user2".to_string());
    }
}    
