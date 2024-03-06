package session_key

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermissionsMerkleRootAndProof(t *testing.T) {
	erc20ABI, err := itoken.IERC20MetaData.GetAbi()
	require.NoError(t, err)

	funcABI := erc20ABI.Methods["approve"]

	tcs := []struct {
		permissions []Permission
		root        string
		proofIndex  int
		proof       [][32]byte
	}{
		{
			permissions: []Permission{
				{
					Target:      common.HexToAddress("0x03A6a84cD762D9707A21605b548aaaB891562aAb"),
					FunctionABI: funcABI,
					ValueLimit:  big.NewInt(0),
					Rules: []ParamRule{
						{
							Condition: EqualParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x000000000000000000000000D6BbDE9174b1CdAa358d2Cf4D57D1a9F7178FBfF")[:32]),
						},
						{
							Condition: GreaterThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
				{
					Target:      common.HexToAddress("0xD6BbDE9174b1CdAa358d2Cf4D57D1a9F7178FBfF"),
					FunctionABI: funcABI,
					ValueLimit:  big.NewInt(0),
					Rules: []ParamRule{
						{
							Condition: NotEqualParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
						{
							Condition: LessThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
				{
					Target:      common.HexToAddress("0x15cF58144EF33af1e14b5208015d11F9143E27b9"),
					FunctionABI: funcABI,
					ValueLimit:  big.NewInt(0),
					Rules: []ParamRule{
						{
							Condition: LessThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
						{
							Condition: GreaterEqualParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
			},
			root:       "0x52cc00752ba452ce5a6e1ca7656d44d53410214dcae2bdddadb1ecbbaea2b268",
			proofIndex: 0,
			proof: [][32]byte{
				[32]byte(hexutil.MustDecode("0xe25a72ba3b51500a0512b23856f0505e284d2872249f6ccfedbc27d683d32cdc")),
				[32]byte(hexutil.MustDecode("0xcb692e6c5e6f67d6ebd8ad8788a98741752260717fc70b1237e754f6d6f71c08")),
			},
		},
	}

	for _, tc := range tcs {
		pt, err := NewPermissionTree(tc.permissions)
		require.NoError(t, err)

		assert.NoError(t, err)
		assert.Equal(t, tc.root, hexutil.Encode(pt.Root))

		proof, err := pt.Proof(tc.permissions[tc.proofIndex].toKernelPermission(uint32(tc.proofIndex)))
		assert.NoError(t, err)
		assert.Equal(t, len(tc.proof), len(proof.Siblings))
		for i, sibling := range proof.Siblings {
			assert.Equal(t, hexutil.Encode(tc.proof[i][:]), hexutil.Encode(sibling))
		}
	}
}
