# Adserver smart contract

Codacy grade A -
https://app.codacy.com/gh/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/file/108971284608/issues/source?bid=34568294&fileBranchId=34568294

# Heavy Transaction Guaranteed Dapps on Daily Basis – A strong candidate for Archway Rewards and Photosynthesis liquid token

https://twitter.com/coolcaptchas/status/1652382322153652225?s=20

https://twitter.com/coolcaptchas/status/1652391219086131200?s=20

https://twitter.com/coolcaptchas/status/1527443243327434752?s=20

https://twitter.com/coolcaptchas/status/1506448375474180098?s=20

https://docs.google.com/document/d/1-pGzhqz7UHgB2O9qD4fcLAhwhTUgUI4A/edit?usp=sharing&ouid=102246369981228451498&rtpof=true&sd=true

Layer 2 chain derivative -
https://github.com/suhasagg/BlockchainBridgeSystemDesign


ArchID:

<img src="https://cdn.discordapp.com/emojis/1014847480539652118.webp?size=96&quality=lossless" width="100">

https://archid.app/domains/iab.arch

https://archid.app/domains/cookieattentiontoken.arch

https://archid.app/domains/cookiepersona.arch

https://archid.app/domains/datamanagementplatform.arch

https://archid.app/domains/crossdevice.arch

https://archid.app/domains/admanager.arch

https://archid.app/domains/adserver.arch


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

Applications:

1)Handling High Traffic: The architecture is designed to handle high traffic volumes of ad data and cookie sync data using benefits of  distributed nature of IPFS.

2)Asynchronous Data Handling: Asynchronous processing of data collection, Merkle Tree generation, and blockchain syncing allows for more efficient use of network resources and can mitigate latency issues.

3)Scheduled Syncing: Regular, scheduled syncing of data and Merkle Tree root hashes prevents sudden spikes in network demand, contributing to smoother network traffic.

4)Data Integrity and Security: The use of Merkle Trees ensures data integrity, as any alteration in the data leads to a different hash. This is crucial for ad data and cookie tracking.

5)Tamper-Proof Audit Trail: Storing Merkle Tree root hashes on a blockchain provides a tamper-proof audit trail, enhancing transparency and trust in the data collected.
```

![GGR7CcFXQAAGZUY](https://github.com/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/assets/3880512/5086affc-61ed-4ac5-9b51-eaf9f768e749)


```
Campaign Delivery System Design for TV, Laptop, Mobile, and VR Devices 📱💻🖥️🕶️

System Objective: To develop a comprehensive, asynchronous campaign delivery system that integrates seamlessly with TV operating systems, laptops, mobile devices, and VR platforms, providing users with a non-intrusive, engaging way to access personalized campaigns.

Target Platforms:
Smart TVs (e.g., Android TV, Roku, Apple TV) 📺
Laptops (Windows, macOS) 💻
Mobile Devices (iOS, Android) 📱
Virtual Reality Devices 🕶️

Core Components:

1. Operating System Integration
   Smart TV OS: A dedicated app leveraging TV OS features for campaign updates. 📺
   Laptop and Mobile OS: App versions for Windows, macOS, iOS, and Android, using each OS's notification systems. 💻📱
   VR Platforms: VR experiences compatible with major VR systems (e.g., Oculus, HTC Vive), allowing users to explore campaigns in a 3D environment. 🕶️
   Amazon, Facebook, Twitter App - Campaign Section Integration support. 😊🌹😊🌹😊🌹🎁🎁🎁🌹🌹🌹
   
2. Campaign Channel/App Design
   User Interface: Intuitive interfaces suitable for each platform, including a 3D VR environment for immersive interaction. 🖌️
   Content Management System: A CMS backend for regular campaign content updates across all platforms. 🗂️
   Personalization Engine: AI to curate content based on user data, enhancing relevance and engagement. 🤖
   
