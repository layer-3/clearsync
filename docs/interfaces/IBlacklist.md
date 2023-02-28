# Solidity API

## IBlacklist

Interface describing functionality of blocking accounts from transferring tokens.

### blacklist

```solidity
function blacklist(address account) external
```

Mark `account` as 'blacklisted' and disallow `transfer` and `transferFrom` operations.

_Require `COMPLIANCE_ROLE` to invoke. Emit `Blacklisted` event`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account to mark 'blacklisted'. |

### removeBlacklisted

```solidity
function removeBlacklisted(address account) external
```

Remove mark 'blacklisted' from `account`, reinstating abilities to invoke `transfer` and `transferFrom`.

_Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedRemoved` event`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account to remove 'blacklisted' mark from. |

### burnBlacklisted

```solidity
function burnBlacklisted(address account) external
```

Burn all tokens from blacklisted `account` specified.

_Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedBurnt` event`. Account specified must be blacklisted._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of 'blacklisted' account to burn funds from. |

### Blacklisted

```solidity
event Blacklisted(address account)
```

`Account` was marked 'blacklisted'.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account to have been marked 'blacklisted'. |

### BlacklistedRemoved

```solidity
event BlacklistedRemoved(address account)
```

Mark 'blacklisted' from `account` was removed.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account 'blacklisted' mark was removed from. |

### BlacklistedBurnt

```solidity
event BlacklistedBurnt(address account, uint256 amount)
```

All tokens from blacklisted `account` specified were burnt.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of 'blacklisted' account funds were burnt from. |
| amount | uint256 |  |

