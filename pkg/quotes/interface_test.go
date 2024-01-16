package quotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/binance"
	"github.com/layer-3/clearsync/pkg/quotes/bitfaker"
	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/kraken"
	"github.com/layer-3/clearsync/pkg/quotes/opendax"
)

func TestNewDriver(t *testing.T) {
	t.Parallel()

	t.Run(common.DriverBinance.String(), func(t *testing.T) {
		t.Parallel()

		config := common.Config{Driver: common.DriverBinance}
		outbox := make(chan<- common.TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*binance.Binance)
		assert.True(t, ok)
	})

	t.Run(common.DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		config := common.Config{Driver: common.DriverKraken}
		outbox := make(chan<- common.TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*kraken.Kraken)
		assert.True(t, ok)
	})

	t.Run(common.DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		config := common.Config{Driver: common.DriverBitfaker}
		outbox := make(chan<- common.TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*bitfaker.Bitfaker)
		assert.True(t, ok)
	})

	t.Run(common.DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		config := common.Config{Driver: common.DriverOpendax}
		outbox := make(chan<- common.TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*opendax.Opendax)
		assert.True(t, ok)
	})
}
