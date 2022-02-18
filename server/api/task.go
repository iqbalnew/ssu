package api

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/account_service"
	announcement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/announcement_service"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/company_service"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/notification_service"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setPagination(v *pb.ListTaskRequest) *pb.PaginationResponse {
	res := &pb.PaginationResponse{
		Limit: 10,
		Page:  1,
	}

	if v == nil || v.Limit == 0 && v.Page == 0 {
		res.Limit = -1
		res.Page = -1
		return res
	} else {
		res.Limit = v.Limit
		res.Page = v.Page
	}

	if res.Page == 0 {
		res.Page = 1
	}

	switch {
	case res.Limit > 100:
		res.Limit = 100
	case res.Limit <= 0:
		res.Limit = 10
	}

	return res
}

func (s *Server) GetTaskByTypeID(ctx context.Context, req *pb.GetTaskByTypeIDReq) (*pb.GetTaskByTypeIDRes, error) {
	res := &pb.GetTaskByTypeIDRes{
		Found: false,
		Data:  nil,
	}

	if req.ID < 1 {
		return res, nil
	}

	filter := pb.TaskORM{
		Type:      req.Type,
		FeatureID: req.ID,
	}

	list, err := s.provider.GetListTask(ctx, &filter, "", "", &pb.PaginationResponse{}, &pb.Sort{})
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		res.Found = true
		data, err := list[0].ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		res.Data = &data
	}

	return res, nil
}

func (s *Server) GetListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
	// logrus.Println("After %v", pb)

	var dataorm pb.TaskORM
	if req.Task != nil {
		dataorm, _ = req.Task.ToORM(ctx)
	}

	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task List",
		Data:    []*pb.Task{},
	}

	result.Pagination = setPagination(req)
	sort := &pb.Sort{
		Column:    req.GetSort(),
		Direction: req.GetDir().Enum().String(),
	}
	list, err := s.provider.GetListTask(ctx, &dataorm, req.Filter, req.Query, result.Pagination, sort)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
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

}

func (s *Server) GetTaskGraphStatus(ctx context.Context, req *pb.GraphStatusRequest) (*pb.GraphStatusResponse, error) {
	stat := req.Status.Number()
	data, err := s.provider.GetGraphStatus(ctx, req.Service, uint(stat))
	if err != nil {
		return nil, err
	}
	res := &pb.GraphStatusResponse{}
	res.Code = 200
	res.Error = false
	res.Message = "Graph Data"
	for _, v := range data {
		val := &pb.GraphStatus{
			Status: pb.Statuses(v.Name),
			Type:   v.Type,
			Total:  v.Total,
		}

		res.Data = append(res.Data, val)
	}

	return res, nil
}

func (s *Server) GraphStatusColumnType(ctx context.Context, req *pb.GraphStatusColumnTypeRequest) (*pb.GraphStatusColumnTypeResponse, error) {
	stat := req.Status.Number()
	data, err := s.provider.GetGraphServiceType(ctx, req.Service, uint(stat), req.Column)
	if err != nil {
		return nil, err
	}
	res := &pb.GraphStatusColumnTypeResponse{}
	res.Code = 200
	res.Error = false
	res.Message = "Graph Data"
	for _, v := range data {
		val := &pb.GraphStatusColumnType{
			Status: v.Name,
			Type:   v.Type,
			Total:  v.Total,
		}

		res.Data = append(res.Data, val)
	}

	return res, nil
}

func (s *Server) GetTaskGraphStep(ctx context.Context, req *pb.GraphStepRequest) (*pb.GraphStepResponse, error) {
	step := req.Step.Number()
	stat := req.Status.Number()

	data, err := s.provider.GetGraphStep(ctx, req.Service, uint(step), uint(stat), req.IsIncludeApprove, req.IsIncludeReject)
	if err != nil {
		return nil, err
	}
	res := &pb.GraphStepResponse{}
	res.Code = 200
	res.Error = false
	res.Message = "Graph Data"
	for _, v := range data {
		val := &pb.GraphStep{
			Step:  pb.Steps(v.Name),
			Type:  v.Type,
			Total: v.Total,
		}

		res.Data = append(res.Data, val)
	}

	return res, nil
}

func (s *Server) GetListAnnouncement(ctx context.Context, req *pb.ListRequest) (*pb.ListTaskResponse, error) {
	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Announcement List",
		Data:    []*pb.Task{},
	}

	list, err := s.provider.GetListTaskWithFilter(ctx, &pb.TaskORM{Type: "Announcement"}, nil, nil)
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
	task.Step = 2
	task.Status = 1
	if req.Task.Type == "Announcement" || req.Task.Type == "Notification" || req.Task.Type == "Menu" {
		task.Step = 3
	}
	if req.TaskID > 0 {
		task.Step = 1
		task.TaskID = req.TaskID
		task.Status = 1
		_, err = s.provider.UpdateTask(ctx, &task)
	} else {
		_, err = s.provider.CreateTask(ctx, &task)
	}

	if err != nil {
		return nil, err
	}

	res := &pb.SaveTaskResponse{
		Success: true,
		Data: &pb.Task{
			TaskID: task.TaskID,
		},
	}

	return res, nil
}

