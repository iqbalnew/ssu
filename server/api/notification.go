package api

import (
	"context"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
)

func TaskNotification(ctx context.Context, task *pb.TaskORM, action string, sendTask bool) {
	eventName := ""

	switch strings.ToLower(action) {
	case "rework":
		// when task rework in company module
		if task.Type == "Company" {
			eventName = "Group sent for rework"
		}
	case "approve":
		// when task approve in company module with status Request for delete
		if task.Type == "Company" && task.Status == 6 {
			eventName = "Request to delete group"
		}
	case "reject":
		// when task rejected in company module for create company
		if task.Type == "Company" {
			eventName = "Group gets rejected"
		}
	case "delete":
	}

	if sendTask {
		switch task.Type {
		case "Company":
			// when task company created
			eventName = "Group data created and/or sent for approval"
			// when task get approval to delete company
			if task.Status == 7 {
				eventName = "Delete request gets approval"
			}
		}
	}

	SendNotification(ctx, task, eventName)
}
