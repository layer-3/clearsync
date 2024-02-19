package keystore

import "github.com/layer-3/clearsync/pkg/signer"

// KeyStore is a signers factory
type KeyStore interface {
	GetOrCreateSigner(uniqueIndex uint32) (signer.Signer, error)
}
