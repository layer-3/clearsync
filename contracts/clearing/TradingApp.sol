// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import '../nitro/interfaces/IForceMoveApp.sol';
import '../nitro/libraries/NitroUtils.sol';
import '../nitro/interfaces/INitroTypes.sol';
import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

/*
    - both Order and Trade contain sigs of user and broker respectively
    - Order and Trade are coupled by some id (that may be calculated based on the order itself)
    - State invariant: for each Order there is a respective Trade, and both are valid (not 0 amounts, etc)
    - State transition rules:
        - amount of Order-Trade pairs can increase without restriction
        - amount of Order-Trade pairs can decrease only after the settlement state
    - settlement state is a state with `settlementData` not
    - to settle using the TradingVault, settlement state must be signed by the Broker and *maybe* by the user
 */

// TODO: optimize for storage slots
interface TradingStructs {
	enum Side {
		BUY,
		SELL
	}

	struct Order {
		Side side;
		address token;
		uint256 amount;
		uint256 ts; // timestamp in seconds
		bytes signature; // user signature
	}

	// NOTE: there is no need in `tradeId` as it is not checked in any way.
	// If you want to signal Trade settlement, then use settlementId in the event instead
	struct Trade {
		bytes32 orderId; // keccak256(abi.encode(Order))
		uint256 amount; // amount executed
		uint256 ts; // timestamp in seconds
		bytes signature; // broker signature
	}

	enum FundingLocation {
		/// @dev funds are taken from the TradingVault account's balance
		TradingVault,
		/// @dev funds are pulled from the account's token balance
		Address
	}

	struct CounterpartiesFundingLocations {
		FundingLocation[] user;
		FundingLocation[] broker;
	}

	struct OrderTradePair {
		Order order;
		Trade trade;
	}

	struct Settlement {
		bytes32 id; // keccak256(abi.encode(OrderTradePair[]));
		CounterpartiesFundingLocations sourceFLs;
		CounterpartiesFundingLocations destinationFLs;
	}

	// is encoded into the `state` field of the variablePart
	struct State {
		OrderTradePair[] pairs;
		bytes settlementData; // 0x if state is not settlement, encoded `Settlement` if it is
	}
}

interface ISettle {
	function settle(
		INitroTypes.FixedPart calldata fixedPart,
		INitroTypes.RecoveredVariablePart[] calldata proof,
		INitroTypes.RecoveredVariablePart calldata candidate
	) external;
}

contract TradingApp is IForceMoveApp {
	function stateIsSupported(
		FixedPart calldata fixedPart,
		RecoveredVariablePart[] calldata proof,
		RecoveredVariablePart calldata candidate
	) external pure override returns (bool, string memory) {}
}
