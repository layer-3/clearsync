// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@layerzerolabs/solidity-examples/contracts/lzApp/NonblockingLzApp.sol';

import '../interfaces/IERC20MintableBurnable.sol';

contract TokenBridge is NonblockingLzApp {
	event BridgeOut(uint16 chainTo, address indexed sender, uint256 amount);
	event BridgeIn(uint16 chainFrom, address indexed receiver, uint256 amount);

	IERC20MintableBurnable tokenContract;
	bool public immutable isRootBridge;

	constructor(
		address endpoint,
		address duckiesAddress,
		bool isRootBridge_
	) NonblockingLzApp(endpoint) {
		tokenContract = IERC20MintableBurnable(duckiesAddress);
		isRootBridge = isRootBridge_;
	}

	function _nonblockingLzReceive(
		uint16, // _srcChainId
		bytes memory, // _srcAddress
		uint64, // _nonce
		bytes memory _payload
	) internal override {
		(address receiver, uint256 amount) = abi.decode(_payload, (address, uint256));

		if (isRootBridge) {
			// NOTE: Bridge should have enough tokens as the only ability for token to appear on other chains is to be transferred to the bridge
			tokenContract.transfer(receiver, amount);
		} else {
			tokenContract.mint(receiver, amount);
		}
	}

	function estimateFee(
		uint16 _dstChainId,
		bool _useZro,
		address receiver,
		uint256 amount,
		bytes calldata _adapterParams
	) public view returns (uint nativeFee, uint zroFee) {
		return
			lzEndpoint.estimateFees(
				_dstChainId,
				address(this),
				abi.encode(receiver, amount),
				_useZro,
				_adapterParams
			);
	}

	// NOTE: chainIds are proprietary to LayerZero protocol and can be found on their docs
	function bridge(
		uint16 chainId,
		bool useZro,
		address receiver,
		uint256 amount
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
			useZro ? msg.sender : address(0), // zroPaymentAddress
			bytes(''), // adapterParams
			msg.value // nativeFee
		);
	}
}
