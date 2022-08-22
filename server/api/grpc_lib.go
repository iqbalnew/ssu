package api

import (
	"context"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	notification_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/notification_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func initNotifService() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	notificationConn, err := grpc.Dial(getEnv("NOTIFICATION_SERVICE", ":9094"), opts...)
	if err != nil {
		logrus.Errorln("Failed connect to Notification Service: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	return notificationConn, nil
}

func SendNotification(ctx context.Context, task *pb.TaskORM, eventName string) error {
	conn, err := initNotifService()
	if err != nil {
		logrus.Errorln("Failed connect to Notification Service: %v", err)
	}
	defer conn.Close()

	service := notification_pb.NewApiServiceClient(conn)

	resp, err := service.SendNotification(ctx, &notification_pb.SendNotificationRequest{
		FeatureName: task.Type,
		CompanyID:   task.CompanyID,
		RoleID:      []uint64{},
		EventName:   eventName,
		TaskID:      task.TaskID,
		Data:        task.Data,
	})

	if err != nil {
		return err
	}

	logrus.Info("Send Notif: %v", resp)

	return nil
}
