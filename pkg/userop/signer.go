package userop

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignerForBiconomy(ecdsaSigner ECDSASigner) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation")

		hash := op.UserOpHash(entryPoint, chainID)
		signedHash, err := signWithECDSA(hash.Bytes(), ecdsaSigner)
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
		signature, err := args.Pack(signedHash, ecdsaOwnershipValidationModuleAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack signature: %w", err)
		}

		slog.Debug("signed user operation for Biconomy",
			"signature", hexutil.Encode(signature),
			"hash", hash.String())
		return signature, nil
	}
}

func SignerForKernel(ecdsaSigner ECDSASigner) Signer {
	return func(op UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation")

		hash := op.UserOpHash(entryPoint, chainID)
		signature, err := signWithECDSA(hash.Bytes(), ecdsaSigner)
		if err != nil {
			return nil, fmt.Errorf("failed to sign user operation: %w", err)
		}

		encodedSig := hexutil.Encode(signature)
		// Add 'use sudo validator' mode to signature
		// See more: https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
		modifiedSig := strings.Replace(encodedSig, "0x", "0x00000000", 1)

		signature, err = hexutil.Decode(modifiedSig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature: %w", err)
		}

		slog.Debug("signed user operation for Kernel",
			"signature", hexutil.Encode(signature),
			"hash", hash.String())
		return signature, nil
	}
}

// signWithECDSA signs the hash with the given private key using the ECDSA algorithm.
func signWithECDSA(hash []byte, ecdsaSigner ECDSASigner) ([]byte, error) {
	ethMessageHash := computeEthSignedMessageHash(hash)

	signature, err := ecdsaSigner.Sign(ethMessageHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign user operation: %w", err)
	}

	// To make the signature compatible with
	// Ethereum's ECDSA recovery, ensure V is 27 or 28.
	if signature[64] < 27 {
		signature[64] += 27
	}

	slog.Debug("user operation signed:", "hash", common.Bytes2Hex(hash), "signature", hexutil.Encode(signature))
	return signature, nil
}

// computeEthSignedMessageHash accepts an arbitrary message, prepends a known message,
// and hashes the result using keccak256. The known message added to the input before hashing is
// "\x19Ethereum Signed Message:\n" + len(message).
func computeEthSignedMessageHash(message []byte) []byte {
	return crypto.Keccak256([]byte(
		fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), string(message)),
	))
}
