package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/testcontainers/testcontainers-go"
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

	a := Account{
		PrivateKey:   pvk,
		Address:      ecrypto.PubkeyToAddress(pvk.PublicKey),
		TransactOpts: opts,
	}

	return a, nil
}

func NewAccountWithPrivateKey(privateKey *ecdsa.PrivateKey, chainID *big.Int) (Account, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return Account{}, err
	}

	a := Account{
		PrivateKey:   privateKey,
		Address:      ecrypto.PubkeyToAddress(privateKey.PublicKey),
		TransactOpts: opts,
	}

	return a, nil
}

func generateSimulatedBackendAccounts(n uint, ethClient ethereum.ChainIDReader) ([]Account, error) {
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, 0, n)
	for i := uint(0); i < n; i++ {
		a, err := NewAccount(chainID)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}

type AccountBalance struct {
	gethContainer testcontainers.Container
	rpcURL        url.URL
}

func NewAccountBalance(gethContainer testcontainers.Container, rpcURL url.URL) *AccountBalance {
	return &AccountBalance{gethContainer: gethContainer, rpcURL: rpcURL}
}

func (a *AccountBalance) IncrementBalance(ctx context.Context, account Account, amount *big.Int) error {
	if amount == nil {
		return fmt.Errorf("amount is nil")
	}
	gethCmd := fmt.Sprintf(
		"eth.sendTransaction({from: eth.coinbase, to: '%s', value: web3.toWei(%d, 'ether')})",
		account.Address, amount.Uint64(),
	)

	exitCode, result, err := a.gethContainer.Exec(ctx, []string{"geth", "attach", "--exec", gethCmd, a.rpcURL.String()})
	if err != nil || exitCode != 0 {
		return fmt.Errorf("failed to exec increment balance: %w (exit code %d)", err, exitCode)
	}

	scanner := bufio.NewScanner(result)
	for scanner.Scan() {
		slog.Debug(scanner.Text())
	}

	return nil
}
