package base

import (
	"fmt"
	"math/big"
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

// SwapParser defines an interface for parsing swap events from decentralized
// exchanges (DEXs). Implementations of this interface extract/compute trade
// details from raw swap events.
type SwapParser interface {
	ParseSwap(logger *zap.SugaredLogger) (common.TradeEvent, error)
}

// SwapV2 represents a generic swap parser
// for events based on Uniswap v2 schema.
type SwapV2[Event any, EventIterator dexEventIterator] struct {
	RawAmount0In  *big.Int
	RawAmount0Out *big.Int
	RawAmount1In  *big.Int
	RawAmount1Out *big.Int
	Pool          *DexPool[Event, EventIterator]
}

func (o *SwapV2[Event, EventIterator]) ParseSwap(logger *zap.SugaredLogger) (trade common.TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorw(common.ErrSwapParsing.Error())
			err = fmt.Errorf("%s: %s", common.ErrSwapParsing, r)
		}
	}()

	if o.Pool.Reversed {
		copyAmount0In, copyAmount0Out := o.RawAmount0In, o.RawAmount0Out
		o.RawAmount0In, o.RawAmount0Out = o.RawAmount1In, o.RawAmount1Out
		o.RawAmount1In, o.RawAmount1Out = copyAmount0In, copyAmount0Out
	}

	var takerType common.TakerType
	var price decimal.Decimal
	var amount decimal.Decimal
	var total decimal.Decimal

	baseDecimals := o.Pool.BaseToken.Decimals
	quoteDecimals := o.Pool.QuoteToken.Decimals

	switch {
	case isValidNonZero(o.RawAmount0In) && isValidNonZero(o.RawAmount1Out):
		amount1Out := decimal.NewFromBigInt(o.RawAmount1Out, 0).Div(ten.Pow(quoteDecimals))
		amount0In := decimal.NewFromBigInt(o.RawAmount0In, 0).Div(ten.Pow(baseDecimals))

		takerType = common.TakerTypeSell
		price = amount1Out.Div(amount0In) // NOTE: may panic here if `amount0In` is zero
		total = amount1Out
		amount = amount0In

	case isValidNonZero(o.RawAmount0Out) && isValidNonZero(o.RawAmount1In):
		amount0Out := decimal.NewFromBigInt(o.RawAmount0Out, 0).Div(ten.Pow(baseDecimals))
		amount1In := decimal.NewFromBigInt(o.RawAmount1In, 0).Div(ten.Pow(quoteDecimals))

		takerType = common.TakerTypeBuy
		price = amount1In.Div(amount0Out) // NOTE: may panic here if `amount0Out` is zero
		total = amount1In
		amount = amount0Out
	default:
		return trade, fmt.Errorf("market %s: unknown swap type", o.Pool.Market)
	}

	trade = common.TradeEvent{
		Market:    o.Pool.Market,
		Price:     price,
		Amount:    amount.Abs(),
		Total:     total,
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return trade, nil
}

// SwapV3 represents a generic swap parser
// for events based on Uniswap v3 schema.
type SwapV3[Event any, EventIterator dexEventIterator] struct {
	RawAmount0      *big.Int
	RawAmount1      *big.Int
	RawSqrtPriceX96 *big.Int
	Pool            *DexPool[Event, EventIterator]
}

func (o *SwapV3[Event, EventIterator]) ParseSwap(logger *zap.SugaredLogger) (trade common.TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorw(common.ErrSwapParsing.Error(), "pool", o.Pool)
			err = fmt.Errorf("%s: %s", common.ErrSwapParsing, r)
		}
	}()

	if !isValidNonZero(o.RawAmount0) {
		return trade, fmt.Errorf("raw amount0 (%s) is not a valid non-zero number", o.RawAmount0)
	}
	amount0 := decimal.NewFromBigInt(o.RawAmount0, 0)

	if !isValidNonZero(o.RawAmount1) {
		return trade, fmt.Errorf("raw amount1 (%s) is not a valid non-zero number", o.RawAmount0)
	}
	amount1 := decimal.NewFromBigInt(o.RawAmount1, 0)

	if !isValidNonZero(o.RawSqrtPriceX96) {
		return trade, fmt.Errorf("raw sqrtPriceX96 (%s) is not a valid non-zero number", o.RawSqrtPriceX96)
	}
	sqrtPriceX96 := decimal.NewFromBigInt(o.RawSqrtPriceX96, 0)

	if o.Pool.Reversed {
		amount0, amount1 = amount1, amount0
	}

	// Normalize swap amounts.
	baseDecimals, quoteDecimals := o.Pool.BaseToken.Decimals, o.Pool.QuoteToken.Decimals
	amount0Normalized := amount0.Div(ten.Pow(baseDecimals)).Abs()
	amount1Normalized := amount1.Div(ten.Pow(quoteDecimals)).Abs()

	// Calculate swap price.
	price := o.calculatePrice(sqrtPriceX96, baseDecimals, quoteDecimals, o.Pool.Reversed)
	// Apply a fallback strategy in case the primary one fails.
	// This should never happen, but just in case.
	if price.IsZero() {
		price = amount1Normalized.Div(amount0Normalized)
	}

	// Calculate trade side, amount and total.
	takerType := common.TakerTypeBuy
	amount, total := amount0Normalized, amount1Normalized
	if (!o.Pool.Reversed && amount0.Sign() < 0) || (o.Pool.Reversed && amount1.Sign() < 0) {
		takerType = common.TakerTypeSell
	}

	trade = common.TradeEvent{
		Market:    o.Pool.Market,
		Price:     price,
		Amount:    amount, // amount of BASE token received
		Total:     total,  // total cost in QUOTE token
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return trade, nil
}

var (
	two      = decimal.NewFromInt(2)
	ten      = decimal.NewFromInt(10)
	priceX96 = two.Pow(decimal.NewFromInt(96))
)

// calculatePrice method calculates the price per token at which the swap was
// performed using the sqrtPriceX96 value supplied with every on-chain swap
// event.
//
// General formula is as follows:
// price = ((sqrtPriceX96 / 2**96)**2) / (10**decimal1 / 10**decimal0)
//
// See the math explained at https://blog.uniswap.org/uniswap-v3-math-primer
func (o *SwapV3[Event, EventIterator]) calculatePrice(sqrtPriceX96, baseDecimals, quoteDecimals decimal.Decimal, reversedPool bool) decimal.Decimal {
	if reversedPool {
		baseDecimals, quoteDecimals = quoteDecimals, baseDecimals
	}

	// Simplification for denominator calculations:
	// 10**decimal1 / 10**decimal0 -> 10**(decimal1 - decimal0)
	decimals := quoteDecimals.Sub(baseDecimals)

	numerator := sqrtPriceX96.Div(priceX96).Pow(two)
	denominator := ten.Pow(decimals)

	if reversedPool {
		return denominator.Div(numerator)
	}
	return numerator.Div(denominator)
}

// isValidNonZero checks if the given big.Int value is a valid number that is
// both non-nil and non-zero. Note that negative values are allowed as they
// represent a reduction in the balance of the pool.
func isValidNonZero(x *big.Int) bool {
	return x != nil && x.Sign() != 0
}
