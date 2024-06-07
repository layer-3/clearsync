# YIP-0014: Staking Protocol on State-Channel using Nitro Framework

## Abstract
This proposal introduces a staking protocol utilizing state channels in the Nitro framework. It enables users to lend liquidity to brokers, who can then register these assets on their balance sheet for use in an Automated Market Maker (AMM) through a dealer module. The protocol also outlines the mechanics of a cooldown period for asset withdrawal and settlement processes in escrow channels.

## Motivation
The motivation behind this proposal is to enhance liquidity in the Yellow Network's AMM system. By allowing users to lend assets to brokers, we can increase market efficiency and provide brokers with additional liquidity to fulfill their role in the AMM. This initiative also aims to incentivize users by sharing trading fees and interest generated from the borrowed capital.

## Specification

### Staking and Lending Mechanism
- Users open a state channel with the Nitro protocol to lend liquidity to a broker.
- The lent assets (e.g., FOXY tokens) are locked in the channel and registered on the broker's balance sheet.

### Broker's Use of Assets
- Brokers use the lent assets for market-making activities in pairs like FOXY/USDC.
- Brokers share a portion of their trading fees and the interest rate of the borrowed capital with the lenders.

### Cooldown Period for Asset Withdrawal
- Lenders can request to unlock their assets from the lending pool.
- A cooldown period of 10 days is implemented before the assets are released to the lender.

### Settlement in Escrow Channels
- In the event of a settlement request, the lent assets (e.g., FOXY tokens) can be moved to the escrow settlement channel.
- The broker provides a quote currency (e.g., USDC) in return to the lending channel, maintaining the balance.

## Rationale
Implementing this protocol facilitates a dynamic and efficient AMM system within the Yellow Network. The cooldown period ensures stability in the liquidity pool, while the escrow channel settlements provide flexibility in handling transactions and settlements.

## Backwards Compatibility
This proposal is compatible with existing protocols in the Yellow Network and leverages the Nitro framework for state channel management.

## Test Cases
- Case 1: Lending FOXY tokens to a broker and registering it on the broker's balance sheet.
- Case 2: Broker utilizing the lent FOXY tokens in the AMM for a FOXY/USDC pair.
- Case 3: Lender requesting asset withdrawal and the initiation of the cooldown period.
- Case 4: Settlement transaction in an escrow channel with a quote currency exchange.

## Implementation
The implementation will require updates to the Nitro protocol state channels and integration with the Yellow Network's AMM and broker-dealer modules.

## Security Considerations
- Ensuring the security of locked assets in the lending channel.
- Robust implementation of the cooldown period mechanism to prevent exploitation.
- Secure handling of asset exchanges in escrow channel settlements.

## References
- YIP-0000: YIP Format and Guidelines