3. Asynchronous Delivery and Duplicate Rendering
   Asynchronous Updates: Schedules periodic content refreshes, ensuring up-to-date campaigns are available on demand. ⏲️
   Duplicate Management: Identify and manage duplicate content, prioritizing new and relevant campaigns for each user. 🔄

4. User Engagement and Control
   Opt-in/Opt-out: Offer clear options for users to control their participation in the system, ensuring privacy and consent. 🔐
   Engagement Tracking: Collects data on how users interact with campaigns, using insights to refine future content. 📊
   Feedback System: Enables users to provide direct feedback on campaigns, further tailoring the experience. 💬

5. Security and Privacy
   Data Protection: Ensures encryption and secure storage of user data, adhering to international privacy standards. 🛡️
   Anonymous Profiling: Employs techniques to minimize personal data use while maintaining personalization. 👥

6. Cross-Platform Synchronization
   Account-Based Sync: Allows users to synchronize preferences and history across devices, providing a cohesive experience. 🔗

VR Specific Enhancements:
Immersive Campaigns: Campaigns leverage VR's immersive capabilities, offering users a novel way to engage with content. 🌐
Interactive Experiences: Interactive VR experiences allow users to engage with campaigns actively, rather than passively viewing them. 🎮


Applications of This Architecture ✨

Long Campaign Running Period: Campaigns can run over extended periods, ensuring maximum exposure. 🕰️

Repeated Views: Users can revisit campaigns, increasing familiarity and reinforcing messaging. 🔁

Full Campaigns Available 24/7: Access to campaigns anytime ensures no missed opportunities for engagement. 🌞🌜

Max Conversion Interactivity: Interactive elements within campaigns boost engagement and potential for conversion.🖱️📈

No Campaign Misses for Users: The asynchronous delivery system ensures users receive all relevant campaigns, tailored to their interests and behaviors. 🚫📵

```

![GGR7CcEXMAADze-](https://github.com/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/assets/3880512/f23fe0d1-9fa4-4e82-959b-1d3cc365ddd5)


```

Integration with Dedicated YouTube Channels 📹 ✨

Objective: Utilize YouTube as a platform for extending the personalized campaign delivery system, providing each user with a dedicated channel filled with curated campaign content based on their interests and behaviors. 🎯


Implementation Strategies:


Channel Creation and Management: 🛠️

Automatically generate a YouTube channel for each user upon opting into the campaign delivery system. Each channel's content is curated based on the user's browsing personas and engagement history. 🧑‍💻

Use a content management system (CMS) to manage video campaigns across these personalized channels, ensuring content remains fresh and relevant. 🔄


Personalization Engine Integration: 🤖

Leverage AI and machine learning algorithms to analyze user data continuously and update the campaign videos on each dedicated YouTube channel, ensuring high relevance and personalization. 🔍

Incorporate feedback loops from user interactions on YouTube (likes, dislikes, comments) to refine and adapt the content further. 🔄👍👎


Asynchronous and Interactive Content Delivery: ⏱️

Schedule video campaigns for asynchronous delivery, allowing users to access content at their convenience, enhancing engagement without intruding on their daily activities. 🕒

Incorporate interactive elements within videos (e.g., polls, clickable links, Q&A sessions) to foster a more engaging and immersive experience. 💡🖱️


NFT-based Creative Verification 🛡️🎨

Convert high-performing campaign creatives into NFTs to certify their authenticity and originality. This step not only adds a layer of security and ownership to digital assets but also creates a unique value proposition for the content. 📜✅

Enhance Content Value: By tokenizing creatives, brands can underscore the unique aspects of their campaigns, making them more appealing and valuable to the audience. 💎


Engagement Incentives 🏆💥

Offer users the opportunity to earn or win NFTs by engaging with campaign content (e.g., watching videos, participating in polls, sharing content across social media channels). This not only incentivizes engagement but also introduces users to the concept of digital collectibles. 🎁👀

Interactive Rewards: Encourage active participation and deeper interaction with campaigns by rewarding users with something of tangible value and exclusivity. 🌟


