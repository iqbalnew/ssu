// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: company_gorm_db.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mwitkow/go-proto-validators"
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
func (this *CompanyGroup) Validate() error {
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
	for _, item := range this.GroupLimits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GroupLimits", err)
			}
		}
	}
	for _, item := range this.CompanyLimits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CompanyLimits", err)
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
	return nil
}
func (this *GroupLimit) Validate() error {
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *CompanyLimit) Validate() error {
	if this.Currency != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Currency); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Currency", err)
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
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *GroupSubsidiary) Validate() error {
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
func (this *CompanyWorkflow) Validate() error {
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
func (this *Currency) Validate() error {
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
