// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@maticnetwork/fx-portal/contracts/tunnel/FxBaseRootTunnel.sol';
import '../YellowToken.sol';

contract YellowTokenRootTunnel is FxBaseRootTunnel {
	bytes32 public constant WITHDRAW = keccak256('WITHDRAW');

	address public childYellowToken;
	YellowToken public rootYellowToken;

	event FxDeposit(address indexed user, uint256 amount);
	event FxWithdraw(address indexed user, uint256 amount);

	constructor(
		address childYellowToken_,
		address rootYellowToken_,
		address checkpointManager_,
		address fxRoot_
	) FxBaseRootTunnel(checkpointManager_, fxRoot_) {
		require(
			childYellowToken_ != address(0) &&
				rootYellowToken_ != address(0) &&
				checkpointManager_ != address(0) &&
				fxRoot_ != address(0),
			'YellowTokenRoot: INVALID_ADDRESS_IN_CONSTRUCTOR'
		);

		childYellowToken = childYellowToken_;
		rootYellowToken = YellowToken(rootYellowToken_);
	}

	function withdraw(uint256 amount, bytes memory data) public {
		_withdraw(msg.sender, amount, data);
	}

	function withdrawTo(address receiver, uint256 amount, bytes memory data) public {
		_withdraw(receiver, amount, data);
	}

	function _withdraw(address receiver, uint256 amount, bytes memory data) internal {
		// burn from withdrawer
		rootYellowToken.burnFrom(receiver, amount);

		// WITHDRAW, encode(withdrawer, receiver, amount, extra data)
		bytes memory message = abi.encode(WITHDRAW, abi.encode(msg.sender, receiver, amount, data));
		_sendMessageToChild(message);
		emit FxWithdraw(receiver, amount);
	}

	// deposit processor
	function _processMessageFromChild(bytes memory data) internal override {
		(address to, uint256 amount) = abi.decode(data, (address, uint256));

		// mint tokens
		rootYellowToken.mint(to, amount);

		emit FxDeposit(to, amount);
	}
}
