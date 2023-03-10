// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: product_api.proto

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
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	ProductByID(ctx context.Context, in *ProductByIDRequest, opts ...grpc.CallOption) (*ProductByIDResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/product.service.v1.ApiService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, "/product.service.v1.ApiService/ListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/product.service.v1.ApiService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ProductByID(ctx context.Context, in *ProductByIDRequest, opts ...grpc.CallOption) (*ProductByIDResponse, error) {
	out := new(ProductByIDResponse)
	err := c.cc.Invoke(ctx, "/product.service.v1.ApiService/ProductByID", in, out, opts...)
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
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	ProductByID(context.Context, *ProductByIDRequest) (*ProductByIDResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) HealthCheck(context.Context, *Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedApiServiceServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedApiServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedApiServiceServer) ProductByID(context.Context, *ProductByIDRequest) (*ProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductByID not implemented")
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
		FullMethod: "/product.service.v1.ApiService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.service.v1.ApiService/ListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.service.v1.ApiService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.service.v1.ApiService/ProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ProductByID(ctx, req.(*ProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.service.v1.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ApiService_HealthCheck_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _ApiService_ListProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ApiService_CreateProduct_Handler,
		},
		{
			MethodName: "ProductByID",
			Handler:    _ApiService_ProductByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_api.proto",
}
