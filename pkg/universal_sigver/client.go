package universal_sigver

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/shopspring/decimal"
)

type Client interface {
	Verify(signer common.Address, messageHash common.Hash, signature []byte) (bool, error)
	Sign(signerAddress common.Address, signer signer.Signer, msg []byte) ([]byte, error)
}

type backend struct {
	provider          *ethclient.Client
	smartWalletConfig smart_wallet.Config
}

func NewUniversalSigver(provider *ethclient.Client, smartWalletConfig smart_wallet.Config) Client {
	return &backend{
		provider:          provider,
		smartWalletConfig: smartWalletConfig,
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
func (b *backend) Sign(signerAddress common.Address, signer signer.Signer, msg []byte) ([]byte, error) {
	sig, err := signer.Sign(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	if signerAddress == signer.CommonAddress() {
		// ECDSA
		return sig.Raw(), nil
	}

	isDeployed, err := smart_wallet.IsAccountDeployed(b.provider, signerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}

	if isDeployed {
		// ERC-1271 signature
		return sig.Raw(), nil
	}

	// ERC-6492 signature
	// FIXME:
	index := decimal.NewFromInt(0)
	factoryCalldata, err := smart_wallet.GetInitCode(b.provider, b.smartWalletConfig, signerAddress, signer.CommonAddress(), index)
	if err != nil {
		return nil, fmt.Errorf("failed to get init code: %w", err)
	}

	return packERC6492Sig(b.smartWalletConfig.Factory, factoryCalldata, sig.Raw()), nil
}
