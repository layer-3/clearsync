## Problem

- **Extensive heap usage**: the current implementation involves significant heap usage due to frequent allocations when operating on `decimal.Decimal` or `big.Int`. When these operations are chained, the inefficiency multiplies.
- **Order book servicing**: representing the order book as a binary tree demands continuous balancing, adding computational overhead.
- **Time complexity**: searching through the order book is log-linear $O(log(n))$ with respect to the size of the order book, making it less efficient as the order book grows.

## Concept

If the price of SHIBUSDT is *`0.00000792`*, a trade would mentally operate on it as just *`792`* and not *`zero, dot, 5 zeros, 792`*. The same idea can be used to streamline the representation of prices by using a fixed number of significant digits – so the price is broken down to 8 significant digits – *`0000 0792`*. 

Let's call this format a Trader Representation, in contrast to Storage Representation (e.g. a `NUMERIC(precision, scale)` type in Postgres) and Blockchain Representation (most ERC-20 tokens use 18 decimals so as to stay comparable to other tokens in terms of supply).

## Order book representation

The order book is represented as an **array of arrays of orders**, where the size is determined by the largest possible number that can be represented with the significant digits, e.g. `9999 9999`. Order book balancing is not required in such representation.

The exact position of price decimal separator doesn't matter, since the number of significant digits is fixed. Thus the only math operations relevant for this representation are less than or equal `≤` and greater than or equal `≥`.

The internal representation of price can be either `[]byte` or `uint32`.

## Order matching

For a *buy* order, the order book is queried for all orders that are ≤ `0000 0792`. The result would look like a slice of orders (`orderbook[0000 0000 : 0000 0792]`).

## Benefits

- **Reduced heap usage**: no need to use `decimal.Decimal` or `big.Int`
- **Efficient order book maintenance**: no time is spent on order book balancing, streamlining the maintenance process.
- **Faster comparisons**: comparing two prices becomes faster.
- **Constant time order matching**: the time complexity of order matching is constant irrespective of the size of the order book.
- **Forced order book depth**: users can only post orders fitting the significant digits, thereby controlling the order book's depth.
- **Network bandwidth optimisation**: by transmitting smaller and simpler price data, the concept ensures less network bandwidth usage.
- **Reduced slippage**: traders get the price they expect with minimal deviation.

