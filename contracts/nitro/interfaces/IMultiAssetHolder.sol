// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @dev The IMultiAssetHolder interface calls for functions that allow assets to be transferred from one channel to other channel and/or external destinations, as well as for guarantees to be claimed.
 */
interface IMultiAssetHolder {
    /**
     * @notice Deposit ETH or erc20 assets against a given destination.
     * @dev Deposit ETH or erc20 assets against a given destination.
     * @param asset erc20 token address, or zero address to indicate ETH
     * @param destination ChannelId to be credited.
     * @param expectedHeld The number of wei the depositor believes are _already_ escrowed against the channelId.
     * @param amount The intended number of wei to be deposited.
     */
    function deposit(
        address asset,
        bytes32 destination,
        uint256 expectedHeld,
        uint256 amount
    ) external payable;

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
    ) external;

    /**
     * @param sourceChannelId Id of a ledger channel containing a guarantee.
     * @param sourceStateHash Hash of the state stored when the source channel finalized.
     * @param sourceOutcomeBytes The abi.encode of source channel outcome
     * @param sourceAssetIndex the index of the targetted asset in the source outcome.
     * @param indexOfTargetInSource The index of the guarantee allocation to the target channel in the source outcome.
     * @param targetStateHash Hash of the state stored when the target channel finalized.
     * @param targetOutcomeBytes The abi.encode of target channel outcome
     * @param targetAssetIndex the index of the targetted asset in the target outcome.
     */
    struct ReclaimArgs {
        bytes32 sourceChannelId;
        bytes32 sourceStateHash;
        bytes sourceOutcomeBytes;
        uint256 sourceAssetIndex;
        uint256 indexOfTargetInSource;
        bytes32 targetStateHash;
        bytes targetOutcomeBytes;
        uint256 targetAssetIndex;
    }

    /**
     * @notice Reclaim moves money from a target channel back into a ledger channel which is guaranteeing it. The guarantee is removed from the ledger channel.
     * @dev Reclaim moves money from a target channel back into a ledger channel which is guaranteeing it. The guarantee is removed from the ledger channel.
     * @param reclaimArgs arguments used in the claim function. Used to avoid stack too deep error.
     */
    function reclaim(ReclaimArgs memory reclaimArgs) external;

    /**
     * @dev Indicates that a deposit has been made.
     * @param destination The channel being deposited into.
     * @param asset The asset being deposited. Zero address indicates the native token (e.g. ETH).
     * @param destinationHoldings The new holdings for `destination`.
     */
    event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings);

    /**
     * @dev Indicates the assetOutcome for this channelId and assetIndex has changed due to a transfer. Includes sufficient data to compute:
     * - the new assetOutcome
     * - the new holdings for this channelId and any others that were transferred to
     * - the payouts to external destinations
     * when combined with the calldata of the transaction causing this event to be emitted.
     * @param channelId The channelId of the funds being withdrawn.
     * @param initialHoldings holdings[asset][channelId] **before** the allocations were updated. The asset in question can be inferred from the calldata of the transaction (it might be "all assets")
     * @param finalHoldings holdings[asset][channelId] **after** the allocations are updated
     */
    event AllocationUpdated(
        bytes32 indexed channelId,
        uint256 assetIndex,
        uint256 initialHoldings,
        uint256 finalHoldings
    );

    /**
     * @dev Indicates the assetOutcome for this channelId and assetIndex has changed due to a reclaim. Includes sufficient data to compute:
     * - the new assetOutcome
     * when combined with the calldata of the transaction causing this event to be emitted.
     * @param channelId The channelId of the funds being withdrawn.
     */
    event Reclaimed(bytes32 indexed channelId, uint256 assetIndex);
}
