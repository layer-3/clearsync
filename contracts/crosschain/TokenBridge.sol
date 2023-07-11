// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@layerzerolabs/solidity-examples/contracts/lzApp/NonblockingLzApp.sol';

import '../interfaces/IERC20MintableBurnable.sol';

contract TokenBridge is NonblockingLzApp {
	event BridgeOut(uint16 chainTo, uint64 nonce, address indexed sender, uint256 amount);
	event BridgeIn(uint16 chainFrom, uint64 nonce, address indexed receiver, uint256 amount);

	IERC20MintableBurnable public tokenContract;
	bool public immutable isRootBridge;

	constructor(
		address endpoint,
		address tokenAddress,
		bool isRootBridge_
	) NonblockingLzApp(endpoint) {
		tokenContract = IERC20MintableBurnable(tokenAddress);
		isRootBridge = isRootBridge_;
	}

	function _nonblockingLzReceive(
		uint16 _srcChainId,
		bytes memory, // _srcAddress
		uint64 _nonce,
		bytes memory _payload
	) internal override {
		(address receiver, uint256 amount) = abi.decode(_payload, (address, uint256));

		if (isRootBridge) {
			// NOTE: Bridge should have enough tokens as the only ability for token to appear on other chains is to be transferred to the bridge
			tokenContract.transfer(receiver, amount);
		} else {
			tokenContract.mint(receiver, amount);
		}

		emit BridgeIn(_srcChainId, _nonce, receiver, amount);
	}

	function estimateFees(
		uint16 _dstChainId,
		address receiver,
		uint256 amount,
		bool payInZRO,
		bytes calldata _adapterParams
	) public view returns (uint nativeFee, uint zroFee) {
		return
			lzEndpoint.estimateFees(
				_dstChainId,
				address(this),
				abi.encode(receiver, amount),
				payInZRO,
				_adapterParams
			);
	}

	// NOTE: chainIds are proprietary to LayerZero protocol and can be found on their docs
	function bridge(
		uint16 chainId,
		address receiver,
		uint256 amount,
		address zroPaymentAddress,
		bytes calldata adapterParams
	) external payable {
		if (isRootBridge) {
			tokenContract.transferFrom(msg.sender, address(this), amount);
		} else {
			tokenContract.burnFrom(msg.sender, amount);
		}

		_lzSend(
			chainId, // chainId
			abi.encode(receiver, amount), // payload
			payable(msg.sender), // refundAddress
			zroPaymentAddress, // zroPaymentAddress
			adapterParams, // adapterParams
			msg.value // nativeFee
		);

		emit BridgeOut(
			chainId,
			lzEndpoint.getOutboundNonce(chainId, address(this)),
			msg.sender,
			amount
		);
	}
}
