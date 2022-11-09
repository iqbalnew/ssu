// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: task.payload_ev.proto

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

type ListTaskRequestEVDirection int32

const (
	ListTaskRequestEV_DESC ListTaskRequestEVDirection = 0
	ListTaskRequestEV_ASC  ListTaskRequestEVDirection = 1
)

// Enum value maps for ListTaskRequestEVDirection.
var (
	ListTaskRequestEVDirection_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	ListTaskRequestEVDirection_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x ListTaskRequestEVDirection) Enum() *ListTaskRequestEVDirection {
	p := new(ListTaskRequestEVDirection)
	*p = x
	return p
}

func (x ListTaskRequestEVDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListTaskRequestEVDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_task_payload_ev_proto_enumTypes[0].Descriptor()
}

func (ListTaskRequestEVDirection) Type() protoreflect.EnumType {
	return &file_task_payload_ev_proto_enumTypes[0]
}

func (x ListTaskRequestEVDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListTaskRequestEVDirection.Descriptor instead.
func (ListTaskRequestEVDirection) EnumDescriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{4, 0}
}

type SaveTaskRequestEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID            string  `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Task              *TaskEV `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	IsDraft           bool    `protobuf:"varint,2,opt,name=isDraft,proto3" json:"isDraft,omitempty"`
	TransactionAmount float64 `protobuf:"fixed64,4,opt,name=transactionAmount,proto3" json:"transactionAmount,omitempty"`
}

func (x *SaveTaskRequestEV) Reset() {
	*x = SaveTaskRequestEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskRequestEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskRequestEV) ProtoMessage() {}

func (x *SaveTaskRequestEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskRequestEV.ProtoReflect.Descriptor instead.
func (*SaveTaskRequestEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{0}
}

func (x *SaveTaskRequestEV) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *SaveTaskRequestEV) GetTask() *TaskEV {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *SaveTaskRequestEV) GetIsDraft() bool {
	if x != nil {
		return x.IsDraft
	}
	return false
}

func (x *SaveTaskRequestEV) GetTransactionAmount() float64 {
	if x != nil {
		return x.TransactionAmount
	}
	return 0
}

type SaveTaskResponseEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Data    *TaskEV `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveTaskResponseEV) Reset() {
	*x = SaveTaskResponseEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskResponseEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskResponseEV) ProtoMessage() {}

func (x *SaveTaskResponseEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskResponseEV.ProtoReflect.Descriptor instead.
func (*SaveTaskResponseEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{1}
}

func (x *SaveTaskResponseEV) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SaveTaskResponseEV) GetData() *TaskEV {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetTaskRequestEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID  string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Action  string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Reasons string `protobuf:"bytes,4,opt,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *SetTaskRequestEV) Reset() {
	*x = SetTaskRequestEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskRequestEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskRequestEV) ProtoMessage() {}

func (x *SetTaskRequestEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskRequestEV.ProtoReflect.Descriptor instead.
func (*SetTaskRequestEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{2}
}

func (x *SetTaskRequestEV) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *SetTaskRequestEV) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SetTaskRequestEV) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SetTaskRequestEV) GetReasons() string {
	if x != nil {
		return x.Reasons
	}
	return ""
}

type SetTaskResponseEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool    `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *TaskEV `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SetTaskResponseEV) Reset() {
	*x = SetTaskResponseEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskResponseEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskResponseEV) ProtoMessage() {}

func (x *SetTaskResponseEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskResponseEV.ProtoReflect.Descriptor instead.
func (*SetTaskResponseEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{3}
}

func (x *SetTaskResponseEV) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *SetTaskResponseEV) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SetTaskResponseEV) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SetTaskResponseEV) GetData() *TaskEV {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListTaskRequestEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *TaskEV                    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Limit    int32                      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page     int32                      `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Sort     string                     `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Dir      ListTaskRequestEVDirection `protobuf:"varint,5,opt,name=dir,proto3,enum=task.service.v1.ListTaskRequestEVDirection" json:"dir,omitempty"`
	Filter   string                     `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
	Query    string                     `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	In       string                     `protobuf:"bytes,8,opt,name=in,proto3" json:"in,omitempty"`
	FilterOr string                     `protobuf:"bytes,9,opt,name=filterOr,proto3" json:"filterOr,omitempty"`
}

