# Solidity API

## DucklingsV1

This contract implements ERC721 and ERC2981 standards, stores and provides basic functionality for Ducklings NFT.
Ducklings expects other Game contracts to define more specific logic for Ducklings NFT.

Game contracts should be granted GAME_ROLE to be able to mint and burn tokens.
Ducklings defines specific query methods for Game contracts to retrieve specific NFT data.

Ducklings can be upgraded by an account with UPGRADER_ROLE to add certain functionality if needed.

### InvalidTokenId

```solidity
error InvalidTokenId(uint256 tokenId)
```

Is thrown when token with given Id does not exist.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Id of the token. |

### InvalidAddress

```solidity
error InvalidAddress(address addr)
```

Is thrown when given address is address(0).

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| addr | address | Invalid address. |

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

### GAME_ROLE

```solidity
bytes32 GAME_ROLE
```

### _royaltiesCollector

```solidity
address _royaltiesCollector
```

### _royaltyFee

```solidity
uint32 _royaltyFee
```

### apiBaseURL

```solidity
string apiBaseURL
```

### nextNewTokenId

```solidity
struct CountersUpgradeable.Counter nextNewTokenId
```

### tokenToDuckling

```solidity
mapping(uint256 => struct IDucklings.Duckling) tokenToDuckling
```

### initialize

```solidity
function initialize() external
```

Initializes the contract.
Grants DEFAULT_ADMIN_ROLE and UPGRADER_ROLE to the deployer.
Sets deployer to be Royalty collector, set royalty fee to 10%.

_This function is called only once during contract deployment._

### _authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Upgrades the contract.

_Requires UPGRADER_ROLE to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| newImplementation | address | Address of the new implementation. |

### _burn

```solidity
function _burn(uint256 tokenId) internal
```

Necessary override to specify what implementation of _burn to use.

_Necessary override to specify what implementation of _burn to use._

### tokenURI

```solidity
function tokenURI(uint256 tokenId) public view returns (string)
```

Composes an API URL for a given token that returns metadata.json.

_Concatenates `apiBaseURL` with token genome, dash (-) and token birthdate._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Id of the token. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | string | uri URL for the token metadata. |

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) public view virtual returns (bool)
```

Checks whether supplied `interface` is supported by the contract.

_Checks whether supplied `interface` is supported by the contract._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| interfaceId | bytes4 | Id of the interface. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | interfaceSupported true if interface is supported, false otherwise. |

### setRoyaltyCollector

```solidity
function setRoyaltyCollector(address account) public
```

Sets royalties collector.

_Requires DEFAULT_ADMIN_ROLE to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| account | address | Address of the royalties collector. |

### getRoyaltyCollector

```solidity
function getRoyaltyCollector() public view returns (address)
```

Returns royalties collector.

_Returns royalties collector._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | address | address Address of the royalties collector. |

### setRoyaltyFee

```solidity
function setRoyaltyFee(uint32 fee) public
```

Sets royalties fee.

_Requires DEFAULT_ADMIN_ROLE to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| fee | uint32 | Royalties fee in permyriad. |

### getRoyaltyFee

```solidity
function getRoyaltyFee() public view returns (uint32)
```

Returns royalties fee.

_Returns royalties fee._

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | uint32 | uint32 Royalties fee in permyriad. |

### setAPIBaseURL

```solidity
function setAPIBaseURL(string apiBaseURL_) external
```

Sets api server endpoint that is prepended to the tokenURI.

_Requires DEFAULT_ADMIN_ROLE to invoke._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| apiBaseURL_ | string | URL of the api server. |

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

### isTransferable

```solidity
function isTransferable(uint256 tokenId) external view returns (bool)
```

Check whether token with `tokenId` is transferable.

_Revert if `tokenId` does not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Token Id to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isTransferable True if token with `tokenId` is transferable, false otherwise. |

### _isTransferable

```solidity
function _isTransferable(uint256 tokenId) internal view returns (bool)
```

Check whether token with `tokenId` is transferable. Internal function.

_Revert if `tokenId` does not exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Token Id to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | isTransferable True if token with `tokenId` is transferable, false otherwise. |

### _beforeTokenTransfer

```solidity
function _beforeTokenTransfer(address from, address to, uint256 firstTokenId, uint256 batchSize) internal
```

Check whether token that is being transferred is transferable. Revert if not.

_Revert if token is not transferable._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| from | address | Address of the sender. |
| to | address | Address of the recipient. |
| firstTokenId | uint256 | Id of the token being transferred. |
| batchSize | uint256 |  |

### mintTo

```solidity
function mintTo(address to, uint256 genome) external returns (uint256 tokenId)
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
| tokenId | uint256 | Id of the minted token. |

### mintBatchTo

```solidity
function mintBatchTo(address to, uint256[] genomes) external returns (uint256[] tokenIds)
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
| tokenIds | uint256[] | Ids of the minted tokens. |

### _mintTo

```solidity
function _mintTo(address to, uint256 genome) internal returns (uint256 tokenId)
```

Mint token with `genome` to `to`. Emits Minted event. Internal function.

_Revert if `to` is address(0) or `genome` has wrong magic number._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| to | address | Address to mint token to. |
| genome | uint256 | Genome of the token to mint. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| tokenId | uint256 | Id of the minted token. |

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

