use cosmwasm_std::{HumanAddr};
use cw721::{Cw721Contract, NftInfoResponse, Metadata};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub struct CampaignNFT {
    pub token_id: String,
    pub owner: HumanAddr,
    pub metadata: Metadata,
}

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> StdResult<Response> {
    match msg {
        ExecuteMsg::MintNft { token_id, metadata } => mint_nft(deps, env, info, token_id, metadata),
        ExecuteMsg::TransferNft { recipient, token_id } => transfer_nft(deps, env, info, recipient, token_id),
        ExecuteMsg::QueryNft { token_id } => query_nft(deps, env, info, token_id),
    }
}

fn mint_nft(deps: DepsMut, _env: Env, info: MessageInfo, token_id: String, metadata: Metadata) -> StdResult<Response> {
    let owner = info.sender.clone();
    let campaign_nft = CampaignNFT {
        token_id,
        owner,
        metadata,
    };
    // Store NFT data in storage
    Ok(Response::new().add_attribute("method", "mint_nft"))
}

fn transfer_nft(deps: DepsMut, _env: Env, info: MessageInfo, recipient: HumanAddr, token_id: String) -> StdResult<Response> {
    // Transfer logic; update the owner in the storage
    Ok(Response::new().add_attribute("method", "transfer_nft"))
}

fn query_nft(deps: Deps, _env: Env, info: MessageInfo, token_id: String) -> StdResult<Response> {
    // Query logic; Fetch NFT details from storage
    Ok(Response::new().add_attribute("method", "query_nft"))
}
