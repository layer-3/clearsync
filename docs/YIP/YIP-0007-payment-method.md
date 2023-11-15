# 0006 - Payment Method

## Status

Proposed

## Context

As Yellow Clearing Network is based on the concept of trading channels, it is necessary to have a payment method interface that enables the swap during the settlement to be performed.
The swap and the following Post-Settlement Margin state checkpoint (more in the YIP about the Settlement) form the 2-step commit, meaning that the swap cannot happen without the margin update, which removes the collateral for the funds being swapped, and vice versa.

### Related YIPs

- [0006 - Settlement](./YIP-0006-settlement.md)

## Decision

### Payment method functionality

The basic logic of the payment method is:

- Receive funds from the parties
- Swap funds
- Deliver swapped funds to the parties

The 2-step commit can fail, and the payment method should be able to:

- Return the deposited funds if any error occurs prior to the swap

However, other errors can occur after the swap, and the payment method should be able to:

- Return the swapped funds if any error occurs after the swap

This YIP does not specify the party initiating the delivery of the swapped funds for the payment method to be more flexible, therefore the payment method should be able to:

- Notify the parties when their swapped funds are delivered

### Payment method interface

> NOTE: for now it has been decided to split the preparation of the payment into several parts to resemble the escrow payment (via state-channels) method more. This is not a final decision and can be changed in the future.

```go
type PaymentMethod interface {
  // Make sure that a payment method is ready, acquire the assets from each party
  Initiate(clearingId ChannelID, settlementLedger SettlementLedger, peer Peer) error
  // Parties agree that assets are deposited and ready to be swapped
  Prepare(clearingId ChannelID) error
  // Swap the assets
  Execute(clearingId ChannelID) error
  // Withdraw deposited assets if any error occurs prior to the swap
  // It's an async operation, so a subscribe on Withdraw event is needed.
  Revert(clearingId ChannelID) error
  // Parties confirm that the assets are swapped, but not yet delivered
  Finalize(clearingId ChannelID) error
  // Retrieve the assets
  Withdraw(clearingId ChannelID) error
  // Withdraw swapped assets without counter-party confirmation if any error occurs after the swap
  ForceWithdraw(clearingId ChannelID) error
  // Subscribe and invoke `handler` when a specified `role` has fully withdrawn the funds from the Payment
  SubscribeOnWithdraw(clearingCid ChannelID, role ProtocolIndex, handler func() error) error
}
```

### Discarded interface

The alternative approach is related to an unhappy case: use `ForceFinalize` instead of `ForceWithdraw`, with a consequent `Withdraw` call, so that

- `ForceFinalize` is called when the counterparty disagree on `Finalize` as well, but it only performs the checks that the swap happened
- `Withdraw` will deliver the assets to the party (parties) in both happy and unhappy scenarios.

However, this approach is discarded as the described logic of `Withdraw` blends the happy and unhappy flows, which makes it harder to understand and maintain.

## Consequences

Implemented interface seems to be too specific for an escrow payment method, so it should be generalized to be more flexible and allow other payment methods to be implemented.

Proposed generalized interface:

```go
type PaymentMethod interface {
  // Both parties deposit the assets, agree to swap them.
  // The swap is performed here, but assets are not yet delivered.
  Execute(clearingId ChannelID) error
  // Withdraw deposited assets if any error occurs prior to the swap
  // It's an async operation, so a subscribe on Withdraw event is needed.
  Revert(clearingId ChannelID) error
  // Both parties agree that the swap was successful, but assets are not yet delivered.
  // Deliver the swapped assets to the parties.
  Withdraw(clearingId ChannelID) error
  // Withdraw swapped assets without counter-party confirmation if any error occurs after the swap
  ForceWithdraw(clearingId ChannelID) error
  // Subscribe and invoke `handler` when a specified `role` has fully withdrawn the funds from the Payment
  SubscribeOnWithdraw(clearingCid ChannelID, role ProtocolIndex, handler func() error) error
}
```
