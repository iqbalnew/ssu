// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: workflow_gorm_db.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *Workflow) Validate() error {
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	for _, item := range this.Logics {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Logics", err)
			}
		}
	}
	return nil
}
func (this *WorkflowLogic) Validate() error {
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
	for _, item := range this.Requirements {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Requirements", err)
			}
		}
	}
	return nil
}
func (this *WorkflowRequirement) Validate() error {
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
func (this *CompanyWorkflows) Validate() error {
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
