# Operator API 1.0.0 documentation

What you can do with the Operator API


## Table of Contents

* [Servers](#servers)
  * [private](#private-server)
* [Operations](#operations)
  * [PUB /](#pub--operation)
  * [SUB /](#sub--operation)

## Servers

### `private` Server

* URL: `not.deployed.yet`
* Protocol: `wss`

The private server is used for internal communication between the Operator and the Operator API.
It is not exposed to the public.



## Operations

### PUB `/` Operation

* Operation ID: `handleMessage`

Send a message to the Operator API.


Accepts **one of** the following messages:

#### Message `authenticate`

*Authenticate with the Operator API*

Authenticate with the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"authenticate"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.challenge | string | - | - | - | - |
| data.signature | string | A signature of the challenge using the private key of the peer. | - | - | - |

> Examples of payload

```json
{
  "event": "authenticate",
  "reqid": 1,
  "data": {
    "challenge": "4ffc0b7f-c93c-4b99-bfa7-7c9886c7fdd2",
    "signature": "0x8e62e91f99b86caa46ea56c4286545d7b63f0e5121c37d319d9bd46076acad4b14cef0e5ffca060c0378ef6135a2ef5efd6617555ee84f6d0890dd5aa5534e011c"
  }
}
```


#### Message `getChannelJwt`

*Request a JWT for access to Finex*

Request a JWT for a channel from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"channel_jwt"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |

> Examples of payload

```json
{
  "event": "channel_jwt",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `openChannel`

*Open a channel*

Open a channel with a peer.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"open_channel"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.peer | object | - | - | - | **additional properties are allowed** |
| data.peer.address | string | - | - | - | - |
| data.peer.url | string | - | - | - | - |
| data.peer.name | string | - | - | - | - |
| data.margin_deposit | string | A decimal string of arbitrary precision. | - | - | - |

> Examples of payload

```json
{
  "event": "open_channel",
  "reqid": 1,
  "data": {
    "peer": {
      "address": "0x55608cdbde2ff4183a81e62da096fa863d8f910d",
      "url": "https://example.com",
      "name": "Alice"
    },
    "margin_deposit": "10000"
  }
}
```


#### Message `recordTrade`

*Record a trade*

Record a trade in the system.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"record_trade"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | A decimal string of arbitrary precision. | - | - | - |
| data.external_id | string | - | - | - | - |
| data.market | string | Market name in format `<BASE>/<QUOTE>`. | - | - | - |
| data.direction | string | - | allowed (`"buy"`, `"sell"`) | - | - |
| data.amount | string | A decimal string of arbitrary precision. | - | - | - |
| data.price | string | A decimal string of arbitrary precision. | - | - | - |
| data.executed_at | integer | - | - | format (`int64`) | - |

> Examples of payload

```json
{
  "event": "record_trade",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "external_id": "123",
    "market": "ETH/USD",
    "direction": "buy",
    "amount": "100",
    "price": "1000",
    "executed_at": 1590000000
  }
}
```


#### Message `subscribeNotifications`

*Subscribe to notifications*

Subscribe to notifications from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"subscribe_notifications"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |

> Examples of payload

```json
{
  "event": "subscribe_notifications",
  "reqid": 1
}
```


#### Message `closeChannel`

*Close a channel*

Close a channel with a peer.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"close_channel"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |

> Examples of payload

```json
{
  "event": "close_channel",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `getPositions`

*Request for positions*

Fetches list of all positions by channel

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"positions"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "positions",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `requestSettlement`

*Request Settlement*

Initiates a settlement request

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"request_settlement"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |
| data.payment_method | string | - | allowed (`"escrow"`, `"mock"`) | - | - |
| data.chain_id | integer | - | - | format (`int64`) | - |
| data.markets | array<string> | - | - | - | - |
| data.markets (single item) | string | Market name in format `<BASE>/<QUOTE>`. | - | - | - |

> Examples of payload

```json
{
  "event": "request_settlement",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "payment_method": "escrow",
    "chain_id": 1,
    "markets": [
      "WBTC/USDT",
      "WETH/USDT"
    ]
  }
}
```



### SUB `/` Operation

* Operation ID: `respondOrNotify`

Receive a message from the Operator API.


Accepts **one of** the following messages:

#### Message `version`

*Receive the API version*

Receive the API version from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"version"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.version | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "version",
  "reqid": 1,
  "data": {
    "version": "1.0.0"
  }
}
```


#### Message `identityAddress`

*Receive an identity address*

Receive an identity address from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"identity_address"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.address | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "identity_address",
  "reqid": 1,
  "data": {
    "address": "0x55608cdbde2ff4183a81e62da096fa863d8f910d"
  }
}
```


#### Message `challenge`

*Receive an auth challenge response*

Response to an auth challenge from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"challenge"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.challenge | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "challenge",
  "reqid": 1,
  "data": {
    "challenge": "4ffc0b7f-c93c-4b99-bfa7-7c9886c7fdd2"
  }
}
```


#### Message `authenticated`

*Successfully authenticated*

Successfully authenticated with the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"authenticate"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |

> Examples of payload

```json
{
  "event": "authenticate",
  "reqid": 1
}
```


#### Message `jwt`

*Receive a JWT*

Respond with a JWT to the Operator.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"channel_jwt"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.jwt | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "channel_jwt",
  "reqid": 1,
  "data": {
    "jwt": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJkb21haW4iOiJzb21lLmRvbWFpbi5jb20iLCJlbWFpbCI6InNvbWVAZW1haWwuY29tIn0.1YXe-ffi3lCr_hLa4585_TIXhGLxMhfA2M8uhpnJVQQQgfTKAaN1VASvaV7NhhSQly2gJfzNrTYecyQJQMKe4Q"
  }
}
```


#### Message `openingChannel`

*Receive a channel opening response*

Receive a channel opening response from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"open_channel"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |

> Examples of payload

```json
{
  "event": "open_channel",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `recordedTrade`

*Recorded a trade*

Recorded a trade in the system.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"record_trade"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |
| data.external_id | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "record_trade",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "external_id": "123"
  }
}
```


#### Message `streamNotification`

*Receive a stream update notification*

Receive a stream update notification from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"stream_notification"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |
| data.action | string | - | allowed (`"channelOpened"`, `"channelClosed"`, `"channelSettled"`, `"marginUpdated"`) | - | - |
| data.channel_status | string | - | allowed (`"opening"`, `"opened"`, `"settling"`, `"closing"`, `"closed"`) | - | - |
| data.my_role | string | - | allowed (`"Initiator"`, `"Responder"`) | - | - |
| data.peer | object | - | - | - | **additional properties are allowed** |
| data.peer.address | string | - | - | - | - |
| data.peer.url | string | - | - | - | - |
| data.peer.name | string | - | - | - | - |
| data.margin_deposit | string | A decimal string of arbitrary precision. | - | - | - |
| data.initiator_margin_balance | string | A decimal string of arbitrary precision. | - | - | - |
| data.follower_margin_balance | string | A decimal string of arbitrary precision. | - | - | - |

> Examples of payload

```json
{
  "event": "stream_notification",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "action": "channelOpened",
    "channel_status": "opened",
    "my_role": "Initiator",
    "peer": {
      "address": "0x55608cdbde2ff4183a81e62da096fa863d8f910d",
      "url": "https://example.com",
      "name": "Alice"
    },
    "margin_deposit": "10000",
    "initiator_margin_balance": "0",
    "follower_margin_balance": "0"
  }
}
```


#### Message `settlementNotification`

*Receive a settlement update notification*

Receive a settlement update notification from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"settlement_notification"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |
| data.state | string | - | allowed (`"proposed"`, `"accepted"`, `"initiated"`, `"prepared"`, `"executed"`, `"completed"`, `"failed"`, `"rejected"`) | - | - |

> Examples of payload

```json
{
  "event": "settlement_notification",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "settlement_state": "proposed"
  }
}
```


#### Message `positionNotification`

*Receive a position update notification*

Receive a position update notification from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"position_notification"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | - | - | - | - |
| data.market | any | - | - | - | **additional properties are allowed** |
| data.direction | string | - | allowed (`"buy"`, `"sell"`) | - | - |
| data.amount | string | A decimal string of arbitrary precision. | - | - | - |
| data.cost | string | A decimal string of arbitrary precision. | - | - | - |
| data.market_value | string | A decimal string of arbitrary precision. | - | - | - |
| data.pnl | string | A decimal string of arbitrary precision. | - | - | - |
| data.status | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "position_notification",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `closeChannel`

*Close a channel*

Close a channel with a peer.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"close_channel"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | Hexadecimal representation of channel ID (length should be 32 bytes). | - | - | - |

> Examples of payload

```json
{
  "event": "close_channel",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62"
  }
}
```


#### Message `positions`

*List of positions*

Contains a list of all positions by requested channel id

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"positions"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.channel_id | string | A decimal string of arbitrary precision. | - | - | - |
| data.positions | array<any> | - | - | - | - |

> Examples of payload

```json
{
  "event": "positions",
  "reqid": 1,
  "data": {
    "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
    "positions": [
      {
        "position": {
          "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
          "market": "WBTC/USDT",
          "direction": "buy",
          "amount": "1.2",
          "cost": "36789",
          "market_value": "35467",
          "pnl": "-1322",
          "status": "open"
        }
      },
      {
        "position": {
          "channel_id": "0x55608cdbde2ff4183a81e62da096fa863d8f910d29d17826124fccc9bcc11f62",
          "market": "ETH/USDT",
          "direction": "sell",
          "amount": "1.5",
          "cost": "3678.9",
          "market_value": "3546.7",
          "pnl": "132.2",
          "status": "open"
        }
      }
    ]
  }
}
```


#### Message `requestedSettlement`

*Requested Settlement*

Response to a settlement request

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"request_settlement"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |

> Examples of payload

```json
{
  "event": "request_settlement",
  "reqid": 1
}
```


#### Message `error`

*Receive an error*

Receive an error from the Operator API.

##### Payload

| Name | Type | Description | Value | Constraints | Notes |
|---|---|---|---|---|---|
| (root) | object | - | - | - | **additional properties are allowed** |
| event | string | - | const (`"error"`) | - | - |
| reqid | integer | client originated ID reflected in response message. | - | - | - |
| data | object | - | - | - | **additional properties are allowed** |
| data.error | string | - | - | - | - |

> Examples of payload

```json
{
  "event": "error",
  "reqid": 1,
  "data": {
    "error": "Invalid signature"
  }
}
```



