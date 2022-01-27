// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: company_api.proto

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
	GenerateTokenByRole(ctx context.Context, in *TempGenToken, opts ...grpc.CallOption) (*JWTTokenResponse, error)
	ListCompany(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyResponse, error)
	ListCompanyGroup(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyGroupResponse, error)
	ListCreateCompanyGroupTask(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyGroupTaskResponse, error)
	ListCurrency(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCurrencyResponse, error)
	CreateCompanyGroup(ctx context.Context, in *CreateCompanyGroupRequest, opts ...grpc.CallOption) (*CreateCompanyGroupResponse, error)
	CreateCompanyGroupTask(ctx context.Context, in *CreateCompanyGroupTaskRequest, opts ...grpc.CallOption) (*CreateCompanyGroupResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GenerateTokenByRole(ctx context.Context, in *TempGenToken, opts ...grpc.CallOption) (*JWTTokenResponse, error) {
	out := new(JWTTokenResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/GenerateTokenByRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListCompany(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyResponse, error) {
	out := new(ListCompanyResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/ListCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListCompanyGroup(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyGroupResponse, error) {
	out := new(ListCompanyGroupResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/ListCompanyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListCreateCompanyGroupTask(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCompanyGroupTaskResponse, error) {
	out := new(ListCompanyGroupTaskResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/ListCreateCompanyGroupTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListCurrency(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListCurrencyResponse, error) {
	out := new(ListCurrencyResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/ListCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateCompanyGroup(ctx context.Context, in *CreateCompanyGroupRequest, opts ...grpc.CallOption) (*CreateCompanyGroupResponse, error) {
	out := new(CreateCompanyGroupResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/CreateCompanyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateCompanyGroupTask(ctx context.Context, in *CreateCompanyGroupTaskRequest, opts ...grpc.CallOption) (*CreateCompanyGroupResponse, error) {
	out := new(CreateCompanyGroupResponse)
	err := c.cc.Invoke(ctx, "/company.service.v1.ApiService/CreateCompanyGroupTask", in, out, opts...)
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
	GenerateTokenByRole(context.Context, *TempGenToken) (*JWTTokenResponse, error)
	ListCompany(context.Context, *ListRequest) (*ListCompanyResponse, error)
	ListCompanyGroup(context.Context, *ListRequest) (*ListCompanyGroupResponse, error)
	ListCreateCompanyGroupTask(context.Context, *ListRequest) (*ListCompanyGroupTaskResponse, error)
	ListCurrency(context.Context, *ListRequest) (*ListCurrencyResponse, error)
	CreateCompanyGroup(context.Context, *CreateCompanyGroupRequest) (*CreateCompanyGroupResponse, error)
	CreateCompanyGroupTask(context.Context, *CreateCompanyGroupTaskRequest) (*CreateCompanyGroupResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) GenerateTokenByRole(context.Context, *TempGenToken) (*JWTTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTokenByRole not implemented")
}
func (UnimplementedApiServiceServer) ListCompany(context.Context, *ListRequest) (*ListCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompany not implemented")
}
func (UnimplementedApiServiceServer) ListCompanyGroup(context.Context, *ListRequest) (*ListCompanyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanyGroup not implemented")
}
func (UnimplementedApiServiceServer) ListCreateCompanyGroupTask(context.Context, *ListRequest) (*ListCompanyGroupTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCreateCompanyGroupTask not implemented")
}
func (UnimplementedApiServiceServer) ListCurrency(context.Context, *ListRequest) (*ListCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrency not implemented")
}
func (UnimplementedApiServiceServer) CreateCompanyGroup(context.Context, *CreateCompanyGroupRequest) (*CreateCompanyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompanyGroup not implemented")
}
func (UnimplementedApiServiceServer) CreateCompanyGroupTask(context.Context, *CreateCompanyGroupTaskRequest) (*CreateCompanyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompanyGroupTask not implemented")
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
		FullMethod: "/company.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GenerateTokenByRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempGenToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GenerateTokenByRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/GenerateTokenByRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GenerateTokenByRole(ctx, req.(*TempGenToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/ListCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListCompany(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListCompanyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListCompanyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/ListCompanyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListCompanyGroup(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListCreateCompanyGroupTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListCreateCompanyGroupTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/ListCreateCompanyGroupTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListCreateCompanyGroupTask(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/ListCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListCurrency(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateCompanyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateCompanyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/CreateCompanyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateCompanyGroup(ctx, req.(*CreateCompanyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateCompanyGroupTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyGroupTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateCompanyGroupTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.service.v1.ApiService/CreateCompanyGroupTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateCompanyGroupTask(ctx, req.(*CreateCompanyGroupTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "company.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "GenerateTokenByRole",
			Handler:    _ApiService_GenerateTokenByRole_Handler,
		},
		{
			MethodName: "ListCompany",
			Handler:    _ApiService_ListCompany_Handler,
		},
		{
			MethodName: "ListCompanyGroup",
			Handler:    _ApiService_ListCompanyGroup_Handler,
		},
		{
			MethodName: "ListCreateCompanyGroupTask",
			Handler:    _ApiService_ListCreateCompanyGroupTask_Handler,
		},
		{
			MethodName: "ListCurrency",
			Handler:    _ApiService_ListCurrency_Handler,
		},
		{
			MethodName: "CreateCompanyGroup",
			Handler:    _ApiService_CreateCompanyGroup_Handler,
		},
		{
			MethodName: "CreateCompanyGroupTask",
			Handler:    _ApiService_CreateCompanyGroupTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company_api.proto",
}
