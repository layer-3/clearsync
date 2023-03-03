# Technical Paper - Yellow Network

[TOC]

## 1. Introduction to Yellow Network: Problems and Solution

### 1.1 Problems facing the crypto industry

As a new form of distributed ledger technology (DLT), Bitcoin has set out to decentralize the issuance of money as well as the transfer thereof through [a peer-to-peer electronic cash system](https://bitcoin.org/bitcoin.pdf)[^1]. Ensuing DLT systems like Ethereum, generally referred to as [smart contract](https://ethereum.org/en/whitepaper/)[^2] or [general-purpose](https://www.oreilly.com/library/view/mastering-ethereum/9781491971932/ch01.html) blockchains,[^3] have been designed to decentralize trading and finance at large.

While blockchain technology has reliably enabled decentralized computation, current manifestations have been far from being able to scale to match the needs of traditional finance. Decentralized computation on a base layer blockchain, although secure, is slow and hinders the scalability of crypto trading. Faced with an inherent [scalability trilemma](https://vitalik.ca/general/2021/04/07/sharding.html)[^4], as of the end of February 2023, [126](https://blockchain-comparison.com/blockchain-protocols/)[^5] layer-1 blockchain protocols have been competing and experimenting in an attempt to find a solution.

Although we are convinced that much innovation will spring from this vibrant competition, the current blockchain environment still suffers from three main problems as a result of this. These are:

- Market and liquidity fragmentation

- Lack of decentralization

- Lack of regulatory frameworks

### 1.2 Market and liquidity fragmentation

Due to the proliferation of blockchains, there has been a fragmentation of assets and liquidity across different layer-1 protocols and increasingly [layer-2 chains](https://l2beat.com/scaling/tvl/)[^6]. Their interoperability is limited, having hitherto been dependent on centralized cross-chain bridges that have emerged as a [top security risk](https://blog.chainalysis.com/reports/cross-chain-bridge-hacks-2022/)[^7] due to the many substantial hacks that have [occurred](https://twitter.com/tokenterminal/status/1582376876143968256/photo/1)[^8] since September 2020.

To counteract market fragmentation, centralized crypto exchanges (CEXs) have rushed onto the scene. Over the last few years, [more than 200](https://coinmarketcap.com/rankings/exchanges/)[^9] CEXs have emerged worldwide. While many of them allow for the trading of digital assets across various blockchains, these CEXs remain isolated silos that trap liquidity. They all have their list of markets, which, unlike in traditional finance, are neither global nor aggregated. Some exchanges even choose a set of blockchains they more closely work with, amplifying the overall market's fragmentation.

The dominance of CEXs in crypto trading has been hailed as one of the [big ironies](https://www.coindesk.com/markets/2019/03/30/the-ultimate-irony-of-crypto-trading/)[^10] of a world that wants to achieve decentralization. Thus, decentralized crypto exchanges (DEXs) have gained in popularity and volume, at one point in time even [driving on-chain transaction volumes past](https://blog.chainalysis.com/reports/defi-dexs-web3/)[^11] that of centralized platforms.

Although DEXs may provide a number of benefits over CEXs in terms of censorship resistance or accessibility, they are not yet of [sufficient quality to compete](https://www.snb.ch/n/mmr/reference/sem_2022_06_03_barbon/source/sem_2022_06_03_barbon.n.pdf)[^12] with the largest CEXs. One of their main problems is the fact that a DEX's blockchain-enabled transparency lends itself to [front-running](https://link.springer.com/chapter/10.1007/978-3-030-43725-1_13)[^13]. Also, they lack speed and have a hard time competing regarding transaction costs and price efficiency.

Due to their on-chain architecture, DEXs remain unsuitable for high-frequency trading. Furthermore, when compared to centralized exchanges, there is a general lack of market depth. Market making is limited because, in the context of DEXs, liquidity providers factor in the volatility and security issues of decentralized exchange protocols and therefore deploy limited capital in accordance with their risk profiles. More generally, because the competition among DEXs is currently high, individual platforms struggle to retain liquidity as most liquidity providers only deploy [mercenary capital](https://www.mechanism.capital/native-token-liquidity/)[^14] that will instantly withdraw funds if competitors bait them with higher short-term returns.

A recent analysis purports to [show](https://blog.bitfinex.com/media-releases/hodlers-put-faith-in-centralised-exchanges-as-platforms-flex-high-tech-security/)[^15] that crypto traders are more comfortable trading on CEXs given the growing threat of hacks that have materialized in the decentralized finance (DeFi) space over the course of this year. In the aftermath of the [FTX collapse](https://www.investopedia.com/what-went-wrong-with-ftx-6828447)[^16], this could be changing, as the inherent count party risk coming with CEXs could increasingly be seen as a major disadvantage. While CEXs are striving to provide transparency by incorporating [proof of reserves](https://niccarter.info/proof-of-reserves/)[^17], in which a custodian transparently attests to the existence of on-chain reserve assets (but not liabilities), DEXs are working towards solving their [blockchain-based scalability issues](https://www.researchgate.net/publication/342639281_Scaling_Blockchains_A_Comprehensive_Survey)[^18].

### 1.3 Lack of decentralization

As the current situation reveals, even though decentralization is considered the driving force within the crypto space, it is far from being a reality. What is referred to as "[decentralization theater](https://www.imf.org/en/Publications/fandd/issues/2022/09/Defi-promise-and-pitfalls-Fabian-Schar)[^19]" creates a risk of deception as, in many cases, DeFi protocols remain heavily centralized.

As a matter of fact, even with most DEXs, there is no real decentralized trading. For example, some DEXs may still rely on a central entity to control the flow of buy and sell orders, which allows them to prevent users from placing orders. Others continue to use third-party accounts as a way to route transactions, making them not really non-custodial in nature. When it comes to their underlying technology blockchain, computing is not fully decentralized, as all machines re-compute the same code in order to reach consensus. Additionally, liquidity remains concentrated on a single chain. As such, trading on decentralized exchanges stands in contrast to something like the shipping industry, which exhibits true decentralization. A vast network of companies works together to ensure global shipping functions effectively.

### 1.4 Lack of regulatory frameworks

Undeniably, DEXs with centralized components must be subject to the regulatory standards that CEXs must follow. However, this is not the case as specific regulation pertaining to the crypto industry is still missing. Consequently, this lack of regulatory frameworks leads to the fact that there is no separation of responsibilities -- mostly among CEXs. In fact, they are doing a bit of everything: Managing their security in-house, doing their own custody, issuing their own stablecoin, launching investment products, or acting as a launchpad. This not only creates a conflict of interest since CEXs act as marketplaces, market makers, liquidity providers, and custodians, but with this concentration in services, a lot of trust ends up being placed in CEXs - the opposite of what the blockchain space is about.

The three problems described above have ailed the crypto industry ever since, leading to inefficient processes and high costs. Running an exchange across many blockchains and markets requires costly market makers that provide the necessary liquidity. The more fragmented the liquidity, the higher the costs to make markets.

As of now, DEXs lack any legal protections as they do not fall under any regulatory scrutiny and oversight so far. CEXs do abide by some regulations, however, they remain [shadow banks](https://corporatefinanceinstitute.com/resources/cryptocurrency/shadow-banking-and-cryptocurrencies/)[^20] in regard to traditional regulations. However, complying with local regulations comes with high costs for a CEX, especially when applying for business in many nations. Some CEXs choose to register in a country that allows them to operate with hardly any regulations, or they operate offshore and offer hardly any consumer protection. This exposes users to the platform's goodwill and can lead to further costs, once the CEX is forced to reduce or even cut user access to non-compliance with regulations. Last but not least, hacks resulting from backdoors due to insufficient protocol decentralization cause costs for users and the industry as a whole.

### 1.5 Solution for the problems

[Yellow Network](https://www.yellow.org/)[^21] is designed to solve these problems. The decentralized model that Yellow is envisioning is one, in which businesses work together, utilizing a shared backbone for liquidity, similar to the way thousands of internet service providers and network firms are interconnected and regulated in their respective countries.

By connecting brokers and exchanges, Yellow Network acts as a blockchain-agnostic mesh network of connected nodes that aggregate liquidity cross-chain, thereby increasing efficiency, lowering slippage, and allowing for best trades execution. To unlock high-speed trading, Yellow Network uses a layer-3 state channel infrastructure, allowing for communicating and trading between brokers and exchanges in a decentralized way. In other words, Yellow Network is a broker-to-broker liquidity channel infrastructure.

As such, Yellow Network is not based on any single blockchain, but a network of different intermediate nodes powered by Yellow Network code and run by brokers and exchanges connected to one another. This way, Yellow Network participants can do high-frequency trading among each other in a peer-to-peer way, using a smart clearing protocol to pool on-chain collateral that minimizes counterparty risks and protects broker-to-broker liabilities exchanged off-chain.

Yellow Network's centerpiece is its state channel smart clearing protocol, which is fully decentralized. The collateral that is locked in a state channel is controlled by a smart contract that is ultimately controlled by the parties that have opened the state channels. Thus, no funds are ever under the control of Yellow Network, making the system non-custodial from the perspective of its participants.

By creating a clearing and settlement system for communicating trading liabilities, updating states within state channels, and carrying out state channel challenges, Yellow Network is a combination of the [SWIFT](https://www.swift.com/about-us/discover-swift/messaging-and-standards)[^22] messaging protocol as well as an [ECN](https://www.angelone.in/knowledge-center/share-market/ecn-electronic-communication-network)[^23] direct order matching protocol. While SWIFT and ECN facilitate a broker's access to global financial markets of traditional markets, Yellow Network does the same for digital assets. And through this clear-cut setup, Yellow Network is contributing to the separation of powers, duties, and responsibilities by mimicking one essential part of the traditional finance stack -- the clearing house.

## 2. Yellow Network: High-level Explanation

To understand Yellow Network and its moving parts, it is important to get a high-level understanding of what is part of Yellow Network and what is not. We can best do this by differentiating three different layers of protocols:

1.  Blockchain layer protocol: This is where the state channels smart clearing protocol called ClearSync resides, responsible for the opening, monitoring, and closing of state channels between trading partners. This happens through the deposition and removal of collateral that serves as mutual protection in a trading relationship. While the collateral resides on the Ethereum blockchain, trading partners can trade any tokenized asset across various different blockchains.

2.  Margin call protocol: This is specific to Yellow Network. The protocol is responsible for updating the collateral state residing in the state channels agreed between trading partners through an off-chain RPC protocol.

3.  Trading protocol: This protocol is not specific to Yellow Network and is proprietary software. It can be FIX, Rest-JSON, Binance API, Bitfinex API, or any banking protocol for that matter.

![Stack Overview](./media/image5.png)

_Yellow Network's different layers visualized._

As we can see, Yellow Network is bringing a smart clearing protocol based on state channels, which is combined with an off-chain RPC protocol that is connected to proprietary trading software. At its core, Yellow Network is simply a protocol based on smart contract technology to track trading liabilities between Yellow Network trading parties. It is designed to be flexible, simple, and agnostic--- all in accordance with the [KISS design principle](https://www.interaction-design.org/literature/article/kiss-keep-it-simple-stupid-a-design-principle)[^24]. This way, existing institutions don't need to change their systems and can use their proprietary trading, settlement, and custody protocols to adopt Yellow Network.

## 3. How does Yellow Network work?

### 3.1 Overlay mesh network for high-frequency trading

As a clearing and settlement network, Yellow Network connects multiple brokers and exchanges into one shared orders platform. This platform spans multiple blockchains. By sitting on top of a network of blockchains, Yellow Network brings all parties like brokers, trading firms, exchanges, and blockchains together, thereby creating a decentralized, global trading network that gives rise to a more efficient trading infrastructure that supports high-frequency trading.

The connection points to each and every blockchain are established through custodians on all the various chains -- they act as gateways, enabling Yellow Network participants to seamlessly trade with one another across chains. The brokers using Yellow Network are connected to these different custodians, and the latter maintains a direct connection to the different blockchains, as they are the ones storing the on-chain assets on behalf of brokers. Thus, through its network of custodians, Yellow Network is integrated with every connected blockchain, which allows for asset virtualization in each chain.

End users that are willing to get a specific asset from any of the supported blockchains, need to withdraw it directly from the relevant custodian integrated with the specific blockchain. Importantly, brokers automate and hide this part away, the same way users can conveniently withdraw any cryptocurrency from Binance for example. This setup makes Yellow Network a layer-3 for fast clearing trades on top of a blockchain. It is not a blockchain itself, but uses blockchain technologies at its core.

Through this blockchain-agnostic layer-3 overlay mesh network of connected nodes run by brokers, exchanges, and trading firms that are aggregating liquidity cross-chain, overall market fragmentation is reduced, market depth is added, and best trade execution is made possible. Ultimately, smaller, more specialized trading venues profit from liquidity aggregation as it helps them to offer better pricing and handle larger trading volumes. Also, because exchanges, brokers, and trading firms are connected through one mesh network, end users can trade additional token pairings.

### 3.2 State channels infrastructure and multi-chain support

Yellow Network's state channel infrastructure and its multi-chain support are two different things that need to be distinguished -- yet they are both part of the overall system. Through its multi-chain support, Yellow Network is not based on a single blockchain but offers the ability to trade, clear, and settle assets in any tokenized asset on various blockchains. Thus, not only can Yellow Network participants trade assets across blockchain, they can trade cryptocurrencies, DeFi tokens, NFTs, or even real-world assets (RWA) tokens.

The state channel infrastructure makes up the backbone or nervous system of Yellow Network and allows for the locking and unlocking of collateral in state channels that are integrated into the Ethereum blockchain (it is possible that the Canary test will be on Polygon). Importantly, liabilities exchanged between state channel trading participants are not written into state channels themselves. It is only the collateral that resides within state channels that is on-chain.

The exchange of liabilities as well as the actual settlement and thus asset exchange is done outside the state channel setup through direct custodian-to-custodian settlement. To indicate that settlement has indeed occurred, both parties of a state channel need to signal agreement by signing the new balance sheet, thereby updating to the latest state within the state channel.

Should any Yellow Network participant refrain from conducting settlement according to the liabilities agreed between parties, collateral locked within a smart contract can be retrieved by the owed-to party. The so-called adjudicator smart contract is on-chain and is holding and releasing collateral based on protocol rules. All parties will agree to use this state channel smart contract protocol called ClearSync for adjudication. Assets held as collateral come in the form of ERC-20 tokens on Ethereum, which means that the clearing and settlement of collateral will only happen with Ethereum-based assets such as WBTC, WETH, USDT, USDC or DAI. However, Yellow Network might likely end up using stablecoins to correct for price volatility.

### 3.3 Decentralized trading through an aggregated order book

To facilitate decentralized trading, Yellow Network applies a multi-layered order book anatomy that aggregates order books across participants into one global network. Starting from their own local order book, each Yellow Network participant can synchronize their own order book with the overall network, giving rise to an aggregated Yellow Network order book.

With its network of brokers and custodians, which are integrated with today's wide variety of blockchains, and its state-channels-powered smart clearing protocol for collateral underwriting trading relationships, Yellow Network can allow for decentralized trading. All types of institutions can easily utilize Yellow Network with their own proprietary trading software, making it possible for them to not only trade digital assets but any other form of tokenized asset like gold, stocks, or bonds.

Making such a decentralized trading infrastructure more accessible and reliable will encourage the development of internet-native finance, giving rise to what is generally called the internet of value across the entire spectrum of today's economic actors and institutions. Not only will traditional players be able to provide products and services related to the internet of value, but such an internet native financial infrastructure will be enabling thousands of startups to tap into, utilize, and innovate on it. This will be similar to how Amazon Web Services spawned the development of the Web 2.0 ecosystem or Google Cloud is contributing to the development of artificial intelligence.

## 4. Technical basis of Yellow Network

### 4.1 Why state channels?

To achieve mass adoption, blockchain scalability is one of the key developments within the crypto industry. As it stands, the Ethereum Network handles \~15 TPS (transactions per second), a far cry from what is needed to efficiently and reliably accommodate millions of users. Still, as the biggest smart contract blockchain, Ethereum is building out a scalable unified settlement and data availability layer.

In doing this, Ethereum [has pivoted](https://members.delphidigital.io/reports/the-hitchhikers-guide-to-ethereum)[^25] to a rollup-centric roadmap. This fact is [reflected](https://www.theblock.co/data/scaling-solutions/scaling-overview/value-locked-of-ethereum-scaling-solutions)[^26] in the value locked in Ethereum scaling solutions, which are dominated by the rollup type. Nonetheless, Yellow Network operates on state channel technology, because its decentralized, global trading network is optimized for low-latency and high-frequency, attributes rollups cannot provide.

In comparison to state channels, rollups are rather limited in what they can optimize for. For example, they are not optimized for ERC-20 transfers, which the majority of transaction volume is in. [Rollups](https://members.delphidigital.io/reports/the-hitchhikers-guide-to-ethereum)[^27] work by moving heavy and resource-intensive smart contract computations from the main chain to a separate layer, or "roll-up". This offloading reduces the load on the main chain and allows for faster processing of transactions. However, ERC20 transfers, which only involve state modifications of a balance, are not considered heavy computations. Thus, rollups do not offer any advantages in that regard.

Furthermore, the performance of rollups is nowhere near what state channels offer. In fact, they are not only slower, but not as scalable as state channels. With state channels, one is limited by hardware and network latency when it comes to transaction throughput. Rollups, on the other hand, need a separate architecture for advanced verification. Also, rollups are not cross-chain, while state channels are. And last but not least, state channels allow transactions between two or more parties without involving any third party. This is unlike rollups, whose transactions are [processed](https://ethereum.org/en/developers/docs/scaling/optimistic-rollups/)[^28] by an operator at present.

The high performance, the absence of third parties, the suitability for a wider variety of use cases as well as the cross-chain aspect, are all attributes Yellow Network is maximizing for, which is why state channel technology is used as the preferred scaling solution type.

### 4.2 What are state channels?

The concept of state channels was [introduced](https://ieeexplore.ieee.org/document/9627997)[^29] in 2015. It is an off-chain scaling mechanism that enables transacting parties to interact without touching the blockchain. Only the final state between them is broadcasted and settled on-chain. The most prominent peer-to-peer state channel network in action is the [Bitcoin Lightning Network](https://lightning.network/)[^30], a second-layer protocol that was [introduced](https://lightning.network/lightning-network-paper.pdf)[^31] in 2016. It is based on payment channels, which are payment-specific state channel applications. Nonetheless, the Bitcoin Lightning Network is an excellent demo of how state channels work.

Typically, a channel can be built upon [threshold signatures](https://ieeexplore.ieee.org/document/4118696)[^32]--- often referred to as multi-signature or multisig --- and instructions for timelocks. Channel participants sign a multi-signature smart contract and lock in funds to participate. Following the lock-up of funds, any number of states can be exchanged among participants. Such states can [represent](https://arxiv.org/pdf/2204.08032.pdf)[^33] any generalized application, enabling state channels to enable many use cases. Because state exchanges happen off-chain, they allow for fast transactions and high transaction throughput.

In the context of trading for example, traders can make x-number of off-chain transactions between each other during the day, and by the evening, all parties involved can mutually settle trades by officially posting the final net balance between them. Every trader would get what is due to him.

### 4.3 The Benefits of State Channels

As a blockchain scaling solution, state channels [offer](https://docs.ethhub.io/ethereum-roadmap/layer-2-scaling/state-channels/)[^34] various benefits. Among them are:

- State channels are ultra-fast and massively scalable

- State channels are blockchain agnostic

- State channels offer confidentiality

- State channels are Web2 compatible

Yellow Network has identified that state channels, thanks to their advantages, provide the necessary solution to address the current challenges Web2 companies are facing when wanting to interact with blockchain technology. Because of path dependency, it is unlikely that today\'s [highly developed internet infrastructure](https://www.internetlivestats.com/)[^35], [with its](https://everysecond.io/the-internet)[^36] billions of websites, millions of servers, and thousands of Web2 companies, will go through a complete overhaul to become blockchain-compatible. It is much more likely that state channels will act as a middleware solution between blockchains and legacy information systems.

Through its state channel based layer-3 technology, Yellow Network enables Web2 companies to have a more straightforward way to connect legacy internet infrastructure to the new internet of value powered by blockchains. This way, smart contracts will make up a bigger part of existing web companies and become easier to operate for them. At the same time, thanks to Yellow Network's integrated state channels technology, these companies do not need to wait for node validation or block creation when having blockchain-based transactions processed.

Instead, any Web2 company can easily integrate with blockchains and use them in their day-to-day business efforts. For anyone to become a peer in Yellow Network's peer-to-peer (P2P) trading platform, they just need to open a state channel with other participants through a Yellow Network node.

## 5. Yellow Network's Solution explained and Architecture deep dive

![Virtual State Channels](./media/image4.png)

_Yellow Network's overlay mesh network of Yellow network nodes._

### 5.1 Running a Yellow Network Node

To make Yellow Network's peer-to-peer trading possible, its participants are connected through Yellow Network nodes, giving rise to a broad trading environment consisting of multiple players providing and pooling liquidity on the platform. A Yellow Network node is a high-performance node that each of the participants is running. They make up a core component of the Yellow Network trading platform and are responsible for searching other Yellow Network nodes, so Yellow Network participants can communicate with each other.

Furthermore, Yellow Network nodes are crucial for setting up state channels, locking collateral, and making the logic eligible for both state channel parties. For communication among Yellow Network nodes, a protocol like [LibP2P](https://libp2p.io/) is most likely used. This protocol gives nodes the ability to discover other nodes and send messages to them. As such, it is a way to connect servers with one another in a mesh network, thereby giving rise to Yellow's own peer-to-peer network.

### 5.2 A world map of Yellow Network's Liquidity

To facilitate decentralized trading, Yellow Network relies on an aggregated orderbook that is shared across network participants. This shared orderbook has unlimited capacity for the participants' requests. Because the orderbook is hosted with individual participants, it is resistant to any work interruptions. Should any of the many network participants go down, the globally shared orderbook would still be accessible to traders on Yellow Network.

This aggregated orderbook is established, when Yellow Network participants synchronize their own orderbook with the network, giving rise to the aggregated order book of Yellow Network. "Aggregated" means that Yellow Network collects and orders feeds that Yellow Network participants get from other participants, their counterparties so to say.

This is accomplished by utilizing LibP2P to create a peer-to-peer network among Yellow Network nodes that facilitates communication and discovery. Within LibP2P, the [publish-subscribe](https://docs.libp2p.io/concepts/pubsub/overview/)[^37] feature is utilized to relay real-time price quotes happening across the network, providing a live stream map of liquidity and prices globally.

This information can be used as a routing table to decide the optimal connections among nodes. Upon connecting to the network, each node subscribes to the quotes it wishes to track, for example, WBTC/USDT. It then receives price updates from all nodes in the global network and can compute a routing table based on the state channels it has already established. A Yellow network node may then decide to create additional direct or virtual channels to access interesting prices or spreads.

![Topic PubSub](./media/image1.png)

_Yellow Network provides a world map of liquidity participants can tap into._

So, by joining Yellow Network, every participant is provided with a world map of liquidity, showing where it is, what prices there are and how this liquidity can be accessed most efficiently. This is similar to how [BitTorrent](https://web.cs.ucla.edu/classes/cs217/05BitTorrent.pdf)[^38] maps all files. Yellow Network does the same, not for files, but for trading data. Joining Yellow Network gives you access to this map, making it easier to find and access liquidity.

### 5.3 State Channels Smart Clearing Protocol: ClearSync

At the heart of Yellow Network is its smart clearing protocol. It exists as a set of automated smart contracts that allow participants to lock and unlock collateral through state channels to protect the clearing and settlement of assets.

It is this state channels protocol that allows Yellow Network participants to minimize risks when exchanging trading liabilities or assets owned and owed from one participant to another. When an exchange trades with another broker, they both exchange liabilities using Yellow Network's protocol, ensuring mutual consent regarding the accounting side of things.

The distribution ratio within the state channels operated by participants is updated on an ongoing basis (preferably every second). This way, both participants can always monitor the most current state within their channel, thereby preventing either participant from being defrauded. Should counterparty settlement risk become imbalanced, the trading party carrying the risk can ask the other party to provide more collateral using a margin call.

![](./media/image3.png)

_Brokers can trade with one another, while Yellow Network is acting as the clearing house between them._

Thus, Yellow Network also provides the option for unilateral collateral settlement, whereby if a broker's balance sheet is \$5,000 higher than that of his trading partner, inequalities can be rebalanced by having the trading partner with the lower balance sheet provide the necessary collateral to settle the difference and re-equalize the trading liabilities. This process would result in a partial settlement, enabling both brokers to readjust the margin of collateral and continue trading.

Should either party not agree to rebalance or settle the liabilities altogether, the defrauded party can refer to the latest state within the state channel and have the smart contract release the collateral in their favor. As a result, the party that has been defrauded will get the necessary funds to cover up for what is entitled to them. The participants will be able to buy the assets that they were not provided with through the other party's settlement.

One has to keep in mind that the actual trading, clearing, and settlement of assets between Yellow Network participants are not directly initiated by the state channels protocol. It is done by either broker, exchange, or trading firm, thereby moving owed funds between one another. This means that technically, liabilities are not written into state channels, but kept in the database of the respective parties. They use a system of gossip-to-gossip communication to determine the amount of collateral required to continue their business activities and ensure that both parties and any potential risks are protected.

![](./media/image2.png)

_Three crucial components making up Yellow Network: Trading and Clearing (off-chain) as well as settlement (on-chain)._

It is only the collateral that resides within state channels. Therefore, state channels do neither know about the liabilities that are being traded nor do they know about their prices. They just know the distribution ratio of the collateral provided and signed by the different state channel participants. Liabilities will be exchanged off-chain using Yellow Network's protocol, not on-chain transaction. However, the state channels smart clearing protocol Clear Sync helps accurately monitor and thus rebalance collateral positions atomically between trading partners.

To perform actual settlement of on-chain assets, brokers will have different ways to proceed:

- Using proper off-chain database and API infrastructure

- Using on-chain escrow smart contract or HTLC

Given two trading partners within Yellow Network have the same custodians, the assets don\'t need to be moved on-chain, but can just be cleared by updating the mutual custodian's accounting database. This is similar to how money is moved between people, who have the same bank and transfer money amongst each other. While the owner changes and accounting states are adjusted, no money is actually moved around.

Using on-chain escrow smart contracts or HTLC instead allows for [atomic settlement](https://r3.com/everyday-blockchain/atomic-settlement-if-you-have-amazon-prime-you-already-understand-the-process/)[^39]. HTLCs are referred to as [hash time-locked contracts (HTLCs)](https://medium.com/hackernoon/what-are-hashed-timelock-contracts-htlcs-application-in-lightning-network-payment-channels-14437eeb9345)[^40] and are popularly used in the context of the Lightning Network. They can be described as a type of smart contract that ensures that a transfer between parties, performed before an expiration time, allows these parties to carry out an atomic swap of assets without a third party. [Using an HTLC](https://www.researchgate.net/publication/358898825_Generalized_HTLC_for_Cross-Chain_Swapping_of_Multiple_Assets_with_Co-Ownerships)[^41] for the settlement process can move funds from one broker\'s custody to another on an atomic level and on multiple blockchains.

## 6. How to Participate in Yellow Network

To participate in Yellow Network, participants have to provide and lock Yellow Network tokens. To fund the collateral in a state channel, participants need to supply the collateral in the form of stablecoins. The amount of collateral deposited into a state channel is called a [state deposit](https://eprint.iacr.org/2019/219)[^42]. Once both parties have an agreement on the collateral amount and they both have deposited the stablecoins to the smart contract, funds will be locked in the adjudicator smart contract called ClearSync. This smart contract sits in between Yellow Network participants. Through the adjudicator smart contract, Yellow Network's smart clearing protocol takes care of brokers\' collateralization levels and liability risk management in real time.

With the funds locked by the adjudicator smart contract, the state channel is opened, and its participants can start trading. With every act of trading, the state channel receives state updates from either participant, upgrading the state channel according to the most recent trading activity. Closing a state channel is done by producing a final state transition. This will cause an automatic settlement of all open positions, and clear liabilities, and lead to the unlocking of each participant's collateral if no dispute about the final state within the channel is in progress.

## 7. Yellow Network Application

As already indicated, there will be an application to run the Yellow Network node, so interested parties can run it on their own servers. This software will handle network-related activities like communicating liabilities, updating states within state channels, and carrying out state channel disputes. Through this gateway, brokers, exchanges, and trading firms can become Yellow Network participants. They are the main target audience that Yellow Network is trying to reach.

While big exchanges and trading firms are welcome to join and profit from the network effects of Yellow Network, and Openware will work on onboarding big players that bring liquidity as well as users, it is smaller- and middle-tier exchanges that represent ideal participants of Yellow Network. Because of increased competition and liquidity fragmentation, smaller players might not be able to keep up with the fierce competition in the mid to long run, making them ideal candidates to use Yellow Network.

Thus, Yellow Network offers a survival solution to them by letting them team up and combine their powers with other exchanges. By uniting different exchange small- and middle-tier exchanges through Yellow Network, this network of exchanges can gain more liquidity, making them fit to compete with behemoth exchanges like Binance or Huobi. As such, Yellow Network's main value proposition is giving smaller to middle-tier exchanges the possibility to sell their liquidity to the network, while simultaneously profiting from the liquidity of other Yellow Network participants, thereby staying competitive.

## 8. Technology Partners

### 8.1. Openware with OpenDAX

[Openware](https://www.openware.com/)[^43], Inc. is a United States multinational blockchain infrastructure development company headquartered in South San Francisco, California, that designs, develops, and sells computer software, and online services. Their focus is on building secure and scalable solutions for Web3 and the internet of finance.

[OpenDAX](https://www.openware.com/product/opendax)[^44] is Openware's flagship product and stands for \'Open-Source Digital Assets Exchange\'. It is a hybrid software consisting of public and private libraries, designed to build a fully-featured exchange service to facilitate the trading of digital assets, cryptocurrencies, and security tokens. OpenDAX^TM^ is cloud-based and while it comes as a plug-and-play solution that can easily be deployed, there is the possibility to customize the solution in accordance with one's needs.

Through operating and selling OpenDAX^TM^, the need for Yellow Network has become apparent. The idea for starting Yellow Network came from Openware's clients that have been starting their own exchange using OpenDAX^TM^, they have always been facing the issue of raising and obtaining sufficient liquidity.

Building out Yellow Network as a built-in solution and natively supported platform for OpenDAX^TM^ users would help brokers and clients efficiently source deep pockets of liquidity, thereby solving their bootstrapping problem that has so far hampered the adoption of OpenDAX^TM^ and thus the creation of more brokers and exchanges within the crypto trading space. While the need for some basic market-making will persist, the creation of Yellow Network will simplify things and cut costs for brokers and smaller exchanges that want to do business in the crypto trading space.

[^1]: Nakamoto Satoshi, "Bitcoin: A Peer-to-Peer Electronic Cash System." [www.bitcoin.org](http://www.bitcoin.org), October 2008: <https://bitcoin.org/bitcoin.pdf>
[^2]: Buterin Vitalik, "Ethereum White paper." [www.ethereum.org](http://www.ethereum.org), November 2014: <https://ethereum.org/en/whitepaper/>
[^3]: Antonopouos Andreas, Wood Gavin, "Mastering Ethereum." O\'Reilly Media, Inc., November 2018: <https://www.oreilly.com/library/view/mastering-ethereum/9781491971932/>
[^4]: Buterin Vitalik, "Why sharding is great: demystifying the technical properties." [www.vitalik.ca](http://www.vitalik.ca), April 2021: <https://vitalik.ca/general/2021/04/07/sharding.html>
[^5]: Blockchain Comparison: <https://blockchain-comparison.com/blockchain-protocols/>
[^6]: L2beat: <https://l2beat.com/scaling/tvl/>
[^7]: Chainalysis Team, "Vulnerabilities in Cross-chain Bridge Protocols Emerge as Top Security Risk." Chainalysis Blog, August 2022: <https://blog.chainalysis.com/reports/cross-chain-bridge-hacks-2022/>
[^8]: Token Terminal, "Bridge exploits account for \~50% of all DeFi exploits, totaling \~\$2.5B in lost assets." Token Terminal Twitter, October 2022: <https://twitter.com/tokenterminal/status/1582376876143968256>
[^9]: Coinmarketcap, "Top Cryptocurrency Spot Exchanges." Coinmarketcap Website: <https://coinmarketcap.com/rankings/exchanges/>
[^10]: Weisberger David, "The Ultimate Irony of Crypto Trading." Coindesk, September 2021: <https://www.coindesk.com/markets/2019/03/30/the-ultimate-irony-of-crypto-trading/>
[^11]: Chainalysis Team, "DeFi-Driven Speculation Pushes Decentralized Exchanges' On-Chain Transaction Volumes Past Centralized Platforms." Chainalysis Blog, June 2022: <https://blog.chainalysis.com/reports/defi-dexs-web3/>
[^12]: Barbon Andrea , Ranaldo Angelo, " On The Quality Of Cryptocurrency Markets: Centralized Versus Decentralized Exchanges." SNP Paper, April 2022: <https://www.snb.ch/n/mmr/reference/sem_2022_06_03_barbon/source/sem_2022_06_03_barbon.n.pdf>
[^13]: Eskandari, S., Moosavi, S., Clark, J.: SoK: Transparent Dishonesty: Front-Running

    Attacks on Blockchain. In: Financial Cryptography. pp. 170--189. Springer International Publishing, Cham (2020): <https://link.springer.com/chapter/10.1007/978-3-030-43725-1_13>

[^14]: Wu, Eva, "The Art and Science of Native Token Liquidity" Mechanism Capital, August 2021: <https://www.mechanism.capital/native-token-liquidity/>
[^15]: Bitfinex, "Hodlers Put Faith in Centralised Exchanges as Platforms Flex High-Tech Security." Bitfinex Blog, October 2022: <https://blog.bitfinex.com/media-releases/hodlers-put-faith-in-centralised-exchanges-as-platforms-flex-high-tech-security/>
[^16]: Reiff Nathan, "The Collapse of FTX: What Went Wrong with the Crypto Exchange?" Investopedia, December 2022: <https://www.investopedia.com/what-went-wrong-with-ftx-6828447>
[^17]: Carter Nic, "Nic's PoR: Wall of Fame" Nic Carter Website: <https://niccarter.info/proof-of-reserves/>
[^18]: Hafid, Abdelatif, Senhaji Hafid Abdelhakim, Samih Mustapha, "Scaling Blockchains: A Comprehensive Survey." IEEE Access, July 2020: <https://www.researchgate.net/publication/342639281_Scaling_Blockchains_A_Comprehensive_Survey>
[^19]: Schär Fabian, "DeFi's Promise and Pitfalls." IMF, September 2022: <https://www.imf.org/en/Publications/fandd/issues/2022/09/Defi-promise-and-pitfalls-Fabian-Schar>
[^20]: CFI Team, "What is Shadow Banking in the Cryptocurrency World?", Corporate Finance Institute, February 2023: <https://corporatefinanceinstitute.com/resources/cryptocurrency/shadow-banking-and-cryptocurrencies/>
[^21]: Yellow Network: <https://www.yellow.org/>
[^22]: Discover Swift: <https://www.swift.com/about-us/discover-swift/messaging-and-standards>
[^23]: What is ECN: <https://www.angelone.in/knowledge-center/share-market/ecn-electronic-communication-network>
[^24]: Kiss Design Principle: <https://www.interaction-design.org/literature/article/kiss-keep-it-simple-stupid-a-design-principle>
[^25]: Charbonneau, Jon, "The Hitchhiker\'s Guide to Ethereum." Delphi Digital, May 2022: <https://members.delphidigital.io/reports/the-hitchhikers-guide-to-ethereum>
[^26]: The Block: "Value Locked in Ethereum Scaling Solutions by Type." <https://www.theblock.co/data/scaling-solutions/scaling-overview/value-locked-of-ethereum-scaling-solutions>
[^27]: Thibault Louis Tremblay, Sarry Tom, Hafid Abdelhakim Senhaji, "Blockchain Scaling Using Rollups: A Comprehensive Survey" IEEE Access, August 2022: <https://ieeexplore.ieee.org/stamp/stamp.jsp?arnumber=9862815>
[^28]: Isthedoom, "Optimistic Rollups." [www.ethereum.org](http://www.ethereum.org), November 2022: <https://ethereum.org/en/developers/docs/scaling/optimistic-rollups/>
[^29]: Negka Lydia, Spathoulas Georgios, "Blockchain State Channels: A State of the Art." IEEE Access, November 2021: <https://ieeexplore.ieee.org/document/9627997>
[^30]: Lightning Network: <https://lightning.network/>
[^31]: Poon Joseph, Dryja Thaddeus, "The Bitcoin Lightning Network: Scalable Off-Chain Instant Payments." January 2016: <https://lightning.network/lightning-network-paper.pdf>
[^32]: Van Der Merwe Johann, Dawoud, McDonald Stephen, "A Fully Distributed Proactively Secure Threshold-Multisignature Scheme." IEEE Access, March 2007: <https://ieeexplore.ieee.org/document/4118696>
[^33]: Gangwal Ankit, Gangavalli Haripriya, Thirupathi Apoorva, "A Survey of Layer-Two Blockchain Protocols." Elsevier Journal of Network and Computer Applications, April 2022: <https://arxiv.org/pdf/2204.08032.pdf>
[^34]: "State Channels." Docs EthHub: <https://docs.ethhub.io/ethereum-roadmap/layer-2-scaling/state-channels/>
[^35]: Internet Live Stats: <https://www.internetlivestats.com/>
[^36]: The Internet -- Every Second: <https://everysecond.io/the-internet>
[^37]: Docs.libp2p - What is Publish/Subscribe: <https://docs.libp2p.io/concepts/pubsub/overview/>
[^38]: Johnsen Jahn, Karlsen Lars, Birkeland Sebjørn, "Peer-to-peer networking with BitTorrent." Department of Telematics, December 2005: <https://web.cs.ucla.edu/classes/cs217/05BitTorrent.pdf>
[^39]: R3, "Atomic settlement: if you have Amazon Prime, you already understand the process", February 2022: <https://r3.com/everyday-blockchain/atomic-settlement-if-you-have-amazon-prime-you-already-understand-the-process/>
[^40]: Vohra Arnav, "What Are Hashed Timelock Contracts (HTLCs)? Application In Lightning Network & Payment Channels." [www.hackernoon.com](http://www.hackernoon.com), May 2018: <https://medium.com/hackernoon/what-are-hashed-timelock-contracts-htlcs-application-in-lightning-network-payment-channels-14437eeb9345>
[^41]: Krishnasuri Narayanam, Ramakrishna Venkatraman, Dhinakaran Vinayagamurthy, Sandeep Nishad, "Generalized HTLC for Cross-Chain Swapping of Multiple Assets with Co-Ownerships." February 2022: <https://www.researchgate.net/publication/358898825_Generalized_HTLC_for_Cross-Chain_Swapping_of_Multiple_Assets_with_Co-Ownerships>
[^42]: Close Tim, "Nitro Protocol" February 2019: <https://eprint.iacr.org/2019/219>
[^43]: Openware: <https://www.openware.com/>
[^44]: OpenDAX: <https://www.openware.com/product/opendax>
