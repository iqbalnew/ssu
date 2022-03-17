// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: menu.api.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TempGenToken_Role int32

const (
	TempGenToken_bankadmin         TempGenToken_Role = 0
	TempGenToken_notificationadmin TempGenToken_Role = 1
	TempGenToken_notificationuser  TempGenToken_Role = 2
)

// Enum value maps for TempGenToken_Role.
var (
	TempGenToken_Role_name = map[int32]string{
		0: "bankadmin",
		1: "notificationadmin",
		2: "notificationuser",
	}
	TempGenToken_Role_value = map[string]int32{
		"bankadmin":         0,
		"notificationadmin": 1,
		"notificationuser":  2,
	}
)

func (x TempGenToken_Role) Enum() *TempGenToken_Role {
	p := new(TempGenToken_Role)
	*p = x
	return p
}

func (x TempGenToken_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TempGenToken_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_menu_api_proto_enumTypes[0].Descriptor()
}

func (TempGenToken_Role) Type() protoreflect.EnumType {
	return &file_menu_api_proto_enumTypes[0]
}

func (x TempGenToken_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TempGenToken_Role.Descriptor instead.
func (TempGenToken_Role) EnumDescriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{0, 0}
}

type TempGenToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role TempGenToken_Role `protobuf:"varint,1,opt,name=role,proto3,enum=menu.service.v1.TempGenToken_Role" json:"role,omitempty"`
}

func (x *TempGenToken) Reset() {
	*x = TempGenToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempGenToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempGenToken) ProtoMessage() {}

func (x *TempGenToken) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TempGenToken.ProtoReflect.Descriptor instead.
func (*TempGenToken) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{0}
}

func (x *TempGenToken) GetRole() TempGenToken_Role {
	if x != nil {
		return x.Role
	}
	return TempGenToken_bankadmin
}

var File_menu_api_proto protoreflect.FileDescriptor

var file_menu_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x72, 0x10, 0x02,
	0x32, 0xd9, 0x13, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x16, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x20, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0xa7, 0x01, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a, 0x0f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x35, 0x12, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x20, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1e, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x20, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x62, 0x00, 0x12, 0xa9, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41, 0x37, 0x12, 0x12, 0x47, 0x65, 0x74,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x1a,
	0x1f, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6d, 0x65, 0x6e, 0x75,
	0x62, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x22, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xcc, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x3a, 0x01, 0x2a, 0x5a, 0x1c, 0x1a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e,
	0x75, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x33, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x54, 0x61,
	0x73, 0x6b, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x1a, 0x1d, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61, 0x73,
	0x6b, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x62, 0x00, 0x12, 0xaf, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92, 0x41,
	0x39, 0x12, 0x13, 0x47, 0x65, 0x74, 0x20, 0x54, 0x61, 0x73, 0x6b, 0x20, 0x6d, 0x65, 0x6e, 0x75,
	0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x47, 0x65, 0x74, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x6d, 0x65,
	0x6e, 0x75, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x62, 0x00, 0x12, 0xa1, 0x01, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x92, 0x41,
	0x2d, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x6e, 0x75,
	0x1a, 0x1a, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x62, 0x00, 0x12, 0x98,
	0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x92, 0x41, 0x31, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x1c, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74,
	0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x62, 0x00, 0x12, 0xaf, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x92, 0x41, 0x3d, 0x12, 0x15,
	0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x6e, 0x75, 0x20, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x22, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x6e,
	0x75, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x00, 0x12, 0xc1, 0x01, 0x0a, 0x15,
	0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x22, 0x52, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x47, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x01,
	0x2a, 0x5a, 0x27, 0x1a, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12,
	0xa1, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x92, 0x41,
	0x02, 0x62, 0x00, 0x12, 0x9e, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x7d, 0x92,
	0x41, 0x02, 0x62, 0x00, 0x12, 0x66, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65,
	0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12, 0xb2, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x3a,
	0x01, 0x2a, 0x5a, 0x24, 0x1a, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12, 0x5d, 0x0a,
	0x0f, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x72, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_api_proto_rawDescOnce sync.Once
	file_menu_api_proto_rawDescData = file_menu_api_proto_rawDesc
)

func file_menu_api_proto_rawDescGZIP() []byte {
	file_menu_api_proto_rawDescOnce.Do(func() {
		file_menu_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_api_proto_rawDescData)
	})
	return file_menu_api_proto_rawDescData
}

