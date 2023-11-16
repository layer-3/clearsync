# 0008 - Settlement

## Status

Proposed

## Context

Settlement is an important part of the trading process. It is the final step of the trading cycle, where the parties exchange the assets they have agreed to trade.

In the context of Yellow Network, several consequent settlements can be performed on the same trading channel, which gives more flexibility to the traders and reduces operation fees.

### Related YIPs

- [0006 - Margin Zones](./YIP-0006-margin-zones.md)
- [0007 - Payment Method](./YIP-0007-payment-method.md)

## Decision

The main role of settlement is for parties to safely exchange the assets they have agreed to trade, and to update their margin by excluding the collateral covering the assets being traded.

The margin excluding the collateral covering the assets being traded is called the **post-settlement margin** (PSM).
The settlement shall be performed as a 2-step commit to ensure that the swap of the assets does not happen without the PSM update, and vice versa.

### 2-state commit process

To link the swap of the assets with PSM update, either of them should include a check that the other one has been performed.

#### Discarded approach

One approach is to put the swap before the PSM update, so that the latter shall have a check that the swap has been performed.

Given that a PSM update is performed inside a state channel application, which has access to on-channel data only, this approach complicates the check that the swap was completed for an off-chain payment methods,
forcing them to store the swap data on-chain, which introduces additional waiting time, complexity and fees.

