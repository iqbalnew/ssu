package db

import (
	"context"

	pb "github.com/ordentco/go-base/server/pb"
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
