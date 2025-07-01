package signer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const (
	ERC1271MagicValue = "0x1626ba7e"
)

func SignerFnFactory(signer Signer, chainID *big.Int) func(common.Address, *types.Transaction) (*types.Transaction, error) {
	signingMethod := types.LatestSignerForChainID(chainID)
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != ecrypto.PubkeyToAddress(*signer.PublicKey()) {
			return nil, bind.ErrNotAuthorized
		}

		hash := signingMethod.Hash(tx).Bytes()
		sig, err := signer.Sign(hash)
		if err != nil {
			return nil, err
		}

		return tx.WithSignature(signingMethod, sig.Raw())
	}
}

func SignEthMessage(signer Signer, msg []byte) (Signature, error) {
	hash := ComputeEthereumSignedMessageHash(msg)
	sig, err := signer.Sign(hash)
	if err != nil {
		return Signature{}, err
	}

	// This step is necessary to remain compatible with the ecrecover precompile
	if int(sig.V) < 27 {
		sig.V = byte(int(sig.V + 27))
	}

	return sig, nil
}

func RecoverEthMessageSigner(signature Signature, message []byte) (*ecdsa.PublicKey, error) {
	sig := signature
	if int(sig.V) >= 27 {
		sig.V = byte(int(sig.V - 27))
	}

	hash := ComputeEthereumSignedMessageHash(message)
	pubKey, err := secp256k1.RecoverPubkey(hash, sig.Raw())
	if err != nil {
		return nil, err
	}
	ecdsaPubKey, err := ecrypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return nil, err
	}

	return ecdsaPubKey, nil
}

func RecoverEthMessageSignerAddress(signature Signature, message []byte) (common.Address, error) {
	ecdsaPubKey, err := RecoverEthMessageSigner(signature, message)
	if err != nil {
		return common.Address{}, err
	}

	return ecrypto.PubkeyToAddress(*ecdsaPubKey), nil
}

// ComputeEthereumSignedMessageHash accepts an arbitrary message, prepends a known message,
// and hashes the result using keccak256. The known message added to the input before hashing is
// "\x19Ethereum Signed Message:\n" + len(message).
func ComputeEthereumSignedMessageHash(message []byte) []byte {
	return ecrypto.Keccak256(
		
			fmt.Appendf(nil, "\x19Ethereum Signed Message:\n%d%s", len(message), string(message)),
		,
	)
}
