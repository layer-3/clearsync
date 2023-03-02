// SPDX-License-Identifier: MIT
// TODO: change version to 0.8.19
pragma solidity 0.8.17;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import '../../interfaces/IVoucher.sol';

contract Garden is IVoucher, Initializable, AccessControlUpgradeable, UUPSUpgradeable {
	using ECDSAUpgradeable for bytes32;

	error CircularReferrers(address target, address base);
	error VoucherAlreadyClaimed(bytes32 voucherCodeHash);
	error InvalidVoucher(Voucher voucher);
	error InvalidRewardParams(RewardParams rewardParams);
	error InsufficientTokenBalance(address token, uint256 expected, uint256 actual);
	error IncorrectSigner(address expected, address actual);

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');

	// Constants
	uint8 public constant REFERRAL_MAX_DEPTH = 5;
	uint8 internal constant _REFERRAL_PAYOUT_DIVIDER = 100;

	enum VoucherType {
		Reward
	}

	struct RewardParams {
		address token; // address of the ERC20 token to pay out
		uint256 amount; // amount of token to be paid
		uint8[REFERRAL_MAX_DEPTH] commissions; // what percentage of the reward will referrer of the level specified get
	}

	// child => parent
	mapping(address => address) internal _referrerOf;

	mapping(bytes32 => bool) internal _claimedVouchers;

	address public issuer;

	// Affiliate is invited by referrer. Referrer receives a tiny part of their affiliate's voucher.
	event AffiliateRegistered(address affiliate, address referrer);
	event VoucherClaimed(
		address wallet,
		uint8 voucherType,
		bytes32 voucherCodeHash,
		uint32 chainId
	);

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

	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	// -------- Partner --------

	function transferTokenBalanceToPartner(
		address tokenAddress,
		address partner
	) public onlyRole(DEFAULT_ADMIN_ROLE) {
		ERC20Upgradeable token = ERC20Upgradeable(tokenAddress);

		uint256 contractTokenBalance = token.balanceOf(address(this));

		if (contractTokenBalance == 0) revert InsufficientTokenBalance(tokenAddress, 1, 0);

		token.transfer(partner, contractTokenBalance);
	}

	// -------- Referrers --------

	function _registerReferrer(address child, address parent) internal {
		_referrerOf[child] = parent;
		emit AffiliateRegistered(child, parent);
	}

	function _requireNotReferrerOf(address target, address base) internal view {
		address curAccount = base;

		for (uint8 i = 0; i < REFERRAL_MAX_DEPTH; i++) {
			if (_referrerOf[curAccount] == target) revert CircularReferrers(target, base);
			curAccount = _referrerOf[curAccount];
		}
	}

	// -------- Vouchers --------

	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_useVoucher(vouchers[i]);
		}
	}

	function useVoucher(Voucher calldata voucher, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_useVoucher(voucher);
	}

	function _useVoucher(Voucher memory voucher) internal {
		_requireValidVoucher(voucher);

		_claimedVouchers[voucher.voucherCodeHash] = true;

		if (voucher.action == uint8(VoucherType.Reward)) {
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

		// check for circular reference and register referrer
		if (voucher.referrer != address(0)) {
			// provided beneficiary has a referrer
			if (voucher.referrer == msg.sender) revert InvalidVoucher(voucher);

			// check if beneficiary is not a referrer of supplied referrer
			_requireNotReferrerOf(msg.sender, voucher.referrer);
			_registerReferrer(voucher.beneficiary, voucher.referrer);
		}

		emit VoucherClaimed(
			voucher.beneficiary,
			voucher.action,
			voucher.voucherCodeHash,
			voucher.chainId
		);
	}

	function _requireValidVoucher(Voucher memory voucher) internal view {
		if (_claimedVouchers[voucher.voucherCodeHash])
			revert VoucherAlreadyClaimed(voucher.voucherCodeHash);

		if (
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
		// check sufficient Garden token balance and pay beneficiary
		ERC20Upgradeable voucherToken = ERC20Upgradeable(tokenAddress);
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

	// -------- Internal --------

	function _requireSufficientContractBalance(
		ERC20Upgradeable token,
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
