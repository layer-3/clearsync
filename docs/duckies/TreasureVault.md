# Solidity API

## TreasureVault

\_This contract allows users to deposit tokens into a vault and redeem vouchers for rewards.

The vouchers can then be used to redeem rewards or to refer others to the platform. Referral commissions are paid out
to referrers of up to 5 levels deep. This contract also allows the issuer to set an authorized address for signing
vouchers and upgrading the contract.\_

### CircularReferrers

```solidity
error CircularReferrers(address target, address base)
```

### VoucherAlreadyUsed

```solidity
error VoucherAlreadyUsed(bytes32 voucherCodeHash)
```

### InvalidVoucher

```solidity
error InvalidVoucher(struct IVoucher.Voucher voucher)
```

### InvalidRewardParams

```solidity
error InvalidRewardParams(struct TreasureVault.RewardParams rewardParams)
```

### InsufficientTokenBalance

```solidity
error InsufficientTokenBalance(address token, uint256 expected, uint256 actual)
```

### IncorrectSigner

```solidity
error IncorrectSigner(address expected, address actual)
```

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

### REFERRAL_MAX_DEPTH

```solidity
uint8 REFERRAL_MAX_DEPTH
```

### \_REFERRAL_PAYOUT_DIVIDER

```solidity
uint8 _REFERRAL_PAYOUT_DIVIDER
```

### VoucherAction

```solidity
enum VoucherAction {
  Reward
}
```

### RewardParams

```solidity
struct RewardParams {
  address token;
  uint256 amount;
  uint8[5] commissions;
}
```

### \_referrerOf

```solidity
mapping(address => address) _referrerOf
```

### \_usedVouchers

```solidity
mapping(bytes32 => bool) _usedVouchers
```

### \_balances

```solidity
mapping(address => mapping(address => uint256)) _balances
```

### issuer

```solidity
address issuer
```

### AffiliateRegistered

```solidity
event AffiliateRegistered(address affiliate, address referrer)
```

### VoucherUsed

```solidity
event VoucherUsed(address wallet, uint8 VoucherAction, bytes32 voucherCodeHash, uint32 chainId)
```

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize() public
```

### setIssuer

```solidity
function setIssuer(address account) external
```

_Sets the issuer address. This function can only be called by accounts with the DEFAULT_ADMIN_ROLE._

#### Parameters

| Name    | Type    | Description                    |
| ------- | ------- | ------------------------------ |
| account | address | The address of the new issuer. |

### deposit

```solidity
function deposit(address token, uint256 amount) public payable
```

_Deposits the specified token into the vault._

#### Parameters

| Name   | Type    | Description                               |
| ------ | ------- | ----------------------------------------- |
| token  | address | The address of the token being deposited. |
| amount | uint256 | The amount of the token being deposited.  |

### withdraw

```solidity
function withdraw(address token, uint256 amount) public
```

_Withdraws the specified token from the vault._

#### Parameters

| Name   | Type    | Description                               |
| ------ | ------- | ----------------------------------------- |
| token  | address | The address of the token being withdrawn. |
| amount | uint256 | The amount of the token to be withdrawn.  |

### getBalance

```solidity
function getBalance(address token) public view returns (uint256)
```

_Returns the balance of the specified token for the caller._

#### Parameters

| Name  | Type    | Description                             |
| ----- | ------- | --------------------------------------- |
| token | address | The address of the token being queried. |

#### Return Values

| Name | Type    | Description                                  |
| ---- | ------- | -------------------------------------------- |
| [0]  | uint256 | The balance of the token held by the caller. |

### \_registerReferrer

```solidity
function _registerReferrer(address child, address parent) internal
```

_Registers a referrer for a child address._

#### Parameters

| Name   | Type    | Description                                     |
| ------ | ------- | ----------------------------------------------- |
| child  | address | The child address to register the referrer for. |
| parent | address | The address of the parent referrer.             |

### \_requireNotReferrerOf

```solidity
function _requireNotReferrerOf(address target, address base) internal view
```

_Checks if the target address is not a referrer of the base address._

#### Parameters

| Name   | Type    | Description                        |
| ------ | ------- | ---------------------------------- |
| target | address | The target address to check.       |
| base   | address | The base address to check against. |

### useVouchers

```solidity
function useVouchers(struct IVoucher.Voucher[] vouchers, bytes signature) external
```

_Uses multiple vouchers at once._

#### Parameters

| Name      | Type                      | Description                                |
| --------- | ------------------------- | ------------------------------------------ |
| vouchers  | struct IVoucher.Voucher[] | An array of Voucher structs to be used.    |
| signature | bytes                     | The signature used to verify the vouchers. |

### useVoucher

```solidity
function useVoucher(struct IVoucher.Voucher voucher, bytes signature) external
```

_Uses a single voucher._

#### Parameters

| Name      | Type                    | Description                               |
| --------- | ----------------------- | ----------------------------------------- |
| voucher   | struct IVoucher.Voucher | The Voucher struct to be used.            |
| signature | bytes                   | The signature used to verify the voucher. |

### \_useVoucher

```solidity
function _useVoucher(struct IVoucher.Voucher voucher) internal
```

### \_requireValidVoucher

```solidity
function _requireValidVoucher(struct IVoucher.Voucher voucher) internal view
```

### \_performPayout

```solidity
function _performPayout(address beneficiary, address tokenAddress, uint256 amount, uint8[5] referrersPayouts) internal
```

### \_requireSufficientContractBalance

```solidity
function _requireSufficientContractBalance(contract IERC20Upgradeable token, uint256 expected) internal view
```

### \_requireCorrectSigner

```solidity
function _requireCorrectSigner(bytes encodedData, bytes signature, address signer) internal pure
```

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

\_Function that should revert when `msg.sender` is not authorized to upgrade the contract. Called by
{upgradeTo} and {upgradeToAndCall}.

Normally, this function will use an xref:access.adoc[access control] modifier such as {Ownable-onlyOwner}.

````solidity
function _authorizeUpgrade(address) internal override onlyOwner {}
```_

````
