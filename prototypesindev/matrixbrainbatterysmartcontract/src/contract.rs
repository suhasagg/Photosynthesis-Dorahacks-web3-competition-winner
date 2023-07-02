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
    // Method for retrieving information about a particular kitty
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
        // The tuple that this function returns represents a kitty's properties.
        // Currently, this function just returns a hardcoded tuple with test values.
        // In a real-world implementation, these values would be fetched from a database
        // or some other kind of storage where information about kitties is kept.
        // For example, the first two booleans might represent whether the kitty
        // is for sale or if it is ready for breeding, and the subsequent `u64`
        // values could represent various stats of the kitty.
        // The last value is a timestamp in Unix format representing some event
        // in the kitty's life (e.g., the time it was last fed).
        // The "_kitty_id" parameter is currently unused.
        (
            true, true, 3, 4, 5, 6, 7, 8, 9, 7_688_748_911_342_991
        )
    }// For example, is the kitty for sale?
    // For example, is the kitty ready for breeding?
     // stat of the kitty
}    // stat of the kitty

// This struct represents a 'MatrixBrainBattery'. 
// This could be a part of a simulation game or a complex system in which a certain entity (represented by the name) has various properties.
// It is annotated with several derive attributes, including ones for common traits like serialization, cloning, partial equality, and debugging.
// The JsonSchema attribute is used for validating JSON data against a predefined schema, which can be useful for parsing JSON data into Rust structures.
#[derive(Serialize, Deserialize, Clone, PartialEq, Debug, JsonSchema)]
pub struct MatrixBrainBattery {
    // 'name' property representing the unique identity of the entity. This could be any descriptive name.
    pub name: String,
     // 'dna' could represent a unique identifier or a series of traits encoded as a single `u64` value.
    pub dna: u64,
    // 'level' could represent the current level or rank of the entity in a game or system.
    pub level: u32,
    // 'ready_time' could represent the time at which this entity will be ready for some action.
    // It could be the end of a cooldown period, for example. 
    // This value might be represented as a Unix timestamp or some other form of time data.
    pub ready_time: u64,
     
    // 'win_count' and 'loss_count' could represent the number of victories and defeats this entity has experienced, respectively.
    // These values could be used in ranking or matchmaking algorithms.
    pub win_count: u16,
    pub loss_count: u16,
}

// 'MatrixBrainBatteryFactory' is a struct that represents a factory or collection for creating and managing multiple instances of MatrixBrainBatteries.
pub struct MatrixBrainBatteryFactory {
    // 'kitty_interface' is a field that represents the interface of an external service or functionality called 'KittyInterface'.
    // It's likely used for some functionality related to 'MatrixBrainBattery' instances.
    pub kitty_interface: KittyInterface,
     // 'matrix_brain_batteries' is a Vector used to store instances of MatrixBrainBatteries.
    // This serves as a storage of all the instances of MatrixBrainBattery that are currently in use or have been created.
    pub matrix_brain_batteries: Vec<MatrixBrainBattery>,
      // 'matrix_brain_battery_to_owner' is a HashMap mapping 'MatrixBrainBattery' identifiers to their respective owner's string identifiers.
    // This is typically used for looking up the owner of a particular MatrixBrainBattery.
    pub matrix_brain_battery_to_owner: HashMap<u64, String>,
      // 'matrix_brain_battery_approvals' is a HashMap mapping 'MatrixBrainBattery' identifiers to the approved users' string identifiers.
    // This is typically used in scenarios where a MatrixBrainBattery could be approved for use by certain users.
    pub matrix_brain_battery_approvals: HashMap<u64, String>,
      // 'owner_matrix_brain_battery_count' is a HashMap mapping owner identifiers to the count of 'MatrixBrainBattery' they own.
    // This can be used to quickly get the number of MatrixBrainBattery instances an owner has.
    pub owner_matrix_brain_battery_count: HashMap<String, u64>,
      
    // 'level_up_fee' is a field representing the cost for leveling up a MatrixBrainBattery. 
    // This is represented as an Uint128 to accommodate very large numbers.
    pub level_up_fee:  Uint128,
       // 'rand_nonce' could be used for generating random numbers or for some cryptographic operation.
    pub rand_nonce: u32,
}

