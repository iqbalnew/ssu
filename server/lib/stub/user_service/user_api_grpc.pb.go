// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: user_api.proto

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
	GetUsers(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	DownloadUsers(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	GetUserTasks(ctx context.Context, in *ListUserWithTaskRequest, opts ...grpc.CallOption) (*ListUserWithTaskResponse, error)
	GetUserTaskByID(ctx context.Context, in *GetTaskByIDReq, opts ...grpc.CallOption) (*GetUserWithTaskRes, error)
	DownloadUserTasks(ctx context.Context, in *ListUserWithTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateUserTask(ctx context.Context, in *CreateUserTaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListUserGroup(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRes, error)
	BRICaMSgetCustomer(ctx context.Context, in *BricamsGetCustomerReq, opts ...grpc.CallOption) (*BricamsGetCustomerRes, error)
	BRICamsGetAllUser(ctx context.Context, in *BricamsGetAddonsUserReq, opts ...grpc.CallOption) (*BricamsGetAddonsUserRes, error)
	BRICamsGetUserByUsername(ctx context.Context, in *BricamsGetAddonsUserByUsernameReq, opts ...grpc.CallOption) (*BricamsGetAddonsUserByUsernameRes, error)
	BRICaMSsvcGetUserList(ctx context.Context, in *BricamsGetAddonsUserReq, opts ...grpc.CallOption) (*BRICaMSSvcUserListRes, error)
	BRICaMSsvcGetUserByUsername(ctx context.Context, in *BricamsGetAddonsUserByUsernameReq, opts ...grpc.CallOption) (*BRICaMSSvcUserRes, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUsers(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadUsers(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/DownloadUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUserTasks(ctx context.Context, in *ListUserWithTaskRequest, opts ...grpc.CallOption) (*ListUserWithTaskResponse, error) {
	out := new(ListUserWithTaskResponse)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/GetUserTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUserTaskByID(ctx context.Context, in *GetTaskByIDReq, opts ...grpc.CallOption) (*GetUserWithTaskRes, error) {
	out := new(GetUserWithTaskRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/GetUserTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadUserTasks(ctx context.Context, in *ListUserWithTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/DownloadUserTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateUserTask(ctx context.Context, in *CreateUserTaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/CreateUserTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListUserGroup(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserRes, error) {
	out := new(ListUserRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/ListUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) BRICaMSgetCustomer(ctx context.Context, in *BricamsGetCustomerReq, opts ...grpc.CallOption) (*BricamsGetCustomerRes, error) {
	out := new(BricamsGetCustomerRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/BRICaMSgetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) BRICamsGetAllUser(ctx context.Context, in *BricamsGetAddonsUserReq, opts ...grpc.CallOption) (*BricamsGetAddonsUserRes, error) {
	out := new(BricamsGetAddonsUserRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/BRICamsGetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) BRICamsGetUserByUsername(ctx context.Context, in *BricamsGetAddonsUserByUsernameReq, opts ...grpc.CallOption) (*BricamsGetAddonsUserByUsernameRes, error) {
	out := new(BricamsGetAddonsUserByUsernameRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/BRICamsGetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) BRICaMSsvcGetUserList(ctx context.Context, in *BricamsGetAddonsUserReq, opts ...grpc.CallOption) (*BRICaMSSvcUserListRes, error) {
	out := new(BRICaMSSvcUserListRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/BRICaMSsvcGetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) BRICaMSsvcGetUserByUsername(ctx context.Context, in *BricamsGetAddonsUserByUsernameReq, opts ...grpc.CallOption) (*BRICaMSSvcUserRes, error) {
	out := new(BRICaMSSvcUserRes)
	err := c.cc.Invoke(ctx, "/user.service.v1.ApiService/BRICaMSsvcGetUserByUsername", in, out, opts...)
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
	GetUsers(context.Context, *CommonRequest) (*ListUserResponse, error)
	DownloadUsers(context.Context, *CommonRequest) (*httpbody.HttpBody, error)
	GetUserTasks(context.Context, *ListUserWithTaskRequest) (*ListUserWithTaskResponse, error)
	GetUserTaskByID(context.Context, *GetTaskByIDReq) (*GetUserWithTaskRes, error)
	DownloadUserTasks(context.Context, *ListUserWithTaskRequest) (*httpbody.HttpBody, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateUserTask(context.Context, *CreateUserTaskRequest) (*CommonResponse, error)
	ListUserGroup(context.Context, *ListUserReq) (*ListUserRes, error)
	BRICaMSgetCustomer(context.Context, *BricamsGetCustomerReq) (*BricamsGetCustomerRes, error)
	BRICamsGetAllUser(context.Context, *BricamsGetAddonsUserReq) (*BricamsGetAddonsUserRes, error)
	BRICamsGetUserByUsername(context.Context, *BricamsGetAddonsUserByUsernameReq) (*BricamsGetAddonsUserByUsernameRes, error)
	BRICaMSsvcGetUserList(context.Context, *BricamsGetAddonsUserReq) (*BRICaMSSvcUserListRes, error)
	BRICaMSsvcGetUserByUsername(context.Context, *BricamsGetAddonsUserByUsernameReq) (*BRICaMSSvcUserRes, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) GetUsers(context.Context, *CommonRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedApiServiceServer) DownloadUsers(context.Context, *CommonRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadUsers not implemented")
}
func (UnimplementedApiServiceServer) GetUserTasks(context.Context, *ListUserWithTaskRequest) (*ListUserWithTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTasks not implemented")
}
func (UnimplementedApiServiceServer) GetUserTaskByID(context.Context, *GetTaskByIDReq) (*GetUserWithTaskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTaskByID not implemented")
}
func (UnimplementedApiServiceServer) DownloadUserTasks(context.Context, *ListUserWithTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadUserTasks not implemented")
}
func (UnimplementedApiServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedApiServiceServer) CreateUserTask(context.Context, *CreateUserTaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserTask not implemented")
}
func (UnimplementedApiServiceServer) ListUserGroup(context.Context, *ListUserReq) (*ListUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserGroup not implemented")
}
func (UnimplementedApiServiceServer) BRICaMSgetCustomer(context.Context, *BricamsGetCustomerReq) (*BricamsGetCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRICaMSgetCustomer not implemented")
}
func (UnimplementedApiServiceServer) BRICamsGetAllUser(context.Context, *BricamsGetAddonsUserReq) (*BricamsGetAddonsUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRICamsGetAllUser not implemented")
}
func (UnimplementedApiServiceServer) BRICamsGetUserByUsername(context.Context, *BricamsGetAddonsUserByUsernameReq) (*BricamsGetAddonsUserByUsernameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRICamsGetUserByUsername not implemented")
}
func (UnimplementedApiServiceServer) BRICaMSsvcGetUserList(context.Context, *BricamsGetAddonsUserReq) (*BRICaMSSvcUserListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRICaMSsvcGetUserList not implemented")
}
func (UnimplementedApiServiceServer) BRICaMSsvcGetUserByUsername(context.Context, *BricamsGetAddonsUserByUsernameReq) (*BRICaMSSvcUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRICaMSsvcGetUserByUsername not implemented")
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
		FullMethod: "/user.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUsers(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/DownloadUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadUsers(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUserTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserWithTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUserTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/GetUserTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUserTasks(ctx, req.(*ListUserWithTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUserTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUserTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/GetUserTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUserTaskByID(ctx, req.(*GetTaskByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadUserTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserWithTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadUserTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/DownloadUserTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadUserTasks(ctx, req.(*ListUserWithTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateUserTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateUserTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/CreateUserTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateUserTask(ctx, req.(*CreateUserTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/ListUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListUserGroup(ctx, req.(*ListUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_BRICaMSgetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BricamsGetCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).BRICaMSgetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/BRICaMSgetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).BRICaMSgetCustomer(ctx, req.(*BricamsGetCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_BRICamsGetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BricamsGetAddonsUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).BRICamsGetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/BRICamsGetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).BRICamsGetAllUser(ctx, req.(*BricamsGetAddonsUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_BRICamsGetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BricamsGetAddonsUserByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).BRICamsGetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/BRICamsGetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).BRICamsGetUserByUsername(ctx, req.(*BricamsGetAddonsUserByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_BRICaMSsvcGetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BricamsGetAddonsUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).BRICaMSsvcGetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/BRICaMSsvcGetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).BRICaMSsvcGetUserList(ctx, req.(*BricamsGetAddonsUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_BRICaMSsvcGetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BricamsGetAddonsUserByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).BRICaMSsvcGetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.service.v1.ApiService/BRICaMSsvcGetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).BRICaMSsvcGetUserByUsername(ctx, req.(*BricamsGetAddonsUserByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _ApiService_GetUsers_Handler,
		},
		{
			MethodName: "DownloadUsers",
			Handler:    _ApiService_DownloadUsers_Handler,
		},
		{
			MethodName: "GetUserTasks",
			Handler:    _ApiService_GetUserTasks_Handler,
		},
		{
			MethodName: "GetUserTaskByID",
			Handler:    _ApiService_GetUserTaskByID_Handler,
		},
		{
			MethodName: "DownloadUserTasks",
			Handler:    _ApiService_DownloadUserTasks_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ApiService_CreateUser_Handler,
		},
		{
			MethodName: "CreateUserTask",
			Handler:    _ApiService_CreateUserTask_Handler,
		},
		{
			MethodName: "ListUserGroup",
			Handler:    _ApiService_ListUserGroup_Handler,
		},
		{
			MethodName: "BRICaMSgetCustomer",
			Handler:    _ApiService_BRICaMSgetCustomer_Handler,
		},
		{
			MethodName: "BRICamsGetAllUser",
			Handler:    _ApiService_BRICamsGetAllUser_Handler,
		},
		{
			MethodName: "BRICamsGetUserByUsername",
			Handler:    _ApiService_BRICamsGetUserByUsername_Handler,
		},
		{
			MethodName: "BRICaMSsvcGetUserList",
			Handler:    _ApiService_BRICaMSsvcGetUserList_Handler,
		},
		{
			MethodName: "BRICaMSsvcGetUserByUsername",
			Handler:    _ApiService_BRICaMSsvcGetUserByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_api.proto",
}
