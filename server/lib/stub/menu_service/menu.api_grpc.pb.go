// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: menu.api.proto

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
	SetTaskMenuAppearance(ctx context.Context, in *SetTaskMenuAppearanceReq, opts ...grpc.CallOption) (*SetTaskMenuAppearanceRes, error)
	GetListTaskMenuAppearance(ctx context.Context, in *GetListTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetListTaskMenuAppearanceRes, error)
	GetTaskMenuAppearance(ctx context.Context, in *GetTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetTaskMenuAppearanceRes, error)
	SaveMenuAppearance(ctx context.Context, in *SaveMenuAppearanceReq, opts ...grpc.CallOption) (*SaveMenuAppearanceRes, error)
	GetMenuAppearance(ctx context.Context, in *GetMenuAppearanceReq, opts ...grpc.CallOption) (*GetMenuAppearanceRes, error)
	GetMenuLicense(ctx context.Context, in *GetMenuAppearanceReq, opts ...grpc.CallOption) (*GetMenuLicenseRes, error)
	SetTaskMenuLicense(ctx context.Context, in *SetTaskMenuLicenseReq, opts ...grpc.CallOption) (*SetTaskMenuLicenseRes, error)
	SaveMenuLicense(ctx context.Context, in *SaveMenuLicenseReq, opts ...grpc.CallOption) (*SaveMenuLicenseRes, error)
	DeleteMenuLicense(ctx context.Context, in *SaveMenuLicenseReq, opts ...grpc.CallOption) (*SaveMenuLicenseRes, error)
	GetListTaskMenuLicense(ctx context.Context, in *GetListTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetListTaskMenuLicenseRes, error)
	DeleteTaskMenuLicense(ctx context.Context, in *GetTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetTaskMenuLicenseRes, error)
	GetTaskMenuLicense(ctx context.Context, in *GetTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetTaskMenuLicenseRes, error)
	ListMenuTask(ctx context.Context, in *ListMenuTaskRequest, opts ...grpc.CallOption) (*ListMenuResponse, error)
	ListModule(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModuleResponse, error)
	ListMenuDisable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListMenuDisableResponse, error)
	GetMenuAppearanceDisable(ctx context.Context, in *GetMenuAppearanceDisableReq, opts ...grpc.CallOption) (*GetMenuAppearanceDisableRes, error)
	FileMenuLicenseTask(ctx context.Context, in *FileMenuLicenseTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	FileMenuAppearanceTask(ctx context.Context, in *FileMenuAppearanceTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetMyTasksMenuLicense(ctx context.Context, in *GetListTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetListTaskMenuLicenseRes, error)
	GetMyTasksMenuAppearance(ctx context.Context, in *GetListTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetListTaskMenuAppearanceRes, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SetTaskMenuAppearance(ctx context.Context, in *SetTaskMenuAppearanceReq, opts ...grpc.CallOption) (*SetTaskMenuAppearanceRes, error) {
	out := new(SetTaskMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/SetTaskMenuAppearance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetListTaskMenuAppearance(ctx context.Context, in *GetListTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetListTaskMenuAppearanceRes, error) {
	out := new(GetListTaskMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetListTaskMenuAppearance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMenuAppearance(ctx context.Context, in *GetTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetTaskMenuAppearanceRes, error) {
	out := new(GetTaskMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetTaskMenuAppearance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SaveMenuAppearance(ctx context.Context, in *SaveMenuAppearanceReq, opts ...grpc.CallOption) (*SaveMenuAppearanceRes, error) {
	out := new(SaveMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/SaveMenuAppearance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMenuAppearance(ctx context.Context, in *GetMenuAppearanceReq, opts ...grpc.CallOption) (*GetMenuAppearanceRes, error) {
	out := new(GetMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetMenuAppearance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMenuLicense(ctx context.Context, in *GetMenuAppearanceReq, opts ...grpc.CallOption) (*GetMenuLicenseRes, error) {
	out := new(GetMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SetTaskMenuLicense(ctx context.Context, in *SetTaskMenuLicenseReq, opts ...grpc.CallOption) (*SetTaskMenuLicenseRes, error) {
	out := new(SetTaskMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/SetTaskMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SaveMenuLicense(ctx context.Context, in *SaveMenuLicenseReq, opts ...grpc.CallOption) (*SaveMenuLicenseRes, error) {
	out := new(SaveMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/SaveMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteMenuLicense(ctx context.Context, in *SaveMenuLicenseReq, opts ...grpc.CallOption) (*SaveMenuLicenseRes, error) {
	out := new(SaveMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/DeleteMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetListTaskMenuLicense(ctx context.Context, in *GetListTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetListTaskMenuLicenseRes, error) {
	out := new(GetListTaskMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetListTaskMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteTaskMenuLicense(ctx context.Context, in *GetTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetTaskMenuLicenseRes, error) {
	out := new(GetTaskMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/DeleteTaskMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetTaskMenuLicense(ctx context.Context, in *GetTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetTaskMenuLicenseRes, error) {
	out := new(GetTaskMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetTaskMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMenuTask(ctx context.Context, in *ListMenuTaskRequest, opts ...grpc.CallOption) (*ListMenuResponse, error) {
	out := new(ListMenuResponse)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/ListMenuTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListModule(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListModuleResponse, error) {
	out := new(ListModuleResponse)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/ListModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListMenuDisable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListMenuDisableResponse, error) {
	out := new(ListMenuDisableResponse)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/ListMenuDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMenuAppearanceDisable(ctx context.Context, in *GetMenuAppearanceDisableReq, opts ...grpc.CallOption) (*GetMenuAppearanceDisableRes, error) {
	out := new(GetMenuAppearanceDisableRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetMenuAppearanceDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) FileMenuLicenseTask(ctx context.Context, in *FileMenuLicenseTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/FileMenuLicenseTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) FileMenuAppearanceTask(ctx context.Context, in *FileMenuAppearanceTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/FileMenuAppearanceTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasksMenuLicense(ctx context.Context, in *GetListTaskMenuLicenseReq, opts ...grpc.CallOption) (*GetListTaskMenuLicenseRes, error) {
	out := new(GetListTaskMenuLicenseRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetMyTasksMenuLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetMyTasksMenuAppearance(ctx context.Context, in *GetListTaskMenuAppearanceReq, opts ...grpc.CallOption) (*GetListTaskMenuAppearanceRes, error) {
	out := new(GetListTaskMenuAppearanceRes)
	err := c.cc.Invoke(ctx, "/menu.service.v1.ApiService/GetMyTasksMenuAppearance", in, out, opts...)
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
	SetTaskMenuAppearance(context.Context, *SetTaskMenuAppearanceReq) (*SetTaskMenuAppearanceRes, error)
	GetListTaskMenuAppearance(context.Context, *GetListTaskMenuAppearanceReq) (*GetListTaskMenuAppearanceRes, error)
	GetTaskMenuAppearance(context.Context, *GetTaskMenuAppearanceReq) (*GetTaskMenuAppearanceRes, error)
	SaveMenuAppearance(context.Context, *SaveMenuAppearanceReq) (*SaveMenuAppearanceRes, error)
	GetMenuAppearance(context.Context, *GetMenuAppearanceReq) (*GetMenuAppearanceRes, error)
	GetMenuLicense(context.Context, *GetMenuAppearanceReq) (*GetMenuLicenseRes, error)
	SetTaskMenuLicense(context.Context, *SetTaskMenuLicenseReq) (*SetTaskMenuLicenseRes, error)
	SaveMenuLicense(context.Context, *SaveMenuLicenseReq) (*SaveMenuLicenseRes, error)
	DeleteMenuLicense(context.Context, *SaveMenuLicenseReq) (*SaveMenuLicenseRes, error)
	GetListTaskMenuLicense(context.Context, *GetListTaskMenuLicenseReq) (*GetListTaskMenuLicenseRes, error)
	DeleteTaskMenuLicense(context.Context, *GetTaskMenuLicenseReq) (*GetTaskMenuLicenseRes, error)
	GetTaskMenuLicense(context.Context, *GetTaskMenuLicenseReq) (*GetTaskMenuLicenseRes, error)
	ListMenuTask(context.Context, *ListMenuTaskRequest) (*ListMenuResponse, error)
	ListModule(context.Context, *Empty) (*ListModuleResponse, error)
	ListMenuDisable(context.Context, *Empty) (*ListMenuDisableResponse, error)
	GetMenuAppearanceDisable(context.Context, *GetMenuAppearanceDisableReq) (*GetMenuAppearanceDisableRes, error)
	FileMenuLicenseTask(context.Context, *FileMenuLicenseTaskRequest) (*httpbody.HttpBody, error)
	FileMenuAppearanceTask(context.Context, *FileMenuAppearanceTaskRequest) (*httpbody.HttpBody, error)
	GetMyTasksMenuLicense(context.Context, *GetListTaskMenuLicenseReq) (*GetListTaskMenuLicenseRes, error)
	GetMyTasksMenuAppearance(context.Context, *GetListTaskMenuAppearanceReq) (*GetListTaskMenuAppearanceRes, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) SetTaskMenuAppearance(context.Context, *SetTaskMenuAppearanceReq) (*SetTaskMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTaskMenuAppearance not implemented")
}
func (UnimplementedApiServiceServer) GetListTaskMenuAppearance(context.Context, *GetListTaskMenuAppearanceReq) (*GetListTaskMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTaskMenuAppearance not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMenuAppearance(context.Context, *GetTaskMenuAppearanceReq) (*GetTaskMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMenuAppearance not implemented")
}
func (UnimplementedApiServiceServer) SaveMenuAppearance(context.Context, *SaveMenuAppearanceReq) (*SaveMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMenuAppearance not implemented")
}
func (UnimplementedApiServiceServer) GetMenuAppearance(context.Context, *GetMenuAppearanceReq) (*GetMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuAppearance not implemented")
}
func (UnimplementedApiServiceServer) GetMenuLicense(context.Context, *GetMenuAppearanceReq) (*GetMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) SetTaskMenuLicense(context.Context, *SetTaskMenuLicenseReq) (*SetTaskMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTaskMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) SaveMenuLicense(context.Context, *SaveMenuLicenseReq) (*SaveMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) DeleteMenuLicense(context.Context, *SaveMenuLicenseReq) (*SaveMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) GetListTaskMenuLicense(context.Context, *GetListTaskMenuLicenseReq) (*GetListTaskMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTaskMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) DeleteTaskMenuLicense(context.Context, *GetTaskMenuLicenseReq) (*GetTaskMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) GetTaskMenuLicense(context.Context, *GetTaskMenuLicenseReq) (*GetTaskMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) ListMenuTask(context.Context, *ListMenuTaskRequest) (*ListMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenuTask not implemented")
}
func (UnimplementedApiServiceServer) ListModule(context.Context, *Empty) (*ListModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModule not implemented")
}
func (UnimplementedApiServiceServer) ListMenuDisable(context.Context, *Empty) (*ListMenuDisableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenuDisable not implemented")
}
func (UnimplementedApiServiceServer) GetMenuAppearanceDisable(context.Context, *GetMenuAppearanceDisableReq) (*GetMenuAppearanceDisableRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuAppearanceDisable not implemented")
}
func (UnimplementedApiServiceServer) FileMenuLicenseTask(context.Context, *FileMenuLicenseTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileMenuLicenseTask not implemented")
}
func (UnimplementedApiServiceServer) FileMenuAppearanceTask(context.Context, *FileMenuAppearanceTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileMenuAppearanceTask not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasksMenuLicense(context.Context, *GetListTaskMenuLicenseReq) (*GetListTaskMenuLicenseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasksMenuLicense not implemented")
}
func (UnimplementedApiServiceServer) GetMyTasksMenuAppearance(context.Context, *GetListTaskMenuAppearanceReq) (*GetListTaskMenuAppearanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasksMenuAppearance not implemented")
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
		FullMethod: "/menu.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SetTaskMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SetTaskMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/SetTaskMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SetTaskMenuAppearance(ctx, req.(*SetTaskMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetListTaskMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTaskMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetListTaskMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetListTaskMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetListTaskMenuAppearance(ctx, req.(*GetListTaskMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetTaskMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMenuAppearance(ctx, req.(*GetTaskMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SaveMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SaveMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/SaveMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SaveMenuAppearance(ctx, req.(*SaveMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMenuAppearance(ctx, req.(*GetMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMenuLicense(ctx, req.(*GetMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SetTaskMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SetTaskMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/SetTaskMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SetTaskMenuLicense(ctx, req.(*SetTaskMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SaveMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SaveMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/SaveMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SaveMenuLicense(ctx, req.(*SaveMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/DeleteMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteMenuLicense(ctx, req.(*SaveMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetListTaskMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTaskMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetListTaskMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetListTaskMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetListTaskMenuLicense(ctx, req.(*GetListTaskMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteTaskMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteTaskMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/DeleteTaskMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteTaskMenuLicense(ctx, req.(*GetTaskMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetTaskMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetTaskMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetTaskMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetTaskMenuLicense(ctx, req.(*GetTaskMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMenuTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMenuTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMenuTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/ListMenuTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMenuTask(ctx, req.(*ListMenuTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/ListModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListModule(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListMenuDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListMenuDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/ListMenuDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListMenuDisable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMenuAppearanceDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuAppearanceDisableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMenuAppearanceDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetMenuAppearanceDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMenuAppearanceDisable(ctx, req.(*GetMenuAppearanceDisableReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_FileMenuLicenseTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMenuLicenseTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).FileMenuLicenseTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/FileMenuLicenseTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).FileMenuLicenseTask(ctx, req.(*FileMenuLicenseTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_FileMenuAppearanceTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMenuAppearanceTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).FileMenuAppearanceTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/FileMenuAppearanceTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).FileMenuAppearanceTask(ctx, req.(*FileMenuAppearanceTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasksMenuLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTaskMenuLicenseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasksMenuLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetMyTasksMenuLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasksMenuLicense(ctx, req.(*GetListTaskMenuLicenseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetMyTasksMenuAppearance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTaskMenuAppearanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetMyTasksMenuAppearance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.service.v1.ApiService/GetMyTasksMenuAppearance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetMyTasksMenuAppearance(ctx, req.(*GetListTaskMenuAppearanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "SetTaskMenuAppearance",
			Handler:    _ApiService_SetTaskMenuAppearance_Handler,
		},
		{
			MethodName: "GetListTaskMenuAppearance",
			Handler:    _ApiService_GetListTaskMenuAppearance_Handler,
		},
		{
			MethodName: "GetTaskMenuAppearance",
			Handler:    _ApiService_GetTaskMenuAppearance_Handler,
		},
		{
			MethodName: "SaveMenuAppearance",
			Handler:    _ApiService_SaveMenuAppearance_Handler,
		},
		{
			MethodName: "GetMenuAppearance",
			Handler:    _ApiService_GetMenuAppearance_Handler,
		},
		{
			MethodName: "GetMenuLicense",
			Handler:    _ApiService_GetMenuLicense_Handler,
		},
		{
			MethodName: "SetTaskMenuLicense",
			Handler:    _ApiService_SetTaskMenuLicense_Handler,
		},
		{
			MethodName: "SaveMenuLicense",
			Handler:    _ApiService_SaveMenuLicense_Handler,
		},
		{
			MethodName: "DeleteMenuLicense",
			Handler:    _ApiService_DeleteMenuLicense_Handler,
		},
		{
			MethodName: "GetListTaskMenuLicense",
			Handler:    _ApiService_GetListTaskMenuLicense_Handler,
		},
		{
			MethodName: "DeleteTaskMenuLicense",
			Handler:    _ApiService_DeleteTaskMenuLicense_Handler,
		},
		{
			MethodName: "GetTaskMenuLicense",
			Handler:    _ApiService_GetTaskMenuLicense_Handler,
		},
		{
			MethodName: "ListMenuTask",
			Handler:    _ApiService_ListMenuTask_Handler,
		},
		{
			MethodName: "ListModule",
			Handler:    _ApiService_ListModule_Handler,
		},
		{
			MethodName: "ListMenuDisable",
			Handler:    _ApiService_ListMenuDisable_Handler,
		},
		{
			MethodName: "GetMenuAppearanceDisable",
			Handler:    _ApiService_GetMenuAppearanceDisable_Handler,
		},
		{
			MethodName: "FileMenuLicenseTask",
			Handler:    _ApiService_FileMenuLicenseTask_Handler,
		},
		{
			MethodName: "FileMenuAppearanceTask",
			Handler:    _ApiService_FileMenuAppearanceTask_Handler,
		},
		{
			MethodName: "GetMyTasksMenuLicense",
			Handler:    _ApiService_GetMyTasksMenuLicense_Handler,
		},
		{
			MethodName: "GetMyTasksMenuAppearance",
			Handler:    _ApiService_GetMyTasksMenuAppearance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu.api.proto",
}
