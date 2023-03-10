package api

import (
	"context"
	"fmt"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/db"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetActivityLogs(ctx context.Context, req *pb.GetActivityLogsReq) (*pb.GetActivityLogsRes, error) {
	if req.Type == "" {
		return nil, status.Error(codes.InvalidArgument, "type is required")
	}

	if req.Limit < 0 {
		req.Limit = 1000000
	} else if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page < 1 {
		req.Page = 1
	}

	if req.DateFrom != "" && req.DateTo != "" {
		req.Limit = 0
		req.Page = 0
	}

	currentUser, _, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	filter := &db.ActivityLogFindReq{
		TaskType: req.Type,
		TaskKey:  req.Key,
		TaskID:   req.TaskID,
		Page:     int(req.Page),
		Limit:    int(req.Limit),
		Sort:     req.Sort,
		Search:   req.Search,
		DateFrom: req.DateFrom,
		DateTo:   req.DateTo,
		Filter:   req.Filter,
		GroupIDs: currentUser.GroupIDs,
	}

	find, err := s.provider.GetActivityLogs(ctx, filter)
	if err != nil {
		logrus.Errorln("Error Get Activity Logs: ", err)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	responses := &pb.GetActivityLogsRes{
		Error:   false,
		Code:    200,
		Message: fmt.Sprintf("Success Get Activity Logs %s", req.Type),
	}

	responses.Pagination = &pb.GetActivityLogsRes_ActivityLogPagination{
		Page:          int32(find.Paginate.Current),
		Limit:         int32(find.Paginate.PerPage),
		TotalRows:     int32(find.Paginate.TotalRecords),
		TotalPages:    int32(find.Paginate.TotalPages),
		RecordsOnPage: int32(find.Paginate.RecordsOnPage),
	}

	// logrus.Println("lenght: ", len(find.Logs))

	for _, v := range find.Logs {
		// logrus.Println("Print task id")
		// logrus.Println(v.TaskID)
		data := &pb.ActivityLog{
			TaskID:      v.TaskID,
			Command:     v.Command,
			Type:        v.Type,
			Key:         v.Key,
			Action:      v.Action,
			Description: v.Description,
			Username:    v.Username,
			CompanyName: v.CompanyName,
			CreatedAt:   timestamppb.New(v.Created),
		}
		if req.TaskID > 0 {
			task, err := v.Data.ToPB(ctx)
			if err != nil {
				logrus.Errorln("Error Get Activity Logs: ", err)
				return nil, status.Error(codes.Internal, "Server Error")
			}
			data.Task = &task
		}
		responses.Data = append(responses.Data, data)
	}

	return responses, nil
}
