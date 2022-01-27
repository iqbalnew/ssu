// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: announcement.api.proto

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
	ListAnnouncement(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAnnouncementResponse, error)
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTask(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	ListAnnouncementActive(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAnnouncementActiveResponse, error)
	ListEventType(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListEventTypeResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GenerateTokenByRole(ctx context.Context, in *TempGenToken, opts ...grpc.CallOption) (*JWTTokenResponse, error) {
	out := new(JWTTokenResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/GenerateTokenByRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncement(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAnnouncementResponse, error) {
	out := new(ListAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/CreateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAnnouncementTask(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/CreateAnnouncementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncementActive(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAnnouncementActiveResponse, error) {
	out := new(ListAnnouncementActiveResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncementActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListEventType(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListEventTypeResponse, error) {
	out := new(ListEventTypeResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListEventType", in, out, opts...)
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
	ListAnnouncement(context.Context, *ListRequest) (*ListAnnouncementResponse, error)
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTask(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	ListAnnouncementActive(context.Context, *ListRequest) (*ListAnnouncementActiveResponse, error)
	ListEventType(context.Context, *ListRequest) (*ListEventTypeResponse, error)
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
func (UnimplementedApiServiceServer) ListAnnouncement(context.Context, *ListRequest) (*ListAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncement not implemented")
}
func (UnimplementedApiServiceServer) CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedApiServiceServer) CreateAnnouncementTask(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncementTask not implemented")
}
func (UnimplementedApiServiceServer) ListAnnouncementActive(context.Context, *ListRequest) (*ListAnnouncementActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncementActive not implemented")
}
func (UnimplementedApiServiceServer) ListEventType(context.Context, *ListRequest) (*ListEventTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventType not implemented")
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
		FullMethod: "/announcement.service.v1.ApiService/HealthCheck",
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
		FullMethod: "/announcement.service.v1.ApiService/GenerateTokenByRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GenerateTokenByRole(ctx, req.(*TempGenToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAnnouncement(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/CreateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAnnouncement(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAnnouncementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAnnouncementTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/CreateAnnouncementTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAnnouncementTask(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncementActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAnnouncementActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListAnnouncementActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAnnouncementActive(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListEventType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListEventType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListEventType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListEventType(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "announcement.service.v1.ApiService",
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
			MethodName: "ListAnnouncement",
			Handler:    _ApiService_ListAnnouncement_Handler,
		},
		{
			MethodName: "CreateAnnouncement",
			Handler:    _ApiService_CreateAnnouncement_Handler,
		},
		{
			MethodName: "CreateAnnouncementTask",
			Handler:    _ApiService_CreateAnnouncementTask_Handler,
		},
		{
			MethodName: "ListAnnouncementActive",
			Handler:    _ApiService_ListAnnouncementActive_Handler,
		},
		{
			MethodName: "ListEventType",
			Handler:    _ApiService_ListEventType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announcement.api.proto",
}
