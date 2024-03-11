package userop

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/signer"
)

func SignerForBiconomy(ecdsaSigner signer.Signer) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation")

		hash := op.UserOpHash(entryPoint, chainID)
		signature, err := signer.SignEthMessage(ecdsaSigner, hash.Bytes())
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}

		// The address is correct for all chains as for Biconomy v2.0
		// see https://docs.biconomy.io/contracts
		ecdsaOwnershipValidationModuleAddress := common.HexToAddress("0x0000001c5b32F37F5beA87BDD5374eB2aC54eA8e")
		args := abi.Arguments{
			{Type: bytes},
			{Type: address},
		}
		// Pack the signature and the ecdsaOwnershipValidationModuleAddress
		// to be used as the signature for the user operation
		// See more: https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L337
		fullSignature, err := args.Pack(signature, ecdsaOwnershipValidationModuleAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack signature: %w", err)
		}

		slog.Debug("signed user operation for Biconomy",
			"signature", hexutil.Encode(fullSignature),
			"hash", hash.String())
		return fullSignature, nil
	}
}

func SignerForKernel(ecdsaSigner signer.Signer) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		// FIXME: not with session key
		slog.Debug("signing user operation with session key")

		hash := op.UserOpHash(entryPoint, chainID)
		signature, err := signer.SignEthMessage(ecdsaSigner, hash.Bytes())
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
