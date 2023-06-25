# Adserver smart contract

#Heavy Transaction Guaranteed Dapps on Daily Basis – A strong candidate for Archway Rewards and Photosynthesis liquid token

https://docs.google.com/document/d/1-pGzhqz7UHgB2O9qD4fcLAhwhTUgUI4A/edit?usp=sharing&ouid=102246369981228451498&rtpof=true&sd=true
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
```
