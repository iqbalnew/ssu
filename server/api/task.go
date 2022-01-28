package api

import (
	"context"
	"encoding/json"
	"os"

	announcement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/announcement_pb"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/company_pb"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/notification_pb"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
	// logrus.Println("After %v", pb)

	dataorm, _ := req.Task.ToORM(ctx)

	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Announcement List",
		Data:    []*pb.Task{},
	}

	list, err := s.provider.GetListTaskWithFilter(ctx, &dataorm)
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

	return &result, err

	// list, err := s.provider.GetListTaskWithFilter(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, v := range list {
	// 	task, err := v.ToPB(ctx)
	// 	if err != nil {
	// 		logrus.Errorln(err)
	// 		return nil, status.Errorf(codes.Internal, "Internal Error")
	// 	}
	// 	result.Data = append(result.Data, &task)
	// }

	// return &result, nil
	return nil, nil

}

func (s *Server) GetListAnnouncement(ctx context.Context, req *pb.ListRequest) (*pb.ListTaskResponse, error) {
	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Announcement List",
		Data:    []*pb.Task{},
	}

	list, err := s.provider.GetListTaskWithFilter(ctx, &pb.TaskORM{Type: "Announcement"})
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
		_, err = s.provider.UpdateTask(ctx, &task)
	} else {
		_, err = s.provider.CreateTask(ctx, &task)
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
				Message: "Task Status Already Approved",
				Data:    &taskPb,
			}, nil
		}

		if task.Type == "Announcement" || task.Type == "Notification" || task.Type == "Menu" {
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
				if task.Type == "Company" || task.Type == "Account" || task.Type == "User" || task.Type == "Role" {
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

	updatedTask, err := s.provider.UpdateTask(ctx, task)
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
		case "Announcement":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			announcementConn, err := grpc.Dial(getEnv("ANNOUNCEMENT_SERVICE", ":9091"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Announcement Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer announcementConn.Close()

			announcementClient := announcement_pb.NewApiServiceClient(announcementConn)

			data := announcement_pb.Announcement{}
			json.Unmarshal([]byte(task.Data), &data)
			send := &announcement_pb.CreateAnnouncementRequest{
				Data: &data,
			}
			res, err := announcementClient.CreateAnnouncement(ctx, send)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Company":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer companyConn.Close()

			companyClient := company_pb.NewApiServiceClient(companyConn)

			data := company_pb.CreateCompanyGroupRequest{}
			json.Unmarshal([]byte(task.Data), &data)

			res, err := companyClient.CreateCompanyGroup(ctx, &data)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Notification":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			notificationConn, err := grpc.Dial(getEnv("NOTIFICATION_SERVICE", ":9094"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer notificationConn.Close()

			companyClient := notification_pb.NewApiServiceClient(notificationConn)

			data := notification_pb.CreateNotificationRequest{}
			json.Unmarshal([]byte(task.Data), &data)

			res, err := companyClient.CreateNotification(ctx, &data)
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