var file_menu_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_menu_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_menu_api_proto_goTypes = []interface{}{
	(TempGenToken_Role)(0),               // 0: menu.service.v1.TempGenToken.Role
	(*TempGenToken)(nil),                 // 1: menu.service.v1.TempGenToken
	(*Empty)(nil),                        // 2: menu.service.v1.Empty
	(*CommonResponse)(nil),               // 3: menu.service.v1.CommonResponse
	(*ListMenuTaskRequest)(nil),          // 4: menu.service.v1.ListMenuTaskRequest
	(*CreateMenuRequest)(nil),            // 5: menu.service.v1.CreateMenuRequest
	(*CreateMenuTaskRequest)(nil),        // 6: menu.service.v1.CreateMenuTaskRequest
	(*GetByTaskID)(nil),                  // 7: menu.service.v1.GetByTaskID
	(*ListMenuDataRequest)(nil),          // 8: menu.service.v1.ListMenuDataRequest
	(*SetTaskMenuAppearanceReq)(nil),     // 9: menu.service.v1.SetTaskMenuAppearanceReq
	(*GetListTaskMenuAppearanceReq)(nil), // 10: menu.service.v1.GetListTaskMenuAppearanceReq
	(*GetTaskMenuAppearanceReq)(nil),     // 11: menu.service.v1.GetTaskMenuAppearanceReq
	(*SaveMenuAppearanceReq)(nil),        // 12: menu.service.v1.SaveMenuAppearanceReq
	(*GetMenuAppearanceReq)(nil),         // 13: menu.service.v1.GetMenuAppearanceReq
	(*SetTaskMenuLicenseReq)(nil),        // 14: menu.service.v1.SetTaskMenuLicenseReq
	(*SaveMenuLicenseReq)(nil),           // 15: menu.service.v1.SaveMenuLicenseReq
	(*HealthCheckResponse)(nil),          // 16: menu.service.v1.HealthCheckResponse
	(*ListMenuResponse)(nil),             // 17: menu.service.v1.ListMenuResponse
	(*CreateMenuResponse)(nil),           // 18: menu.service.v1.CreateMenuResponse
	(*ListMenuActiveResponse)(nil),       // 19: menu.service.v1.ListMenuActiveResponse
	(*ListModuleResponse)(nil),           // 20: menu.service.v1.ListModuleResponse
	(*ListMenuDisableResponse)(nil),      // 21: menu.service.v1.ListMenuDisableResponse
	(*SetTaskMenuAppearanceRes)(nil),     // 22: menu.service.v1.SetTaskMenuAppearanceRes
	(*GetListTaskMenuAppearanceRes)(nil), // 23: menu.service.v1.GetListTaskMenuAppearanceRes
	(*GetTaskMenuAppearanceRes)(nil),     // 24: menu.service.v1.GetTaskMenuAppearanceRes
	(*SaveMenuAppearanceRes)(nil),        // 25: menu.service.v1.SaveMenuAppearanceRes
	(*GetMenuAppearanceRes)(nil),         // 26: menu.service.v1.GetMenuAppearanceRes
	(*SetTaskMenuLicenseRes)(nil),        // 27: menu.service.v1.SetTaskMenuLicenseRes
	(*SaveMenuLicenseRes)(nil),           // 28: menu.service.v1.SaveMenuLicenseRes
}
var file_menu_api_proto_depIdxs = []int32{
	0,  // 0: menu.service.v1.TempGenToken.role:type_name -> menu.service.v1.TempGenToken.Role
	2,  // 1: menu.service.v1.ApiService.HealthCheck:input_type -> menu.service.v1.Empty
	3,  // 2: menu.service.v1.ApiService.UpdateMenuAdmin:input_type -> menu.service.v1.CommonResponse
	4,  // 3: menu.service.v1.ApiService.ListMenuTask:input_type -> menu.service.v1.ListMenuTaskRequest
	5,  // 4: menu.service.v1.ApiService.CreateMenu:input_type -> menu.service.v1.CreateMenuRequest
	6,  // 5: menu.service.v1.ApiService.CreateMenuTask:input_type -> menu.service.v1.CreateMenuTaskRequest
	7,  // 6: menu.service.v1.ApiService.GetMenuTaskByID:input_type -> menu.service.v1.GetByTaskID
	8,  // 7: menu.service.v1.ApiService.ListMenu:input_type -> menu.service.v1.ListMenuDataRequest
	2,  // 8: menu.service.v1.ApiService.ListModule:input_type -> menu.service.v1.Empty
	2,  // 9: menu.service.v1.ApiService.ListMenuDisable:input_type -> menu.service.v1.Empty
	9,  // 10: menu.service.v1.ApiService.SetTaskMenuAppearance:input_type -> menu.service.v1.SetTaskMenuAppearanceReq
	10, // 11: menu.service.v1.ApiService.GetListTaskMenuAppearance:input_type -> menu.service.v1.GetListTaskMenuAppearanceReq
	11, // 12: menu.service.v1.ApiService.GetTaskMenuAppearance:input_type -> menu.service.v1.GetTaskMenuAppearanceReq
	12, // 13: menu.service.v1.ApiService.SaveMenuAppearance:input_type -> menu.service.v1.SaveMenuAppearanceReq
	13, // 14: menu.service.v1.ApiService.GetMenuAppearance:input_type -> menu.service.v1.GetMenuAppearanceReq
	14, // 15: menu.service.v1.ApiService.SetTaskMenuLicense:input_type -> menu.service.v1.SetTaskMenuLicenseReq
	15, // 16: menu.service.v1.ApiService.SaveMenuLicense:input_type -> menu.service.v1.SaveMenuLicenseReq
	16, // 17: menu.service.v1.ApiService.HealthCheck:output_type -> menu.service.v1.HealthCheckResponse
	3,  // 18: menu.service.v1.ApiService.UpdateMenuAdmin:output_type -> menu.service.v1.CommonResponse
	17, // 19: menu.service.v1.ApiService.ListMenuTask:output_type -> menu.service.v1.ListMenuResponse
	18, // 20: menu.service.v1.ApiService.CreateMenu:output_type -> menu.service.v1.CreateMenuResponse
	18, // 21: menu.service.v1.ApiService.CreateMenuTask:output_type -> menu.service.v1.CreateMenuResponse
	17, // 22: menu.service.v1.ApiService.GetMenuTaskByID:output_type -> menu.service.v1.ListMenuResponse
	19, // 23: menu.service.v1.ApiService.ListMenu:output_type -> menu.service.v1.ListMenuActiveResponse
	20, // 24: menu.service.v1.ApiService.ListModule:output_type -> menu.service.v1.ListModuleResponse
	21, // 25: menu.service.v1.ApiService.ListMenuDisable:output_type -> menu.service.v1.ListMenuDisableResponse
	22, // 26: menu.service.v1.ApiService.SetTaskMenuAppearance:output_type -> menu.service.v1.SetTaskMenuAppearanceRes
	23, // 27: menu.service.v1.ApiService.GetListTaskMenuAppearance:output_type -> menu.service.v1.GetListTaskMenuAppearanceRes
	24, // 28: menu.service.v1.ApiService.GetTaskMenuAppearance:output_type -> menu.service.v1.GetTaskMenuAppearanceRes
	25, // 29: menu.service.v1.ApiService.SaveMenuAppearance:output_type -> menu.service.v1.SaveMenuAppearanceRes
	26, // 30: menu.service.v1.ApiService.GetMenuAppearance:output_type -> menu.service.v1.GetMenuAppearanceRes
	27, // 31: menu.service.v1.ApiService.SetTaskMenuLicense:output_type -> menu.service.v1.SetTaskMenuLicenseRes
	28, // 32: menu.service.v1.ApiService.SaveMenuLicense:output_type -> menu.service.v1.SaveMenuLicenseRes
	17, // [17:33] is the sub-list for method output_type
	1,  // [1:17] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_menu_api_proto_init() }
func file_menu_api_proto_init() {
	if File_menu_api_proto != nil {
		return
	}
	file_menu_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_menu_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TempGenToken); i {
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
			RawDescriptor: file_menu_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_api_proto_goTypes,
		DependencyIndexes: file_menu_api_proto_depIdxs,
		EnumInfos:         file_menu_api_proto_enumTypes,
		MessageInfos:      file_menu_api_proto_msgTypes,
	}.Build()
	File_menu_api_proto = out.File
	file_menu_api_proto_rawDesc = nil
	file_menu_api_proto_goTypes = nil
	file_menu_api_proto_depIdxs = nil
}
