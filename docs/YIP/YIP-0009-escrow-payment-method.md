# 0009 - Escrow Payment Method

## Status

Proposed

## Context

One of the first payment methods should be simple enough to introduce to the public, but yet secure and robust to be used in production. The escrow payment method is a good candidate for that.

### Related YIPs

- [0007 - Payment Method](./YIP-0007-payment-method.md)
- [0008 - Settlement](./YIP-0008-settlement.md)
- TODO: channel applications YIP

## Decision

The escrow payment method uses State channels to perform the swap of the assets.

### State channels

The table below describes the state channel protocol (SCP) logic used in the payment method functions.

| Interface method | Off-chain SCP               | On-chain SCP                          |
| ---------------- | --------------------------- | ------------------------------------- |
| Initiate         | Sign the prefund state (1)  | Fund the channel                      |
| Prepare          | Sign the postfund state (2) | -                                     |
| Execute          | Sign the swap state (3)     | -                                     |
| Revert           | -                           | Challenge with (2), TransferAllAssets |
| Finalize         | Sign the final state (4)    | -                                     |
| Withdraw         | -                           | ConcludeAndTransferAllAssets          |
| ForceWithdraw    | -                           | Challenge with (3), TransferAllAssets |

Note:

1. The PSM state checkpoint ([YIP 0008 - Settlement](./YIP-0008-settlement.md)), which guarantees 2-step commit, is performed after `Execute` and before `Finalize`. If it fails, the `Revert` is called.
2. In `Finalize` both parties must agree the swapped state was signed and PSMS checkpoint was performed. The channel is not concluded in this step to unblock the underlying clearing channel and allow performing margin calls or settlements.
3. The logic of withdrawing the assets without counter-party confirmation if any error after the swap happened is extracted to a `ForceWithdraw` (and not put to `Finalize`) to emphasize the unhappy path the settlement is now undertaking.
   The counter-party of the `ForceWithdraw` invoker will be seen as malicious and will be punished by the protocol.

### Data

To operate, the escrow payment method must store the following data:

- the mapping between the clearing cid and the peer
- the mapping between the clearing cid and the settlement ledger
- the mapping between the clearing cid and the escrow channel

## Consequences

1. Escrow payment method shall implement the Payment method interface [YIP 0007 - Payment Method](./YIP-0007-payment-method.md).
2. Escrow payment method shall guarantee the atomicity of the swap.

### Reflections

As the escrow payment method store its data in-memory and its functions are almost always invoked by the Initiator, when the Responder invoke `SubscribeForWithdraw`, there is not enough data to create an event subscription.

The lack of data is caused by the fact that incoming channel requests (open, update, close) are handled by the escrow channel application (YIP link), which stores data about channel.

Therefore, to fix this issue, the escrow payment method should be merged with the escrow channel application, and share the same channel data.
