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
		require(
			_getBalance(tokenAddress) >= amount * recipients.length,
			'Contract has insufficient balance.'
		);

		for (uint256 i = 0; i < recipients.length; i++) {
			if (!_transfer(tokenAddress, recipients[i], amount)) revert TransferFailed(recipients[i]);
		}
	}

	/**
	 * @notice Receives ETH transfers.
	 * @dev Required to receive ETH transfers.
	 */
	receive() external payable {}

	/**
	 * @notice Withdraws all tokens of `tokenAddress` from the contract.
	 * @dev Can only be called by the contract owner.
	 * @param tokenAddress The address of the ERC20 token to be withdrawn.
	 */
	function withdraw(address tokenAddress) external onlyOwner {
		uint256 tokenBalance = _getBalance(tokenAddress);
		require(tokenBalance > 0, 'Contract has no balance of such token.');
		
        require(_transfer(tokenAddress, msg.sender, tokenBalance), 'Could not transfer tokens');
	}

	/**
	 * @notice Returns the balance of `tokenAddress` held by the contract.
	 * @dev Can only be called by the contract owner.
	 * @param tokenAddress The address of the ERC20 token to be withdrawn.
	 */
	function _getBalance(address tokenAddress) internal view returns (uint256) {
		if (tokenAddress == address(0)) {
			return address(this).balance;
		} else {
			return IERC20(tokenAddress).balanceOf(address(this));
		}
	}

	/**
	 * @notice Transfers `amount` tokens of `tokenAddress` to the sender.
	 * @dev Can only be called by the contract owner.
	 * @param tokenAddress The address of the token to be transferred.
	 * @param amount The amount of tokens to be transferred.
	 */
	function _transfer(address tokenAddress, address recipient, uint256 amount) internal returns (bool success) {
		if (tokenAddress == address(0)) {
			(success, ) = recipient.call{value: amount}(''); //solhint-disable-line avoid-low-level-calls
		} else {
			success = IERC20(tokenAddress).transfer(recipient, amount);
		}
	}
}
