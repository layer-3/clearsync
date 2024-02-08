package wallet

import (
	"crypto/ecdsa"
	"errors"
	"sync"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Account represents a derived account.
type Account struct {
	path       accounts.DerivationPath
	privateKey *ecdsa.PrivateKey
}

// Path returns the account's derivation path.
func (a *Account) Path() accounts.DerivationPath {
	return a.path
}

// Address returns the Ethereum address of the account.
func (a *Account) Address() common.Address {
	return crypto.PubkeyToAddress(a.privateKey.PublicKey)
}

// PrivateKey returns the account's private key.
func (a *Account) PrivateKey() *ecdsa.PrivateKey {
	return a.privateKey
}

// PublicKey returns the account's public key.
func (a *Account) PublicKey() *ecdsa.PublicKey {
	return &a.privateKey.PublicKey
}

// Wallet represents a master wallet that can generate accounts.
type Wallet struct {
	masterKey *hdkeychain.ExtendedKey
	mx        sync.RWMutex
}

// NewFromSeed creates a new Wallet instance from a given seed.
func NewFromSeed(seed []byte) (*Wallet, error) {
	if len(seed) == 0 {
		return nil, errors.New("empty seed")
	}
	key, err := hdkeychain.NewMaster(seed, &chaincfg.Params{})
	if err != nil {
		return nil, err
	}
	return &Wallet{
		masterKey: key,
	}, nil
}

// Derive generates a new Account based on a given derivation path.
func (w *Wallet) Derive(path accounts.DerivationPath) (Account, error) {
	w.mx.RLock()
	priv, err := w.derivePrivateKey(path)
	w.mx.RUnlock()
	if err != nil {
		return Account{}, err
	}

	return Account{
		path:       path,
		privateKey: priv,
	}, nil
}

// derivePrivateKey is a helper function to derive an ECDSA private key
func (w *Wallet) derivePrivateKey(path accounts.DerivationPath) (*ecdsa.PrivateKey, error) {
	var err error
	key := w.masterKey
	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return privateKey.ToECDSA(), nil
}