func (x *ListTaskRequestEV) Reset() {
	*x = ListTaskRequestEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRequestEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRequestEV) ProtoMessage() {}

func (x *ListTaskRequestEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRequestEV.ProtoReflect.Descriptor instead.
func (*ListTaskRequestEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{4}
}

func (x *ListTaskRequestEV) GetTask() *TaskEV {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ListTaskRequestEV) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTaskRequestEV) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTaskRequestEV) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListTaskRequestEV) GetDir() ListTaskRequestEVDirection {
	if x != nil {
		return x.Dir
	}
	return ListTaskRequestEV_DESC
}

func (x *ListTaskRequestEV) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListTaskRequestEV) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListTaskRequestEV) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *ListTaskRequestEV) GetFilterOr() string {
	if x != nil {
		return x.FilterOr
	}
	return ""
}

type ListTaskResponseEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      bool                `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code       uint32              `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message    string              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data       []*TaskEV           `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,5,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListTaskResponseEV) Reset() {
	*x = ListTaskResponseEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskResponseEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskResponseEV) ProtoMessage() {}

func (x *ListTaskResponseEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskResponseEV.ProtoReflect.Descriptor instead.
func (*ListTaskResponseEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{5}
}

func (x *ListTaskResponseEV) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ListTaskResponseEV) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListTaskResponseEV) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListTaskResponseEV) GetData() []*TaskEV {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListTaskResponseEV) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type AssignaTypeIDRequestEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID    string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	FeatureID uint64 `protobuf:"varint,2,opt,name=featureID,proto3" json:"featureID,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *AssignaTypeIDRequestEV) Reset() {
	*x = AssignaTypeIDRequestEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignaTypeIDRequestEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignaTypeIDRequestEV) ProtoMessage() {}

func (x *AssignaTypeIDRequestEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignaTypeIDRequestEV.ProtoReflect.Descriptor instead.
func (*AssignaTypeIDRequestEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{6}
}

func (x *AssignaTypeIDRequestEV) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *AssignaTypeIDRequestEV) GetFeatureID() uint64 {
	if x != nil {
		return x.FeatureID
	}
	return 0
}

func (x *AssignaTypeIDRequestEV) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetTaskByTypeIDReqEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ID   string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetTaskByTypeIDReqEV) Reset() {
	*x = GetTaskByTypeIDReqEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByTypeIDReqEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByTypeIDReqEV) ProtoMessage() {}

func (x *GetTaskByTypeIDReqEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByTypeIDReqEV.ProtoReflect.Descriptor instead.
func (*GetTaskByTypeIDReqEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{7}
}

func (x *GetTaskByTypeIDReqEV) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetTaskByTypeIDReqEV) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetTaskByTypeIDResEV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool    `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Data  *TaskEV `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTaskByTypeIDResEV) Reset() {
	*x = GetTaskByTypeIDResEV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_ev_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByTypeIDResEV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByTypeIDResEV) ProtoMessage() {}

func (x *GetTaskByTypeIDResEV) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_ev_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByTypeIDResEV.ProtoReflect.Descriptor instead.
func (*GetTaskByTypeIDResEV) Descriptor() ([]byte, []int) {
	return file_task_payload_ev_proto_rawDescGZIP(), []int{8}
}

