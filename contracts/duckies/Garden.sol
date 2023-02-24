// SPDX-License-Identifier: MIT
// TODO: change version to 0.8.18
pragma solidity ^0.8.4;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import '../ducklings/DucklingsNFT.sol';

contract Garden is Initializable, AccessControlUpgradeable, UUPSUpgradeable {
	using ECDSAUpgradeable for bytes32;

	error CircularReferrers(address target, address base);
	error VoucherAlreadyClaimed(bytes32 voucherCodeHash);
	error InvalidVoucher(Voucher voucher);
	error InvalidPayoutParams(PayoutParams payoutParams);
	error InvalidMintNFTsParams(MintNFTsParams mintNFTsParams);
	error InvalidMeldNFTsParams(MeldNFTsParams meldNFTsParams);
	error InsufficientTokenBalance(address token, uint256 expected, uint256 actual);
	error IncorrectSigner(address expected, address actual);

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');

	// Constants
	uint8 public constant REFERRAL_MAX_DEPTH = 5;
	uint8 internal constant _REFERRAL_PAYOUT_DIVIDER = 100;

	enum VoucherType {
		Payout,
		MintNFTs,
		MeldNFTs
	}

	struct PayoutParams {
		address token; // address of token to pay out
		uint256 amount; // amount of token to pay out
		uint8[REFERRAL_MAX_DEPTH] referrersPayouts; // what percentage of bounty will referrer of the level specified get
	}

	// TODO: add collection support
	// TODO: add baseGenes support
	struct MintNFTsParams {
		uint8 collection; // collection index
		uint8 amount; // card pack size
		uint256 baseGene; // preset gene values (if any)
	}

	struct MeldNFTsParams {
		uint256[5] meldingTokenIds; // token Ids to meld
	}

	// Voucher Message for signature verification
	struct Voucher {
		VoucherType type_; // voucher type
		bytes encodedData; // voucher type specific encoded data
		address beneficiary; // beneficiary of voucher
		address referrer; // address of the parent
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId; // chain id of the voucher
		bytes32 voucherCodeHash; // hash of voucherCode
	}

	// child => parent
	mapping(address => address) internal _referrerOf;

	mapping(bytes32 => bool) internal _claimedVouchers;

	address public issuer;

	DucklingsNFT public ducklingsNFT;

	// Affiliate is invited by referrer. Referrer receives a tiny part of their affiliate's voucher.
	event AffiliateRegistered(address affiliate, address referrer);
	event VoucherClaimed(
		address wallet,
		VoucherType voucherType,
		bytes32 voucherCodeHash,
		uint32 chainId
	);

	// disallow calling implementation directly (not via proxy)
	/// @custom:oz-upgrades-unsafe-allow constructor
	constructor() {
		_disableInitializers();
	}

	function initialize(address ducklingsNFTAddress) public initializer {
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);

		ducklingsNFT = DucklingsNFT(ducklingsNFTAddress);
	}

	// -------- Issuer, DucklingsNFT --------

	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	function setDucklingsNFT(address ducklingsNFTAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
		ducklingsNFT = DucklingsNFT(ducklingsNFTAddress);
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

	// TODO: think if there can be damage if circular referrers more than in REFERRAL_MAX_DEPTH levels
	function _requireNotReferrerOf(address target, address base) internal view {
		address curAccount = base;

		for (uint8 i = 0; i < REFERRAL_MAX_DEPTH; i++) {
			if (_referrerOf[curAccount] == target) revert CircularReferrers(target, base);
			curAccount = _referrerOf[curAccount];
		}
	}

	// -------- Vouchers --------

	function claimVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_claimVoucher(vouchers[i]);
		}
	}

	function claimVoucher(Voucher calldata voucher, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_claimVoucher(voucher);
	}

	function _claimVoucher(Voucher memory voucher) internal {
		_requireValidVoucher(voucher);

		_claimedVouchers[voucher.voucherCodeHash] = true;

		if (voucher.type_ == VoucherType.Payout) {
			PayoutParams memory payoutParams = abi.decode(voucher.encodedData, (PayoutParams));

			// payoutParams checks
			if (payoutParams.token == address(0) || payoutParams.amount == 0)
				revert InvalidPayoutParams(payoutParams);

			_performPayout(
				voucher.beneficiary,
				payoutParams.token,
				payoutParams.amount,
				payoutParams.referrersPayouts
			);
		} else if (voucher.type_ == VoucherType.MintNFTs) {
			MintNFTsParams memory mintNFTsParams = abi.decode(
				voucher.encodedData,
				(MintNFTsParams)
			);

			// mintNFTsParams checks
			if (mintNFTsParams.amount == 0) revert InvalidMintNFTsParams(mintNFTsParams);

			ducklingsNFT.freeMintPackTo(voucher.beneficiary, mintNFTsParams.amount);
		} else if (voucher.type_ == VoucherType.MeldNFTs) {
			MeldNFTsParams memory meldNFTsParams = abi.decode(
				voucher.encodedData,
				(MeldNFTsParams)
			);

			// meldNFTsParams checks
			if (!_areIdsUnique(meldNFTsParams.meldingTokenIds))
				revert InvalidMeldNFTsParams(meldNFTsParams);

			ducklingsNFT.freeMeldOf(voucher.beneficiary, meldNFTsParams.meldingTokenIds);
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
			voucher.type_,
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

	function _areIdsUnique(uint256[5] memory ids) internal pure returns (bool) {
		for (uint8 i = 0; i < ids.length; i++) {
			for (uint8 j = i + 1; j < ids.length; j++) {
				if (ids[i] == ids[j]) {
					return false;
				}
			}
		}

		return true;
	}

	// -------- Upgrading --------

	function _authorizeUpgrade(
		address newImplementation
	) internal override(UUPSUpgradeable) onlyRole(UPGRADER_ROLE) {}
}
