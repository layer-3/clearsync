package session_key

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
)

func SignerForKernel(sessionKeySigner signer.Signer) userop.Signer {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation with session key")

		hash := op.UserOpHash(entryPoint, chainID)
		signature, err := signer.SignEthMessage(sessionKeySigner, hash.Bytes())
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}

		// Add 'use sudo validator' mode to signature
		// See more: https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
		modifiedSig := strings.Replace(signature.String(), "0x", "0x00000000", 1)

		fullSignature, err := hexutil.Decode(modifiedSig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature: %w", err)
		}

		slog.Debug("signed user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSignature),
			"hash", hash.String())
		return fullSignature, nil
	}
}
