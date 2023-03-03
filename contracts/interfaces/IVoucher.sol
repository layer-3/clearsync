// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

/**
 * @notice Interface describing Voucher for redeeming game items
 *
 * @dev The Voucher type must have a strict implementation on backend
 *
 * A Voucher is a document signed from the server IssuerKey and allows the execution
 * of actions on the game generally for creating game items, such as Booster Packs, Meld or reward tokens
 *
 */
interface IVoucher {
	/**
	 * @dev Build and encode the Voucher from server side
	 *
	 * Voucher structure will be valid only in chainId until expire timestamp
	 * the beneficiary MUST be the same as the user redeeming the Voucher.
	 *
	 */
	struct Voucher {
		address target; // contract address which the voucher is meant for
		uint8 action; // voucher type defined by the implementation
		address beneficiary; // beneficiary account which voucher will redeem to
		address referrer; // address of the parent
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId; // chain id of the voucher
		bytes32 voucherCodeHash; // hash of voucherCode
		bytes encodedParams; // voucher type specific encoded params
	}

	/**
	 * @notice Use vouchers that were issued and signed by the Back-end to receive game items.
	 * @param vouchers Vouchers issued by the Back-end.
	 * @param signature Vouchers signed by the Back-end.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external;

	/**
	 * @notice Use the voucher that was signed by the Back-end to receive game items.
	 * @param voucher Voucher issued by the Back-end.
	 * @param signature Voucher signed by the Back-end.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external;
}
