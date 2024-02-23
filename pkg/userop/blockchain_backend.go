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

type RPCBackend interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

type rpcBackendImpl struct {
	*rpc.Client
}

func NewRPCBackend(rpcURL url.URL) (RPCBackend, error) {
	if rpcURL.Scheme != "http" && rpcURL.Scheme != "https" && rpcURL.Scheme != "ws" && rpcURL.Scheme != "wss" {
		return nil, fmt.Errorf("RPC uURL must be an HTTP or WS url")
	}

	client, err := rpc.Dial(rpcURL.String())
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

func NewEthBackend(rpcURL url.URL) (EthBackend, error) {
	if rpcURL.Scheme != "ws" && rpcURL.Scheme != "wss" {
		return nil, fmt.Errorf("RPC URL must be a WS url")
	}

	client, err := ethclient.Dial(rpcURL.String())
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
