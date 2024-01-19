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

var logger = log.Logger("uniswap")

type uniswapV3 struct {
	once       *once
	url        string
	outbox     chan<- TradeEvent
	windowSize time.Duration
	streams    sync.Map
}

func newUniswapV3(config Config, outbox chan<- TradeEvent) *uniswapV3 {
	url := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"
	if config.URL != "" {
		url = config.URL
	}

	return &uniswapV3{
		once:       newOnce(),
		url:        url,
		outbox:     outbox,
		windowSize: 2 * time.Second,
	}
}

func (u *uniswapV3) Start() error {
	u.once.Start(func() {})
	return nil
}

func (u *uniswapV3) Stop() error {
	u.once.Stop(func() {
		u.streams.Range(func(market, stream any) bool {
			stopCh := stream.(chan struct{})
			stopCh <- struct{}{}
			close(stopCh)
			return true
		})

		u.streams = sync.Map{} // delete all stopped streams
	})
	return nil
}

func (u *uniswapV3) Subscribe(market Market) error {
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
		for {
			if stream, ok := u.streams.Load(market); ok {
				stopCh := stream.(chan struct{})
				select {
				case <-stopCh:
					loggerBinance.Infof("market %s is stopped", symbol)
					return
				default:
				}
			}

			select {
			case <-time.After(u.windowSize):
				to := from.Add(u.windowSize)
				swaps, err := u.fetchSwaps(market, from, to)
				if err != nil {
					err = fmt.Errorf("%s: %w", market, err)
					loggerBinance.Warn(err)
				}

				for _, swap := range swaps {
					price := swap.price()
					createdAt, err := swap.time()
					if err != nil {
						loggerBinance.Warnf("failed to get swap timestamp: %s", err)
					}

					u.outbox <- TradeEvent{
						Source:    DriverUniswapV3,
						Market:    market.QuoteUnit,
						Price:     price,
						Amount:    swap.Amount0,
						Total:     price.Mul(swap.Amount0),
						TakerType: TakerTypeSell,
						CreatedAt: createdAt,
					}
				}
			default:
			}
		}
	}()

	return nil
}

func (u *uniswapV3) Unsubscribe(market Market) error {
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

const tokenTemplate = `query {
 pools(where: {token0_: {symbol:"%s"}, token1_: {symbol:"%s"}}) {
  token0 { symbol }
  token1 { symbol }
 }
}`

func (u *uniswapV3) isMarketAvailable(market Market) (bool, error) {
	query := fmt.Sprintf(tokenTemplate,
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	pools, err := runUniswapV3GraphqlRequest[uniswapV3Pools](u.url, query)
	if err != nil {
		return false, err
	}
	return len(pools.Pools) > 0, nil
}

// NOTE: GraphQL Query is used here because
// Uniswap V3 does not support Subscriptions
const swapsTemplate = `query {
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

func (u *uniswapV3) fetchSwaps(market Market, from, to time.Time) ([]uniswapV3Swap, error) {
	query := fmt.Sprintf(swapsTemplate,
		from.Unix(),
		to.Unix(),
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	swaps, err := runUniswapV3GraphqlRequest[uniswapV3Swaps](u.url, query)
	if err != nil {
		return nil, err
	}
	return swaps.Swaps, nil
}

func runUniswapV3GraphqlRequest[T uniswapV3Pools | uniswapV3Swaps](url, query string) (*T, error) {
	requestBody, err := json.Marshal(uniswapV3GraphqlRequest{Query: query})
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
			loggerBinance.Errorf("error closing HTTP response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var parsed uniswapV3GraphqlResponse[T]
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &parsed.Data, nil
}

type uniswapV3GraphqlRequest struct {
	Query string `json:"query"`
}

type uniswapV3GraphqlResponse[T any] struct {
	Data   T      `json:"data"`
	Errors string `json:"errors"`
}

type uniswapV3Pools struct {
	Pools []uniswapV3Market `json:"pools"`
}

type uniswapV3Market struct {
	Token0 struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		Symbol string `json:"symbol"`
	} `json:"token"`
}

type uniswapV3Swaps struct {
	Swaps []uniswapV3Swap `json:"swaps"`
}

type uniswapV3Swap struct {
	Timestamp    string          `json:"timestamp"`
	Token0       token           `json:"token0"`
	Token1       token           `json:"token1"`
	Amount0      decimal.Decimal `json:"amount0"`
	Amount1      decimal.Decimal `json:"amount1"`
	SqrtPriceX96 decimal.Decimal `json:"sqrtPriceX96"`
}

var priceX96 = decimal.NewFromInt(2).Pow(decimal.NewFromInt(96))

// price method calculates the price per token at which the swap was performed.
// General formula is as follows: ((sqrtPriceX96 / 2**96)**2) / (10**decimal1 / 10**decimal0)
// See the math explained at https://blog.uniswap.org/uniswap-v3-math-primer
func (swap *uniswapV3Swap) price() decimal.Decimal {
	ten := decimal.NewFromInt(10)
	decimals := swap.Token1.Decimals.Sub(swap.Token0.Decimals)

	numerator := swap.SqrtPriceX96.Div(priceX96).Pow(decimal.NewFromInt(2))
	denominator := ten.Pow(decimals)
	return numerator.Div(denominator)
}

// time method parses string representation of Unix timestamp into a stdlib Time object.
func (swap *uniswapV3Swap) time() (time.Time, error) {
	unixTimestamp, err := strconv.ParseInt(swap.Timestamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(unixTimestamp, 0), nil
}

type token struct {
	Decimals decimal.Decimal `json:"decimals"`
}
