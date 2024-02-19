package keystore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromSeed(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		seed := "example seed phrase"
		w, err := NewHdWallet(seed)
		require.NoError(t, err)
		require.NotNil(t, w)
	})

	t.Run("Empty seed", func(t *testing.T) {
		w, err := NewHdWallet("")
		require.EqualError(t, err, "empty seed")
		require.Nil(t, w)
	})
}

func TestWallet_Derive(t *testing.T) {
	t.Run("Default derivation path", func(t *testing.T) {
		seed := "example seed phrase"
		w, err := NewHdWallet(seed)
		require.NoError(t, err)

		signer, err := w.GetOrCreateSigner(623)
		require.NoError(t, err)

		require.Equal(t, "0x900FD92779c0714908b70636Cfda05832277C3E4", signer.CommonAddress().Hex())
	})
}
