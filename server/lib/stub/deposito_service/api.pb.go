// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: api.proto

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

type ErrorBodyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorBodyResponse) Reset() {
	*x = ErrorBodyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorBodyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorBodyResponse) ProtoMessage() {}

func (x *ErrorBodyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorBodyResponse.ProtoReflect.Descriptor instead.
func (*ErrorBodyResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorBodyResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ErrorBodyResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorBodyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62,
	0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd5, 0x13, 0x0a, 0x0f, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96,
	0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x22,
	0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x1f, 0x1a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x5d, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61,
	0x73, 0x6b, 0x1a, 0x43, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x6f,
	0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x2c, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x69, 0x64, 0x20, 0x69, 0x73, 0x20, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x12, 0xf7, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x8f, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x1a, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x68, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x53, 0x65, 0x74, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x4a, 0x73, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20,
	0x27, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0xe2, 0x80, 0x99, 0x2c, 0x20, 0x27, 0x72, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0xe2, 0x80, 0x99, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0xe2, 0x80, 0x99, 0x2c, 0x20, 0xe2, 0x80, 0x98, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0xe2, 0x80,
	0x99, 0x12, 0xbc, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2c, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x92, 0x41, 0x29, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x1a, 0x0d, 0x67, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0xb9, 0x01, 0x0a, 0x18, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x28, 0x2e,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x5d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b,
	0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x7d, 0x92, 0x41, 0x2e, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x74,
	0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x0d, 0x67,
	0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xd4, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x2f, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12,
	0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x7d, 0x92, 0x41, 0x2d, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x0f, 0x67, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0xa3, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x12, 0x24, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x92, 0x41, 0x29,
	0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0d, 0x47, 0x65, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x67, 0x65, 0x74, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xb4, 0x01, 0x0a, 0x0e, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x92, 0x41, 0x2d, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f,
	0x47, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a,
	0x0f, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0xa4, 0x01, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x12, 0x2a, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x19, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x09, 0x44, 0x41, 0x54, 0x41, 0x20,
	0x53, 0x41, 0x56, 0x45, 0x1a, 0x01, 0x2d, 0x12, 0xd1, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x22, 0x5b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2d, 0x72, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x33, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x0a, 0x03, 0x45,
	0x53, 0x42, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x1a, 0x16, 0x67, 0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x20, 0x72, 0x61, 0x74, 0x65, 0x12, 0xd3, 0x01, 0x0a, 0x13,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x2f, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x74, 0x61, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x33, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x0a, 0x03, 0x45, 0x53, 0x42, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x44, 0x61,
	0x74, 0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x16, 0x67, 0x65, 0x74, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x20, 0x72, 0x61, 0x74,
	0x65, 0x12, 0xdd, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x22, 0x5e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x34, 0x0a, 0x03, 0x41,
	0x6c, 0x6c, 0x0a, 0x03, 0x45, 0x53, 0x42, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74,
	0x61, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0xce, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x36, 0x0a, 0x03, 0x41, 0x6c,
	0x6c, 0x0a, 0x03, 0x45, 0x53, 0x42, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61,
	0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x19, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0xba, 0x01, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0xb0, 0x01, 0x12,
	0x1e, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x20, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a,
	0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e,
	0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*ErrorBodyResponse)(nil),            // 0: deposito.service.v1.errorBodyResponse
	(*Empty)(nil),                        // 1: deposito.service.v1.Empty
	(*CreateDepositoTaskRequest)(nil),    // 2: deposito.service.v1.CreateDepositoTaskRequest
	(*TaskActionRequest)(nil),            // 3: deposito.service.v1.TaskActionRequest
	(*DepositoListTaskRequest)(nil),      // 4: deposito.service.v1.DepositoListTaskRequest
	(*FileTaskListRequest)(nil),          // 5: deposito.service.v1.FileTaskListRequest
	(*GetDepositoTaskByIDRequest)(nil),   // 6: deposito.service.v1.GetDepositoTaskByIDRequest
	(*DataListRequest)(nil),              // 7: deposito.service.v1.DataListRequest
	(*DataDetailRequest)(nil),            // 8: deposito.service.v1.DataDetailRequest
	(*CreateDepositoRequest)(nil),        // 9: deposito.service.v1.CreateDepositoRequest
	(*DepositInquiryRateRequest)(nil),    // 10: deposito.service.v1.DepositInquiryRateRequest
	(*DepositoMaintananceRequest)(nil),   // 11: deposito.service.v1.DepositoMaintananceRequest
	(*DepositoCreateAccountRequest)(nil), // 12: deposito.service.v1.DepositoCreateAccountRequest
	(*DepositoPlacementRequest)(nil),     // 13: deposito.service.v1.DepositoPlacementRequest
	(*CreateDepositoTaskResponse)(nil),   // 14: deposito.service.v1.CreateDepositoTaskResponse
	(*TaskActionResponse)(nil),           // 15: deposito.service.v1.TaskActionResponse
	(*DepositoListTaskResponse)(nil),     // 16: deposito.service.v1.DepositoListTaskResponse
	(*httpbody.HttpBody)(nil),            // 17: google.api.HttpBody
	(*GetDepositoTaskByIDResponse)(nil),  // 18: deposito.service.v1.GetDepositoTaskByIDResponse
	(*DataListResponse)(nil),             // 19: deposito.service.v1.DataListResponse
	(*DataDetailResponse)(nil),           // 20: deposito.service.v1.DataDetailResponse
	(*CreateDepositoResponse)(nil),       // 21: deposito.service.v1.CreateDepositoResponse
	(*DepositInquiryRateRespons)(nil),    // 22: deposito.service.v1.DepositInquiryRateRespons
	(*DepositoMaintananceRespons)(nil),   // 23: deposito.service.v1.DepositoMaintananceRespons
	(*DepositoCreateAccountRespons)(nil), // 24: deposito.service.v1.DepositoCreateAccountRespons
	(*DepositoPlacementRespons)(nil),     // 25: deposito.service.v1.DepositoPlacementRespons
}
var file_api_proto_depIdxs = []int32{
	2,  // 0: deposito.service.v1.DepositoService.CreateDepositoTask:input_type -> deposito.service.v1.CreateDepositoTaskRequest
	3,  // 1: deposito.service.v1.DepositoService.DepositoActionTask:input_type -> deposito.service.v1.TaskActionRequest
	4,  // 2: deposito.service.v1.DepositoService.DepositoListTask:input_type -> deposito.service.v1.DepositoListTaskRequest
	5,  // 3: deposito.service.v1.DepositoService.DownloadDepositoListTask:input_type -> deposito.service.v1.FileTaskListRequest
	6,  // 4: deposito.service.v1.DepositoService.GetDepositoTaskByID:input_type -> deposito.service.v1.GetDepositoTaskByIDRequest
	7,  // 5: deposito.service.v1.DepositoService.ListDeposito:input_type -> deposito.service.v1.DataListRequest
	8,  // 6: deposito.service.v1.DepositoService.DepositoDetail:input_type -> deposito.service.v1.DataDetailRequest
	9,  // 7: deposito.service.v1.DepositoService.createDeposito:input_type -> deposito.service.v1.CreateDepositoRequest
	10, // 8: deposito.service.v1.DepositoService.DepositInquiryRate:input_type -> deposito.service.v1.DepositInquiryRateRequest
	11, // 9: deposito.service.v1.DepositoService.DepositoMaintanance:input_type -> deposito.service.v1.DepositoMaintananceRequest
	12, // 10: deposito.service.v1.DepositoService.DepositoCreateAccount:input_type -> deposito.service.v1.DepositoCreateAccountRequest
	13, // 11: deposito.service.v1.DepositoService.DepositoPlacement:input_type -> deposito.service.v1.DepositoPlacementRequest
	14, // 12: deposito.service.v1.DepositoService.CreateDepositoTask:output_type -> deposito.service.v1.CreateDepositoTaskResponse
	15, // 13: deposito.service.v1.DepositoService.DepositoActionTask:output_type -> deposito.service.v1.TaskActionResponse
	16, // 14: deposito.service.v1.DepositoService.DepositoListTask:output_type -> deposito.service.v1.DepositoListTaskResponse
	17, // 15: deposito.service.v1.DepositoService.DownloadDepositoListTask:output_type -> google.api.HttpBody
	18, // 16: deposito.service.v1.DepositoService.GetDepositoTaskByID:output_type -> deposito.service.v1.GetDepositoTaskByIDResponse
	19, // 17: deposito.service.v1.DepositoService.ListDeposito:output_type -> deposito.service.v1.DataListResponse
	20, // 18: deposito.service.v1.DepositoService.DepositoDetail:output_type -> deposito.service.v1.DataDetailResponse
	21, // 19: deposito.service.v1.DepositoService.createDeposito:output_type -> deposito.service.v1.CreateDepositoResponse
	22, // 20: deposito.service.v1.DepositoService.DepositInquiryRate:output_type -> deposito.service.v1.DepositInquiryRateRespons
	23, // 21: deposito.service.v1.DepositoService.DepositoMaintanance:output_type -> deposito.service.v1.DepositoMaintananceRespons
	24, // 22: deposito.service.v1.DepositoService.DepositoCreateAccount:output_type -> deposito.service.v1.DepositoCreateAccountRespons
	25, // 23: deposito.service.v1.DepositoService.DepositoPlacement:output_type -> deposito.service.v1.DepositoPlacementRespons
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorBodyResponse); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
