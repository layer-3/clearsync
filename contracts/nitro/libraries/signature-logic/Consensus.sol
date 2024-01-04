// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {NitroUtils} from '../NitroUtils.sol';
import {INitroTypes} from '../../interfaces/INitroTypes.sol';

/**
 * @dev Library for consensus signatures logic, which implies that all participants have signed the candidate state, while supplying proof as empty.
 */
library Consensus {
    /**
     * @notice Require supplied arguments to comply with consensus signatures logic, i.e. each participant has signed the candidate state.
     * @dev Require supplied arguments to comply with consensus signatures logic, i.e. each participant has signed the candidate state.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. The proof is a validation for the supplied candidate. Must be empty.
     * @param candidate Recovered variable part the proof was supplied for. The candidate state is supported by proof states. All participants must have signed this state.
     */
    function requireConsensus(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart[] memory proof,
        INitroTypes.RecoveredVariablePart memory candidate
    ) internal pure {
        require(proof.length == 0, '|proof|!=0');
        require(
            NitroUtils.getClaimedSignersNum(candidate.signedBy) == fixedPart.participants.length,
            '!unanimous'
        );
    }
}
