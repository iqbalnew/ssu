package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/pb"
	"gorm.io/gorm"
)

type Provider interface {
	GetStaticUser(context.Context) (*pb.UserORM, error)
	CreateUser(context.Context, *pb.UserORM) error
	GetUserByUserName(context.Context, string) (*pb.UserORM, error)
}

type GormProvider struct {
	db_main *gorm.DB
}

func NewProvider(db *gorm.DB) *GormProvider {
	return &GormProvider{db_main: db}
}
