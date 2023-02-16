// SPDX-License-Identifier: MIT
// TODO: change version to 0.8.18
pragma solidity ^0.8.4;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20CappedUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

contract Garden is Initializable, AccessControlUpgradeable, UUPSUpgradeable {
	using ECDSAUpgradeable for bytes32;

	error CircularReferrers(address target, address base);
	error BountyAlreadyClaimed(bytes32 bountyCodeHash);
	error InvalidBounty(Bounty bounty);
	error InsufficientTokenBalance(address token, uint256 expected, uint256 actual);
	error IncorrectSigner(address expected, address actual);

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');

	// Constants
	uint8 public constant REFERRAL_MAX_DEPTH = 5;
	uint256 private _DUCKIES_SUPPLY_CAP;
	uint8 private constant _MAX_HALVING_STEP = 5;
	uint8 internal constant _REFERRAL_PAYOUT_DIVIDER = 100;

	// Bounty Message for signature verification
	struct Bounty {
		uint256 amount;
		address tokenAddress;
		address beneficiary; // beneficiary of bounty
		bool isPaidToReferrers; // whether bounty is payed to referrers
		address referrer; // address of the parent
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId;
		bytes32 bountyCodeHash;
	}

	// child => parent
	mapping(address => address) private _referrerOf;

	uint16[REFERRAL_MAX_DEPTH] internal _baseReferralPayouts;

	ERC20CappedUpgradeable internal _duckies;

	address private _issuer;

	mapping(bytes32 => mapping(uint32 => bool)) private _claimedBounties;

	// Affiliate is invited by referrer. Referrer receives a tiny part of their affiliate's bounty.
	event AffiliateRegistered(address affiliate, address referrer);
	event BountyClaimed(
		address wallet,
		bytes32 bountyCodeHash,
		uint32 chainId,
		address tokenAddress
	);

	// disallow calling implementation directly (not via proxy)
	/// @custom:oz-upgrades-unsafe-allow constructor
	constructor() {
		_disableInitializers();
	}

	function initialize(address duckies) public initializer {
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);

		_baseReferralPayouts = [500, 125, 80, 50, 20];

		_duckies = ERC20CappedUpgradeable(duckies);
		_DUCKIES_SUPPLY_CAP = _duckies.cap();
	}

	// -------- Issuer --------

	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		_issuer = account;
	}

	function getIssuer() external view returns (address) {
		return _issuer;
	}

	// -------- Partner --------

	function transferTokenBalanceToPartner(
		address tokenAddress,
		address partner
	) public onlyRole(DEFAULT_ADMIN_ROLE) {
		ERC20Upgradeable token = ERC20Upgradeable(tokenAddress);

		uint256 contractTokenBalance = token.balanceOf(address(this));

		if (contractTokenBalance > 0) {
			token.transfer(partner, contractTokenBalance);
		}
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

	// -------- Payouts --------

	function getReferralPayouts() external view returns (uint16[REFERRAL_MAX_DEPTH] memory) {
		return _getReferralPayouts();
	}

	function setBaseReferralPayouts(
		uint8[REFERRAL_MAX_DEPTH] memory referralPayouts
	) public onlyRole(DEFAULT_ADMIN_ROLE) {
		_baseReferralPayouts = referralPayouts;
	}

	function _getReferralPayouts() internal view returns (uint16[REFERRAL_MAX_DEPTH] memory) {
		uint16[REFERRAL_MAX_DEPTH] memory referralPayouts = _baseReferralPayouts;
		uint8 halvingDivider = halving();

		if (halvingDivider >= 2 ** _MAX_HALVING_STEP) {
			// return constant literal to save gas. Array length depends on REFERRAL_MAX_DEPTH value.
			return [uint16(5), 4, 3, 2, 1];
		}

		for (uint i = 0; i < referralPayouts.length; i++) {
			referralPayouts[i] /= halvingDivider;
		}
		return referralPayouts;
	}

	// -------- Halving --------

	function halving() public view returns (uint8) {
		uint16 nextHalvingStep = 1;

		while (
			nextHalvingStep < _MAX_HALVING_STEP &&
			_duckies.totalSupply() >= _halvingThreshold(nextHalvingStep)
		) {
			nextHalvingStep++;
		}

		return uint8(2 ** (nextHalvingStep - 1));
	}

	function _halvingThreshold(uint16 halvingNum) internal view returns (uint256) {
		return _DUCKIES_SUPPLY_CAP - (_DUCKIES_SUPPLY_CAP / 2 ** halvingNum);
	}

	// -------- Bounties --------

	function claimBounties(Bounty[] calldata bounties, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(bounties), signature, _issuer);
		for (uint8 i = 0; i < bounties.length; i++) {
			_claimBounty(bounties[i]);
		}
	}

	function claimBounty(Bounty calldata bounty, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(bounty), signature, _issuer);
		_claimBounty(bounty);
	}

	function _claimBounty(Bounty memory bounty) internal {
		_requireValidBounty(bounty);

		_claimedBounties[bounty.bountyCodeHash][bounty.chainId] = true;

		ERC20Upgradeable bountyToken = ERC20Upgradeable(bounty.tokenAddress);

		_requireSufficientContractBalance(bountyToken, bounty.amount);
		bountyToken.transfer(bounty.beneficiary, bounty.amount);

		// provided beneficiary has a referrer
		if (bounty.referrer != address(0)) {
			if (bounty.referrer == msg.sender) revert InvalidBounty(bounty);

			// check if beneficiary is not a referrer of supplied referrer
			_requireNotReferrerOf(msg.sender, bounty.referrer);
			_registerReferrer(bounty.beneficiary, bounty.referrer);
		}

		if (bounty.isPaidToReferrers) {
			uint16[REFERRAL_MAX_DEPTH] memory referralRates = _getReferralPayouts();
			address currReferrer = _referrerOf[bounty.beneficiary];

			for (uint8 i = 0; i < REFERRAL_MAX_DEPTH && currReferrer != address(0); i++) {
				uint256 referralAmount = (bounty.amount * referralRates[i]) /
					_REFERRAL_PAYOUT_DIVIDER;

				_requireSufficientContractBalance(bountyToken, referralAmount);
				bountyToken.transfer(currReferrer, referralAmount);

				currReferrer = _referrerOf[currReferrer];
			}
		}

		emit BountyClaimed(
			bounty.beneficiary,
			bounty.bountyCodeHash,
			bounty.chainId,
			bounty.tokenAddress
		);
	}

	function _requireValidBounty(Bounty memory bounty) internal view {
		if (_claimedBounties[bounty.bountyCodeHash][bounty.chainId])
			revert BountyAlreadyClaimed(bounty.bountyCodeHash);

		if (
			bounty.amount == 0 ||
			bounty.beneficiary != msg.sender ||
			block.timestamp > bounty.expire ||
			bounty.chainId != block.chainid
		) revert InvalidBounty(bounty);
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
