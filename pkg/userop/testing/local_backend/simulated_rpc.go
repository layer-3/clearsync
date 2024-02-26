package local_backend

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

var (
	SimulatedChainID                     = big.NewInt(1337)
	DefaultSimulatedBackendFundingAmount = big.NewInt(0).SetUint64(1000000000000000000) // 1 ETH
)

type SimulatedBackend struct {
	*simulated.Backend

	accounts        []Account
	deployer        Account
	waitMinedPeriod time.Duration
}

func NewSimulatedBackend() (*SimulatedBackend, error) {
	// Generate 10 accounts with 10 eth each
	accounts, err := generateSimulatedBackendAccounts(10)
	if err != nil {
		return nil, err
	}

	alloc := map[common.Address]types.Account{}
	for _, a := range accounts {
		alloc[a.CommonAddress] = types.Account{
			Balance: big.NewInt(0).Mul(DefaultSimulatedBackendFundingAmount, big.NewInt(2)),
		}
	}

	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(15_000_000))
	waitMinedPeriod := 10 * time.Millisecond

	// Simulates mining of blocks within some interval
	go func() {
		for {
			backend.Commit()
			<-time.After(waitMinedPeriod)
		}
	}()

	return &SimulatedBackend{
		Backend:         backend,
		accounts:        accounts,
		deployer:        accounts[0],
		waitMinedPeriod: waitMinedPeriod,
	}, nil
}

func (sb *SimulatedBackend) ChainID(_ context.Context) (*big.Int, error) {
	return SimulatedChainID, nil
}

func (sb *SimulatedBackend) WaitMinedPeriod() time.Duration {
	return sb.waitMinedPeriod
}
