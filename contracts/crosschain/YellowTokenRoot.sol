// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@maticnetwork/fx-portal/contracts/tunnel/FxBaseRootTunnel.sol';
import '../YellowToken.sol';

contract YellowTokenRoot is FxBaseRootTunnel {
	bytes32 public constant DEPOSIT = keccak256('DEPOSIT');

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

	// TODO: swap deposit and withdraw between root and child
	// so that when depositing on child to root, tokens are locked on child and minted on root
	// and when withdrawing from root to child, tokens are burnt on root and unlocked on child
	function deposit(address user, uint256 amount, bytes memory data) public {
		// transfer from depositor to this contract
		rootYellowToken.transferFrom(
			msg.sender, // depositor
			address(this), // manager contract
			amount
		);

		// DEPOSIT, encode(rootToken, depositor, user, amount, extra data)
		bytes memory message = abi.encode(DEPOSIT, abi.encode(msg.sender, user, amount, data));
		_sendMessageToChild(message);
		emit FxDeposit(user, amount);
	}

	// exit processor
	function _processMessageFromChild(bytes memory data) internal override {
		(address to, uint256 amount) = abi.decode(data, (address, uint256));

		// check if current balance for token is less than amount,
		// mint remaining amount for this address
		uint256 balanceOf = rootYellowToken.balanceOf(address(this));
		if (balanceOf < amount) {
			rootYellowToken.mint(address(this), amount - balanceOf);
		}

		// approve token transfer
		rootYellowToken.approve(address(this), amount);

		// transfer from tokens to
		rootYellowToken.transferFrom(address(this), to, amount);
		emit FxWithdraw(to, amount);
	}
}
