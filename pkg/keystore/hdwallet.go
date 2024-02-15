package keystore

import "github.com/layer-3/clearsync/pkg/signer"

// TODO: remove wallet pkg and implement it here

type HdWallet struct {
	seed string
}

func NewHdWallet(seed string) HdWallet {
	return HdWallet{
		seed: seed,
	}
}

// TODO:
// 1. Generate derivation path based on unique index
// 2. Derive account from derivation path
// 3. Create signer from account's private key
func (w *HdWallet) GetOrCreateSigner(uniqueIndex uint32) (signer.Signer, error) {
	return signer.NewLocalSigner(nil), nil
}
