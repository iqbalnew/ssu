// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: auth_api.proto

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
	return file_auth_api_proto_enumTypes[0].Descriptor()
}

func (TempGenToken_Role) Type() protoreflect.EnumType {
	return &file_auth_api_proto_enumTypes[0]
}

func (x TempGenToken_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TempGenToken_Role.Descriptor instead.
func (TempGenToken_Role) EnumDescriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{3, 0}
}

type BricamsVerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	USERID    string `protobuf:"bytes,1,opt,name=USERID,proto3" json:"USERID,omitempty"`
	SESSIONID string `protobuf:"bytes,2,opt,name=SESSIONID,proto3" json:"SESSIONID,omitempty"`
}

func (x *BricamsVerifyReq) Reset() {
	*x = BricamsVerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BricamsVerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BricamsVerifyReq) ProtoMessage() {}

func (x *BricamsVerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BricamsVerifyReq.ProtoReflect.Descriptor instead.
func (*BricamsVerifyReq) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{0}
}

func (x *BricamsVerifyReq) GetUSERID() string {
	if x != nil {
		return x.USERID
	}
	return ""
}

func (x *BricamsVerifyReq) GetSESSIONID() string {
	if x != nil {
		return x.SESSIONID
	}
	return ""
}

type BricamsLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	USERID    string `protobuf:"bytes,1,opt,name=USERID,proto3" json:"USERID,omitempty"`
	SESSIONID string `protobuf:"bytes,2,opt,name=SESSIONID,proto3" json:"SESSIONID,omitempty"`
	DTTIME    string `protobuf:"bytes,3,opt,name=DTTIME,proto3" json:"DTTIME,omitempty"`
}

func (x *BricamsLoginReq) Reset() {
	*x = BricamsLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BricamsLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BricamsLoginReq) ProtoMessage() {}

func (x *BricamsLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BricamsLoginReq.ProtoReflect.Descriptor instead.
func (*BricamsLoginReq) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{1}
}

func (x *BricamsLoginReq) GetUSERID() string {
	if x != nil {
		return x.USERID
	}
	return ""
}

func (x *BricamsLoginReq) GetSESSIONID() string {
	if x != nil {
		return x.SESSIONID
	}
	return ""
}

func (x *BricamsLoginReq) GetDTTIME() string {
	if x != nil {
		return x.DTTIME
	}
	return ""
}

type BricamsVerifyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BricamsVerifyRes) Reset() {
	*x = BricamsVerifyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BricamsVerifyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BricamsVerifyRes) ProtoMessage() {}

func (x *BricamsVerifyRes) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BricamsVerifyRes.ProtoReflect.Descriptor instead.
func (*BricamsVerifyRes) Descriptor() ([]byte, []int) {
	return file_auth_api_proto_rawDescGZIP(), []int{2}
}

type TempGenToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role TempGenToken_Role `protobuf:"varint,1,opt,name=role,proto3,enum=auth.service.v1.TempGenToken_Role" json:"role,omitempty"`
}

func (x *TempGenToken) Reset() {
	*x = TempGenToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempGenToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempGenToken) ProtoMessage() {}

func (x *TempGenToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_api_proto_msgTypes[3]
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
	return file_auth_api_proto_rawDescGZIP(), []int{3}
}

func (x *TempGenToken) GetRole() TempGenToken_Role {
	if x != nil {
		return x.Role
	}
	return TempGenToken_bankadmin
}

var File_auth_api_proto protoreflect.FileDescriptor

var file_auth_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48,
	0x0a, 0x10, 0x42, 0x72, 0x69, 0x63, 0x61, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x53, 0x45, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x53, 0x45, 0x52, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x44, 0x22, 0x5f, 0x0a, 0x0f, 0x42, 0x72, 0x69, 0x63,
	0x61, 0x6d, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x53, 0x45, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x53, 0x45,
	0x52, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x54, 0x54, 0x49, 0x4d, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x44, 0x54, 0x54, 0x49, 0x4d, 0x45, 0x22, 0x12, 0x0a, 0x10, 0x42, 0x72, 0x69,
	0x63, 0x61, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x22, 0x8a, 0x01,
	0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d,
	0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x72, 0x10, 0x02, 0x32, 0x84, 0x09, 0x0a, 0x0a, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x92, 0x41, 0x32, 0x12, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x1a, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x00, 0x12, 0xa3, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x35, 0x12, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x20, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x6f, 0x6e,
	0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x00, 0x12, 0x94, 0x01, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x92, 0x41, 0x37, 0x12, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x20, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x1a, 0x26, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x00, 0x12, 0xa3, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x92, 0x41, 0x3b, 0x12, 0x10,
	0x47, 0x65, 0x74, 0x20, 0x6d, 0x65, 0x20, 0x75, 0x73, 0x65, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x00, 0x12, 0xb4, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x92, 0x41, 0x3d, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x6d, 0x65, 0x20, 0x75, 0x73, 0x65, 0x20,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61,
	0x64, 0x64, 0x6f, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x00,
	0x12, 0x9a, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x69, 0x63, 0x61, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x69, 0x63, 0x61, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x50, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x92, 0x41, 0x35, 0x12, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64, 0x64, 0x6f,
	0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x00, 0x12, 0xa1, 0x01,
	0x0a, 0x0c, 0x42, 0x72, 0x69, 0x63, 0x61, 0x6d, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x72, 0x69, 0x63, 0x61, 0x6d, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x92, 0x41, 0x35, 0x12, 0x0a, 0x41, 0x75,
	0x74, 0x68, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x25, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x6f, 0x20,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x00, 0x42, 0x72, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x92, 0x41, 0x69, 0x5a, 0x59, 0x0a, 0x57,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_api_proto_rawDescOnce sync.Once
	file_auth_api_proto_rawDescData = file_auth_api_proto_rawDesc
)

