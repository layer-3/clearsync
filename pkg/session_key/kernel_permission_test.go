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
			encodedPermission: "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003a6a84cd762d9707a21605b548aaab891562aab2091af26000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000064",
			hash:              "0x226bf7aa427c6a527d702ffb392572d0472acff05af6f82a4fce3aa489867d04",
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
			encodedPermission: "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000b9000000000000000000000000a0075dddf74b904842c323a2a8189e643beff4da2091af26000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064",
			hash:              "0xa57df7451d868cb2bc41b57b4ab8ce08124608658b9653c7fb7944c7a1b202fb",
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
			root:       "0x6ced4afaeb72d8244957527b852461f761fc1bb79c5480b72db742158241ba50",
			proofIndex: 0,
			proof: [][32]byte{
				[32]byte(hexutil.MustDecode("0xfad01e583666b3c252233d7ada3bd2d8cd5cb3ab5c43184748dbd3d8ef5478b4")),
				[32]byte(hexutil.MustDecode("0x8615504ab603e1e23d6804c93bb06e477dbcecf1851549e5547070aaa84a6526")),
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
