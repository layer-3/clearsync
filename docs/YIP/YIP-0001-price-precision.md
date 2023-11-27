# 0001 - Asset Price Precision

## Status

Approved

## Context

In trading industry most users understand PIP, or price-points which is the number of significant digits composing a price.

In FOREX a pip is 0.1 Cent, and this is the smallest price increment. Matching engine should match using only those price steps and not lower values.

In the context of YN, we must have a network global configuration to make members join easily and have harmony of configuration, we also aim in having configless deployments or network wide configuration files using IPFS.

## Decision

### Definitions

**Price** - asset price in quote currency.

**Quantity** - in trading, usually refers to the number of units or shares
of a particular security or asset that is being bought or sold.
For instance, if an investor buys 100 shares of a company's stock, the quantity in this case is 100.

| Number      | Round to |
| ----------- | ----------- |
| Price       | 18       |
| Quantity    | 8        |

**Significant figures** (also known as the precision of a number in positional notation) - are digits in the number that are reliable and necessary to indicate the quantity of something.

The precision level of all trading prices is calculated based on significant figures. Yellow members shall use five (5) significant digits for the prices of all pairs of the network.
Alternative options for significant figures are in range of 5 - 8.

Some examples of _five significant digits_ are 1.0234, 10.234, 120.34, 1234.5, 0.012345, and 0.00012340.

This is similar to how traditional global markets handle the precision of small, medium, and large values. The reasoning behind this is that the number of decimals becomes less important as the quantity increases. The opposite is true for minimal amounts, where greater precision is more valuable.

Quantity precision allows up to 8 decimals. Anything exceeding this shall be rounded to the 8th decimal.

## Consequences

1. All API outputting prices shall be truncated and formation using five (5) significant figures
2. All API input shall return an error for orders where prices or amounts exceed 18 and 8 precision respectfully
3. All frontend and UI clients shall truncate order price input to five (5) significant figures
4. Backend and UI pkgs shall support easy change of significant figures in range of 5 - 8.
