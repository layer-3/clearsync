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

	struct SwapCall {
		address[2] brokers;
		Liability[][2] swaps;
		uint64 version;
		uint64 expire;
		uint256 chainId;
		MarginCall adjustedMargin;
	}

	struct SignedSwapCall {
		SwapCall swapCall;
		INitroTypes.Signature[2] sigs;
	}

	/**
	 * @notice Swap the Leader and Follower assets atomically.
	 * It will result in a valid margin adjustment in the channel.
	 * @param sSC SignedSwapCall struct.
	 * @param channelID Id of the channel swap is being performed in.
	 */
	function swap(SignedSwapCall calldata sSC, bytes32 channelID) external payable;
}
