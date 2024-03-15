package universal_sigver

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
	"github.com/stretchr/testify/require"
)

func TestIsERC6492Sig(t *testing.T) {
	// empty sig
	require.False(t, IsERC6492Sig([]byte{}))

	// ECDSA sig
	pvk, err := crypto.GenerateKey()
	require.NoError(t, err)
	signer := signer_pkg.NewLocalSigner(pvk)

	msg := []byte("hello")
	sig, err := signer_pkg.SignEthMessage(signer, msg)
	require.NoError(t, err)
	require.False(t, IsERC6492Sig(sig.Raw()))

	// ERC-6492 sig
	factoryCallData := hexutil.MustDecode("0xdeadbeef")
	erc6492Sig := PackERC6492Sig(common.HexToAddress("0x4242"), factoryCallData, sig.Raw())
	require.True(t, IsERC6492Sig(erc6492Sig))
}

func TestPackERC6492Sig(t *testing.T) {
	factoryAddress := common.HexToAddress("0x4242")
	pvk, err := crypto.GenerateKey()
	require.NoError(t, err)
	signer := signer_pkg.NewLocalSigner(pvk)

	msg := []byte("hello again")
	sig, err := signer_pkg.SignEthMessage(signer, msg)
	require.NoError(t, err)

	factoryCallData := hexutil.MustDecode("0xdeadbeef")

	calldata := PackERC6492Sig(factoryAddress, factoryCallData, sig.Raw())
	len := len(calldata)
	require.Equal(t, erc6492MagicValue, hexutil.Encode(calldata[len-32:len]))

	args := abi.Arguments{
		{Name: "factory", Type: address},
		{Name: "calldata", Type: bytes},
		{Name: "signature", Type: bytes},
	}

	unpacked, err := args.Unpack(calldata[:len-32])
	require.NoError(t, err)

	unpackedFactory, ok := unpacked[0].(common.Address)
	require.True(t, ok)
	require.Equal(t, factoryAddress, unpackedFactory)

	unpackedFactoryCallData, ok := unpacked[1].([]byte)
	require.True(t, ok)
	require.Equal(t, hexutil.Encode(factoryCallData), hexutil.Encode(unpackedFactoryCallData))

	unpackedSig, ok := unpacked[2].([]byte)
	require.True(t, ok)
	require.Equal(t, sig.Raw(), unpackedSig)
}

func TestUnpackERC6492(t *testing.T) {
	// valid sig
	factoryAddress := common.HexToAddress("0x4242")
	pvk, err := crypto.GenerateKey()
	require.NoError(t, err)
	signer := signer_pkg.NewLocalSigner(pvk)

	msg := []byte("hello again")
	sig, err := signer_pkg.SignEthMessage(signer, msg)
	require.NoError(t, err)

	factoryCallData := hexutil.MustDecode("0xc001beef")
	erc6492Sig := PackERC6492Sig(factoryAddress, factoryCallData, sig.Raw())

	unpackedFactoryAddress, unpackedFactoryCallData, unpackedSig, err := UnpackERC6492Sig(erc6492Sig)
	require.NoError(t, err)
	require.Equal(t, factoryAddress, unpackedFactoryAddress)
	require.Equal(t, factoryCallData, unpackedFactoryCallData)
	require.Equal(t, sig.Raw(), unpackedSig)

	// not ERC6492 sig
	_, _, _, err = UnpackERC6492Sig(sig.Raw())
	require.EqualError(t, err, ErrNotERC6492Sig.Error())

	// Corrupted sig
	corERC6492Sig := hexutil.MustDecode("0xdeadbeef" + erc6492MagicValue[2:])
	_, _, _, err = UnpackERC6492Sig(corERC6492Sig)
	require.EqualError(t, err, ErrCorruptedERC6492Sig.Error())
}
