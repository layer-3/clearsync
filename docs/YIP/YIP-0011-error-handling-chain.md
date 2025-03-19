# YIP-0011 Error Handling Chain

## Status

Proposed

## Context

As the complexity of the protocol grows, the number of possible errors and edge cases increases. It is important to have a clear and consistent way of handling errors and exceptions.
This proposal describes how the errors should be handled in the protocol.

## Decision

### Definitions

**Handling chain** - is a sequence of handlers that are executed one after another. In the context of error handling chain, the next handler is executed is the previous one failed to handle the error.

This way, the behavior after an error handler failed is consistent across the protocol.

### Handler impact

Error handling chain should start with a handler with the biggest impact on a system in case of successful handling, e.g. `recovery` handler should be run first as it can potentially recover the system from the error.

Next handler in the chain should be have the next biggest impact on the system, but smaller than the previous one, e.g. `rollback` handler should be run after `recovery` handler has failed as it can potentially rollback the system to the previous state.

The last handler in the chain should be the one with the smallest impact on the system, e.g. `log` handler should be run after all previous handlers have failed as it can only log the error.

### Handler impact levels

This proposal defines the following impact levels:

- `recovery` - handler that can recover the system from the error, making the performed state transition valid
- `rollback` - handler that can rollback the system to the previous state
- `store` - handler that can store the error in the database, so that it can be processed and reviewed later
- `log` - handler that can log the error

It is worth noting that the `recovery` and `rollback` handlers can include `store` and `log` handlers **logic**.

Nevertheless, error handlers chain shall guarantee that the lower level handlers _are executed_ if the higher level handlers have failed.

### Error handling chain example

Let's consider the following example of the error handling chain.

Initiator initiates finalizing the channel, but it fails. The error handling chain should be executed in the following order:

1. `recovery`: retry the finalization. If retry fails, then
2. `rollback`: try to rollback the channel to operational state. If rollback fails, then
3. `store`: store the error in the database. If storing fails, then
4. `log`: log the error.

## Consequences

- The protocol shall implement the error handling chain as described in this proposal.
- Error handling chains for different errors should contain all levels described above (if possible).
