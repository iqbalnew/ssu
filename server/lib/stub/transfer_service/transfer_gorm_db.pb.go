// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: transfer_gorm_db.proto

// import "google/protobuf/timestamp.proto";
// import "google/api/field_behavior.proto";
// import "protoc-gen-gorm/options/gorm.proto";
// import "protoc-gen-openapiv2/options/annotations.proto";
// import "mwitkow/go-proto-validators/validator.proto";

package pb

import (
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

var File_transfer_gorm_db_proto protoreflect.FileDescriptor

var file_transfer_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_transfer_gorm_db_proto_goTypes = []interface{}{}
var file_transfer_gorm_db_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_gorm_db_proto_init() }
func file_transfer_gorm_db_proto_init() {
	if File_transfer_gorm_db_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transfer_gorm_db_proto_goTypes,
		DependencyIndexes: file_transfer_gorm_db_proto_depIdxs,
	}.Build()
	File_transfer_gorm_db_proto = out.File
	file_transfer_gorm_db_proto_rawDesc = nil
	file_transfer_gorm_db_proto_goTypes = nil
	file_transfer_gorm_db_proto_depIdxs = nil
}
