package api

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/db"
	manager "bitbucket.bri.co.id/scm/addons/addons-task-service/server/jwt"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	addonsLogger "bitbucket.bri.co.id/scm/addons/addons-task-service/server/logger"
	mongoClient "bitbucket.bri.co.id/scm/addons/addons-task-service/server/mongodb"
	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/redis"

	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

const TaskServicePath string = "/task.service.v1.TaskService/"

type dataPublish struct {
	DataType string
	Data     string
}

// Server represents the server implementation of the SW API.
type Server struct {
	provider         *db.GormProvider
	manager          *manager.JWTManager
	announcementConn *grpc.ClientConn
	logger           *addonsLogger.Logger

	sid *shortid.Shortid
	rdb *redis.Redis

	pb.TaskServiceServer
}

func New(
	db01 *gorm.DB,
	authManager *manager.JWTManager,
	sid *shortid.Shortid,
	conn01 *grpc.ClientConn,
	mongo01 *mongoClient.MongoDB,
	logger *addonsLogger.Logger,
	rdb *redis.Redis,
) *Server {
	server := &Server{
		provider:         db.NewProvider(db01, mongo01, rdb),
		announcementConn: conn01,
		manager:          authManager,
		logger:           logger,

		sid:               sid,
		rdb:               rdb,
		TaskServiceServer: nil,
	}

	return server
}

func (s *Server) TestLogger(ctx context.Context, req *pb.LoggerTestReq) (*pb.LoggerTestRes, error) {
	meMD, err := s.manager.GetUserMD(ctx)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	logrus.Println(meMD["user-username"][0])

	s.logger.Info(req.Message)
	return &pb.LoggerTestRes{
		Message: req.Message,
	}, nil
}

func (s *Server) GetManager() *manager.JWTManager {
	return s.manager
}

func (s *Server) notImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}

func (s *Server) HealthCheck(ctx context.Context, _ *pb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Message: "API Running !"}, nil
}
