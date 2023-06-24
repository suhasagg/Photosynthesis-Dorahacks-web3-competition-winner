use cosmwasm_std::{
    Deps, DepsMut, Env, BlockInfo,ContractInfo,MessageInfo,Response, StdResult, Uint128, to_binary, ContractResult, WasmQuery
};
use cw_storage_plus::Item;
use cw20::{Cw20ExecuteMsg, Cw20ReceiveMsg};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use cosmwasm_std::StdError;
use cw2::set_contract_version;
use schemars::JsonSchema;
use sha2::{Sha256, Digest};
use rand::Rng;
use cosmwasm_std::Timestamp;
use cosmwasm_std::TransactionInfo;
use cosmwasm_std::Addr;
use std::time::{SystemTime, UNIX_EPOCH};

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

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug, JsonSchema)]
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
    pub matrix_brain_battery_approvals: HashMap<u64, String>,
    pub owner_matrix_brain_battery_count: HashMap<String, u64>,
    pub level_up_fee:  Uint128,
    pub rand_nonce: u32,
}

impl Default for MatrixBrainBattery {
    fn default() -> Self {
        Self {
            name: "".to_string(),
            dna: 0, // Or any default value
            level: 0, // Or any default value
            ready_time: 0, // Or any default value
            win_count: 0, // Or any default value
            loss_count: 0, // Or any default value
        }
    }
}

impl MatrixBrainBatteryFactory {
    pub fn new() -> Self {
        MatrixBrainBatteryFactory {
            kitty_interface: KittyInterface,
            matrix_brain_batteries: Vec::new(),
            matrix_brain_battery_to_owner: HashMap::new(),
            matrix_brain_battery_approvals:  HashMap::new(),
            owner_matrix_brain_battery_count: HashMap::new(),
            level_up_fee: Uint128::zero(),
            rand_nonce: 0,
        }
    }

    pub fn generate_random_dna(&self, _input_str: &str) -> u64 {
        //let mut rng = rand::thread_rng();
        //rng.gen_range(0, DNA_MODULUS)
        let mut rng = rand::thread_rng();
        rng.gen_range(0..DNA_MODULUS)
        
    }

    pub fn create_random_matrix_brain_battery(&mut self, name: String, owner: String, env: Env) -> StdResult<MatrixBrainBattery> {
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
        let new_battery_clone = new_battery.clone();
        
        //self.matrix_brain_batteries.push(new_battery);
        self.matrix_brain_batteries.push(new_battery_clone);
        let id = self.matrix_brain_batteries.len() - 1;
        self.matrix_brain_battery_to_owner.insert(id as u64, owner.clone());
        self.owner_matrix_brain_battery_count.insert(owner, 1);
        Ok(new_battery)
    }
     
    pub fn trigger_cooldown(&mut self, index: usize, env: &Env) {
    let matrix_brain_battery = self.matrix_brain_batteries
        .get_mut(index)
        .expect("MatrixBrainBattery not found"); 
     matrix_brain_battery.ready_time = env.block.time.plus_seconds(COOLDOWN_TIME).seconds();
    }
   
   pub fn is_ready(matrix_brain_battery: &MatrixBrainBattery, env: &Env) -> bool {
    matrix_brain_battery.ready_time <= env.block.time.seconds()
   }


    pub fn feed_and_multiply(
        &mut self,
        matrix_brain_battery_id: u64,
        //matrix_brain_battery: &mut MatrixBrainBattery,
        target_dna: u64,
        species: &str,
        owner: &str,
        env: Env
    ) -> StdResult<()> {
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
        if !MatrixBrainBatteryFactory::is_ready(&matrix_brain_battery, &env) {
            return Err(StdError::generic_err("MatrixBrainBattery not ready"));
        }

        let new_dna = (matrix_brain_battery.dna + target_dna % DNA_MODULUS) / 2;
        let new_dna = if species == "kitty" {
            new_dna - new_dna % 100 + 99
        } else {
            new_dna
        };
       
     let new_battery = MatrixBrainBattery {
        name: "BrainAxie".to_string(),
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
   
    self.trigger_cooldown(matrix_brain_battery_id as usize, &env);
    Ok(())
       

    }

    pub fn feed_on_kitty(
        &mut self,
        matrix_brain_battery_id: u64,
        kitty_id: u64,
        owner: &str,
        env: Env
    ) -> StdResult<()> {
        let (_, _, _, _, _, _, _, _, _, kitty_dna) = self.kitty_interface.get_kitty(kitty_id);
        self.feed_and_multiply(matrix_brain_battery_id, kitty_dna, "kitty", owner, env)
    }

    pub fn print_matrix_brain_batteries(&self) {
        println!("{:#?}", self.matrix_brain_batteries);
    }
}

use cw721::{Cw721, Expiration, Approval};

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

    fn rand_mod(&mut self, modulo: u64, env: &Env) -> u64 {
        self.rand_nonce += 1;
        let seed = format!("{}{}", env.block.time.seconds(), self.rand_nonce);
        //let hash = u64::from_be_bytes(Sha256::digest(seed.as_bytes()));
        let hash = Sha256::digest(seed.as_bytes());
        let mut array = [0u8; 8];
        array.copy_from_slice(&hash[0..8]);
        let num = u64::from_be_bytes(array);
        num
    }

