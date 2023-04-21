// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @notice Interface for custody allowing users to deposit funds to their brokers.
 */
interface ICustody {
    enum Actions {
        deposit,
        withdraw
    }

    struct Payload {
        Actions action;
        uint64 nonce;
        address broker;
        address account;
        address asset;
        uint256 amount;
        uint256 chainId;
        uint256 expire;
    }

    /**
     * @notice Deposit assets with given payload from the caller. Emits `Deposited` event.
     * @param payload Deposit payload.
     * @param signatures Payload signed by the CoSigner service.
     */
    function deposit(
        Payload calldata payload,
        bytes calldata brokerSignature,
        bytes[] calldata signatures
    ) external payable;

    /**
     * @notice Withdraw assets with given payload to the destination specified in the payload. Emits `Withdrawn` event.
     * @param payload Withdraw payload.
     * @param signatures Payload signed by the CoSigner service.
     */
    function withdraw(
        Payload calldata payload,
        bytes calldata brokerSignature,
        bytes[] calldata signatures
    ) external payable;

    event Deposited(
        uint64 nonce,
        address indexed broker,
        address indexed account,
        address indexed asset,
        uint256 amount
    );

    event Withdrawn(
        uint64 nonce,
        address indexed broker,
        address indexed destination,
        address indexed asset,
        uint256 amount
    );
}