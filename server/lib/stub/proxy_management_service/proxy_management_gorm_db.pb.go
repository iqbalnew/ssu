// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proxy_management_gorm_db.proto

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

type ProxyManagementType int32

const (
	ProxyManagementType_NullProxyManagementType ProxyManagementType = 0
	ProxyManagementType_Email                   ProxyManagementType = 1
	ProxyManagementType_MobileNumber            ProxyManagementType = 2
)

// Enum value maps for ProxyManagementType.
var (
	ProxyManagementType_name = map[int32]string{
		0: "NullProxyManagementType",
		1: "Email",
		2: "MobileNumber",
	}
	ProxyManagementType_value = map[string]int32{
		"NullProxyManagementType": 0,
		"Email":                   1,
		"MobileNumber":            2,
	}
)

func (x ProxyManagementType) Enum() *ProxyManagementType {
	p := new(ProxyManagementType)
	*p = x
	return p
}

func (x ProxyManagementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyManagementType) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_management_gorm_db_proto_enumTypes[0].Descriptor()
}

func (ProxyManagementType) Type() protoreflect.EnumType {
	return &file_proxy_management_gorm_db_proto_enumTypes[0]
}

func (x ProxyManagementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyManagementType.Descriptor instead.
func (ProxyManagementType) EnumDescriptor() ([]byte, []int) {
	return file_proxy_management_gorm_db_proto_rawDescGZIP(), []int{0}
}

type ProxyManagementAccountType int32

const (
	ProxyManagementAccountType_NullAccountType ProxyManagementAccountType = 0
	ProxyManagementAccountType_Deposit         ProxyManagementAccountType = 1
)

// Enum value maps for ProxyManagementAccountType.
var (
	ProxyManagementAccountType_name = map[int32]string{
		0: "NullAccountType",
		1: "Deposit",
	}
	ProxyManagementAccountType_value = map[string]int32{
		"NullAccountType": 0,
		"Deposit":         1,
	}
)

func (x ProxyManagementAccountType) Enum() *ProxyManagementAccountType {
	p := new(ProxyManagementAccountType)
	*p = x
	return p
}

func (x ProxyManagementAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyManagementAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_management_gorm_db_proto_enumTypes[1].Descriptor()
}

func (ProxyManagementAccountType) Type() protoreflect.EnumType {
	return &file_proxy_management_gorm_db_proto_enumTypes[1]
}

func (x ProxyManagementAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyManagementAccountType.Descriptor instead.
func (ProxyManagementAccountType) EnumDescriptor() ([]byte, []int) {
	return file_proxy_management_gorm_db_proto_rawDescGZIP(), []int{1}
}

type ProxyManagementAccountStatus int32

const (
	ProxyManagementAccountStatus_NullAccountStatus ProxyManagementAccountStatus = 0
	ProxyManagementAccountStatus_Active            ProxyManagementAccountStatus = 1
)

// Enum value maps for ProxyManagementAccountStatus.
var (
	ProxyManagementAccountStatus_name = map[int32]string{
		0: "NullAccountStatus",
		1: "Active",
	}
	ProxyManagementAccountStatus_value = map[string]int32{
		"NullAccountStatus": 0,
		"Active":            1,
	}
)

func (x ProxyManagementAccountStatus) Enum() *ProxyManagementAccountStatus {
	p := new(ProxyManagementAccountStatus)
	*p = x
	return p
}

func (x ProxyManagementAccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyManagementAccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_management_gorm_db_proto_enumTypes[2].Descriptor()
}

func (ProxyManagementAccountStatus) Type() protoreflect.EnumType {
	return &file_proxy_management_gorm_db_proto_enumTypes[2]
}

func (x ProxyManagementAccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyManagementAccountStatus.Descriptor instead.
func (ProxyManagementAccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_proxy_management_gorm_db_proto_rawDescGZIP(), []int{2}
}

type ProxyManagement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskID    uint64                 `protobuf:"varint,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Data      string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ProxyManagement) Reset() {
	*x = ProxyManagement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_management_gorm_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyManagement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyManagement) ProtoMessage() {}

func (x *ProxyManagement) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_management_gorm_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyManagement.ProtoReflect.Descriptor instead.
func (*ProxyManagement) Descriptor() ([]byte, []int) {
	return file_proxy_management_gorm_db_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyManagement) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProxyManagement) GetTaskID() uint64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *ProxyManagement) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ProxyManagement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProxyManagement) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_proxy_management_gorm_db_proto protoreflect.FileDescriptor

var file_proxy_management_gorm_db_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0xb9, 0x19, 0x09, 0x0a, 0x07, 0x12, 0x05,
	0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x19, 0xba, 0xb9, 0x19, 0x15, 0x08,
	0x01, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2a, 0x4f, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4e,
	0x75, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x10, 0x02, 0x2a, 0x3e, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x75, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x10, 0x01, 0x2a, 0x41, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x75, 0x6c, 0x6c, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_management_gorm_db_proto_rawDescOnce sync.Once
	file_proxy_management_gorm_db_proto_rawDescData = file_proxy_management_gorm_db_proto_rawDesc
)

func file_proxy_management_gorm_db_proto_rawDescGZIP() []byte {
	file_proxy_management_gorm_db_proto_rawDescOnce.Do(func() {
		file_proxy_management_gorm_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_management_gorm_db_proto_rawDescData)
	})
	return file_proxy_management_gorm_db_proto_rawDescData
}

var file_proxy_management_gorm_db_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proxy_management_gorm_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_management_gorm_db_proto_goTypes = []interface{}{
	(ProxyManagementType)(0),          // 0: proxy_management.service.v1.ProxyManagementType
	(ProxyManagementAccountType)(0),   // 1: proxy_management.service.v1.ProxyManagementAccountType
	(ProxyManagementAccountStatus)(0), // 2: proxy_management.service.v1.ProxyManagementAccountStatus
	(*ProxyManagement)(nil),           // 3: proxy_management.service.v1.ProxyManagement
	(*timestamppb.Timestamp)(nil),     // 4: google.protobuf.Timestamp
}
var file_proxy_management_gorm_db_proto_depIdxs = []int32{
	4, // 0: proxy_management.service.v1.ProxyManagement.createdAt:type_name -> google.protobuf.Timestamp
	4, // 1: proxy_management.service.v1.ProxyManagement.updatedAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proxy_management_gorm_db_proto_init() }
func file_proxy_management_gorm_db_proto_init() {
	if File_proxy_management_gorm_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_management_gorm_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyManagement); i {
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
			RawDescriptor: file_proxy_management_gorm_db_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_management_gorm_db_proto_goTypes,
		DependencyIndexes: file_proxy_management_gorm_db_proto_depIdxs,
		EnumInfos:         file_proxy_management_gorm_db_proto_enumTypes,
		MessageInfos:      file_proxy_management_gorm_db_proto_msgTypes,
	}.Build()
	File_proxy_management_gorm_db_proto = out.File
	file_proxy_management_gorm_db_proto_rawDesc = nil
	file_proxy_management_gorm_db_proto_goTypes = nil
	file_proxy_management_gorm_db_proto_depIdxs = nil
}