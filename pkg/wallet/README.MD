# Wallet

The package provides a Go implementation for creating Ethereum wallets.

Generate Ethereum accounts, including private and public keys and addresses, using a seed and a derivation path.

## Types

- **Wallet:** Main structure holding the master key. It includes methods for account derivation.
- **Account:** Represents a derived account, including its path, address, and private key, with methods to access these fields.

## Usage

### Creating a New Wallet

To create a new wallet from a seed, use the `NewFromSeed` function. It requires a seed byte slice as input and returns a new `Wallet` instance.

```go
seed := []byte("your secure seed here")
wallet, err := wallet.NewFromSeed(seed)
if err != nil {
    // handle error
}
```

### Deriving an Account

Once you have a `Wallet` instance, you can derive new accounts using the `Derive` method. It accepts a derivation path and returns an `Account` instance containing the address, private key, and path.

```go
path := accounts.DefaultBaseDerivationPath
account, err := wallet.Derive(path)
if err != nil {
    // handle error
}

address := account.Address()
privateKey := account.PrivateKey()
publicKey := account.PublicKey()
```
