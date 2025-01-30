package driver

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestBinance_HistoricalData(t *testing.T) {
	config, err := NewConfigFromEnv()
	require.NoError(t, err)
	outbox := make(chan common.TradeEvent, 128)

	driver, err := newBinance(config.Binance, outbox, nil)
	binance := driver.(*binance)
	require.NoError(t, err)

	market := common.NewMarket("btc", "usdt")
	window := 15 * time.Minute
	limit := uint64(500)

	trades, err := binance.HistoricalData(context.Background(), market, window, limit)
	require.NoError(t, err)
	require.NotEmpty(t, trades)
}
