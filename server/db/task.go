package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GraphResult struct {
	Name  uint32
	Type  string
	Total uint64
}

type GraphResultColumnType struct {
	Name  string
	Type  string
	Total uint64
}

type GraphResultWorkflowType struct {
	Name  string
	Type  string
	Total uint64
}

type ProductAccountFilter struct {
	ProductName          string
	AccountIDs           []uint64
	HasAuthorityMaker    bool
	HasAuthorityChecker  bool
	HasAuthoritySigner   bool
	HasAuthorityReleaser bool
}

func (p *GormProvider) GetGraphStepAll(ctx context.Context, idCompany string) (result *GraphResult, err error) {

	selectOpt := "count(*) as total"
	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	// if service != "" {
	// 	whereOpt = fmt.Sprintf("type = '%v'", service)
	// }
	if idCompany != "" {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + `( ("data" -> 'user'->> 'companyID' = '` + idCompany + `' OR "data" -> 'companyID' = '` + idCompany + `' OR "data" -> 'company' ->> 'companyID' = '` + idCompany + `'
		OR  "data" @> '[{"companyID":` + idCompany + `}]') 
		or ("type" = 'BG Issuing'  AND  company_id = '` + idCompany + `'))`
	}

	if whereOpt != "" {
		whereOpt = whereOpt + " AND "
	}
	whereOpt = fmt.Sprintf("%v (status = %v   or status = %v)", whereOpt, 1, 3)

	if whereOpt != "" {
		whereOpt = whereOpt + " AND status != 7 AND is_parent_active = false"
	} else {
		whereOpt = "status != 7 AND is_parent_active = false"
	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	return result, nil

}

func (p *GormProvider) GetGraphPendingTaskWithWorkflow(ctx context.Context, companyIDFilter uint64, workflowUserIDFilter uint64, workflowRoleIDFilter []uint64, workflowAccountIDFilter []*ProductAccountFilter) (result []*GraphResultWorkflowType, err error) {

	if len(workflowRoleIDFilter) < 1 {
		return []*GraphResultWorkflowType{}, nil
	}

	selectOpt := `CASE WHEN workflow_doc->'workflow'->>'currentStep' IS NULL THEN 'maker' ELSE workflow_doc->'workflow'->>'currentStep' END AS name, "type", count(*) as total`

	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select(selectOpt).Where("status NOT IN (0, 4, 5, 7, 8)")

	if companyIDFilter > 0 {
		query = query.Where(fmt.Sprintf(`("data"->'user'->>'companyID' = '%d' OR "data"->'companyID' = '%d' OR "data"->'company'->>'companyID' = '%d' OR "data" @> '[{"companyID":%d}]' OR "company_id" = '%d')`, companyIDFilter, companyIDFilter, companyIDFilter, companyIDFilter, companyIDFilter))
	}

	customQuery := ""

	if len(workflowRoleIDFilter) > 0 {
		value := strings.ReplaceAll(fmt.Sprint(workflowRoleIDFilter), " ", "','")
		value = strings.ReplaceAll(value, "[", "'")
		value = strings.ReplaceAll(value, "]", "'")
		if customQuery == "" {
			customQuery = fmt.Sprintf("ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", value)
		} else {
			customQuery = fmt.Sprintf("%s AND ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", customQuery, value)
		}
	}

	accountIDQuery := ""
	makerQuery := ""

	for _, v := range workflowAccountIDFilter {

		accountIDs := ""

		for i, accountID := range v.AccountIDs {
			if i > 0 {
				accountIDs = fmt.Sprintf("%s,%d", accountIDs, accountID)
			} else {
				accountIDs = fmt.Sprint(accountID)
			}
		}

		logrus.Printf("[db][func: GetGraphPendingTaskWithWorkflow] Product Name: %s, UserID: %v, Has Authority Maker: %v", v.ProductName, workflowUserIDFilter, v.HasAuthorityMaker)

		logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Account IDs:", v.AccountIDs)
		logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Account ID String:", accountIDs)

		if accountIDs != "" {

			csrQuery := ""

			if v.HasAuthorityChecker {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier')", csrQuery)
				}
			}

			if v.HasAuthoritySigner {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'signer')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'signer')", csrQuery)
				}
			}

			if v.HasAuthorityReleaser {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'releaser')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'releaser')", csrQuery)
				}
			}

			if csrQuery != "" {
				csrQuery = fmt.Sprintf("AND (%s)", csrQuery)
			}

			if accountIDQuery == "" {
				accountIDQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT IN (%s)) %s)", v.ProductName, accountIDs, csrQuery)
			} else {
				accountIDQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT IN (%s)) %s)", accountIDQuery, v.ProductName, accountIDs, csrQuery)
			}

			if v.HasAuthorityMaker {

				payrollQuery := ""
				if v.ProductName == "Payroll Transfer" {
					payrollQuery = "AND data->>'status' = 'Ready to Submit'"
				}

				if makerQuery == "" {
					makerQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)) %s)", v.ProductName, accountIDs, payrollQuery)
				} else {
					makerQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)) %s)", makerQuery, v.ProductName, accountIDs, payrollQuery)
				}

			}

		}

	}

	if accountIDQuery != "" {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s)", accountIDQuery)
		} else {
			customQuery = fmt.Sprintf("%s AND (%s)", customQuery, accountIDQuery)
		}
	}

	logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Account ID Query:", accountIDQuery)

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", workflowUserIDFilter)
		} else {
			customQuery = fmt.Sprintf("%s AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", customQuery, workflowUserIDFilter)
		}
	}

	if customQuery != "" {
		customQuery = fmt.Sprintf("(%s AND status = '1')", customQuery)
	}

	logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Maker Query:", makerQuery)

	if makerQuery != "" {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s)", makerQuery)
		} else {
			customQuery = fmt.Sprintf(`%s OR (%s)`, customQuery, makerQuery)
		}
	}

	logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Custom Query list ========== ========== ==========")
	logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Custom Query list:", customQuery)
	logrus.Println("[db][func: GetGraphPendingTaskWithWorkflow] Custom Query list ========== ========== ==========")

	if customQuery != "" {
		query = query.Where(fmt.Sprintf("(%s)", customQuery))
	}

	query = query.Group("name, type")

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	return result, nil

}

