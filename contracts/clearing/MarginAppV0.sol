// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@statechannels/nitro-protocol/contracts/interfaces/IForceMoveApp.sol';
import '@statechannels/nitro-protocol/contracts/libraries/NitroUtils.sol';
import '@statechannels/nitro-protocol/contracts/interfaces/INitroTypes.sol';
import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

// NOTE: Attack:
// Bob can submit a convenient candidate, when Alice in trouble (Way back machine attack)

// Possible solutions:
// 1: Alice does checkpoint periodically
// 2: Alice hire a WatchTower, which replicates Alice's states,
// and challenge in the case of challenge event and missing heartbeat

// TODO: change `revert` to `return (false, 'error')` in all require statements

/**
 * @dev The MarginApp contract complies with the ForceMoveApp interface and allows payments to be made virtually from Initiator to Follower (participants[0] to participants[n+1], where n is the number of intermediaries).
 */
contract MarginAppV0 is IForceMoveApp {
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
		// This channel has only 5 states which can be supported:
		// 0    prefund
		// 1    postfund
		// 2+   margin call
		// TODO: 3+   settlement request (requires explicit margin call to be agreed on before)
		// 3+   final

		uint8 nParticipants = uint8(fixedPart.participants.length);
		require(nParticipants == 2, 'only 2 participants supported');

		require(NitroUtils.getClaimedSignersNum(candidate.signedBy) == nParticipants, '!unanimous');

		// states 0,1,3+:
		if (proof.length == 0) {
			if (candidate.variablePart.turnNum == 0) return (true, ''); // prefund
			if (candidate.variablePart.turnNum == 1) return (true, ''); // postfund

			// final
			if (candidate.variablePart.turnNum >= 3) {
				// final (note: there is a core protocol escape hatch for this, too, so it could be removed)
				require(candidate.variablePart.isFinal, '!final; turnNum>=3 && |proof|=0');
				return (true, '');
			}

			revert('bad candidate turnNum; |proof|=0');
		}

		// state 2+ requires postfund state to be supplied
		if (proof.length == 1) {
			require(candidate.variablePart.turnNum >= 2, 'turnNum<2 && |proof|=1');

			// postfund checks
			require(proof[0].variablePart.turnNum == 1, 'postfund.turnNum != 1');
			require(
				NitroUtils.getClaimedSignersNum(proof[0].signedBy) == nParticipants,
				'postfund !unanimous'
			);

			// postfund - margin call checks
			_requireValidOutcomeTransition(
				proof[0].variablePart.outcome,
				candidate.variablePart.outcome
			);
			return (true, '');
		}

		revert('bad proof length');
	}

	function _requireValidOutcomeTransition(
		Outcome.SingleAssetExit[] memory oldOutcome,
		Outcome.SingleAssetExit[] memory newOutcome
	) internal pure {
		// TODO: change to support 2 collateral assets
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
