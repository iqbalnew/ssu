// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: workflow_gorm_db.proto

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

type WorkflowRequirementSteps int32

const (
	WorkflowRequirement_maker    WorkflowRequirementSteps = 0
	WorkflowRequirement_checker  WorkflowRequirementSteps = 1
	WorkflowRequirement_signer   WorkflowRequirementSteps = 2
	WorkflowRequirement_releaser WorkflowRequirementSteps = 3
)

// Enum value maps for WorkflowRequirementSteps.
var (
	WorkflowRequirementSteps_name = map[int32]string{
		0: "maker",
		1: "checker",
		2: "signer",
		3: "releaser",
	}
	WorkflowRequirementSteps_value = map[string]int32{
		"maker":    0,
		"checker":  1,
		"signer":   2,
		"releaser": 3,
	}
)

func (x WorkflowRequirementSteps) Enum() *WorkflowRequirementSteps {
	p := new(WorkflowRequirementSteps)
	*p = x
	return p
}

func (x WorkflowRequirementSteps) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowRequirementSteps) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_gorm_db_proto_enumTypes[0].Descriptor()
}

func (WorkflowRequirementSteps) Type() protoreflect.EnumType {
	return &file_workflow_gorm_db_proto_enumTypes[0]
}

func (x WorkflowRequirementSteps) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowRequirementSteps.Descriptor instead.
func (WorkflowRequirementSteps) EnumDescriptor() ([]byte, []int) {
	return file_workflow_gorm_db_proto_rawDescGZIP(), []int{3, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Workflow  string                 `protobuf:"bytes,4,opt,name=workflow,proto3" json:"workflow,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_gorm_db_proto_msgTypes[0]
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
	return file_workflow_gorm_db_proto_rawDescGZIP(), []int{0}
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

func (x *User) GetWorkflow() string {
	if x != nil {
		return x.Workflow
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

type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowID   uint64                 `protobuf:"varint,1,opt,name=workflowID,proto3" json:"workflowID,omitempty"`
	ModuleID     uint64                 `protobuf:"varint,2,opt,name=moduleID,proto3" json:"moduleID,omitempty"`
	CompanyID    uint64                 `protobuf:"varint,3,opt,name=companyID,proto3" json:"companyID,omitempty"`
	AccountID    uint64                 `protobuf:"varint,4,opt,name=accountID,proto3" json:"accountID,omitempty"`
	CreatedByID  uint64                 `protobuf:"varint,5,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID  uint64                 `protobuf:"varint,6,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	DeletedByID  uint64                 `protobuf:"varint,7,opt,name=deletedByID,proto3" json:"deletedByID,omitempty"`
	AccountAlias string                 `protobuf:"bytes,8,opt,name=accountAlias,proto3" json:"accountAlias,omitempty"`
	WorkflowCode string                 `protobuf:"bytes,9,opt,name=workflowCode,proto3" json:"workflowCode,omitempty"`
	Description  string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,53,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	Logics       []*WorkflowLogic       `protobuf:"bytes,101,rep,name=logics,proto3" json:"logics,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_gorm_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_gorm_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_workflow_gorm_db_proto_rawDescGZIP(), []int{1}
}

func (x *Workflow) GetWorkflowID() uint64 {
	if x != nil {
		return x.WorkflowID
	}
	return 0
}

func (x *Workflow) GetModuleID() uint64 {
	if x != nil {
		return x.ModuleID
	}
	return 0
}

func (x *Workflow) GetCompanyID() uint64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Workflow) GetAccountID() uint64 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *Workflow) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *Workflow) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *Workflow) GetDeletedByID() uint64 {
	if x != nil {
		return x.DeletedByID
	}
	return 0
}

func (x *Workflow) GetAccountAlias() string {
	if x != nil {
		return x.AccountAlias
	}
	return ""
}

func (x *Workflow) GetWorkflowCode() string {
	if x != nil {
		return x.WorkflowCode
	}
	return ""
}

func (x *Workflow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Workflow) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Workflow) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Workflow) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Workflow) GetLogics() []*WorkflowLogic {
	if x != nil {
		return x.Logics
	}
	return nil
}

type WorkflowLogic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowLogicID uint64                 `protobuf:"varint,1,opt,name=workflowLogicID,proto3" json:"workflowLogicID,omitempty"`
	WorkflowID      uint64                 `protobuf:"varint,2,opt,name=workflowID,proto3" json:"workflowID,omitempty"`
	BottomRange     uint64                 `protobuf:"varint,3,opt,name=bottomRange,proto3" json:"bottomRange,omitempty"`
	TopRange        uint64                 `protobuf:"varint,4,opt,name=topRange,proto3" json:"topRange,omitempty"`
	CreatedByID     uint64                 `protobuf:"varint,5,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID     uint64                 `protobuf:"varint,6,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Requirements    []*WorkflowRequirement `protobuf:"bytes,101,rep,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *WorkflowLogic) Reset() {
	*x = WorkflowLogic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_gorm_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowLogic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowLogic) ProtoMessage() {}

func (x *WorkflowLogic) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_gorm_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowLogic.ProtoReflect.Descriptor instead.
func (*WorkflowLogic) Descriptor() ([]byte, []int) {
	return file_workflow_gorm_db_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowLogic) GetWorkflowLogicID() uint64 {
	if x != nil {
		return x.WorkflowLogicID
	}
	return 0
}

func (x *WorkflowLogic) GetWorkflowID() uint64 {
	if x != nil {
		return x.WorkflowID
	}
	return 0
}

func (x *WorkflowLogic) GetBottomRange() uint64 {
	if x != nil {
		return x.BottomRange
	}
	return 0
}

func (x *WorkflowLogic) GetTopRange() uint64 {
	if x != nil {
		return x.TopRange
	}
	return 0
}

func (x *WorkflowLogic) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *WorkflowLogic) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *WorkflowLogic) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WorkflowLogic) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *WorkflowLogic) GetRequirements() []*WorkflowRequirement {
	if x != nil {
		return x.Requirements
	}
	return nil
}

type WorkflowRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRequirementID uint64                 `protobuf:"varint,1,opt,name=workflowRequirementID,proto3" json:"workflowRequirementID,omitempty"`
	WorkflowLogicID       uint64                 `protobuf:"varint,2,opt,name=workflowLogicID,proto3" json:"workflowLogicID,omitempty"`
	Step                  string                 `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	RoleID                uint64                 `protobuf:"varint,4,opt,name=roleID,proto3" json:"roleID,omitempty"`
	Priority              uint32                 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	MinimumNumber         uint32                 `protobuf:"varint,6,opt,name=minimumNumber,proto3" json:"minimumNumber,omitempty"`
	CreatedByID           uint64                 `protobuf:"varint,7,opt,name=createdByID,proto3" json:"createdByID,omitempty"`
	UpdatedByID           uint64                 `protobuf:"varint,8,opt,name=updatedByID,proto3" json:"updatedByID,omitempty"`
	CreatedAt             *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt             *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *WorkflowRequirement) Reset() {
	*x = WorkflowRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_gorm_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRequirement) ProtoMessage() {}

func (x *WorkflowRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_gorm_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRequirement.ProtoReflect.Descriptor instead.
func (*WorkflowRequirement) Descriptor() ([]byte, []int) {
	return file_workflow_gorm_db_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowRequirement) GetWorkflowRequirementID() uint64 {
	if x != nil {
		return x.WorkflowRequirementID
	}
	return 0
}

func (x *WorkflowRequirement) GetWorkflowLogicID() uint64 {
	if x != nil {
		return x.WorkflowLogicID
	}
	return 0
}

func (x *WorkflowRequirement) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

func (x *WorkflowRequirement) GetRoleID() uint64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *WorkflowRequirement) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *WorkflowRequirement) GetMinimumNumber() uint32 {
	if x != nil {
		return x.MinimumNumber
	}
	return 0
}

