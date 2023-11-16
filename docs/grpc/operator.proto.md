# gRPC Protocol Documentation

This file describes the content of operator.proto.


### Operator

protolint:disable MAX_LINE_LENGTH

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVersion | [VersionRequest](#versionrequest) | [VersionResponse](#versionresponse) |  |
| GetIdentityAddress | [GetIdentityAddressRequest](#getidentityaddressrequest) | [GetIdentityAddressResponse](#getidentityaddressresponse) |  |
| GetChallenge | [GetChallengeRequest](#getchallengerequest) | [GetChallengeResponse](#getchallengeresponse) |  |
| Authenticate | [AuthenticateRequest](#authenticaterequest) | [AuthenticateResponse](#authenticateresponse) |  |
| OpenChannel | [OpenChannelRequest](#openchannelrequest) | [OpenChannelResponse](#openchannelresponse) |  |
| GetChannelJwt | [GetJwtRequest](#getjwtrequest) | [GetJwtResponse](#getjwtresponse) |  |
| GetPositions | [GetPositionsRequest](#getpositionsrequest) | [GetPositionsResponse](#getpositionsresponse) |  |
| RecordTrade | [TradeRequest](#traderequest) | [TradeResponse](#traderesponse) |  |
| RecordTrades | [TradesRequest](#tradesrequest) | [TradesResponse](#tradesresponse) |  |
| RequestSettlement | [SettlementRequest](#settlementrequest) | [SettlementResponse](#settlementresponse) |  |
| CloseChannel | [CloseChannelRequest](#closechannelrequest) | [CloseChannelResponse](#closechannelresponse) |  |
| SubscribeChannelsEvents | [SubscribeRequest](#subscriberequest) | [Notification](#notification) stream |  |

 <!-- end services -->



### AuthenticateRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| challenge | [string](#string) |  |
| signature | [string](#string) |  |







### AuthenticateResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| jwt | [string](#string) |  |







### CloseChannelRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### CloseChannelResponse








### ErrorMetadata



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| tx_id | [string](#string) |  |







### ErrorNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| code | [string](#string) |  |
| msg | [string](#string) |  |
| err | [string](#string) |  |
| metadata | [ErrorMetadata](#errormetadata) |  |







### GetChallengeRequest








### GetChallengeResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| challenge | [string](#string) |  |







### GetIdentityAddressRequest








### GetIdentityAddressResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| address | [core.Address](core.proto.md#address) |  |







### GetJwtRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### GetJwtResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| jwt | [string](#string) |  |







### GetPositionsRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### GetPositionsResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| positions | [core.Position](core.proto.md#position) |  |







### Notification



| Field | Type | Description |
| ----- | ---- | ----------- |
| stream_notification | [StreamNotification](#streamnotification) |  |
| settlement_notification | [SettlementNotification](#settlementnotification) |  |
| position_notification | [PositionNotification](#positionnotification) |  |
| error_notification | [ErrorNotification](#errornotification) |  |







### OpenChannelRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| peer | [auth.Peer](auth.proto.md#peer) |  |
| margin_deposit | [string](#string) |  |







### OpenChannelResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### PositionNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| position | [core.Position](core.proto.md#position) |  |







### SettlementNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| settlement_state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |







### SettlementRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| payment_method | [settlement.PaymentMethod](settlement.proto.md#paymentmethod) |  |
| chain_id | [uint64](#uint64) |  |
| markets | [core.Market](core.proto.md#market) |  |







### SettlementResponse








### StreamNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| notification_type | [NotificationType](#notificationtype) |  |
| channel_status | [ChannelStatus](#channelstatus) |  |
| my_role | [core.ProtocolIndex](core.proto.md#protocolindex) |  |
| peer | [auth.Peer](auth.proto.md#peer) |  |
| margin_limit_type | [core.MarginLimitType](core.proto.md#marginlimittype) |  |
| margin_deposit | [string](#string) |  |
| initiator_margin_balance | [string](#string) | margin updates will be reflected here |
| follower_margin_balance | [string](#string) | margin updates will be reflected here |
| turn_num | [uint64](#uint64) |  |







### SubscribeRequest








### TradeRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| trade | [core.Trade](core.proto.md#trade) |  |







### TradeResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| trade | [core.Trade](core.proto.md#trade) |  |







### TradesRequest
Maximum number of trades per request is 6765


| Field | Type | Description |
| ----- | ---- | ----------- |
| trades | [core.Trade](core.proto.md#trade) |  |







### TradesResponse








### VersionRequest








### VersionResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| version | [string](#string) |  |





 <!-- end messages -->



### ChannelStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANNEL_STATUS_UNSPECIFIED | 0 |  |
| CHANNEL_STATUS_OPENING | 1 |  |
| CHANNEL_STATUS_OPEN | 2 |  |
| CHANNEL_STATUS_CHALLENGING | 3 |  |
| CHANNEL_STATUS_CLOSED | 4 |  |




### NotificationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NOTIFICATION_TYPE_UNSPECIFIED | 0 |  |
| NOTIFICATION_TYPE_CHANNEL_OPENING | 1 |  |
| NOTIFICATION_TYPE_CHANNEL_OPENED | 2 |  |
| NOTIFICATION_TYPE_CHANNEL_CLOSED | 3 |  |
| NOTIFICATION_TYPE_CHALLENGE_STARTED | 4 |  |
| NOTIFICATION_TYPE_CHALLENGE_CLEARED | 5 |  |
| NOTIFICATION_TYPE_CHALLENGE_FINISHED | 6 |  |
| NOTIFICATION_TYPE_MARGIN_UPDATED | 7 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
