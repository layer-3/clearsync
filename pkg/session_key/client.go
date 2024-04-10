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
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
	mt "github.com/layer-3/go-merkletree"

	"github.com/layer-3/clearsync/pkg/artifacts/session_key_validator_v2_4"
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
	permissions                []Permission   // for now, permissions are the same for all session keys
	permTree                   *mt.MerkleTree // root (hash) of the permission tree over all parameters as leaves
}

func NewClient(config Config) (Client, error) {
	provider, err := NewEthClient(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
	}

	chainId, err := provider.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	executionSig := KernelExecuteSig
	if config.ExecuteInBatch {
		executionSig = KernelExecuteBatchSig
	}

	permTree, err := NewPermissionTree(config.Permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to create permission tree: %w", err)
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
		permissions:                config.Permissions,
		permTree:                   permTree,
	}, nil
}

func (b *backend) GetEnableDataDigest(kernelAddress, sessionKey common.Address) ([]byte, error) {
	sessionData, err := b.getSessionData(kernelAddress, sessionKey)
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

		sessionData, err := b.getSessionData(op.Sender, sessionSigner.CommonAddress())
		if err != nil {
			return nil, err
		}

		enableData := sessionData.PackEnableData()
		fullSig := PackEnableValidatorSignature(enableData, b.sessionKeyValidatorAddress, b.executorAddress, enableSig)

		userOpHash, err := op.UserOpHash(entryPoint, chainID)
		if err != nil {
			return nil, err
		}

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

		userOpHash, err := op.UserOpHash(entryPoint, chainID)
		if err != nil {
			return nil, err
		}

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

func (b *backend) getSessionData(smartWallet, sessionKey common.Address) (SessionData, error) {
	sessionKeyValidator, err := session_key_validator_v2_4.NewSessionKeyValidator(b.sessionKeyValidatorAddress, b.provider)
	if err != nil {
		return SessionData{}, fmt.Errorf("failed to connect to session key validator: %w", err)
	}

	nonces, err := sessionKeyValidator.Nonces(nil, smartWallet)
	if err != nil {
		return SessionData{}, fmt.Errorf("failed to get nonces: %w", err)
	}

	sessionData := SessionData{
		SessionKey: sessionKey,
		ValidAfter: time.Unix(int64(b.sessionKeyValidAfter), 0),
		ValidUntil: time.Unix(int64(b.sessionKeyValidUntil), 0),
		MerkleRoot: b.permTree.Root,
		Paymaster:  b.paymasterAddress,
		Nonce:      big.NewInt(0).Add(nonces.LastNonce, big.NewInt(1)),
	}

	return sessionData, nil
}

func (b *backend) getUseSessionKeySig(sessionSigner signer.Signer, userOpCallData []byte, userOpHash common.Hash) ([]byte, error) {
	calls, err := smart_wallet.UnpackCallsForKernel(userOpCallData)
	if err != nil {
		return nil, err
	}

	if len(calls) == 0 {
		return nil, fmt.Errorf("no calls found in user operation")
	}

	kernelPermissions, err := b.filterPermissions(calls)
	if err != nil {
		return nil, err
	}

	proofs, err := b.getProofs(kernelPermissions)
	if err != nil {
		return nil, err
	}

	signature, err := signer.SignEthMessage(sessionSigner, userOpHash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to sign user operation with session key: %w", err)
	}

	fullSignature, err := PackUseSessionKeySignature(sessionSigner.CommonAddress(), signature, kernelPermissions, proofs)
	if err != nil {
		return nil, err
	}

	return fullSignature, nil
}

func (b *backend) filterPermissions(calls smart_wallet.Calls) ([]kernelPermission, error) {
	permissions := make([]kernelPermission, len(calls))
	for i, call := range calls {
		permissionFound := false

		for index, perm := range b.permissions {
			if perm.Target != call.To && perm.Target != (common.Address{}) {
				continue
			}

			if call.Value.Cmp(perm.ValueLimit) > 0 {
				continue
			}

			// TODO: add native send support
			if !bytes.HasPrefix(call.CallData, perm.FunctionABI.ID) {
				continue
			}

			// TODO: check param rules

			permissions[i] = perm.toKernelPermission(uint32(index))
			permissionFound = true
		}

		if !permissionFound {
			return nil, fmt.Errorf("no permission found for call: %s %s", call.To.String(), hexutil.Encode(call.CallData[:4]))
		}
	}

	return permissions, nil
}

func (b *backend) getProofs(kernelPermissions []kernelPermission) ([]mt.Proof, error) {
	proofs := make([]mt.Proof, len(kernelPermissions))
	for i, perm := range kernelPermissions {
		proof, err := b.permTree.Proof(perm)
		if err != nil {
			return nil, fmt.Errorf("failed to get proof for permission: %w", err)
		}
		proofs[i] = *proof
	}

	return proofs, nil
}
