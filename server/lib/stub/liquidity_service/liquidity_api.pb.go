// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.16.1
// source: liquidity_api.proto

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

var File_liquidity_api_proto protoreflect.FileDescriptor

var file_liquidity_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x23, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0xe0, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x2e, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x3f, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x74,
	0x61, 0x73, 0x6b, 0x1a, 0x24, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0xe4, 0x01, 0x0a, 0x19, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x36, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61,
	0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x7b, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x7d, 0x92, 0x41, 0x41, 0x12,
	0x17, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x24, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x62, 0x00,
	0x12, 0x85, 0x02, 0x0a, 0x13, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x30, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x88, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f,
	0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x4d, 0x12, 0x1e, 0x47, 0x65, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x2b, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61,
	0x74, 0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x9b, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x30, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5a, 0x22, 0x20,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63,
	0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x3a, 0x01, 0x2a, 0x5a, 0x33, 0x22, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3b, 0x12, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0xe8, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x30,
	0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x30, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68,
	0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x31,
	0x12, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x1a, 0x1d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x12, 0xca, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x30, 0x12, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x6d, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x62,
	0x79, 0x92, 0x41, 0x3a, 0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d,
	0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x20, 0x54, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0xcc,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12,
	0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f,
	0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x2d, 0x62, 0x79,
	0x92, 0x41, 0x3a, 0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x79,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x20, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0xd0, 0x01,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x12, 0x2c, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d,
	0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x33, 0x12, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x1a,
	0x1d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x62, 0x00,
	0x12, 0x6e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x2c, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0xcd, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x25, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f,
	0x6c, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x49, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x1a, 0x2c, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x62, 0x00,
	0x12, 0xc5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x42, 0x41, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2a, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x42, 0x41,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x62, 0x61, 0x92, 0x41, 0x3f, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x54, 0x42, 0x41, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x27, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x54, 0x42, 0x41,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x35, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xf7, 0x01, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x69, 0x71,
	0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2f, 0x72, 0x75, 0x6e, 0x2d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x2d, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x92, 0x41, 0x57, 0x12, 0x1e, 0x52, 0x75, 0x6e, 0x20, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73,
	0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x64, 0x61, 0x79, 0x1a, 0x33, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x72, 0x75, 0x6e, 0x20, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x64, 0x61, 0x79, 0x62, 0x00,
	0x12, 0xf1, 0x01, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x2d, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75,
	0x6e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x26,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63,
	0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x2d,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x92, 0x41, 0x4b, 0x12, 0x18, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x20, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x1a, 0x2d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73,
	0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x20, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x62, 0x00, 0x12, 0xd0, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x65, 0x92,
	0x41, 0x35, 0x12, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x44, 0x61, 0x74,
	0x65, 0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75,
	0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x44, 0x61, 0x74, 0x65, 0x62, 0x00, 0x12, 0xf4, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x12, 0x2e, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68,
	0x2d, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x92, 0x41, 0x49, 0x12, 0x1c, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x1a, 0x29, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0xff,
	0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x9d, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70,
	0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x68, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0f, 0x53, 0x65, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x4a, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x27,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0xe2, 0x80, 0x99, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0xe2, 0x80, 0x99, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0xe2,
	0x80, 0x99, 0x2c, 0x20, 0xe2, 0x80, 0x98, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0xe2, 0x80, 0x99,
	0x12, 0x8f, 0x02, 0x0a, 0x19, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x36,
	0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x80, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x6f,
	0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x64, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x3e, 0x12, 0x19, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a,
	0x21, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x12, 0x8f, 0x02, 0x0a, 0x19, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x36, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x80, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2d,
	0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3e, 0x12, 0x19, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x20, 0x74, 0x61, 0x73,
	0x6b, 0x1a, 0x21, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x20, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x12, 0x70, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x43, 0x61, 0x73, 0x68,
	0x50, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x43, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x43, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41,
	0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08,
	0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_liquidity_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                             // 0: liquidity.service.v1.Empty
	(*ListTaskLiquidityRequest)(nil),          // 1: liquidity.service.v1.ListTaskLiquidityRequest
	(*DownloadListTaskLiquidityRequest)(nil),  // 2: liquidity.service.v1.DownloadListTaskLiquidityRequest
	(*DetailLiquidityTaskRequest)(nil),        // 3: liquidity.service.v1.DetailLiquidityTaskRequest
	(*CreateTaskLiquidityRequest)(nil),        // 4: liquidity.service.v1.CreateTaskLiquidityRequest
	(*CreateLiquidityRequest)(nil),            // 5: liquidity.service.v1.CreateLiquidityRequest
	(*DeleteLiquidityRequest)(nil),            // 6: liquidity.service.v1.DeleteLiquidityRequest
	(*ListDataRequest)(nil),                   // 7: liquidity.service.v1.ListDataRequest
	(*CreateLiquiditySchedulesRequest)(nil),   // 8: liquidity.service.v1.CreateLiquiditySchedulesRequest
	(*RunDailyScheduleRequest)(nil),           // 9: liquidity.service.v1.RunDailyScheduleRequest
	(*RunLiquidityTaskRequest)(nil),           // 10: liquidity.service.v1.RunLiquidityTaskRequest
	(*ValidateDateRequest)(nil),               // 11: liquidity.service.v1.ValidateDateRequest
	(*TaskActionRequest)(nil),                 // 12: liquidity.service.v1.TaskActionRequest
	(*DeactivateLiquiditySchemeRequest)(nil),  // 13: liquidity.service.v1.DeactivateLiquiditySchemeRequest
	(*ExecCashPoolingRequest)(nil),            // 14: liquidity.service.v1.ExecCashPoolingRequest
	(*HealthCheckResponse)(nil),               // 15: liquidity.service.v1.HealthCheckResponse
	(*ListLiquidityTaskResponse)(nil),         // 16: liquidity.service.v1.ListLiquidityTaskResponse
	(*httpbody.HttpBody)(nil),                 // 17: google.api.HttpBody
	(*DetailLiquidityTaskResponse)(nil),       // 18: liquidity.service.v1.DetailLiquidityTaskResponse
	(*CreateTaskLiquidityResponse)(nil),       // 19: liquidity.service.v1.CreateTaskLiquidityResponse
	(*ArrayString)(nil),                       // 20: liquidity.service.v1.ArrayString
	(*DeleteLiquidityResponse)(nil),           // 21: liquidity.service.v1.DeleteLiquidityResponse
	(*ListDataResponse)(nil),                  // 22: liquidity.service.v1.ListDataResponse
	(*ListTBAValueResponse)(nil),              // 23: liquidity.service.v1.ListTBAValueResponse
	(*CreateLiquiditySchedulesResponse)(nil),  // 24: liquidity.service.v1.CreateLiquiditySchedulesResponse
	(*RunDailyScheduleResponse)(nil),          // 25: liquidity.service.v1.RunDailyScheduleResponse
	(*RunLiquidityTaskResponse)(nil),          // 26: liquidity.service.v1.RunLiquidityTaskResponse
	(*ValidateDateResponse)(nil),              // 27: liquidity.service.v1.ValidateDateResponse
	(*TaskActionResponse)(nil),                // 28: liquidity.service.v1.TaskActionResponse
	(*DeactivateLiquiditySchemeResponse)(nil), // 29: liquidity.service.v1.DeactivateLiquiditySchemeResponse
	(*ExecCashPoolingResponse)(nil),           // 30: liquidity.service.v1.ExecCashPoolingResponse
}
var file_liquidity_api_proto_depIdxs = []int32{
	0,  // 0: liquidity.service.v1.ApiService.HealthCheck:input_type -> liquidity.service.v1.Empty
	1,  // 1: liquidity.service.v1.ApiService.ListLiquidityTask:input_type -> liquidity.service.v1.ListTaskLiquidityRequest
	2,  // 2: liquidity.service.v1.ApiService.DownloadListLiquidityTask:input_type -> liquidity.service.v1.DownloadListTaskLiquidityRequest
	3,  // 3: liquidity.service.v1.ApiService.DetailLiquidityTask:input_type -> liquidity.service.v1.DetailLiquidityTaskRequest
	4,  // 4: liquidity.service.v1.ApiService.CreateLiquidityTask:input_type -> liquidity.service.v1.CreateTaskLiquidityRequest
	3,  // 5: liquidity.service.v1.ApiService.DeleteLiquidityTask:input_type -> liquidity.service.v1.DetailLiquidityTaskRequest
	0,  // 6: liquidity.service.v1.ApiService.GetMyTasksCreatedBy:input_type -> liquidity.service.v1.Empty
	0,  // 7: liquidity.service.v1.ApiService.GetMyTasksApprovedBy:input_type -> liquidity.service.v1.Empty
	5,  // 8: liquidity.service.v1.ApiService.CreateLiquidity:input_type -> liquidity.service.v1.CreateLiquidityRequest
	6,  // 9: liquidity.service.v1.ApiService.DeleteLiquidity:input_type -> liquidity.service.v1.DeleteLiquidityRequest
	7,  // 10: liquidity.service.v1.ApiService.GetListData:input_type -> liquidity.service.v1.ListDataRequest
	0,  // 11: liquidity.service.v1.ApiService.GetListTBAValue:input_type -> liquidity.service.v1.Empty
	8,  // 12: liquidity.service.v1.ApiService.CreateLiquiditySchedules:input_type -> liquidity.service.v1.CreateLiquiditySchedulesRequest
	9,  // 13: liquidity.service.v1.ApiService.RunDailySchedule:input_type -> liquidity.service.v1.RunDailyScheduleRequest
	10, // 14: liquidity.service.v1.ApiService.RunLiquidityScheme:input_type -> liquidity.service.v1.RunLiquidityTaskRequest
	11, // 15: liquidity.service.v1.ApiService.ValidateDate:input_type -> liquidity.service.v1.ValidateDateRequest
	1,  // 16: liquidity.service.v1.ApiService.ListLiquidityAuthorize:input_type -> liquidity.service.v1.ListTaskLiquidityRequest
	12, // 17: liquidity.service.v1.ApiService.TaskAction:input_type -> liquidity.service.v1.TaskActionRequest
	13, // 18: liquidity.service.v1.ApiService.DeactivateLiquidityScheme:input_type -> liquidity.service.v1.DeactivateLiquiditySchemeRequest
	13, // 19: liquidity.service.v1.ApiService.ReactivateLiquidityScheme:input_type -> liquidity.service.v1.DeactivateLiquiditySchemeRequest
	14, // 20: liquidity.service.v1.ApiService.ExecCashPooling:input_type -> liquidity.service.v1.ExecCashPoolingRequest
	15, // 21: liquidity.service.v1.ApiService.HealthCheck:output_type -> liquidity.service.v1.HealthCheckResponse
	16, // 22: liquidity.service.v1.ApiService.ListLiquidityTask:output_type -> liquidity.service.v1.ListLiquidityTaskResponse
	17, // 23: liquidity.service.v1.ApiService.DownloadListLiquidityTask:output_type -> google.api.HttpBody
	18, // 24: liquidity.service.v1.ApiService.DetailLiquidityTask:output_type -> liquidity.service.v1.DetailLiquidityTaskResponse
	19, // 25: liquidity.service.v1.ApiService.CreateLiquidityTask:output_type -> liquidity.service.v1.CreateTaskLiquidityResponse
	18, // 26: liquidity.service.v1.ApiService.DeleteLiquidityTask:output_type -> liquidity.service.v1.DetailLiquidityTaskResponse
	20, // 27: liquidity.service.v1.ApiService.GetMyTasksCreatedBy:output_type -> liquidity.service.v1.ArrayString
	20, // 28: liquidity.service.v1.ApiService.GetMyTasksApprovedBy:output_type -> liquidity.service.v1.ArrayString
	19, // 29: liquidity.service.v1.ApiService.CreateLiquidity:output_type -> liquidity.service.v1.CreateTaskLiquidityResponse
	21, // 30: liquidity.service.v1.ApiService.DeleteLiquidity:output_type -> liquidity.service.v1.DeleteLiquidityResponse
	22, // 31: liquidity.service.v1.ApiService.GetListData:output_type -> liquidity.service.v1.ListDataResponse
	23, // 32: liquidity.service.v1.ApiService.GetListTBAValue:output_type -> liquidity.service.v1.ListTBAValueResponse
	24, // 33: liquidity.service.v1.ApiService.CreateLiquiditySchedules:output_type -> liquidity.service.v1.CreateLiquiditySchedulesResponse
	25, // 34: liquidity.service.v1.ApiService.RunDailySchedule:output_type -> liquidity.service.v1.RunDailyScheduleResponse
	26, // 35: liquidity.service.v1.ApiService.RunLiquidityScheme:output_type -> liquidity.service.v1.RunLiquidityTaskResponse
	27, // 36: liquidity.service.v1.ApiService.ValidateDate:output_type -> liquidity.service.v1.ValidateDateResponse
	16, // 37: liquidity.service.v1.ApiService.ListLiquidityAuthorize:output_type -> liquidity.service.v1.ListLiquidityTaskResponse
	28, // 38: liquidity.service.v1.ApiService.TaskAction:output_type -> liquidity.service.v1.TaskActionResponse
	29, // 39: liquidity.service.v1.ApiService.DeactivateLiquidityScheme:output_type -> liquidity.service.v1.DeactivateLiquiditySchemeResponse
	29, // 40: liquidity.service.v1.ApiService.ReactivateLiquidityScheme:output_type -> liquidity.service.v1.DeactivateLiquiditySchemeResponse
	30, // 41: liquidity.service.v1.ApiService.ExecCashPooling:output_type -> liquidity.service.v1.ExecCashPoolingResponse
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_liquidity_api_proto_init() }
func file_liquidity_api_proto_init() {
	if File_liquidity_api_proto != nil {
		return
	}
	file_liquidity_payload_proto_init()
	file_liquidity_core_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_liquidity_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liquidity_api_proto_goTypes,
		DependencyIndexes: file_liquidity_api_proto_depIdxs,
	}.Build()
	File_liquidity_api_proto = out.File
	file_liquidity_api_proto_rawDesc = nil
	file_liquidity_api_proto_goTypes = nil
	file_liquidity_api_proto_depIdxs = nil
}
