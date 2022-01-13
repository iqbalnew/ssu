package api

import (
	"context"

	pb "github.com/ordentco/go-base/server/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsCorrectPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.provider.GetStaticUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !IsCorrectPassword(user.Password, req.Password) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := s.manager.Generate(pb.User{Username: user.Username, Role: user.Role})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}

	return res, nil
}
