# gRPC Protocol Documentation

This file describes the content of core.proto.

 <!-- end services -->



### Address
A 42-character hexadecimal address
derived from the last 20 bytes of the public key


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [string](#string) |  |







### BigInt



| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [string](#string) |  |







### Decimal
Represent Decimal as a string
Can be changed to 2 numbers
Due to compatibility issues we may need
To create a new Decimal type compatible with ExitFormat


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [string](#string) |  |







### Market



| Field | Type | Description |
| ----- | ---- | ----------- |
| base | [string](#string) |  |
| quote | [string](#string) |  |







### Position



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [string](#string) |  |
| channel_id | [string](#string) |  |
| market | [core.Market](core.proto.md#market) |  |
| direction | [core.Direction](core.proto.md#direction) |  |
| amount | [string](#string) |  |
| cost | [string](#string) |  |
| market_value | [string](#string) |  |
| pnl | [string](#string) |  |
| status | [core.PositionStatus](core.proto.md#positionstatus) |  |







### Signature
A 132-character hexadecimal string


| Field | Type | Description |
| ----- | ---- | ----------- |
| v | [uint32](#uint32) |  |
| r | [bytes](#bytes) | 32 bytes |
| s | [bytes](#bytes) | 32 bytes |







### Trade



| Field | Type | Description |
| ----- | ---- | ----------- |
| channel_id | [string](#string) |  |
| external_id | [string](#string) |  |
| market | [core.Market](core.proto.md#market) |  |
| direction | [core.Direction](core.proto.md#direction) |  |
| amount | [core.Decimal](core.proto.md#decimal) |  |
| price | [core.Decimal](core.proto.md#decimal) |  |
| executed_at | [int64](#int64) |  |





 <!-- end messages -->



### Direction


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIRECTION_UNSPECIFIED | 0 |  |
| DIRECTION_BUY | 1 |  |
| DIRECTION_SELL | 2 |  |




### MarginLimitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| MARGIN_LIMIT_TYPE_UNSPECIFIED | 0 |  |
| MARGIN_LIMIT_TYPE_NONE | 1 |  |
| MARGIN_LIMIT_TYPE_SOFT | 2 |  |
| MARGIN_LIMIT_TYPE_HARD | 3 |  |




### PositionStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| POSITION_STATUS_UNSPECIFIED | 0 |  |
| POSITION_STATUS_OPEN | 1 |  |
| POSITION_STATUS_IN_SETTLEMENT | 2 |  |
| POSITION_STATUS_SETTLED | 3 |  |
| POSITION_STATUS_CLOSED | 4 |  |




### ProtocolIndex


| Name | Number | Description |
| ---- | ------ | ----------- |
| PROTOCOL_INDEX_UNSPECIFIED | 0 |  |
| PROTOCOL_INDEX_INITIATOR | 1 |  |
| PROTOCOL_INDEX_RESPONDER | 2 |  |


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
