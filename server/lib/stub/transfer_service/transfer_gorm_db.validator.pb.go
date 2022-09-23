// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer_gorm_db.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (this *TransferJob) Validate() error {
	if this.RunAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RunAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RunAt", err)
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
func (this *MassInquiryJob) Validate() error {
	if this.RunAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RunAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RunAt", err)
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
func (this *InternalSingleTemplate) Validate() error {
	if !(this.Amount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Amount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.Amount))
	}
	if !(this.ExchangeRate >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ExchangeRate", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ExchangeRate))
	}
	if !(this.ReceivedAmount >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ReceivedAmount", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ReceivedAmount))
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
	if this.ScheduledAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScheduledAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScheduledAt", err)
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
