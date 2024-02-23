package userop

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type RpcBackend interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

type rpcBackendImpl struct {
	*rpc.Client
}

func NewRpcBackend(rpcUrl url.URL) (RpcBackend, error) {
	if rpcUrl.Scheme != "http" && rpcUrl.Scheme != "https" {
		return nil, fmt.Errorf("rpcUrl must be an HTTP url")
	}

	client, err := rpc.Dial(rpcUrl.String())
	if err != nil {
		return nil, err
	}

	return &rpcBackendImpl{Client: client}, nil
}

type EthBackend interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	ethereum.TransactionSender
	ethereum.ContractCaller

	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)
	WaitMinedPeriod() time.Duration
	RPC() *rpc.Client
	bind.ContractBackend
}

type ethBackendImpl struct {
	*ethclient.Client
}

func NewEthBackend(rpcUrl url.URL) (EthBackend, error) {
	if rpcUrl.Scheme != "ws" && rpcUrl.Scheme != "wss" {
		return nil, fmt.Errorf("rpcUrl must be a WS url")
	}

	client, err := ethclient.Dial(rpcUrl.String())
	if err != nil {
		return nil, err
	}

	return &ethBackendImpl{client}, nil
}

func (n *ethBackendImpl) WaitMinedPeriod() time.Duration {
	return time.Second
}

func (n *ethBackendImpl) RPC() *rpc.Client {
	return n.Client.Client()
}
