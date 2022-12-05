// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: online_transfer_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OnlineTransferData) Validate() error {
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
func (this *TaskOnlineTransferData) Validate() error {
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
func (this *GetTaskOnlineTransferFileRequest) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferFileResponse) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferRequest) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferResponse) Validate() error {
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
func (this *GetTaskOnlineTransferDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SetTaskOnlineTransferRequest) Validate() error {
	return nil
}
func (this *SetTaskOnlineTransferResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskOnlineTransferSingleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskOnlineTransferSingleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskOnlineTransferMultipleRequest) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskOnlineTransferMultipleResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateOnlineTransferTransactionRequest) Validate() error {
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
func (this *CreateOnlineTransferTransactionResponse) Validate() error {
	return nil
}
func (this *CancelOnlineTransferTransactionRequest) Validate() error {
	return nil
}
func (this *CancelOnlineTransferTransactionResponse) Validate() error {
	return nil
}
func (this *ExecOnlineTransferRequest) Validate() error {
	return nil
}
func (this *ExecOnlineTransferResponse) Validate() error {
	return nil
}
func (this *FundTransferToOtherBankRequest) Validate() error {
	return nil
}
func (this *FundTransferToOtherBankResponse) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferSingleTemplateRequest) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferSingleTemplateResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskOnlineTransferSingleTemplateDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskOnlineTransferSingleTemplateDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskOnlineTransferSingleTemplateRequest) Validate() error {
	return nil
}
func (this *CreateTaskOnlineTransferSingleTemplateResponse) Validate() error {
	return nil
}
func (this *DeleteTaskOnlineTransferSingleTemplateRequest) Validate() error {
	return nil
}
func (this *DeleteTaskOnlineTransferSingleTemplateResponse) Validate() error {
	return nil
}
