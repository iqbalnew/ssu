package api

import (
	"context"
	"encoding/json"
	"os"

	announcement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/announcement_pb"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetListTask(ctx context.Context, req *pb.ListRequest) (*pb.ListTaskResponse, error) {
	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Announcement List",
		Data:    []*pb.Task{},
	}

	list, err := s.provider.GetListTask(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		result.Data = append(result.Data, &task)
	}

	return &result, nil

}

func (s *Server) SaveTaskWithData(ctx context.Context, req *pb.SaveTaskRequest) (*pb.SaveTaskResponse, error) {
	task, _ := req.Task.ToORM(ctx)
	var err error
	req.Task.Step = 1
	if req.Task.Type == "announcement" || req.Task.Type == "notification" || req.Task.Type == "menu" {
		req.Task.Step = 2
	}
	if req.TaskID > 0 {
		req.Task.Step = 0
		_, err = s.provider.UpdateTask(ctx, &task, req.Data)
	} else {
		_, err = s.provider.CreateTask(ctx, &task, req.Data)
	}

	if err != nil {
		return nil, err
	}

	res := &pb.SaveTaskResponse{
		Success: true,
	}

	return res, nil
}

func (s *Server) SetTask(ctx context.Context, req *pb.SetTaskRequest) (*pb.SetTaskResponse, error) {
	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	sendTask := false
	currentStep := task.Step
	currentStatus := task.Status
	switch req.Action {
	case "rework":
		task.Status = 1
		task.Step = 0
	case "approve":

		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 3 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Alredy Approved",
				Data:    &taskPb,
			}, nil
		}

		if task.Type == "announcement" || task.Type == "notification" || task.Type == "menu" {
			if currentStep == 0 {
				task.Status = 0
				task.Step = 2
			}
			if currentStep == 2 {
				sendTask = true
				task.Status = 3
			}
		} else {
			if currentStep >= 3 {
				if task.Type == "company" || task.Type == "account" || task.Type == "user" || task.Type == "role" {
					if currentStep == 4 {
						sendTask = true
						task.Status = 3
					} else {
						task.Status = 0
						task.Step++
					}
				} else {
					sendTask = true
					task.Status = 3
				}
			} else {
				task.Status = 0
				task.Step++
			}
		}
	case "reject":
		task.Status = 4
	}

	updatedTask, err := s.provider.UpdateTask(ctx, task, "")
	if err != nil {
		return nil, err
	}

	taskPb, _ := updatedTask.ToPB(ctx)

	result := &pb.SetTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task Updated",
		Data:    &taskPb,
	}

	if sendTask {
		switch updatedTask.Type {
		case "announcement":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			announcementConn, err := grpc.Dial(getEnv("ANNOUNCEMENT_SERVICE", ":9091"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Announcement Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer announcementConn.Close()

			announcementClient := announcement_pb.NewApiServiceClient(announcementConn)

			annoncement, err := s.provider.GetAnnouncementTaskById(ctx, req.TaskID)
			if err != nil {
				return nil, err
			}
			data := announcement_pb.Announcement{}
			json.Unmarshal([]byte(annoncement.Data), &data)
			send := &announcement_pb.CreateAnnouncementRequest{
				Data: &data,
			}
			res, err := announcementClient.CreateAnnouncement(ctx, send)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		}
	}

	return result, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
