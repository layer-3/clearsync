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
  - Subgraph API
  - go-ethereum

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

Index Price is used mainly in risk management for portfolio evaluation and is aggregated from multiple sources.

```
index_price = NumEMA/DenEMA

NumEMA = EMA20((Volume)*(SourceMultiplier)*(Price))

DenEMA = EMA20((Volume)*(SourceMultiplier))

SourceMultiplier = (TradeSourceWeight/ActiveWeights))

ActiveWeights(market) - Selects all drivers having the market of the received trade active, and returns the sum of their weights.
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

If trade price or amount is zero, the trade is getting skipped.

2. Select active drivers for the market. By default, all drivers are not active. When the first trade from a driver is received, a market becomes active. 
Let's say, we are receiving btcusdt trades only from Binance and Uniswap.
```	
ActiveWeights(btcusdt) = {driverWeights[Binance] + driverWeights[Uniswap] = 20 + 10} = 30
```
1. Select the driver's weight
```	
tradeSourceWeight = driverWeights[Binance] = 20
```
1. Select the previous value of numEMA and denEMA. If there is no initial data (on the startup), the values of the first received trade are used to set the initial EMA.
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
index_price = numEMA/denEMA
```

**IMPORTANT!** You should understand how EMA20 function works. It calculates the exponential moving average price of the last 20 trades, but it doesn't use the last 20 trades for calculation. Similar to a hash function, it takes the previous EMA and influences it with the price and importance of an incoming trade.

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

**Example 1:**

Drivers Weights Config: 
```
driverWeights = [{Binance: 2}, {Uniswap: 2}]
```
**Trade 1 received ```{binance, 41000, 0.3}```:**
```
activeWeights(btcusdt) = driverWeights[Binance] = 2
tradeSourceWeight = driverWeights[Binance] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 1

<!-- Cache is empty. Setting initial values -->
prevNumEMA = event.Price.Mul(event.Amount).Mul(sourceMultiplier)
prevDenEMA = event.Amount.Mul(sourceWeight).Div(activeWeights)

prevNumEMA = 41000*0.3*1 = 12300
prevDenEMA = 0.3*1

NumEMA = EMA20((Volume)*(SourceMultiplier)*(Price))
DenEMA = EMA20((Volume)*(SourceMultiplier))

numEMA = EMA20(prevNumEMA, (Trade.Volume * sourceMultiplier * Trade.Price))
denEMA = EMA20(prevDenEMA, (Trade.Volume * sourceMultiplier))

<!-- EMA calculation -->
smoothingFactor := 2/(intervals+1) = 2/21 = 0.095
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((41000*0.3*1-12300) x 0.95) + 12300 = 12300
denEMA = ((0.3*1-0.3) x 0.95) + 0.3 = 0.3

<!-- Save values in cache -->
priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 12300/0.3 = 41000
```
**Trade 2 received ```{binance, 42500, 0.5}```:**
```
activeWeights(btcusdt) = driverWeights[Binance] = 2
tradeSourceWeight = driverWeights[Binance] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 1

<!-- Get previous values from cache -->
prevNumEMA = priceCache.Get(btcusdt) = 12300
prevDenEMA = priceCache.Get(btcusdt) = 0.3

<!-- EMA calculation -->
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((42500*0.5*1-12300) x 0.095) + 12300 = 13150.25
denEMA = ((0.5*1-0.3) x 0.095) + 0.3 = 0.319

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 13150.25/0.319 = 41223
```

**Trade 3 received ```{uniswap, 55000, 0.6}```:** 
```
activeWeights(btcusdt) = driverWeights[Binance] + driverWeights[Uniswap] = 4
tradeSourceWeight = driverWeights[Uniswap] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 2/4 = 0.5

prevNumEMA = priceCache.Get(btcusdt) = 13150.25
prevDenEMA = priceCache.Get(btcusdt) = 0.319

<!-- EMA calculation -->
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((55000*0.6*0.5-13150.25) x 0.095) + 13150.25 = 13468.476
denEMA = ((0.6*0.5-0.319) x 0.095) + 0.319 = 0.317

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 13468.476/0.317 = 42487
```

