// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @notice Interface describing functionality of blocking accounts from transferring tokens.
 * This limitation is going to be applied only to hackers and malicious users, who was confirmed to had stolen funds from any exchanges.
 *
 * In order to ensure that no individual can manipulate user funds, it is necessary to assign specific roles to different accounts.
 * One account with a particular role should be authorized to blacklist other accounts, while another account with a different role should be authorized to burn funds.
 * This separation of responsibilities between two different accounts mitigate the risk of misuse of this functionality.
 *
 */
interface IBlacklist {
	/**
	 * @notice Mark `account` as 'blacklisted' and disallow `transfer` and `transferFrom` operations.
	 * @dev Require `COMPLIANCE_ROLE` to invoke. Emit `Blacklisted` event`.
	 * @param account Address of account to mark 'blacklisted'.
	 */
	function blacklist(address account) external;

	/**
	 * @notice Remove mark 'blacklisted' from `account`, reinstating abilities to invoke `transfer` and `transferFrom`.
	 * @dev Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedRemoved` event`.
	 * @param account Address of account to remove 'blacklisted' mark from.
	 */
	function removeBlacklisted(address account) external;

	/**
	 * @notice Burn all tokens from blacklisted `account` specified.
	 * @dev Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedBurnt` event`. Account specified must be blacklisted.
	 * @param account Address of 'blacklisted' account to burn funds from.
	 */
	function burnBlacklisted(address account) external;

	/// Events

	/**
	 * @notice `Account` was marked 'blacklisted'.
	 * @param account Address of account to have been marked 'blacklisted'.
	 */
	event Blacklisted(address indexed account);

	/**
	 * @notice Mark 'blacklisted' from `account` was removed.
	 * @param account Address of account 'blacklisted' mark was removed from.
	 */
	event BlacklistedRemoved(address indexed account);

	/**
	 * @notice All tokens from blacklisted `account` specified were burnt.
	 * @param account Address of 'blacklisted' account funds were burnt from.
	 */
	event BlacklistedBurnt(address indexed account, uint256 amount);
}
