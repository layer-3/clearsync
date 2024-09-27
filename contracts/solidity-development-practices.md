# Solidity Development Practices

Solidity Development Practices is a comprehensive guide dedicated to enhancing the quality of smart contracts by focusing on development techniques that impact the logic and functionality of the code.
Unlike a [Style Guide](./solidity-style-guide.md) that emphasizes syntax, naming conventions, and formatting for readability and consistency, this document delves into practices that make smart contracts more secure, gas-efficient, and operationally robust.

This document contains:

- Security Best Practices
- Gas Optimization and Efficiency technics
- Operational Advice

## 1. Security Best Practices

### A. Checks, Effects, Interactions Pattern

A function should first perform all necessary checks or validations, then execute the intended effects by updating the contract's state, and only then proceed with any interactions involving external contracts or addresses.
Adhering to this sequence minimizes the risk of external calls manipulating the contract's state in unintended ways, preventing reentrancy attacks.

YES:

```solidity
function withdraw(uint256 _amount) public {
  // checks
  require(_amount <= balances[msg.sender], 'Insufficient balance');
  // effects
  balances[msg.sender] -= _amount;
  // interactions
  (bool success, ) = msg.sender.call{ value: _amount }('');
  require(success, 'Transfer failed');
}
```

NO:

```solidity
function withdraw(uint256 _amount) public {
  // interactions
  (bool success, ) = msg.sender.call{ value: _amount }('');
  require(success, 'Transfer failed');
  // effects
  balances[msg.sender] -= _amount;
  // checks
  require(_amount <= balances[msg.sender], 'Insufficient balance');
}
```

### B. Validation

#### 1. Zero address validation

Always include a special handling for zero addresses passed as parameters.

YES:

```solidity
function mint(address to, uint256 amount) public {
  require(to != address(0), 'Invalid address');
  _mint_(to, amount);
}
```

NO:

```solidity
function mint(address to, uint256 amount) public {
  _mint(to, amount);
}
```

### C. Access Control

#### 1. Use Two-step ownership transfer pattern

When transferring ownership of a contract, use a two-step process that requires the new owner to accept the transfer before it is finalized.

This can be done using `@openzeppelin/contracts/access/Ownable2Step.sol` contract.

YES:

```solidity
contract Token is Ownable2Step {
  ...
}
```

NO:

```solidity
contract Token is Ownable {
  ...
}
```

## 2. Gas Optimization and Efficiency

### A. Storage

#### 1. Where possible, struct values should be packed to minimize SLOADs and SSTOREs

YES:

```solidity
struct Position {
  // slot 1:
  uint256 amount;
  // slot 2:
  address beneficiary;
  uint64 id;
  uint32 validUntil;
  // slot 3*:
  bytes data;
}
```

NO:

```solidity
struct Position {
  // slot 1:
  address beneficiary;
  // slot 2:
  uint256 amount;
  // slot 3:
  uint64 id;
  // slot 4*:
  bytes data;
  // slot 5:
  uint32 validUntil;
}
```

\* - in fact, dynamic data takes more that one slot and is not stored conqecutively with other fields.

#### 2. Timestamp fields in a struct should be at least uint32 and ideally be uint40

`uint32` will give the contract ~82 years of validity `(2^32 / (60*60*24*365)) - (2024 - 1970)`. If space allows, uint40 is the preferred size.

## 3. Operational Advice

### A. Interoperability

#### 1. Use `call` instead of `transfer` to send Ether

It is recommended to use built-in call() function instead of transfer() to transfer native assets. This method does not impose a 2300 Gas limit, it provides greater flexibility and compatibility with contracts having more complex business logic upon receiving the native tokens.
When working with the call() function ensure that its execution is successful by checking the returned boolean value.

YES:

```solidity
(bool success, ) = msg.sender.call{ value: _amount }('');
require(success, 'Transfer failed');
```

NO:

```solidity
msg.sender.transfer(_amount);
```

#### 2. Avoid using `msg.sender` for permissionless functions

Using `msg.sender` instead of an explicit address parameter can preclude certain communication logic and use cases.
It is recommended to use an explicit address parameter for functions that don't have a permission-based access, e.g. view functions, deposits, etc.

YES:

```solidity
function deposit(address beneficiary, address token, uint256 amount) public {
  require(amount > 0, 'Invalid amount');
  require(beneficiary != address(0), 'Invalid beneficiary');
  // deposit logic to beneficiary
}
```

NO:

```solidity
function deposit(address token, uint256 amount) public {
  require(amount > 0, 'Invalid amount');
  // deposit logic to msg.sender
}
```
