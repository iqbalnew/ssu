// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.16.1
// source: liquidity_api.proto

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
	ListLiquidityTask(ctx context.Context, in *ListTaskLiquidityRequest, opts ...grpc.CallOption) (*ListLiquidityTaskResponse, error)
	DownloadListLiquidityTask(ctx context.Context, in *DownloadListTaskLiquidityRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	DetailLiquidityTask(ctx context.Context, in *DetailLiquidityTaskRequest, opts ...grpc.CallOption) (*DetailLiquidityTaskResponse, error)
	CreateLiquidityTask(ctx context.Context, in *CreateTaskLiquidityRequest, opts ...grpc.CallOption) (*CreateTaskLiquidityResponse, error)
	DeleteLiquidityTask(ctx context.Context, in *DetailLiquidityTaskRequest, opts ...grpc.CallOption) (*DetailLiquidityTaskResponse, error)
	GetMyTasksCreatedBy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArrayString, error)
	GetMyTasksApprovedBy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArrayString, error)
	CreateLiquidity(ctx context.Context, in *CreateLiquidityRequest, opts ...grpc.CallOption) (*CreateTaskLiquidityResponse, error)
	DeleteLiquidity(ctx context.Context, in *DeleteLiquidityRequest, opts ...grpc.CallOption) (*DeleteLiquidityResponse, error)
	GetListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataResponse, error)
	GetListTBAValue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTBAValueResponse, error)
	CreateLiquiditySchedules(ctx context.Context, in *CreateLiquiditySchedulesRequest, opts ...grpc.CallOption) (*CreateLiquiditySchedulesResponse, error)
	RunDailySchedule(ctx context.Context, in *RunDailyScheduleRequest, opts ...grpc.CallOption) (*RunDailyScheduleResponse, error)
	RunLiquidityScheme(ctx context.Context, in *RunLiquidityTaskRequest, opts ...grpc.CallOption) (*RunLiquidityTaskResponse, error)
	ValidateDate(ctx context.Context, in *ValidateDateRequest, opts ...grpc.CallOption) (*ValidateDateResponse, error)
	ListLiquidityAuthorize(ctx context.Context, in *ListTaskLiquidityRequest, opts ...grpc.CallOption) (*ListLiquidityTaskResponse, error)
	TaskAction(ctx context.Context, in *TaskActionRequest, opts ...grpc.CallOption) (*TaskActionResponse, error)
	DeactivateLiquidityScheme(ctx context.Context, in *DeactivateLiquiditySchemeRequest, opts ...grpc.CallOption) (*DeactivateLiquiditySchemeResponse, error)
	ReactivateLiquidityScheme(ctx context.Context, in *DeactivateLiquiditySchemeRequest, opts ...grpc.CallOption) (*DeactivateLiquiditySchemeResponse, error)
	ExecCashPooling(ctx context.Context, in *ExecCashPoolingRequest, opts ...grpc.CallOption) (*ExecCashPoolingResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListLiquidityTask(ctx context.Context, in *ListTaskLiquidityRequest, opts ...grpc.CallOption) (*ListLiquidityTaskResponse, error) {
	out := new(ListLiquidityTaskResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/ListLiquidityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadListLiquidityTask(ctx context.Context, in *DownloadListTaskLiquidityRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/DownloadListLiquidityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DetailLiquidityTask(ctx context.Context, in *DetailLiquidityTaskRequest, opts ...grpc.CallOption) (*DetailLiquidityTaskResponse, error) {
	out := new(DetailLiquidityTaskResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/DetailLiquidityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateLiquidityTask(ctx context.Context, in *CreateTaskLiquidityRequest, opts ...grpc.CallOption) (*CreateTaskLiquidityResponse, error) {
	out := new(CreateTaskLiquidityResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/CreateLiquidityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteLiquidityTask(ctx context.Context, in *DetailLiquidityTaskRequest, opts ...grpc.CallOption) (*DetailLiquidityTaskResponse, error) {
	out := new(DetailLiquidityTaskResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/DeleteLiquidityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasksCreatedBy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArrayString, error) {
	out := new(ArrayString)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/GetMyTasksCreatedBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasksApprovedBy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArrayString, error) {
	out := new(ArrayString)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/GetMyTasksApprovedBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateLiquidity(ctx context.Context, in *CreateLiquidityRequest, opts ...grpc.CallOption) (*CreateTaskLiquidityResponse, error) {
	out := new(CreateTaskLiquidityResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/CreateLiquidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteLiquidity(ctx context.Context, in *DeleteLiquidityRequest, opts ...grpc.CallOption) (*DeleteLiquidityResponse, error) {
	out := new(DeleteLiquidityResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/DeleteLiquidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetListData(ctx context.Context, in *ListDataRequest, opts ...grpc.CallOption) (*ListDataResponse, error) {
	out := new(ListDataResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/GetListData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetListTBAValue(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTBAValueResponse, error) {
	out := new(ListTBAValueResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/GetListTBAValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateLiquiditySchedules(ctx context.Context, in *CreateLiquiditySchedulesRequest, opts ...grpc.CallOption) (*CreateLiquiditySchedulesResponse, error) {
	out := new(CreateLiquiditySchedulesResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/CreateLiquiditySchedules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RunDailySchedule(ctx context.Context, in *RunDailyScheduleRequest, opts ...grpc.CallOption) (*RunDailyScheduleResponse, error) {
	out := new(RunDailyScheduleResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/RunDailySchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RunLiquidityScheme(ctx context.Context, in *RunLiquidityTaskRequest, opts ...grpc.CallOption) (*RunLiquidityTaskResponse, error) {
	out := new(RunLiquidityTaskResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/RunLiquidityScheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ValidateDate(ctx context.Context, in *ValidateDateRequest, opts ...grpc.CallOption) (*ValidateDateResponse, error) {
	out := new(ValidateDateResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/ValidateDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListLiquidityAuthorize(ctx context.Context, in *ListTaskLiquidityRequest, opts ...grpc.CallOption) (*ListLiquidityTaskResponse, error) {
	out := new(ListLiquidityTaskResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/ListLiquidityAuthorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) TaskAction(ctx context.Context, in *TaskActionRequest, opts ...grpc.CallOption) (*TaskActionResponse, error) {
	out := new(TaskActionResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/TaskAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeactivateLiquidityScheme(ctx context.Context, in *DeactivateLiquiditySchemeRequest, opts ...grpc.CallOption) (*DeactivateLiquiditySchemeResponse, error) {
	out := new(DeactivateLiquiditySchemeResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/DeactivateLiquidityScheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ReactivateLiquidityScheme(ctx context.Context, in *DeactivateLiquiditySchemeRequest, opts ...grpc.CallOption) (*DeactivateLiquiditySchemeResponse, error) {
	out := new(DeactivateLiquiditySchemeResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/ReactivateLiquidityScheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ExecCashPooling(ctx context.Context, in *ExecCashPoolingRequest, opts ...grpc.CallOption) (*ExecCashPoolingResponse, error) {
	out := new(ExecCashPoolingResponse)
	err := c.cc.Invoke(ctx, "/liquidity.service.v1.ApiService/ExecCashPooling", in, out, opts...)
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
	ListLiquidityTask(context.Context, *ListTaskLiquidityRequest) (*ListLiquidityTaskResponse, error)
	DownloadListLiquidityTask(context.Context, *DownloadListTaskLiquidityRequest) (*httpbody.HttpBody, error)
	DetailLiquidityTask(context.Context, *DetailLiquidityTaskRequest) (*DetailLiquidityTaskResponse, error)
	CreateLiquidityTask(context.Context, *CreateTaskLiquidityRequest) (*CreateTaskLiquidityResponse, error)
	DeleteLiquidityTask(context.Context, *DetailLiquidityTaskRequest) (*DetailLiquidityTaskResponse, error)
	GetMyTasksCreatedBy(context.Context, *Empty) (*ArrayString, error)
	GetMyTasksApprovedBy(context.Context, *Empty) (*ArrayString, error)
	CreateLiquidity(context.Context, *CreateLiquidityRequest) (*CreateTaskLiquidityResponse, error)
	DeleteLiquidity(context.Context, *DeleteLiquidityRequest) (*DeleteLiquidityResponse, error)
	GetListData(context.Context, *ListDataRequest) (*ListDataResponse, error)
	GetListTBAValue(context.Context, *Empty) (*ListTBAValueResponse, error)
	CreateLiquiditySchedules(context.Context, *CreateLiquiditySchedulesRequest) (*CreateLiquiditySchedulesResponse, error)
	RunDailySchedule(context.Context, *RunDailyScheduleRequest) (*RunDailyScheduleResponse, error)
	RunLiquidityScheme(context.Context, *RunLiquidityTaskRequest) (*RunLiquidityTaskResponse, error)
	ValidateDate(context.Context, *ValidateDateRequest) (*ValidateDateResponse, error)
	ListLiquidityAuthorize(context.Context, *ListTaskLiquidityRequest) (*ListLiquidityTaskResponse, error)
	TaskAction(context.Context, *TaskActionRequest) (*TaskActionResponse, error)
	DeactivateLiquidityScheme(context.Context, *DeactivateLiquiditySchemeRequest) (*DeactivateLiquiditySchemeResponse, error)
	ReactivateLiquidityScheme(context.Context, *DeactivateLiquiditySchemeRequest) (*DeactivateLiquiditySchemeResponse, error)
	ExecCashPooling(context.Context, *ExecCashPoolingRequest) (*ExecCashPoolingResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListLiquidityTask(context.Context, *ListTaskLiquidityRequest) (*ListLiquidityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiquidityTask not implemented")
}
func (UnimplementedApiServiceServer) DownloadListLiquidityTask(context.Context, *DownloadListTaskLiquidityRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListLiquidityTask not implemented")
}
func (UnimplementedApiServiceServer) DetailLiquidityTask(context.Context, *DetailLiquidityTaskRequest) (*DetailLiquidityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailLiquidityTask not implemented")
}
func (UnimplementedApiServiceServer) CreateLiquidityTask(context.Context, *CreateTaskLiquidityRequest) (*CreateTaskLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiquidityTask not implemented")
}
func (UnimplementedApiServiceServer) DeleteLiquidityTask(context.Context, *DetailLiquidityTaskRequest) (*DetailLiquidityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiquidityTask not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasksCreatedBy(context.Context, *Empty) (*ArrayString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasksCreatedBy not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasksApprovedBy(context.Context, *Empty) (*ArrayString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasksApprovedBy not implemented")
}
func (UnimplementedApiServiceServer) CreateLiquidity(context.Context, *CreateLiquidityRequest) (*CreateTaskLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiquidity not implemented")
}
func (UnimplementedApiServiceServer) DeleteLiquidity(context.Context, *DeleteLiquidityRequest) (*DeleteLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiquidity not implemented")
}
func (UnimplementedApiServiceServer) GetListData(context.Context, *ListDataRequest) (*ListDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListData not implemented")
}
func (UnimplementedApiServiceServer) GetListTBAValue(context.Context, *Empty) (*ListTBAValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTBAValue not implemented")
}
func (UnimplementedApiServiceServer) CreateLiquiditySchedules(context.Context, *CreateLiquiditySchedulesRequest) (*CreateLiquiditySchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiquiditySchedules not implemented")
}
func (UnimplementedApiServiceServer) RunDailySchedule(context.Context, *RunDailyScheduleRequest) (*RunDailyScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunDailySchedule not implemented")
}
func (UnimplementedApiServiceServer) RunLiquidityScheme(context.Context, *RunLiquidityTaskRequest) (*RunLiquidityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunLiquidityScheme not implemented")
}
func (UnimplementedApiServiceServer) ValidateDate(context.Context, *ValidateDateRequest) (*ValidateDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDate not implemented")
}
func (UnimplementedApiServiceServer) ListLiquidityAuthorize(context.Context, *ListTaskLiquidityRequest) (*ListLiquidityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiquidityAuthorize not implemented")
}
func (UnimplementedApiServiceServer) TaskAction(context.Context, *TaskActionRequest) (*TaskActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskAction not implemented")
}
func (UnimplementedApiServiceServer) DeactivateLiquidityScheme(context.Context, *DeactivateLiquiditySchemeRequest) (*DeactivateLiquiditySchemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateLiquidityScheme not implemented")
}
func (UnimplementedApiServiceServer) ReactivateLiquidityScheme(context.Context, *DeactivateLiquiditySchemeRequest) (*DeactivateLiquiditySchemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReactivateLiquidityScheme not implemented")
}
func (UnimplementedApiServiceServer) ExecCashPooling(context.Context, *ExecCashPoolingRequest) (*ExecCashPoolingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecCashPooling not implemented")
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
		FullMethod: "/liquidity.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListLiquidityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListLiquidityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/ListLiquidityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListLiquidityTask(ctx, req.(*ListTaskLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadListLiquidityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadListTaskLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListLiquidityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/DownloadListLiquidityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListLiquidityTask(ctx, req.(*DownloadListTaskLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DetailLiquidityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailLiquidityTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DetailLiquidityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/DetailLiquidityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DetailLiquidityTask(ctx, req.(*DetailLiquidityTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateLiquidityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateLiquidityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/CreateLiquidityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateLiquidityTask(ctx, req.(*CreateTaskLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteLiquidityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailLiquidityTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteLiquidityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/DeleteLiquidityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteLiquidityTask(ctx, req.(*DetailLiquidityTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasksCreatedBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasksCreatedBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/GetMyTasksCreatedBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasksCreatedBy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasksApprovedBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasksApprovedBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/GetMyTasksApprovedBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasksApprovedBy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateLiquidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateLiquidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/CreateLiquidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateLiquidity(ctx, req.(*CreateLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteLiquidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteLiquidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/DeleteLiquidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteLiquidity(ctx, req.(*DeleteLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/GetListData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetListData(ctx, req.(*ListDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetListTBAValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetListTBAValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/GetListTBAValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetListTBAValue(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateLiquiditySchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLiquiditySchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateLiquiditySchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/CreateLiquiditySchedules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateLiquiditySchedules(ctx, req.(*CreateLiquiditySchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RunDailySchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunDailyScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RunDailySchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/RunDailySchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RunDailySchedule(ctx, req.(*RunDailyScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RunLiquidityScheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunLiquidityTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RunLiquidityScheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/RunLiquidityScheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RunLiquidityScheme(ctx, req.(*RunLiquidityTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ValidateDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ValidateDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/ValidateDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ValidateDate(ctx, req.(*ValidateDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListLiquidityAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListLiquidityAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/ListLiquidityAuthorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListLiquidityAuthorize(ctx, req.(*ListTaskLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_TaskAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).TaskAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/TaskAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).TaskAction(ctx, req.(*TaskActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeactivateLiquidityScheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateLiquiditySchemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeactivateLiquidityScheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/DeactivateLiquidityScheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeactivateLiquidityScheme(ctx, req.(*DeactivateLiquiditySchemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ReactivateLiquidityScheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateLiquiditySchemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ReactivateLiquidityScheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/ReactivateLiquidityScheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ReactivateLiquidityScheme(ctx, req.(*DeactivateLiquiditySchemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ExecCashPooling_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecCashPoolingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ExecCashPooling(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liquidity.service.v1.ApiService/ExecCashPooling",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ExecCashPooling(ctx, req.(*ExecCashPoolingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "liquidity.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListLiquidityTask",
			Handler:    _ApiService_ListLiquidityTask_Handler,
		},
		{
			MethodName: "DownloadListLiquidityTask",
			Handler:    _ApiService_DownloadListLiquidityTask_Handler,
		},
		{
			MethodName: "DetailLiquidityTask",
			Handler:    _ApiService_DetailLiquidityTask_Handler,
		},
		{
			MethodName: "CreateLiquidityTask",
			Handler:    _ApiService_CreateLiquidityTask_Handler,
		},
		{
			MethodName: "DeleteLiquidityTask",
			Handler:    _ApiService_DeleteLiquidityTask_Handler,
		},
		{
			MethodName: "GetMyTasksCreatedBy",
			Handler:    _ApiService_GetMyTasksCreatedBy_Handler,
		},
		{
			MethodName: "GetMyTasksApprovedBy",
			Handler:    _ApiService_GetMyTasksApprovedBy_Handler,
		},
		{
			MethodName: "CreateLiquidity",
			Handler:    _ApiService_CreateLiquidity_Handler,
		},
		{
			MethodName: "DeleteLiquidity",
			Handler:    _ApiService_DeleteLiquidity_Handler,
		},
		{
			MethodName: "GetListData",
			Handler:    _ApiService_GetListData_Handler,
		},
		{
			MethodName: "GetListTBAValue",
			Handler:    _ApiService_GetListTBAValue_Handler,
		},
		{
			MethodName: "CreateLiquiditySchedules",
			Handler:    _ApiService_CreateLiquiditySchedules_Handler,
		},
		{
			MethodName: "RunDailySchedule",
			Handler:    _ApiService_RunDailySchedule_Handler,
		},
		{
			MethodName: "RunLiquidityScheme",
			Handler:    _ApiService_RunLiquidityScheme_Handler,
		},
		{
			MethodName: "ValidateDate",
			Handler:    _ApiService_ValidateDate_Handler,
		},
		{
			MethodName: "ListLiquidityAuthorize",
			Handler:    _ApiService_ListLiquidityAuthorize_Handler,
		},
		{
			MethodName: "TaskAction",
			Handler:    _ApiService_TaskAction_Handler,
		},
		{
			MethodName: "DeactivateLiquidityScheme",
			Handler:    _ApiService_DeactivateLiquidityScheme_Handler,
		},
		{
			MethodName: "ReactivateLiquidityScheme",
			Handler:    _ApiService_ReactivateLiquidityScheme_Handler,
		},
		{
			MethodName: "ExecCashPooling",
			Handler:    _ApiService_ExecCashPooling_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liquidity_api.proto",
}
