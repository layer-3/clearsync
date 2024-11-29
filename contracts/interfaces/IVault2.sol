// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title IVault
 * @notice Interface for a vault contract that allows users to deposit, withdraw, and check balances of tokens and ETH.
 */
interface IVault {
    /**
     * @notice Error thrown when the address supplied with the function call is invalid.
     */
    error InvalidAddress();

    /**
     * @notice Emitted when a user deposits tokens or ETH into the vault.
     * @param user The address of the user that deposited the tokens.
     * @param token The address of the token deposited or address(0) for ETH.
     * @param amount The amount of tokens or ETH deposited.
     */
    event Deposited(address indexed user, address indexed token, uint256 amount);

    /**
     * @notice Emitted when a user withdraws tokens or ETH from the vault.
     * @param user The address of the user that withdrew the tokens.
     * @param token The address of the token withdrawn or address(0) for ETH.
     * @param amount The amount of tokens or ETH withdrawn.
     */
    event Withdrawn(address indexed user, address indexed token, uint256 amount);

    /**
     * @notice Error thrown when the value supplied with the function call is incorrect.
     */
    error IncorrectValue();

    /**
     * @notice Error thrown when the user has insufficient balance to perform an action.
     * @param token The address of the token that user lacks.
     * @param required The amount of tokens that is required to perform the action.
     * @param available The amount of tokens that the user has.
     */
    error InsufficientBalance(address token, uint256 required, uint256 available);

    /**
     * @notice Error thrown when the transfer of Eth fails.
     */
    error NativeTransferFailed();

    /**
     * @dev Returns the balance of a specified token for a user.
     * @param user The address of the user.
     * @param token The address of the token. Use address(0) for ETH.
     * @return The balance of the specified token for the user.
     */
    function balanceOf(address token) external view returns (uint256);

    /**
     * @dev Returns the balances of multiple tokens for a user.
     * @param user The address of the user.
     * @param tokens The addresses of the tokens. Use address(0) for ETH.
     * @return The balances of the specified tokens for the user.
     */
    function balancesOfTokens(address[] calldata tokens) external view returns (uint256[] memory);

    /**
     * @dev Deposits a specified amount of tokens or ETH into the vault.
     * @param token The address of the token to deposit. Use address(0) for ETH.
     * @param amount The amount of tokens or ETH to deposit.
     * @param to The address to send the tokens to.
     */
    function deposit(address token, uint256 amount, address to) external payable;

    /**
     * @dev Withdraws a specified amount of tokens or ETH from the vault.
     * @param token The address of the token to withdraw. Use address(0) for ETH.
     * @param amount The amount of tokens or ETH to withdraw.
     * @param to The address to send the tokens to.
     */
    function withdraw(address token, uint256 amount, address to) external;
}
