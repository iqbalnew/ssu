package main

import (
	"context"
	"strings"

	manager "bitbucket.bri.co.id/scm/addons/addons-task-service/server/jwt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientInterceptor struct {
}

func NewClientInterceptor() *ClientInterceptor {
	return &ClientInterceptor{}
}

func (c *ClientInterceptor) UnaryIntercetoprSetUserData(authManager *manager.JWTManager) grpc.UnaryClientInterceptor {

	fn := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// check if authorization exist
		var accessToken string = ""
		md, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			values := md["authorization"]
			if len(values) > 0 {
				split := strings.Split(values[0], " ")
				accessToken = split[0]
				if len(split) > 1 {
					accessToken = split[1]
				}
			}

		}
		// if accesstoken exist
		if accessToken != "" {
			userMd, err := authManager.GetUserMD(ctx)
			if err != nil {
				logrus.Errorln("Failed to get user metadata")
				return err
			}

			logrus.Println("Metadata: ", md)
			logrus.Println("User Metadata: ", userMd)

			ctx = metadata.NewOutgoingContext(ctx, metadata.Join(md, userMd))
		}
		// Invoke the original method call
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}

	return fn

}

func (c *ClientInterceptor) StreamIntercetoprSetUserData(authManager *manager.JWTManager) grpc.StreamClientInterceptor {

	fn := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		// check if authorization exist
		var accessToken string = ""
		md, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			values := md["authorization"]
			if len(values) > 0 {
				split := strings.Split(values[0], " ")
				accessToken = split[0]
				if len(split) > 1 {
					accessToken = split[1]
				}
			}

		}
		// if accesstoken exist
		if accessToken != "" {
			userMd, err := authManager.GetUserMD(ctx)
			if err != nil {
				logrus.Errorln("Failed to get user metadata")
				logrus.Errorln("Error: ", err)
				return nil, err
			}
			ctx = metadata.NewOutgoingContext(ctx, metadata.Join(md, userMd))
		}
		// Invoke the original method call
		return streamer(ctx, desc, cc, method)
	}

	return fn

}