func (x *WorkflowRequirement) GetCreatedByID() uint64 {
	if x != nil {
		return x.CreatedByID
	}
	return 0
}

func (x *WorkflowRequirement) GetUpdatedByID() uint64 {
	if x != nil {
		return x.UpdatedByID
	}
	return 0
}

func (x *WorkflowRequirement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WorkflowRequirement) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_workflow_gorm_db_proto protoreflect.FileDescriptor

var file_workflow_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x81,
	0x08, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x42, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x22, 0xba, 0xb9, 0x19, 0x14, 0x0a, 0x12, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x12,
	0x35, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a, 0xba, 0xb9, 0x19, 0x0f, 0x0a,
	0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41,
	0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x12, 0x36, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x18, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xba,
	0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xba,
	0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1d, 0xba,
	0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x44, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x46, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x22, 0xba, 0xb9, 0x19, 0x1e, 0x0a, 0x1c, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x0c, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32,
	0x35, 0x35, 0x29, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xba, 0xb9, 0x19, 0x20, 0x0a, 0x1e, 0x0a,
	0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x76,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x40, 0x01, 0x92, 0x41, 0x04,
	0x9a, 0x02, 0x01, 0x07, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xba, 0xb9, 0x19, 0x1d, 0x0a, 0x1b, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x76, 0x61,
	0x72, 0x63, 0x68, 0x61, 0x72, 0x28, 0x32, 0x35, 0x35, 0x29, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a,
	0x0d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x40, 0x01, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19,
	0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x40, 0x01,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x62, 0x0a, 0x06, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x42,
	0x26, 0xba, 0xb9, 0x19, 0x22, 0x2a, 0x20, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x44, 0x40, 0x01, 0x48, 0x01, 0x58, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x73, 0x3a,
	0x11, 0xba, 0xb9, 0x19, 0x0d, 0x08, 0x01, 0x1a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x22, 0xd6, 0x05, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x12, 0x51, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x27, 0xba,
	0xb9, 0x19, 0x19, 0x0a, 0x17, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x92, 0x41, 0x04, 0x9a,
	0x02, 0x01, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1e, 0xba, 0xb9, 0x19,
	0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x40,
	0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9,
	0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0b, 0x62, 0x6f, 0x74, 0x74,
	0x6f, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0e, 0x0a,
	0x0c, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x40, 0x01, 0x92, 0x41, 0x04,
	0x9a, 0x02, 0x01, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3e,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x3e,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x50,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0xe0,
	0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x50, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x34, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x16, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x7e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x30, 0xba, 0xb9, 0x19, 0x2c, 0x2a, 0x2a, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x12, 0x13, 0x0a, 0x0f, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x40, 0x01,
	0x48, 0x01, 0x58, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x3a, 0x17, 0xba, 0xb9, 0x19, 0x13, 0x08, 0x01, 0x1a, 0x0f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x73, 0x22, 0xa0, 0x06, 0x0a, 0x13,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2d, 0xba, 0xb9, 0x19, 0x1f, 0x0a, 0x1d, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x28, 0x01, 0x40, 0x01, 0x48, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x4d, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x23, 0xba, 0xb9, 0x19, 0x15, 0x0a, 0x13, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a,
	0x02, 0x01, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x04, 0x53,
	0x74, 0x65, 0x70, 0x3a, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x40, 0x01, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x17, 0xba, 0xb9, 0x19, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x19, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x0a, 0x08, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01, 0x03,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0d, 0x6d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x1e, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a, 0x02, 0x01,
	0x03, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a,
	0x02, 0x01, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x3e, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x11, 0x0a, 0x0f, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44, 0x40, 0x01, 0x92, 0x41, 0x04, 0x9a,
	0x02, 0x01, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x50, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x16, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x16, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x09, 0x0a,
	0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x10, 0x03, 0x3a,
	0x1d, 0xba, 0xb9, 0x19, 0x19, 0x08, 0x01, 0x1a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_gorm_db_proto_rawDescOnce sync.Once
	file_workflow_gorm_db_proto_rawDescData = file_workflow_gorm_db_proto_rawDesc
)

