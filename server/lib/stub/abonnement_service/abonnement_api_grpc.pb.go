// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: abonnement_api.proto

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
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	ListAbonnement(ctx context.Context, in *ListAbonnementRequest, opts ...grpc.CallOption) (*ListAbonnementResponse, error)
	DeleteAbonnement(ctx context.Context, in *CreateAbonnementRequest, opts ...grpc.CallOption) (*CreateAbonnementResponse, error)
	CreateAbonnement(ctx context.Context, in *CreateAbonnementRequest, opts ...grpc.CallOption) (*CreateAbonnementResponse, error)
	RequestDeleteAbonnementTask(ctx context.Context, in *AbonnementTaskDetailRequest, opts ...grpc.CallOption) (*AbonnementTaskDetailResponse, error)
	CreateAbonnementTask(ctx context.Context, in *CreateAbonnementTaskRequest, opts ...grpc.CallOption) (*CreateAbonnementTaskResponse, error)
	DownloadListAbonnementTasks(ctx context.Context, in *DownloadListAbonnementTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	ListAbonnementTask(ctx context.Context, in *ListAbonnementTaskRequest, opts ...grpc.CallOption) (*ListAbonnementTaskResponse, error)
	DownloadListAbonnementInvoice(ctx context.Context, in *DownloadListAbonnementInvoiceRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	AbonnementTaskDetail(ctx context.Context, in *AbonnementTaskDetailRequest, opts ...grpc.CallOption) (*AbonnementTaskDetailResponse, error)
	ListAbonnementInvoice(ctx context.Context, in *ListAbonnementInvoiceRequest, opts ...grpc.CallOption) (*ListAbonnementInvoiceResponse, error)
	CekAvaibility(ctx context.Context, in *CekAvaibilityReq, opts ...grpc.CallOption) (*CekAvaibilityRes, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAbonnement(ctx context.Context, in *ListAbonnementRequest, opts ...grpc.CallOption) (*ListAbonnementResponse, error) {
	out := new(ListAbonnementResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/ListAbonnement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteAbonnement(ctx context.Context, in *CreateAbonnementRequest, opts ...grpc.CallOption) (*CreateAbonnementResponse, error) {
	out := new(CreateAbonnementResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/DeleteAbonnement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAbonnement(ctx context.Context, in *CreateAbonnementRequest, opts ...grpc.CallOption) (*CreateAbonnementResponse, error) {
	out := new(CreateAbonnementResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/CreateAbonnement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RequestDeleteAbonnementTask(ctx context.Context, in *AbonnementTaskDetailRequest, opts ...grpc.CallOption) (*AbonnementTaskDetailResponse, error) {
	out := new(AbonnementTaskDetailResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/RequestDeleteAbonnementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAbonnementTask(ctx context.Context, in *CreateAbonnementTaskRequest, opts ...grpc.CallOption) (*CreateAbonnementTaskResponse, error) {
	out := new(CreateAbonnementTaskResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/CreateAbonnementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadListAbonnementTasks(ctx context.Context, in *DownloadListAbonnementTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/DownloadListAbonnementTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAbonnementTask(ctx context.Context, in *ListAbonnementTaskRequest, opts ...grpc.CallOption) (*ListAbonnementTaskResponse, error) {
	out := new(ListAbonnementTaskResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/ListAbonnementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadListAbonnementInvoice(ctx context.Context, in *DownloadListAbonnementInvoiceRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/DownloadListAbonnementInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AbonnementTaskDetail(ctx context.Context, in *AbonnementTaskDetailRequest, opts ...grpc.CallOption) (*AbonnementTaskDetailResponse, error) {
	out := new(AbonnementTaskDetailResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/AbonnementTaskDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAbonnementInvoice(ctx context.Context, in *ListAbonnementInvoiceRequest, opts ...grpc.CallOption) (*ListAbonnementInvoiceResponse, error) {
	out := new(ListAbonnementInvoiceResponse)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/ListAbonnementInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CekAvaibility(ctx context.Context, in *CekAvaibilityReq, opts ...grpc.CallOption) (*CekAvaibilityRes, error) {
	out := new(CekAvaibilityRes)
	err := c.cc.Invoke(ctx, "/abonnement.service.v1.ApiService/CekAvaibility", in, out, opts...)
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
	ListAbonnement(context.Context, *ListAbonnementRequest) (*ListAbonnementResponse, error)
	DeleteAbonnement(context.Context, *CreateAbonnementRequest) (*CreateAbonnementResponse, error)
	CreateAbonnement(context.Context, *CreateAbonnementRequest) (*CreateAbonnementResponse, error)
	RequestDeleteAbonnementTask(context.Context, *AbonnementTaskDetailRequest) (*AbonnementTaskDetailResponse, error)
	CreateAbonnementTask(context.Context, *CreateAbonnementTaskRequest) (*CreateAbonnementTaskResponse, error)
	DownloadListAbonnementTasks(context.Context, *DownloadListAbonnementTaskRequest) (*httpbody.HttpBody, error)
	ListAbonnementTask(context.Context, *ListAbonnementTaskRequest) (*ListAbonnementTaskResponse, error)
	DownloadListAbonnementInvoice(context.Context, *DownloadListAbonnementInvoiceRequest) (*httpbody.HttpBody, error)
	AbonnementTaskDetail(context.Context, *AbonnementTaskDetailRequest) (*AbonnementTaskDetailResponse, error)
	ListAbonnementInvoice(context.Context, *ListAbonnementInvoiceRequest) (*ListAbonnementInvoiceResponse, error)
	CekAvaibility(context.Context, *CekAvaibilityReq) (*CekAvaibilityRes, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListAbonnement(context.Context, *ListAbonnementRequest) (*ListAbonnementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAbonnement not implemented")
}
func (UnimplementedApiServiceServer) DeleteAbonnement(context.Context, *CreateAbonnementRequest) (*CreateAbonnementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAbonnement not implemented")
}
func (UnimplementedApiServiceServer) CreateAbonnement(context.Context, *CreateAbonnementRequest) (*CreateAbonnementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAbonnement not implemented")
}
func (UnimplementedApiServiceServer) RequestDeleteAbonnementTask(context.Context, *AbonnementTaskDetailRequest) (*AbonnementTaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDeleteAbonnementTask not implemented")
}
func (UnimplementedApiServiceServer) CreateAbonnementTask(context.Context, *CreateAbonnementTaskRequest) (*CreateAbonnementTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAbonnementTask not implemented")
}
func (UnimplementedApiServiceServer) DownloadListAbonnementTasks(context.Context, *DownloadListAbonnementTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListAbonnementTasks not implemented")
}
func (UnimplementedApiServiceServer) ListAbonnementTask(context.Context, *ListAbonnementTaskRequest) (*ListAbonnementTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAbonnementTask not implemented")
}
func (UnimplementedApiServiceServer) DownloadListAbonnementInvoice(context.Context, *DownloadListAbonnementInvoiceRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListAbonnementInvoice not implemented")
}
func (UnimplementedApiServiceServer) AbonnementTaskDetail(context.Context, *AbonnementTaskDetailRequest) (*AbonnementTaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbonnementTaskDetail not implemented")
}
func (UnimplementedApiServiceServer) ListAbonnementInvoice(context.Context, *ListAbonnementInvoiceRequest) (*ListAbonnementInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAbonnementInvoice not implemented")
}
func (UnimplementedApiServiceServer) CekAvaibility(context.Context, *CekAvaibilityReq) (*CekAvaibilityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CekAvaibility not implemented")
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
		FullMethod: "/abonnement.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAbonnement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAbonnementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAbonnement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/ListAbonnement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAbonnement(ctx, req.(*ListAbonnementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteAbonnement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAbonnementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteAbonnement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/DeleteAbonnement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteAbonnement(ctx, req.(*CreateAbonnementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAbonnement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAbonnementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAbonnement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/CreateAbonnement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAbonnement(ctx, req.(*CreateAbonnementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RequestDeleteAbonnementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbonnementTaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RequestDeleteAbonnementTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/RequestDeleteAbonnementTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RequestDeleteAbonnementTask(ctx, req.(*AbonnementTaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAbonnementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAbonnementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAbonnementTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/CreateAbonnementTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAbonnementTask(ctx, req.(*CreateAbonnementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadListAbonnementTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadListAbonnementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListAbonnementTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/DownloadListAbonnementTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListAbonnementTasks(ctx, req.(*DownloadListAbonnementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAbonnementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAbonnementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAbonnementTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/ListAbonnementTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAbonnementTask(ctx, req.(*ListAbonnementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadListAbonnementInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadListAbonnementInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListAbonnementInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/DownloadListAbonnementInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListAbonnementInvoice(ctx, req.(*DownloadListAbonnementInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AbonnementTaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbonnementTaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AbonnementTaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/AbonnementTaskDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AbonnementTaskDetail(ctx, req.(*AbonnementTaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAbonnementInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAbonnementInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAbonnementInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/ListAbonnementInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAbonnementInvoice(ctx, req.(*ListAbonnementInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CekAvaibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CekAvaibilityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CekAvaibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abonnement.service.v1.ApiService/CekAvaibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CekAvaibility(ctx, req.(*CekAvaibilityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abonnement.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListAbonnement",
			Handler:    _ApiService_ListAbonnement_Handler,
		},
		{
			MethodName: "DeleteAbonnement",
			Handler:    _ApiService_DeleteAbonnement_Handler,
		},
		{
			MethodName: "CreateAbonnement",
			Handler:    _ApiService_CreateAbonnement_Handler,
		},
		{
			MethodName: "RequestDeleteAbonnementTask",
			Handler:    _ApiService_RequestDeleteAbonnementTask_Handler,
		},
		{
			MethodName: "CreateAbonnementTask",
			Handler:    _ApiService_CreateAbonnementTask_Handler,
		},
		{
			MethodName: "DownloadListAbonnementTasks",
			Handler:    _ApiService_DownloadListAbonnementTasks_Handler,
		},
		{
			MethodName: "ListAbonnementTask",
			Handler:    _ApiService_ListAbonnementTask_Handler,
		},
		{
			MethodName: "DownloadListAbonnementInvoice",
			Handler:    _ApiService_DownloadListAbonnementInvoice_Handler,
		},
		{
			MethodName: "AbonnementTaskDetail",
			Handler:    _ApiService_AbonnementTaskDetail_Handler,
		},
		{
			MethodName: "ListAbonnementInvoice",
			Handler:    _ApiService_ListAbonnementInvoice_Handler,
		},
		{
			MethodName: "CekAvaibility",
			Handler:    _ApiService_CekAvaibility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "abonnement_api.proto",
}