func (x *GetTaskByTypeIDResEV) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *GetTaskByTypeIDResEV) GetData() *TaskEV {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_task_payload_ev_proto protoreflect.FileDescriptor

var file_task_payload_ev_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65,
	0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x12, 0x53, 0x61, 0x76,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x56, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0x92, 0x41, 0x27, 0x32, 0x25, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x20, 0x27, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x27, 0x2c, 0x20, 0x27, 0x72, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x27, 0x2c, 0x20, 0x27, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x27, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x56, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xc5, 0x0a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x54, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x3e, 0x92, 0x41, 0x3b, 0x32, 0x39, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x70, 0x65, 0x72, 0x20, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x20,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x27, 0x2d,
	0x31, 0x27, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x55, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x41, 0x92, 0x41, 0x3e, 0x32, 0x3c,
	0x70, 0x61, 0x67, 0x65, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x20, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x27, 0x2d, 0x31, 0x27, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x20, 0x62, 0x79, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x54, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x56, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x14, 0x92, 0x41, 0x11, 0x32, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x20, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x8e, 0x04, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0xf5, 0x03,
	0x92, 0x41, 0xf1, 0x03, 0x32, 0xee, 0x03, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x27, 0x41,
	0x4e, 0x44, 0x27, 0x20, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x20, 0x0a, 0x20, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x3a, 0x20, 0x5b, 0x6b, 0x65, 0x79, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2c, 0x6b, 0x65, 0x79, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5d, 0x20, 0x0a, 0x20, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3d, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x3a, 0x31, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x3d, 0x3e, 0x20, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x20, 0x3d, 0x20, 0x31, 0x20, 0x41,
	0x4e, 0x44, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x3d, 0x20, 0x27, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x27, 0x20, 0x0a, 0x20, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79,
	0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20,
	0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65,
	0x79, 0x3a, 0x25, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b,
	0x65, 0x79, 0x20, 0x4c, 0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27,
	0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x21, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x49, 0x4c, 0x49, 0x4b, 0x45,
	0x20, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b,
	0x65, 0x79, 0x3a, 0x3e, 0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27,
	0x6b, 0x65, 0x79, 0x20, 0x3e, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20,
	0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x3c, 0x3d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20,
	0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x3c, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x3a, 0x3e, 0x3d, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x31, 0x2c, 0x6b, 0x65, 0x79, 0x32, 0x3a, 0x3c, 0x3d, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x20, 0x3e, 0x3d, 0x20,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x20, 0x41, 0x4e, 0x44, 0x20, 0x6b, 0x65, 0x79, 0x32, 0x20,
	0x3c, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x27, 0x20, 0x42, 0x45, 0x54, 0x57, 0x45,
	0x45, 0x4e, 0x20, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27,
	0x6b, 0x65, 0x79, 0x31, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x3a,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x2d,
	0x3e, 0x3e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x20, 0x3d, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x27, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0xee, 0x02,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0xd7, 0x02,
	0x92, 0x41, 0xd3, 0x02, 0x32, 0xd0, 0x02, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x27, 0x4f,
	0x52, 0x27, 0x20, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x20, 0x0a, 0x20, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x3a, 0x20, 0x5b, 0x6b, 0x65, 0x79, 0x31, 0x2c, 0x6b, 0x65, 0x79, 0x32, 0x2c,
	0x6b, 0x65, 0x79, 0x33, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5d, 0x20, 0x0a, 0x20, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3d, 0x6e, 0x61, 0x6d,
	0x65, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x3d, 0x3e, 0x20, 0x57, 0x68, 0x65, 0x72, 0x65, 0x20, 0x6e, 0x61, 0x6d,
	0x65, 0x20, 0x3d, 0x20, 0x27, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x27, 0x20, 0x4f, 0x52, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x3d, 0x20, 0x27, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x27, 0x20, 0x0a, 0x20, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27,
	0x6b, 0x65, 0x79, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b,
	0x65, 0x79, 0x20, 0x3d, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20,
	0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e,
	0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x4c, 0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x3a, 0x25, 0x21, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79, 0x20, 0x49, 0x4c,
	0x49, 0x4b, 0x45, 0x20, 0x25, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x25, 0x27, 0x20, 0x0a, 0x20, 0x2d,
	0x20, 0x27, 0x6b, 0x65, 0x79, 0x31, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79,
	0x31, 0x3a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x20, 0x3d, 0x3e, 0x20, 0x27, 0x6b, 0x65, 0x79,
	0x31, 0x2d, 0x3e, 0x3e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x31, 0x20, 0x3d,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x27, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x22, 0x1e, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x56, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x16, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x56, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x45, 0x56, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x45, 0x56, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x56, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_task_payload_ev_proto_rawDescOnce sync.Once
	file_task_payload_ev_proto_rawDescData = file_task_payload_ev_proto_rawDesc
)

