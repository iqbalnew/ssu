// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: announcement.gorm_db.proto

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
		mi := &file_announcement_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_gorm_db_proto_msgTypes[0]
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
	return file_announcement_gorm_db_proto_rawDescGZIP(), []int{0}
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

type Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnnouncementID uint64                 `protobuf:"varint,1,opt,name=announcementID,proto3" json:"announcementID,omitempty"`
	EventTypeID    uint64                 `protobuf:"varint,2,opt,name=eventTypeID,proto3" json:"eventTypeID,omitempty"`
	CompanyID      uint64                 `protobuf:"varint,3,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Description    string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Code           string                 `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Title          string                 `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Content        string                 `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	StatusLevel    string                 `protobuf:"bytes,8,opt,name=statusLevel,proto3" json:"statusLevel,omitempty"`
	IsAutoClose    bool                   `protobuf:"varint,9,opt,name=isAutoClose,proto3" json:"isAutoClose,omitempty"`
	IsEnabled      bool                   `protobuf:"varint,10,opt,name=isEnabled,proto3" json:"isEnabled,omitempty"`
	StartFrom      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=startFrom,proto3" json:"startFrom,omitempty"`
	EndTo          *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=endTo,proto3" json:"endTo,omitempty"`
	CreatedByID    uint64                 `protobuf:"varint,13,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID    uint64                 `protobuf:"varint,14,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	DeletedByID    uint64                 `protobuf:"varint,15,opt,name=deletedByID,proto3" json:"deletedByID,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt      *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	EventType      *AnnouncementEventType `protobuf:"bytes,51,opt,name=EventType,proto3" json:"EventType,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announcement_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_announcement_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Announcement) GetAnnouncementID() uint64 {
	if x != nil {
		return x.AnnouncementID
	}
	return 0
}

func (x *Announcement) GetEventTypeID() uint64 {
	if x != nil {
		return x.EventTypeID
	}
	return 0
}

func (x *Announcement) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Announcement) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Announcement) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Announcement) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Announcement) GetStatusLevel() string {
	if x != nil {
		return x.StatusLevel
	}
	return ""
}

func (x *Announcement) GetIsAutoClose() bool {
	if x != nil {
		return x.IsAutoClose
	}
	return false
}

func (x *Announcement) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *Announcement) GetStartFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.StartFrom
	}
	return nil
}

func (x *Announcement) GetEndTo() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTo
	}
	return nil
}

func (x *Announcement) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Announcement) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *Announcement) GetDeletedByID() uint64 {
	if x != nil {
		return x.DeletedByID
	}
	return 0
}

func (x *Announcement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Announcement) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Announcement) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Announcement) GetEventType() *AnnouncementEventType {
	if x != nil {
		return x.EventType
	}
	return nil
}

type AnnouncementEventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypeID uint64                 `protobuf:"varint,1,opt,name=eventTypeID,proto3" json:"eventTypeID,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,103,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *AnnouncementEventType) Reset() {
	*x = AnnouncementEventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announcement_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementEventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementEventType) ProtoMessage() {}

func (x *AnnouncementEventType) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementEventType.ProtoReflect.Descriptor instead.
func (*AnnouncementEventType) Descriptor() ([]byte, []int) {
	return file_announcement_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *AnnouncementEventType) GetEventTypeID() uint64 {
	if x != nil {
		return x.EventTypeID
	}
	return 0
}

func (x *AnnouncementEventType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnnouncementEventType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AnnouncementEventType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AnnouncementEventType) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_announcement_gorm_db_proto protoreflect.FileDescriptor

