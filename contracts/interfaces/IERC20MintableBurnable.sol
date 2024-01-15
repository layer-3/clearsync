// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.22;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

interface IERC20MintableBurnable is IERC20 {
	function mint(address to, uint256 amount) external;

	function burnFrom(address account, uint256 amount) external;
}
