// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import '../nitro/interfaces/IForceMoveApp.sol';
import '../nitro/libraries/NitroUtils.sol';

/**
 * @dev The EscrowApp contracts complies with the ForceMoveApp interface and uses consensus signatures logic.
 */
contract EscrowApp is IForceMoveApp {
	/**
	 * @notice Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @dev Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @param fixedPart Fixed part of the state channel.
	 * @param proof Array of recovered variable parts which constitutes a support proof for the candidate.
	 * @param candidate Recovered variable part the proof was supplied for.
	 */
	function stateIsSupported(
		FixedPart calldata fixedPart,
		RecoveredVariablePart[] calldata proof,
		RecoveredVariablePart calldata candidate
	) external pure override returns (bool, string memory) {
		// This channel has only 4 states which can be supported:
		// 0    prefund
		// 1    postfund
		// 2 	settlement
		// 3 	final

		uint8 nParticipants = uint8(fixedPart.participants.length);

		if (nParticipants != 2) return (false, 'nParticipants != 2');

		// prefund, postfund, final:
		if (proof.length == 0) {
			require(
				NitroUtils.getClaimedSignersNum(candidate.signedBy) == nParticipants,
				'!unanimous; |proof|=0'
			);

			if (candidate.variablePart.turnNum == 0) return (true, ''); // prefund
			if (candidate.variablePart.turnNum == 1) return (true, ''); // postfund

			// final
			if (candidate.variablePart.turnNum == 3) {
				// final (note: there is a core protocol escape hatch for this, too, so it could be removed)
				require(candidate.variablePart.isFinal, '!final; turnNum>=3 && |proof|=0');
				return (true, '');
			}

			revert('bad candidate turnNum; |proof|=0');
		}

		// settlement. Requires a postfund state to be supplied
		if (proof.length == 1) {
			// postfund checks
			_requireProofOfUnanimousConsensusOnPostFund(proof[0], nParticipants);

			_requireOneAllocationPerAsset(proof[0].variablePart.outcome);
			_requireOneAllocationPerAsset(candidate.variablePart.outcome);

			_requireAllocationSwapped(
				proof[0].variablePart.outcome,
				candidate.variablePart.outcome
			);

			return (true, '');
		}

		return (false, 'bad proof length');
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

	function _requireOneAllocationPerAsset(Outcome.SingleAssetExit[] memory outcome) internal pure {
		for (uint256 i = 0; i < outcome.length; i++) {
			Outcome.SingleAssetExit memory exit = outcome[i];
			require(exit.allocations.length == 1, '|allocations| != 1');
		}
	}

	function _requireAllocationSwapped(
		Outcome.SingleAssetExit[] memory outcomeA,
		Outcome.SingleAssetExit[] memory outcomeB
	) internal pure {
		require(outcomeA.length == outcomeB.length, '|outcomeA| != |outcomeB|');

		// TODO: verify that there are only 2 destinations

		for (uint256 i = 0; i < outcomeA.length; i++) {
			require(outcomeA[i].asset == outcomeB[i].asset, 'asset mismatch');

			Outcome.Allocation memory allocA = outcomeA[i].allocations[0];
			Outcome.Allocation memory allocB = outcomeB[i].allocations[0];

			require(allocA.destination != allocB.destination, 'destination must swap');
			require(allocA.amount == allocB.amount, 'amount mismatch');
		}
	}
}
