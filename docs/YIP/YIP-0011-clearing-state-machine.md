# 00011 - Clearing State Machine

## Status

Proposed

## Context

Clearing process become more and more complex as the protocol evolves. The clearing state machine (CSM) is a good candidate to encapsulate the logic of the clearing process. And make sure that we handle (or reject to handle) all the possible events of the clearing process.

## Decision

Create a state machine that handles the clearing process using prepared framework.

### CSM diagram

```mermaid
---
title: Clearing State Machine
---

stateDiagram-v2
    [*] --> Instantiated: Instantiate
    Instantiated --> Accepted: Accept
    Instantiated --> Failed: Fail

    Accepted --> InitiatorFunded: InitiatorFunds
    InitiatorFunded --> Funded: Fund
    Funded --> Operational: MoveToOperational

    Operational --> ActiveSettlement: PerepareSettlement
    Operational --> Finalized: Finalize
    Operational --> IssuedMarginCall: IssueMarginCall
    
    IssuedMarginCall --> Operational: MoveToOperational
    IssuedMarginCall --> Challenged: Challenge

    ActiveSettlement --> Operational: CompleteSettlement
    ActiveSettlement --> Operational: FailSettlement
    
    Challenged --> Operational: ClearChallenge
    Challenged --> Withdrawn: TimeoutChallenge
    
    Finalized --> Concluded: Withdraw
    Concluded --> [*]
    Withdrawn --> [*]
```


## Consequences

TODO
