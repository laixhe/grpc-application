// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/pbuser/list.proto

package pbuser

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

// 请求-用户列表
type ListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 每页数量（分页）
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 偏移id（分页）
	OffsetId      int64 `protobuf:"varint,2,opt,name=offset_id,json=offsetId,proto3" json:"offset_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_proto_pbuser_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbuser_list_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbuser_list_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequest) GetOffsetId() int64 {
	if x != nil {
		return x.OffsetId
	}
	return 0
}

// 响应-用户列表
type ListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 列表
	Users         []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_proto_pbuser_list_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbuser_list_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbuser_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_pbuser_list_proto protoreflect.FileDescriptor

const file_proto_pbuser_list_proto_rawDesc = "" +
	"\n" +
	"\x17proto/pbuser/list.proto\x12\x06pbuser\x1a\x17proto/pbuser/user.proto\"G\n" +
	"\vListRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x03R\bpageSize\x12\x1b\n" +
	"\toffset_id\x18\x02 \x01(\x03R\boffsetId\"H\n" +
	"\fListResponse\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x03R\x05total\x12\"\n" +
	"\x05users\x18\x02 \x03(\v2\f.pbuser.UserR\x05usersB!Z\x1fgo-kratos/api/gen/pbuser;pbuserb\x06proto3"

var (
	file_proto_pbuser_list_proto_rawDescOnce sync.Once
	file_proto_pbuser_list_proto_rawDescData []byte
)

func file_proto_pbuser_list_proto_rawDescGZIP() []byte {
	file_proto_pbuser_list_proto_rawDescOnce.Do(func() {
		file_proto_pbuser_list_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pbuser_list_proto_rawDesc), len(file_proto_pbuser_list_proto_rawDesc)))
	})
	return file_proto_pbuser_list_proto_rawDescData
}

var file_proto_pbuser_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pbuser_list_proto_goTypes = []any{
	(*ListRequest)(nil),  // 0: pbuser.ListRequest
	(*ListResponse)(nil), // 1: pbuser.ListResponse
	(*User)(nil),         // 2: pbuser.User
}
var file_proto_pbuser_list_proto_depIdxs = []int32{
	2, // 0: pbuser.ListResponse.users:type_name -> pbuser.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_pbuser_list_proto_init() }
func file_proto_pbuser_list_proto_init() {
	if File_proto_pbuser_list_proto != nil {
		return
	}
	file_proto_pbuser_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pbuser_list_proto_rawDesc), len(file_proto_pbuser_list_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_pbuser_list_proto_goTypes,
		DependencyIndexes: file_proto_pbuser_list_proto_depIdxs,
		MessageInfos:      file_proto_pbuser_list_proto_msgTypes,
	}.Build()
	File_proto_pbuser_list_proto = out.File
	file_proto_pbuser_list_proto_goTypes = nil
	file_proto_pbuser_list_proto_depIdxs = nil
}