func (p *GormProvider) GetGraphStep(ctx context.Context, idCompany uint64, service string, step uint, stat uint, isIncludeApprove bool, isIncludeReject bool, userType string) (result []*GraphResult, err error) {

	selectOpt := "step as name, type, count(*) as total"
	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	if service != "" {
		whereOpt = fmt.Sprintf("type = '%v'", service)
	}
	if idCompany > 0 {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf(`%s ( 
			(
				"data" -> 'user'->> 'companyID' = '%d' 
				OR "data" -> 'companyID' = '%d' 
				OR "data" -> 'company' ->> 'companyID' = '%d'
				OR  "data" @> '[{"companyID":%d}]'
				OR "company_id" = '%d'
			)
		)`, whereOpt, idCompany, idCompany, idCompany, idCompany, idCompany)
	}

	if userType == "ca" {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf(`%s ("data"->'userTypeID' = '4' OR type != 'Role') AND (data->'user'->>'userTypeName' = 'Customer User' OR type != 'User') `, whereOpt)
	}

	if !isIncludeApprove {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + "status != 4"
	}
	if !isIncludeReject {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + "status != 5"
	}
	if stat > 0 {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf("%v status = %v", whereOpt, stat)
	}
	if step > 0 {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf("%v step = %v", whereOpt, step)
	}

	// elemintate deleted task and child
	if whereOpt != "" {
		whereOpt = whereOpt + " AND status != 7 AND is_parent_active = false"
	} else {
		whereOpt = "status != 7 AND is_parent_active = false"
	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group("step, type")

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	return result, nil

}

func (p *GormProvider) GetGraphServiceType(ctx context.Context, service string, stat uint, filter string) (result []*GraphResultColumnType, err error) {

	if filter == "" {
		return nil, status.Errorf(codes.InvalidArgument, "column cant be empty")
	}

	column := columnNameBuilder(filter, false)
	selectOpt := fmt.Sprintf("%s as name, type, count(*) as total", column)
	query := p.db_main.Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	if service != "" {
		whereOpt = fmt.Sprintf("type = '%v'", service)
	}
	if stat > 0 {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf("%v status = %v", whereOpt, stat)
	}

	// elemintate deleted task
	if whereOpt != "" {
		whereOpt = whereOpt + " AND status != 7"
	} else {
		whereOpt = "status != 7"
	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group(fmt.Sprintf("%s, type", column))

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}
	return result, nil

}

func (p *GormProvider) GetGraphStatus(ctx context.Context, service string, stat uint) (result []*GraphResult, err error) {

	selectOpt := "status as name, type, count(*) as total"
	query := p.db_main.Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	if service != "" {
		whereOpt = fmt.Sprintf("type = '%v'", service)
	}
	if stat > 0 {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = fmt.Sprintf("%v status = %v", whereOpt, stat)
	}

	// elemintate deleted task
	if whereOpt != "" {
		whereOpt = whereOpt + " AND status != 7"
	} else {
		whereOpt = "status != 7"
	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group("status, type")

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}
	return result, nil

}

func (p *GormProvider) CreateTask(ctx context.Context, task *pb.TaskORM) (*pb.TaskORM, error) {
	query := p.db_main.Clauses(clause.OnConflict{
		UpdateAll: true,
	})

	childs := task.Childs
	task.Childs = nil

	if err := query.Debug().Create(&task).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	for i := range childs {
		childs[i].ParentID = &task.TaskID
	}

	if len(childs) > 0 {
		if err := query.Debug().Model(&childs).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "task_id"}},
			UpdateAll: true,
		}).Create(&childs).Error; err != nil {
			if !errors.Is(err, gorm.ErrModelValueRequired) {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
			}
		}
	}

	return task, nil
}

