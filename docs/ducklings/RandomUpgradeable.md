# Solidity API

## RandomUpgradeable

### salt

```solidity
bytes32 salt
```

### nonce

```solidity
uint256 nonce
```

### __Random_init

```solidity
function __Random_init() internal
```

### __Random_init_unchained

```solidity
function __Random_init_unchained() internal
```

### Random

```solidity
modifier Random()
```

### UseRandom

```solidity
modifier UseRandom()
```

### _updateNonce

```solidity
function _updateNonce() private
```

### _updateSalt

```solidity
function _updateSalt() private
```

### _randomMaxNumber

```solidity
function _randomMaxNumber(uint256 max) internal returns (uint256)
```

