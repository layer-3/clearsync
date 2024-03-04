package session_key

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	mt "github.com/layer-3/go-merkletree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermissionHash(t *testing.T) {
	tcs := []struct {
		permission        kernelPermission
		encodedPermission string
		hash              string
	}{
		{
			permission: kernelPermission{
				Index:      0,
				Target:     common.HexToAddress("0x03A6a84cD762D9707A21605b548aaaB891562aAb"),
				Sig:        [4]byte{0x20, 0x91, 0xaf, 0x26},
				ValueLimit: big.NewInt(0),
				ExecutionRule: kernelExecutionRule{
					ValidAfter: big.NewInt(3),
					Interval:   big.NewInt(2),
					Runs:       big.NewInt(9),
				},
				Rules: []kernelParamRule{
					{
						Offset:    big.NewInt(0),
						Condition: EqualParamCondition,
						Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
					},
					{
						Offset:    big.NewInt(32),
						Condition: GreaterThanParamCondition,
						Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
					},
				},
			},
			encodedPermission: "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003a6a84cd762d9707a21605b548aaab891562aab2091af2600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000064",
			hash:              "0xe5486fabe8cd4128ee54244f940fc7477b2652ab4fdfa13e9927bf252e4d66da",
		},
		{
			permission: kernelPermission{
				Index:      185,
				Target:     common.HexToAddress("0xa0075DDDF74b904842c323A2a8189E643beFF4DA"),
				Sig:        [4]byte{0x20, 0x91, 0xaf, 0x26},
				ValueLimit: big.NewInt(0),
				ExecutionRule: kernelExecutionRule{
					ValidAfter: big.NewInt(3),
					Interval:   big.NewInt(2),
					Runs:       big.NewInt(9),
				},
				Rules: []kernelParamRule{
					{
						Offset:    big.NewInt(0),
						Condition: NotEqualParamCondition,
						Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
					},
					{
						Offset:    big.NewInt(32),
						Condition: EqualParamCondition,
						Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
					},
				},
			},
			encodedPermission: "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000b9000000000000000000000000a0075dddf74b904842c323a2a8189e643beff4da2091af2600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064",
			hash:              "0x651fe941928597ecb0406ec9678509c0913fc735ff48f24194f08ff364c92499",
		},
	}

	for _, tc := range tcs {
		encodedPermission, err := tc.permission.Encode()
		assert.NoError(t, err)
		assert.Equal(t, tc.encodedPermission, hexutil.Encode(encodedPermission))

		hash := crypto.Keccak256(encodedPermission)
		assert.NoError(t, err)
		assert.Equal(t, tc.hash, hexutil.Encode(hash))
	}
}

func Test_abiPermissionsMerkleRootAndProof(t *testing.T) {
	tcs := []struct {
		permissions []kernelPermission
		root        string
		proofIndex  int
		proof       [][32]byte
	}{
		{
			permissions: []kernelPermission{
				{
					Index:      0,
					Target:     common.HexToAddress("0x03A6a84cD762D9707A21605b548aaaB891562aAb"),
					Sig:        [4]byte{0x20, 0x91, 0xaf, 0x26},
					ValueLimit: big.NewInt(0),
					ExecutionRule: kernelExecutionRule{
						ValidAfter: big.NewInt(177),
						Interval:   big.NewInt(5850),
						Runs:       big.NewInt(3),
					},
					Rules: []kernelParamRule{
						{
							Offset:    big.NewInt(0),
							Condition: EqualParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
						{
							Offset:    big.NewInt(32),
							Condition: GreaterThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
				{
					Index:      1,
					Target:     common.HexToAddress("0xD6BbDE9174b1CdAa358d2Cf4D57D1a9F7178FBfF"),
					Sig:        [4]byte{0x20, 0x91, 0xaf, 0x26},
					ValueLimit: big.NewInt(0),
					ExecutionRule: kernelExecutionRule{
						ValidAfter: big.NewInt(177),
						Interval:   big.NewInt(5850),
						Runs:       big.NewInt(3),
					},
					Rules: []kernelParamRule{
						{
							Offset:    big.NewInt(0),
							Condition: GreaterThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
						{
							Offset:    big.NewInt(32),
							Condition: LessThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
				{
					Index:      2,
					Target:     common.HexToAddress("0x15cF58144EF33af1e14b5208015d11F9143E27b9"),
					Sig:        [4]byte{0x20, 0x91, 0xaf, 0x26},
					ValueLimit: big.NewInt(0),
					ExecutionRule: kernelExecutionRule{
						ValidAfter: big.NewInt(177),
						Interval:   big.NewInt(5850),
						Runs:       big.NewInt(3),
					},
					Rules: []kernelParamRule{
						{
							Offset:    big.NewInt(0),
							Condition: LessThanParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
						{
							Offset:    big.NewInt(32),
							Condition: GreaterEqualParamCondition,
							Param:     [32]byte(hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000064")[:32]),
						},
					},
				},
			},
			root:       "0x8d5b5624af55afe4c927b5139d4dbb8e72b8e4ad844f8a20745a4700a7533edf",
			proofIndex: 0,
			proof: [][32]byte{
				[32]byte(hexutil.MustDecode("0x3e1b0fd674a588568c3ca9ffcafc2fd125cc6e2b6b2b133977c02047d262b690")),
				[32]byte(hexutil.MustDecode("0x3001620487f821a0b18b4a3db22bea23f12abf535fac8e90064127ff10b9dbbc")),
			},
		},
	}

	for _, tc := range tcs {
		contents := make([]mt.DataBlock, len(tc.permissions))
		for i, permission := range tc.permissions {
			contents[i] = permission
		}

		hashFunc := func(data []byte) ([]byte, error) {
			return crypto.Keccak256(data), nil
		}

		tree, err := mt.New(&mt.Config{
			HashFunc:         hashFunc,
			Mode:             mt.ModeTreeBuild,
			SortSiblingPairs: true,
		}, contents)
		require.NoError(t, err)

		assert.NoError(t, err)
		assert.Equal(t, tc.root, hexutil.Encode(tree.Root))

		proof, err := tree.Proof(tc.permissions[tc.proofIndex])
		assert.NoError(t, err)
		assert.Equal(t, len(tc.proof), len(proof.Siblings))
		for i, sibling := range proof.Siblings {
			assert.Equal(t, hexutil.Encode(tc.proof[i][:]), hexutil.Encode(sibling))
		}
	}
}
