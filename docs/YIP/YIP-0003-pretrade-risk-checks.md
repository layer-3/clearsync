# 0003 - Pre-Trade risk checks

## Status

In Work

## Context

We have to develop the mechanism that will limit users possibility to create orders that will put parties at risk of not been able to cover the net exposure by the colleteral. Exposing the parameters by API so it will be visible and understood by the user of the system. Like a Buying power = colleteral * levarage - open order amount, and Equity = colleteral + net exposure.
https://www.b2bits.com/trading_solutions/pre-trade-risk-check-solutions


## Decision

**Pre-Trade risk checks vol.1**

### Definitions

Trading channel's ```buying power``` is equal to ```collateral * margin - locked_in_open_orders``` in quote currency.

### Order creation checks

- In Neodax each open channel has its session to control orders creation. A session contains channel's margin and locked balances, leverage, and rate limit restrictions.
- Neodax and ClearPort are synchronized - channels' margin balances are updated in real-time on each operation
- Each order to create goes through a balance check. Neodax compares if the channel's up-to-date buying power is sufficient to allow order creation. In a successful case, the order amount is added to a channel's locked balance, thus the buying power gets decreased. In the case of market orders, as they are executed immediately, the lock isn't happening, and the created order affects the margin balance.

### Open orders cancellation when buying power goes below zero

- If the buying power goes below zero (for example, position losses), Neodax starts cancelling open orders until the buying power is recovered.
Cancellation starts with orders with the smallest locked amount in quote currency and contunues until the buying power is more or equal to zero.

## Consequences

- Channel balances are updated in real-time. Buying power is a complex dynamic channel balance acting as a risk mitigation mechanism.
- Each order to create goes through a balance check.
- If the buying power goes below zero, open orders are sequentially canceled until the balance is recovered.
- Clear communication of the consequences is a must. Users should be notified about the system's actions, for instance, automatically canceled orders.
