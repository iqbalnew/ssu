// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: bifast_core.proto

package pb

import (
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

type Direction int32

const (
	Direction_DESC Direction = 0
	Direction_ASC  Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	Direction_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_bifast_core_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_bifast_core_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_bifast_core_proto_rawDescGZIP(), []int{0}
}

type FormatTemplate int32

const (
	FormatTemplate_null FormatTemplate = 0
	FormatTemplate_csv  FormatTemplate = 1
	FormatTemplate_xls  FormatTemplate = 2
	FormatTemplate_xlsx FormatTemplate = 3
	FormatTemplate_pdf  FormatTemplate = 4
)

// Enum value maps for FormatTemplate.
var (
	FormatTemplate_name = map[int32]string{
		0: "null",
		1: "csv",
		2: "xls",
		3: "xlsx",
		4: "pdf",
	}
	FormatTemplate_value = map[string]int32{
		"null": 0,
		"csv":  1,
		"xls":  2,
		"xlsx": 3,
		"pdf":  4,
	}
)

func (x FormatTemplate) Enum() *FormatTemplate {
	p := new(FormatTemplate)
	*p = x
	return p
}

func (x FormatTemplate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FormatTemplate) Descriptor() protoreflect.EnumDescriptor {
	return file_bifast_core_proto_enumTypes[1].Descriptor()
}

func (FormatTemplate) Type() protoreflect.EnumType {
	return &file_bifast_core_proto_enumTypes[1]
}

func (x FormatTemplate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FormatTemplate.Descriptor instead.
func (FormatTemplate) EnumDescriptor() ([]byte, []int) {
	return file_bifast_core_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bifast_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bifast_core_proto_msgTypes[0]
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
	return file_bifast_core_proto_rawDescGZIP(), []int{0}
}

type ErrorBodyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorBodyResponse) Reset() {
	*x = ErrorBodyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bifast_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorBodyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorBodyResponse) ProtoMessage() {}

func (x *ErrorBodyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bifast_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorBodyResponse.ProtoReflect.Descriptor instead.
func (*ErrorBodyResponse) Descriptor() ([]byte, []int) {
	return file_bifast_core_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorBodyResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ErrorBodyResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorBodyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column    string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bifast_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_bifast_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_bifast_core_proto_rawDescGZIP(), []int{2}
}

func (x *Sort) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *Sort) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page       int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	TotalRows  int64 `protobuf:"varint,3,opt,name=totalRows,proto3" json:"totalRows,omitempty"`
	TotalPages int32 `protobuf:"varint,4,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bifast_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bifast_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_bifast_core_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationResponse) GetTotalRows() int64 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

func (x *PaginationResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_bifast_core_proto protoreflect.FileDescriptor

var file_bifast_core_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x69, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x57, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x2a, 0x1e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x53, 0x43, 0x10, 0x01, 0x2a, 0x3f, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x63, 0x73, 0x76, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x78, 0x6c, 0x73,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x78, 0x6c, 0x73, 0x78, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x70, 0x64, 0x66, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bifast_core_proto_rawDescOnce sync.Once
	file_bifast_core_proto_rawDescData = file_bifast_core_proto_rawDesc
)

func file_bifast_core_proto_rawDescGZIP() []byte {
	file_bifast_core_proto_rawDescOnce.Do(func() {
		file_bifast_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_bifast_core_proto_rawDescData)
	})
	return file_bifast_core_proto_rawDescData
}

var file_bifast_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bifast_core_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bifast_core_proto_goTypes = []interface{}{
	(Direction)(0),             // 0: bifast.service.v1.Direction
	(FormatTemplate)(0),        // 1: bifast.service.v1.FormatTemplate
	(*Empty)(nil),              // 2: bifast.service.v1.Empty
	(*ErrorBodyResponse)(nil),  // 3: bifast.service.v1.ErrorBodyResponse
	(*Sort)(nil),               // 4: bifast.service.v1.Sort
	(*PaginationResponse)(nil), // 5: bifast.service.v1.PaginationResponse
}
var file_bifast_core_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bifast_core_proto_init() }
func file_bifast_core_proto_init() {
	if File_bifast_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bifast_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_bifast_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorBodyResponse); i {
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
		file_bifast_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
		file_bifast_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationResponse); i {
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
			RawDescriptor: file_bifast_core_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bifast_core_proto_goTypes,
		DependencyIndexes: file_bifast_core_proto_depIdxs,
		EnumInfos:         file_bifast_core_proto_enumTypes,
		MessageInfos:      file_bifast_core_proto_msgTypes,
	}.Build()
	File_bifast_core_proto = out.File
	file_bifast_core_proto_rawDesc = nil
	file_bifast_core_proto_goTypes = nil
	file_bifast_core_proto_depIdxs = nil
}