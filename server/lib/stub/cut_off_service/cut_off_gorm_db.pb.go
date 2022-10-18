// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.16.1
// source: cut_off_gorm_db.proto

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

type CutOff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleID uint64 `protobuf:"varint,1,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`
	// string DEBIT_ACCOUNT = 2 [(gorm.field).tag = {type: 'varchar(100)'}]; // "DEBIT_ACCOUNT":"001901000162530",
	ScheduleName        string                 `protobuf:"bytes,2,opt,name=scheduleName,proto3" json:"scheduleName,omitempty"`
	ScheduleDescription string                 `protobuf:"bytes,3,opt,name=scheduleDescription,proto3" json:"scheduleDescription,omitempty"`
	ModuleID            string                 `protobuf:"bytes,19,opt,name=moduleID,proto3" json:"moduleID,omitempty"`
	ScheduleTime        string                 `protobuf:"bytes,4,opt,name=scheduleTime,proto3" json:"scheduleTime,omitempty"`
	RepeatOn            string                 `protobuf:"bytes,5,opt,name=repeatOn,proto3" json:"repeatOn,omitempty"`
	RepeatPattern       string                 `protobuf:"bytes,6,opt,name=repeatPattern,proto3" json:"repeatPattern,omitempty"`
	Saved               bool                   `protobuf:"varint,7,opt,name=saved,proto3" json:"saved,omitempty"`
	Status              string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedByID         uint64                 `protobuf:"varint,15,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	LastApprovedByID    uint64                 `protobuf:"varint,16,opt,name=lastApprovedByID,proto3" json:"lastApprovedByID,omitempty"`
	LastRejectedByID    uint64                 `protobuf:"varint,17,opt,name=lastRejectedByID,proto3" json:"lastRejectedByID,omitempty"`
	CreatedAt           *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt           *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt           *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *CutOff) Reset() {
	*x = CutOff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cut_off_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CutOff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CutOff) ProtoMessage() {}

func (x *CutOff) ProtoReflect() protoreflect.Message {
	mi := &file_cut_off_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CutOff.ProtoReflect.Descriptor instead.
func (*CutOff) Descriptor() ([]byte, []int) {
	return file_cut_off_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *CutOff) GetScheduleID() uint64 {
	if x != nil {
		return x.ScheduleID
	}
	return 0
}

func (x *CutOff) GetScheduleName() string {
	if x != nil {
		return x.ScheduleName
	}
	return ""
}

func (x *CutOff) GetScheduleDescription() string {
	if x != nil {
		return x.ScheduleDescription
	}
	return ""
}

func (x *CutOff) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *CutOff) GetScheduleTime() string {
	if x != nil {
		return x.ScheduleTime
	}
	return ""
}

func (x *CutOff) GetRepeatOn() string {
	if x != nil {
		return x.RepeatOn
	}
	return ""
}

func (x *CutOff) GetRepeatPattern() string {
	if x != nil {
		return x.RepeatPattern
	}
	return ""
}

func (x *CutOff) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

func (x *CutOff) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CutOff) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *CutOff) GetLastApprovedByID() uint64 {
	if x != nil {
		return x.LastApprovedByID
	}
	return 0
}

func (x *CutOff) GetLastRejectedByID() uint64 {
	if x != nil {
		return x.LastRejectedByID
	}
	return 0
}

func (x *CutOff) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CutOff) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CutOff) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Holiday struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleID          uint64                 `protobuf:"varint,1,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`
	ScheduleName        string                 `protobuf:"bytes,2,opt,name=scheduleName,proto3" json:"scheduleName,omitempty"`
	ScheduleDescription string                 `protobuf:"bytes,3,opt,name=scheduleDescription,proto3" json:"scheduleDescription,omitempty"`
	Module              string                 `protobuf:"bytes,4,opt,name=module,proto3" json:"module,omitempty"`
	ScheduleDate        string                 `protobuf:"bytes,5,opt,name=scheduleDate,proto3" json:"scheduleDate,omitempty"`
	Saved               bool                   `protobuf:"varint,6,opt,name=saved,proto3" json:"saved,omitempty"`
	Status              string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	TaskID              uint64                 `protobuf:"varint,8,opt,name=taskID,proto3" json:"taskID,omitempty"`
	CreatedByID         uint64                 `protobuf:"varint,15,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	LastApprovedByID    uint64                 `protobuf:"varint,16,opt,name=lastApprovedByID,proto3" json:"lastApprovedByID,omitempty"`
	LastRejectedByID    uint64                 `protobuf:"varint,17,opt,name=lastRejectedByID,proto3" json:"lastRejectedByID,omitempty"`
	CreatedAt           *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt           *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt           *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Holiday) Reset() {
	*x = Holiday{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cut_off_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Holiday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holiday) ProtoMessage() {}

func (x *Holiday) ProtoReflect() protoreflect.Message {
	mi := &file_cut_off_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Holiday.ProtoReflect.Descriptor instead.
func (*Holiday) Descriptor() ([]byte, []int) {
	return file_cut_off_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Holiday) GetScheduleID() uint64 {
	if x != nil {
		return x.ScheduleID
	}
	return 0
}

func (x *Holiday) GetScheduleName() string {
	if x != nil {
		return x.ScheduleName
	}
	return ""
}

func (x *Holiday) GetScheduleDescription() string {
	if x != nil {
		return x.ScheduleDescription
	}
	return ""
}

func (x *Holiday) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Holiday) GetScheduleDate() string {
	if x != nil {
		return x.ScheduleDate
	}
	return ""
}

func (x *Holiday) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

func (x *Holiday) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Holiday) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *Holiday) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Holiday) GetLastApprovedByID() uint64 {
	if x != nil {
		return x.LastApprovedByID
	}
	return 0
}

func (x *Holiday) GetLastRejectedByID() uint64 {
	if x != nil {
		return x.LastRejectedByID
	}
	return 0
}

func (x *Holiday) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Holiday) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Holiday) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Calendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date             *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Day              string                 `protobuf:"bytes,3,opt,name=day,proto3" json:"day,omitempty"`
	IsHoliday        bool                   `protobuf:"varint,4,opt,name=isHoliday,proto3" json:"isHoliday,omitempty"`
	Notes            string                 `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	CreatedByID      uint64                 `protobuf:"varint,15,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	LastApprovedByID uint64                 `protobuf:"varint,16,opt,name=lastApprovedByID,proto3" json:"lastApprovedByID,omitempty"`
	LastRejectedByID uint64                 `protobuf:"varint,17,opt,name=lastRejectedByID,proto3" json:"lastRejectedByID,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt        *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Calendar) Reset() {
	*x = Calendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cut_off_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calendar) ProtoMessage() {}

