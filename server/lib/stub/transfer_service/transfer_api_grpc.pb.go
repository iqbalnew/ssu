// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: transfer_api.proto

package pb

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetPairRate(ctx context.Context, in *GetPairRateRequest, opts ...grpc.CallOption) (*GetPairRateResponse, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error)
	GetTaskInternalSingleFile(ctx context.Context, in *GetTaskInternalSingleFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetTaskInternalSingle(ctx context.Context, in *GetTaskInternalSingleRequest, opts ...grpc.CallOption) (*GetTaskInternalSingleResponse, error)
	GetTaskInternalSingleDetail(ctx context.Context, in *GetTaskInternalSingleDetailRequest, opts ...grpc.CallOption) (*GetTaskInternalSingleDetailResponse, error)
	CreateTaskInternalSingle(ctx context.Context, in *CreateTaskInternalSingleRequest, opts ...grpc.CallOption) (*CreateTaskInternalSingleResponse, error)
	GetTaskInternalMultiple(ctx context.Context, in *GetTaskInternalMultipleRequest, opts ...grpc.CallOption) (*GetTaskInternalMultipleResponse, error)
	GetTaskInternalMultipleDetail(ctx context.Context, in *GetTaskInternalMultipleDetailRequest, opts ...grpc.CallOption) (*GetTaskInternalMultipleDetailResponse, error)
	CreateTaskInternalMultiple(ctx context.Context, in *CreateTaskInternalMultipleRequest, opts ...grpc.CallOption) (*CreateTaskInternalMultipleResponse, error)
	DecodeFile(ctx context.Context, in *DecodeFileRequest, opts ...grpc.CallOption) (*DecodeFileResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetPairRate(ctx context.Context, in *GetPairRateRequest, opts ...grpc.CallOption) (*GetPairRateResponse, error) {
	out := new(GetPairRateResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetPairRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error) {
	out := new(CreateTransferResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskInternalSingleFile(ctx context.Context, in *GetTaskInternalSingleFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetTaskInternalSingleFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskInternalSingle(ctx context.Context, in *GetTaskInternalSingleRequest, opts ...grpc.CallOption) (*GetTaskInternalSingleResponse, error) {
	out := new(GetTaskInternalSingleResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetTaskInternalSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskInternalSingleDetail(ctx context.Context, in *GetTaskInternalSingleDetailRequest, opts ...grpc.CallOption) (*GetTaskInternalSingleDetailResponse, error) {
	out := new(GetTaskInternalSingleDetailResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetTaskInternalSingleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTaskInternalSingle(ctx context.Context, in *CreateTaskInternalSingleRequest, opts ...grpc.CallOption) (*CreateTaskInternalSingleResponse, error) {
	out := new(CreateTaskInternalSingleResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/CreateTaskInternalSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskInternalMultiple(ctx context.Context, in *GetTaskInternalMultipleRequest, opts ...grpc.CallOption) (*GetTaskInternalMultipleResponse, error) {
	out := new(GetTaskInternalMultipleResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetTaskInternalMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskInternalMultipleDetail(ctx context.Context, in *GetTaskInternalMultipleDetailRequest, opts ...grpc.CallOption) (*GetTaskInternalMultipleDetailResponse, error) {
	out := new(GetTaskInternalMultipleDetailResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/GetTaskInternalMultipleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateTaskInternalMultiple(ctx context.Context, in *CreateTaskInternalMultipleRequest, opts ...grpc.CallOption) (*CreateTaskInternalMultipleResponse, error) {
	out := new(CreateTaskInternalMultipleResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/CreateTaskInternalMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DecodeFile(ctx context.Context, in *DecodeFileRequest, opts ...grpc.CallOption) (*DecodeFileResponse, error) {
	out := new(DecodeFileResponse)
	err := c.cc.Invoke(ctx, "/transfer.service.v1.ApiService/DecodeFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetPairRate(context.Context, *GetPairRateRequest) (*GetPairRateResponse, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error)
	GetTaskInternalSingleFile(context.Context, *GetTaskInternalSingleFileRequest) (*httpbody.HttpBody, error)
	GetTaskInternalSingle(context.Context, *GetTaskInternalSingleRequest) (*GetTaskInternalSingleResponse, error)
	GetTaskInternalSingleDetail(context.Context, *GetTaskInternalSingleDetailRequest) (*GetTaskInternalSingleDetailResponse, error)
	CreateTaskInternalSingle(context.Context, *CreateTaskInternalSingleRequest) (*CreateTaskInternalSingleResponse, error)
	GetTaskInternalMultiple(context.Context, *GetTaskInternalMultipleRequest) (*GetTaskInternalMultipleResponse, error)
	GetTaskInternalMultipleDetail(context.Context, *GetTaskInternalMultipleDetailRequest) (*GetTaskInternalMultipleDetailResponse, error)
	CreateTaskInternalMultiple(context.Context, *CreateTaskInternalMultipleRequest) (*CreateTaskInternalMultipleResponse, error)
	DecodeFile(context.Context, *DecodeFileRequest) (*DecodeFileResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) GetPairRate(context.Context, *GetPairRateRequest) (*GetPairRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPairRate not implemented")
}
func (UnimplementedApiServiceServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedApiServiceServer) GetTaskInternalSingleFile(context.Context, *GetTaskInternalSingleFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInternalSingleFile not implemented")
}
func (UnimplementedApiServiceServer) GetTaskInternalSingle(context.Context, *GetTaskInternalSingleRequest) (*GetTaskInternalSingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInternalSingle not implemented")
}
func (UnimplementedApiServiceServer) GetTaskInternalSingleDetail(context.Context, *GetTaskInternalSingleDetailRequest) (*GetTaskInternalSingleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInternalSingleDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTaskInternalSingle(context.Context, *CreateTaskInternalSingleRequest) (*CreateTaskInternalSingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskInternalSingle not implemented")
}
func (UnimplementedApiServiceServer) GetTaskInternalMultiple(context.Context, *GetTaskInternalMultipleRequest) (*GetTaskInternalMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInternalMultiple not implemented")
}
func (UnimplementedApiServiceServer) GetTaskInternalMultipleDetail(context.Context, *GetTaskInternalMultipleDetailRequest) (*GetTaskInternalMultipleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInternalMultipleDetail not implemented")
}
func (UnimplementedApiServiceServer) CreateTaskInternalMultiple(context.Context, *CreateTaskInternalMultipleRequest) (*CreateTaskInternalMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskInternalMultiple not implemented")
}
func (UnimplementedApiServiceServer) DecodeFile(context.Context, *DecodeFileRequest) (*DecodeFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeFile not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetPairRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPairRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetPairRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetPairRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetPairRate(ctx, req.(*GetPairRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskInternalSingleFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskInternalSingleFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskInternalSingleFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetTaskInternalSingleFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskInternalSingleFile(ctx, req.(*GetTaskInternalSingleFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskInternalSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskInternalSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskInternalSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetTaskInternalSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskInternalSingle(ctx, req.(*GetTaskInternalSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskInternalSingleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskInternalSingleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskInternalSingleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetTaskInternalSingleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskInternalSingleDetail(ctx, req.(*GetTaskInternalSingleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTaskInternalSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskInternalSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTaskInternalSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/CreateTaskInternalSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTaskInternalSingle(ctx, req.(*CreateTaskInternalSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskInternalMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskInternalMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskInternalMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetTaskInternalMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskInternalMultiple(ctx, req.(*GetTaskInternalMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskInternalMultipleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskInternalMultipleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskInternalMultipleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/GetTaskInternalMultipleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskInternalMultipleDetail(ctx, req.(*GetTaskInternalMultipleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateTaskInternalMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskInternalMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateTaskInternalMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/CreateTaskInternalMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateTaskInternalMultiple(ctx, req.(*CreateTaskInternalMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DecodeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DecodeFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfer.service.v1.ApiService/DecodeFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DecodeFile(ctx, req.(*DecodeFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transfer.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "GetPairRate",
			Handler:    _ApiService_GetPairRate_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _ApiService_CreateTransfer_Handler,
		},
		{
			MethodName: "GetTaskInternalSingleFile",
			Handler:    _ApiService_GetTaskInternalSingleFile_Handler,
		},
		{
			MethodName: "GetTaskInternalSingle",
			Handler:    _ApiService_GetTaskInternalSingle_Handler,
		},
		{
			MethodName: "GetTaskInternalSingleDetail",
			Handler:    _ApiService_GetTaskInternalSingleDetail_Handler,
		},
		{
			MethodName: "CreateTaskInternalSingle",
			Handler:    _ApiService_CreateTaskInternalSingle_Handler,
		},
		{
			MethodName: "GetTaskInternalMultiple",
			Handler:    _ApiService_GetTaskInternalMultiple_Handler,
		},
		{
			MethodName: "GetTaskInternalMultipleDetail",
			Handler:    _ApiService_GetTaskInternalMultipleDetail_Handler,
		},
		{
			MethodName: "CreateTaskInternalMultiple",
			Handler:    _ApiService_CreateTaskInternalMultiple_Handler,
		},
		{
			MethodName: "DecodeFile",
			Handler:    _ApiService_DecodeFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer_api.proto",
}