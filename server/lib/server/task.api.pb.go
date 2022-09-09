// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: task.api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoggerTestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LoggerTestReq) Reset() {
	*x = LoggerTestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerTestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerTestReq) ProtoMessage() {}

func (x *LoggerTestReq) ProtoReflect() protoreflect.Message {
	mi := &file_task_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerTestReq.ProtoReflect.Descriptor instead.
func (*LoggerTestReq) Descriptor() ([]byte, []int) {
	return file_task_api_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerTestReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LoggerTestRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LoggerTestRes) Reset() {
	*x = LoggerTestRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerTestRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerTestRes) ProtoMessage() {}

func (x *LoggerTestRes) ProtoReflect() protoreflect.Message {
	mi := &file_task_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerTestRes.ProtoReflect.Descriptor instead.
func (*LoggerTestRes) Descriptor() ([]byte, []int) {
	return file_task_api_proto_rawDescGZIP(), []int{1}
}

func (x *LoggerTestRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_task_api_proto protoreflect.FileDescriptor

var file_task_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x29, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfc, 0x1b, 0x0a, 0x0b,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x53,
	0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x56, 0x12, 0x22, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56,
	0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x45, 0x56, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x57, 0x6f, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x2f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x57, 0x69, 0x74, 0x68, 0x57, 0x6f, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x12, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xe0, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22,
	0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x71, 0x12, 0x19, 0x53, 0x65, 0x74, 0x20, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x54, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69,
	0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74,
	0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x0a, 0x20,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20,
	0x0a, 0x20, 0x2d, 0x20, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xe9, 0x01, 0x0a, 0x09, 0x53,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x22, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x56, 0x22,
	0x94, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x2f, 0x65, 0x76, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x71, 0x12, 0x19, 0x53, 0x65, 0x74, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x54, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20,
	0x62, 0x79, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x0a, 0x20, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x20, 0x0a, 0x20, 0x2d, 0x20, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x0a, 0x20, 0x2d, 0x20,
	0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xb5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x92, 0x41,
	0x3d, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x27, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x52,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x6c, 0x75, 0x63, 0x6b, 0x12, 0x25, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x6c, 0x75, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x6c, 0x75, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xab, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x1a, 0x23, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x56,
	0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x92, 0x41, 0x3d, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x27, 0x54, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0xda, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x92, 0x41, 0x58, 0x12, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x37, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x61, 0x73,
	0x65, 0x20, 0x6f, 0x6e, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0xf6, 0x01, 0x0a, 0x15, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x92, 0x41,
	0x58, 0x12, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x62, 0x79, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x37, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x61, 0x73, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xcf, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x12, 0x21,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x92, 0x41, 0x54, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x79, 0x20, 0x73, 0x74, 0x65, 0x70, 0x20,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x35, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69,
	0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x62, 0x61, 0x73, 0x65, 0x20, 0x6f,
	0x6e, 0x20, 0x73, 0x74, 0x65, 0x70, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x12, 0xb1, 0x02, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b,
	0x57, 0x69, 0x74, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x39, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x57,
	0x69, 0x74, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x94, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6d, 0x79, 0x2d,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x92, 0x41,
	0x6e, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20, 0x6d, 0x79, 0x20, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x20, 0x1a, 0x50, 0x54,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x77, 0x6e,
	0x20, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x12,
	0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x45, 0x56, 0x12, 0x27, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x56, 0x1a, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x23, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x0e, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x12, 0xc4, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x6c, 0x6f, 0x67, 0x73,
	0x2f, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3b, 0x12, 0x11, 0x47,
	0x65, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73,
	0x1a, 0x26, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0xe0, 0x01, 0x0a, 0x14, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x28, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x87, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x34, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x6c,
	0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x7d,
	0x3a, 0x01, 0x2a, 0x92, 0x41, 0x45, 0x12, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x1a, 0x2b,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x71, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x00, 0x12, 0x5a,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0xaf, 0x01, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x92, 0x41, 0xa5, 0x01, 0x12, 0x13, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a, 0x01, 0x01, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_api_proto_rawDescOnce sync.Once
	file_task_api_proto_rawDescData = file_task_api_proto_rawDesc
)

func file_task_api_proto_rawDescGZIP() []byte {
	file_task_api_proto_rawDescOnce.Do(func() {
		file_task_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_api_proto_rawDescData)
	})
	return file_task_api_proto_rawDescData
}

var file_task_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_task_api_proto_goTypes = []interface{}{
	(*LoggerTestReq)(nil),                             // 0: task.service.v1.LoggerTestReq
	(*LoggerTestRes)(nil),                             // 1: task.service.v1.LoggerTestRes
	(*SaveTaskRequest)(nil),                           // 2: task.service.v1.SaveTaskRequest
	(*SaveTaskRequestEV)(nil),                         // 3: task.service.v1.SaveTaskRequestEV
	(*GenerateTaskWithWokflowRequest)(nil),            // 4: task.service.v1.GenerateTaskWithWokflowRequest
	(*SetTaskRequest)(nil),                            // 5: task.service.v1.SetTaskRequest
	(*SetTaskRequestEV)(nil),                          // 6: task.service.v1.SetTaskRequestEV
	(*ListTaskRequest)(nil),                           // 7: task.service.v1.ListTaskRequest
	(*ListTaskPluckRequest)(nil),                      // 8: task.service.v1.ListTaskPluckRequest
	(*ListTaskRequestEV)(nil),                         // 9: task.service.v1.ListTaskRequestEV
	(*GraphStatusRequest)(nil),                        // 10: task.service.v1.GraphStatusRequest
	(*GraphStatusColumnTypeRequest)(nil),              // 11: task.service.v1.GraphStatusColumnTypeRequest
	(*GraphStepRequest)(nil),                          // 12: task.service.v1.GraphStepRequest
	(*GetMyPendingTaskWithWorkflowGraphRequest)(nil),  // 13: task.service.v1.GetMyPendingTaskWithWorkflowGraphRequest
	(*ListRequest)(nil),                               // 14: task.service.v1.ListRequest
	(*AssignaTypeIDRequest)(nil),                      // 15: task.service.v1.AssignaTypeIDRequest
	(*AssignaTypeIDRequestEV)(nil),                    // 16: task.service.v1.AssignaTypeIDRequestEV
	(*GetTaskByIDReq)(nil),                            // 17: task.service.v1.GetTaskByIDReq
	(*GetTaskByTypeIDReq)(nil),                        // 18: task.service.v1.GetTaskByTypeIDReq
	(*RejectBySystemReq)(nil),                         // 19: task.service.v1.RejectBySystemReq
	(*GetActivityLogsReq)(nil),                        // 20: task.service.v1.GetActivityLogsReq
	(*DownloadActivityLogsReq)(nil),                   // 21: task.service.v1.DownloadActivityLogsReq
	(*UpdateTaskDataReq)(nil),                         // 22: task.service.v1.UpdateTaskDataReq
	(*SaveTaskResponse)(nil),                          // 23: task.service.v1.SaveTaskResponse
	(*SaveTaskResponseEV)(nil),                        // 24: task.service.v1.SaveTaskResponseEV
	(*SetTaskResponse)(nil),                           // 25: task.service.v1.SetTaskResponse
	(*SetTaskResponseEV)(nil),                         // 26: task.service.v1.SetTaskResponseEV
	(*ListTaskResponse)(nil),                          // 27: task.service.v1.ListTaskResponse
	(*ListTaskPluckResponse)(nil),                     // 28: task.service.v1.ListTaskPluckResponse
	(*ListTaskResponseEV)(nil),                        // 29: task.service.v1.ListTaskResponseEV
	(*GraphStatusResponse)(nil),                       // 30: task.service.v1.GraphStatusResponse
	(*GraphStatusColumnTypeResponse)(nil),             // 31: task.service.v1.GraphStatusColumnTypeResponse
	(*GraphStepResponse)(nil),                         // 32: task.service.v1.GraphStepResponse
	(*GetMyPendingTaskWithWorkflowGraphResponse)(nil), // 33: task.service.v1.GetMyPendingTaskWithWorkflowGraphResponse
	(*AssignaTypeIDResponse)(nil),                     // 34: task.service.v1.AssignaTypeIDResponse
	(*GetTaskByIDRes)(nil),                            // 35: task.service.v1.GetTaskByIDRes
	(*GetTaskByTypeIDRes)(nil),                        // 36: task.service.v1.GetTaskByTypeIDRes
	(*RejectBySystemRes)(nil),                         // 37: task.service.v1.RejectBySystemRes
	(*GetActivityLogsRes)(nil),                        // 38: task.service.v1.GetActivityLogsRes
	(*httpbody.HttpBody)(nil),                         // 39: google.api.HttpBody
	(*UpdateTaskDataRes)(nil),                         // 40: task.service.v1.UpdateTaskDataRes
}
var file_task_api_proto_depIdxs = []int32{
	2,  // 0: task.service.v1.TaskService.SaveTaskWithData:input_type -> task.service.v1.SaveTaskRequest
	3,  // 1: task.service.v1.TaskService.SaveTaskWithDataEV:input_type -> task.service.v1.SaveTaskRequestEV
	4,  // 2: task.service.v1.TaskService.GenerateTaskWithWokflow:input_type -> task.service.v1.GenerateTaskWithWokflowRequest
	2,  // 3: task.service.v1.TaskService.SaveTaskWithWorkflow:input_type -> task.service.v1.SaveTaskRequest
	5,  // 4: task.service.v1.TaskService.SetTask:input_type -> task.service.v1.SetTaskRequest
	6,  // 5: task.service.v1.TaskService.SetTaskEV:input_type -> task.service.v1.SetTaskRequestEV
	7,  // 6: task.service.v1.TaskService.GetListTaskWithToken:input_type -> task.service.v1.ListTaskRequest
	7,  // 7: task.service.v1.TaskService.GetListTask:input_type -> task.service.v1.ListTaskRequest
	8,  // 8: task.service.v1.TaskService.GetListTaskPluck:input_type -> task.service.v1.ListTaskPluckRequest
	9,  // 9: task.service.v1.TaskService.GetListTaskEV:input_type -> task.service.v1.ListTaskRequestEV
	10, // 10: task.service.v1.TaskService.GetTaskGraphStatus:input_type -> task.service.v1.GraphStatusRequest
	11, // 11: task.service.v1.TaskService.GraphStatusColumnType:input_type -> task.service.v1.GraphStatusColumnTypeRequest
	12, // 12: task.service.v1.TaskService.GetTaskGraphStep:input_type -> task.service.v1.GraphStepRequest
	13, // 13: task.service.v1.TaskService.GetMyPendingTaskWithWorkflowGraph:input_type -> task.service.v1.GetMyPendingTaskWithWorkflowGraphRequest
	14, // 14: task.service.v1.TaskService.GetListAnnouncement:input_type -> task.service.v1.ListRequest
	15, // 15: task.service.v1.TaskService.AssignTypeID:input_type -> task.service.v1.AssignaTypeIDRequest
	16, // 16: task.service.v1.TaskService.AssignTypeIDEV:input_type -> task.service.v1.AssignaTypeIDRequestEV
	17, // 17: task.service.v1.TaskService.GetTaskByID:input_type -> task.service.v1.GetTaskByIDReq
	18, // 18: task.service.v1.TaskService.GetTaskByTypeID:input_type -> task.service.v1.GetTaskByTypeIDReq
	19, // 19: task.service.v1.TaskService.RejectBySystem:input_type -> task.service.v1.RejectBySystemReq
	20, // 20: task.service.v1.TaskService.GetActivityLogs:input_type -> task.service.v1.GetActivityLogsReq
	21, // 21: task.service.v1.TaskService.DownloadActivityLogs:input_type -> task.service.v1.DownloadActivityLogsReq
	0,  // 22: task.service.v1.TaskService.TestLogger:input_type -> task.service.v1.LoggerTestReq
	22, // 23: task.service.v1.TaskService.UpdateTaskData:input_type -> task.service.v1.UpdateTaskDataReq
	23, // 24: task.service.v1.TaskService.SaveTaskWithData:output_type -> task.service.v1.SaveTaskResponse
	24, // 25: task.service.v1.TaskService.SaveTaskWithDataEV:output_type -> task.service.v1.SaveTaskResponseEV
	23, // 26: task.service.v1.TaskService.GenerateTaskWithWokflow:output_type -> task.service.v1.SaveTaskResponse
	23, // 27: task.service.v1.TaskService.SaveTaskWithWorkflow:output_type -> task.service.v1.SaveTaskResponse
	25, // 28: task.service.v1.TaskService.SetTask:output_type -> task.service.v1.SetTaskResponse
	26, // 29: task.service.v1.TaskService.SetTaskEV:output_type -> task.service.v1.SetTaskResponseEV
	27, // 30: task.service.v1.TaskService.GetListTaskWithToken:output_type -> task.service.v1.ListTaskResponse
	27, // 31: task.service.v1.TaskService.GetListTask:output_type -> task.service.v1.ListTaskResponse
	28, // 32: task.service.v1.TaskService.GetListTaskPluck:output_type -> task.service.v1.ListTaskPluckResponse
	29, // 33: task.service.v1.TaskService.GetListTaskEV:output_type -> task.service.v1.ListTaskResponseEV
	30, // 34: task.service.v1.TaskService.GetTaskGraphStatus:output_type -> task.service.v1.GraphStatusResponse
	31, // 35: task.service.v1.TaskService.GraphStatusColumnType:output_type -> task.service.v1.GraphStatusColumnTypeResponse
	32, // 36: task.service.v1.TaskService.GetTaskGraphStep:output_type -> task.service.v1.GraphStepResponse
	33, // 37: task.service.v1.TaskService.GetMyPendingTaskWithWorkflowGraph:output_type -> task.service.v1.GetMyPendingTaskWithWorkflowGraphResponse
	27, // 38: task.service.v1.TaskService.GetListAnnouncement:output_type -> task.service.v1.ListTaskResponse
	34, // 39: task.service.v1.TaskService.AssignTypeID:output_type -> task.service.v1.AssignaTypeIDResponse
	34, // 40: task.service.v1.TaskService.AssignTypeIDEV:output_type -> task.service.v1.AssignaTypeIDResponse
	35, // 41: task.service.v1.TaskService.GetTaskByID:output_type -> task.service.v1.GetTaskByIDRes
	36, // 42: task.service.v1.TaskService.GetTaskByTypeID:output_type -> task.service.v1.GetTaskByTypeIDRes
	37, // 43: task.service.v1.TaskService.RejectBySystem:output_type -> task.service.v1.RejectBySystemRes
	38, // 44: task.service.v1.TaskService.GetActivityLogs:output_type -> task.service.v1.GetActivityLogsRes
	39, // 45: task.service.v1.TaskService.DownloadActivityLogs:output_type -> google.api.HttpBody
	1,  // 46: task.service.v1.TaskService.TestLogger:output_type -> task.service.v1.LoggerTestRes
	40, // 47: task.service.v1.TaskService.UpdateTaskData:output_type -> task.service.v1.UpdateTaskDataRes
	24, // [24:48] is the sub-list for method output_type
	0,  // [0:24] is the sub-list for method input_type
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
	if !protoimpl.UnsafeEnabled {
		file_task_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerTestReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerTestRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_api_proto_goTypes,
		DependencyIndexes: file_task_api_proto_depIdxs,
		MessageInfos:      file_task_api_proto_msgTypes,
	}.Build()
	File_task_api_proto = out.File
	file_task_api_proto_rawDesc = nil
	file_task_api_proto_goTypes = nil
	file_task_api_proto_depIdxs = nil
}
