package session_key

import (
	"bytes"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
	mt "github.com/layer-3/go-merkletree"
)

type Client interface {
	// GetIncompleteEnablingUserOpSigner returns a user operation signer that assembles enable session data,
	// but does not sign it. `enableSigLength` is set to 65 and `enableSig` is zeroed. This Signer also signs
	// the user operation with the session key.
	//
	// Parameters:
	// - sessionSigner: the session key signer
	//
	// Returns:
	// - a user operation signer function
	// - an error if the signer could not be created
	GetIncompleteEnablingUserOpSigner(sessionSigner signer.Signer) (userop.Signer, error)

	// GetUserOpSigner returns a user operation signer that signs the user operation with the session key.
	//
	// Parameters:
	// - sessionSigner: the session key signer
	//
	// Returns:
	// - a user operation signer function
	// - an error if the signer could not be created
	GetUserOpSigner(sessionSigner signer.Signer) (userop.Signer, error)
}

type backend struct {
	provider                   ethBackend
	sessionKeyValidAfter       uint64
	sessionKeyValidUntil       uint64
	sessionKeyValidatorAddress common.Address
	executorAddress            common.Address
	paymasterAddress           common.Address
	Permissions                []Permission   // for now, permissions are the same for all session keys
	PermTree                   *mt.MerkleTree // root (hash) of the permission tree over all parameters as leaves
}

func NewClient(config Config) (Client, error) {
	provider, err := NewEthBackend(config.ProviderUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to eth backend: %w", err)
	}

	permTree, err := NewPermissionTree(config.Permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to create permission tree: %w", err)
	}

	return &backend{
		provider:                   provider,
		sessionKeyValidAfter:       config.SessionKeyValidAfter,
		sessionKeyValidUntil:       config.SessionKeyValidUntil,
		sessionKeyValidatorAddress: config.SessionKeyValidatorAddress,
		executorAddress:            config.ExecutorAddress,
		paymasterAddress:           config.PaymasterAddress,
		Permissions:                config.Permissions,
		PermTree:                   permTree,
	}, nil
}

func (b *backend) GetIncompleteEnablingUserOpSigner(sessionSigner signer.Signer) (userop.Signer, error) {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing enable session key + user operation with session key")

		userOpHash := op.UserOpHash(entryPoint, chainID)
		validationSig, err := b.getValidationSig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, fmt.Errorf("failed to build validation sig: %w", err)
		}

		// "enable validator" mode
		// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
		validatorMode := "0x00000002"
		uint48Zero := "000000000000"

		// TODO:
		enableData := make([]byte, 32)
		// enableData length padded to 32 bytes
		enableDataLength := fmt.Sprintf("%032x", len(enableData))

		// TODO:
		enableSig := strings.Repeat("0", 65)
		// enableSig length padded to 32 bytes
		enableSigLength := fmt.Sprintf("%032x", len(enableSig))

		// validatorMode + validatorValidAfter + validatorValidUntil + validatorAddress + executorAddress (zero) + enableDataLength (padded to 32 bytes) + enableData
		// TODO: verify zero address is correct
		fullSigStr := validatorMode + uint48Zero + uint48Zero + ECDSAValidatorAddress[2:] + common.BytesToAddress(nil).String()[2:] + enableDataLength + hexutil.Encode(enableData)[2:] + enableSigLength + enableSig + hexutil.Encode(validationSig)[2:]

		fullSignature, err := hexutil.Decode(fullSigStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature: %w", err)
		}

		slog.Debug("signed enable session key + user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSignature),
			"hash", userOpHash.String())
		return fullSignature, nil
	}, nil
}

func (b *backend) GetUserOpSigner(sessionSigner signer.Signer) (userop.Signer, error) {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation with session key")

		userOpHash := op.UserOpHash(entryPoint, chainID)
		validationSig, err := b.getValidationSig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, fmt.Errorf("failed to build validation sig: %w", err)
		}

		// "use given validator" mode
		// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
		validatorMode := "0x00000001"
		fullSigStr := validatorMode + hexutil.Encode(validationSig)[2:]

		fullSignature, err := hexutil.Decode(fullSigStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signature: %w", err)
		}

		slog.Debug("signed user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSignature),
			"hash", userOpHash.String())
		return fullSignature, nil
	}, nil
}

func (b *backend) getValidationSig(sessionSigner signer.Signer, userOpCallData []byte, userOpHash common.Hash) ([]byte, error) {
	calls, err := userop.UnpackKernelCalls(userOpCallData)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack user operation call data: %w", err)
	}

	permissions := make([]Permission, len(calls))
	proofs := make([]mt.Proof, len(calls))
	for i, call := range calls {
		permissionIndex := -1
		for index, perm := range b.Permissions {
			if perm.Target != call.To && call.To != (common.Address{}) {
				continue
			}

			if call.Value.Cmp(perm.ValueLimit) > 0 {
				continue
			}

			if !bytes.HasPrefix(call.CallData, perm.FunctionABI.ID) {
				continue
			}

			permissions[i] = perm
			permissionIndex = index
		}

		if permissionIndex == -1 {
			return nil, fmt.Errorf("no permission found for call: %s %x", call.To.String(), call.CallData[:4])
		}

		proof, err := b.PermTree.Proof(permissions[i].toKernelPermission(uint32(permissionIndex)))
		if err != nil {
			return nil, fmt.Errorf("failed to get proof for permission: %w", err)
		}
		proofs[i] = *proof
	}

	signature, err := signer.SignEthMessage(sessionSigner, userOpHash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to sign user operation: %w", err)
	}

	fullSignature, err := PackUseSessionKeySignature(sessionSigner.CommonAddress(), signature, permissions, proofs)
	if err != nil {
		return nil, fmt.Errorf("failed to pack session key signature: %w", err)
	}

	return fullSignature, nil
}
