// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: task.api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
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

var File_task_api_proto protoreflect.FileDescriptor

var file_task_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65,
	0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x10, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x56, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x23, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x56, 0x22, 0x00, 0x12, 0xe0, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x91, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x71, 0x12, 0x19, 0x53, 0x65, 0x74, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x54, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20,
	0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x0a, 0x20, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x20, 0x0a, 0x20, 0x2d, 0x20, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x0a, 0x20, 0x2d, 0x20,
	0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe9, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x56, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x56, 0x22, 0x94, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x2f, 0x65, 0x76, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x71, 0x12, 0x19, 0x53, 0x65, 0x74, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x74, 0x61,
	0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x54, 0x54, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x0a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a,
	0x20, 0x0a, 0x20, 0x2d, 0x20, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x20, 0x0a, 0x20, 0x2d,
	0x20, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0xac, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x92, 0x41, 0x3d, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x27, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x74, 0x61, 0x73,
	0x6b, 0x12, 0xab, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x45, 0x56, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x56, 0x22, 0x51, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92,
	0x41, 0x3d, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x27, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12,
	0xda, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x92, 0x41, 0x58, 0x12, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x74,
	0x61, 0x73, 0x6b, 0x1a, 0x37, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73,
	0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x61, 0x73, 0x65, 0x20, 0x6f, 0x6e,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0xf6, 0x01, 0x0a,
	0x15, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x92, 0x41, 0x58, 0x12, 0x1d, 0x47,
	0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x79,
	0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x37, 0x54, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x62, 0x61, 0x73, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x20, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xcf, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x92, 0x41, 0x54, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x62, 0x79, 0x20, 0x73, 0x74, 0x65, 0x70, 0x20, 0x74, 0x61, 0x73, 0x6b,
	0x1a, 0x35, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x61, 0x73, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x73, 0x74,
	0x65, 0x70, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5d, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x25, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61,
	0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x45, 0x56,
	0x12, 0x27, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x42,
	0x46, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x3d, 0x12, 0x13, 0x0a, 0x0c, 0x54, 0x61,
	0x73, 0x6b, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a,
	0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_task_api_proto_goTypes = []interface{}{
	(*SaveTaskRequest)(nil),               // 0: task.service.v1.SaveTaskRequest
	(*SaveTaskRequestEV)(nil),             // 1: task.service.v1.SaveTaskRequestEV
	(*SetTaskRequest)(nil),                // 2: task.service.v1.SetTaskRequest
	(*SetTaskRequestEV)(nil),              // 3: task.service.v1.SetTaskRequestEV
	(*ListTaskRequest)(nil),               // 4: task.service.v1.ListTaskRequest
	(*ListTaskRequestEV)(nil),             // 5: task.service.v1.ListTaskRequestEV
	(*GraphStatusRequest)(nil),            // 6: task.service.v1.GraphStatusRequest
	(*GraphStatusColumnTypeRequest)(nil),  // 7: task.service.v1.GraphStatusColumnTypeRequest
	(*GraphStepRequest)(nil),              // 8: task.service.v1.GraphStepRequest
	(*ListRequest)(nil),                   // 9: task.service.v1.ListRequest
	(*AssignaTypeIDRequest)(nil),          // 10: task.service.v1.AssignaTypeIDRequest
	(*AssignaTypeIDRequestEV)(nil),        // 11: task.service.v1.AssignaTypeIDRequestEV
	(*GetTaskByIDReq)(nil),                // 12: task.service.v1.GetTaskByIDReq
	(*GetTaskByTypeIDReq)(nil),            // 13: task.service.v1.GetTaskByTypeIDReq
	(*SaveTaskResponse)(nil),              // 14: task.service.v1.SaveTaskResponse
	(*SaveTaskResponseEV)(nil),            // 15: task.service.v1.SaveTaskResponseEV
	(*SetTaskResponse)(nil),               // 16: task.service.v1.SetTaskResponse
	(*SetTaskResponseEV)(nil),             // 17: task.service.v1.SetTaskResponseEV
	(*ListTaskResponse)(nil),              // 18: task.service.v1.ListTaskResponse
	(*ListTaskResponseEV)(nil),            // 19: task.service.v1.ListTaskResponseEV
	(*GraphStatusResponse)(nil),           // 20: task.service.v1.GraphStatusResponse
	(*GraphStatusColumnTypeResponse)(nil), // 21: task.service.v1.GraphStatusColumnTypeResponse
	(*GraphStepResponse)(nil),             // 22: task.service.v1.GraphStepResponse
	(*AssignaTypeIDResponse)(nil),         // 23: task.service.v1.AssignaTypeIDResponse
	(*GetTaskByIDRes)(nil),                // 24: task.service.v1.GetTaskByIDRes
	(*GetTaskByTypeIDRes)(nil),            // 25: task.service.v1.GetTaskByTypeIDRes
}
var file_task_api_proto_depIdxs = []int32{
	0,  // 0: task.service.v1.TaskService.SaveTaskWithData:input_type -> task.service.v1.SaveTaskRequest
	1,  // 1: task.service.v1.TaskService.SaveTaskWithDataEV:input_type -> task.service.v1.SaveTaskRequestEV
	2,  // 2: task.service.v1.TaskService.SetTask:input_type -> task.service.v1.SetTaskRequest
	3,  // 3: task.service.v1.TaskService.SetTaskEV:input_type -> task.service.v1.SetTaskRequestEV
	4,  // 4: task.service.v1.TaskService.GetListTask:input_type -> task.service.v1.ListTaskRequest
	5,  // 5: task.service.v1.TaskService.GetListTaskEV:input_type -> task.service.v1.ListTaskRequestEV
	6,  // 6: task.service.v1.TaskService.GetTaskGraphStatus:input_type -> task.service.v1.GraphStatusRequest
	7,  // 7: task.service.v1.TaskService.GraphStatusColumnType:input_type -> task.service.v1.GraphStatusColumnTypeRequest
	8,  // 8: task.service.v1.TaskService.GetTaskGraphStep:input_type -> task.service.v1.GraphStepRequest
	9,  // 9: task.service.v1.TaskService.GetListAnnouncement:input_type -> task.service.v1.ListRequest
	10, // 10: task.service.v1.TaskService.AssignTypeID:input_type -> task.service.v1.AssignaTypeIDRequest
	11, // 11: task.service.v1.TaskService.AssignTypeIDEV:input_type -> task.service.v1.AssignaTypeIDRequestEV
	12, // 12: task.service.v1.TaskService.GetTaskByID:input_type -> task.service.v1.GetTaskByIDReq
	13, // 13: task.service.v1.TaskService.GetTaskByTypeID:input_type -> task.service.v1.GetTaskByTypeIDReq
	14, // 14: task.service.v1.TaskService.SaveTaskWithData:output_type -> task.service.v1.SaveTaskResponse
	15, // 15: task.service.v1.TaskService.SaveTaskWithDataEV:output_type -> task.service.v1.SaveTaskResponseEV
	16, // 16: task.service.v1.TaskService.SetTask:output_type -> task.service.v1.SetTaskResponse
	17, // 17: task.service.v1.TaskService.SetTaskEV:output_type -> task.service.v1.SetTaskResponseEV
	18, // 18: task.service.v1.TaskService.GetListTask:output_type -> task.service.v1.ListTaskResponse
	19, // 19: task.service.v1.TaskService.GetListTaskEV:output_type -> task.service.v1.ListTaskResponseEV
	20, // 20: task.service.v1.TaskService.GetTaskGraphStatus:output_type -> task.service.v1.GraphStatusResponse
	21, // 21: task.service.v1.TaskService.GraphStatusColumnType:output_type -> task.service.v1.GraphStatusColumnTypeResponse
	22, // 22: task.service.v1.TaskService.GetTaskGraphStep:output_type -> task.service.v1.GraphStepResponse
	18, // 23: task.service.v1.TaskService.GetListAnnouncement:output_type -> task.service.v1.ListTaskResponse
	23, // 24: task.service.v1.TaskService.AssignTypeID:output_type -> task.service.v1.AssignaTypeIDResponse
	23, // 25: task.service.v1.TaskService.AssignTypeIDEV:output_type -> task.service.v1.AssignaTypeIDResponse
	24, // 26: task.service.v1.TaskService.GetTaskByID:output_type -> task.service.v1.GetTaskByIDRes
	25, // 27: task.service.v1.TaskService.GetTaskByTypeID:output_type -> task.service.v1.GetTaskByTypeIDRes
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_task_api_proto_init() }
func file_task_api_proto_init() {
	if File_task_api_proto != nil {
		return
	}
	file_task_payload_proto_init()
	file_task_payload_ev_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_api_proto_goTypes,
		DependencyIndexes: file_task_api_proto_depIdxs,
	}.Build()
	File_task_api_proto = out.File
	file_task_api_proto_rawDesc = nil
	file_task_api_proto_goTypes = nil
	file_task_api_proto_depIdxs = nil
}
