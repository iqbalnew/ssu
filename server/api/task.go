package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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
	bifast_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/bifast_service"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/company_service"
	deposito_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/deposito_service"
	liquidity_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/liquidity_service"
	menu_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/menu_service"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/notification_service"
	online_transfer_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/online_transfer_service"
	payroll_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/payroll_service"
	product_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/product_service"
	role_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/role_service"
	sso_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/sso_service"
	swift_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/swift_service"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/system_service"
	transfer_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/transfer_service"
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
	list, err := s.provider.GetListTask(ctx, &filter, &pb.PaginationResponse{}, sqlBuilder, 0, []uint64{}, []uint64{})
	if err != nil {
		logrus.Errorln("[api][func: GetTaskByTypeID] Failed when execute GetListTask:", err)
		return nil, err
	}

	if len(list) > 0 {

		res.Found = true

		data, err := list[0].ToPB(ctx)
		if err != nil {
			logrus.Errorln("[api][func: GetTaskByTypeID] Failed convert ORM to PB:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
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
		logrus.Errorln("[api][func: GetListTaskEV] Failed when execute taskEVtoPB:", err)
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
		logrus.Errorln("[api][func: GetListTaskEV] Failed when execute GetListTask:", err)
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
			logrus.Errorln("[api][func: GetListTaskEV] Failed when execute taskPBtoEV:", err)
			return nil, err
		}

		res.Data = append(res.Data, task)

	}

	return res, nil
}

