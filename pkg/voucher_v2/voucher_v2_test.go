package voucher_v2

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/clearsync/pkg/abi/ivoucher_v2"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
)

func TestEncodeDecode(t *testing.T) {
	var codeHash [32]byte
	for i := 0; i < len(codeHash); i++ {
		codeHash[i] = byte(i)
	}

	tests := []struct {
		name    string
		voucher ivoucher_v2.IVoucherVoucher
	}{
		{
			name: "BasicVoucher",
			voucher: ivoucher_v2.IVoucherVoucher{
				ChainId:     59144,
				Router:      common.HexToAddress("0xabc123abc123abc123abc123abc123abc123abc1"),
				Executor:    common.HexToAddress("0xdef456def456def456def456def456def456def4"),
				Beneficiary: common.HexToAddress("0xfef456def456def456def456def456def456def5"),
				ExpireAt:    1715100785,
				Nonce:       big.NewInt(1234567890),
				Data:        []byte("paramData"),
				Signature:   []byte("paramSignature"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Encoding
			encodedData, err := Encode(tt.voucher)
			require.NoError(t, err)
			require.NotEmpty(t, encodedData, "Encode() encoded data is empty")
			require.NotEmpty(t, common.Bytes2Hex(encodedData), "Encode() encoded data is empty")

			// Test Decoding
			decodedVoucher, err := Decode(encodedData)
			require.NoError(t, err)
			require.NotNil(t, decodedVoucher, "Decode() decoded voucher is nil")
			require.True(t, reflect.DeepEqual(decodedVoucher, tt.voucher), "Decode() got = %v, want %v", decodedVoucher, tt.voucher)
		})

		t.Run(tt.name+"_Sign", func(t *testing.T) {
			privateKey, err := ecrypto.GenerateKey()
			require.NoError(t, err)

			signer := signer_pkg.NewLocalSigner(privateKey)
			encodedData, err := SignAndEncode(tt.voucher, signer)
			require.NoError(t, err)
			require.NotEmpty(t, encodedData, "Encode() encoded data is empty")
			require.NotEmpty(t, common.Bytes2Hex(encodedData), "Encode() encoded data is empty")

			// Test Decoding
			decodedVoucher, err := Decode(encodedData)
			require.NoError(t, err)
			require.NotNil(t, decodedVoucher, "Decode() decoded voucher is nil")

			signature := signer_pkg.NewSignatureFromBytes(decodedVoucher.Signature)

			tt.voucher.Signature = nil
			unsignedEncodedData, err := Encode(tt.voucher)
			require.NoError(t, err)

			signedHashedBytes := ecrypto.Keccak256Hash(unsignedEncodedData).Bytes()
			recoveredPubKey, err := signer_pkg.RecoverEthMessageSigner(signature, signedHashedBytes)
			require.NoError(t, err)
			require.Equal(t, *recoveredPubKey, privateKey.PublicKey)
		})
	}
}
