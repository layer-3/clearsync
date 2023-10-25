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
| RecordTrade | [TradeRequest](#traderequest) | [TradeResponse](#traderesponse) |  |
| RequestSettlement | [SettlementRequest](#settlementrequest) | [SettlementResponse](#settlementresponse) |  |
| CloseChannel | [CloseChannelRequest](#closechannelrequest) | [CloseChannelResponse](#closechannelresponse) | rpc GetPosition() returns (); rpc SettleChannel() returns (); |
| SubscribeChannelsEvents | [SubscribeRequest](#subscriberequest) | [ChannelNotification](#channelnotification) stream |  |

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







### ChannelNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| notification_type | [string](#string) | "channel_opened", "channel_closed", "channel_settled", "margin_updated" |
| channel_status | [string](#string) | "opening", "open", "settling", "closing", "closed" |
| my_role | [string](#string) | "initiator" or "follower" |
| peer | [auth.Peer](auth.proto.md#peer) |  |
| margin_deposit | [string](#string) |  |
| initiator_margin_balance | [string](#string) | margin updates will be reflected here |
| follower_margin_balance | [string](#string) | margin updates will be reflected here |







### CloseChannelRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### CloseChannelResponse








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







### Market



| Field | Type | Description |
| ----- | ---- | ----------- |
| base | [string](#string) |  |
| quote | [string](#string) |  |







### OpenChannelRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| peer | [auth.Peer](auth.proto.md#peer) |  |
| margin_deposit | [string](#string) |  |







### OpenChannelResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |







### SettlementRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| payment_method | [PaymentMethod](#paymentmethod) |  |
| chain_id | [uint64](#uint64) |  |
| markets | [Market](#market) |  |







### SettlementResponse








### SubscribeRequest








### Trade



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| external_id | [string](#string) |  |
| market | [Market](#market) |  |
| direction | [Direction](#direction) |  |
| amount | [core.Decimal](core.proto.md#decimal) |  |
| price | [core.Decimal](core.proto.md#decimal) |  |
| executed_at | [int64](#int64) |  |







### TradeRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| trade | [Trade](#trade) |  |







### TradeResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| trade | [Trade](#trade) |  |







### VersionRequest








### VersionResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| version | [string](#string) |  |





 <!-- end messages -->



### Direction


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIRECTION_UNSPECIFIED | 0 |  |
| DIRECTION_BUY | 1 |  |
| DIRECTION_SELL | 2 |  |




### PaymentMethod


| Name | Number | Description |
| ---- | ------ | ----------- |
| PAYMENT_METHOD_UNSPECIFIED | 0 |  |
| PAYMENT_METHOD_ESCROW | 1 |  |


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