**Trade 4 received ```{uniswap, 50000, 0.4}```:** 
```
activeWeights(btcusdt) = driverWeights[Binance] + driverWeights[Uniswap] = 4
tradeSourceWeight = driverWeights[Uniswap] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 2/4 = 0.5

prevNumEMA = priceCache.Get(btcusdt) = 13468.476
prevDenEMA = priceCache.Get(btcusdt) = 0.317

<!-- EMA calculation -->
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((50000*0.4*0.5-13468.476) x 0.095) + 13468.476 = 13138.971
denEMA = ((0.4*0.5-0.317) x 0.095) + 0.317 = 0.3056

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 13138.971/0.3056 = 42994
```

**Trade 5 received ```{binance, 40000, 1.0}```:** 
```
activeWeights(btcusdt) = driverWeights[Binance] + driverWeights[Uniswap] = 4
tradeSourceWeight = driverWeights[Binance] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 2/4 = 0.5

prevNumEMA = priceCache.Get(btcusdt) = 13138.971
prevDenEMA = priceCache.Get(btcusdt) = 0.3056

<!-- EMA calculation -->
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((40000*1.0*0.5-13138.971) x 0.095) + 13138.971 = 13790.769
denEMA = ((1.0*0.5-0.3056) x 0.095) + 0.3056 = 0.324

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 13790.769/0.324 = 42564
```

```Weighted EMA price change with each trade: 41000, 41223, 42487, 42994, 42564```

**Example 2:**

```
driverWeights = [{Binance: 2}, {Uniswap: 0}]
```
```
<!-- The first two trades are the same as in the previous example. -->
```
**Trade 3 received ```{uniswap, 55000, 0.6}```:** 
```
activeWeights(btcusdt) = driverWeights[Binance] + driverWeights[Uniswap] = 2
tradeSourceWeight = driverWeights[Uniswap] = 0
sourceMultiplier = tradeSourceWeight/ActiveWeights = 0/2 = 0

NumEMA = EMA20((Volume)*(SourceMultiplier)*(Price))
DenEMA = EMA20((Volume)*(SourceMultiplier))

<!-- Get previous values from cache -->
prevNumEMA = priceCache.Get(btcusdt) = 13150.25
prevDenEMA = priceCache.Get(btcusdt) = 0.319

<!-- EMA calculation -->
EMA = ((newValue − previous EMA) × smoothingFactor) + (previous EMA).

numEMA = ((55000*0.6*0-13150.25) x 0.095) + 13150.25 = 11900.97625
denEMA = ((0.6*0-0.319) x 0.095) + 0.319 = 0.288695

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 11900.97625/0.288695 = 41223
```

**Trade 4 received ```{uniswap, 50000, 0.4}```:** 
```
sourceMultiplier = tradeSourceWeight/ActiveWeights = 0/2 = 0

prevNumEMA = priceCache.Get(btcusdt) = 11900.97625
prevDenEMA = priceCache.Get(btcusdt) = 0.288695

<!-- EMA calculation -->
numEMA = ((50000*0.6*0-11900.97625) x 0.095) + 11900.97625 = 10770.3835063
denEMA = ((0.4*0-0.288695) x 0.095) + 0.288695 = 0.261268975

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 10770.3835063/0.261268975 = 41223
```

