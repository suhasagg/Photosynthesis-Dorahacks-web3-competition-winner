use cosmwasm_std::{
    Deps, DepsMut, Env, MessageInfo, Response, StdResult, Uint128, to_binary, ContractResult, WasmQuery
};
use cw_storage_plus::Item;
use cw20::{Cw20CoinHuman, Cw20ExecuteMsg, Cw20ReceiveMsg};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

use cw2::set_contract_version;

use crate::error::ContractError;
use crate::msg::{ExecuteMsg, GetCountResponse, InstantiateMsg, QueryMsg};
use crate::state::{State, STATE};

// version info for migration info
const CONTRACT_NAME: &str = "crates.io:dapps";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");




pub const DNA_DIGITS: u32 = 16;
pub const DNA_MODULUS: u64 = 10u64.pow(DNA_DIGITS);
pub const COOLDOWN_TIME: u64 = 0;

pub struct KittyInterface;

impl KittyInterface {
    pub fn get_kitty(&self, _kitty_id: u64) -> (
        bool,
        bool,
        u64,
        u64,
        u64,
        u64,
        u64,
        u64,
        u64,
        u64,
    ) {
        (
            true, true, 3, 4, 5, 6, 7, 8, 9, 7_688_748_911_342_991
        )
    }
}

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema)]
pub struct MatrixBrainBattery {
    pub name: String,
    pub dna: u64,
    pub level: u32,
    pub ready_time: u64,
    pub win_count: u16,
    pub loss_count: u16,
}

pub struct MatrixBrainBatteryFactory {
    pub kitty_interface: KittyInterface,
    pub matrix_brain_batteries: Vec<MatrixBrainBattery>,
    pub matrix_brain_battery_to_owner: HashMap<u64, String>,
    pub owner_matrix_brain_battery_count: HashMap<String, u64>,
}

impl MatrixBrainBatteryFactory {
    pub fn new() -> Self {
        MatrixBrainBatteryFactory {
            kitty_interface: KittyInterface,
            matrix_brain_batteries: Vec::new(),
            matrix_brain_battery_to_owner: HashMap::new(),
            owner_matrix_brain_battery_count: HashMap::new(),
        }
    }

    pub fn generate_random_dna(&self, _input_str: &str) -> u64 {
        let mut rng = rand::thread_rng();
        rng.gen_range(0, DNA_MODULUS)
    }

    pub fn create_random_matrix_brain_battery(&mut self, name: String, owner: String) -> StdResult<()> {
        let count = self.owner_matrix_brain_battery_count.get(&owner).cloned().unwrap_or(0);
        if count > 0 {
            return Err(StdError::generic_err("Owner already has a MatrixBrainBattery"));
        }

        let rand_dna = self.generate_random_dna(&name);
        let new_battery = MatrixBrainBattery {
            name,
            dna: rand_dna,
            level: 1,
            ready_time: env.block.time.plus_seconds(COOLDOWN_TIME).seconds(),
            win_count: 0,
            loss_count: 0,
        };
        self.matrix_brain_batteries.push(new_battery);
        let id = self.matrix_brain_batteries.len() - 1;
        self.matrix_brain_battery_to_owner.insert(id as u64, owner.clone());
        self.owner_matrix_brain_battery_count.insert(owner, 1);
        Ok(())
    }

      pub fn trigger_cooldown(&mut self, matrix_brain_battery: &mut MatrixBrainBattery) {
        matrix_brain_battery.ready_time = env.block.time.plus_seconds(COOLDOWN_TIME).seconds();
    }

    pub fn is_ready(&self, matrix_brain_battery: &MatrixBrainBattery) -> bool {
        matrix_brain_battery.ready_time <= env.block.time.seconds()
    }

    pub fn feed_and_multiply(
        &mut self,
        matrix_brain_battery_id: u64,
        target_dna: u64,
        species: &str,
        owner: &str,
    ) -> StdResult<()> {
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        if !self.is_ready(&matrix_brain_battery) {
            return Err(StdError::generic_err("MatrixBrainBattery not ready"));
        }

        let new_dna = (matrix_brain_battery.dna + target_dna % DNA_MODULUS) / 2;
        let new_dna = if species == "kitty" {
            new_dna - new_dna % 100 + 99
        } else {
            new_dna
        };

        let new_battery = MatrixBrainBattery {
            name: "NoName".to_string(),
            dna: new_dna,
            level: 1,
            ready_time: env.block.time.plus_seconds(COOLDOWN_TIME).seconds(),
            win_count: 0,
            loss_count: 0,
        };
        self.matrix_brain_batteries.push(new_battery);
        let id = self.matrix_brain_batteries.len() - 1;
        self.matrix_brain_battery_to_owner.insert(id as u64, owner.to_string());
        self.owner_matrix_brain_battery_count
            .entry(owner.to_string())
            .and_modify(|count| *count += 1)
            .or_insert(1);
        self.trigger_cooldown(&mut matrix_brain_battery);
        Ok(())
    }

