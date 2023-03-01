// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract RandomUpgradeable is Initializable {
    bytes32 private salt;
    uint256 private nonce;

    function __Random_init() internal onlyInitializing {
    }

    function __Random_init_unchained() internal onlyInitializing {
    }

    // internal
    modifier Random {
        _;
        _updateNonce();
    }

    // specifies an external function which uses Random logic
    modifier UseRandom {
      _;
      _updateSalt();
    }

    function _updateNonce() private {
        unchecked {
            nonce++;
        }
    }

    function _updateSalt() private {
        salt = keccak256(abi.encode(msg.sender, block.timestamp));
    }

    function _randomMaxNumber(uint256 max) internal Random returns (uint256) {
        return uint256(keccak256(abi.encode(salt, nonce, msg.sender, block.timestamp))) % max;
    }
}
