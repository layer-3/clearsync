// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

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

    enum VoucherType {
		Reward,
		MintPacks,
		MeldFlock
	}

	struct RewardParams {
		address token;          // address of the ERC20 token to payout
		uint256 amount;         // amount of token to be pay
		uint8[5] commissions;   // what percentage of bounty will referrer of the level specified get
	}

	struct MintPacksParams {
		address token;      // address of ERC721 token to mint
		uint8 collection;   // collection index
		uint8 size;         // card booster pack size
	}

	struct MeldFlockParams {
		uint256[5] meldingTokenIds; // token Ids to meld
	}

    /**
     * @dev Build and encode the Voucher from server side
     *
     * Voucher structure will be valid only in chainId until expire timestamp
     * the beneficiary MUST be the same as the user redeeming the Voucher.
     * 
     */
	struct Voucher {
		VoucherType action;             // voucher type
		address     beneficiary;        // beneficiary of voucher
		address     referrer;           // address of the parent
		uint64      expire;             // expiration time in seconds UTC
		uint32      chainId;            // chain id of the voucher
		bytes32     voucherCodeHash;    // hash of voucherCode
		bytes       encodedParams;      // voucher type specific encoded params
	}
}