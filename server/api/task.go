package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/db"
	customAES "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/aes"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	abonnement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/abonnement_service"
	account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/account_service"
	announcement_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/announcement_service"
	beneficiary_account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/beneficiary_account_service"
	bg_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/bg_service"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/company_service"
	deposito_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/deposito_service"
	liquidity_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/liquidity_service"
	menu_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/menu_service"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/notification_service"
	product_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/product_service"
	role_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/role_service"
	sso_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/sso_service"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/system_service"
	users_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/user_service"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/workflow_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	sqlBuilder := &db.QueryBuilder{
		Filter:        "",
		FilterOr:      "",
		CollectiveAnd: "",
		In:            "",
		CustomOrder:   "",
		Sort:          &pb.Sort{},
	}
	list, err := s.provider.GetListTask(ctx, &filter, &pb.PaginationResponse{}, sqlBuilder)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		res.Found = true
		data, err := list[0].ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			// s.logger.Error("GetTaskByTypeID", fmt.Sprintf("%v", err))
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
		Task:     taskPB,
		Limit:    req.Limit,
		Page:     req.Page,
		Sort:     req.Sort,
		Dir:      pb.ListTaskRequestDirection(req.Dir.Number()),
		Filter:   req.Filter,
		Query:    req.Query,
		In:       req.In,
		FilterOr: req.FilterOr,
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

