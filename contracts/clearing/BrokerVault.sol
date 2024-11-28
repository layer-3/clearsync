// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import {IVault2} from '../interfaces/IVault2.sol';
import {TradingApp, ISettle} from './TradingApp.sol';
import {Ownable2Step} from '@openzeppelin/contracts/access/Ownable2Step.sol';
import {ReentrancyGuard} from '@openzeppelin/contracts/utils/ReentrancyGuard.sol';
import {NitroUtils} from '../nitro/libraries/NitroUtils.sol';

contract BrokerVault is IVault2, ISettle, Ownable2Step, ReentrancyGuard {
	/// @dev Using SafeERC20 to support non fully ERC20-compliant tokens,
	/// that may not return a boolean value on success.
	using SafeERC20 for IERC20;

	mapping(address token => uint256 balance) internal _balances;
	TradingApp internal tradingApp;
	mapping(uint256 channel_id => bool done) internal performedSettlements;
	address public _broker;

	// ====== Constructor ======

	constructor(address owner, address broker_) Ownable(owner) {
		broker = broker_;
	}

	// ---------- View functions ----------

	function balanceOf(address token) external view returns (uint256) {
		return _balances[token];
	}

	function balancesOfTokens(address[] calldata tokens) external view returns (uint256[] memory) {
		uint256[] memory balances = new uint256[](tokens.length);
		for (uint256 i = 0; i < tokens.length; i++) {
			balances[i] = _balances[tokens[i]];
		}
		return balances;
	}

	// ---------- Owner functions ----------

	function setBroker(address broker_) external onlyOwner {
		broker = broker_;
	}

	function setAuthorizer(IAuthorizeV2 newAuthorizer) external onlyOwner {
		if (address(newAuthorizer) == address(0)) {
			revert InvalidAddress();
		}

		authorizer = newAuthorizer;
		emit AuthorizerChanged(newAuthorizer);
	}

	// ---------- Write functions ----------

	function deposit(address token, uint256 amount, address to) external payable {
		require(msg.value == 0, 'IncorrectValue');
		require(token != address(0), 'InvalidAddress');

		IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
		_balances[token] += amount;

		emit Deposited(msg.sender, token, amount);
	}

	function withdraw(address token, uint256 amount, address to) external {
		require(token != address(0), 'InvalidAddress');

		if (_balances[token] < amount) {
			revert InsufficientBalance(token, amount, _balances[token]);
		}

		_balances[token] -= amount;
		IERC20(token).safeTransfer(to, amount);

		emit Withdrawn(msg.sender, token, amount);
	}

	function settle(
		INitroTypes.FixedPart calldata fixedPart,
		INitroTypes.RecoveredVariablePart[] calldata proof,
		INitroTypes.RecoveredVariablePart calldata candidate
	) external {
		uint256 channelId = NitroUtils.getChannelId(fixedPart);

		require(!performedSettlements[channelId], 'Settlement already performed');
		require(fixedPart.participants[1] == broker, 'Broker is not a participant');
		require(
			tradingApp.isStateTransitionValid(fixedPart, proof, candidate),
			'Invalid state transition'
		);

		ITradingStructs.Settlement memory settlement = abi.decode(
			candidate.variablePart.appData,
			(ITradingStructs.Settlement)
		);

		for (uint256 i = 0; i < settlement.toTrader.length; i++) {
			address token = settlement.toTrader[i].asset;
			uint256 amount = settlement.toTrader[i].amount;
			require(_balances[token] >= amount, 'Insufficient balance');
			_balances[token] -= amount;
			IERC20(token).safeTransfer(fixedPart.participants[0], amount);
		}

		for (uint256 i = 0; i < settlement.toBroker.length; i++) {
			address token = settlement.toBroker[i].asset;
			uint256 amount = settlement.toBroker[i].amount;
			_balances[token] += amount;
		}
		performedSettlements[channelId] = true;
	}
}
