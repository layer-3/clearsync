package quotes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

var loggerUniswapV3Api = log.Logger("uniswap_v3_api")

type uniswapV3Api struct {
	once       *once
	url        string
	outbox     chan<- TradeEvent
	windowSize time.Duration
	streams    safe.Map[Market, chan struct{}]
}

func newUniswapV3Api(config UniswapV3ApiConfig, outbox chan<- TradeEvent) Driver {
	return &uniswapV3Api{
		once:       newOnce(),
		url:        config.URL,
		outbox:     outbox,
		windowSize: config.WindowSize,
		streams:    safe.NewMap[Market, chan struct{}](),
	}
}

func (u *uniswapV3Api) Type() DriverType {
	return DriverUniswapV3Api
}

func (u *uniswapV3Api) Start() error {
	if started := u.once.Start(func() {}); !started {
		return errAlreadyStarted
	}
	return nil
}

func (u *uniswapV3Api) Stop() error {
	stopped := u.once.Stop(func() {
		u.streams.Range(func(_ Market, stopCh chan struct{}) bool {
			stopCh <- struct{}{}
			close(stopCh)
			return true
		})

		u.streams = safe.NewMap[Market, chan struct{}]() // delete all stopped streams
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (u *uniswapV3Api) Subscribe(market Market) error {
	if !u.once.Subscribe() {
		return errNotStarted
	}

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	exists, err := u.isMarketAvailable(market)
	if err != nil {
		return fmt.Errorf("failed to check if market %s exists: %s", market.String(), err)
	}
	if !exists {
		return fmt.Errorf("market %s does not exist", market.String())
	}

	u.streams.Store(market, make(chan struct{}, 1))

	go func() {
		from := time.Now()
		timer := time.After(u.windowSize)
		for {
			if stopCh, ok := u.streams.Load(market); ok {
				select {
				case <-stopCh:
					loggerUniswapV3Api.Infof("market %s is stopped", market.String())
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
					loggerUniswapV3Api.Warn(err)
				}

				for _, swap := range swaps {
					amount := swap.Amount0.Abs()
					price := calculatePrice(swap.SqrtPriceX96, swap.Token0.Decimals, swap.Token1.Decimals)
					createdAt, err := swap.time()
					if err != nil {
						loggerUniswapV3Api.Warnf("failed to get swap timestamp: %s", err)
					}
					takerType := TakerTypeBuy
					if swap.Amount0.Sign() < 0 {
						// When amount0 is negative (and amount1 is positive),
						// it means token0 is leaving the pool in exchange for token1.
						// This is equivalent to a "sell" of token0 (or a "buy" of token1).
						takerType = TakerTypeSell
					}

					u.outbox <- TradeEvent{
						Source:    DriverUniswapV3Api,
						Market:    market,
						Price:     price,
						Amount:    amount,
						Total:     price.Mul(amount),
						TakerType: takerType,
						CreatedAt: createdAt,
					}
				}
				timer = time.After(u.windowSize)
				from = to
			default:
			}
		}
	}()

	return nil
}

func (u *uniswapV3Api) Unsubscribe(market Market) error {
	if !u.once.Unsubscribe() {
		return errNotStarted
	}

	stopCh, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

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

func (u *uniswapV3Api) isMarketAvailable(market Market) (bool, error) {
	query := fmt.Sprintf(tokenTemplate,
		strings.ToUpper(market.Base()),
		strings.ToUpper(market.Quote()),
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

func (u *uniswapV3Api) fetchSwaps(market Market, from, to time.Time) ([]uniswapV3Swap, error) {
	query := fmt.Sprintf(swapsTemplate,
		from.Unix(),
		to.Unix(),
		strings.ToUpper(market.Base()),
		strings.ToUpper(market.Quote()),
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
			loggerUniswapV3Api.Errorf("error closing HTTP response body: %v", err)
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
	Token0       graphqlToken    `json:"token0"`
	Token1       graphqlToken    `json:"token1"`
	Amount0      decimal.Decimal `json:"amount0"`
	Amount1      decimal.Decimal `json:"amount1"`
	SqrtPriceX96 decimal.Decimal `json:"sqrtPriceX96"`
}

var priceX96 = decimal.NewFromInt(2).Pow(decimal.NewFromInt(96))
var ten = decimal.NewFromInt(10)

// calculatePrice method calculates the price per token at which the swap was performed.
// General formula is as follows: ((sqrtPriceX96 / 2**96)**2) / (10**decimal1 / 10**decimal0)
// See the math explained at https://blog.uniswap.org/uniswap-v3-math-primer
func calculatePrice(sqrtPriceX96, baseTokenDecimals, quoteTokenDecimals decimal.Decimal) decimal.Decimal {
	decimals := quoteTokenDecimals.Sub(baseTokenDecimals)

	numerator := sqrtPriceX96.Div(priceX96).Pow(decimal.NewFromInt(2))
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

type graphqlToken struct {
	Decimals decimal.Decimal `json:"decimals"`
}