func (p *GormProvider) FindAndSetStatus(ctx context.Context, taskID uint64, stat int) (*pb.TaskORM, error) {
	task, err := p.FindTaskById(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// updateData := map[string]interface{}{
	// 	"status": stat,
	// }

	updateData := &pb.TaskORM{Status: int32(stat)}

	listIDs := []uint64{task.TaskID}
	for _, v := range task.Childs {
		if task.IsParentActive {
			listIDs = append(listIDs, v.TaskID)
		}
	}

	if err := p.db_main.Model(&pb.TaskORM{}).Where("task_id IN ?", listIDs).Updates(&updateData).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	return task, nil
}

func (p *GormProvider) UpdateTask(ctx context.Context, task *pb.TaskORM, updateChild bool) (*pb.TaskORM, error) {

	if task.Status == 5 {
		updateData := map[string]interface{}{
			"step": task.Step,
		}

		listIDs := []uint64{task.TaskID}
		for _, v := range task.Childs {
			if task.IsParentActive {
				listIDs = append(listIDs, v.TaskID)
			}
		}

		if err := p.db_main.Debug().Model(&pb.TaskORM{}).Where("task_id IN ?", listIDs).Updates(&updateData).Error; err != nil {
			return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
		}

	}

	taskModel := pb.TaskORM{TaskID: task.TaskID}
	query := p.db_main.Clauses(clause.OnConflict{
		UpdateAll: true,
	})
	if task != nil || task.Type != "" {
		childs := task.Childs
		task.Childs = []*pb.TaskORM{}
		if err := query.Model(&taskModel).Updates(&task).Error; err != nil {
			if !errors.Is(err, gorm.ErrModelValueRequired) {
				return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
			} else {
				return nil, status.Errorf(codes.NotFound, "Task Not Found")
			}
		}

		if len(childs) > 0 && (updateChild || task.Type == "Menu:Appearance" || task.Type == "Menu:License") {

			for i := range childs {
				childs[i].ParentID = &task.TaskID
				task.Childs = append(task.Childs, childs[i])
			}

			if task.Type == "Menu:Appearance" || task.Type == "Menu:License" {
				if err := p.db_main.Where("parent_id", task.TaskID).Delete(&pb.TaskORM{}).Error; err != nil {
					return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
				}
			}

			if err := p.db_main.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "task_id"}},
				UpdateAll: true,
			}).Create(&childs).Error; err != nil {
				if !errors.Is(err, gorm.ErrModelValueRequired) {
					return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
				}
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
				}
			}

		}
	}

	taskModel = pb.TaskORM{TaskID: task.TaskID}
	return &taskModel, nil
}