func (s *Server) GetListTaskWithToken(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {

	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task List",
		Data:    []*pb.Task{},
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	if currentUser == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Permission Denied")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var trailer metadata.MD

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	accountConn, err := grpc.Dial(getEnv("ACCOUNT_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("[api][func: GetListTaskWithToken] Unable to connect Account Service:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	defer accountConn.Close()

	accountClient := account_pb.NewApiServiceClient(accountConn)

	userID := uint64(0)
	companyID := uint64(0)
	roleIDs := []uint64{}
	accountIDs := []uint64{}

	if currentUser.UserType == "cu" {

		listAccountByRoleReq := &account_pb.ListAccountRequest{}

		if req.GetTask() != nil && req.GetTask().GetType() != "" {

			productConn, err := grpc.Dial(getEnv("PRODUCT_SERVICE", ":9090"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to connect Product Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer productConn.Close()

			productClient := product_pb.NewApiServiceClient(productConn)

			listProductRes, err := productClient.ListProduct(ctx, &product_pb.ListProductRequest{
				Product: &product_pb.Product{
					Name: req.GetTask().GetType(),
				},
			})
			if err != nil {
				logrus.Println("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to Get List Product:", err.Error())
				return nil, err
			}

			if len(listProductRes.GetData()) < 1 {
				return nil, status.Errorf(codes.NotFound, "Product Not Found")
			}

			listAccountByRoleReq = &account_pb.ListAccountRequest{
				ProductID: listProductRes.GetData()[0].GetProductID(),
			}

		}

		listAccountRes, err := accountClient.ListAccountByRole(ctx, listAccountByRoleReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			logrus.Println("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to Get Account By Role:", err.Error())
			return nil, err
		}

		for _, v := range listAccountRes.Data {
			accountIDs = append(accountIDs, v.AccountID)
		}

		userID = currentUser.UserID
		roleIDs = currentUser.RoleIDs

	}

	dataORM := &pb.TaskORM{}
	if req.GetTask() != nil {
		dataORM.Type = req.GetTask().GetType()
	}

	sort := &pb.Sort{
		Column:    req.GetSort(),
		Direction: req.GetDir().Enum().String(),
	}

	sqlBuilder := &db.QueryBuilder{
		Filter:   req.GetFilter(),
		FilterOr: req.GetFilterOr(),
		Sort:     sort,
	}

	stepFilter := ""

	if currentUser.UserType == "cu" {

		companyID = currentUser.CompanyID

		switch req.GetTask().GetStep() {
		case pb.Steps_Maker:
			stepFilter = "status:<>0,status:<>1,status:<>4,status:<>5,status:<>6,status:<>7"
		case pb.Steps_Checker:
			stepFilter = "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7,workflow_doc.workflow.currentStep:checker"
		case pb.Steps_Signer:
			stepFilter = "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7,workflow_doc.workflow.currentStep:signer"
		case pb.Steps_Releaser:
			stepFilter = "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7,workflow_doc.workflow.currentStep:releaser"
		}

		if sqlBuilder.Filter != "" {
			sqlBuilder.Filter = fmt.Sprintf("%s,%s", sqlBuilder.Filter, stepFilter)
		} else {
			sqlBuilder.Filter = stepFilter
		}

	} else if currentUser.UserType == "ca" {

		companyID = currentUser.CompanyID

		if req.GetTask().GetStep() != pb.Steps_NullStep {

			dataORM.Step = int32(req.GetTask().GetStep())

			if req.GetTask().GetStep() == pb.Steps_Checker || req.GetTask().GetStep() == pb.Steps_Signer || req.GetTask().GetStep() == pb.Steps_Releaser {

				if sqlBuilder.Filter != "" {
					sqlBuilder.Filter = fmt.Sprintf("%s,%s", sqlBuilder.Filter, "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7")
				} else {
					sqlBuilder.Filter = "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7"
				}

			} else if req.GetTask().GetStep() == pb.Steps_Maker {

				if sqlBuilder.Filter != "" {
					sqlBuilder.Filter = fmt.Sprintf("%s,%s", sqlBuilder.Filter, "status:<>0,status:<>1,status:<>4,status:<>5,status:<>6,status:<>7")
				} else {
					sqlBuilder.Filter = "status:<>0,status:<>1,status:<>4,status:<>5,status:<>6,status:<>7"
				}

			}

		}

	} else {

		if req.GetTask().GetStep() != pb.Steps_NullStep {

			dataORM.Step = int32(req.GetTask().GetStep())

			if req.GetTask().GetStep() == pb.Steps_Checker || req.GetTask().GetStep() == pb.Steps_Signer || req.GetTask().GetStep() == pb.Steps_Releaser {

				if sqlBuilder.Filter != "" {
					sqlBuilder.Filter = fmt.Sprintf("%s,%s", sqlBuilder.Filter, "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7")
				} else {
					sqlBuilder.Filter = "status:<>0,status:<>2,status:<>3,status:<>4,status:<>5,status:<>7"
				}

			} else if req.GetTask().GetStep() == pb.Steps_Maker {

				if sqlBuilder.Filter != "" {
					sqlBuilder.Filter = fmt.Sprintf("%s,%s", sqlBuilder.Filter, "status:<>0,status:<>1,status:<>4,status:<>5,status:<>6,status:<>7")
				} else {
					sqlBuilder.Filter = "status:<>0,status:<>1,status:<>4,status:<>5,status:<>6,status:<>7"
				}

			}

		}

	}

	result.Pagination = setPagination(req)

	list, err := s.provider.GetListTaskNormal(ctx, dataORM, result.Pagination, sqlBuilder, companyID, userID, roleIDs, accountIDs)
	if err != nil {
		logrus.Errorln("[api][func: GetListTask] Failed when execute GetListTaskNormal:", err)
		return nil, err
	}

	for _, v := range list {

		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln("[api][func: GetListTask] Failed convert ORM to PB:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		result.Data = append(result.Data, &task)

	}

	return &result, err

}

func (s *Server) GetListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskResponse, error) {

	result := pb.ListTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task List",
		Data:    []*pb.Task{},
	}

	var dataorm pb.TaskORM
	if req.Task != nil {
		dataorm, _ = req.Task.ToORM(ctx)
	}

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
		FilterNot:     req.GetFilterNot(),
	}

	result.Pagination = setPagination(req)

	list, err := s.provider.GetListTask(ctx, &dataorm, result.Pagination, sqlBuilder, req.UserIDFilter, req.RoleIDFilter, req.AccountIDFilter)
	if err != nil {
		return nil, err
	}

	for _, v := range list {

		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln("[api][func: GetListTask] Failed convert ORM to PB:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		result.Data = append(result.Data, &task)

	}

	return &result, err

}

func (s *Server) GetListTaskPluck(ctx context.Context, req *pb.ListTaskPluckRequest) (*pb.ListTaskPluckResponse, error) {

	result := &pb.ListTaskPluckResponse{
		Data: []string{},
	}

	if req.PluckKey == "" {
		return result, nil
	}

	var dataORM pb.TaskORM
	var err error

	if req.Task != nil {

		dataORM, err = req.Task.ToORM(ctx)
		if err != nil {
			logrus.Errorln("[api][func: GetListTaskPluck] Failed when convert PB to ORM:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

	}

	sqlBuilder := &db.QueryBuilder{
		Filter:        req.GetFilter(),
		FilterOr:      req.GetFilterOr(),
		CollectiveAnd: req.GetQuery(),
		In:            req.GetIn(),
		Distinct:      req.GetDistinctKey(),
	}

	list, err := s.provider.GetListTaskPluck(ctx, req.GetPluckKey(), &dataORM, sqlBuilder)
	if err != nil {
		logrus.Errorln("[api][func: GetListTaskPluck] Failed when execute GetListTaskPluck:", err)
		return nil, err
	}

	result.Data = list

	return result, nil

}

func (s *Server) GetTaskGraphStatus(ctx context.Context, req *pb.GraphStatusRequest) (*pb.GraphStatusResponse, error) {

	res := &pb.GraphStatusResponse{
		Error:   false,
		Code:    200,
		Message: "Graph Data",
	}

	stat := req.Status.Number()

	data, err := s.provider.GetGraphStatus(ctx, req.Service, uint(stat))
	if err != nil {
		logrus.Errorln("[api][func: GetTaskGraphStatus] Failed when execute GetGraphStatus:", err)
		return nil, err
	}

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

	res := &pb.GraphStatusColumnTypeResponse{
		Error:   false,
		Code:    200,
		Message: "Graph Data",
	}

	stat := req.Status.Number()

	data, err := s.provider.GetGraphServiceType(ctx, req.Service, uint(stat), req.Column)
	if err != nil {
		logrus.Errorln("[api][func: GraphStatusColumnType] Failed when execute GetGraphServiceType:", err)
		return nil, err
	}

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

	res := &pb.GraphStepResponse{
		Error:   false,
		Code:    200,
		Message: "Graph Data",
	}

	me, _, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		logrus.Errorln("[api][func: GetTaskGraphStep] Failed when execute GetMeFromJWT:", err)
		return nil, err
	}

	step := req.Step.Number()
	stat := req.Status.Number()

	if me.UserType == "ba" {
		me.CompanyID = 0
	}

	switch req.GetStep() {
	case pb.Steps_Checker:
		stat = 1
	case pb.Steps_Signer:
		stat = 1
	case pb.Steps_Releaser:
		stat = 1
	}

	data, err := s.provider.GetGraphStep(ctx, me.CompanyID, req.Service, uint(step), uint(stat), req.IsIncludeApprove, req.IsIncludeReject, me.UserType)
	if err != nil {
		logrus.Errorln("[api][func: GetTaskGraphStep] Failed when execute GetGraphStep:", err)
		return nil, err
	}

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

func (s *Server) GetMyPendingTaskWithWorkflowGraph(ctx context.Context, req *pb.GetMyPendingTaskWithWorkflowGraphRequest) (*pb.GetMyPendingTaskWithWorkflowGraphResponse, error) {

	res := &pb.GetMyPendingTaskWithWorkflowGraphResponse{
		Code:    200,
		Error:   false,
		Message: "Graph Data",
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		logrus.Errorln("[api][func: GetMyPendingTaskWithWorkflowGraph] Failed when execute GetMeFromMD:", err)
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var trailer metadata.MD

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	accountConn, err := grpc.Dial(getEnv("ACCOUNT_SERVICE", ":9090"), opts...)
	if err != nil {
		logrus.Errorln("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to connect Account Service:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	defer accountConn.Close()

	accountClient := account_pb.NewApiServiceClient(accountConn)

	listAccountByRoleReq := &account_pb.ListAccountRequest{}

	if req.GetService() != "" {

		productConn, err := grpc.Dial(getEnv("PRODUCT_SERVICE", ":9090"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to connect Product Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer productConn.Close()

		productClient := product_pb.NewApiServiceClient(productConn)

		listProductRes, err := productClient.ListProduct(ctx, &product_pb.ListProductRequest{
			Product: &product_pb.Product{
				Name: req.GetService(),
			},
		})
		if err != nil {
			logrus.Println("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to Get List Product:", err.Error())
			return nil, err
		}

		if len(listProductRes.GetData()) < 1 {
			return nil, status.Errorf(codes.NotFound, "Product Not Found")
		}

		listAccountByRoleReq = &account_pb.ListAccountRequest{
			ProductID: listProductRes.GetData()[0].GetProductID(),
		}

	}

	listAccountRes, err := accountClient.ListAccountByRole(ctx, listAccountByRoleReq, grpc.Header(&userMD), grpc.Trailer(&trailer))
	if err != nil {
		logrus.Println("[api][func: GetMyPendingTaskWithWorkflowGraph] Unable to Get Account By Role:", err.Error())
		return nil, err
	}

	accountIDs := []uint64{}
	for _, v := range listAccountRes.Data {
		accountIDs = append(accountIDs, v.AccountID)
	}

	data, err := s.provider.GetGraphPendingTaskWithWorkflow(ctx, req.Service, currentUser.RoleIDs, accountIDs, currentUser.UserID, currentUser.CompanyID)
	if err != nil {
		logrus.Errorln("[api][func: GetMyPendingTaskWithWorkflowGraph] Failed when execute GetGraphPendingTaskWithWorkflow:", err)
		return nil, err
	}

	for _, v := range data {

		switch v.Name {
		case "verifier":
			v.Name = "Checker"
		case "checker":
			v.Name = "Checker"
		case "approver":
			v.Name = "Signer"
		case "signer":
			v.Name = "Signer"
		case "releaser":
			v.Name = "Releaser"
		case "maker":
			v.Name = "Maker"
		}

		val := &pb.GraphStepWorkflow{
			Step:  v.Name,
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
		logrus.Errorln("[api][func: GetListAnnouncement] Failed when execute GetListTaskWithFilter:", err)
		return nil, err
	}

	for _, v := range list {

		task, err := v.ToPB(ctx)
		if err != nil {
			logrus.Errorln("[api][func: GetListAnnouncement] Failed when convert ORM to PB:", err)
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
		logrus.Errorln("[api][func: SaveTaskWithDataEV] Failed when execute Decrypt:", err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
	}

	taskID, err := strconv.Atoi(text)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithDataEV] Failed when execute Atoi:", err)
		return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
	}

	taskPB, err := taskEVtoPB(req.Task, aes)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithDataEV] Failed when execute taskEVtoPB:", err)
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
		logrus.Errorln("[api][func: SaveTaskWithDataEV] Failed when execute SaveTaskWithData:", err)
		return nil, err
	}

	taskEV, err := taskPBtoEV(response.Data, aes)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithDataEV] Failed when execute taskPBtoEV:", err)
		return nil, err
	}

	res := &pb.SaveTaskResponseEV{
		Success: response.Success,
		Data:    taskEV,
	}

	return res, nil

}

func (s *Server) SaveTaskWithData(ctx context.Context, req *pb.SaveTaskRequest) (*pb.SaveTaskResponse, error) {

	logrus.Printf("first req =======> %v", req)
	logrus.Println("[api][func: SaveTaskWithData] Task Type:", req.Task.Type)

	var err error

	task, err := req.Task.ToORM(ctx)
	logrus.Printf("req after orm =======> %v", task)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Unable to convert PB to ORM:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	} else {
		ctx = metadata.NewOutgoingContext(context.Background(), userMD)
	}
	var trailer metadata.MD

	task.Step = 3
	task.Status = 1
	task.LastApprovedByID = 0
	task.LastApprovedByName = ""
	task.LastRejectedByID = 0
	task.LastRejectedByName = ""
	task.DataBak = "{}"

	logrus.Println("[api][func: SaveTaskWithData] Step 2")

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

	companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Unable to connect Company Service:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	defer companyConn.Close()

	companyClient := company_pb.NewApiServiceClient(companyConn)

	if task.Type == "Menu:License" {

		// if strings.Contains(task.Data, `"isParent": true`) {

		// 	for i := range task.Childs {

		// 		menu := menu_pb.MenuLicenseSave{}

		// 		err = json.Unmarshal([]byte(task.Childs[i].Data), &menu)
		// 		if err != nil {
		// 			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
		// 			return nil, status.Errorf(codes.Internal, "Internal Error")
		// 		}

		// 		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
		// 			CompanyID: menu.CompanyID,
		// 		})
		// 		if err != nil {
		// 			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
		// 			return nil, status.Errorf(codes.Internal, "Internal Error")
		// 		}

		// 		if len(company.Data) == 0 {
		// 			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
		// 			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		// 		}

		// 	}

		// } else {

		menu := menu_pb.MenuLicenseSave{}

		json.Unmarshal([]byte(task.Data), &menu)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: menu.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

		// }
	}

	if task.Type == "Account" {

		if strings.Contains(task.Data, `"isParent": true`) {

			for i := range task.Childs {

				account := account_pb.Account{}
				err = json.Unmarshal([]byte(task.Childs[i].Data), &account)
				if err != nil {
					logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: account.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

		} else {

			account := account_pb.Account{}

			json.Unmarshal([]byte(task.Data), &account)
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
				CompanyID: account.CompanyID,
			})
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			if len(company.Data) == 0 {
				logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
				return nil, status.Errorf(codes.NotFound, "Company does not exist")
			}

		}

	}

	if task.Type == "Subscription" {

		abonnement := abonnement_pb.ListTaskAbonnementRes{}

		err = json.Unmarshal([]byte(task.Data), &abonnement)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: abonnement.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "Beneficiary Account" {

		beneficiaryAccount := beneficiary_account_pb.BeneficiaryAccount{}

		err = json.Unmarshal([]byte(task.Data), &beneficiaryAccount)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: beneficiaryAccount.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "Role" {

		role := role_pb.Role{}

		err = json.Unmarshal([]byte(task.Data), &role)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: role.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "Workflow" {

		workflow := workflow_pb.WorkflowTask{}
		err = json.Unmarshal([]byte(task.Data), &workflow)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		companyID, _ := strconv.ParseUint(workflow.Company.CompanyID, 10, 64)

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: companyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "User" {

		users := users_pb.UserTaskData{}

		err = json.Unmarshal([]byte(task.Data), &users)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: users.User.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "Notification" {

		notification := notification_pb.NotificationCompany{}

		err = json.Unmarshal([]byte(task.Data), &notification)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: notification.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	if task.Type == "Company" {

		company := company_pb.CreateCompanyTaskReq{}

		err = json.Unmarshal([]byte(task.Data), &company)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		companyBRICaMS, err := companyClient.BRICaMSgetCustomerByIDV2(ctx, &company_pb.BricamsGetCustomerByIdReq{
			Id: fmt.Sprintf("%v", company.Company.CompanyID),
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute BRICaMSgetCustomerByIDV2:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if companyBRICaMS.ResponseCode == "0002" {
			logrus.Errorln("[api][func: SaveTaskWithData] Company Data Not Found From BRICaMS")
			return nil, status.Errorf(codes.NotFound, "Company Data Not Found From BRICaMS")
		}

	}

	if task.Type == "Announcement" {

		announcement := announcement_pb.Announcement{}

		err = json.Unmarshal([]byte(task.Data), &announcement)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable Unmarshal Data:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
			CompanyID: announcement.CompanyID,
		})
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListCompanyData:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		if len(company.Data) == 0 {
			logrus.Errorln("[api][func: SaveTaskWithData] Company does not exist")
			return nil, status.Errorf(codes.NotFound, "Company does not exist")
		}

	}

	productConn, err := grpc.Dial(getEnv("PRODUCT_SERVICE", ":9097"), opts...)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Unable to connect Product Service:", err)
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
		logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute ListProduct:", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	logrus.Println("[api][func: SaveTaskWithData] Step 3")

	if len(productData.Data) < 1 {
		return nil, status.Errorf(codes.NotFound, "This task type product, not found")
	} else {
		if productData.Data[0].Name != task.Type {
			return nil, status.Errorf(codes.NotFound, "This task type product, not found")
		}
	}

	product := productData.Data[0]

	taskType := []string{"Swift", "Cash Pooling", "BG Issuing", "Import LC", "Internal Fund Transfer", "BI-Fast", "Payroll Transfer",
		"Amend Cancel LC", "Deposito", "Online Transfer"}

	if product.IsTransactional && contains(taskType, task.Type) && !req.IsDraft { //skip for difference variable name, revisit later

		logrus.Println("[api][func: SaveTaskWithData] Step 4")

		if req.TransactionAmount == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Transaction amount is required")
		}

		workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Unable to connect Workflow Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer workflowConn.Close()

		workflowClient := workflow_pb.NewApiServiceClient(workflowConn)

		// Implement Workflow STP Here -- Start

		companyWorkflow, err := workflowClient.GetCompanyWorkflow(ctx, &workflow_pb.GetCompanyWorkflowRequest{
			CompanyID: currentUser.CompanyID,
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute GetCompanyWorkflow", err)
			return nil, err
		}

		if !companyWorkflow.Data.IsTransactionSTP {

			getWorkflow, err := workflowClient.GenerateWorkflow(ctx, &workflow_pb.GenerateWorkflowRequest{
				ProductID:           product.ProductID,
				CompanyID:           currentUser.CompanyID,
				TransactionalNumber: uint64(req.TransactionAmount),
				Currency:            req.GetTransactionCurrency(),
				SelectedAccountID:   req.GetSelectedAccountID(),
			}, grpc.Header(&userMD), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute GenerateWorkflow", err)
				return nil, err
			}

			if getWorkflow.Data == nil {
				return nil, status.Errorf(codes.NotFound, "workflow for this task type not found")
			}

			workflow, err := json.Marshal(getWorkflow.Data)
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Unable to Marshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			task.WorkflowDoc = string(workflow)

		} else {

			// Auto approve task if Workflow STP

			task.Status = 4

		}

		// Implement Workflow STP Here -- End

		logrus.Println("[api][func: SaveTaskWithData] Step 5")

	}

	if req.TaskID > 0 {

		findTask, err := s.provider.FindTaskById(ctx, req.TaskID)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute FindTaskById:", err)
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

	logrus.Println("[api][func: SaveTaskWithData] Step 6")

	if req.IsDraft {
		task.Step = 1
		task.Status = 2
	}

	command := "Create"

	logrus.Println("[api][func: SaveTaskWithData] Task Type:", task.Type)
	logrus.Println("[api][func: SaveTaskWithData] Task ID:", task.TaskID)

	var savedTask *pb.TaskORM

	if req.TaskID > 0 {

		command = "Update"
		task.TaskID = req.TaskID
		task.UpdatedByID = currentUser.UserID
		task.UpdatedByName = currentUser.Username

		savedTask, err = s.provider.UpdateTask(ctx, &task, true)

	} else {

		if req.GetTask().GetCreatedByID() > 0 {
			logrus.Println("[api][func: SaveTaskWithData] Created by IDs:", req.GetTask().GetCreatedByID())

			userConn, err := grpc.Dial(getEnv("USER_SERVICE", ":9095"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Unable to connect User Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer userConn.Close()

			userClient := users_pb.NewApiServiceClient(userConn)

			userRes, err := userClient.GetUserByUserID(ctx, &users_pb.GetUserIDArrayRequest{
				UserID: []uint64{req.GetTask().GetCreatedByID()},
			})
			if err != nil {
				logrus.Errorln("[api][func: SaveTaskWithData] Unable to Get User By User ID:", err)
				return nil, err
			}

			if len(userRes.GetData()) > 0 {

				task.CreatedByID = userRes.GetData()[0].UserID
				task.CreatedByName = userRes.GetData()[0].Username
				task.UpdatedByID = userRes.GetData()[0].UserID
				task.UpdatedByName = userRes.GetData()[0].Username

			} else {

				logrus.Errorln("[api][func: SaveTaskWithData] User not found")
				return nil, status.Errorf(codes.NotFound, "User not found")

			}

		} else {

			task.CreatedByID = currentUser.UserID
			task.CreatedByName = currentUser.Username
			task.UpdatedByID = currentUser.UserID
			task.UpdatedByName = currentUser.Username

		}

		savedTask, err = s.provider.CreateTask(ctx, &task)

	}

	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute UpdateTask or CreateTask:", err)
		return nil, err
	}

	saved, err := savedTask.ToPB(ctx)
	if err != nil {
		logrus.Errorln("[api][func: SaveTaskWithData] Unable to convert ORM to PB:", err)
	}

	res := &pb.SaveTaskResponse{
		Success: true,
		Data: &pb.Task{
			TaskID:        saved.TaskID,
			Data:          saved.Data,
			WorkflowDoc:   saved.WorkflowDoc,
			CompanyID:     saved.CompanyID,
			CreatedByID:   saved.CreatedByID,
			CreatedByName: saved.CreatedByName,
			CreatedAt:     saved.CreatedAt,
			UpdatedAt:     saved.UpdatedAt,
		},
	}

	logrus.Println("[api][func: SaveTaskWithData] Step 7")

	logrus.Println("[api][func: SaveTaskWithData] Save Log for Task Type:", task.Type)

	if getEnv("ENV", "LOCAL") != "LOCAL" {
		action := "save"
		if req.IsDraft {
			action = "draft"
		}

		logrus.Println("[api][func: SaveTaskWithData] Set Save Log for Task Type:", task.Type)

		activityLog, err := GenerateActivityLog(&task, currentUser, action, command)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed to generate activity log: %v", err)
			return nil, err
		}

		err = s.provider.SaveLog(ctx, activityLog)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed when execute SaveLog:", err)
		}

	}

	logrus.Println("[api][func: SaveTaskWithData] Step 8")

	go TaskNotificationCreateOrUpdate(ctx, &task, command, task.Step, task.Status)

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
	skipProduct := []string{"BG Mapping", "BG Mapping Digital", "Internal Fund Transfer", "BI-Fast", "Payroll Transfer", "Online Transfer"}

	logrus.Print(taskType)

	for _, v := range skipProduct {
		if v == taskType {
			return true
		}
	}

	logrus.Println("check md Task Type:", taskType)

	productName := strings.Replace(taskType, ":", "_", -1)
	productName = strings.Replace(productName, " ", "_", -1)
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
	logrus.Println("MD +==============>", md)
	logrus.Println("Product +===============>", productName)
	if len(md[productName]) > 0 {
		result := strings.Split(md[productName][0], ",")
		logrus.Print("result md %s", result)
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

func (s *Server) SetTaskWithWorkflow(ctx context.Context, req *pb.SetTaskWithWorkflowRequest) (*pb.SetTaskResponse, error) {

	result := &pb.SetTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	var err error

	_, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	var newCtx context.Context

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		newCtx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var trailer metadata.MD

	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	taskPB, err := task.ToPB(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	switch task.Type {
	case "Internal Fund Transfer":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		transferConn, err := grpc.Dial(getEnv("TRANSFER_SERVICE", ":9125"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Transfer Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer transferConn.Close()

		transferClient := transfer_pb.NewApiServiceClient(transferConn)

		_, err = transferClient.SetTaskInternalTransfer(newCtx, &transfer_pb.SetTaskInternalTransferRequest{
			TaskID:   req.GetTaskID(),
			Action:   req.GetAction(),
			Comment:  req.GetComment(),
			Reasons:  req.GetReasons(),
			PassCode: req.GetPassCode(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "BI-Fast":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		bifastConn, err := grpc.Dial(getEnv("BIFAST_SERVICE", ":9127"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect BiFast Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer bifastConn.Close()

		bifastClient := bifast_pb.NewApiServiceClient(bifastConn)

		_, err = bifastClient.SetTaskExternalTransfer(newCtx, &bifast_pb.SetTaskExternalTransferRequest{
			TaskID:  req.GetTaskID(),
			Action:  req.GetAction(),
			Comment: req.GetComment(),
			Reasons: req.GetReasons(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "Beneficiary Account":
		logrus.Println("TEST masuk ga????")
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		benefConn, err := grpc.Dial(getEnv("BENEFICIARY_ACCOUNT_SERVICE", ":9107"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Beneficiary Account Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer benefConn.Close()

		benefClient := beneficiary_account_pb.NewApiServiceClient(benefConn)

		_, err = benefClient.TaskAction(newCtx, &beneficiary_account_pb.TaskActionRequest{
			TaskID:  req.GetTaskID(),
			Action:  req.GetAction(),
			Comment: req.GetComment(),
			Reasons: req.GetReasons(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "Payroll Transfer":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		payrollConn, err := grpc.Dial(getEnv("PAYROLL_SERVICE", ":9126"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Payroll Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer payrollConn.Close()

		payrollClient := payroll_pb.NewApiServiceClient(payrollConn)

		if req.GetAction() == "submit" {

			taskData := &payroll_pb.PayrollDataJob{}
			err = json.Unmarshal([]byte(taskPB.Data), taskData)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err.Error())
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			_, err = payrollClient.CreateTaskPayroll(newCtx, &payroll_pb.CreateTaskPayrollRequest{
				TaskID:              req.GetTaskID(),
				FileName:            taskData.GetUploadFileName(),
				FileDescription:     taskData.GetFileDescription(),
				TransactionType:     taskData.GetTransactionType(),
				TransactionSchedule: taskData.GetTransactionSchedule(),
				ScheduledAt:         taskData.GetScheduledAt(),
				IsDraft:             false,
			})
			if err != nil {
				return nil, err
			}

		} else {

			_, err = payrollClient.SetTaskPayroll(newCtx, &payroll_pb.SetTaskPayrollRequest{
				TaskID:  req.GetTaskID(),
				Action:  req.GetAction(),
				Comment: req.GetComment(),
				Reasons: req.GetReasons(),
			}, grpc.Header(&userMD), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}

		}

	case "Swift":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		swiftConn, err := grpc.Dial(getEnv("SWIFT_SERVICE", ":9211"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Swift Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer swiftConn.Close()

		swiftClient := swift_pb.NewSwiftServiceClient(swiftConn)

		_, err = swiftClient.TaskAction(newCtx, &swift_pb.TaskActionRequest{
			TaskID:   req.GetTaskID(),
			Action:   req.GetAction(),
			Comment:  req.GetComment(),
			Reasons:  req.GetReasons(),
			PassCode: req.GetPassCode(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "BG Issuing":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Transfer Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer bgConn.Close()

		bgClient := bg_pb.NewApiServiceClient(bgConn)

		_, err = bgClient.TaskIssuingAction(newCtx, &bg_pb.TaskIssuingActionRequest{
			TaskID:   req.GetTaskID(),
			Action:   req.GetAction(),
			Comment:  req.GetComment(),
			Reasons:  req.GetReasons(),
			PassCode: req.GetPassCode(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "Deposito":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		depositoConn, err := grpc.Dial(getEnv("DEPOSITO_SERVICE", ":9202"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Transfer Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer depositoConn.Close()

		depositoClient := deposito_pb.NewDepositoServiceClient(depositoConn)

		depositoClient.DepositoActionTask(newCtx, &deposito_pb.TaskActionRequest{
			TaskID:   req.GetTaskID(),
			Action:   req.GetAction(),
			Comment:  req.GetComment(),
			Reasons:  req.GetReasons(),
			PassCode: req.GetPassCode(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	case "Online Transfer":

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())

		onlineTransferConn, err := grpc.Dial(getEnv("ONLINE_TRANSFER_SERVICE", ":9128"), opts...)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Unable to connect Online Transfer Service:", err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
		defer onlineTransferConn.Close()

		onlineTransferClient := online_transfer_pb.NewApiServiceClient(onlineTransferConn)

		_, err = onlineTransferClient.SetTaskOnlineTransfer(newCtx, &online_transfer_pb.SetTaskOnlineTransferRequest{
			TaskID:  req.GetTaskID(),
			Action:  req.GetAction(),
			Comment: req.GetComment(),
			Reasons: req.GetReasons(),
		}, grpc.Header(&userMD), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

	default:

		_, err = s.SetTask(ctx, &pb.SetTaskRequest{
			TaskID:  req.GetTaskID(),
			Action:  req.GetAction(),
			Comment: req.GetComment(),
			Reasons: req.GetReasons(),
		})
		if err != nil {
			return nil, err
		}

	}

	result.Data = &taskPB

	return result, nil

}

func (s *Server) SetTask(ctx context.Context, req *pb.SetTaskRequest) (*pb.SetTaskResponse, error) {

	var err error

	currentUser, userMd, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	originalCtx := ctx

	re := regexp.MustCompile(`(<[a-z,\/]+.*?>)`)

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var header, trailer metadata.MD

	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	if task.Type != "System" && task.Type != "Menu:License" {

		if strings.ToLower(req.Action) == "delete" {

			allowed := checkAllowedApproval(userMd, task.Type, "data_entry:maker")
			if !allowed {
				logrus.Errorln("[api][func: SetTask] Permission Denied")
				return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
			}

		} else {

			allowed := checkAllowedApproval(userMd, task.Type, "approve:signer")
			if !allowed {
				logrus.Errorln("[api][func: SetTask] Permission Denied")
				return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
			}

		}

	}

	if strings.ToLower(req.Action) == "approve" {

		allowed := checkAllowedApproval(userMd, task.Type, "approve:signer")
		if !allowed {
			logrus.Errorln("[api][func: SetTask] Permission Denied")
			return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
		}

		if currentUser.UserType != "ba" {

			if currentUser.CompanyID != task.CompanyID {
				logrus.Errorln("[api][func: SetTask] Permission Denied")
				return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
			}

		}

	}

	if task.IsParentActive {
		logrus.Errorln("[api][func: SetTask] This is child task with active parent, please refer to parent for change status")
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
	currentData := task.Data
	currentDataBak := task.DataBak

	if currentStep == 3 { // If Current Step is Signer

		if currentStatus == 1 || currentStatus == 6 { // If Current Status is Pending Or DeleteRequest

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Company Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer companyConn.Close()

			companyClient := company_pb.NewApiServiceClient(companyConn)

			if task.Type == "Menu:License" {
				logrus.Printf("child menu =======> %v", task.Childs)

				if strings.Contains(task.Data, `"isParent": true`) {

					for i := range task.Childs {
						menu := menu_pb.MenuLicenseSave{}
						json.Unmarshal([]byte(task.Childs[i].Data), &menu)

						company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
							CompanyID: menu.CompanyID,
						})
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						if len(company.Data) == 0 {
							logrus.Errorln("[api][func: SetTask] Company does not exist")
							return nil, status.Errorf(codes.NotFound, "Company does not exist")
						}

					}

				} else {

					menu := menu_pb.MenuLicenseSave{}
					err = json.Unmarshal([]byte(task.Data), &menu)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
						CompanyID: menu.CompanyID,
					})
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					if len(company.Data) == 0 {
						logrus.Errorln("[api][func: SetTask] Company does not exist")
						return nil, status.Errorf(codes.NotFound, "Company does not exist")
					}

				}

			}

			if task.Type == "Account" {

				if strings.Contains(task.Data, `"isParent": true`) {

					for i := range task.Childs {

						account := account_pb.Account{}
						err = json.Unmarshal([]byte(task.Childs[i].Data), &account)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						if currentUser.UserType != "ba" && currentUser.CompanyID != account.CompanyID {
							logrus.Errorln("[api][func: SetTask] Permission Denied")
							return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
						}

						company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
							CompanyID: account.CompanyID,
						})
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						if len(company.Data) == 0 {
							logrus.Errorln("[api][func: SetTask] Company does not exist")
							return nil, status.Errorf(codes.NotFound, "Company does not exist")
						}

					}

				} else {

					account := account_pb.Account{}
					err = json.Unmarshal([]byte(task.Data), &account)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					if currentUser.UserType != "ba" && currentUser.CompanyID != account.CompanyID {
						logrus.Errorln("[api][func: SetTask] Permission Denied")
						return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
					}

					company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
						CompanyID: account.CompanyID,
					})
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					if len(company.Data) == 0 {
						logrus.Errorln("[api][func: SetTask] Company does not exist")
						return nil, status.Errorf(codes.NotFound, "Company does not exist")
					}

				}

			}

			if task.Type == "Subscription" {

				abonnement := abonnement_pb.ListTaskAbonnementRes{}
				err = json.Unmarshal([]byte(task.Data), &abonnement)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != abonnement.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: abonnement.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "Beneficiary Account" {

				beneficiaryAccount := beneficiary_account_pb.BeneficiaryAccount{}
				err = json.Unmarshal([]byte(task.Data), &beneficiaryAccount)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != beneficiaryAccount.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: beneficiaryAccount.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "Role" {

				role := role_pb.Role{}
				err = json.Unmarshal([]byte(task.Data), &role)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != role.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: role.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "Workflow" {

				workflow := workflow_pb.WorkflowTask{}
				json.Unmarshal([]byte(task.Data), &workflow)
				companyID, _ := strconv.ParseUint(workflow.Company.CompanyID, 10, 64)

				if currentUser.UserType != "ba" && currentUser.CompanyID != companyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: companyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "User" {

				users := users_pb.UserTaskData{}
				err = json.Unmarshal([]byte(task.Data), &users)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != users.User.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: users.User.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "Notification" {

				notification := notification_pb.NotificationCompany{}
				err = json.Unmarshal([]byte(task.Data), &notification)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != notification.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: notification.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

			if task.Type == "Company" {

				company := company_pb.CreateCompanyTaskReq{}
				err = json.Unmarshal([]byte(task.Data), &company)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != company.Company.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				companyBRICaMS, err := companyClient.BRICaMSgetCustomerByIDV2(ctx, &company_pb.BricamsGetCustomerByIdReq{
					Id: fmt.Sprintf("%v", company.Company.CompanyID),
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when BRICaMSgetCustomerByIDV2:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if companyBRICaMS.ResponseCode == "0002" {
					logrus.Errorln("[api][func: SetTask] Company Data Not Found From BRICaMS")
					return nil, status.Errorf(codes.NotFound, "Company Data Not Found From BRICaMS")
				}

			}

			if task.Type == "Announcement" {

				announcement := announcement_pb.Announcement{}
				err = json.Unmarshal([]byte(task.Data), &announcement)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Umarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if currentUser.UserType != "ba" && currentUser.CompanyID != announcement.CompanyID {
					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
				}

				company, err := companyClient.ListCompanyData(ctx, &company_pb.ListCompanyDataReq{
					CompanyID: announcement.CompanyID,
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when ListCompanyData:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				if len(company.Data) == 0 {
					logrus.Errorln("[api][func: SetTask] Company does not exist")
					return nil, status.Errorf(codes.NotFound, "Company does not exist")
				}

			}

		}

	}

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

			if re.MatchString(req.Comment) {

				logrus.Errorln("[api][func: SetTask] Invalid Reword Comment Characters:", req.Comment)
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")

			} else {

				task.Comment = req.Comment

			}

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

		// if task.Type == "BG Issuing" {

		// 	var opts []grpc.DialOption
		// 	opts = append(opts, grpc.WithInsecure())

		// 	bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
		// 	if err != nil {
		// 		logrus.Errorln("[api][func: SetTask] Unable to connect BG Service:", err)
		// 		return nil, status.Errorf(codes.Internal, "Internal Error")
		// 	}
		// 	defer bgConn.Close()

		// 	bgClient := bg_pb.NewApiServiceClient(bgConn)

		// 	taskData := bg_pb.IssuingData{}
		// 	err = json.Unmarshal([]byte(currentData), &taskData)
		// 	if err != nil {
		// 		logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
		// 		return nil, status.Errorf(codes.Internal, "Internal Error", err)
		// 	}

		// 	bgGrpcReq := &bg_pb.CreateIssuingRequest{
		// 		TaskID: task.TaskID,
		// 		Data:   &taskData,
		// 	}

		// 	result, err := bgClient.CreateIssuing(ctx, bgGrpcReq, grpc.Header(&header), grpc.Trailer(&trailer))
		// 	if err != nil {
		// 		logrus.Errorln("[api][func: SetTask] Failed when CreateIssuing:", err)
		// 		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		// 	}

		// 	taskData.RegistrationNo = result.Data.GetRegistrationNo()
		// 	taskData.ReferenceNo = result.Data.GetReferenceNo()

		// 	data, _ := json.Marshal(&taskData)
		// 	task.Data = string(data)

		// }

		if currentStep == 2 {

			currentStep = 3

		}

		if currentStatus == 2 {

			task.Step = 3
			task.Status = 1

		} else {

			if currentStep == 1 {

				task.Status = 1
				task.Step = 3

				if currentStatus == 6 {

					task.Status = currentStatus

				}

			}

			if currentStep == 3 {

				if currentStatus == 1 {

					if task.Type == "Company" {

						company := &company_pb.CreateCompanyTaskReq{}
						json.Unmarshal([]byte(task.Data), &company)

						companyWorkflow := company.Workflow.TansactionalStpStep
						assignedStep := []string{}
						if companyWorkflow.Approver {
							assignedStep = append(assignedStep, "signer")
						}
						if companyWorkflow.Releaser {
							assignedStep = append(assignedStep, "releaser")
						}
						if companyWorkflow.Verifier {
							assignedStep = append(assignedStep, "checker")
						}

						var opts []grpc.DialOption
						opts = append(opts, grpc.WithInsecure())

						workFlowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Unable to connect Workflow Service:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}
						defer workFlowConn.Close()

						workflowClient := workflow_pb.NewApiServiceClient(workFlowConn)
						resWorkflow, err := workflowClient.ListWorkflow(ctx, &workflow_pb.ListWorkflowRequest{
							Workflow: &workflow_pb.Workflow{
								CompanyID: company.Company.CompanyID,
							},
						}, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							return nil, err
						}
						deletedID := []uint64{}
						if len(resWorkflow.Data) > 0 {
							for _, w := range resWorkflow.Data {
								for _, l := range w.Logics {
									for _, r := range l.Requirements {
										deleted := true
										for _, s := range assignedStep {
											if r.Step == s {
												deleted = false
											}
										}
										if deleted {
											deletedID = append(deletedID, r.WorkflowRequirementID)
										}
									}
								}
							}
						}
						if len(deletedID) > 0 {
							workflowClient.DeleteRequirement(ctx, &workflow_pb.DeleteRequirementRequest{
								RequirementID: deletedID,
								CompanyID:     company.Company.CompanyID,
								AssignedSteps: assignedStep,
							}, grpc.Header(&header), grpc.Trailer(&trailer))
						}

					}

				}

				task.Status = 1
				task.Status = 4
				task.Step = 3
				sendTask = true

				if currentStatus == 6 {

					if task.Type == "Company" {

						err := s.DeleteCompany(originalCtx, task.Data)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Failed when DeleteCompany:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
						}

					}

					task.Status = 7

				}

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
						logrus.Errorln("[api][func: SetTask] Unable to connect Workflow Service:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}
					defer workflowConn.Close()

					workflowClient := workflow_pb.NewApiServiceClient(workflowConn)

					company := company_pb.CreateCompanyReq{}
					err = json.Unmarshal([]byte(task.Data), &company)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					data := workflow_pb.CreateCompanyWorkflowRequest{
						TaskID: task.TaskID,
						Data: &workflow_pb.CompanyWorkflows{
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
						},
					}

					_, err = workflowClient.CreateCompanyWorkflow(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when CreateCompanyWorkflow:", err)
						return nil, err
					}

				}

			}

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

			if re.MatchString(req.Comment) {

				logrus.Errorln("[api][func: SetTask] Invalid Reject Comment Characters:", req.Comment)
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")

			} else {

				task.Comment = req.Comment

			}

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
			"Deposito", "Subscription", "Notification", "Announcement", "Liquidity", "Swift", "CBM",
			"Abonnement"}

		if contains(taskType, task.Type) {

			if task.DataBak != "" && task.DataBak != "{}" {

				task.Status = 4
				task.Step = 3

				if task.Type == "Subscription" {

					taskSubscription := abonnement_pb.CreateAbonnementTaskReq{}

					err = json.Unmarshal([]byte(task.Data), &taskSubscription)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					lastTransStat := taskSubscription.BillingStatus

					err = json.Unmarshal([]byte(task.DataBak), &taskSubscription)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					taskSubscription.BillingStatus = lastTransStat

					marsData, err := json.Marshal(&taskSubscription)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Marshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					task.Data = string(marsData)

				} else {

					task.Data = task.DataBak

				}

			}

		}

		if task.Type == "Menu:Appearance" || task.Type == "Menu:License" {

			if task.ChildBak != "" && task.ChildBak != "[]" {

				taskChilds := []*pb.TaskORM{}

				err := json.Unmarshal([]byte(task.ChildBak), &taskChilds)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
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

		if task.Type == "Company" {

			company := company_pb.CreateCompanyTaskReq{}
			err = json.Unmarshal([]byte(task.Data), &company)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data, err := s.provider.GetGraphStepAll(ctx, fmt.Sprint(company.Company.CompanyID))
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when GetGraphStepAll:", err)
				return nil, err
			}

			if data.Total > 0 {
				logrus.Errorln("[api][func: SetTask] Cannot delete company, some task on progress")
				return nil, status.Errorf(codes.Canceled, "Cannot delete company, some task on progress")
			}

		}

		if task.Type == "Role" {

			role := role_pb.Role{}

			err = json.Unmarshal([]byte(task.Data), &role)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			companyID := role.CompanyID

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			workFlowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Workflow Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer workFlowConn.Close()

			workflowClient := workflow_pb.NewApiServiceClient(workFlowConn)

			workFlow, err := workflowClient.ListWorkflow(ctx, &workflow_pb.ListWorkflowRequest{
				Filter: fmt.Sprintf("CompanyID:%d", companyID),
			}, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when ListWorkflow:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			used := []string{}
			for _, v := range workFlow.Data {

				for _, v2 := range v.Logics {

					for _, v3 := range v2.Requirements {

						if v3.RoleID == role.RoleID {

							used = append(used, v.WorkflowCode)

						}

					}

				}

			}

			if len(used) > 0 {
				return nil, status.Errorf(codes.Canceled, "Delete Role Failed: Role mapping to workflow [ %v ]", strings.Join(used, ", "))
			}

			userTask, err := s.GetListTask(ctx, &pb.ListTaskRequest{
				Filter: "type:User",
			})
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when GetListTask:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			used2 := []string{}
			for _, v := range userTask.Data {

				if strings.Contains(v.Data, fmt.Sprintf(`"roleID": %v`, role.RoleID)) {

					userDat := users_pb.UserTaskData{}

					err = json.Unmarshal([]byte(v.Data), &userDat)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					used2 = append(used2, userDat.User.Username)

				}

			}

			if len(used2) > 0 {
				return nil, status.Errorf(codes.Canceled, "Delete Role Failed: Role mapping to User [ %v ]", strings.Join(used2, ", "))
			}

		}

		if task.Type == "Account" {

			if strings.Contains(task.Data, `"isParent": true`) {

				for i := range task.Childs {

					account := account_pb.Account{}
					err = json.Unmarshal([]byte(task.Childs[i].Data), &account)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					roleTask, err := s.GetListTask(ctx, &pb.ListTaskRequest{
						Filter: "type:Role",
						In:     fmt.Sprintf("data.companyID:%v", account.CompanyID),
					})
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when GetListTask:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					used := []string{}
					for _, v := range roleTask.Data {

						if strings.Contains(v.Data, fmt.Sprintf(`"accountNumber": "%v"`, account.AccountNumber)) {

							role := role_pb.Role{}

							err = json.Unmarshal([]byte(v.Data), &role)
							if err != nil {
								logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
								return nil, status.Errorf(codes.Internal, "Internal Error")
							}

							used = append(used, role.Name)

						}

					}

					if len(used) > 0 {
						return nil, status.Errorf(codes.Canceled, "Delete Account Failed: Account mapping to role [ %v ]", strings.Join(used, ", "))
					}

				}

			} else {

				account := account_pb.Account{}

				err = json.Unmarshal([]byte(task.Data), &account)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				roleTask, err := s.GetListTask(ctx, &pb.ListTaskRequest{
					Filter: "type:Role",
					In:     fmt.Sprintf("data.companyID:%v", account.CompanyID),
				})
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when GetListTask:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				used := []string{}
				for _, v := range roleTask.Data {

					if strings.Contains(v.Data, fmt.Sprintf(`"accountNumber": "%v"`, account.AccountNumber)) {

						role := role_pb.Role{}

						err = json.Unmarshal([]byte(v.Data), &role)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						used = append(used, role.Name)

					}

				}

				if len(used) > 0 {
					return nil, status.Errorf(codes.Canceled, "Delete Account Failed: Account mapping to role [ %v ]", strings.Join(used, ", "))
				}

			}

		}

		if task.Type == "Beneficiary Account" {

			beneficiaryAccount := beneficiary_account_pb.BeneficiaryAccount{}

			err = json.Unmarshal([]byte(task.Data), &beneficiaryAccount)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			roleTask, err := s.GetListTask(ctx, &pb.ListTaskRequest{
				Filter: "type:Role",
				In:     fmt.Sprintf("data.companyID:%v", beneficiaryAccount.CompanyID),
			})
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when GetListTask:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			used := []string{}
			for _, v := range roleTask.Data {

				if strings.Contains(v.Data, fmt.Sprintf(`"accountNumber": "%v"`, beneficiaryAccount.AccountNumber)) {

					role := role_pb.Role{}

					err = json.Unmarshal([]byte(v.Data), &role)
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}

					used = append(used, role.Name)

				}

			}

			if len(used) > 0 {
				return nil, status.Errorf(codes.Canceled, "Delete Account Failed: Account mapping to role [ %v ]", strings.Join(used, ", "))
			}

		}

		if task.Type == "BG Mapping" {

			if currentUser.UserType != "ba" {

				if currentUser.CompanyID != task.CompanyID {

					logrus.Errorln("[api][func: SetTask] Permission Denied")
					return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")

				}

			}

			mappingDigitalTask, err := s.provider.GetListTask(ctx, &pb.TaskORM{
				Type:      "BG Mapping Digital",
				CompanyID: task.CompanyID,
			}, &pb.PaginationResponse{}, &db.QueryBuilder{Filter: "status:<>5,status:<>7"}, 0, []uint64{}, []uint64{})
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when GetListTask:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			used := []string{}
			for _, v := range mappingDigitalTask {

				used = append(used, strconv.FormatUint(v.TaskID, 10))

			}

			if len(mappingDigitalTask) > 0 {
				return nil, status.Errorf(codes.Canceled, "Bad Request: Delete BG Mapping Failed: BG Mapping used to Mapping Digital BG [ %v ]", strings.Join(used, ", "))
			}

		}

		task.LastApprovedByID = 0
		task.LastApprovedByName = ""
		task.LastRejectedByID = 0
		task.LastRejectedByName = ""

		task.Status = 6
		task.Step = 3

		if currentStatus == 2 {

			if !(task.DataBak == "" || task.DataBak == "{}") {

				task.Status = 4
				task.Step = 3
				task.Data = task.DataBak

				if contains([]string{"Internal Fund Transfer", "Payroll Transfer", "BI-Fast", "Online Transfer"}, task.Type) {

					task.Status = 7
					task.Step = 1

				}

			} else {

				task.Status = 7
				task.Step = 1

			}

		}

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

	if sendTask {

		if task.Data != "" && task.Data != "{}" {

			logrus.Println("[api][func: SetTask] Save Backup")
			task.DataBak = task.Data

		}

		if task.DataBak == "" {

			task.DataBak = "{}"

		}

		if (task.Type == "Menu:License" || task.Type == "Menu:Appearance") && len(task.Childs) > 0 {

			jsonChilds, err := json.Marshal(task.Childs)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Marshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			task.ChildBak = string(jsonChilds)

		} else {

			task.ChildBak = "[]"

		}

	}

	updatedTask, err := s.provider.UpdateTask(ctx, task, false)
	if err != nil {
		logrus.Errorln("[api][func: SetTask] Failed when UpdateTask:", err)
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

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			announcementConn, err := grpc.Dial(getEnv("ANNOUNCEMENT_SERVICE", ":9091"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Announcement Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer announcementConn.Close()

			announcementClient := announcement_pb.NewApiServiceClient(announcementConn)

			data := announcement_pb.Announcement{}

			err = json.Unmarshal([]byte(task.Data), &data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			send := &announcement_pb.CreateAnnouncementRequest{
				TaskID: task.TaskID,
				Data:   &data,
			}
			send.Data.AnnouncementID = task.FeatureID

			if task.Status == 7 {

				_, err = announcementClient.DeleteAnnouncement(ctx, send, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteAnnouncement:", err)
					return nil, err
				}

			} else {

				_, err = announcementClient.CreateAnnouncement(ctx, send, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateAnnouncement:", err)
					return nil, err
				}

			}

		case "Company":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			companyConn, err := grpc.Dial(getEnv("COMPANY_SERVICE", ":9092"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Company Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer companyConn.Close()

			companyClient := company_pb.NewApiServiceClient(companyConn)

			data := company_pb.CreateCompanyReq{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID

			if task.Status == 7 {

				_, err := companyClient.DeleteCompany(ctx, data.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteCompany:", err)
					return nil, err
				}

			} else {

				_, err := companyClient.CreateCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateCompany:", err)
					return nil, err
				}

			}

		case "Deposito":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			DepositoConn, err := grpc.Dial(getEnv("DEPOSITO_SERVICE", ":9201"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Deposito Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer DepositoConn.Close()

			depositoClient := deposito_pb.NewDepositoServiceClient(DepositoConn)

			data := &deposito_pb.CreateDepositoRequest{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

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

			_, err = depositoClient.CreateDeposito(ctx, data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when CreateDeposito:", err)
				return nil, err
			}

		case "Account":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			accountConn, err := grpc.Dial(getEnv("ACCOUNT_SERVICE", ":9093"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Account Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer accountConn.Close()

			accountClient := account_pb.NewApiServiceClient(accountConn)

			if isParent {

				for i := range task.Childs {

					if task.Childs[i].IsParentActive {

						data := account_pb.CreateAccountRequest{}
						account := account_pb.Account{}

						err = json.Unmarshal([]byte(task.Childs[i].Data), &account)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						data.Data = &account
						data.TaskID = task.Childs[i].TaskID

						_, err := accountClient.CreateAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Failed when CreateAccount:", err)
							return nil, err
						}

						task.Childs[i].IsParentActive = false
						reUpdate = true

					}

				}

			} else {

				data := account_pb.CreateAccountRequest{}
				account := account_pb.Account{}

				err = json.Unmarshal([]byte(task.Data), &account)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

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
					Cif:              account.Cif,
					ProductCode:      account.ProductCode,
					StatusCode:       account.StatusCode,
				}

				data.TaskID = task.TaskID

				if task.Status == 7 {

					_, err := accountClient.DeleteAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when DeleteAccount:", err)
						return nil, err
					}

				} else {

					_, err := accountClient.CreateAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
					if err != nil {
						logrus.Errorln("[api][func: SetTask] Failed when CreateAccount:", err)
						return nil, err
					}

				}

			}

		case "Notification":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			notificationConn, err := grpc.Dial(getEnv("NOTIFICATION_SERVICE", ":9094"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Notification Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer notificationConn.Close()

			companyClient := notification_pb.NewApiServiceClient(notificationConn)

			data := notification_pb.CreateNotificationRequest{}

			err = json.Unmarshal([]byte(task.Data), &data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID

			if task.Status == 7 {

				deleteReq := &notification_pb.CreateNotificationRequest{NotificationID: task.FeatureID}

				_, err := companyClient.DeleteNotification(ctx, deleteReq, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteNotification:", err)
					return nil, err
				}

			} else {

				data.NotificationID = task.FeatureID

				_, err := companyClient.CreateNotification(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateNotification:", err)
					return nil, err
				}

			}

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
				logrus.Errorln("[api][func: SetTask] Unable to connect Menu Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)

			if strings.Contains(task.Data, `"isParent": true`) {

				beforeSave := true

				for i := range task.Childs {

					if task.Childs[i].IsParentActive {

						data := menu_pb.SaveMenuAppearanceReq{}
						menu := menu_pb.MenuAppearance{}

						err = json.Unmarshal([]byte(task.Childs[i].Data), &menu)
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
							return nil, status.Errorf(codes.Internal, "Internal Error")
						}

						if beforeSave {

							_, err := menuClient.BeforeSaveMenuAppearance(ctx, &menu_pb.BeforeSaveMenuAppearanceReq{
								UserType: menu.UserType,
							}, grpc.Header(&header), grpc.Trailer(&trailer))
							if err != nil {
								logrus.Errorln("[api][func: SetTask] Failed when BeforeSaveMenuAppearance:", err)
								return nil, err
							}

							beforeSave = false

						}

						data.Data = &menu
						data.TaskID = task.Childs[i].TaskID

						_, err = menuClient.SaveMenuAppearance(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
						if err != nil {
							logrus.Errorln("[api][func: SetTask] Failed when SaveMenuAppearance:", err)
							return nil, err
						}

					}

				}

			} else {

				data := menu_pb.SaveMenuAppearanceReq{}
				menu := menu_pb.MenuAppearance{}

				err = json.Unmarshal([]byte(task.Data), &menu)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				data.Data = &menu
				data.TaskID = task.TaskID

				_, err = menuClient.SaveMenuAppearance(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when SaveMenuAppearance:", err)
					return nil, err
				}

			}

		case "Menu:License":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			menuConn, err := grpc.Dial(getEnv("MENU_SERVICE", ":9096"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Menu Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer menuConn.Close()

			menuClient := menu_pb.NewApiServiceClient(menuConn)

			isDeleted := false
			logrus.Printf("child menu 2 =======> %v", task.Childs)

			// if strings.Contains(task.Data, `"isParent": true`) {

			// 	for i := range task.Childs {

			// 		if task.Childs[i].IsParentActive {

			// 			data := menu_pb.SaveMenuLicenseReq{}
			// 			menu := menu_pb.MenuLicenseSave{}

			// 			err = json.Unmarshal([]byte(task.Childs[i].Data), &menu)
			// 			if err != nil {
			// 				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
			// 				return nil, status.Errorf(codes.Internal, "Internal Error")
			// 			}

			// 			data.Data = &menu
			// 			data.TaskID = task.Childs[i].TaskID

			// 			if !isDeleted {

			// 				_, err := menuClient.DeleteMenuLicenseCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			// 				if err != nil {
			// 					logrus.Errorln("[api][func: SetTask] Failed when DeleteMenuLicenseCompany:", err)
			// 					return nil, err
			// 				}

			// 				isDeleted = true

			// 			}

			// 			_, err := menuClient.SaveMenuLicense(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			// 			if err != nil {
			// 				logrus.Errorln("[api][func: SetTask] Failed when SaveMenuLicense:", err)
			// 				return nil, err
			// 			}

			// 		}

			// 	}

			// } else {

			data := menu_pb.SaveMenuLicenseReq{}
			// menu := []menu_pb.MenuLicenseSave{}
			parent := menu_pb.MenuParent{}

			err = json.Unmarshal([]byte(task.Data), &parent)

			// err = json.Unmarshal([]byte(task.Data), &menu)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.Data = parent.Menus
			data.TaskID = task.TaskID

			if !isDeleted {

				_, err := menuClient.DeleteMenuLicenseCompany(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteMenuLicenseCompany:", err)
					return nil, err
				}

				isDeleted = true

			}

			_, err = menuClient.SaveMenuLicense(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when SaveMenuLicense:", err)
				return nil, err
			}

			// }

		case "Role":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			roleConn, err := grpc.Dial(getEnv("ROLE_SERVICE", ":9098"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Role Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer roleConn.Close()

			client := role_pb.NewApiServiceClient(roleConn)

			data := role_pb.CreateRoleRequest{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID
			data.Data.RoleID = task.FeatureID

			if task.Status == 7 {

				_, err := client.DeleteRole(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteRole:", err)
					return nil, err
				}

			} else {

				res, err := client.CreateRole(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateRole:", err)
					return nil, err
				}

				task.FeatureID = res.Data.RoleID

				Role, err := json.Marshal(res.Data)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Marshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				task.Data = string(Role)

				reUpdate = true

			}

		case "Workflow":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Workflow Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer workflowConn.Close()

			client := workflow_pb.NewApiServiceClient(workflowConn)

			data := workflow_pb.CreateWorkflowRequest{}
			workflowTask := workflow_pb.WorkflowTask{}

			err = json.Unmarshal([]byte(task.Data), &workflowTask)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

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

			if task.Status == 7 {

				_, err := client.DeleteWorkflow(ctx, &workflowTask, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteWorkflow:", err)
					return nil, err
				}

			} else {

				_, err := client.CreateWorkflow(ctx, &workflowTask, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateWorkflow:", err)
					return nil, err
				}

			}

		case "Cash Pooling":

			logrus.Println("Liquidity")
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			liquidityConn, err := grpc.Dial(getEnv("LIQUIDITY_SERVICE", ":9010"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Liquidity Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer liquidityConn.Close()

			client := liquidity_pb.NewApiServiceClient(liquidityConn)

			data := liquidity_pb.CreateLiquidityRequest{}
			liquidityTask := liquidity_pb.CreateTaskLiquidityRequest{}

			err = json.Unmarshal([]byte(task.Data), &liquidityTask)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.Data = &liquidityTask
			data.TaskID = task.TaskID

			_, err = client.CreateLiquidity(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Failed when CreateLiquidity:", err)
				return nil, err
			}

		case "SSO:User":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			ssoConn, err := grpc.Dial(getEnv("SSO_SERVICE", ":9106"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect SSO Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer ssoConn.Close()

			ssoClient := sso_pb.NewApiServiceClient(ssoConn)

			data := sso_pb.WriteSyncUserTask{}

			err = json.Unmarshal([]byte(task.Data), &data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			send := &sso_pb.CreateSyncUserTaskReq{
				Data:   &data,
				TaskID: task.TaskID,
			}

			if task.Status == 7 {

				send.Data.User.UserSyncID = task.FeatureID

				_, err := ssoClient.DeleteSyncUser(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteSyncUser:", err)
					return nil, err
				}

			}

		case "SSO:Company":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			ssoConn, err := grpc.Dial(getEnv("SSO_SERVICE", ":9106"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect SSO Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer ssoConn.Close()

			ssoClient := sso_pb.NewApiServiceClient(ssoConn)

			data := sso_pb.WriteSyncCompanyTask{}

			err = json.Unmarshal([]byte(task.Data), &data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			send := &sso_pb.CreateSyncCompanyTaskReq{
				Data:   &data,
				TaskID: task.TaskID,
			}

			if task.Status == 7 {

				send.Data.Company.CompanySyncID = task.FeatureID

				_, err := ssoClient.DeleteSyncCompany(ctx, send.Data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteSyncCompany:", err)
					return nil, err
				}

			}

		case "System":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			systemConn, err := grpc.Dial(getEnv("SYSTEM_SERVICE", ":9101"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect System Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer systemConn.Close()

			systemClient := system_pb.NewApiServiceClient(systemConn)

			data := system_pb.CreateRequest{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID

			_, err = systemClient.CreateSystem(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
			if err != nil {
				return nil, err
			}

		case "Subscription":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			abonnementConn, err := grpc.Dial(getEnv("ABONNEMENT_SERVICE", ":9100"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Abonnement Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer abonnementConn.Close()

			abonnementClient := abonnement_pb.NewApiServiceClient(abonnementConn)

			data := abonnement_pb.CreateAbonnementRequest{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID

			isFirst := false

			if len(data.Data.BillingStatus) < 1 {

				isFirst = true
				data.Data.BillingStatus = "Waiting Schedule"

			}

			if task.Status == 7 {

				_, err := abonnementClient.DeleteAbonnement(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteAbonnement:", err)
					return nil, err
				}

			} else {

				res, err := abonnementClient.CreateAbonnement(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateAbonnement:", err)
					return nil, err
				}

				data.Data.Id = res.Data.Id

				dataUpdate, err := json.Marshal(data.Data)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Marshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}

				task.Data = string(dataUpdate)

				if isFirst {
					task.DataBak = task.Data
				}

				task.FeatureID = res.Data.Id
				reUpdate = true

			}

		case "Beneficiary Account":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			beneficiaryAccountConn, err := grpc.Dial(getEnv("BENEFICIARY_ACCOUNT_SERVICE", ":9107"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect Beneficiary Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer beneficiaryAccountConn.Close()

			beneficiaryAccountClient := beneficiary_account_pb.NewApiServiceClient(beneficiaryAccountConn)

			data := beneficiary_account_pb.CreateBeneficiaryAccountRequest{}

			err = json.Unmarshal([]byte(task.Data), &data.Data)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			data.TaskID = task.TaskID

			if task.Status == 7 {

				data.Data.BeneficiaryAccountID = task.FeatureID

				_, err := beneficiaryAccountClient.DeleteBeneficiaryAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteBeneficiaryAccount:", err)
					return nil, err
				}

			} else {

				res, err := beneficiaryAccountClient.CreateBeneficiaryAccount(ctx, &data, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateBeneficiaryAccount:", err)
					return nil, err
				}

				task.FeatureID = res.Data.BeneficiaryAccountID
				reUpdate = true

			}

		case "BG Mapping":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect BG Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer bgConn.Close()

			bgClient := bg_pb.NewApiServiceClient(bgConn)

			var taskData []*bg_pb.MappingData
			err = json.Unmarshal([]byte(currentData), &taskData)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			var taskDataBak []*bg_pb.MappingData
			if currentDataBak != "" && currentDataBak != "{}" {
				err = json.Unmarshal([]byte(currentDataBak), &taskDataBak)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}
			}

			if task.Status == 7 {

				_, err = bgClient.DeleteTransaction(ctx, &bg_pb.DeleteTransactionRequest{Type: "BG Mapping", MappingData: taskData, MappingDataBackup: taskDataBak}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteTransaction:", err)
					return nil, err
				}

			} else {

				_, err = bgClient.CreateTransaction(ctx, &bg_pb.CreateTransactionRequest{Type: "BG Mapping", MappingData: taskData, MappingDataBackup: taskDataBak}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateTransaction:", err)
					return nil, err
				}

			}

		case "BG Mapping Digital":

			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			bgConn, err := grpc.Dial(getEnv("BG_SERVICE", ":9124"), opts...)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to connect BG Service:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}
			defer bgConn.Close()

			bgClient := bg_pb.NewApiServiceClient(bgConn)

			var taskData []*bg_pb.MappingDigitalData
			err = json.Unmarshal([]byte(currentData), &taskData)
			if err != nil {
				logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
				return nil, status.Errorf(codes.Internal, "Internal Error")
			}

			var taskDataBak []*bg_pb.MappingDigitalData
			if currentDataBak != "" && currentDataBak != "{}" {
				err = json.Unmarshal([]byte(currentDataBak), &taskDataBak)
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Unable to Unmarshal Data:", err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				}
			}

			if task.Status == 7 {

				_, err = bgClient.DeleteTransaction(ctx, &bg_pb.DeleteTransactionRequest{Type: "BG Mapping Digital", MappingDigitalData: taskData, MappingDigitalDataBackup: taskDataBak}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when DeleteTransaction:", err)
					return nil, err
				}

			} else {

				_, err = bgClient.CreateTransaction(ctx, &bg_pb.CreateTransactionRequest{Type: "BG Mapping Digital", MappingDigitalData: taskData, MappingDigitalDataBackup: taskDataBak}, grpc.Header(&header), grpc.Trailer(&trailer))
				if err != nil {
					logrus.Errorln("[api][func: SetTask] Failed when CreateTransaction:", err)
					return nil, err
				}

			}

		}

	}

	if reUpdate {

		updatedTask, err = s.provider.UpdateTask(ctx, task, false)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Failed when UpdateTask:", err)
			return nil, err
		}

	}

	if updatedTask.Reasons != "" && req.Action == "approve" {
		updatedTask.Reasons = ""
	}

	// Save activity Log
	if getEnv("ENV", "LOCAL") != "LOCAL" {

		activityLog, err := GenerateActivityLog(task, currentUser, req.Action, "Update")
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Failed when GenerateActivityLog:", err)
			return nil, err
		}

		err = s.provider.SaveLog(ctx, activityLog)
		if err != nil {
			logrus.Errorln("[api][func: SetTask] Failed when SaveLog:", err)
		}

	}

	go TaskNotification(ctx, task, req.Action, sendTask, currentStatus, currentStep)

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

func (s *Server) GetTaskByIDNoFilter(ctx context.Context, req *pb.GetTaskByIDReq) (*pb.GetTaskByIDRes, error) {

	find, err := s.provider.FindTaskById(ctx, req.GetID())
	if err != nil {
		return nil, err
	}

	resData, err := find.ToPB(ctx)
	if err != nil {
		logrus.Errorln(err)
		err = nil
	}

	return &pb.GetTaskByIDRes{
		Found: false,
		Data:  &resData,
	}, err

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

	list, err := s.provider.GetListTask(ctx, &filter, &pb.PaginationResponse{}, sqlBuilder, 0, []uint64{}, []uint64{})
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
	task.DataBak = req.Data
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
			Currency:            req.GetTransactionCurrency(),
			SelectedAccountID:   req.GetSelectedAccountID(),
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
		activityLog, err := GenerateActivityLog(&task, nil, action, command)
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed to generate activity log: %v", err)
			return nil, err
		}
		err = s.provider.SaveLog(ctx, activityLog)
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
		}
	}

	return res, nil
}

func (s *Server) UpdateTaskRaw(ctx context.Context, req *pb.UpdateTaskRawReq) (*pb.SetTaskResponse, error) {

	result := &pb.SetTaskResponse{
		Error:   false,
		Code:    200,
		Message: "Task Updated",
	}

	var err error

	currentUser, userMd, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}

	if strings.ToLower(req.Action) == "delete" {
		allowed := checkAllowedApproval(userMd, req.Type, "data_entry:maker")
		if !allowed {
			return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
		}
	} else {
		allowed := checkAllowedApproval(userMd, req.Type, "approve:signer")
		if !allowed {
			return nil, status.Errorf(codes.PermissionDenied, "Permission Denied")
		}
	}

	task, _ := req.Task.ToORM(ctx)

	updatedTask, err := s.provider.UpdateTask(ctx, &task, req.UpdateChild)
	if err != nil {
		return nil, err
	}

	newestTask, err := s.provider.FindTaskById(ctx, updatedTask.TaskID)
	if err != nil {
		return nil, err
	}

	taskPb, _ := newestTask.ToPB(ctx)

	result.Data = &taskPb

	logrus.Println("Save LOG task 'update', type: ", task.Type)

	// Save activity Log
	if getEnv("ENV", "LOCAL") != "LOCAL" {

		logrus.Println("Set to save log")

		taskORM, err := req.Task.ToORM(ctx)
		if err != nil {
			return nil, err
		}

		activityLog, err := GenerateActivityLog(&taskORM, currentUser, req.Action, "Update")
		if err != nil {
			logrus.Errorln("Failed to generate activity log: %v", err)
			return nil, err
		}

		err = s.provider.SaveLog(ctx, activityLog)
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
		}

	}

	return result, nil

}

func (s *Server) DeleteDraftTask(ctx context.Context, req *pb.DeleteDraftTaskRequest) (*pb.DeleteDraftTaskResponse, error) {
	if req.GetTaskID() < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid task id")
	}
	var err error
	currentUser, _, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return nil, err
	}

	allowedProducts := []string{"Swift"}

	task, err := s.provider.FindTaskById(ctx, req.GetTaskID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error find task, error:", err.Error())
	}
	if task == nil {
		return nil, status.Errorf(codes.NotFound, "task not found")
	}

	isAllowedProduct := false
	for _, v := range allowedProducts {
		if v == task.Type {
			isAllowedProduct = true
		}
	}
	if !isAllowedProduct {
		return nil, status.Errorf(codes.InvalidArgument, "task type not allowed")
	}

	if task.Status != int32(pb.Statuses_Draft) {
		return nil, status.Errorf(codes.PermissionDenied, "task status is not draft")

	}

	if task.CreatedByID != currentUser.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "task not allowed to delete")
	}

	task.Status = int32(pb.Statuses_Deleted)

	_, err = s.provider.UpdateTask(ctx, task, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error delete task, error:", err.Error())

	}

	return &pb.DeleteDraftTaskResponse{
		Success: true,
	}, nil

}
