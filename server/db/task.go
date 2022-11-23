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
	Name   string
	Type   string
	Status int
	Total  uint64
}

func (p *GormProvider) GetGraphStepAll(ctx context.Context, idCompany string) (result *GraphResult, err error) {
	selectOpt := fmt.Sprintf("count(*) as total")
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

func (p *GormProvider) GetGraphPendingTaskWithWorkflow(ctx context.Context, service string, roleids []uint64, accountids []uint64, isMaker bool, createdByID uint64) (result []*GraphResultWorkflowType, err error) {

	if len(roleids) < 1 {
		return []*GraphResultWorkflowType{}, nil
	}

	selectOpt := `workflow_doc->'workflow'->>'currentStep' as name, "type", "status", count(*) as total`

	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select(selectOpt)

	whereOpt := ""

	if isMaker {

		whereOpt = ""

		statusFilter := []string{
			"status != 0",
			"status != 1",
			"status != 4",
			"status != 5",
			"status != 6",
			"status != 7",
		}

		whereOpt = strings.Join(statusFilter, " AND ")

	} else {

		whereOpt = "workflow_doc != '{}'"

		statusFilter := []string{
			"status != 0",
			"status != 2",
			"status != 3",
			"status != 4",
			"status != 5",
			"status != 7",
		}

		whereOpt = whereOpt + " AND " + strings.Join(statusFilter, " AND ")

	}

	if service != "" {

		whereOpt = fmt.Sprintf("%s AND type = '%v'", whereOpt, service)

	}

	if isMaker {

		whereOpt = fmt.Sprintf("%s AND created_by_id = '%d'", whereOpt, createdByID)

	} else {

		roleidstring := ""
		for i, roleid := range roleids {
			if i > 0 {
				roleidstring = fmt.Sprintf("%s,%d", roleidstring, roleid)
			} else {
				roleidstring = fmt.Sprintf("%s%d", roleidstring, roleid)
			}
		}
		roleidstring = "[" + roleidstring + "]"
		whereOpt = fmt.Sprintf("%s AND TRANSLATE(workflow_doc->'workflow'->>'currentRoleIDs', '[]','{}')::INT[] && ARRAY%s", whereOpt, roleidstring)

		accountidstring := ""
		for i, accountid := range accountids {
			if i > 0 {
				accountidstring = fmt.Sprintf("%s,%d", accountidstring, accountid)
			} else {
				accountidstring = fmt.Sprintf("%s%d", accountidstring, accountid)
			}
		}
		whereOpt = fmt.Sprintf("%s AND (workflow_doc->'workflow'->'header'->'uaID' in (%s))", whereOpt, accountidstring)

		if createdByID > 0 {
			whereOpt = fmt.Sprintf("%s AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR workflow_doc->'workflow'->>'participantUserIDs' NOT LIKE '%s')", whereOpt, "%"+fmt.Sprint(createdByID)+"%")
		}

	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group("workflow_doc->'workflow'->>'currentStep', type, status")

	if err = query.Find(&result).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
	}

	return result, nil
}

func (p *GormProvider) GetGraphStep(ctx context.Context, idCompany string, service string, step uint, stat uint, isIncludeApprove bool, isIncludeReject bool, userType string) (result []*GraphResult, err error) {
	selectOpt := "step as name, type, count(*) as total"
	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	if service != "" {
		whereOpt = fmt.Sprintf("type = '%v'", service)
	}
	if idCompany != "" {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + `( ("data" -> 'user'->> 'companyID' = '` + idCompany + `' OR "data" -> 'companyID' = '` + idCompany + `' OR "data" -> 'company' ->> 'companyID' = '` + idCompany + `'
		OR  "data" @> '[{"companyID":` + idCompany + `}]') 
		or ("type" IN ('BG Issuing', 'Internal Fund Transfer', 'External Fund Transfer', 'Payroll Transfer')  AND  company_id = '` + idCompany + `'))`
	}

	if userType == "ca" {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + ` ("data"->'userTypeID' = '4' OR type != 'Role') AND (data->'user'->>'userTypeName' = 'Customer User' OR type != 'User') `
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

func (p *GormProvider) GetListTask(ctx context.Context, filter *pb.TaskORM, pagination *pb.PaginationResponse, sql *QueryBuilder, workflowUserIDFilter uint64, workflowRoleIDFilter []uint64, workflowAccountIDFilter []uint64) (tasks []*pb.TaskORM, err error) {

	query := p.db_main
	if filter.Type != "" {
		query = query.Debug()
	}

	query = query.Select("*", " CASE WHEN status = '3' or status = '5' THEN last_rejected_by_name ELSE last_approved_by_name END AS reviewed_by").Where("status != 7")
	if filter != nil {
		query = query.Where(&filter)
	}

	customQuery := ""

	if len(workflowRoleIDFilter) > 0 {
		value := strings.ReplaceAll(fmt.Sprint(workflowRoleIDFilter), " ", "','")
		value = strings.ReplaceAll(value, "[", "'")
		value = strings.ReplaceAll(value, "]", "'")
		customQuery = fmt.Sprintf("array(select jsonb_array_elements_text(workflow_doc->'workflow'->'currentRoleIDs')) && array[%s]", value)
	}

	if len(workflowAccountIDFilter) > 0 {
		valueAccount := strings.ReplaceAll(fmt.Sprint(workflowAccountIDFilter), " ", "','")
		valueAccount = strings.ReplaceAll(valueAccount, "[", "'")
		valueAccount = strings.ReplaceAll(valueAccount, "]", "'")
		if customQuery == "" {
			customQuery = fmt.Sprintf("workflow_doc->'workflow'->'header'->'uaID' in (%s)", valueAccount)
		} else {
			customQuery = customQuery + " AND (workflow_doc->'workflow'->'header'->'uaID' in (" + valueAccount + "))"
		}
	}

	if workflowUserIDFilter > 0 {
		if customQuery == "" {
			customQuery = "(workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR workflow_doc->'workflow'->>'participantUserIDs' NOT LIKE '%" + fmt.Sprint(workflowUserIDFilter) + "%') and workflow_doc != '{}' "
		} else {
			customQuery = customQuery + " AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR workflow_doc->'workflow'->>'participantUserIDs' NOT LIKE '%" + fmt.Sprint(workflowUserIDFilter) + "%') and workflow_doc != '{}' "
		}

	}

	logrus.Println("Custom Query list: ==> %s", customQuery)
	query = query.Scopes(FilterScoope(sql.Filter))
	query = query.Scopes(FilterOrScoope(sql.FilterOr, customQuery))

	query = query.Scopes(QueryScoop(sql.CollectiveAnd), WhereInScoop(sql.In), WhereInScoop(sql.MeFilterIn), NotConditionalScoope(sql.FilterNot))
	if sql.CompanyID != "" {
		query = query.Where(`("company_id" = ? OR "data" ->> 'companyID' = ?)`, sql.CompanyID, sql.CompanyID)
	}
	query = query.Scopes(DistinctScoope(sql.Distinct))
	query = query.Scopes(Paginate(tasks, pagination, query), CustomOrderScoop(sql.CustomOrder), Sort(sql.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))
	logrus.Println("Query list: ==> %s", query)
	if err := query.Preload(clause.Associations).Debug().Find(&tasks).Error; err != nil {
		if !errors.Is(err, gorm.ErrModelValueRequired) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
	}

	return tasks, nil

}

func (p *GormProvider) GetListTaskNormal(ctx context.Context, filter *pb.TaskORM, pagination *pb.PaginationResponse, sql *QueryBuilder, workflowUserIDFilter uint64, workflowRoleIDFilter []uint64, workflowAccountIDFilter []uint64) (tasks []*pb.TaskORM, err error) {

	query := p.db_main.Debug().Model(&pb.TaskORM{}).Select("*", " CASE WHEN status = '3' or status = '5' THEN last_rejected_by_name ELSE last_approved_by_name END AS reviewed_by")

	whereOpt := ""

	if filter.Step == int32(pb.Steps_Maker) {

		whereOpt = ""

		statusFilter := []string{
			"status != 0",
			"status != 1",
			"status != 4",
			"status != 5",
			"status != 6",
			"status != 7",
		}

		whereOpt = strings.Join(statusFilter, " AND ")

	} else {

		whereOpt = "workflow_doc != '{}'"

		statusFilter := []string{
			"status != 0",
			"status != 2",
			"status != 3",
			"status != 4",
			"status != 5",
			"status != 7",
		}

		whereOpt = whereOpt + " AND " + strings.Join(statusFilter, " AND ")

	}

	if filter.Type != "" {

		whereOpt = fmt.Sprintf("%s AND type = '%v'", whereOpt, filter.Type)

	}

	if filter.Step == int32(pb.Steps_Maker) {

		whereOpt = fmt.Sprintf("%s AND created_by_id = '%d'", whereOpt, workflowUserIDFilter)

	} else {

		roleidstring := ""
		for i, roleid := range workflowRoleIDFilter {
			if i > 0 {
				roleidstring = fmt.Sprintf("%s,%d", roleidstring, roleid)
			} else {
				roleidstring = fmt.Sprintf("%s%d", roleidstring, roleid)
			}
		}
		roleidstring = "[" + roleidstring + "]"
		whereOpt = fmt.Sprintf("%s AND TRANSLATE(workflow_doc->'workflow'->>'currentRoleIDs', '[]','{}')::INT[] && ARRAY%s", whereOpt, roleidstring)

		if workflowUserIDFilter > 0 {
			whereOpt = fmt.Sprintf("%s AND (workflow_doc->'workflow'->>'participantUserIDs' IS NULL OR workflow_doc->'workflow'->>'participantUserIDs' NOT LIKE '%s')", whereOpt, "%"+fmt.Sprint(workflowUserIDFilter)+"%")
		}

		if filter.Step == int32(pb.Steps_Checker) {
			whereOpt = fmt.Sprintf("%s AND workflow_doc->'workflow'->>'currentStep' = 'checker'", whereOpt)
		} else if filter.Step == int32(pb.Steps_Signer) {
			whereOpt = fmt.Sprintf("%s AND workflow_doc->'workflow'->>'currentStep' = 'signer'", whereOpt)
		} else if filter.Step == int32(pb.Steps_Releaser) {
			whereOpt = fmt.Sprintf("%s AND workflow_doc->'workflow'->>'currentStep' = 'releaser'", whereOpt)
		}

	}

	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Scopes(Paginate(tasks, pagination, query), CustomOrderScoop(sql.CustomOrder), Sort(sql.Sort), Sort(&pb.Sort{Column: "updated_at", Direction: "DESC"}))

	if err = query.Preload(clause.Associations).Debug().Find(&tasks).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "DB Internal Error: %v", err)
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
