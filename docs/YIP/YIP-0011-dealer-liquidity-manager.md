# YIP-0011 - Dealer Liquidity Manager

## Status

Proposed

## Context

Managing liquidity and mitigating risks associated with unsettled positions is crucial in the dynamic and volatile world of cryptocurrency trading. A Dealer Liquidity Manager (DLM) is proposed to periodically review unsettled positions and implement risk management strategies to secure against volatility.

## Decision

The Dealer Liquidity Manager will employ predefined strategies to manage the risks associated with unsettled positions. These strategies include:

1. **Swap on Affordable Blockchain on a DEX**: Utilizing decentralized exchanges (DEXs) on blockchains with lower transaction fees for swapping positions to more stable assets.
2. **Hedge on Futures Market**: Engaging in futures contracts to hedge against potential losses in the spot market.
3. **Buy/Sell on a CEX Spot Market**: Executing buy or sell orders on centralized exchange (CEX) spot markets to balance the portfolio.

### Strategy Routing

The DLM will determine the appropriate strategy based on the following criteria:

1. **Default Strategy**: A pre-set, general approach that applies to most situations.
2. **Fixed Strategy for a Market**: Specific strategies tailored to particular markets or assets.
3. **Conditional Strategy**: Strategies triggered based on certain criteria, such as the average leverage used for the open positions.

### Implementation

The DLM will be integrated into the existing trading system, with periodic checks on unsettled positions. The system will:

- Analyze each position based on predefined criteria.
- Select the appropriate risk management strategy.
- Execute the strategy in a timely and efficient manner.

## Consequences

1. The DLM will enhance the platform's ability to manage risks associated with unsettled positions.
2. It will provide flexibility in strategy selection, adapting to different market conditions and specific needs of each position.
3. The DLM will aim to minimize potential losses due to market volatility and improve overall portfolio performance.

### Reflections

While the DLM aims to mitigate risks, it's important to note that:

- The effectiveness of each strategy may vary based on market conditions.
- There is a need to continuously monitor and adjust the strategies to ensure they remain effective.
- The DLM should be part of a broader risk management framework, complementing other tools and strategies.

Integrating the DLM into the trading system represents a proactive approach to liquidity management and risk mitigation in the fast-paced and often unpredictable cryptocurrency market.
