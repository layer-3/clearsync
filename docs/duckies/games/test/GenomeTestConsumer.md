# Solidity API

## GenomeTestConsumer

### getFlags

```solidity
function getFlags(uint256 genome) external pure returns (uint8)
```

### getFlag

```solidity
function getFlag(uint256 genome, uint8 flag) external pure returns (bool)
```

### setFlag

```solidity
function setFlag(uint256 genome, uint8 flag, bool value) external pure returns (uint256)
```

### setGene

```solidity
function setGene(uint256 genome, uint8 gene, uint8 value) external pure returns (uint256)
```

### getGene

```solidity
function getGene(uint256 genome, uint8 gene) external pure returns (uint8)
```

### maxGene

```solidity
function maxGene(uint256[] genomes, uint8 gene) external pure returns (uint8)
```

### geneValuesAreEqual

```solidity
function geneValuesAreEqual(uint256[] genomes, uint8 gene) external pure returns (bool)
```

### geneValuesAreUnique

```solidity
function geneValuesAreUnique(uint256[] genomes, uint8 gene) external pure returns (bool)
```

