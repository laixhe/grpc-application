// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/pbuser/server.proto

package pbuser

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_pbuser_server_proto protoreflect.FileDescriptor

const file_proto_pbuser_server_proto_rawDesc = "" +
	"\n" +
	"\x19proto/pbuser/server.proto\x12\x06pbuser\x1a\x1cgoogle/api/annotations.proto\x1a\x17proto/pbuser/user.proto\x1a\x17proto/pbuser/info.proto\x1a\x17proto/pbuser/list.proto\x1a\x19proto/pbuser/update.proto2\xea\x01\n" +
	"\x05SUser\x12D\n" +
	"\x04Info\x12\x13.pbuser.InfoRequest\x1a\f.pbuser.User\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/api/v1/user/info\x12L\n" +
	"\x04List\x12\x13.pbuser.ListRequest\x1a\x14.pbuser.ListResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/api/v1/user/list\x12M\n" +
	"\x06Update\x12\x15.pbuser.UpdateRequest\x1a\f.pbuser.User\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/api/v1/user/updateB!Z\x1fgo-kratos/api/gen/pbuser;pbuserb\x06proto3"

var file_proto_pbuser_server_proto_goTypes = []any{
	(*InfoRequest)(nil),   // 0: pbuser.InfoRequest
	(*ListRequest)(nil),   // 1: pbuser.ListRequest
	(*UpdateRequest)(nil), // 2: pbuser.UpdateRequest
	(*User)(nil),          // 3: pbuser.User
	(*ListResponse)(nil),  // 4: pbuser.ListResponse
}
var file_proto_pbuser_server_proto_depIdxs = []int32{
	0, // 0: pbuser.SUser.Info:input_type -> pbuser.InfoRequest
	1, // 1: pbuser.SUser.List:input_type -> pbuser.ListRequest
	2, // 2: pbuser.SUser.Update:input_type -> pbuser.UpdateRequest
	3, // 3: pbuser.SUser.Info:output_type -> pbuser.User
	4, // 4: pbuser.SUser.List:output_type -> pbuser.ListResponse
	3, // 5: pbuser.SUser.Update:output_type -> pbuser.User
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pbuser_server_proto_init() }
func file_proto_pbuser_server_proto_init() {
	if File_proto_pbuser_server_proto != nil {
		return
	}
	file_proto_pbuser_user_proto_init()
	file_proto_pbuser_info_proto_init()
	file_proto_pbuser_list_proto_init()
	file_proto_pbuser_update_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pbuser_server_proto_rawDesc), len(file_proto_pbuser_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pbuser_server_proto_goTypes,
		DependencyIndexes: file_proto_pbuser_server_proto_depIdxs,
	}.Build()
	File_proto_pbuser_server_proto = out.File
	file_proto_pbuser_server_proto_goTypes = nil
	file_proto_pbuser_server_proto_depIdxs = nil
}
