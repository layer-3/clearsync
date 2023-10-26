# Quotes service

The Quotes Service is a centralized system responsible for recording and disseminating real-time trade data 
from various trading platforms.

Available drivers:
- binance
- kraken
- opendax
- bitfaker

## Interface to connect

```
type Driver interface {
	Init(markets cache.Market, outbox chan trade.Event, output chan<- event.Event, config config.QuoteFeed, dialer client.WSDialer) error
	Start() error
	Subscribe(base, quote string) error
	Close() error
}
```

## Type of price source

| Top Tier CEX | Altcoins CEX | FIAT       | DEX         |
| ------------ | ------------ | ---------- | ----------- |
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



## TODO:

- remove Finex dependencies
- add specs or amendments to current interface
