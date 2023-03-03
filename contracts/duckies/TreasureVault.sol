// SPDX-License-Identifier: MIT
// TODO: change version to 0.8.19
pragma solidity 0.8.17;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import '../interfaces/IVoucher.sol';
import '../interfaces/IVault.sol';

contract TreasureVault is
    IVoucher,
    IVault,
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    using ECDSAUpgradeable for bytes32;

    error CircularReferrers(address target, address base);
    error VoucherAlreadyUsed(bytes32 voucherCodeHash);
    error InvalidVoucher(Voucher voucher);
    error InvalidRewardParams(RewardParams rewardParams);
    error InsufficientTokenBalance(
        address token,
        uint256 expected,
        uint256 actual
    );
    error IncorrectSigner(address expected, address actual);

    // Roles
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

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

    // child => parent
    mapping(address => address) internal _referrerOf;

    mapping(bytes32 => bool) internal _usedVouchers;

	mapping (address => mapping(address => uint256)) private _balances;

    address public issuer;

    // Affiliate is invited by referrer. Referrer receives a tiny part of their affiliate's voucher.
    event AffiliateRegistered(address affiliate, address referrer);
    event VoucherUsed(
        address wallet,
        uint8 VoucherAction,
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

	/**
     * @dev Deposits the specified token into the vault.
     * @param token The address of the token being deposited.
     * @param amount The amount of the token being deposited.
     */
    function deposit(address token, uint256 amount) public payable {
        if (token == address(0)) {
            require(msg.value > 0, "You must deposit a non-zero amount of Ether.");
            _balances[token][msg.sender] += msg.value;
        } else {
            require(amount > 0, "You must deposit a non-zero amount of tokens.");
            IERC20Upgradeable(token).transferFrom(msg.sender, address(this), amount);
            _balances[token][msg.sender] += amount;
        }
    }

    /**
     * @dev Withdraws the specified token from the vault.
     * @param token The address of the token being withdrawn.
     * @param amount The amount of the token to be withdrawn.
     */
    function withdraw(address token, uint256 amount) public {
        require(amount > 0, "You must withdraw a non-zero amount.");
        if(_balances[token][msg.sender] < amount) {
			revert InsufficientTokenBalance(token, amount, _balances[token][msg.sender]);
		}

        _balances[token][msg.sender] -= amount;
        if (token == address(0)) {
            payable(msg.sender).transfer(amount);
        } else {
            IERC20Upgradeable(token).transfer(msg.sender, amount);
        }
    }

    /**
     * @dev Returns the balance of the specified token for the caller.
     * @param token The address of the token being queried.
     * @return The balance of the token held by the caller.
     */
    function getBalance(address token) public view returns (uint256) {
        return _balances[token][msg.sender];
    }

    // -------- Referrers --------

    function _registerReferrer(address child, address parent) internal {
        _referrerOf[child] = parent;
        emit AffiliateRegistered(child, parent);
    }

    function _requireNotReferrerOf(address target, address base) internal view {
        address curAccount = base;

        for (uint8 i = 0; i < REFERRAL_MAX_DEPTH; i++) {
            if (_referrerOf[curAccount] == target)
                revert CircularReferrers(target, base);
            curAccount = _referrerOf[curAccount];
        }
    }

    // -------- Vouchers --------

    function useVouchers(
        Voucher[] calldata vouchers,
        bytes calldata signature
    ) external {
        _requireCorrectSigner(abi.encode(vouchers), signature, issuer);
        for (uint8 i = 0; i < vouchers.length; i++) {
            _useVoucher(vouchers[i]);
        }
    }

    function useVoucher(
        Voucher calldata voucher,
        bytes calldata signature
    ) external {
        _requireCorrectSigner(abi.encode(voucher), signature, issuer);
        _useVoucher(voucher);
    }

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
            RewardParams memory rewardParams = abi.decode(
                voucher.encodedParams,
                (RewardParams)
            );

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
        // check sufficient Garden token balance and pay beneficiary
        IERC20Upgradeable voucherToken = IERC20Upgradeable(tokenAddress);
        _requireSufficientContractBalance(voucherToken, amount);
        voucherToken.transfer(beneficiary, amount);

        // pay referrers
        address currReferrer = _referrerOf[beneficiary];

        for (
            uint8 i = 0;
            i < REFERRAL_MAX_DEPTH && currReferrer != address(0);
            i++
        ) {
            if (referrersPayouts[i] != 0) {
                uint256 referralAmount = (amount * referrersPayouts[i]) /
                    _REFERRAL_PAYOUT_DIVIDER;

                _requireSufficientContractBalance(voucherToken, referralAmount);
                voucherToken.transfer(currReferrer, referralAmount);
            }

            currReferrer = _referrerOf[currReferrer];
        }
    }

    // -------- Internal --------

    function _requireSufficientContractBalance(
        IERC20Upgradeable token,
        uint256 expected
    ) internal view {
        uint256 actual = token.balanceOf(address(this));
        if (actual < expected)
            revert InsufficientTokenBalance(address(token), expected, actual);
    }

    function _requireCorrectSigner(
        bytes memory encodedData,
        bytes memory signature,
        address signer
    ) internal pure {
        address actualSigner = keccak256(encodedData)
            .toEthSignedMessageHash()
            .recover(signature);
        if (actualSigner != signer)
            revert IncorrectSigner(signer, actualSigner);
    }

    // -------- Upgrading --------

    function _authorizeUpgrade(
        address newImplementation
    ) internal override(UUPSUpgradeable) onlyRole(UPGRADER_ROLE) {}

}
