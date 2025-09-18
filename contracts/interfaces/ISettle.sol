// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {INitroTypes} from '../nitro/interfaces/INitroTypes.sol';

/**
 * @title ISettle
 * @notice Interface for a contract that allows users to settle a channel.
 */
interface ISettle {
	// ========== Events ==========

	/**
	 * @notice Emitted when a channel is settled.
	 * @param trader The address of the trader.
	 * @param broker The address of the broker.
	 * @param channelId The ID of the channel.
	 */
	event Settled(
		address indexed trader,
		address indexed broker,
		bytes32 indexed channelId,
		bytes32 settlementId
	);

	// ========== Errors ==========

	error InvalidStateTransition(string reason);

	// ========== Functions ==========

	/**
	 * @notice Settle a channel.
	 * @param fixedPart The fixed part of the state.
	 * @param proof The proof of the state.
	 * @param candidate The candidate state.
	 */
	function settle(
		INitroTypes.FixedPart calldata fixedPart,
		INitroTypes.RecoveredVariablePart[] calldata proof,
		INitroTypes.RecoveredVariablePart calldata candidate
	) external;
}
