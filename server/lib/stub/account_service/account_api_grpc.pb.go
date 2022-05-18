// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: account_api.proto

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
	ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateAccountTask(ctx context.Context, in *CreateAccountTaskRequest, opts ...grpc.CallOption) (*CreateAccountTaskResponse, error)
	CreateAccountTaskMultiple(ctx context.Context, in *CreateAccountTaskBulkRequest, opts ...grpc.CallOption) (*CreateAccountTaskResponse, error)
	CreateAccountTaskMultipleNoParrent(ctx context.Context, in *CreateAccountTaskBulkRequest, opts ...grpc.CallOption) (*CreateAccountTaskMultipleResponse, error)
	ListAccountTask(ctx context.Context, in *ListAccountTaskRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error)
	GetAccountTaskByID(ctx context.Context, in *GetAccountTaskByIDRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error)
	RequestDeleteAccountTask(ctx context.Context, in *GetAccountTaskByIDRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error)
	ValidateAccount(ctx context.Context, in *ValidateAccountRequest, opts ...grpc.CallOption) (*ValidateAccountResponse, error)
	AccountDetail(ctx context.Context, in *AccountDetailRequest, opts ...grpc.CallOption) (*AccountDetailResponse, error)
	UpdateAccountRole(ctx context.Context, in *UpdateAccountRoleRequest, opts ...grpc.CallOption) (*UpdateAccountRoleResponse, error)
	DownloadListAccountTasks(ctx context.Context, in *FileListAccountTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	DownloadTemplate(ctx context.Context, in *FileListTemplateRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	CekAccountAvaibility(ctx context.Context, in *CekAccountAvaibilityReq, opts ...grpc.CallOption) (*CekAccountAvaibilityRes, error)
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

func (c *apiServiceClient) CreateAccountTaskMultiple(ctx context.Context, in *CreateAccountTaskBulkRequest, opts ...grpc.CallOption) (*CreateAccountTaskResponse, error) {
	out := new(CreateAccountTaskResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/CreateAccountTaskMultiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateAccountTaskMultipleNoParrent(ctx context.Context, in *CreateAccountTaskBulkRequest, opts ...grpc.CallOption) (*CreateAccountTaskMultipleResponse, error) {
	out := new(CreateAccountTaskMultipleResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/CreateAccountTaskMultipleNoParrent", in, out, opts...)
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

func (c *apiServiceClient) GetAccountTaskByID(ctx context.Context, in *GetAccountTaskByIDRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error) {
	out := new(ListAccountTaskResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/GetAccountTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RequestDeleteAccountTask(ctx context.Context, in *GetAccountTaskByIDRequest, opts ...grpc.CallOption) (*ListAccountTaskResponse, error) {
	out := new(ListAccountTaskResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/RequestDeleteAccountTask", in, out, opts...)
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

func (c *apiServiceClient) AccountDetail(ctx context.Context, in *AccountDetailRequest, opts ...grpc.CallOption) (*AccountDetailResponse, error) {
	out := new(AccountDetailResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/AccountDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateAccountRole(ctx context.Context, in *UpdateAccountRoleRequest, opts ...grpc.CallOption) (*UpdateAccountRoleResponse, error) {
	out := new(UpdateAccountRoleResponse)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/UpdateAccountRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadListAccountTasks(ctx context.Context, in *FileListAccountTaskRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/DownloadListAccountTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DownloadTemplate(ctx context.Context, in *FileListTemplateRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/DownloadTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CekAccountAvaibility(ctx context.Context, in *CekAccountAvaibilityReq, opts ...grpc.CallOption) (*CekAccountAvaibilityRes, error) {
	out := new(CekAccountAvaibilityRes)
	err := c.cc.Invoke(ctx, "/account.service.v1.ApiService/CekAccountAvaibility", in, out, opts...)
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
	CreateAccountTaskMultiple(context.Context, *CreateAccountTaskBulkRequest) (*CreateAccountTaskResponse, error)
	CreateAccountTaskMultipleNoParrent(context.Context, *CreateAccountTaskBulkRequest) (*CreateAccountTaskMultipleResponse, error)
	ListAccountTask(context.Context, *ListAccountTaskRequest) (*ListAccountTaskResponse, error)
	GetAccountTaskByID(context.Context, *GetAccountTaskByIDRequest) (*ListAccountTaskResponse, error)
	RequestDeleteAccountTask(context.Context, *GetAccountTaskByIDRequest) (*ListAccountTaskResponse, error)
	ValidateAccount(context.Context, *ValidateAccountRequest) (*ValidateAccountResponse, error)
	AccountDetail(context.Context, *AccountDetailRequest) (*AccountDetailResponse, error)
	UpdateAccountRole(context.Context, *UpdateAccountRoleRequest) (*UpdateAccountRoleResponse, error)
	DownloadListAccountTasks(context.Context, *FileListAccountTaskRequest) (*httpbody.HttpBody, error)
	DownloadTemplate(context.Context, *FileListTemplateRequest) (*httpbody.HttpBody, error)
	CekAccountAvaibility(context.Context, *CekAccountAvaibilityReq) (*CekAccountAvaibilityRes, error)
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
func (UnimplementedApiServiceServer) CreateAccountTaskMultiple(context.Context, *CreateAccountTaskBulkRequest) (*CreateAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountTaskMultiple not implemented")
}
func (UnimplementedApiServiceServer) CreateAccountTaskMultipleNoParrent(context.Context, *CreateAccountTaskBulkRequest) (*CreateAccountTaskMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountTaskMultipleNoParrent not implemented")
}
func (UnimplementedApiServiceServer) ListAccountTask(context.Context, *ListAccountTaskRequest) (*ListAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountTask not implemented")
}
func (UnimplementedApiServiceServer) GetAccountTaskByID(context.Context, *GetAccountTaskByIDRequest) (*ListAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountTaskByID not implemented")
}
func (UnimplementedApiServiceServer) RequestDeleteAccountTask(context.Context, *GetAccountTaskByIDRequest) (*ListAccountTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDeleteAccountTask not implemented")
}
func (UnimplementedApiServiceServer) ValidateAccount(context.Context, *ValidateAccountRequest) (*ValidateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAccount not implemented")
}
func (UnimplementedApiServiceServer) AccountDetail(context.Context, *AccountDetailRequest) (*AccountDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountDetail not implemented")
}
func (UnimplementedApiServiceServer) UpdateAccountRole(context.Context, *UpdateAccountRoleRequest) (*UpdateAccountRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountRole not implemented")
}
func (UnimplementedApiServiceServer) DownloadListAccountTasks(context.Context, *FileListAccountTaskRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadListAccountTasks not implemented")
}
func (UnimplementedApiServiceServer) DownloadTemplate(context.Context, *FileListTemplateRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadTemplate not implemented")
}
func (UnimplementedApiServiceServer) CekAccountAvaibility(context.Context, *CekAccountAvaibilityReq) (*CekAccountAvaibilityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CekAccountAvaibility not implemented")
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

func _ApiService_CreateAccountTaskMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountTaskBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAccountTaskMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/CreateAccountTaskMultiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAccountTaskMultiple(ctx, req.(*CreateAccountTaskBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateAccountTaskMultipleNoParrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountTaskBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateAccountTaskMultipleNoParrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/CreateAccountTaskMultipleNoParrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateAccountTaskMultipleNoParrent(ctx, req.(*CreateAccountTaskBulkRequest))
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

func _ApiService_GetAccountTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountTaskByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAccountTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/GetAccountTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAccountTaskByID(ctx, req.(*GetAccountTaskByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RequestDeleteAccountTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountTaskByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RequestDeleteAccountTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/RequestDeleteAccountTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RequestDeleteAccountTask(ctx, req.(*GetAccountTaskByIDRequest))
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

func _ApiService_AccountDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AccountDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/AccountDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AccountDetail(ctx, req.(*AccountDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_UpdateAccountRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).UpdateAccountRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/UpdateAccountRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).UpdateAccountRole(ctx, req.(*UpdateAccountRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadListAccountTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListAccountTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadListAccountTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/DownloadListAccountTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadListAccountTasks(ctx, req.(*FileListAccountTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DownloadTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DownloadTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/DownloadTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DownloadTemplate(ctx, req.(*FileListTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CekAccountAvaibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CekAccountAvaibilityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CekAccountAvaibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.service.v1.ApiService/CekAccountAvaibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CekAccountAvaibility(ctx, req.(*CekAccountAvaibilityReq))
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
			MethodName: "CreateAccountTaskMultiple",
			Handler:    _ApiService_CreateAccountTaskMultiple_Handler,
		},
		{
			MethodName: "CreateAccountTaskMultipleNoParrent",
			Handler:    _ApiService_CreateAccountTaskMultipleNoParrent_Handler,
		},
		{
			MethodName: "ListAccountTask",
			Handler:    _ApiService_ListAccountTask_Handler,
		},
		{
			MethodName: "GetAccountTaskByID",
			Handler:    _ApiService_GetAccountTaskByID_Handler,
		},
		{
			MethodName: "RequestDeleteAccountTask",
			Handler:    _ApiService_RequestDeleteAccountTask_Handler,
		},
		{
			MethodName: "ValidateAccount",
			Handler:    _ApiService_ValidateAccount_Handler,
		},
		{
			MethodName: "AccountDetail",
			Handler:    _ApiService_AccountDetail_Handler,
		},
		{
			MethodName: "UpdateAccountRole",
			Handler:    _ApiService_UpdateAccountRole_Handler,
		},
		{
			MethodName: "DownloadListAccountTasks",
			Handler:    _ApiService_DownloadListAccountTasks_Handler,
		},
		{
			MethodName: "DownloadTemplate",
			Handler:    _ApiService_DownloadTemplate_Handler,
		},
		{
			MethodName: "CekAccountAvaibility",
			Handler:    _ApiService_CekAccountAvaibility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_api.proto",
}
