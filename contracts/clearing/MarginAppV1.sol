// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@statechannels/nitro-protocol/contracts/interfaces/IForceMoveApp.sol';
import '@statechannels/nitro-protocol/contracts/libraries/NitroUtils.sol';
import '@statechannels/nitro-protocol/contracts/interfaces/INitroTypes.sol';
import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

import '../interfaces/IClearingTypes.sol';

// NOTE: Attack:
// Bob can submit a convenient candidate, when Alice in trouble (Way back machine attack)

// Possible solutions:
// 1: Alice does checkpoint periodically
// 2: Alice hire a WatchTower, which replicates Alice's states,
// and challenge in the case of challenge event and missing heartbeat

/**
 * @dev The MarginApp contract complies with the ForceMoveApp interface and allows payments to be made virtually from Initiator to Follower (participants[0] to participants[n+1], where n is the number of intermediaries).
 */
contract MarginAppV1 is IForceMoveApp {
	/**
	 * @notice Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @dev Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @param fixedPart Fixed part of the state channel.
	 * @param proof Array of recovered variable parts which constitutes a support proof for the candidate.
	 * @param candidate Recovered variable part the proof was supplied for.
	 */
	function requireStateSupported(
		FixedPart calldata fixedPart,
		RecoveredVariablePart[] calldata proof,
		RecoveredVariablePart calldata candidate
	) external pure override {
		// This channel has only 4 states which can be supported:
		// 0    prefund
		// 1    postfund
		// 2+   margin call
		// 3+   settlement request (requires explicit margin call to be agreed on before)
		// 3+   final

		uint8 nParticipants = uint8(fixedPart.participants.length);

		// states 0,1,3+:
		if (proof.length == 0) {
			require(
				NitroUtils.getClaimedSignersNum(candidate.signedBy) == nParticipants,
				'!unanimous; |proof|=0'
			);

			if (candidate.variablePart.turnNum == 0) return; // prefund
			if (candidate.variablePart.turnNum == 1) return; // postfund

			// final
			if (candidate.variablePart.turnNum >= 3) {
				// final (note: there is a core protocol escape hatch for this, too, so it could be removed)
				require(candidate.variablePart.isFinal, '!final; turnNum>=3 && |proof|=0');
				return;
			}

			revert('bad candidate turnNum; |proof|=0');
		}

		// state 2+ requires postfund state to be supplied
		if (proof.length == 1) {
			// postfund checks
			_requireProofOfUnanimousConsensusOnPostFund(proof[0], nParticipants);

			// postfund - margin call checks
			_requireValidOutcomeTransition(
				proof[0].variablePart.outcome,
				candidate.variablePart.outcome
			);
			_requireValidMarginCall(fixedPart, candidate);
			return;
		}

		// state 2+ margin in settlement request is supported
		// proof[0] - postfund
		// proof[1] - margin call preceding settlement request
		// candidate - settlement request
		if (proof.length == 2) {
			// postfund checks
			_requireProofOfUnanimousConsensusOnPostFund(proof[0], nParticipants);

			// postfund - margin call checks
			_requireValidOutcomeTransition(
				proof[0].variablePart.outcome,
				proof[1].variablePart.outcome
			);
			_requireValidMarginCall(fixedPart, proof[1]);

			// margin call - settlement request checks
			require(
				proof[1].variablePart.turnNum + 1 == candidate.variablePart.turnNum,
				'settlementRequest not direct successor of marginCall'
			);
			_requireValidOutcomeTransition(
				proof[1].variablePart.outcome,
				candidate.variablePart.outcome
			);
			_requireValidSettlementRequest(fixedPart, candidate);
			return;
		}

		revert('bad proof length');
	}

	function _requireProofOfUnanimousConsensusOnPostFund(
		RecoveredVariablePart memory rVP,
		uint256 numParticipants
	) internal pure {
		require(rVP.variablePart.turnNum == 1, 'postfund.turnNum != 1');
		require(
			NitroUtils.getClaimedSignersNum(rVP.signedBy) == numParticipants,
			'postfund !unanimous'
		);
	}

	// margin call in app data
	function _requireValidMarginCall(
		FixedPart memory fixedPart,
		RecoveredVariablePart memory marginCallState
	) internal pure {
		uint8 nParticipants = uint8(fixedPart.participants.length);

		require(marginCallState.variablePart.turnNum >= 2, 'marginCall.turnNum < 2');

		// supplied state must be signed by either party
		require(
			NitroUtils.isClaimedSignedBy(marginCallState.signedBy, 0) ||
				NitroUtils.isClaimedSignedBy(marginCallState.signedBy, nParticipants - 1),
			'no identity proof on margin call'
		);

		IClearingTypes.SignedMarginCall memory sMC = abi.decode(
			marginCallState.variablePart.appData,
			(IClearingTypes.SignedMarginCall)
		);
		_requireVariablePartFitsMarginCall(marginCallState.variablePart, sMC.marginCall);
		_requireValidSigs(
			abi.encode(NitroUtils.getChannelId(fixedPart), sMC.marginCall),
			sMC.sigs,
			[fixedPart.participants[0], fixedPart.participants[nParticipants - 1]]
		);
	}

	// settlement request in app data, margin call part of settlement request
	function _requireValidSettlementRequest(
		FixedPart memory fixedPart,
		RecoveredVariablePart memory settlementRequestState
	) internal pure {
		uint8 nParticipants = uint8(fixedPart.participants.length);

		require(settlementRequestState.variablePart.turnNum >= 3, 'settlementRequest.turnNum < 3');

		// supplied state must be signed by either party
		require(
			NitroUtils.isClaimedSignedBy(settlementRequestState.signedBy, 0) ||
				NitroUtils.isClaimedSignedBy(settlementRequestState.signedBy, nParticipants - 1),
			'no identity proof on settlement request'
		);

		IClearingTypes.SignedSettlementRequest memory sSR = abi.decode(
			settlementRequestState.variablePart.appData,
			(IClearingTypes.SignedSettlementRequest)
		);
		_requireValidSettlementRequest(
			fixedPart.participants,
			settlementRequestState.variablePart,
			sSR.settlementRequest
		);
		_requireValidSigs(
			abi.encode(NitroUtils.getChannelId(fixedPart), sSR.settlementRequest),
			sSR.sigs,
			[fixedPart.participants[0], fixedPart.participants[nParticipants - 1]]
		);
	}

	function _requireVariablePartFitsMarginCall(
		VariablePart memory variablePart,
		IClearingTypes.MarginCall memory marginCall
	) internal pure {
		// correct margin version
		require(marginCall.version == variablePart.turnNum, 'marginCall.version != turnNum');

		uint256 initiatorIdx = uint256(IClearingTypes.MarginIndices.Initiator);
		uint256 followerIdx = uint256(IClearingTypes.MarginIndices.Follower);

		// correct outcome adjustments
		require(
			variablePart.outcome[0].allocations[initiatorIdx].amount ==
				marginCall.margin[initiatorIdx],
			'incorrect initiator margin'
		);
		require(
			variablePart.outcome[0].allocations[followerIdx].amount ==
				marginCall.margin[followerIdx],
			'incorrect follower margin'
		);
	}

	// check signed settlement request and included margin call
	function _requireValidSettlementRequest(
		address[] memory participants,
		VariablePart memory variablePart,
		IClearingTypes.SettlementRequest memory settlementRequest
	) internal pure {
		// brokers are participants
		require(
			settlementRequest.brokers[uint256(IClearingTypes.MarginIndices.Initiator)] ==
				participants[0],
			'1st broker not initiator'
		);
		require(
			settlementRequest.brokers[uint256(IClearingTypes.MarginIndices.Follower)] ==
				participants[participants.length - 1],
			'2nd broker not follower'
		);

		// correct settlement version
		require(
			settlementRequest.version == variablePart.turnNum,
			'settlementRequest.version != turnNum'
		);

		// FIXME: `NitroUtils.getChainID()` is view, but `requireStateSupported` is pure
		// correct chainId
		// require(settlementRequest.chainId == NitroUtils.getChainID(), 'incorrect chainId');

		// correct adjusted margin call, outcome
		_requireVariablePartFitsMarginCall(variablePart, settlementRequest.adjustedMargin);
	}

	function _requireValidSigs(
		bytes memory signedData,
		Signature[2] memory sigs,
		address[2] memory signers
	) internal pure {
		// correct initiator signature
		uint256 initiatorIdx = uint256(IClearingTypes.MarginIndices.Initiator);
		address recoveredInitiator = NitroUtils.recoverSigner(
			keccak256(signedData),
			sigs[initiatorIdx]
		);
		require(recoveredInitiator == signers[initiatorIdx], 'invalid initiator signature'); // could be incorrect channelId or incorrect signature

		// correct follower signature
		uint256 followerIdx = uint256(IClearingTypes.MarginIndices.Follower);
		address recoveredFollower = NitroUtils.recoverSigner(
			keccak256(signedData),
			sigs[followerIdx]
		);
		require(recoveredFollower == signers[followerIdx], 'invalid follower signature'); // could be incorrect channelId or incorrect signature
	}

	function _requireValidOutcomeTransition(
		Outcome.SingleAssetExit[] memory oldOutcome,
		Outcome.SingleAssetExit[] memory newOutcome
	) internal pure {
		// only 1 collateral asset (USDT)
		require(oldOutcome.length == 1 && newOutcome.length == 1, 'incorrect number of assets');

		// only 2 allocations
		require(
			oldOutcome[0].allocations.length == 2 && newOutcome[0].allocations.length == 2,
			'incorrect number of allocations'
		);

		require(
			oldOutcome[0].allocations[0].destination == newOutcome[0].allocations[0].destination &&
				oldOutcome[0].allocations[1].destination ==
				newOutcome[0].allocations[1].destination,
			'destinations cannot change'
		);

		// equal sums
		uint256 oldAllocationSum;
		uint256 newAllocationSum;
		for (uint256 i = 0; i < 2; i++) {
			oldAllocationSum += oldOutcome[0].allocations[i].amount;
			newAllocationSum += newOutcome[0].allocations[i].amount;
		}
		require(oldAllocationSum == newAllocationSum, 'total allocated cannot change');
	}
}
