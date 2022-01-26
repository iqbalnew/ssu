package db

import (
	"context"
	"errors"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *GormProvider) CreateTask(ctx context.Context, task *pb.TaskORM, data string) (*pb.TaskORM, error) {
	if err := p.db_main.Create(&task).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	switch task.Type {
	case "announcement":
		if err := p.db_main.Create(&pb.AnnouncementTaskORM{
			TaskID: task.TaskID,
			Data:   data,
			Type:   task.Type,
		}).Error; err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

	case "company":
		if err := p.db_main.Create(&pb.CompanyTaskORM{
			TaskID: task.TaskID,
			Data:   data,
			Type:   task.Type,
		}).Error; err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

	case "account":
		if err := p.db_main.Create(&pb.AccountTaskORM{
			TaskID: task.TaskID,
			Data:   data,
			Type:   task.Type,
		}).Error; err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}
	}

	return task, nil
}

func (p *GormProvider) UpdateTask(ctx context.Context, task *pb.TaskORM, data string) (*pb.TaskORM, error) {
	taskModel := pb.TaskORM{TaskID: task.TaskID}
	if err := p.db_main.Model(&taskModel).Updates(&task).Error; err != nil {
		logrus.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	if len(data) > 0 {

		switch task.Type {
		case "announcement":
			var model *pb.AnnouncementTaskORM
			if err := p.db_main.Where("task_id = ?", task.TaskID).First(&model).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logrus.Errorln(err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				} else {
					if err := p.db_main.Create(&pb.AnnouncementTaskORM{
						TaskID: task.TaskID,
						Data:   data,
						Type:   task.Type,
					}).Error; err != nil {
						logrus.Errorln(err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}
				}
			} else {
				model.Data = data
				p.db_main.Save(&model)
			}

		case "company":
			var model *pb.CompanyTaskORM
			if err := p.db_main.Where("task_id = ?", task.TaskID).First(&model).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logrus.Errorln(err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				} else {
					if err := p.db_main.Create(&pb.CompanyTaskORM{
						TaskID: task.TaskID,
						Data:   data,
						Type:   task.Type,
					}).Error; err != nil {
						logrus.Errorln(err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}
				}
			} else {
				model.Data = data
				p.db_main.Save(&model)
			}

		case "account":
			var model *pb.AccountTaskORM
			if err := p.db_main.Where("task_id = ?", task.TaskID).First(&model).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logrus.Errorln(err)
					return nil, status.Errorf(codes.Internal, "Internal Error")
				} else {
					if err := p.db_main.Create(&pb.AccountTaskORM{
						TaskID: task.TaskID,
						Data:   data,
						Type:   task.Type,
					}).Error; err != nil {
						logrus.Errorln(err)
						return nil, status.Errorf(codes.Internal, "Internal Error")
					}
				}
			} else {
				model.Data = data
				p.db_main.Save(&model)
			}
		}
	}

	return task, nil
}

func (p *GormProvider) FindTaskById(ctx context.Context, id uint64) (*pb.TaskORM, error) {
	task := &pb.TaskORM{TaskID: id}
	if err := p.db_main.Find(&task).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Task Not Found")
	}
	return task, nil
}

func (p *GormProvider) GetListTask(ctx context.Context) (tasks []*pb.TaskORM, err error) {
	if err := p.db_main.Find(&tasks).Error; err != nil {
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
