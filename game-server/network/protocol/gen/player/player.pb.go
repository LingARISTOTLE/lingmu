// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: proto/player.proto

package player

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建用户包
type CSCreateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *CSCreateUser) Reset() {
	*x = CSCreateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSCreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSCreateUser) ProtoMessage() {}

func (x *CSCreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSCreateUser.ProtoReflect.Descriptor instead.
func (*CSCreateUser) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

func (x *CSCreateUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CSCreateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SCCreateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SCCreateUser) Reset() {
	*x = SCCreateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCCreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCCreateUser) ProtoMessage() {}

func (x *SCCreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCCreateUser.ProtoReflect.Descriptor instead.
func (*SCCreateUser) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

// 用户登录
type CSLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uint64 UId = 1;//玩家Id
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *CSLogin) Reset() {
	*x = CSLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSLogin) ProtoMessage() {}

func (x *CSLogin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSLogin.ProtoReflect.Descriptor instead.
func (*CSLogin) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

func (x *CSLogin) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CSLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SCLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"` //请求成功与否描述
}

func (x *SCLogin) Reset() {
	*x = SCLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCLogin) ProtoMessage() {}

func (x *SCLogin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCLogin.ProtoReflect.Descriptor instead.
func (*SCLogin) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{3}
}

func (x *SCLogin) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type CSAddFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId uint64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"` //玩家Id
}

func (x *CSAddFriend) Reset() {
	*x = CSAddFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSAddFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSAddFriend) ProtoMessage() {}

func (x *CSAddFriend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSAddFriend.ProtoReflect.Descriptor instead.
func (*CSAddFriend) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{4}
}

func (x *CSAddFriend) GetUId() uint64 {
	if x != nil {
		return x.UId
	}
	return 0
}

type SCAddFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc string `protobuf:"bytes,1,opt,name=Desc,proto3" json:"Desc,omitempty"` //请求成功与否描述
}

func (x *SCAddFriend) Reset() {
	*x = SCAddFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCAddFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCAddFriend) ProtoMessage() {}

func (x *SCAddFriend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCAddFriend.ProtoReflect.Descriptor instead.
func (*SCAddFriend) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{5}
}

func (x *SCAddFriend) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CSDelFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId uint64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"` //玩家Id
}

func (x *CSDelFriend) Reset() {
	*x = CSDelFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSDelFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSDelFriend) ProtoMessage() {}

func (x *CSDelFriend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSDelFriend.ProtoReflect.Descriptor instead.
func (*CSDelFriend) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{6}
}

func (x *CSDelFriend) GetUId() uint64 {
	if x != nil {
		return x.UId
	}
	return 0
}

type SCDelFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc string `protobuf:"bytes,1,opt,name=Desc,proto3" json:"Desc,omitempty"` //请求成功与否描述
}

func (x *SCDelFriend) Reset() {
	*x = SCDelFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCDelFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCDelFriend) ProtoMessage() {}

func (x *SCDelFriend) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCDelFriend.ProtoReflect.Descriptor instead.
func (*SCDelFriend) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{7}
}

func (x *SCDelFriend) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CSSendChatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId      uint64       `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`           //玩家Id
	Msg      *ChatMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`            //消息本体
	Category int32        `protobuf:"varint,3,opt,name=Category,proto3" json:"Category,omitempty"` //聊天类型
}

func (x *CSSendChatMsg) Reset() {
	*x = CSSendChatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSSendChatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSSendChatMsg) ProtoMessage() {}

func (x *CSSendChatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSSendChatMsg.ProtoReflect.Descriptor instead.
func (*CSSendChatMsg) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{8}
}

func (x *CSSendChatMsg) GetUId() uint64 {
	if x != nil {
		return x.UId
	}
	return 0
}

func (x *CSSendChatMsg) GetMsg() *ChatMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *CSSendChatMsg) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

type SCSendChatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SCSendChatMsg) Reset() {
	*x = SCSendChatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCSendChatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCSendChatMsg) ProtoMessage() {}

func (x *SCSendChatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCSendChatMsg.ProtoReflect.Descriptor instead.
func (*SCSendChatMsg) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{9}
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"` //私信消息内容
	Extra   [][]byte `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{10}
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetExtra() [][]byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x0c,
	0x43, 0x53, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x43, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x07, 0x43, 0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x53, 0x43, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x4f, 0x6b, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x53, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x55, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x43, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x53, 0x44, 0x65, 0x6c, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x55, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x43, 0x44, 0x65, 0x6c,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x22, 0x64, 0x0a, 0x0d, 0x43, 0x53,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x55, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x43, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73,
	0x67, 0x22, 0x3d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_player_proto_goTypes = []interface{}{
	(*CSCreateUser)(nil),  // 0: player.CSCreateUser
	(*SCCreateUser)(nil),  // 1: player.SCCreateUser
	(*CSLogin)(nil),       // 2: player.CSLogin
	(*SCLogin)(nil),       // 3: player.SCLogin
	(*CSAddFriend)(nil),   // 4: player.CSAddFriend
	(*SCAddFriend)(nil),   // 5: player.SCAddFriend
	(*CSDelFriend)(nil),   // 6: player.CSDelFriend
	(*SCDelFriend)(nil),   // 7: player.SCDelFriend
	(*CSSendChatMsg)(nil), // 8: player.CSSendChatMsg
	(*SCSendChatMsg)(nil), // 9: player.SCSendChatMsg
	(*ChatMessage)(nil),   // 10: player.ChatMessage
}
var file_proto_player_proto_depIdxs = []int32{
	10, // 0: player.CSSendChatMsg.msg:type_name -> player.ChatMessage
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSCreateUser); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCCreateUser); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSLogin); i {
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
		file_proto_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCLogin); i {
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
		file_proto_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSAddFriend); i {
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
		file_proto_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCAddFriend); i {
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
		file_proto_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSDelFriend); i {
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
		file_proto_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCDelFriend); i {
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
		file_proto_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSSendChatMsg); i {
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
		file_proto_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCSendChatMsg); i {
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
		file_proto_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}