Moreover, even on-chain payment methods may delete the swap data after it is completed (e.g. state-channels [remove on-chain channel data](https://github.com/statechannels/go-nitro/blob/ca23778ca28ac93113ec227a342932256242509c/packages/nitro-protocol/contracts/NitroAdjudicator.sol#L89) after all funds have been withdrawn), which makes the check impossible.

#### Accepted approach

The other approach is the opposite: perform the PSM update before the swap, so that the latter shall have a check that the PSM update has been performed.

The PSM update can performed inside a state channel application via a [`checkpoint`](https://docs.statechannels.org/protocol-tutorial/0070-finalizing-a-channel/?extract-info-from-adjudicator-events#call-checkpoint), thus this data is stored on-chain is accessible both for on-chain and off-chain applications.

#### Linking the swap with PSM update

The swap should include a PSM state turn-number and before performing the swap, using an on-chain channel data the payment method should check that the channel supports a state with a turn number equal or greater than the PSM turn number specified in the swap.
Otherwise, the swap should be rejected.

Then both participants should agree on the swap with the specified PSM turn number, e.g. by signing some data.

Only after both parties agree on the swap with the PSM turn number, they can sign the PSM state with the specified turn number and checkpoint it in a clearing channel.

This way, the swap is not executed until the PSM state is checkpointed, and the PSM state should not be signed until a swap, which includes the PSM turn number, is agreed.

Such a 2-step commit logic creates a knot between the swap and the PSM update, which is untied by the PSM support check performed in a payment method.

// TODO: picture of the 2-step commit

### Settlement configuration

#### PSM state turn number

Before initiating the settlement, participants shall agree on the PSM state turn number to guarantee the atomicity of the settlement.
Later during the settlement process, a PSM state will be created and this turn number will be used.

PSM state turn number shall be bigger than the turn number of latest supported state.

#### Settled markets

A settlement is performed on a trading channel for the _selected markets_. The markets are selected by the initiator of the settlement and then proposed to the Responder.

Note, that net exposure of the selected markets shall not move the margin to yellow or red zone, otherwise the settlement should be rejected.

#### Settlement type

Settlement can be of 2 types: available and forced.

##### Available

Available settlement is initiated by the user, who selects the markets to be settled, and sends a settlement request to the counter-party.
The counter-party can freely reject the request, and the parties can negotiate the list of markets to be settled.

##### Forced

Forced settlement can be initiated by the user or by the Clearport, but the main difference is that it signals the counter-party that if they are to reject the settlement, the channel is going to be challenged.

In other words, by proposing a forced settlement, the initiator states that if the settlement is rejected, the channel is going to be challenged.

There are 2 main use-cases for Forced Settlement:

- **"Urgent settle all"** initiated by the Initiator when they want to settle all markets as soon as possible, e.g. when they can no longer be online.
- **"System precaution"** created by the system when the margin transitions to the red zone (see [Margin Zones YIP](./YIP-0006-margin-zones.md)), so that the channel is not left undercollateralized.

If the Forced Settlement is rejected by the Responder, the Clearport on the Initiator side is going to challenge and close the channel automatically.

#### Settlement payment method

One of the goals of the settlement is to swap the assets being traded, that can be done via a payment method, which can include an escrow payment (non-custodial solution), 3-rd party custodian, or even an off-chain solutions.

The payment method is specified in the settlement request, and the Responder can accept or reject the request based on their preferences regarding the payment method.

Payment method is outside of the scope of this proposal and is going to be specified in a separate YIP.

#### Settlement state

To keep track of a settlement flow, a state is used that resembles the progress of funds swap between the participants:

```proto
enum SettlementState {
  SETTLEMENT_STATE_UNSPECIFIED = 0;
  SETTLEMENT_STATE_PROPOSED = 1;
  SETTLEMENT_STATE_ACCEPTED = 2;
  SETTLEMENT_STATE_INITIATED = 3;
  SETTLEMENT_STATE_PREPARED = 4;
  SETTLEMENT_STATE_EXECUTED = 5;
  SETTLEMENT_STATE_COMPLETED = 6;
  SETTLEMENT_STATE_WITHDRAWN = 7;
  SETTLEMENT_STATE_FAILED = 8;
  SETTLEMENT_STATE_REJECTED = 9;
}
```

##### Settlement state updates

Both parties shall have the same settlement state and update it only after the other party has agreed on the settlement state change.

For that purpose, the Clearport shall expose an endpoint (e.g. `UpdateSettlementState`), which accepts the clearing channel id and a settlement state to transition to:

```proto
message SettlementStateUpdate {
  string cid = 1;
  SettlementState to_state = 2;
}
```

### Settlement proposition

After all settlement parameters are assembled, the Initiator shall send a settlement proposal to the Responder.
If the latter wants to accept the request, they shall specify the state of the settlement as `Accepted` and include the proposed list of markets. Otherwise, to reject the request, they shall specify the state of the settlement as `Rejected` and provide a list of desired markets if any.

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

### Margin calls and PSM state turn number

As drastic price changes can happen during the settlement, the latest mutually signed margin distribution can become obsolete. For that reason, even during the settlement, participants can perform margin calls to update the margin distribution, which still includes the net exposure from markets that are being settled.
With each supported margin call, its turn number increases, but it can not be bigger than the PSM state turn number agreed before the settlement.

Therefore, before creating a settlement, Initiator can assess the time required to perform the payment, and propose such a PSM state turn number, which will allow participants to perform enough margin calls during the settlement.

### Margin zones and settlement

As the yellow and red margin zones are undesirable, the settlement should not result in the post-settlement margin in these zones.

This means that Clearpost should be able to identify the zone of the post-settlement margin when users select markets and disallow settling if the resulting zones are yellow and red.

During the settlement participants can perform margin calls, which can lead to transition to a different margin zone.
Therefore, if a post-settlement margin calculation transitions to the yellow zone during the settlement process, the users are notified that the current settlement should be accelerated and a new one should be prepared soon after the current is completed.
If the a margin call transitions to the red zone, the settlement is interrupted and a new Forced-one is created.

### Settlement flow

#### Happy case

1. Initiator sends a settlement request to the responder, specifying the markets to be settled and state `Proposed`.
2. Responder accepts the settlement request, returning specified markets and state `Accepted`.
3. Initiator updates settlement state to `Accepted`.
4. Positions of the settled markets are locked at both sides, so that if a trade on those markets arrive, it is put into a separate position.
5. Initiator calls `PaymentMethod.Initiate`, which locks the assets to be swapped.
6. Initiator requests the Responder to change the settlement state to `Initiated`.
7. Responder accepts the settlement state change.
8. Initiator updates settlement state to `Initiated`.
9. Initiator calls `PaymentMethod.Prepare`, stating that the funds are ready to be swapped.
10. Initiator requests the Responder to change the settlement state to `Prepared`.
11. Responder accepts the settlement state change.
12. Initiator updates settlement state to `Prepared`.
13. Initiator calls `PaymentMethod.Execute`, which swaps the assets.
14. Initiator calculates the PSM with PSM state turn number taken from the settlement, and proposes it to the responder.
15. Responder accepts the PSM.
16. Initiator checkpoints the adjudicator with PSM.
17. Initiator requests the Responder to change the settlement state to `Executed`.
18. Responder accepts the settlement state change.
19. Initiator updates settlement state to `Executed`.
20. Initiator calls `PaymentMethod.Finalize`, stating that the assets are swapped, but not yet delivered.
21. Initiator requests the Responder to change the settlement state to `Completed`.
22. Responder accepts the settlement state change.
23. Initiator updates settlement state to `Completed`.
24. Initiator calls `PaymentMethod.Withdraw`.
25. Parties, if subscribed to `PaymentMethod.SubscribeOnWithdraw`, are notified that the assets are delivered.

#### Unhappy case - Responder rejected the settlement

1. Initiator sends a settlement request to the responder, specifying the markets to be settled and state `Proposed`.
2. Responder rejects the settlement request, returning desired markets (if any) and state `Rejected`.
3. Initiator update settlement state to `Rejected`.
4. Initiator and Responder can keep ping ponging the settlement request until they agree on the markets to be settled.

#### Unhappy case - swap preparations failed

1. Initiator sends a settlement request to the responder, specifying the markets to be settled and state `Proposed`.
2. Responder accepts the settlement request, returning specified markets and state `Accepted`.
3. Initiator updates settlement state to `Accepted`.
4. Positions of the settled markets are locked at both sides, so that if a trade on those markets arrive, it is put into a separate position.
5. Payment initiation, preparation or execution failed.
6. Initiator calls `PaymentMethod.Revert`, reclaiming the locked assets back for both parties.
7. Initiator requests the Responder to change the settlement state to `Failed`.
8. Responder accepts the settlement state change.
9. Initiator updates settlement state to `Failed`.

#### Unhappy case - swap withdrawal failed

1. Initiator sends a settlement request to the responder, specifying the markets to be settled and state `Proposed`.
2. Responder accepts the settlement request, returning specified markets and state `Accepted`.
3. Initiator updates settlement state to `Accepted`.
4. Positions of the settled markets are locked at both sides, so that if a trade on those markets arrive, it is put into a separate position.
5. Initiator calls `PaymentMethod.Initiate`, which locks the assets to be swapped.
6. Initiator requests the Responder to change the settlement state to `Initiated`.
7. Responder accepts the settlement state change.
8. Initiator updates settlement state to `Initiated`.
9. Initiator calls `PaymentMethod.Prepare`, stating that the funds are ready to be swapped.
10. Initiator requests the Responder to change the settlement state to `Prepared`.
11. Responder accepts the settlement state change.
12. Initiator updates settlement state to `Prepared`.
13. Initiator calls `PaymentMethod.Execute`, which swaps the assets.
14. Initiator calculates the PSM with PSM state turn number taken from the settlement, and proposes it to the responder.
15. Responder accepts the PSM.
16. Initiator checkpoints the adjudicator with PSM.
17. Initiator requests the Responder to change the settlement state to `Executed`.
18. Responder accepts the settlement state change.
19. Initiator updates settlement state to `Executed`.
20. Initiator calls `PaymentMethod.Finalize`, stating that the assets are swapped, but not yet delivered.
21. Finalization fails.
22. Initiator calls `PaymentMethod.ForceWithdraw`, reclaiming the locked assets back for both parties.

## Consequences

1. Settlement shall be performed as a 2-step commit, where the swap of the assets does not happen without the PSM update, and vice versa. The atomicity is guaranteed by including PSM state turn number in the swap, and performing a checkpoint of PSM state before the swap.
2. To assemble the settlement, Initiator should specify the markets to be settled, the PSM state turn number, and the payment method.
3. Settlement state shall be updated only after the other party has agreed on the settlement state change.
4. Participants can perform margin calls during the settlement, but the turn number of the latest supported margin call shall not be bigger than the PSM state turn number.
5. Clearport shall disallow proposing settlement if the post-settlement margin calculation for selected markets is in the yellow or red zone.
6. Clearport shall interrupt ongoing settlement if the current post-settlement margin calculation is in the red zone, and start a new settlement according to these rules.
