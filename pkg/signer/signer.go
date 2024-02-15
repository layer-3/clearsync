package signer

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Signer interface {
	Sign(msg []byte) (Signature, error)

	PublicKey() *ecdsa.PublicKey
	CommonAddress() common.Address
}
