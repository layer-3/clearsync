package uniswap

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

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

var logger = log.Logger("uniswap")

type UniswapV3 struct {
	once       *common.Once
	url        string
	outbox     chan<- common.TradeEvent
	windowSize time.Duration
	streams    sync.Map
}

func NewUniswapV3(config common.Config, outbox chan<- common.TradeEvent) *UniswapV3 {
	url := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"
	if config.URL != "" {
		url = config.URL
	}

	return &UniswapV3{
		once:       common.NewOnce(),
		url:        url,
		outbox:     outbox,
		windowSize: 2 * time.Second,
	}
}

func (u *UniswapV3) Start() error {
	u.once.Start(func() {})
	return nil
}

func (u *UniswapV3) Stop() error {
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

func (u *UniswapV3) Subscribe(market common.Market) error {
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
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
					logger.Infof("market %s is stopped", symbol)
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
					logger.Warn(err)
				}

				for _, swap := range swaps {
					price := swap.price()
					createdAt, err := swap.time()
					if err != nil {
						logger.Warnf("failed to get swap timestamp: %s", err)
					}

					u.outbox <- common.TradeEvent{
						Source:    common.DriverUniswapV3,
						Market:    market.QuoteUnit,
						Price:     price,
						Amount:    swap.Amount0,
						Total:     price.Mul(swap.Amount0),
						TakerType: common.TakerTypeSell,
						CreatedAt: createdAt,
					}
				}
			default:
			}
		}
	}()

	return nil
}

func (u *UniswapV3) Unsubscribe(market common.Market) error {
	stream, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
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

func (u *UniswapV3) isMarketAvailable(market common.Market) (bool, error) {
	query := fmt.Sprintf(tokenTemplate,
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	pools, err := runGraphqlRequest[uniswapPools](u.url, query)
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

func (u *UniswapV3) fetchSwaps(market common.Market, from, to time.Time) ([]uniswapSwap, error) {
	query := fmt.Sprintf(swapsTemplate,
		from.Unix(),
		to.Unix(),
		strings.ToUpper(market.BaseUnit),
		strings.ToUpper(market.QuoteUnit),
	)

	swaps, err := runGraphqlRequest[uniswapSwaps](u.url, query)
	if err != nil {
		return nil, err
	}
	return swaps.Swaps, nil
}

func runGraphqlRequest[T uniswapPools | uniswapSwaps](url, query string) (*T, error) {
	requestBody, err := json.Marshal(graphqlRequest{Query: query})
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
			logger.Errorf("error closing HTTP response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var parsed graphqlResponse[T]
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &parsed.Data, nil
}

type graphqlRequest struct {
	Query string `json:"query"`
}

type graphqlResponse[T any] struct {
	Data   T      `json:"data"`
	Errors string `json:"errors"`
}

type uniswapPools struct {
	Pools []uniswapMarket `json:"pools"`
}

type uniswapMarket struct {
	Token0 struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		Symbol string `json:"symbol"`
	} `json:"token"`
}

type uniswapSwaps struct {
	Swaps []uniswapSwap `json:"swaps"`
}

type uniswapSwap struct {
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
func (swap *uniswapSwap) price() decimal.Decimal {
	ten := decimal.NewFromInt(10)
	decimals := swap.Token1.Decimals.Sub(swap.Token0.Decimals)

	numerator := swap.SqrtPriceX96.Div(priceX96).Pow(decimal.NewFromInt(2))
	denominator := ten.Pow(decimals)
	return numerator.Div(denominator)
}

// time method parses string representation of Unix timestamp into a stdlib Time object.
func (swap *uniswapSwap) time() (time.Time, error) {
	unixTimestamp, err := strconv.ParseInt(swap.Timestamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(unixTimestamp, 0), nil
}

type token struct {
	Decimals decimal.Decimal `json:"decimals"`
}
