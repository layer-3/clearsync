package session_key

import (
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const SessionKeyValidatorAddress = "0x5C06CE2b673fD5E6e56076e40DD46aB67f5a72A5"
const ECDSAValidatorAddress = "0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"

type Config struct {
	ProviderUrl                string
	SessionKeyValidAfter       uint64
	SessionKeyValidUntil       uint64
	SessionKeyValidatorAddress common.Address
	ExecutorAddress            common.Address
	PaymasterAddress           common.Address
	Permissions                []Permission
}

func NewEthClient(rpcURLString string) (*ethclient.Client, error) {
	rpcURL, err := url.Parse(rpcURLString)
	if err != nil {
		return nil, fmt.Errorf("invalid RPC URL: %s", rpcURLString)
	}

	if rpcURL.Scheme != "http" && rpcURL.Scheme != "https" && rpcURL.Scheme != "ws" && rpcURL.Scheme != "wss" {
		return nil, fmt.Errorf("RPC URL must be an HTTP or WS url")
	}

	client, err := ethclient.Dial(rpcURL.String())
	if err != nil {
		return nil, err
	}

	return client, nil
}
