package wallet

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/stretchr/testify/require"
)

func TestNewFromSeed(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		seed := []byte("example seed phrase")
		w, err := NewFromSeed(seed)
		require.NoError(t, err)
		require.NotNil(t, w)
	})

	t.Run("Empty seed", func(t *testing.T) {
		w, err := NewFromSeed([]byte{})
		require.EqualError(t, err, "empty seed")
		require.Nil(t, w)
	})
}

func TestWallet_Derive(t *testing.T) {
	t.Run("Default derivation path", func(t *testing.T) {
		seed := []byte("example seed phrase")
		w, err := NewFromSeed(seed)
		require.NoError(t, err)

		path := accounts.DefaultBaseDerivationPath
		account, err := w.Derive(path)
		require.NoError(t, err)

		require.Equal(t, "0x86FCB07ac2E29B897C4AD3632605161FD4c907e9", account.Address().Hex())
		require.Equal(t, "551fa8ca244673bcdf53d7b031704a5728e2b15c04de05d9d6aa14b9ecf89d42", hex.EncodeToString(account.privateKey.D.Bytes()))
	})
}
