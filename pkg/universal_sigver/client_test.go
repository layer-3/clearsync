package universal_sigver

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/userop"
	"github.com/layer-3/clearsync/pkg/userop/testing/local_blockchain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// FIXME: rename `local_blockchain` to something better
func TestVerify(t *testing.T) {
	ctx := context.Background()

	// 1. Setup node, contracts, signer
	node := local_blockchain.NewEthNode(ctx, t)
	contracts := local_blockchain.SetupContracts(ctx, t, node)
	eoa, err := local_blockchain.NewAccountWithBalance(ctx, big.NewInt(42), node)
	require.NoError(t, err)
	signer := signer_pkg.NewLocalSigner(eoa.PrivateKey)

	// 2. Setup smart wallet config, sigver and sign message
	swCfg := smart_wallet.Config{
		Type:           &smart_wallet.KernelType,
		Factory:        contracts.Factory,
		Logic:          contracts.Logic,
		ECDSAValidator: contracts.ECDSAValidator,
	}
	sigver := NewUniversalSigver(node.Client, &swCfg, &contracts.EntryPoint)
	msg := []byte("hello")
	sig, err := signer_pkg.SignEthMessage(signer, msg)
	require.NoError(t, err)
	msgHash := common.BytesToHash(signer_pkg.ComputeEthereumSignedMessageHash(msg))
	rawSig := sig.Raw()

	t.Run("Successfuly verify ECDSA signature", func(t *testing.T) {
		ok, err := sigver.Verify(ctx, signer.CommonAddress(), msgHash, rawSig)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("Successfuly verify ERC-1271 signature", func(t *testing.T) {
		index := decimal.NewFromInt(0)

		// 3. Start bundler, create client and kernel signer
		bundlerUrl := local_blockchain.NewBundler(ctx, t, node, contracts.EntryPoint)
		useropCfg := userop.ClientConfig{}
		useropCfg.Init()
		useropCfg.ProviderURL = node.LocalURL.String()
		useropCfg.BundlerURL = bundlerUrl.String()
		useropCfg.EntryPoint = contracts.EntryPoint
		useropCfg.SmartWallet = swCfg
		useropCfg.Paymaster = userop.PaymasterConfig{
			Type: &userop.PaymasterDisabled,
		}

		client, err := userop.NewClient(useropCfg)
		require.NoError(t, err)
		kernelSigner := userop.SignerForKernel(signer)

		// 4. Calculate SW address and transfer some funds to it
		swAddr, err := client.GetAccountAddress(ctx, eoa.Address, index)
		require.NoError(t, err)
		local_blockchain.SendNative(ctx, t, node, eoa, local_blockchain.Account{Address: swAddr}, decimal.NewFromInt(1e18)) // send 100 wei

		// 5. Create and send user operation
		calls := smart_wallet.Calls{{To: signer.CommonAddress(), Value: big.NewInt(1)}}
		wdo := &userop.WalletDeploymentOpts{Index: index, Owner: signer.CommonAddress()}
		op, err := client.NewUserOp(ctx, swAddr, kernelSigner, calls, wdo, nil)
		require.NoError(t, err)
		done, err := client.SendUserOp(ctx, op)
		require.NoError(t, err)
		receipt := <-done
		require.True(t, receipt.Success)

		// 6. Verify ERC-1271 signature
		ok, err := sigver.Verify(ctx, swAddr, msgHash, rawSig)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("Successfuly verify ERC-6492 signature", func(t *testing.T) {
		index := decimal.NewFromInt(1)

		// 3. Calculate smart wallet address
		swAddr, err := smart_wallet.GetAccountAddress(ctx, node.Client, swCfg, contracts.EntryPoint, signer.CommonAddress(), index)
		require.NoError(t, err)

		// 4. Pack ERC-6492 signature
		erc6492Sig, err := sigver.PackERC6492Sig(ctx, signer.CommonAddress(), index, rawSig)
		require.NoError(t, err)

		// 5. Verify ERC-6492 signature
		ok, err := sigver.Verify(ctx, swAddr, msgHash, erc6492Sig)
		require.NoError(t, err)
		require.True(t, ok)
	})
}

func TestPackIsValidSigCall(t *testing.T) {
	pvk, err := crypto.GenerateKey()
	require.NoError(t, err)
	signer := signer_pkg.NewLocalSigner(pvk)

	msg := []byte("hello again")
	msgHash := common.BytesToHash(signer_pkg.ComputeEthereumSignedMessageHash(msg))
	sig, err := signer_pkg.SignEthMessage(signer, msg)
	require.NoError(t, err)

	calldata := packIsValidSigCall(signer.CommonAddress(), msgHash, sig.Raw())
	bytecodeLen := len(hexutil.MustDecode(validateSigOffchainBytecode))
	require.Equal(t, validateSigOffchainBytecode, hexutil.Encode(calldata[:bytecodeLen]))

	args := abi.Arguments{
		{Name: "signer", Type: address},
		{Name: "hash", Type: bytes32},
		{Name: "signature", Type: bytes},
	}

	unpacked, err := args.Unpack(calldata[bytecodeLen:])
	require.NoError(t, err)

	unpackedAddress, ok := unpacked[0].(common.Address)
	require.True(t, ok)
	require.Equal(t, signer.CommonAddress(), unpackedAddress)

	unpackedHash, ok := unpacked[1].([32]byte)
	require.True(t, ok)
	require.Equal(t, msgHash.Hex(), hexutil.Encode(unpackedHash[:]))

	unpackedSig, ok := unpacked[2].([]byte)
	require.True(t, ok)
	require.Equal(t, sig.Raw(), unpackedSig)
}
