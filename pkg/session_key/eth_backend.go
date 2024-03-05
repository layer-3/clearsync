package session_key

import (
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: use EthBackend from userop package and extract it to a common package?
type ethBackend interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	ethereum.TransactionSender
	ethereum.ContractCaller

	bind.ContractBackend
}

type ethBackendImpl struct {
	*ethclient.Client
}

func NewEthBackend(rpcURL url.URL) (ethBackend, error) {
	if rpcURL.Scheme != "http" && rpcURL.Scheme != "https" && rpcURL.Scheme != "ws" && rpcURL.Scheme != "wss" {
		return nil, fmt.Errorf("RPC URL must be an HTTP or WS url")
	}

	client, err := ethclient.Dial(rpcURL.String())
	if err != nil {
		return nil, err
	}

	return &ethBackendImpl{client}, nil
}