// The 'Default' trait provides a function 'default' to create a default value for types.
// Here we implement the 'Default' trait for our 'MatrixBrainBattery' struct.
impl Default for MatrixBrainBattery {
    // The 'default' function is used to generate a default value for a type.
    fn default() -> Self {
        Self {
            name: "".to_string(),
                // We set the default dna to be 0. 
            // Note: this is just a placeholder and should be replaced with actual default logic as required.
            dna: 0, // Or any default value
              // The default level is set to 0. 
            // Note: this could be adjusted based on the actual requirements of the game mechanics.
            level: 0, // Or any default value
               // The default ready_time is set to 0. 
            // This could represent that the battery is ready for use immediately after being created.
            ready_time: 0, // Or any default value
              // The default win_count is set to 0, as a new battery has not yet participated in any battles.
            win_count: 0, // Or any default value
               // Similarly, the default loss_count is also 0, as a new battery has not yet lost any battles.
            loss_count: 0, // Or any default value
        }
    }
}
// Implementing methods for MatrixBrainBatteryFactory struct
impl MatrixBrainBatteryFactory {
   // This method is used to create a new instance of MatrixBrainBatteryFactory with initial values
    pub fn new() -> Self {
        MatrixBrainBatteryFactory {
            kitty_interface: KittyInterface,
             
            // Initialize an empty vector to store matrix_brain_batteries. 
            // This will hold all the MatrixBrainBatteries created in the game.
            matrix_brain_batteries: Vec::new(),
            // Initialize an empty HashMap to store the association between MatrixBrainBatteries and their owners.
            // The keys are the ids of the MatrixBrainBatteries and the values are the addresses of the owners.
            matrix_brain_battery_to_owner: HashMap::new(),
             // Initialize an empty HashMap to store the approvals of the MatrixBrainBatteries.
            // An approval allows someone else to transfer ownership of a MatrixBrainBattery. 
            // The keys are the ids of the MatrixBrainBatteries and the values are the addresses of the approved accounts.
            matrix_brain_battery_approvals:  HashMap::new(),
            
            // Initialize an empty HashMap to keep track of how many MatrixBrainBatteries each address owns.
            // The keys are the addresses of the owners and the values are the number of MatrixBrainBatteries they own.
            owner_matrix_brain_battery_count: HashMap::new(),
            
            // Initialize the level_up_fee with a value of zero.
            // This fee is required to level up a MatrixBrainBattery.
            level_up_fee: Uint128::zero(),
             // Initialize the rand_nonce (used for randomness in the game) to zero.
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
            // Get the count of MatrixBrainBatteries owned by the specified owner, or 0 if the owner does not exist in the HashMap
        let count = self.owner_matrix_brain_battery_count.get(&owner).cloned().unwrap_or(0);
         // If the owner already owns a MatrixBrainBattery, return an error
        if count > 0 {
            return Err(StdError::generic_err("Owner already has a MatrixBrainBattery"));
        }
 // Generate a random dna for the new MatrixBrainBattery using the provided name
        let rand_dna = self.generate_random_dna(&name);
     // Create a new MatrixBrainBattery with the specified name, generated dna, initial level 1, 
    // set ready time to current block time plus COOLDOWN_TIME seconds, and initial win and loss counts to 0.
        let new_battery = MatrixBrainBattery {
            name,
            dna: rand_dna,
            level: 1,
            ready_time: env.block.time.plus_seconds(COOLDOWN_TIME).seconds(),
            win_count: 0,
            loss_count: 0,
        };
        let new_battery_clone = new_battery.clone();
         // Add the newly created MatrixBrainBattery to the vector of all MatrixBrainBatteries
        //self.matrix_brain_batteries.push(new_battery);
        self.matrix_brain_batteries.push(new_battery_clone);
       
    // Get the id of the newly created MatrixBrainBattery, which is the index of the last element in the vector
        let id = self.matrix_brain_batteries.len() - 1;
       // Associate the new MatrixBrainBattery with its owner in the HashMap
        self.matrix_brain_battery_to_owner.insert(id as u64, owner.clone());
       // Increment the count of MatrixBrainBatteries owned by the specified owner in the HashMap
        self.owner_matrix_brain_battery_count.insert(owner, 1);
        Ok(new_battery)
    }
     
    pub fn trigger_cooldown(&mut self, index: usize, env: &Env) {
    let matrix_brain_battery = self.matrix_brain_batteries
        .get_mut(index)
        .expect("MatrixBrainBattery not found"); 
     
    // Update the ready_time of the MatrixBrainBattery to be the current block time plus COOLDOWN_TIME seconds.
    // This effectively puts the MatrixBrainBattery in a "cooldown" state, 
    // where it will not be ready for another action until after the cooldown period has passed.
     matrix_brain_battery.ready_time = env.block.time.plus_seconds(COOLDOWN_TIME).seconds();
    }
   
   pub fn is_ready(matrix_brain_battery: &MatrixBrainBattery, env: &Env) -> bool {
    // The function takes a reference to a MatrixBrainBattery instance and an Env instance.
    // It checks if the current time (in seconds) is greater than or equal to the MatrixBrainBattery's ready_time.
    // If the current time is greater than the ready_time, it means the MatrixBrainBattery has finished its cooldown and is "ready". Therefore, it returns true.
    // If the current time is less than the ready_time, it means the MatrixBrainBattery is still on cooldown and is not "ready". Therefore, it returns false.
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
        // The function takes an existing MatrixBrainBattery id, target DNA, species type, 
    // owner and the environment details.

    // Accesses the matrix_brain_battery from the collection of batteries 
    // using provided matrix_brain_battery_id. If the id is invalid, return a corresponding error.
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;
   // Check if the MatrixBrainBattery is ready. If not, return a corresponding error.
        if !MatrixBrainBatteryFactory::is_ready(&matrix_brain_battery, &env) {
            return Err(StdError::generic_err("MatrixBrainBattery not ready"));
        }
 // Create new DNA by averaging the current MatrixBrainBattery DNA and the target DNA,
    // with the target DNA modulo DNA_MODULUS to ensure it stays within expected bounds.
        let new_dna = (matrix_brain_battery.dna + target_dna % DNA_MODULUS) / 2;
    // If the species of the MatrixBrainBattery is "kitty", modify the DNA slightly.
        let new_dna = if species == "kitty" {
            new_dna - new_dna % 100 + 99
        } else {
            new_dna
        };
     // Create a new MatrixBrainBattery with the new DNA, initial level 1, and set its ready_time to now + COOLDOWN_TIME.   
     let new_battery = MatrixBrainBattery {
        name: "BrainAxie".to_string(),
        dna: new_dna,
        level: 1,
        ready_time: env.block.time.plus_seconds(COOLDOWN_TIME).seconds(),
        win_count: 0,
        loss_count: 0,
    };
    
 // Push the new MatrixBrainBattery to the batteries vector.
    self.matrix_brain_batteries.push(new_battery);
   // Compute the id for the newly created MatrixBrainBattery as the last index in the vector.
    let id = self.matrix_brain_batteries.len() - 1;
     // Insert a new entry in the matrix_brain_battery_to_owner hashmap where 
    // key is the id of the newly created MatrixBrainBattery and value is the owner string.
    self.matrix_brain_battery_to_owner.insert(id as u64, owner.to_string());
     // Insert or update the owner_matrix_brain_battery_count hashmap.
    // If the owner already exists in the hashmap, increment the count by 1.
    // If not, insert the owner with a count of 1.
    self.owner_matrix_brain_battery_count
        .entry(owner.to_string())
        .and_modify(|count| *count += 1)
        .or_insert(1);
   // Trigger the cooldown period for the original MatrixBrainBattery.
    // This essentially updates its ready_time to current block time plus the COOLDOWN_TIME.
    self.trigger_cooldown(matrix_brain_battery_id as usize, &env);
    Ok(())
       

    }
// The `feed_on_kitty` method is a specialized version of `feed_and_multiply`.
// It uses a provided `kitty_id` to query the DNA of a CryptoKitty from the `kitty_interface`,
// then it feeds this DNA to a given `MatrixBrainBattery`.
    pub fn feed_on_kitty(
        &mut self,  // Mutable reference to the current `MatrixBrainBatteryFactory` instance.
        matrix_brain_battery_id: u64, // The id of the MatrixBrainBattery that is going to "feed" on a kitty.
        kitty_id: u64,  // The id of the kitty from which we will get the DNA.
        owner: &str, // The owner of the MatrixBrainBattery.
        env: Env // A snapshot of the current blockchain environment.
    ) -> StdResult<()> {
          // Fetch the data of the kitty, including its DNA, from the `kitty_interface`.
        let (_, _, _, _, _, _, _, _, _, kitty_dna) = self.kitty_interface.get_kitty(kitty_id);
         // Use the kitty's DNA to feed and multiply the MatrixBrainBattery, indicating that the species of the target DNA is 'kitty'.
        self.feed_and_multiply(matrix_brain_battery_id, kitty_dna, "kitty", owner, env)
    }
    // This function is for debugging purposes. It prints out all the MatrixBrainBatteries in the current factory instance.
    pub fn print_matrix_brain_batteries(&self) {  // Immutable reference to the current `MatrixBrainBatteryFactory` instance.
        println!("{:#?}", self.matrix_brain_batteries);
    }
}

use cw721::{Cw721, Expiration, Approval};

// MatrixBrainBatteryHelper
impl MatrixBrainBatteryFactory {
  // LEVEL_UP_FEE is a constant value that indicates the cost for a MatrixBrainBattery to level up.
    pub const LEVEL_UP_FEE: u128 = 1_000_000_000; // 0.001 SCRT
    
