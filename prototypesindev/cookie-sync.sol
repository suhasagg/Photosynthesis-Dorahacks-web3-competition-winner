// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CookieContract {
    struct Cookie {
        uint256 id;
        string domain;
        string data;
        uint256 expiration;
    }

    struct CookiePacketData {
        Cookie cookie;
        address sourcePub;
        address destPub;
    }

    Cookie[] public cookies;
    mapping(string => Cookie) private cookiesById;

    event CookieCreated(uint256 id, string domain, string data, uint256 expiration);
    event CookieSynced(uint256 id, address sourcePub, address destPub);
    event BatchSynced(uint256[] ids, address sourcePub, address destPub);

    function createCookie(string memory domain, string memory data, uint256 expiration) public {
        uint256 newId = cookies.length + 1;
        Cookie memory newCookie = Cookie(newId, domain, data, expiration);
        cookies.push(newCookie);
        cookiesById[uint2str(newId)] = newCookie;

        emit CookieCreated(newId, domain, data, expiration);
    }

    function syncCookie(CookiePacketData memory packet) public {
        require(packet.cookie.id > 0 && packet.cookie.id <= cookies.length, "Cookie does not exist");
        emit CookieSynced(packet.cookie.id, packet.sourcePub, packet.destPub);
    }

    function batchSyncCookies(CookiePacketData[] memory packets) public {
        uint256[] memory syncedIds;
        for (uint256 i = 0; i < packets.length; i++) {
            require(packets[i].cookie.id > 0 && packets[i].cookie.id <= cookies.length, "Cookie does not exist");
            syncedIds[i] = packets[i].cookie.id;
        }
        emit BatchSynced(syncedIds, packets[0].sourcePub, packets[0].destPub);
    }

    function getCookie(uint256 id) public view returns (Cookie memory) {
        require(id > 0 && id <= cookies.length, "Cookie does not exist");
        return cookies[id - 1];
    }

    function listCookies() public view returns (Cookie[] memory) {
        return cookies;
    }

    function uint2str(uint256 _i) private pure returns (string memory) {
        if (_i == 0) {
            return "0";
        }
        uint256 j = _i;
        uint256 length;
        while (j != 0) {
            length++;
            j /= 10;
        }
        bytes memory bstr = new bytes(length);
        uint256 k = length - 1;
        while (_i != 0) {
            bstr[k--] = bytes1(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);
    }
}

