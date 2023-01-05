// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proxy_management_core_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (this *WorkflowHeader) Validate() error {
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
func (this *ValidateWorkflowData) Validate() error {
	if this.Workflow != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Workflow); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Workflow", err)
		}
	}
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
func (this *HealthCheckRequest) Validate() error {
	return nil
}
func (this *HealthCheckResponse) Validate() error {
	return nil
}
func (this *NotificationData) Validate() error {
	return nil
}
