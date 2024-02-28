// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '../clearing/YellowAdjudicator.sol';
import '../interfaces/IClearpool.sol';

// TODO: change all reverts to custom errors
contract ClearPool is IClearpool {
	struct Account {
		uint256 balance;
		uint256 reward;
		uint256 lastClaimAt;
	}

	// TODO: make sure bidders have enough balance to cover the rewardRate
	// But where in the code do bidders use the liquidity?
	struct Bidder {
		uint256 rewardRate;
	}

	struct Asset {
		uint256 totalRewardRate;
		uint256 totalHolding;
		mapping(address => Bidder) bidders;
		mapping(address => Account) accounts;
	}

	mapping(address => Asset) public pools;

	mapping(bytes32 => bool) public isSettlementExecuted;

	YellowAdjudicator public adjudicator;

	uint256 public constant COOLDOWN_PERIOD = 24 hours;

	// TODO: OWNER can change adjudicator address
	constructor(YellowAdjudicator adjudicator_) {
		adjudicator = adjudicator_;
	}

	// Deposit function
	function deposit(address asset, uint256 amount) external {
		require(asset != address(0), 'Invalid asset address');
		require(amount > 0, 'Amount must be greater than 0');

		IERC20(asset).transferFrom(msg.sender, address(this), amount);
		pools[asset].accounts[msg.sender].balance += amount;
		pools[asset].totalHolding += amount;
		emit Deposited(msg.sender, asset, amount);
	}

	function claim(address asset) external {
		require(asset != address(0), 'Invalid asset address');
		if (pools[asset].accounts[msg.sender].lastClaimAt + COOLDOWN_PERIOD >= block.timestamp) {
			revert('Cooldown period not passed');
		}

		uint256 daysSinceLastClaim = (block.timestamp -
			pools[asset].accounts[msg.sender].lastClaimAt) / COOLDOWN_PERIOD;
		uint256 reward = (daysSinceLastClaim *
			pools[asset].totalRewardRate *
			pools[asset].accounts[msg.sender].balance) / pools[asset].totalHolding;

		// TODO: decrease bidders internal balance?
		pools[asset].accounts[msg.sender].reward += reward;
		pools[asset].accounts[msg.sender].lastClaimAt = block.timestamp;
		emit Claimed(msg.sender, asset, reward);
	}

	// Withdraw function
	function withdraw(address asset, uint256 amount) external {
		require(asset != address(0), 'Invalid asset address');
		require(amount > 0, 'Amount must be greater than 0');
		if (pools[asset].bidders[msg.sender].rewardRate > 0) {
			revert('Bidder can not withdraw');
		}

		require(amount <= pools[asset].accounts[msg.sender].balance, 'Insufficient balance');
		pools[asset].accounts[msg.sender].balance -= amount;
		pools[asset].totalHolding -= amount;
		IERC20(asset).transfer(msg.sender, amount);
		emit Withdrawn(msg.sender, asset, amount);
	}

	// Function to set the access rate
	function setRewardRate(address asset, uint256 rate) external {
		require(asset != address(0), 'Invalid asset address');
		require(rate > 0, 'Rate must be greater than 0');

		pools[asset].bidders[msg.sender].rewardRate = rate;
		pools[asset].totalRewardRate += rate;
		emit RewardRateSet(msg.sender, asset, rate);
	}

	// Execute settlement
	// TODO: add possibility to deposit to an escrow state channel
	function execute(PoolSettlement memory settlement, bytes[] calldata sigs) external {
		if (isSettlementExecuted[keccak256(abi.encode(settlement))]) {
			revert('Settlement already executed');
		}

		if (!_includesBidder(settlement.fixedPart.participants)) {
			revert('No bidder in participants');
		}

		bytes32 channelId = _getChannelId(settlement.fixedPart);
		uint48 turnNumRecorded;
		(turnNumRecorded, , ) = adjudicator.unpackStatus(channelId);

		if (turnNumRecorded < settlement.settlementTurnNum) {
			revert('Settlement turnNum not checkpointed');
		}

		// TODO:
		// check all parties have assets to be swapped
		// check that supplied settlement is signed by both supplied participants
		// Swap all balances internally given the outcome
	}

	function _includesBidder(address[] memory participants) internal pure returns (bool) {
		for (uint256 i = 0; i < participants.length; i++) {
			// TODO: implement
			if (participants[i] == address(0)) {
				return true;
			}
		}
		return false;
	}

	function _getChannelId(
		INitroTypes.FixedPart memory fixedPart
	) internal pure returns (bytes32 channelId) {
		channelId = keccak256(
			abi.encode(
				fixedPart.participants,
				fixedPart.channelNonce,
				fixedPart.appDefinition,
				fixedPart.challengeDuration
			)
		);
	}
}
