# 0006 - Margin Zones

## Status

Proposed

## Context

One of the main goals of a clearing channel is to increase the speed and reduce the fees of the trading process by collecting trades into batches and performing settlements on them.

However, as the trades are not settled immediately, this exposes users to the market price volatility risks, which can be mitigated by collateralizing the trades with the concept of margin.

When opening a clearing channel, participants must provide enough margin (collateral) to cover price changes of the trades they are going to perform.
Nevertheless, if the market is volatile, user's net exposure can exceed the margin provided, which is a huge security downside that the Yellow Network should mitigate.

## Decision

For the Clearport to be able to avoid undercollateralized trades, the margin should be divided into zones, and when transitioned to, each zone should trigger a special action.

### Margin zones

- **Green** zone (default) - the margin distribution is enough for both parties to cover possible price changes that they have performed.
- **Yellow** zone - the margin distribution shifts the bigger part of a margin for one party, meaning the other one can become undercollaterallized if the market moves against them.
  When the margin transitions to the _yellow zone_, users can no longer create trades and are adviced to perform settlement to move margin to a _green zone_. Note, that parties can select what markets to settle themselves.
- **Red** zone - the margin is almost depleted for one of the parties, meaning if the market continues to move against them, they will become undercollaterallized soon.
  When the margin transitions to the _red zone_, users also are not able to create trades, and an automatic settlement is created by the Clearport to move margin to a _green zone_.
  Note, that there can be several approaches for Clearport to select the markets to settle.

#### Margin zones limits

When described in a context of one party, margin zone limits are represented as a percentage of their margin, e.g. green zone is 100% - 25%, yellow zone is 25% - 10%, and red zone is 10% - 0%.

"Channel margin is not in a green zone" means that margin for one of participants is not in a green zone, and this participant should be specified, e.g. `Alice [88, 12] Bob` means that margin is in yellow zone for Bob (given zone limits above).

Zone limits can be configurable by users and this proposal does not define the exact values.

#### Market selection for automatic settlement

This proposal does not define an algorithm for Clearport to select markets for an automatic settlement, but it can be one of the following:

- **Settle all** - the Clearport selects all markets for the settlement. The resulting margin distribution is 50/50.
- **Move to center with the least amount of most impactful** - the Clearport selects the least amount of the markets with the highest net impact, such that after settlement margin would be the closest to 50/50.
- **Move to green with the least amount of most impactful** - alike the previous, but the resulting margin is shifted more to the red zone it was before.

> The difference between the last two is that the _"Move to green ..."_ changes the margin distribution less. For example, before the settlement the margin disrtibution was `[93, 7]` with the green zone at 15 and higher,
> the _"Move to center ..."_ market selection algorithm can change the margin to be `[52, 48]`, and the _"Move to green ..."_ to `[83, 17]`.

### Restrictions on post-settlement margin zone

As the yellow and red zones are undesirable, the settlement should not result in the post-settlement margin in these zones.

This means that Clearpost should be able to identify the zone of the post-settlement margin when users select markets and disallow settling if the resulting zones are yellow and red.

It is worth noting that the post-settlement margin is calculated periodically during the settlement to allow participants to agree on the margin distribution which indeed resembles the market.
Therefore, if a post-settlement margin calculation transitions to the yellow zone during the settlement process, the users are notified that the current settlement should be accelerated and a new one should be prepared soon after the current is completed.
If the PSM calculation transitions to the red zone, the settlement is automatically interrupted and a new automatic one is prepared by the Clearport.

## Consequences

1. Margin zone limits shall be defined in the application configuration.
2. Clearport shall notify other components that the margin zone has changed.
3. Clearport shall disallow adding new trades if the margin is in the yellow or red zone.
4. Clearport shall initiate automatic settlement if the margin is in the red zone.
5. Clearport shall implement at least one, and should implement several market selection algorithm for automatic settlement, which may be one of the described in this document or a custom one(s).
6. Clearport shall disallow proposing settlement if the post-settlement margin calculation for selected markets is in the yellow or red zone.
7. Clearport shall interrupt ongoing settlement if the current post-settlement margin calculation is in the red zone, and start a new settlement according to these rules.
