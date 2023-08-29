// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@statechannels/nitro-protocol/contracts/interfaces/INitroTypes.sol';
import './IMarginTypes.sol';

/**
 * @notice Interface
 */
interface ISettlementTypes {
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
		IMarginTypes.MarginCall adjustedMargin;
	}

	struct SignedSettlementRequest {
		SettlementRequest settlementRequest;
		INitroTypes.Signature[2] sigs;
	}
}
