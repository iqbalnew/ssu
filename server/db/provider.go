package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	mongoClient "bitbucket.bri.co.id/scm/addons/addons-task-service/server/mongodb"
	"gorm.io/gorm"
)

type Provider interface {
	GetStaticUser(context.Context) (*pb.UserORM, error)
	CreateUser(context.Context, *pb.UserORM) error
	GetUserByUserName(context.Context, string) (*pb.UserORM, error)

	CreateTask(context.Context, *pb.TaskORM, string) (*pb.TaskORM, error)
	UpdateTask(context.Context, *pb.TaskORM, string) (*pb.TaskORM, error)
	FindTaskById(context.Context, uint64) (*pb.TaskORM, error)
	GetListTask(context.Context) ([]*pb.TaskORM, error)
	SaveTask(context.Context, *pb.TaskORM) (*pb.TaskORM, error)
}

type GormProvider struct {
	db_main *gorm.DB
	mongo   *mongoClient.MongoDB
}

func NewProvider(db *gorm.DB, mongo *mongoClient.MongoDB) *GormProvider {
	return &GormProvider{
		db_main: db.Debug(),
		mongo:   mongo,
	}
}
