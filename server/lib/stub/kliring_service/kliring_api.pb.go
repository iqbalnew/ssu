// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: kliring_api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
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

var File_kliring_api_proto protoreflect.FileDescriptor

var file_kliring_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x1a, 0x0a, 0x0a, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb0, 0x01, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x6b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0x9a, 0x02, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e,
	0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x86, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x92, 0x41, 0x5f, 0x0a, 0x16, 0x4b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a,
	0x28, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0xc0, 0x02, 0x0a, 0x22, 0x47, 0x65,
	0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x3d, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3e, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9a, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x6e, 0x0a, 0x17,
	0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x22, 0x47, 0x65, 0x74, 0x20, 0x4b, 0x6c, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x20, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x2f, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x4b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0xd8, 0x02, 0x0a,
	0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x29, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x3b, 0x2e, 0x6b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcc, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x49, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x5a, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7a, 0x0a, 0x17, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x28, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6f,
	0x72, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x35, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0xb6, 0x02, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x6b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x99, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x28, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x92, 0x41, 0x66, 0x0a, 0x17, 0x4b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x1a, 0x2b, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x29, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x44, 0x0a, 0x07, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x15, 0x47, 0x65, 0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x54,
	0x61, 0x73, 0x6b, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xe8, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x48, 0x0a,
	0x07, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x4b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x20, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x1a, 0x24, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x67, 0x65, 0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61, 0x73, 0x6b,
	0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0xde, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d,
	0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0x82, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x7d, 0x92, 0x41, 0x52, 0x0a, 0x07, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20,
	0x54, 0x61, 0x73, 0x6b, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x2d, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x74, 0x6f, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x12, 0xf3, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2c,
	0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x80, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x58, 0x0a, 0x07, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20,
	0x54, 0x61, 0x73, 0x6b, 0x20, 0x57, 0x69, 0x74, 0x68, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x2c, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61,
	0x73, 0x6b, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xbc,
	0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x32, 0x2e, 0x6b, 0x6c, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xb7, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x22, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x69, 0x0a, 0x0e, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x24, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4f, 0x72, 0x20,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x31, 0x54, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x12, 0xa2, 0x02,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x34, 0x2e, 0x6b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a,
	0x92, 0x41, 0x6f, 0x0a, 0x10, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4f, 0x72,
	0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x33, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x4b, 0x6c, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x74, 0x61,
	0x73, 0x6b, 0x12, 0x60, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x26, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x4b, 0x6c, 0x69, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b,
	0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x6c, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x87,
	0x01, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x6b, 0x6c,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x6b, 0x6c, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4b, 0x6c, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_kliring_api_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),                         // 0: kliring.service.v1.HealthCheckRequest
	(*GetKliringTaskSingleTemplateRequest)(nil),        // 1: kliring.service.v1.GetKliringTaskSingleTemplateRequest
	(*GetKliringTaskSingleTemplateDetailRequest)(nil),  // 2: kliring.service.v1.GetKliringTaskSingleTemplateDetailRequest
	(*KliringSingleTemplate)(nil),                      // 3: kliring.service.v1.KliringSingleTemplate
	(*DeleteKliringTaskSingleTemplateRequest)(nil),     // 4: kliring.service.v1.DeleteKliringTaskSingleTemplateRequest
	(*GetKliringTaskRequest)(nil),                      // 5: kliring.service.v1.GetKliringTaskRequest
	(*GetKliringTaskDetailRequest)(nil),                // 6: kliring.service.v1.GetKliringTaskDetailRequest
	(*GetKliringTaskFileRequest)(nil),                  // 7: kliring.service.v1.GetKliringTaskFileRequest
	(*UpdateKliringTaskRequest)(nil),                   // 8: kliring.service.v1.UpdateKliringTaskRequest
	(*CreateKliringTaskSingleRequest)(nil),             // 9: kliring.service.v1.CreateKliringTaskSingleRequest
	(*CreateKliringTaskMultipleRequest)(nil),           // 10: kliring.service.v1.CreateKliringTaskMultipleRequest
	(*ExecKliringRequest)(nil),                         // 11: kliring.service.v1.ExecKliringRequest
	(*CreateKliringTransactionRequest)(nil),            // 12: kliring.service.v1.CreateKliringTransactionRequest
	(*CancelKliringTransactionRequest)(nil),            // 13: kliring.service.v1.CancelKliringTransactionRequest
	(*HealthCheckResponse)(nil),                        // 14: kliring.service.v1.HealthCheckResponse
	(*GetKliringTaskSingleTemplateResponse)(nil),       // 15: kliring.service.v1.GetKliringTaskSingleTemplateResponse
	(*GetKliringTaskSingleTemplateDetailResponse)(nil), // 16: kliring.service.v1.GetKliringTaskSingleTemplateDetailResponse
	(*CreateKliringTaskSingleTemplateResponse)(nil),    // 17: kliring.service.v1.CreateKliringTaskSingleTemplateResponse
	(*DeleteKliringTaskSingleTemplateResponse)(nil),    // 18: kliring.service.v1.DeleteKliringTaskSingleTemplateResponse
	(*GetKliringTaskResponse)(nil),                     // 19: kliring.service.v1.GetKliringTaskResponse
	(*GetKliringTaskDetailResponse)(nil),               // 20: kliring.service.v1.GetKliringTaskDetailResponse
	(*httpbody.HttpBody)(nil),                          // 21: google.api.HttpBody
	(*UpdateKliringTaskResponse)(nil),                  // 22: kliring.service.v1.UpdateKliringTaskResponse
	(*CreateKliringTaskSingleResponse)(nil),            // 23: kliring.service.v1.CreateKliringTaskSingleResponse
	(*CreateKliringTaskMultipleResponse)(nil),          // 24: kliring.service.v1.CreateKliringTaskMultipleResponse
	(*ExecKliringResponse)(nil),                        // 25: kliring.service.v1.ExecKliringResponse
	(*CreateKliringTransactionResponse)(nil),           // 26: kliring.service.v1.CreateKliringTransactionResponse
	(*CancelKliringTransactionResponse)(nil),           // 27: kliring.service.v1.CancelKliringTransactionResponse
}
var file_kliring_api_proto_depIdxs = []int32{
	0,  // 0: kliring.service.v1.ApiService.HealthCheck:input_type -> kliring.service.v1.HealthCheckRequest
	1,  // 1: kliring.service.v1.ApiService.GetKliringTaskSingleTemplate:input_type -> kliring.service.v1.GetKliringTaskSingleTemplateRequest
	2,  // 2: kliring.service.v1.ApiService.GetKliringTaskSingleTemplateDetail:input_type -> kliring.service.v1.GetKliringTaskSingleTemplateDetailRequest
	3,  // 3: kliring.service.v1.ApiService.CreateKliringTaskSingleTemplate:input_type -> kliring.service.v1.KliringSingleTemplate
	4,  // 4: kliring.service.v1.ApiService.DeleteKliringTaskSingleTemplate:input_type -> kliring.service.v1.DeleteKliringTaskSingleTemplateRequest
	5,  // 5: kliring.service.v1.ApiService.GetKliringTask:input_type -> kliring.service.v1.GetKliringTaskRequest
	6,  // 6: kliring.service.v1.ApiService.GetKliringTaskDetail:input_type -> kliring.service.v1.GetKliringTaskDetailRequest
	7,  // 7: kliring.service.v1.ApiService.GetKliringTaskFile:input_type -> kliring.service.v1.GetKliringTaskFileRequest
	8,  // 8: kliring.service.v1.ApiService.UpdateKliringTask:input_type -> kliring.service.v1.UpdateKliringTaskRequest
	9,  // 9: kliring.service.v1.ApiService.CreateKliringTaskSingle:input_type -> kliring.service.v1.CreateKliringTaskSingleRequest
	10, // 10: kliring.service.v1.ApiService.CreateKliringTaskMultiple:input_type -> kliring.service.v1.CreateKliringTaskMultipleRequest
	11, // 11: kliring.service.v1.ApiService.ExecKliring:input_type -> kliring.service.v1.ExecKliringRequest
	12, // 12: kliring.service.v1.ApiService.CreateKliringTransaction:input_type -> kliring.service.v1.CreateKliringTransactionRequest
	13, // 13: kliring.service.v1.ApiService.CancelKliringTransaction:input_type -> kliring.service.v1.CancelKliringTransactionRequest
	14, // 14: kliring.service.v1.ApiService.HealthCheck:output_type -> kliring.service.v1.HealthCheckResponse
	15, // 15: kliring.service.v1.ApiService.GetKliringTaskSingleTemplate:output_type -> kliring.service.v1.GetKliringTaskSingleTemplateResponse
	16, // 16: kliring.service.v1.ApiService.GetKliringTaskSingleTemplateDetail:output_type -> kliring.service.v1.GetKliringTaskSingleTemplateDetailResponse
	17, // 17: kliring.service.v1.ApiService.CreateKliringTaskSingleTemplate:output_type -> kliring.service.v1.CreateKliringTaskSingleTemplateResponse
	18, // 18: kliring.service.v1.ApiService.DeleteKliringTaskSingleTemplate:output_type -> kliring.service.v1.DeleteKliringTaskSingleTemplateResponse
	19, // 19: kliring.service.v1.ApiService.GetKliringTask:output_type -> kliring.service.v1.GetKliringTaskResponse
	20, // 20: kliring.service.v1.ApiService.GetKliringTaskDetail:output_type -> kliring.service.v1.GetKliringTaskDetailResponse
	21, // 21: kliring.service.v1.ApiService.GetKliringTaskFile:output_type -> google.api.HttpBody
	22, // 22: kliring.service.v1.ApiService.UpdateKliringTask:output_type -> kliring.service.v1.UpdateKliringTaskResponse
	23, // 23: kliring.service.v1.ApiService.CreateKliringTaskSingle:output_type -> kliring.service.v1.CreateKliringTaskSingleResponse
	24, // 24: kliring.service.v1.ApiService.CreateKliringTaskMultiple:output_type -> kliring.service.v1.CreateKliringTaskMultipleResponse
	25, // 25: kliring.service.v1.ApiService.ExecKliring:output_type -> kliring.service.v1.ExecKliringResponse
	26, // 26: kliring.service.v1.ApiService.CreateKliringTransaction:output_type -> kliring.service.v1.CreateKliringTransactionResponse
	27, // 27: kliring.service.v1.ApiService.CancelKliringTransaction:output_type -> kliring.service.v1.CancelKliringTransactionResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_kliring_api_proto_init() }
func file_kliring_api_proto_init() {
	if File_kliring_api_proto != nil {
		return
	}
	file_kliring_gorm_db_proto_init()
	file_kliring_core_payload_proto_init()
	file_kliring_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kliring_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kliring_api_proto_goTypes,
		DependencyIndexes: file_kliring_api_proto_depIdxs,
	}.Build()
	File_kliring_api_proto = out.File
	file_kliring_api_proto_rawDesc = nil
	file_kliring_api_proto_goTypes = nil
	file_kliring_api_proto_depIdxs = nil
}
