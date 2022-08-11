package api

import (
	"encoding/json"
	"fmt"

	"bitbucket.bri.co.id/scm/addons/addons-task-service/server/db"
	manager "bitbucket.bri.co.id/scm/addons/addons-task-service/server/jwt"
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
	role_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/role_service"
	sso_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/sso_service"
	system_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/system_service"
	users_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/user_service"
	workflow_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/workflow_service"
)

// generate activity log then Assign TaskID, Type, Description, Data and Key from input task data
func ActivityLogSetKey(task *pb.TaskORM) (*db.ActivityLog, error) {
	key := ""
	var err error = nil
	switch task.Type {
	case "Announcement":
		_, key, err = TaskDataAnnouncementToPB(task.Data)

	case "Company":
		_, key, err = TaskDataCompanyToPB(task.Data)

	case "Deposito":
		_, key, err = TaskDataDepositoToPB(task.Data)

	case "Account":
		_, key, err = TaskDataAccountToPB(task.Data)

	case "Notification":
		_, key, err = TaskDataNotificationToPB(task.Data)

	case "User":
		_, key, err = TaskDataUserToPB(task.Data)

	case "Menu:Appearance":
		_, key, err = TaskDataMenuAppearanceToPB(task.Data)

	case "Menu:License":
		_, key, err = TaskDataMenuLicenseToPB(task.Data)

	case "Role":
		_, key, err = TaskDataRoleToPB(task.Data)

	case "Workflow":
		_, key, err = TaskDataWorkflowToPB(task.Data)

	case "Liquidity":
		_, key, err = TaskDataLiquidityToPB(task.Data)

	case "SSO:User":
		_, key, err = TaskDataSSOUserToPB(task.Data)

	case "SSO:Company":
		_, key, err = TaskDataSSOCompanyToPB(task.Data)

	case "System":
		_, key, err = TaskDataSystemToPB(task.Data)

	case "Subscription":
		_, key, err = TaskDataSubscriptionToPB(task.Data)

	case "Beneficiary Account":
		_, key, err = TaskDataBeneficiaryAccountToPB(task.Data)

	case "BG Mapping":
		_, key, err = TaskDataBGMappingToPB(task.Data)

	case "BG Mapping Digital":
		_, key, err = TaskDataBGMappingDigitalToPB(task.Data)
	}

	if err != nil {
		return nil, err
	}

	activityLog := &db.ActivityLog{
		TaskID:      task.TaskID,
		Type:        task.Type,
		Description: task.Reasons,
		Data:        task,
		Key:         key,
	}

	return activityLog, nil
}

func GenerateActivityLog(task *pb.TaskORM, user *manager.UserData, action string, command string) (*db.ActivityLog, error) {
	activityLog, err := ActivityLogSetKey(task)
	if err != nil {
		return nil, err
	}
	activityLog.Command = command
	activityLog.Action = action
	if user != nil {
		activityLog.UserID = user.UserID
		activityLog.Username = user.Username
		activityLog.CompanyID = user.CompanyID
		activityLog.CompanyName = user.CompanyName
		activityLog.RoleIDs = user.RoleIDs
	}

	return activityLog, nil
}

func TaskDataUserToPB(data string) (val *users_pb.UserTaskData, key string, err error) {
	user := users_pb.UserTaskData{}
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, "", err
	}

	return &user, user.GetUser().GetUsername(), nil
}

func TaskDataCompanyToPB(data string) (val *company_pb.CreateCompanyTaskReq, key string, err error) {
	company := company_pb.CreateCompanyTaskReq{}
	err = json.Unmarshal([]byte(data), &company)
	if err != nil {
		return nil, "", err
	}
	return &company, company.Company.GetGroupName(), nil
}

func TaskDataAnnouncementToPB(data string) (val *announcement_pb.Announcement, key string, err error) {
	announcement := announcement_pb.Announcement{}
	err = json.Unmarshal([]byte(data), &announcement)
	if err != nil {
		return nil, "", err
	}
	return &announcement, announcement.GetCode(), nil
}

func TaskDataDepositoToPB(data string) (val *deposito_pb.CreateDepositoRequest, key string, err error) {
	deposito := deposito_pb.CreateDepositoRequest{}
	err = json.Unmarshal([]byte(data), &deposito)
	if err != nil {
		return nil, "", err
	}

	return &deposito, deposito.GetData().GetDeposito().GetCode(), nil
}

func TaskDataMenuAppearanceToPB(data string) (val *menu_pb.MenuAppearance, key string, err error) {
	menu := menu_pb.MenuAppearance{}
	err = json.Unmarshal([]byte(data), &menu)
	if err != nil {
		return nil, "", err
	}

	return &menu, menu.GetName(), nil
}

