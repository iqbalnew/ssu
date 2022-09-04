package api

import (
	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/common/constant"
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

// TaskNotificationCreateOrUpdate is function to send notification when task created or updated
func TaskNotificationCreateOrUpdate(ctx context.Context, task *pb.TaskORM, action string, currentStatus, currentStep int32) {
	eventName := ""
	switch action {
	case constant.Create:
		// module company
		if task.Type == constant.ModuleCompany {
			eventName = constant.CreateCompanySentApproval
		}

		// module user
		if task.Type == constant.ModuleUser {
			eventName = constant.CreateUserSentApproval
		}

		// module role
		if task.Type == constant.ModuleRole {
			eventName = constant.CreateRoleSentApproval
		}

		// module beneficiary account
		if task.Type == constant.ModuleBeneficiaryAccount {
			eventName = constant.CreateBeneficiaryAccountSentApproval
		}
	case constant.Update:
		// module company
		if task.Type == constant.ModuleCompany {
			eventName = constant.UpdateCompanySentApproval
		}

		// module user
		if task.Type == constant.ModuleUser {
			eventName = constant.UpdateUserSentApproval
		}

		// module role
		if task.Type == constant.ModuleRole {
			eventName = constant.UpdateRoleSentApproval
		}

		// module beneficiary account
		if task.Type == constant.ModuleBeneficiaryAccount {
			eventName = constant.UpdateBeneficiaryAccountSentApproval
		}
	}

	if eventName != "" {
		SendNotification(ctx, task, eventName)
	}
}
