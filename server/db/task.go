package db

import (
	"context"
	"errors"
	"fmt"

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

func (p *GormProvider) GetGraphStep(ctx context.Context, service string, step uint, stat uint, isIncludeApprove bool, isIncludeReject bool) (result []*GraphResult, err error) {
	selectOpt := fmt.Sprintf("step as name, type, count(*) as total")
	query := p.db_main.Model(&pb.TaskORM{}).Select(selectOpt)
	whereOpt := ""
	if service != "" {
		whereOpt = fmt.Sprintf("type = '%v'", service)
	}
	if !isIncludeApprove {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + "status != 3"
	}
	if !isIncludeReject {
		if whereOpt != "" {
			whereOpt = whereOpt + " AND "
		}
		whereOpt = whereOpt + "status != 4"
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
	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group("step, type")

	if err = query.Find(&result).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	return result, nil
}

func (p *GormProvider) GetGraphServiceType(ctx context.Context, service string, stat uint, filter string) (result []*GraphResultColumnType, err error) {
	if filter == "" {
		return nil, status.Errorf(codes.InvalidArgument, "column cant be empty")
	}

	column := columnNameBuilder(filter)
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
	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group(fmt.Sprintf("%s, type", column))

	if err = query.Find(&result).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	return result, nil
}

func (p *GormProvider) GetGraphStatus(ctx context.Context, service string, stat uint) (result []*GraphResult, err error) {
	selectOpt := fmt.Sprintf("status as name, type, count(*) as total")
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
	if whereOpt != "" {
		query = query.Where(whereOpt)
	}

	query = query.Group("status, type")

	if err = query.Find(&result).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}
	return result, nil
}

func (p *GormProvider) CreateTask(ctx context.Context, task *pb.TaskORM) (*pb.TaskORM, error) {
	query := p.db_main.Clauses(clause.OnConflict{
		UpdateAll: true,
	})
	if err := query.Create(&task).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
	}

	return task, nil
}

func (p *GormProvider) UpdateTask(ctx context.Context, task *pb.TaskORM) (*pb.TaskORM, error) {
	taskModel := pb.TaskORM{TaskID: task.TaskID}
	query := p.db_main.Clauses(clause.OnConflict{
		UpdateAll: true,
	})
	if err := query.Model(&taskModel).Updates(&task).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		} else {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.NotFound, "Task Not Found")
		}
	}

	return task, nil
}

func (p *GormProvider) FindTaskById(ctx context.Context, id uint64) (*pb.TaskORM, error) {
	task := &pb.TaskORM{TaskID: id}
	if err := p.db_main.Preload(clause.Associations).Find(&task).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}
	return task, nil
}

func (p *GormProvider) GetListTask(ctx context.Context, filter *pb.TaskORM, f string, q string, pagination *pb.PaginationResponse, sort *pb.Sort) (tasks []*pb.TaskORM, err error) {
	query := p.db_main
	if filter != nil {
		query = query.Where(&filter)
	}
	query = query.Scopes(FilterScoope(f), QueryScoop(q))
	query = query.Scopes(Paginate(tasks, pagination, query), Sort(sort))
	result := query.Preload(clause.Associations).Find(&tasks)
	if err := result.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error: %v", err)
		}
	}
	return tasks, nil
}

func (p *GormProvider) GetListTaskWithFilter(ctx context.Context, task *pb.TaskORM, pagination *pb.PaginationResponse, sort *pb.Sort) (tasks []*pb.TaskORM, err error) {
	res := p.db_main.Debug().Where(&task).Scopes(Paginate(tasks, pagination, p.db_main), Sort(sort))
	err = res.Find(&tasks).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}
	return tasks, nil
}

func (p *GormProvider) SaveTask(ctx context.Context, task *pb.TaskORM) (*pb.TaskORM, error) {
	if err := p.db_main.Save(&task).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
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
