// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: auth_gorm_db.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Authentication) Reset() {
	*x = Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication) ProtoMessage() {}

func (x *Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_auth_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication.ProtoReflect.Descriptor instead.
func (*Authentication) Descriptor() ([]byte, []int) {
	return file_auth_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *Authentication) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Authentication) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Authentication) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        uint64                 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	LastLoginTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
}

func (x *UserLoginTime) Reset() {
	*x = UserLoginTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginTime) ProtoMessage() {}

func (x *UserLoginTime) ProtoReflect() protoreflect.Message {
	mi := &file_auth_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginTime.ProtoReflect.Descriptor instead.
func (*UserLoginTime) Descriptor() ([]byte, []int) {
	return file_auth_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginTime) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserLoginTime) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginTime) GetLastLoginTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginTime
	}
	return nil
}

type Jail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JailID    uint64                 `protobuf:"varint,1,opt,name=jailID,proto3" json:"jailID,omitempty"`
	Token     string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TokenType string                 `protobuf:"bytes,3,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Jail) Reset() {
	*x = Jail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jail) ProtoMessage() {}

func (x *Jail) ProtoReflect() protoreflect.Message {
	mi := &file_auth_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jail.ProtoReflect.Descriptor instead.
func (*Jail) Descriptor() ([]byte, []int) {
	return file_auth_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *Jail) GetJailID() uint64 {
	if x != nil {
		return x.JailID
	}
	return 0
}

func (x *Jail) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Jail) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Jail) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *Jail) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TokenStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserID    uint64                 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Username  string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	SessionID string                 `protobuf:"bytes,4,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	DateTime  string                 `protobuf:"bytes,5,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *TokenStore) Reset() {
	*x = TokenStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_gorm_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenStore) ProtoMessage() {}

func (x *TokenStore) ProtoReflect() protoreflect.Message {
	mi := &file_auth_gorm_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenStore.ProtoReflect.Descriptor instead.
func (*TokenStore) Descriptor() ([]byte, []int) {
	return file_auth_gorm_db_proto_rawDescGZIP(), []int{3}
}

func (x *TokenStore) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenStore) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TokenStore) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenStore) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *TokenStore) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *TokenStore) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TokenStore) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SessionIDPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PairID    uint64                 `protobuf:"varint,1,opt,name=pairID,proto3" json:"pairID,omitempty"`
	SessionID string                 `protobuf:"bytes,2,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Token     string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *SessionIDPair) Reset() {
	*x = SessionIDPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_gorm_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionIDPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionIDPair) ProtoMessage() {}

func (x *SessionIDPair) ProtoReflect() protoreflect.Message {
	mi := &file_auth_gorm_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionIDPair.ProtoReflect.Descriptor instead.
func (*SessionIDPair) Descriptor() ([]byte, []int) {
	return file_auth_gorm_db_proto_rawDescGZIP(), []int{4}
}

func (x *SessionIDPair) GetPairID() uint64 {
	if x != nil {
		return x.PairID
	}
	return 0
}

func (x *SessionIDPair) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *SessionIDPair) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SessionIDPair) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *SessionIDPair) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_auth_gorm_db_proto protoreflect.FileDescriptor

var file_auth_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1d, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x28,
	0x01, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x2c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x17,
	0xba, 0xb9, 0x19, 0x13, 0x08, 0x01, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1d, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x28, 0x01, 0x40,
	0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59,
	0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x17, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x18, 0xba, 0xb9, 0x19, 0x14, 0x08,
	0x01, 0x1a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x22, 0xc2, 0x02, 0x0a, 0x04, 0x4a, 0x61, 0x69, 0x6c, 0x12, 0x34, 0x0a, 0x06,
	0x6a, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x06, 0x4a, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x28,
	0x01, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06, 0x6a, 0x61, 0x69, 0x6c,
	0x49, 0x44, 0x12, 0x30, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1a, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x30, 0x01, 0x40, 0x01, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x40, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x1b,
	0x0a, 0x19, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x76,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x52, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x3a, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x08, 0x01, 0x1a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x6a, 0x61, 0x69, 0x6c, 0x22, 0xc2, 0x03, 0x0a, 0x0a, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x15, 0x0a,
	0x13, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x28, 0x01,
	0x30, 0x01, 0x40, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x13, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x40, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x40,
	0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c,
	0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x15, 0x0a, 0x13, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x40, 0x01, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x08, 0x01,
	0x1a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xcc, 0x02,
	0x0a, 0x0d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x70, 0x61, 0x69, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1c, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x06, 0x50, 0x61, 0x69, 0x72,
	0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x69, 0x72, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19,
	0x17, 0x0a, 0x15, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x30, 0x01, 0x40, 0x01, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1a, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x30, 0x01, 0x40, 0x01, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x3a, 0x17, 0xba, 0xb9, 0x19, 0x13, 0x08, 0x01, 0x1a, 0x0f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_gorm_db_proto_rawDescOnce sync.Once
	file_auth_gorm_db_proto_rawDescData = file_auth_gorm_db_proto_rawDesc
)

func file_auth_gorm_db_proto_rawDescGZIP() []byte {
	file_auth_gorm_db_proto_rawDescOnce.Do(func() {
		file_auth_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_gorm_db_proto_rawDescData)
	})
	return file_auth_gorm_db_proto_rawDescData
}

var file_auth_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_auth_gorm_db_proto_goTypes = []interface{}{
	(*Authentication)(nil),        // 0: auth.service.v1.Authentication
	(*UserLoginTime)(nil),         // 1: auth.service.v1.UserLoginTime
	(*Jail)(nil),                  // 2: auth.service.v1.Jail
	(*TokenStore)(nil),            // 3: auth.service.v1.TokenStore
	(*SessionIDPair)(nil),         // 4: auth.service.v1.SessionIDPair
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_auth_gorm_db_proto_depIdxs = []int32{
	5, // 0: auth.service.v1.UserLoginTime.lastLoginTime:type_name -> google.protobuf.Timestamp
	5, // 1: auth.service.v1.Jail.expiredAt:type_name -> google.protobuf.Timestamp
	5, // 2: auth.service.v1.Jail.createdAt:type_name -> google.protobuf.Timestamp
	5, // 3: auth.service.v1.TokenStore.createdAt:type_name -> google.protobuf.Timestamp
	5, // 4: auth.service.v1.TokenStore.updatedAt:type_name -> google.protobuf.Timestamp
	5, // 5: auth.service.v1.SessionIDPair.expiredAt:type_name -> google.protobuf.Timestamp
	5, // 6: auth.service.v1.SessionIDPair.createdAt:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_auth_gorm_db_proto_init() }
func file_auth_gorm_db_proto_init() {
	if File_auth_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication); i {
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
		file_auth_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginTime); i {
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
		file_auth_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jail); i {
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
		file_auth_gorm_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenStore); i {
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
		file_auth_gorm_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionIDPair); i {
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
			RawDescriptor: file_auth_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_gorm_db_proto_goTypes,
		DependencyIndexes: file_auth_gorm_db_proto_depIdxs,
		MessageInfos:      file_auth_gorm_db_proto_msgTypes,
	}.Build()
	File_auth_gorm_db_proto = out.File
	file_auth_gorm_db_proto_rawDesc = nil
	file_auth_gorm_db_proto_goTypes = nil
	file_auth_gorm_db_proto_depIdxs = nil
}