**Trade 5 received ```{binance, 40000, 1.0}```:** 
```
activeWeights(btcusdt) = driverWeights[Binance] + driverWeights[Uniswap] = 2
tradeSourceWeight = driverWeights[Binance] = 2
sourceMultiplier = tradeSourceWeight/ActiveWeights = 2/2 = 1

prevNumEMA = priceCache.Get(btcusdt) = 10770.3835063
prevDenEMA = priceCache.Get(btcusdt) = 0.261268975

<!-- EMA calculation -->
numEMA = ((40000*1.0*1-10770.3835063) x 0.095) + 10770.3835063 = 13547.197
denEMA = ((1.0*1-0.261268975) x 0.095) + 0.261268975 = 0.332

priceCache.Update(numEMA, denEMA)

index_price = numEMA/denEMA = 13547.1970732/0.33144842237 = 40872
```

```Weighted EMA price change with each trade: 41000, 41223, 41223, 41223, 40872```

**Example 3:**

- Only Binance driver is active
- The first 4 trades amount is ```1```
- The 5-th trade amount is ```10```
- The price increases by 2000 with each trade:

```
{source, price, amount}
{binance, 40000, 1.0}
{binance, 42000, 1.0}
{binance, 44000, 1.0}
{binance, 46000, 1.0}
{binance, 48000, 10.0}
```

Calculations flow is the same as in the previous examples, ```sourceMultiplier``` stays at ```1``` because only the Binance driver is active.

```Weighted EMA price change with each trade: 40000, 40190, 40553, 41072, 44624```

The last trade with a larger amount has influenced the final price more significantly than all other trades.

**Example 4:**

- Initial price is 41000
- Two drivers are active
- These drivers are sequentially sending 40000 and 42000 trades with the same amount

In this case, the index price will get smoothed, alternating between 41050 and 40950 instead of the initial 40000 and 42000.
If the initial range were 38000 and 44000, the smoothed price range would also increase to 40850 and 41150.

- After that these drivers start sending trades with the same amount and the price of 41000.

In this case, the ingex price will stabilyse at 41000.

<!-- TEMP DRAFT -->
**EMA20 vs VWA20 comparison**
|          | Amount  | Price  |   EMA20   |   VWA20   |
|----------|---------|--------|-----------|-----------|
|    1     |   1.0   | 40000  |   40000   |   40000   |
|    2     |   1.1   | 42000  |   40207   |   41047   |
|    3     |   2.4   | 41500  |   40466   |   41288   |
|    4     |   0.2   | 44000  |   40530   |   41404   |
|    5     |   2.0   | 43000  |   40941   |   41880   |
|    6     |   3.3   | 40000  |   40722   |   41260   |
|    7     |   4.0   | 41000  |   40788   |   41185   |
|    8     |   2.9   | 42000  |   40982   |   41325   |
|    9     |   1.0   | 43000  |   41098   |   41418   |
|   10     |   0.1   | 42000  |   41104   |   41422   |
|   11     |  0.01   | 45500  |   41107   |   41424   |
|   12     |  0.04   | 41000  |   41107   |   41423   |
|   13     |   9.0   | 41500  |   41277   |   41448   |
|   14     |   0.4   | 42000  |   41292   |   41457   |
|   15     |   4.4   | 44000  |   41839   |   41808   |
|   16     |   5.0   | 46000  |   42682   |   42377   |
|   17     |   6.0   | 47000  |   43596   |   43024   |
|   18     |   0.1   | 46000  |   43605   |   43031   |
|   19     |   2.0   | 44000  |   43637   |   43074   |
|   20     |   1.0   | 42000  |   43568   |   43051   |

**Described test scenarios can be found in ```index_test.go```. The final price in the test examples may be different because of higher calculation precision.**
*PriceCache*

```
type PriceCache interface {
	Get(market string) (decimal.Decimal, decimal.Decimal)        // Returns numEMA and denEMA EMAs for a market
	Update(market string, priceWeight, weight decimal.Decimal)   // Updates numEMA and denEMA EMAs for a market with a new value
	ActivateDriver(driver DriverType, market string)             // ActivateDriver makes the driver active for the market.
	ActiveWeights(market string) decimal.Decimal                 // Returns the sum of active driver weights for the market.
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

- [ ] add specs or amendments to current interface
