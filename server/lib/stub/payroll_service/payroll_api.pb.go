// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payroll_api.proto

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

var File_payroll_api_proto protoreflect.FileDescriptor

var file_payroll_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x61,
	0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x13, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb0, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0xc3, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x70,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x7d, 0x92, 0x41, 0x36, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x0e,
	0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x20, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x1b,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x20, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0xe6, 0x01, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0x7f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x7d, 0x92, 0x41, 0x44,
	0x0a, 0x07, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x15, 0x50, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x20, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x50,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0xad, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x28, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x07, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x1a, 0x14,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x12, 0xcf, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x2e,
	0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x2f, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c,
	0x12, 0x07, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x1a, 0x1b, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x20,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0xe8, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x2c, 0x2e, 0x70,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72,
	0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x37, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x36, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x1a, 0x1b, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x12, 0xf2, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x29, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x72,
	0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x88, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x73, 0x65, 0x74, 0x2f, 0x7b, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x5c, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x72,
	0x6f, 0x6c, 0x6c, 0x12, 0x0f, 0x53, 0x65, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x40, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x27,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x27, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x27, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x27, 0x2c, 0x20, 0x27, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x27, 0x12, 0x87, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x87, 0x01, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e,
	0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61,
	0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x13, 0x45, 0x78,
	0x65, 0x63, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x2e, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x50, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x50, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8a, 0x01, 0x0a, 0x19, 0x45, 0x78, 0x65, 0x63, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x8d, 0x01, 0x0a, 0x1a, 0x45, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x4a, 0x6f, 0x62,
	0x12, 0x35, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79,
	0x72, 0x6f, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x75, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x4d, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x12, 0x2d, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x4d, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x4d, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63,
	0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x12,
	0x2e, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x75, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f,
	0x6c, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x61, 0x79, 0x72, 0x6f, 0x6c,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62,
	0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_payroll_api_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),                 // 0: payroll.service.v1.HealthCheckRequest
	(*GetTaskPayrollFileRequest)(nil),          // 1: payroll.service.v1.GetTaskPayrollFileRequest
	(*GetTaskPayrollDetailFileRequest)(nil),    // 2: payroll.service.v1.GetTaskPayrollDetailFileRequest
	(*GetTaskPayrollRequest)(nil),              // 3: payroll.service.v1.GetTaskPayrollRequest
	(*GetTaskPayrollDetailRequest)(nil),        // 4: payroll.service.v1.GetTaskPayrollDetailRequest
	(*CreateTaskPayrollRequest)(nil),           // 5: payroll.service.v1.CreateTaskPayrollRequest
	(*SetTaskPayrollRequest)(nil),              // 6: payroll.service.v1.SetTaskPayrollRequest
	(*CreatePayrollTransactionRequest)(nil),    // 7: payroll.service.v1.CreatePayrollTransactionRequest
	(*CancelPayrollTransactionRequest)(nil),    // 8: payroll.service.v1.CancelPayrollTransactionRequest
	(*ExecPayrollTransferRequest)(nil),         // 9: payroll.service.v1.ExecPayrollTransferRequest
	(*ExecFailedPayrollTransferRequest)(nil),   // 10: payroll.service.v1.ExecFailedPayrollTransferRequest
	(*ExecFileCheckingPayrollJobRequest)(nil),  // 11: payroll.service.v1.ExecFileCheckingPayrollJobRequest
	(*ExecMassInquiryJobRequest)(nil),          // 12: payroll.service.v1.ExecMassInquiryJobRequest
	(*ExecMassTransferJobRequest)(nil),         // 13: payroll.service.v1.ExecMassTransferJobRequest
	(*CreateMassTransferRequest)(nil),          // 14: payroll.service.v1.CreateMassTransferRequest
	(*HealthCheckResponse)(nil),                // 15: payroll.service.v1.HealthCheckResponse
	(*httpbody.HttpBody)(nil),                  // 16: google.api.HttpBody
	(*GetTaskPayrollResponse)(nil),             // 17: payroll.service.v1.GetTaskPayrollResponse
	(*GetTaskPayrollDetailResponse)(nil),       // 18: payroll.service.v1.GetTaskPayrollDetailResponse
	(*CreateTaskPayrollResponse)(nil),          // 19: payroll.service.v1.CreateTaskPayrollResponse
	(*SetTaskPayrollResponse)(nil),             // 20: payroll.service.v1.SetTaskPayrollResponse
	(*CreatePayrollTransactionResponse)(nil),   // 21: payroll.service.v1.CreatePayrollTransactionResponse
	(*CancelPayrollTransactionResponse)(nil),   // 22: payroll.service.v1.CancelPayrollTransactionResponse
	(*ExecPayrollTransferResponse)(nil),        // 23: payroll.service.v1.ExecPayrollTransferResponse
	(*ExecFailedPayrollTransferResponse)(nil),  // 24: payroll.service.v1.ExecFailedPayrollTransferResponse
	(*ExecFileCheckingPayrollJobResponse)(nil), // 25: payroll.service.v1.ExecFileCheckingPayrollJobResponse
	(*ExecMassInquiryJobResponse)(nil),         // 26: payroll.service.v1.ExecMassInquiryJobResponse
	(*ExecMassTransferJobResponse)(nil),        // 27: payroll.service.v1.ExecMassTransferJobResponse
	(*CreateMassTransferResponse)(nil),         // 28: payroll.service.v1.CreateMassTransferResponse
}
var file_payroll_api_proto_depIdxs = []int32{
	0,  // 0: payroll.service.v1.ApiService.HealthCheck:input_type -> payroll.service.v1.HealthCheckRequest
	1,  // 1: payroll.service.v1.ApiService.GetTaskPayrollFile:input_type -> payroll.service.v1.GetTaskPayrollFileRequest
	2,  // 2: payroll.service.v1.ApiService.GetTaskPayrollDetailFile:input_type -> payroll.service.v1.GetTaskPayrollDetailFileRequest
	3,  // 3: payroll.service.v1.ApiService.GetTaskPayroll:input_type -> payroll.service.v1.GetTaskPayrollRequest
	4,  // 4: payroll.service.v1.ApiService.GetTaskPayrollDetail:input_type -> payroll.service.v1.GetTaskPayrollDetailRequest
	5,  // 5: payroll.service.v1.ApiService.CreateTaskPayroll:input_type -> payroll.service.v1.CreateTaskPayrollRequest
	6,  // 6: payroll.service.v1.ApiService.SetTaskPayroll:input_type -> payroll.service.v1.SetTaskPayrollRequest
	7,  // 7: payroll.service.v1.ApiService.CreatePayrollTransaction:input_type -> payroll.service.v1.CreatePayrollTransactionRequest
	8,  // 8: payroll.service.v1.ApiService.CancelPayrollTransaction:input_type -> payroll.service.v1.CancelPayrollTransactionRequest
	9,  // 9: payroll.service.v1.ApiService.ExecPayrollTransfer:input_type -> payroll.service.v1.ExecPayrollTransferRequest
	10, // 10: payroll.service.v1.ApiService.ExecFailedPayrollTransfer:input_type -> payroll.service.v1.ExecFailedPayrollTransferRequest
	11, // 11: payroll.service.v1.ApiService.ExecFileCheckingPayrollJob:input_type -> payroll.service.v1.ExecFileCheckingPayrollJobRequest
	12, // 12: payroll.service.v1.ApiService.ExecMassInquiryJob:input_type -> payroll.service.v1.ExecMassInquiryJobRequest
	13, // 13: payroll.service.v1.ApiService.ExecMassTransferJob:input_type -> payroll.service.v1.ExecMassTransferJobRequest
	14, // 14: payroll.service.v1.ApiService.CreateMassTransfer:input_type -> payroll.service.v1.CreateMassTransferRequest
	15, // 15: payroll.service.v1.ApiService.HealthCheck:output_type -> payroll.service.v1.HealthCheckResponse
	16, // 16: payroll.service.v1.ApiService.GetTaskPayrollFile:output_type -> google.api.HttpBody
	16, // 17: payroll.service.v1.ApiService.GetTaskPayrollDetailFile:output_type -> google.api.HttpBody
	17, // 18: payroll.service.v1.ApiService.GetTaskPayroll:output_type -> payroll.service.v1.GetTaskPayrollResponse
	18, // 19: payroll.service.v1.ApiService.GetTaskPayrollDetail:output_type -> payroll.service.v1.GetTaskPayrollDetailResponse
	19, // 20: payroll.service.v1.ApiService.CreateTaskPayroll:output_type -> payroll.service.v1.CreateTaskPayrollResponse
	20, // 21: payroll.service.v1.ApiService.SetTaskPayroll:output_type -> payroll.service.v1.SetTaskPayrollResponse
	21, // 22: payroll.service.v1.ApiService.CreatePayrollTransaction:output_type -> payroll.service.v1.CreatePayrollTransactionResponse
	22, // 23: payroll.service.v1.ApiService.CancelPayrollTransaction:output_type -> payroll.service.v1.CancelPayrollTransactionResponse
	23, // 24: payroll.service.v1.ApiService.ExecPayrollTransfer:output_type -> payroll.service.v1.ExecPayrollTransferResponse
	24, // 25: payroll.service.v1.ApiService.ExecFailedPayrollTransfer:output_type -> payroll.service.v1.ExecFailedPayrollTransferResponse
	25, // 26: payroll.service.v1.ApiService.ExecFileCheckingPayrollJob:output_type -> payroll.service.v1.ExecFileCheckingPayrollJobResponse
	26, // 27: payroll.service.v1.ApiService.ExecMassInquiryJob:output_type -> payroll.service.v1.ExecMassInquiryJobResponse
	27, // 28: payroll.service.v1.ApiService.ExecMassTransferJob:output_type -> payroll.service.v1.ExecMassTransferJobResponse
	28, // 29: payroll.service.v1.ApiService.CreateMassTransfer:output_type -> payroll.service.v1.CreateMassTransferResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_payroll_api_proto_init() }
func file_payroll_api_proto_init() {
	if File_payroll_api_proto != nil {
		return
	}
	file_payroll_core_payload_proto_init()
	file_payroll_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payroll_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payroll_api_proto_goTypes,
		DependencyIndexes: file_payroll_api_proto_depIdxs,
	}.Build()
	File_payroll_api_proto = out.File
	file_payroll_api_proto_rawDesc = nil
	file_payroll_api_proto_goTypes = nil
	file_payroll_api_proto_depIdxs = nil
}
