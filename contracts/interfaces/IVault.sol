// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @title IVault Interface
 * @dev Interface for a smart contract vault that can hold multiple currencies, allowing users to deposit and withdraw funds in any supported token.
 */
interface IVault {
	/**
	 * @dev Deposits the specified token into the vault.
	 * @param token The address of the token being deposited.
	 * @param amount The amount of the token being deposited.
	 */
	function deposit(address token, uint256 amount) external payable;

	/**
	 * @dev Withdraws the specified token from the vault.
	 * @param token The address of the token being withdrawn.
	 * @param amount The amount of the token to be withdrawn.
	 */
	function withdraw(address token, uint256 amount) external;

	/**
	 * @dev Returns the balance of the specified token for the caller.
	 * @param token The address of the token being queried.
	 * @return The balance of the token held by the caller.
	 */
	function getBalance(address token) external view returns (uint256);
}
