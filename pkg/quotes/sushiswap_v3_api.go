package quotes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
)

var loggerSushiswapV3Api = log.Logger("sushiswap_v3_api")

type sushiswapV3Api struct {
	once         *once
	url          string
	outbox       chan<- TradeEvent
	windowSize   time.Duration
	streams      sync.Map
	tradeSampler tradeSampler
}

func newSushiswapV3Api(config SushiswapV3ApiConfig, outbox chan<- TradeEvent) *sushiswapV3Api {
	return &sushiswapV3Api{
		once:         newOnce(),
		url:          config.URL,
		outbox:       outbox,
		windowSize:   config.WindowSize,
		tradeSampler: *newTradeSampler(config.TradeSampler),
	}
}

func (u *sushiswapV3Api) Start() error {
	if started := u.once.Start(func() {}); !started {
		return errAlreadyStarted
	}
	return nil
}

func (u *sushiswapV3Api) Stop() error {
	stopped := u.once.Stop(func() {
		u.streams.Range(func(market, stream any) bool {
			stopCh := stream.(chan struct{})
			stopCh <- struct{}{}
			close(stopCh)
			return true
		})

		u.streams = sync.Map{} // delete all stopped streams
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (u *sushiswapV3Api) Subscribe(market Market) error {
	if !u.once.Subscribe() {
		return errNotStarted
	}
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	exists, err := u.isMarketAvailable(market)
	if err != nil {
		return fmt.Errorf("failed to check if market %s exists: %s", symbol, err)
	}
	if !exists {
		return fmt.Errorf("market %s does not exist", symbol)
	}

	u.streams.Store(market, make(chan struct{}, 1))

	go func() {
		from := time.Now()
		timer := time.After(u.windowSize)
		for {
			if stream, ok := u.streams.Load(market); ok {
				stopCh := stream.(chan struct{})
				select {
				case <-stopCh:
					loggerSushiswapV3Api.Infof("market %s is stopped", symbol)
					return
				default:
				}
			}

			select {
			case <-timer:
				to := from.Add(u.windowSize)
				swaps, err := u.fetchSwaps(market, from, to)
				if err != nil {
					err = fmt.Errorf("%s: %w", market, err)
					loggerSushiswapV3Api.Warn(err)
				}

				for _, swap := range swaps {
					amount := swap.Amount0.Abs()
					price := calculatePrice(swap.SqrtPriceX96, swap.Token0.Decimals, swap.Token1.Decimals)
					createdAt, err := swap.time()
					if err != nil {
						loggerSushiswapV3Api.Warnf("failed to get swap timestamp: %s", err)
					}
					takerType := TakerTypeBuy
					if swap.Amount0.Sign() < 0 {
						// When amount0 is negative (and amount1 is positive),
						// it means token0 is leaving the pool in exchange for token1.
						// This is equivalent to a "sell" of token0 (or a "buy" of token1).
						takerType = TakerTypeSell
					}

					tr := TradeEvent{
						Source:    DriverSushiswapV3Api,
						Market:    market.QuoteUnit,
						Price:     price,
						Amount:    amount,
						Total:     price.Mul(amount),
						TakerType: takerType,
						CreatedAt: createdAt,
					}

					if !u.tradeSampler.allow(tr) {
						continue
					}
					u.outbox <- tr
				}
				timer = time.After(u.windowSize)
				from = to
			default:
			}
		}
	}()

	return nil
}

func (u *sushiswapV3Api) Unsubscribe(market Market) error {
	if !u.once.Unsubscribe() {
		return errNotStarted
	}

	stream, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	stopCh := stream.(chan struct{})
	stopCh <- struct{}{}
	close(stopCh)

	u.streams.Delete(market)
	return nil
}

const sushiswapV3ApiTokenTemplate = `query {
  pools(where: {token0_: {symbol:"%s"}, token1_: {symbol:"%s"}}) {
    token0 { symbol }
    token1 { symbol }
  }
}`

func (u *sushiswapV3Api) isMarketAvailable(market Market) (bool, error) {
	query := fmt.Sprintf(sushiswapV3ApiTokenTemplate,
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	pools, err := runSushiswapV3GraphqlRequest[sushiswapV3Pools](u.url, query)
	if err != nil {
		return false, err
	}
	return len(pools.Pools) > 0, nil
}

// NOTE: Query is used here because
// Sushiswap V3 GraphQL API does not support Subscriptions
const sushiswapV3ApiSwapsTemplate = `query {
  swaps(
    orderBy: timestamp
    orderDirection: desc
    where: {
      timestamp_gte: %d
      timestamp_lt: %d
      token0_: {symbol: "%s"}
      token1_: {symbol: "%s"}
    }
  ) {
    timestamp
    token0 { decimals }
    token1 { decimals }
    amount0
    amount1
    sqrtPriceX96
  }
}`

func (u *sushiswapV3Api) fetchSwaps(market Market, from, to time.Time) ([]sushiswapV3Swap, error) {
	query := fmt.Sprintf(sushiswapV3ApiSwapsTemplate,
		from.Unix(),
		to.Unix(),
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	swaps, err := runSushiswapV3GraphqlRequest[sushiswapV3Swaps](u.url, query)
	if err != nil {
		return nil, err
	}
	return swaps.Swaps, nil
}

func runSushiswapV3GraphqlRequest[T sushiswapV3Pools | sushiswapV3Swaps](url, query string) (*T, error) {
	requestBody, err := json.Marshal(sushiswapV3GraphqlRequest{Query: query})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to request data: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			loggerSushiswapV3Api.Errorf("error closing HTTP response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var parsed sushiswapV3GraphqlResponse[T]
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w (body: `%s`)", err, string(body))
	}

	return &parsed.Data, nil
}

type sushiswapV3GraphqlRequest struct {
	Query string `json:"query"`
}

type sushiswapV3GraphqlResponse[T any] struct {
	Data   T      `json:"data"`
	Errors string `json:"errors"`
}

type sushiswapV3Pools struct {
	Pools []sushiswapV3Market `json:"pools"`
}

type sushiswapV3Market struct {
	Token0 struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		Symbol string `json:"symbol"`
	} `json:"token"`
}

type sushiswapV3Swaps struct {
	Swaps []sushiswapV3Swap `json:"swaps"`
}

type sushiswapV3Swap struct {
	Timestamp    string                     `json:"timestamp"`
	Token0       sushiswapV3ApiGraphqlToken `json:"token0"`
	Token1       sushiswapV3ApiGraphqlToken `json:"token1"`
	Amount0      decimal.Decimal            `json:"amount0"`
	Amount1      decimal.Decimal            `json:"amount1"`
	SqrtPriceX96 decimal.Decimal            `json:"sqrtPriceX96"`
}

// time method parses string representation of Unix timestamp into a stdlib Time object.
func (swap *sushiswapV3Swap) time() (time.Time, error) {
	unixTimestamp, err := strconv.ParseInt(swap.Timestamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(unixTimestamp, 0), nil
}

type sushiswapV3ApiGraphqlToken struct {
	Decimals decimal.Decimal `json:"decimals"`
}
