use sha2::{Sha256, Digest};
use std::collections::HashMap;
use chrono::prelude::*; // For engagement time tracking
use serde::{Serialize, Deserialize}; // For JSON serialization/deserialization

// Cookie struct
#[derive(Debug, Clone, Hash, Serialize, Deserialize)]
struct Cookie {
    id: u64,
    data: String,
    engagement_time: u64, // In seconds
}

// Function to hash cookie data
fn hash_cookie(cookie: &Cookie) -> Vec<u8> {
    let mut hasher = Sha256::new();
    hasher.update(format!("{:?}", cookie));
    hasher.finalize().to_vec()4
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

// function for storing data in IPFS storage and returning a hash reference
fn store_data_ipfs(data: &str) -> String {
    format!("ipfs_hash_{}", hex::encode(Sha256::digest(data.as_bytes())))
}

// Function to list cookie data in the marketplace
fn list_cookie_data(cookies: &[Cookie]) -> String {
    serde_json::to_string(&cookies).unwrap()   
}

// Function to purchase and access cookie data from the marketplace
fn purchase_cookie_data(ipfs_hash: &str) -> Option<Vec<Cookie>> {
    //sample retrieval from IPFS using the hash 
    let data = match ipfs_hash {
        "ipfs_hash" => r#"[{"id":1,"data":"Cookie data 1","engagement_time":120},{"id":2,"data":"Cookie data 2","engagement_time":300}]"#,
        _ => return None,
    };
    let cookies: Vec<Cookie> = serde_json::from_str(data).unwrap();
    Some(cookies)
}


// Function to sync cookies across publishers
fn sync_cookies(cookies: &[Cookie], publisher_id: &str) -> HashMap<u64, String> {
    let mut synced_cookies = HashMap::new();
    for cookie in cookies {
        let synced_data = format!("{}:{}", publisher_id, cookie.data);
        synced_cookies.insert(cookie.id, synced_data);
    }
       
    synced_cookies
}

fn main() {
    let mut cookies = vec![
        Cookie { id: 1, data: "Cookie data1".to_string(), engagement_time: 0 },
        Cookie { id: 2, data: "Cookie data2".to_string(), engagement_time: 0 },
        // Add more cookies in a batch
    ];

    // Simulate tracking engagement
    track_engagement(&mut cookies[0], 120); // 120 seconds engagement for cookie 1
    track_engagement(&mut cookies[1], 300); // 300 seconds engagement for cookie 2

    // Build the Merkle Tree and get the root hash
    if let Some(merkle_tree) = build_merkle_tree(&cookies) {
        let root_hash = get_merkle_root(&merkle_tree);
        println!("Merkle Tree Root Hash: {:?}", root_hash);

        // Store the data on IPFS and get the hash reference
        let ipfs_hash = store_data_ipfs(&serde_json::to_string(&cookies).unwrap());
        println!("IPFS Hash: {:?}", ipfs_hash);

        // List cookie data in the marketplace
        let listed_data = list_cookie_data(&cookies);
        println!("Listed Data: {:?}", listed_data);
        // Purchase and access cookie data from the marketplace - ipfs hash based on segments - cookie behavioral segments persona (support.google.com/admanager/answer/4349785?hl=en)
        if let Some(purchased_cookies) = purchase_cookie_data(&ipfs_hash) {
            println!("Purchased Cookies: {:?}", purchased_cookies);
        } else {
            println!("Failed to purchase cookie data.");
        }

        // Calculate rewards
        let reward_rate = 0.01; // reward rate
        let rewards = calculate_rewards(&cookies, reward_rate);
        println!("Rewards: {:?}", rewards);
       

        // Distribute rewards
        for (id, reward) in rewards {
            println!("Distributing {} tokens to user with Cookie ID: {}", reward, id);
        }

        // Sync cookies across publishers
        let publisher_id = "publisher_123";
        let synced_cookies = sync_cookies(&cookies, publisher_id);
        println!("Synced Cookies: {:?}", synced_cookies);
    } else {
        println!("No cookies to build Merkle Tree.");
    }
}
