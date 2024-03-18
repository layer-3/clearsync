package signer

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type Signature struct {
	R []byte
	S []byte
	V byte
}

func NewSignature(R, S []byte, V byte) Signature {
	return Signature{R: R, S: S, V: V}
}

func NewSignatureFromBytes(sig []byte) Signature {
	r := sig[:32]
	s := sig[32:64]
	v := sig[64]

	return NewSignature(r, s, v)
}

func (sig Signature) String() string {
	return hexutil.Encode(sig.Raw())
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
