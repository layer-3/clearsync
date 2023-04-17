# Solidity API

## TreasureVault

This contract allows users to deposit tokens into a vault and redeem vouchers for rewards.

The vouchers can then be used to redeem rewards or to refer others to the platform. Referral commissions are paid out
to referrers of up to 5 levels deep. This contract also allows the issuer to set an authorized address for signing
vouchers and upgrading the contract.

### CircularReferrers

```solidity
error CircularReferrers(address target, address base)
```

### InvalidRewardParams

```solidity
error InvalidRewardParams(struct TreasureVault.RewardParams rewardParams)
```

### InsufficientTokenBalance

```solidity
error InsufficientTokenBalance(address token, uint256 expected, uint256 actual)
```

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

### TREASURY_ROLE

```solidity
bytes32 TREASURY_ROLE
```

### REFERRAL_MAX_DEPTH

```solidity
uint8 REFERRAL_MAX_DEPTH
```

### _REFERRAL_PAYOUT_DIVIDER

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

### _referrerOf

```solidity
mapping(address => address) _referrerOf
```

### _usedVouchers

```solidity
mapping(bytes32 => bool) _usedVouchers
```

### issuer

```solidity
address issuer
```

### AffiliateRegistered

```solidity
event AffiliateRegistered(address affiliate, address referrer)
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

Set the address of vouchers issuer.

_Require `DEFAULT_ADMIN_ROLE` to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | The address of the new issuer. |

### withdraw

```solidity
function withdraw(address tokenAddress, address beneficiary, uint256 amount) public
```

Withdraw the specified token from the vault. The risk of single account withdrawing all balances of this contract
is mitigated by requiring `TREASURY_ROLE` to invoke, which is granted to a Quorum account.

_Require `TREASURY_ROLE` to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenAddress | address | The address of the token being withdrawn. |
| beneficiary | address | The address of the account receiving the amount. |
| amount | uint256 | The amount of the token to be withdrawn. |

### _registerReferrer

```solidity
function _registerReferrer(address child, address parent) internal
```

Register a referrer for a child address. Internal function.

_Emit `AffiliateRegistered` event._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| child | address | The child address to register the referrer for. |
| parent | address | The address of the parent referrer. |

### _requireNotReferrerOf

```solidity
function _requireNotReferrerOf(address target, address base) internal view
```

Check if the target address is not a referrer of the base address. Internal function.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| target | address | The target address to check. |
| base | address | The base address to check against. |

### useVouchers

```solidity
function useVouchers(struct IVoucher.Voucher[] vouchers, bytes signature) external
```

Use multiple vouchers at once.

_Emit `VoucherUsed` event for each voucher used._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| vouchers | struct IVoucher.Voucher[] | An array of Voucher structs to be used. |
| signature | bytes | Array of Vouchers signed by the Issuer. |

### useVoucher

```solidity
function useVoucher(struct IVoucher.Voucher voucher, bytes signature) external
```

Use a single voucher.

_Emit `VoucherUsed` event._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | The Voucher struct to be used. |
| signature | bytes | Voucher signed by the Issuer. |

### _useVoucher

```solidity
function _useVoucher(struct IVoucher.Voucher voucher) internal
```

Use a single voucher. Internal function.

_Emit `VoucherUsed` event._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | Voucher to be used. |

### _requireValidVoucher

```solidity
function _requireValidVoucher(struct IVoucher.Voucher voucher) internal view
```

Check voucher for being valid. Internal function.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | Voucher to check for validity. |

### _performPayout

```solidity
function _performPayout(address beneficiary, address tokenAddress, uint256 amount, uint8[5] referrersPayouts) internal
```

Perform reward payout, including commissions. Internal function.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| beneficiary | address | The address receiving the payout. |
| tokenAddress | address | The token to be paid. |
| amount | uint256 | Amount to be paid. |
| referrersPayouts | uint8[5] | Commissions to be paid to the referrers of the beneficiary, if any. |

### _requireSufficientContractBalance

```solidity
function _requireSufficientContractBalance(contract IERC20Upgradeable token, uint256 expected) internal view
```

Require this contract has not less than `expected` amount of the `token` deposited. Internal function.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| token | contract IERC20Upgradeable | Token to be deposited to the address of this contract. |
| expected | uint256 | Minimal amount of the `token` to be on this contract. |

### _requireCorrectSigner

```solidity
function _requireCorrectSigner(bytes encodedData, bytes signature, address signer) internal pure
```

Require `encodedData` was signed by the `signer`. Internal function.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| encodedData | bytes | Encoded data signed. |
| signature | bytes | Signature produced by the `signer` signing `encodedData`. |
| signer | address | Signer to have signed `encodedData`. |

### _authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Restrict upgrading this contract to address with `UPGRADER_ROLE`.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newImplementation | address | Address of the new implementation. |

