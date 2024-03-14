package universal_sigver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/shopspring/decimal"
)

var entryPointV0_6Address = common.HexToAddress("0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789")

type Client interface {
	Verify(signer common.Address, messageHash common.Hash, signature []byte) (bool, error)
	SignERC6492(ctx context.Context, owner signer.Signer, index decimal.Decimal, msg []byte) ([]byte, error)
	PackERC6492Sig(ctx context.Context, ownerAddress common.Address, index decimal.Decimal, sig []byte) ([]byte, error)
}

type backend struct {
	provider          *ethclient.Client
	smartWalletConfig *smart_wallet.Config
	entryPointAddress *common.Address
}

func NewUniversalSigver(provider *ethclient.Client, smartWalletConfig *smart_wallet.Config, entryPointAddress *common.Address) Client {
	var entryPointAddress_ = entryPointAddress
	if entryPointAddress_ == nil {
		entryPointAddress_ = &entryPointV0_6Address
	}
	return &backend{
		provider:          provider,
		smartWalletConfig: smartWalletConfig,
		entryPointAddress: entryPointAddress_,
	}
}

func (b *backend) Verify(signer common.Address, messageHash common.Hash, signature []byte) (bool, error) {
	calldata := packIsValidSigCall(signer, messageHash, signature)

	var res []byte
	err := b.provider.Client().Call(&res, "eth_call", map[string]interface{}{
		"data:": calldata,
	})
	if err != nil {
		return false, fmt.Errorf("failed to call ValidateSigOffchain: %w", err)
	}

	return hexutil.Encode(res) == validateSigOffchainSuccess, nil
}

// NOTE: no support for contract being deployed but not ready
// TODO: check for ERC-712 support
func (b *backend) SignERC6492(ctx context.Context, owner signer.Signer, index decimal.Decimal, msg []byte) ([]byte, error) {
	sig, err := owner.Sign(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	return b.PackERC6492Sig(ctx, owner.CommonAddress(), index, sig.Raw())
}

func (b *backend) PackERC6492Sig(ctx context.Context, ownerAddress common.Address, index decimal.Decimal, sig []byte) ([]byte, error) {
	swAddress, err := smart_wallet.GetAccountAddress(ctx, b.provider, *b.smartWalletConfig, *b.entryPointAddress, ownerAddress, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get smart wallet address: %w", err)
	}

	if isDeployed, err := smart_wallet.IsAccountDeployed(ctx, b.provider, swAddress); err != nil {
		return nil, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	} else if isDeployed {
		// use ERC-1271 signature
		return nil, fmt.Errorf("smart wallet already deployed")
	}

	factoryCalldata, err := smart_wallet.GetInitCode(b.provider, *b.smartWalletConfig, ownerAddress, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get init code: %w", err)
	}

	return packERC6492Sig(b.smartWalletConfig.Factory, factoryCalldata, sig), nil
}
