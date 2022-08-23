// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.16.1
// source: beneficiary_account_gorm_db.proto

package pb

import (
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

// Example User
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role      string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beneficiary_account_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_beneficiary_account_gorm_db_proto_msgTypes[0]
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
	return file_beneficiary_account_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
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

type BeneficiaryAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeneficiaryAccountID uint64                 `protobuf:"varint,1,opt,name=beneficiaryAccountID,proto3" json:"beneficiaryAccountID,omitempty"`
	CompanyID            uint64                 `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Country              string                 `protobuf:"bytes,17,opt,name=country,proto3" json:"country,omitempty"`
	Bic                  string                 `protobuf:"bytes,18,opt,name=bic,proto3" json:"bic,omitempty"`
	MasterBankID         string                 `protobuf:"bytes,16,opt,name=masterBankID,proto3" json:"masterBankID,omitempty"`
	AccountNumber        string                 `protobuf:"bytes,3,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	AccountAlias         string                 `protobuf:"bytes,4,opt,name=accountAlias,proto3" json:"accountAlias,omitempty"`
	AccountName          string                 `protobuf:"bytes,5,opt,name=accountName,proto3" json:"accountName,omitempty"`
	AccountType          string                 `protobuf:"bytes,6,opt,name=accountType,proto3" json:"accountType,omitempty"`
	AccountStatus        string                 `protobuf:"bytes,7,opt,name=accountStatus,proto3" json:"accountStatus,omitempty"`
	AccountCurrency      string                 `protobuf:"bytes,8,opt,name=accountCurrency,proto3" json:"accountCurrency,omitempty"`
	AccessLevel          string                 `protobuf:"bytes,9,opt,name=accessLevel,proto3" json:"accessLevel,omitempty"`
	IsOwnedByCompany     string                 `protobuf:"bytes,10,opt,name=isOwnedByCompany,proto3" json:"isOwnedByCompany,omitempty"`
	CreatedByID          uint64                 `protobuf:"varint,11,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID          uint64                 `protobuf:"varint,12,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	DeletedByID          uint64                 `protobuf:"varint,13,opt,name=deletedByID,proto3" json:"deletedByID,omitempty"`
	CompanyName          string                 `protobuf:"bytes,14,opt,name=companyName,proto3" json:"companyName,omitempty"`
	MasterBankName       string                 `protobuf:"bytes,15,opt,name=masterBankName,proto3" json:"masterBankName,omitempty"`
	RoleID               uint64                 `protobuf:"varint,19,opt,name=roleID,proto3" json:"roleID,omitempty"`
	Disabled             bool                   `protobuf:"varint,20,opt,name=disabled,proto3" json:"disabled,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *BeneficiaryAccount) Reset() {
	*x = BeneficiaryAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beneficiary_account_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeneficiaryAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeneficiaryAccount) ProtoMessage() {}

func (x *BeneficiaryAccount) ProtoReflect() protoreflect.Message {
	mi := &file_beneficiary_account_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeneficiaryAccount.ProtoReflect.Descriptor instead.
func (*BeneficiaryAccount) Descriptor() ([]byte, []int) {
	return file_beneficiary_account_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *BeneficiaryAccount) GetBeneficiaryAccountID() uint64 {
	if x != nil {
		return x.BeneficiaryAccountID
	}
	return 0
}

func (x *BeneficiaryAccount) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *BeneficiaryAccount) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *BeneficiaryAccount) GetBic() string {
	if x != nil {
		return x.Bic
	}
	return ""
}

func (x *BeneficiaryAccount) GetMasterBankID() string {
	if x != nil {
		return x.MasterBankID
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountAlias() string {
	if x != nil {
		return x.AccountAlias
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountStatus() string {
	if x != nil {
		return x.AccountStatus
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccountCurrency() string {
	if x != nil {
		return x.AccountCurrency
	}
	return ""
}

func (x *BeneficiaryAccount) GetAccessLevel() string {
	if x != nil {
		return x.AccessLevel
	}
	return ""
}

func (x *BeneficiaryAccount) GetIsOwnedByCompany() string {
	if x != nil {
		return x.IsOwnedByCompany
	}
	return ""
}

func (x *BeneficiaryAccount) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *BeneficiaryAccount) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *BeneficiaryAccount) GetDeletedByID() uint64 {
	if x != nil {
		return x.DeletedByID
	}
	return 0
}

func (x *BeneficiaryAccount) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *BeneficiaryAccount) GetMasterBankName() string {
	if x != nil {
		return x.MasterBankName
	}
	return ""
}

func (x *BeneficiaryAccount) GetRoleID() uint64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *BeneficiaryAccount) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *BeneficiaryAccount) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BeneficiaryAccount) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BeneficiaryAccount) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type BeneficiaryBank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeneficiaryBankID uint64 `protobuf:"varint,1,opt,name=beneficiaryBankID,proto3" json:"beneficiaryBankID,omitempty"`
	BankCode          string `protobuf:"bytes,2,opt,name=bankCode,proto3" json:"bankCode,omitempty"`
	BankName          string `protobuf:"bytes,3,opt,name=bankName,proto3" json:"bankName,omitempty"`
	BankDesc          string `protobuf:"bytes,4,opt,name=bankDesc,proto3" json:"bankDesc,omitempty"`
	// string bic = 5 [(gorm.field).tag = {column: "BIC"}, (validator.field) = {msg_exists : true}];
	Country     string                 `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	CreatedByID uint64                 `protobuf:"varint,11,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID uint64                 `protobuf:"varint,12,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *BeneficiaryBank) Reset() {
	*x = BeneficiaryBank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beneficiary_account_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeneficiaryBank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeneficiaryBank) ProtoMessage() {}

