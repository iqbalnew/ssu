// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swift_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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
func (this *SwiftRemittanceRequest) Validate() error {
	return nil
}
func (this *SwiftRemittance) Validate() error {
	return nil
}
func (this *SwiftRemittanceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CounterpartTransactionRequest) Validate() error {
	return nil
}
func (this *CounterpartTransaction) Validate() error {
	return nil
}
func (this *CounterpartTransactionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SwiftRoutePartnerRequest) Validate() error {
	return nil
}
func (this *KursBRIEFXRequest) Validate() error {
	return nil
}
func (this *BRIGateBRIefxRequest) Validate() error {
	return nil
}
func (this *BRIGateBRIefxData) Validate() error {
	return nil
}
func (this *BRIGateBRIefxResponse) Validate() error {
	for _, item := range this.ResponseData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ResponseData", err)
			}
		}
	}
	return nil
}
func (this *BRIGateBRIefxAddons) Validate() error {
	return nil
}
func (this *BRIGateBRIefxAddonsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *KursBRIEFX) Validate() error {
	return nil
}
func (this *KursBRIEFXResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SwiftRoutePartnerData) Validate() error {
	if this.ENUMERATED_VALUE != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ENUMERATED_VALUE); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ENUMERATED_VALUE", err)
		}
	}
	return nil
}
func (this *SwiftRoutePartnerDataEnumeratedValue) Validate() error {
	return nil
}
func (this *SwiftRoutePartner) Validate() error {
	for _, item := range this.DATA {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DATA", err)
			}
		}
	}
	return nil
}
func (this *SwiftRoutePartnerResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SwiftStatusTransactionRequest) Validate() error {
	return nil
}
func (this *SwiftStatusTransactionData) Validate() error {
	return nil
}
func (this *SwiftStatusTransaction) Validate() error {
	if this.DATA != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DATA); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DATA", err)
		}
	}
	return nil
}
func (this *SwiftStatusTransactionResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaskCreateRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.SelectedRoutePartner != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectedRoutePartner); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectedRoutePartner", err)
		}
	}
	return nil
}
func (this *TaskDetail) Validate() error {
	if this.RoutePartnerReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RoutePartnerReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RoutePartnerReq", err)
		}
	}
	if this.RoutePartnerData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RoutePartnerData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RoutePartnerData", err)
		}
	}
	if this.SwiftRemittanceReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SwiftRemittanceReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SwiftRemittanceReq", err)
		}
	}
	if this.SwiftRemittanceData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SwiftRemittanceData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SwiftRemittanceData", err)
		}
	}
	return nil
}
func (this *TaskData) Validate() error {
	if this.Transaction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Transaction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Transaction", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Detail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Detail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Detail", err)
		}
	}
	return nil
}
func (this *TaskCreateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
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
func (this *TaskListRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *TaskListRequest_TaskListFilter) Validate() error {
	return nil
}
func (this *FileTaskListRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *TaskListResponse) Validate() error {
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
func (this *TaskDetailRequest) Validate() error {
	return nil
}
func (this *TaskDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *DataListRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DataListRequest_DataListFilter) Validate() error {
	return nil
}
func (this *DataListDownloadRequest) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	return nil
}
func (this *DataListResponse) Validate() error {
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
func (this *DataDetailRequest) Validate() error {
	return nil
}
func (this *DataDetailResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SaveDataRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SaveDataResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *BRIGateHardTokenValidationRequest) Validate() error {
	return nil
}
func (this *BRIGateHardTokenValidation) Validate() error {
	return nil
}
func (this *BRIGateHardTokenValidationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *NostroPriorityRequest) Validate() error {
	return nil
}
func (this *NostroPriorityResponse) Validate() error {
	return nil
}
func (this *TransactionsCheckerResponse) Validate() error {
	return nil
}
func (this *TransactionsCheckerSingelRequest) Validate() error {
	return nil
}
func (this *TransactionsCheckerSingelResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TransactionsCheckerBulkRequest) Validate() error {
	return nil
}
func (this *TransactionsCheckerBulkResponse) Validate() error {
	if this.Additional != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Additional); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Additional", err)
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
func (this *DelayedHandlerRes) Validate() error {
	return nil
}
func (this *ExecTransactionReq) Validate() error {
	return nil
}
func (this *ExecTransactionRes) Validate() error {
	return nil
}
func (this *ApprovalCacheReq) Validate() error {
	return nil
}
func (this *ApprovalCacheRes) Validate() error {
	return nil
}