func file_auth_api_proto_rawDescGZIP() []byte {
	file_auth_api_proto_rawDescOnce.Do(func() {
		file_auth_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_api_proto_rawDescData)
	})
	return file_auth_api_proto_rawDescData
}

var file_auth_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_api_proto_goTypes = []interface{}{
	(TempGenToken_Role)(0),      // 0: auth.service.v1.TempGenToken.Role
	(*BricamsVerifyReq)(nil),    // 1: auth.service.v1.BricamsVerifyReq
	(*BricamsLoginReq)(nil),     // 2: auth.service.v1.BricamsLoginReq
	(*BricamsVerifyRes)(nil),    // 3: auth.service.v1.BricamsVerifyRes
	(*TempGenToken)(nil),        // 4: auth.service.v1.TempGenToken
	(*Empty)(nil),               // 5: auth.service.v1.Empty
	(*LoginRequest)(nil),        // 6: auth.service.v1.LoginRequest
	(*LogoutRequest)(nil),       // 7: auth.service.v1.LogoutRequest
	(*VerifyTokenReq)(nil),      // 8: auth.service.v1.VerifyTokenReq
	(*VerifySessionReq)(nil),    // 9: auth.service.v1.VerifySessionReq
	(*HealthCheckResponse)(nil), // 10: auth.service.v1.HealthCheckResponse
	(*LoginResponse)(nil),       // 11: auth.service.v1.LoginResponse
	(*VerifyTokenRes)(nil),      // 12: auth.service.v1.VerifyTokenRes
	(*httpbody.HttpBody)(nil),   // 13: google.api.HttpBody
}
var file_auth_api_proto_depIdxs = []int32{
	0,  // 0: auth.service.v1.TempGenToken.role:type_name -> auth.service.v1.TempGenToken.Role
	5,  // 1: auth.service.v1.ApiService.HealthCheck:input_type -> auth.service.v1.Empty
	6,  // 2: auth.service.v1.ApiService.Login:input_type -> auth.service.v1.LoginRequest
	7,  // 3: auth.service.v1.ApiService.Logout:input_type -> auth.service.v1.LogoutRequest
	8,  // 4: auth.service.v1.ApiService.VerifyToken:input_type -> auth.service.v1.VerifyTokenReq
	9,  // 5: auth.service.v1.ApiService.GetTokenBySession:input_type -> auth.service.v1.VerifySessionReq
	1,  // 6: auth.service.v1.ApiService.BricamsVerify:input_type -> auth.service.v1.BricamsVerifyReq
	2,  // 7: auth.service.v1.ApiService.BricamsLogin:input_type -> auth.service.v1.BricamsLoginReq
	10, // 8: auth.service.v1.ApiService.HealthCheck:output_type -> auth.service.v1.HealthCheckResponse
	11, // 9: auth.service.v1.ApiService.Login:output_type -> auth.service.v1.LoginResponse
	5,  // 10: auth.service.v1.ApiService.Logout:output_type -> auth.service.v1.Empty
	12, // 11: auth.service.v1.ApiService.VerifyToken:output_type -> auth.service.v1.VerifyTokenRes
	11, // 12: auth.service.v1.ApiService.GetTokenBySession:output_type -> auth.service.v1.LoginResponse
	13, // 13: auth.service.v1.ApiService.BricamsVerify:output_type -> google.api.HttpBody
	11, // 14: auth.service.v1.ApiService.BricamsLogin:output_type -> auth.service.v1.LoginResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_auth_api_proto_init() }
func file_auth_api_proto_init() {
	if File_auth_api_proto != nil {
		return
	}
	file_auth_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BricamsVerifyReq); i {
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
		file_auth_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BricamsLoginReq); i {
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
		file_auth_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BricamsVerifyRes); i {
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
		file_auth_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_auth_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_api_proto_goTypes,
		DependencyIndexes: file_auth_api_proto_depIdxs,
		EnumInfos:         file_auth_api_proto_enumTypes,
		MessageInfos:      file_auth_api_proto_msgTypes,
	}.Build()
	File_auth_api_proto = out.File
	file_auth_api_proto_rawDesc = nil
	file_auth_api_proto_goTypes = nil
	file_auth_api_proto_depIdxs = nil
}
