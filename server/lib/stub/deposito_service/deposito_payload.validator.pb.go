// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deposito_payload.proto

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

func (this *PaginationResponse) Validate() error {
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
func (this *DepositoDataTask) Validate() error {
	if this.WithdrawAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WithdrawAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WithdrawAt", err)
		}
	}
	if this.OpenDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OpenDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OpenDate", err)
		}
	}
	if this.LastRenewalDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastRenewalDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastRenewalDate", err)
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
func (this *ManageDeposit) Validate() error {
	return nil
}
func (this *ManageDepositTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *WithdrawDepositTaskRequest) Validate() error {
	return nil
}
func (this *CreateDepositoTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaskData) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Deposito != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deposito); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deposito", err)
		}
	}
	if this.Company != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Company); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
		}
	}
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	for _, item := range this.WorkflowHistory {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkflowHistory", err)
			}
		}
	}
	if this.DepositoPrev != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DepositoPrev); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DepositoPrev", err)
		}
	}
	if this.DepositoNew != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DepositoNew); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DepositoNew", err)
		}
	}
	return nil
}
func (this *Account) Validate() error {
	return nil
}
func (this *Accounts) Validate() error {
	if this.AccTD != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AccTD); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AccTD", err)
		}
	}
	if this.AccPrincipal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AccPrincipal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AccPrincipal", err)
		}
	}
	if this.AccInterest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AccInterest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AccInterest", err)
		}
	}
	return nil
}
func (this *Company) Validate() error {
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
func (this *CreateDepositoTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositoListTaskRequest) Validate() error {
	return nil
}
func (this *DepositoListTaskResponse) Validate() error {
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
func (this *FileTaskListRequest) Validate() error {
	return nil
}
func (this *GetDepositoTaskByIDRequest) Validate() error {
	return nil
}
func (this *GetDepositoTaskByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositoDataListRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DepositoDataListRequest_DataListFilter) Validate() error {
	return nil
}
func (this *DepositoDataks) Validate() error {
	if this.Deposito != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deposito); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deposito", err)
		}
	}
	if this.Accounts != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Accounts); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
		}
	}
	return nil
}
func (this *DepositoDataListResponse) Validate() error {
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
func (this *DepositoDataDetailRequest) Validate() error {
	return nil
}
func (this *DepositoDataDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateDepositoRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateDepositoResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetProductDepositoRequest) Validate() error {
	if this.Product != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Product); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Product", err)
		}
	}
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *GetProductDepositoRequest_DataListFilter) Validate() error {
	return nil
}
func (this *GetProductDepositoRespons) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ExecTransactionDepositoReq) Validate() error {
	return nil
}
func (this *ExecTransactionDepositoRes) Validate() error {
	return nil
}
func (this *DepositInquiryRateRequest) Validate() error {
	return nil
}
func (this *DepositInquiryRate) Validate() error {
	return nil
}
func (this *DepositInquiryRateRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositoMaintananceRequest) Validate() error {
	return nil
}
func (this *DepositoMaintanance) Validate() error {
	return nil
}
func (this *DepositoMaintananceRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositoCreateAccountRequest) Validate() error {
	return nil
}
func (this *DepositoCreateAccount) Validate() error {
	return nil
}
func (this *DepositoCreateAccountRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositoPlacementRequest) Validate() error {
	return nil
}
func (this *DepositoPlacement) Validate() error {
	return nil
}
func (this *DepositoPlacementRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositVoucherNegoRequest) Validate() error {
	return nil
}
func (this *DepositInquiryRateNegoRequest) Validate() error {
	return nil
}
func (this *DepositInquiryRateNego) Validate() error {
	return nil
}
func (this *DepositInquiryTDPinaltiRequest) Validate() error {
	return nil
}
func (this *DepositInquiryTDPinalti) Validate() error {
	return nil
}
func (this *DepositInquiryTDPinaltiResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositInquiryRateNegoRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositInquiryCIFRequest) Validate() error {
	return nil
}
func (this *DepositInquiryCIFRespons) Validate() error {
	return nil
}
func (this *DepositWithdrawRequest) Validate() error {
	return nil
}
func (this *DepositoWithdrwal) Validate() error {
	return nil
}
func (this *DepositWithdrawRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DepositInquiryTDRequest) Validate() error {
	return nil
}
func (this *DepositInquiryTD) Validate() error {
	return nil
}
func (this *DepositInquiryTDRespons) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
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
func (this *MakerData) Validate() error {
	return nil
}
func (this *WorkflowHeader) Validate() error {
	if this.Maker != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Maker); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Maker", err)
		}
	}
	return nil
}
func (this *RejectedBy) Validate() error {
	if this.RejectedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RejectedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RejectedAt", err)
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
func (this *WorkflowDataStatus) Validate() error {
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
	return nil
}
func (this *DepositoNotificationData) Validate() error {
	return nil
}
