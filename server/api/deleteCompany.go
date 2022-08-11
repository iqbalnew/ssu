package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	company_pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/company_service"
)

func (s *Server) DeleteCompany(ctx context.Context, dataString string) error {

	taskData := company_pb.CreateCompanyTaskReq{}
	json.Unmarshal([]byte(dataString), &taskData)
	companyID := taskData.Company.CompanyID

	// get task with company id
	listRes, err := s.GetListTask(ctx, &pb.ListTaskRequest{
		// Filter: "step:>=3",
		FilterOr: fmt.Sprintf("data.companyID:%v|data.company.companyID:%v", companyID,companyID),
		In: "type:Account,Beneficiary Account,Menu:License,Role,Workflow",
	})
	if err != nil {
		logrus.Errorf("error get list task: %v", err)
		return err
	}
	listTask := listRes.Data
	listTaskORM := []pb.TaskORM{}

	for _, task := range listTask {
		orm, _ := task.ToORM(ctx)
		listTaskORM = append(listTaskORM, orm)
	}

	err = s.provider.DeleteCompanyTask(ctx, listTaskORM)
	if err != nil {
		logrus.Errorf("error updpate company task: %v", err)
		return err
	}
	logrus.Infof("company delete =====> update status to 6")

	for _, task := range listTask {
		// delete task
		if !task.IsParentActive {
			_, err = s.SetTask(ctx, &pb.SetTaskRequest{
				TaskID: task.TaskID,
				Action: "approve",
				Comment: "this company is deleted",
			})
			if err != nil {
				logrus.Errorf("error delete task: %v", err)
				return err
			}
			logrus.Infof("company delete =====> delete task: %v", task.TaskID)
		}
	}

	return nil
}
