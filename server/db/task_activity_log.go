package db

import (
	"context"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type ActivityLog struct {
	bongo.DocumentBase `bson:",inline"`
	TaskID             uint64      `bson:"taskid" json:"taskID"`
	Command            string      `bson:"command" json:"command"`
	Type               string      `bson:"type" json:"type"`
	Action             string      `bson:"action" json:"action"`
	Description        string      `bson:"description" json:"description"`
	UserID             uint64      `bson:"userid" json:"userID"`
	Username           string      `bson:"username" json:"username"`
	CompanyID          uint64      `bson:"companyid" json:"companyID"`
	CompanyName        string      `bson:"companyname" json:"companyName"`
	RoleIDs            []uint64    `bson:"roleids" json:"roleIDs"`
	Data               *pb.TaskORM `bson:"data" json:"data"`
}

func (p *GormProvider) SaveLog(ctx context.Context, log *ActivityLog) error {
	if log.Type == "" && log.Data != nil {
		log.Type = log.Data.Type
	}

	log.Type = strings.Replace(log.Type, ":", "_", -1)
	log.Type = strings.ToLower(log.Type)

	err := p.mongo.Collection.Save(log)
	if err != nil {
		return err
	}

	return nil
}

type ActivityLogFindRes struct {
	Logs     []*ActivityLog
	Paginate *bongo.PaginationInfo
}

type ActivityLogFindReq struct {
	TaskID   uint64   `json:"taskID"`
	TaskType string   `json:"taskType"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	Sort     string   `json:"sort"`
	GroupIDs []uint64 `json:"groupIDs"`
}

func (p *GormProvider) GetActivityLogs(ctx context.Context, req *ActivityLogFindReq) (*ActivityLogFindRes, error) {

	var logs []*ActivityLog
	query := bson.M{
		"type": req.TaskType,
	}
	if req.TaskID > 0 {
		query["taskid"] = req.TaskID
	}
	if len(req.GroupIDs) > 0 {
		query["companyid"] = bson.M{"$in": req.GroupIDs}
	}

	results := p.mongo.Collection.Find(query)
	if req.Sort == "" {
		results.Query = results.Query.Sort("-_created")
	} else {
		results.Query = results.Query.Sort(req.Sort)
	}

	pagination, err := results.Paginate(req.Limit, req.Page)
	if err != nil {
		logrus.Errorf("Failed to paginate data log: %v", err)
		return nil, err
	}

	log := &ActivityLog{}

	for results.Next(log) {
		logs = append(logs, log)
	}

	return &ActivityLogFindRes{
		Logs:     logs,
		Paginate: pagination,
	}, nil
}
