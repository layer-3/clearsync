# YIP-0012: Consistent Asset Price Verification for Margin Calls

## Status

Draft

## Context

In the context of trading platforms, specifically when handling margin calls, a significant challenge arises due to the volatility of asset price indices. Even if two brokers employ the same price index formula, discrepancies can occur when one broker sends a margin call to another. This is primarily because the asset price index can change rapidly, introducing potential inconsistencies in the pricing information.

Using a fixed price range for accepting offers is not a foolproof solution, as it may not always be effective and can introduce a vulnerability to attacks.

## Decision

To address the challenges associated with margin calls and ensure consistency in asset pricing, we propose the following solution:

### Price INDEX Definition

- **Definition:** Each broker must adhere to a common definition of the price index. This includes using the same price feeds, weights, and formula for calculating the index.

### Price History Buffer

- **Buffer Implementation:** Each broker node will maintain a buffer of price history for a specified time range (e.g., the last 2 seconds). This buffer stores historical prices, providing a reference for verifying the accuracy of received margin calls. The timestamp of prices must be provided by the price sources to ensure every broker have the same data. 

### Timestamp Inclusion

- **Inclusion in Margin Calls:** When a broker sends a margin call to another, it must include the timestamp of the price used in the calculation. This timestamp serves as a reference point for the receiving broker to validate the accuracy of the provided price.

### Verification Process

- **Timestamp Check:** Upon receiving a margin call, the receiving broker checks the included timestamp against its stored price history. It accepts prices within the specified buffer time range, ensuring consistency in the asset pricing.

- **Price Movement Criteria:** Additionally, the receiving broker can apply a criterion to reject a price if the asset price has moved significantly since the timestamp used by the counterparty. In such cases, the broker rejects the proposal and initiates a new margin call based on the latest known price.

## Consequences

Implementing this solution ensures that brokers have a common understanding of asset prices, reducing the risk of discrepancies during margin calls. The inclusion of timestamps and the verification process based on historical price buffers enhance the reliability and accuracy of the margin call process.

This approach aims to mitigate the challenges posed by rapidly changing asset prices, providing a more robust foundation for margin call transactions. Further discussion and refinement of this proposal within the community are encouraged before moving it to the "Accepted" status.