func (x *BeneficiaryBank) ProtoReflect() protoreflect.Message {
	mi := &file_beneficiary_account_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeneficiaryBank.ProtoReflect.Descriptor instead.
func (*BeneficiaryBank) Descriptor() ([]byte, []int) {
	return file_beneficiary_account_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *BeneficiaryBank) GetBeneficiaryBankID() uint64 {
	if x != nil {
		return x.BeneficiaryBankID
	}
	return 0
}

func (x *BeneficiaryBank) GetBankCode() string {
	if x != nil {
		return x.BankCode
	}
	return ""
}

func (x *BeneficiaryBank) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *BeneficiaryBank) GetBankDesc() string {
	if x != nil {
		return x.BankDesc
	}
	return ""
}

func (x *BeneficiaryBank) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *BeneficiaryBank) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *BeneficiaryBank) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *BeneficiaryBank) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BeneficiaryBank) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_beneficiary_account_gorm_db_proto protoreflect.FileDescriptor

var file_beneficiary_account_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x13, 0xe0, 0x41, 0x02, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x30, 0x01, 0x40, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x40, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x0d, 0xba, 0xb9, 0x19, 0x09, 0x08, 0x01,
	0x1a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xfd, 0x0b, 0x0a, 0x12, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x59,
	0x0a, 0x14, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x25, 0xe0, 0x41,
	0x03, 0xba, 0xb9, 0x19, 0x1e, 0x0a, 0x1c, 0x0a, 0x14, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x28, 0x01, 0x40,
	0x01, 0x48, 0x01, 0x52, 0x14, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x19, 0xba, 0xb9,
	0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x40,
	0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x44, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x03, 0x62, 0x69, 0x63, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x42, 0x69, 0x63, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x03, 0x62, 0x69, 0x63, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a,
	0xba, 0xb9, 0x19, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61,
	0x6e, 0x6b, 0x49, 0x44, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1d, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x40, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x0a, 0x0c, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x3b, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x0a, 0x0f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0f,
	0x0a, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0xe2,
	0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x4a, 0x0a, 0x10, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xba, 0xb9,
	0x19, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x10, 0x69, 0x73,
	0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3d,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3b, 0x0a, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b,
	0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1c, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e,
	0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x42, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a,
	0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x3a, 0x21, 0xba, 0xb9, 0x19, 0x1d, 0x08, 0x01, 0x1a, 0x19, 0x62, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xed, 0x04, 0x0a, 0x0f, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x50, 0x0a, 0x11, 0x62,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x22, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x1b, 0x0a,
	0x19, 0x0a, 0x11, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x42, 0x61,
	0x6e, 0x6b, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x52, 0x11, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x44, 0x12, 0x32, 0x0a,
	0x08, 0x62, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x16, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64,
	0x65, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x08, 0x42, 0x61, 0x6e,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x62,
	0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xba, 0xb9, 0x19, 0x0e, 0x0a,
	0x0c, 0x0a, 0x08, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x40, 0x01, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x12, 0x2f, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15,
	0xba, 0xb9, 0x19, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0xe2,
	0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3d,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3d, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x4e, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03,
	0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x19, 0xba, 0xb9,
	0x19, 0x15, 0x08, 0x01, 0x1a, 0x11, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72,
	0x79, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beneficiary_account_gorm_db_proto_rawDescOnce sync.Once
	file_beneficiary_account_gorm_db_proto_rawDescData = file_beneficiary_account_gorm_db_proto_rawDesc
)

