package universal_sigver

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/shopspring/decimal"
)

var entryPointV0_6Address = common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")

type Client interface {
	// Verify checks if the signature was signed by the signer.
	// This function supports ECDSA, ERC-1271 and ERC-6492 signatures.
	//
	// Parameters:
	// - ctx: the context
	// - signer: the address of the signer. EOA address if ECDSA signature is supplied, smart wallet address if ERC-1271 or ERC-6492 signature is supplied.
	// - messageHash: the hash of the message that was signed
	// - signature: the signature
	//
	// Returns:
	// - true if the signature is valid, false otherwise
	// - an error if an error occurred during the verification
	Verify(ctx context.Context, signer common.Address, messageHash common.Hash, signature []byte) (bool, error)

	// PackERC6492Sig packs an ERC-6492 signature by adding smart wallet factory address, calldata and ERC-6492 suffix.
	//
	// Parameters:
	// - ctx: the context
	// - ownerAddress: the address of the owner of the smart wallet
	// - index: the index of the smart wallet
	// - sig: the signature
	//
	// Returns:
	// - the packed signature
	// - an error if an error occurred during the packing
	PackERC6492Sig(ctx context.Context, ownerAddress common.Address, index decimal.Decimal, sig []byte) ([]byte, error)
}

type backend struct {
	provider          *ethclient.Client
	smartWalletConfig smart_wallet.Config
	entryPointAddress common.Address
}

func NewUniversalSigVer(providerURL string, smartWalletConfig smart_wallet.Config, entryPointAddress common.Address) (Client, error) {
	provider, err := ethclient.Dial(providerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	if entryPointAddress.Cmp(common.Address{}) == 0 {
		entryPointAddress = entryPointV0_6Address
	}

	return &backend{
		provider:          provider,
		smartWalletConfig: smartWalletConfig,
		entryPointAddress: entryPointAddress,
	}, nil
}

func (b *backend) Verify(ctx context.Context, signer common.Address, messageHash common.Hash, signature []byte) (bool, error) {
	calldata := packIsValidSigCall(signer, messageHash, signature)

	var res string
	err := b.provider.Client().CallContext(ctx, &res, "eth_call", map[string]interface{}{
		"data": hexutil.Encode(calldata),
	},
		"latest")
	if err != nil {
		var scError rpc.DataError
		if ok := errors.As(err, &scError); !ok {
			return false, fmt.Errorf("could not unpack error data: unexpected error type '%T' containing message %w)", err, err)
		}
		return false, fmt.Errorf("failed to call ValidateSigOffchain: %w, errorData: %s", err, scError.ErrorData())
	}

	return res == validateSigOffchainSuccess, nil
}

func (b *backend) PackERC6492Sig(ctx context.Context, ownerAddress common.Address, index decimal.Decimal, sig []byte) ([]byte, error) {
	swAddress, err := smart_wallet.GetAccountAddress(ctx, b.provider, b.smartWalletConfig, b.entryPointAddress, ownerAddress, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get smart wallet address: %w", err)
	}

	if isDeployed, err := smart_wallet.IsAccountDeployed(ctx, b.provider, swAddress); err != nil {
		return nil, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	} else if isDeployed {
		// use ERC-1271 signature instead
		return nil, fmt.Errorf("smart wallet already deployed")
	}

	factoryCalldata, err := smart_wallet.GetFactoryCallData(b.smartWalletConfig, ownerAddress, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get init code: %w", err)
	}

	return PackERC6492Sig(b.smartWalletConfig.Factory, factoryCalldata, sig), nil
}
