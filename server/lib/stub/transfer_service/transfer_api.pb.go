// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: transfer_api.proto

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

var File_transfer_api_proto protoreflect.FileDescriptor

var file_transfer_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdc, 0x17, 0x0a, 0x0a, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb3, 0x01, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12,
	0xaf, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x92, 0x41,
	0x30, 0x0a, 0x07, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x09, 0x50, 0x61, 0x69, 0x72,
	0x20, 0x52, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x70, 0x61, 0x69, 0x72, 0x20, 0x72, 0x61, 0x74,
	0x65, 0x12, 0xc5, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x92, 0x41, 0x42, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x1c, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x97, 0x02, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0xac, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x12, 0x38, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x7d, 0x92, 0x41, 0x69, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x1f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x2c, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x8e, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x31, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x92, 0x41, 0x60, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x1d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x25, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0xab, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x37, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x98, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x12, 0x2b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x62,
	0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x18, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x2c, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0xe2, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12,
	0x34, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd8, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x59, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x30, 0x1a, 0x2b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x76, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x29, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x4f, 0x72, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x2f, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x9c, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x12, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x66,
	0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x1f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x27, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0xb9, 0x02, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xa0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x68, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x1a, 0x2e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x20, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0xf2, 0x02, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x12, 0x36, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xe2, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5d, 0x22, 0x24, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x32, 0x1a, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7c, 0x0a, 0x1a, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x2b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x4f, 0x72, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x1a, 0x31, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x8f, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaf, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27,
	0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x2f, 0x64, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7f, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x42, 0x75,
	0x6c, 0x6b, 0x12, 0x2b, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x26, 0x20, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x42, 0x75, 0x6c, 0x6b, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x38, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x20, 0x26, 0x20, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x42, 0x75, 0x6c, 0x6b, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62,
	0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_transfer_api_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),                    // 0: transfer.service.v1.HealthCheckRequest
	(*GetPairRateRequest)(nil),                    // 1: transfer.service.v1.GetPairRateRequest
	(*CreateTransferRequest)(nil),                 // 2: transfer.service.v1.CreateTransferRequest
	(*GetTaskInternalSingleFileRequest)(nil),      // 3: transfer.service.v1.GetTaskInternalSingleFileRequest
	(*GetTaskInternalSingleRequest)(nil),          // 4: transfer.service.v1.GetTaskInternalSingleRequest
	(*GetTaskInternalSingleDetailRequest)(nil),    // 5: transfer.service.v1.GetTaskInternalSingleDetailRequest
	(*CreateTaskInternalSingleRequest)(nil),       // 6: transfer.service.v1.CreateTaskInternalSingleRequest
	(*GetTaskInternalMultipleRequest)(nil),        // 7: transfer.service.v1.GetTaskInternalMultipleRequest
	(*GetTaskInternalMultipleDetailRequest)(nil),  // 8: transfer.service.v1.GetTaskInternalMultipleDetailRequest
	(*CreateTaskInternalMultipleRequest)(nil),     // 9: transfer.service.v1.CreateTaskInternalMultipleRequest
	(*DecodeFileRequest)(nil),                     // 10: transfer.service.v1.DecodeFileRequest
	(*HealthCheckResponse)(nil),                   // 11: transfer.service.v1.HealthCheckResponse
	(*GetPairRateResponse)(nil),                   // 12: transfer.service.v1.GetPairRateResponse
	(*CreateTransferResponse)(nil),                // 13: transfer.service.v1.CreateTransferResponse
	(*httpbody.HttpBody)(nil),                     // 14: google.api.HttpBody
	(*GetTaskInternalSingleResponse)(nil),         // 15: transfer.service.v1.GetTaskInternalSingleResponse
	(*GetTaskInternalSingleDetailResponse)(nil),   // 16: transfer.service.v1.GetTaskInternalSingleDetailResponse
	(*CreateTaskInternalSingleResponse)(nil),      // 17: transfer.service.v1.CreateTaskInternalSingleResponse
	(*GetTaskInternalMultipleResponse)(nil),       // 18: transfer.service.v1.GetTaskInternalMultipleResponse
	(*GetTaskInternalMultipleDetailResponse)(nil), // 19: transfer.service.v1.GetTaskInternalMultipleDetailResponse
	(*CreateTaskInternalMultipleResponse)(nil),    // 20: transfer.service.v1.CreateTaskInternalMultipleResponse
	(*DecodeFileResponse)(nil),                    // 21: transfer.service.v1.DecodeFileResponse
}
var file_transfer_api_proto_depIdxs = []int32{
	0,  // 0: transfer.service.v1.ApiService.HealthCheck:input_type -> transfer.service.v1.HealthCheckRequest
	1,  // 1: transfer.service.v1.ApiService.GetPairRate:input_type -> transfer.service.v1.GetPairRateRequest
	2,  // 2: transfer.service.v1.ApiService.CreateTransfer:input_type -> transfer.service.v1.CreateTransferRequest
	3,  // 3: transfer.service.v1.ApiService.GetTaskInternalSingleFile:input_type -> transfer.service.v1.GetTaskInternalSingleFileRequest
	4,  // 4: transfer.service.v1.ApiService.GetTaskInternalSingle:input_type -> transfer.service.v1.GetTaskInternalSingleRequest
	5,  // 5: transfer.service.v1.ApiService.GetTaskInternalSingleDetail:input_type -> transfer.service.v1.GetTaskInternalSingleDetailRequest
	6,  // 6: transfer.service.v1.ApiService.CreateTaskInternalSingle:input_type -> transfer.service.v1.CreateTaskInternalSingleRequest
	7,  // 7: transfer.service.v1.ApiService.GetTaskInternalMultiple:input_type -> transfer.service.v1.GetTaskInternalMultipleRequest
	8,  // 8: transfer.service.v1.ApiService.GetTaskInternalMultipleDetail:input_type -> transfer.service.v1.GetTaskInternalMultipleDetailRequest
	9,  // 9: transfer.service.v1.ApiService.CreateTaskInternalMultiple:input_type -> transfer.service.v1.CreateTaskInternalMultipleRequest
	10, // 10: transfer.service.v1.ApiService.DecodeFile:input_type -> transfer.service.v1.DecodeFileRequest
	11, // 11: transfer.service.v1.ApiService.HealthCheck:output_type -> transfer.service.v1.HealthCheckResponse
	12, // 12: transfer.service.v1.ApiService.GetPairRate:output_type -> transfer.service.v1.GetPairRateResponse
	13, // 13: transfer.service.v1.ApiService.CreateTransfer:output_type -> transfer.service.v1.CreateTransferResponse
	14, // 14: transfer.service.v1.ApiService.GetTaskInternalSingleFile:output_type -> google.api.HttpBody
	15, // 15: transfer.service.v1.ApiService.GetTaskInternalSingle:output_type -> transfer.service.v1.GetTaskInternalSingleResponse
	16, // 16: transfer.service.v1.ApiService.GetTaskInternalSingleDetail:output_type -> transfer.service.v1.GetTaskInternalSingleDetailResponse
	17, // 17: transfer.service.v1.ApiService.CreateTaskInternalSingle:output_type -> transfer.service.v1.CreateTaskInternalSingleResponse
	18, // 18: transfer.service.v1.ApiService.GetTaskInternalMultiple:output_type -> transfer.service.v1.GetTaskInternalMultipleResponse
	19, // 19: transfer.service.v1.ApiService.GetTaskInternalMultipleDetail:output_type -> transfer.service.v1.GetTaskInternalMultipleDetailResponse
	20, // 20: transfer.service.v1.ApiService.CreateTaskInternalMultiple:output_type -> transfer.service.v1.CreateTaskInternalMultipleResponse
	21, // 21: transfer.service.v1.ApiService.DecodeFile:output_type -> transfer.service.v1.DecodeFileResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_api_proto_init() }
func file_transfer_api_proto_init() {
	if File_transfer_api_proto != nil {
		return
	}
	file_transfer_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_api_proto_goTypes,
		DependencyIndexes: file_transfer_api_proto_depIdxs,
	}.Build()
	File_transfer_api_proto = out.File
	file_transfer_api_proto_rawDesc = nil
	file_transfer_api_proto_goTypes = nil
	file_transfer_api_proto_depIdxs = nil
}
