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

These root hashes are synced at regular intervals (5-minute, hourly, daily) according to Merkle Tree generation schedule above which can be tuned.

Applications:

1)Handling High Traffic: The architecture is designed to handle high traffic volumes of ad data and cookie sync data using benefits of  distributed nature of IPFS.

2)Asynchronous Data Handling: Asynchronous processing of data collection, Merkle Tree generation, and blockchain syncing allows for more efficient use of network resources and can mitigate latency issues.

3)Scheduled Syncing: Regular, scheduled syncing of data and Merkle Tree root hashes prevents sudden spikes in network demand, contributing to smoother network traffic.

4)Data Integrity and Security: The use of Merkle Trees ensures data integrity, as any alteration in the data leads to a different hash. This is crucial for ad data and cookie tracking.

5)Tamper-Proof Audit Trail: Storing Merkle Tree root hashes on a blockchain provides a tamper-proof audit trail, enhancing transparency and trust in the data collected.

```
