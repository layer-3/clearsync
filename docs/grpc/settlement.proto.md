# gRPC Protocol Documentation

This file describes the content of settlement.proto.


### SettlementService



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Prepare | [settlement.PrepareRequest](settlement.proto.md#preparerequest) | [settlement.PrepareResponse](settlement.proto.md#prepareresponse) |  |
| Update | [settlement.SettlementStateUpdate](settlement.proto.md#settlementstateupdate) | [settlement.StateConfirmation](settlement.proto.md#stateconfirmation) |  |
| Finalize | [settlement.FinalizeRequest](settlement.proto.md#finalizerequest) | [settlement.FinalizeResponse](settlement.proto.md#finalizeresponse) |  |

 <!-- end services -->



### Asset



| Field | Type | Description |
| ----- | ---- | ----------- |
| symbol | [string](#string) |  |
| chain_id | [uint64](#uint64) |  |
| address | [core.Address](core.proto.md#address) |  |
| decimals | [uint32](#uint32) |  |







### FinalizeRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| cid | [string](#string) |  |







### FinalizeResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |







### Liability



| Field | Type | Description |
| ----- | ---- | ----------- |
| asset | [settlement.Asset](settlement.proto.md#asset) |  |
| amount | [string](#string) |  |







### PrepareRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| settlement | [settlement.Settlement](settlement.proto.md#settlement) |  |







### PrepareResponse



| Field | Type | Description |
| ----- | ---- | ----------- |
| state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |
| markets | [string](#string) |  |







### Settlement



| Field | Type | Description |
| ----- | ---- | ----------- |
| cid | [string](#string) |  |
| type | [settlement.SettlementType](settlement.proto.md#settlementtype) |  |
| state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |
| ledger | [settlement.SettlementLedger](settlement.proto.md#settlementledger) |  |
| psm_turn_num | [uint64](#uint64) |  |
| payment_method | [settlement.PaymentMethod](settlement.proto.md#paymentmethod) |  |
| markets | [string](#string) |  |
| chain_id | [uint64](#uint64) |  |







### SettlementLedger



| Field | Type | Description |
| ----- | ---- | ----------- |
| initiator_entries | [settlement.Liability](settlement.proto.md#liability) |  |
| responder_entries | [settlement.Liability](settlement.proto.md#liability) |  |
| next_margin | [core.MarginCall](core.proto.md#margincall) |  |







### SettlementStateUpdate



| Field | Type | Description |
| ----- | ---- | ----------- |
| cid | [string](#string) |  |
| to_state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |







### StateConfirmation



| Field | Type | Description |
| ----- | ---- | ----------- |
| state | [settlement.SettlementState](settlement.proto.md#settlementstate) |  |





 <!-- end messages -->



### PaymentMethod


| Name | Number | Description |
| ---- | ------ | ----------- |
| PAYMENT_METHOD_UNSPECIFIED | 0 |  |
| PAYMENT_METHOD_ESCROW | 1 |  |
| PAYMENT_METHOD_MOCK | 2 |  |




### SettlementState


| Name | Number | Description |
| ---- | ------ | ----------- |
| SETTLEMENT_STATE_UNSPECIFIED | 0 |  |
| SETTLEMENT_STATE_PROPOSED | 1 |  |
| SETTLEMENT_STATE_ACCEPTED | 2 |  |
| SETTLEMENT_STATE_INITIATED | 3 |  |
| SETTLEMENT_STATE_PREPARED | 4 |  |
| SETTLEMENT_STATE_EXECUTED | 5 |  |
| SETTLEMENT_STATE_COMPLETED | 6 |  |
| SETTLEMENT_STATE_WITHDRAWN | 7 |  |
| SETTLEMENT_STATE_FAILED | 8 |  |
| SETTLEMENT_STATE_REJECTED | 9 |  |




### SettlementType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SETTLEMENT_TYPE_UNSPECIFIED | 0 |  |
| SETTLEMENT_TYPE_AVAILABLE | 1 |  |
| SETTLEMENT_TYPE_FORCE | 2 |  |


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
