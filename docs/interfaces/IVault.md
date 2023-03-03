# Solidity API

## IVault

_Interface for a smart contract vault that can hold multiple currencies, allowing users to deposit and withdraw funds in any supported token._

### deposit

```solidity
function deposit(address token, uint256 amount) external payable
```

_Deposits the specified token into the vault._

#### Parameters

| Name   | Type    | Description                               |
| ------ | ------- | ----------------------------------------- |
| token  | address | The address of the token being deposited. |
| amount | uint256 | The amount of the token being deposited.  |

### withdraw

```solidity
function withdraw(address token, uint256 amount) external
```

_Withdraws the specified token from the vault._

#### Parameters

| Name   | Type    | Description                               |
| ------ | ------- | ----------------------------------------- |
| token  | address | The address of the token being withdrawn. |
| amount | uint256 | The amount of the token to be withdrawn.  |

### getBalance

```solidity
function getBalance(address token) external view returns (uint256)
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
