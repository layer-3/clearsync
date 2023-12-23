// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol';

/**
 * @notice Yellow and Canary utility token is a simple ERC20 token with permit functionality.
 * All the supply is minted to the deployer.
 */
contract SimpleYellowToken is ERC20Permit {
	/// @dev Token maximum supply
	uint256 public immutable TOKEN_SUPPLY_CAP;

	/**
	 * @dev Simple constructor, passing arguments to ERC20Permit and ERC20 constructors.
	 * Mints the supply cap to the deployer.
	 * @param name Name of the Token.
	 * @param symbol Symbol of the Token.
	 * @param supplyCap Maximum supply of the Token.
	 */
	constructor(
		string memory name,
		string memory symbol,
		uint256 supplyCap
	) ERC20Permit(name) ERC20(name, symbol) {
		TOKEN_SUPPLY_CAP = supplyCap;
		_mint(msg.sender, supplyCap);
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
	 * @notice Return the cap on the token's total supply.
	 * @return uint256 Token supply cap.
	 */
	function cap() external view returns (uint256) {
		return TOKEN_SUPPLY_CAP;
	}
}