func (s *Server) GetListTaskWithToken(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {
	// logrus.Println("After %v", pb)

	module := ""
	if req.Task != nil {
		if len(req.Task.Type) > 0 {
			module = req.Task.Type
		}
	}
	logrus.Printf("<==== result =======>> %s", module)
	me, err := s.manager.GetMeFromJWT(ctx, "", module)
	if err != nil {
		return nil, err
	}
	logrus.Printf("<==== result =======>> %s", me)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	logrus.Printf("<==== result =======>> %s", ctx)
	var dataorm pb.TaskORM
	if req.Task != nil {
		dataorm, _ = req.Task.ToORM(ctx)
	}
	logrus.Printf("<==== result =======>> %s", dataorm)
	logrus.Printf("<==== result =======>> %s", me.TaskFilter)

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
	sqlBuilder := &db.QueryBuilder{
		Filter:        req.GetFilter(),
		FilterOr:      req.GetFilterOr(),
		CollectiveAnd: req.GetQuery(),
		In:            req.GetIn(),
		MeFilterIn:    me.TaskFilter,
		CustomOrder:   req.GetCustomOrder(),
		Sort:          sort,
	}
	list, err := s.provider.GetListTask(ctx, &dataorm, result.Pagination, sqlBuilder)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			// s.logger.Error("GetListTask", fmt.Sprintf("%v", err))
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		result.Data = append(result.Data, &task)
	}

	return &result, err

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
	sqlBuilder := &db.QueryBuilder{
		Filter:        req.GetFilter(),
		FilterOr:      req.GetFilterOr(),
		CollectiveAnd: req.GetQuery(),
		In:            req.GetIn(),
		CustomOrder:   req.GetCustomOrder(),
		Sort:          sort,
	}
	list, err := s.provider.GetListTask(ctx, &dataorm, result.Pagination, sqlBuilder)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			// s.logger.Error("GetListTask", fmt.Sprintf("%v", err))
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

func (s *Server) GetListTaskPluck(ctx context.Context, req *pb.ListTaskPluckRequest) (*pb.ListTaskPluckResponse, error) {
	result := &pb.ListTaskPluckResponse{
		Data: []string{},
	}

	if req.PluckKey == "" {
		return result, nil
	}

	var dataorm pb.TaskORM
	if req.Task != nil {
		dataorm, _ = req.Task.ToORM(ctx)
	}

	sqlBuilder := &db.QueryBuilder{
		Filter:        req.GetFilter(),
		FilterOr:      req.GetFilterOr(),
		CollectiveAnd: req.GetQuery(),
		In:            req.GetIn(),
		Distinct:      req.GetDistinctKey(),
	}

	list, err := s.provider.GetListTaskPluck(ctx, req.GetPluckKey(), &dataorm, sqlBuilder)
	if err != nil {
		return nil, err
	}

	result.Data = list

	return result, nil
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
	me, err := s.manager.GetMeFromJWT(ctx, "", "")
	if err != nil {
		return nil, err
	}
	step := req.Step.Number()
	stat := req.Status.Number()
	if me.UserType == "ba" {
		me.CompanyID = ""
	}

	data, err := s.provider.GetGraphStep(ctx, me.CompanyID, req.Service, uint(step), uint(stat), req.IsIncludeApprove, req.IsIncludeReject, me.UserType)
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
			// s.logger.Error("GetListAnnouncement", fmt.Sprintf("%v", err))
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
		// s.logger.Error("SaveTaskWithDataEV", fmt.Sprintf("Failed to decrypt taskID, val: %v | %v", req.TaskID, err))
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		// s.logger.Error("SaveTaskWithDataEV", fmt.Sprintf("failed to convert to int: %v", err))
		return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
	}

	taskPB, err := taskEVtoPB(req.Task, aes)
	if err != nil {
		return nil, err
	}

	request := &pb.SaveTaskRequest{
		TaskID:            uint64(taskID),
		Task:              taskPB,
		IsDraft:           req.IsDraft,
		TransactionAmount: req.TransactionAmount,
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
	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 1")
	}
	task, _ := req.Task.ToORM(ctx)
	var err error

	// logrus.Println("==> 01: ", task.Type)
	// logrus.Println("==> taskID: ", task.TaskID)
	// if len(task.Childs) > 0 {
	// 	for i, v := range task.Childs {
	// 		logrus.Println("==> 01: ", i, ": ", v.Type)
	// 		logrus.Println("ParentID: ", v.ParentID)
	// 	}
	// }

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	} else {
		ctx = metadata.NewOutgoingContext(context.Background(), userMD)
		// me, err := s.manager.GetMeFromJWT(ctx, "")

		// if err == nil {
		// 	if getEnv("ENV", "DEV") != "LOCAL" {
		// 		logrus.Println("Send Log to fluentd")
		// 		s.logger.InfoUser(
		// 			"task-save",
		// 			me.UserID,
		// 			me.CompanyID,
		// 			fmt.Sprintf("taskID: %d", req.TaskID),
		// 		)
		// 	}
		// }
	}
	var trailer metadata.MD

	task.Step = 3
	task.Status = 1
	task.LastApprovedByID = 0
	task.LastApprovedByName = ""
	task.LastRejectedByID = 0
	task.LastRejectedByName = ""
	task.DataBak = "{}"

	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 2")
	}

	if task.CompanyID < 1 {
		if currentUser.UserType != "ba" {
			task.CompanyID = currentUser.CompanyID
		}
	}
	if task.HoldingID < 1 {
		if currentUser.UserType != "ba" {
			task.HoldingID = currentUser.HoldingID
		}
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	productConn, err := grpc.Dial(getEnv("PRODUCT_SERVICE", ":9097"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Product Service: %v", err)
		// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Announcement Service: %v", err))

		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	defer productConn.Close()

	productClient := product_pb.NewApiServiceClient(productConn)
	productData, err := productClient.ListProduct(ctx, &product_pb.ListProductRequest{
		Limit: 1,
		Page:  1,
		Product: &product_pb.Product{
			Name: task.Type,
		},
	})
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Failed to get product data: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 3")
	}

	errorProduct := status.Errorf(codes.NotFound, "This task type product, not found")
	if len(productData.Data) < 1 {
		return nil, errorProduct
	} else {
		if productData.Data[0].Name != task.Type {
			return nil, errorProduct
		}
	}

	product := productData.Data[0]

	taskType := []string{"Swift", "Cash Pooling"}

	if product.IsTransactional && contains(taskType, task.Type) && !req.IsDraft { //skip for difference variable name, revisit later
		if req.Task.Type == "Swift" {
			logrus.Println("SaveTaskWithData =================> 4")
		}
		if req.TransactionAmount == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Transaction amount is required")
		}
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
		if err != nil {
			logrus.Errorln("Failed connect to Workflow Service: %v", err)
			// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Workflow Service: %v", err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer workflowConn.Close()

		client := workflow_pb.NewApiServiceClient(workflowConn)
		getWorkflow, err := client.GenerateWorkflow(ctx, &workflow_pb.GenerateWorkflowRequest{
			ProductID:           product.ProductID,
			CompanyID:           currentUser.CompanyID,
			TransactionalNumber: uint64(req.TransactionAmount),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed to generate workflow: %v", err)
			return nil, err
		}

		if getWorkflow.Data == nil {
			return nil, status.Errorf(codes.NotFound, "workflow for this task type not found")
		}

		workflow, err := json.Marshal(getWorkflow.Data)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed to marshal workflow: %v", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		task.WorkflowDoc = string(workflow)
		if req.Task.Type == "Swift" {
			logrus.Println("SaveTaskWithData =================> 5")
		}
	}

	if req.TaskID > 0 {
		findTask, err := s.provider.FindTaskById(ctx, req.TaskID)
		if err != nil {
			return nil, err
		}
		if findTask.DataBak != "{}" || findTask.Data != "" {
			task.DataBak = findTask.DataBak
		}
	}

	if task.DataBak == "" {
		task.DataBak = "{}"
	}

	if len(task.Childs) > 0 {
		for i := range task.Childs {
			if task.Childs[i].DataBak == "" {
				task.Childs[i].DataBak = "{}"
			}
		}
	}

	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 6")
	}

	// if req.Task.Type == "Announcement" || req.Task.Type == "Notification" || req.Task.Type == "Menu:Appearance" || req.Task.Type == "Menu:License" {
	// 	task.Step = 3
	// }

	if req.IsDraft {
		task.Step = 1
		task.Status = 2
	}

	command := "Create"

	logrus.Println("Task save ==> Task Type: ", task.Type)
	logrus.Println("Task save ==> Task ID: ", task.TaskID)

	if req.TaskID > 0 {
		command = "Update"
		task.TaskID = req.TaskID
		task.UpdatedByID = currentUser.UserID
		task.UpdatedByName = currentUser.Username

		_, err = s.provider.UpdateTask(ctx, &task, true)
	} else {
		task.CreatedByID = currentUser.UserID
		task.CreatedByName = currentUser.Username
		task.UpdatedByID = currentUser.UserID
		task.UpdatedByName = currentUser.Username

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

	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 7")
	}

	logrus.Println("Save LOG task, type: ", task.Type)
	// Save activity Log
	if getEnv("ENV", "LOCAL") != "LOCAL" {
		action := "save"
		if req.IsDraft {
			action = "draft"
		}
		logrus.Println("Set to save log")
		err = s.provider.SaveLog(ctx, &db.ActivityLog{
			TaskID:      task.TaskID,
			Command:     command,
			Type:        task.Type,
			Action:      action,
			Description: task.Reasons,
			UserID:      currentUser.UserID,
			Username:    currentUser.Username,
			CompanyID:   currentUser.CompanyID,
			CompanyName: currentUser.CompanyName,
			RoleIDs:     currentUser.RoleIDs,
			Data:        &task,
		})
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
		}
	}

	if req.Task.Type == "Swift" {
		logrus.Println("SaveTaskWithData =================> 8")
	}

	return res, nil
}

