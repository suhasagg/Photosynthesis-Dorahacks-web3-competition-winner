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

Applications:

1)Handling High Traffic: The architecture is designed to handle high traffic volumes of ad data and cookie sync data using benefits of  distributed nature of IPFS.

2)Asynchronous Data Handling: Asynchronous processing of data collection, Merkle Tree generation, and blockchain syncing allows for more efficient use of network resources and can mitigate latency issues.

3)Scheduled Syncing: Regular, scheduled syncing of data and Merkle Tree root hashes prevents sudden spikes in network demand, contributing to smoother network traffic.

4)Data Integrity and Security: The use of Merkle Trees ensures data integrity, as any alteration in the data leads to a different hash. This is crucial for ad data and cookie tracking.

5)Tamper-Proof Audit Trail: Storing Merkle Tree root hashes on a blockchain provides a tamper-proof audit trail, enhancing transparency and trust in the data collected.


Campaign Delivery System Design for TV, Laptop, Mobile, and VR Devices 📱💻🖥️🕶️

![Campaign TV/App](https://photos.app.goo.gl/t4yieWo3RRx7Vtgo9?raw=true)

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


```
