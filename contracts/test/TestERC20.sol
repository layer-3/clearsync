// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.22;

import '@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol';

contract TestERC20 is ERC20Capped {
	uint8 private immutable _decimals;

	constructor(
		string memory name_,
		string memory symbol_,
		uint8 decimals_,
		uint256 cap_
	) ERC20(name_, symbol_) ERC20Capped(cap_) {
		_decimals = decimals_;
	}

	function mint(address account, uint256 amount) external virtual {
		_mint(account, amount);
	}

	function burn(address account, uint256 amount) external virtual {
		_burn(account, amount);
	}

	function decimals() public view virtual override returns (uint8) {
		return _decimals;
	}
}
