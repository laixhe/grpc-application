// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/config/cmongodb/mongodb.proto

package cmongodb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// MongoDB数据库配置
type MongoDB struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 连接地址
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// 指定数据库
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// 最大连接的数量
	MaxPoolSize uint64 `protobuf:"varint,3,opt,name=max_pool_size,json=maxPoolSize,proto3" json:"max_pool_size,omitempty"`
	// 最小连接的数量
	MinPoolSize uint64 `protobuf:"varint,4,opt,name=min_pool_size,json=minPoolSize,proto3" json:"min_pool_size,omitempty"`
	// 最大连接的空闲时间(设置了连接可复用的最大时间)(单位秒)
	MaxConnIdleTime *durationpb.Duration `protobuf:"bytes,5,opt,name=max_conn_idle_time,json=maxConnIdleTime,proto3" json:"max_conn_idle_time,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MongoDB) Reset() {
	*x = MongoDB{}
	mi := &file_proto_config_cmongodb_mongodb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoDB) ProtoMessage() {}

func (x *MongoDB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_cmongodb_mongodb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoDB.ProtoReflect.Descriptor instead.
func (*MongoDB) Descriptor() ([]byte, []int) {
	return file_proto_config_cmongodb_mongodb_proto_rawDescGZIP(), []int{0}
}

func (x *MongoDB) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *MongoDB) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MongoDB) GetMaxPoolSize() uint64 {
	if x != nil {
		return x.MaxPoolSize
	}
	return 0
}

func (x *MongoDB) GetMinPoolSize() uint64 {
	if x != nil {
		return x.MinPoolSize
	}
	return 0
}

func (x *MongoDB) GetMaxConnIdleTime() *durationpb.Duration {
	if x != nil {
		return x.MaxConnIdleTime
	}
	return nil
}

var File_proto_config_cmongodb_mongodb_proto protoreflect.FileDescriptor

var file_proto_config_cmongodb_mongodb_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc7, 0x01, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x46, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64,
	0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e,
	0x6e, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x3b, 0x63, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_config_cmongodb_mongodb_proto_rawDescOnce sync.Once
	file_proto_config_cmongodb_mongodb_proto_rawDescData []byte
)

func file_proto_config_cmongodb_mongodb_proto_rawDescGZIP() []byte {
	file_proto_config_cmongodb_mongodb_proto_rawDescOnce.Do(func() {
		file_proto_config_cmongodb_mongodb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_config_cmongodb_mongodb_proto_rawDesc), len(file_proto_config_cmongodb_mongodb_proto_rawDesc)))
	})
	return file_proto_config_cmongodb_mongodb_proto_rawDescData
}

var file_proto_config_cmongodb_mongodb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_config_cmongodb_mongodb_proto_goTypes = []any{
	(*MongoDB)(nil),             // 0: cmongodb.MongoDB
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_proto_config_cmongodb_mongodb_proto_depIdxs = []int32{
	1, // 0: cmongodb.MongoDB.max_conn_idle_time:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_config_cmongodb_mongodb_proto_init() }
func file_proto_config_cmongodb_mongodb_proto_init() {
	if File_proto_config_cmongodb_mongodb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_config_cmongodb_mongodb_proto_rawDesc), len(file_proto_config_cmongodb_mongodb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_cmongodb_mongodb_proto_goTypes,
		DependencyIndexes: file_proto_config_cmongodb_mongodb_proto_depIdxs,
		MessageInfos:      file_proto_config_cmongodb_mongodb_proto_msgTypes,
	}.Build()
	File_proto_config_cmongodb_mongodb_proto = out.File
	file_proto_config_cmongodb_mongodb_proto_goTypes = nil
	file_proto_config_cmongodb_mongodb_proto_depIdxs = nil
}
