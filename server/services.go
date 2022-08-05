package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	announcementConn *grpc.ClientConn
	accountConn      *grpc.ClientConn
	companyConn      *grpc.ClientConn
)

func initOtherServicesConn() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	announcementConn, err := grpc.Dial(config.AnnouncementService, opts...)
	if err != nil {
		logrus.Fatalf("Failed connect to Announcement Service: %v", err)
		os.Exit(1)
		return
	}
	logrus.Println("Announcement Service Connected, on %s", config.AnnouncementService)

	defer announcementConn.Close()
}
