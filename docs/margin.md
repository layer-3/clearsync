# Margin operations documentation

## Definition

**Margin** is a collateral covering the net exposure of a participant in a trading channel.

## Operations

Users deposit agreed amount of margin when they open a clearing channel between each other, e.g. `m0 = [A, B]`.
After recording some trades and price changes, the net exposure of each participant changes, thus they agree on redestributing the margin: `m1 = [A + x, B - x]`.

### Margin limits

It is obvious, that drastic price changes can dramatically impact margin distribution. To ensure participant positions are always covered with collateral, we use the concept of _margin limits_ or _margin zones_.

Limits are defined as a percantage of initial margin deposited by pariticipant. For example, if the initial margin is 1000 USDT and the soft limit is 20%, then the soft limit is 200 USDT.

TBD: users can specify their own soft limits.

#### No limits

From the start, the margin is in a _green zone_ or within no limit. This means that it is safe to add new trades.

#### Soft limits

After adding some trades and some medium price changes, the margin can be in a _yellow zone_ or within soft limit.
This means that participants can no longer add new trades. To return to a _green zone_, participants can either cancel their position or request a settlement.

#### Hard limits

After a very drastic price changes or when a soft limit was ignored, the margin can transition to a _red zone_ or a _hard limit_.
This is a dangerous situation, because momentarily it can result in margin being not enough to cover the net exposure, thus it is in participant's best interest to return to a _green zone_.

If it is your margin that is in a _red zone_, your Clearport will request a Forced Settlement from your counterparty.
Only the minimum number of high-margin-impact positions required to return the margin to a _green zone_ are included to a Forced Settlement.

If your counterparty does not respond to a Forced Settlement request or it is them who are in a _red zone_, after a short period of time the Clearport will challenge the clearing channel, cancelling orders and withdrawing the margin.
