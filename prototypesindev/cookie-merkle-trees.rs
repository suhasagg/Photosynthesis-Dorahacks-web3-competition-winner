use sha2::{Sha256, Digest};
use std::collections::HashMap;

// Cookie struct
#[derive(Debug, Clone, Hash)]
struct Cookie {
    id: u64,
    data: String,
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

fn main() {
    let cookies = vec![
        Cookie { id: 1, data: "ff6ea16d_b456_41b5_b760_b079194204cf:@seoul@@south_korea@@samsung_galaxy_s20_fe_5g_2021_march@@	automotive and vehicles vehicle brands lamborghini	Seg1525	science medicine physiology	Seg2463	style and fashion clothing blazers	Seg1202	education alumni and reunions@@Seg826	religion and spirituality buddhism  Seg1525	science medicine physiology		automotive and vehicles vehicle brands lamborghini Seg2463 style and fashion clothing blazers".to_string() },
        Cookie { id: 2, data: "80048892_3de8_4266_8c8f_1fe5a8a5814c:@nainital@@india@@samsung_sm_g531f_2014_october@@Seg1410	technology and computing consumer electronics game systems and consoles playstation@@Seg1410	technology and computing consumer electronics game systems and consoles playstation	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0@@technology	india	website1	search	Quarter4  Topic Engagement time scores	0	0	0	0	0	0	0	0	0	0	4	2	2	1	1	2	1	1	4	1	1	1	1	2	1	4	2	1	1	3	2	5	3	2	2	1	5	1	2	2	1	2	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	0	4	6	0	0	0	0	0	0	0	0	0	0	0	0	0	130	14	1	loyal	0@@male@@3@@high@@lenovo legion phone viral video queen elizabeth husband us navy amoled display health and lowfat cooking tutorials friends reunion special".to_string()},
        // Add more cookies in a batch
    ];

    if let Some(merkle_tree) = build_merkle_tree(&cookies) {
        let root_hash = get_merkle_root(&merkle_tree);
        println!("Merkle Tree Root Hash: {:?}", root_hash);
    } else {
        println!("No cookies to build Merkle Tree.");
    }
}
