package session_key

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/abi/itoken"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/stretchr/testify/require"
)

func TestGetUseSessionKeySig(t *testing.T) {
	be := backend{}

	t.Run("Error when no calls", func(t *testing.T) {
		noCallsCallData, err := userop.Calls{}.PackForKernel()
		require.NoError(t, err)

		_, err = be.getUseSessionKeySig(signer_pkg.LocalSigner{}, noCallsCallData, common.Hash{})
		require.EqualError(t, err, "no calls found in user operation")
	})
}

func TestFilterPermissions(t *testing.T) {
	ierc20ABI, err := itoken.IERC20MetaData.GetAbi()
	require.NoError(t, err)

	permissions := []Permission{
		{
			Target:      common.HexToAddress("0x1"),
			FunctionABI: ierc20ABI.Methods["approve"],
			ValueLimit:  big.NewInt(0),
			Rules: []ParamRule{
				{
					Condition: EqualParamCondition,
					Param:     [32]byte(append(make([]byte, 12), common.HexToAddress("0x11").Bytes()...)),
				},
				{
					Condition: GreaterThanParamCondition,
					Param:     [32]byte(big.NewInt(10).FillBytes(make([]byte, 32))),
				},
			},
		},
		{
			Target:      common.Address{},
			FunctionABI: ierc20ABI.Methods["transferFrom"],
			ValueLimit:  big.NewInt(0),
			Rules: []ParamRule{
				{
					Condition: EqualParamCondition,
					Param:     [32]byte(append(make([]byte, 12), common.HexToAddress("0x21").Bytes()...)),
				},
			},
		},
		{
			Target:      common.HexToAddress("0x3"),
			FunctionABI: ierc20ABI.Methods["transfer"],
			ValueLimit:  big.NewInt(0),
			Rules:       []ParamRule{},
		},
		// uncomment when native send support is added
		// {
		// 	Target:     common.HexToAddress("0xdcbee058bCd0723559DA80000cb791a1Ee1023e0"),
		// 	ValueLimit: big.NewInt(10),
		// },
	}

	be := backend{
		permissions: permissions,
	}

	approveCD, err := ierc20ABI.Pack("approve", common.HexToAddress("0x11"), big.NewInt(15))
	require.NoError(t, err)
	transferFromCD, err := ierc20ABI.Pack("transferFrom", common.HexToAddress("0x21"), common.HexToAddress("0x42"), big.NewInt(42))
	require.NoError(t, err)
	transferCD, err := ierc20ABI.Pack("transfer", common.HexToAddress("0x3"), big.NewInt(42))
	require.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		tcs := []struct {
			permissions []Permission
			calls       userop.Calls
			expected    []kernelPermission
		}{
			{ // single call
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0x1"),
						CallData: approveCD,
						Value:    big.NewInt(0),
					},
				},
				expected: []kernelPermission{permissions[0].toKernelPermission(0)},
			},
			{ // single call with zero address permission
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0x42"),
						CallData: transferFromCD,
						Value:    big.NewInt(0),
					},
				},
				expected: []kernelPermission{permissions[1].toKernelPermission(1)},
			},
			// uncomment when native send support is added
			// { // native transfer call
			// 	permissions: permissions,
			// 	calls: userop.Calls{
			// 		{
			// 			To:       common.HexToAddress("0xdcbee058bCd0723559DA80000cb791a1Ee1023e0"),
			// 			Value:    big.NewInt(1),
			// 		},
			// 	},
			// },
			{ // multiple calls
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0x1"),
						CallData: approveCD,
						Value:    big.NewInt(0),
					},
					{
						To:       common.HexToAddress("0x3"),
						CallData: transferCD,
						Value:    big.NewInt(0),
					},
				},
				expected: []kernelPermission{permissions[0].toKernelPermission(0), permissions[2].toKernelPermission(2)},
			},
		}

		for _, tc := range tcs {
			kernelPermissions, err := be.filterPermissions(tc.calls)
			require.NoError(t, err)
			require.Equal(t, tc.expected, kernelPermissions)
		}
	})

	t.Run("Error", func(t *testing.T) {
		errStr := "no permission found for call:"

		tcs := []struct {
			permissions []Permission
			calls       userop.Calls
		}{
			{ // single call and it is not in permissions
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0xdeadbeef"),
						CallData: approveCD,
						Value:    big.NewInt(0),
					},
				},
			},
			// TODO: uncomment when param rule checks are added
			// { // single call, it is in permissions, but its param rules are not satisfied
			// 	permissions: permissions,
			// 	calls: userop.Calls{
			// 		{
			// 			To:       common.HexToAddress("0x03A6a84cD762D9707A21605b548aaaB891562aAb"),
			// 			CallData: approveCD,
			// 			Value:    big.NewInt(0),
			// 		},
			// 	},
			// 	expected: []kernelPermission{permissions[0].toKernelPermission(0)},
			// },
			{ // multiple calls, one of them is not in permissions
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0x1"),
						CallData: approveCD,
						Value:    big.NewInt(0),
					},
					{
						To:       common.HexToAddress("0x1"),
						CallData: hexutil.MustDecode("0xdeadbeef"),
						Value:    big.NewInt(0),
					},
				},
			},
			{ // multiple calls, all of them are not in permissions
				permissions: permissions,
				calls: userop.Calls{
					{
						To:       common.HexToAddress("0x2"),
						CallData: hexutil.MustDecode("0xdeadbeef"),
						Value:    big.NewInt(0),
					},
					{
						To:       common.HexToAddress("0x2"),
						CallData: hexutil.MustDecode("0xdeadbeef"),
						Value:    big.NewInt(0),
					},
				},
			},
		}

		for _, tc := range tcs {
			_, err := be.filterPermissions(tc.calls)
			require.ErrorContains(t, err, errStr)
		}
	})

}
