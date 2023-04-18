// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: proto/messageId.proto

package messageId

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

// 定义Message消息的id
type MessageId int32

const (
	MessageId_None MessageId = 0
	// 用户创建
	MessageId_CSCreatePlayer MessageId = 50001
	MessageId_SCCreatePlayer MessageId = 50002
	// 客户端登录请求
	MessageId_CSLogin MessageId = 50003
	// 服务器响应登录请求
	MessageId_SCLogin MessageId = 50004
	// player
	// 添加好友
	MessageId_CSAddFriend MessageId = 100001
	MessageId_SCAddFriend MessageId = 100002
	// 删除好友
	MessageId_CSDelFriend MessageId = 100003
	MessageId_SCDelFriend MessageId = 100004
	// 聊天包
	MessageId_CSSendChatMsg                   MessageId = 100005
	MessageId_SCSendChatMsg                   MessageId = 100006
	MessageId_ClientRequestObjectSync         MessageId = 200010
	MessageId_ServerRequestObjectSync         MessageId = 200011
	MessageId_ServerRequestObjectSyncComplete MessageId = 200012
	MessageId_Instantiate                     MessageId = 200020
	MessageId_Destroy                         MessageId = 200021
	MessageId_DestroyNetworkObjects           MessageId = 200022
	MessageId_SyncTransform                   MessageId = 200030
)

// Enum value maps for MessageId.
var (
	MessageId_name = map[int32]string{
		0:      "None",
		50001:  "CSCreatePlayer",
		50002:  "SCCreatePlayer",
		50003:  "CSLogin",
		50004:  "SCLogin",
		100001: "CSAddFriend",
		100002: "SCAddFriend",
		100003: "CSDelFriend",
		100004: "SCDelFriend",
		100005: "CSSendChatMsg",
		100006: "SCSendChatMsg",
		200010: "ClientRequestObjectSync",
		200011: "ServerRequestObjectSync",
		200012: "ServerRequestObjectSyncComplete",
		200020: "Instantiate",
		200021: "Destroy",
		200022: "DestroyNetworkObjects",
		200030: "SyncTransform",
	}
	MessageId_value = map[string]int32{
		"None":                            0,
		"CSCreatePlayer":                  50001,
		"SCCreatePlayer":                  50002,
		"CSLogin":                         50003,
		"SCLogin":                         50004,
		"CSAddFriend":                     100001,
		"SCAddFriend":                     100002,
		"CSDelFriend":                     100003,
		"SCDelFriend":                     100004,
		"CSSendChatMsg":                   100005,
		"SCSendChatMsg":                   100006,
		"ClientRequestObjectSync":         200010,
		"ServerRequestObjectSync":         200011,
		"ServerRequestObjectSyncComplete": 200012,
		"Instantiate":                     200020,
		"Destroy":                         200021,
		"DestroyNetworkObjects":           200022,
		"SyncTransform":                   200030,
	}
)

func (x MessageId) Enum() *MessageId {
	p := new(MessageId)
	*p = x
	return p
}

func (x MessageId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messageId_proto_enumTypes[0].Descriptor()
}

func (MessageId) Type() protoreflect.EnumType {
	return &file_proto_messageId_proto_enumTypes[0]
}

func (x MessageId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageId.Descriptor instead.
func (MessageId) EnumDescriptor() ([]byte, []int) {
	return file_proto_messageId_proto_rawDescGZIP(), []int{0}
}

var File_proto_messageId_proto protoreflect.FileDescriptor

var file_proto_messageId_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x2a, 0x8e, 0x03, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x0e, 0x43, 0x53,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0xd1, 0x86, 0x03,
	0x12, 0x14, 0x0a, 0x0e, 0x53, 0x43, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x0d, 0x0a, 0x07, 0x43, 0x53, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x10, 0xd3, 0x86, 0x03, 0x12, 0x0d, 0x0a, 0x07, 0x53, 0x43, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x10, 0xd4, 0x86, 0x03, 0x12, 0x11, 0x0a, 0x0b, 0x43, 0x53, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x53, 0x43, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa2, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x43, 0x53,
	0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa3, 0x8d, 0x06, 0x12, 0x11, 0x0a,
	0x0b, 0x53, 0x43, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa4, 0x8d, 0x06,
	0x12, 0x13, 0x0a, 0x0d, 0x43, 0x53, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73,
	0x67, 0x10, 0xa5, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x53, 0x43, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xa6, 0x8d, 0x06, 0x12, 0x1d, 0x0a, 0x17, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xca, 0x9a, 0x0c, 0x12, 0x1d, 0x0a, 0x17, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x10, 0xcb, 0x9a, 0x0c, 0x12, 0x25, 0x0a, 0x1f, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0xcc, 0x9a, 0x0c, 0x12,
	0x11, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x10, 0xd4,
	0x9a, 0x0c, 0x12, 0x0d, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x10, 0xd5, 0x9a,
	0x0c, 0x12, 0x1b, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x10, 0xd6, 0x9a, 0x0c, 0x12, 0x13,
	0x0a, 0x0d, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x10,
	0xde, 0x9a, 0x0c, 0x42, 0x18, 0x5a, 0x0a, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0xaa, 0x02, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messageId_proto_rawDescOnce sync.Once
	file_proto_messageId_proto_rawDescData = file_proto_messageId_proto_rawDesc
)

func file_proto_messageId_proto_rawDescGZIP() []byte {
	file_proto_messageId_proto_rawDescOnce.Do(func() {
		file_proto_messageId_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messageId_proto_rawDescData)
	})
	return file_proto_messageId_proto_rawDescData
}

var file_proto_messageId_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_messageId_proto_goTypes = []interface{}{
	(MessageId)(0), // 0: messageId.MessageId
}
var file_proto_messageId_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_messageId_proto_init() }
func file_proto_messageId_proto_init() {
	if File_proto_messageId_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_messageId_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_messageId_proto_goTypes,
		DependencyIndexes: file_proto_messageId_proto_depIdxs,
		EnumInfos:         file_proto_messageId_proto_enumTypes,
	}.Build()
	File_proto_messageId_proto = out.File
	file_proto_messageId_proto_rawDesc = nil
	file_proto_messageId_proto_goTypes = nil
	file_proto_messageId_proto_depIdxs = nil
}
