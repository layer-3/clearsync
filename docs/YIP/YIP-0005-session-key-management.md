# 0005 - Session Key Management for ClearSync

## Status

Work in progress

## Context

The adoption of session keys enhances security and operational efficiency.
By limiting the validity of keys and associating them with specific counterparts, we significantly reduce risks associated with key exposure and unauthorized access.

This proposal introduces a new strategy for key management in the ClearSync project, shifting from signing sub-keys to registering session keys.
Each Clearport will generate a key pair limited to a default duration of 24 hours, termed a clearing session.

## Decision

### Key Generation and Registration

- **Generation**: Each Clearport generates a key pair valid for a default duration of 24 hours.
- **Registration**: Participants must register the clearing public key on the adjudicator, specifying the duration, counterpart, and owner (Smart-Account).
- **Duration**: 24 hours
- **Counterpart**: Smart-Account
- **Owner**: Smart-Account

### Key Features

- **Spending Limit**: A spending limit can be added to the clearing key.
- **Collateral Association**: The collateral is associated with this clearing key.
- **Revocation**: The Clearing-Key can be revoked at any time.
- **NeoDAX Access**: The same key can access NeoDAX, with the trading session having the same duration.

### Renewal and Expiration
- **Renewal**: Participants must extend the expiration of the key every 24 hours.
- **Invalidity Post-Expiration**: Past the period, the key is considered invalid for ForceMove, Finalization, and signing states.

### Session Expiration Cases

- Counterpart server crash
- Potential key compromise
- Discontinuation of clearing by the counterpart
- Liquidation trigger by the counterpart

### Liquidation of Positions
- **Automatic Closure**: If a SessionKey expires, all positions are automatically closed, and the channel is concluded.
- **Under-collateralization**: If positions are under-collateralized, all positions are closed, the channel is concluded, and the Clearing-Key is invalidated.

### Security Considerations
- **Vault and KMS**: In Clearport production, a grade Vault with Key Management Service (KMS) is recommended.
- **Desktop Use**: When used on a desktop, keys should be persisted in SQLite or PGSQL, preferably with encryption at rest.

## Consequences

1. Clearport shall use process to generate key pairs at regular intervals
2. Clearport interface shall enable participants to seamlessly register clearing public keys on the adjudicator, including specifying duration, counterpart, and owner details
3. Renewal process should ensure continuous validity of keys
4. Adjudicator should effectively validate and store registered keys along with associated metadata
5. Adjudicator should to accurately track key expiration and enforce associated actions
