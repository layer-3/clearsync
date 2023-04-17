# Solidity API

## Genome

The library to work with NFT genomes.

Genome is a number with a special structure that defines Duckling genes.
All genes are packed consequently in the reversed order in the Genome, meaning the first gene being stored in the last Genome bits.
Each gene takes up the block of 8 bits in genome, thus having 256 possible values.

Example of genome, following genes Rarity, Head and Body are defined:

00000001|00000010|00000011
  Body    Head     Rarity

This genome can be represented in uint24 as 66051.
Genes have the following values: Body = 1, Head = 2, Rarity = 3.

### BITS_PER_GENE

```solidity
uint8 BITS_PER_GENE
```

Number of bits each gene constitutes. Thus, each gene can have 2^8 = 256 possible values.

### COLLECTION_GENE_IDX

```solidity
uint8 COLLECTION_GENE_IDX
```

### FLAGS_GENE_IDX

```solidity
uint8 FLAGS_GENE_IDX
```

Reserve 30th gene for bool flags, which are stored as a bit field.

### FLAG_TRANSFERABLE

```solidity
uint8 FLAG_TRANSFERABLE
```

### MAGIC_NUMBER_GENE_IDX

```solidity
uint8 MAGIC_NUMBER_GENE_IDX
```

Reserve 31th gene for magic number, which is used as an extension for genomes.
Genomes with wrong extension are considered invalid.

### BASE_MAGIC_NUMBER

```solidity
uint8 BASE_MAGIC_NUMBER
```

### MYTHIC_MAGIC_NUMBER

```solidity
uint8 MYTHIC_MAGIC_NUMBER
```

### getFlags

```solidity
function getFlags(uint256 self) internal pure returns (uint8)
```

Read flags gene from genome.

_Read flags gene from genome._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | uint256 | Genome to get flags gene from. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | flags Flags gene. |

### getFlag

```solidity
function getFlag(uint256 self, uint8 flag) internal pure returns (bool)
```

Read specific bit mask flag from genome.

_Read specific bit mask flag from genome._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | uint256 | Genome to read flag from. |
| flag | uint8 | Bit mask flag to read. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | value Value of the flag. |

### setFlag

```solidity
function setFlag(uint256 self, uint8 flag, bool value) internal pure returns (uint256)
```

Set specific bit mask flag in genome.

_Set specific bit mask flag in genome._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | uint256 | Genome to set flag in. |
| flag | uint8 | Bit mask flag to set. |
| value | bool | Value of the flag. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Genome with the flag set. |

### setGene

```solidity
function setGene(uint256 self, uint8 gene, uint8 value) internal pure returns (uint256)
```

Set `value` to `gene` in genome.

_Set `value` to `gene` in genome._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | uint256 | Genome to set gene in. |
| gene | uint8 | Gene to set. |
| value | uint8 | Value to set. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Genome with the gene set. |

### getGene

```solidity
function getGene(uint256 self, uint8 gene) internal pure returns (uint8)
```

Get `gene` value from genome.

_Get `gene` value from genome._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| self | uint256 | Genome to get gene from. |
| gene | uint8 | Gene to get. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | geneValue Gene value. |

### _maxGene

```solidity
function _maxGene(uint256[] genomes, uint8 gene) internal pure returns (uint8)
```

Get largest value of a `gene` in `genomes`.

_Get largest value of a `gene` in `genomes`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Genomes to get gene from. |
| gene | uint8 | Gene to get. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | maxValue Largest value of a `gene` in `genomes`. |

### _geneValuesAreEqual

```solidity
function _geneValuesAreEqual(uint256[] genomes, uint8 gene) internal pure returns (bool)
```

Check if values of `gene` in `genomes` are equal.

_Check if values of `gene` in `genomes` are equal._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Genomes to check. |
| gene | uint8 | Gene to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isEqual True if values of `gene` in `genomes` are equal, false otherwise. |

### _geneValuesAreUnique

```solidity
function _geneValuesAreUnique(uint256[] genomes, uint8 gene) internal pure returns (bool)
```

Check if values of `gene` in `genomes` are unique.

_Check if values of `gene` in `genomes` are unique._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Genomes to check. |
| gene | uint8 | Gene to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isUnique True if values of `gene` in `genomes` are unique, false otherwise. |

