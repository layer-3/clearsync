# Quotes service

The Quotes Service is a centralized system responsible for recording and disseminating real-time trade data 
from various trading platforms.

Available drivers:
- Index Price
- Binance
- Kraken
- Opendax
- Bitfaker
- Uniswap V3

## Interface to connect

```go
type Driver interface {
	Subscribe(market Market) error
	Start(markets []Market) error
	Stop() error
}
```

## Types of price sources

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

## Index Price

Index Price is used mainly in risk management for portfolio evaluation, and is aggregated from multiple sources.

```
index_price = priceWeightEMA/weightEMA

priceWeightEMA = EMA20((Volume)*(Price)*(DriverWeight/activeWeights))

weightEMA = EMA20(Volume*(DriverWeight/activeWeights))
```
```
activeWeights is the sum of weight where this market is active (ex: KuCoin:5 + uniswap:50)
```

This formula calculates a weighted EMA price, using trade volume, and active driver weight as additional weights.

EMA20 is calculated based on the last 20 trades. If on the startup there is no initial data, the values of the first trade received are used as the initial EMA.

Let's assume 5 trades have been received:

```
{source, price, amount}
{binance, 41000, 0.3}
{binance, 42500, 0.5}
{uniswap, 55000, 0.6}
{uniswap, 50000, 0.4}
{binance, 40000, 1.0}
```

In the first example, drivers' weights are equal: ```2``` for the Binance driver, and ```2``` for the Uniswap.

```Weighted EMA price change with each trade: 41000, 41223.8806, 43500.32787, 43873.32054, 43343.1976```

In the second example, the Uniswap driver has ```0``` weight:

```Weighted EMA price change with each trade: 41000, 41223.8806, 41223.8806, 41223.8806, 40872.3039```

In the third example, we are setting the amound of the first 4 trades to ```1```, and the last trade to ```10```, and this time the price increases by 2000 with each trade:

```
{source, price, amount}
{binance, 40000, 1.0}
{binance, 42000, 1.0}
{binance, 44000, 1.0}
{binance, 46000, 1.0}
{binance, 48000, 10.0}
```

```Weighted EMA price change with each trade: 40000, 40190.47619, 40553.28798, 41072.02246, 44624.83145```

We observe that the trade with 10X amount influenced the final EMA more significantly than all the other trades.

**Described test scenarios can be found in ```index_test.go```.**

*PriceCache*

```
type PriceCache interface {
	Get(market string) (decimal.Decimal, decimal.Decimal)                         // Returns priceWeight and weight EMAs for a market
	Update(driver DriverType, market string, priceWeight, weight decimal.Decimal) // Updates priceWeight and weight EMAs for a market with a new value
	ActiveWeights(market string) decimal.Decimal                                  // Returns the sum of active driver weights for the market.
}
```

- PriceCache stores previous priceWeight and weight EMAs.
- It also stores a map of active drivers for a market. By default, no drivers are active. 
  When the cache receives the first price, it makes the source driver active for the market.
	This approach solves the problem when different drivers may support different markets.

## How Uniswap adapter calculates swap price

### Method 1: from sqrtPriceX96

#### Motivation

Uniswap V3 uses [Q notation](https://en.wikipedia.org/wiki/Q_(number_format)), which is a type of fixed-point arithmetics, to encode swap price.
 
Q notation allows variables to remain integers, but function similarly to floating point numbers.
Additionally piping the price through square root allows to reduce dimentionality of the number.
Predictable size of the number encoded this way enables efficient caching and fast retrieval of data from chain.

#### How to calculate price?

Actually this is a two step process:
1. Decode price
2. Convert the price into wei/wei ratio

##### Step 1

Here's the formula:
```
sqrtPrice = sqrtPriceX96 / (2^96)
price = sqrtPrice^2
```

##### Step 2

ERC20 tokens have built in decimal values.
For example, 1 WETH actually represents WETH in the contract whereas USDC is 10^6.
USDC has 6 decimals and WETH has 18.
So the price calculated on step 1 actually depicts [wei of token0] / [unit token of token1].
Now let's convert that into [wei] / [wei] ratio:

```
price0 = price / (10^(decimal1-decimal0))
price1 = 1 / price0
```

### Method 2: ticks

This method requires a different set of inputs to calculate.

#### Motivation

Ticks are related directly to price and enable simpler calculations compared to [[#Method 1]].
This requires just a tick value and a one-shot formula.

#### How to calculate price?

To convert from tick t, to price, take 1.0001^t to get the corresponding price.

```
price  = 1.0001^tick
price0 = price / (10^(decimal1-decimal0))
price1 = 1 / price0
```

It's also possible to convert sqrtPriceX96 to tick:
```
tick = floor[ log((sqrtPriceX96 / (2^96))^2) / log(1.0001) ]
```

## TODO:

- [x] remove Finex dependencies
- [ ] add specs or amendments to current interface
