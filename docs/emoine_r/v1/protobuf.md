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
    - [DeleteMeetingRequest](#emoine_r-v1-DeleteMeetingRequest)
    - [GenerateTokenRequest](#emoine_r-v1-GenerateTokenRequest)
    - [GenerateTokenResponse](#emoine_r-v1-GenerateTokenResponse)
    - [GetTokensRequest](#emoine_r-v1-GetTokensRequest)
    - [GetTokensResponse](#emoine_r-v1-GetTokensResponse)
    - [RevokeTokenRequest](#emoine_r-v1-RevokeTokenRequest)
    - [UpdateMeetingRequest](#emoine_r-v1-UpdateMeetingRequest)
  
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
| id | [string](#string) |  | コメントのID (UUID) |
| meeting_id | [string](#string) |  | コメントが送信された集会のID |
| username | [string](#string) |  | コメントの送信者名 |
| text | [string](#string) |  | コメントの本文 |
| is_anonymous | [bool](#bool) |  | コメントが匿名かどうか |
| color | [string](#string) |  | コメントの色 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | コメントの作成日時 |






<a name="emoine_r-v1-Meeting"></a>

### Meeting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 集会ID (UUID) |
| video_id | [string](#string) |  | YouTubeの動画ID |
| description | [string](#string) |  | 集会の説明 |
| title | [string](#string) |  | 動画タイトル (YouTubeから取得) |
| thumbnail | [string](#string) |  | 動画サムネイル (YouTubeから取得) |
| started_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 集会の開始日時 (YouTubeから取得) |
| ended_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 集会の終了日時 (YouTubeから取得) |






<a name="emoine_r-v1-Reaction"></a>

### Reaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | リアクションのID (UUID) |
| meeting_id | [string](#string) |  | リアクションが送信された集会のID |
| username | [string](#string) |  | リアクションの送信者名 |
| stamp_id | [string](#string) |  | 送信されたリアクションのスタンプID (UUID, traQと同期) |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | リアクションの作成日時 |






<a name="emoine_r-v1-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | トークン文字列 |
| username | [string](#string) |  | トークンの所有者名 |
| meeting_id | [string](#string) |  | トークンが有効な集会のID |
| creator_id | [string](#string) |  | トークン発行者のtraQID (traQと同期) |
| description | [string](#string) |  | トークンの説明 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | トークンの作成日時 |
| expire_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | トークンの有効期限 |





 

 

 

 



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






<a name="emoine_r-v1-DeleteMeetingRequest"></a>

### DeleteMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GenerateTokenRequest"></a>

### GenerateTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| username | [string](#string) |  |  |
| description | [string](#string) |  |  |
| expire_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="emoine_r-v1-GenerateTokenResponse"></a>

### GenerateTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#emoine_r-v1-Token) |  |  |






<a name="emoine_r-v1-GetTokensRequest"></a>

### GetTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetTokensResponse"></a>

### GetTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [Token](#emoine_r-v1-Token) | repeated |  |






<a name="emoine_r-v1-RevokeTokenRequest"></a>

### RevokeTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| meeting_id | [string](#string) |  |  |






<a name="emoine_r-v1-UpdateMeetingRequest"></a>

### UpdateMeetingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meeting_id | [string](#string) |  |  |
| video_id | [string](#string) | optional |  |
| description | [string](#string) | optional |  |





 

 

 


<a name="emoine_r-v1-AdminAPIService"></a>

### AdminAPIService
管理者権限が必要なAPIを提供します

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateMeeting | [CreateMeetingRequest](#emoine_r-v1-CreateMeetingRequest) | [CreateMeetingResponse](#emoine_r-v1-CreateMeetingResponse) | 集会を作成します |
| UpdateMeeting | [UpdateMeetingRequest](#emoine_r-v1-UpdateMeetingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 集会情報を更新します |
| DeleteMeeting | [DeleteMeetingRequest](#emoine_r-v1-DeleteMeetingRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 集会を削除します |
| GetTokens | [GetTokensRequest](#emoine_r-v1-GetTokensRequest) | [GetTokensResponse](#emoine_r-v1-GetTokensResponse) | 該当する集会のトークン一覧を取得します |
| GenerateToken | [GenerateTokenRequest](#emoine_r-v1-GenerateTokenRequest) | [GenerateTokenResponse](#emoine_r-v1-GenerateTokenResponse) | 集会用のトークンを生成します |
| RevokeToken | [RevokeTokenRequest](#emoine_r-v1-RevokeTokenRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 集会用のトークンを無効化します |

 



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