func (s *Server) AssignTypeID(ctx context.Context, req *pb.AssignaTypeIDRequest) (*pb.AssignaTypeIDResponse, error) {
	data, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	data.FeatureID = req.FeatureID
	_, err = s.provider.UpdateTask(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	return &pb.AssignaTypeIDResponse{
		Error:   false,
		Code:    200,
		Message: "Created",
	}, nil
}

func (s *Server) SetTask(ctx context.Context, req *pb.SetTaskRequest) (*pb.SetTaskResponse, error) {
	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}
	if task.IsParentActive {
		return nil, status.Errorf(codes.InvalidArgument, "This is child task with active parent, please refer to parent for change status")
	}

	isParent := false
	if task.Data == "[]" {
		isParent = true
	}

	if req.Comment != "" {
		task.Comment = req.Comment
	}

	sendTask := false
	currentStep := task.Step
	currentStatus := task.Status
	switch strings.ToLower(req.Action) {
	case "rework":
		task.Status = 2
		task.Step = 1
	case "approve":

		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 4 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Approved",
				Data:    &taskPb,
			}, nil
		}

		if task.Type == "Announcement" || task.Type == "Notification" || task.Type == "Menu" {
			if currentStep == 1 {
				task.Status = 1
				task.Step = 4
			}
			if currentStep == 3 {
				task.Status = 1
				task.Step = 4
			}
			if currentStep == 4 {
				sendTask = true
				task.Status = 4
			}
		} else {
			if currentStep >= 3 {
				if task.Type == "Company" || task.Type == "Account" || task.Type == "User" || task.Type == "Role" {
					if currentStep == 4 {
						sendTask = true
						task.Status = 4
					} else {
						task.Status = 1
						task.Step++
					}
				} else {
					sendTask = true
					task.Status = 4
				}
			} else {
				task.Status = 1
				task.Step++
			}
		}
	case "reject":
		task.Status = 5
	}

	for i := range task.Childs {
		task.Childs[i].LastApprovedByID = task.LastApprovedByID
		task.Childs[i].LastRejectedByID = task.LastRejectedByID
		if sendTask {
			task.Childs[i].Status = task.Status
			task.Childs[i].Step = task.Step
		}
	}

	updatedTask, err := s.provider.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	reUpdate := false

	taskPb, _ := updatedTask.ToPB(ctx)

	result := &pb.SetTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task Updated",
		Data:    &taskPb,
	}

	if sendTask {
		switch task.Type {
		case "Announcement":
			// data := &dataPublish{
			// 	DataType: "create-task",
			// 	Data:     task.Data,
			// }

			// out, err := json.Marshal(data)
			// if err != nil {
			// 	panic(err)
			// }

			// amqpConfig := amqp.NewDurableQueueConfig(getEnv("AMQP_URI", "amqp://user:bitnami@localhost:5672"))

			// publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
			// if err != nil {
			// 	logrus.Errorf("Failed to create publisher: %s", err)
			// 	return nil, status.Errorf(codes.Internal, "Failed to create publisher: %s", err)
			// }

			// msg := message.NewMessage(watermill.NewUUID(), out)

			// if err := publisher.Publish("announcement.topic", msg); err != nil {
			// 	logrus.Errorf("Failed to publish: %s", err)
			// 	return nil, status.Errorf(codes.Internal, "Failed to publish: %s", err)
			// }
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
				TaskID: task.TaskID,
				Data:   &data,
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
			data.TaskID = task.TaskID

			res, err := companyClient.CreateCompanyGroup(ctx, &data)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Account":
			// data := &dataPublish{
			// 	DataType: "create-account",
			// 	Data:     task.Data,
			// }

			// out, err := json.Marshal(data)
			// if err != nil {
			// 	panic(err)
			// }

			// amqpConfig := amqp.NewDurableQueueConfig(getEnv("AMQP_URI", "amqp://user:bitnami@localhost:5672"))

			// publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
			// if err != nil {
			// 	logrus.Errorf("Failed to create publisher: %s", err)
			// 	return nil, status.Errorf(codes.Internal, "Failed to create publisher: %s", err)
			// }

			// msg := message.NewMessage(watermill.NewUUID(), out)

			// if err := publisher.Publish("account.topic", msg); err != nil {
			// 	logrus.Errorf("Failed to publish: %s", err)
			// 	return nil, status.Errorf(codes.Internal, "Failed to publish: %s", err)
			// }

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			accountConn, err := grpc.Dial(getEnv("ACCOUNT_SERVICE", ":9093"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer accountConn.Close()

			companyClient := account_pb.NewApiServiceClient(accountConn)

			if isParent {
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := account_pb.CreateAccountRequest{}
						account := account_pb.Account{}
						json.Unmarshal([]byte(task.Childs[i].Data), &account)

						data.Data = &account
						data.TaskID = task.Childs[i].TaskID

						res, err := companyClient.CreateAccount(ctx, &data)
						if err != nil {
							return nil, err
						}
						logrus.Println(res)

						task.Childs[i].IsParentActive = false
						reUpdate = true
					}
				}
			} else {
				data := account_pb.CreateAccountRequest{}
				account := account_pb.Account{}
				json.Unmarshal([]byte(task.Data), &account)

				data.Data = &account
				data.TaskID = task.TaskID

				res, err := companyClient.CreateAccount(ctx, &data)
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

		case "Notification":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			notificationConn, err := grpc.Dial(getEnv("NOTIFICATION_SERVICE", ":9094"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Notification Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer notificationConn.Close()

			companyClient := notification_pb.NewApiServiceClient(notificationConn)

			data := notification_pb.CreateNotificationRequest{}
			json.Unmarshal([]byte(task.Data), &data)

			data.TaskID = task.TaskID
			res, err := companyClient.CreateNotification(ctx, &data)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		}
	}

	if reUpdate {
		updatedTask, err = s.provider.UpdateTask(ctx, task)
		if err != nil {
			return nil, err
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
