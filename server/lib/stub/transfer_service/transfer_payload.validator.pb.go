// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InternalTransferData) Validate() error {
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
func (this *TaskInternalTransferData) Validate() error {
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
func (this *GetTaskInternalTransferFileRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferFileResponse) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferResponse) Validate() error {
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
func (this *GetTaskInternalTransferDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SetTaskInternalTransferRequest) Validate() error {
	return nil
}
func (this *SetTaskInternalTransferResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferSingleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferSingleResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferMultipleRequest) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferMultipleResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
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
func (this *CancelInternalTransferTransactionRequest) Validate() error {
	return nil
}
func (this *CancelInternalTransferTransactionResponse) Validate() error {
	return nil
}
func (this *ExecInternalTransferRequest) Validate() error {
	return nil
}
func (this *ExecInternalTransferResponse) Validate() error {
	return nil
}
func (this *ExecFailedInternalTransferRequest) Validate() error {
	return nil
}
func (this *ExecFailedInternalTransferResponse) Validate() error {
	return nil
}
func (this *InternalTransferSingleTemplateData) Validate() error {
	if this.Template != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Template); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Template", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTaskInternalTransferSingleTemplateRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferSingleTemplateResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskInternalTransferSingleTemplateDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskInternalTransferSingleTemplateDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferSingleTemplateRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskInternalTransferSingleTemplateResponse) Validate() error {
	return nil
}
func (this *DeleteTaskInternalTransferSingleTemplateRequest) Validate() error {
	return nil
}
func (this *DeleteTaskInternalTransferSingleTemplateResponse) Validate() error {
	return nil
}
