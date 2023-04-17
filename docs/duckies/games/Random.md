# Solidity API

## Random

A contract that provides pseudo random number generation.
Pseudo random number generation is based on the block timestamp, sender address, salt and nonce.
Salt is based on block timestamp and msg sender, and is calculated every time a user-function that uses Random logic is called.
Nonce is incremented every time a random number is generated.

### InvalidWeights

```solidity
error InvalidWeights(uint32[] weights)
```

Invalid weights error while trying to generate a weighted random number.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| weights | uint32[] | Empty weights array. |

### UseRandom

```solidity
modifier UseRandom()
```

Specifies that calling function uses random number generation.

_Modifier that updates salt after calling function is invoked._

### _randomMaxNumber

```solidity
function _randomMaxNumber(uint256 max) internal returns (uint256)
```

Generates a random number in range [0, max).

_Cast hash of encoded salt, nonce, msg sender block timestamp to the number, and returns modulo `max`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| max | uint256 | Upper bound of the range. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | Random number in range [0, max). |

### _randomWeightedNumber

```solidity
function _randomWeightedNumber(uint32[] weights) internal returns (uint8)
```

Generates a weighted random number given the `weights` array in range [0, weights.length).

_Number `x` is generated with probability `weights[x] / sum(weights)`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| weights | uint32[] | Array of weights. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | Random number in range [0, weights.length). |

### _sum

```solidity
function _sum(uint32[] numbers) internal pure returns (uint256 sum)
```

Calculates sum of all elements in array.

_Calculates sum of all elements in array._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| numbers | uint32[] | Array of numbers. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| sum | uint256 | Sum of all elements in array. |

