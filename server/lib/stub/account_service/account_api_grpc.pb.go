// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: account_api.proto

package pb

import (
	context "context"
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
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateAccountTask(ctx context.Context, in *CreateAccountTaskRequest, opts ...grpc.CallOption) (*CreateAccountTaskResponse, error)
	ListAccountTask(ctx context.Context, in *ListAccountTaskRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error)
	ValidateAccount(ctx context.Context, in *ValidateAccountRequest, opts ...grpc.CallOption) (*ValidateAccountResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountResponse, error) {
	out := new(ListAccountResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/ListAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAccountTask(ctx context.Context, in *CreateAccountTaskRequest, opts ...grpc.CallOption) (*CreateAccountTaskResponse, error) {
	out := new(CreateAccountTaskResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/CreateAccountTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAccountTask(ctx context.Context, in *ListAccountTaskRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error) {
	out := new(ListAccountTaskResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/ListAccountTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ValidateAccount(ctx context.Context, in *ValidateAccountRequest, opts ...grpc.CallOption) (*ValidateAccountResponse, error) {
	out := new(ValidateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/ValidateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error)
	ListAccount(context.Context, *ListAccountRequest) (*ListAccountResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	CreateAccountTask(context.Context, *CreateAccountTaskRequest) (*CreateAccountTaskResponse, error)
	ListAccountTask(context.Context, *ListAccountTaskRequest) (*ListAccountTaskResponse, error)
	ValidateAccount(context.Context, *ValidateAccountRequest) (*ValidateAccountResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListAccount(context.Context, *ListAccountRequest) (*ListAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedApiServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedApiServiceServer) CreateAccountTask(context.Context, *CreateAccountTaskRequest) (*CreateAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountTask not implemented")
}
func (UnimplementedApiServiceServer) ListAccountTask(context.Context, *ListAccountTaskRequest) (*ListAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountTask not implemented")
}
func (UnimplementedApiServiceServer) ValidateAccount(context.Context, *ValidateAccountRequest) (*ValidateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAccount not implemented")
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
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/ListAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAccount(ctx, req.(*ListAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAccountTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAccountTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/CreateAccountTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAccountTask(ctx, req.(*CreateAccountTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAccountTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAccountTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/ListAccountTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAccountTask(ctx, req.(*ListAccountTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ValidateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ValidateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/ValidateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ValidateAccount(ctx, req.(*ValidateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _ApiService_ListAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _ApiService_CreateAccount_Handler,
		},
		{
			MethodName: "CreateAccountTask",
			Handler:    _ApiService_CreateAccountTask_Handler,
		},
		{
			MethodName: "ListAccountTask",
			Handler:    _ApiService_ListAccountTask_Handler,
		},
		{
			MethodName: "ValidateAccount",
			Handler:    _ApiService_ValidateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_api.proto",
}