// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: announcement.api.proto

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
	TempGenToken_bankadmin    TempGenToken_Role = 0
	TempGenToken_companyadmin TempGenToken_Role = 1
	TempGenToken_companyuser  TempGenToken_Role = 2
)

// Enum value maps for TempGenToken_Role.
var (
	TempGenToken_Role_name = map[int32]string{
		0: "bankadmin",
		1: "companyadmin",
		2: "companyuser",
	}
	TempGenToken_Role_value = map[string]int32{
		"bankadmin":    0,
		"companyadmin": 1,
		"companyuser":  2,
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
	return file_announcement_api_proto_enumTypes[0].Descriptor()
}

func (TempGenToken_Role) Type() protoreflect.EnumType {
	return &file_announcement_api_proto_enumTypes[0]
}

func (x TempGenToken_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TempGenToken_Role.Descriptor instead.
func (TempGenToken_Role) EnumDescriptor() ([]byte, []int) {
	return file_announcement_api_proto_rawDescGZIP(), []int{0, 0}
}

type TempGenToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role TempGenToken_Role `protobuf:"varint,1,opt,name=role,proto3,enum=announcement.service.v1.TempGenToken_Role" json:"role,omitempty"`
}

func (x *TempGenToken) Reset() {
	*x = TempGenToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announcement_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempGenToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempGenToken) ProtoMessage() {}

func (x *TempGenToken) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_api_proto_msgTypes[0]
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
	return file_announcement_api_proto_rawDescGZIP(), []int{0}
}

func (x *TempGenToken) GetRole() TempGenToken_Role {
	if x != nil {
		return x.Role
	}
	return TempGenToken_bankadmin
}

var File_announcement_api_proto protoreflect.FileDescriptor

var file_announcement_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0c,
	0x54, 0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x75, 0x73, 0x65, 0x72, 0x10, 0x02, 0x32, 0x8f, 0x08, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12, 0xa2, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x29, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x57, 0x54, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x7d, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12,
	0x8b, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12, 0x7d, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x32, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9f, 0x01, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x32, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x98,
	0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x92, 0x41, 0x02, 0x62, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x92, 0x41, 0x02, 0x62, 0x00, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_announcement_api_proto_rawDescOnce sync.Once
	file_announcement_api_proto_rawDescData = file_announcement_api_proto_rawDesc
)

func file_announcement_api_proto_rawDescGZIP() []byte {
	file_announcement_api_proto_rawDescOnce.Do(func() {
		file_announcement_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_announcement_api_proto_rawDescData)
	})
	return file_announcement_api_proto_rawDescData
}

var file_announcement_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_announcement_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_announcement_api_proto_goTypes = []interface{}{
	(TempGenToken_Role)(0),             // 0: announcement.service.v1.TempGenToken.Role
	(*TempGenToken)(nil),               // 1: announcement.service.v1.TempGenToken
	(*Empty)(nil),                      // 2: announcement.service.v1.Empty
	(*ListRequest)(nil),                // 3: announcement.service.v1.ListRequest
	(*CreateAnnouncementRequest)(nil),  // 4: announcement.service.v1.CreateAnnouncementRequest
	(*HealthCheckResponse)(nil),        // 5: announcement.service.v1.HealthCheckResponse
	(*JWTTokenResponse)(nil),           // 6: announcement.service.v1.JWTTokenResponse
	(*ListAnnouncementResponse)(nil),   // 7: announcement.service.v1.ListAnnouncementResponse
	(*CreateAnnouncementResponse)(nil), // 8: announcement.service.v1.CreateAnnouncementResponse
	(*ListEventTypeResponse)(nil),      // 9: announcement.service.v1.ListEventTypeResponse
}
var file_announcement_api_proto_depIdxs = []int32{
	0, // 0: announcement.service.v1.TempGenToken.role:type_name -> announcement.service.v1.TempGenToken.Role
	2, // 1: announcement.service.v1.ApiService.HealthCheck:input_type -> announcement.service.v1.Empty
	1, // 2: announcement.service.v1.ApiService.GenerateTokenByRole:input_type -> announcement.service.v1.TempGenToken
	3, // 3: announcement.service.v1.ApiService.ListAnnouncement:input_type -> announcement.service.v1.ListRequest
	4, // 4: announcement.service.v1.ApiService.CreateAnnouncement:input_type -> announcement.service.v1.CreateAnnouncementRequest
	4, // 5: announcement.service.v1.ApiService.CreateAnnouncementTask:input_type -> announcement.service.v1.CreateAnnouncementRequest
	3, // 6: announcement.service.v1.ApiService.ListAnnouncementActive:input_type -> announcement.service.v1.ListRequest
	3, // 7: announcement.service.v1.ApiService.ListEventType:input_type -> announcement.service.v1.ListRequest
	5, // 8: announcement.service.v1.ApiService.HealthCheck:output_type -> announcement.service.v1.HealthCheckResponse
	6, // 9: announcement.service.v1.ApiService.GenerateTokenByRole:output_type -> announcement.service.v1.JWTTokenResponse
	7, // 10: announcement.service.v1.ApiService.ListAnnouncement:output_type -> announcement.service.v1.ListAnnouncementResponse
	8, // 11: announcement.service.v1.ApiService.CreateAnnouncement:output_type -> announcement.service.v1.CreateAnnouncementResponse
	8, // 12: announcement.service.v1.ApiService.CreateAnnouncementTask:output_type -> announcement.service.v1.CreateAnnouncementResponse
	7, // 13: announcement.service.v1.ApiService.ListAnnouncementActive:output_type -> announcement.service.v1.ListAnnouncementResponse
	9, // 14: announcement.service.v1.ApiService.ListEventType:output_type -> announcement.service.v1.ListEventTypeResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_announcement_api_proto_init() }
func file_announcement_api_proto_init() {
	if File_announcement_api_proto != nil {
		return
	}
	file_announcement_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_announcement_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_announcement_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_announcement_api_proto_goTypes,
		DependencyIndexes: file_announcement_api_proto_depIdxs,
		EnumInfos:         file_announcement_api_proto_enumTypes,
		MessageInfos:      file_announcement_api_proto_msgTypes,
	}.Build()
	File_announcement_api_proto = out.File
	file_announcement_api_proto_rawDesc = nil
	file_announcement_api_proto_goTypes = nil
	file_announcement_api_proto_depIdxs = nil
}
