package userop

import (
	"math/big"
	"sync"
)

// Implements a nonce key rotation for 2d nonces.
// https://docs.stackup.sh/docs/useroperation-nonce
type NonceKeyRotator struct {
	key *big.Int
	mu  sync.Mutex

	max *big.Int
}

func NewNonceKeyRotator() NonceKeyRotator {
	return NonceKeyRotator{
		key: big.NewInt(0),

		// The biggest possible key value is ^uint192, but we use 2^64 for simplicity.
		max: big.NewInt(0).Exp(big.NewInt(2), big.NewInt(64), nil),
	}
}

func (r *NonceKeyRotator) Next() *big.Int {
	r.mu.Lock()
	defer r.mu.Unlock()

	incKey := new(big.Int).Add(r.key, big.NewInt(1))
	modKey := new(big.Int).Mod(incKey, r.max)

	r.key = modKey
	return modKey
}
