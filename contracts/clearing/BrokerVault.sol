// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import {IVault2} from '../interfaces/IVault2.sol';
import {TradingApp, ISettle} from './TradingApp.sol';
import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {Ownable2Step} from '@openzeppelin/contracts/access/Ownable2Step.sol';
import {ReentrancyGuard} from '@openzeppelin/contracts/utils/ReentrancyGuard.sol';
import {NitroUtils} from '../nitro/libraries/NitroUtils.sol';

contract BrokerVault is IVault2, ISettle, Ownable2Step, ReentrancyGuard {
	/// @dev Using SafeERC20 to support non fully ERC20-compliant tokens,
	/// that may not return a boolean value on success.
	using SafeERC20 for IERC20;

	// ====== Variables ======

	address public broker;
	TradingApp public tradingApp;
	mapping(bytes32 channelId => bool done) public performedSettlements;

	mapping(address token => uint256 balance) internal _balances;

	// ====== Errors ======

	error InsufficientBalance(address token, uint256 required, uint256 available);
	error InvalidAddress();
	error InvalidAmount(uint256 amount);
	error SettlementAlreadyPerformed(bytes32 channelId);
	error BrokerNotParticipant(address actual, address expectedBroker);

	// ====== Constructor ======

	constructor(address owner, address broker_, TradingApp tradingApp_) Ownable(owner) {
		broker = broker_;
		tradingApp = tradingApp_;
	}

	// ---------- View functions ----------

	function balanceOf(address user, address token) external view returns (uint256) {
		if (user != broker) {
			return 0;
		}
		return _balances[token];
	}

	function balancesOfTokens(
		address user,
		address[] calldata tokens
	) external view returns (uint256[] memory) {
		if (user != broker) {
			return new uint256[](tokens.length);
		}

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

	// ---------- Write functions ----------

	function deposit(address token, uint256 amount, address to) external payable {
		require(msg.value != 0, InvalidAmount(msg.value));
		require(token != address(0), InvalidAddress());
		require(to != broker, InvalidAddress());

		IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
		_balances[token] += amount;

		emit Deposited(to, token, amount);
	}

	function withdraw(address token, uint256 amount, address to) external {
		require(token != address(0), InvalidAddress());
		require(_balances[token] >= amount, InsufficientBalance(token, amount, _balances[token]));

		_balances[token] -= amount;
		IERC20(token).safeTransfer(to, amount);

		emit Withdrawn(to, token, amount);
	}

	function settle(
		INitroTypes.FixedPart calldata fixedPart,
		INitroTypes.RecoveredVariablePart[] calldata proof,
		INitroTypes.RecoveredVariablePart calldata candidate
	) external {
		uint256 channelId = NitroUtils.getChannelId(fixedPart);
		require(!performedSettlements[channelId], SettlementAlreadyPerformed(channelId));

		require(
			fixedPart.participants[0] == broker,
			BrokerNotParticipant(fixedPart.participants[1], broker)
		);
		address trader = fixedPart.participants[0];

		(bool isStateValid, string memory reason) = tradingApp.isStateTransitionValid(
			fixedPart,
			proof,
			candidate
		);
		require(isStateValid, InvalidStateTransition(reason));

		ITradingStructs.Settlement memory settlement = abi.decode(
			candidate.variablePart.appData,
			(ITradingStructs.Settlement)
		);

		for (uint256 i = 0; i < settlement.toTrader.length; i++) {
			address token = settlement.toTrader[i].asset;
			uint256 amount = settlement.toTrader[i].amount;
			require(
				_balances[token] >= amount,
				InsufficientBalance(token, amount, _balances[token])
			);
			IERC20(token).safeTransfer(trader, amount);
			_balances[token] -= amount;
		}

		for (uint256 i = 0; i < settlement.toBroker.length; i++) {
			address token = settlement.toBroker[i].asset;
			uint256 amount = settlement.toBroker[i].amount;
			IERC20(token).safeTransferFrom(trader, broker, amount);
			_balances[token] += amount;
		}

		performedSettlements[channelId] = true;
		emit Settled(trader, broker, channelId);
	}
}
