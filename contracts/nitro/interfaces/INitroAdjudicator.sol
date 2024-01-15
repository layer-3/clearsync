// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';
import {IMultiAssetHolder} from './IMultiAssetHolder.sol';
import {IForceMove} from './IForceMove.sol';

/**
 * @dev The INitroAdjudicator defines an interface for a contract that adjudicates state channels. It is based on IMultiAssetHolder and IForceMove, extending them with some functions.
 */
interface INitroAdjudicator is IMultiAssetHolder, IForceMove {
    /**
     * @notice Finalizes a channel according to the given candidate, and liquidates all assets for the channel.
     * @dev Finalizes a channel according to the given candidate, and liquidates all assets for the channel.
     * @param fixedPart Data describing properties of the state channel that do not change with state updates.
     * @param candidate Variable part of the state to change to.
     */
    function concludeAndTransferAllAssets(
        FixedPart memory fixedPart,
        SignedVariablePart memory candidate
    ) external;

    /**
     * @notice Liquidates all assets for the channel
     * @dev Liquidates all assets for the channel
     * @param channelId Unique identifier for a state channel
     * @param outcome An array of SingleAssetExit[] items.
     * @param stateHash stored state hash for the channel
     */
    function transferAllAssets(
        bytes32 channelId,
        Outcome.SingleAssetExit[] memory outcome,
        bytes32 stateHash
    ) external;

    /**
     * @notice Checks whether an application-specific rules for a particular ForceMove-compliant state channel are enforced in supplied states.
     * @dev Checks whether an application-specific rules for a particular ForceMove-compliant state channel are enforced in supplied states.
     * @param fixedPart Fixed part of the state channel.
     * @param proof Variable parts of the states with signatures in the support proof. The proof is a validation for the supplied candidate.
     * @param candidate Variable part of the state to change to. The candidate state is supported by proof states.
     */
    function stateIsSupported(
        FixedPart calldata fixedPart,
        SignedVariablePart[] calldata proof,
        SignedVariablePart calldata candidate
    ) external view returns (bool, string memory);
}
