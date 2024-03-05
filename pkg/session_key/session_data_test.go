package session_key

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestPackEnableData(t *testing.T) {
	tcs := []struct {
		sessionData SessionData
		enableData  string
	}{
		{
			sessionData: SessionData{
				SessionKey: common.HexToAddress("0x4C3C9C9fE28eA197cC260491393B8f6ED48e732f"),
				ValidAfter: time.Unix(177, 0),
				ValidUntil: time.Unix(0, 0),
				MerkleRoot: hexutil.MustDecode("0x8d5b5624af55afe4c927b5139d4dbb8e72b8e4ad844f8a20745a4700a7533edf"),
				Paymaster:  common.HexToAddress("0x0000000000000000000000000000000000000001"),
				Nonce:      big.NewInt(1),
			},
			enableData: "0x4c3c9c9fe28ea197cc260491393b8f6ed48e732f8d5b5624af55afe4c927b5139d4dbb8e72b8e4ad844f8a20745a4700a7533edf0000000000b100000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001",
		},
	}

	for _, tc := range tcs {
		enableData := tc.sessionData.PackEnableData()
		assert.Equal(t, tc.enableData, hexutil.Encode(enableData))

		incompleteSignature := make([]byte, 4+6+6+20+20+32)
		incompleteSignature = append(incompleteSignature, enableData...)

		sessionData, err := UnpackEnableData(incompleteSignature)
		assert.NoError(t, err)
		assert.Equal(t, tc.sessionData, sessionData)
	}
}

func Test_getKernelSessionDataHash(t *testing.T) {
	tcs := []struct {
		sessionData   SessionData
		sig           [4]byte
		chainId       *big.Int
		kernelAddress common.Address
		validator     common.Address
		executor      common.Address
		hash          string
	}{
		{
			sessionData: SessionData{
				SessionKey: common.HexToAddress("0x4C3C9C9fE28eA197cC260491393B8f6ED48e732f"),
				ValidAfter: time.Unix(177, 0),
				ValidUntil: time.Unix(0, 0),
				MerkleRoot: hexutil.MustDecode("0x8d5b5624af55afe4c927b5139d4dbb8e72b8e4ad844f8a20745a4700a7533edf"),
				Paymaster:  common.HexToAddress("0x0000000000000000000000000000000000000001"),
				Nonce:      big.NewInt(1),
			},
			sig:           KernelExecuteSig,
			chainId:       big.NewInt(31337),
			kernelAddress: common.HexToAddress("0xBf1ca3AF628e173b067629F007c4860593779D79"),
			validator:     common.HexToAddress("0xa0Cb889707d426A7A386870A03bc70d1b0697598"),
			executor:      common.HexToAddress("0x"),
			hash:          "0x1ebf9db3933b552ad1d8f6927dccdb6d0f7cd61a89affb0de0144f125f796dea",
		},
	}

	for _, tc := range tcs {
		hash := getKernelSessionDataHash(
			tc.sessionData,
			tc.sig,
			tc.chainId,
			tc.kernelAddress,
			tc.validator,
			tc.executor,
		)

		assert.Equal(t, tc.hash, hexutil.Encode(hash))
	}
}

func Test_buildKernelDomainSeparator(t *testing.T) {
	tcs := []struct {
		chainId       *big.Int
		kernelAddress common.Address
		hash          string
	}{
		{
			chainId:       big.NewInt(31337),
			kernelAddress: common.HexToAddress("0xBf1ca3AF628e173b067629F007c4860593779D79"),
			hash:          "0xff233fe31a7c621c000cd12803c14902809135522ffe1d656ef68a329e21c6b6",
		},
	}

	for _, tc := range tcs {
		hash := getKernelDomainSeparator(tc.chainId, tc.kernelAddress)
		assert.Equal(t, tc.hash, hexutil.Encode(hash))
	}
}

func Test_buildEnableDataHash(t *testing.T) {
	tcs := []struct {
		enableData []byte
		sig        [4]byte
		validator  common.Address
		executor   common.Address
		hash       string
	}{
		{
			sig:        [4]byte{0x51, 0x94, 0x54, 0x47},
			validator:  common.HexToAddress("0xa0Cb889707d426A7A386870A03bc70d1b0697598"),
			executor:   common.HexToAddress("0x"),
			enableData: hexutil.MustDecode("0x4c3c9c9fe28ea197cc260491393b8f6ed48e732f8d5b5624af55afe4c927b5139d4dbb8e72b8e4ad844f8a20745a4700a7533edf0000000000b100000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001"),
			hash:       "0x33272c8ad9202d555926d9517120c63eab21ec7969958e787a34cfde1bb9d776",
		},
	}

	for _, tc := range tcs {
		hash := getEnableDataHash(tc.enableData, tc.sig, tc.validator, tc.executor)
		assert.Equal(t, tc.hash, hexutil.Encode(hash))
	}
}
