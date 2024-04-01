// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EvmosAdSystem {
    struct Ad {
        uint256 id;
        string imageURL;
        string targetURL;
        uint256 views;
        address rewardAddress;
        uint256 pltAmount;
    }

    Ad[] private ads;
    uint256 private totalViews;
    address public pltAddress; 

    event AdAdded(uint256 indexed id, address indexed rewardAddress, uint256 pltAmount);
    event AdDeleted(uint256 indexed id);
    event RewardsDistributed(address[] recipients, uint256[] amounts);

  
    address owner;
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    constructor(address _pltAddress) {
        owner = msg.sender;
        pltAddress = _pltAddress;
    }

    function addAd(
        string memory imageURL,
        string memory targetURL,
        address rewardAddress,
        uint256 pltAmount
    ) public onlyOwner {
        uint256 newAdId = ads.length + 1;
        ads.push(Ad(newAdId, imageURL, targetURL, 0, rewardAddress, pltAmount));
        emit AdAdded(newAdId, rewardAddress, pltAmount);
    }

    function deleteAd(uint256 adId) public onlyOwner {
        for (uint256 i = 0; i < ads.length; i++) {
            if (ads[i].id == adId) {
                delete ads[i];
                emit AdDeleted(adId);
                break;
            }
        }
    }

    function distributeRewards() public onlyOwner {
        address[] memory recipients = new address[](ads.length);
        uint256[] memory amounts = new uint256[](ads.length);

        for (uint256 i = 0; i < ads.length; i++) {
            Ad storage ad = ads[i];
            uint256 reward = ad.views * ad.pltAmount;
            totalViews += ad.views;
            ads[i].views = 0; // Reset views after distributing rewards
            recipients[i] = ad.rewardAddress;
            amounts[i] = reward;
            // Transfer rewards to the ad's reward address
            payable(ad.rewardAddress).transfer(reward);
        }
        emit RewardsDistributed(recipients, amounts);
    }

    function queryAd(uint256 adId) public view returns (Ad memory) {
        for (uint256 i = 0; i < ads.length; i++) {
            if (ads[i].id == adId) {
                return ads[i];
            }
        }
        revert("Ad not found");
    }

    function queryAllAds() public view returns (Ad[] memory) {
        return ads;
    }

    function queryTotalViews() public view returns (uint256) {
        return totalViews;
    }

    // Allow the contract to receive ETH
    receive() external payable {}

    // Function to retrieve funds from the contract
    function withdraw(uint256 amount) public onlyOwner {
        require(address(this).balance >= amount, "Insufficient balance");
        payable(msg.sender).transfer(amount);
    }
}

