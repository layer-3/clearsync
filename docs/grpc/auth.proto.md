# gRPC Protocol Documentation

This file describes the content of auth.proto.


### Auth



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetChallenge | [auth.ChallengeRequest](auth.proto.md#challengerequest) | [auth.ChallengeResponse](auth.proto.md#challengeresponse) |  |
| Authenticate | [auth.SignedChallenge](auth.proto.md#signedchallenge) | [auth.AuthToken](auth.proto.md#authtoken) |  |
| TokenRefresh | [auth.AuthToken](auth.proto.md#authtoken) | [auth.AuthToken](auth.proto.md#authtoken) |  |
| Signout | [auth.AuthToken](auth.proto.md#authtoken) | [auth.SignoutResponse](auth.proto.md#signoutresponse) |  |

 <!-- end services -->



### AuthToken
AuthToken can be attached in headers


| Field | Type | Description |
| ----- | ---- | ----------- |
| token | [string](#string) |  |







### ChallengeRequest



| Field | Type | Description |
| ----- | ---- | ----------- |
| client | [auth.Peer](auth.proto.md#peer) |  |
| operator_address | [string](#string) |  |







### ChallengeResponse
Server signs the client's challenge


| Field | Type | Description |
| ----- | ---- | ----------- |
| server | [auth.Peer](auth.proto.md#peer) |  |
| server_challenge | [string](#string) |  |







### Peer



| Field | Type | Description |
| ----- | ---- | ----------- |
| url | [string](#string) |  |
| name | [string](#string) |  |
| participant_address | [string](#string) |  |
| operator_address | [string](#string) |  |







### SignedChallenge
This should be signed with the client's private key


| Field | Type | Description |
| ----- | ---- | ----------- |
| client | [auth.Peer](auth.proto.md#peer) |  |
| signed_server_challenge | [bytes](#bytes) |  |
| operator_address | [string](#string) |  |







### SignoutResponse






 <!-- end messages -->

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
