// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';
import {StrictTurnTaking} from '../libraries/signature-logic/StrictTurnTaking.sol';
import {IForceMoveApp} from '../interfaces/IForceMoveApp.sol';

/**
 * @dev The SingleAssetPayments contract complies with the ForceMoveApp interface, uses strict turn taking logic and implements a simple payment channel with a single asset type only.
 */
contract SingleAssetPayments is IForceMoveApp {
    /**
     * @notice Encodes application-specific rules for a particular ForceMove-compliant state channel. Must revert when invalid support proof and a candidate are supplied.
     * @dev Encodes application-specific rules for a particular ForceMove-compliant state channel. Must revert when invalid support proof and a candidate are supplied.
     * @param fixedPart Fixed part of the state channel.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate.
     * @param candidate Recovered variable part the proof was supplied for.
     */
    function stateIsSupported(
        FixedPart calldata fixedPart,
        RecoveredVariablePart[] calldata proof,
        RecoveredVariablePart calldata candidate
    ) external pure override returns (bool, string memory) {
        StrictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate);

        for (uint256 i = 0; i < proof.length; i++) {
            _requireValidOutcome(fixedPart.participants.length, proof[i].variablePart.outcome);

            if (i > 0) {
                _requireValidTransition(
                    fixedPart.participants.length,
                    proof[i - 1].variablePart,
                    proof[i].variablePart
                );
            }
        }

        _requireValidOutcome(fixedPart.participants.length, candidate.variablePart.outcome);

        _requireValidTransition(
            fixedPart.participants.length,
            proof[proof.length - 1].variablePart,
            candidate.variablePart
        );
        return (true, '');
    }

    /**
     * @notice Require specific rules in outcome are followed.
     * @dev Require specific rules in outcome are followed.
     * @param nParticipants Number of participants in a channel.
     * @param outcome Outcome to check.
     */
    function _requireValidOutcome(
        uint256 nParticipants,
        Outcome.SingleAssetExit[] memory outcome
    ) internal pure {
        // Throws if more than one asset
        require(outcome.length == 1, 'outcome: Only one asset allowed');

        // Throws unless that allocation has exactly n outcomes
        Outcome.Allocation[] memory allocations = outcome[0].allocations;

        require(allocations.length == nParticipants, '|Allocation|!=|participants|');

        for (uint256 i = 0; i < nParticipants; i++) {
            require(
                allocations[i].allocationType == uint8(Outcome.AllocationType.simple),
                'not a simple allocation'
            );
        }
    }

    /**
     * @notice Require specific rules in variable parts are followed when progressing state.
     * @dev Require specific rules in variable parts are followed when progressing state.
     * @param nParticipants Number of participants in a channel.
     * @param a Variable part to progress from.
     * @param b Variable part to progress to.
     */
    function _requireValidTransition(
        uint256 nParticipants,
        VariablePart memory a,
        VariablePart memory b
    ) internal pure {
        Outcome.Allocation[] memory allocationsA = a.outcome[0].allocations;
        Outcome.Allocation[] memory allocationsB = b.outcome[0].allocations;

        // Interprets the nth outcome as benefiting participant n
        // checks the destinations have not changed
        // Checks that the sum of assets hasn't changed
        // And that for all non-movers
        // the balance hasn't decreased
        uint256 allocationSumA;
        uint256 allocationSumB;
        for (uint256 i = 0; i < nParticipants; i++) {
            require(
                allocationsB[i].destination == allocationsA[i].destination,
                'Destinations may not change'
            );
            allocationSumA += allocationsA[i].amount;
            allocationSumB += allocationsB[i].amount;
            if (i != b.turnNum % nParticipants) {
                require(
                    allocationsB[i].amount >= allocationsA[i].amount,
                    'Nonmover balance decreased'
                );
            }
        }
        require(allocationSumA == allocationSumB, 'Total allocated cannot change');
    }
}
