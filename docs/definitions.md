# Domain Definitions

**Sidecar** — is an implementation of decentralised clearing protocol in Yellow Network ecosystem, enabling secure and efficient transaction clearing and settlement without centralised intermediaries. Additionally, the Core domain includes secondary objectives such as market making and liquidity finding to support a vibrant and liquid trading ecosystem within the network.

**Trading Channel** —  a fundamental and pivotal concept within the Core domain of peer-to-peer trading systems. It serves as a virtual conduit established between two peers, with the primary purpose of facilitating the exchange and maintenance of their respective positions in a secure and efficient manner. The essence of a Trading Channel lies in its ability to enable seamless and real-time interactions between the involved parties, allowing them to transact assets, commodities, or financial instruments without the need for every individual trade to be executed on the underlying blockchain or main network.

**State Channel —** can be thought of as a private ledger containing balances and other arbitrary data housed in a data structure which we call a "state". The state of the channel is updated, committed to and exchanged between a fixed set of actors (which we call participants), together with some execution rules. (https://docs.statechannels.org/protocol-tutorial/0010-states-channels/)

**Position** (Deal, Obligation, Trade, Liability) — A position is the amount of a security, asset, or property that is owned (or sold short) by some individual or other entity. A trader or investor takes a position when they make a purchase through a buy order, signaling bullish intent; or if they sell short securities with bearish intent. (https://www.investopedia.com/terms/p/position.asp)

**Market Price** — The market price is the current price at which an asset or service can be bought or sold. The market price of an asset or service is determined by the forces of supply and demand. The price at which quantity supplied equals quantity demanded is the market price. (https://www.investopedia.com/terms/m/market-price.asp)

**Exposure** (Balance Sheet, Portfolio Evaluation) — Net exposure is the difference between a hedge fund’s long positions and its short positions. Expressed as a percentage, this number is a measure of the extent to which a fund’s trading book is exposed to market fluctuations. (https://www.investopedia.com/terms/n/net-exposure.asp)

**Portfolio** — the collective holdings of assets, commodities, or financial instruments that a participant currently possesses within that specific trading channel.

**Settlement** (Swap) — stage in a Trading Channel where both parties involved mutually agree to close all their positions, thereby concluding their trading activities within the channel.

## Margin Systems

- Channel Margin system - All cross positions share the channel margin balance.
- Position Margin system - The isolated margin is assigned to a position.

## Margin Definitions


| Term               |  Example | Description                                                                                                  |
|--------------------|--|--------------------------------------------------------------------------------------------------------------|
| **IMR**            | 20% | Initial Margin Rate (%) required for opening positions.                                                      |
| **MMR**            | 10% | Maintenance Margin Rate (%) to keep positions.                                                               |
| **Channel Margin** | $100 | On-chain collateral locked on the state-channel.                                                             |
| **Leverage**       | 5x | Leverage the user uses to create a position.                                                                 |
| **Position size**  | 0.018 | Position amount.                                                                                             |
| **Entry Price**    | $3,000 | Market price of an asset at the moment of opening the position.                                              |
| **Current Price**  | $2,000 | Current market price of an asset.                                                                            |
| **Position costs** | $270 | Position Size × Position Average Entry Price.                                                                |
| **Unrealized PNL** | -$18 | Position Size x Current Price - Position costs                                                               |
| **Initial Margin** | $54 | Down payment to open a position.                                                                             |
| **Margin Balance** | $82 | Channel Margin + Unrealized PNL.                                                                             |
| **Available Balance** | $28 | Channel Margin - Initial Margin + Unrealized PNL.                                                            |
| **Maintenance Margin** | $27.54 | Minimum amount of margin that must before liquidation (Positions Costs × MMR + Close Positions Fee). |

Your positions will be liquidated once ```Margin Balance <= Maintenance Margin```.