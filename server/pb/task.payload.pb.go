// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: task.payload.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckResponse) GetMessage() string {
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
		mi := &file_task_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[3]
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
	return file_task_payload_proto_rawDescGZIP(), []int{3}
}

type InvalidKeyError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InvalidKeyError) Reset() {
	*x = InvalidKeyError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidKeyError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidKeyError) ProtoMessage() {}

func (x *InvalidKeyError) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidKeyError.ProtoReflect.Descriptor instead.
func (*InvalidKeyError) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{4}
}

func (x *InvalidKeyError) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *InvalidKeyError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SaveTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID uint64 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Task   *Task  `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *SaveTaskRequest) Reset() {
	*x = SaveTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskRequest) ProtoMessage() {}

func (x *SaveTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskRequest.ProtoReflect.Descriptor instead.
func (*SaveTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{5}
}

func (x *SaveTaskRequest) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *SaveTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type SaveTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SaveTaskResponse) Reset() {
	*x = SaveTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskResponse) ProtoMessage() {}

func (x *SaveTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskResponse.ProtoReflect.Descriptor instead.
func (*SaveTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{6}
}

func (x *SaveTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID uint64 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *SetTaskRequest) Reset() {
	*x = SetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskRequest) ProtoMessage() {}

func (x *SetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskRequest.ProtoReflect.Descriptor instead.
func (*SetTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{7}
}

func (x *SetTaskRequest) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *SetTaskRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type SetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Task  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SetTaskResponse) Reset() {
	*x = SetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskResponse) ProtoMessage() {}

func (x *SetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskResponse.ProtoReflect.Descriptor instead.
func (*SetTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{8}
}

func (x *SetTaskResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *SetTaskResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SetTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SetTaskResponse) GetData() *Task {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{9}
}

type ListTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *ListTaskRequest) Reset() {
	*x = ListTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRequest) ProtoMessage() {}

func (x *ListTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRequest.ProtoReflect.Descriptor instead.
func (*ListTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{10}
}

func (x *ListTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type ListTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool    `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Task `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListTaskResponse) Reset() {
	*x = ListTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_payload_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskResponse) ProtoMessage() {}

func (x *ListTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_payload_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskResponse.ProtoReflect.Descriptor instead.
func (*ListTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_payload_proto_rawDescGZIP(), []int{11}
}

func (x *ListTaskResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ListTaskResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListTaskResponse) GetData() []*Task {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_task_payload_proto protoreflect.FileDescriptor

var file_task_payload_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x67, 0x6f, 0x72,
	0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x32, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x2f, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x0a,
	0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4b, 0x65, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x0f,
	0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_payload_proto_rawDescOnce sync.Once
	file_task_payload_proto_rawDescData = file_task_payload_proto_rawDesc
)

func file_task_payload_proto_rawDescGZIP() []byte {
	file_task_payload_proto_rawDescOnce.Do(func() {
		file_task_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_payload_proto_rawDescData)
	})
	return file_task_payload_proto_rawDescData
}

var file_task_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_task_payload_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),        // 0: task.service.v1.LoginRequest
	(*LoginResponse)(nil),       // 1: task.service.v1.LoginResponse
	(*HealthCheckResponse)(nil), // 2: task.service.v1.HealthCheckResponse
	(*Empty)(nil),               // 3: task.service.v1.Empty
	(*InvalidKeyError)(nil),     // 4: task.service.v1.InvalidKeyError
	(*SaveTaskRequest)(nil),     // 5: task.service.v1.SaveTaskRequest
	(*SaveTaskResponse)(nil),    // 6: task.service.v1.SaveTaskResponse
	(*SetTaskRequest)(nil),      // 7: task.service.v1.SetTaskRequest
	(*SetTaskResponse)(nil),     // 8: task.service.v1.SetTaskResponse
	(*ListRequest)(nil),         // 9: task.service.v1.ListRequest
	(*ListTaskRequest)(nil),     // 10: task.service.v1.ListTaskRequest
	(*ListTaskResponse)(nil),    // 11: task.service.v1.ListTaskResponse
	(*Task)(nil),                // 12: task.service.v1.Task
}
var file_task_payload_proto_depIdxs = []int32{
	12, // 0: task.service.v1.SaveTaskRequest.task:type_name -> task.service.v1.Task
	12, // 1: task.service.v1.SetTaskResponse.data:type_name -> task.service.v1.Task
	12, // 2: task.service.v1.ListTaskRequest.task:type_name -> task.service.v1.Task
	12, // 3: task.service.v1.ListTaskResponse.data:type_name -> task.service.v1.Task
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_task_payload_proto_init() }
func file_task_payload_proto_init() {
	if File_task_payload_proto != nil {
		return
	}
	file_task_gorm_db_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_task_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_task_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_task_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_task_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_task_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidKeyError); i {
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
		file_task_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskRequest); i {
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
		file_task_payload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskResponse); i {
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
		file_task_payload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskRequest); i {
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
		file_task_payload_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskResponse); i {
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
		file_task_payload_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_task_payload_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRequest); i {
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
		file_task_payload_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskResponse); i {
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
			RawDescriptor: file_task_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_payload_proto_goTypes,
		DependencyIndexes: file_task_payload_proto_depIdxs,
		MessageInfos:      file_task_payload_proto_msgTypes,
	}.Build()
	File_task_payload_proto = out.File
	file_task_payload_proto_rawDesc = nil
	file_task_payload_proto_goTypes = nil
	file_task_payload_proto_depIdxs = nil
}