func (p *GormProvider) FindTaskById(ctx context.Context, id uint64) (*pb.TaskORM, error) {
	task := &pb.TaskORM{TaskID: id}
	if err := p.db_main.Preload(clause.Associations).First(&task).Error; err != nil {
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
		} else {
			return nil, status.Errorf(codes.NotFound, "Task Not Found")
		}
	}
	return task, nil
}

func (p *GormProvider) GetListTask(ctx context.Context, filter *pb.TaskORM, pagination *pb.PaginationResponse, sql *QueryBuilder, workflowUserIDFilter uint64, workflowRoleIDFilter []uint64, workflowAccountIDFilter []*ProductAccountFilter) (tasks []*pb.TaskORM, err error) {

	query := p.db_main
	if filter.Type != "" {
		query = query.Debug()
	}

	query = query.Select("*", " CASE WHEN status = '3' or status = '5' THEN last_rejected_by_name ELSE last_approved_by_name END AS reviewed_by").Where("status != 7")
	if filter != nil {
		query = query.Where(&filter)
	}

	logrus.Println("[db][func: GetListTask] Product IN Length:", len(sql.ProductIn))

	if len(sql.ProductIn) > 0 && sql.ProductIn[0] != "" {
		query = query.Where("type IN ?", sql.ProductIn)
	}

	customQuery := ""

	if len(workflowRoleIDFilter) > 0 {
		value := strings.ReplaceAll(fmt.Sprint(workflowRoleIDFilter), " ", "','")
		value = strings.ReplaceAll(value, "[", "'")
		value = strings.ReplaceAll(value, "]", "'")
		if customQuery == "" {
			customQuery = fmt.Sprintf("ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", value)
		} else {
			customQuery = fmt.Sprintf("%s AND ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", customQuery, value)
		}
	}

	accountIDQuery := ""
	makerQuery := ""

	for _, v := range workflowAccountIDFilter {

		accountIDs := ""

		for i, accountID := range v.AccountIDs {
			if i > 0 {
				accountIDs = fmt.Sprintf("%s,%d", accountIDs, accountID)
			} else {
				accountIDs = fmt.Sprint(accountID)
			}
		}

		logrus.Printf("[db][func: GetListTask] Product Name: %s, UserID: %v, Has Authority Maker: %v", v.ProductName, workflowUserIDFilter, v.HasAuthorityMaker)

		logrus.Println("[db][func: GetListTask] Account ID String:", accountIDs)

		if accountIDs != "" {

			csrQuery := ""

			if v.HasAuthorityChecker {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier'))", csrQuery)
				}
			}

			if v.HasAuthoritySigner {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'signer')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'signer')", csrQuery)
				}
			}

			if v.HasAuthorityReleaser {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'releaser')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'releaser')", csrQuery)
				}
			}

			if csrQuery != "" {
				csrQuery = fmt.Sprintf("AND (%s)", csrQuery)
			}

			if accountIDQuery == "" {
				accountIDQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT  IN (%s)) %s)", v.ProductName, accountIDs, csrQuery)
			} else {
				accountIDQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT IN (%s)) %s)", accountIDQuery, v.ProductName, accountIDs, csrQuery)
			}

			if v.HasAuthorityMaker {

				if makerQuery == "" {
					makerQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)))", v.ProductName, accountIDs)
				} else {
					makerQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)))", makerQuery, v.ProductName, accountIDs)
				}

				if makerQuery == "" {
					makerQuery = fmt.Sprintf("(type = '%s' AND ('%d' = ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[])))", v.ProductName, workflowUserIDFilter)
				} else {
					makerQuery = fmt.Sprintf("%s OR (type = '%s' AND ('%d' = ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[])))", makerQuery, v.ProductName, workflowUserIDFilter)
				}

			}

		}

	}

	if accountIDQuery != "" {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s)", accountIDQuery)
		} else {
			customQuery = fmt.Sprintf("%s AND (%s)", customQuery, accountIDQuery)
		}
	}

	logrus.Println("[db][func: GetListTask] Account ID Query:", accountIDQuery)

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", workflowUserIDFilter)
		} else {
			customQuery = fmt.Sprintf("%s AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", customQuery, workflowUserIDFilter)
		}
	}

	if customQuery != "" {
		customQuery = fmt.Sprintf("(%s)", customQuery)
	}

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(('%d' = ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[])) AND status IN (1, 4, 5))", workflowUserIDFilter)
		} else {
			customQuery = fmt.Sprintf("%s OR (('%d' = ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[])) AND status IN (1, 4, 5))", customQuery, workflowUserIDFilter)
		}
	}

	if makerQuery != "" {
		makerQuery = fmt.Sprintf("(%s) OR", makerQuery)
	}

	logrus.Println("[db][func: GetListTask] Maker Query:", makerQuery)

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s workflow_doc->'workflow'->'createdBy'->>'userID' = '%d')", makerQuery, workflowUserIDFilter)
		} else {
			customQuery = fmt.Sprintf(`%s OR (%s workflow_doc->'workflow'->'createdBy'->>'userID' = '%d')`, customQuery, makerQuery, workflowUserIDFilter)
		}
	}

	logrus.Println("[db][func: GetListTask] Custom Query list ========== ========== ==========")
	logrus.Println("[db][func: GetListTask] Custom Query list:", customQuery)
	logrus.Println("[db][func: GetListTask] Custom Query list ========== ========== ==========")

	query = query.Scopes(FilterScoope(sql.Filter))
	query = query.Scopes(FilterOrScoope(sql.FilterOr, ""))

	if customQuery != "" {
		query = query.Where(fmt.Sprintf("(%s)", customQuery))
	}

	query = query.Scopes(QueryScoop(sql.CollectiveAnd), WhereInScoop(sql.In), WhereInScoop(sql.MeFilterIn), NotConditionalScoope(sql.FilterNot))

	if sql.CompanyID != "" {
		query = query.Where(`("company_id" = $1 OR "data" ->> 'companyID' = $2)`, sql.CompanyID, sql.CompanyID)
	}

	query = query.Scopes(DistinctScoope(sql.Distinct))
	query = query.Scopes(Paginate(tasks, pagination, query), CustomOrderScoop(sql.CustomOrder), Sort(sql.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))

	if err := query.Preload(clause.Associations).Debug().Find(&tasks).Error; err != nil {
		logrus.Errorln("[db][func: GetListTask] Failed:", err.Error())
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
	}

	return tasks, nil

}

