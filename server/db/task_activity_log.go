package db

import (
	"context"
	"fmt"
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

	switch log.Action {
	case "save":
		log.Action = "save & send for approval"

	case "draft":
		log.Action = "save as draft"
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

	req.TaskType = strings.Replace(req.TaskType, ":", "_", -1)
	req.TaskType = strings.ToLower(req.TaskType)

	query := bson.D{
		{
			Name:  "type",
			Value: req.TaskType,
		},
	}

	if req.TaskID > 0 {
		query = append(query, bson.DocElem{
			Name:  "taskid",
			Value: req.TaskID,
		})
	}

	if len(req.GroupIDs) > 0 {
		query = append(query, bson.DocElem{
			Name:  "companyid",
			Value: bson.DocElem{Name: "$in", Value: req.GroupIDs},
		})
	}

	if req.Filter != nil {
		if req.Filter.Command != "" {
			query = append(query, bson.DocElem{
				Name:  "command",
				Value: req.Filter.Command,
			})
		}

		if req.Filter.Action != "" {
			query = append(query, bson.DocElem{
				Name:  "action",
				Value: req.Filter.Action,
			})
		}

		if req.Filter.Description != "" {
			query = append(query, bson.DocElem{
				Name:  "description",
				Value: req.Filter.Description,
			})
		}

		if req.Filter.Username != "" {
			query = append(query, bson.DocElem{
				Name:  "username",
				Value: req.Filter.Username,
			})
		}

		if req.Filter.CompanyName != "" {
			query = append(query, bson.DocElem{
				Name:  "companyname",
				Value: req.Filter.CompanyName,
			})
		}
	}

	if req.Search != "" {

		query = append(query, bson.DocElem{
			Name: "$or",
			Value: []interface{}{
				bson.DocElem{Name: "action", Value: fmt.Sprintf("/%s/i", req.Search)},
				bson.DocElem{Name: "description", Value: fmt.Sprintf("/%s/i", req.Search)},
				bson.DocElem{Name: "username", Value: fmt.Sprintf("/%s/i", req.Search)},
				bson.DocElem{Name: "companyname", Value: fmt.Sprintf("/%s/i", req.Search)},
			},
		})
	}

	logrus.Println("===<Mongo Find Filter>===")
	logrus.Println(query)

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

	result := &ActivityLogFindRes{
		Logs:     logs,
		Paginate: pagination,
	}

	log := &ActivityLog{}

	for results.Next(log) {
		data := &ActivityLog{
			DocumentBase: log.DocumentBase,
			TaskID:       log.TaskID,
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
	}

	return result, nil
}
