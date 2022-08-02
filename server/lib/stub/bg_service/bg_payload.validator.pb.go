// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bg_payload.proto

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
func (this *Company) Validate() error {
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
func (this *ThirdParty) Validate() error {
	return nil
}
func (this *Transaction) Validate() error {
	return nil
}
func (this *TransactionRequest) Validate() error {
	return nil
}
func (this *TaskMappingData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *MappingData) Validate() error {
	return nil
}
func (this *TaskMappingDigitalData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *MappingDigitalData) Validate() error {
	return nil
}
func (this *Beneficiary) Validate() error {
	return nil
}
func (this *IssuingPortal) Validate() error {
	return nil
}
func (this *TaskIssuingData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *IssuingData) Validate() error {
	if this.Publishing != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Publishing); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Publishing", err)
		}
	}
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	if this.Applicant != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Applicant); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Applicant", err)
		}
	}
	if this.Project != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Project); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Project", err)
		}
	}
	if this.Document != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Document); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Document", err)
		}
	}
	return nil
}
func (this *PublishingData) Validate() error {
	return nil
}
func (this *AccountData) Validate() error {
	return nil
}
func (this *ApplicantData) Validate() error {
	return nil
}
func (this *ProjectData) Validate() error {
	if !(this.ProjectAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ProjectAmount))
	}
	if !(this.BgAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("BgAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.BgAmount))
	}
	if !(this.CashAccountAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CashAccountAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.CashAccountAmount))
	}
	if !(this.NonCashAccountAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("NonCashAccountAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.NonCashAccountAmount))
	}
	if !(this.CounterGuaranteeAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CounterGuaranteeAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.CounterGuaranteeAmount))
	}
	if !(this.HoldAccountAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("HoldAccountAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.HoldAccountAmount))
	}
	if !(this.ConsumerLimitAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ConsumerLimitAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ConsumerLimitAmount))
	}
	return nil
}
func (this *DocumentData) Validate() error {
	return nil
}
func (this *ApiPaginationResponse) Validate() error {
	return nil
}
func (this *HealthCheckRequest) Validate() error {
	return nil
}
func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *GetBranchRequest) Validate() error {
	return nil
}
func (this *GetBranchResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *BeneficiaryName) Validate() error {
	return nil
}
func (this *GetBeneficiaryNameRequest) Validate() error {
	return nil
}
func (this *GetBeneficiaryNameResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ApplicantName) Validate() error {
	return nil
}
func (this *GetApplicantNameRequest) Validate() error {
	return nil
}
func (this *GetApplicantNameResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ThirdPartyName) Validate() error {
	return nil
}
func (this *GetThirdPartyRequest) Validate() error {
	return nil
}
func (this *GetThirdPartyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskMappingFilterCompanyRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingFilterCompanyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskMappingFileRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingFileResponse) Validate() error {
	return nil
}
func (this *GetTaskMappingRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingResponse) Validate() error {
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
func (this *GetTaskMappingDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskMappingRequest) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskMappingResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTaskMappingDigitalFileRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingDigitalFileResponse) Validate() error {
	return nil
}
func (this *GetTaskMappingDigitalRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingDigitalResponse) Validate() error {
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
func (this *GetTaskMappingDigitalDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskMappingDigitalDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskMappingDigitalRequest) Validate() error {
	for _, item := range this.Beneficiary {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Beneficiary", err)
			}
		}
	}
	return nil
}
func (this *CreateTaskMappingDigitalResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTransactionAttachmentRequest) Validate() error {
	return nil
}
func (this *GetTransactionAttachmentResponse) Validate() error {
	return nil
}
func (this *GetTransactionFileRequest) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	return nil
}
func (this *GetTransactionFileResponse) Validate() error {
	return nil
}
func (this *GetTransactionRequest) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	return nil
}
func (this *GetTransactionResponse) Validate() error {
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
func (this *GetTransactionDetailRequest) Validate() error {
	return nil
}
func (this *GetTransactionDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTransactionRequest) Validate() error {
	for _, item := range this.MappingData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingData", err)
			}
		}
	}
	for _, item := range this.MappingDataBackup {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDataBackup", err)
			}
		}
	}
	for _, item := range this.MappingDigitalData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDigitalData", err)
			}
		}
	}
	for _, item := range this.MappingDigitalDataBackup {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDigitalDataBackup", err)
			}
		}
	}
	return nil
}
func (this *CreateTransactionResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *DeleteTransactionRequest) Validate() error {
	for _, item := range this.MappingData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingData", err)
			}
		}
	}
	for _, item := range this.MappingDataBackup {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDataBackup", err)
			}
		}
	}
	for _, item := range this.MappingDigitalData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDigitalData", err)
			}
		}
	}
	for _, item := range this.MappingDigitalDataBackup {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MappingDigitalDataBackup", err)
			}
		}
	}
	return nil
}
func (this *DeleteTransactionResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTaskIssuingRequest) Validate() error {
	return nil
}
func (this *GetTaskIssuingResponse) Validate() error {
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
func (this *GetTaskIssuingDetailRequest) Validate() error {
	return nil
}
func (this *GetTaskIssuingDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskIssuingRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateTaskIssuingResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetTaskIssuingFileRequest) Validate() error {
	return nil
}
func (this *GetTaskIssuingFileResponse) Validate() error {
	return nil
}
func (this *CreateIssuingRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateIssuingResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CheckIssuingRequest) Validate() error {
	return nil
}
func (this *CheckIssuingResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
