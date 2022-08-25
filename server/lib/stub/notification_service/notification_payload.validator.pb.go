// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notification_payload.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
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
func (this *NotificationFilter) Validate() error {
	return nil
}
func (this *GetNotificationsReq) Validate() error {
	if this.Notification != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Notification); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Notification", err)
		}
	}
	return nil
}
func (this *Sort) Validate() error {
	return nil
}
func (this *ListNotificationRes) Validate() error {
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
func (this *ListNotificationFCMRes) Validate() error {
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
func (this *ErrorBodyResponse) Validate() error {
	return nil
}
func (this *CreateNotificationRequest) Validate() error {
	if this.Notification != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Notification); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Notification", err)
		}
	}
	for _, item := range this.Company {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Company", err)
			}
		}
	}
	if this.NotificationEmail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotificationEmail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotificationEmail", err)
		}
	}
	if this.NotificationSMS != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NotificationSMS); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NotificationSMS", err)
		}
	}
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	if this.Client != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Client); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Client", err)
		}
	}
	return nil
}
func (this *CreateNotificationTaskRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *SendPushNotificationRequest) Validate() error {
	return nil
}
func (this *CommonResponse) Validate() error {
	return nil
}
func (this *CreateNotificationTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateNotificationTaskResponseData) Validate() error {
	return nil
}
func (this *GetNotificationTaskByIDRequest) Validate() error {
	return nil
}
func (this *GetNotificationTaskRequest) Validate() error {
	return nil
}
func (this *PaginationResponse) Validate() error {
	return nil
}
func (this *GetNotificationTaskResponse) Validate() error {
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
func (this *GetDetailNotificationTaskResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateNotificationTask) Validate() error {
	if this.Notification != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Notification); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Notification", err)
		}
	}
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *NotificationTaskDto) Validate() error {
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
func (this *ListNotificationModuleResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ListModuleEventResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *TempGenToken) Validate() error {
	return nil
}
func (this *SendEmailRequest) Validate() error {
	return nil
}
func (this *SendSmsRequest) Validate() error {
	return nil
}
func (this *BRIGateResponseDetail) Validate() error {
	return nil
}
func (this *BRIGateNotificationResponse) Validate() error {
	if this.Detail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Detail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Detail", err)
		}
	}
	return nil
}
func (this *FileListNotificationTaskRequest) Validate() error {
	return nil
}
func (this *SendNotificationRequest) Validate() error {
	return nil
}
func (this *SendNotificationResponse) Validate() error {
	return nil
}
func (this *HistoryNotificationRequest) Validate() error {
	return nil
}
func (this *HistoryNotificationResponse) Validate() error {
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
func (this *PaginationHistory) Validate() error {
	return nil
}
func (this *HistoryNotificationDetail) Validate() error {
	return nil
}
func (this *UpdateLogHistoryNotifiationRequest) Validate() error {
	return nil
}
func (this *UpdateLogHistoryNotifiationResponse) Validate() error {
	return nil
}
func (this *ModuleVarMessage) Validate() error {
	if this.Module != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Module); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Module", err)
		}
	}
	return nil
}
func (this *ModuleVarMessage_Module) Validate() error {
	return nil
}
func (this *ListModuleVariableReq) Validate() error {
	return nil
}
func (this *ListModuleVariableResp) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
