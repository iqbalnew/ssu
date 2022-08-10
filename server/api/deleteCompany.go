package api

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/account_service"
	beneficiary_account_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/beneficiary_account_service"
	role_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/role_service"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/workflow_service"
	menu_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/menu_service"

)

func (s *Server) DeleteCompany(ctx context.Context, companyID uint64) error {
	var err error

	currentUser, userMD, err := s.manager.GetMeFromMD(ctx)
	if err != nil {
		return err
	}

	if currentUser.UserType != "ba" {
		return status.Error(codes.PermissionDenied, "You are not authorized to delete company")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(context.Background(), md)
	}
	var trailer metadata.MD

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	// call service account, beneficiary, role, workflow, menu license
	accountConn, err := grpc.Dial(getEnv("ACCOUNT_SERVICE", ":9093"), opts...)
	if err != nil {
		logrus.Errorf("fail to dial account service: %v", err)
		return status.Errorf(codes.Internal, "fail to dial account service: %v", err)
	}
	defer accountConn.Close()
	accountClient := account_pb.NewApiServiceClient(accountConn)

	beneficiaryConn, err := grpc.Dial(getEnv("BENEFICIARY_ACCOUNT_SERVICE", ":9107"), opts...)
	if err != nil {
		logrus.Errorf("fail to dial beneficiary service: %v", err)
		return status.Errorf(codes.Internal, "fail to dial beneficiary service: %v", err)
	}
	defer beneficiaryConn.Close()
	beneficiaryClient := beneficiary_account_pb.NewApiServiceClient(beneficiaryConn)

	roleConn, err := grpc.Dial(getEnv("ROLE_SERVICE", ":9098"), opts...)
	if err != nil {
		logrus.Errorf("fail to dial role service: %v", err)
		return status.Errorf(codes.Internal, "fail to dial role service: %v", err)
	}
	defer roleConn.Close()
	roleClient := role_pb.NewApiServiceClient(roleConn)

	workflowConn, err := grpc.Dial(getEnv("WORKFLOW_SERVICE", ":9099"), opts...)
	if err != nil {
		logrus.Errorf("fail to dial workflow service: %v", err)
		return status.Errorf(codes.Internal, "fail to dial workflow service: %v", err)
	}
	defer workflowConn.Close()
	workflowClient := workflow_pb.NewApiServiceClient(workflowConn)

	menuConn, err := grpc.Dial(getEnv("MENU_SERVICE", ":9096"), opts...)
	if err != nil {
		logrus.Errorf("fail to dial menu service: %v", err)
		return status.Errorf(codes.Internal, "fail to dial menu service: %v", err)
	}
	defer menuConn.Close()
	menuClient := menu_pb.NewApiServiceClient(menuConn)

	return nil
}
