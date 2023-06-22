// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@maticnetwork/fx-portal/contracts/tunnel/FxBaseChildTunnel.sol';
import '../YellowToken.sol';

contract YellowTokenChildTunnel is FxBaseChildTunnel {
	bytes32 public constant DEPOSIT = keccak256('DEPOSIT');
	YellowToken public childYellowToken;
	address public rootYellowToken;

	event FxDeposit(address indexed user, uint256 amount);
	event FxWithdraw(address indexed user, uint256 amount);

	constructor(
		address childYellowToken_,
		address rootYellowToken_,
		address fxChild_
	) FxBaseChildTunnel(fxChild_) {
		require(
			childYellowToken_ != address(0) &&
				rootYellowToken_ != address(0) &&
				fxChild_ != address(0),
			'YellowTokenChild: INVALID_ADDRESS_IN_CONSTRUCTOR'
		);

		childYellowToken = YellowToken(childYellowToken_);
		require(
			childYellowToken.hasRole(childYellowToken.MINTER_ROLE(), address(this)),
			'YellowTokenChild: INVALID_MINTER_ROLE'
		);

		rootYellowToken = rootYellowToken_;
	}

	function withdraw(uint256 amount) public {
		_withdraw(msg.sender, amount);
	}

	function withdrawTo(address receiver, uint256 amount) public {
		_withdraw(receiver, amount);
	}

	//
	// Internal methods
	//

	function _processMessageFromRoot(
		uint256 /* stateId */,
		address sender,
		bytes memory data
	) internal override validateSender(sender) {
		// decode incoming data
		(bytes32 syncType, bytes memory syncData) = abi.decode(data, (bytes32, bytes));

		if (syncType == DEPOSIT) {
			_syncDeposit(syncData);
		} else {
			revert('FxERC20ChildTunnel: INVALID_SYNC_TYPE');
		}
	}

	function _syncDeposit(bytes memory syncData) internal {
		(address depositor, address to, uint256 amount, bytes memory depositData) = abi.decode(
			syncData,
			(address, address, uint256, bytes)
		);

		// deposit tokens
		childYellowToken.mint(to, amount);

		// call onTokenTransfer() on `to` with limit and ignore error
		if (_isContract(to)) {
			uint256 txGas = 2000000;
			bool success = false;
			bytes memory data = abi.encodeWithSignature(
				'onTokenTransfer(address,address,address,address,uint256,bytes)',
				rootYellowToken,
				childYellowToken,
				depositor,
				to,
				amount,
				depositData
			);
			// solhint-disable-next-line security/no--assembly
			assembly {
				success := call(txGas, to, 0, add(data, 0x20), mload(data), 0, 0)
			}
		}

		emit FxDeposit(to, amount);
	}

	function _withdraw(address receiver, uint256 amount) internal {
		// withdraw tokens
		childYellowToken.burnFrom(msg.sender, amount);

		// send message to root regarding token burn
		_sendMessageToRoot(abi.encode(receiver, amount));

		emit FxWithdraw(receiver, amount);
	}

	// check if address is contract
	function _isContract(address _addr) private view returns (bool) {
		uint32 size;
		assembly {
			size := extcodesize(_addr)
		}
		return (size > 0);
	}
}