func file_workflow_gorm_db_proto_rawDescGZIP() []byte {
	file_workflow_gorm_db_proto_rawDescOnce.Do(func() {
		file_workflow_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_gorm_db_proto_rawDescData)
	})
	return file_workflow_gorm_db_proto_rawDescData
}

var file_workflow_gorm_db_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_workflow_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_workflow_gorm_db_proto_goTypes = []interface{}{
	(WorkflowRequirementSteps)(0), // 0: workflow.service.v1.WorkflowRequirement.steps
	(*User)(nil),                  // 1: workflow.service.v1.User
	(*Workflow)(nil),              // 2: workflow.service.v1.Workflow
	(*WorkflowLogic)(nil),         // 3: workflow.service.v1.WorkflowLogic
	(*WorkflowRequirement)(nil),   // 4: workflow.service.v1.WorkflowRequirement
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_workflow_gorm_db_proto_depIdxs = []int32{
	5,  // 0: workflow.service.v1.User.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: workflow.service.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 2: workflow.service.v1.User.deleted_at:type_name -> google.protobuf.Timestamp
	5,  // 3: workflow.service.v1.Workflow.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 4: workflow.service.v1.Workflow.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 5: workflow.service.v1.Workflow.deletedAt:type_name -> google.protobuf.Timestamp
	3,  // 6: workflow.service.v1.Workflow.logics:type_name -> workflow.service.v1.WorkflowLogic
	5,  // 7: workflow.service.v1.WorkflowLogic.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 8: workflow.service.v1.WorkflowLogic.updatedAt:type_name -> google.protobuf.Timestamp
	4,  // 9: workflow.service.v1.WorkflowLogic.requirements:type_name -> workflow.service.v1.WorkflowRequirement
	5,  // 10: workflow.service.v1.WorkflowRequirement.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 11: workflow.service.v1.WorkflowRequirement.updatedAt:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_workflow_gorm_db_proto_init() }
func file_workflow_gorm_db_proto_init() {
	if File_workflow_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_gorm_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_workflow_gorm_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowLogic); i {
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
		file_workflow_gorm_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowRequirement); i {
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
			RawDescriptor: file_workflow_gorm_db_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflow_gorm_db_proto_goTypes,
		DependencyIndexes: file_workflow_gorm_db_proto_depIdxs,
		EnumInfos:         file_workflow_gorm_db_proto_enumTypes,
		MessageInfos:      file_workflow_gorm_db_proto_msgTypes,
	}.Build()
	File_workflow_gorm_db_proto = out.File
	file_workflow_gorm_db_proto_rawDesc = nil
	file_workflow_gorm_db_proto_goTypes = nil
	file_workflow_gorm_db_proto_depIdxs = nil
}
