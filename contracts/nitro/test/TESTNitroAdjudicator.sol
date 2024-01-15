// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {NitroAdjudicator} from '../NitroAdjudicator.sol';
import {TESTForceMove} from './TESTForceMove.sol';

/**
 * @dev This contract extends the NitroAdjudicator contract to enable it to be more easily unit-tested. It exposes public or external functions that set storage variables or wrap otherwise internal functions. It should not be deployed in a production environment.
 */
contract TESTNitroAdjudicator is NitroAdjudicator, TESTForceMove {
    /**
     * @dev Manually set the holdings mapping to a given amount for a given channelId.  Shortcuts the deposit workflow (ONLY USE IN A TESTING ENVIRONMENT)
     * @param channelId Unique identifier for a state channel.
     * @param amount The number of assets that should now be "escrowed: against channelId
     */
    function setHoldings(address asset, bytes32 channelId, uint256 amount) external {
        holdings[asset][channelId] = amount;
    }

    /**
     * @dev Wrapper for otherwise internal function. Checks if a given destination is external (and can therefore have assets transferred to it) or not.
     * @param destination Destination to be checked.
     * @return True if the destination is external, false otherwise.
     */
    function isExternalDestination(bytes32 destination) public pure returns (bool) {
        return _isExternalDestination(destination);
    }

    /**
     * @dev Wrapper for otherwise internal function. Converts an ethereum address to a nitro external destination.
     * @param participant The address to be converted.
     * @return The input address left-padded with zeros.
     */
    function addressToBytes32(address participant) public pure returns (bytes32) {
        return _addressToBytes32(participant);
    }
}
