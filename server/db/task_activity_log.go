package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
)

type ActivityLog struct {
	TaskID  uint64
	Command string
	Data    *pb.Task
}

func (p *GormProvider) SaveLog(ctx context.Context, log *ActivityLog) error {
	save, err := p.mongo.Collection.InsertOne(ctx, log)
	if err != nil {
		return err
	}

	logrus.Println(save)

	return nil
}
