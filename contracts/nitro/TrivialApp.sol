// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IForceMoveApp} from './interfaces/IForceMoveApp.sol';

/**
 * @dev The TrivialApp contracts complies with the ForceMoveApp interface and allows all transitions, regardless of the data. Used for testing purposes.
 */
contract TrivialApp is IForceMoveApp {
    /**
     * @notice Encodes trivial rules.
     * @dev Encodes trivial rules.
     */
    function stateIsSupported(
        FixedPart calldata, // fixedPart, unused
        RecoveredVariablePart[] calldata, // proof, unused
        RecoveredVariablePart calldata // candidate, unused
    ) external pure override returns (bool, string memory) {
        return (true, '');
    }
}
