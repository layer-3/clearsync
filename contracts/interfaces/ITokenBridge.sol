// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface ITokenBridge {
	event BridgeOut(uint16 chainTo, uint64 nonce, address indexed sender, uint256 amount);
	event BridgeIn(uint16 chainFrom, uint64 nonce, address indexed receiver, uint256 amount);

	function bridge(
		uint16 chainId,
		address receiver,
		uint256 amount,
		address zroPaymentAddress,
		bytes calldata adapterParams
	) external payable;
}
