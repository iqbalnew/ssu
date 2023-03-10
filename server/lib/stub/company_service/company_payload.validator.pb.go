// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: company_payload.proto

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
func (this *DetailCurrencyRequest) Validate() error {
	return nil
}
func (this *DetailCurrencyResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListCompanyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CompanyResponse) Validate() error {
	return nil
}
func (this *ListCurrencyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ListCompanyDataReq) Validate() error {
	return nil
}
func (this *ListGroupLimitReq) Validate() error {
	return nil
}
func (this *ListLimitReq) Validate() error {
	return nil
}
func (this *ListCurrencyReq) Validate() error {
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
func (this *ListGroupLimitRes) Validate() error {
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
func (this *ListCompanyDataRes) Validate() error {
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
func (this *ListLimitRes) Validate() error {
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
func (this *ListCurrencyRes) Validate() error {
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
func (this *CompanyGroupDetailRequest) Validate() error {
	return nil
}
func (this *CompanyGroupDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CompanyGroupTask) Validate() error {
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
func (this *CreateCompanyGroupTask) Validate() error {
	if this.CompanyGroupRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroupRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroupRequest", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *ListCompanyGroupTaskRequest) Validate() error {
	return nil
}
func (this *PaginationResponse) Validate() error {
	return nil
}
func (this *ListCompanyGroupTaskResponse) Validate() error {
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
func (this *CreateCompanyGroupTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCompanyGroupRequest) Validate() error {
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	for _, item := range this.GroupLimit {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GroupLimit", err)
			}
		}
	}
	for _, item := range this.CompanyLimit {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CompanyLimit", err)
			}
		}
	}
	for _, item := range this.GroupSubsidiaries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GroupSubsidiaries", err)
			}
		}
	}
	if this.CompanyWorkflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyWorkflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyWorkflow", err)
		}
	}
	return nil
}
func (this *CreateCompanyLimitRequest) Validate() error {
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
		}
	}
	return nil
}
func (this *CreateGroupLimitRequest) Validate() error {
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
		}
	}
	return nil
}
func (this *CurrencyRequest) Validate() error {
	return nil
}
func (this *CreateCompanyWorkflowRequest) Validate() error {
	return nil
}
func (this *CreateGroupSubsdiaryRequest) Validate() error {
	return nil
}
func (this *CreateCompanyGroupResponse) Validate() error {
	return nil
}
func (this *BricamsGetCustomerReq) Validate() error {
	return nil
}
func (this *BricamsGetCustomerGroupReq) Validate() error {
	return nil
}
func (this *BricamsGetCustomerByIdReq) Validate() error {
	return nil
}
func (this *BricamsGetCustomerByUserReq) Validate() error {
	return nil
}
func (this *BricamsGetCustomerResMulti) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Detail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Detail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Detail", err)
		}
	}
	return nil
}
func (this *BricamsGetCustomerResMulti_Detail) Validate() error {
	return nil
}
func (this *BricamsGetCustomerResSingel) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Detail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Detail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Detail", err)
		}
	}
	return nil
}
func (this *BricamsGetCustomerResSingel_Detail) Validate() error {
	return nil
}
func (this *BricamsResponseDatas) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *BricamsResponseData) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *BricamsGetCustomerData) Validate() error {
	return nil
}
func (this *WorkflowStep) Validate() error {
	return nil
}
func (this *Workflow) Validate() error {
	if this.TansactionalStpStep != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TansactionalStpStep); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TansactionalStpStep", err)
		}
	}
	if this.NonTansactionalStpStep != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NonTansactionalStpStep); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NonTansactionalStpStep", err)
		}
	}
	return nil
}
func (this *CreateCompanyTaskReq) Validate() error {
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *CompanyParams) Validate() error {
	return nil
}
func (this *CreateCompanyTask) Validate() error {
	return nil
}
func (this *CreateCompanyTaskRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DeleteCompanyTaskReq) Validate() error {
	return nil
}
func (this *DeleteCompanyTaskRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetCompanyTaskByHoldingIDReq) Validate() error {
	return nil
}
func (this *GetCompanyTaskByIDReq) Validate() error {
	return nil
}
func (this *GetCompanyTaskByIDRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ListCompanyTaskReq) Validate() error {
	return nil
}
func (this *GetCompanyTaskByIDRequest) Validate() error {
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
func (this *CompanyTask) Validate() error {
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *ListCompanyTaskRes) Validate() error {
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
func (this *CreateCompanyReq) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateCompanyRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *BRICaMSSvcCompanyV2) Validate() error {
	if this.CompanyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroup", err)
		}
	}
	return nil
}
func (this *GetGroupIDStoreReq) Validate() error {
	return nil
}
func (this *GetGroupIDStoreRes) Validate() error {
	return nil
}
func (this *DeleteCompanyRes) Validate() error {
	return nil
}
func (this *DetailCompanyRes) Validate() error {
	return nil
}
func (this *CekCompanyIDAvaibilityReq) Validate() error {
	return nil
}
func (this *CompanySubsidiaryValidationReq) Validate() error {
	return nil
}
func (this *CompanySubsidiaryValidationRes) Validate() error {
	return nil
}
func (this *CheckCompanyIDTaskAndData) Validate() error {
	return nil
}
func (this *CekCompanyIDAvaibilityRes) Validate() error {
	return nil
}
func (this *TempGenToken) Validate() error {
	return nil
}
func (this *BRICaMSSvcCompanyV2CG) Validate() error {
	if this.CompanyGroup != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CompanyGroup); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CompanyGroup", err)
		}
	}
	return nil
}
func (this *CompanyGroupBRICaMS) Validate() error {
	return nil
}
func (this *BRICaMSSvcPagionationV2) Validate() error {
	return nil
}
func (this *BRICaMSSvcMultipleResV2) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
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
func (this *BRICaMSSvcMultipleResV2CG) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
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
func (this *BRICaMSSvcSingleResV2) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *FileListCompanyTaskRequest) Validate() error {
	return nil
}
func (this *ErrorBodyResponse) Validate() error {
	return nil
}