    pub fn feed_on_kitty(
        &mut self,
        matrix_brain_battery_id: u64,
        kitty_id: u64,
        owner: &str,
    ) -> StdResult<()> {
        let (_, _, _, _, _, _, _, _, _, kitty_dna) = self.kitty_interface.get_kitty(kitty_id);
        self.feed_and_multiply(matrix_brain_battery_id, kitty_dna, "kitty", owner)
    }
}

use cosmwasm_std::{
    // Other required imports
    Uint128, StdResult,
};
use cw721::{Cw721, Cw721Contract, Expiration, Approval};

// MatrixBrainBatteryHelper
impl MatrixBrainBatteryFactory {
    pub const LEVEL_UP_FEE: u128 = 1_000_000_000; // 0.001 SCRT
    
    
    pub fn set_level_up_fee(&mut self, fee: Uint128) {
        self.level_up_fee = fee;
    }

    pub fn level_up(&mut self, matrix_brain_battery_id: u64) -> StdResult<()> {
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        matrix_brain_battery.level += 1;
        Ok(())
    }

    pub fn change_name(&mut self, matrix_brain_battery_id: u64, new_name: &str) -> StdResult<()> {
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        matrix_brain_battery.name = new_name.to_string();
        Ok(())
    }

    pub fn change_dna(&mut self, matrix_brain_battery_id: u64, new_dna: u64) -> StdResult<()> {
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        matrix_brain_battery.dna = new_dna;
        Ok(())
    }

    pub fn get_matrix_brain_batteries_by_owner(&self, owner: &str) -> Vec<u64> {
        let mut result = Vec::new();
        for (i, matrix_brain_battery) in self.matrix_brain_batteries.iter().enumerate() {
            if let Some(matrix_brain_battery_owner) = self.matrix_brain_battery_to_owner.get(&(i as u64)) {
                if matrix_brain_battery_owner == owner {
                    result.push(i as u64);
                }
            }
        }
        result
    }
}

// MatrixBrainBatteryAttack
impl MatrixBrainBatteryFactory {
    pub const ATTACK_VICTORY_PROBABILITY: u64 = 70;

    pub fn rand_mod(&mut self, modulus: u64) -> u64 {
        self.rand_nonce += 1;
        let seed = format!("{}{}{}", env.block.time.seconds(), env.message.sender, self.rand_nonce);
        let hash = std::u64::from_be_bytes(sha256(seed.as_bytes()));
        hash % modulus
    }

    pub fn attack(&mut self, matrix_brain_battery_id: u64, target_id: u64) -> StdResult<()> {
        let mut my_matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        let enemy_matrix_brain_battery = self.matrix_brain_batteries
            .get(target_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid target MatrixBrainBattery ID"))?;
        let rand = self.rand_mod(100);
                if rand <= Self::ATTACK_VICTORY_PROBABILITY {
            my_matrix_brain_battery.win_count += 1;
            my_matrix_brain_battery.level += 1;
            enemy_matrix_brain_battery.loss_count += 1;
            self.feed_and_multiply(matrix_brain_battery_id, enemy_matrix_brain_battery.dna, "matrixbrainbattery")?;
        } else {
            my_matrix_brain_battery.loss_count += 1;
            enemy_matrix_brain_battery.win_count += 1;
            self.trigger_cooldown(my_matrix_brain_battery)?;
        }
        Ok(())
    }
}

// MatrixBrainBatteryOwnership
impl MatrixBrainBatteryFactory {
    pub fn transfer_ownership(&mut self, from: &str, to: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
        if self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id) != Some(&from.to_string()) {
            return Err(StdError::generic_err("Only owner can transfer ownership"));
        }

        self.matrix_brain_battery_to_owner.insert(matrix_brain_battery_id, to.to_string());
        Ok(())
    }

    pub fn approve(&mut self, approver: &str, to: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
        if self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id) != Some(&approver.to_string()) {
            return Err(StdError::generic_err("Only owner can approve"));
        }

        self.matrix_brain_battery_approvals.insert(matrix_brain_battery_id, to.to_string());
        Ok(())
    }

    pub fn take_ownership(&mut self, new_owner: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
        let approval = self.matrix_brain_battery_approvals.get(&matrix_brain_battery_id);

        if approval != Some(&new_owner.to_string()) {
            return Err(StdError::generic_err("Only approved address can take ownership"));
        }

        let owner = self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id).ok_or_else(|| {
            StdError::generic_err("MatrixBrainBattery not found")
        })?;

        self.transfer_ownership(&owner, new_owner, matrix_brain_battery_id)
    }
}