Marketplace Integration 💹🔗

Establish a marketplace or partner with existing NFT platforms to allow users to trade or sell the NFTs they earn. This can create a new revenue stream for content creators and advertisers, while also increasing the longevity and interaction with campaign content. 🛒🔄

Economic Ecosystem: Create a vibrant ecosystem where digital assets can be traded, adding a new dimension to user engagement and monetization strategies for creators. 💰🌍



Benefits of Creative Experimentation Using This System 🎨✨

Personalized User Experiences: By continuously refining creatives based on user interaction and feedback, the system ensures that each user receives the most engaging and relevant content 🎯, enhancing satisfaction and brand perception. 🌈

Increased Engagement and Conversion: Through targeted, optimized, and interactive campaign content, the system aims to maximize user engagement and conversion rates 💥, driving higher ROI for advertisers. 📈

Innovative Brand Interaction: The use of NFTs for rewarding engagement introduces users to new forms of digital interaction 🚀, positioning brands as forward-thinking and technologically savvy. 🌍

Data-Driven Insights: The wealth of data generated from creative testing and user interactions provides deep insights into consumer behavior and preferences 🔍, aiding in the development of future campaigns and marketing strategies. 📊


Achieve the best organic CPM, CTR, and Engagement Time with a component of nostalgia 🕰️.

Campaigns get sorted according to the calendar 📅, and users can view them whenever they desire, saving advertisers costs and eliminating campaign expiry date hassles at no extra cost.

Campaigns get archived in the mobile app 📱 and can be reviewed by users at any point in the future depending on the IAB taxonomy of the campaign. 📚🔍

This system ensures that every ad has its time to shine, again and again, bringing back cherished memories and maximizing user engagement!! 🌟💬🔄


Campaign Sorting Steps:

Step 1: Data Collection 📊

Collect real-time and historical data on user interactions with campaigns.

Categorize content and user interactions into IAB segments, geography, device, and more segments. 🌐
Update the topic cloud for each user. ☁️


Step 2: Calculate Interest Scores 📈

Frequency Score (F): Count how often the user interacts with campaigns within each IAB segment. 🔢

Recency Score (R): Calculate the time decay factor for each segment based on the last interaction. Older interactions have less weight, and recent interactions score higher. ⏳


Step 3: Combine Scores Using a Weighted Formula 🔍

Interest Score (I): Combine frequency and recency scores to compute an overall interest score for each IAB segment for a user.
```

![equation](https://github.com/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/assets/3880512/58b20900-bd67-4e75-b43a-de6a6666f9ba) 🧮

Where:

    • w1​,w2​ represent the weights assigned to frequency and recency.
    
    • λ is the decay constant indicating how quickly past interactions lose relevance.
    
    • t measures the time elapsed since the last interaction.
 
```
Step 4: Sort Campaigns 🗂️

Sort available campaigns based on their match with the user’s highest Interest Scores across IAB segments.
Prioritize campaigns that not only fit the highest Interest Scores but also complement the user's topic cloud. 🎯


Implementation and Benefits 🌟

This algorithm can be implemented within the existing CMS and personalization engine. By leveraging this strategy, campaigns will be more effectively targeted, enhancing user engagement because the content will resonate more deeply on a personal level. ❤️

Additionally, this sorting algorithm helps in minimizing ad fatigue, as it dynamically adjusts to changing interests and preferences, thus optimizing both advertiser spend and user experience. 🔄

By integrating this advanced sorting algorithm, the campaign delivery system ensures that nostalgia and other user preferences are leveraged effectively, campaigns are delivered with an understanding of user behavior over time, and advertisers can achieve better engagement metrics at reduced costs. ✅📉

```



```


🚀 Async Campaign Chatting Feature 🎉 for Best Campaign User Experience ✨



1. Chat Feature Development 🛠️
   - Custom chat module or integrate a third-party service to enable messaging within the app. This module supports starting new conversations, displaying existing conversations, and managing multiple chat threads.


