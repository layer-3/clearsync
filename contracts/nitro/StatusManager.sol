// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IStatusManager} from './interfaces/IStatusManager.sol';

/**
 * @dev The StatusManager is responsible for on-chain storage of the status of active channels
 */
contract StatusManager is IStatusManager {
    mapping(bytes32 => bytes32) public statusOf;

    /**
     * @notice Computes the ChannelMode for a given channelId.
     * @dev Computes the ChannelMode for a given channelId.
     * @param channelId Unique identifier for a channel.
     */
    function _mode(bytes32 channelId) internal view returns (ChannelMode) {
        // Note that _unpackStatus(someRandomChannelId) returns (0,0,0), which is
        // correct when nobody has written to storage yet.

        (, uint48 finalizesAt, ) = _unpackStatus(channelId);
        if (finalizesAt == 0) {
            return ChannelMode.Open;
            // solhint-disable-next-line not-rely-on-time
        } else if (finalizesAt <= block.timestamp) {
            return ChannelMode.Finalized;
        } else {
            return ChannelMode.Challenge;
        }
    }

    /**
     * @notice Formats the input data for on chain storage.
     * @dev Formats the input data for on chain storage.
     * @param channelData ChannelData data.
     */
    function _generateStatus(
        ChannelData memory channelData
    ) internal pure returns (bytes32 status) {
        // The hash is constructed from left to right.
        uint256 result;
        uint16 cursor = 256;

        // Shift turnNumRecord 208 bits left to fill the first 48 bits
        result = uint256(channelData.turnNumRecord) << (cursor -= 48);

        // logical or with finalizesAt padded with 160 zeros to get the next 48 bits
        result |= (uint256(channelData.finalizesAt) << (cursor -= 48));

        // logical or with the last 160 bits of the hash the remaining channelData fields
        // (we call this the fingerprint)
        result |= uint256(_generateFingerprint(channelData.stateHash, channelData.outcomeHash));

        status = bytes32(result);
    }

    function _generateFingerprint(
        bytes32 stateHash,
        bytes32 outcomeHash
    ) internal pure returns (uint160) {
        return uint160(uint256(keccak256(abi.encode(stateHash, outcomeHash))));
    }

    /**
     * @notice Unpacks turnNumRecord, finalizesAt and fingerprint from the status of a particular channel.
     * @dev Unpacks turnNumRecord, finalizesAt and fingerprint from the status of a particular channel.
     * @param channelId Unique identifier for a state channel.
     * @return turnNumRecord A turnNum that (the adjudicator knows) is supported by a signature from each participant.
     * @return finalizesAt The unix timestamp when `channelId` will finalize.
     * @return fingerprint The last 160 bits of kecca256(stateHash, outcomeHash)
     */
    function _unpackStatus(
        bytes32 channelId
    ) internal view returns (uint48 turnNumRecord, uint48 finalizesAt, uint160 fingerprint) {
        bytes32 status = statusOf[channelId];
        uint16 cursor = 256;
        turnNumRecord = uint48(uint256(status) >> (cursor -= 48));
        finalizesAt = uint48(uint256(status) >> (cursor -= 48));
        fingerprint = uint160(uint256(status));
    }
}
