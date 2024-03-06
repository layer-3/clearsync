package session_key

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
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
	// - sessionKey: the address of session key
	//
	// Returns:
	// - the hash of the enable session data
	GetEnableDataDigest(kernelAddress, sessionKey common.Address) ([]byte, error)

	// GetEnablingUserOpSigner returns a user operation signer that signs the user operation
	// with the session key and the enable signature.
	//
	// Parameters:
	// - sessionSigner: the session key signer
	// - enableSig: the signature of the enable session data
	//
	// Returns:
	// - a user operation signer function
	// - an error if the signer could not be created
	GetEnablingUserOpSigner(sessionSigner signer.Signer, enableSig signer.Signature) userop.Signer

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
	provider                   *ethclient.Client
	chainId                    *big.Int
	executionSig               [4]byte
	sessionKeyValidAfter       uint64
	sessionKeyValidUntil       uint64
	sessionKeyValidatorAddress common.Address
	executorAddress            common.Address
	paymasterAddress           common.Address
	Permissions                []Permission   // for now, permissions are the same for all session keys
	PermTree                   *mt.MerkleTree // root (hash) of the permission tree over all parameters as leaves
}

func NewClient(config Config) (Client, error) {
	provider, err := NewEthClient(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
	}

	permTree, err := NewPermissionTree(config.Permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to create permission tree: %w", err)
	}

	executionSig := KernelExecuteSig
	if config.ExecuteInBatch {
		executionSig = KernelExecuteBatchSig
	}

	chainId, err := provider.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return &backend{
		provider:                   provider,
		chainId:                    chainId,
		executionSig:               executionSig,
		sessionKeyValidAfter:       config.SessionKeyValidAfter,
		sessionKeyValidUntil:       config.SessionKeyValidUntil,
		sessionKeyValidatorAddress: config.SessionKeyValidatorAddress,
		executorAddress:            config.ExecutorAddress,
		paymasterAddress:           config.PaymasterAddress,
		Permissions:                config.Permissions,
		PermTree:                   permTree,
	}, nil
}

func (b *backend) GetEnableDataDigest(kernelAddress, sessionKey common.Address) ([]byte, error) {
	sessionData, err := b.getSessionData(sessionKey)
	if err != nil {
		return nil, err
	}

	return getKernelSessionDataHash(
		sessionData,
		b.executionSig,
		b.chainId,
		kernelAddress,
		b.sessionKeyValidatorAddress,
		b.executorAddress,
	), nil
}

func (b *backend) GetEnablingUserOpSigner(sessionSigner signer.Signer, enableSig signer.Signature) userop.Signer {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing enable session key + user operation with session key")

		sessionData, err := b.getSessionData(sessionSigner.CommonAddress())
		if err != nil {
			return nil, err
		}

		enableData := sessionData.PackEnableData()
		fullSig := PackEnableValidatorSignature(enableData, b.sessionKeyValidatorAddress, b.executorAddress, enableSig)

		userOpHash := op.UserOpHash(entryPoint, chainID)
		useSessionKeySig, err := b.getUseSessionKeySig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, err
		}

		fullSig = append(fullSig, useSessionKeySig...)

		slog.Debug("signed enable session key + user operation with session key for Kernel",
			"signature", hexutil.Encode(fullSig),
			"hash", userOpHash.String())
		return fullSig, nil
	}
}

func (b *backend) GetUserOpSigner(sessionSigner signer.Signer) userop.Signer {
	return func(op userop.UserOperation, entryPoint common.Address, chainID *big.Int) ([]byte, error) {
		slog.Debug("signing user operation with session key")

		userOpHash := op.UserOpHash(entryPoint, chainID)
		useSessionKeySig, err := b.getUseSessionKeySig(sessionSigner, op.CallData, userOpHash)
		if err != nil {
			return nil, err
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

func (b *backend) getSessionData(sessionKey common.Address) (SessionData, error) {
	sessionKeyValidator, err := session_key_validator.NewSessionKeyValidator(b.sessionKeyValidatorAddress, b.provider)
	if err != nil {
		return SessionData{}, fmt.Errorf("failed to connect to session key validator: %w", err)
	}

	nonces, err := sessionKeyValidator.Nonces(nil, sessionKey)
	if err != nil {
		return SessionData{}, fmt.Errorf("failed to get nonces: %w", err)
	}

	sessionData := SessionData{
		SessionKey: sessionKey,
		ValidAfter: time.Unix(int64(b.sessionKeyValidAfter), 0),
		ValidUntil: time.Unix(int64(b.sessionKeyValidUntil), 0),
		MerkleRoot: b.PermTree.Root,
		Paymaster:  b.paymasterAddress,
		Nonce:      big.NewInt(0).Add(nonces.LastNonce, big.NewInt(1)),
	}

	return sessionData, nil
}

func (b *backend) getUseSessionKeySig(sessionSigner signer.Signer, userOpCallData []byte, userOpHash common.Hash) ([]byte, error) {
	calls, err := userop.UnpackCallsForKernel(userOpCallData)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return fullSignature, nil
}
