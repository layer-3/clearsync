package local_blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

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
	amount *big.Int, // in wei
	node *EthNode,
) (Account, error) {
	account, err := NewAccount(ctx, node)
	if err != nil {
		return Account{}, err
	}

	if amount == nil {
		return Account{}, fmt.Errorf("specified balance is nil")
	}

	if err := node.FundAccount(ctx, account, amount); err != nil {
		return Account{}, err
	}

	return account, nil
}
