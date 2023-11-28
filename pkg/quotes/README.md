# Quotes service

The Quotes Service is a centralized system responsible for recording and disseminating real-time trade data 
from various trading platforms.

Available drivers:
- Binance
- Kraken
- Opendax
- Bitfaker

## Interface to connect

```go
type Driver interface {
	Subscribe(market Market) error
	Start(markets []Market) error
	Stop() error
}
```

## Type of price source

| Top Tier CEX | Altcoins CEX | FIAT       | DEX         |
|--------------|--------------|------------|-------------|
| BitFinex     | Gate         | UpBit      | Sushi       |
| OKX          | MEXC         | Kraken     | PancakeSwap |
| Binance      | KuCoin       | Coinbase   | Uniswap     |
| weight: 20   | weight: 5    | weight: 15 | weight: 50  |

## Last Price

For candle sticks, Recent trade, tickers last price is calculated as follows:

```
last_price = price
```

## Market Price

Used mainly in risk management for portfolio evaluation:

```
index_price = EMA20(price x ( trade_size x weight / active_weights ))
# active_weight being the sum of weight where this market exists (ex: KuCoin:5 + uniswap:50)
# EMA20 is likely 20 at 1 min scale
```

## How Uniswap adapter calculates swap price?

### Motivation

Uniswap V3 uses [Q notation](https://en.wikipedia.org/wiki/Q_(number_format)), which is a type of fixed-point arithmetics, to encode swap price.
 
Q notation allows variables to remain integers, but function similarly to floating point numbers.
Additionally piping the price through square root allows to reduce dimentionality of the number.
Predictable size of the number encoded this way enables efficient caching and fast retrieval of data from chain.

### How to decode price?

Actually this is a two step process:
1. Decode price
2. Convert the price into wei/wei ratio

#### Step 1

Here's the formula:
```
sqrtPrice = sqrtPriceX96 / (2^96)
price = sqrtPrice^2
```

#### Step 2

ERC20 tokens have built in decimal values.
For example, 1 WETH actually represents WETH in the contract whereas USDC is 10^6.
USDC has 6 decimals and WETH has 18.
So the price calculated on step 1 actually depicts [wei of token0] / [unit token of token1].
Now let's convert that into [wei] / [wei] ratio:

```
price / (10^(decimal1-decimal0))
```

## TODO:

- [x] remove Finex dependencies
- [ ] add specs or amendments to current interface