package quotes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/precision"
)

type uniswapV3 struct {
	url        string
	outbox     chan<- TradeEvent
	windowSize time.Duration

	mu      sync.Mutex
	streams map[Market]chan struct{}
}

func newUniswapV3(config Config, outbox chan<- TradeEvent) *uniswapV3 {
	url := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"
	if config.URL != "" {
		url = config.URL
	}

	return &uniswapV3{
		url:        url,
		outbox:     outbox,
		windowSize: 2 * time.Second,

		streams: make(map[Market]chan struct{}),
	}
}

func (u *uniswapV3) Subscribe(market Market) error {
	symbol := market.BaseUnit + market.QuoteUnit
	exists, err := u.isMarketAvailable(market)
	if err != nil {
		return fmt.Errorf("failed to check if market %s exists: %s", symbol, err)
	}
	if !exists {
		return fmt.Errorf("market %s does not exist", symbol)
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	u.streams[market] = make(chan struct{}, 1)

	go func() {
		from := time.Now()
		for {
			select {
			case <-u.streams[market]:
				logger.Infof("market %s is stopped", symbol)
				return
			case <-time.After(u.windowSize):
				to := from.Add(u.windowSize)
				swaps, err := u.fetchSwaps(market, from, to)
				if err != nil {
					logger.Warn("failed to fetch swaps")
				}

				for _, swap := range swaps {
					price := precision.ToSignificant(swap.price(), 8)
					createdAt, err := swap.time()
					if err != nil {
						logger.Warnf("failed to get swap timestamp: %s", err)
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
			}
		}
	}()

	return nil
}

func (u *uniswapV3) Start(markets []Market) error {
	if len(markets) == 0 {
		return errors.New("no markets specified")
	}

	for _, m := range markets {
		m := m
		go func() {
			err := u.Subscribe(m)
			if err != nil {
				symbol := m.BaseUnit + m.QuoteUnit
				logger.Warnf("failed to subscribe to market %s: %s", symbol, err)
			}
		}()
	}

	return nil
}

func (u *uniswapV3) Stop() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, stopCh := range u.streams {
		stopCh <- struct{}{}
		close(stopCh)
	}

	u.streams = make(map[Market]chan struct{}) // delete all stopped streams
	return nil
}

const tokenTemplate = `query {
 pools(where: {token0_: {symbol:"%s"}, token1_: {symbol:"%s"}}) {
  token0 { symbol }
  token1 { symbol }
 }
}`

func (u *uniswapV3) isMarketAvailable(m Market) (bool, error) {
	query := fmt.Sprintf(tokenTemplate,
		strings.ToUpper(m.BaseUnit),
		strings.ToUpper(m.QuoteUnit),
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

func (u *uniswapV3) fetchSwaps(m Market, from, to time.Time) ([]uniswapSwap, error) {
	query := fmt.Sprintf(swapsTemplate,
		from.Unix(),
		to.Unix(),
		strings.ToUpper(m.BaseUnit),
		strings.ToUpper(m.QuoteUnit),
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

	defer resp.Body.Close()
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
