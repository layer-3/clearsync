package local_backend

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	PrivateKey    *ecdsa.PrivateKey
	CommonAddress common.Address
	TransactOpts  *bind.TransactOpts
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
		PrivateKey:    pvk,
		CommonAddress: ecrypto.PubkeyToAddress(pvk.PublicKey),
		TransactOpts:  opts,
	}

	return a, nil
}

func NewAccountWithPrivateKey(privateKey *ecdsa.PrivateKey, chainID *big.Int) (Account, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return Account{}, err
	}

	a := Account{
		PrivateKey:    privateKey,
		CommonAddress: ecrypto.PubkeyToAddress(privateKey.PublicKey),
		TransactOpts:  opts,
	}

	return a, nil
}

func generateSimulatedBackendAccounts(n uint) ([]Account, error) {
	accounts := make([]Account, 0, n)
	for i := uint(0); i < n; i++ {
		a, err := NewAccount(SimulatedChainID)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}
