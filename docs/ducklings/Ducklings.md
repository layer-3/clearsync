# Solidity API

## Ducklings

### Duckling

```solidity
struct Duckling {
  uint256 gene;
  uint64 birthdate;
}
```

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

### API_SETTER_ROLE

```solidity
bytes32 API_SETTER_ROLE
```

### ROYALTIES_COLLECTOR_ROLE

```solidity
bytes32 ROYALTIES_COLLECTOR_ROLE
```

### ROYALTY_FEE

```solidity
uint32 ROYALTY_FEE
```

### MAX_MINT_PACK_SIZE

```solidity
uint8 MAX_MINT_PACK_SIZE
```

### BASE_DUCKIES_PER_MINT

```solidity
uint256 BASE_DUCKIES_PER_MINT
```

### _royaltiesCollector

```solidity
address _royaltiesCollector
```

### apiBaseURL

```solidity
string apiBaseURL
```

### salt

```solidity
bytes32 salt
```

### nextNewTokenId

```solidity
struct CountersUpgradeable.Counter nextNewTokenId
```

### tokenIdToDuckling

```solidity
mapping(uint256 => struct Ducklings.Duckling) tokenIdToDuckling
```

### traitWeights

```solidity
uint8[][] traitWeights
```

### meldWeights

```solidity
uint8[] meldWeights
```

### meldingZombeakWeights

```solidity
uint8[][3] meldingZombeakWeights
```

### duckiesContract

```solidity
contract ERC20BurnableUpgradeable duckiesContract
```

### Minted

```solidity
event Minted(uint256 mintedTokenId, uint256 mintedGene, address owner, uint256 chainId)
```

### Melded

```solidity
event Melded(uint256[5] meldingTokenIds, uint256 meldedTokenId, uint256 meldedGene, address owner, uint256 chainId)
```

### initialize

```solidity
function initialize(address ducklingsAddress) public
```

### _authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

_Function that should revert when `msg.sender` is not authorized to upgrade the contract. Called by
{upgradeTo} and {upgradeToAndCall}.

Normally, this function will use an xref:access.adoc[access control] modifier such as {Ownable-onlyOwner}.

```solidity
function _authorizeUpgrade(address) internal override onlyOwner {}
```_

### _burn

```solidity
function _burn(uint256 tokenId) internal
```

### tokenURI

```solidity
function tokenURI(uint256 tokenId) public view returns (string)
```

_See {IERC721Metadata-tokenURI}._

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) public view virtual returns (bool)
```

### setRoyaltyCollector

```solidity
function setRoyaltyCollector(address account) public
```

### getRoyaltyCollector

```solidity
function getRoyaltyCollector() public view returns (address)
```

### setAPIBaseURL

```solidity
function setAPIBaseURL(string apiBaseURL_) external
```

### ducklingsPerMint

```solidity
function ducklingsPerMint() external view returns (uint256)
```

### mintPack

```solidity
function mintPack(uint8 amount) external
```

### meld

```solidity
function meld(uint256[5] meldingTokenIds) external
```

### _mint

```solidity
function _mint() internal returns (uint256 tokenId, uint256 gene)
```

### _generateGene

```solidity
function _generateGene() internal returns (uint256)
```

### _generateTrait

```solidity
function _generateTrait(enum Gene.Traits trait) internal returns (uint8)
```

### _meldGenes

```solidity
function _meldGenes(uint256[5] genes) internal returns (uint256)
```

### _meldTraits

```solidity
function _meldTraits(uint256[5] genes, enum Gene.Traits trait) internal returns (uint8)
```

### _requireCorrectMeldingClasses

```solidity
function _requireCorrectMeldingClasses(uint256[5] genes) internal pure
```

### _requireCorrectMeldingTraits

```solidity
function _requireCorrectMeldingTraits(uint256[5] genes) internal pure
```

### _checkZombeak

```solidity
function _checkZombeak(enum Gene.Classes class) internal returns (bool)
```

### _ducklingsPerMint

```solidity
function _ducklingsPerMint() internal view returns (uint256)
```

### _requireCallerIsOwner

```solidity
function _requireCallerIsOwner(uint256[5] tokenIds) internal view
```

### _burnTokensAndGetGenes

```solidity
function _burnTokensAndGetGenes(uint256[5] tokenIds) internal returns (uint256[5])
```

### _maxTrait

```solidity
function _maxTrait(uint256[5] genes, enum Gene.Traits trait) internal pure returns (uint8)
```

### _randomWeightedNumber

```solidity
function _randomWeightedNumber(uint8[] weights) internal returns (uint8)
```

### _sum

```solidity
function _sum(uint8[] numbers) internal pure returns (uint256)
```

### _getChainId

```solidity
function _getChainId() internal view returns (uint256 id)
```