var file_announcement_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0,
	0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x13, 0xe0, 0x41, 0x02, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19,
	0x06, 0x0a, 0x04, 0x30, 0x01, 0x40, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x3a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x40, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x0d, 0xba, 0xb9, 0x19, 0x09,
	0x08, 0x01, 0x1a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xd4, 0x08, 0x0a, 0x0c, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40,
	0x01, 0x52, 0x0e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x2f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04,
	0x28, 0x01, 0x40, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12,
	0x33, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x40, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x19, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x12, 0x0c, 0x76,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x40, 0x01, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x19, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x12, 0x0c, 0x76,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x40, 0x01, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10,
	0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x40, 0x01,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19,
	0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68,
	0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x40, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0xe0, 0x41, 0x01,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x45, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x0b, 0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x6e, 0x64, 0x54, 0x6f,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x05, 0x65,
	0x6e, 0x64, 0x54, 0x6f, 0x12, 0x2b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x02, 0x0a, 0x00, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x2b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x02, 0x0a,
	0x00, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2b,
	0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x68, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x1a, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x13, 0x22, 0x11, 0x0a, 0x0d, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x38, 0x01, 0x52, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x08,
	0x01, 0x1a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xec, 0x02, 0x0a, 0x15, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x0b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xe0, 0x41, 0x01, 0xba, 0xb9,
	0x19, 0x12, 0x0a, 0x10, 0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x31, 0x35,
	0x30, 0x29, 0x40, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x09, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x02,
	0x0a, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x20, 0xba,
	0xb9, 0x19, 0x1c, 0x08, 0x01, 0x1a, 0x18, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_announcement_gorm_db_proto_rawDescOnce sync.Once
	file_announcement_gorm_db_proto_rawDescData = file_announcement_gorm_db_proto_rawDesc
)

func file_announcement_gorm_db_proto_rawDescGZIP() []byte {
	file_announcement_gorm_db_proto_rawDescOnce.Do(func() {
		file_announcement_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_announcement_gorm_db_proto_rawDescData)
	})
	return file_announcement_gorm_db_proto_rawDescData
}

var file_announcement_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_announcement_gorm_db_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: announcement.service.v1.User
	(*Announcement)(nil),          // 1: announcement.service.v1.Announcement
	(*AnnouncementEventType)(nil), // 2: announcement.service.v1.AnnouncementEventType
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_announcement_gorm_db_proto_depIdxs = []int32{
	3,  // 0: announcement.service.v1.User.created_at:type_name -> google.protobuf.Timestamp
	3,  // 1: announcement.service.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 2: announcement.service.v1.User.deleted_at:type_name -> google.protobuf.Timestamp
	3,  // 3: announcement.service.v1.Announcement.startFrom:type_name -> google.protobuf.Timestamp
	3,  // 4: announcement.service.v1.Announcement.endTo:type_name -> google.protobuf.Timestamp
	3,  // 5: announcement.service.v1.Announcement.createdAt:type_name -> google.protobuf.Timestamp
	3,  // 6: announcement.service.v1.Announcement.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 7: announcement.service.v1.Announcement.deleted_at:type_name -> google.protobuf.Timestamp
	2,  // 8: announcement.service.v1.Announcement.EventType:type_name -> announcement.service.v1.AnnouncementEventType
	3,  // 9: announcement.service.v1.AnnouncementEventType.createdAt:type_name -> google.protobuf.Timestamp
	3,  // 10: announcement.service.v1.AnnouncementEventType.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 11: announcement.service.v1.AnnouncementEventType.deletedAt:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_announcement_gorm_db_proto_init() }
func file_announcement_gorm_db_proto_init() {
	if File_announcement_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_announcement_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_announcement_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
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
		file_announcement_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementEventType); i {
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
			RawDescriptor: file_announcement_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_announcement_gorm_db_proto_goTypes,
		DependencyIndexes: file_announcement_gorm_db_proto_depIdxs,
		MessageInfos:      file_announcement_gorm_db_proto_msgTypes,
	}.Build()
	File_announcement_gorm_db_proto = out.File
	file_announcement_gorm_db_proto_rawDesc = nil
	file_announcement_gorm_db_proto_goTypes = nil
	file_announcement_gorm_db_proto_depIdxs = nil
}