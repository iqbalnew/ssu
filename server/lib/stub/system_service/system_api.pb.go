// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: system_api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_system_api_proto protoreflect.FileDescriptor

var file_system_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x05, 0x0a, 0x0a, 0x41, 0x70,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x82, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x26, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x1e, 0x1a, 0x19, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x59, 0x12, 0x32,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x2f, 0x20, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x64, 0x64, 0x20, 0x28, 0x69, 0x64, 0x29, 0x20, 0x74, 0x6f, 0x20,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x74, 0x61,
	0x73, 0x6b, 0x1a, 0x21, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x00, 0x12, 0xcb, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x13, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6d, 0x65,
	0x5a, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x3e, 0x12, 0x18, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x6d, 0x79, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x69,
	0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02,
	0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_system_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: system.service.v1.Empty
	(*CreateRequest)(nil),            // 1: system.service.v1.CreateRequest
	(*CreateSystemRequest)(nil),      // 2: system.service.v1.CreateSystemRequest
	(*ListSystemRequest)(nil),        // 3: system.service.v1.ListSystemRequest
	(*HealthCheckResponse)(nil),      // 4: system.service.v1.HealthCheckResponse
	(*CreateResponse)(nil),           // 5: system.service.v1.CreateResponse
	(*CreateTaskSystemResponse)(nil), // 6: system.service.v1.CreateTaskSystemResponse
	(*ListSystemResponse)(nil),       // 7: system.service.v1.ListSystemResponse
}
var file_system_api_proto_depIdxs = []int32{
	0, // 0: system.service.v1.ApiService.HealthCheck:input_type -> system.service.v1.Empty
	1, // 1: system.service.v1.ApiService.CreateSystem:input_type -> system.service.v1.CreateRequest
	2, // 2: system.service.v1.ApiService.CreateSystemParam:input_type -> system.service.v1.CreateSystemRequest
	3, // 3: system.service.v1.ApiService.GetMyTasks:input_type -> system.service.v1.ListSystemRequest
	4, // 4: system.service.v1.ApiService.HealthCheck:output_type -> system.service.v1.HealthCheckResponse
	5, // 5: system.service.v1.ApiService.CreateSystem:output_type -> system.service.v1.CreateResponse
	6, // 6: system.service.v1.ApiService.CreateSystemParam:output_type -> system.service.v1.CreateTaskSystemResponse
	7, // 7: system.service.v1.ApiService.GetMyTasks:output_type -> system.service.v1.ListSystemResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_api_proto_init() }
func file_system_api_proto_init() {
	if File_system_api_proto != nil {
		return
	}
	file_system_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_api_proto_goTypes,
		DependencyIndexes: file_system_api_proto_depIdxs,
	}.Build()
	File_system_api_proto = out.File
	file_system_api_proto_rawDesc = nil
	file_system_api_proto_goTypes = nil
	file_system_api_proto_depIdxs = nil
}
