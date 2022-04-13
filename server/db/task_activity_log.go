package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
)

type ActivityLog struct {
	TaskID  uint64
	Command string
	Type    string
	Data    *pb.Task
}

func (p *GormProvider) SaveLog(ctx context.Context, log *ActivityLog) error {
	if log.Type == "" && log.Data != nil {
		log.Type = log.Data.Type
	}
	save, err := p.mongo.Collection.InsertOne(ctx, log)
	if err != nil {
		return err
	}

	logrus.Println(save)

	return nil
}
