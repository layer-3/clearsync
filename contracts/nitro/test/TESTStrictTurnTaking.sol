// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {INitroTypes} from '../interfaces/INitroTypes.sol';
import {StrictTurnTaking} from '../libraries/signature-logic/StrictTurnTaking.sol';

/**
 * @dev This contract uses the StrictTurnTaking library to make it more easily unit-tested. It exposes public or external functions which call into internal functions. It should not be deployed in a production environment.
 */
contract TESTStrictTurnTaking {
    /**
     * @notice Wrapper for otherwise internal function. Require supplied arguments to comply with turn taking logic, i.e. each participant signed the one state, they were mover for.
     * @dev Require supplied arguments to comply with turn taking logic, i.e. each participant signed the one state, they were mover for.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. The proof is a validation for the supplied candidate.
     * @param candidate Recovered variable part the proof was supplied for. The candidate state is supported by proof states.
     */
    function requireValidTurnTaking(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart[] memory proof,
        INitroTypes.RecoveredVariablePart memory candidate
    ) public pure {
        StrictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate);
    }

    /**
     * @notice Wrapper for otherwise internal function. Require supplied state is signed by its corresponding moving participant.
     * @dev Require supplied state is signed by its corresponding moving participant.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param recoveredVariablePart A struct describing dynamic properties of the state channel, that must be signed by moving participant.
     */
    function isSignedByMover(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart memory recoveredVariablePart
    ) public pure {
        StrictTurnTaking.isSignedByMover(fixedPart, recoveredVariablePart);
    }

    /**
     * @notice Wrapper for otherwise internal function. Find moving participant address based on state turn number.
     * @dev Find moving participant address based on state turn number.
     * @param participants Array of participant addresses.
     * @param turnNum State turn number.
     * @return address Moving partitipant address.
     */
    function moverAddress(
        address[] memory participants,
        uint48 turnNum
    ) public pure returns (address) {
        return StrictTurnTaking._moverAddress(participants, turnNum);
    }

    /**
     * @notice Wrapper for otherwise internal function. Validate input for turn taking logic.
     * @dev Validate input for turn taking logic.
     * @param numParticipants Number of participants in a channel.
     * @param numProofStates Number of proof states submitted.
     */
    function requireValidInput(uint256 numParticipants, uint256 numProofStates) public pure {
        StrictTurnTaking._requireValidInput(numParticipants, numProofStates);
    }
}
