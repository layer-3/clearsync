// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@maticnetwork/fx-portal/contracts/tunnel/FxBaseChildTunnel.sol';
import '../YellowToken.sol';

contract YellowTokenChildTunnel is FxBaseChildTunnel {
	bytes32 public constant WITHDRAW = keccak256('WITHDRAW');
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
		rootYellowToken = rootYellowToken_;
	}

	function deposit(uint256 amount) public {
		_deposit(msg.sender, amount);
	}

	function depositTo(address to, uint256 amount) public {
		_deposit(to, amount);
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

		if (syncType == WITHDRAW) {
			_syncWithdraw(syncData);
		} else {
			revert('FxERC20ChildTunnel: INVALID_SYNC_TYPE');
		}
	}

	function _syncWithdraw(bytes memory syncData) internal {
		(address withdrawer, address receiver, uint256 amount, bytes memory depositData) = abi
			.decode(syncData, (address, address, uint256, bytes));

		// deposit tokens
		childYellowToken.transfer(receiver, amount);

		// call onTokenTransfer() on `to` with limit and ignore error
		if (_isContract(receiver)) {
			uint256 txGas = 2000000;
			bool success = false;
			bytes memory data = abi.encodeWithSignature(
				'onTokenTransfer(address,address,address,address,uint256,bytes)',
				rootYellowToken,
				childYellowToken,
				withdrawer,
				receiver,
				amount,
				depositData
			);
			// solhint-disable-next-line security/no--assembly
			assembly {
				success := call(txGas, receiver, 0, add(data, 0x20), mload(data), 0, 0)
			}
		}

		emit FxWithdraw(receiver, amount);
	}

	function _deposit(address to, uint256 amount) internal {
		// withdraw tokens
		childYellowToken.transferFrom(msg.sender, address(this), amount);

		// send message to root regarding token burn
		_sendMessageToRoot(abi.encode(to, amount));

		emit FxDeposit(to, amount);
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