    // This function sets the level up fee for the MatrixBrainBatteryFactory.
    pub fn set_level_up_fee(&mut self, fee: Uint128) { // fee: the new level up fee.
        self.level_up_fee = fee;  // Assign the new fee to the current instance's level_up_fee.
    }
      // This function increases the level of a specified MatrixBrainBattery.
    pub fn level_up(&mut self, matrix_brain_battery_id: u64) -> StdResult<()> { // matrix_brain_battery_id: the id of the MatrixBrainBattery to be levelled up.
         // Attempt to get a mutable reference to the MatrixBrainBattery to be levelled up.
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)   // Convert the id to an index.
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?; // If the MatrixBrainBattery does not exist, return an error.
        matrix_brain_battery.level += 1; // If it exists, increment its level.
        Ok(())
    }
  // This function changes the name of a specified MatrixBrainBattery.
    pub fn change_name(&mut self, matrix_brain_battery_id: u64, new_name: &str) -> StdResult<()> {// matrix_brain_battery_id: the id of the MatrixBrainBattery to be renamed, new_name: the new name for the MatrixBrainBattery.
         // Attempt to get a mutable reference to the MatrixBrainBattery to be renamed.
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize) // Convert the id to an index.
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?; // If the MatrixBrainBattery does not exist, return an error.