func (x *Calendar) ProtoReflect() protoreflect.Message {
	mi := &file_cut_off_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calendar.ProtoReflect.Descriptor instead.
func (*Calendar) Descriptor() ([]byte, []int) {
	return file_cut_off_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *Calendar) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Calendar) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Calendar) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Calendar) GetIsHoliday() bool {
	if x != nil {
		return x.IsHoliday
	}
	return false
}

func (x *Calendar) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Calendar) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Calendar) GetLastApprovedByID() uint64 {
	if x != nil {
		return x.LastApprovedByID
	}
	return 0
}

func (x *Calendar) GetLastRejectedByID() uint64 {
	if x != nil {
		return x.LastRejectedByID
	}
	return 0
}

func (x *Calendar) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Calendar) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Calendar) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_cut_off_gorm_db_proto protoreflect.FileDescriptor

var file_cut_off_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x75, 0x74, 0x2e, 0x6f, 0x66, 0x66,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4,
	0x07, 0x0a, 0x06, 0x43, 0x75, 0x74, 0x4f, 0x66, 0x66, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x19, 0xe0,
	0x41, 0x03, 0xba, 0xb9, 0x19, 0x12, 0x0a, 0x10, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xba, 0xb9, 0x19, 0x10,
	0x0a, 0x0e, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x21, 0xba, 0xb9, 0x19, 0x17, 0x0a, 0x15, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x13, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xba, 0xb9, 0x19,
	0x12, 0x0a, 0x10, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x49, 0x44, 0x12, 0x05, 0x6a, 0x73,
	0x6f, 0x6e, 0x62, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x45, 0x0a,
	0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x21, 0xba, 0xb9, 0x19, 0x17, 0x0a, 0x15, 0x0a, 0x0c, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x4f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x08,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x4f, 0x6e, 0x12, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0xe2,
	0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x4f, 0x6e, 0x12,
	0x41, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x53, 0x61, 0x76, 0x65, 0x64,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba,
	0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x1a, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0d,
	0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x10, 0x6c,
	0x61, 0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02,
	0x01, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x08, 0x01, 0x1a, 0x06, 0x43,
	0x75, 0x74, 0x4f, 0x66, 0x66, 0x22, 0xbd, 0x06, 0x0a, 0x07, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04,
	0x28, 0x01, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x39, 0x0a,
	0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x13, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0xe2,
	0xdf, 0x1f, 0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x13, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0e, 0xba, 0xb9,
	0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x05, 0x73, 0x61,
	0x76, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1a, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
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
	0x65, 0x64, 0x41, 0x74, 0x3a, 0x0f, 0xba, 0xb9, 0x19, 0x0b, 0x08, 0x01, 0x1a, 0x07, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x22, 0xbe, 0x05, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a,
	0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0xe2, 0xdf, 0x1f,
	0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02,
	0x01, 0x07, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x69, 0x73, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x04,
	0x0a, 0x02, 0x40, 0x01, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x09, 0x69, 0x73, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x92, 0x41, 0x04, 0x9a,
	0x02, 0x01, 0x07, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1a, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x52, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x4e, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
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
	0x65, 0x64, 0x41, 0x74, 0x3a, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x08, 0x01, 0x1a, 0x08, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cut_off_gorm_db_proto_rawDescOnce sync.Once
	file_cut_off_gorm_db_proto_rawDescData = file_cut_off_gorm_db_proto_rawDesc
)

