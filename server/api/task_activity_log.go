package api

import (
	"context"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/db"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
)

func (s *Server) TestActivityLog(ctx context.Context, req *pb.ActivityLogTestReq) (*pb.ActivityLogTestRes, error) {
	
	task, err := s.GetTaskByID(ctx, &pb.GetTaskByIDReq{
		ID:   req.TaskID,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	err = s.provider.SaveLog(ctx, &db.ActivityLog{
		TaskID:  req.TaskID,
		Command: req.Command,
		Data:    task.Data,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ActivityLogTestRes{
		Message: "Success",
	}, nil
}
