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
index_price = NumEMA/DenEMA

NumEMA = EMA20((Volume)*(SourceMultiplier)*(Price))

DenEMA = EMA20((Volume)*(SourceMultiplier))

SourceMultiplier = (TradeSourceWeight/ActiveWeights))

ActiveWeights(market) - Selects all drivers having the market of the reeceived trade active, and returns the sum of their weights.
```

### Example flow

Drivers Weights Config: 
```
driverWeights = [{Binance: 20}, {Uniswap: 10}, {Kraken: 15}]
```

Trades from all the sources are coming into one queue.
IndexAggregator reads trades one by one, each time influencing EMA based on trade price, trade importance (volume), and source driver importance (weight):

1. Trade received:

```
Trade {
	Source: Binance
	Market: btcusdt
	Volume: 0.5
	Price: 44000
}
```
		
2. Select active drivers for the market. Lets say, we are receiving btcusdt trades only from Binance and Uniswap.
```	
ActiveWeights(btcusdt) = {driverWeights[Binance] + driverWeights[Uniswap] = 20 + 10} = 30
```
3. Select the driver weight
```	
tradeSourceWeight = driverWeights[Binance] = 20
```
4. Select Previous value of numEMA and denEMA. If there is no initial data (on the startup), the values of the first received trade are used to set the initial EMA.
```
prevNumEMA, prevDenEMA = cache.Get(btcusdt)

if cache.IsEmpty() {
	prevNumEMA = Trade.Volume * sourceMultiplier * Trade.Price
	prevNumEMA = Trade.Volume * sourceMultiplier
}
```
5. Calculate new numEMA and denEMA.
```
sourceMultiplier = tradeSourceWeight/ActiveWeights
numEMA = EMA20(prevNumEMA, (Trade.Volume * sourceMultiplier * Trade.Price))
denEMA = EMA20(prevDenEMA, (Trade.Volume * sourceMultiplier))
```
6. Store new EMAs in the cache

7. Return Index Price
```
return index_price = numEMA/denEMA
```


**IMPORTANT!** You should understand how EMA20 function works. It calculates exponential moving average price of last 20 trades, but it doesnt use last 20 trades for calculation. Like a hash function, it takes the previously saved EMA, and influences it with the price and importance of an incoming trade.

[Exponential Moving Average](https://www.investopedia.com/terms/e/ema.asp)

[Weighted Exponential Moving Average](https://www.financialwisdomforum.org/gummy-stuff/EMA.htm)

### Calculation example

Let's assume 5 trades have been received:

```
{source, price, amount}
{binance, 41000, 0.3}
{binance, 42500, 0.5}
{uniswap, 55000, 0.6}
{uniswap, 50000, 0.4}
{binance, 40000, 1.0}
```

Example 1:

Drivers Weights Config: 
```
driverWeights = [{Binance: 2}, {Uniswap: 2}]
```

```Weighted EMA price change with each trade: 41000, 41223.8806, 43500.32787, 43873.32054, 43343.1976```

Example 2:

driverWeights = [{Binance: 2}, {Uniswap: 0}]

```Weighted EMA price change with each trade: 41000, 41223.8806, 41223.8806, 41223.8806, 40872.3039```

In the third example, we are setting the amount of the first 4 trades to ```1```, and the last trade to ```10```, and this time the price increases by 2000 with each trade:

```
{source, price, amount}
{binance, 40000, 1.0}
{binance, 42000, 1.0}
{binance, 44000, 1.0}
{binance, 46000, 1.0}
{binance, 48000, 10.0}
```

```Weighted EMA price change with each trade: 40000, 40190.47619, 40553.28798, 41072.02246, 44624.83145```

We observe the last trade with ten times the amount has influenced the final EMA more significantly than all the other trades.

**Described test scenarios can be found in ```index_test.go```.**

*PriceCache*

```
type PriceCache interface {
	Get(market string) (decimal.Decimal, decimal.Decimal)                         // Returns numEMA and denEMA EMAs for a market
	Update(driver DriverType, market string, numEMA, denEMA decimal.Decimal) // Updates numEMA and denEMA EMAs for a market with a new value
	ActiveWeights(market string) decimal.Decimal                                  // Returns the sum of active driver weights for the market.
}
```

- PriceCache stores previous numEMA and denEMA EMAs.
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
