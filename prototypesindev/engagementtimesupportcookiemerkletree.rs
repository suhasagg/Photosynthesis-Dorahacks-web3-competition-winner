use sha2::{Sha256, Digest};
use std::collections::HashMap;
use chrono::prelude::*; //For engagement time tracking
use serde::{Serialize, Deserialize}; //For JSON serialization/deserialization

// Cookie struct
#[derive(Debug, Clone, Hash, Serialize, Deserialize)]
struct Cookie {
    id: u64,
    data: String,
    engagement_time: u64, //In seconds
}

// Function to hash cookie data
fn hash_cookie(cookie: &Cookie) -> Vec<u8> {
    let mut hasher = Sha256::new();
    hasher.update(format!("{:?}", cookie));
    hasher.finalize().to_vec()
}

// Merkle Tree node
#[derive(Clone, Debug)]
enum Node {
    Leaf(Vec<u8>),
    Internal(Vec<u8>, Box<Node>, Box<Node>),
}

// Function to generate Merkle Tree from cookies
fn build_merkle_tree(cookies: &[Cookie]) -> Option<Node> {
    match cookies.len() {
        0 => None,
        1 => Some(Node::Leaf(hash_cookie(&cookies[0]))),
        _ => {
            let mid = cookies.len() / 2;
            let left = build_merkle_tree(&cookies[..mid])?;
            let right = build_merkle_tree(&cookies[mid..])?;
            let mut hasher = Sha256::new();
            if let Node::Leaf(ref left_hash) = left {
                hasher.update(left_hash);
            }
            if let Node::Leaf(ref right_hash) = right {
                hasher.update(right_hash);
            }
            Some(Node::Internal(hasher.finalize().to_vec(), Box::new(left), Box::new(right)))
        }
    }
}

// Function to retrieve the root hash of the Merkle Tree
fn get_merkle_root(node: &Node) -> Vec<u8> {
    match node {
        Node::Leaf(hash) | Node::Internal(hash, _, _) => hash.clone(),
    }
}


// Function to track engagement time 
fn track_engagement(cookie: &mut Cookie, additional_time: u64) {
    cookie.engagement_time += additional_time;
}

// Function to calculate rewards based on engagement time
fn calculate_rewards(cookies: &[Cookie], reward_rate: f64) -> HashMap<u64, f64> {
    let mut rewards = HashMap::new();
    for cookie in cookies {
        let reward = (cookie.engagement_time as f64) * reward_rate;
        rewards.insert(cookie.id, reward);
    }
    rewards
}


fn store_data_ipfs(data: &str) -> String {
    format!("ipfs_hash_{}", Sha256::digest(data.as_bytes()))
}

fn main() {
    let mut cookies = vec![
        Cookie { id: 1, data: "Cookie data1".to_string(), engagement_time: 0 },
        Cookie { id: 2, data: "Cookie data2".to_string(), engagement_time: 0 },
        // Add more cookies in a batch
    ];
    //tracking engagement
    track_engagement(&mut cookies[0], 120); // 120 seconds engagement for cookie 1
    track_engagement(&mut cookies[1], 300); // 300 seconds engagement for cookie 2

    // Build the Merkle Tree and get the root hash
    if let Some(merkle_tree) = build_merkle_tree(&cookies) {
        let root_hash = get_merkle_root(&merkle_tree);
        println!("Merkle Tree Root Hash: {:?}", root_hash);

        // Store the data on IPFS and get the hash reference
        let ipfs_hash = store_data_ipfs(&serde_json::to_string(&cookies).unwrap());
        println!("IPFS Hash: {:?}", ipfs_hash);

        // Calculate rewards
        let reward_rate = 0.01; //reward rate
        let rewards = calculate_rewards(&cookies, reward_rate);
        println!("Rewards: {:?}", rewards);

        // Distribute rewards
        for (id, reward) in rewards {
            println!("Distributing {} tokens to user with Cookie ID: {}", reward, id);
        }
    } else {
        println!("No cookies to build Merkle Tree.");
    }
}
