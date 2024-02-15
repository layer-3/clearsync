package signer

import (
	"crypto/ecdsa"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	nc "github.com/statechannels/go-nitro/crypto"
)

type Signature struct{ nc.Signature }

func NewSignature(R, S []byte, V byte) Signature {
	return Signature{nc.Signature{R: R, S: S, V: V}}
}

func NewSignatureFromBytes(sig []byte) Signature {
	return Signature{nc.SplitSignature(sig)}
}

func (sig Signature) String() string {
	return sig.Signature.ToHexString()
}

func (sig Signature) Raw() (concatenatedSignature []byte) {
	concatenatedSignature = append(concatenatedSignature, sig.R...)
	concatenatedSignature = append(concatenatedSignature, sig.S...)
	concatenatedSignature = append(concatenatedSignature, sig.V)
	return
}

func (sig Signature) RecoverPublicKey(msg []byte) (*ecdsa.PublicKey, error) {
	pubKey, err := secp256k1.RecoverPubkey(msg, sig.Raw())
	if err != nil {
		return nil, err
	}
	ecdsaPubKey, err := ecrypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return nil, err
	}
	return ecdsaPubKey, nil
}
