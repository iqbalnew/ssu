// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *LoginRequest) Validate() error {
	return nil
}
func (this *LoginResponse) Validate() error {
	return nil
}
func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
func (this *InvalidKeyError) Validate() error {
	return nil
}
func (this *GenerateTaskWithWokflowRequest) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *GenerateTaskWithWokflowResponse) Validate() error {
	return nil
}
func (this *SaveTaskRequest) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *SaveTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SetTaskRequest) Validate() error {
	return nil
}
func (this *SetTaskWithWorkflowRequest) Validate() error {
	return nil
}
func (this *ErrorBodyResponse) Validate() error {
	return nil
}
func (this *SetTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListRequest) Validate() error {
	return nil
}
func (this *GetActivityLogsReq) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DownloadActivityLogsReq) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ActivityLogFilter) Validate() error {
	return nil
}
func (this *ActivityLog) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *GetActivityLogsRes) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *GetActivityLogsRes_ActivityLogPagination) Validate() error {
	return nil
}
func (this *ListTaskRequest) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *ListTaskPluckRequest) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *ListTaskPluckResponse) Validate() error {
	return nil
}
func (this *Pagination) Validate() error {
	return nil
}
func (this *PaginationResponse) Validate() error {
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
func (this *Search) Validate() error {
	return nil
}
func (this *ListTaskResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *GraphStatusRequest) Validate() error {
	return nil
}
func (this *GraphStatusColumnTypeRequest) Validate() error {
	return nil
}
func (this *GraphStatusColumnTypeResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GraphStatusColumnType) Validate() error {
	return nil
}
func (this *GraphStatus) Validate() error {
	return nil
}
func (this *GraphStatusResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GraphStepRequest) Validate() error {
	return nil
}
func (this *GraphStep) Validate() error {
	return nil
}
func (this *GraphStepResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetMyPendingTaskWithWorkflowGraphRequest) Validate() error {
	return nil
}
func (this *GraphStepWorkflow) Validate() error {
	return nil
}
func (this *GetMyPendingTaskWithWorkflowGraphResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *AssignaTypeIDRequest) Validate() error {
	return nil
}
func (this *AssignaTypeIDResponse) Validate() error {
	return nil
}
func (this *GetTaskByTypeIDReq) Validate() error {
	return nil
}
func (this *GetTaskByTypeIDRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *RejectBySystemReq) Validate() error {
	return nil
}
func (this *RejectBySystemRes) Validate() error {
	return nil
}
func (this *GetTaskByIDReq) Validate() error {
	return nil
}
func (this *GetTaskByIDRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateTaskDataReq) Validate() error {
	return nil
}
func (this *UpdateTaskDataRes) Validate() error {
	return nil
}
func (this *UpdateTaskWorkflowReq) Validate() error {
	return nil
}
func (this *UpdateTaskWorkflowRes) Validate() error {
	return nil
}
func (this *UpdateTaskRawReq) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *DeleteDraftTaskRequest) Validate() error {
	return nil
}
func (this *DeleteDraftTaskResponse) Validate() error {
	return nil
}
