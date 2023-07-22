// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: emoine_r/v1/schema.proto

package emoine_rv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Meeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 集会ID (UUID)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// YouTubeの動画ID
	VideoId string `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	// 集会の説明
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// 動画タイトル (YouTubeから取得)
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// 動画サムネイル (YouTubeから取得)
	Thumbnail string `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// 集会の開始日時 (YouTubeから取得)
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// 集会の終了日時 (YouTubeから取得)
	EndedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
}

func (x *Meeting) Reset() {
	*x = Meeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meeting) ProtoMessage() {}

func (x *Meeting) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meeting.ProtoReflect.Descriptor instead.
func (*Meeting) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Meeting) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meeting) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *Meeting) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Meeting) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Meeting) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Meeting) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Meeting) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// トークン文字列
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// トークンの所有者名
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// トークンが有効な集会のID
	MeetingId string `protobuf:"bytes,3,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// トークン発行者のtraQID (traQと同期)
	CreatorId string `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	// トークンの説明
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// トークンの作成日時
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// トークンの有効期限
	ExpireAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Token) GetMeetingId() string {
	if x != nil {
		return x.MeetingId
	}
	return ""
}

func (x *Token) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Token) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Token) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Token) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// コメントのID (UUID)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// コメントが送信された集会のID
	MeetingId string `protobuf:"bytes,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// コメントの送信者名
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// コメントの本文
	Text string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// コメントが匿名かどうか
	IsAnonymous bool `protobuf:"varint,5,opt,name=is_anonymous,json=isAnonymous,proto3" json:"is_anonymous,omitempty"`
	// コメントの色
	Color string `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	// コメントの作成日時
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetMeetingId() string {
	if x != nil {
		return x.MeetingId
	}
	return ""
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetIsAnonymous() bool {
	if x != nil {
		return x.IsAnonymous
	}
	return false
}

func (x *Comment) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// リアクションのID (UUID)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// リアクションが送信された集会のID
	MeetingId string `protobuf:"bytes,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// リアクションの送信者名
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// 送信されたリアクションのスタンプID (UUID, traQと同期)
	StampId string `protobuf:"bytes,4,opt,name=stamp_id,json=stampId,proto3" json:"stamp_id,omitempty"`
	// リアクションの作成日時
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *Reaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reaction) GetMeetingId() string {
	if x != nil {
		return x.MeetingId
	}
	return ""
}

func (x *Reaction) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Reaction) GetStampId() string {
	if x != nil {
		return x.StampId
	}
	return ""
}

func (x *Reaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_emoine_r_v1_schema_proto protoreflect.FileDescriptor

var file_emoine_r_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6d, 0x6f, 0x69,
	0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8d, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0xa9, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6d, 0x6f,
	0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x50, 0x74, 0x69, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x45,
	0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x52, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x66, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x45,
	0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x45, 0x6d, 0x6f, 0x69,
	0x6e, 0x65, 0x52, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emoine_r_v1_schema_proto_rawDescOnce sync.Once
	file_emoine_r_v1_schema_proto_rawDescData = file_emoine_r_v1_schema_proto_rawDesc
)

func file_emoine_r_v1_schema_proto_rawDescGZIP() []byte {
	file_emoine_r_v1_schema_proto_rawDescOnce.Do(func() {
		file_emoine_r_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_emoine_r_v1_schema_proto_rawDescData)
	})
	return file_emoine_r_v1_schema_proto_rawDescData
}

var file_emoine_r_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_emoine_r_v1_schema_proto_goTypes = []interface{}{
	(*Meeting)(nil),               // 0: emoine_r.v1.Meeting
	(*Token)(nil),                 // 1: emoine_r.v1.Token
	(*Comment)(nil),               // 2: emoine_r.v1.Comment
	(*Reaction)(nil),              // 3: emoine_r.v1.Reaction
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_emoine_r_v1_schema_proto_depIdxs = []int32{
	4, // 0: emoine_r.v1.Meeting.started_at:type_name -> google.protobuf.Timestamp
	4, // 1: emoine_r.v1.Meeting.ended_at:type_name -> google.protobuf.Timestamp
	4, // 2: emoine_r.v1.Token.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: emoine_r.v1.Token.expire_at:type_name -> google.protobuf.Timestamp
	4, // 4: emoine_r.v1.Comment.created_at:type_name -> google.protobuf.Timestamp
	4, // 5: emoine_r.v1.Reaction.created_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_emoine_r_v1_schema_proto_init() }
func file_emoine_r_v1_schema_proto_init() {
	if File_emoine_r_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emoine_r_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meeting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_emoine_r_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_emoine_r_v1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_emoine_r_v1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_emoine_r_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_emoine_r_v1_schema_proto_goTypes,
		DependencyIndexes: file_emoine_r_v1_schema_proto_depIdxs,
		MessageInfos:      file_emoine_r_v1_schema_proto_msgTypes,
	}.Build()
	File_emoine_r_v1_schema_proto = out.File
	file_emoine_r_v1_schema_proto_rawDesc = nil
	file_emoine_r_v1_schema_proto_goTypes = nil
	file_emoine_r_v1_schema_proto_depIdxs = nil
}
