package quotes

import (
	"testing"

	"github.com/shopspring/decimal"
)

// TODO: add tests and make sure everything works perfectly

func TestOpendax_EMA(t *testing.T) {
	// TODO: test the EMA formula
	t.Run("Successful test", func(t *testing.T) {
		_ = EMA(decimal.Zero, decimal.Zero, 5)
	})
}
