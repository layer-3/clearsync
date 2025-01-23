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
		// TODO: refactor by extracting logic into several functions
		// TODO: do we want to continue operating this channel after settlement? If so, we need to support such state change. Changes to liquidation validation are required.
		// turn nums:
		// 0 - prefund
		// 1 - postfund
		// 2 - order
		// 2n+1 - order response
		// 2n+2 - order / settlement / liquidation

		require(fixedPart.participants.length == 2, 'invalid number of participants, expected 2');

		uint48 candTurnNum = candidate.variablePart.turnNum;

		// prefund or postfund
		if (candTurnNum == 0 || candTurnNum == 1) {
			// no proof, candidate consensus
			Consensus.requireConsensus(fixedPart, proof, candidate);
			return (true, '');
		}

		bytes memory candidateData = candidate.variablePart.appData;

		// order or orderResponse
		if (proof.length == 1) {
			_requireSingleAllocation(proof[0].variablePart.outcome);
			_requireSingleAllocation(candidate.variablePart.outcome);
			_requireNoAllocationAmountChange(
				proof[0].variablePart.outcome,
				candidate.variablePart.outcome
			);

			// first order
			if (candidate.variablePart.turnNum == 2) {
				require(
					proof[0].variablePart.turnNum == 1,
					'invalid proof turn num on first order'
				);
				// check consensus of postfund
				Consensus.requireConsensus(fixedPart, new RecoveredVariablePart[](0), proof[0]);
				StrictTurnTaking.isSignedByMover(fixedPart, candidate);
				// NOTE: used just to check the data structure validity
				ITradingTypes.Order memory _candOrder = abi.decode(
					candidateData,
					(ITradingTypes.Order)
				);
				return (true, '');
			}

			// participant 0 signs even turns
			// participant 1 signs odd turns
			StrictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate);
			VariablePart memory proof0 = proof[0].variablePart;

			// order
			if (candTurnNum % 2 == 0) {
				// NOTE: used just to check the data structure validity
				ITradingTypes.OrderResponse memory _prevOrderResponse = abi.decode(
					proof0.appData,
					(ITradingTypes.OrderResponse)
				);
				// NOTE: used just to check the data structure validity
				ITradingTypes.Order memory _candOrder = abi.decode(
					candidateData,
					(ITradingTypes.Order)
				);
				return (true, '');
			}

			// orderResponse
			ITradingTypes.Order memory order = abi.decode(proof0.appData, (ITradingTypes.Order));
			ITradingTypes.OrderResponse memory orderResponse = abi.decode(
				candidateData,
				(ITradingTypes.OrderResponse)
			);
			if (orderResponse.responseType == ITradingTypes.OrderResponseType.ACCEPT) {
				require(orderResponse.orderID == order.orderID, 'order and response IDs mismatch');
			}
			return (true, '');
		}
		// settlement or liquidation
		else if (proof.length >= 2) {
			// both can only happen after an OrderResponse
			require(candTurnNum % 2 == 0, 'invalid candidate turn num');

			// liquidation
			if (NitroUtils.getClaimedSignersNum(candidate.signedBy) == 1) {
				require(proof.length == 2, 'liquidation proof too long');
				// check proof[0] - order
				StrictTurnTaking.isSignedByMover(fixedPart, proof[0]);
				ITradingTypes.Order memory order = abi.decode(
					proof[0].variablePart.appData,
					(ITradingTypes.Order)
				);

				// check proof[1] - ACCEPT orderResponse
				StrictTurnTaking.isSignedByMover(fixedPart, proof[1]);
				require(
					proof[1].variablePart.turnNum == proof[0].variablePart.turnNum + 1,
					'turns are not consecutive'
				);
				ITradingTypes.OrderResponse memory orderResponse = abi.decode(
					proof[1].variablePart.appData,
					(ITradingTypes.OrderResponse)
				);
				require(orderResponse.orderID == order.orderID, 'order and response IDs mismatch');
				require(
					orderResponse.responseType == ITradingTypes.OrderResponseType.ACCEPT,
					'order not accepted'
				);

				// check candidate - liquidation state
				require(
					// NOTE: liquidation can be not a direct successor of the ACCEPT orderResponse to allow
					// for liquidation after REJECT orderResponse
					candidate.variablePart.turnNum > proof[1].variablePart.turnNum,
					'invalid liquidation turn num'
				);
				require(
					// trader is mover #0, broker is mover #1
					NitroUtils.isClaimedSignedOnlyBy(candidate.signedBy, 1),
					'not signed by broker'
				);

				// outcomes
				_requireSingleAllocation(proof[0].variablePart.outcome);
				_requireSingleAllocation(proof[1].variablePart.outcome);
				_requireNoAllocationAmountChange(
					proof[0].variablePart.outcome,
					proof[1].variablePart.outcome
				);
				_requireValidFundsSplit(
					proof[1].variablePart.outcome,
					candidate.variablePart.outcome
				);
				return (true, '');
			}
			// settlement
			else {
				require(proof.length % 2 == 0, 'settlement proof contains dangling values');
				// check consensus of candidate
				Consensus.requireConsensus(fixedPart, new RecoveredVariablePart[](0), candidate);
				// Check the settlement data structure validity
				ITradingTypes.Settlement memory settlement = abi.decode(
					candidateData,
					(ITradingTypes.Settlement)
				);
				_verifyProofForSettlement(fixedPart, settlement, proof);
				return (true, '');
			}
		}

		revert('invalid proof length');
	}

	function _requireSingleAllocation(Outcome.SingleAssetExit[] memory outcome) internal pure {
		require(outcome.length == 1, 'not 1 asset');
		require(outcome[0].allocations.length == 1, 'not 1 allocation');
	}

	function _requireNoAllocationAmountChange(
		Outcome.SingleAssetExit[] memory prevOutcome,
		Outcome.SingleAssetExit[] memory nextOutcome
	) internal pure {
		require(
			prevOutcome[0].allocations[0].destination == nextOutcome[0].allocations[0].destination,
			'destination changed in allocation'
		);
		require(
			prevOutcome[0].allocations[0].amount == nextOutcome[0].allocations[0].amount,
			'amount changed in allocation'
		);
	}

	function _requireValidFundsSplit(
		Outcome.SingleAssetExit[] memory prevOutcome,
		Outcome.SingleAssetExit[] memory nextOutcome
	) internal pure {
		require(
			prevOutcome[0].allocations[0].amount ==
				nextOutcome[0].allocations[0].amount + nextOutcome[0].allocations[1].amount,
			'amounts sum mismatch'
		);
	}

	function _verifyProofForSettlement(
		FixedPart calldata fixedPart,
		ITradingTypes.Settlement memory settlement,
		RecoveredVariablePart[] calldata proof
	) internal pure {
		bytes32[] memory proofDataHashes = new bytes32[](proof.length);
		uint256 prevTurnNum = 1; // postfund state
		for (uint256 i = 0; i < proof.length - 1; i += 2) {
			VariablePart memory currProof = proof[i].variablePart;
			VariablePart memory nextProof = proof[i + 1].variablePart;

			StrictTurnTaking.isSignedByMover(fixedPart, proof[i]);
			StrictTurnTaking.isSignedByMover(fixedPart, proof[i + 1]);
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
			require(orderResponse.orderID == order.orderID, 'order and response IDs mismatch');

			// outcomes
			if (i != 0) {
				_requireNoAllocationAmountChange(
					proof[i - 1].variablePart.outcome,
					proof[i].variablePart.outcome
				);
			}
			_requireSingleAllocation(proof[i].variablePart.outcome);
			_requireSingleAllocation(proof[i + 1].variablePart.outcome);
			_requireNoAllocationAmountChange(
				proof[i].variablePart.outcome,
				proof[i + 1].variablePart.outcome
			);

			proofDataHashes[i] = keccak256(currProof.appData);
			proofDataHashes[i + 1] = keccak256(nextProof.appData);
			prevTurnNum = nextProof.turnNum;
		}

		bytes32 ordersChecksum = keccak256(abi.encode(proofDataHashes));
		require(ordersChecksum == settlement.ordersChecksum, 'settlement checksum mismatch');
	}
}
