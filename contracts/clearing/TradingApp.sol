// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

import {StrictTurnTaking} from '../nitro/libraries/signature-logic/StrictTurnTaking.sol';
import {Consensus} from '../nitro/libraries/signature-logic/Consensus.sol';
import {IForceMoveApp} from '../nitro/interfaces/IForceMoveApp.sol';
import {NitroUtils} from '../nitro/libraries/NitroUtils.sol';
import {INitroTypes} from '../nitro/interfaces/INitroTypes.sol';

interface ITradingStructs {
	struct Order {
		bytes32 orderID; // tradeID
	}

	enum OrderResponseType {
		ACCEPT,
		REJECT
	}

	struct OrderResponse {
		OrderResponseType responseType;
		bytes32 orderID; // orderID making the trade
	}

	struct AssetAndAmount {
		address asset;
		uint256 amount;
	}

	struct Settlement {
		AssetAndAmount[] toTrader;
		AssetAndAmount[] toBroker;
	}
}

// FIXME: should Vault support multiple brokers?
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
	) external pure override returns (bool, string memory) {
		// FIXME: does the Broker deposit to the Adjudicator?
		// turn nums:
		// 0 - prefund
		// 1 - postfund
		// 2 - order
		// 2n+1 - order response
		// 2n - order or settlement

		uint48 candTurnNum = candidate.variablePart.turnNum;

		// prefund or postfund
		if (candTurnNum == 0 || candTurnNum == 1) {
			Consensus.requireConsensus(fixedPart, proof, candidate);
			return (true, '');
		}

		bytes memory candidateData = candidate.variablePart.appData;
		// settlement
		// TODO: unsure whether we should check the proof when consensus is reached
		if (candTurnNum % 2 == 0 && proof.length == 0) {
			Consensus.requireConsensus(fixedPart, proof, candidate);
			// NOTE: used just to check the data structure validity
			ITradingStructs.Settlement memory _unused = abi.decode(
				candidateData,
				(ITradingStructs.Settlement)
			);
			return (true, '');
		}

		// participant 0 signs even turns
		// participant 1 signs odd turns
		StrictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate);
		require(proof.length == 2, 'proof.length < 2');
		(VariablePart memory proof0, VariablePart memory proof1) = (
			proof[0].variablePart,
			proof[1].variablePart
		);
		require(proof0.turnNum == candTurnNum - 2, 'proof0.turnNum != candTurnNum - 1');
		require(proof1.turnNum == candTurnNum - 1, 'proof1.turnNum != candTurnNum - 1');

		// order
		if (candTurnNum % 2 == 0) {
			ITradingStructs.Order memory prevOrder = abi.decode(
				proof0.appData,
				(ITradingStructs.Order)
			);
			ITradingStructs.OrderResponse memory prevOrderResponse = abi.decode(
				proof1.appData,
				(ITradingStructs.OrderResponse)
			);
			if (prevOrderResponse.responseType == ITradingStructs.OrderResponseType.ACCEPT) {
				require(
					prevOrderResponse.orderID == prevOrder.orderID,
					'orderResponse.orderID != prevOrder.orderID, candidate is order'
				);
			}
			// NOTE: used just to check the data structure validity
			ITradingStructs.Order memory _candOrder = abi.decode(
				candidateData,
				(ITradingStructs.Order)
			);
			return (true, '');
		}

		// orderResponse
		// NOTE: used just to check the data structure validity
		ITradingStructs.OrderResponse memory _prevOrderResponse = abi.decode(
			proof0.appData,
			(ITradingStructs.OrderResponse)
		);

		ITradingStructs.Order memory order = abi.decode(proof1.appData, (ITradingStructs.Order));
		ITradingStructs.OrderResponse memory orderResponse = abi.decode(
			candidateData,
			(ITradingStructs.OrderResponse)
		);
		if (orderResponse.responseType == ITradingStructs.OrderResponseType.ACCEPT) {
			require(
				orderResponse.orderID == order.orderID,
				'orderResponse.orderID != order.orderID, candidate is orderResponse'
			);
		}
		return (true, '');
	}
}
