# 0006 - Settlement

## Status

Accepted

## Context

Settlement is an important part of the trading process. It is the final step of the trading cycle, where the parties exchange the assets they have agreed to trade.

In the context of Yellow Network, several consequent settlements can be performed on the same trading channel, which gives more flexibility to the traders and reduces operation fees.

## Decision

The main role of settlement is for parties to safely exchange the assets they have agreed to trade, and to update their margin by excluding the collateral covering the assets being traded.

The margin excluding the collateral covering the assets being traded is called the **post-settlement margin** (PSM).
The settlement must be performed as a 2-step commit to ensure that the swap of the assets does not happen without the PSM update, and vice versa.

### 2-state commit process

To link the swap of the assets with PSM update, either of them should include a check that the other one has been performed.

#### Discarded approach

One approach is to put the swap before the PSM update, so that the latter must have a check that the swap has been performed.

Given that a PSM update is performed inside a state channel application, which has access to on-channel data only, this approach complicates the check that the swap was completed for an off-chain payment methods,
forcing them to store the swap data on-chain, which introduces additional waiting time, complexity and fees.

Moreover, even on-chain payment methods may delete the swap data after it is completed (e.g. state-channels [remove on-chain channel data](https://github.com/statechannels/go-nitro/blob/ca23778ca28ac93113ec227a342932256242509c/packages/nitro-protocol/contracts/NitroAdjudicator.sol#L89) after all funds have been withdrawn), which makes the check impossible.

#### Accepted approach

The other approach is the opposite: perform the PSM update before the swap, so that the latter must have a check that the PSM update has been performed.

The PSM update can performed inside a state channel application via a [`checkpoint`](https://docs.statechannels.org/protocol-tutorial/0070-finalizing-a-channel/?extract-info-from-adjudicator-events#call-checkpoint), thus this data is stored on-chain is accessible both for on-chain and off-chain applications.

#### Linking the swap with PSM update

The swap should include a PSM state turn-number and before performing the swap, using an on-chain channel data the payment method should check that the channel supports a state with a turn number equal or greater than the PSM turn number specified in the swap.
Otherwise, the swap should be rejected.

Then both participants should agree on the swap with the specified PSM turn number, e.g. by signing some data.

Only after both parties agree on the swap with the PSM turn number, they can sign the PSM state with the specified turn number and checkpoint it in a clearing channel.

This way, the swap is not executed until the PSM state is checkpointed, and the PSM state should not be signed until a swap, which includes the PSM turn number, is agreed.

Such a 2-step commit logic creates a knot between the swap and the PSM update, which is untied by the PSM support check performed in a payment method.

// TODO: picture of the 2-step commit

### Settlement state

// TODO:

### Settled markets

A settlement is performed on a trading channel for the _selected markets_. The markets are selected by the initiator of the settlement and then proposed to the Responder.
If the latter wants to accept the request, they must specify the state of the settlement as `Accepted` and include the proposed list of markets. Otherwise, to reject the request, they must specify the state of the settlement as `Rejected` and provide a list of desired markets if any.

```proto
// Settlement proposal request
message PrepareRequest {
  Settlement settlement = 2;
}

// Settlement proposal response
message PrepareResponse {
  SettlementState state = 1;
  repeated string markets = 2;
}
```

### Settlement type

Settlement can be of 2 types: available and forced.

#### Available

Available settlement is initiated by the user, who selects the markets to be settled, and sends a settlement request to the counter-party.
The counter-party can freely reject the request, and the parties can negotiate the list of markets to be settled.

#### Forced

Forced settlement can be initiated by the user or by the Clearport, but the main difference is that it signals the counter-party that if they are to reject the settlement, the channel is going to be challenged.

In other words, by proposing a forced settlement, the initiator states that if the settlement is rejected, the channel is going to be challenged.

There are 2 main use-cases for Forced Settlement:

- **"Urgent settle all"** initiated by the Initiator when they want to settle all markets as soon as possible, e.g. when they can no longer be online.
- **"System precaution"** created by the system when the margin transitions to the red zone (TODO: margin zone YIP)

If the Forced Settlement is rejected by the Responder, the Clearport on the Initiator side is going to challenge and close the channel automatically.

### Settlement payment method

One of the goals of the settlement is to swap the assets being traded, that can be done via a payment method, which can include an escrow payment (non-custodial solution), 3-rd party custodian, or even an off-chain solutions.

The payment method is specified in the settlement request, and the Responder can accept or reject the request based on their preferences regarding the payment method.

Payment method is outside of the scope of this proposal and is going to be specified in a separate YIP.

### Settlement flow

The settlement flow is as follows:

1. The initiator of the settlement sends a settlement request to the responder, specifying the markets to be settled.
2. The responder accepts the settlement request.
3. Positions of the settled markets are locked at both sides, so that if a trade on those markets arrive, it is put into a separate position.
4.

## Consequences

TODO:
