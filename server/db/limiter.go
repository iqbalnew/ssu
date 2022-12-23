package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	redisV8 "github.com/go-redis/redis/v8"
)

type TaskRedisCount struct {
	TaskID   uint64
	TypeTask string
}
type RedisCount struct {
	Task []TaskRedisCount
}

type APILimitObject struct {
	FuncRPC         string
	Method          string
	Duration        string
	CountMax        string
	DurationBlocked string
	CountBlockedMax string
}

func (p *GormProvider) SuspiciousActivityLimiter(ctx context.Context, md metadata.MD, api APILimitObject, task TaskRedisCount) error {

	ips := strings.Split(md["x-forwarded-for"][0], ",")
	ip := ""
	if len(ips) > 0 {
		ip = ips[0]
	} else {
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] internal server error [1]")
		return status.Error(codes.Internal, "internal server error [1]")
	}

	if getEnv("ENV", "DEV") == "DEV" {
		ip = "0.0.0.0"
	}

	userToken := ""
	if md["authorization"][0] != "" {
		userToken = strings.Split(md["authorization"][0], " ")[1]
	}

	logrus.Infoln("[DEBUG IP] Client IP address:", ip)

	data := RedisCount{}

	_, err := p.rdb.Get(ctx, fmt.Sprintf("suspicious_limiter:%v:%v:%v-%v:blocked", api.FuncRPC, api.Method, ip, userToken))
	if err != redisV8.Nil {
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] your request have suspicious activity")
		return status.Error(codes.PermissionDenied, "you have suspicious activity")
	}

	val, err := p.rdb.Get(ctx, fmt.Sprintf("suspicious_limiter:%v:%v:%v-%v", api.FuncRPC, api.Method, ip, userToken))
	switch {
	case err == redisV8.Nil:
		data.Task = append(data.Task, task)

		dura := "1m"
		if api.Duration != "" {
			dura = api.Duration
		}
		limitDuration, err := time.ParseDuration(dura)
		if err != nil {
			logrus.Panic(err)
		}

		if err = p.rdb.Set(
			ctx,
			fmt.Sprintf("suspicious_limiter:%v:%v:%v-%v", api.FuncRPC, api.Method, ip, userToken),
			data,
			limitDuration,
		); err != nil {
			logrus.Errorln("[db][redis][SuspiciousActivityLimiter] failed set check limit dleet")
			return status.Error(codes.Internal, "internal server error [2]")
		}

		return nil
	case err != nil:
		logrus.Errorln("Filed to find key from RDB: ", err)
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] internal server error [3]")
		return status.Error(codes.Internal, "internal server error [3]")
	}

	if err = json.Unmarshal([]byte(val), &data); err != nil {
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] Failed to unmarshal RDB value")
		return status.Error(codes.Internal, "internal server error [4]")
	}

	countMax, err := strconv.Atoi(api.CountMax)
	if err != nil {
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] Failed to convet count max")
		return status.Error(codes.Internal, "internal server error [5]")
	}

	// logrus.Infoln("Limit Count:", data.Count)

	data.Task = append(data.Task, task)
	suspiciousCount := 0

	countBlockedMax := 0
	if api.CountBlockedMax != "" {
		countBlockedMax, err = strconv.Atoi(api.CountBlockedMax)
		if err != nil {
			logrus.Errorln("[db][redis][SuspiciousActivityLimiter] Failed to convet count blocked max")
			return status.Error(codes.Internal, "internal server error [6]")
		}
	}

	types := make(map[string]string)
	if len(data.Task) >= countMax {

		for i, v := range data.Task {

			if i+1 < len(data.Task) {

				if v.TaskID+1 == data.Task[i+1].TaskID {

					suspiciousCount += 1
					types[v.TypeTask] = v.TypeTask

					if suspiciousCount >= countBlockedMax {
						suspiciousCount += 1
						types[data.Task[0].TypeTask] = data.Task[0].TypeTask
						break
					}

				} else {

					suspiciousCount = 0
					types = make(map[string]string)

				}
			}
		}
	}

	if suspiciousCount >= countBlockedMax {
		if len(types) > 1 {
			// if data.Count >= countMax {
			dura := "15m"
			if api.Duration != "" {
				dura = api.DurationBlocked
			}
			limitDuration, _ := time.ParseDuration(dura)
			if err = p.rdb.Set(
				ctx,
				fmt.Sprintf("suspicious_limiter:%v:%v:%v-%v:blocked", api.FuncRPC, api.Method, ip, userToken),
				data,
				limitDuration,
			); err != nil {
				logrus.Errorln("[db][redis][SuspiciousActivityLimiter] failed set blocker delete")
				return status.Error(codes.Internal, "internal server error [7]")
			}

			logrus.Errorln("[db][redis][SuspiciousActivityLimiter] your request have suspicious activity")
			return status.Error(codes.PermissionDenied, "you have suspicious activity")
			//}
		}
	}

	logrus.Println("[db][redis][SuspiciousActivityLimiter] suspicious count:", suspiciousCount)
	logrus.Println("[db][redis][SuspiciousActivityLimiter] count task: ", len(data.Task))
	logrus.Println("[db][redis][SuspiciousActivityLimiter] count type: ", len(types))
	logrus.Println("[db][redis][SuspiciousActivityLimiter] type map:", types)

	if err = p.rdb.Update(
		ctx,
		fmt.Sprintf("suspicious_limiter:%v:%v:%v-%v", api.FuncRPC, api.Method, ip, userToken),
		data,
	); err != nil {
		logrus.Errorln(err.Error())
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] internal server error [8]")
		return status.Error(codes.Internal, "internal server error [8]")
	}

	if suspiciousCount >= countMax {
		logrus.Errorln("[db][redis][SuspiciousActivityLimiter] Request is exceeding limit len:", len(data.Task))
		return status.Error(codes.InvalidArgument, "Request is exceeding limit")
	}

	return nil
}
