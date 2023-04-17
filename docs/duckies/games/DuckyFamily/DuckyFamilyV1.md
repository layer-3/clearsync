# Solidity API

## DuckyFamilyV1

DuckyFamily contract defines rules of Ducky Family game, which is a card game similar to Happy Families and Uno games.
This game also incorporates vouchers as defined in IVoucher interface.

DuckyFamily operates on Ducklings NFT, which is defined in a corresponding contract. DuckyFamily can mint, burn and query information about NFTs
by calling Ducklings contract.

Users can buy NFT (card) packs of different size. When a pack is bought, a number of cards is generated and assigned to the user.
The packs can be bought with Duckies token, so user should approve DuckyFamily contract to spend Duckies on his behalf.

Each card has a unique genome, which is a 256-bit number. The genome is a combination of different genes, which describe the card and its properties.
There are 3 types of cards introduced in this game, which are differentiated by the 'collection' gene: Duckling, Zombeak and Mythic.
Duckling and Zombeak NFTs have a class system, which is defined by 'rarity' gene: Common, Rare, Epic and Legendary.
Mythic NFTs are not part of the class system and are considered to be the most rare and powerful cards in the game.

All cards have a set of generative genes, which are used to describe the card, its rarity and image.
There are 2 types of generative genes: with even and uneven chance for each value of that gene.
All values of even genes are generated with equal probability, while uneven genes have a higher chance for the first values and lower for the last values.
Thus, while even genes can describe the card, uneven can set the rarity of the card.

Note: don't confuse 'rarity' gene with rarity of the card. 'Rarity' gene is a part of the game logic, while rarity of the card is a value this card represents.
Henceforth, if a 'Common' rarity gene card has uneven generative genes with high values (which means this card has a tiny chance to being generated),
then this card can be more rare than some 'Rare' rarity gene cards.
So, when we mean 'rarity' gene, we will use quotes, while when we mean rarity of the card, we will use it without quotes.

Duckling are the main cards in the game, as they are the only way users can get Mythic cards.
However, users are not obliged to use every Duckling cards to help them get Mythic, they can improve them and collect the rarest ones.
Users can get Duckling cards from minting packs.

Users can improve the 'rarity' of the card by melding them. Melding is a process of combining a flock of 5 cards to create a new one.
The new card will have the same 'collection' gene as the first card in the flock, but the 'rarity' gene will be incremented.
However, users must oblige to specific rules when melding cards:
1. All cards in the flock must have the same 'collection' gene.
2. All cards in the flock must have the same 'rarity' gene.
3a. When melding Common cards, all cards in the flock must have either the same Color or Family gene values.
3b. When melding Rare and Epic cards, all cards in the flock must have both the same Color and Family gene values.
3c. When melding Legendary cards, all cards in the flock must have the same Color and different Family gene values.
4. Mythic cards cannot be melded.
5. Legendary Zombeak cards cannot be melded.

Other generative genes of the melded card are not random, but are calculated from the genes of the source cards.
This process is called 'inheritance' and is the following:
1. Each generative gene is inherited separately
2. A gene has a high chance of being inherited from the first card in the flock, and this chance is lower for each next card in the flock.
3. A gene has a mere chance of 'positive mutation', which sets inherited gene value to be bigger than the biggest value of this gene in the flock.

Melding is not free and has a different cost for each 'rarity' of the cards being melded.

Zombeak are secondary cards, that you can only get when melding mutates. There is a different chance (defined in Config section below) for each 'rarity' of the Duckling cards that are being melded,
that the melding result card will mutate to Zombeak. If the melding mutates, then the new card will have the same 'rarity' gene as the source cards.
This logic makes Zombeak cards more rare than some Duckling cards, as they can only be obtained by melding mutating.
However, Zombeak cards cannot be melded into Mythic, which means their main value is rarity.

