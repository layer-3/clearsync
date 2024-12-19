// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

import {StrictTurnTaking} from '../nitro/libraries/signature-logic/StrictTurnTaking.sol';
import {Consensus} from '../nitro/libraries/signature-logic/Consensus.sol';
import {IForceMoveApp} from '../nitro/interfaces/IForceMoveApp.sol';
import {ITradingTypes} from '../interfaces/ITradingTypes.sol';

contract TradingApp is IForceMoveApp {
	// TODO: add errors

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
		if (candTurnNum % 2 == 0 && proof.length > 0) {
			Consensus.requireConsensus(fixedPart, proof, candidate);
			// Check the settlement data structure validity
			ITradingTypes.Settlement memory settlement = abi.decode(
				candidateData,
				(ITradingTypes.Settlement)
			);

			verifyProofForSettlement(settlement, proof);
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
			ITradingTypes.Order memory prevOrder = abi.decode(
				proof0.appData,
				(ITradingTypes.Order)
			);
			ITradingTypes.OrderResponse memory prevOrderResponse = abi.decode(
				proof1.appData,
				(ITradingTypes.OrderResponse)
			);
			if (prevOrderResponse.responseType == ITradingTypes.OrderResponseType.ACCEPT) {
				require(
					prevOrderResponse.orderID == prevOrder.orderID,
					'orderResponse.orderID != prevOrder.orderID, candidate is order'
				);
			}
			// NOTE: used just to check the data structure validity
			ITradingTypes.Order memory _candOrder = abi.decode(
				candidateData,
				(ITradingTypes.Order)
			);
			return (true, '');
		}

		// orderResponse
		// NOTE: used just to check the data structure validity
		ITradingTypes.OrderResponse memory _prevOrderResponse = abi.decode(
			proof0.appData,
			(ITradingTypes.OrderResponse)
		);

		ITradingTypes.Order memory order = abi.decode(proof1.appData, (ITradingTypes.Order));
		ITradingTypes.OrderResponse memory orderResponse = abi.decode(
			candidateData,
			(ITradingTypes.OrderResponse)
		);
		if (orderResponse.responseType == ITradingTypes.OrderResponseType.ACCEPT) {
			require(
				orderResponse.orderID == order.orderID,
				'orderResponse.orderID != order.orderID, candidate is orderResponse'
			);
		}
		return (true, '');
	}

	function verifyProofForSettlement(
		ITradingTypes.Settlement memory settlement,
		RecoveredVariablePart[] calldata proof
	) internal pure {
		if (proof.length == 1) {
			VariablePart memory proof = proof[0].variablePart;
			ITradingTypes.Order memory order = abi.decode(proof.appData, (ITradingTypes.Order));
			return;
		}

		bytes32[] memory orderIDs = new bytes32[](proof.length);
		for (uint256 i = 0; i < proof.length - 1; i++) {
			VariablePart memory currProof = proof[i].variablePart;
			VariablePart memory nextProof = proof[i + 1].variablePart;

			require(currProof.turnNum + 1 == nextProof.turnNum, 'turns are not consecutive');

			// Verify validity of orders and responses
			if (i % 2 == 0) {
				// If current proof contains an order,
				// then the next one must contain a response
				// with the same order ID
				ITradingTypes.Order memory order = abi.decode(
					currProof.appData,
					(ITradingTypes.Order)
				);
				ITradingTypes.OrderResponse memory orderResponse = abi.decode(
					nextProof.appData,
					(ITradingTypes.OrderResponse)
				);
				require(
					orderResponse.orderID == order.orderID,
					'orderResponse.orderID != order.orderID'
				);
				orderIDs[i] = order.orderID;
			} else {
				// If current proof contains a response,
				// then the next one must be an order
				// with a different order ID, since they are not related
				ITradingTypes.OrderResponse memory orderResponse = abi.decode(
					currProof.appData,
					(ITradingTypes.OrderResponse)
				);
				ITradingTypes.Order memory order = abi.decode(
					nextProof.appData,
					(ITradingTypes.Order)
				);
				require(
					order.orderID != orderResponse.orderID,
					'order.orderID == orderResponse.orderID'
				);
				orderIDs[i] = order.orderID;
			}
		}

		bytes32 proofHash = keccak256(abi.encode(orderIDs));
		require(proofHash == settlement.proofHash, 'proof has been tampered with');
	}
}
