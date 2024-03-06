package userop

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackCallsForKernel(t *testing.T) {
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
		calldata, err := tc.calls.PackForKernel()
		assert.NoError(t, err)

		unpackedCalls, err := UnpackCallsForKernel(calldata)
		assert.NoError(t, err)

		for i := range tc.calls {
			assert.Equal(t, tc.calls[i].To, unpackedCalls[i].To)
			assert.Zero(t, tc.calls[i].Value.Cmp(unpackedCalls[i].Value))
			assert.Equal(t, tc.calls[i].CallData, unpackedCalls[i].CallData)
		}
	}
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
