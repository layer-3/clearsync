package session_key

import (
	"bytes"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/userop"
	mt "github.com/layer-3/go-merkletree"

	"github.com/layer-3/clearsync/pkg/artifacts/session_key_validator"
)

type Client interface {
	// GetEnableDataDigest returns the hash of the enable session data, which is used to validate the session key.
	//
	// Parameters:
	// - kernelAddress: the address of the kernel contract
	// - sessionData: the session data
	// - sig: the selector of the kernel function that will be called
	// - chainId: the chain id
	//
	// Returns:
	// - the hash of the enable session data
	GetEnableDataDigest(kernelAddress common.Address, sessionData SessionData, sig [4]byte, chainId *big.Int) []byte

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
	GetUserOpSigner(sessionSigner signer.Signer) userop.Signer
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

func (b *backend) GetEnableDataDigest(kernelAddress common.Address, sessionData SessionData, sig [4]byte, chainId *big.Int) []byte {
	return getKernelSessionDataHash(sessionData, sig, chainId, kernelAddress, b.sessionKeyValidatorAddress, b.executorAddress)
}

func (b *backend) GetIncompleteEnablingUserOpSigner(sessionSigner signer.Signer) (userop.Signer, error) {
	sessionKeyValidator, err := session_key_validator.NewSessionKeyValidator(b.sessionKeyValidatorAddress, b.provider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to session key validator: %w", err)
	}

	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing enable session key + user operation with session key")

		nonces, err := sessionKeyValidator.Nonces(nil, sessionSigner.CommonAddress())
		if err != nil {
			return nil, fmt.Errorf("failed to get nonces: %w", err)
		}

		sessionData := SessionData{
			SessionKey: sessionSigner.CommonAddress(),
			ValidAfter: time.Unix(int64(b.sessionKeyValidAfter), 0),
			ValidUntil: time.Unix(int64(b.sessionKeyValidUntil), 0),
			MerkleRoot: b.PermTree.Root,
			Paymaster:  b.paymasterAddress,
			Nonce:      big.NewInt(0).Add(nonces.LastNonce, big.NewInt(1)),
		}

		enableData := sessionData.PackEnableData()
		emptySig := signer.NewSignatureFromBytes(make([]byte, 65)) // placeholder for enableSig

		fullSig := PackEnableValidatorSignature(enableData, b.sessionKeyValidatorAddress, b.executorAddress, emptySig)

		userOpHash := op.UserOpHash(entryPoint, chainID)
		useSessionKeySig, err := b.getUseSessionKeySig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, fmt.Errorf("failed to build validation sig: %w", err)
		}

		fullSig = append(fullSig, useSessionKeySig...)

		slog.Debug("signed enable session key + user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSig),
			"hash", userOpHash.String())
		return fullSig, nil
	}, nil
}

func (b *backend) GetUserOpSigner(sessionSigner signer.Signer) userop.Signer {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation with session key")

		userOpHash := op.UserOpHash(entryPoint, chainID)
		useSessionKeySig, err := b.getUseSessionKeySig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, fmt.Errorf("failed to build use session key sig: %w", err)
		}

		fullSig := make([]byte, 0, 4+len(useSessionKeySig))
		// "use given validator" (0x00000001) mode
		// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
		fullSig = append(fullSig, []byte{0x00, 0x00, 0x00, 0x01}...)
		fullSig = append(fullSig, useSessionKeySig...)

		slog.Debug("signed user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSig),
			"hash", userOpHash.String())
		return fullSig, nil
	}
}

func (b *backend) getUseSessionKeySig(sessionSigner signer.Signer, userOpCallData []byte, userOpHash common.Hash) ([]byte, error) {
	calls, err := userop.UnpackCallsForKernel(userOpCallData)
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
		return nil, fmt.Errorf("failed to pack use session key signature: %w", err)
	}

	return fullSignature, nil
}
