package api

import (
	"context"
	"strings"

	manager "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/jwt"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func IsCorrectPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// user, err := s.provider.GetStaticUser(ctx)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	// }

	// if user == nil || !IsCorrectPassword(user.Password, req.Password) {
	// 	return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	// }

	// token, err := s.manager.Generate(pb.User{Username: user.Username, Role: user.Role})
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "cannot generate access token")
	// }

	// res := &pb.LoginResponse{
	// 	AccessToken: token,
	// }

	return nil, status.Errorf(codes.Unimplemented, "Not Implemented")
}

func (s *Server) getCurrentUser(ctx context.Context) (*manager.VerifyTokenRes, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	split := strings.Split(values[0], " ")
	accessToken := split[0]
	if len(split) > 1 {
		accessToken = split[1]
	}

	getUser, err := s.manager.GetMeFromAuthService(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	if getUser.IsExpired && !getUser.IsValid {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return getUser, nil
}
