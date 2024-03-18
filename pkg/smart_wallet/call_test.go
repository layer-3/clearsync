package smart_wallet

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func randomAddress() common.Address {
	return common.BigToAddress(big.NewInt(rand.Int63()))
}

func TestPackUnpackCallsForKernel(t *testing.T) {
	tcs := []struct {
		calls Calls
	}{
		{
			// empty calls
			calls: Calls{},
		},
		{ // single call
			calls: Calls{
				{
					To:       randomAddress(),
					Value:    big.NewInt(0),
					CallData: []byte{1, 2, 3},
				},
			},
		},
		{ // single call without value
			calls: Calls{
				{
					To:       randomAddress(),
					CallData: []byte{1, 2, 3},
				},
			},
		},
		{ // single call with empty calldata
			calls: Calls{
				{
					To:    randomAddress(),
					Value: big.NewInt(0),
				},
			},
		},
		{ // multiple calls
			calls: Calls{
				{
					To:       randomAddress(),
					Value:    big.NewInt(0),
					CallData: []byte{1, 2, 3},
				}, {
					To:       randomAddress(),
					Value:    big.NewInt(1),
					CallData: []byte{41, 42, 43},
				},
			},
		},
	}

	for _, tc := range tcs {
		calldata, err := tc.calls.PackForKernel()
		assert.NoError(t, err)

		unpackedCalls, err := UnpackCallsForKernel(calldata)
		assert.NoError(t, err)

		if len(tc.calls) == 0 {
			assert.Empty(t, unpackedCalls)
			continue
		}

		for i := range tc.calls {
			expValue := big.NewInt(0)
			if tc.calls[i].Value != nil {
				expValue.Set(tc.calls[i].Value)
			}

			expCallData := make([]byte, len(tc.calls[i].CallData))
			copy(expCallData, tc.calls[i].CallData)

			assert.Equal(t, tc.calls[i].To, unpackedCalls[i].To)
			assert.Zero(t, expValue.Cmp(unpackedCalls[i].Value))
			assert.Equal(t, expCallData, unpackedCalls[i].CallData)
		}
	}
}

func TestErrorUnpackCallsForKernel(t *testing.T) {
	_, err := UnpackCallsForKernel([]byte{}) // less than 4 bytes
	assert.EqualError(t, err, "invalid data length")
}

func TestPackCallsForSimpleAccount(t *testing.T) {
	tcs := []struct {
		calls Calls
	}{
		{
			calls: Calls{
				{
					To:       randomAddress(),
					Value:    big.NewInt(0),
					CallData: []byte{1, 2, 3},
				},
			},
		},
	}

	for _, tc := range tcs {
		calldata, err := tc.calls.PackForSimpleAccount()
		assert.NoError(t, err)

		unpackedCalls, err := UnpackCallsForSimpleAccount(calldata)
		assert.NoError(t, err)

		for i := range tc.calls {
			assert.Equal(t, tc.calls[i].To, unpackedCalls[i].To)
			assert.Zero(t, tc.calls[i].Value.Cmp(unpackedCalls[i].Value))
			assert.Equal(t, tc.calls[i].CallData, unpackedCalls[i].CallData)
		}
	}
}
