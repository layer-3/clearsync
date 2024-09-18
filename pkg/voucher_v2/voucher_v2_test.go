package voucher_v2

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/clearsync/pkg/abi/ivoucher_v2"
	"github.com/layer-3/clearsync/pkg/signer"
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
	}
}

func TestVoucherSigner(t *testing.T) {
	signerPK, err := ecrypto.HexToECDSA("02016836A56B71F0D02689E69E326F4F4C1B9057164EF592671CF0D37C8040C0")
	require.NoError(t, err)

	signer := signer.NewLocalSigner(signerPK)

	voucher := ivoucher_v2.IVoucherVoucher{
		ChainId:     42,
		Router:      common.HexToAddress("0xb8Fff74a193B180096e3908F1c4D40c44b6Fefdd"),
		Executor:    common.HexToAddress("0x0"),
		Beneficiary: common.HexToAddress("0x0"),
		ExpireAt:    0,
		Nonce:       big.NewInt(0),
		Data:        []byte(""),
	}

	vs := NewVoucherSigner().WithSigner(signer).AddVoucher(voucher)
	err = vs.Sign()
	require.NoError(t, err)

	sv, err := vs.First()
	require.NoError(t, err)

	require.Equal(t, "0xf7e2f5d6db05c6aed92d79e1a47e9c14dccd7853129832cfb4da986d077b195941da674d8551890e36c20e9234145e79c1a6c6913c89e63fb31b4ce981cd20b81b", hexutil.Encode(sv.Signature))
}
