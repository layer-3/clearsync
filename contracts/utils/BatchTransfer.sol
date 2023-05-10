// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract BatchTransfer is Ownable {
	function batchTransferUniqueAmounts(
		address token,
		address[] memory recipients,
		uint256[] memory amounts,
		uint256 totalAmount
	) external onlyOwner {
		require(recipients.length == amounts.length, 'Arrays should have the same length.');

		IERC20 erc20Token = IERC20(token);

		require(
			erc20Token.allowance(msg.sender, address(this)) >= totalAmount,
			'Contract has insufficient allowance.'
		);

		for (uint256 i = 0; i < recipients.length; i++) {
			require(
				erc20Token.transferFrom(msg.sender, recipients[i], amounts[i]),
				'Token transfer failed.'
			);
		}
	}

	function batchTransferSameAmount(
		address token,
		address[] memory recipients,
		uint256 amount
	) external onlyOwner {
		IERC20 erc20Token = IERC20(token);

		require(
			erc20Token.allowance(msg.sender, address(this)) >= amount * recipients.length,
			'Contract has insufficient allowance.'
		);

		for (uint256 i = 0; i < recipients.length; i++) {
			require(
				erc20Token.transferFrom(msg.sender, recipients[i], amount),
				'Token transfer failed.'
			);
		}
	}
}