func file_cut_off_gorm_db_proto_rawDescGZIP() []byte {
	file_cut_off_gorm_db_proto_rawDescOnce.Do(func() {
		file_cut_off_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_cut_off_gorm_db_proto_rawDescData)
	})
	return file_cut_off_gorm_db_proto_rawDescData
}

var file_cut_off_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cut_off_gorm_db_proto_goTypes = []interface{}{
	(*CutOff)(nil),                // 0: cut.off.service.v1.CutOff
	(*Holiday)(nil),               // 1: cut.off.service.v1.Holiday
	(*Calendar)(nil),              // 2: cut.off.service.v1.Calendar
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_cut_off_gorm_db_proto_depIdxs = []int32{
	3,  // 0: cut.off.service.v1.CutOff.createdAt:type_name -> google.protobuf.Timestamp
	3,  // 1: cut.off.service.v1.CutOff.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 2: cut.off.service.v1.CutOff.deletedAt:type_name -> google.protobuf.Timestamp
	3,  // 3: cut.off.service.v1.Holiday.createdAt:type_name -> google.protobuf.Timestamp
	3,  // 4: cut.off.service.v1.Holiday.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 5: cut.off.service.v1.Holiday.deletedAt:type_name -> google.protobuf.Timestamp
	3,  // 6: cut.off.service.v1.Calendar.date:type_name -> google.protobuf.Timestamp
	3,  // 7: cut.off.service.v1.Calendar.createdAt:type_name -> google.protobuf.Timestamp
	3,  // 8: cut.off.service.v1.Calendar.updatedAt:type_name -> google.protobuf.Timestamp
	3,  // 9: cut.off.service.v1.Calendar.deletedAt:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cut_off_gorm_db_proto_init() }
func file_cut_off_gorm_db_proto_init() {
	if File_cut_off_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cut_off_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CutOff); i {
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
		file_cut_off_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Holiday); i {
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
		file_cut_off_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calendar); i {
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
			RawDescriptor: file_cut_off_gorm_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cut_off_gorm_db_proto_goTypes,
		DependencyIndexes: file_cut_off_gorm_db_proto_depIdxs,
		MessageInfos:      file_cut_off_gorm_db_proto_msgTypes,
	}.Build()
	File_cut_off_gorm_db_proto = out.File
	file_cut_off_gorm_db_proto_rawDesc = nil
	file_cut_off_gorm_db_proto_goTypes = nil
	file_cut_off_gorm_db_proto_depIdxs = nil
}
