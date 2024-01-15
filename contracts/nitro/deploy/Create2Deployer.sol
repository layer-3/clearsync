// SPDX-License-Identifier: MIT
// Taken from https://github.com/OpenZeppelin/openzeppelin-contracts/blob/7201e6707f6631d9499a569f492870ebdd4133cf/contracts/mocks/Create2Impl.sol

pragma solidity ^0.8.20;

import {Create2} from '@openzeppelin/contracts/utils/Create2.sol';

contract Create2Deployer {
    function deploy(uint256 value, bytes32 salt, bytes memory code) public {
        Create2.deploy(value, salt, code);
    }

    function computeAddress(bytes32 salt, bytes32 codeHash) public view returns (address) {
        return Create2.computeAddress(salt, codeHash);
    }

    function computeAddressWithDeployer(
        bytes32 salt,
        bytes32 codeHash,
        address deployer
    ) public pure returns (address) {
        return Create2.computeAddress(salt, codeHash, deployer);
    }
}
