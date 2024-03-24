package signer

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
)

type LocalSigner struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func NewLocalSigner(privateKey *ecdsa.PrivateKey) LocalSigner {
	publicKey := privateKey.PublicKey
	return LocalSigner{
		privateKey: privateKey,
		address:    ecrypto.PubkeyToAddress(publicKey),
	}
}

func (s LocalSigner) Sign(msg []byte) (Signature, error) {
	sigBytes, err := ecrypto.Sign(msg, s.privateKey)
	if err != nil {
		return Signature{}, err
	}

	return NewSignatureFromBytes(sigBytes), nil
}

func (s LocalSigner) PublicKey() *ecdsa.PublicKey {
	return &s.privateKey.PublicKey
}

func (s LocalSigner) CommonAddress() common.Address {
	return s.address
}

func (s LocalSigner) PrivateKey() *ecdsa.PrivateKey {
	return s.privateKey
}
