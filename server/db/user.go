package db

import (
	"context"
	"errors"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *GormProvider) GetUserByUserName(ctx context.Context, username string) (*pb.UserORM, error) {
	var user *pb.UserORM
	if err := p.db_main.Where(&pb.UserORM{Username: username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NotFound("user", username)
		}
		return nil, status.Errorf(codes.Unauthenticated, err.Error())

	}

	return user, nil
}

func (p *GormProvider) CreateUser(ctx context.Context, data *pb.UserORM) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	data.Password = string(hashedPassword)
	if data.Role == "" {
		data.Role = "User"
	}
	return p.db_main.Create(&data).Error
}

func (p *GormProvider) GetStaticUser(ctx context.Context) (*pb.UserORM, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("staticuser"), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	user := &pb.UserORM{
		Username: "staticuser",
		Password: string(hashedPassword),
		Role:     "user",
	}

	return user, nil
}
