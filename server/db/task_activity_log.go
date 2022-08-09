package db

import (
	"context"
	"strings"
	"time"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

type ActivityLog struct {
	bongo.DocumentBase `bson:",inline"`
	TaskID             uint64      `bson:"taskid" json:"taskID"`
	Command            string      `bson:"command" json:"command"`
	Type               string      `bson:"type" json:"type"`
	Key                string      `bson:"key" json:"key"`
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

	switch log.Action {
	case "save":
		log.Action = "save & send for approval"

	case "draft":
		log.Action = "save as draft"

	case "delete":
		log.Action = "request for deleted"
	}

	if log.Data.Status == 7 {
		log.Action = "deleted"
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
	TaskType string                `json:"taskType"`
	TaskKey  string                `json:"key"`
	TaskID   uint64                `json:"taskID"`
	Page     int                   `json:"page"`
	Limit    int                   `json:"limit"`
	Sort     string                `json:"sort"`
	Search   string                `json:"search"`
	DateFrom string                `json:"dateFrom"`
	DateTo   string                `json:"dateTo"`
	Filter   *pb.ActivityLogFilter `json:"filter"`
	GroupIDs []uint64              `json:"groupIDs"`
}

func (p *GormProvider) GetActivityLogs(ctx context.Context, req *ActivityLogFindReq) (*ActivityLogFindRes, error) {

	var logs []*ActivityLog

	if p.mongo == nil {
		logrus.Println("Mongo Connection is nil")
		return &ActivityLogFindRes{
			Logs: logs,
			Paginate: &bongo.PaginationInfo{
				Current:       0,
				TotalPages:    0,
				PerPage:       0,
				TotalRecords:  0,
				RecordsOnPage: 0,
			},
		}, nil
	}

	req.TaskType = strings.Replace(req.TaskType, ":", "_", -1)
	req.TaskType = strings.ToLower(req.TaskType)

	query := bson.M{
		"type": req.TaskType,
	}

	if req.DateFrom != "" && req.DateTo != "" {
		dateLayout := "2006-01-02"
		dateFrom, err := time.Parse(dateLayout, req.DateFrom)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid dateFrom")
		}
		dateTo, err := time.Parse(dateLayout, req.DateTo)
		dateTo = dateTo.AddDate(0, 0, 1)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid dateTo")
		}
		query["_created"] = bson.M{"$gte": dateFrom, "$lte": dateTo}
	}

	if req.TaskKey != "" {
		query["key"] = req.TaskKey
	}

	if req.TaskID > 0 {
		query["taskid"] = req.TaskID
	}

	if len(req.GroupIDs) > 0 {
		query["companyid"] = bson.M{"$in": req.GroupIDs}
	}

	if req.Filter != nil {
		if req.Filter.Command != "" {
			query["command"] = req.Filter.Command
		}

		if req.Filter.Action != "" {
			query["action"] = req.Filter.Action
		}

		if req.Filter.Description != "" {
			query["description"] = req.Filter.Description
		}

		if req.Filter.Username != "" {
			query["username"] = req.Filter.Username
		}

		if req.Filter.CompanyName != "" {
			query["companyname"] = req.Filter.CompanyName
		}
	}

	if req.Search != "" {
		query["$or"] = []interface{}{
			bson.M{"command": bson.M{"$regex": req.Search, "$options": "i"}},
			bson.M{"action": bson.M{"$regex": req.Search, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": req.Search, "$options": "i"}},
			bson.M{"username": bson.M{"$regex": req.Search, "$options": "i"}},
			bson.M{"companyname": bson.M{"$regex": req.Search, "$options": "i"}},
		}
	}

	logrus.Println("===<Mongo Find Filter>===")
	queryS, _ := bson.Marshal(query)
	logrus.Println(string(queryS))

	results := p.mongo.Collection.Find(query)
	if req.Sort == "" {
		results.Query = results.Query.Sort("-_created")
	} else {
		results.Query = results.Query.Sort(req.Sort)
	}

	// pagination := &bongo.PaginationInfo{
	// 	Current:       -1,
	// 	TotalPages:    -1,
	// 	PerPage:       -1,
	// 	TotalRecords:  -1,
	// 	RecordsOnPage: -1,
	// }
	pagination, err := results.Paginate(req.Limit, req.Page)
	if err != nil {
		logrus.Errorf("Failed to paginate data log: %v", err)
		return nil, err
	}

	result := &ActivityLogFindRes{
		Logs:     logs,
		Paginate: pagination,
	}

	log := &ActivityLog{}

	for results.Next(log) {

		data := &ActivityLog{
			DocumentBase: log.DocumentBase,
			TaskID:       log.TaskID,
			Key:          log.Key,
			Command:      log.Command,
			Type:         log.Type,
			Action:       log.Action,
			Description:  log.Description,
			UserID:       log.UserID,
			Username:     log.Username,
			CompanyID:    log.CompanyID,
			CompanyName:  log.CompanyName,
			RoleIDs:      log.RoleIDs,
			Data:         log.Data,
		}
		result.Logs = append(result.Logs, data)
		logrus.Println(data)
	}

	return result, nil
}