Mythic are the most rare and powerful cards in the game. They can only be obtained by melding Legendary Duckling cards with special rules described above.
The rarity of the Mythic card is defined by the 'UniqId' gene, which corresponds to the picture of the card. The higher the 'UniqId' gene value, the rarer the card.
The 'UniqId' value is correlated with the 'peculiarity' of the flock that creates the Mythic: the higher the peculiarity, the higher the 'UniqId' value.
Peculiarity of the card is a sum of all uneven gene values of this card, and peculiarity of the flock is a sum of peculiarities of all cards in the flock.

Mythic cards give bonuses to their owned depending on their rarity. These bonuses will be revealed in the future, but they may include
free Yellow tokens (with vesting claim mechanism), an ability to change existing cards, stealing / fighting other cards, etc.

### MAINTAINER_ROLE

```solidity
bytes32 MAINTAINER_ROLE
```

### issuer

```solidity
address issuer
```

### _usedVouchers

```solidity
mapping(bytes32 => bool) _usedVouchers
```

### ducklingCollectionId

```solidity
uint8 ducklingCollectionId
```

### zombeakCollectionId

```solidity
uint8 zombeakCollectionId
```

### mythicCollectionId

```solidity
uint8 mythicCollectionId
```

### RARITIES_NUM

```solidity
uint8 RARITIES_NUM
```

### MAX_PACK_SIZE

```solidity
uint8 MAX_PACK_SIZE
```

### FLOCK_SIZE

```solidity
uint8 FLOCK_SIZE
```

### collectionGeneIdx

```solidity
uint8 collectionGeneIdx
```

### rarityGeneIdx

```solidity
uint8 rarityGeneIdx
```

### flagsGeneIdx

```solidity
uint8 flagsGeneIdx
```

### generativeGenesOffset

```solidity
uint8 generativeGenesOffset
```

### collectionsGeneValuesNum

```solidity
uint8[][3] collectionsGeneValuesNum
```

### collectionsGeneDistributionTypes

```solidity
uint32[3] collectionsGeneDistributionTypes
```

### maxPeculiarity

```solidity
uint16 maxPeculiarity
```

### MYTHIC_DISPERSION

```solidity
uint8 MYTHIC_DISPERSION
```

### mythicAmount

```solidity
uint8 mythicAmount
```

### rarityChances

```solidity
uint32[] rarityChances
```

### collectionMutationChances

```solidity
uint32[] collectionMutationChances
```

### geneMutationChance

```solidity
uint32[] geneMutationChance
```

### geneInheritanceChances

```solidity
uint32[] geneInheritanceChances
```

### duckiesContract

```solidity
contract ERC20Burnable duckiesContract
```

### ducklingsContract

```solidity
contract IDucklings ducklingsContract
```

### treasureVaultAddress

```solidity
address treasureVaultAddress
```

### mintPrice

```solidity
uint256 mintPrice
```

### meldPrices

```solidity
uint256[4] meldPrices
```

### constructor

```solidity
constructor(address duckiesAddress, address ducklingsAddress, address treasureVaultAddress_) public
```

Sets Duckies, Ducklings and Treasure Vault addresses, minting and melding prices and other game config.

_Grants DEFAULT_ADMIN_ROLE and MAINTAINER_ROLE to the deployer._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| duckiesAddress | address | Address of Duckies ERC20 contract. |
| ducklingsAddress | address | Address of Ducklings ERC721 contract. |
| treasureVaultAddress_ | address | Address of Treasure Vault contract. |

### setIssuer

```solidity
function setIssuer(address account) external
```

Sets the issuer of Vouchers.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of a new issuer. |

### useVouchers

```solidity
function useVouchers(struct IVoucher.Voucher[] vouchers, bytes signature) external
```

Use multiple Vouchers. Check the signature and invoke internal function for each voucher.

_Vouchers are issued by the Back-End and signed by the issuer._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| vouchers | struct IVoucher.Voucher[] | Array of Vouchers to use. |
| signature | bytes | Vouchers signed by the issuer. |

### useVoucher

```solidity
function useVoucher(struct IVoucher.Voucher voucher, bytes signature) external
```