        matrix_brain_battery.name = new_name.to_string(); // If it exists, change its name to the provided new name.
        Ok(())
    }
 // This function changes the DNA of a specified MatrixBrainBattery.
    pub fn change_dna(&mut self, matrix_brain_battery_id: u64, new_dna: u64) -> StdResult<()> {// matrix_brain_battery_id: the id of the MatrixBrainBattery whose DNA will be changed, new_dna: the new DNA for the MatrixBrainBattery.
       // Attempt to get a mutable reference to the MatrixBrainBattery whose DNA will be changed.
        let mut matrix_brain_battery = self.matrix_brain_batteries
            .get_mut(matrix_brain_battery_id as usize)// Convert the id to an index.
            .ok_or_else(|| StdError::generic_err("Invalid MatrixBrainBattery ID"))?;  // If the MatrixBrainBattery does not exist, return an error.
        matrix_brain_battery.dna = new_dna; // If it exists, change its DNA to the provided new DNA.
        Ok(())
    }
  // This function returns all the MatrixBrainBattery IDs that are owned by a specific owner.
    pub fn get_matrix_brain_batteries_by_owner(&self, owner: &str) -> Vec<u64> {// owner: The name of the owner to look up.
        let mut result = Vec::new();// Create an empty vector to store the IDs.
            // Iterate over the matrix_brain_batteries list, with the index 'i' and the battery 'matrix_brain_battery'.
        for (i, matrix_brain_battery) in self.matrix_brain_batteries.iter().enumerate() {
          // Try to get the owner of the current MatrixBrainBattery by checking in the matrix_brain_battery_to_owner hashmap.
            if let Some(matrix_brain_battery_owner) = self.matrix_brain_battery_to_owner.get(&(i as u64)) {
               // If the current MatrixBrainBattery is owned by the owner we're looking for, 
                // push the index (which serves as the ID in this context) into the result vector.
                if matrix_brain_battery_owner == owner {
                    result.push(i as u64);
                }
            }
        }
        // After iterating over all MatrixBrainBatteries, return the result vector which contains all the IDs owned by the specified owner.
        result
    }
}

