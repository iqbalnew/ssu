package api

import (
	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/common/constant"
	"context"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
)

func TaskNotification(ctx context.Context, task *pb.TaskORM, action string, sendTask bool, currentStatus, currentStep int32) {
	eventName := ""
	switch strings.ToLower(action) {
	// trigger by action rework
	case "rework":
		// module company
		if task.Type == constant.ModuleCompany {
			// company created for rework
			eventName = constant.CompanyCreatedForRework
			if task.LastApprovedByID != 0 && task.LastRejectedByID != 0 {
				// company edit for rework
				eventName = constant.CompanyEditForRework
			}
		}

		// module beneficiary account
		if task.Type == constant.ModuleBeneficiaryAccount {
			// beneficiary account created for rework
			eventName = constant.BeneficiaryAccountCreatedForRework
			if task.LastApprovedByID != 0 && task.LastRejectedByID != 0 {
				// beneficiary account edit for rework
				eventName = constant.BeneficiaryAccountEditForRework
			}
		}
	// trigger by action rejected
	case "reject":
		// when task rejected in company module for create company
		if task.Type == constant.ModuleCompany {
			// company create gets rejected
			eventName = constant.CompanyCreateGetsRejected
			if task.LastApprovedByID != 0 && task.LastRejectedByID != 0 {
				// company edit gets rejected
				eventName = constant.CompanyEditGetsRejected
			} else if currentStatus == constant.RequestForDelete {
				// company delete gets rejected
				eventName = constant.CompanyDeleteGetsRejected
			}
		}

		// module beneficiary account
		if task.Type == constant.ModuleBeneficiaryAccount {
			// beneficiary account create gets rejected
			eventName = constant.BeneficiaryAccountCreateGetsRejected
			if task.LastApprovedByID != 0 && task.LastRejectedByID != 0 {
				// beneficiary account edit gets rejected
				eventName = constant.BeneficiaryAccountEditGetsRejected
			} else if currentStatus == constant.RequestForDelete {
				// beneficiary account delete gets rejected
				eventName = constant.BeneficiaryAccountDeleteGetsRejected
			}
		}
	// trigger by action request for delete
	case "delete":
		// module company
		if task.Type == constant.ModuleCompany {
			// company request for delete
			eventName = constant.CompanyRequestForDelete
		}

		// module beneficiary account
		if task.Type == constant.ModuleBeneficiaryAccount {
			// beneficiary account request for delete
			eventName = constant.BeneficiaryAccountRequestForDelete
		}
	}

	// trigger by status approved deleted dan created
	if sendTask {
		switch task.Type {
		case constant.ModuleCompany:
			// company gets approved created
			eventName = constant.CompanyGetsApprovedCreated
			if task.Status == constant.Deleted {
				// company gets approved deleted
				eventName = constant.CompanyGetsApprovedDeleted
			}

		// module beneficiary account
		case constant.ModuleBeneficiaryAccount:
			// beneficiary account gets approved created
			eventName = constant.BeneficiaryAccountGetsApprovedCreated
			if task.Status == constant.Deleted {
				// beneficiary account gets approved deleted
				eventName = constant.BeneficiaryAccountGetsApprovedDeleted
			}
		}
	}

	if eventName != "" {
		SendNotification(ctx, task, eventName)
	}
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
