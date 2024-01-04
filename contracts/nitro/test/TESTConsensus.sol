// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Consensus} from '../libraries/signature-logic/Consensus.sol';
import {INitroTypes} from '../interfaces/INitroTypes.sol';

/**
 * @dev This contract uses the Consensus library to make it more easily unit-tested. It exposes public or external functions which call into internal functions. It should not be deployed in a production environment.
 */
contract TESTConsensus {
    /**
     * @notice Wrapper for otherwise internal function. Require supplied arguments to comply with consensus signatures logic, i.e. each participant has signed the candidate state.
     * @dev Require supplied arguments to comply with consensus signatures logic, i.e. each participant has signed the candidate state.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. The proof is a validation for the supplied candidate. Must be empty.
     * @param candidate Recovered variable part the proof was supplied for. The candidate state is supported by proof states. All participants must have signed this state.
     */
    function requireConsensus(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart[] memory proof,
        INitroTypes.RecoveredVariablePart memory candidate
    ) public pure {
        Consensus.requireConsensus(fixedPart, proof, candidate);
    }
}
