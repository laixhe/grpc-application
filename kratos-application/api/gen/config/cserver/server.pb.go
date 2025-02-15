// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/config/cserver/server.proto

package cserver

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

// 服务器配置
type Server struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 运行IP
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// 运行端口
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// 超时时间
	Timeout       *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_proto_config_cserver_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_cserver_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_proto_config_cserver_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Server) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Server) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// 服务器组
type Servers struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server                `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server                `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Servers) Reset() {
	*x = Servers{}
	mi := &file_proto_config_cserver_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Servers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servers) ProtoMessage() {}

func (x *Servers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_cserver_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Servers.ProtoReflect.Descriptor instead.
func (*Servers) Descriptor() ([]byte, []int) {
	return file_proto_config_cserver_server_proto_rawDescGZIP(), []int{1}
}

func (x *Servers) GetHttp() *Server {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Servers) GetGrpc() *Server {
	if x != nil {
		return x.Grpc
	}
	return nil
}

var File_proto_config_cserver_server_proto protoreflect.FileDescriptor

var file_proto_config_cserver_server_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x06,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x53, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12,
	0x23, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x42, 0x33, 0x5a, 0x31, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x3b, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_proto_config_cserver_server_proto_rawDescOnce sync.Once
	file_proto_config_cserver_server_proto_rawDescData []byte
)

func file_proto_config_cserver_server_proto_rawDescGZIP() []byte {
	file_proto_config_cserver_server_proto_rawDescOnce.Do(func() {
		file_proto_config_cserver_server_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_config_cserver_server_proto_rawDesc), len(file_proto_config_cserver_server_proto_rawDesc)))
	})
	return file_proto_config_cserver_server_proto_rawDescData
}

var file_proto_config_cserver_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_config_cserver_server_proto_goTypes = []any{
	(*Server)(nil),              // 0: cserver.Server
	(*Servers)(nil),             // 1: cserver.Servers
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_proto_config_cserver_server_proto_depIdxs = []int32{
	2, // 0: cserver.Server.timeout:type_name -> google.protobuf.Duration
	0, // 1: cserver.Servers.http:type_name -> cserver.Server
	0, // 2: cserver.Servers.grpc:type_name -> cserver.Server
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_config_cserver_server_proto_init() }
func file_proto_config_cserver_server_proto_init() {
	if File_proto_config_cserver_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_config_cserver_server_proto_rawDesc), len(file_proto_config_cserver_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_cserver_server_proto_goTypes,
		DependencyIndexes: file_proto_config_cserver_server_proto_depIdxs,
		MessageInfos:      file_proto_config_cserver_server_proto_msgTypes,
	}.Build()
	File_proto_config_cserver_server_proto = out.File
	file_proto_config_cserver_server_proto_goTypes = nil
	file_proto_config_cserver_server_proto_depIdxs = nil
}
