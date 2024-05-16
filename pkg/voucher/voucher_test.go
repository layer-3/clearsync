package voucher

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/ivoucher"
)

func TestEncodeDecode(t *testing.T) {
	var codeHash [32]byte
	for i := 0; i < len(codeHash); i++ {
		codeHash[i] = byte(i)
	}

	tests := []struct {
		name    string
		voucher ivoucher.IVoucherVoucher
	}{
		{
			name: "Basic Voucher",
			voucher: ivoucher.IVoucherVoucher{
				Target:          common.HexToAddress("0xabc123abc123abc123abc123abc123abc123abc1"),
				Action:          3,
				Beneficiary:     common.HexToAddress("0xdef456def456def456def456def456def456def4"),
				ExpireAt:        1715100785,
				ChainId:         59144,
				VoucherCodeHash: codeHash,
				EncodedParams:   []byte("paramData"),
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
