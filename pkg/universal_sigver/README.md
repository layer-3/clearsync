# Universal Signer Verifier package

## Overview

Universal Signer Verifier is a golang library that provides helper functions to verify signatures.

### Verify

You can use `Verify(ctx, signer, messageHash, signature) (bool, error)` to verify an ECDSA, ERC-1271 or ERC-6492 signature.

### Sign

You can use `PackERC6492Sig(ctx, ownerAddress, index, sig) ([]byte, error)` to pack a smart wallet signature into an ERC-6492 signature.

## Usage

To use the Universal Signer Verifier package, you need to create a client, providing an ethclient, smart wallet config and an address of an EntryPoint contract.
If you don't provide an address of the latter one, the default (v0.6) address will be used.

```go
import (
  "github.com/layer-3/clearsync/pkg/universal_sigver"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
  sigver, err := universal_sigver.NewUniversalSigver(client, config, entryPointAddress)
  if err != nil {
    panic(err)
  }

  // Use verifier to verify signatures
  sigver.Verify(...)
}
```
