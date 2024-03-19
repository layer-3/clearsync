package local_blockchain

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

		deployerAccount := Account{
			PrivateKey: privateKey,
			Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		}

		err = SendNative(ctx, node, deployerAccount, account, balance)
		if err != nil {
			return Account{}, fmt.Errorf("failed to send native: %w", err)
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
