package local_blockchain

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	PrivateKey   *ecdsa.PrivateKey
	Address      common.Address
	TransactOpts *bind.TransactOpts
}

func NewAccount(ctx context.Context, node *EthNode) (Account, error) {
	chainID, err := node.Client.ChainID(ctx)
	if err != nil {
		return Account{}, err
	}

	pvk, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, fmt.Errorf("failed to generate deployer private key: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(pvk, chainID)
	if err != nil {
		return Account{}, err
	}

	return Account{
		PrivateKey:   pvk,
		Address:      crypto.PubkeyToAddress(pvk.PublicKey),
		TransactOpts: opts,
	}, nil
}

func NewAccountWithBalance(
	ctx context.Context,
	balance *big.Int, // in ether, not wei or gwei
	node *EthNode,
) (Account, error) {
	account, err := NewAccount(ctx, node)
	if err != nil {
		return Account{}, err
	}

	if balance == nil {
		return Account{}, fmt.Errorf("specified balance is nil")
	}
	gethCmd := fmt.Sprintf(
		"eth.sendTransaction({from: eth.coinbase, to: '%s', value: web3.toWei(%d, 'wei')})",
		account.Address, balance.Uint64(),
	)

	if pkStr := os.Getenv("DEPLOYER_PK"); pkStr != "" {
		privateKey, err := crypto.HexToECDSA(pkStr)
		if err != nil {
			return Account{}, fmt.Errorf("failed to parse deployer private key: %w", err)
		}

		deployerAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

		nonce, err := node.Client.PendingNonceAt(context.Background(), deployerAddress)
		if err != nil {
			return Account{}, fmt.Errorf("failed to get nonce: %w", err)
		}

		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &account.Address,
			Value:    balance,
			Gas:      21000, // default gas limit for native transfer
			GasPrice: big.NewInt(50000000000),
		})

		chainID, err := node.Client.NetworkID(context.Background())
		if err != nil {
			return Account{}, fmt.Errorf("failed to get chain ID: %w", err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return Account{}, fmt.Errorf("failed to sign tx: %w", err)
		}

		err = node.Client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			return Account{}, fmt.Errorf("failed to send tx: %w", err)
		}

		_, err = waitMinedV2(ctx, node, signedTx)
		if err != nil {
			return Account{}, fmt.Errorf("failed to wait for tx: %w", err)
		}

		return account, nil
	}

	exitCode, result, err := node.Container.Exec(ctx, []string{"geth", "attach", "--exec", gethCmd, node.LocalURL.String()})
	if err != nil || exitCode != 0 {
		return Account{}, fmt.Errorf("failed to exec increment balance: %w (exit code %d)", err, exitCode)
	}

	scanner := bufio.NewScanner(result)
	for scanner.Scan() {
		slog.Debug(scanner.Text())
	}

	return account, nil
}

func waitMinedV2(ctx context.Context, node *EthNode, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(1 * time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := node.Client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}

		bn, _ := node.Client.BlockNumber(context.Background())
		fmt.Println(bn)

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
