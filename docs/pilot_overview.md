# Pilot Release Architecture

![pilot_overview](./pilot_overview.png)

## Sub-Systems

### Clearport

The component handle all state-channels operations, blockchain funding and defunding and has a built-in EscrowPayment mechanism for Settlement of Trades. It handle the communication point-to-point for Channel state exchange.

This service can be working locally on desktop, or remotely on a cloud for production environments. It manipulates funds and the use of Key Management Service is mandatory.

#### Operator API

GRPC API Exposed for Clearport user.

Terminal, clients, or server such as NeoDAX must authenticate to the Operator API using a whitelisted public key signature and communicate using a localhost or trusted SSL domain.

Through the Operator API, client and server can make the following operations:

1. Open a Clearing channel and fund it on chain.
2. Get the margin balance associated to a channel ID
3. Get all the positions
4. Request a Settlement on those positions
5. Disconnect and Force Close channels
6. Record new Trades for initiating the clearing and later settlement process.

### User Interface

React/Typescript user interface working on several platforms such as desktop (MacOSX, Windows), but also iOS/Android or Web Application. The User interface allow the following features:

- Open a clearing channel, by depositing stablecoin (USDT) and YELLOW Tokens
- Create Limit/Market Order on the connected channel Peer matching engine
- List unsettled positions (Trades)
- Request a settlement for the positions
- Close channels which also close all positions and requires a final settlement.

### NeoDAX

Broker-Dealer microservice accepting orders per channels. The Limits of order size and trades depends on the margin utilization on the clearing system.

Margin account in NeoDAX is initialized from a JWT at the authentification steps, The Margin-Account is then updated in memory for the duration of the trading session.

Once the JWT expire, the trading session is also expired and requires a new token (with most recent margin). We can also assume clearsync will need to send a regular heartbeat with the updated margin according to market prices and settlements states.