func TaskDataMenuLicenseToPB(data string) (val *menu_pb.MenuParent, key string, err error) {
	menu := menu_pb.MenuParent{}
	err = json.Unmarshal([]byte(data), &menu)
	if err != nil {
		return nil, "", err
	}

	return &menu, menu.GetCompanyName(), nil
}

func TaskDataLiquidityToPB(data string) (val *liquidity_pb.CreateTaskLiquidityRequest, key string, err error) {
	liquidity := liquidity_pb.CreateTaskLiquidityRequest{}
	err = json.Unmarshal([]byte(data), &liquidity)
	if err != nil {
		return nil, "", err
	}

	return &liquidity, liquidity.GetCode(), nil
}

func TaskDataAccountToPB(data string) (val *account_pb.Account, key string, err error) {
	account := account_pb.Account{}
	err = json.Unmarshal([]byte(data), &account)
	if err != nil {
		return nil, "", err
	}

	return &account, account.GetAccountName(), nil
}

func TaskDataSubscriptionToPB(data string) (val *abonnement_pb.ListTaskAbonnementRes, key string, err error) {
	subscription := abonnement_pb.ListTaskAbonnementRes{}
	err = json.Unmarshal([]byte(data), &subscription)
	if err != nil {
		return nil, "", err
	}

	return &subscription, subscription.GetCompany().GetName(), nil
}

func TaskDataBeneficiaryAccountToPB(data string) (val *beneficiary_account_pb.BeneficiaryAccount, key string, err error) {
	beneficiary := beneficiary_account_pb.BeneficiaryAccount{}
	err = json.Unmarshal([]byte(data), &beneficiary)
	if err != nil {
		return nil, "", err
	}

	return &beneficiary, beneficiary.GetAccountName(), nil
}

func TaskDataRoleToPB(data string) (val *role_pb.Role, key string, err error) {
	role := role_pb.Role{}
	err = json.Unmarshal([]byte(data), &role)
	if err != nil {
		return nil, "", err
	}

	return &role, role.GetName(), nil
}

func TaskDataWorkflowToPB(data string) (val *workflow_pb.WorkflowTask, key string, err error) {
	workflow := workflow_pb.WorkflowTask{}
	err = json.Unmarshal([]byte(data), &workflow)
	if err != nil {
		return nil, "", err
	}

	return &workflow, workflow.GetWorkflow().GetWorkflowCode(), nil
}

func TaskDataNotificationToPB(data string) (val *notification_pb.NotificationTask, key string, err error) {
	notification := notification_pb.NotificationTask{}
	err = json.Unmarshal([]byte(data), &notification)
	if err != nil {
		return nil, "", err
	}

	return &notification, notification.GetNotification().GetCode(), nil
}

func TaskDataSSOUserToPB(data string) (val *sso_pb.WriteSyncUserTask, key string, err error) {
	user := sso_pb.WriteSyncUserTask{}
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, "", err
	}

	return &user, fmt.Sprintf("%s:%d", user.GetClient(), user.GetUser().GetAddonsUserID()), nil
}

func TaskDataSSOCompanyToPB(data string) (val *sso_pb.WriteSyncCompanyTask, key string, err error) {
	company := sso_pb.WriteSyncCompanyTask{}
	err = json.Unmarshal([]byte(data), &company)
	if err != nil {
		return nil, "", err
	}

	return &company, fmt.Sprintf("%s:%d", company.GetClient(), company.Company.GetCompanyID()), nil
}

func TaskDataSystemToPB(data string) (val *system_pb.CreateRequest, key string, err error) {
	system := system_pb.CreateRequest{}
	err = json.Unmarshal([]byte(data), &system)
	if err != nil {
		return nil, "", err
	}

	return &system, system.GetData().GetKey(), nil
}

func TaskDataBGMappingToPB(data string) (val *bg_pb.MappingData, key string, err error) {
	mapping := bg_pb.MappingData{}
	err = json.Unmarshal([]byte(data), &mapping)
	if err != nil {
		return nil, "", err
	}

	return &mapping, mapping.GetCompanyName(), nil
}

func TaskDataBGMappingDigitalToPB(data string) (val *bg_pb.MappingDigitalData, key string, err error) {
	mapping := bg_pb.MappingDigitalData{}
	err = json.Unmarshal([]byte(data), &mapping)
	if err != nil {
		return nil, "", err
	}

	return &mapping, mapping.GetCompanyName(), nil
}
