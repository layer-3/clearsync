package userop

import (
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func SignerForBiconomy(privateKey *ecdsa.PrivateKey) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Info("signing user operation")

		hash := op.UserOpHash(entryPoint, chainID)
		signedHash, err := op.SignWithECDSA(hash.Bytes(), privateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}

		ecdsaOwnershipValidationModuleAddress := common.HexToAddress("0x0000001c5b32F37F5beA87BDD5374eB2aC54eA8e")
		args := abi.Arguments{
			{Type: Bytes},
			{Type: Address},
		}
		signature, err := args.Pack(signedHash, ecdsaOwnershipValidationModuleAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack signature: %w", err)
		}

		slog.Info("signed user operation for Biconomy", "signature", hexutil.Encode(signature))
		return signature, nil
	}
}

func SignerForKernel(privateKey *ecdsa.PrivateKey) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Info("signing user operation")

		hash := op.UserOpHash(entryPoint, chainID)
		signature, err := op.SignWithECDSA(hash.Bytes(), privateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}
		fmt.Println("signed hash", signature)

		encodedSig := hexutil.Encode(signature)
		modifiedSig := strings.Replace(encodedSig, "0x", "0x00000000", 1)

		signature, err = hexutil.Decode(modifiedSig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature: %w", err)
		}

		slog.Info("signed user operation for Biconomy", "signature", hexutil.Encode(signature))
		return signature, nil
	}
}
