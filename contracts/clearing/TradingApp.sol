// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

import {StrictTurnTaking} from '../nitro/libraries/signature-logic/StrictTurnTaking.sol';
import {Consensus} from '../nitro/libraries/signature-logic/Consensus.sol';
import {IForceMoveApp} from '../nitro/interfaces/IForceMoveApp.sol';
import {ITradingTypes} from '../interfaces/ITradingTypes.sol';
import {NitroUtils} from '../nitro/libraries/NitroUtils.sol';

contract TradingApp is IForceMoveApp {
	// TODO: add custom errors after contract logic is finalized

	function stateIsSupported(
		FixedPart calldata fixedPart,
		RecoveredVariablePart[] calldata proof,
		RecoveredVariablePart calldata candidate
	) external pure override returns (bool, string memory) {
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
		uint8 signaturesNum = NitroUtils.getClaimedSignersNum(candidate.signedBy);
		if (
			candTurnNum % 2 == 0 /* is either order or settlement */ &&
			signaturesNum == 2 /* is settlement */ &&
			proof.length >= 2 /* contains at least one order+response pair */ &&
			proof.length % 2 == 0 /* contains full pairs only, no dangling values */
		) {
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
		require(signaturesNum == 1, 'signaturesNum != 1');
		require(proof.length == 2, 'proof.length != 2');
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
		bytes32[] memory proofDataHashes = new bytes32[](proof.length);
		uint256 prevTurnNum = 1; // postfund state
		for (uint256 i = 0; i < proof.length - 1; i += 2) {
			VariablePart memory currProof = proof[i].variablePart;
			VariablePart memory nextProof = proof[i + 1].variablePart;

			require(prevTurnNum + 1 == currProof.turnNum, 'turns are not consecutive');
			require(currProof.turnNum + 1 == nextProof.turnNum, 'turns are not consecutive');

			// Verify validity of orders and responses
			ITradingTypes.Order memory order = abi.decode(currProof.appData, (ITradingTypes.Order));
			ITradingTypes.OrderResponse memory orderResponse = abi.decode(
				nextProof.appData,
				(ITradingTypes.OrderResponse)
			);

			// If current proof contains an order,
			// then the next one must contain a response
			// with the same order ID
			require(orderResponse.orderID == order.orderID, 'order and response IDs do not match');

			proofDataHashes[i] = keccak256(currProof.appData);
			proofDataHashes[i + 1] = keccak256(nextProof.appData);
			prevTurnNum = nextProof.turnNum;
		}

		bytes32 ordersChecksum = keccak256(abi.encode(proofDataHashes));
		require(ordersChecksum == settlement.ordersChecksum, 'proof has been tampered with');
	}
}
