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
	var err error

	taskID := 0
	if val.TaskID != "" {
		taskID, err = strconv.Atoi(aes.Decrypt(val.TaskID))
		if err != nil {
			// handle error
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
		}
	}

	createdByID := 0
	if len(val.CreatedByID) > 0 {

		createdByID, err = strconv.Atoi(aes.Decrypt(val.CreatedByID))
		if err != nil {
			// handle error
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt taskID")
		}
	}

	lastApprovedByID := 0
	if len(val.LastApprovedByID) > 0 {
		lastApprovedByID, err = strconv.Atoi(aes.Decrypt(val.LastApprovedByID))
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastApprovedByID")
		}
	}

	lastRejectedByID := 0
	if len(val.LastRejectedByID) > 0 {
		lastRejectedByID, err = strconv.Atoi(aes.Decrypt(val.LastRejectedByID))
		if err != nil {
			logrus.Errorln(err)
			return nil, status.Errorf(codes.Internal, "Failed to decrypt LastRejectedByID")
		}
	}

	featureID := 0
	if len(val.FeatureID) > 0 {
		featureID, err = strconv.Atoi(aes.Decrypt(val.FeatureID))
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

	data.TaskID = aes.Encrypt(fmt.Sprint(val.TaskID))
	data.Type = val.Type
	data.Status = val.Status
	data.Step = val.Step
	data.CreatedByID = aes.Encrypt(fmt.Sprint(val.CreatedByID))
	data.LastApprovedByID = aes.Encrypt(fmt.Sprint(val.LastApprovedByID))
	data.LastRejectedByID = aes.Encrypt(fmt.Sprint(val.LastRejectedByID))
	data.Data = val.Data
	data.Reasons = val.Reasons
	data.Comment = val.Comment
	data.FeatureID = aes.Encrypt(fmt.Sprint(val.FeatureID))
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
