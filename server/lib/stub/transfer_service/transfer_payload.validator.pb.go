// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Participant) Validate() error {
	if this.ApprovedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ApprovedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ApprovedAt", err)
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
func (this *WorkflowHeader) Validate() error {
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
func (this *ValidateWorkflowData) Validate() error {
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *Task) Validate() error {
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
func (this *HealthCheckRequest) Validate() error {
	return nil
}
func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *GetPairRateRequest) Validate() error {
	return nil
}
func (this *GetPairRateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Rate) Validate() error {
	return nil
}
func (this *InternalMultipleReceiverData) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.ExchangeRate >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ExchangeRate", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ExchangeRate))
	}
	if !(this.ReceivedAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ReceivedAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ReceivedAmount))
	}
	return nil
}
func (this *InternalBulkDataRequest) Validate() error {
	for _, item := range this.Receivers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Receivers", err)
			}
		}
	}
	return nil
}
func (this *InternalBulkReceiverData) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.ExchangeRate >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ExchangeRate", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ExchangeRate))
	}
	if !(this.ReceivedAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ReceivedAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ReceivedAmount))
	}
	return nil
}
func (this *InternalBulkData) Validate() error {
	return nil
}
func (this *TaskInternalBulkReceiverData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTaskInternalFileRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalFileResponse) Validate() error {
	return nil
}
func (this *GetTaskInternalRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalResponse) Validate() error {
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
func (this *GetTaskInternalDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SetTaskInternalRequest) Validate() error {
	return nil
}
func (this *SetTaskInternalResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateInternalTransferTransactionRequest) Validate() error {
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
func (this *CreateInternalTransferTransactionResponse) Validate() error {
	return nil
}
func (this *ExecInternalTransferRequest) Validate() error {
	return nil
}
func (this *ExecInternalTransferResponse) Validate() error {
	return nil
}
func (this *CancelInternalTransferTransactionRequest) Validate() error {
	return nil
}
func (this *CancelInternalTransferTransactionResponse) Validate() error {
	return nil
}
func (this *CreateMassInquiryRequest) Validate() error {
	return nil
}
func (this *CreateMassInquiryResponse) Validate() error {
	return nil
}
func (this *CreateMassTransferRequest) Validate() error {
	return nil
}
func (this *CreateMassTransferResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateMassTransferResponseData) Validate() error {
	return nil
}
func (this *TaskInternalSingleData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *InternalSingleData) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.ExchangeRate >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ExchangeRate", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ExchangeRate))
	}
	if !(this.ReceivedAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ReceivedAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ReceivedAmount))
	}
	if this.RecurringPeriodStart != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecurringPeriodStart); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecurringPeriodStart", err)
		}
	}
	if this.RecurringPeriodEnd != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecurringPeriodEnd); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecurringPeriodEnd", err)
		}
	}
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalSingleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalSingleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *InternalMultipleData) Validate() error {
	for _, item := range this.Receivers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Receivers", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskInternalMultipleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalMultipleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PayrollDataJob) Validate() error {
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	if this.InquiryResult != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.InquiryResult); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("InquiryResult", err)
		}
	}
	if this.TransferResult != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TransferResult); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TransferResult", err)
		}
	}
	return nil
}
func (this *PayrollData) Validate() error {
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	return nil
}
func (this *PayrollDataDetail) Validate() error {
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	for _, item := range this.Validation {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Validation", err)
			}
		}
	}
	return nil
}
func (this *PayrollDataList) Validate() error {
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	return nil
}
func (this *PayrollItemDetail) Validate() error {
	return nil
}
func (this *PayrollItem) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.Fee >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Fee", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Fee))
	}
	return nil
}
func (this *PayrollItemValidation) Validate() error {
	return nil
}
func (this *TaskPayrollData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *TaskPayrollDataList) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *DecodePayrollFileRequest) Validate() error {
	return nil
}
func (this *DecodePayrollFileResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DecodePayrollData) Validate() error {
	for _, item := range this.Rows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rows", err)
			}
		}
	}
	return nil
}
func (this *DecodePayrollRow) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	return nil
}
func (this *GetTaskPayrollFileRequest) Validate() error {
	return nil
}
func (this *GetTaskPayrollFileResponse) Validate() error {
	return nil
}
func (this *GetTaskPayrollRequest) Validate() error {
	return nil
}
func (this *GetTaskPayrollResponse) Validate() error {
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
func (this *GetTaskPayrollDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskPayrollDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *CreateTaskPayrollRequest) Validate() error {
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	return nil
}
func (this *CreateTaskPayrollResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskPayrollResponseData) Validate() error {
	return nil
}
func (this *CancelTransferPayrollRequest) Validate() error {
	return nil
}
func (this *CancelTransferPayrollResponse) Validate() error {
	return nil
}
func (this *DecodeBulkFileRequest) Validate() error {
	return nil
}
func (this *DecodeBulkFileResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DecodeBulkData) Validate() error {
	for _, item := range this.Rows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rows", err)
			}
		}
	}
	return nil
}
func (this *DecodeBulkRow) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.Fee >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Fee", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Fee))
	}
	return nil
}
func (this *GetTaskInternalBulkDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalBulkDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *TaskInternalBulkDetailData) Validate() error {
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	for _, item := range this.Receivers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Receivers", err)
			}
		}
	}
	return nil
}
func (this *GetTaskInternalBulkRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalBulkResponse) Validate() error {
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
func (this *TaskInternalBulkData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalBulkRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalBulkResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTaskInternalSingleTemplateRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalSingleTemplateResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskInternalSingleTemplateDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalSingleTemplateDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalSingleTemplateRequest) Validate() error {
	return nil
}
func (this *CreateTaskInternalSingleTemplateResponse) Validate() error {
	return nil
}
func (this *DeleteTaskInternalSingleTemplateRequest) Validate() error {
	return nil
}
func (this *DeleteTaskInternalSingleTemplateResponse) Validate() error {
	return nil
}
func (this *RunTransferJobRequest) Validate() error {
	return nil
}
func (this *RunTransferJobResponse) Validate() error {
	return nil
}
func (this *RunMassInquiryJobRequest) Validate() error {
	return nil
}
func (this *RunMassInquiryJobResponse) Validate() error {
	return nil
}
func (this *RunMassTransferJobRequest) Validate() error {
	return nil
}
func (this *RunMassTransferJobResponse) Validate() error {
	return nil
}
func (this *RunMassTransferScheduledJobRequest) Validate() error {
	return nil
}
func (this *RunMassTransferScheduledJobResponse) Validate() error {
	return nil
}
func (this *SetTaskPayrollRequest) Validate() error {
	return nil
}
func (this *SetTaskPayrollResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaskExternalTransferData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *ExternalTransferData) Validate() error {
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	if this.Beneficiary != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Beneficiary); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Beneficiary", err)
		}
	}
	if !(this.FeeAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("FeeAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.FeeAmount))
	}
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
		}
	}
	if this.RecurringPeriodStart != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecurringPeriodStart); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecurringPeriodStart", err)
		}
	}
	if this.RecurringPeriodEnd != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RecurringPeriodEnd); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RecurringPeriodEnd", err)
		}
	}
	return nil
}
func (this *SenderData) Validate() error {
	if !(this.AccountBalance >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountBalance", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.AccountBalance))
	}
	return nil
}
func (this *BeneficiaryData) Validate() error {
	if !(this.AccountBalance >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountBalance", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.AccountBalance))
	}
	return nil
}
func (this *GetTaskExternalTransferRequest) Validate() error {
	return nil
}
func (this *GetTaskExternalTransferResponse) Validate() error {
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
func (this *GetTaskExternalTransferDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskExternalTransferDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskExternalTransferSingleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskExternalTransferSingleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskExternalTransferMultipleRequest) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskExternalTransferMultipleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *InquiryResult) Validate() error {
	return nil
}
func (this *InquiryResultItem) Validate() error {
	return nil
}
func (this *TransferResult) Validate() error {
	return nil
}
func (this *TransferResultItem) Validate() error {
	return nil
}
