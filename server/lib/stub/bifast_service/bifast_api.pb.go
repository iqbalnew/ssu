// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: bifast_api.proto

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

var File_bifast_api_proto protoreflect.FileDescriptor

var file_bifast_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x22, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x69,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92,
	0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x00, 0x12, 0xd1, 0x02, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3f,
	0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x40, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa4, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x92, 0x41, 0x7e, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x47,
	0x65, 0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x1a, 0x32, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x47, 0x65, 0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0xe8, 0x02, 0x0a, 0x2b, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x45, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x46, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x12, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x92, 0x41, 0x7e, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x47, 0x65, 0x74, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a,
	0x32, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x47, 0x65,
	0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0xf0, 0x02, 0x0a, 0x28, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x31, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x1a, 0x43, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcb, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x47, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x5a, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7b, 0x0a, 0x21, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x2c, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0xe3, 0x02, 0x0a, 0x28, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x42, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61,
	0x73, 0x74, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x92, 0x41,
	0x7b, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x28, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x2c,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0xe8, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x62, 0x69,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x66, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x66, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x4b, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12,
	0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0xf2, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x85, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x7d, 0x92, 0x41, 0x54, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x85, 0x02, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x37,
	0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x71, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x4d, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x25, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0xaf, 0x01, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x12,
	0x40, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x41, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x62,
	0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x62, 0x69, 0x66, 0x61,
	0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x21, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x62,
	0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x90, 0x02, 0x0a,
	0x17, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x62, 0x69,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x8d, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x69, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x66, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x53, 0x65,
	0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x2e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65,
	0x74, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x79, 0x0a, 0x14, 0x45, 0x78, 0x65, 0x63, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9d, 0x01, 0x0a, 0x20, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12,
	0x3a, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x62, 0x69,
	0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xeb, 0x02, 0x0a, 0x24, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48,
	0x74, 0x74, 0x70, 0x12, 0x3a, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc9, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x43, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61,
	0x73, 0x74, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01,
	0x2a, 0x5a, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7d, 0x0a, 0x18, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x29, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4f, 0x72,
	0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x1a, 0x36, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0xca, 0x02, 0x0a, 0x22, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12,
	0x3c, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61,
	0x73, 0x74, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x7f, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x12, 0x29, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x4f, 0x72, 0x20, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x36, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x84, 0x02, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9f, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x77, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x20,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12,
	0x25, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x20,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x34, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x66,
	0x75, 0x6e, 0x64, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x72, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bifast_api_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),                                  // 0: bifast.service.v1.HealthCheckRequest
	(*GetTaskExternalTransferSingleTemplateRequest)(nil),        // 1: bifast.service.v1.GetTaskExternalTransferSingleTemplateRequest
	(*GetTaskExternalTransferSingleTemplateDetailRequest)(nil),  // 2: bifast.service.v1.GetTaskExternalTransferSingleTemplateDetailRequest
	(*ExternalTransferSingleTemplate)(nil),                      // 3: bifast.service.v1.ExternalTransferSingleTemplate
	(*DeleteTaskExternalTransferSingleTemplateRequest)(nil),     // 4: bifast.service.v1.DeleteTaskExternalTransferSingleTemplateRequest
	(*GetTaskExternalTransferRequest)(nil),                      // 5: bifast.service.v1.GetTaskExternalTransferRequest
	(*GetTaskExternalTransferFileRequest)(nil),                  // 6: bifast.service.v1.GetTaskExternalTransferFileRequest
	(*GetTaskExternalTransferDetailRequest)(nil),                // 7: bifast.service.v1.GetTaskExternalTransferDetailRequest
	(*GetTaskExternalTransferDetailByReffNumRequest)(nil),       // 8: bifast.service.v1.GetTaskExternalTransferDetailByReffNumRequest
	(*CreateExternalTransferTransactionRequest)(nil),            // 9: bifast.service.v1.CreateExternalTransferTransactionRequest
	(*CancelExternalTransferTransactionRequest)(nil),            // 10: bifast.service.v1.CancelExternalTransferTransactionRequest
	(*SetTaskExternalTransferRequest)(nil),                      // 11: bifast.service.v1.SetTaskExternalTransferRequest
	(*ExecExternalTransferRequest)(nil),                         // 12: bifast.service.v1.ExecExternalTransferRequest
	(*CreateTaskExternalTransferSingleRequest)(nil),             // 13: bifast.service.v1.CreateTaskExternalTransferSingleRequest
	(*CreateTaskExternalTransferMultipleRequest)(nil),           // 14: bifast.service.v1.CreateTaskExternalTransferMultipleRequest
	(*ValidateProxyRequest)(nil),                                // 15: bifast.service.v1.ValidateProxyRequest
	(*HealthCheckResponse)(nil),                                 // 16: bifast.service.v1.HealthCheckResponse
	(*GetTaskExternalTransferSingleTemplateResponse)(nil),       // 17: bifast.service.v1.GetTaskExternalTransferSingleTemplateResponse
	(*GetTaskExternalTransferSingleTemplateDetailResponse)(nil), // 18: bifast.service.v1.GetTaskExternalTransferSingleTemplateDetailResponse
	(*CreateTaskExternalTransferSingleTemplateResponse)(nil),    // 19: bifast.service.v1.CreateTaskExternalTransferSingleTemplateResponse
	(*DeleteTaskExternalTransferSingleTemplateResponse)(nil),    // 20: bifast.service.v1.DeleteTaskExternalTransferSingleTemplateResponse
	(*GetTaskExternalTransferResponse)(nil),                     // 21: bifast.service.v1.GetTaskExternalTransferResponse
	(*httpbody.HttpBody)(nil),                                   // 22: google.api.HttpBody
	(*GetTaskExternalTransferDetailResponse)(nil),               // 23: bifast.service.v1.GetTaskExternalTransferDetailResponse
	(*GetTaskExternalTransferDetailByReffNumResponse)(nil),      // 24: bifast.service.v1.GetTaskExternalTransferDetailByReffNumResponse
	(*CreateExternalTransferTransactionResponse)(nil),           // 25: bifast.service.v1.CreateExternalTransferTransactionResponse
	(*CancelExternalTransferTransactionResponse)(nil),           // 26: bifast.service.v1.CancelExternalTransferTransactionResponse
	(*SetTaskExternalTransferResponse)(nil),                     // 27: bifast.service.v1.SetTaskExternalTransferResponse
	(*ExecExternalTransferResponse)(nil),                        // 28: bifast.service.v1.ExecExternalTransferResponse
	(*CreateTaskExternalTransferSingleResponse)(nil),            // 29: bifast.service.v1.CreateTaskExternalTransferSingleResponse
	(*CreateTaskExternalTransferMultipleResponse)(nil),          // 30: bifast.service.v1.CreateTaskExternalTransferMultipleResponse
	(*ValidateProxyResponse)(nil),                               // 31: bifast.service.v1.ValidateProxyResponse
}
var file_bifast_api_proto_depIdxs = []int32{
	0,  // 0: bifast.service.v1.ApiService.HealthCheck:input_type -> bifast.service.v1.HealthCheckRequest
	1,  // 1: bifast.service.v1.ApiService.GetTaskExternalTransferSingleTemplate:input_type -> bifast.service.v1.GetTaskExternalTransferSingleTemplateRequest
	2,  // 2: bifast.service.v1.ApiService.GetTaskExternalTransferSingleTemplateDetail:input_type -> bifast.service.v1.GetTaskExternalTransferSingleTemplateDetailRequest
	3,  // 3: bifast.service.v1.ApiService.CreateTaskExternalTransferSingleTemplate:input_type -> bifast.service.v1.ExternalTransferSingleTemplate
	4,  // 4: bifast.service.v1.ApiService.DeleteTaskExternalTransferSingleTemplate:input_type -> bifast.service.v1.DeleteTaskExternalTransferSingleTemplateRequest
	5,  // 5: bifast.service.v1.ApiService.GetTaskExternalTransfer:input_type -> bifast.service.v1.GetTaskExternalTransferRequest
	6,  // 6: bifast.service.v1.ApiService.GetTaskExternalTransferFile:input_type -> bifast.service.v1.GetTaskExternalTransferFileRequest
	7,  // 7: bifast.service.v1.ApiService.GetTaskExternalTransferDetail:input_type -> bifast.service.v1.GetTaskExternalTransferDetailRequest
	8,  // 8: bifast.service.v1.ApiService.GetTaskExternalTransferDetailByReffNum:input_type -> bifast.service.v1.GetTaskExternalTransferDetailByReffNumRequest
	9,  // 9: bifast.service.v1.ApiService.CreateExternalTransferTransaction:input_type -> bifast.service.v1.CreateExternalTransferTransactionRequest
	10, // 10: bifast.service.v1.ApiService.CancelExternalTransferTransaction:input_type -> bifast.service.v1.CancelExternalTransferTransactionRequest
	11, // 11: bifast.service.v1.ApiService.SetTaskExternalTransfer:input_type -> bifast.service.v1.SetTaskExternalTransferRequest
	12, // 12: bifast.service.v1.ApiService.ExecExternalTransfer:input_type -> bifast.service.v1.ExecExternalTransferRequest
	13, // 13: bifast.service.v1.ApiService.CreateTaskExternalTransferSingle:input_type -> bifast.service.v1.CreateTaskExternalTransferSingleRequest
	13, // 14: bifast.service.v1.ApiService.CreateTaskExternalTransferSingleHttp:input_type -> bifast.service.v1.CreateTaskExternalTransferSingleRequest
	14, // 15: bifast.service.v1.ApiService.CreateTaskExternalTransferMultiple:input_type -> bifast.service.v1.CreateTaskExternalTransferMultipleRequest
	15, // 16: bifast.service.v1.ApiService.ValidateProxy:input_type -> bifast.service.v1.ValidateProxyRequest
	16, // 17: bifast.service.v1.ApiService.HealthCheck:output_type -> bifast.service.v1.HealthCheckResponse
	17, // 18: bifast.service.v1.ApiService.GetTaskExternalTransferSingleTemplate:output_type -> bifast.service.v1.GetTaskExternalTransferSingleTemplateResponse
	18, // 19: bifast.service.v1.ApiService.GetTaskExternalTransferSingleTemplateDetail:output_type -> bifast.service.v1.GetTaskExternalTransferSingleTemplateDetailResponse
	19, // 20: bifast.service.v1.ApiService.CreateTaskExternalTransferSingleTemplate:output_type -> bifast.service.v1.CreateTaskExternalTransferSingleTemplateResponse
	20, // 21: bifast.service.v1.ApiService.DeleteTaskExternalTransferSingleTemplate:output_type -> bifast.service.v1.DeleteTaskExternalTransferSingleTemplateResponse
	21, // 22: bifast.service.v1.ApiService.GetTaskExternalTransfer:output_type -> bifast.service.v1.GetTaskExternalTransferResponse
	22, // 23: bifast.service.v1.ApiService.GetTaskExternalTransferFile:output_type -> google.api.HttpBody
	23, // 24: bifast.service.v1.ApiService.GetTaskExternalTransferDetail:output_type -> bifast.service.v1.GetTaskExternalTransferDetailResponse
	24, // 25: bifast.service.v1.ApiService.GetTaskExternalTransferDetailByReffNum:output_type -> bifast.service.v1.GetTaskExternalTransferDetailByReffNumResponse
	25, // 26: bifast.service.v1.ApiService.CreateExternalTransferTransaction:output_type -> bifast.service.v1.CreateExternalTransferTransactionResponse
	26, // 27: bifast.service.v1.ApiService.CancelExternalTransferTransaction:output_type -> bifast.service.v1.CancelExternalTransferTransactionResponse
	27, // 28: bifast.service.v1.ApiService.SetTaskExternalTransfer:output_type -> bifast.service.v1.SetTaskExternalTransferResponse
	28, // 29: bifast.service.v1.ApiService.ExecExternalTransfer:output_type -> bifast.service.v1.ExecExternalTransferResponse
	29, // 30: bifast.service.v1.ApiService.CreateTaskExternalTransferSingle:output_type -> bifast.service.v1.CreateTaskExternalTransferSingleResponse
	29, // 31: bifast.service.v1.ApiService.CreateTaskExternalTransferSingleHttp:output_type -> bifast.service.v1.CreateTaskExternalTransferSingleResponse
	30, // 32: bifast.service.v1.ApiService.CreateTaskExternalTransferMultiple:output_type -> bifast.service.v1.CreateTaskExternalTransferMultipleResponse
	31, // 33: bifast.service.v1.ApiService.ValidateProxy:output_type -> bifast.service.v1.ValidateProxyResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_bifast_api_proto_init() }
func file_bifast_api_proto_init() {
	if File_bifast_api_proto != nil {
		return
	}
	file_bifast_gorm_db_proto_init()
	file_bifast_core_payload_proto_init()
	file_bifast_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bifast_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bifast_api_proto_goTypes,
		DependencyIndexes: file_bifast_api_proto_depIdxs,
	}.Build()
	File_bifast_api_proto = out.File
	file_bifast_api_proto_rawDesc = nil
	file_bifast_api_proto_goTypes = nil
	file_bifast_api_proto_depIdxs = nil
}
