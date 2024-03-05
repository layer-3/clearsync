package session_key

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	mt "github.com/layer-3/go-merkletree"
)

func NewPermissionTree(permissions []Permission) (*mt.MerkleTree, error) {
	contents := make([]mt.DataBlock, len(permissions))
	for i, permission := range permissions {
		contents[i] = permission.toKernelPermission(uint32(i))
	}

	hashFunc := func(data []byte) ([]byte, error) {
		return crypto.Keccak256(data), nil
	}

	tree, err := mt.New(&mt.Config{
		HashFunc:         hashFunc,
		Mode:             mt.ModeTreeBuild,
		SortSiblingPairs: true,
	}, contents)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

type Permission struct {
	Target      common.Address `json:"target"`
	FunctionABI abi.Method     `json:"functionABI"`
	ValueLimit  *big.Int       `json:"valueLimit"`
	Rules       []ParamRule    `json:"rules"`
}

func (p Permission) toKernelPermission(index uint32) kernelPermission {
	rules := make([]kernelParamRule, len(p.Rules))
	offset := 0
	for i, rule := range p.Rules {
		rules[i] = kernelParamRule{
			Offset:    big.NewInt(int64(offset)),
			Condition: rule.Condition,
			Param:     rule.Param,
		}

		offset += p.FunctionABI.Inputs[i].Type.Size
	}

	return kernelPermission{
		Index:      index,
		Target:     p.Target,
		Sig:        [4]byte(p.FunctionABI.ID),
		ValueLimit: p.ValueLimit,
		ExecutionRule: kernelExecutionRule{
			ValidAfter: big.NewInt(0),
			Interval:   big.NewInt(0),
			Runs:       big.NewInt(0),
		},
		Rules: rules,
	}
}

type ParamRule struct {
	Condition ParamCondition `json:"condition"`
	Param     [32]byte       `json:"param"`
}

type ParamCondition uint8

const (
	EqualParamCondition ParamCondition = iota
	GreaterThanParamCondition
	LessThanParamCondition
	GreaterEqualParamCondition
	LessEqualParamCondition
	NotEqualParamCondition
)
