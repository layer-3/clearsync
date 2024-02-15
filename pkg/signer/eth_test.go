package signer

import (
	"strings"
	"testing"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestLocalSignerSignEthMessage(t *testing.T) {
	t.Parallel()

	signer := generateSigner(t)
	message := []byte(strings.Repeat("0", 32))

	sig, err := SignEthMessage(signer, message)
	assert.NoError(t, err)

	pubkey, err := RecoverEthMessageSigner(sig, message)
	assert.NoError(t, err)
	assert.Equal(t, signer.PublicKey(), pubkey)

	address, err := RecoverEthMessageSignerAddress(sig, message)
	assert.NoError(t, err)
	assert.Equal(t, signer.CommonAddress(), address)
}

func generateSigner(t *testing.T) LocalSigner {
	privateKey, err := ecrypto.GenerateKey()
	assert.NoError(t, err)

	return NewLocalSigner(privateKey)
}
