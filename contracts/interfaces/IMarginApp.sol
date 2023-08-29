// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@statechannels/nitro-protocol/contracts/interfaces/INitroTypes.sol';

/**
 * @notice Interface
 */
interface IMarginApp {
	enum MarginIndices {
		Leader,
		Follower
	}

	struct MarginCall {
		uint64 version;
		uint256[2] margin;
	}

	struct SignedMarginCall {
		MarginCall marginCall;
		INitroTypes.Signature[2] sigs;
	}

	struct Liability {
		address token;
		uint256 amount;
	}

	struct SettlementRequest {
		address[2] brokers;
		Liability[][2] settlements;
		uint64 version;
		uint64 expire;
		uint256 chainId;
		MarginCall adjustedMargin;
	}

	struct SignedSettlementRequest {
		SettlementRequest settlementRequest;
		INitroTypes.Signature[2] sigs;
	}

	/**
	 * @notice Settlement the Leader and Follower assets atomically.
	 * It will result in a valid margin adjustment in the channel.
	 * @param sSR SignedSettlementRequest struct.
	 * @param channelID Id of the channel settlement is being performed in.
	 */
	function settle(SignedSettlementRequest calldata sSR, bytes32 channelID) external payable;
}
