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

## TODO:
- remove Finex dependencies
- add specs or amendments to current interface
