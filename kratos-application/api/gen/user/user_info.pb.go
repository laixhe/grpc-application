// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/user/user_info.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求-用户信息
type UserInfoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 用户ID
	Uid           uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	mi := &file_proto_user_user_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// 响应-用户信息
type UserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	mi := &file_proto_user_user_info_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_info_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_user_user_info_proto protoreflect.FileDescriptor

var file_proto_user_user_info_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x32,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x42, 0x26, 0x5a, 0x24, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_proto_user_user_info_proto_rawDescOnce sync.Once
	file_proto_user_user_info_proto_rawDescData []byte
)

func file_proto_user_user_info_proto_rawDescGZIP() []byte {
	file_proto_user_user_info_proto_rawDescOnce.Do(func() {
		file_proto_user_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_user_info_proto_rawDesc), len(file_proto_user_user_info_proto_rawDesc)))
	})
	return file_proto_user_user_info_proto_rawDescData
}

var file_proto_user_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_user_info_proto_goTypes = []any{
	(*UserInfoRequest)(nil),  // 0: user.UserInfoRequest
	(*UserInfoResponse)(nil), // 1: user.UserInfoResponse
	(*User)(nil),             // 2: user.User
}
var file_proto_user_user_info_proto_depIdxs = []int32{
	2, // 0: user.UserInfoResponse.user:type_name -> user.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_user_info_proto_init() }
func file_proto_user_user_info_proto_init() {
	if File_proto_user_user_info_proto != nil {
		return
	}
	file_proto_user_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_user_info_proto_rawDesc), len(file_proto_user_user_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_user_info_proto_goTypes,
		DependencyIndexes: file_proto_user_user_info_proto_depIdxs,
		MessageInfos:      file_proto_user_user_info_proto_msgTypes,
	}.Build()
	File_proto_user_user_info_proto = out.File
	file_proto_user_user_info_proto_goTypes = nil
	file_proto_user_user_info_proto_depIdxs = nil
}
