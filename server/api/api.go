package api

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	manager "github.com/ordentco/go-base/server/api/jwt"
	pb "github.com/ordentco/go-base/server/pb"
	"github.com/sirupsen/logrus"
)

// Server represents the server implementation of the SW API.
type Server struct {
	// provider db.Provider
	manager manager.JWTManager
	pb.GoBaseServiceServer
}

func New(
// db01 *gorm.DB,
// db02 *gorm.DB,
// db03 *gorm.DB,
) *Server {
	secret := os.Getenv("JWT_SECRET")
	tokenDuration, err := time.ParseDuration(os.Getenv("JWT_DURATION"))
	if err != nil {
		logrus.Panic(err)
	}

	return &Server{
		// provider:           db.NewProvider(db01),
		manager:             *manager.NewJWTManager(secret, tokenDuration),
		GoBaseServiceServer: nil,
	}
}

func (s *Server) notImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}

func (s *Server) HealthCheck(ctx context.Context, _ *pb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Message: "API Running !"}, nil
}
