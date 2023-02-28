# Solidity API

## Token

Yellow and Canary utility token inheriting AccessControl and implementing Cap and Blacklist.

_Blacklist feature is using OpenZeppelin AccessControl._

### MINTER_ROLE

```solidity
bytes32 MINTER_ROLE
```

_Role given to the DAO snapshot_

### COMPLIANCE_ROLE

```solidity
bytes32 COMPLIANCE_ROLE
```

_Role for managing the blacklist process chosen by the DAO_

### BLACKLISTED_ROLE

```solidity
bytes32 BLACKLISTED_ROLE
```

_Role for user blacklisted_

### TOKEN_SUPPLY_CAP

```solidity
uint256 TOKEN_SUPPLY_CAP
```

_Token maximum supply_

### activatedAt

```solidity
uint256 activatedAt
```

_Activation must be called at Token Listing Event._

### Activated

```solidity
event Activated(uint256 premint)
```

Activated event. Emitted when `activate` function is invoked.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| premint | uint256 | Amount of tokes pre-minted during activation. |

### constructor

```solidity
constructor(string name, string symbol, uint256 supplyCap) public
```

_Simple constructor, passing arguments to ERC20 constructor.
Grants `DEFAULT_ADMIN_ROLE`, `COMPLIANCE_ROLE` and `MINTER_ROLE` to deployer._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| name | string | Name of the Token. |
| symbol | string | Symbol of the Token. |
| supplyCap | uint256 |  |

### decimals

```solidity
function decimals() public view virtual returns (uint8)
```

Return the number of decimals used to get its user representation.

_Overrides ERC20 default value of 18;_

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | uint8 Number of decimals of Token. |

### cap

```solidity
function cap() external view returns (uint256)
```

Return the cap on the token's total supply.

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | uint256 Token supply cap. |

### activate

```solidity
function activate(uint256 premint, address account) external
```

Activate token, minting `premint` amount to `account` address.

_Require `DEFAULT_ADMIN_ROLE` to invoke. Premint must satisfy these conditions: 0 < premint < token supply cap. Can be called only once._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| premint | uint256 | Amount of tokens to premint. |
| account | address | Address of account to premint to. |

### mint

```solidity
function mint(address to, uint256 amount) external
```

Increase balance of address `to` by `amount`.

_Require `MINTER_ROLE` to invoke. Emit `Transfer` event.
Require Token to be activated.
The following conditions must be satisfied: `totalSupply + amount <= supplyCap`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to mint tokens to. |
| amount | uint256 | Amount of tokens to mint. |

### burn

```solidity
function burn(uint256 amount) external
```

Destroys `amount` tokens from caller's account. Emit `Transfer` event.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amount | uint256 | Amount of tokens to burn. |

### burnFrom

```solidity
function burnFrom(address account, uint256 amount) external
```

Destroys `amount` tokens from `account`, deducting from the caller's allowance. Emit `Transfer` event.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account to burn tokens from. |
| amount | uint256 | Amount of tokens to burn. |

### transfer

```solidity
function transfer(address to, uint256 amount) public returns (bool)
```

Transfer `amount` of tokens to `to` address from caller.

_Require caller is not marked 'blacklisted'._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to transfer tokens to. |
| amount | uint256 | Amount of tokens to transfer. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if transfer succeeded. |

### transferFrom

```solidity
function transferFrom(address from, address to, uint256 amount) public virtual returns (bool)
```

Transfer `amount` of tokens from `from` to `to` address.

_Require `from` is not marked 'blacklisted'._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| from | address | Address to transfer tokens from. |
| to | address | Address to transfer tokens to. |
| amount | uint256 | Amount of tokens to transfer. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | bool true if transfer succeeded. |

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

_Require `COMPLIANCE_ROLE` to invoke. Emit `BlacklistedBurnt` event`.
Account specified must be blacklisted._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of 'blacklisted' account to burn funds from. |

### _requireAccountNotBlacklisted

```solidity
function _requireAccountNotBlacklisted(address account) internal view
```

Internal Function

_Require `account` is not marked 'blacklisted'._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of account to require not marked 'blacklisted'. |

