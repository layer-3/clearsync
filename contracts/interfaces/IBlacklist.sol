// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.20;

/**
 * @title The IBlacklist interface outlines the ability to prevent certain accounts from transferring tokens.
 *
 * @notice This feature is blocking transfers of reported stolen funds from exchanges or engaged in malicious activities.
 *
 * To safeguard user funds against any potential manipulation, specific roles are assigned to different accounts.
 * One account is authorized to blacklist other accounts while another account is authorized to burn funds.
 * By dividing these responsibilities between two different accounts, the risk of misuse of this functionality is reduced.
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