func (p *GormProvider) GetListTaskNormal(ctx context.Context, filter *pb.TaskORM, pagination *pb.PaginationResponse, sql *QueryBuilder, companyIDFilter uint64, workflowUserIDFilter uint64, workflowRoleIDFilter []uint64, workflowAccountIDFilter []*ProductAccountFilter) (tasks []*pb.TaskORM, err error) {

	query := p.db_main
	if filter.Type != "" {
		query = query.Debug()
	}

	query = query.Select("*", " CASE WHEN status = '3' or status = '5' THEN last_rejected_by_name ELSE last_approved_by_name END AS reviewed_by").Where("status != 7")
	if filter != nil {
		query = query.Where(&filter)
	}

	if len(sql.ProductIn) > 0 && sql.ProductIn[0] != "" {
		query = query.Where("type IN ?", sql.ProductIn)
	}

	if companyIDFilter > 0 {
		query = query.Where(fmt.Sprintf(`"data"->'user'->>'companyID' = '%d' OR "data"->'companyID' = '%d' OR "data"->'company'->>'companyID' = '%d' OR "data" @> '[{"companyID":%d}]' OR "company_id" = '%d'`, companyIDFilter, companyIDFilter, companyIDFilter, companyIDFilter, companyIDFilter))
	}

	customQuery := ""

	if len(workflowRoleIDFilter) > 0 {
		value := strings.ReplaceAll(fmt.Sprint(workflowRoleIDFilter), " ", "','")
		value = strings.ReplaceAll(value, "[", "'")
		value = strings.ReplaceAll(value, "]", "'")
		if customQuery == "" {
			customQuery = fmt.Sprintf("ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", value)
		} else {
			customQuery = fmt.Sprintf("%s AND ARRAY(SELECT jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && ARRAY[%s]", customQuery, value)
		}
	}

	accountIDQuery := ""
	makerQuery := ""

	for _, v := range workflowAccountIDFilter {

		accountIDs := ""

		for i, accountID := range v.AccountIDs {
			if i > 0 {
				accountIDs = fmt.Sprintf("%s,%d", accountIDs, accountID)
			} else {
				accountIDs = fmt.Sprint(accountID)
			}
		}

		logrus.Printf("[db][func: GetListTaskNormal] Product Name: %s, UserID: %v, Has Authority Maker: %v", v.ProductName, workflowUserIDFilter, v.HasAuthorityMaker)

		logrus.Println("[db][func: GetListTaskNormal] Account ID String:", accountIDs)

		if accountIDs != "" {

			csrQuery := ""

			if v.HasAuthorityChecker {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'checker' OR workflow_doc->'workflow'->>'currentStep' = 'verifier')", csrQuery)
				}
			}

			if v.HasAuthoritySigner {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'signer')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'signer')", csrQuery)
				}
			}

			if v.HasAuthorityReleaser {
				if csrQuery == "" {
					csrQuery = "(workflow_doc->'workflow'->>'currentStep' = 'releaser')"
				} else {
					csrQuery = fmt.Sprintf("%s OR (workflow_doc->'workflow'->>'currentStep' = 'releaser')", csrQuery)
				}
			}

			if csrQuery != "" {
				csrQuery = fmt.Sprintf("AND (%s)", csrQuery)
			}

			if accountIDQuery == "" {
				accountIDQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT IN (%s)) %s)", v.ProductName, accountIDs, csrQuery)
			} else {
				accountIDQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'header'->'uaID' IS NULL OR (workflow_doc->'workflow'->'header'->'uaID')::INT IN (%s)) %s)", accountIDQuery, v.ProductName, accountIDs, csrQuery)
			}

			if v.HasAuthorityMaker {

				payrollQuery := ""
				if v.ProductName == "Payroll Transfer" {
					payrollQuery = "AND data->>'status' = 'Ready to Submit'"
				}

				if makerQuery == "" {
					makerQuery = fmt.Sprintf("(type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)) %s)", v.ProductName, accountIDs, payrollQuery)
				} else {
					makerQuery = fmt.Sprintf("%s OR (type = '%s' AND (workflow_doc->'workflow'->'createdBy'->>'userID' IS NULL OR workflow_doc->>'nextStatus' = 'returned') AND (data->'uaID' IS NULL OR (data->'uaID')::INT IN (%s)) %s)", makerQuery, v.ProductName, accountIDs, payrollQuery)
				}

			}

		}

	}

	if accountIDQuery != "" {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s)", accountIDQuery)
		} else {
			customQuery = fmt.Sprintf("%s AND (%s)", customQuery, accountIDQuery)
		}
	}

	logrus.Println("[db][func: GetListTaskNormal] Account ID Query:", accountIDQuery)

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", workflowUserIDFilter)
		} else {
			customQuery = fmt.Sprintf("%s AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR '%d' != ANY(TRANSLATE(workflow_doc->'workflow'->>'participantUserIDs', '[]', '{}')::INT[]))", customQuery, workflowUserIDFilter)
		}
	}

	if customQuery != "" {
		customQuery = fmt.Sprintf("(%s AND status = '1')", customQuery)
	}

	logrus.Println("[db][func: GetListTaskNormal] Maker Query:", makerQuery)

	if makerQuery != "" {
		if customQuery == "" {
			customQuery = fmt.Sprintf("(%s)", makerQuery)
		} else {
			customQuery = fmt.Sprintf(`%s OR (%s)`, customQuery, makerQuery)
		}
	}

	logrus.Println("[db][func: GetListTaskNormal] Custom Query list ========== ========== ==========")
	logrus.Println("[db][func: GetListTaskNormal] Custom Query list:", customQuery)
	logrus.Println("[db][func: GetListTaskNormal] Custom Query list ========== ========== ==========")

	query = query.Scopes(FilterScoope(sql.Filter))
	query = query.Scopes(FilterOrScoope(sql.FilterOr, ""))

	if customQuery != "" {
		query = query.Where(fmt.Sprintf("(%s)", customQuery))
	}

	query = query.Scopes(QueryScoop(sql.CollectiveAnd), WhereInScoop(sql.In), WhereInScoop(sql.MeFilterIn), NotConditionalScoope(sql.FilterNot))

	if sql.CompanyID != "" {
		query = query.Where(`("company_id" = $1 OR "data" ->> 'companyID' = $2)`, sql.CompanyID, sql.CompanyID)
	}

	query = query.Scopes(DistinctScoope(sql.Distinct))
	query = query.Scopes(Paginate(tasks, pagination, query), CustomOrderScoop(sql.CustomOrder), Sort(sql.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))

	if err := query.Preload(clause.Associations).Debug().Find(&tasks).Error; err != nil {
		logrus.Errorln("[db][func: GetListTaskNormal] Failed:", err.Error())
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
	}

	return tasks, nil

}

