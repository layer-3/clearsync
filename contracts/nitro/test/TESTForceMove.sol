// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ForceMove} from '../ForceMove.sol';

/**
 * @dev This contract extends the ForceMove contract to enable it to be more easily unit-tested. It exposes public or external functions that set storage variables or wrap otherwise internal functions. It should not be deployed in a production environment.
 */
contract TESTForceMove is ForceMove {
    // Public wrappers for internal methods:

    /**
     * @dev Wrapper for otherwise internal function. Tests whether a given address is in a given array of addresses.
     * @param suspect A single address of interest.
     * @param addresses A line-up of possible perpetrators.
     * @return true if the address is in the array, false otherwise
     */
    function isAddressInArray(
        address suspect,
        address[] memory addresses
    ) public pure returns (bool) {
        return _isAddressInArray(suspect, addresses);
    }

    // public setter for statusOf

    /**
     * @dev Manually set the fingerprint for a given channelId.  Shortcuts the public methods (ONLY USE IN A TESTING ENVIRONMENT).
     * @param channelId Unique identifier for a state channel.
     * @param channelData The channelData to be formatted and stored against the channelId
     */
    function setStatusFromChannelData(bytes32 channelId, ChannelData memory channelData) public {
        if (channelData.finalizesAt == 0) {
            require(
                channelData.stateHash == bytes32(0) && channelData.outcomeHash == bytes32(0),
                'Invalid open channel'
            );
        }

        statusOf[channelId] = _generateStatus(channelData);
    }

    /**
     * @dev Manually set the fingerprint for a given channelId.  Shortcuts the public methods (ONLY USE IN A TESTING ENVIRONMENT).
     * @param channelId Unique identifier for a state channel.
     * @param f The fingerprint to store against the channelId
     */
    function setStatus(bytes32 channelId, bytes32 f) public {
        statusOf[channelId] = f;
    }

    /**
     * @dev Wrapper for otherwise internal function. Hashes the input data and formats it for on chain storage.
     * @param channelData ChannelData data.
     */
    function generateStatus(
        ChannelData memory channelData
    ) public pure returns (bytes32 newStatus) {
        return _generateStatus(channelData);
    }

    /**
     * @dev Wrapper for otherwise internal function. Checks that a given ChannelData struct matches a supplied bytes32 when formatted for storage.
     * @param cs A given ChannelData data structure.
     * @param f Some data in on-chain storage format.
     */
    function matchesStatus(ChannelData memory cs, bytes32 f) public pure returns (bool) {
        return _matchesStatus(cs, f);
    }

    /**
     * @dev Wrapper for otherwise internal function. Checks that a given channel is in the Challenge mode.
     * @param channelId Unique identifier for a channel.
     */
    function requireChannelOpen(bytes32 channelId) public view {
        _requireChannelOpen(channelId);
    }
}
