// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {INitroTypes} from '../interfaces/INitroTypes.sol';
import {NitroUtils} from '../libraries/NitroUtils.sol';

/**
 * @dev This contract uses the NitroUtils library to enable it to be more easily unit-tested. It exposes public or external functions which call into internal functions. It should not be deployed in a production environment.
 */
contract TESTNitroUtils {
    /**
     * @dev Wrapper for otherwise internal function. Given a digest and digital signature, recover the signer
     * @param _d message digest
     * @param sig ethereum digital signature
     * @return signer
     */
    function recoverSigner(
        bytes32 _d,
        INitroTypes.Signature memory sig
    ) public pure returns (address) {
        return NitroUtils.recoverSigner(_d, sig);
    }

    /**
     * @notice Check if supplied participantIndex bit is set to 1 in signedBy bit mask.
     * @dev Check if supplied partitipationIndex bit is set to 1 in signedBy bit mask.
     * @param signedBy Bit mask field to check.
     * @param participantIndex Bit to check.
     */
    function isClaimedSignedBy(
        uint256 signedBy,
        uint8 participantIndex
    ) public pure returns (bool) {
        return NitroUtils.isClaimedSignedBy(signedBy, participantIndex);
    }

    /**
     * @notice Check if supplied participantIndex is the only bit set to 1 in signedBy bit mask.
     * @dev Check if supplied participantIndex is the only bit set to 1 in signedBy bit mask.
     * @param signedBy Bit mask field to check.
     * @param participantIndex Bit to check.
     */
    function isClaimedSignedOnlyBy(
        uint256 signedBy,
        uint8 participantIndex
    ) public pure returns (bool) {
        return NitroUtils.isClaimedSignedOnlyBy(signedBy, participantIndex);
    }

    /**
     * @notice Count number of bits set to '1', specifying the number of participants which have signed the state.
     * @dev Count number of bits set to '1', specifying the number of participants which have signed the state.
     * @param signedBy Bit mask field specifying which participants have signed the state.
     * @return amount of signers, which have signed the state.
     */
    function getClaimedSignersNum(uint256 signedBy) public pure returns (uint8) {
        return NitroUtils.getClaimedSignersNum(signedBy);
    }

    /**
     * @notice Determine indices of participants who have signed the state.
     * @dev Determine indices of participants who have signed the state.
     * @param signedBy Bit mask field specifying which participants have signed the state.
     * @return signerIndices
     */
    function getClaimedSignersIndices(uint256 signedBy) public pure returns (uint8[] memory) {
        return NitroUtils.getClaimedSignersIndices(signedBy);
    }
}
