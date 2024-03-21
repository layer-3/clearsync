package signer

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

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

func (s1 Signature) Equal(s2 Signature) bool {
	return bytes.Equal(s1.S, s2.S) && bytes.Equal(s1.R, s2.R) && s1.V == s2.V
}

func (s Signature) MarshalJSON() ([]byte, error) {
	hex := hexutil.Encode(s.Raw())
	return json.Marshal(hex)
}

func (s *Signature) UnmarshalJSON(b []byte) error {
	var hex string
	err := json.Unmarshal(b, &hex)
	if err != nil {
		return err
	}
	joined, err := hexutil.Decode(hex)
	if err != nil {
		return err
	}

	// If the signature is all zeros, we consider it to be the empty signature
	if allZero(joined) {
		return nil
	}

	if len(joined) != 65 {
		return fmt.Errorf("signature must be 65 bytes long or a zero string, received %d bytes", len(joined))
	}

	s.R = joined[:32]
	s.S = joined[32:64]
	s.V = joined[64]
	return nil
}

// allZero returns true if all bytes in the slice are zero false otherwise
func allZero(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
