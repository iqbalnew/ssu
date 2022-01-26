// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: task.api.proto

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	SaveTaskWithData(ctx context.Context, in *SaveTaskRequest, opts ...grpc.CallOption) (*SaveTaskResponse, error)
	SetTask(ctx context.Context, in *SetTaskRequest, opts ...grpc.CallOption) (*SetTaskResponse, error)
	GetListTask(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTaskResponse, error)
	GetTaskGraph(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTaskResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SaveTaskWithData(ctx context.Context, in *SaveTaskRequest, opts ...grpc.CallOption) (*SaveTaskResponse, error) {
	out := new(SaveTaskResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/SaveTaskWithData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SetTask(ctx context.Context, in *SetTaskRequest, opts ...grpc.CallOption) (*SetTaskResponse, error) {
	out := new(SetTaskResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/SetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetListTask(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTaskResponse, error) {
	out := new(ListTaskResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/GetListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTaskGraph(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListTaskResponse, error) {
	out := new(ListTaskResponse)
	err := c.cc.Invoke(ctx, "/task.service.v1.TaskService/GetTaskGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error)
	SaveTaskWithData(context.Context, *SaveTaskRequest) (*SaveTaskResponse, error)
	SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error)
	GetListTask(context.Context, *ListRequest) (*ListTaskResponse, error)
	GetTaskGraph(context.Context, *ListRequest) (*ListTaskResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedTaskServiceServer) SaveTaskWithData(context.Context, *SaveTaskRequest) (*SaveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTaskWithData not implemented")
}
func (UnimplementedTaskServiceServer) SetTask(context.Context, *SetTaskRequest) (*SetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTask not implemented")
}
func (UnimplementedTaskServiceServer) GetListTask(context.Context, *ListRequest) (*ListTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTaskGraph(context.Context, *ListRequest) (*ListTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskGraph not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SaveTaskWithData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SaveTaskWithData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/SaveTaskWithData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SaveTaskWithData(ctx, req.(*SaveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/SetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SetTask(ctx, req.(*SetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/GetListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetListTask(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTaskGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTaskGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.service.v1.TaskService/GetTaskGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTaskGraph(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.service.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _TaskService_HealthCheck_Handler,
		},
		{
			MethodName: "SaveTaskWithData",
			Handler:    _TaskService_SaveTaskWithData_Handler,
		},
		{
			MethodName: "SetTask",
			Handler:    _TaskService_SetTask_Handler,
		},
		{
			MethodName: "GetListTask",
			Handler:    _TaskService_GetListTask_Handler,
		},
		{
			MethodName: "GetTaskGraph",
			Handler:    _TaskService_GetTaskGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.api.proto",
}
