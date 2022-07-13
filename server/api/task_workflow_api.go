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

	isSave := false

	logrus.Println("[api][SaveTaskWithWorkflow] workflow current step: ", workflow.Workflow.CurrentStep)

	// TODO: set task status basse on next step dari workflow
	switch workflow.Workflow.CurrentStep {
	case "complete":
		if task.Status != pb.Statuses_Approved {
			logrus.Println("[api][SaveTaskWithWorkflow] set task status to approved")
			// process to approve
			isSave = true
			task.Status = pb.Statuses_Approved
		}
	default:
		if task.Status != pb.Statuses_Pending {
			logrus.Println("[api][SaveTaskWithWorkflow] task status is not pending, set to pending")
			// process to pending
			isSave = true
			task.Status = pb.Statuses_Pending
		}
	}

	result := &pb.SaveTaskResponse{
		Success: true,
		Data:    &pb.Task{},
	}

	logrus.Println("[api][SaveTaskWithWorkflow] isSave: ", isSave)
	if isSave {
		logrus.Println("[api][SaveTaskWithWorkflow] task will be save to database")
		taskORM, err := task.ToORM(ctx)
		if err != nil {
			logrus.Errorln("[api][func:SaveTaskWithWorkflow] Error ToORM: ", err)
			return nil, status.Error(codes.Internal, "Server Error")
		}
		logrus.Println("[api][func:SaveTaskWithWorkflow] taskORM ID: ", taskORM)
		logrus.Println("[api][func:SaveTaskWithWorkflow] taskORM Workflow: ", taskORM.WorkflowDoc)
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

	}

	return result, nil
}
