package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	customAES "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/aes"
	manager "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/jwt"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/account_service"
	announcement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/announcement_service"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/company_service"
	liquidity_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/liquidity_service"
	menu_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/menu_service"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/notification_service"
	role_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/role_service"
	sso_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/sso_service"
	users_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/user_service"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/workflow_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

func (s *Server) GetListTaskEV(ctx context.Context, req *pb.ListTaskRequestEV) (*pb.ListTaskResponseEV, error) {
	key := getEnv("AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	taskPB, err := taskEVtoPB(req.Task, aes)
	if err != nil {
		return nil, err
	}

	reqPB := &pb.ListTaskRequest{
		Task:   taskPB,
		Limit:  req.Limit,
		Page:   req.Page,
		Sort:   req.Sort,
		Dir:    pb.ListTaskRequestDirection(req.Dir.Number()),
		Filter: req.Filter,
		Query:  req.Query,
	}

	resPB, err := s.GetListTask(ctx, reqPB)
	if err != nil {
		return nil, err
	}

	res := &pb.ListTaskResponseEV{
		Error:      resPB.Error,
		Code:       resPB.Code,
		Message:    resPB.Message,
		Pagination: resPB.Pagination,
	}

	for _, v := range resPB.Data {
		task, err := taskPBtoEV(v, aes)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, task)
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

func (s *Server) SaveTaskWithDataEV(ctx context.Context, req *pb.SaveTaskRequestEV) (*pb.SaveTaskResponseEV, error) {
	key := getEnv("AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	text, err := aes.Decrypt(req.TaskID)
	if err != nil {
		logrus.Errorf("val: %v | %v", req.TaskID, err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
	}

	taskPB, err := taskEVtoPB(req.Task, aes)
	if err != nil {
		return nil, err
	}

	request := &pb.SaveTaskRequest{
		TaskID:  uint64(taskID),
		Task:    taskPB,
		IsDraft: req.IsDraft,
	}

	response, err := s.SaveTaskWithData(ctx, request)
	if err != nil {
		return nil, err
	}

	taskEV, _ := taskPBtoEV(response.Data, aes)

	res := &pb.SaveTaskResponseEV{
		Success: response.Success,
		Data:    taskEV,
	}

	return res, nil
}

func (s *Server) SaveTaskWithData(ctx context.Context, req *pb.SaveTaskRequest) (*pb.SaveTaskResponse, error) {
	task, _ := req.Task.ToORM(ctx)
	var err error

	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		if getEnv("ENV", "DEV") == "PROD" {
			return nil, status.Errorf(codes.Unauthenticated, "%v", err)
		} else {
			currentUser = &manager.VerifyTokenRes{
				UserID:   1,
				Username: "",
			}
		}
	}
	fmt.Println(">>")
	fmt.Printf("Data Current User =>>>>>>>>>>>>>>>>>>>>> %v", currentUser)
	fmt.Println(">>")

	task.Step = 3
	task.Status = 1

	// if req.Task.Type == "Announcement" || req.Task.Type == "Notification" || req.Task.Type == "Menu:Appearance" || req.Task.Type == "Menu:License" {
	// 	task.Step = 3
	// }

	if req.IsDraft {
		task.Step = 1
		task.Status = 2
	}

	if req.TaskID > 0 {
		task.TaskID = req.TaskID
		task.UpdatedByID = currentUser.UserID
		task.UpdatedByName = currentUser.Username

		fmt.Println(">>")
		fmt.Printf("Data Task Update => %d|%s", task.UpdatedByID, task.UpdatedByName)
		fmt.Println(">>")

		_, err = s.provider.UpdateTask(ctx, &task)
	} else {
		task.CreatedByID = currentUser.UserID
		task.CreatedByName = currentUser.Username
		task.UpdatedByID = currentUser.UserID
		task.UpdatedByName = currentUser.Username

		fmt.Println(">>")
		fmt.Printf("Data Task Create => %d|%s|%d|%s", task.CreatedByID, task.CreatedByName, task.UpdatedByID, task.UpdatedByName)
		fmt.Println(">>")

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

func (s *Server) AssignTypeIDEV(ctx context.Context, req *pb.AssignaTypeIDRequestEV) (*pb.AssignaTypeIDResponse, error) {
	key := getEnv("AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	text, err := aes.Decrypt(req.TaskID)
	if err != nil {
		logrus.Errorf("val: %v | %v", req.TaskID, err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
	}

	reqPB := &pb.AssignaTypeIDRequest{
		TaskID:    uint64(taskID),
		FeatureID: req.FeatureID,
		Type:      req.Type,
	}

	return s.AssignTypeID(ctx, reqPB)
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

func (s *Server) SetTaskEV(ctx context.Context, req *pb.SetTaskRequestEV) (*pb.SetTaskResponseEV, error) {
	key := getEnv("AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	text, err := aes.Decrypt(req.TaskID)
	if err != nil {
		logrus.Errorf("val: %v | %v", req.TaskID, err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
	}

	reqPB := &pb.SetTaskRequest{
		TaskID:  uint64(taskID),
		Action:  req.Action,
		Comment: req.Comment,
		Reasons: req.Reasons,
	}

	resPB, err := s.SetTask(ctx, reqPB)
	if err != nil {
		return nil, err
	}

	taskEV, _ := taskPBtoEV(resPB.Data, aes)

	res := &pb.SetTaskResponseEV{
		Error:   resPB.Error,
		Code:    resPB.Code,
		Message: resPB.Message,
		Data:    taskEV,
	}

	return res, nil
}

func (s *Server) SetTask(ctx context.Context, req *pb.SetTaskRequest) (*pb.SetTaskResponse, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		if getEnv("ENV", "DEV") == "PROD" {
			return nil, status.Errorf(codes.Unauthenticated, "%v", err)
		} else {
			currentUser = &manager.VerifyTokenRes{
				UserID:   1,
				Username: "",
			}
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

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

	if req.Reasons != "" {
		task.Reasons = req.Reasons
	}

	sendTask := false
	currentStep := task.Step
	currentStatus := task.Status
	switch strings.ToLower(req.Action) {
	case "rework":
		task.LastRejectedByID = currentUser.UserID
		task.LastRejectedByName = currentUser.Username

		task.Status = 3
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

		if currentStatus == 7 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Already Deleted",
				Data:    &taskPb,
			}, nil
		}

		task.LastApprovedByID = currentUser.UserID
		task.LastApprovedByName = currentUser.Username

		// Fix Workflow Step 2
		if currentStep == 2 {
			currentStep = 3
		}

		if currentStatus == 2 {
			task.Step = 3
			task.Status = 1

			// if task.Type == "Announcement" || task.Type == "Notification" || task.Type == "Menu:Appearance" || task.Type == "Menu:License" {
			// 	task.Step = 3
			// }

		} else {

			// if task.Type == "Announcement" || task.Type == "Notification" || task.Type == "Menu:Appearance" || task.Type == "Menu:License" {
			if currentStep == 1 {
				task.Status = 1
				task.Step = 3
				if currentStatus == 6 {
					task.Status = currentStatus
				}
			}
			if currentStep == 3 {
				task.Status = 1
				// task.Step = 4
				// if task.Type == "Announcement" {
				task.Status = 4
				task.Step = 3
				sendTask = true
				if currentStatus == 6 {
					task.Status = 7
				}
				// }
				// if currentStatus == 6 {
				// 	task.Status = currentStatus
				// }
			}
			if currentStep == 4 {
				sendTask = true
				task.Status = 4
				if currentStatus == 6 {
					task.Status = 7
				}
			}
			// } else {
			// 	if currentStep >= 3 {
			// 		if task.Type == "Company" || task.Type == "Account" || task.Type == "User" || task.Type == "Role" || task.Type == "Workflow" || task.Type == "Liquidity" {
			// 			if currentStep == 4 {
			// 				sendTask = true
			// 				task.Status = 4
			// 				if currentStatus == 6 {
			// 					task.Status = 7
			// 				}
			// 			} else {
			// 				task.Status = 1
			// 				task.Step++
			// 				if currentStatus == 6 {
			// 					task.Status = currentStatus
			// 				}
			// 			}
			// 		} else {
			// 			sendTask = true
			// 			task.Status = 4
			// 			if currentStatus == 6 {
			// 				task.Status = 7
			// 			}
			// 		}
			// 	} else {
			// 		task.Status = 1
			// 		task.Step++
			// 		if currentStatus == 6 {
			// 			task.Status = currentStatus
			// 		}
			// 	}
			// }
		}

	case "reject":
		task.LastRejectedByID = currentUser.UserID
		task.LastRejectedByName = currentUser.Username

		task.Status = 5
		task.Step = 0

	case "delete":
		task.LastApprovedByID = currentUser.UserID
		task.LastApprovedByName = currentUser.Username

		task.Status = 6
		task.Step = 3

	}

	for i := range task.Childs {
		task.Childs[i].LastApprovedByID = task.LastApprovedByID
		task.Childs[i].LastApprovedByName = task.LastApprovedByName
		task.Childs[i].LastRejectedByID = task.LastRejectedByID
		task.Childs[i].LastRejectedByName = task.LastRejectedByName
		task.Childs[i].CreatedByID = task.CreatedByID
		task.Childs[i].CreatedByName = task.CreatedByName
		task.Childs[i].UpdatedByID = task.UpdatedByID
		task.Childs[i].UpdatedByName = task.UpdatedByName
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
			res, err := announcementClient.CreateAnnouncement(ctx, send, grpc.Header(&header), grpc.Trailer(&trailer))
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

			data := company_pb.CreateCompanyReq{}
			json.Unmarshal([]byte(task.Data), &data.Data)
			data.TaskID = task.TaskID

			if task.Status == 7 {
				// deleteReq := data.
				res, err := companyClient.DeleteCompany(ctx, data.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Printf("[Delete Company] data : %v", res)
			} else {
				res, err := companyClient.CreateCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

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

						res, err := companyClient.CreateAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
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

				res, err := companyClient.CreateAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
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
			res, err := companyClient.CreateNotification(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "User":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			usersConn, err := grpc.Dial(getEnv("USER_SERVICE", ":9095"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to User Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer usersConn.Close()

			client := users_pb.NewApiServiceClient(usersConn)

			data := users_pb.CreateUserRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)

			data.TaskID = task.TaskID
			if task.Status == 7 {
				deleteReq := &users_pb.DeleteUserReq{UserID: data.Data.User.UserID}
				res, err := client.DeleteUser(ctx, deleteReq, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Printf("[Delete User] data : %v", res)
			} else {
				res, err := client.CreateUser(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

		case "Menu:Appearance":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			menuConn, err := grpc.Dial(getEnv("MENU_SERVICE", ":9093"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)

			if isParent {
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := menu_pb.SaveMenuAppearanceReq{}
						menu := menu_pb.MenuAppearance{}
						json.Unmarshal([]byte(task.Childs[i].Data), &menu)

						data.Data = &menu
						data.TaskID = task.Childs[i].TaskID

						res, err := menuClient.SaveMenuAppearance(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							return nil, err
						}
						logrus.Println(res)

						task.Childs[i].IsParentActive = false
						reUpdate = true
					}
				}
			} else {
				data := menu_pb.SaveMenuAppearanceReq{}
				menu := menu_pb.MenuAppearance{}
				json.Unmarshal([]byte(task.Data), &menu)

				data.Data = &menu
				data.TaskID = task.TaskID

				res, err := menuClient.SaveMenuAppearance(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

		case "Menu:License":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			menuConn, err := grpc.Dial(getEnv("MENU_SERVICE", ":9096"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)

			if isParent {
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := menu_pb.SaveMenuLicenseReq{}
						menu := menu_pb.MenuLicenseSave{}
						json.Unmarshal([]byte(task.Childs[i].Data), &menu)

						data.Data = &menu
						data.TaskID = task.Childs[i].TaskID

						res, err := menuClient.SaveMenuLicense(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							return nil, err
						}
						logrus.Println(res)

						task.Childs[i].IsParentActive = false
						reUpdate = true
					}
				}
			} else {
				data := menu_pb.SaveMenuLicenseReq{}
				menu := menu_pb.MenuLicenseSave{}
				json.Unmarshal([]byte(task.Data), &menu)

				data.Data = &menu
				data.TaskID = task.TaskID

				res, err := menuClient.SaveMenuLicense(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

		case "Role":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			roleConn, err := grpc.Dial(getEnv("ROLE_SERVICE", ":9098"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Role Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer roleConn.Close()

			client := role_pb.NewApiServiceClient(roleConn)

			data := role_pb.CreateRoleRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)

			data.TaskID = task.TaskID
			res, err := client.CreateRole(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Workflow":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Workflow Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer workflowConn.Close()

			client := workflow_pb.NewApiServiceClient(workflowConn)

			data := workflow_pb.CreateWorkflowRequest{}
			workflowTask := workflow_pb.WorkflowTask{}
			json.Unmarshal([]byte(task.Data), &workflowTask)

			data.Data = workflowTask.Workflow
			data.TaskID = task.TaskID

			res, err := client.CreateWorkflow(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Liquidity":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			liquidityConn, err := grpc.Dial(getEnv("LIQUIDITY_SERVICE", ":9010"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Workflow Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer liquidityConn.Close()

			client := liquidity_pb.NewApiServiceClient(liquidityConn)

			data := liquidity_pb.CreateLiquidityRequest{}
			liquidityTask := liquidity_pb.CreateTaskLiquidityRequest{}
			json.Unmarshal([]byte(task.Data), &liquidityTask)

			data.Data = &liquidityTask
			data.TaskID = task.TaskID

			res, err := client.CreateLiquidity(ctx, &data)
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "SSO:User":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			ssoConn, err := grpc.Dial(getEnv("SSO_SERVICE", ":9106"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to SSO Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer ssoConn.Close()

			ssoClient := sso_pb.NewApiServiceClient(ssoConn)

			data := sso_pb.WriteSyncUserTask{}
			json.Unmarshal([]byte(task.Data), &data)
			send := &sso_pb.CreateSyncUserTaskReq{
				Data:   &data,
				TaskID: task.TaskID,
			}
			if task.Status == 7 {
				send.Data.User.UserSyncID = task.FeatureID
				res, err := ssoClient.DeleteSyncUser(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			} else {
				if task.FeatureID > 0 {
					send.Data.User.UserSyncID = task.FeatureID
				}
				res, err := ssoClient.SaveSyncUser(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

		case "SSO:Company":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			ssoConn, err := grpc.Dial(getEnv("SSO_SERVICE", ":9106"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to SSO Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer ssoConn.Close()

			ssoClient := sso_pb.NewApiServiceClient(ssoConn)

			data := sso_pb.WriteSyncCompanyTask{}
			json.Unmarshal([]byte(task.Data), &data)
			send := &sso_pb.CreateSyncCompanyTaskReq{
				Data:   &data,
				TaskID: task.TaskID,
			}
			if task.Status == 7 {
				send.Data.Company.CompanySyncID = task.FeatureID
				res, err := ssoClient.DeleteSyncCompany(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			} else {
				if task.FeatureID > 0 {
					send.Data.Company.CompanySyncID = task.FeatureID
				}
				res, err := ssoClient.SaveSyncCompany(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
			}

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

func (s *Server) RejectBySystem(ctx context.Context, req *pb.RejectBySystemReq) (res *pb.RejectBySystemRes, err error) {
	res = &pb.RejectBySystemRes{
		Success: true,
		Code:    "200",
	}
	_, err = s.provider.FindAndSetStatus(ctx, req.TaskID, req.Status.Descriptor().Index())
	if err != nil {
		res.Success = false
		res.Code = "000"
	}

	return res, nil
}

func (s *Server) GetTaskByID(ctx context.Context, req *pb.GetTaskByIDReq) (*pb.GetTaskByIDRes, error) {
	res := &pb.GetTaskByIDRes{
		Found: false,
		Data:  nil,
	}

	if req.ID < 1 {
		return res, nil
	}

	filter := pb.TaskORM{
		Type:   req.Type,
		TaskID: req.ID,
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
