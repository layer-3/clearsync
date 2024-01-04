// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {INitroTypes} from './INitroTypes.sol';

/**
 * @dev The IForceMoveApp interface calls for its children to implement an application-specific stateIsSupported function, defining the state machine of a ForceMove state channel DApp.
 */
interface IForceMoveApp is INitroTypes {
    /**
     * @notice Encodes application-specific rules for a particular ForceMove-compliant state channel. Must revert or return false when invalid support proof and a candidate are supplied.
     * @dev Depending on the application, it might be desirable to narrow the state mutability of an implementation to 'pure' to make security analysis easier.
     * @param fixedPart Fixed part of the state channel.
     * @param proof Array of recovered variable parts which constitutes a support proof for the candidate. May be omitted when `candidate` constitutes a support proof itself.
     * @param candidate Recovered variable part the proof was supplied for. Also may constitute a support proof itself.
     */
    function stateIsSupported(
        FixedPart calldata fixedPart,
        RecoveredVariablePart[] calldata proof,
        RecoveredVariablePart calldata candidate
    ) external view returns (bool, string memory);
}
