# Solidity API

## Gene

### BYTES_PER_TRAIT

```solidity
uint8 BYTES_PER_TRAIT
```

### Classes

```solidity
enum Classes {
  Common,
  Rare,
  Epic,
  Legendary,
  SuperLegendary,
  Zombie
}
```

### Traits

```solidity
enum Traits {
  Class,
  Body,
  Head,
  Background,
  Element,
  Eyes,
  Beak,
  Wings,
  Firstname,
  Lastname,
  Temper,
  Peculiarity
}
```

### setTrait

```solidity
function setTrait(uint256 gene, enum Gene.Traits trait, uint8 value) internal pure returns (uint256)
```

### getTrait

```solidity
function getTrait(uint256 gene, enum Gene.Traits trait) internal pure returns (uint8)
```

