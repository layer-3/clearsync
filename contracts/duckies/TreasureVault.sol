// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import '../interfaces/IVoucher.sol';

/**
 * @title TreasureVault
 * @dev This contract allows users to deposit tokens into a vault and redeem vouchers for rewards.
 *
 * The vouchers can then be used to redeem rewards or to refer others to the platform. Referral commissions are paid out
 * to referrers of up to 5 levels deep. This contract also allows the issuer to set an authorized address for signing
 * vouchers and upgrading the contract.
 */
contract TreasureVault is IVoucher, Initializable, AccessControlUpgradeable, UUPSUpgradeable {
	using ECDSAUpgradeable for bytes32;

	error CircularReferrers(address target, address base);
	error VoucherAlreadyUsed(bytes32 voucherCodeHash);
	error InvalidVoucher(Voucher voucher);
	error InvalidRewardParams(RewardParams rewardParams);
	error InsufficientTokenBalance(address token, uint256 expected, uint256 actual);
	error IncorrectSigner(address expected, address actual);

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant TREASURY_ROLE = keccak256('TREASURY_ROLE');

	// Constants
	uint8 public constant REFERRAL_MAX_DEPTH = 5;
	uint8 internal constant _REFERRAL_PAYOUT_DIVIDER = 100;

	enum VoucherAction {
		Reward
	}

	struct RewardParams {
		address token; // address of the ERC20 token to pay out
		uint256 amount; // amount of token to be paid
		uint8[REFERRAL_MAX_DEPTH] commissions; // what percentage of the reward will referrer of the level specified get
	}

	// Referral tree child => parent
	mapping(address => address) internal _referrerOf;

	// Store the vouchers to avoid replay attacks
	mapping(bytes32 => bool) internal _usedVouchers;

	// Address signing vouchers
	address public issuer;

	// Affiliate is invited by referrer. Referrer receives a tiny part of their affiliate's voucher.
	event AffiliateRegistered(address affiliate, address referrer);

	// disallow calling implementation directly (not via proxy)
	/// @custom:oz-upgrades-unsafe-allow constructor
	constructor() {
		_disableInitializers();
	}

	function initialize() public initializer {
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);
	}

	// -------- Issuer --------

	/**
	 * @dev Sets the issuer address. This function can only be called by accounts with the DEFAULT_ADMIN_ROLE.
	 * @param account The address of the new issuer.
	 */
	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	// -------- Withdraw --------

	/**
	 * @dev Withdraws the specified token from the vault.
	 * @param tokenAddress The address of the token being withdrawn.
	 * @param beneficiary The address of the account receiving the amount.
	 * @param amount The amount of the token to be withdrawn.
	 */
	function withdraw(
		address tokenAddress,
		address beneficiary,
		uint256 amount
	) public onlyRole(TREASURY_ROLE) {
		if (amount == 0) revert InsufficientTokenBalance(tokenAddress, 1, 0);

		IERC20Upgradeable token = IERC20Upgradeable(tokenAddress);
		uint256 tokenBalance = token.balanceOf(address(this));

		if (amount > tokenBalance)
			revert InsufficientTokenBalance(tokenAddress, amount, tokenBalance);

		token.transfer(beneficiary, amount);
	}

	// -------- Referrers --------

	/**
	 * @dev Registers a referrer for a child address.
	 * @param child The child address to register the referrer for.
	 * @param parent The address of the parent referrer.
	 */
	function _registerReferrer(address child, address parent) internal {
		_referrerOf[child] = parent;
		emit AffiliateRegistered(child, parent);
	}

	/**
	 * @dev Checks if the target address is not a referrer of the base address.
	 * @param target The target address to check.
	 * @param base The base address to check against.
	 */
	function _requireNotReferrerOf(address target, address base) internal view {
		address curAccount = base;

		for (uint8 i = 0; i < REFERRAL_MAX_DEPTH; i++) {
			if (_referrerOf[curAccount] == target) revert CircularReferrers(target, base);
			curAccount = _referrerOf[curAccount];
		}
	}

	// -------- Vouchers --------

	/**
	 * @dev Uses multiple vouchers at once.
	 * @param vouchers An array of Voucher structs to be used.
	 * @param signature The signature used to verify the vouchers.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_useVoucher(vouchers[i]);
		}
	}

	/**
	 * @dev Uses a single voucher.
	 * @param voucher The Voucher struct to be used.
	 * @param signature The signature used to verify the voucher.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_useVoucher(voucher);
	}

	// -------- Internal --------

	function _useVoucher(Voucher memory voucher) internal {
		_requireValidVoucher(voucher);

		_usedVouchers[voucher.voucherCodeHash] = true;

		// check for circular reference and register referrer
		if (voucher.referrer != address(0)) {
			// provided beneficiary has a referrer
			if (voucher.referrer == msg.sender)
				revert CircularReferrers(msg.sender, voucher.referrer);

			// check if beneficiary is not a referrer of supplied referrer
			_requireNotReferrerOf(msg.sender, voucher.referrer);
			_registerReferrer(voucher.beneficiary, voucher.referrer);
		}

		// parse & process Voucher
		if (voucher.action == uint8(VoucherAction.Reward)) {
			RewardParams memory rewardParams = abi.decode(voucher.encodedParams, (RewardParams));

			// rewardParams checks
			if (rewardParams.token == address(0) || rewardParams.amount == 0)
				revert InvalidRewardParams(rewardParams);

			_performPayout(
				voucher.beneficiary,
				rewardParams.token,
				rewardParams.amount,
				rewardParams.commissions
			);
		} else {
			revert InvalidVoucher(voucher);
		}

		emit VoucherUsed(
			voucher.beneficiary,
			voucher.action,
			voucher.voucherCodeHash,
			voucher.chainId
		);
	}

	function _requireValidVoucher(Voucher memory voucher) internal view {
		if (_usedVouchers[voucher.voucherCodeHash])
			revert VoucherAlreadyUsed(voucher.voucherCodeHash);

		if (
			voucher.target != address(this) ||
			voucher.beneficiary != msg.sender ||
			block.timestamp > voucher.expire ||
			voucher.chainId != block.chainid
		) revert InvalidVoucher(voucher);
	}

	function _performPayout(
		address beneficiary,
		address tokenAddress,
		uint256 amount,
		uint8[REFERRAL_MAX_DEPTH] memory referrersPayouts
	) internal {
		// check sufficient Vault token balance and pay beneficiary
		IERC20Upgradeable voucherToken = IERC20Upgradeable(tokenAddress);
		_requireSufficientContractBalance(voucherToken, amount);
		voucherToken.transfer(beneficiary, amount);

		// pay referrers
		address currReferrer = _referrerOf[beneficiary];

		for (uint8 i = 0; i < REFERRAL_MAX_DEPTH && currReferrer != address(0); i++) {
			if (referrersPayouts[i] != 0) {
				uint256 referralAmount = (amount * referrersPayouts[i]) / _REFERRAL_PAYOUT_DIVIDER;

				_requireSufficientContractBalance(voucherToken, referralAmount);
				voucherToken.transfer(currReferrer, referralAmount);
			}

			currReferrer = _referrerOf[currReferrer];
		}
	}

	function _requireSufficientContractBalance(
		IERC20Upgradeable token,
		uint256 expected
	) internal view {
		uint256 actual = token.balanceOf(address(this));
		if (actual < expected) revert InsufficientTokenBalance(address(token), expected, actual);
	}

	function _requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) internal pure {
		address actualSigner = keccak256(encodedData).toEthSignedMessageHash().recover(signature);
		if (actualSigner != signer) revert IncorrectSigner(signer, actualSigner);
	}

	// -------- Upgrading --------

	function _authorizeUpgrade(
		address newImplementation
	) internal override(UUPSUpgradeable) onlyRole(UPGRADER_ROLE) {}
}
