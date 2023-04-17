# Solidity API

## TESTDuckyFamilyV1

### GenomeReturned

```solidity
event GenomeReturned(uint256 genome)
```

### GeneReturned

```solidity
event GeneReturned(uint8 gene)
```

### BoolReturned

```solidity
event BoolReturned(bool returnedBool)
```

### Uint8Returned

```solidity
event Uint8Returned(uint8 returnedUint8)
```

### constructor

```solidity
constructor(address duckiesAddress, address ducklingsAddress, address treasureVaultAddress) public
```

### setRarityChances

```solidity
function setRarityChances(uint32[] chances) external
```

### setCollectionMutationChances

```solidity
function setCollectionMutationChances(uint32[] chances) external
```

### setGeneMutationChance

```solidity
function setGeneMutationChance(uint32 chance) external
```

### setGeneInheritanceChances

```solidity
function setGeneInheritanceChances(uint32[] chances) external
```

### generateGenome

```solidity
function generateGenome(uint8 collectionId) external
```

### generateAndSetGenes

```solidity
function generateAndSetGenes(uint256 genome, uint8 collectionId) external
```

### generateAndSetGene

```solidity
function generateAndSetGene(uint256 genome, uint8 geneIdx, uint8 geneValuesNum, enum IDuckyFamily.GeneDistributionTypes distrType) external
```

### generateMythicGenome

```solidity
function generateMythicGenome(uint256[] genomes) external
```

### requireGenomesSatisfyMelding

```solidity
function requireGenomesSatisfyMelding(uint256[] genomes) external pure
```

### meldGenomes

```solidity
function meldGenomes(uint256[] genomes) external
```

### isCollectionMutating

```solidity
function isCollectionMutating(enum IDuckyFamily.Rarities rarity) external
```

### meldGenes

```solidity
function meldGenes(uint256[] genomes, uint8 gene, uint8 maxGeneValue, enum IDuckyFamily.GeneDistributionTypes geneDistrType) external
```

### getDistributionType

```solidity
function getDistributionType(uint32 distributionTypes, uint8 idx) external pure returns (enum IDuckyFamily.GeneDistributionTypes)
```

### generateUnevenGeneValue

```solidity
function generateUnevenGeneValue(uint8 valuesNum) external
```

### calcMaxPeculiarity

```solidity
function calcMaxPeculiarity() external view returns (uint16)
```

### calcPeculiarity

```solidity
function calcPeculiarity(uint256 genome) external view returns (uint16)
```

### calcUniqIdGenerationParams

```solidity
function calcUniqIdGenerationParams(uint16 pivotalUniqId, uint16 maxUniqId) external pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength)
```