    pub fn attack(&mut self,matrix_brain_battery_id: u64, my_matrix_brain_battery:&mut MatrixBrainBattery,enemy_matrix_brain_battery:&mut MatrixBrainBattery, owner: &str, env: Env) -> StdResult<()> {     
      
      let rand = self.rand_mod(100, &env);  
        if rand <= Self::ATTACK_VICTORY_PROBABILITY {
            my_matrix_brain_battery.win_count += 1;
            my_matrix_brain_battery.level += 1;
            enemy_matrix_brain_battery.loss_count += 1;
            self.feed_and_multiply(matrix_brain_battery_id, enemy_matrix_brain_battery.dna, "matrixbrainbattery", owner, env)?;
            //Self::feed_and_multiply(my_matrix_brain_battery, enemy_matrix_brain_battery.dna, "matrixbrainbattery", owner, env)?;

        } else {
            my_matrix_brain_battery.loss_count += 1;
            enemy_matrix_brain_battery.win_count += 1;
       
          let target_name = "BrainAxie".to_string();
          let my_matrix_brain_battery_index = self.matrix_brain_batteries.iter()
            .position(|battery| battery.name == target_name)
            .expect("Battery not found");

           self.trigger_cooldown(my_matrix_brain_battery_index, &env);
        
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
    
        let owner = self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id)
        .ok_or_else(|| {
            StdError::generic_err("MatrixBrainBattery not found")
        })?.clone();  // note the clone() here
    
        self.transfer_ownership(&owner, new_owner, matrix_brain_battery_id)
    }  
}

fn main() {
    // Initialize the factory  
    let mut factory = MatrixBrainBatteryFactory::new();
    let now = SystemTime::now();
    // Assume the following initial parameters
    let owner = "Owner1".to_string();
    let name = "Battery1".to_string();
    let timestamp = now.duration_since(UNIX_EPOCH)
    .expect("System time is before the UNIX epoch")
    .as_secs();

    let owner2 = "Owner2".to_string();
    let name2 = "Battery2".to_string();
    let timestamp = now.duration_since(UNIX_EPOCH)
    .expect("System time is before the UNIX epoch")
    .as_secs(); 
    // Initialize the environment
    // Note: You'll need to replace this with a real Env initialization
    let env = Env {
            block: BlockInfo {
                height: 0,
                time: Timestamp::from_nanos(1_571_797_419_879_305_533),
                chain_id: "localnet".to_string(),
                },
            transaction: Some(TransactionInfo { index: 3 }),
            contract: ContractInfo {
                   address: Addr::unchecked("contract"),
            },
        };

    let mut matrix_brain_battery1 = MatrixBrainBattery::default();

    let result1 = factory.create_random_matrix_brain_battery(name, owner.clone(), env.clone());
    
     if let Ok(battery) = result1 {
       // Copy the battery to your placeholder
         matrix_brain_battery1 = battery;
     } else {
    // Handle error if result1 is Err
    println!("Error creating battery");
    }
    factory.print_matrix_brain_batteries();

    let mut matrix_brain_battery2 = MatrixBrainBattery::default();

    let result2 = factory.create_random_matrix_brain_battery(name2, owner2.clone(), env.clone()); 
    ;
     if let Ok(battery) = result2 {
       // Copy the battery to your placeholder
         matrix_brain_battery2 = battery;
     } else {
    // Handle error if result1 is Err
    println!("Error creating battery");
    }
    factory.print_matrix_brain_batteries();

    // Feed on Kitty and multiply
    let kitty_id: u64 = 1; // Assume an existing Kitty id
    let matrix_brain_battery_id: u64 = 0; // Assume the first MatrixBrainBattery's id
    let result = factory.feed_on_kitty(matrix_brain_battery_id, kitty_id, &owner, env.clone());
    let kitty_id_usize = kitty_id as usize;
    let matrix_brain_battery_id_usize = matrix_brain_battery_id as usize;
    match result {
        Ok(()) => println!("MatrixBrainBattery fed on Kitty and multiplied successfully"),
        Err(e) => eprintln!("Failed to feed on Kitty and multiply: {}", e),
    }
    factory.print_matrix_brain_batteries();

    // Perform an attack
    let target_id: u64 = 1; // Assume another MatrixBrainBattery id for target

    let result = factory.attack(matrix_brain_battery_id, &mut matrix_brain_battery1, &mut matrix_brain_battery2, &owner, env.clone());
    
    match result {
        Ok(()) => println!("Attack was successful"),
        Err(e) => eprintln!("Attack failed: {}", e),
    }
    factory.print_matrix_brain_batteries();
    
    // Level up the MatrixBrainBattery
    let result = factory.level_up(matrix_brain_battery_id);

    match result {
        Ok(()) => println!("MatrixBrainBattery leveled up successfully"),
        Err(e) => eprintln!("Failed to level up MatrixBrainBattery: {}", e),
    }
    factory.print_matrix_brain_batteries();

    // Change the name of the MatrixBrainBattery
    let new_name = "Super Battery";
    let result = factory.change_name(matrix_brain_battery_id, new_name);

    match result {
        Ok(()) => println!("MatrixBrainBattery name changed successfully"),
        Err(e) => eprintln!("Failed to change MatrixBrainBattery name: {}", e),
    }
    factory.print_matrix_brain_batteries();

    // List all MatrixBrainBatteries by owner
    let owned_matrix_brain_batteries = factory.get_matrix_brain_batteries_by_owner(&owner);

    println!("{} owns the following MatrixBrainBatteries: {:?}", owner, owned_matrix_brain_batteries);
    factory.print_matrix_brain_batteries();
}
