// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';
import {StrictTurnTaking} from './libraries/signature-logic/StrictTurnTaking.sol';
import {IForceMoveApp} from './interfaces/IForceMoveApp.sol';

/**
 * @dev The CountingApp contract complies with the ForceMoveApp interface and strict turn taking logic and allows only for a simple counter to be incremented. Used for testing purposes.
 */
contract CountingApp is IForceMoveApp {
    struct CountingAppData {
        uint256 counter;
    }

    /**
     * @notice Decodes the appData.
     * @dev Decodes the appData.
     * @param appDataBytes The abi.encode of a CountingAppData struct describing the application-specific data.
     * @return A CountingAppData struct containing the application-specific data.
     */
    function appData(bytes memory appDataBytes) internal pure returns (CountingAppData memory) {
        return abi.decode(appDataBytes, (CountingAppData));
    }

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
        StrictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate);

        require(proof.length != 0, '|proof| = 0');

        // validate the proof
        for (uint256 i = 1; i < proof.length; i++) {
            _requireIncrementedCounter(proof[i], proof[i - 1]);
            _requireEqualOutcomes(proof[i], proof[i - 1]);
        }

        _requireIncrementedCounter(candidate, proof[proof.length - 1]);
        _requireEqualOutcomes(candidate, proof[proof.length - 1]);

        return (true, '');
    }

    /**
     * @notice Checks that counter encoded in first variable part equals an incremented counter in second variable part.
     * @dev Checks that counter encoded in first variable part equals an incremented counter in second variable part.
     * @param b RecoveredVariablePart with incremented counter.
     * @param a RecoveredVariablePart with counter before incrementing.
     */
    function _requireIncrementedCounter(
        RecoveredVariablePart memory b,
        RecoveredVariablePart memory a
    ) internal pure {
        require(
            appData(b.variablePart.appData).counter == appData(a.variablePart.appData).counter + 1,
            'Counter must be incremented'
        );
    }

    /**
     * @notice Checks that supplied signed variable parts contain the same outcome.
     * @dev Checks that supplied signed variable parts contain the same outcome.
     * @param a First RecoveredVariablePart.
     * @param b Second RecoveredVariablePart.
     */
    function _requireEqualOutcomes(
        RecoveredVariablePart memory a,
        RecoveredVariablePart memory b
    ) internal pure {
        require(
            Outcome.exitsEqual(a.variablePart.outcome, b.variablePart.outcome),
            'Outcome must not change'
        );
    }
}