2. User Authentication 🔒
   - Use OAuth or a similar protocol to authenticate users. This ensures that the chat feature is accessible only to registered and logged-in users, enhancing security and providing personalized interactions.


3. Real-Time Messaging ⚡
   - Implement WebSocket for a persistent connection between the client and the server or use Firebase Cloud Messaging (FCM) to handle the delivery and reception of messages in real-time.


4. Notification System 📬
   - Develop or integrate a push notification system to notify users about new messages or actions needed. This can be done through Android’s native notification system or FCM.


5. UI/UX Design 🎨
   -  Clean, intuitive user interface for the chat feature. Include essentials such as message bubbles, timestamps, sender identification, and multimedia message capabilities (like sending images and videos).


6. Integration with Campaign Platform 🔄
   - Chat feature is integrated with the rest of the campaign platform functionalities. This involves linking chat sessions with specific campaign issues or queries to provide contextual support.

```


```

Enhancements for Asynchronous Chat Support 🚀


Message Queueing System 📥:
Implement a message queueing system to manage asynchronous communication. This system stores messages sent by users when the support team is unavailable and ensures they are delivered when the team is ready to respond.


Offline Messaging Capability 📶:
Enable users to send messages even when they are offline or when the support team is not immediately available. These messages are stored locally and synced with the server once the connection is re-established.


Scheduled Messages 📆:
Allow users to schedule messages for future delivery, useful for inquiries about upcoming promotions or products. This feature can also notify users when a scheduled time for a discount or early bird offer is approaching.


Read Receipts and Last Seen 👀:
Incorporate read receipts and indicators showing when the support team last viewed the chat. This keeps users informed about the status of their queries and manages expectations regarding response times.


Automated Responses 🤖:
Develop an automated response system that can provide immediate basic answers to common questions or acknowledge the receipt of messages. 


Push Notifications for Updates 🔔:
Ensure that users receive push notifications for new messages or updates from the advertiser, especially for ongoing discussions about product demos or special offers. This keeps the user engaged and informed without needing to check the app constantly.


Rich Media Support 🎥:
Enhance the chat interface to support rich media, such as videos, images, and documents. This is useful for product demos and detailed explanations of offers or discounts.


Searchable Chat Archives 🔍:
Implement a feature that allows users to search their chat archives. This can help them quickly find information about past conversations regarding offers, product details, or support interactions.


User Feedback Collection ✍️:
Integrate a mechanism for collecting user feedback directly through the chat interface. This feedback can be about the chat experience itself or the products and offers discussed within the chat.

```


```

YouTube Channel Interface Integration 🎥


1. YouTube Messaging API 📝
   - Utilize the YouTube Data API to integrate messaging features directly on the channel. This involve solutions like linking to external chat platforms or embedding chat in the channel page.



2. User Permissions 🛂
   - Handle permission requests carefully to access user data for messaging. Clearly communicate which permissions are needed and why, ensuring compliance with YouTube’s policies and user privacy standards.
   


3. Real-Time Communication 🗨️
   - Implement real-time messaging on the YouTube interface, which may involve using WebSockets or polling mechanisms to retrieve messages without refreshing the page.



4. Notification Integration 🔔
   - Integrate with YouTube’s notification system or develop a custom notification feature to alert users about new messages directly on YouTube.



5. User Experience Enhancement 💡
   - Create a seamless integration that does not disrupt the user’s browsing experience on YouTube. The chat interface should be minimally invasive but easily accessible.


6. Integration with Campaign Content 🔗
   - Link chat functionalities with specific campaign content on the YouTube channel, such as enabling chat prompts on particular videos or campaign-related posts.

```

![dancing-robot-lego-boost](https://github.com/suhasagg/Photosynthesis-Dorahacks-web3-competition-winner/assets/3880512/07fcba66-4ec4-4427-a4ad-2e2f3a188fd4)

