package db

import (
	"context"
	"os"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	mongoClient "bitbucket.bri.co.id/scm/addons/addons-task-service/server/mongodb"
	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/redis"
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
	rdb     *redis.Redis
}

func NewProvider(db *gorm.DB, mongo *mongoClient.MongoDB, rdb *redis.Redis) *GormProvider {
	db_main := db
	if getEnv("DEBUG", "false") == "true" {
		db_main = db_main.Debug()
	}
	return &GormProvider{
		db_main: db_main,
		mongo:   mongo,
		rdb:     rdb,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