// MatrixBrainBatteryAttack
impl MatrixBrainBatteryFactory {
    pub const ATTACK_VICTORY_PROBABILITY: u64 = 70;
 // Constant defining the likelihood of winning an attack, as a percentage out of 100.
    // In this case, a MatrixBrainBattery has a 70% chance of winning an attack.
    
    // rand_mod is a helper function to generate a pseudo-random number within a certain range (modulo).
    // This is used to generate unpredictable results in certain game operations.
    // This method is using Sha256 hash of the combination of current block time and an incrementing nonce to generate a random number.
    
    
    fn rand_mod(&mut self, modulo: u64, env: &Env) -> u64 {
             // Increment the nonce by 1 each time this function is called. 
        // This ensures that the seed for the hashing function is different every time.
        self.rand_nonce += 1;
         // Create a seed string by concatenating the current block time in seconds and the nonce.
        let seed = format!("{}{}", env.block.time.seconds(), self.rand_nonce);
        //let hash = u64::from_be_bytes(Sha256::digest(seed.as_bytes()));
        let hash = Sha256::digest(seed.as_bytes());
        // Only keep the first 8 bytes of the hash and convert them to a u64 number.
        let mut array = [0u8; 8];
        array.copy_from_slice(&hash[0..8]);
        let num = u64::from_be_bytes(array);
        // Return the u64 number modulo the input value. This creates a pseudo-random number in the range 0 to modulo.
        num
    }
   // The attack function allows one MatrixBrainBattery to attack another. 
    // The outcome of the attack is determined randomly, but the attacker has a higher chance of winning.
    // If the attacker wins, it increases its level, win count and also multiplies, creating a new MatrixBrainBattery with a mixed DNA.
    // If the attacker loses, its loss count is increased and it goes into a cooldown period.
    pub fn attack(&mut self, matrix_brain_battery_id: u64, my_matrix_brain_battery: &mut MatrixBrainBattery, enemy_matrix_brain_battery: &mut MatrixBrainBattery, owner: &str, env: Env) -> StdResult<()> {     

        // Finding the index of my_matrix_brain_battery and enemy_matrix_brain_battery in the matrix_brain_batteries vector
        let my_index = self.matrix_brain_batteries.iter()
            .position(|battery| battery.name == my_matrix_brain_battery.name)
            .expect("My battery not found");
    
        let enemy_index = self.matrix_brain_batteries.iter()
            .position(|battery| battery.name == enemy_matrix_brain_battery.name)
            .expect("Enemy battery not found");
    
        // Generating a random number between 0 to 100.
        // If the number is less than or equal to ATTACK_VICTORY_PROBABILITY (70), my_matrix_brain_battery wins the attack.
        // Otherwise, enemy_matrix_brain_battery wins.
        let rand = self.rand_mod(100, &env);  
        // Generating a random number between 0 to 100.
        // If the number is less than or equal to ATTACK_VICTORY_PROBABILITY (70), my_matrix_brain_battery wins the attack.
        // Otherwise, enemy_matrix_brain_battery wins.
        if rand <= Self::ATTACK_VICTORY_PROBABILITY {
          // Update win count and level of my_matrix_brain_battery, and loss count of enemy_matrix_brain_battery
            my_matrix_brain_battery.win_count += 1;
            my_matrix_brain_battery.level += 1;
            enemy_matrix_brain_battery.loss_count += 1;
            // Call feed_and_multiply to create a new MatrixBrainBattery with DNA from both my_matrix_brain_battery and enemy_matrix_brain_battery
            self.feed_and_multiply(matrix_brain_battery_id, enemy_matrix_brain_battery.dna, "matrixbrainbattery", owner, env)?;
        } else {
              // Update loss count of my_matrix_brain_battery and win count of enemy_matrix_brain_battery
            my_matrix_brain_battery.loss_count += 1;
            enemy_matrix_brain_battery.win_count += 1;
            // Trigger cooldown on my_matrix_brain_battery
            let target_name = "BrainAxie".to_string();
            let my_battery_index = self.matrix_brain_batteries.iter()
                .position(|battery| battery.name == target_name)
                .expect("Battery not found");
    
            self.trigger_cooldown(my_battery_index, &env);
        }
     // Update the my_matrix_brain_battery and enemy_matrix_brain_battery in the matrix_brain_batteries vector
        self.matrix_brain_batteries[my_index] = my_matrix_brain_battery.clone();
        self.matrix_brain_batteries[enemy_index] = enemy_matrix_brain_battery.clone();
    
        Ok(())
    }
    
}   
//This attack function defines how a MatrixBrainBattery can attack another one. This process is influenced by randomness - each attack has a set chance of being successful based on the constant ATTACK_VICTORY_PROBABILITY. If the attack is successful, the attacking MatrixBrainBattery increases its win count and level and produces a new MatrixBrainBattery. If the attack fails, the attacking MatrixBrainBattery's loss count increases, and it enters a cooldown period where it cannot attack again for some time


