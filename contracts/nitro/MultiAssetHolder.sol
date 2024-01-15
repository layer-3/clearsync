// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';
import {IERC20} from '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {IMultiAssetHolder} from './interfaces/IMultiAssetHolder.sol';
import {StatusManager} from './StatusManager.sol';

/**
@dev An implementation of the IMultiAssetHolder interface. The AssetHolder contract escrows ETH or tokens against state channels. It allows assets to be internally accounted for, and ultimately prepared for transfer from one channel to other channels and/or external destinations, as well as for guarantees to be reclaimed.
 */
contract MultiAssetHolder is IMultiAssetHolder, StatusManager {
    using SafeERC20 for IERC20;

    // *******
    // Storage
    // *******

    /**
     * holdings[asset][channelId] is the amount of asset held against channel channelId. 0 address implies ETH
     */
    mapping(address => mapping(bytes32 => uint256)) public holdings;

    // **************
    // External methods
    // **************

    /**
     * @notice Deposit ETH or erc20 tokens against a given channelId.
     * @dev Deposit ETH or erc20 tokens against a given channelId.
     * @param asset erc20 token address, or zero address to indicate ETH
     * @param channelId ChannelId to be credited.
     * @param expectedHeld The number of wei/tokens the depositor believes are _already_ escrowed against the channelId.
     * @param amount The intended number of wei/tokens to be deposited.
     */
    function deposit(
        address asset,
        bytes32 channelId,
        uint256 expectedHeld,
        uint256 amount
    ) external payable virtual override {
        require(!_isExternalDestination(channelId), 'Deposit to external destination');
        // this allows participants to reduce the wait between deposits, while protecting them from losing funds by depositing too early. Specifically it protects against the scenario:
        // 1. Participant A deposits
        // 2. Participant B sees A's deposit, which means it is now safe for them to deposit
        // 3. Participant B submits their deposit
        // 4. The chain re-orgs, leaving B's deposit in the chain but not A's
        uint256 held = holdings[asset][channelId];
        require(held == expectedHeld, 'held != expectedHeld');

        // require successful deposit before updating holdings (protect against reentrancy)
        if (asset == address(0)) {
            require(msg.value == amount, 'Incorrect msg.value for deposit');
        } else {
            IERC20(asset).safeTransferFrom(msg.sender, address(this), amount);
        }

        held += amount;

        holdings[asset][channelId] = held;
        emit Deposited(channelId, asset, held);
    }

    /**
     * @notice Transfers as many funds escrowed against `channelId` as can be afforded for a specific destination. Assumes no repeated entries.
     * @dev Transfers as many funds escrowed against `channelId` as can be afforded for a specific destination. Assumes no repeated entries.
     * @param assetIndex Will be used to slice the outcome into a single asset outcome.
     * @param fromChannelId Unique identifier for state channel to transfer funds *from*.
     * @param outcomeBytes The encoded Outcome of this state channel
     * @param stateHash The hash of the state stored when the channel finalized.
     * @param indices Array with each entry denoting the index of a destination to transfer funds to. An empty array indicates "all".
     */
    function transfer(
        uint256 assetIndex, // TODO consider a uint48?
        bytes32 fromChannelId,
        bytes memory outcomeBytes,
        bytes32 stateHash,
        uint256[] memory indices
    ) external override {
        (
            Outcome.SingleAssetExit[] memory outcome,
            address asset,
            uint256 initialAssetHoldings
        ) = _apply_transfer_checks(assetIndex, indices, fromChannelId, stateHash, outcomeBytes); // view

        (
            Outcome.Allocation[] memory newAllocations,
            ,
            Outcome.Allocation[] memory exitAllocations,
            uint256 totalPayouts
        ) = compute_transfer_effects_and_interactions(
                initialAssetHoldings,
                outcome[assetIndex].allocations,
                indices
            ); // pure, also performs checks

        _apply_transfer_effects(
            assetIndex,
            asset,
            fromChannelId,
            stateHash,
            outcome,
            newAllocations,
            initialAssetHoldings,
            totalPayouts
        );
        _apply_transfer_interactions(outcome[assetIndex], exitAllocations);
    }

    function _apply_transfer_checks(
        uint256 assetIndex,
        uint256[] memory indices,
        bytes32 channelId,
        bytes32 stateHash,
        bytes memory outcomeBytes
    )
        internal
        view
        returns (
            Outcome.SingleAssetExit[] memory outcome,
            address asset,
            uint256 initialAssetHoldings
        )
    {
        _requireIncreasingIndices(indices); // This assumption is relied on by compute_transfer_effects_and_interactions
        _requireChannelFinalized(channelId);
        _requireMatchingFingerprint(stateHash, keccak256(outcomeBytes), channelId);

        outcome = Outcome.decodeExit(outcomeBytes);
        asset = outcome[assetIndex].asset;
        initialAssetHoldings = holdings[asset][channelId];
    }

    function compute_transfer_effects_and_interactions(
        uint256 initialHoldings,
        Outcome.Allocation[] memory allocations,
        uint256[] memory indices
    )
        public
        pure
        returns (
            Outcome.Allocation[] memory newAllocations,
            bool allocatesOnlyZeros,
            Outcome.Allocation[] memory exitAllocations,
            uint256 totalPayouts
        )
    {
        // `indices == []` means "pay out to all"
        // Note: by initializing exitAllocations to be an array of fixed length, its entries are initialized to be `0`
        exitAllocations = new Outcome.Allocation[](
            indices.length > 0 ? indices.length : allocations.length
        );
        totalPayouts = 0;
        newAllocations = new Outcome.Allocation[](allocations.length);
        allocatesOnlyZeros = true; // switched to false if there is an item remaining with amount > 0
        uint256 surplus = initialHoldings; // tracks funds available during calculation
        uint256 k = 0; // indexes the `indices` array

        // loop over allocations and decrease surplus
        for (uint256 i = 0; i < allocations.length; i++) {
            // copy destination, allocationType and metadata parts
            newAllocations[i].destination = allocations[i].destination;
            newAllocations[i].allocationType = allocations[i].allocationType;
            newAllocations[i].metadata = allocations[i].metadata;
            // compute new amount part
            uint256 affordsForDestination = min(allocations[i].amount, surplus);
            if ((indices.length == 0) || ((k < indices.length) && (indices[k] == i))) {
                if (allocations[k].allocationType == uint8(Outcome.AllocationType.guarantee))
                    revert('cannot transfer a guarantee');
                // found a match
                // reduce the current allocationItem.amount
                newAllocations[i].amount = allocations[i].amount - affordsForDestination;
                // increase the relevant exit allocation
                exitAllocations[k] = Outcome.Allocation(
                    allocations[i].destination,
                    affordsForDestination,
                    allocations[i].allocationType,
                    allocations[i].metadata
                );
                totalPayouts += affordsForDestination;
                // move on to the next supplied index
                ++k;
            } else {
                newAllocations[i].amount = allocations[i].amount;
            }
            if (newAllocations[i].amount != 0) allocatesOnlyZeros = false;
            // decrease surplus by the current amount if possible, else surplus goes to zero
            surplus -= affordsForDestination;
        }
    }

    function _apply_transfer_effects(
        uint256 assetIndex,
        address asset,
        bytes32 channelId,
        bytes32 stateHash,
        Outcome.SingleAssetExit[] memory outcome,
        Outcome.Allocation[] memory newAllocations,
        uint256 initialHoldings,
        uint256 totalPayouts
    ) internal {
        // update holdings
        holdings[asset][channelId] -= totalPayouts;

        // store fingerprint of modified outcome
        outcome[assetIndex].allocations = newAllocations;
        _updateFingerprint(channelId, stateHash, keccak256(abi.encode(outcome)));

        // emit the information needed to compute the new outcome stored in the fingerprint
        emit AllocationUpdated(channelId, assetIndex, initialHoldings, holdings[asset][channelId]);
    }

    function _apply_transfer_interactions(
        Outcome.SingleAssetExit memory singleAssetExit,
        Outcome.Allocation[] memory exitAllocations
    ) internal {
        // create a new tuple to avoid mutating singleAssetExit
        _executeSingleAssetExit(
            Outcome.SingleAssetExit(
                singleAssetExit.asset,
                singleAssetExit.assetMetadata,
                exitAllocations
            )
        );
    }

    /**
     * @notice Reclaim moves money from a target channel back into a ledger channel which is guaranteeing it. The guarantee is removed from the ledger channel.
     * @dev Reclaim moves money from a target channel back into a ledger channel which is guaranteeing it. The guarantee is removed from the ledger channel.
     * @param reclaimArgs arguments used in the reclaim function. Used to avoid stack too deep error.
     */
    function reclaim(ReclaimArgs memory reclaimArgs) external override {
        (
            Outcome.SingleAssetExit[] memory sourceOutcome,
            Outcome.SingleAssetExit[] memory targetOutcome
        ) = _apply_reclaim_checks(reclaimArgs); // view

        Outcome.Allocation[] memory newSourceAllocations;
        {
            Outcome.Allocation[] memory sourceAllocations = sourceOutcome[
                reclaimArgs.sourceAssetIndex
            ].allocations;
            Outcome.Allocation[] memory targetAllocations = targetOutcome[
                reclaimArgs.targetAssetIndex
            ].allocations;
            newSourceAllocations = compute_reclaim_effects(
                sourceAllocations,
                targetAllocations,
                reclaimArgs.indexOfTargetInSource
            ); // pure
        }

        _apply_reclaim_effects(reclaimArgs, sourceOutcome, newSourceAllocations);
    }

    /**
     * @dev Checks that the source and target channels are finalized; that the supplied outcomes match the stored fingerprints; that the asset is identical in source and target. Computes and returns the decoded outcomes.
     */
    function _apply_reclaim_checks(
        ReclaimArgs memory reclaimArgs
    )
        internal
        view
        returns (
            Outcome.SingleAssetExit[] memory sourceOutcome,
            Outcome.SingleAssetExit[] memory targetOutcome
        )
    {
        (
            bytes32 sourceChannelId,
            bytes memory sourceOutcomeBytes,
            uint256 sourceAssetIndex,
            bytes memory targetOutcomeBytes,
            uint256 targetAssetIndex
        ) = (
                reclaimArgs.sourceChannelId,
                reclaimArgs.sourceOutcomeBytes,
                reclaimArgs.sourceAssetIndex,
                reclaimArgs.targetOutcomeBytes,
                reclaimArgs.targetAssetIndex
            );

        // source checks
        _requireChannelFinalized(sourceChannelId);
        _requireMatchingFingerprint(
            reclaimArgs.sourceStateHash,
            keccak256(sourceOutcomeBytes),
            sourceChannelId
        );

        sourceOutcome = Outcome.decodeExit(sourceOutcomeBytes);
        targetOutcome = Outcome.decodeExit(targetOutcomeBytes);
        address asset = sourceOutcome[sourceAssetIndex].asset;
        require(
            sourceOutcome[sourceAssetIndex]
                .allocations[reclaimArgs.indexOfTargetInSource]
                .allocationType == uint8(Outcome.AllocationType.guarantee),
            'not a guarantee allocation'
        );

        bytes32 targetChannelId = sourceOutcome[sourceAssetIndex]
            .allocations[reclaimArgs.indexOfTargetInSource]
            .destination;

        // target checks
        require(targetOutcome[targetAssetIndex].asset == asset, 'targetAsset != guaranteeAsset');
        _requireChannelFinalized(targetChannelId);
        _requireMatchingFingerprint(
            reclaimArgs.targetStateHash,
            keccak256(targetOutcomeBytes),
            targetChannelId
        );
    }

    /**
     * @dev Computes side effects for the reclaim function. Returns updated allocations for the source, computed by finding the guarantee in the source for the target, and moving money out of the guarantee and back into the ledger channel as regular allocations for the participants.
     */
    function compute_reclaim_effects(
        Outcome.Allocation[] memory sourceAllocations,
        Outcome.Allocation[] memory targetAllocations,
        uint256 indexOfTargetInSource
    ) public pure returns (Outcome.Allocation[] memory) {
        Outcome.Allocation[] memory newSourceAllocations = new Outcome.Allocation[](
            sourceAllocations.length - 1 // is one slot shorter as we remove the guarantee
        );

        Outcome.Allocation memory guarantee = sourceAllocations[indexOfTargetInSource];
        Guarantee memory guaranteeData = decodeGuaranteeData(guarantee.metadata);

        bool foundTarget = false;
        bool foundLeft = false;
        bool foundRight = false;
        uint256 totalReclaimed;

        uint256 k = 0;
        for (uint256 i = 0; i < sourceAllocations.length; i++) {
            if (i == indexOfTargetInSource) {
                foundTarget = true;
                continue;
            }
            newSourceAllocations[k] = Outcome.Allocation({
                destination: sourceAllocations[i].destination,
                amount: sourceAllocations[i].amount,
                allocationType: sourceAllocations[i].allocationType,
                metadata: sourceAllocations[i].metadata
            });

            if (!foundLeft && sourceAllocations[i].destination == guaranteeData.left) {
                newSourceAllocations[k].amount += targetAllocations[0].amount;
                totalReclaimed += targetAllocations[0].amount;
                foundLeft = true;
            }
            if (!foundRight && sourceAllocations[i].destination == guaranteeData.right) {
                newSourceAllocations[k].amount += targetAllocations[1].amount;
                totalReclaimed += targetAllocations[1].amount;
                foundRight = true;
            }
            k++;
        }

        require(foundTarget, 'could not find target');
        require(foundLeft, 'could not find left');
        require(foundRight, 'could not find right');
        require(totalReclaimed == guarantee.amount, 'totalReclaimed!=guarantee.amount');

        return newSourceAllocations;
    }

    /**
     * @dev Updates the fingerprint of the outcome for the source channel and emit an event for it.
     */
    function _apply_reclaim_effects(
        ReclaimArgs memory reclaimArgs,
        Outcome.SingleAssetExit[] memory sourceOutcome,
        Outcome.Allocation[] memory newSourceAllocations
    ) internal {
        (bytes32 sourceChannelId, uint256 sourceAssetIndex) = (
            reclaimArgs.sourceChannelId,
            reclaimArgs.sourceAssetIndex
        );

        // store fingerprint of modified source outcome
        sourceOutcome[sourceAssetIndex].allocations = newSourceAllocations;
        _updateFingerprint(
            sourceChannelId,
            reclaimArgs.sourceStateHash,
            keccak256(abi.encode(sourceOutcome))
        );

        // emit the information needed to compute the new source outcome stored in the fingerprint
        emit Reclaimed(reclaimArgs.sourceChannelId, reclaimArgs.sourceAssetIndex);

        // Note: no changes are made to the target channel.
    }

    /**
     * @notice Executes a single asset exit by paying out the asset and calling external contracts, as well as updating the holdings stored in this contract.
     * @dev Executes a single asset exit by paying out the asset and calling external contracts, as well as updating the holdings stored in this contract.
     * @param singleAssetExit The single asset exit to be paid out.
     */
    function _executeSingleAssetExit(Outcome.SingleAssetExit memory singleAssetExit) internal {
        address asset = singleAssetExit.asset;
        for (uint256 j = 0; j < singleAssetExit.allocations.length; j++) {
            bytes32 destination = singleAssetExit.allocations[j].destination;
            uint256 amount = singleAssetExit.allocations[j].amount;
            if (_isExternalDestination(destination)) {
                _transferAsset(asset, _bytes32ToAddress(destination), amount);
            } else {
                holdings[asset][destination] += amount;
            }
        }
    }

    /**
     * @notice Transfers the given amount of this AssetHolders's asset type to a supplied ethereum address.
     * @dev Transfers the given amount of this AssetHolders's asset type to a supplied ethereum address.
     * @param destination ethereum address to be credited.
     * @param amount Quantity of assets to be transferred.
     */
    function _transferAsset(address asset, address destination, uint256 amount) internal {
        if (asset == address(0)) {
            (bool success, ) = destination.call{value: amount}(''); //solhint-disable-line avoid-low-level-calls
            require(success, 'Could not transfer ETH');
        } else {
            IERC20(asset).transfer(destination, amount);
        }
    }

    /**
     * @notice Checks if a given destination is external (and can therefore have assets transferred to it) or not.
     * @dev Checks if a given destination is external (and can therefore have assets transferred to it) or not.
     * @param destination Destination to be checked.
     * @return True if the destination is external, false otherwise.
     */
    function _isExternalDestination(bytes32 destination) internal pure returns (bool) {
        return uint96(bytes12(destination)) == 0;
    }

    /**
     * @notice Converts an ethereum address to a nitro external destination.
     * @dev Converts an ethereum address to a nitro external destination.
     * @param participant The address to be converted.
     * @return The input address left-padded with zeros.
     */
    function _addressToBytes32(address participant) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(participant)));
    }

    /**
     * @notice Converts a nitro destination to an ethereum address.
     * @dev Converts a nitro destination to an ethereum address.
     * @param destination The destination to be converted.
     * @return The rightmost 160 bits of the input string.
     */
    function _bytes32ToAddress(bytes32 destination) internal pure returns (address payable) {
        return payable(address(uint160(uint256(destination))));
    }

    // **************
    // Requirers
    // **************

    /**
     * @notice Checks that a given variables hash to the data stored on chain.
     * @dev Checks that a given variables hash to the data stored on chain.
     */
    function _requireMatchingFingerprint(
        bytes32 stateHash,
        bytes32 outcomeHash,
        bytes32 channelId
    ) internal view {
        (, , uint160 fingerprint) = _unpackStatus(channelId);
        require(
            fingerprint == _generateFingerprint(stateHash, outcomeHash),
            'incorrect fingerprint'
        );
    }

    /**
     * @notice Checks that a given channel is in the Finalized mode.
     * @dev Checks that a given channel is in the Finalized mode.
     * @param channelId Unique identifier for a channel.
     */
    function _requireChannelFinalized(bytes32 channelId) internal view {
        require(_mode(channelId) == ChannelMode.Finalized, 'Channel not finalized.');
    }

    function _updateFingerprint(
        bytes32 channelId,
        bytes32 stateHash,
        bytes32 outcomeHash
    ) internal {
        (uint48 turnNumRecord, uint48 finalizesAt, ) = _unpackStatus(channelId);

        bytes32 newStatus = _generateStatus(
            ChannelData(turnNumRecord, finalizesAt, stateHash, outcomeHash)
        );
        statusOf[channelId] = newStatus;
    }

    /**
     * @notice Checks that the supplied indices are strictly increasing.
     * @dev Checks that the supplied indices are strictly increasing. This allows us allows us to write a more efficient claim function.
     */
    function _requireIncreasingIndices(uint256[] memory indices) internal pure {
        for (uint256 i = 0; i + 1 < indices.length; i++) {
            require(indices[i] < indices[i + 1], 'Indices must be sorted');
        }
    }

    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? b : a;
    }

    struct Guarantee {
        bytes32 left;
        bytes32 right;
    }

    function decodeGuaranteeData(bytes memory data) internal pure returns (Guarantee memory) {
        return abi.decode(data, (Guarantee));
    }
}
