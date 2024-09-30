# Solidity Style Guide

This style guide is a set of rules and conventions for writing Solidity code. It is based on the [Solidity Style Guide](https://docs.soliditylang.org/en/latest/style-guide.html), [Coinbase Style Guide](https://github.com/coinbase/solidity-style-guide/blob/7e29d890be81b7df9beefb38f04ff212cc32fb6c/README.md) and includes some additional rules and conventions specific to Yellow Network.

The document contains "syntax" and "semantics" rules for naming, structure, and formatting of Solidity code and tests. If not applied, they will not cause a compilation error or make the code more prone to bugs or vulnerabilities. However, they will make the code consistent, more readable, and easier to maintain.

On the other hand, the [Development Practices](./solidity-development-practices.md) document contains development practises and technics that do change the logic and make the contract more secure, gas-efficient or increase operability.

## 1. Style

### A. Unless an exception or addition is specifically noted, we follow the [Solidity Style Guide](https://docs.soliditylang.org/en/latest/style-guide.html)

### B. Exceptions

#### 1. Names of internal functions in a library should not have an underscore prefix

The style guide states

> Underscore Prefix for Non-external Functions and Variables

One of the motivations for this rule is that it is a helpful visual clue.

> Leading underscores allow you to immediately recognize the intent of such functions...

We agree that a leading underscore is a useful visual clue, and this is why we oppose using them for internal library functions that can be called from other contracts. Visually, it looks wrong.

```solidity
Library._function()
```

or

```solidity
using Library for bytes
bytes._function()
```

Note, we cannot remedy this by insisting on the use of public functions. Whether a library functions are internal or external has important implications. From the [Solidity documentation](https://docs.soliditylang.org/en/latest/contracts.html#libraries)

> ... the code of internal library functions that are called from a contract and all functions called from therein will at compile time be included in the calling contract, and a regular JUMP call will be used instead of a DELEGATECALL.

Developers may prefer internal functions because they are more gas efficient to call.

If a function should never be called from another contract, it should be marked private and its name should have a leading underscore.

### C. Additions

#### 1. Errors

##### A. Prefer custom errors

Custom errors are in some cases more gas efficient and allow passing useful information.

##### B. Custom error names should be CapWords style

YES:

```solidity
error InsufficientBalance();
```

NO:

```solidity
error insufficient_balance();
```

##### A. Do not use `override` keyword when implementing an interface function

Functions defined in an interface are inherently abstract and must be implemented by the inheriting contract, which means there is nothing to override in the traditional sense.

#### 2. Events

##### A. Events names should be past tense

Events should track things that _happened_ and so should be past tense. Using past tense also helps prevent naming collisions with structs or functions.

We are aware this does not follow precedent from early ERCs, like [ERC-20](https://eips.ethereum.org/EIPS/eip-20). However it does align with some more recent high profile Solidity, e.g. [1](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/976a3d53624849ecaef1231019d2052a16a39ce4/contracts/access/Ownable.sol#L33), [2](https://github.com/aave/aave-v3-core/blob/724a9ef43adf139437ba87dcbab63462394d4601/contracts/interfaces/IAaveOracle.sol#L25-L31), [3](https://github.com/ProjectOpenSea/seaport/blob/1d12e33b71b6988cbbe955373ddbc40a87bd5b16/contracts/zones/interfaces/PausableZoneEventsAndErrors.sol#L25-L41).

YES:

```solidity
event OwnerUpdated(address newOwner);
```

NO:

```solidity
event OwnerUpdate(address newOwner);
```

##### B. Prefer `SubjectVerb` naming format

YES:

```solidity
event OwnerUpdated(address newOwner);
```

NO:

```solidity
event UpdatedOwner(address newOwner);
```

#### 3. Functions

##### A. `public` functions not used internally must be marked `external`

YES:

```solidity
/// @dev Not used anywhere else
function withdraw(address to, uint256 amount) external {
  ...
}
```

NO:

```solidity
/// @dev Not used anywhere else
function withdraw(address to, uint256 amount) public {
  ...
}
```

#### 4. Named arguments and parameters

##### A. Avoid unnecessary named return arguments

In short functions, named return arguments are unnecessary.

NO:

```solidity
function add(uint a, uint b) public returns (uint result) {
  result = a + b;
}
```

Named return arguments can be helpful in functions with multiple returned values.

```solidity
function validate(UserOperation calldata userOp) external returns (bytes memory context, uint256 validationData)
```

However, it is important to be explicit when returning early.

YES:

```solidity
function validate(
  UserOperation calldata userOp
) external returns (bytes memory context, uint256 validationData) {
  context = '';
  validationData = 1;

  if (condition) {
    return (context, validationData);
  }
}
```

NO:

```solidity
function validate(
  UserOperation calldata userOp
) external returns (bytes memory context, uint256 validationData) {
  context = '';
  validationData = 1;
  if (condition) {
    return;
  }
}
```

##### B. Prefer named arguments

Passing arguments to functions, events, and errors with explicit naming is helpful for clarity, when the name of the variable passed does not provide sufficient context.

YES:

```solidity
pow({base: x, exponent: y, scalar: v});
pow(base, exp, scal);
```

NO:

```solidity
pow(x, y, v);
```

##### C. Prefer named parameters in mapping types

Explicitly naming parameters in mapping types is helpful for clarity, especially when nesting is used.

YES:

```solidity
mapping(address account => mapping(address asset => uint256 amount)) public balances;
```

NO:

```solidity
mapping(uint256 => mapping(address => uint256)) public balances;
```

##### D. Prefer contract type for function arguments

When passing a contract as an argument, use the contract type instead of `address`, as the former provides type safety, allows direct interaction with the contract, and makes the code more readable and maintainable.

#### 5. Structure of a Contract

##### A. Prefer composition over inheritance

If a function or set of functions could reasonably be defined as its own contract or as a part of a larger contract, prefer defining it as part of a larger contract. This makes the code easier to understand and audit.

Note this _does not_ mean that we should avoid inheritance, in general. Inheritance is useful at times, most especially when building on existing, trusted contracts. For example, _do not_ reimplement `Ownable` functionality to avoid inheritance. Inherit `Ownable` from a trusted vendor, such as [OpenZeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts/) or [Solady](https://github.com/Vectorized/solady).

##### B. Avoid using assembly

Assembly code is hard to read and audit. We should avoid it unless the gas savings are very consequential, e.g. > 25%.

#### 6. Versioning

##### A. Avoid unnecessary version Pragma constraints

While the main contracts we deploy should specify a single Solidity version, all supporting contracts and libraries should have as open a Pragma as possible. A rule of thumb is to specify compatibility with the next major version. For example:

```solidity
pragma solidity ^0.8.0;
```

#### 7. Struct and Error Definitions

##### A. Prefer declaring structs and errors within the interface, contract, or library where they are used

##### B. If a struct or error is used across many files, with no interface, contract, or library reasonably being the "owner," then define them in their own file. Multiple structs and errors can be defined together in one file

#### 8. Imports

##### A. Use named imports

Named imports help readers understand what exactly is being used and where it is originally declared.

YES:

```solidity
import {Contract} from "./contract.sol"
```

NO:

```solidity
import "./contract.sol"
```

For convenience, named imports do not have to be used in test files.

##### B. Order imports alphabetically (A to Z) by file name

YES:

```solidity
import {A} from './A.sol'
import {B} from './B.sol'
```

NO:

```solidity
import {B} from './B.sol'
import {A} from './A.sol'
```

##### C. Group imports by external and local with a new line in between

YES:

```solidity
import { Math } from '/solady/Math.sol';

import { MyHelper } from './MyHelper.sol';
```

In test files, imports from `/test` should be their own group, as well.

YES:

```solidity
import { Math } from '/solady/Math.sol';

import { MyHelper } from '../src/MyHelper.sol';

import { Mock } from './mocks/Mock.sol';
```

NO:

```solidity
import { MyHelper } from './MyHelper.sol';

import { Math } from '/solady/Math.sol';

// or

import { Math } from '/solady/Math.sol';
import { MyHelper } from './MyHelper.sol';
```

#### 9. Comments

##### A. Commenting to group sections of the code is permitted

Sometimes authors and readers find it helpful to comment dividers between groups of functions. This permitted, however ensure the style guide [ordering of functions](https://docs.soliditylang.org/en/latest/style-guide.html#order-of-functions) is still followed.

YES:

```solidity
/// External Functions ///
```

```solidity
/*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
/*                   VALIDATION OPERATIONS                    */
/*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/
```

## 2. Development

### A. Use [Forge](https://github.com/foundry-rs/foundry/tree/master/crates/forge) for testing and dependency management if possible

### B. Forge testing

This section applies if Forge is used for testing.

#### 1. Test files and contracts

##### A. Test file names should follow Solidity Style Guide conventions for files names and also have `.t` before `.sol`

YES:

- AuthorizerTest.t.sol
- ERC20Test.t.sol

##### B. Harness contract names should start from "Test", followed by the contract under test

YES:

- `TestERC20`
- `TestDailyClaim`

##### C. Integration tests should be in a separate file with a name that follows the convention `<ContractName>Test_integration_<functionName>.t.sol`

Integration test files can also contain tests for multiple functions, in which case the `<functionName>` part can be omitted.

YES:

- `ERC20Test_integration_transfer.t.sol`
- `DailyClaimTest_integration.t.sol`

##### D. Test contract names should include the name of the contract being tested, followed by "Test", with an optional suffix for the function being tested

YEs:

- `ERC20Test`
- `ERC20Test_transfer`
- `DailyClaimTest`

Note, that if a mocked contract is used for testing, the test contract should be named after the contract being tested, not the mock.

NO:

- `TestERC20Test`
- `TestDailyClaimTest`

##### E. Prefer separate test contracts for complex functions

Unit tests for complex functions should be in their own test contract. This makes it easier to understand what is being tested and to debug failures.

Simple functions, on the other hand, can be tested in the same root test contract.

For example, the contract "ERC20" has a complex function `transferFrom`. This function should have its own test contract, `ERC20Test_transferFrom`.
In the same time, the function `decimals` can be tested in the root test contract, `ERC20Test`.

##### F. Test contracts/functions should be written in the same order as the original functions in the contract-under-test

#### 2. Test functions

##### A. In harness contract, each internal function that is tested should be exposed via an external one with a name that follows the pattern `exposed\_<function_name>`

YES:

```solidity
// file: src/MyContract.sol
contract MyContract {
  function myInternalMethod() internal returns (uint) {
    return 42;
  }

  function _mySecondInternalMethod() internal returns (uint) {
    return 442;
  }
}

// file: test/MyContract.t.sol
import { MyContract } from 'src/MyContract.sol';

contract MyContractHarness is MyContract {
  function exposed_myInternalMethod() external returns (uint) {
    return myInternalMethod();
  }

  function exposed_mySecondInternalMethod() external returns (uint) {
    return _mySecondInternalMethod();
  }
}
```

NO:

```solidity
// file: src/MyContract.sol
contract MyContract {
  function myInternalMethod() internal returns (uint) {
    return 42;
  }

  function _mySecondInternalMethod() internal returns (uint) {
    return 442;
  }
}

// file: test/MyContract.t.sol
import { MyContract } from 'src/MyContract.sol';

contract MyContractHarness is MyContract {
  function getMyInternalMethod() external returns (uint) {
    return myInternalMethod();
  }

  function mySecondInternalMethod() external returns (uint) {
    return _mySecondInternalMethod();
  }
}
```

##### B. In harness contracts, workaround functions should be named `workaround_<function_name>`

Harnesses can also be used to expose functionality or information otherwise unavailable in the original smart contract.
The most straightforward example is when we want to test the length of a public array.

YES:

- `workaround_queueLength()`
- `workaround_isElementInMapping()`

NO:

- `getQueueLength()`
- `setAuthorizerWithoutOwner()`

##### C. Test names should follow the convention `test_functionName_outcome_optionalContext`

YES:

- `test_transferFrom_debitsFromAccountBalance`
- `test_transferFrom_debitsFromAccountBalance_whenCalledViaPermit`
- `test_transferFrom_revert_ifAmountExceedsBalance`

NO:

- `testTransfer`
- `test_transferFrom`
- `test_success_transfer`
- `test_transferFromWhenNoBalance`

If the contract is named after a function, then function name can be omitted.

YES:

```solidity
contract TransferFromTest {
  function test_debitsFromAccountBalance() ...
}
```

NO:

```solidity
contract TransferFromTest {
  function test_transferFrom_debitsFromAccountBalance() ...
}
```

##### D. Reverting tests should follow the convention `test_functionName_revert_[if|when]Condition`

YES:

- `test_transferFrom_revert_ifAmountExceedsBalance`
- `test_transferFrom_revert_ifNotOwner`

NO:

- `test_transferFrom_revert`
- `test_transferFrom_revertWhenAmountExceedsBalance`
- `test_transferFromRevertIfNotOwner`

##### E. Prefer tests that test one thing

This is generally good practice, but especially so because Forge does not give line numbers on assertion failures. This makes it hard to track down what, exactly, failed if a test has many assertions.

YES:

```solidity
function test_transferFrom_debitsFrom() {
  ...
}

function test_transferFrom_creditsTo() {
  ...
}

function test_transferFrom_emitsCorrectly() {
  ...
}

function test_transferFrom_reverts_whenAmountExceedsBalance() {
  ...
}
```

NO:

```solidity
function test_transferFrom_works() {
  // debits correctly
  // credits correctly
  // emits correctly
  // reverts correctly
}
```

Note, this does not mean a test should only ever have one assertion. Sometimes having multiple assertions is helpful for certainty on what is being tested.

YES:

```solidity
function test_transferFrom_creditsTo() {
  assertEq(balanceOf(to), 0);
  ...
  assertEq(balanceOf(to), amount);
}
```

##### F. Use variables for important values in tests

YES:

```solidity
function test_transferFrom_creditsTo() {
  assertEq(balanceOf(to), 0);
  uint amount = 10;
  transferFrom(from, to, amount);
  assertEq(balanceOf(to), amount);
}
```

NO:

```solidity
function test_transferFrom_creditsTo() {
  assertEq(balanceOf(to), 0);
  transferFrom(from, to, 10);
  assertEq(balanceOf(to), 10);
}
```

##### G. When testing events, prefer using `vm.expectEmit()`

YES:

```solidity
function test_transferFrom_emitsCorrectly() {
  vm.expectEmit();
  emit IERC20.Transfer(from, to, 10);
  token.transferFrom(from, to, 10);
}
```

NO:

```solidity
function test_transferFrom_emitsCorrectly() {
  vm.expectEmit(true, true, true, true);
  token.transferFrom(from, to, 10);
  vm.expectEmit();
}
```

Benefits:

- `vm.expectEmit()` is equal to `vm.expectEmit(true,true,true,true)`, but takes less space.

Benefits:

- This ensures you test everything in your event.
- If you add a topic (i.e. a new indexed parameter), it’s now tested by default.
- Even if you only have 1 topic, the extra true arguments don’t hurt.

##### H. Prefer fuzz tests

All else being equal, prefer fuzz tests.

YES:

```solidity
function test_transferFrom_creditsTo(uint amount) {
  assertEq(balanceOf(to), 0);
  transferFrom(from, to, amount);
  assertEq(balanceOf(to), amount);
}
```

NO:

```solidity
function test_transferFrom_creditsTo() {
  assertEq(balanceOf(to), 0);
  uint amount = 10;
  transferFrom(from, to, amount);
  assertEq(balanceOf(to), amount);
}
```

### C. Project Setup

#### 1. Avoid custom remappings

[Remappings](https://book.getfoundry.sh/projects/dependencies?#remapping-dependencies) help Forge find dependencies based on import statements. Forge will automatically deduce some remappings, for example

```rust
forge-std/=lib/forge-std/src/
solmate/=lib/solmate/src/
```

We should avoid adding to these or defining any remappings explicitly, as it makes our project harder for others to use as a dependency. For example, if our project depends on Solmate and so does theirs, we want to avoid our project having some irregular import naming, resolved with a custom remapping, which will conflict with their import naming.

### D. Upgradability

#### 1. Prefer [ERC-7201](https://eips.ethereum.org/EIPS/eip-7201) "Namespaced Storage Layout" convention to prevent storage collisions

## 3. NatSpec

### A. Unless an exception or addition is specifically noted, follow [Solidity NatSpec](https://docs.soliditylang.org/en/latest/natspec-format.html)

### B. Additions

#### 1. All external functions, events, and errors should have complete NatSpec

Minimally including `@notice`. `@param` and `@return` should be present if there are parameters or return values.

#### 2. Struct NatSpec

Structs can be documented with a `@notice` above and, if desired, `@dev` for each field.

YES:

```solidity
/// @notice A struct describing an accounts position
struct Position {
  /// @dev The unix timestamp (seconds) of the block when the position was created.
  uint created;
  /// @dev The amount of ETH in the position
  uint amount;
}
```

#### 3. Newlines between tag types

For easier reading, add a new line between tag types, when multiple are present and there are three or more lines.

YES:

```solidity
/// @notice ...
///
/// @dev ...
/// @dev ...
///
/// @param ...
/// @param ...
///
/// @return
```

NO:

```solidity
/// @notice ...
/// @dev ...
/// @dev ...
/// @param ...
/// @param ...
/// @return
```
