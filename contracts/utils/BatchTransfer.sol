// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract BatchTransfer is Ownable {
	/**
	 * @dev Emitted when a token transfer fails.
	 * @param recipient The address of the recipient that failed to receive the tokens.
	 */
	error TransferFailed(address recipient);

	/**
	 * @notice Transfers `amount` tokens of `tokenAddress` to each of the `recipients`.
	 * @dev Can only be called by the contract owner.
	 * @param tokenAddress The address of the ERC20 token to be transferred.
	 * @param recipients The addresses of the recipients.
	 * @param amount The amount of tokens to be transferred to each recipient.
	 */
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
			bool success = token.transfer(recipients[i], amount);
			if (!success) {
				revert TransferFailed(recipients[i]);
			}
		}
	}

	/**
	 * @notice Withdraws all tokens of `tokenAddress` from the contract.
	 * @dev Can only be called by the contract owner.
	 * @param tokenAddress The address of the ERC20 token to be withdrawn.
	 */
	function withdraw(address tokenAddress) external onlyOwner {
		IERC20 token = IERC20(tokenAddress);

		require(token.balanceOf(address(this)) > 0, 'Contract has no balance of such token.');

		token.transfer(msg.sender, token.balanceOf(address(this)));
	}
}
