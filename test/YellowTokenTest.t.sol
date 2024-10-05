// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Test, Vm} from 'forge-std/Test.sol';
import {YellowToken} from '../contracts/YellowToken.sol';

contract YellowTokenTest is Test {
	YellowToken token;
	address deployer = vm.createWallet('deployer').addr;

	string constant NAME = 'YellowToken';
	string constant SYMBOL = 'YELLOW';
	uint8 constant DECIMALS = 8;
	uint256 constant TOKEN_SUPPLY = 10_000_000_000;

	function setUp() public {
		vm.prank(deployer);
		token = new YellowToken(NAME, SYMBOL, TOKEN_SUPPLY);
	}

	function test_constructor() public view {
		assertEq(token.name(), NAME);
		assertEq(token.symbol(), SYMBOL);
		assertEq(token.decimals(), DECIMALS);
		assertEq(token.totalSupply(), TOKEN_SUPPLY);
		assertEq(token.balanceOf(deployer), TOKEN_SUPPLY);
	}
}