func (p *GormProvider) GetListTaskPluck(ctx context.Context, key string, filter *pb.TaskORM, sql *QueryBuilder) ([]string, error) {
	var data []string
	task := &pb.TaskORM{}
	query := p.db_main.Model(&task)
	if filter != nil {
		query = query.Where(&filter)
	}
	query = query.Scopes(FilterScoope(sql.Filter), FilterOrScoope(sql.FilterOr, ""), QueryScoop(sql.CollectiveAnd), WhereInScoop(sql.In))
	query = query.Scopes(DistinctScoope(sql.Distinct))
	if err := query.Pluck(columnNameBuilder(key, false), &data).Error; err != nil {
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
		}
	}
	return data, nil
}

func (p *GormProvider) GetListTaskWithFilter(ctx context.Context, task *pb.TaskORM, pagination *pb.PaginationResponse, sort *pb.Sort) (tasks []*pb.TaskORM, err error) {
	res := p.db_main.Where(&task).Scopes(Paginate(tasks, pagination, p.db_main), Sort(sort))
	err = res.Find(&tasks).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "DB Internal Error")
		}
	}
	return tasks, nil
}

func (p *GormProvider) SaveTask(ctx context.Context, task *pb.TaskORM) (*pb.TaskORM, error) {
	if err := p.db_main.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		UpdateAll: true,
	}).Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (p *GormProvider) GetAnnouncementTaskById(ctx context.Context, id uint64) (*pb.AnnouncementTaskORM, error) {
	task := &pb.AnnouncementTaskORM{TaskID: id}
	if err := p.db_main.Find(&task).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}
	return task, nil
}

func (p *GormProvider) GetCompanyTaskById(ctx context.Context, id uint64) (*pb.CompanyTaskORM, error) {
	task := &pb.CompanyTaskORM{TaskID: id}
	if err := p.db_main.Find(&task).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}
	return task, nil
}

func (p *GormProvider) GetAccountTaskById(ctx context.Context, id uint64) (*pb.AccountTaskORM, error) {
	task := &pb.AccountTaskORM{TaskID: id}
	if err := p.db_main.Find(&task).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}
	return task, nil
}

func (p *GormProvider) DeleteCompanyTask(ctx context.Context, listTask []pb.TaskORM) error {
	updateData := map[string]interface{}{
		"status": 6,
		"step":   3,
	}
	listID := []uint64{}
	for _, task := range listTask {
		listID = append(listID, task.TaskID)
	}

	if err := p.db_main.Debug().Model(&pb.TaskORM{}).Where("task_id IN ?", listID).Updates(&updateData).Error; err != nil {
		return status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}
	return nil
}
