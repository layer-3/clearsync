// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {NitroUtils} from '../libraries/NitroUtils.sol';
import {INitroTypes} from '../interfaces/INitroTypes.sol';

/**
 * @notice NOTE: Development of this signature logic was discontinued, but may be renewed in future. Use at your own risk.
 * @dev Signatures implied by `signedBy` part of `RecoveredVariablePart` must be in ascending order relative to participant index, which has created the signature.
 */
library ShortcuttingTurnTaking {
    /**
     * @notice Require supplied arguments to comply with shortcutting turn taking logic, i.e. there is a signature for each participant, either on the hash of the state for which they are a mover, or on the hash of a state that appears after that state in the array..
     * @dev Require supplied arguments to comply with shortcutting turn taking logic, i.e. there is a signature for each participant, either on the hash of the state for which they are a mover, or on the hash of a state that appears after that state in the array..
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. The proof is a validation for the supplied candidate.
     * @param candidate Recovered variable part the proof was supplied for. The candidate state is supported by proof states.
     */
    function requireValidTurnTaking(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart[] memory proof,
        INitroTypes.RecoveredVariablePart memory candidate
    ) internal pure {
        uint256 nParticipants = fixedPart.participants.length;
        uint48 largestTurnNum = candidate.variablePart.turnNum;

        _requireValidInput(nParticipants, proof, candidate);

        // The difference between the support proof candidate turn number (aka largestTurnNum) and the round robin cycle last turn number.
        uint256 roundRobinShift = (largestTurnNum + 1) % nParticipants;
        uint48 prevTurnNum = 0;

        // validating the proof
        for (uint256 i = 0; i < proof.length; i++) {
            requireValidSignatures(fixedPart, proof[i], roundRobinShift);

            if (i != 0) {
                requireIncreasedTurnNum(prevTurnNum, proof[i].variablePart.turnNum);
            }

            prevTurnNum = proof[i].variablePart.turnNum;
        }

        // validating the candidate
        requireValidSignatures(fixedPart, candidate, roundRobinShift);
        requireIncreasedTurnNum(prevTurnNum, candidate.variablePart.turnNum);
    }

    /**
     * @notice Given a state, checks the validity of the supplied signatures. Valid means each signature correspond to a participant, who is either a mover for this state or was a mover for some preceding state.
     * @dev Given a state, checks the validity of the supplied signatures. Valid means each signature correspond to a participant, who is either a mover for this state or was a mover for some preceding state.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param recoveredVariablePart A struct describing dynamic properties of the state channel, that must be signed either by a participant, who is either a mover for this state or was a mover for some preceding state.
     * @param roundRobinShift Difference between a turn number of the last state, which have a last participant as a mover, and supplied largest turn number.
     */
    function requireValidSignatures(
        INitroTypes.FixedPart memory fixedPart,
        INitroTypes.RecoveredVariablePart memory recoveredVariablePart,
        uint256 roundRobinShift
    ) internal pure {
        require(
            NitroUtils.getClaimedSignersNum(recoveredVariablePart.signedBy) > 0,
            'Insufficient signatures'
        );

        _requireAcceptableSigsOrder(
            recoveredVariablePart.signedBy,
            recoveredVariablePart.variablePart.turnNum,
            roundRobinShift,
            fixedPart.participants.length
        );

        uint8[] memory signerIndices = NitroUtils.getClaimedSignersIndices(
            recoveredVariablePart.signedBy
        );

        for (uint256 i = 0; i < signerIndices.length; i++) {
            require(
                NitroUtils.isClaimedSignedBy(recoveredVariablePart.signedBy, signerIndices[i]),
                'Invalid signer'
            );
        }
    }

    /**
     * @notice Given a declaration of which participant have signed the supplied state, check if this declaration is acceptable. Acceptable means there is a signature for each participant, either on the hash of the state for which they are a mover, or on the hash of a state that appears after that state in the array.
     * @dev Given a declaration of which participant have signed the supplied state, check if this declaration is acceptable. Acceptable means there is a signature for each participant, either on the hash of the state for which they are a mover, or on the hash of a state that appears after that state in the array.
     * @param signedBy Bit mask field specifying which participants have signed the state.
     * @param turnNum Turn number of the state to check.
     * @param shift Difference between a turn number of the last state, which have a last participant as a mover, and supplied largest turn number.
     * @param nParticipants Number of participants in a channel.
     */
    function _requireAcceptableSigsOrder(
        uint256 signedBy,
        uint48 turnNum,
        uint256 shift,
        uint256 nParticipants
    ) internal pure {
        uint8[] memory signerIndices = NitroUtils.getClaimedSignersIndices(signedBy);

        for (uint256 i = 0; i < signerIndices.length; i++) {
            require(
                (signerIndices[i] + nParticipants - shift) % nParticipants <=
                    (turnNum - shift) % nParticipants,
                'Unacceptable sigs order'
            );
        }
    }

    /**
     * @notice Require supplied newTurnNum is greater than prevTurnNum.
     * @dev Require supplied newTurnNum is greater than prevTurnNum.
     * @param prevTurnNum Previous turn number.
     * @param newTurnNum New turn number.
     */
    function requireIncreasedTurnNum(uint48 prevTurnNum, uint48 newTurnNum) internal pure {
        require(prevTurnNum < newTurnNum, 'turnNum not increased');
    }

    /**
     * @notice Validate input for turn taking logic.
     * @dev Validate input for turn taking logic.
     * @param nParticipants Number of participants in a channel.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. The proof is a validation for the supplied candidate.
     * @param candidate Recovered variable part the proof was supplied for. The candidate state is supported by proof states.
     */
    function _requireValidInput(
        uint256 nParticipants,
        INitroTypes.RecoveredVariablePart[] memory proof,
        INitroTypes.RecoveredVariablePart memory candidate
    ) internal pure {
        uint256 numStates = proof.length + 1;
        require((nParticipants >= numStates) && (numStates > 0), 'Insufficient or excess states');

        uint256 largestTurnNum = candidate.variablePart.turnNum;
        require(largestTurnNum + 1 >= nParticipants, 'largestTurnNum too low');

        // no more than 255 participants
        require(nParticipants <= type(uint8).max, 'Too many participants'); // type(uint8).max = 2**8 - 1 = 255

        // when proof.length = 0, turnNumDelta is always = 0 and thus <= nParticipants
        if (proof.length != 0) {
            uint256 turnNumDelta = largestTurnNum - proof[0].variablePart.turnNum;
            require(turnNumDelta <= nParticipants, 'Only one round-robin allowed');
        }

        uint256 signedSoFar = 0;
        uint256 hasTwoSigs = 0;

        // processing the proof
        for (uint256 i = 0; i < proof.length; i++) {
            hasTwoSigs = signedSoFar & proof[i].signedBy;
            require(hasTwoSigs == 0, 'Excess sigs from one participant');
            signedSoFar |= proof[i].signedBy;
        }

        // processing candidate
        hasTwoSigs = signedSoFar & candidate.signedBy;
        require(hasTwoSigs == 0, 'Excess sigs from one participant');
        signedSoFar |= candidate.signedBy;

        require(signedSoFar == 2 ** nParticipants - 1, 'Lacking participant signature');
    }
}
