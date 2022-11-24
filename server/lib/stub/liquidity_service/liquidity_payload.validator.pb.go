// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: liquidity_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
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
func (this *ArrayString) Validate() error {
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
func (this *PaginationResponse) Validate() error {
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
func (this *ErrorBodyResponse) Validate() error {
	return nil
}
func (this *ListLiquidityFilter) Validate() error {
	return nil
}
func (this *ListTaskLiquidityRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DownloadListTaskLiquidityRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ListLiquidityTaskResponse) Validate() error {
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
func (this *ListTaskLiquidityRes) Validate() error {
	if this.CompanyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroup", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
		}
	}
	for _, item := range this.Cashflow {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Cashflow", err)
			}
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *CompanyGroupListTask) Validate() error {
	return nil
}
func (this *CompanyListTask) Validate() error {
	return nil
}
func (this *CurrencyListTask) Validate() error {
	return nil
}
func (this *CashflowsListTask) Validate() error {
	for _, item := range this.Source {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
			}
		}
	}
	for _, item := range this.Beneficiary {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Beneficiary", err)
			}
		}
	}
	return nil
}
func (this *TaskListTask) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *DetailLiquidityTaskRequest) Validate() error {
	return nil
}
func (this *DetailLiquidityTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DetailTaskLiquidityRes) Validate() error {
	if this.CompanyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroup", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
		}
	}
	for _, item := range this.Cashflow {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Cashflow", err)
			}
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *CompanyGroupDetailTask) Validate() error {
	return nil
}
func (this *CompanyDetailTask) Validate() error {
	return nil
}
func (this *CurrencyDetailTask) Validate() error {
	return nil
}
func (this *CashflowsDetailTask) Validate() error {
	for _, item := range this.Source {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
			}
		}
	}
	for _, item := range this.Beneficiary {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Beneficiary", err)
			}
		}
	}
	return nil
}
func (this *CreateLiquidityRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskLiquidityRequest) Validate() error {
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.CompanyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroup", err)
		}
	}
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
		}
	}
	for _, item := range this.Cashflow {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Cashflow", err)
			}
		}
	}
	for _, item := range this.PriorityUpdates {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PriorityUpdates", err)
			}
		}
	}
	return nil
}
func (this *LiquiditySourceReq) Validate() error {
	return nil
}
func (this *LiquidityBeneficiaryReq) Validate() error {
	return nil
}
func (this *LiquiditySourceRes) Validate() error {
	return nil
}
func (this *LiquidityBeneficiaryRes) Validate() error {
	return nil
}
func (this *CompanyCreateTask) Validate() error {
	return nil
}
func (this *CompanyGroupCreateTask) Validate() error {
	return nil
}
func (this *CurrencyCreateTask) Validate() error {
	return nil
}
func (this *PriorityUpdatesCreateTask) Validate() error {
	return nil
}
func (this *CashflowsCreateTask) Validate() error {
	for _, item := range this.Source {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Source", err)
			}
		}
	}
	for _, item := range this.Beneficiary {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Beneficiary", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskLiquidityResponse) Validate() error {
	return nil
}
func (this *DeleteTaskLiquidityResponse) Validate() error {
	return nil
}
func (this *DeleteLiquidityTaskRequest) Validate() error {
	return nil
}
func (this *DeleteLiquidityTaskResponse) Validate() error {
	return nil
}
func (this *DeleteLiquidityRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteLiquidityResponse) Validate() error {
	return nil
}
func (this *ListDataRequest) Validate() error {
	if this.Liquidity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Liquidity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Liquidity", err)
		}
	}
	return nil
}
func (this *Pagination) Validate() error {
	return nil
}
func (this *Search) Validate() error {
	return nil
}
func (this *ListDataResponse) Validate() error {
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
func (this *ListTBAValueResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *RunLiquidityTaskRequest) Validate() error {
	return nil
}
func (this *RunLiquidityTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ValidateDateRequest) Validate() error {
	return nil
}
func (this *ValidateDateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AvailableTime) Validate() error {
	return nil
}
func (this *CreateLiquiditySchedulesRequest) Validate() error {
	return nil
}
func (this *CreateLiquiditySchedulesResponse) Validate() error {
	for _, item := range this.Schedules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Schedules", err)
			}
		}
	}
	return nil
}
func (this *CreateLiquiditySchedulesRes) Validate() error {
	if this.ScheduleTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduleTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduleTime", err)
		}
	}
	return nil
}
func (this *RunDailyScheduleResponse) Validate() error {
	return nil
}
func (this *TaskActionRequest) Validate() error {
	return nil
}
func (this *TaskActionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateLiquidityTransactionReq) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.CurrentWorkflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CurrentWorkflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CurrentWorkflow", err)
		}
	}
	return nil
}
func (this *CreateLiquidityTransactionRes) Validate() error {
	return nil
}
func (this *ValidateWorkflowData) Validate() error {
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *WorkflowPayload) Validate() error {
	if this.Header != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Header); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Header", err)
		}
	}
	if this.Records != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Records); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
		}
	}
	if this.CreatedBy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedBy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	return nil
}
func (this *WorkflowHeader) Validate() error {
	return nil
}
func (this *WorkflowRecords) Validate() error {
	if this.LastUpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUpdatedAt", err)
		}
	}
	for _, item := range this.Flows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Flows", err)
			}
		}
	}
	return nil
}
func (this *UserData) Validate() error {
	return nil
}
func (this *Flow) Validate() error {
	if this.Verifier != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Verifier); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Verifier", err)
		}
	}
	if this.Approver != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Approver); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Approver", err)
		}
	}
	if this.Releaser != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Releaser); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Releaser", err)
		}
	}
	if this.CompletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompletedAt", err)
		}
	}
	return nil
}
func (this *Participants) Validate() error {
	for _, item := range this.Participants {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Participants", err)
			}
		}
	}
	return nil
}
func (this *Participant) Validate() error {
	if this.ApprovedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ApprovedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ApprovedAt", err)
		}
	}
	return nil
}
