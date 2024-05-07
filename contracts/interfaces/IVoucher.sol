// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.18;

/**
 * @notice Interface describing Voucher for redeeming rewards
 *
 * @dev The Voucher type must have a strict implementation on backend
 *
 * A Voucher is a document signed by the server's Issuer private key and allows the execution
 * of redeeming rewards actions on the blockchain.
 */
interface IVoucher {
	/**
	 * @notice Custom error specifying that voucher has already been used.
	 * @param voucherCodeHash Hash of the code of the voucher that has been used.
	 */
	error VoucherAlreadyUsed(bytes32 voucherCodeHash);

	/**
	 * @notice Custom error specifying that voucher has not passed general voucher checks and is invalid.
	 * @param voucher Voucher that is invalid.
	 */
	error InvalidVoucher(Voucher voucher);

	/**
	 * @notice Custom error specifying that the message was expected to be signed by `expected` address, but was signed by `actual`.
	 * @param expected Expected address to have signed the message.
	 * @param actual Actual address that has signed the message.
	 */
	error IncorrectSigner(address expected, address actual);

	/**
	 * @dev Build and encode the Voucher from server side
	 * Voucher structure will be valid only in chainId until expire timestamp
	 */
	struct Voucher {
		address target; // contract address which the voucher is meant for
		uint8 action; // voucher type defined by the implementation
		address beneficiary; // beneficiary account which voucher will redeem to
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId; // chain id of the voucher
		bytes32 voucherCodeHash; // hash of voucherCode
		bytes encodedParams; // voucher type specific encoded params
	}

	/**
	 * @notice Use vouchers that were issued and signed by the Back-end to receive rewards.
	 * @param vouchers Vouchers issued by the Back-end.
	 * @param signature Vouchers signed by the Back-end.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external;

	/**
	 * @notice Use the voucher that was signed by the Back-end to receive rewards.
	 * @param voucher Voucher issued by the Back-end.
	 * @param signature Voucher signed by the Back-end.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external;

	/**
	 * @notice Event specifying that a voucher has been used.
	 * @param wallet Wallet that used a voucher.
	 * @param action The action of the voucher used.
	 * @param voucherCodeHash The code hash of the voucher used.
	 * @param chainId Id of the chain the voucher was used on.
	 */
	event VoucherUsed(address wallet, uint8 action, bytes32 voucherCodeHash, uint32 chainId);
}
