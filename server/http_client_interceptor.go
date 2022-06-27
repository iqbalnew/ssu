package main

import (
	"context"
	"log"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientInterceptor struct {
}

func NewClientInterceptor() *ClientInterceptor {
	return &ClientInterceptor{}
}

func (c *ClientInterceptor) SetUserData() grpc.UnaryClientInterceptor {

	fn := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Get the operating system the client is running on
		cos := runtime.GOOS
		// Append the OS info to the outgoing request
		ctx = metadata.AppendToOutgoingContext(ctx, "client-os", cos)
		// Invoke the original method call
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("client interceptor hit: appending OS: '%v' to metadata", cos)
		return err
	}

	return fn

}