func (s *Server) AssignTypeIDEV(ctx context.Context, req *pb.AssignaTypeIDRequestEV) (*pb.AssignaTypeIDResponse, error) {
	key := getEnv("AES_KEY", "Odj12345*12345678901234567890123")
	aes := customAES.NewCustomAES(key)

	text, err := aes.Decrypt(req.TaskID)
	if err != nil {
		logrus.Errorf("val: %v | %v", req.TaskID, err)
		// s.logger.Error("AssignTypeIDEV", fmt.Sprintf("val: %v | %v", req.TaskID, err))
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		// s.logger.Error("AssignTypeIDEV", fmt.Sprintf("%v", err))
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
		// s.logger.Error("AssignTypeID", fmt.Sprintf("%v", err))
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	data.FeatureID = req.FeatureID
	_, err = s.provider.UpdateTask(ctx, data, false)
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
		// s.logger.Error("SetTaskEV", fmt.Sprintf("val: %v | %v", req.TaskID, err))
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}
	taskID, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		fmt.Println(err)
		// s.logger.Error("SetTaskEV", fmt.Sprintf("%v", err))
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

func checkAllowedApproval(md metadata.MD, taskType string, permission string) bool {
	allowed := false
	authorities := []string{}
	//TODO: REVISIT LATTER, skip beneficary and cash polling
	skipProduct := []string{"SSO:User", "SSO:Company", "SSO:Client", "Menu:Appearance", "Menu:License", "Cash Pooling", "Liquidity", "Beneficiary Account", "BG Mapping", "BG Mapping Digital", "Deposito"}

	for _, v := range skipProduct {
		if v == taskType {
			return true
		}
	}

	productName := strings.Replace(taskType, ":", "_", -1)
	productName = strings.ToLower(productName)
	productName = fmt.Sprintf("user-product-%s", productName)

	// typeSplit := strings.Split(taskType, ":")
	// if len(typeSplit) > 1 {
	// 	taskType = typeSplit[0]
	// }

	// for _, v := range user.ProductRoles {
	// 	if v.ProductName == taskType {
	// 		authorities = v.Authorities
	// 		break
	// 	}
	// }

	if len(md[productName]) > 0 {
		result := strings.Split(md[productName][0], ",")
		if len(result) > 0 {
			authorities = result
		}
	}

	if len(authorities) > 0 {
		for _, v := range authorities {

			if v == permission {
				allowed = true
				break
			}
		}
	}

	return allowed

}

func (s *Server) SetTask(ctx context.Context, req *pb.SetTaskRequest) (*pb.SetTaskResponse, error) {
	var err error
	currentUser, userMd, err := s.manager.GetMeFromMD(ctx)
	// currentUser, userMd, err := manager.UserData{
	// 	UserID: 6,
	// }, metadata.MD{}, err
	logrus.Printf("<@@ result @@>1 %s", req)
	if err != nil {
		return nil, err
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

	// if task.Type == "Menu:License" {
	// 	logrus.Println("Menu License Child length 101 : ", len(task.Childs))
	// }

	if task.Type != "System" && task.Type != "Menu:License" {

		if strings.ToLower(req.Action) == "delete" {
			allowed := checkAllowedApproval(userMd, task.Type, "data_entry:maker")
			if !allowed {
				return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
			}
		} else {
			allowed := checkAllowedApproval(userMd, task.Type, "approve:signer")
			if !allowed {
				return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
			}
		}
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
		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 3 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Returned",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 4 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Approved",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 5 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Rejected",
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

		task.LastApprovedByID = 0
		task.LastApprovedByName = ""
		task.LastRejectedByID = currentUser.UserID
		task.LastRejectedByName = currentUser.Username

		if req.Comment != "" {
			task.Comment = req.Comment
		} else {
			task.Comment = "-"
		}

		if req.Reasons != "" {
			task.Reasons = req.Reasons
		} else {
			task.Reasons = "-"
		}

		task.Status = 3
		task.Step = 1
	case "approve":
		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 3 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Returned",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 4 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Approved",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 5 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Rejected",
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
		task.LastRejectedByID = 0
		task.LastRejectedByName = ""

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

				if task.Type == "Company" {

					var opts []grpc.DialOption
					opts = append(opts, grpc.WithInsecure())

					workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9097"), opts...)
					if err != nil {
						logrus.Errorln("Failed connect to Workflow Service: %v", err)
						// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Workflow Service: %v", err))

						return nil, status.Errorf(codes.Internal, "Internal Error")
					}
					defer workflowConn.Close()

					workflowClient := workflow_pb.NewApiServiceClient(workflowConn)

					company := company_pb.CreateCompanyReq{}
					json.Unmarshal([]byte(task.Data), company)

					data := workflow_pb.CreateCompanyWorkflowRequest{}
					data.TaskID = task.TaskID
					data.Data = &workflow_pb.CompanyWorkflows{
						CompanyID:                company.GetData().Company.CompanyID,
						IsTransactionSTP:         company.GetData().Workflow.IsTansactionalStp,
						IsTransactionChecker:     company.GetData().Workflow.TansactionalStpStep.Verifier,
						IsTransactionSigner:      company.GetData().Workflow.TansactionalStpStep.Approver,
						IsTransactionReleaser:    company.GetData().Workflow.TansactionalStpStep.Releaser,
						IsNonTransactionSTP:      company.GetData().Workflow.IsNonTansactionalStp,
						IsNonTransactionChecker:  company.GetData().Workflow.NonTansactionalStpStep.Verifier,
						IsNonTransactionSigner:   company.GetData().Workflow.NonTansactionalStpStep.Approver,
						IsNonTransactionReleaser: company.GetData().Workflow.NonTansactionalStpStep.Releaser,
						CreatedByID:              currentUser.UserID,
					}

					res, err := workflowClient.CreateCompanyWorkflow(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						return nil, err
					}
					logrus.Println(res)

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
		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 3 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Returned",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 4 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Approved",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 5 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Rejected",
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

		task.LastApprovedByID = 0
		task.LastApprovedByName = ""
		task.LastRejectedByID = currentUser.UserID
		task.LastRejectedByName = currentUser.Username

		if req.Comment != "" {
			task.Comment = req.Comment
		} else {
			task.Comment = "-"
		}

		if req.Reasons != "" {
			task.Reasons = req.Reasons
		} else {
			task.Reasons = "-"
		}

		if currentStatus == 6 && (task.DataBak != "" && task.DataBak != "{}") {
			task.Status = 4
			task.Step = 3
			if task.DataBak != "" && task.DataBak != "{}" {
				task.Data = task.DataBak
			}
		} else {
			task.Status = 5
			task.Step = 0
		}

		taskType := []string{"System", "Account", "Beneficiary Account", "Company", "User",
			"Role", "Workflow", "Menu:Appearance", "Menu:License", "BG Mapping", "BG Mapping Digital",
			"Deposito"}

		if contains(taskType, task.Type) {
			if task.DataBak != "" && task.DataBak != "{}" {
				task.Status = 4
				task.Step = 3
				task.Data = task.DataBak
			}
		}

		if task.Type == "Menu:Appearance" || task.Type == "Menu:License" {
			if task.ChildBak != "" && task.ChildBak != "[]" {
				taskChilds := []*pb.TaskORM{}
				err := json.Unmarshal([]byte(task.ChildBak), &taskChilds)
				if err != nil {
					logrus.Errorln("Error Unmarshal Task Childs: ", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}
				task.Childs = taskChilds
			}
		}

	case "delete":
		taskPb, _ := task.ToPB(ctx)
		if currentStatus == 3 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Returned",
				Data:    &taskPb,
			}, nil
		}

		if currentStatus == 4 {
			taskType := []string{"BG Mapping", "BG Mapping Digital"}
			if !contains(taskType, task.Type) {
				return &pb.SetTaskResponse{
					Error:   false,
					Code:    200,
					Message: "Task Status Already Approved",
					Data:    &taskPb,
				}, nil
			}
		}

		if currentStatus == 5 {
			return &pb.SetTaskResponse{
				Error:   false,
				Code:    200,
				Message: "Task Status Already Rejected",
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

		task.LastApprovedByID = 0
		task.LastApprovedByName = ""
		task.LastRejectedByID = 0
		task.LastRejectedByName = ""

		task.Status = 6
		task.Step = 3

		if currentStatus == 2 {
			if task.Type == "BG Mapping" || task.Type == "BG Mapping Digital" {
				if !(task.DataBak == "" || task.DataBak == "{}") {
					task.Status = 4
					task.Step = 1
					task.Data = task.DataBak
				} else {
					task.Status = 7
					task.Step = 1
				}
			}
		}
	}

	logrus.Println("Input Comment" + req.Comment)
	logrus.Println("Input Reasons" + req.Reasons)
	logrus.Println("End Val Comment" + task.Comment)
	logrus.Println("End Val Reasons" + task.Reasons)

	if task.Type == "Menu:License" {
		logrus.Println("Menu License Child length 102 : ", len(task.Childs))
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

	if task.Type == "Menu:License" {
		logrus.Println("Menu License Child length 103 : ", len(task.Childs))
	}

	if sendTask {
		if task.Data != "" && task.Data != "{}" {
			logrus.Println("Save Backup")
			task.DataBak = task.Data
		}
		if task.DataBak == "" {
			task.DataBak = "{}"
		}

		if (task.Type == "Menu:License" || task.Type == "Menu:Appearance") && len(task.Childs) > 0 {
			jsonChilds, err := json.Marshal(task.Childs)
			if err != nil {
				logrus.Println("Error Marshal Childs")
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			task.ChildBak = string(jsonChilds)
		} else {
			task.ChildBak = "[]"
		}
	}

	if task.Type == "Menu:License" {
		logrus.Println("Menu License Child length 104 : ", len(task.Childs))
	}

	updatedTask, err := s.provider.UpdateTask(ctx, task, false)
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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Announcement Service: %v", err))

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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Company Service: %v", err))

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

		case "Deposito":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			DepositoConn, err := grpc.Dial(getEnv("DEPOSITO_SERVICE", ":9201"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Company Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Company Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer DepositoConn.Close()
			logrus.Println("dicek ini", task.TaskID)

			depositoClient := deposito_pb.NewDepositoServiceClient(DepositoConn)

			data := &deposito_pb.CreateDepositoRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)
			data.Data.Deposito.DepositoID = task.FeatureID
			data.Data.Task = &deposito_pb.Task{
				TaskID:             task.TaskID,
				FeatureID:          task.FeatureID,
				LastApprovedByName: task.LastApprovedByName,
				LastRejectedByName: task.LastRejectedByName,
				CreatedByName:      task.CreatedByName,
				UpdatedByName:      task.UpdatedByName,
				Reasons:            task.Reasons,
				Comment:            task.Comment,
			}

			logrus.Println(data, "cek data")

			res, err := depositoClient.CreateDeposito(ctx, data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Printf("[create deposito data] data : %v", res)

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
				logrus.Errorln("Failed connect to Account Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Account Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer accountConn.Close()

			companyClient := account_pb.NewApiServiceClient(accountConn)

			if isParent {
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := account_pb.CreateAccountRequest{}
						// account := account_pb.AccountTaskDataString{}
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

				data.Data = &account_pb.Account{
					AccountID:        account.AccountID,
					CompanyID:        account.CompanyID,
					AccountNumber:    account.AccountNumber,
					AccountAlias:     account.AccountAlias,
					AccountName:      account.AccountName,
					AccountType:      account.AccountType,
					AccountStatus:    account.AccountStatus,
					AccountCurrency:  account.AccountCurrency,
					AccessLevel:      account.AccessLevel,
					IsOwnedByCompany: account.IsOwnedByCompany,
					CreatedByID:      account.CreatedByID,
					UpdatedByID:      account.UpdatedByID,
					DeletedByID:      account.DeletedByID,
					CreatedAt:        account.CreatedAt,
					UpdatedAt:        account.UpdatedAt,
					DeletedAt:        account.DeletedAt,
				}
				data.TaskID = task.TaskID

				logrus.Printf("Task Account for save ===>: %v", data)

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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Notification Service: %v", err))

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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to User Service: %v", err))

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
				logrus.Errorln("Failed connect to Menu Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Menu Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)
			logrus.Println("@@@@@@@@> Task Type: ", task.Data)

			if strings.Contains(task.Data, `"isParent": true`) {
				// isParent = true
				beforeSave := true

				logrus.Println("@@@@@@@@***> Task Type: ", task.Childs)
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := menu_pb.SaveMenuAppearanceReq{}
						menu := menu_pb.MenuAppearance{}
						json.Unmarshal([]byte(task.Childs[i].Data), &menu)

						if beforeSave {

							_, err := menuClient.BeforeSaveMenuAppearance(ctx, &menu_pb.BeforeSaveMenuAppearanceReq{
								UserType: menu.UserType,
							}, grpc.Header(&header), grpc.Trailer(&trailer))
							if err != nil {
								return nil, err
							}
							beforeSave = false
						}

						data.Data = &menu
						data.TaskID = task.Childs[i].TaskID
						logrus.Println("@@@@@@@@000> Task Type: ", data)

						res, err := menuClient.SaveMenuAppearance(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							return nil, err
						}
						logrus.Println("@@@@@@@@111> Task Type: ", res)
						// logrus.Println(res)

						// task.Childs[i].IsParentActive = false
						// reUpdate = true
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
				logrus.Println("@@@@@@@@222> Task Type: ", res)
				// logrus.Println(res)
			}

		case "Menu:License":
			fmt.Println("Menu:License")
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			menuConn, err := grpc.Dial(getEnv("MENU_SERVICE", ":9096"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Menu Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Menu Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)

			fmt.Println("result", strings.Contains(task.Data, `"isParent": true`))
			isDeleted := false
			if strings.Contains(task.Data, `"isParent": true`) {
				// isParent = true
				for i := range task.Childs {
					if task.Childs[i].IsParentActive {
						data := menu_pb.SaveMenuLicenseReq{}
						menu := menu_pb.MenuLicenseSave{}
						json.Unmarshal([]byte(task.Childs[i].Data), &menu)

						data.Data = &menu
						data.TaskID = task.Childs[i].TaskID
						fmt.Println("data ", data.TaskID)
						if !isDeleted {
							_, err := menuClient.DeleteMenuLicenseCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
							if err != nil {
								return nil, err
							}
							isDeleted = true
						}
						res, err := menuClient.SaveMenuLicense(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							return nil, err
						}
						logrus.Println(res)
						logrus.Printf("3-3 create =======> %v", "done")

						// task.Childs[i].IsParentActive = false
						// reUpdate = true
					}
				}
			} else {
				data := menu_pb.SaveMenuLicenseReq{}
				menu := menu_pb.MenuLicenseSave{}
				json.Unmarshal([]byte(task.Data), &menu)

				data.Data = &menu
				data.TaskID = task.TaskID

				if !isDeleted {
					_, err := menuClient.DeleteMenuLicenseCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						return nil, err
					}
					isDeleted = true
				}

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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Role Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer roleConn.Close()

			client := role_pb.NewApiServiceClient(roleConn)

			data := role_pb.CreateRoleRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)

			data.TaskID = task.TaskID
			data.Data.RoleID = task.FeatureID

			logrus.Println("==> 0create role: %s", data)
			if task.Status == 7 {
				res, err := client.DeleteRole(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Printf("[Delete User] data : %v", res)
			} else {
				res, err := client.CreateRole(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				task.FeatureID = res.Data.RoleID
				Role, _ := json.Marshal(res.Data)
				task.Data = string(Role)
				logrus.Println(task.Data, "hasil Data")
				logrus.Println(res, "hasil res")
				reUpdate = true
			}

		case "Workflow":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Workflow Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Workflow Service: %v", err))
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer workflowConn.Close()

			client := workflow_pb.NewApiServiceClient(workflowConn)

			data := workflow_pb.CreateWorkflowRequest{}
			workflowTask := workflow_pb.WorkflowTask{}
			json.Unmarshal([]byte(task.Data), &workflowTask)
			workflowTask.Workflow.WorkflowID = task.FeatureID
			workflowTask.Task = &workflow_pb.Task{
				TaskID: task.TaskID,
			}

			modul := uint64(0)
			if len(workflowTask.Workflow.ModuleID) > 0 {
				modul = workflowTask.Workflow.ModuleID[0]
			}

			data.Data = &workflow_pb.Workflow{
				WorkflowID:            workflowTask.Workflow.WorkflowID,
				ModuleID:              modul,
				CompanyID:             workflowTask.Workflow.CompanyID,
				CurrencyID:            workflowTask.Workflow.CurrencyID,
				CreatedByID:           currentUser.UserID,
				UpdatedByID:           currentUser.UserID,
				CurrencyName:          workflowTask.Workflow.CurrencyName,
				WorkflowCode:          workflowTask.Workflow.WorkflowCode,
				Description:           workflowTask.Workflow.Description,
				CreatedAt:             &timestamppb.Timestamp{},
				UpdatedAt:             &timestamppb.Timestamp{},
				IsCreatedInputAccount: workflowTask.Workflow.IsCreatedInputAccount,
				IsCustomInputAccount:  workflowTask.Workflow.IsCustomInputAccount,
				Logics:                []*workflow_pb.WorkflowLogic{},
			}
			data.TaskID = task.TaskID

			res, err := client.CreateWorkflow(ctx, &workflowTask, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

		case "Liquidity":
			logrus.Println("Liquidity")
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			liquidityConn, err := grpc.Dial(getEnv("LIQUIDITY_SERVICE", ":9010"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Liquidity Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Liquidity Service: %v", err))
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer liquidityConn.Close()

			client := liquidity_pb.NewApiServiceClient(liquidityConn)

			data := liquidity_pb.CreateLiquidityRequest{}
			liquidityTask := liquidity_pb.CreateTaskLiquidityRequest{}
			json.Unmarshal([]byte(task.Data), &liquidityTask)

			data.Data = &liquidityTask
			data.TaskID = task.TaskID

			res, err := client.CreateLiquidity(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
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
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to SSO Service: %v", err))
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
			}

		case "SSO:Company":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			ssoConn, err := grpc.Dial(getEnv("SSO_SERVICE", ":9106"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to SSO Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to SSO Service: %v", err))

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
			}
		case "System":
			logrus.Println("System Creataion Triggered ========>")
			logrus.Println()
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			systemConn, err := grpc.Dial(getEnv("SYSTEM_SERVICE", ":9101"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to system Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to system Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer systemConn.Close()

			systemClient := system_pb.NewApiServiceClient(systemConn)

			data := system_pb.CreateRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)
			data.TaskID = task.TaskID
			res, err := systemClient.CreateSystem(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)
		case "Subscription":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			abonnementConn, err := grpc.Dial(getEnv("ABONNEMENT_SERVICE", ":9100"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Abonnement Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Abonnement Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer abonnementConn.Close()

			abonnementClient := abonnement_pb.NewApiServiceClient(abonnementConn)

			data := abonnement_pb.CreateAbonnementRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)
			data.TaskID = task.TaskID
			data.Data.Id = task.FeatureID
			if len(data.Data.BillingStatus) < 1 {
				data.Data.BillingStatus = "Waiting Schedule"
			}

			res, err := abonnementClient.CreateAbonnement(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}
			logrus.Println(res)

			// update task billing status
			dataUpdate, err := json.Marshal(data.Data)
			if err != nil {
				logrus.Errorln("Failed to marshal data: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			task.Data = string(dataUpdate)
			task.FeatureID = res.Data.Id
			reUpdate = true

		case "Beneficiary Account":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			beneficiaryAccountConn, err := grpc.Dial(getEnv("BENEFICIARY_ACCOUNT_SERVICE", ":9093"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to Account Service: %v", err)
				// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Account Service: %v", err))

				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer beneficiaryAccountConn.Close()

			beneficiaryAccountClient := beneficiary_account_pb.NewApiServiceClient(beneficiaryAccountConn)

			data := beneficiary_account_pb.CreateBeneficiaryAccountRequest{}
			json.Unmarshal([]byte(task.Data), &data.Data)
			data.TaskID = task.TaskID

			if task.Status == 7 {
				// deleteReq := data.
				data.Data.BeneficiaryAccountID = task.FeatureID
				res, err := beneficiaryAccountClient.DeleteBeneficiaryAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
				logrus.Printf("[Delete Bneficiary] data : %v", res)
			} else {

				res, err := beneficiaryAccountClient.CreateBeneficiaryAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					return nil, err
				}
				logrus.Println(res)
				// update task billing status
				task.FeatureID = res.Data.BeneficiaryAccountID
				reUpdate = true
			}
		case "BG Mapping":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to BG Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer bgConn.Close()

			bgClient := bg_pb.NewApiServiceClient(bgConn)

			data := &bg_pb.CreateTransactionRequest{
				TaskID: task.TaskID,
			}
			_, err = bgClient.CreateTransaction(ctx, data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("Failed connect to create transaction: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
		case "BG Mapping Digital":
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
			if err != nil {
				logrus.Errorln("Failed connect to BG Service: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer bgConn.Close()

			bgClient := bg_pb.NewApiServiceClient(bgConn)

			data := &bg_pb.CreateTransactionRequest{
				TaskID: task.TaskID,
			}
			_, err = bgClient.CreateTransaction(ctx, data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("Failed connect to create transaction: %v", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
		}

	}

	if reUpdate {
		// logrus.Println(task.Data)
		updatedTask, err = s.provider.UpdateTask(ctx, task, false)
		if err != nil {
			return nil, err
		}
	}
	if updatedTask.Reasons != "" && req.Action == "approve" {
		updatedTask.Reasons = ""
	}
	logrus.Println("Save LOG task 'update', type: ", task.Type)
	// Save activity Log
	if getEnv("ENV", "LOCAL") != "LOCAL" {
		logrus.Println("Set to save log")
		err = s.provider.SaveLog(ctx, &db.ActivityLog{
			TaskID:      updatedTask.TaskID,
			Command:     "Update",
			Type:        updatedTask.Type,
			Action:      req.Action,
			Description: updatedTask.Reasons,
			UserID:      currentUser.UserID,
			Username:    currentUser.Username,
			CompanyID:   currentUser.CompanyID,
			CompanyName: currentUser.CompanyName,
			RoleIDs:     currentUser.RoleIDs,
			Data:        updatedTask,
		})
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
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

	sqlBuilder := &db.QueryBuilder{
		Filter:        "",
		FilterOr:      "",
		CollectiveAnd: "",
		In:            "",
		CustomOrder:   "",
		Sort:          &pb.Sort{},
	}
	list, err := s.provider.GetListTask(ctx, &filter, &pb.PaginationResponse{}, sqlBuilder)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		res.Found = true
		data, err := list[0].ToPB(ctx)
		if err != nil {
			logrus.Errorln(err)
			// s.logger.Error("GetTaskByID", fmt.Sprintf("Error: %v", err))
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
		res.Data = &data
	}

	return res, nil
}

func (s *Server) UpdateTaskData(ctx context.Context, req *pb.UpdateTaskDataReq) (*pb.UpdateTaskDataRes, error) {
	res := &pb.UpdateTaskDataRes{
		Success: false,
	}

	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	if task.Type != req.Type {
		return res, nil
	}

	task.Data = req.Data
	_, err = s.provider.UpdateTask(ctx, task, false)
	if err != nil {
		return nil, err
	}

	res.Success = true
	return res, nil
}

func (s *Server) UpdateTaskPlain(ctx context.Context, req *pb.SaveTaskRequest) (*pb.SaveTaskResponse, error) {
	task, _ := req.Task.ToORM(ctx)
	var err error

	// if len(task.Childs) > 0 {
	// 	for i, v := range task.Childs {
	// 	}
	// }

	task.Step = 3
	task.Status = 1
	task.LastApprovedByID = 0
	task.LastApprovedByName = ""
	task.LastRejectedByID = 0
	task.LastRejectedByName = ""
	task.DataBak = "{}"

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	productConn, err := grpc.Dial(getEnv("PRODUCT_SERVICE", ":9097"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Product Service: %v", err)
		// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Announcement Service: %v", err))

		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	defer productConn.Close()

	productClient := product_pb.NewApiServiceClient(productConn)
	productData, err := productClient.ListProduct(ctx, &product_pb.ListProductRequest{
		Limit: 1,
		Page:  1,
		Product: &product_pb.Product{
			Name: task.Type,
		},
	})
	if err != nil {
		logrus.Errorln("[api][func: UpdateTaskPlain] Failed to get product data: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	errorProduct := status.Errorf(codes.NotFound, "This task type product, not found")
	if len(productData.Data) < 1 {
		return nil, errorProduct
	} else {
		if productData.Data[0].Name != task.Type {
			return nil, errorProduct
		}
	}

	product := productData.Data[0]

	taskType := []string{"Swift", "Cash Pooling"}

	if product.IsTransactional && contains(taskType, task.Type) && !req.IsDraft { //skip for difference variable name, revisit later
		if req.TransactionAmount == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Transaction amount is required")
		}
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
		if err != nil {
			logrus.Errorln("Failed connect to Workflow Service: %v", err)
			// s.logger.Error("SetTask", fmt.Sprintf("Failed connect to Workflow Service: %v", err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer workflowConn.Close()

		client := workflow_pb.NewApiServiceClient(workflowConn)
		getWorkflow, err := client.GenerateWorkflow(ctx, &workflow_pb.GenerateWorkflowRequest{
			ProductID:           product.ProductID,
			TransactionalNumber: uint64(req.TransactionAmount),
		})
		if err != nil {
			logrus.Errorln("[api][func: UpdateTaskPlain] Failed to generate workflow: %v", err)
			return nil, err
		}

		if getWorkflow.Data == nil {
			return nil, status.Errorf(codes.NotFound, "workflow for this task type not found")
		}

		workflow, err := json.Marshal(getWorkflow.Data)
		if err != nil {
			logrus.Errorln("[api][func: UpdateTaskPlain] Failed to marshal workflow: %v", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		task.WorkflowDoc = string(workflow)
	}

	if req.TaskID > 0 {
		findTask, err := s.provider.FindTaskById(ctx, req.TaskID)
		if err != nil {
			return nil, err
		}
		if findTask.DataBak != "{}" || findTask.Data != "" {
			task.DataBak = findTask.DataBak
		}
	}

	if task.DataBak == "" {
		task.DataBak = "{}"
	}

	if len(task.Childs) > 0 {
		for i := range task.Childs {
			if task.Childs[i].DataBak == "" {
				task.Childs[i].DataBak = "{}"
			}
		}
	}

	// if req.Task.Type == "Announcement" || req.Task.Type == "Notification" || req.Task.Type == "Menu:Appearance" || req.Task.Type == "Menu:License" {
	// 	task.Step = 3
	// }

	if req.IsDraft {
		task.Step = 1
		task.Status = 2
	}

	command := "Create"

	logrus.Println("Task save ==> Task Type: ", task.Type)
	logrus.Println("Task save ==> Task ID: ", task.TaskID)

	if req.TaskID > 0 {
		command = "Update"
		task.TaskID = req.TaskID

		_, err = s.provider.UpdateTask(ctx, &task, true)
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

	logrus.Println("Save LOG task, type: ", task.Type)
	// Save activity Log
	if getEnv("ENV", "LOCAL") != "LOCAL" {
		action := "save"
		if req.IsDraft {
			action = "draft"
		}
		logrus.Println("Set to save log")
		err = s.provider.SaveLog(ctx, &db.ActivityLog{
			TaskID:      task.TaskID,
			Command:     command,
			Type:        task.Type,
			Action:      action,
			Description: task.Reasons,
			Data:        &task,
		})
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
		}
	}

	return res, nil
}
