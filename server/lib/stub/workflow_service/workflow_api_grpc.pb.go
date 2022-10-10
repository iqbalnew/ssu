// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: workflow_api.proto

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
	ListWorkflow(ctx context.Context, in *ListWorkflowRequest, opts ...grpc.CallOption) (*ListWorkflowResponse, error)
	CreateWorkflow(ctx context.Context, in *WorkflowTask, opts ...grpc.CallOption) (*CreateWorkflowResponse, error)
	DeleteWorkflow(ctx context.Context, in *WorkflowTask, opts ...grpc.CallOption) (*CreateWorkflowResponse, error)
	CreateWorkflowTask(ctx context.Context, in *CreateWorkflowTaskRequest, opts ...grpc.CallOption) (*CreateWorkflowTaskResponse, error)
	ListWorkflowTask(ctx context.Context, in *ListWorkflowTaskRequest, opts ...grpc.CallOption) (*ListWorkflowTaskResponse, error)
	GetWorkflowTask(ctx context.Context, in *GetWorkflowTaskRequest, opts ...grpc.CallOption) (*GetWorkflowTaskResponse, error)
	RequestDeleteWorkflowTask(ctx context.Context, in *GetWorkflowTaskRequest, opts ...grpc.CallOption) (*GetWorkflowTaskResponse, error)
	DownloadListWorkflowTask(ctx context.Context, in *DownloadListWorkflowTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	ListMyWorkflowTask(ctx context.Context, in *ListWorkflowTaskRequest, opts ...grpc.CallOption) (*ListWorkflowTaskResponse, error)
	ValidateWorkflow(ctx context.Context, in *ValidateWorkflowRequest, opts ...grpc.CallOption) (*ValidateWorkflowResponse, error)
	GenerateWorkflow(ctx context.Context, in *GenerateWorkflowRequest, opts ...grpc.CallOption) (*ValidateWorkflowResponse, error)
	ListCompanyWorkflow(ctx context.Context, in *ListCompanyWorkflowRequest, opts ...grpc.CallOption) (*ListCompanyWorkflowResponse, error)
	GetCompanyWorkflow(ctx context.Context, in *GetCompanyWorkflowRequest, opts ...grpc.CallOption) (*GetCompanyWorkflowResponse, error)
	CreateCompanyWorkflow(ctx context.Context, in *CreateCompanyWorkflowRequest, opts ...grpc.CallOption) (*CreateCompanyWorkflowResponse, error)
	GetAvailableCurrency(ctx context.Context, in *GetAvailableCurrencyRequest, opts ...grpc.CallOption) (*GetAvailableCurrencyResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListWorkflow(ctx context.Context, in *ListWorkflowRequest, opts ...grpc.CallOption) (*ListWorkflowResponse, error) {
	out := new(ListWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/ListWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateWorkflow(ctx context.Context, in *WorkflowTask, opts ...grpc.CallOption) (*CreateWorkflowResponse, error) {
	out := new(CreateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteWorkflow(ctx context.Context, in *WorkflowTask, opts ...grpc.CallOption) (*CreateWorkflowResponse, error) {
	out := new(CreateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/DeleteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateWorkflowTask(ctx context.Context, in *CreateWorkflowTaskRequest, opts ...grpc.CallOption) (*CreateWorkflowTaskResponse, error) {
	out := new(CreateWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/CreateWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListWorkflowTask(ctx context.Context, in *ListWorkflowTaskRequest, opts ...grpc.CallOption) (*ListWorkflowTaskResponse, error) {
	out := new(ListWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/ListWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetWorkflowTask(ctx context.Context, in *GetWorkflowTaskRequest, opts ...grpc.CallOption) (*GetWorkflowTaskResponse, error) {
	out := new(GetWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/GetWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RequestDeleteWorkflowTask(ctx context.Context, in *GetWorkflowTaskRequest, opts ...grpc.CallOption) (*GetWorkflowTaskResponse, error) {
	out := new(GetWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/RequestDeleteWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadListWorkflowTask(ctx context.Context, in *DownloadListWorkflowTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/DownloadListWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMyWorkflowTask(ctx context.Context, in *ListWorkflowTaskRequest, opts ...grpc.CallOption) (*ListWorkflowTaskResponse, error) {
	out := new(ListWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/ListMyWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ValidateWorkflow(ctx context.Context, in *ValidateWorkflowRequest, opts ...grpc.CallOption) (*ValidateWorkflowResponse, error) {
	out := new(ValidateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/ValidateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GenerateWorkflow(ctx context.Context, in *GenerateWorkflowRequest, opts ...grpc.CallOption) (*ValidateWorkflowResponse, error) {
	out := new(ValidateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/GenerateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListCompanyWorkflow(ctx context.Context, in *ListCompanyWorkflowRequest, opts ...grpc.CallOption) (*ListCompanyWorkflowResponse, error) {
	out := new(ListCompanyWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/ListCompanyWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetCompanyWorkflow(ctx context.Context, in *GetCompanyWorkflowRequest, opts ...grpc.CallOption) (*GetCompanyWorkflowResponse, error) {
	out := new(GetCompanyWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/GetCompanyWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateCompanyWorkflow(ctx context.Context, in *CreateCompanyWorkflowRequest, opts ...grpc.CallOption) (*CreateCompanyWorkflowResponse, error) {
	out := new(CreateCompanyWorkflowResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/CreateCompanyWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetAvailableCurrency(ctx context.Context, in *GetAvailableCurrencyRequest, opts ...grpc.CallOption) (*GetAvailableCurrencyResponse, error) {
	out := new(GetAvailableCurrencyResponse)
	err := c.cc.Invoke(ctx, "/workflow.service.v1.ApiService/GetAvailableCurrency", in, out, opts...)
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
	ListWorkflow(context.Context, *ListWorkflowRequest) (*ListWorkflowResponse, error)
	CreateWorkflow(context.Context, *WorkflowTask) (*CreateWorkflowResponse, error)
	DeleteWorkflow(context.Context, *WorkflowTask) (*CreateWorkflowResponse, error)
	CreateWorkflowTask(context.Context, *CreateWorkflowTaskRequest) (*CreateWorkflowTaskResponse, error)
	ListWorkflowTask(context.Context, *ListWorkflowTaskRequest) (*ListWorkflowTaskResponse, error)
	GetWorkflowTask(context.Context, *GetWorkflowTaskRequest) (*GetWorkflowTaskResponse, error)
	RequestDeleteWorkflowTask(context.Context, *GetWorkflowTaskRequest) (*GetWorkflowTaskResponse, error)
	DownloadListWorkflowTask(context.Context, *DownloadListWorkflowTaskRequest) (*httpbody.HttpBody, error)
	ListMyWorkflowTask(context.Context, *ListWorkflowTaskRequest) (*ListWorkflowTaskResponse, error)
	ValidateWorkflow(context.Context, *ValidateWorkflowRequest) (*ValidateWorkflowResponse, error)
	GenerateWorkflow(context.Context, *GenerateWorkflowRequest) (*ValidateWorkflowResponse, error)
	ListCompanyWorkflow(context.Context, *ListCompanyWorkflowRequest) (*ListCompanyWorkflowResponse, error)
	GetCompanyWorkflow(context.Context, *GetCompanyWorkflowRequest) (*GetCompanyWorkflowResponse, error)
	CreateCompanyWorkflow(context.Context, *CreateCompanyWorkflowRequest) (*CreateCompanyWorkflowResponse, error)
	GetAvailableCurrency(context.Context, *GetAvailableCurrencyRequest) (*GetAvailableCurrencyResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListWorkflow(context.Context, *ListWorkflowRequest) (*ListWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflow not implemented")
}
func (UnimplementedApiServiceServer) CreateWorkflow(context.Context, *WorkflowTask) (*CreateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (UnimplementedApiServiceServer) DeleteWorkflow(context.Context, *WorkflowTask) (*CreateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedApiServiceServer) CreateWorkflowTask(context.Context, *CreateWorkflowTaskRequest) (*CreateWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) ListWorkflowTask(context.Context, *ListWorkflowTaskRequest) (*ListWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) GetWorkflowTask(context.Context, *GetWorkflowTaskRequest) (*GetWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) RequestDeleteWorkflowTask(context.Context, *GetWorkflowTaskRequest) (*GetWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDeleteWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) DownloadListWorkflowTask(context.Context, *DownloadListWorkflowTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) ListMyWorkflowTask(context.Context, *ListWorkflowTaskRequest) (*ListWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMyWorkflowTask not implemented")
}
func (UnimplementedApiServiceServer) ValidateWorkflow(context.Context, *ValidateWorkflowRequest) (*ValidateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateWorkflow not implemented")
}
func (UnimplementedApiServiceServer) GenerateWorkflow(context.Context, *GenerateWorkflowRequest) (*ValidateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWorkflow not implemented")
}
func (UnimplementedApiServiceServer) ListCompanyWorkflow(context.Context, *ListCompanyWorkflowRequest) (*ListCompanyWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanyWorkflow not implemented")
}
func (UnimplementedApiServiceServer) GetCompanyWorkflow(context.Context, *GetCompanyWorkflowRequest) (*GetCompanyWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyWorkflow not implemented")
}
func (UnimplementedApiServiceServer) CreateCompanyWorkflow(context.Context, *CreateCompanyWorkflowRequest) (*CreateCompanyWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompanyWorkflow not implemented")
}
func (UnimplementedApiServiceServer) GetAvailableCurrency(context.Context, *GetAvailableCurrencyRequest) (*GetAvailableCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableCurrency not implemented")
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
		FullMethod: "/workflow.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/ListWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListWorkflow(ctx, req.(*ListWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateWorkflow(ctx, req.(*WorkflowTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/DeleteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteWorkflow(ctx, req.(*WorkflowTask))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/CreateWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateWorkflowTask(ctx, req.(*CreateWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/ListWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListWorkflowTask(ctx, req.(*ListWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/GetWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetWorkflowTask(ctx, req.(*GetWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RequestDeleteWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RequestDeleteWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/RequestDeleteWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RequestDeleteWorkflowTask(ctx, req.(*GetWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadListWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadListWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/DownloadListWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListWorkflowTask(ctx, req.(*DownloadListWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMyWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMyWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/ListMyWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMyWorkflowTask(ctx, req.(*ListWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ValidateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ValidateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/ValidateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ValidateWorkflow(ctx, req.(*ValidateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GenerateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GenerateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/GenerateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GenerateWorkflow(ctx, req.(*GenerateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListCompanyWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompanyWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListCompanyWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/ListCompanyWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListCompanyWorkflow(ctx, req.(*ListCompanyWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetCompanyWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetCompanyWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/GetCompanyWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetCompanyWorkflow(ctx, req.(*GetCompanyWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateCompanyWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateCompanyWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/CreateCompanyWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateCompanyWorkflow(ctx, req.(*CreateCompanyWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetAvailableCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAvailableCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.service.v1.ApiService/GetAvailableCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAvailableCurrency(ctx, req.(*GetAvailableCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListWorkflow",
			Handler:    _ApiService_ListWorkflow_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _ApiService_CreateWorkflow_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _ApiService_DeleteWorkflow_Handler,
		},
		{
			MethodName: "CreateWorkflowTask",
			Handler:    _ApiService_CreateWorkflowTask_Handler,
		},
		{
			MethodName: "ListWorkflowTask",
			Handler:    _ApiService_ListWorkflowTask_Handler,
		},
		{
			MethodName: "GetWorkflowTask",
			Handler:    _ApiService_GetWorkflowTask_Handler,
		},
		{
			MethodName: "RequestDeleteWorkflowTask",
			Handler:    _ApiService_RequestDeleteWorkflowTask_Handler,
		},
		{
			MethodName: "DownloadListWorkflowTask",
			Handler:    _ApiService_DownloadListWorkflowTask_Handler,
		},
		{
			MethodName: "ListMyWorkflowTask",
			Handler:    _ApiService_ListMyWorkflowTask_Handler,
		},
		{
			MethodName: "ValidateWorkflow",
			Handler:    _ApiService_ValidateWorkflow_Handler,
		},
		{
			MethodName: "GenerateWorkflow",
			Handler:    _ApiService_GenerateWorkflow_Handler,
		},
		{
			MethodName: "ListCompanyWorkflow",
			Handler:    _ApiService_ListCompanyWorkflow_Handler,
		},
		{
			MethodName: "GetCompanyWorkflow",
			Handler:    _ApiService_GetCompanyWorkflow_Handler,
		},
		{
			MethodName: "CreateCompanyWorkflow",
			Handler:    _ApiService_CreateCompanyWorkflow_Handler,
		},
		{
			MethodName: "GetAvailableCurrency",
			Handler:    _ApiService_GetAvailableCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow_api.proto",
}
