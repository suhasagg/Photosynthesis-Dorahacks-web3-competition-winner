# Adserver smart contract

Codacy grade A -
https://app.codacy.com/gh/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/file/93328436253/issues/source?bid=34568294&fileBranchId=34568294

# Heavy Transaction Guaranteed Dapps on Daily Basis – A strong candidate for Archway Rewards and Photosynthesis liquid token

https://twitter.com/coolcaptchas/status/1652382322153652225?s=20

https://twitter.com/coolcaptchas/status/1652391219086131200?s=20

https://twitter.com/coolcaptchas/status/1527443243327434752?s=20

https://twitter.com/coolcaptchas/status/1506448375474180098?s=20

https://docs.google.com/document/d/1-pGzhqz7UHgB2O9qD4fcLAhwhTUgUI4A/edit?usp=sharing&ouid=102246369981228451498&rtpof=true&sd=true

Layer 2 chain derivative -
https://github.com/suhasagg/BlockchainBridgeSystemDesign

```
Saving state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
State instantiated successfully.
Loaded state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id",
            image_url: "test_image_url",
            target_url: "test_target_url",
            views: 0,
            reward_address: "test_reward_address",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "add_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "add_ad",
        },
        Attribute {
            key: "reward_address",
            value: "test_reward_address",
        },
        Attribute {
            key: "id",
            value: "test_id",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url",
        },
    ],
}
Event created: Event {
    ty: "reward",
    attributes: [
        Attribute {
            key: "action",
            value: "reward",
        },
        Attribute {
            key: "recipient",
            value: "0x123abc...",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Event created: Event {
    ty: "staking",
    attributes: [
        Attribute {
            key: "action",
            value: "staking",
        },
        Attribute {
            key: "contract_address",
            value: "0x123abc...",
        },
        Attribute {
            key: "staker_address",
            value: "cosmos2contract",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Ad added successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id",
            image_url: "test_image_url",
            target_url: "test_target_url",
            views: 0,
            reward_address: "test_reward_address",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: Ad {
    id: "test_id",
    image_url: "test_image_url",
    target_url: "test_target_url",
    views: 1,
    reward_address: "test_reward_address",
}
Event created: Event {
    ty: "serve_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "serve_ad",
        },
        Attribute {
            key: "id",
            value: "test_id",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url",
        },
    ],
}
Ad served successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id",
            image_url: "test_image_url",
            target_url: "test_target_url",
            views: 0,
            reward_address: "test_reward_address",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "delete_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "delete_ad",
        },
        Attribute {
            key: "id",
            value: "test_id",
        },
    ],
}
Ad deleted successfully.

Batch Ad Serving Demo for gas and Archway rewards optimisation

Saving state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
State instantiated successfully.
Loaded state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "add_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "add_ad",
        },
        Attribute {
            key: "reward_address",
            value: "test_reward_address1",
        },
        Attribute {
            key: "id",
            value: "test_id1",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url1",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url1",
        },
    ],
}
Event created: Event {
    ty: "reward",
    attributes: [
        Attribute {
            key: "action",
            value: "reward",
        },
        Attribute {
            key: "recipient",
            value: "0x123abc...",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Event created: Event {
    ty: "staking",
    attributes: [
        Attribute {
            key: "action",
            value: "staking",
        },
        Attribute {
            key: "contract_address",
            value: "0x123abc...",
        },
        Attribute {
            key: "staker_address",
            value: "cosmos2contract",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Ad test_id1 added successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "add_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "add_ad",
        },
        Attribute {
            key: "reward_address",
            value: "test_reward_address2",
        },
        Attribute {
            key: "id",
            value: "test_id2",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url2",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url2",
        },
    ],
}
Event created: Event {
    ty: "reward",
    attributes: [
        Attribute {
            key: "action",
            value: "reward",
        },
        Attribute {
            key: "recipient",
            value: "0x123abc...",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Event created: Event {
    ty: "staking",
    attributes: [
        Attribute {
            key: "action",
            value: "staking",
        },
        Attribute {
            key: "contract_address",
            value: "0x123abc...",
        },
        Attribute {
            key: "staker_address",
            value: "cosmos2contract",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Ad test_id2 added successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "add_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "add_ad",
        },
        Attribute {
            key: "reward_address",
            value: "test_reward_address3",
        },
        Attribute {
            key: "id",
            value: "test_id3",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url3",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url3",
        },
    ],
}
Event created: Event {
    ty: "reward",
    attributes: [
        Attribute {
            key: "action",
            value: "reward",
        },
        Attribute {
            key: "recipient",
            value: "0x123abc...",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Event created: Event {
    ty: "staking",
    attributes: [
        Attribute {
            key: "action",
            value: "staking",
        },
        Attribute {
            key: "contract_address",
            value: "0x123abc...",
        },
        Attribute {
            key: "staker_address",
            value: "cosmos2contract",
        },
        Attribute {
            key: "amount",
            value: "20",
        },
    ],
}
Ad test_id3 added successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: Ad {
    id: "test_id1",
    image_url: "test_image_url1",
    target_url: "test_target_url1",
    views: 1,
    reward_address: "test_reward_address1",
}
Event created: Event {
    ty: "serve_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "serve_ad",
        },
        Attribute {
            key: "id",
            value: "test_id1",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url1",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url1",
        },
    ],
}
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: Ad {
    id: "test_id2",
    image_url: "test_image_url2",
    target_url: "test_target_url2",
    views: 1,
    reward_address: "test_reward_address2",
}
Event created: Event {
    ty: "serve_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "serve_ad",
        },
        Attribute {
            key: "id",
            value: "test_id2",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url2",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url2",
        },
    ],
}
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: Ad {
    id: "test_id3",
    image_url: "test_image_url3",
    target_url: "test_target_url3",
    views: 1,
    reward_address: "test_reward_address3",
}
Event created: Event {
    ty: "serve_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "serve_ad",
        },
        Attribute {
            key: "id",
            value: "test_id3",
        },
        Attribute {
            key: "image_url",
            value: "test_image_url3",
        },
        Attribute {
            key: "target_url",
            value: "test_target_url3",
        },
    ],
}
Ads served successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id1",
            image_url: "test_image_url1",
            target_url: "test_target_url1",
            views: 0,
            reward_address: "test_reward_address1",
        },
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "delete_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "delete_ad",
        },
        Attribute {
            key: "id",
            value: "test_id1",
        },
    ],
}
Ad test_id1 deleted successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id2",
            image_url: "test_image_url2",
            target_url: "test_target_url2",
            views: 0,
            reward_address: "test_reward_address2",
        },
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "delete_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "delete_ad",
        },
        Attribute {
            key: "id",
            value: "test_id2",
        },
    ],
}
Ad test_id2 deleted successfully.
Loaded state: State {
    ads: [
        Ad {
            id: "test_id3",
            image_url: "test_image_url3",
            target_url: "test_target_url3",
            views: 0,
            reward_address: "test_reward_address3",
        },
    ],
    total_views: 0,
    plt_address: "",
}
Saving state: State {
    ads: [],
    total_views: 0,
    plt_address: "",
}
Event created: Event {
    ty: "delete_ad",
    attributes: [
        Attribute {
            key: "action",
            value: "delete_ad",
        },
        Attribute {
            key: "id",
            value: "test_id3",
        },
    ],
}
Ad test_id3 deleted successfully.


Protocol Scaling Algorithm
                
+------------------+      +------------------+      +------------------+
| Website          |      | Data Processing  |      | Merkle Tree      |
| - Collect Ad     | ---> | and Aggregation  | ---> | Generation       |
|   Clicks/        |      | - Data Formatting|      | - 5-min/hourly/  |
|   Impressions    |      | - Data Encryption|      |   daily trees    |
| - Ad Data as     |      | - Stream to IPFS |      | - Ad data as     |
|   Input          |      |                  |      |   leaves         |
+------------------+      +------------------+      +------------------+
                                                               |
                                                               V
                                                      +------------------+
                                                      | IPFS             |
                                                      | - Data Storage   |
                                                      | - CID Generation |
                                                      +------------------+
                                                               |
                                                               V
                                                      +------------------+
                                                      | CosmWasm Contract|
                                                      | Photosynthesis-Archway Chain|
                                                      | - Sync Merkle    |
                                                      |   Tree Root Hash |
                                                      +------------------+

These root hashes are synced at regular intervals (5-minute, hourly, daily) according to Merkle Tree generation schedule above which can be tuned.



```