func file_beneficiary_account_gorm_db_proto_rawDescGZIP() []byte {
	file_beneficiary_account_gorm_db_proto_rawDescOnce.Do(func() {
		file_beneficiary_account_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_beneficiary_account_gorm_db_proto_rawDescData)
	})
	return file_beneficiary_account_gorm_db_proto_rawDescData
}

var file_beneficiary_account_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_beneficiary_account_gorm_db_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: beneficiary.account.service.v1.User
	(*BeneficiaryAccount)(nil),    // 1: beneficiary.account.service.v1.BeneficiaryAccount
	(*BeneficiaryBank)(nil),       // 2: beneficiary.account.service.v1.BeneficiaryBank
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_beneficiary_account_gorm_db_proto_depIdxs = []int32{
	3, // 0: beneficiary.account.service.v1.User.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: beneficiary.account.service.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: beneficiary.account.service.v1.User.deleted_at:type_name -> google.protobuf.Timestamp
	3, // 3: beneficiary.account.service.v1.BeneficiaryAccount.createdAt:type_name -> google.protobuf.Timestamp
	3, // 4: beneficiary.account.service.v1.BeneficiaryAccount.updatedAt:type_name -> google.protobuf.Timestamp
	3, // 5: beneficiary.account.service.v1.BeneficiaryAccount.deletedAt:type_name -> google.protobuf.Timestamp
	3, // 6: beneficiary.account.service.v1.BeneficiaryBank.createdAt:type_name -> google.protobuf.Timestamp
	3, // 7: beneficiary.account.service.v1.BeneficiaryBank.updatedAt:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_beneficiary_account_gorm_db_proto_init() }
func file_beneficiary_account_gorm_db_proto_init() {
	if File_beneficiary_account_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beneficiary_account_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_beneficiary_account_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeneficiaryAccount); i {
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
		file_beneficiary_account_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeneficiaryBank); i {
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
			RawDescriptor: file_beneficiary_account_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beneficiary_account_gorm_db_proto_goTypes,
		DependencyIndexes: file_beneficiary_account_gorm_db_proto_depIdxs,
		MessageInfos:      file_beneficiary_account_gorm_db_proto_msgTypes,
	}.Build()
	File_beneficiary_account_gorm_db_proto = out.File
	file_beneficiary_account_gorm_db_proto_rawDesc = nil
	file_beneficiary_account_gorm_db_proto_goTypes = nil
	file_beneficiary_account_gorm_db_proto_depIdxs = nil
}
