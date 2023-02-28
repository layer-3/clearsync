# Solidity API

## Garden

### CircularReferrers

```solidity
error CircularReferrers(address target, address base)
```

### BountyAlreadyClaimed

```solidity
error BountyAlreadyClaimed(bytes32 bountyCodeHash)
```

### InvalidBounty

```solidity
error InvalidBounty(struct Garden.Bounty bounty)
```

### InsufficientTokenBalance

```solidity
error InsufficientTokenBalance(address token, uint256 expected, uint256 actual)
```

### IncorrectSigner

```solidity
error IncorrectSigner(address expected, address actual)
```

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

### REFERRAL_MAX_DEPTH

```solidity
uint8 REFERRAL_MAX_DEPTH
```

### _REFERRAL_PAYOUT_DIVIDER

```solidity
uint8 _REFERRAL_PAYOUT_DIVIDER
```

### Bounty

```solidity
struct Bounty {
  uint256 amount;
  address tokenAddress;
  address beneficiary;
  bool isPaidToReferrers;
  uint8[5] referrersPayouts;
  address referrer;
  uint64 expire;
  uint32 chainId;
  bytes32 bountyCodeHash;
}
```

### _referrerOf

```solidity
mapping(address => address) _referrerOf
```

### _issuer

```solidity
address _issuer
```

### _claimedBounties

```solidity
mapping(bytes32 => bool) _claimedBounties
```

### AffiliateRegistered

```solidity
event AffiliateRegistered(address affiliate, address referrer)
```

### BountyClaimed

```solidity
event BountyClaimed(address wallet, bytes32 bountyCodeHash, uint32 chainId, address tokenAddress)
```

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize() public
```

### setIssuer

```solidity
function setIssuer(address account) external
```

### getIssuer

```solidity
function getIssuer() external view returns (address)
```

### transferTokenBalanceToPartner

```solidity
function transferTokenBalanceToPartner(address tokenAddress, address partner) public
```

### _registerReferrer

```solidity
function _registerReferrer(address child, address parent) internal
```

### _requireNotReferrerOf

```solidity
function _requireNotReferrerOf(address target, address base) internal view
```

### claimBounties

```solidity
function claimBounties(struct Garden.Bounty[] bounties, bytes signature) external
```

### claimBounty

```solidity
function claimBounty(struct Garden.Bounty bounty, bytes signature) external
```

### _claimBounty

```solidity
function _claimBounty(struct Garden.Bounty bounty) internal
```

### _requireValidBounty

```solidity
function _requireValidBounty(struct Garden.Bounty bounty) internal view
```

### _requireSufficientContractBalance

```solidity
function _requireSufficientContractBalance(contract ERC20Upgradeable token, uint256 expected) internal view
```

### _requireCorrectSigner

```solidity
function _requireCorrectSigner(bytes encodedData, bytes signature, address signer) internal pure
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

