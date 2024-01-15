// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.22;

import '@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol';

/**
 * @notice Yellow and Canary utility token is a simple ERC20 token with permit functionality.
 * All the supply is minted to the deployer.
 */
contract YellowToken is ERC20Permit {
	/**
	 * @dev Simple constructor, passing arguments to ERC20Permit and ERC20 constructors.
	 * Mints the supply to the deployer.
	 * @param name Name of the Token.
	 * @param symbol Symbol of the Token.
	 * @param supply Maximum supply of the Token.
	 */
	constructor(
		string memory name,
		string memory symbol,
		uint256 supply
	) ERC20Permit(name) ERC20(name, symbol) {
		_mint(msg.sender, supply);
	}

	/**
	 * @notice Return the number of decimals used to get its user representation.
	 * @dev Overrides ERC20 default value of 18;
	 * @return uint8 Number of decimals of Token.
	 */
	function decimals() public pure override returns (uint8) {
		return 8;
	}
}
