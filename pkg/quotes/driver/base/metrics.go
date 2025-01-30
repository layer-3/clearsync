package base

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

var MarketSubscriptions = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "price_feed_value",
		Help: "Current trades subscriptions by provider and market.",
	},
	[]string{"provider", "market"}, // labels
)

func RecordSubscribed(provider common.DriverType, market common.Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Inc()
}

func RecordUnsubscribed(provider common.DriverType, market common.Market) {
	MarketSubscriptions.WithLabelValues(provider.String(), market.String()).Dec()
}
