# gRPC Protocol Documentation

This file describes the content of operator.proto.


### Operator

protolint:disable MAX_LINE_LENGTH

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVersion | [VersionRequest](#versionrequest) | [VersionResponse](#versionresponse) |  |
| GetChallenge | [GetChallengeRequest](#getchallengerequest) | [GetChallengeResponse](#getchallengeresponse) |  |
| Authenticate | [AuthenticateRequest](#authenticaterequest) | [AuthenticateResponse](#authenticateresponse) |  |
| OpenChannel | [OpenChannelRequest](#openchannelrequest) | [OpenChannelResponse](#openchannelresponse) |  |
| GetChannelJwt | [GetJwtRequest](#getjwtrequest) | [GetJwtResponse](#getjwtresponse) |  |
| GetPositions | [GetPositionsRequest](#getpositionsrequest) | [GetPositionsResponse](#getpositionsresponse) |  |
| GetActiveClearStreams | [GetActiveClearStreamsRequest](#getactiveclearstreamsrequest) | [GetActiveClearStreamsResponse](#getactiveclearstreamsresponse) |  |
| RecordTrade | [TradeRequest](#traderequest) | [TradeResponse](#traderesponse) |  |
| RecordTrades | [TradesRequest](#tradesrequest) | [TradesResponse](#tradesresponse) |  |
| RequestSettlement | [SettlementRequest](#settlementrequest) | [SettlementResponse](#settlementresponse) |  |
| CloseChannel | [CloseChannelRequest](#closechannelrequest) | [CloseChannelResponse](#closechannelresponse) |  |
| SubscribeChannelsEvents | [SubscribeRequest](#subscriberequest) | [Notification](#notification) stream |  |

 <!-- end services -->



### AuthenticateRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| signature | [string](#string) |  |
| address | [core.Address](core.proto.md#address) |  |







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
| state | [string](#string) |  |
| event | [string](#string) |  |







### ErrorNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| msg | [string](#string) |  |
| metadata | [ErrorMetadata](#errormetadata) |  |
| action | [string](#string) |  |







### GetActiveClearStreamsRequest








### GetActiveClearStreamsResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| clear_streams | [clear_stream.ClearStream](clear_stream.proto.md#clearstream) |  |







### GetChallengeRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| address | [core.Address](core.proto.md#address) |  |
| name | [string](#string) |  |







### GetChallengeResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| challenge | [string](#string) |  |







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
| transaction_notification | [TransactionNotification](#transactionnotification) |  |
| session_key_transaction_notification | [SessionKeyTransactionNotification](#sessionkeytransactionnotification) |  |
| record_trade_notification | [RecordTradeNotification](#recordtradenotification) |  |







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







### RecordTradeNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| trade | [core.Trade](core.proto.md#trade) |  |
| revert_reason | [string](#string) |  |







### SessionKeyTransactionNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| incomplete_userop | [IncompleteUserOp](#incompleteuserop) |  |
| enable_sig_offset | [uint64](#uint64) |  |
| digest_hash | [string](#string) |  |
| permissions | [SessionKeyPermission](#sessionkeypermission) |  |







### SettlementNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| settlement_state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |
| markets | [core.Market](core.proto.md#market) |  |
| ledger | [settlement.SettlementLedger](settlement.proto.md#settlementledger) |  |







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
| event | [ClearingEvent](#clearingevent) |  |
| state | [ClearingState](#clearingstate) |  |
| my_role | [core.ProtocolIndex](core.proto.md#protocolindex) |  |
| peer | [auth.Peer](auth.proto.md#peer) |  |
| margin_limit_type | [core.MarginLimitType](core.proto.md#marginlimittype) |  |
| margin_deposit | [string](#string) |  |
| initiator_margin_balance | [string](#string) | margin updates will be reflected here |
| follower_margin_balance | [string](#string) | margin updates will be reflected here |
| turn_num | [uint64](#uint64) |  |
| clearing_sm_state | [state_machine.ClearingSMState](state_machine.proto.md#clearingsmstate) |  |







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








### TransactionNotification



| Field | Type | Description |
| ----- | ---- | ----------- |
| to | [string](#string) |  |
| data | [string](#string) |  |







### VersionRequest








### VersionResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| version | [string](#string) |  |





 <!-- end messages -->



### ClearingEvent


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLEARING_EVENT_UNSPECIFIED | 0 |  |
| CLEARING_EVENT_INSTANTIATED | 1 |  |
| CLEARING_EVENT_ACCEPTED | 2 |  |
| CLEARING_EVENT_FAILED | 3 |  |
| CLEARING_EVENT_INITIATOR_FUNDED | 4 |  |
| CLEARING_EVENT_RESPONDER_FUNDED | 5 |  |
| CLEARING_EVENT_POSTFUND_PROPOSED | 6 |  |
| CLEARING_EVENT_AGREED_ON_POSTFUND | 7 |  |
| CLEARING_EVENT_MOVE_TO_OPERATIONAL | 8 |  |
| CLEARING_EVENT_PROCESS_MARGIN_CALL | 9 |  |
| CLEARING_EVENT_STARTED_SETTLEMENT | 10 |  |
| CLEARING_EVENT_PROCESS_POST_SETTLEMENT_MARGIN | 11 |  |
| CLEARING_EVENT_FINALIZE_SETTLEMENT | 12 |  |
| CLEARING_EVENT_FAILED_SETTLEMENT | 13 |  |
| CLEARING_EVENT_CHALLENGE | 14 |  |
| CLEARING_EVENT_FINALIZE | 15 |  |
| CLEARING_EVENT_WITHDRAW | 16 |  |
| CLEARING_EVENT_CLEAR_CHALLENGE | 17 |  |
| CLEARING_EVENT_TIMEOUT_CHALLENGE | 18 |  |




### ClearingState


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLEARING_STATE_UNSPECIFIED | 0 |  |
| CLEARING_STATE_INSTANTIATED | 1 |  |
| CLEARING_STATE_ACCEPTED | 2 |  |
| CLEARING_STATE_FAILED | 3 |  |
| CLEARING_STATE_INITIATOR_FUNDED | 4 |  |
| CLEARING_STATE_FUNDED | 5 |  |
| CLEARING_STATE_OPERATIONAL | 6 |  |
| CLEARING_STATE_ISSUING_MARGIN_CALL | 7 |  |
| CLEARING_STATE_PROCESSING_MARGIN_CALL | 8 |  |
| CLEARING_STATE_ACTIVE_SETTLEMENT | 9 |  |
| CLEARING_STATE_ISSUING_POST_SETTLEMENT_MARGIN | 10 |  |
| CLEARING_STATE_PROCESSING_POST_SETTLEMENT_MARGIN | 11 |  |
| CLEARING_STATE_CHALLENGING | 12 |  |
| CLEARING_STATE_FINALIZING | 13 |  |
| CLEARING_STATE_WITHDRAWING | 14 |  |
| CLEARING_STATE_CONCLUDING | 15 |  |
| CLEARING_STATE_DEFAULT | 16 |  |


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
