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
			root:       "0xbc1b64afc7aed802815dcf65a6fe5b91fec933352d4392c741142873627c6dcc",
			proofIndex: 0,
			proof: [][32]byte{
				[32]byte(hexutil.MustDecode("0x7b33528d92deb6fb4f510a22d8de3f1eddff8d15be3cd7bcc0bfc4b907d1d1da")),
				[32]byte(hexutil.MustDecode("0xb42da658b0d0fa36b677dc6053edbe05f471a958bc8a053505ad3321d2df8ef0")),
			},
		},
	}

	for _, tc := range tcs {
		pt, err := NewPermissionTree(tc.permissions)
		require.NoError(t, err)

		assert.NoError(t, err)
		assert.Equal(t, tc.root, hexutil.Encode(pt.Tree.Root))

		proof, err := pt.Tree.Proof(tc.permissions[tc.proofIndex].toABI(uint32(tc.proofIndex)))
		assert.NoError(t, err)
		assert.Equal(t, len(tc.proof), len(proof.Siblings))
		for i, sibling := range proof.Siblings {
			assert.Equal(t, hexutil.Encode(tc.proof[i][:]), hexutil.Encode(sibling))
		}
	}
}