Use a single Voucher. Check the signature and invoke internal function.

_Vouchers are issued by the Back-End and signed by the issuer._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | Voucher to use. |
| signature | bytes | Voucher signed by the issuer. |

### _useVoucher

```solidity
function _useVoucher(struct IVoucher.Voucher voucher) internal
```

Check the validity of a voucher, decode voucher params and mint or meld tokens depending on voucher's type. Emits VoucherUsed event. Internal function.

_Vouchers are issued by the Back-End and signed by the issuer._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | Voucher to use. |

### _requireValidVoucher

```solidity
function _requireValidVoucher(struct IVoucher.Voucher voucher) internal view
```

Check the validity of a voucher, reverts if invalid.

_Voucher address must be this contract, beneficiary must be msg.sender, voucher must not be used before, voucher must not be expired._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| voucher | struct IVoucher.Voucher | Voucher to check. |

### _requireCorrectSigner

```solidity
function _requireCorrectSigner(bytes encodedData, bytes signature, address signer) internal pure
```

Check that `signatures is `encodedData` signed by `signer`. Reverts if not.

_Check that `signatures is `encodedData` signed by `signer`. Reverts if not._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| encodedData | bytes | Data to check. |
| signature | bytes | Signature to check. |
| signer | address | Address of the signer. |

### getMintPrice

```solidity
function getMintPrice() external view returns (uint256)
```

Get the mint price in Duckies with decimals.

_Get the mint price in Duckies with decimals._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | mintPrice Mint price in Duckies with decimals. |

### setMintPrice

```solidity
function setMintPrice(uint256 price) external
```

Set the mint price in Duckies without decimals.

_Require MAINTAINER_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| price | uint256 | Mint price in Duckies without decimals. |

### getMeldPrices

```solidity
function getMeldPrices() external view returns (uint256[4])
```

Get the meld price for each 'rarity' in Duckies with decimals.

_Get the meld price for each 'rarity' in Duckies with decimals._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256[4] | meldPrices Array of meld prices in Duckies with decimals. |

### setMeldPrices

```solidity
function setMeldPrices(uint256[4] prices) external
```

Set the meld price for each 'rarity' in Duckies without decimals.

_Require MAINTAINER_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| prices | uint256[4] | Array of meld prices in Duckies without decimals. |

### getCollectionsGeneValues

```solidity
function getCollectionsGeneValues() external view returns (uint8[][3], uint8)
```

Get number of gene values for all collections and a number of different Mythic tokens.

_Get number of gene values for all collections and a number of different Mythic tokens._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8[][3] | collectionsGeneValuesNum Arrays of number of gene values for all collections and a mythic amount. |
| [1] | uint8 |  |

### getCollectionsGeneDistributionTypes

```solidity
function getCollectionsGeneDistributionTypes() external view returns (uint32[3])
```

Get gene distribution types for all collections.

_Get gene distribution types for all collections._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint32[3] | collectionsGeneDistributionTypes Arrays of gene distribution types for all collections. |

### setDucklingGeneValues

```solidity
function setDucklingGeneValues(uint8[] duckingGeneValuesNum) external
```

Set gene values number for each gene for Duckling collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| duckingGeneValuesNum | uint8[] | Array of gene values number for each gene for Duckling collection. |

### setDucklingGeneDistributionTypes

```solidity
function setDucklingGeneDistributionTypes(uint32 ducklingGeneDistrTypes) external
```

Set gene distribution types for Duckling collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| ducklingGeneDistrTypes | uint32 | Gene distribution types for Duckling collection. |

### setZombeakGeneValues

```solidity
function setZombeakGeneValues(uint8[] zombeakGeneValuesNum) external
```

Set gene values number for each gene for Zombeak collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| zombeakGeneValuesNum | uint8[] | Array of gene values number for each gene for Duckling collection. |

### setZombeakGeneDistributionTypes

```solidity
function setZombeakGeneDistributionTypes(uint32 zombeakGeneDistrTypes) external
```

