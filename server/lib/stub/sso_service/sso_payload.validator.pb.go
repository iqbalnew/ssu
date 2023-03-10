// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sso_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
func (this *CBMTokenRequest) Validate() error {
	return nil
}
func (this *CBMTokenData) Validate() error {
	return nil
}
func (this *CBMTokenResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CBMSessionValidationRequest) Validate() error {
	return nil
}
func (this *CBMSessionValidationData) Validate() error {
	return nil
}
func (this *CBMSessionValidationResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SSOGetLoginURLReq) Validate() error {
	return nil
}
func (this *SSOGetLoginURLData) Validate() error {
	if this.MessageResponse != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MessageResponse); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MessageResponse", err)
		}
	}
	return nil
}
func (this *CBMMessageResponse) Validate() error {
	return nil
}
func (this *SSOGetLoginURLRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SSOLogoutData) Validate() error {
	return nil
}
func (this *SSOLogoutRes) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ProductAuthority) Validate() error {
	return nil
}
func (this *ErrorBodyResponse) Validate() error {
	return nil
}
func (this *CBMCompany) Validate() error {
	return nil
}
func (this *CBMUser) Validate() error {
	return nil
}
func (this *CBMSyncUserReq) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CBMSyncCompanyReq) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CBMSyncUserRes) Validate() error {
	if this.ResponseData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResponseData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResponseData", err)
		}
	}
	if this.DataToSave != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DataToSave); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DataToSave", err)
		}
	}
	return nil
}
func (this *CBMSyncUserResData) Validate() error {
	return nil
}
func (this *CBMSyncCompanyRes) Validate() error {
	if this.ResponseData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResponseData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResponseData", err)
		}
	}
	if this.DataToSave != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DataToSave); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DataToSave", err)
		}
	}
	return nil
}
func (this *CBMSyncCompanyResData) Validate() error {
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
