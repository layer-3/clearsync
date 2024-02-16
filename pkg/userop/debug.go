//go:build dev

package userop

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (op *UserOperation) ToArray() string {
	return fmt.Sprintf(
		`["%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s"]`,
		op.Sender.String(),
		hexutil.EncodeBig(op.Nonce.BigInt()),
		hexutil.Encode(op.InitCode),
		hexutil.Encode(op.CallData),
		hexutil.EncodeBig(op.CallGasLimit.BigInt()),
		hexutil.EncodeBig(op.VerificationGasLimit.BigInt()),
		hexutil.EncodeBig(op.PreVerificationGas.BigInt()),
		hexutil.EncodeBig(op.MaxFeePerGas.BigInt()),
		hexutil.EncodeBig(op.MaxPriorityFeePerGas.BigInt()),
		hexutil.Encode(op.PaymasterAndData),
		hexutil.Encode(op.Signature),
	)
}
