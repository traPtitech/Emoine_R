# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [emoine_r/v1/schema.proto](#emoine_r_v1_schema-proto)
    - [Comment](#emoine_r-v1-Comment)
    - [Event](#emoine_r-v1-Event)
    - [Reaction](#emoine_r-v1-Reaction)
    - [Token](#emoine_r-v1-Token)
  
- [emoine_r/v1/admin_api.proto](#emoine_r_v1_admin_api-proto)
    - [CreateEventRequest](#emoine_r-v1-CreateEventRequest)
    - [CreateEventResponse](#emoine_r-v1-CreateEventResponse)
    - [DeleteEventRequest](#emoine_r-v1-DeleteEventRequest)
    - [GenerateTokenRequest](#emoine_r-v1-GenerateTokenRequest)
    - [GenerateTokenResponse](#emoine_r-v1-GenerateTokenResponse)
    - [GetTokensRequest](#emoine_r-v1-GetTokensRequest)
    - [GetTokensResponse](#emoine_r-v1-GetTokensResponse)
    - [RevokeTokenRequest](#emoine_r-v1-RevokeTokenRequest)
    - [UpdateEventRequest](#emoine_r-v1-UpdateEventRequest)
  
    - [AdminAPIService](#emoine_r-v1-AdminAPIService)
  
- [emoine_r/v1/general_api.proto](#emoine_r_v1_general_api-proto)
    - [ConnectToEventStreamRequest](#emoine_r-v1-ConnectToEventStreamRequest)
    - [ConnectToEventStreamResponse](#emoine_r-v1-ConnectToEventStreamResponse)
    - [GetEventCommentsRequest](#emoine_r-v1-GetEventCommentsRequest)
    - [GetEventCommentsResponse](#emoine_r-v1-GetEventCommentsResponse)
    - [GetEventReactionsRequest](#emoine_r-v1-GetEventReactionsRequest)
    - [GetEventReactionsResponse](#emoine_r-v1-GetEventReactionsResponse)
    - [GetEventRequest](#emoine_r-v1-GetEventRequest)
    - [GetEventResponse](#emoine_r-v1-GetEventResponse)
    - [GetEventsRequest](#emoine_r-v1-GetEventsRequest)
    - [GetEventsResponse](#emoine_r-v1-GetEventsResponse)
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
| event_id | [string](#string) |  | コメントが送信されたイベントのID |
| username | [string](#string) |  | コメントの送信者名 |
| text | [string](#string) |  | コメントの本文 |
| is_anonymous | [bool](#bool) |  | コメントが匿名かどうか |
| color | [string](#string) |  | コメントの色 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | コメントの作成日時 |






<a name="emoine_r-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | イベントID (UUID) |
| video_id | [string](#string) |  | YouTubeの動画ID |
| description | [string](#string) |  | イベントの説明 |
| title | [string](#string) |  | 動画タイトル (YouTubeから取得) |
| thumbnail | [string](#string) |  | 動画サムネイル (YouTubeから取得) |
| started_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | イベントの開始日時 (YouTubeから取得) |
| ended_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | イベントの終了日時 (YouTubeから取得) |






<a name="emoine_r-v1-Reaction"></a>

### Reaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | リアクションのID (UUID) |
| event_id | [string](#string) |  | リアクションが送信されたイベントのID |
| username | [string](#string) |  | リアクションの送信者名 |
| stamp_id | [string](#string) |  | 送信されたリアクションのスタンプID (UUID, traQと同期) |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | リアクションの作成日時 |






<a name="emoine_r-v1-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | トークンのID (UUID) |
| raw | [string](#string) |  | トークン文字列 |
| username | [string](#string) |  | トークンの所有者名 |
| event_id | [string](#string) |  | トークンが有効なイベントのID |
| creator_id | [string](#string) |  | トークン発行者のtraQID (traQと同期) |
| description | [string](#string) |  | トークンの説明 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | トークンの作成日時 |
| expire_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | トークンの有効期限 |





 

 

 

 



<a name="emoine_r_v1_admin_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## emoine_r/v1/admin_api.proto



<a name="emoine_r-v1-CreateEventRequest"></a>

### CreateEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| video_id | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="emoine_r-v1-CreateEventResponse"></a>

### CreateEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#emoine_r-v1-Event) |  |  |






<a name="emoine_r-v1-DeleteEventRequest"></a>

### DeleteEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-GenerateTokenRequest"></a>

### GenerateTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |
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
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetTokensResponse"></a>

### GetTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [Token](#emoine_r-v1-Token) | repeated |  |






<a name="emoine_r-v1-RevokeTokenRequest"></a>

### RevokeTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_id | [string](#string) |  |  |
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-UpdateEventRequest"></a>

### UpdateEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |
| video_id | [string](#string) | optional |  |
| description | [string](#string) | optional |  |





 

 

 


<a name="emoine_r-v1-AdminAPIService"></a>

### AdminAPIService
管理者権限が必要なAPIを提供します

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEvent | [CreateEventRequest](#emoine_r-v1-CreateEventRequest) | [CreateEventResponse](#emoine_r-v1-CreateEventResponse) | イベントを作成します |
| UpdateEvent | [UpdateEventRequest](#emoine_r-v1-UpdateEventRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | イベント情報を更新します |
| DeleteEvent | [DeleteEventRequest](#emoine_r-v1-DeleteEventRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | イベントを削除します |
| GetTokens | [GetTokensRequest](#emoine_r-v1-GetTokensRequest) | [GetTokensResponse](#emoine_r-v1-GetTokensResponse) | 該当するイベントのトークン一覧を取得します |
| GenerateToken | [GenerateTokenRequest](#emoine_r-v1-GenerateTokenRequest) | [GenerateTokenResponse](#emoine_r-v1-GenerateTokenResponse) | イベント用のトークンを生成します |
| RevokeToken | [RevokeTokenRequest](#emoine_r-v1-RevokeTokenRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | イベント用のトークンを無効化します |

 



<a name="emoine_r_v1_general_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## emoine_r/v1/general_api.proto



<a name="emoine_r-v1-ConnectToEventStreamRequest"></a>

### ConnectToEventStreamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-ConnectToEventStreamResponse"></a>

### ConnectToEventStreamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#emoine_r-v1-Event) |  |  |
| comment | [Comment](#emoine_r-v1-Comment) |  |  |
| reaction | [Reaction](#emoine_r-v1-Reaction) |  |  |






<a name="emoine_r-v1-GetEventCommentsRequest"></a>

### GetEventCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetEventCommentsResponse"></a>

### GetEventCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| comments | [Comment](#emoine_r-v1-Comment) | repeated |  |






<a name="emoine_r-v1-GetEventReactionsRequest"></a>

### GetEventReactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |






<a name="emoine_r-v1-GetEventReactionsResponse"></a>

### GetEventReactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reactions | [Reaction](#emoine_r-v1-Reaction) | repeated | reactions are sorted by created_at asc (oldest first) |






<a name="emoine_r-v1-GetEventRequest"></a>

### GetEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="emoine_r-v1-GetEventResponse"></a>

### GetEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#emoine_r-v1-Event) |  |  |






<a name="emoine_r-v1-GetEventsRequest"></a>

### GetEventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) | optional | default: 10 |
| offset | [int32](#int32) | optional | default: 0 |






<a name="emoine_r-v1-GetEventsResponse"></a>

### GetEventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  |  |
| events | [Event](#emoine_r-v1-Event) | repeated |  |






<a name="emoine_r-v1-SendCommentRequest"></a>

### SendCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_id | [string](#string) |  |  |
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
| event_id | [string](#string) |  |  |
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
| GetEvents | [GetEventsRequest](#emoine_r-v1-GetEventsRequest) | [GetEventsResponse](#emoine_r-v1-GetEventsResponse) | イベント一覧を取得します |
| GetEvent | [GetEventRequest](#emoine_r-v1-GetEventRequest) | [GetEventResponse](#emoine_r-v1-GetEventResponse) | 該当するイベントを取得します |
| GetEventComments | [GetEventCommentsRequest](#emoine_r-v1-GetEventCommentsRequest) | [GetEventCommentsResponse](#emoine_r-v1-GetEventCommentsResponse) | 該当するイベントのコメント一覧を取得します |
| GetEventReactions | [GetEventReactionsRequest](#emoine_r-v1-GetEventReactionsRequest) | [GetEventReactionsResponse](#emoine_r-v1-GetEventReactionsResponse) | 該当するイベントのリアクション一覧を取得します |
| ConnectToEventStream | [ConnectToEventStreamRequest](#emoine_r-v1-ConnectToEventStreamRequest) | [ConnectToEventStreamResponse](#emoine_r-v1-ConnectToEventStreamResponse) stream | イベントのストリームに接続します |
| SendComment | [SendCommentRequest](#emoine_r-v1-SendCommentRequest) | [SendCommentResponse](#emoine_r-v1-SendCommentResponse) | (コメントはイベントのストリームに反映されます) |
| SendReaction | [SendReactionRequest](#emoine_r-v1-SendReactionRequest) | [SendReactionResponse](#emoine_r-v1-SendReactionResponse) | (リアクションはイベントのストリームに反映されます) |

 



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

