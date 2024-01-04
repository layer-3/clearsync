// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {INitroTypes} from '../interfaces/INitroTypes.sol';
import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

library NitroUtils {
    // *****************
    // Signature methods:
    // *****************

    /**
     * @notice Require supplied stateHash is signed by signer.
     * @dev Require supplied stateHash is signed by signer.
     * @param stateHash State hash to check.
     * @param sig Signed state signature.
     * @param signer Address which must have signed the state.
     * @return true if signer with sig has signed stateHash.
     */
    function isSignedBy(
        bytes32 stateHash,
        INitroTypes.Signature memory sig,
        address signer
    ) internal pure returns (bool) {
        return signer == NitroUtils.recoverSigner(stateHash, sig);
    }

    /**
     * @notice Check if supplied participantIndex bit is set to 1 in signedBy bit mask.
     * @dev Check if supplied partitipationIndex bit is set to 1 in signedBy bit mask.
     * @param signedBy Bit mask field to check.
     * @param participantIndex Bit to check.
     * @return true if supplied partitipationIndex bit is set to 1 in signedBy bit mask.
     */
    function isClaimedSignedBy(
        uint256 signedBy,
        uint8 participantIndex
    ) internal pure returns (bool) {
        return ((signedBy >> participantIndex) % 2 == 1);
    }

    /**
     * @notice Check if supplied participantIndex is the only bit set to 1 in signedBy bit mask.
     * @dev Check if supplied participantIndex is the only bit set to 1 in signedBy bit mask.
     * @param signedBy Bit mask field to check.
     * @param participantIndex Bit to check.
     * @return true if supplied partitipationIndex bit is the only bit set to 1 in signedBy bit mask.
     */
    function isClaimedSignedOnlyBy(
        uint256 signedBy,
        uint8 participantIndex
    ) internal pure returns (bool) {
        return (signedBy == (2 ** participantIndex));
    }

    /**
     * @notice Given a digest and ethereum digital signature, recover the signer.
     * @dev Given a digest and digital signature, recover the signer.
     * @param _d message digest.
     * @param sig ethereum digital signature.
     * @return signer
     */
    function recoverSigner(
        bytes32 _d,
        INitroTypes.Signature memory sig
    ) internal pure returns (address) {
        bytes32 prefixedHash = keccak256(abi.encodePacked('\x19Ethereum Signed Message:\n32', _d));
        address a = ecrecover(prefixedHash, sig.v, sig.r, sig.s);
        require(a != address(0), 'Invalid signature');
        return (a);
    }

    /**
     * @notice Count number of bits set to '1', specifying the number of participants which have signed the state.
     * @dev Count number of bits set to '1', specifying the number of participants which have signed the state.
     * @param signedBy Bit mask field specifying which participants have signed the state.
     * @return amount of signers, which have signed the state.
     */
    function getClaimedSignersNum(uint256 signedBy) internal pure returns (uint8) {
        uint8 amount = 0;

        for (; signedBy > 0; amount++) {
            // for reference: Kernighan's Bit Counting Algorithm
            signedBy &= signedBy - 1;
        }

        return amount;
    }

    /**
     * @notice Determine indices of participants who have signed the state.
     * @dev Determine indices of participants who have signed the state.
     * @param signedBy Bit mask field specifying which participants have signed the state.
     * @return signerIndices
     */
    function getClaimedSignersIndices(uint256 signedBy) internal pure returns (uint8[] memory) {
        uint8[] memory signerIndices = new uint8[](getClaimedSignersNum(signedBy));
        uint8 signerNum = 0;
        uint8 acceptedSigners = 0;

        for (; signedBy > 0; signerNum++) {
            if (signedBy % 2 == 1) {
                signerIndices[acceptedSigners] = signerNum;
                acceptedSigners++;
            }
            signedBy >>= 1;
        }

        return signerIndices;
    }

    // *****************
    // ID methods:
    // *****************

    /**
     * @notice Computes the unique id of a channel.
     * @dev Computes the unique id of a channel.
     * @param fixedPart Part of the state that does not change
     * @return channelId
     */
    function getChannelId(
        INitroTypes.FixedPart memory fixedPart
    ) internal pure returns (bytes32 channelId) {
        channelId = keccak256(
            abi.encode(
                fixedPart.participants,
                fixedPart.channelNonce,
                fixedPart.appDefinition,
                fixedPart.challengeDuration
            )
        );
    }

    // *****************
    // Hash methods:
    // *****************

    /**
     * @notice Computes the hash of the state corresponding to the input data.
     * @dev Computes the hash of the state corresponding to the input data.
     * @param turnNum Turn number
     * @param isFinal Is the state final?
     * @param channelId Unique identifier for the channel
     * @param appData Application specific data.
     * @param outcome Outcome structure.
     * @return The stateHash
     */
    function hashState(
        bytes32 channelId,
        bytes memory appData,
        Outcome.SingleAssetExit[] memory outcome,
        uint48 turnNum,
        bool isFinal
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(channelId, appData, outcome, turnNum, isFinal));
    }

    /**
     * @notice Computes the hash of the state corresponding to the input data.
     * @dev Computes the hash of the state corresponding to the input data.
     * @param fp The FixedPart of the state
     * @param vp The VariablePart of the state
     * @return The stateHash
     */
    function hashState(
        INitroTypes.FixedPart memory fp,
        INitroTypes.VariablePart memory vp
    ) internal pure returns (bytes32) {
        return
            keccak256(abi.encode(getChannelId(fp), vp.appData, vp.outcome, vp.turnNum, vp.isFinal));
    }

    /**
     * @notice Hashes the outcome structure. Internal helper.
     * @dev Hashes the outcome structure. Internal helper.
     * @param outcome Outcome structure to encode hash.
     * @return bytes32 Hash of encoded outcome structure.
     */
    function hashOutcome(Outcome.SingleAssetExit[] memory outcome) internal pure returns (bytes32) {
        return keccak256(Outcome.encodeExit(outcome));
    }

    // *****************
    // Equality methods:
    // *****************

    /**
     * @notice Check for equality of two byte strings
     * @dev Check for equality of two byte strings
     * @param _preBytes One bytes string
     * @param _postBytes The other bytes string
     * @return true if the bytes are identical, false otherwise.
     */
    function bytesEqual(
        bytes memory _preBytes,
        bytes memory _postBytes
    ) internal pure returns (bool) {
        // copied from https://www.npmjs.com/package/solidity-bytes-utils/v/0.1.1
        bool success = true;

        /* solhint-disable no-inline-assembly */
        assembly {
            let length := mload(_preBytes)

            // if lengths don't match the arrays are not equal
            switch eq(length, mload(_postBytes))
            case 1 {
                // cb is a circuit breaker in the for loop since there's
                //  no said feature for inline assembly loops
                // cb = 1 - don't breaker
                // cb = 0 - break
                let cb := 1

                let mc := add(_preBytes, 0x20)
                let end := add(mc, length)

                for {
                    let cc := add(_postBytes, 0x20)
                    // the next line is the loop condition:
                    // while(uint256(mc < end) + cb == 2)
                } eq(add(lt(mc, end), cb), 2) {
                    mc := add(mc, 0x20)
                    cc := add(cc, 0x20)
                } {
                    // if any of these checks fails then arrays are not equal
                    if iszero(eq(mload(mc), mload(cc))) {
                        // unsuccess:
                        success := 0
                        cb := 0
                    }
                }
            }
            default {
                // unsuccess:
                success := 0
            }
        }
        /* solhint-disable no-inline-assembly */

        return success;
    }
}
