# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [emoine_r/v1/schema.proto](#emoine_r_v1_schema-proto)
    - [Comment](#emoine_r-v1-Comment)
    - [Meeting](#emoine_r-v1-Meeting)
    - [Reaction](#emoine_r-v1-Reaction)
    - [Token](#emoine_r-v1-Token)
  
- [emoine_r/v1/admin_api.proto](#emoine_r_v1_admin_api-proto)
    - [CreateMeetingRequest](#emoine_r-v1-CreateMeetingRequest)
    - [CreateMeetingResponse](#emoine_r-v1-CreateMeetingResponse)
    - [CreateTokenRequest](#emoine_r-v1-CreateTokenRequest)
    - [CreateTokenResponse](#emoine_r-v1-CreateTokenResponse)
    - [DeleteMeetingRequest](#emoine_r-v1-DeleteMeetingRequest)
    - [GetMeetingTokensRequest](#emoine_r-v1-GetMeetingTokensRequest)
    - [GetMeetingTokensResponse](#emoine_r-v1-GetMeetingTokensResponse)
    - [GetTokenRequest](#emoine_r-v1-GetTokenRequest)
    - [GetTokenResponse](#emoine_r-v1-GetTokenResponse)
    - [GetTokensResponse](#emoine_r-v1-GetTokensResponse)
    - [UpdateMeetingRequest](#emoine_r-v1-UpdateMeetingRequest)
    - [UpdateTokenRequest](#emoine_r-v1-UpdateTokenRequest)
  
    - [AdminAPIService](#emoine_r-v1-AdminAPIService)
  
- [emoine_r/v1/general_api.proto](#emoine_r_v1_general_api-proto)
    - [ConnectToMeetingStreamRequest](#emoine_r-v1-ConnectToMeetingStreamRequest)
    - [ConnectToMeetingStreamResponse](#emoine_r-v1-ConnectToMeetingStreamResponse)
    - [GetMeetingCommentsRequest](#emoine_r-v1-GetMeetingCommentsRequest)
    - [GetMeetingCommentsResponse](#emoine_r-v1-GetMeetingCommentsResponse)
    - [GetMeetingReactionsRequest](#emoine_r-v1-GetMeetingReactionsRequest)
    - [GetMeetingReactionsResponse](#emoine_r-v1-GetMeetingReactionsResponse)
    - [GetMeetingRequest](#emoine_r-v1-GetMeetingRequest)
    - [GetMeetingResponse](#emoine_r-v1-GetMeetingResponse)
    - [GetMeetingsRequest](#emoine_r-v1-GetMeetingsRequest)
    - [GetMeetingsResponse](#emoine_r-v1-GetMeetingsResponse)
    - [OAuth2AuthorizeRequest](#emoine_r-v1-OAuth2AuthorizeRequest)
    - [OAuth2CallbackRequest](#emoine_r-v1-OAuth2CallbackRequest)
    - [SendCommentRequest](#emoine_r-v1-SendCommentRequest)
    - [SendCommentResponse](#emoine_r-v1-SendCommentResponse)
    - [SendReactionRequest](#emoine_r-v1-SendReactionRequest)
    - [SendReactionResponse](#emoine_r-v1-SendReactionResponse)
  
    - [GeneralAPIService](#emoine_r-v1-GeneralAPIService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="emoine_r_v1_schema-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## emoine_r/v1/schema.proto



<a name="emoine_r-v1-Comment"></a>

### Comment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| meeting_id | [string](#string) |  |  |
| username | [string](#string) |  |  |
| text | [string](#string) |  |  |
| is_anonymous | [bool](#bool) |  |  |
| color | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="emoine_r-v1-Meeting"></a>

### Meeting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| video_id | [string](#string) |  |  |
| thumbnail | [string](#string) |  |  |
| description | [string](#string) |  |  |
| started_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| ended_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="emoine_r-v1-Reaction"></a>

### Reaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| meeting_id | [string](#string) |  |  |
| username | [string](#string) |  |  |
| stamp_id | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="emoine_r-v1-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| username | [string](#string) |  |  |
| meeting_id | [string](#string) |  |  |
| creator_id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| expire_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="emoine_r_v1_admin_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## emoine_r/v1/admin_api.proto



<a name="emoine_r-v1-CreateMeetingRequest"></a>

### CreateMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| video_id | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="emoine_r-v1-CreateMeetingResponse"></a>

### CreateMeetingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting | [Meeting](#emoine_r-v1-Meeting) |  |  |






<a name="emoine_r-v1-CreateTokenRequest"></a>

### CreateTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="emoine_r-v1-CreateTokenResponse"></a>

### CreateTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#emoine_r-v1-Token) |  |  |






<a name="emoine_r-v1-DeleteMeetingRequest"></a>

### DeleteMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetMeetingTokensRequest"></a>

### GetMeetingTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetMeetingTokensResponse"></a>

### GetMeetingTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [Token](#emoine_r-v1-Token) | repeated |  |






<a name="emoine_r-v1-GetTokenRequest"></a>

### GetTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetTokenResponse"></a>

### GetTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#emoine_r-v1-Token) |  |  |






<a name="emoine_r-v1-GetTokensResponse"></a>

### GetTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [Token](#emoine_r-v1-Token) | repeated |  |






<a name="emoine_r-v1-UpdateMeetingRequest"></a>

### UpdateMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| video_id | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="emoine_r-v1-UpdateTokenRequest"></a>

### UpdateTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| meeting_id | [string](#string) |  |  |
| description | [string](#string) |  |  |





 

 

 


<a name="emoine_r-v1-AdminAPIService"></a>

### AdminAPIService
管理者権限が必要なAPIを提供します

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateMeeting | [CreateMeetingRequest](#emoine_r-v1-CreateMeetingRequest) | [CreateMeetingResponse](#emoine_r-v1-CreateMeetingResponse) | 集会を作成します |
| UpdateMeeting | [UpdateMeetingRequest](#emoine_r-v1-UpdateMeetingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 集会情報を更新します |
| DeleteMeeting | [DeleteMeetingRequest](#emoine_r-v1-DeleteMeetingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 集会を削除します |
| GetMeetingTokens | [GetMeetingTokensRequest](#emoine_r-v1-GetMeetingTokensRequest) | [GetMeetingTokensResponse](#emoine_r-v1-GetMeetingTokensResponse) | 該当する集会のトークン一覧を取得します |
| GetTokens | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetTokensResponse](#emoine_r-v1-GetTokensResponse) | トークン一覧を取得します |
| GetToken | [GetTokenRequest](#emoine_r-v1-GetTokenRequest) | [GetTokenResponse](#emoine_r-v1-GetTokenResponse) | 該当するトークンを取得します |
| CreateToken | [CreateTokenRequest](#emoine_r-v1-CreateTokenRequest) | [CreateTokenResponse](#emoine_r-v1-CreateTokenResponse) | トークンを作成します |
| UpdateToken | [UpdateTokenRequest](#emoine_r-v1-UpdateTokenRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | トークン情報を更新します |

 



<a name="emoine_r_v1_general_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## emoine_r/v1/general_api.proto



<a name="emoine_r-v1-ConnectToMeetingStreamRequest"></a>

### ConnectToMeetingStreamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-ConnectToMeetingStreamResponse"></a>

### ConnectToMeetingStreamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting | [Meeting](#emoine_r-v1-Meeting) |  |  |
| comment | [Comment](#emoine_r-v1-Comment) |  |  |
| reaction | [Reaction](#emoine_r-v1-Reaction) |  |  |






<a name="emoine_r-v1-GetMeetingCommentsRequest"></a>

### GetMeetingCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetMeetingCommentsResponse"></a>

### GetMeetingCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| comments | [Comment](#emoine_r-v1-Comment) | repeated |  |






<a name="emoine_r-v1-GetMeetingReactionsRequest"></a>

### GetMeetingReactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetMeetingReactionsResponse"></a>

### GetMeetingReactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reactions | [Reaction](#emoine_r-v1-Reaction) | repeated |  |






<a name="emoine_r-v1-GetMeetingRequest"></a>

### GetMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="emoine_r-v1-GetMeetingResponse"></a>

### GetMeetingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting | [Meeting](#emoine_r-v1-Meeting) |  |  |






<a name="emoine_r-v1-GetMeetingsRequest"></a>

### GetMeetingsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |






<a name="emoine_r-v1-GetMeetingsResponse"></a>

### GetMeetingsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| meetings | [Meeting](#emoine_r-v1-Meeting) | repeated |  |






<a name="emoine_r-v1-OAuth2AuthorizeRequest"></a>

### OAuth2AuthorizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [string](#string) |  |  |






<a name="emoine_r-v1-OAuth2CallbackRequest"></a>

### OAuth2CallbackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| state | [string](#string) |  |  |






<a name="emoine_r-v1-SendCommentRequest"></a>

### SendCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| text | [string](#string) |  |  |
| is_anonymous | [bool](#bool) |  |  |
| color | [string](#string) |  |  |






<a name="emoine_r-v1-SendCommentResponse"></a>

### SendCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| comment | [Comment](#emoine_r-v1-Comment) |  |  |






<a name="emoine_r-v1-SendReactionRequest"></a>

### SendReactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| stamp_id | [string](#string) |  |  |






<a name="emoine_r-v1-SendReactionResponse"></a>

### SendReactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reaction | [Reaction](#emoine_r-v1-Reaction) |  |  |





 

 

 


<a name="emoine_r-v1-GeneralAPIService"></a>

### GeneralAPIService
権限が必要ないAPIを提供します

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetMeetings | [GetMeetingsRequest](#emoine_r-v1-GetMeetingsRequest) | [GetMeetingsResponse](#emoine_r-v1-GetMeetingsResponse) | 集会一覧を取得します |
| GetMeeting | [GetMeetingRequest](#emoine_r-v1-GetMeetingRequest) | [GetMeetingResponse](#emoine_r-v1-GetMeetingResponse) | 該当する集会を取得します |
| GetMeetingComments | [GetMeetingCommentsRequest](#emoine_r-v1-GetMeetingCommentsRequest) | [GetMeetingCommentsResponse](#emoine_r-v1-GetMeetingCommentsResponse) | 該当する集会のコメント一覧を取得します |
| GetMeetingReactions | [GetMeetingReactionsRequest](#emoine_r-v1-GetMeetingReactionsRequest) | [GetMeetingReactionsResponse](#emoine_r-v1-GetMeetingReactionsResponse) | 該当する集会のリアクション一覧を取得します |
| ConnectToMeetingStream | [ConnectToMeetingStreamRequest](#emoine_r-v1-ConnectToMeetingStreamRequest) | [ConnectToMeetingStreamResponse](#emoine_r-v1-ConnectToMeetingStreamResponse) stream | 集会のストリームに接続します |
| SendComment | [SendCommentRequest](#emoine_r-v1-SendCommentRequest) | [SendCommentResponse](#emoine_r-v1-SendCommentResponse) | (コメントは集会のストリームに反映されます) |
| SendReaction | [SendReactionRequest](#emoine_r-v1-SendReactionRequest) | [SendReactionResponse](#emoine_r-v1-SendReactionResponse) | (リアクションは集会のストリームに反映されます) |
| OAuth2Authorize | [OAuth2AuthorizeRequest](#emoine_r-v1-OAuth2AuthorizeRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | OAuth2による認可を行います |
| OAuth2Callback | [OAuth2CallbackRequest](#emoine_r-v1-OAuth2CallbackRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | OAuth2のコールバックを受け取ります |

 



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

