//SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ILegacyERC20 {
    function transfer(address to, uint256 amount) external;

    function transferFrom(address from, address to, uint256 amount) external;
}