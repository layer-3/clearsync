// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract BatchTransfer is Ownable {
	function batchTransfer(
		address tokenAddress,
		address[] memory recipients,
		uint256 amount
	) external onlyOwner {
		IERC20 token = IERC20(tokenAddress);

		require(
			token.balanceOf(address(this)) >= amount * recipients.length,
			'Contract has insufficient balance.'
		);

		for (uint256 i = 0; i < recipients.length; i++) {
			require(token.transfer(recipients[i], amount), 'Token transfer failed.');
		}
	}

	function withdraw(address tokenAddress) external onlyOwner {
		IERC20 token = IERC20(tokenAddress);

		require(token.balanceOf(address(this)) > 0, 'Contract has no balance of such token.');

		token.transfer(msg.sender, token.balanceOf(address(this)));
	}

	function balanceOf(address tokenAddress) public view returns (uint256) {
		IERC20 token = IERC20(tokenAddress);
		return token.balanceOf(address(this));
	}
}
