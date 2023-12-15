package portfolio

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAddTrade(t *testing.T) {
	p := &Portfolio{
		Positions:   make(map[MarketSymbol]Position),
		Prices:      make(map[MarketSymbol]decimal.Decimal),
		Exposure: decimal.Zero,
	}

	trade := Trade{
		Symbol:    "BTCUSD",
		Amount:    decimal.NewFromInt(2),
		Price:     decimal.NewFromInt(50000),
		Direction: Long,
	}
	p.AddTrade(trade)

	if p.Positions["BTCUSD"].Amount.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("Expected amount to be 2, got %s", p.Positions["BTCUSD"].Amount.String())
	}

	if p.Positions["BTCUSD"].Spent.Cmp(decimal.NewFromInt(100000)) != 0 {
		t.Errorf("Expected spent to be 100000, got %s", p.Positions["BTCUSD"].Spent.String())
	}
}

func TestUpdatePrice(t *testing.T) {
	p := &Portfolio{
		Positions:   make(map[MarketSymbol]Position),
		Prices:      make(map[MarketSymbol]decimal.Decimal),
		Exposure: decimal.Zero,
	}

	p.Positions["BTCUSD"] = Position{
		Amount:    decimal.NewFromInt(2),
		Spent:     decimal.NewFromInt(100000),
		Direction: Long,
	}

	quote := Quote{
		Symbol: "BTCUSD",
		Price:  decimal.NewFromInt(51000),
	}
	p.UpdatePrice(quote)

	if p.Prices["BTCUSD"].Cmp(decimal.NewFromInt(51000)) != 0 {
		t.Errorf("Expected price to be 51000, got %s", p.Prices["BTCUSD"].String())
	}
}

func TestNetExposure(t *testing.T) {
	p := &Portfolio{
		Positions:   make(map[MarketSymbol]Position),
		Prices:      make(map[MarketSymbol]decimal.Decimal),
		Exposure: decimal.Zero,
	}

	p.Positions["BTCUSD"] = Position{
		Amount:    decimal.NewFromInt(2),
		Spent:     decimal.NewFromInt(100000),
		Direction: Long,
	}
	p.Prices["BTCUSD"] = decimal.NewFromInt(51000)
	p.calculateNetExposure("BTCUSD")

	netExposure := p.NetExposure()

	expectedExposure := decimal.NewFromInt(2000) // 2 * (51000 - 50000)
	if netExposure.Cmp(expectedExposure) != 0 {
		t.Errorf("Expected net exposure to be 2000, got %s", netExposure.String())
	}
}