Set gene distribution types for Zombeak collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| zombeakGeneDistrTypes | uint32 | Gene distribution types for Zombeak collection. |

### setMythicAmount

```solidity
function setMythicAmount(uint8 amount) external
```

Set number of different Mythic tokens.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| amount | uint8 | Number of different Mythic tokens. |

### setMythicGeneValues

```solidity
function setMythicGeneValues(uint8[] mythicGeneValuesNum) external
```

Set gene values number for each gene for Mythic collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| mythicGeneValuesNum | uint8[] | Array of gene values number for each gene for Mythic collection. |

### setMythicGeneDistributionTypes

```solidity
function setMythicGeneDistributionTypes(uint32 mythicGeneDistrTypes) external
```

Set gene distribution types for Mythic collection.

_Require DEFAULT_ADMIN_ROLE to call._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| mythicGeneDistrTypes | uint32 | Gene distribution types for Mythic collection. |

### mintPack

```solidity
function mintPack(uint8 size) external
```

Mint a pack with `size` of Ducklings. Transfer Duckies from the sender to the TreasureVault.

_`Size` must be less than or equal to `MAX_PACK_SIZE`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| size | uint8 | Number of Ducklings in the pack. |

### _mintPackTo

```solidity
function _mintPackTo(address to, uint8 amount, bool isTransferable) internal returns (uint256[] tokenIds)
```

Mint a pack with `amount` of Ducklings to `to` and set transferable flag for each token. Internal function.

_`amount` must be less than or equal to `MAX_PACK_SIZE`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to mint the pack to. |
| amount | uint8 | Number of Ducklings in the pack. |
| isTransferable | bool | Transferable flag for each token. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenIds | uint256[] | Array of minted token IDs. |

### _generateGenome

```solidity
function _generateGenome(uint8 collectionId) internal returns (uint256)
```

Generate genome for Duckling or Zombeak.

_Generate and set all genes from a corresponding collection._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| collectionId | uint8 | Collection ID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Generated genome. |

### _generateRarity

```solidity
function _generateRarity() internal returns (enum IDuckyFamily.Rarities)
```

Generate rarity for a token.

_Generate rarity using rarity chances._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | enum IDuckyFamily.Rarities | rarity Generated rarity. |

### _generateAndSetGenes

```solidity
function _generateAndSetGenes(uint256 genome, uint8 collectionId) internal returns (uint256)
```

Generate and set all genes from a corresponding collection to `genome`.

_Generate and set all genes from a corresponding collection to `genome`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genome | uint256 | Genome to set genes to. |
| collectionId | uint8 | Collection ID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Genome with set genes. |

### _generateAndSetGene

```solidity
function _generateAndSetGene(uint256 genome, uint8 geneIdx, uint8 geneValuesNum, enum IDuckyFamily.GeneDistributionTypes distrType) internal returns (uint256)
```

Generate and set a gene with `geneIdx` to `genome`.

_Generate and set a gene with `geneIdx` to `genome`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genome | uint256 | Genome to set a gene to. |
| geneIdx | uint8 | Gene index. |
| geneValuesNum | uint8 | Number of gene values. |
| distrType | enum IDuckyFamily.GeneDistributionTypes | Gene distribution type. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Genome with set gene. |

### _generateMythicGenome

```solidity
function _generateMythicGenome(uint256[] genomes) internal returns (uint256)
```

Generate mythic genome based on melding `genomes`.

_Calculates flock peculiarity, and randomizes UniqId corresponding to the peculiarity._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Array of genomes to meld into Mythic. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Generated Mythic genome. |

### meldFlock

```solidity
function meldFlock(uint256[] meldingTokenIds) external
```

Meld tokens with `meldingTokenIds` into a new token. Calls internal function.

_Meld tokens with `meldingTokenIds` into a new token._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| meldingTokenIds | uint256[] | Array of token IDs to meld. |

### _meldOf

```solidity
function _meldOf(address owner, uint256[] meldingTokenIds, bool isTransferable) internal returns (uint256)
```

