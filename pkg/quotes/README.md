# Quotes service

The Quotes Service is a centralized system responsible for recording and disseminating real-time trade data 
from various trading platforms.

Available drivers:
- Index Price
- Binance
- Kraken
- Opendax
- Bitfaker
- Uniswap v3
  - based on Subgraph API
  - based on go-ethereum
- Syncswap

## Interface to connect

```go
type Driver interface {
	Start() error
	Stop() error
	Subscribe(market Market) error
	Unsubscribe(market Market) error
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

`Index Price` is used mainly in risk management for portfolio evaluation and is aggregated from multiple sources.

Trades from all sources are coming into one queue. IndexAggregator reads trades sequentially, calling calculateIndexPrice() method from a configured strategy.

> You can leverage your index price calculation strategy by implementing the priceCalculator interface and passing it as an argument to the IndexAggregator constructor.

```go
type priceCalculator interface {
	calculateIndexPrice(trade TradeEvent) (decimal.Decimal, bool)
}
```

The default strategy for index price calculation is Volume Weighted Average Price (VWAP), additionally weighted by trade source importance.

```go
sourceMultiplier = (tradeSourceWeight/(activeSourcesWeightsSum(market)))

var totalPriceVolume, totalVolume num
for trade in range(N trades(market)) {
		totalPriceVolume += (Price * Volume * sourceMultiplier)
		totalVolume += (Volume * sourceMultiplier)
}

index_price = totalPriceVolume / totalVolume
```

In the VWAP strategy `Price cache` is used to store a queue containing the last N (default=20) trades for each market.

> Drivers may support different markets. That's why the price cache additionally stores active drivers for each market. By default, no drivers are active. When a trade is added, the source of the trade becomes active for a market.

### Calculation flow

Drivers weights config example: 
```go
driverWeights = [{Binance: 20}, {Uniswap: 10}, {Kraken: 15}]
```


Trade received:

```go
Trade {
	Source: Binance
	Market: btcusdt
	Volume: 0.5
	Price: 44000
}
```

The trade is skipped if the trade price or volume is zero.

1. Calculate sourceMultiplier. Select active drivers for the market. By default, all drivers are not active. When the first trade from a driver is received, a market becomes active. 
Let's say, we are receiving btcusdt trades only from Binance and Uniswap.
```go
activeSourcesWeightsSum(btcusdt) = {driversWeights[Binance] + driversWeights[Uniswap] = 2 + 1} = 3
tradeSourceWeight = driverWeights[Binance] = 2

sourceMultiplier = (tradeSourceWeight/(activeSourcesWeightsSum(market)))
```
2. Add trade data to the price cache.
```go
priceCache.AddTrade(event.Market, event.Price, event.Amount, sourceMultiplier)
```
3. Fetch the last n trades from PriceCache.
```go
var totalPriceVolume, totalVolume num
for trade in range(priceCache.trades(btcusdt)) {
		totalPriceVolume += (Price * Volume * sourceMultiplier)
		totalVolume += (Volume * sourceMultiplier)
}
```
4. Return Index Price
```go
index_price = totalPriceVolume / totalVolume 
```


**VWA20 Index Price Example Over 20 trades**

|        | 1     | 2     | 3     | 4     | 5     | 6     | 7     | 8     | 9     | 10    | 11    | 12    | 13    | 14    | 15    | 16    | 17    | 18    | 19    | 20    |
|--------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| Amount | 1.0   | 1.1   | 2.4   | 0.2   | 2.0   | 3.3   | 4.0   | 2.9   | 1.0   | 0.1   | 0.01  | 0.04  | 9.0   | 0.4   | 4.4   | 5.0   | 6.0   | 0.1   | 2.0   | 1.0   |
| Price  | 40000 | 42000 | 41500 | 44000 | 43000 | 40000 | 41000 | 42000 | 43000 | 42000 | 45500 | 41000 | 41500 | 42000 | 44000 | 46000 | 47000 | 46000 | 44000 | 42000 |
| VWA20  | 40000 | 41047 | 41288 | 41404 | 41880 | 41260 | 41185 | 41325 | 41418 | 41422 | 41424 | 41423 | 41448 | 41457 | 41808 | 42377 | 43024 | 43031 | 43074 | 43051 |

> More examples and test scenarios can be found in `index_test.go`.

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

- [ ] add specs or amendments to current interface
