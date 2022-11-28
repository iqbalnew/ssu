package api

import (
	"context"
	"encoding/json"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	workflowPb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/stub/workflow_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// rpc GenerateTaskWithWokflow(GenerateTaskWithWokflowRequest) returns (SaveTaskResponse) {};
// rpc SaveTaskWithWorkflow(SaveTaskRequest) returns (SaveTaskResponse) {};

func (s *Server) GenerateTaskWithWokflow(ctx context.Context, req *pb.GenerateTaskWithWokflowRequest) (*pb.SaveTaskResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) SaveTaskWithWorkflow(ctx context.Context, req *pb.SaveTaskRequest) (*pb.SaveTaskResponse, error) {
	if req.Task == nil {
		return nil, status.Error(codes.InvalidArgument, "task is required")
	}

	task := req.Task
	if req.TaskID > 0 {
		task.TaskID = req.TaskID
	}

	if task.WorkflowDoc == "" || task.WorkflowDoc == "{}" {
		return nil, status.Error(codes.InvalidArgument, "workflow doc is required")
	}

	var workflow *workflowPb.ValidateWorkflowData
	err := json.Unmarshal([]byte(task.WorkflowDoc), &workflow)
	if err != nil {
		logrus.Errorln()
	}

	// isSave := false

	// logrus.Println("[api][SaveTaskWithWorkflow] workflow current step: ", workflow.Workflow.CurrentStep)

	// TODO: set task status basse on next step dari workflow
	switch workflow.Workflow.CurrentStep {
	case "complete":
		if task.Status != pb.Statuses_Approved {
			// logrus.Println("[api][SaveTaskWithWorkflow] set task status to approved")
			// process to approve
			// isSave = true
			task.Status = pb.Statuses_Approved
			task.Step = pb.Steps_Releaser
		}
	case "maker":
		// isSave = true
		if task.Type == "Deposito" {
			if task.DataBak != "{}" && workflow.NextStatus == "rejected" {
				task.Status = pb.Statuses_Approved
				task.Step = pb.Steps_Releaser
				//task.Data = task.DataBak
			} else {
				task.Step = pb.Steps_Maker
				if workflow.NextStatus == "rejected" {
					task.Status = pb.Statuses_Rejected
				} else if workflow.NextStatus == "returned" {
					task.Status = pb.Statuses_Returned
				}
			}
		} else {
			task.Step = pb.Steps_Maker
			if workflow.NextStatus == "rejected" {
				task.Status = pb.Statuses_Rejected
			} else if workflow.NextStatus == "returned" {
				task.Status = pb.Statuses_Returned
			}
		}
	default:
		if task.Status != pb.Statuses_Pending {
			// logrus.Println("[api][SaveTaskWithWorkflow] task status is not pending, set to pending")
			// process to pending
			task.Status = pb.Statuses_Pending
		}
		// isSave = true
	}

	logrus.Println("task ===> ", task)

	result := &pb.SaveTaskResponse{
		Success: true,
		Data:    &pb.Task{},
	}

	// logrus.Println("[api][SaveTaskWithWorkflow] isSave: ", isSave)
	// if isSave {
	// logrus.Println("[api][SaveTaskWithWorkflow] task will be save to database")
	taskORM, err := task.ToORM(ctx)
	if err != nil {
		logrus.Errorln("[api][func:SaveTaskWithWorkflow] Error ToORM: ", err)
		return nil, status.Error(codes.Internal, "Server Error")
	}
	// logrus.Println("[api][func:SaveTaskWithWorkflow] taskORM ID: ", taskORM)
	// logrus.Println("[api][func:SaveTaskWithWorkflow] taskORM Workflow: ", taskORM.WorkflowDoc)
	saved, err := s.provider.SaveTask(ctx, &taskORM)
	if err != nil {
		logrus.Errorln("[api][func:SaveTaskWithWorkflow] Error SaveTask: ", err)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	savedTask, err := saved.ToPB(ctx)
	if err != nil {
		logrus.Errorln("[api][func:SaveTaskWithWorkflow] Error ToPB: ", err)
		return nil, status.Error(codes.Internal, "Server Error")
	}

	result.Success = true
	result.Data = &savedTask

	// }

	if getEnv("ENV", "LOCAL") != "LOCAL" {
		action := "save"
		if req.IsDraft {
			action = "draft"
		}
		logrus.Println("Set to save log")
		activityLog, err := GenerateActivityLog(&taskORM, nil, action, "Update")
		if err != nil {
			logrus.Errorln("[api][func: SaveTaskWithData] Failed to generate activity log: %v", err)
			return nil, err
		}
		err = s.provider.SaveLog(ctx, activityLog)
		if err != nil {
			logrus.Errorln("Error SaveActivityLog: ", err)
		}
	}

	return result, nil
}

func (s *Server) UpdateTaskWorkflow(ctx context.Context, req *pb.UpdateTaskWorkflowReq) (*pb.UpdateTaskWorkflowRes, error) {
	res := &pb.UpdateTaskWorkflowRes{
		Success: false,
	}

	task, err := s.provider.FindTaskById(ctx, req.TaskID)
	if err != nil {
		return nil, err
	}

	if task.Type != req.Type {
		return res, nil
	}

	task.WorkflowDoc = req.Workflow
	_, err = s.provider.UpdateTask(ctx, task, false)
	if err != nil {
		return nil, err
	}

	res.Success = true
	return res, nil
}