// MatrixBrainBatteryOwnership
impl MatrixBrainBatteryFactory {

    // This function transfers the ownership of a MatrixBrainBattery from one user to another
    pub fn transfer_ownership(&mut self, from: &str, to: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
         // Check if the `from` user is the current owner of the MatrixBrainBattery
        if self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id) != Some(&from.to_string()) {
            return Err(StdError::generic_err("Only owner can transfer ownership"));
        }
 // If `from` user is the owner, change the owner to `to` user
        self.matrix_brain_battery_to_owner.insert(matrix_brain_battery_id, to.to_string());
        Ok(())
    }
  // This function lets the current owner of a MatrixBrainBattery approve another user to take the ownership
    pub fn approve(&mut self, approver: &str, to: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
     // Check if the `approver` is the current owner of the MatrixBrainBattery
        if self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id) != Some(&approver.to_string()) {
            return Err(StdError::generic_err("Only owner can approve"));
        }
 // If `approver` is the owner, insert the approval for `to` user in `matrix_brain_battery_approvals` map
        self.matrix_brain_battery_approvals.insert(matrix_brain_battery_id, to.to_string());
        Ok(())
    }

   
     // This function lets a user who has been approved to take the ownership of a MatrixBrainBattery
    pub fn take_ownership(&mut self, new_owner: &str, matrix_brain_battery_id: u64) -> StdResult<()> {
        let approval = self.matrix_brain_battery_approvals.get(&matrix_brain_battery_id);
        // Check if the `new_owner` user has been approved to take the ownership
        if approval != Some(&new_owner.to_string()) {
            return Err(StdError::generic_err("Only approved address can take ownership"));
        }
      // If `new_owner` has been approved, find the current owner
        let owner = self.matrix_brain_battery_to_owner.get(&matrix_brain_battery_id)
        .ok_or_else(|| {
            StdError::generic_err("MatrixBrainBattery not found")
        })?.clone();  // note the clone() here
    // Call `transfer_ownership` function to transfer the ownership from current owner to `new_owner`
        self.transfer_ownership(&owner, new_owner, matrix_brain_battery_id)
    }  
}

fn main() {
    // Initialize the factory  
    // First, a new MatrixBrainBatteryFactory is instantiated
    let mut factory = MatrixBrainBatteryFactory::new();
   // SystemTime and Timestamp are used to generate unique IDs for MatrixBrainBatteries
    let now = SystemTime::now();
    // Assume the following initial parameters
    // Assumed initial parameters for the first battery
    let owner = "Owner1".to_string();
    let name = "Battery1".to_string();
    // This timestamp is used as a unique identifier for the battery
    let timestamp = now.duration_since(UNIX_EPOCH)
    .expect("System time is before the UNIX epoch")
    .as_secs();
// Assumed initial parameters for the second battery
    let owner2 = "Owner2".to_string();
    let name2 = "Battery2".to_string();
    let timestamp = now.duration_since(UNIX_EPOCH)
    .expect("System time is before the UNIX epoch")
    .as_secs(); 
    // env provides environment information for the contract execution
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
// Create the first battery and handle the result
    let mut matrix_brain_battery1 = MatrixBrainBattery::default();

    let result1 = factory.create_random_matrix_brain_battery(name, owner.clone(), env.clone());
    
     if let Ok(battery) = result1 {
         // If successful, copy the battery to your placeholder
       // Copy the battery to your placeholder
         matrix_brain_battery1 = battery;
     } else {
    // Handle error if result1 is Err
    println!("Error creating battery");
    }
    // Print all current batteries
    factory.print_matrix_brain_batteries();
// Create the second battery and handle the result
    let mut matrix_brain_battery2 = MatrixBrainBattery::default();
// Feed the first battery on a Kitty and multiply it
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
    // Print all current batteries
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