Meld tokens with `meldingTokenIds` into a new token. Internal function.

_Check `owner` is indeed the owner of `meldingTokenIds`. Burn NFTs with `meldingTokenIds`. Transfers Duckies to the TreasureVault._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| owner | address |  |
| meldingTokenIds | uint256[] | Array of token IDs to meld. |
| isTransferable | bool | Whether the new token is transferable. |

### _requireGenomesSatisfyMelding

```solidity
function _requireGenomesSatisfyMelding(uint256[] genomes) internal pure
```

Check that `genomes` satisfy melding rules. Reverts if not.

_Check that `genomes` satisfy melding rules. Reverts if not._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Array of genomes to check. |

### _meldGenomes

```solidity
function _meldGenomes(uint256[] genomes) internal returns (uint256)
```

Meld `genomes` into a new genome.

_Meld `genomes` into a new genome gene by gene. Set the corresponding collection_

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Array of genomes to meld. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | meldedGenome Melded genome. |

### _isCollectionMutating

```solidity
function _isCollectionMutating(enum IDuckyFamily.Rarities rarity) internal returns (bool)
```

Randomize if collection is mutating.

_Randomize if collection is mutating._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| rarity | enum IDuckyFamily.Rarities | Rarity of the collection. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isMutating True if mutating, false otherwise. |

### _meldGenes

```solidity
function _meldGenes(uint256[] genomes, uint8 gene, uint8 maxGeneValue, enum IDuckyFamily.GeneDistributionTypes geneDistrType) internal returns (uint8)
```

Meld `gene` from `genomes` into a new gene value.

_Meld `gene` from `genomes` into a new gene value. Gene mutation and inheritance are applied._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genomes | uint256[] | Array of genomes to meld. |
| gene | uint8 | Gene to be meld. |
| maxGeneValue | uint8 | Max gene value. |
| geneDistrType | enum IDuckyFamily.GeneDistributionTypes | Gene distribution type. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | geneValue Melded gene value. |

### _getDistributionType

```solidity
function _getDistributionType(uint32 distributionTypes, uint8 idx) internal pure returns (enum IDuckyFamily.GeneDistributionTypes)
```

Get gene distribution type.

_Get gene distribution type._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| distributionTypes | uint32 | Distribution types. |
| idx | uint8 | Index of the gene. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | enum IDuckyFamily.GeneDistributionTypes | Gene distribution type. |

### _generateUnevenGeneValue

```solidity
function _generateUnevenGeneValue(uint8 valuesNum) internal returns (uint8)
```

Generate uneven gene value given the maximum number of values.

_Generate uneven gene value using quadratic algorithm described below._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| valuesNum | uint8 | Maximum number of gene values. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint8 | geneValue Gene value. |

### _calcMaxPeculiarity

```solidity
function _calcMaxPeculiarity() internal view returns (uint16)
```

Calculate max peculiarity for a current Duckling config.

_Sum up number of uneven gene values._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint16 | maxPeculiarity Max peculiarity. |

### _calcPeculiarity

```solidity
function _calcPeculiarity(uint256 genome) internal view returns (uint16)
```

Calculate peculiarity for a given genome.

_Sum up number of uneven gene values._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| genome | uint256 | Genome. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint16 | peculiarity Peculiarity. |

### _calcUniqIdGenerationParams

```solidity
function _calcUniqIdGenerationParams(uint16 pivotalUniqId, uint16 maxUniqId) internal pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength)
```

Calculate `leftEndUniqId` and `uniqIdSegmentLength` for UniqId generation.

_Then UniqId is generated by adding a random number [0, `uniqIdSegmentLength`) to `leftEndUniqId`._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pivotalUniqId | uint16 | Pivotal UniqId. |
| maxUniqId | uint16 | Max UniqId. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| leftEndUniqId | uint16 | Left end of the UniqId segment. |
| uniqIdSegmentLength | uint16 | Length of the UniqId segment. |

