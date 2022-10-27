// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: abonnement_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ErrorBodyResponse) Validate() error {
	return nil
}
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
func (this *JWTTokenResponse) Validate() error {
	return nil
}
func (this *ListRequest) Validate() error {
	return nil
}
func (this *Sort) Validate() error {
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
func (this *PaginationResponse) Validate() error {
	return nil
}
func (this *ListAbonnementRequest) Validate() error {
	if this.Abonnement != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Abonnement); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Abonnement", err)
		}
	}
	return nil
}
func (this *ListAbonnementResponse) Validate() error {
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
func (this *CreateAbonnementRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAbonnementTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAbonnementResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *AbonnementTask) Validate() error {
	if this.Abonnement != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Abonnement); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Abonnement", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *CreateAbonnementTaskResponse) Validate() error {
	return nil
}
func (this *ListFilterAbannoment) Validate() error {
	return nil
}
func (this *ListAbonnementTaskRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ListAbonnementTaskResponse) Validate() error {
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
func (this *DownloadListAbonnementTaskRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *AbonnementTaskDetailRequest) Validate() error {
	return nil
}
func (this *AbonnementTaskDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateAbonnementTaskReq) Validate() error {
	for _, item := range this.MechanismTransaction {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismTransaction", err)
			}
		}
	}
	for _, item := range this.MechanismFrequency {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismFrequency", err)
			}
		}
	}
	for _, item := range this.MechanismBalance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismBalance", err)
			}
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.AbonnementAccounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementAccounts", err)
			}
		}
	}
	for _, item := range this.AbonnementRetries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementRetries", err)
			}
		}
	}
	for _, item := range this.AbonnementInvoices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementInvoices", err)
			}
		}
	}
	for _, item := range this.CustomAbonnements {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CustomAbonnements", err)
			}
		}
	}
	return nil
}
func (this *DetailTaskAbonnementRes) Validate() error {
	if this.DeadlineDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeadlineDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeadlineDate", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.AbonnementAccounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementAccounts", err)
			}
		}
	}
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
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *Mechanism) Validate() error {
	return nil
}
func (this *ListTaskAbonnementRes) Validate() error {
	for _, item := range this.MechanismTransaction {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismTransaction", err)
			}
		}
	}
	for _, item := range this.MechanismFrequency {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismFrequency", err)
			}
		}
	}
	for _, item := range this.MechanismBalance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MechanismBalance", err)
			}
		}
	}
	if this.Mechanism != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mechanism); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mechanism", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.AbonnementAccounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementAccounts", err)
			}
		}
	}
	for _, item := range this.AbonnementRetries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementRetries", err)
			}
		}
	}
	for _, item := range this.AbonnementInvoices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AbonnementInvoices", err)
			}
		}
	}
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
func (this *AbonnmentAccountRes) Validate() error {
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
func (this *AbonnementAccountCreateTask) Validate() error {
	return nil
}
func (this *CustomAbonnementCreateTask) Validate() error {
	return nil
}
func (this *CompanyCreateTask) Validate() error {
	return nil
}
func (this *CompanyListTask) Validate() error {
	return nil
}
func (this *ListFilterAbannomentInvoice) Validate() error {
	return nil
}
func (this *ListAbonnementInvoiceRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DownloadListAbonnementInvoiceRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *ListAbonnementInvoiceResponse) Validate() error {
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
func (this *AbonnementInvoiceDetailRequest) Validate() error {
	return nil
}
func (this *CekAvaibilityReq) Validate() error {
	return nil
}
func (this *CekAvaibilityRes) Validate() error {
	return nil
}
func (this *ChargeCompaniesReq) Validate() error {
	return nil
}
func (this *ChargeCompaniesRes) Validate() error {
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
func (this *ChargeCompanyResponse) Validate() error {
	return nil
}
func (this *ListPendingAbonnementInvoiceRequest) Validate() error {
	return nil
}
func (this *ListPendingAbonnementInvoiceResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateAbonnementInvoiceRequest) Validate() error {
	if this.DeadlineDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeadlineDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeadlineDate", err)
		}
	}
	return nil
}
func (this *CreateAbonnementInvoiceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateAbonnementInvoiceRequest) Validate() error {
	return nil
}
func (this *UpdateAbonnementInvoiceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GenerateBulkInvoiceReq) Validate() error {
	return nil
}
func (this *GenerateBulkInvoiceRes) Validate() error {
	for _, item := range this.SuccessData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SuccessData", err)
			}
		}
	}
	for _, item := range this.FailedData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FailedData", err)
			}
		}
	}
	return nil
}
func (this *GenerateSingleInvoiceReq) Validate() error {
	return nil
}
func (this *GenerateSingleInvoiceRes) Validate() error {
	if this.SuccessData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SuccessData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SuccessData", err)
		}
	}
	if this.FailedData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FailedData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FailedData", err)
		}
	}
	return nil
}
func (this *ChargeRes) Validate() error {
	for _, item := range this.SuccessData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SuccessData", err)
			}
		}
	}
	for _, item := range this.FailedData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FailedData", err)
			}
		}
	}
	return nil
}
func (this *ArrayString) Validate() error {
	return nil
}
func (this *ChargeCompanyRequest) Validate() error {
	return nil
}
func (this *BulkInvoiceRequest) Validate() error {
	return nil
}
