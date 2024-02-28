// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.22;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';
import {INitroTypes} from '../nitro/interfaces/INitroTypes.sol';

/**
 * @title The IClearpool interface outlines the functionality of the Clearpool liquidity sharing pool smart contract.
 * @author Nikita Sazonov (nksazonov)
 * @notice The IClearpool interface outlines the ability for Users to deposit liquidity, claim the rewards and
 * withdraw the liquidity and rewards, and an ability for Brokers set reward rates for tokens, borrow those
 * tokens to conduct settlements.
 *
 * A User who sets the reward rate becomes a Broker, and can borrow tokens from the pool to conduct settlements.
 * When a User becomes a Broker, they can no longer withdraw their liquidity from the pool or update the reward rate.
 * The reward rate can only be updated << WHEN? >>
 */
interface IClearpool {
	struct PoolSettlement {
		INitroTypes.FixedPart fixedPart;
		uint48 settlementTurnNum;
		// TODO: change to simpler type
		Outcome.SingleAssetExit[] outcome;
	}

	event Deposited(address indexed user, address indexed token, uint256 amount);
	event Claimed(address indexed user, address indexed token, uint256 amount);
	event Withdrawn(address indexed user, address indexed token, uint256 amount);
	event RewardRateSet(address indexed user, address indexed token, uint256 rate);
	event SettlementExecuted(address indexed user, bytes settlement);

	/**
	 * @notice Deposit `amount` of `asset` into the pool.
	 * @dev Require `amount` to be greater than 0. Invokes `transferFrom`. Emit `Deposited` event.
	 * @param asset The address of the token to deposit.
	 * @param amount The amount of tokens to deposit.
	 */
	function deposit(address asset, uint256 amount) external;

	/**
	 * @notice Debit the reward to the user's internal balance and update the last claim timestamp.
	 * The reward is calculated as days since last claim * token balance / total tokens in pool * rewardRate.
	 * @dev Emit `Claimed` event.
	 * @param asset The address of the token to claim.
	 */
	function claim(address asset) external;

	/**
	 * @notice Withdraw `amount` of `asset` from the pool. This decreases the user's pool share,
	 * reducing the reward they are entitled to.
	 * @dev Require `amount` to be less than or equal to the user's balance. Invokes `transfer`. Emit `Withdrawn` event.
	 */
	function withdraw(address asset, uint256 amount) external;

	/**
	 * @notice Set the reward rate for `asset` to `rate` to be paid each 24h out of the setter's balance in the pool.
	 * @dev Require `rate` to be greater than 0. Emit `RewardRateSet` event.
	 * @param asset The address of the token to set the reward rate for.
	 * @param rate The reward rate to set.
	 */
	function setRewardRate(address asset, uint256 rate) external;

	/**
	 * @notice Execute a settlement with the given `settlement` and `signature`.
	 * @dev Require the settlement to be valid. Emit `SettlementExecuted` event.
	 * @param settlement The settlement to execute.
	 * @param sigs The signatures to validate the settlement.
	 */
	function execute(PoolSettlement calldata settlement, bytes[] calldata sigs) external;
}
