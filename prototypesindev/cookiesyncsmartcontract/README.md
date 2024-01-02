# Cookie Sync Smart Contract

# Heavy Transaction Guaranteed Dapps on Daily Basis – A strong candidate for Archway Rewards and Photosynthesis liquid token

Codacy grade A -
https://app.codacy.com/gh/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/file/93328450487/issues/source?bid=34568294&fileBranchId=34568294

https://twitter.com/coolcaptchas/status/1652382322153652225?s=20

https://twitter.com/coolcaptchas/status/1652391219086131200?s=20

https://twitter.com/coolcaptchas/status/1527443243327434752?s=20

https://twitter.com/coolcaptchas/status/1506448375474180098?s=20

https://docs.google.com/document/d/1-pGzhqz7UHgB2O9qD4fcLAhwhTUgUI4A/edit?usp=sharing\&ouid=102246369981228451498&rtpof=true&sd=true

Layer 2 chain derivative -
https://github.com/suhasagg/BlockchainBridgeSystemDesign

```
Saving state: CookiePacketData {
    cookie: Cookie {
        id: "default_id",
        domain: "default_domain",
        data: "default_data",
        expiration: 0,
    },
    source_pub: "new york times",
    dest_pub: "cnn",
}
State instantiated successfully.
Event created: Event {
    ty: "sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "123",
        },
    ],
}
Cookie Synced successfully.
Event created: Event {
    ty: "sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "456",
        },
    ],
}
Cookie Synced successfully.

Batch Cookie Sync for optimised gas usage and Archway rewards distribution

Saving state: CookiePacketData {
    cookie: Cookie {
        id: "default_id",
        domain: "default_domain",
        data: "default_data",
        expiration: 0,
    },
    source_pub: "new york times",
    dest_pub: "cnn",
}
State instantiated successfully.
Event created: Event {
    ty: "batch-sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "123",
        },
    ],
}
Event created: Event {
    ty: "batch-sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "456",
        },
    ],
}
Event created: Event {
    ty: "batch-sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "789",
        },
    ],
}
Event created: Event {
    ty: "batch-sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "012",
        },
    ],
}
Event created: Event {
    ty: "batch-sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "345",
        },
    ],
}
Cookies batch synced successfully.
Event created: Event {
    ty: "sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "123",
        },
    ],
}
Cookie Synced successfully.
Event created: Event {
    ty: "sync",
    attributes: [
        Attribute {
            key: "action",
            value: "sync",
        },
        Attribute {
            key: "from",
            value: "new york times",
        },
        Attribute {
            key: "to",
            value: "cnn",
        },
        Attribute {
            key: "cookie",
            value: "456",
        },
    ],
}
Cookie Synced successfully.

Protocol Scaling Algorithm

+------------------+      +------------------+      +------------------+
| Website          |      | Data Processing  |      | Merkle Tree      |
| - Collect Cookie | ---> | and Aggregation  | ---> | Generation       |
|   Data           |      | - Data Formatting|      | - 5-min/hourly/  |
|                  |      | - Data Encryption|      |   daily trees    |
|                  |      | - Stream to IPFS |      | - Cookie data as |
+------------------+      +------------------+      |   leaves         |
                                                    +------------------+
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
                                                      | on Cosmos Chain  |
                                                      | - Sync Merkle    |
                                                      |   Tree Root Hash |
                                                      +------------------+


```
