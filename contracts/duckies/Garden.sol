// SPDX-License-Identifier: MIT
// TODO: change version to 0.8.18
pragma solidity ^0.8.4;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20CappedUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import './DuckiesNFT.sol';

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
	uint8 internal constant _REFERRAL_PAYOUT_DIVIDER = 100;

	enum BountyType {
		TransferBountyToken, // encodedData: abi.encode(address, uint256, uint8[5]) - encoded token address, amount to pay, referrer payouts
		MintDucklingsNFT // encodedData: 0x (empty). May be extended in the future.
	}

	// Bounty Message for signature verification
	struct Bounty {
		BountyType type_;
		bytes encodedData; // specific to BountyType
		address beneficiary; // beneficiary of bounty
		address referrer; // address of the parent
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId;
		bytes32 bountyCodeHash;
	}

	// child => parent
	mapping(address => address) internal _referrerOf;

	mapping(bytes32 => bool) internal _claimedBounties;

	address public _issuer;

	DuckiesNFT public _ducklingsNFT;

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
	}

	// -------- Issuer, DucklingsNFT --------

	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		_issuer = account;
	}

	function setDucklingsNFTAddress(
		address ducklingsNFTAddress
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		_ducklingsNFT = DuckiesNFT(ducklingsNFTAddress);
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

		_claimedBounties[bounty.bountyCodeHash] = true;

		if (bounty.type_ == BountyType.TransferBountyToken) {
			_claimTransferBountyToken();
		} else if (bounty.type_ == BountyType.MintDucklingsNFT) {
			_claimMintDucklingsNFT();
		} else {
			revert InvalidBounty(bounty);
		}

		// event
		emit BountyClaimed(
			bounty.beneficiary,
			bounty.bountyCodeHash,
			bounty.chainId,
			bounty.tokenAddress
		);
	}

	function _claimTransferBountyToken() internal {
		// // check Garden has enough tokens and pay bounty to beneficiary
		// ERC20Upgradeable bountyToken = ERC20Upgradeable(tokenAddress);
		// _requireSufficientContractBalance(bountyToken, amount);
		// bountyToken.transfer(beneficiary, amount);
		// // register referrer is it is not a circular one
		// // provided beneficiary has a referrer
		// if (bounty.referrer != address(0)) {
		// 	if (bounty.referrer == msg.sender) revert InvalidBounty(bounty);
		// 	// check if beneficiary is not a referrer of supplied referrer
		// 	_requireNotReferrerOf(msg.sender, bounty.referrer);
		// 	_registerReferrer(beneficiary, bounty.referrer);
		// }
		// // pay referrers if specified
		// address currReferrer = _referrerOf[beneficiary];
		// for (uint8 i = 0; i < REFERRAL_MAX_DEPTH && currReferrer != address(0); i++) {
		// 	if (referrersPayouts[i] != 0) {
		// 		uint256 referralAmount = (amount * referrersPayouts[i]) / _REFERRAL_PAYOUT_DIVIDER;
		// 		_requireSufficientContractBalance(bountyToken, referralAmount);
		// 		bountyToken.transfer(currReferrer, referralAmount);
		// 	}
		// 	currReferrer = _referrerOf[currReferrer];
		// }
	}

	function _claimMintDucklingsNFT() internal {}

	function _requireValidBounty(Bounty memory bounty) internal view {
		if (_claimedBounties[bounty.bountyCodeHash])
			revert BountyAlreadyClaimed(bounty.bountyCodeHash);

		if (
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
