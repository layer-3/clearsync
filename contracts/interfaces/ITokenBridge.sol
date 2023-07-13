// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface ITokenBridge {
	error BridgingUnauthorized(address sender, address token);
	error TokenAlreadySupported(address token);
	error TokenNotSupported(address token);
	error NoDstToken(address token, uint16 dstChainId);
	error InvalidToken(address token);

	event BridgeOut(
		uint16 chainTo,
		uint64 nonce,
		address token,
		address indexed sender,
		uint256 amount
	);
	event BridgeIn(
		uint16 chainFrom,
		uint64 nonce,
		address token,
		address indexed receiver,
		uint256 amount
	);

	function bridge(
		uint16 chainId,
		address token,
		address receiver,
		uint256 amount,
		address zroPaymentAddress,
		bytes calldata adapterParams
	) external payable;
}
