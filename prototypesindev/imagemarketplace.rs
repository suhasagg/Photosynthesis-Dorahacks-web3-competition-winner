use sha2::{Sha256, Digest};
use std::collections::HashMap;
use chrono::prelude::*; // For engagement time tracking
use serde::{Serialize, Deserialize}; // For JSON serialization/deserialization
use hex;

// Image struct
#[derive(Debug, Clone, Hash, Serialize, Deserialize)]
struct Image {
    id: u64,
    metadata: String,
    engagement_time: u64, // In seconds
}

// Function to hash image metadata
fn hash_image(image: &Image) -> Vec<u8> {
    let mut hasher = Sha256::new();
    hasher.update(format!("{:?}", image));
    hasher.finalize().to_vec()
}

// Merkle Tree node
#[derive(Clone, Debug)]
enum Node {
    Leaf(Vec<u8>),
    Internal(Vec<u8>, Box<Node>, Box<Node>),
}

// Function to generate Merkle Tree from images
fn build_merkle_tree(images: &[Image]) -> Option<Node> {
    match images.len() {
        0 => None,
        1 => Some(Node::Leaf(hash_image(&images[0]))),
        _ => {
            let mid = images.len() / 2;
            let left = build_merkle_tree(&images[..mid])?;
            let right = build_merkle_tree(&images[mid..])?;
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
fn track_engagement(image: &mut Image, additional_time: u64) {
    image.engagement_time += additional_time;
}

// Function to calculate rewards based on engagement time
fn calculate_rewards(images: &[Image], reward_rate: f64) -> HashMap<u64, f64> {
    let mut rewards = HashMap::new();
    for image in images {
        let reward = (image.engagement_time as f64) * reward_rate;
        rewards.insert(image.id, reward);
    }
    rewards
}

// Function for storing data in IPFS storage and returning a hash reference
fn store_data_ipfs(data: &str) -> String {
    format!("ipfs_hash_{}", hex::encode(Sha256::digest(data.as_bytes())))
}

// Function to list image data in the marketplace
fn list_image_data(images: &[Image]) -> String {
    serde_json::to_string(&images).unwrap()   
}

// Function to purchase and access image data from the marketplace
fn purchase_image_data(ipfs_hash: &str) -> Option<Vec<Image>> {
    // Sample retrieval from IPFS using the hash 
    let data = match ipfs_hash {
        "ipfs_hash" => r#"[{"id":1,"metadata":"Image metadata 1","engagement_time":120},{"id":2,"metadata":"Image metadata 2","engagement_time":300}]"#,
        _ => return None,
    };
    let images: Vec<Image> = serde_json::from_str(data).unwrap();
    Some(images)
}

// Function to sync images across publishers
fn sync_images(images: &[Image], publisher_id: &str) -> HashMap<u64, String> {
    let mut synced_images = HashMap::new();
    for image in images {
        let synced_data = format!("{}:{}", publisher_id, image.metadata);
        synced_images.insert(image.id, synced_data);
    }
    synced_images
}

fn main() {
    let mut images = vec![
        Image { id: 1, metadata: "Image metadata 1".to_string(), engagement_time: 0 },
        Image { id: 2, metadata: "Image metadata 2".to_string(), engagement_time: 0 },
        // Add more images in a batch
    ];

    // Simulate tracking engagement
    track_engagement(&mut images[0], 120); // 120 seconds engagement for image 1
    track_engagement(&mut images[1], 300); // 300 seconds engagement for image 2

    // Build the Merkle Tree and get the root hash
    if let Some(merkle_tree) = build_merkle_tree(&images) {
        let root_hash = get_merkle_root(&merkle_tree);
        println!("Merkle Tree Root Hash: {:?}", root_hash);

        // Store the data on IPFS and get the hash reference
        let ipfs_hash = store_data_ipfs(&serde_json::to_string(&images).unwrap());
        println!("IPFS Hash: {:?}", ipfs_hash);

        // List image data in the marketplace
        let listed_data = list_image_data(&images);
        println!("Listed Data: {:?}", listed_data);

        // Purchase and access image data from the marketplace
        if let Some(purchased_images) = purchase_image_data(&ipfs_hash) {
            println!("Purchased Images: {:?}", purchased_images);
        } else {
            println!("Failed to purchase image data.");
        }

        // Calculate rewards
        let reward_rate = 0.01; // reward rate
        let rewards = calculate_rewards(&images, reward_rate);
        println!("Rewards: {:?}", rewards);

        // Distribute rewards
        for (id, reward) in rewards {
            println!("Distributing {} tokens to user with Image ID: {}", reward, id);
        }

        // Sync images across publishers
        let publisher_id = "publisher_123";
        let synced_images = sync_images(&images, publisher_id);
        println!("Synced Images: {:?}", synced_images);
    } else {
        println!("No images to build Merkle Tree.");
    }
}
