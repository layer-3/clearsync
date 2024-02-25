package portfolio

import (
    "github.com/shopspring/decimal"
)

type Direction string

const (
    Long  Direction = "Long"
    Short Direction = "Short"
)

type Position struct {
    Amount    decimal.Decimal
    Spent     decimal.Decimal
    Direction Direction
}

type MarketSymbol string

type Trade struct {
    Symbol    MarketSymbol
    Amount    decimal.Decimal
    Price     decimal.Decimal
    Direction Direction
}

type Quote struct {
    Symbol MarketSymbol
    Price  decimal.Decimal
}

type Portfolio struct {
    Positions    map[MarketSymbol]Position
    Prices       map[MarketSymbol]decimal.Decimal
    Exposure  decimal.Decimal
}

// AddTrade adds a trade to the portfolio.
func (p *Portfolio) AddTrade(trade Trade) {
    // Initialize the position if it does not exist.
    if _, exists := p.Positions[trade.Symbol]; !exists {
        p.Positions[trade.Symbol] = Position{}
    }

    position := p.Positions[trade.Symbol]

    // Calculate new amount and spent based on trade direction.
    if trade.Direction == Long {
        position.Amount = position.Amount.Add(trade.Amount)
        position.Spent = position.Spent.Add(trade.Amount.Mul(trade.Price))
    } else { // Short
        position.Amount = position.Amount.Sub(trade.Amount)
        position.Spent = position.Spent.Sub(trade.Amount.Mul(trade.Price))
    }

    position.Direction = trade.Direction
    p.Positions[trade.Symbol] = position
}

// UpdatePrice updates the price of a symbol in the portfolio.
func (p *Portfolio) UpdatePrice(quote Quote) {
    p.Prices[quote.Symbol] = quote.Price
    p.calculateNetExposure(quote.Symbol)
}

// calculateNetExposure calculates the net exposure for a given market symbol.
func (p *Portfolio) calculateNetExposure(symbol MarketSymbol) {
    position, exists := p.Positions[symbol]
    if !exists {
        return
    }

    currentPrice := p.Prices[symbol]
    var netExposure decimal.Decimal

    if position.Direction == Long {
        netExposure = currentPrice.Mul(position.Amount).Sub(position.Spent)
    } else { // Short
        netExposure = position.Spent.Sub(currentPrice.Mul(position.Amount))
    }

    p.Exposure = netExposure
}

// NetExposure returns the net exposure of the entire portfolio.
func (p Portfolio) NetExposure() decimal.Decimal {
    return p.Exposure
}

