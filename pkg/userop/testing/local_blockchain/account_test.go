package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	PrivateKey   *ecdsa.PrivateKey
	Address      common.Address
	TransactOpts *bind.TransactOpts
}

func NewAccount(chainID *big.Int) (Account, error) {
	pvk, err := ecrypto.GenerateKey()
	if err != nil {
		return Account{}, fmt.Errorf("failed to generate deployer private key: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(pvk, chainID)
	if err != nil {
		return Account{}, err
	}

	return Account{
		PrivateKey:   pvk,
		Address:      ecrypto.PubkeyToAddress(pvk.PublicKey),
		TransactOpts: opts,
	}, nil
}

func NewAccountWithBalance(ctx context.Context, chainID, balance *big.Int, node *EthNode) (Account, error) {
	account, err := NewAccount(chainID)
	if err != nil {
		return Account{}, err
	}

	if balance == nil {
		return Account{}, fmt.Errorf("amount is nil")
	}
	gethCmd := fmt.Sprintf(
		"eth.sendTransaction({from: eth.coinbase, to: '%s', value: web3.toWei(%d, 'ether')})",
		account.Address, balance.Uint64(),
	)

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
