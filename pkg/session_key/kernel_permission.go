package session_key

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	permissionABI = "[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"valueLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"enum ParamCondition\",\"name\":\"condition\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"param\",\"type\":\"bytes32\"}],\"internalType\":\"struct ParamRule[]\",\"name\":\"rules\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"ValidAfter\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"interval\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"runs\",\"type\":\"uint48\"}],\"internalType\":\"struct ExecutionRule\",\"name\":\"executionRule\",\"type\":\"tuple\"}],\"internalType\":\"struct Permission\",\"name\":\"permission\",\"type\":\"tuple\"}]"
)

type kernelPermission struct {
	Index         uint32              `json:"index"` // ? -> used in permissionKey to track executions
	Target        common.Address      `json:"target"`
	Sig           [4]byte             `json:"sig"` // 4 bytes of function signature
	ValueLimit    *big.Int            `json:"valueLimit"`
	ExecutionRule kernelExecutionRule `json:"executionRule"`
	Rules         []kernelParamRule   `json:"rules"` // if empty - no rules
}

func (p kernelPermission) Encode() ([]byte, error) {
	var args abi.Arguments
	dec := json.NewDecoder(strings.NewReader(permissionABI))
	if err := dec.Decode(&args); err != nil {
		return nil, err
	}

	encodedPermission, err := args.Pack(p)
	if err != nil {
		return nil, err
	}

	return encodedPermission, nil
}

func (p kernelPermission) Serialize() ([]byte, error) {
	return p.Encode()
}

type kernelParamRule struct {
	Offset    *big.Int       `json:"offset"`
	Condition ParamCondition `json:"condition"`
	Param     [32]byte       `json:"param"`
}

type kernelExecutionRule struct {
	ValidAfter *big.Int `json:"validAfter"`
	Interval   *big.Int `json:"interval"` // if zero - unlimited
	Runs       *big.Int `json:"runs"`     // if zero - unlimited
}
