// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: role_api.proto

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

var File_role_api_proto protoreflect.FileDescriptor

var file_role_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe7, 0x12, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12,
	0x97, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x92, 0x41, 0x2d, 0x12, 0x0d, 0x47, 0x65,
	0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x1a, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x00, 0x12, 0xcc, 0x01, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x2a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x92, 0x41, 0x39, 0x12,
	0x13, 0x47, 0x65, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x20, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x62, 0x00, 0x12, 0xc6, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x29,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x92, 0x41, 0x37, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20,
	0x72, 0x6f, 0x6c, 0x65, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x1f,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x62,
	0x00, 0x12, 0xb2, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x92, 0x41, 0x37, 0x12,
	0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74,
	0x79, 0x70, 0x65, 0x1a, 0x1f, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x62, 0x00, 0x12, 0xb2, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x72, 0x6f, 0x6c,
	0x65, 0x92, 0x41, 0x37, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x1f, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0xd0, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31,
	0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x3a, 0x01, 0x2a, 0x5a, 0x1c, 0x1a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x33, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x72, 0x6f, 0x6c, 0x65,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x00, 0x12, 0xad, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x37,
	0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1f, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x6f, 0x6c, 0x65,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x27, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x12, 0x5e, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0xbe, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41, 0x39, 0x12, 0x13, 0x47, 0x65, 0x74, 0x20, 0x72,
	0x6f, 0x6c, 0x65, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x1a, 0x20,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64,
	0x62, 0x00, 0x12, 0xaf, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x92, 0x41, 0x38, 0x12, 0x19, 0x47, 0x65,
	0x74, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x19, 0x47, 0x65, 0x74, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x62, 0x00, 0x12, 0xe6, 0x01, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x27, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x22, 0x7d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x92, 0x41, 0x54, 0x12, 0x27, 0x47, 0x65, 0x74, 0x20, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x1a, 0x27, 0x47, 0x65, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x20, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x61,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x62, 0x00, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_role_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                      // 0: role.service.v1.Empty
	(*ListRoleRequest)(nil),            // 1: role.service.v1.ListRoleRequest
	(*ListAuthorityLevelRequest)(nil),  // 2: role.service.v1.ListAuthorityLevelRequest
	(*ListRoleAuthorityRequest)(nil),   // 3: role.service.v1.ListRoleAuthorityRequest
	(*ListUserTypeRequest)(nil),        // 4: role.service.v1.ListUserTypeRequest
	(*ListUserRoleRequest)(nil),        // 5: role.service.v1.ListUserRoleRequest
	(*CreateRoleRequest)(nil),          // 6: role.service.v1.CreateRoleRequest
	(*CreateRoleTaskRequest)(nil),      // 7: role.service.v1.CreateRoleTaskRequest
	(*ListRoleTaskRequest)(nil),        // 8: role.service.v1.ListRoleTaskRequest
	(*GetRoleUserByUserIDReq)(nil),     // 9: role.service.v1.GetRoleUserByUserIDReq
	(*AssignUserRolesRequest)(nil),     // 10: role.service.v1.AssignUserRolesRequest
	(*RoleTaskDetailRequest)(nil),      // 11: role.service.v1.RoleTaskDetailRequest
	(*UserRoleCountReq)(nil),           // 12: role.service.v1.UserRoleCountReq
	(*RoleAuthorityLevelsReq)(nil),     // 13: role.service.v1.RoleAuthorityLevelsReq
	(*HealthCheckResponse)(nil),        // 14: role.service.v1.HealthCheckResponse
	(*ListRoleResponse)(nil),           // 15: role.service.v1.ListRoleResponse
	(*ListAuthorityLevelResponse)(nil), // 16: role.service.v1.ListAuthorityLevelResponse
	(*ListRoleAuthorityResponse)(nil),  // 17: role.service.v1.ListRoleAuthorityResponse
	(*ListUserTypeResponse)(nil),       // 18: role.service.v1.ListUserTypeResponse
	(*ListUserRoleResponse)(nil),       // 19: role.service.v1.ListUserRoleResponse
	(*CreateRoleResponse)(nil),         // 20: role.service.v1.CreateRoleResponse
	(*CreateRoleTaskResponse)(nil),     // 21: role.service.v1.CreateRoleTaskResponse
	(*ListRoleTaskResponse)(nil),       // 22: role.service.v1.ListRoleTaskResponse
	(*GetRoleUserByUserIDRes)(nil),     // 23: role.service.v1.GetRoleUserByUserIDRes
	(*ErrorBodyResponse)(nil),          // 24: role.service.v1.errorBodyResponse
	(*RoleTaskDetailResponse)(nil),     // 25: role.service.v1.RoleTaskDetailResponse
	(*UserRoleCountRes)(nil),           // 26: role.service.v1.UserRoleCountRes
	(*RoleAuthorityLevelsRes)(nil),     // 27: role.service.v1.RoleAuthorityLevelsRes
}
var file_role_api_proto_depIdxs = []int32{
	0,  // 0: role.service.v1.ApiService.HealthCheck:input_type -> role.service.v1.Empty
	1,  // 1: role.service.v1.ApiService.ListRole:input_type -> role.service.v1.ListRoleRequest
	2,  // 2: role.service.v1.ApiService.ListAuthorityLevel:input_type -> role.service.v1.ListAuthorityLevelRequest
	3,  // 3: role.service.v1.ApiService.ListRoleAuthority:input_type -> role.service.v1.ListRoleAuthorityRequest
	4,  // 4: role.service.v1.ApiService.ListUserType:input_type -> role.service.v1.ListUserTypeRequest
	5,  // 5: role.service.v1.ApiService.ListUserRole:input_type -> role.service.v1.ListUserRoleRequest
	6,  // 6: role.service.v1.ApiService.CreateRole:input_type -> role.service.v1.CreateRoleRequest
	7,  // 7: role.service.v1.ApiService.CreateRoleTask:input_type -> role.service.v1.CreateRoleTaskRequest
	8,  // 8: role.service.v1.ApiService.ListRoleTask:input_type -> role.service.v1.ListRoleTaskRequest
	9,  // 9: role.service.v1.ApiService.GetRoleUserByUserID:input_type -> role.service.v1.GetRoleUserByUserIDReq
	10, // 10: role.service.v1.ApiService.AssignUserRoles:input_type -> role.service.v1.AssignUserRolesRequest
	11, // 11: role.service.v1.ApiService.RoleTaskDetail:input_type -> role.service.v1.RoleTaskDetailRequest
	12, // 12: role.service.v1.ApiService.UserRoleCount:input_type -> role.service.v1.UserRoleCountReq
	13, // 13: role.service.v1.ApiService.RoleAuthorityLevels:input_type -> role.service.v1.RoleAuthorityLevelsReq
	14, // 14: role.service.v1.ApiService.HealthCheck:output_type -> role.service.v1.HealthCheckResponse
	15, // 15: role.service.v1.ApiService.ListRole:output_type -> role.service.v1.ListRoleResponse
	16, // 16: role.service.v1.ApiService.ListAuthorityLevel:output_type -> role.service.v1.ListAuthorityLevelResponse
	17, // 17: role.service.v1.ApiService.ListRoleAuthority:output_type -> role.service.v1.ListRoleAuthorityResponse
	18, // 18: role.service.v1.ApiService.ListUserType:output_type -> role.service.v1.ListUserTypeResponse
	19, // 19: role.service.v1.ApiService.ListUserRole:output_type -> role.service.v1.ListUserRoleResponse
	20, // 20: role.service.v1.ApiService.CreateRole:output_type -> role.service.v1.CreateRoleResponse
	21, // 21: role.service.v1.ApiService.CreateRoleTask:output_type -> role.service.v1.CreateRoleTaskResponse
	22, // 22: role.service.v1.ApiService.ListRoleTask:output_type -> role.service.v1.ListRoleTaskResponse
	23, // 23: role.service.v1.ApiService.GetRoleUserByUserID:output_type -> role.service.v1.GetRoleUserByUserIDRes
	24, // 24: role.service.v1.ApiService.AssignUserRoles:output_type -> role.service.v1.errorBodyResponse
	25, // 25: role.service.v1.ApiService.RoleTaskDetail:output_type -> role.service.v1.RoleTaskDetailResponse
	26, // 26: role.service.v1.ApiService.UserRoleCount:output_type -> role.service.v1.UserRoleCountRes
	27, // 27: role.service.v1.ApiService.RoleAuthorityLevels:output_type -> role.service.v1.RoleAuthorityLevelsRes
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_role_api_proto_init() }
func file_role_api_proto_init() {
	if File_role_api_proto != nil {
		return
	}
	file_role_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_role_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_api_proto_goTypes,
		DependencyIndexes: file_role_api_proto_depIdxs,
	}.Build()
	File_role_api_proto = out.File
	file_role_api_proto_rawDesc = nil
	file_role_api_proto_goTypes = nil
	file_role_api_proto_depIdxs = nil
}
