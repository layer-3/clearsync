# gRPC Protocol Documentation

This file describes the content of exit.proto.

 <!-- end services -->



### Allocation



| Field | Type | Description |
| ----- | ---- | ----------- |
| destination | [bytes](#bytes) | bytes32 in solidity |
| amount | [string](#string) | big.Int cast to string |
| allocation_type | [outcome.AllocationType](outcome.proto.md#allocationtype) |  |
| metadata | [bytes](#bytes) |  |







### AssetMetadata



| Field | Type | Description |
| ----- | ---- | ----------- |
| asset_type | [outcome.AssetType](outcome.proto.md#assettype) |  |
| metadata | [bytes](#bytes) |  |







### Exit



| Field | Type | Description |
| ----- | ---- | ----------- |
| single_asset_exits | [outcome.SingleAssetExit](outcome.proto.md#singleassetexit) |  |







### SingleAssetExit



| Field | Type | Description |
| ----- | ---- | ----------- |
| asset | [core.Address](core.proto.md#address) | Either the zero address (implying the native token) or the address of an ERC20 contract |
| asset_metadata | [outcome.AssetMetadata](outcome.proto.md#assetmetadata) |  |
| allocations | [outcome.Allocation](outcome.proto.md#allocation) |  |





 <!-- end messages -->



### AllocationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ALLOCATION_TYPE_UNSPECIFIED | 0 |  |
| ALLOCATION_TYPE_WITHDRAW_HELPER | 1 |  |
| ALLOCATION_TYPE_GUARANTEE | 2 |  |




### AssetType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASSET_TYPE_UNSPECIFIED | 0 |  |
| ASSET_TYPE_ERC721 | 1 |  |
| ASSET_TYPE_ERC1155 | 2 |  |
| ASSET_TYPE_QUALIFIED | 3 |  |


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
