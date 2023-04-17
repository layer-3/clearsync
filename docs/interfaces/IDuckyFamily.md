# Solidity API

## IDuckyFamily

### InvalidMintParams

```solidity
error InvalidMintParams(struct IDuckyFamily.MintParams mintParams)
```

### InvalidMeldParams

```solidity
error InvalidMeldParams(struct IDuckyFamily.MeldParams meldParams)
```

### MintingRulesViolated

```solidity
error MintingRulesViolated(uint8 collectionId, uint8 amount)
```

### MeldingRulesViolated

```solidity
error MeldingRulesViolated(uint256[] tokenIds)
```

### IncorrectGenomesForMelding

```solidity
error IncorrectGenomesForMelding(uint256[] genomes)
```

### Melded

```solidity
event Melded(address owner, uint256[] meldingTokenIds, uint256 meldedTokenId, uint256 chainId)
```

### VoucherActions

```solidity
enum VoucherActions {
  MintPack,
  MeldFlock
}
```

### MintParams

```solidity
struct MintParams {
  address to;
  uint8 size;
  bool isTransferable;
}
```

### MeldParams

```solidity
struct MeldParams {
  address owner;
  uint256[] tokenIds;
  bool isTransferable;
}
```

### Rarities

```solidity
enum Rarities {
  Common,
  Rare,
  Epic,
  Legendary
}
```

### GeneDistributionTypes

```solidity
enum GeneDistributionTypes {
  Even,
  Uneven
}
```

### GenerativeGenes

```solidity
enum GenerativeGenes {
  Collection,
  Rarity,
  Color,
  Family,
  Body,
  Head
}
```

### MythicGenes

```solidity
enum MythicGenes {
  Collection,
  UniqId
}
```

### getMintPrice

```solidity
function getMintPrice() external view returns (uint256)
```

### getMeldPrices

```solidity
function getMeldPrices() external view returns (uint256[4])
```

### getCollectionsGeneValues

```solidity
function getCollectionsGeneValues() external view returns (uint8[][3], uint8)
```

### getCollectionsGeneDistributionTypes

```solidity
function getCollectionsGeneDistributionTypes() external view returns (uint32[3])
```

### mintPack

```solidity
function mintPack(uint8 size) external
```

### meldFlock

```solidity
function meldFlock(uint256[] meldingTokenIds) external
```

