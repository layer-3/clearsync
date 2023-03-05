// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/access/AccessControl.sol';

import './interfaces/IBlacklist.sol';

/**
 * @notice Yellow and Canary utility token inheriting AccessControl and implementing Cap and Blacklist.
 * This smart contract is an ERC20 used by both YELLOW and DUCKIES tokens.
 * The YELLOW token is a collateral to open a state channel with another network entity.
 * Additionally, it is used to pay the settlement fees on the network.
 *
 * After deployment, DEFAULT_ADMIN_ROLE will be transferred to a DAO, which will govern the token.
 * This is done not to give too much token governance power to once account, which will definitely be a vector of attack.
 *
 * The similar applies to COMPLIANCE_ROLE. It is going to be granted to a multisig account, which will govern hackers and malicious users by blacklisting them.
 *
 * @dev Blacklist feature is using OpenZeppelin AccessControl.
 */
contract YellowToken is ERC20, AccessControl, IBlacklist {
	/// @dev Role for managing the blacklist process chosen by the DAO
	bytes32 public constant COMPLIANCE_ROLE = keccak256('COMPLIANCE_ROLE');

	/// @dev Role for user blacklisted
	bytes32 public constant BLACKLISTED_ROLE = keccak256('BLACKLISTED_ROLE');

	/// @dev Role given to the DAO snapshot
	bytes32 public constant MINTER_ROLE = keccak256('MINTER_ROLE');

	/// @dev Activation must be called at Token Listing Event.
	uint256 public activatedAt;

	/// @dev Token maximum supply
	uint256 public immutable TOKEN_SUPPLY_CAP;

	/**
	 * @notice Activated event. Emitted when `activate` function is invoked.
	 * @param premint Amount of tokes pre-minted during activation.
	 */
	event Activated(uint256 premint);

	/**
	 * @dev Simple constructor, passing arguments to ERC20 constructor.
	 * Grants `DEFAULT_ADMIN_ROLE` and `MINTER_ROLE` to deployer.
	 * @param name Name of the Token.
	 * @param symbol Symbol of the Token.
	 */
	constructor(string memory name, string memory symbol, uint256 supplyCap) ERC20(name, symbol) {
		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(MINTER_ROLE, msg.sender);
		TOKEN_SUPPLY_CAP = supplyCap;
	}

	/// Token functions

	/**
	 * @notice Return the cap on the token's total supply.
	 * @return uint256 Token supply cap.
	 */
	function cap() external view returns (uint256) {
		return TOKEN_SUPPLY_CAP;
	}

	/**
	 * @notice Return the number of decimals used to get its user representation.
	 * @dev Overrides ERC20 default value of 18;
	 * @return uint8 Number of decimals of Token.
	 */
	function decimals() public pure override returns (uint8) {
		return 8;
	}

	/**
	 * @notice Activate token, minting `premint` amount to `account` address.
	 * @dev Require `DEFAULT_ADMIN_ROLE` to invoke. Premint must satisfy these conditions: 0 < premint < token supply cap. Can be called only once.
	 * @param premint Amount of tokens to premint.
	 * @param account Address of account to premint to.
	 */
	function activate(uint256 premint, address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		require(premint > 0, 'Zero premint');
		require(premint <= TOKEN_SUPPLY_CAP, 'Premint exceeds cap');
		require(activatedAt == 0, 'Already activated');

		activatedAt = block.timestamp;
		_mint(account, premint);

		emit Activated(premint);
	}

	/**
	 * @notice Increase balance of address `to` by `amount`.
	 * @dev Require `MINTER_ROLE` to invoke. Emit `Transfer` event.
	 * Require Token to be activated.
	 * The following conditions must be satisfied: `totalSupply + amount <= supplyCap`.
	 * @param to Address to mint tokens to.
	 * @param amount Amount of tokens to mint.
	 */
	function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
		require(activatedAt > 0, 'Not activated');
		require(totalSupply() + amount <= TOKEN_SUPPLY_CAP, 'Mint exceeds cap');

		_mint(to, amount);
	}

	/**
	 * @notice Destroys `amount` tokens from caller's account. Emit `Transfer` event.
	 * @param amount Amount of tokens to burn.
	 */
	function burn(uint256 amount) external {
		_burn(msg.sender, amount);
	}

	/**
	 * @notice Destroys `amount` tokens from `account`, deducting from the caller's allowance. Emit `Transfer` event.
	 * @param account Address of account to burn tokens from.
	 * @param amount Amount of tokens to burn.
	 */
	function burnFrom(address account, uint256 amount) external {
		_spendAllowance(account, msg.sender, amount);
		_burn(account, amount);
	}

	/**
	 * @notice Transfer `amount` of tokens to `to` address from caller.
	 * @dev Require caller is not marked 'blacklisted'.
	 * @param to Address to transfer tokens to.
	 * @param amount Amount of tokens to transfer.
	 * @return bool true if transfer succeeded.
	 */
	function transfer(address to, uint256 amount) public override returns (bool) {
		_requireAccountNotBlacklisted(msg.sender);

		return ERC20.transfer(to, amount);
	}

	/**
	 * @notice Transfer `amount` of tokens from `from` to `to` address.
	 * @dev Require `from` is not marked 'blacklisted'.
	 * @param from Address to transfer tokens from.
	 * @param to Address to transfer tokens to.
	 * @param amount Amount of tokens to transfer.
	 * @return bool true if transfer succeeded.
	 */
	function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
		_requireAccountNotBlacklisted(from);

		return ERC20.transferFrom(from, to, amount);
	}

	/// Blacklist Implementation

	/**
	 * @notice Mark `account` as 'blacklisted' and disallow `transfer` and `transferFrom` operations.
	 * @dev Require `COMPLIANCE_ROLE` to invoke. Emit `Blacklisted` event`.
	 * @param account Address of account to mark 'blacklisted'.
	 */
	function blacklist(address account) external onlyRole(COMPLIANCE_ROLE) {
		_grantRole(BLACKLISTED_ROLE, account);
		emit Blacklisted(account);
	}

	/**
	 * @notice Remove mark 'blacklisted' from `account`, reinstating abilities to invoke `transfer` and `transferFrom`.
	 * @dev Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedRemoved` event`.
	 * @param account Address of account to remove 'blacklisted' mark from.
	 */
	function removeBlacklisted(address account) external onlyRole(COMPLIANCE_ROLE) {
		_revokeRole(BLACKLISTED_ROLE, account);
		emit BlacklistedRemoved(account);
	}

	/**
	 * @notice Burn all tokens from blacklisted `account` specified.
	 * @dev Require `DEFAULT_ADMIN_ROLE` to invoke. Emit `BlacklistedBurnt` event`.
	 * Account specified must be blacklisted.
	 * @param account Address of 'blacklisted' account to burn funds from.
	 */
	function burnBlacklisted(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		require(hasRole(BLACKLISTED_ROLE, account), 'Account is not blacklisted');

		uint256 blackFundsAmount = balanceOf(account);
		_burn(account, blackFundsAmount);

		emit BlacklistedBurnt(account, blackFundsAmount);
	}

	/**
	 * @notice Internal Function
	 * @dev Require `account` is not marked 'blacklisted'.
	 * @param account Address of account to require not marked 'blacklisted'.
	 */
	function _requireAccountNotBlacklisted(address account) internal view {
		require(!hasRole(BLACKLISTED_ROLE, account), 'Account is blacklisted');
	}
}
