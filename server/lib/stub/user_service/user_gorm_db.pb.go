// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: user_gorm_db.proto

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

type User_IDType int32

const (
	User_NULL     User_IDType = 0
	User_KTP      User_IDType = 1
	User_SIM      User_IDType = 2
	User_PASSPORT User_IDType = 3
	User_KITAS    User_IDType = 4
	User_KITAPI   User_IDType = 5
)

// Enum value maps for User_IDType.
var (
	User_IDType_name = map[int32]string{
		0: "NULL",
		1: "KTP",
		2: "SIM",
		3: "PASSPORT",
		4: "KITAS",
		5: "KITAPI",
	}
	User_IDType_value = map[string]int32{
		"NULL":     0,
		"KTP":      1,
		"SIM":      2,
		"PASSPORT": 3,
		"KITAS":    4,
		"KITAPI":   5,
	}
)

func (x User_IDType) Enum() *User_IDType {
	p := new(User_IDType)
	*p = x
	return p
}

func (x User_IDType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_IDType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_gorm_db_proto_enumTypes[0].Descriptor()
}

func (User_IDType) Type() protoreflect.EnumType {
	return &file_user_gorm_db_proto_enumTypes[0]
}

func (x User_IDType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_IDType.Descriptor instead.
func (User_IDType) EnumDescriptor() ([]byte, []int) {
	return file_user_gorm_db_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       uint64                 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username     string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IdType       User_IDType            `protobuf:"varint,3,opt,name=idType,proto3,enum=user.service.v1.User_IDType" json:"idType,omitempty"`
	IdNumber     string                 `protobuf:"bytes,4,opt,name=idNumber,proto3" json:"idNumber,omitempty"`
	UserTypeID   uint64                 `protobuf:"varint,5,opt,name=userTypeID,proto3" json:"userTypeID,omitempty"`
	UserTypeName string                 `protobuf:"bytes,6,opt,name=userTypeName,proto3" json:"userTypeName,omitempty"`
	CompanyID    uint64                 `protobuf:"varint,7,opt,name=companyID,proto3" json:"companyID,omitempty"`
	CompanyName  string                 `protobuf:"bytes,8,opt,name=companyName,proto3" json:"companyName,omitempty"`
	IsBlocked    bool                   `protobuf:"varint,9,opt,name=isBlocked,proto3" json:"isBlocked,omitempty"`
	CreatedByID  uint64                 `protobuf:"varint,10,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID  uint64                 `protobuf:"varint,11,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	DeletedByID  uint64                 `protobuf:"varint,12,opt,name=deletedByID,proto3" json:"deletedByID,omitempty"`
	Authority    string                 `protobuf:"bytes,13,opt,name=authority,proto3" json:"authority,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIdType() User_IDType {
	if x != nil {
		return x.IdType
	}
	return User_NULL
}

func (x *User) GetIdNumber() string {
	if x != nil {
		return x.IdNumber
	}
	return ""
}

func (x *User) GetUserTypeID() uint64 {
	if x != nil {
		return x.UserTypeID
	}
	return 0
}

func (x *User) GetUserTypeName() string {
	if x != nil {
		return x.UserTypeName
	}
	return ""
}

func (x *User) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *User) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *User) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *User) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *User) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *User) GetDeletedByID() uint64 {
	if x != nil {
		return x.DeletedByID
	}
	return 0
}

func (x *User) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_user_gorm_db_proto protoreflect.FileDescriptor

var file_user_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x33, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x10, 0x0a, 0x0e, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x40, 0x01, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x44,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x49, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x40, 0x01, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x69, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x08, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x40, 0x01, 0x52, 0x08, 0x69, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x40, 0x01, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x92, 0x41, 0x04,
	0x9a, 0x02, 0x01, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12,
	0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x09, 0x69, 0x73,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1f, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1f, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3f, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x1d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2f,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x45, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b,
	0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x49, 0x0a, 0x06, 0x49, 0x44, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x54, 0x50, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x53, 0x49, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53,
	0x50, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x49, 0x54, 0x41, 0x53, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x49, 0x54, 0x41, 0x50, 0x49, 0x10, 0x05, 0x3a, 0x0d, 0xba,
	0xb9, 0x19, 0x09, 0x08, 0x01, 0x1a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_gorm_db_proto_rawDescOnce sync.Once
	file_user_gorm_db_proto_rawDescData = file_user_gorm_db_proto_rawDesc
)

func file_user_gorm_db_proto_rawDescGZIP() []byte {
	file_user_gorm_db_proto_rawDescOnce.Do(func() {
		file_user_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_gorm_db_proto_rawDescData)
	})
	return file_user_gorm_db_proto_rawDescData
}

var file_user_gorm_db_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_gorm_db_proto_goTypes = []interface{}{
	(User_IDType)(0),              // 0: user.service.v1.User.IDType
	(*User)(nil),                  // 1: user.service.v1.User
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_user_gorm_db_proto_depIdxs = []int32{
	0, // 0: user.service.v1.User.idType:type_name -> user.service.v1.User.IDType
	2, // 1: user.service.v1.User.createdAt:type_name -> google.protobuf.Timestamp
	2, // 2: user.service.v1.User.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 3: user.service.v1.User.deletedAt:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_gorm_db_proto_init() }
func file_user_gorm_db_proto_init() {
	if File_user_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_gorm_db_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_gorm_db_proto_goTypes,
		DependencyIndexes: file_user_gorm_db_proto_depIdxs,
		EnumInfos:         file_user_gorm_db_proto_enumTypes,
		MessageInfos:      file_user_gorm_db_proto_msgTypes,
	}.Build()
	File_user_gorm_db_proto = out.File
	file_user_gorm_db_proto_rawDesc = nil
	file_user_gorm_db_proto_goTypes = nil
	file_user_gorm_db_proto_depIdxs = nil
}
