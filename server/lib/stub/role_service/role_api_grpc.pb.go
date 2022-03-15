// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.16.0
// source: role_api.proto

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
	ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleResponse, error)
	ListAuthorityLevel(ctx context.Context, in *ListAuthorityLevelRequest, opts ...grpc.CallOption) (*ListAuthorityLevelResponse, error)
	ListRoleAuthority(ctx context.Context, in *ListRoleAuthorityRequest, opts ...grpc.CallOption) (*ListRoleAuthorityResponse, error)
	ListUserType(ctx context.Context, in *ListUserTypeRequest, opts ...grpc.CallOption) (*ListUserTypeResponse, error)
	ListUserRole(ctx context.Context, in *ListUserRoleRequest, opts ...grpc.CallOption) (*ListUserRoleResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	CreateRoleTask(ctx context.Context, in *CreateRoleTaskRequest, opts ...grpc.CallOption) (*CreateRoleTaskResponse, error)
	ListRoleTask(ctx context.Context, in *ListRoleTaskRequest, opts ...grpc.CallOption) (*ListRoleTaskResponse, error)
	GetRoleUserByUserID(ctx context.Context, in *GetRoleUserByUserIDReq, opts ...grpc.CallOption) (*GetRoleUserByUserIDRes, error)
	AssignUserRoles(ctx context.Context, in *AssignUserRolesRequest, opts ...grpc.CallOption) (*ErrorBodyResponse, error)
	RoleTaskDetail(ctx context.Context, in *RoleTaskDetailRequest, opts ...grpc.CallOption) (*RoleTaskDetailResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleResponse, error) {
	out := new(ListRoleResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAuthorityLevel(ctx context.Context, in *ListAuthorityLevelRequest, opts ...grpc.CallOption) (*ListAuthorityLevelResponse, error) {
	out := new(ListAuthorityLevelResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListAuthorityLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListRoleAuthority(ctx context.Context, in *ListRoleAuthorityRequest, opts ...grpc.CallOption) (*ListRoleAuthorityResponse, error) {
	out := new(ListRoleAuthorityResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListRoleAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListUserType(ctx context.Context, in *ListUserTypeRequest, opts ...grpc.CallOption) (*ListUserTypeResponse, error) {
	out := new(ListUserTypeResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListUserType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListUserRole(ctx context.Context, in *ListUserRoleRequest, opts ...grpc.CallOption) (*ListUserRoleResponse, error) {
	out := new(ListUserRoleResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateRoleTask(ctx context.Context, in *CreateRoleTaskRequest, opts ...grpc.CallOption) (*CreateRoleTaskResponse, error) {
	out := new(CreateRoleTaskResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/CreateRoleTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListRoleTask(ctx context.Context, in *ListRoleTaskRequest, opts ...grpc.CallOption) (*ListRoleTaskResponse, error) {
	out := new(ListRoleTaskResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/ListRoleTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetRoleUserByUserID(ctx context.Context, in *GetRoleUserByUserIDReq, opts ...grpc.CallOption) (*GetRoleUserByUserIDRes, error) {
	out := new(GetRoleUserByUserIDRes)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/GetRoleUserByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AssignUserRoles(ctx context.Context, in *AssignUserRolesRequest, opts ...grpc.CallOption) (*ErrorBodyResponse, error) {
	out := new(ErrorBodyResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/AssignUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RoleTaskDetail(ctx context.Context, in *RoleTaskDetailRequest, opts ...grpc.CallOption) (*RoleTaskDetailResponse, error) {
	out := new(RoleTaskDetailResponse)
	err := c.cc.Invoke(ctx, "/role.service.v1.ApiService/RoleTaskDetail", in, out, opts...)
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
	ListRole(context.Context, *ListRoleRequest) (*ListRoleResponse, error)
	ListAuthorityLevel(context.Context, *ListAuthorityLevelRequest) (*ListAuthorityLevelResponse, error)
	ListRoleAuthority(context.Context, *ListRoleAuthorityRequest) (*ListRoleAuthorityResponse, error)
	ListUserType(context.Context, *ListUserTypeRequest) (*ListUserTypeResponse, error)
	ListUserRole(context.Context, *ListUserRoleRequest) (*ListUserRoleResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	CreateRoleTask(context.Context, *CreateRoleTaskRequest) (*CreateRoleTaskResponse, error)
	ListRoleTask(context.Context, *ListRoleTaskRequest) (*ListRoleTaskResponse, error)
	GetRoleUserByUserID(context.Context, *GetRoleUserByUserIDReq) (*GetRoleUserByUserIDRes, error)
	AssignUserRoles(context.Context, *AssignUserRolesRequest) (*ErrorBodyResponse, error)
	RoleTaskDetail(context.Context, *RoleTaskDetailRequest) (*RoleTaskDetailResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListRole(context.Context, *ListRoleRequest) (*ListRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedApiServiceServer) ListAuthorityLevel(context.Context, *ListAuthorityLevelRequest) (*ListAuthorityLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthorityLevel not implemented")
}
func (UnimplementedApiServiceServer) ListRoleAuthority(context.Context, *ListRoleAuthorityRequest) (*ListRoleAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoleAuthority not implemented")
}
func (UnimplementedApiServiceServer) ListUserType(context.Context, *ListUserTypeRequest) (*ListUserTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserType not implemented")
}
func (UnimplementedApiServiceServer) ListUserRole(context.Context, *ListUserRoleRequest) (*ListUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRole not implemented")
}
func (UnimplementedApiServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedApiServiceServer) CreateRoleTask(context.Context, *CreateRoleTaskRequest) (*CreateRoleTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleTask not implemented")
}
func (UnimplementedApiServiceServer) ListRoleTask(context.Context, *ListRoleTaskRequest) (*ListRoleTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoleTask not implemented")
}
func (UnimplementedApiServiceServer) GetRoleUserByUserID(context.Context, *GetRoleUserByUserIDReq) (*GetRoleUserByUserIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleUserByUserID not implemented")
}
func (UnimplementedApiServiceServer) AssignUserRoles(context.Context, *AssignUserRolesRequest) (*ErrorBodyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignUserRoles not implemented")
}
func (UnimplementedApiServiceServer) RoleTaskDetail(context.Context, *RoleTaskDetailRequest) (*RoleTaskDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleTaskDetail not implemented")
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
		FullMethod: "/role.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListRole(ctx, req.(*ListRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAuthorityLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorityLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAuthorityLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListAuthorityLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAuthorityLevel(ctx, req.(*ListAuthorityLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListRoleAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListRoleAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListRoleAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListRoleAuthority(ctx, req.(*ListRoleAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListUserType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListUserType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListUserType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListUserType(ctx, req.(*ListUserTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListUserRole(ctx, req.(*ListUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateRoleTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateRoleTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/CreateRoleTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateRoleTask(ctx, req.(*CreateRoleTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListRoleTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListRoleTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/ListRoleTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListRoleTask(ctx, req.(*ListRoleTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetRoleUserByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleUserByUserIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetRoleUserByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/GetRoleUserByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetRoleUserByUserID(ctx, req.(*GetRoleUserByUserIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AssignUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AssignUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/AssignUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AssignUserRoles(ctx, req.(*AssignUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RoleTaskDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleTaskDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RoleTaskDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.service.v1.ApiService/RoleTaskDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RoleTaskDetail(ctx, req.(*RoleTaskDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _ApiService_ListRole_Handler,
		},
		{
			MethodName: "ListAuthorityLevel",
			Handler:    _ApiService_ListAuthorityLevel_Handler,
		},
		{
			MethodName: "ListRoleAuthority",
			Handler:    _ApiService_ListRoleAuthority_Handler,
		},
		{
			MethodName: "ListUserType",
			Handler:    _ApiService_ListUserType_Handler,
		},
		{
			MethodName: "ListUserRole",
			Handler:    _ApiService_ListUserRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _ApiService_CreateRole_Handler,
		},
		{
			MethodName: "CreateRoleTask",
			Handler:    _ApiService_CreateRoleTask_Handler,
		},
		{
			MethodName: "ListRoleTask",
			Handler:    _ApiService_ListRoleTask_Handler,
		},
		{
			MethodName: "GetRoleUserByUserID",
			Handler:    _ApiService_GetRoleUserByUserID_Handler,
		},
		{
			MethodName: "AssignUserRoles",
			Handler:    _ApiService_AssignUserRoles_Handler,
		},
		{
			MethodName: "RoleTaskDetail",
			Handler:    _ApiService_RoleTaskDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_api.proto",
}
