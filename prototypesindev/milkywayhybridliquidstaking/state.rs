use cosmwasm_std::Uint128;
use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakeInfo {
    pub amount: Uint128,
    pub last_updated_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct RewardInfo {
    pub total_rewards: Uint128,
    pub last_claim_epoch: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct State {
    pub current_epoch: u64,
    pub staking_info: HashMap<String, StakeInfo>,
    pub reward_info: HashMap<String, RewardInfo>,
    pub reward_records: HashMap<String, Vec<RewardRecord>>,
    pub total_liquid_tokens: Uint128,
    pub total_redemption_tokens: Uint128,
    pub contract_addresses: Vec<String>,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct RewardRecord {
    pub rewards: Uint128,
    pub calculated_time: String,
    pub calculated_height: u64,
}

pub const STATE: Item<State> = Item::new("state");

