package keystore

import (
	"crypto/ecdsa"
	"errors"
	"sync"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/layer-3/clearsync/pkg/signer"
)

// HdWallet represents a master wallet that can create signers.
type HdWallet struct {
	masterKey *hdkeychain.ExtendedKey
	mx        sync.RWMutex
}

// NewHdWallet creates a new NewHdWallet based on seed.
func NewHdWallet(seed string) (*HdWallet, error) {
	if seed == "" {
		return nil, errors.New("empty seed")
	}
	key, err := hdkeychain.NewMaster([]byte(seed), &chaincfg.Params{})
	if err != nil {
		return &HdWallet{}, err
	}

	return &HdWallet{
		masterKey: key,
	}, err
}

// GetOrCreateSigner creates a signer instance based on unique account index.
func (w *HdWallet) GetOrCreateSigner(uniqueIndex uint32) (signer.Signer, error) {
	// Define a derivation path and set the account component to be the uniqueIndex.
	path := accounts.DefaultBaseDerivationPath
	path[2] = uniqueIndex
	w.mx.RLock()
	priv, err := w.derivePrivateKey(path)
	w.mx.RUnlock()
	if err != nil {
		return nil, err
	}

	return signer.NewLocalSigner(priv), nil
}

// derivePrivateKey is a helper function to derive an ECDSA private key
func (w *HdWallet) derivePrivateKey(path accounts.DerivationPath) (*ecdsa.PrivateKey, error) {
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
