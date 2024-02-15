package signer

import (
	"strings"
	"testing"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var _ Signer = &LocalSigner{}

func TestSignerSign(t *testing.T) {
	t.Parallel()

	privateKey, err := ecrypto.GenerateKey()
	assert.NoError(t, err)

	signer := NewLocalSigner(privateKey)
	msg := []byte(strings.Repeat("a", 32))
	sig, err := signer.Sign(msg)
	assert.NoError(t, err)

	recoveredPubKey, err := sig.RecoverPublicKey(msg)
	assert.NoError(t, err)
	assert.Equal(t, *recoveredPubKey, privateKey.PublicKey)
}
