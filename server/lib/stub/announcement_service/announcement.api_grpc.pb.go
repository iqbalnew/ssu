// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: announcement.api.proto

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
	DownloadListAnnouncementTasks(ctx context.Context, in *FileListAnnouncementTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	ListAnnouncementTask(ctx context.Context, in *ListAnnouncementTaskRequest, opts ...grpc.CallOption) (*ListAnnouncementResponse, error)
	ListAnnouncementTaskEV(ctx context.Context, in *ListAnnouncementTaskRequest, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error)
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	DeleteAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTask(ctx context.Context, in *CreateAnnouncementTaskRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTaskEV(ctx context.Context, in *CreateAnnouncementTaskRequestEV, opts ...grpc.CallOption) (*CreateAnnouncementResponseEV, error)
	GetAnnouncementTaskByID(ctx context.Context, in *GetByTaskID, opts ...grpc.CallOption) (*ListAnnouncementResponse, error)
	GetAnnouncementTaskByIDEV(ctx context.Context, in *GetByTaskIDEV, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error)
	DeleteAnnouncementTaskByIDEV(ctx context.Context, in *GetByTaskIDEV, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error)
	ListAnnouncement(ctx context.Context, in *ListAnnouncementDataRequest, opts ...grpc.CallOption) (*ListAnnouncementActiveResponse, error)
	ListAnnouncementEV(ctx context.Context, in *ListAnnouncementDataRequestEV, opts ...grpc.CallOption) (*ListAnnouncementActiveResponseEV, error)
	ListEventType(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEventTypeResponse, error)
	ListEventTypeEV(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEventTypeResponseEV, error)
	AnnouncementGraph(ctx context.Context, in *AnnouncementGraphReq, opts ...grpc.CallOption) (*AnnouncementGraphRes, error)
	AnnouncementActivate(ctx context.Context, in *ActivatedRequest, opts ...grpc.CallOption) (*ActivatedResponse, error)
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

func (c *apiServiceClient) DownloadListAnnouncementTasks(ctx context.Context, in *FileListAnnouncementTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/DownloadListAnnouncementTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncementTask(ctx context.Context, in *ListAnnouncementTaskRequest, opts ...grpc.CallOption) (*ListAnnouncementResponse, error) {
	out := new(ListAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncementTaskEV(ctx context.Context, in *ListAnnouncementTaskRequest, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error) {
	out := new(ListAnnouncementResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncementTaskEV", in, out, opts...)
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

func (c *apiServiceClient) DeleteAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/DeleteAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAnnouncementTask(ctx context.Context, in *CreateAnnouncementTaskRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/CreateAnnouncementTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAnnouncementTaskEV(ctx context.Context, in *CreateAnnouncementTaskRequestEV, opts ...grpc.CallOption) (*CreateAnnouncementResponseEV, error) {
	out := new(CreateAnnouncementResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/CreateAnnouncementTaskEV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetAnnouncementTaskByID(ctx context.Context, in *GetByTaskID, opts ...grpc.CallOption) (*ListAnnouncementResponse, error) {
	out := new(ListAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/GetAnnouncementTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetAnnouncementTaskByIDEV(ctx context.Context, in *GetByTaskIDEV, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error) {
	out := new(ListAnnouncementResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/GetAnnouncementTaskByIDEV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteAnnouncementTaskByIDEV(ctx context.Context, in *GetByTaskIDEV, opts ...grpc.CallOption) (*ListAnnouncementResponseEV, error) {
	out := new(ListAnnouncementResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/DeleteAnnouncementTaskByIDEV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncement(ctx context.Context, in *ListAnnouncementDataRequest, opts ...grpc.CallOption) (*ListAnnouncementActiveResponse, error) {
	out := new(ListAnnouncementActiveResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAnnouncementEV(ctx context.Context, in *ListAnnouncementDataRequestEV, opts ...grpc.CallOption) (*ListAnnouncementActiveResponseEV, error) {
	out := new(ListAnnouncementActiveResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListAnnouncementEV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListEventType(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEventTypeResponse, error) {
	out := new(ListEventTypeResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListEventType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListEventTypeEV(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListEventTypeResponseEV, error) {
	out := new(ListEventTypeResponseEV)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/ListEventTypeEV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AnnouncementGraph(ctx context.Context, in *AnnouncementGraphReq, opts ...grpc.CallOption) (*AnnouncementGraphRes, error) {
	out := new(AnnouncementGraphRes)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/AnnouncementGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AnnouncementActivate(ctx context.Context, in *ActivatedRequest, opts ...grpc.CallOption) (*ActivatedResponse, error) {
	out := new(ActivatedResponse)
	err := c.cc.Invoke(ctx, "/announcement.service.v1.ApiService/AnnouncementActivate", in, out, opts...)
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
	DownloadListAnnouncementTasks(context.Context, *FileListAnnouncementTaskRequest) (*httpbody.HttpBody, error)
	ListAnnouncementTask(context.Context, *ListAnnouncementTaskRequest) (*ListAnnouncementResponse, error)
	ListAnnouncementTaskEV(context.Context, *ListAnnouncementTaskRequest) (*ListAnnouncementResponseEV, error)
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	DeleteAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTask(context.Context, *CreateAnnouncementTaskRequest) (*CreateAnnouncementResponse, error)
	CreateAnnouncementTaskEV(context.Context, *CreateAnnouncementTaskRequestEV) (*CreateAnnouncementResponseEV, error)
	GetAnnouncementTaskByID(context.Context, *GetByTaskID) (*ListAnnouncementResponse, error)
	GetAnnouncementTaskByIDEV(context.Context, *GetByTaskIDEV) (*ListAnnouncementResponseEV, error)
	DeleteAnnouncementTaskByIDEV(context.Context, *GetByTaskIDEV) (*ListAnnouncementResponseEV, error)
	ListAnnouncement(context.Context, *ListAnnouncementDataRequest) (*ListAnnouncementActiveResponse, error)
	ListAnnouncementEV(context.Context, *ListAnnouncementDataRequestEV) (*ListAnnouncementActiveResponseEV, error)
	ListEventType(context.Context, *Empty) (*ListEventTypeResponse, error)
	ListEventTypeEV(context.Context, *Empty) (*ListEventTypeResponseEV, error)
	AnnouncementGraph(context.Context, *AnnouncementGraphReq) (*AnnouncementGraphRes, error)
	AnnouncementActivate(context.Context, *ActivatedRequest) (*ActivatedResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) DownloadListAnnouncementTasks(context.Context, *FileListAnnouncementTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListAnnouncementTasks not implemented")
}
func (UnimplementedApiServiceServer) ListAnnouncementTask(context.Context, *ListAnnouncementTaskRequest) (*ListAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncementTask not implemented")
}
func (UnimplementedApiServiceServer) ListAnnouncementTaskEV(context.Context, *ListAnnouncementTaskRequest) (*ListAnnouncementResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncementTaskEV not implemented")
}
func (UnimplementedApiServiceServer) CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedApiServiceServer) DeleteAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncement not implemented")
}
func (UnimplementedApiServiceServer) CreateAnnouncementTask(context.Context, *CreateAnnouncementTaskRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncementTask not implemented")
}
func (UnimplementedApiServiceServer) CreateAnnouncementTaskEV(context.Context, *CreateAnnouncementTaskRequestEV) (*CreateAnnouncementResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncementTaskEV not implemented")
}
func (UnimplementedApiServiceServer) GetAnnouncementTaskByID(context.Context, *GetByTaskID) (*ListAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncementTaskByID not implemented")
}
func (UnimplementedApiServiceServer) GetAnnouncementTaskByIDEV(context.Context, *GetByTaskIDEV) (*ListAnnouncementResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncementTaskByIDEV not implemented")
}
func (UnimplementedApiServiceServer) DeleteAnnouncementTaskByIDEV(context.Context, *GetByTaskIDEV) (*ListAnnouncementResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncementTaskByIDEV not implemented")
}
func (UnimplementedApiServiceServer) ListAnnouncement(context.Context, *ListAnnouncementDataRequest) (*ListAnnouncementActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncement not implemented")
}
func (UnimplementedApiServiceServer) ListAnnouncementEV(context.Context, *ListAnnouncementDataRequestEV) (*ListAnnouncementActiveResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnnouncementEV not implemented")
}
func (UnimplementedApiServiceServer) ListEventType(context.Context, *Empty) (*ListEventTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventType not implemented")
}
func (UnimplementedApiServiceServer) ListEventTypeEV(context.Context, *Empty) (*ListEventTypeResponseEV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventTypeEV not implemented")
}
func (UnimplementedApiServiceServer) AnnouncementGraph(context.Context, *AnnouncementGraphReq) (*AnnouncementGraphRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnouncementGraph not implemented")
}
func (UnimplementedApiServiceServer) AnnouncementActivate(context.Context, *ActivatedRequest) (*ActivatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnouncementActivate not implemented")
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

func _ApiService_DownloadListAnnouncementTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListAnnouncementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListAnnouncementTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/DownloadListAnnouncementTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListAnnouncementTasks(ctx, req.(*FileListAnnouncementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnnouncementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAnnouncementTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListAnnouncementTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAnnouncementTask(ctx, req.(*ListAnnouncementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncementTaskEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnnouncementTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAnnouncementTaskEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListAnnouncementTaskEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAnnouncementTaskEV(ctx, req.(*ListAnnouncementTaskRequest))
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

func _ApiService_DeleteAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/DeleteAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteAnnouncement(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAnnouncementTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementTaskRequest)
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
		return srv.(ApiServiceServer).CreateAnnouncementTask(ctx, req.(*CreateAnnouncementTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAnnouncementTaskEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementTaskRequestEV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAnnouncementTaskEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/CreateAnnouncementTaskEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAnnouncementTaskEV(ctx, req.(*CreateAnnouncementTaskRequestEV))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetAnnouncementTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAnnouncementTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/GetAnnouncementTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAnnouncementTaskByID(ctx, req.(*GetByTaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetAnnouncementTaskByIDEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTaskIDEV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAnnouncementTaskByIDEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/GetAnnouncementTaskByIDEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAnnouncementTaskByIDEV(ctx, req.(*GetByTaskIDEV))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteAnnouncementTaskByIDEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTaskIDEV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteAnnouncementTaskByIDEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/DeleteAnnouncementTaskByIDEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteAnnouncementTaskByIDEV(ctx, req.(*GetByTaskIDEV))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnnouncementDataRequest)
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
		return srv.(ApiServiceServer).ListAnnouncement(ctx, req.(*ListAnnouncementDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAnnouncementEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnnouncementDataRequestEV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAnnouncementEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListAnnouncementEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAnnouncementEV(ctx, req.(*ListAnnouncementDataRequestEV))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListEventType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(ApiServiceServer).ListEventType(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListEventTypeEV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListEventTypeEV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/ListEventTypeEV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListEventTypeEV(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AnnouncementGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnouncementGraphReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AnnouncementGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/AnnouncementGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AnnouncementGraph(ctx, req.(*AnnouncementGraphReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AnnouncementActivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AnnouncementActivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/announcement.service.v1.ApiService/AnnouncementActivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AnnouncementActivate(ctx, req.(*ActivatedRequest))
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
			MethodName: "DownloadListAnnouncementTasks",
			Handler:    _ApiService_DownloadListAnnouncementTasks_Handler,
		},
		{
			MethodName: "ListAnnouncementTask",
			Handler:    _ApiService_ListAnnouncementTask_Handler,
		},
		{
			MethodName: "ListAnnouncementTaskEV",
			Handler:    _ApiService_ListAnnouncementTaskEV_Handler,
		},
		{
			MethodName: "CreateAnnouncement",
			Handler:    _ApiService_CreateAnnouncement_Handler,
		},
		{
			MethodName: "DeleteAnnouncement",
			Handler:    _ApiService_DeleteAnnouncement_Handler,
		},
		{
			MethodName: "CreateAnnouncementTask",
			Handler:    _ApiService_CreateAnnouncementTask_Handler,
		},
		{
			MethodName: "CreateAnnouncementTaskEV",
			Handler:    _ApiService_CreateAnnouncementTaskEV_Handler,
		},
		{
			MethodName: "GetAnnouncementTaskByID",
			Handler:    _ApiService_GetAnnouncementTaskByID_Handler,
		},
		{
			MethodName: "GetAnnouncementTaskByIDEV",
			Handler:    _ApiService_GetAnnouncementTaskByIDEV_Handler,
		},
		{
			MethodName: "DeleteAnnouncementTaskByIDEV",
			Handler:    _ApiService_DeleteAnnouncementTaskByIDEV_Handler,
		},
		{
			MethodName: "ListAnnouncement",
			Handler:    _ApiService_ListAnnouncement_Handler,
		},
		{
			MethodName: "ListAnnouncementEV",
			Handler:    _ApiService_ListAnnouncementEV_Handler,
		},
		{
			MethodName: "ListEventType",
			Handler:    _ApiService_ListEventType_Handler,
		},
		{
			MethodName: "ListEventTypeEV",
			Handler:    _ApiService_ListEventTypeEV_Handler,
		},
		{
			MethodName: "AnnouncementGraph",
			Handler:    _ApiService_AnnouncementGraph_Handler,
		},
		{
			MethodName: "AnnouncementActivate",
			Handler:    _ApiService_AnnouncementActivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announcement.api.proto",
}
