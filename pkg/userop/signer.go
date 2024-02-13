package userop

import (
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func SignerForBiconomy(privateKey *ecdsa.PrivateKey) func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Info("signing user operation")

		signature, err := op.SignWithECDSA(privateKey, entryPoint, chainID)
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}

		ecdsaOwnershipValidationModuleAddress := common.HexToAddress("0x0000001c5b32F37F5beA87BDD5374eB2aC54eA8e")
		args := abi.Arguments{
			{Type: Bytes},
			{Type: Address},
		}
		biconomySignature, err := args.Pack(signature, ecdsaOwnershipValidationModuleAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack signature: %w", err)
		}

		slog.Info("signed user operation for Biconomy", "signature", hexutil.Encode(biconomySignature))
		return biconomySignature, nil
	}
}
