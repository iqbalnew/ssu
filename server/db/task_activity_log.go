package db

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type ActivityLog struct {
	bongo.DocumentBase `bson:",inline"`
	TaskID             uint64   `json:"task_id"`
	Command            string   `json:"command"`
	Type               string   `json:"type"`
	Data               *pb.Task `json:"data"`
}

func (p *GormProvider) SaveLog(ctx context.Context, log *ActivityLog) error {
	if log.Type == "" && log.Data != nil {
		log.Type = log.Data.Type
	}
	err := p.mongo.Collection.Save(log)
	if err != nil {
		return err
	}

	return nil
}

type ActivityLogFind struct {
	logs     []*ActivityLog
	paginate *bongo.PaginationInfo
}

func (p *GormProvider) GetLogByTaskID(ctx context.Context, taskID uint64, perPage int, page int) (*ActivityLogFind, error) {
	var logs []*ActivityLog
	results := p.mongo.Collection.Find(bson.M{"task_id": taskID})

	log := &ActivityLog{}

	pagination, err := results.Paginate(perPage, page)
	if err != nil {
		logrus.Errorf("Failed to paginate data log: %v", err)
		return nil, err
	}

	for results.Next(log) {
		logs = append(logs, log)
	}

	return &ActivityLogFind{
		logs:     logs,
		paginate: pagination,
	}, nil
}

func (p *GormProvider) GetLogByTaskType(ctx context.Context, taskType string, perPage int, page int) (*ActivityLogFind, error) {
	var logs []*ActivityLog
	results := p.mongo.Collection.Find(bson.M{"type": taskType})

	log := &ActivityLog{}

	pagination, err := results.Paginate(perPage, page)
	if err != nil {
		logrus.Errorf("Failed to paginate data log: %v", err)
		return nil, err
	}

	for results.Next(log) {
		logs = append(logs, log)
	}

	return &ActivityLogFind{
		logs:     logs,
		paginate: pagination,
	}, nil
}
