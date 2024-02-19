# Keystore

The package provides a Go implementation for deriving Ethereum wallets and creating Signers.

## Types

- **KeyStore Interface:** A factory for `signer.Signer` instances.
- **HdWallet:** Represents a master wallet capable of generating Ethereum signers. It encapsulates the master key and provides thread-safe access to key derivation and signer creation.

## Usage

### Creating a New HD Wallet

To create a new `HdWallet` instance from a seed, use the `NewHdWallet` function. This function requires a seed string as input and returns a new `HdWallet` instance or an error if the seed is empty.

```go
seed := "your secure seed here"
hdWallet, err := keystore.NewHdWallet(seed)
if err != nil {
    // handle error
}
```

### Creating a Signer

With an `HdWallet` instance, you can generate or retrieve a signer for transaction signing using the `GetOrCreateSigner` method. This method accepts a unique account index (uint32) and returns a `signer.Signer` instance.

```go
uniqueIndex := uint32(0) // Example index
signer, err := hdWallet.GetOrCreateSigner(uniqueIndex)
if err != nil {
    // handle error
}

// Use the signer for signing transactions
```

### Deriving a Private Key

The `HdWallet` structure includes a private method `derivePrivateKey`, used internally to derive an ECDSA private key from the master key based on a specified derivation path. The `HdWallet` ensures thread-safe operations using a read-write mutex, allowing concurrent access to key derivation functions.