func file_task_payload_ev_proto_rawDescGZIP() []byte {
	file_task_payload_ev_proto_rawDescOnce.Do(func() {
		file_task_payload_ev_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_payload_ev_proto_rawDescData)
	})
	return file_task_payload_ev_proto_rawDescData
}

var file_task_payload_ev_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_payload_ev_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_task_payload_ev_proto_goTypes = []interface{}{
	(ListTaskRequestEVDirection)(0), // 0: task.service.v1.ListTaskRequestEV.direction
	(*SaveTaskRequestEV)(nil),       // 1: task.service.v1.SaveTaskRequestEV
	(*SaveTaskResponseEV)(nil),      // 2: task.service.v1.SaveTaskResponseEV
	(*SetTaskRequestEV)(nil),        // 3: task.service.v1.SetTaskRequestEV
	(*SetTaskResponseEV)(nil),       // 4: task.service.v1.SetTaskResponseEV
	(*ListTaskRequestEV)(nil),       // 5: task.service.v1.ListTaskRequestEV
	(*ListTaskResponseEV)(nil),      // 6: task.service.v1.ListTaskResponseEV
	(*AssignaTypeIDRequestEV)(nil),  // 7: task.service.v1.AssignaTypeIDRequestEV
	(*GetTaskByTypeIDReqEV)(nil),    // 8: task.service.v1.GetTaskByTypeIDReqEV
	(*GetTaskByTypeIDResEV)(nil),    // 9: task.service.v1.GetTaskByTypeIDResEV
	(*TaskEV)(nil),                  // 10: task.service.v1.TaskEV
	(*PaginationResponse)(nil),      // 11: task.service.v1.PaginationResponse
}
var file_task_payload_ev_proto_depIdxs = []int32{
	10, // 0: task.service.v1.SaveTaskRequestEV.task:type_name -> task.service.v1.TaskEV
	10, // 1: task.service.v1.SaveTaskResponseEV.data:type_name -> task.service.v1.TaskEV
	10, // 2: task.service.v1.SetTaskResponseEV.data:type_name -> task.service.v1.TaskEV
	10, // 3: task.service.v1.ListTaskRequestEV.task:type_name -> task.service.v1.TaskEV
	0,  // 4: task.service.v1.ListTaskRequestEV.dir:type_name -> task.service.v1.ListTaskRequestEV.direction
	10, // 5: task.service.v1.ListTaskResponseEV.data:type_name -> task.service.v1.TaskEV
	11, // 6: task.service.v1.ListTaskResponseEV.pagination:type_name -> task.service.v1.PaginationResponse
	10, // 7: task.service.v1.GetTaskByTypeIDResEV.data:type_name -> task.service.v1.TaskEV
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_task_payload_ev_proto_init() }
func file_task_payload_ev_proto_init() {
	if File_task_payload_ev_proto != nil {
		return
	}
	file_task_gorm_db_proto_init()
	file_task_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_task_payload_ev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskRequestEV); i {
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
		file_task_payload_ev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskResponseEV); i {
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
		file_task_payload_ev_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskRequestEV); i {
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
		file_task_payload_ev_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskResponseEV); i {
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
		file_task_payload_ev_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRequestEV); i {
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
		file_task_payload_ev_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskResponseEV); i {
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
		file_task_payload_ev_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignaTypeIDRequestEV); i {
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
		file_task_payload_ev_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskByTypeIDReqEV); i {
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
		file_task_payload_ev_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskByTypeIDResEV); i {
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
			RawDescriptor: file_task_payload_ev_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_payload_ev_proto_goTypes,
		DependencyIndexes: file_task_payload_ev_proto_depIdxs,
		EnumInfos:         file_task_payload_ev_proto_enumTypes,
		MessageInfos:      file_task_payload_ev_proto_msgTypes,
	}.Build()
	File_task_payload_ev_proto = out.File
	file_task_payload_ev_proto_rawDesc = nil
	file_task_payload_ev_proto_goTypes = nil
	file_task_payload_ev_proto_depIdxs = nil
}
