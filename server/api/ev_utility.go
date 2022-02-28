package api

import (
	"fmt"
	"strconv"

	customAES "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/aes"
	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func taskEVtoPB(val *pb.TaskEV, aes *customAES.CustomAES) (*pb.Task, error) {
	if val == nil {
		return nil, nil
	}

	data := &pb.Task{}

	taskID := 0
	if val.TaskID != "" {
		text, err := aes.Decrypt(val.TaskID)
		if err != nil {
			logrus.Errorf("val: %v | %v", val.TaskID, err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
		}
		taskID, err = strconv.Atoi(text)
		if err != nil {
			// handle error
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt TaskID")
		}
	}

	createdByID := 0
	if val.CreatedByID != "" {
		text, err := aes.Decrypt(val.CreatedByID)
		if err != nil {
			logrus.Errorf("val: %v | %v", val.CreatedByID, err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt CreatedByID")
		}
		createdByID, err = strconv.Atoi(text)
		if err != nil {
			// handle error
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt CreatedByID")
		}
	}

	lastApprovedByID := 0
	if len(val.LastApprovedByID) > 0 {
		text, err := aes.Decrypt(val.LastApprovedByID)
		if err != nil {
			logrus.Errorf("val: %v | %v", val.LastApprovedByID, err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastApprovedByID")
		}
		lastApprovedByID, err = strconv.Atoi(text)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastApprovedByID")
		}
	}

	lastRejectedByID := 0
	if len(val.LastRejectedByID) > 0 {
		text, err := aes.Decrypt(val.LastRejectedByID)
		if err != nil {
			logrus.Errorf("val: %v | %v", val.LastRejectedByID, err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastRejectedByID")
		}
		lastRejectedByID, err = strconv.Atoi(text)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastRejectedByID")
		}
	}

	featureID := 0
	if len(val.FeatureID) > 0 {
		text, err := aes.Decrypt(val.FeatureID)
		if err != nil {
			logrus.Errorf("val: %v | %v", val.FeatureID, err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt FeatureID")
		}
		featureID, err = strconv.Atoi(text)
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt FeatureID")
		}
	}

	data.TaskID = uint64(taskID)
	data.Type = val.Type
	data.Status = val.Status
	data.Step = val.Step
	data.CreatedByID = uint64(createdByID)
	data.LastApprovedByID = uint64(lastApprovedByID)
	data.LastRejectedByID = uint64(lastRejectedByID)
	data.Data = val.Data
	data.Reasons = val.Reasons
	data.Comment = val.Comment
	data.FeatureID = uint64(featureID)
	data.IsParentActive = val.IsParentActive
	data.CreatedAt = val.CreatedAt
	data.UpdatedAt = val.UpdatedAt
	data.DeletedAt = val.DeletedAt

	for _, v := range val.Childs {
		vr, err := taskEVtoPB(v, aes)
		if err != nil {
			return nil, err
		}
		data.Childs = append(data.Childs, vr)
	}

	return data, nil
}

func taskPBtoEV(val *pb.Task, aes *customAES.CustomAES) (*pb.TaskEV, error) {
	if val == nil {
		return nil, nil
	}

	data := &pb.TaskEV{}
	var err error

	data.TaskID, err = aes.Encrypt(fmt.Sprint(val.TaskID))
	if err != nil {
		logrus.Errorf("val: %v | %v", val.TaskID, err)
		return nil, status.Errorf(codes.Internal, "Failed to encrypt TaskID")
	}
	data.Type = val.Type
	data.Status = val.Status
	data.Step = val.Step
	data.CreatedByID, err = aes.Encrypt(fmt.Sprint(val.CreatedByID))
	if err != nil {
		logrus.Errorf("val: %v | %v", val.CreatedByID, err)
		return nil, status.Errorf(codes.Internal, "Failed to encrypt CreatedByID")
	}
	data.LastApprovedByID, err = aes.Encrypt(fmt.Sprint(val.LastApprovedByID))
	if err != nil {
		logrus.Errorf("val: %v | %v", val.LastApprovedByID, err)
		return nil, status.Errorf(codes.Internal, "Failed to encrypt LastApprovedByID")
	}
	data.LastRejectedByID, err = aes.Encrypt(fmt.Sprint(val.LastRejectedByID))
	if err != nil {
		logrus.Errorf("val: %v | %v", val.LastRejectedByID, err)
		return nil, status.Errorf(codes.Internal, "Failed to encrypt LastRejectedByID")
	}
	data.Data = val.Data
	data.Reasons = val.Reasons
	data.Comment = val.Comment
	data.FeatureID, err = aes.Encrypt(fmt.Sprint(val.FeatureID))
	if err != nil {
		logrus.Errorf("val: %v | %v", val.FeatureID, err)
		return nil, status.Errorf(codes.Internal, "Failed to encrypt FeatureID")
	}
	data.IsParentActive = val.IsParentActive
	data.CreatedAt = val.CreatedAt
	data.UpdatedAt = val.UpdatedAt
	data.DeletedAt = val.DeletedAt

	for _, v := range val.Childs {
		vr, err := taskPBtoEV(v, aes)
		if err != nil {
			return nil, err
		}
		data.Childs = append(data.Childs, vr)
	}

	return data, nil
}
