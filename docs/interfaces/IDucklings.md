# Solidity API

## IDucklings

This interface defines the Ducklings ERC721-compatible contract,
which provides basic functionality for minting, burning and querying information about the tokens.

### TokenNotTransferable

```solidity
error TokenNotTransferable(uint256 tokenId)
```

Token not transferable error. Is used when trying to transfer a token that is not transferable.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Token Id that is not transferable. |

### InvalidMagicNumber

```solidity
error InvalidMagicNumber(uint8 magicNumber)
```

Invalid magic number error. Is used when trying to mint a token with an invalid magic number.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| magicNumber | uint8 | Magic number that is invalid. |

### Duckling

```solidity
struct Duckling {
  uint256 genome;
  uint64 birthdate;
}
```

### Minted

```solidity
event Minted(address to, uint256 tokenId, uint256 genome, uint64 birthdate, uint256 chainId)
```

Minted event. Is emitted when a token is minted.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address of the token owner. |
| tokenId | uint256 | Id of the minted token. |
| genome | uint256 | Genome of the minted token. |
| birthdate | uint64 | Birthdate of the minted token. |
| chainId | uint256 | Id of the chain where the token was minted. |

### isOwnerOf

```solidity
function isOwnerOf(address account, uint256 tokenId) external view returns (bool)
```

Check whether `account` is owner of `tokenId`.

_Revert if `account` is address(0) or `tokenId` does not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address to check. |
| tokenId | uint256 | Token Id to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isOwnerOf True if `account` is owner of `tokenId`, false otherwise. |

### isOwnerOfBatch

```solidity
function isOwnerOfBatch(address account, uint256[] tokenIds) external view returns (bool)
```

Check whether `account` is owner of `tokenIds`.

_Revert if `account` is address(0) or any of `tokenIds` do not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address to check. |
| tokenIds | uint256[] | Token Ids to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isOwnerOfBatch True if `account` is owner of `tokenIds`, false otherwise. |

### getGenome

```solidity
function getGenome(uint256 tokenId) external view returns (uint256)
```

Get genome of `tokenId`.

_Revert if `tokenId` does not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Token Id to get the genome of. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | genome Genome of `tokenId`. |

### getGenomes

```solidity
function getGenomes(uint256[] tokenIds) external view returns (uint256[])
```

Get genomes of `tokenIds`.

_Revert if any of `tokenIds` do not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenIds | uint256[] | Token Ids to get the genomes of. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256[] | genomes Genomes of `tokenIds`. |

### mintTo

```solidity
function mintTo(address to, uint256 genome) external returns (uint256)
```

Mint token with `genome` to `to`. Emits Minted event.

_Revert if `to` is address(0) or `genome` has wrong magic number._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to mint token to. |
| genome | uint256 | Genome of the token to mint. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256 | tokenId Id of the minted token. |

### mintBatchTo

```solidity
function mintBatchTo(address to, uint256[] genomes) external returns (uint256[])
```

Mint tokens with `genomes` to `to`. Emits Minted event for each token.

_Revert if `to` is address(0) or any of `genomes` has wrong magic number._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to mint tokens to. |
| genomes | uint256[] | Genomes of the tokens to mint. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint256[] | tokenIds Ids of the minted tokens. |

### burn

```solidity
function burn(uint256 tokenId) external
```

Burn token with `tokenId`.

_Revert if `tokenId` does not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Id of the token to burn. |

### burnBatch

```solidity
function burnBatch(uint256[] tokenIds) external
```

Burn tokens with `tokenIds`.

_Revert if any of `tokenIds` do not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenIds | uint256[] | Ids of the tokens to burn. |

