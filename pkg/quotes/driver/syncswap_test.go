package driver

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/quotes/common"
)

// Mock data for testing
var (
	baseTokenDecimals  = big.NewInt(18) // Example: 18 decimals for ETH
	quoteTokenDecimals = big.NewInt(6)  // Example: 6 decimals for USDC
	pool               = dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator]{
		BaseToken:  poolToken{Decimals: decimal.NewFromBigInt(baseTokenDecimals, 0)},
		QuoteToken: poolToken{Decimals: decimal.NewFromBigInt(quoteTokenDecimals, 0)},
	}
	market = common.Market{} // Assuming market is correctly initialized for the test
)

func TestParseSwapSellETHUSDC(t *testing.T) {
	// Example sell swap: selling Amount0In of quoteToken for Amount1Out of baseToken
	swap := &isyncswap_pool.ISyncSwapPoolSwap{
		Amount0In:  big.NewInt(26272675352021868), // USDC in example, with 6 decimals
		Amount1Out: big.NewInt(103073083),         // ETH out, with 18 decimals
	}

	expectedPrice, _ := decimal.NewFromString("3923.204683913839496")
	expectedAmount := decimal.NewFromFloat(0.0262726753520219)
	expectedTotal := decimal.NewFromFloat(103.073083)

	s := syncswap{} // Assuming s is correctly initialized for the test
	tradeEvent, err := s.parseSwap(swap, &pool)

	assert.Nil(t, err)
	assert.Equal(t, common.TakerTypeSell, tradeEvent.TakerType)
	assert.True(t, tradeEvent.Price.Equals(expectedPrice))
	assert.True(t, tradeEvent.Amount.Equals(expectedAmount))
	assert.True(t, tradeEvent.Total.Equals(expectedTotal))
}

func TestParseSwapBuyETHUSDC(t *testing.T) {
	// Example buy swap: buying Amount0Out of quoteToken with Amount1In of baseToken
	swap := &isyncswap_pool.ISyncSwapPoolSwap{
		Amount0Out: big.NewInt(21958435519730050), // USDC out, with 6 decimals
		Amount1In:  big.NewInt(85000000),          // ETH in, with 18 decimals
	}

	expectedPrice, _ := decimal.NewFromString("3870.9497278904854025")
	expectedAmount := decimal.NewFromFloat(0.0219584355197301)
	expectedTotal := decimal.NewFromInt(85)

	s := syncswap{} // Assuming s is correctly initialized for the test
	tradeEvent, err := s.parseSwap(swap, &pool)

	assert.Nil(t, err)
	assert.Equal(t, common.TakerTypeBuy, tradeEvent.TakerType)
	assert.True(t, tradeEvent.Price.Equals(expectedPrice))
	assert.True(t, tradeEvent.Amount.Equals(expectedAmount))
	assert.True(t, tradeEvent.Total.Equals(expectedTotal))
}

func TestParseSwapBuyLINDAWETH(t *testing.T) {
	baseTokenDecimals = big.NewInt(18)  // Example: 18 decimals for LINDA
	quoteTokenDecimals = big.NewInt(18) // Example: 18 decimals for WETH
	pool = dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator]{
		BaseToken:  poolToken{Decimals: decimal.NewFromBigInt(baseTokenDecimals, 0)},
		QuoteToken: poolToken{Decimals: decimal.NewFromBigInt(quoteTokenDecimals, 0)},
	}

	value := new(big.Int)
	value.SetString("5552128050580394634028004", 10)
	// Example buy swap: buying Amount0Out of quoteToken with Amount1In of baseToken
	swap := &isyncswap_pool.ISyncSwapPoolSwap{
		Amount0Out: value,                         // Linda out example, with 18 decimals
		Amount1In:  big.NewInt(41092000000000000), // ETH in, with 18 decimals
	}

	expectedPrice := decimal.NewFromFloat(0.0000000074011261)
	expectedAmount, _ := decimal.NewFromString("5552128.050580394634028")
	expectedTotal := decimal.NewFromFloat(0.041092)

	s := syncswap{} // Assuming s is correctly initialized for the test
	tradeEvent, err := s.parseSwap(swap, &pool)

	assert.Nil(t, err)
	assert.Equal(t, common.TakerTypeBuy, tradeEvent.TakerType)
	assert.True(t, tradeEvent.Price.Equals(expectedPrice))
	assert.True(t, tradeEvent.Amount.Equals(expectedAmount))
	assert.True(t, tradeEvent.Total.Equals(expectedTotal))
}

func TestParseSwapSellLINDAWETH(t *testing.T) {
	baseTokenDecimals = big.NewInt(18)  // Example: 18 decimals for LINDA
	quoteTokenDecimals = big.NewInt(18) // Example: 18 decimals for WETH
	pool = dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator]{
		BaseToken:  poolToken{Decimals: decimal.NewFromBigInt(baseTokenDecimals, 0)},
		QuoteToken: poolToken{Decimals: decimal.NewFromBigInt(quoteTokenDecimals, 0)},
	}

	value := new(big.Int)
	value.SetString("26372443657387303085144227", 10)
	// Example sell swap: selling Amount0In of quoteToken for Amount1Out of baseToken
	swap := &isyncswap_pool.ISyncSwapPoolSwap{
		Amount0In:  value,                          // Linda in example, with 18 decimals
		Amount1Out: big.NewInt(193601651777829038), // ETH out, with 18 decimals
	}

	expectedPrice := decimal.NewFromFloat(0.0000000073410585)
	expectedAmount, _ := decimal.NewFromString("26372443.6573873030851442")
	expectedTotal := decimal.NewFromFloat(0.193601651777829)

	s := syncswap{} // Assuming s is correctly initialized for the test
	tradeEvent, err := s.parseSwap(swap, &pool)

	assert.Nil(t, err)
	assert.Equal(t, common.TakerTypeSell, tradeEvent.TakerType)
	assert.True(t, tradeEvent.Price.Equals(expectedPrice))
	assert.True(t, tradeEvent.Amount.Equals(expectedAmount))
	assert.True(t, tradeEvent.Total.Equals(expectedTotal))
}
