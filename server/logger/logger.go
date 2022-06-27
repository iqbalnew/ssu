package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	fluent *fluent.Fluent
	tag    string
}

func NewLogger(port string, host string, tag string) *Logger {
	logrus.Println("Logger host: ", host)
	logrus.Println("Logger port: ", port)
	logrus.Println("Logger tag: ", tag)

	if getEnv("ENV", "DEV") != "LOCAL" {
		portVal, _ := strconv.Atoi(port)

		config := fluent.Config{
			FluentPort:    portVal,
			FluentHost:    host,
			MarshalAsJSON: true,
			RequestAck:    true,
		}

		logrus.Println("MarshalAsJson : ", config.MarshalAsJSON)

		fluentd, err := fluent.New(config)
		if err != nil {
			panic(err)
		}

		logger := &Logger{
			fluent: fluentd,
		}

		if tag != "" {
			logger.tag = tag
		} else {
			logger.tag = "default"
		}

		return logger
	}
	return nil
}

func (l *Logger) Close() {
	l.fluent.Close()
}

func (l *Logger) PostWithTime(tag string, data map[string]interface{}, time time.Time) {
	l.fluent.PostWithTime(tag, time, data)
}

func (l *Logger) Error(function string, text string) {
	data := map[string]string{
		"level":    "error",
		"function": function,
		"error":    text,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Errorln(data)
	}
}

func (l *Logger) ErrorWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "info",
		"error": text,
		"data":  x,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Errorln(data)
	}
}

func (l *Logger) Info(text string) {
	data := map[string]interface{}{
		"level":     "info",
		"info":      text,
		"timestamp": time.Now().String(),
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		now := time.Now()
		logrus.Println("Logger send: ", data)
		err := l.fluent.PostWithTime(l.tag, now, data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Infoln(data)
	}
}

func (l *Logger) InfoWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "info",
		"info":  text,
		"data":  x,
	}
	logrus.Println("HttpLog logger ver", data)
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Infoln(data)
	}
}

func (l *Logger) InfoWithDataMap(text string, x map[string]interface{}) {
	now := time.Now()

	data := map[string]interface{}{
		"level":     "info",
		"info":      text,
		"timestamp": now.Format(time.RFC3339Nano),
	}
	for k, v := range x {
		data[k] = v
	}

	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, now, data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Infoln(data)
	}
}

func (l *Logger) InfoUser(text string, userID string, companyID string, action string) {
	data := map[string]string{
		"level":     "info",
		"info":      text,
		"userID":    userID,
		"companyID": companyID,
		"action":    action,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Infoln(data)
	}
}

type UserData struct {
	UserID    string `json:"userID"`
	UserName  string `json:"userName"`
	CompanyID string `json:"companyID"`
	Action    string `json:"action"`
}

func (l *Logger) InfoUserWithData(text string, user UserData, x interface{}, logID string) {
	data := map[string]interface{}{
		"level": "info",
		"logID": logID,
		"info":  text,
		"user":  user,
		"data":  x,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		// logrus.Infoln(data)
	}
}

func (l *Logger) ErrorUserWithData(text string, user UserData, x interface{}, logID string) {
	data := map[string]interface{}{
		"level": "error",
		"logID": logID,
		"info":  text,
		"user":  user,
		"data":  x,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		// logrus.Infoln(data)
	}
}

func (l *Logger) Debug(text string) {
	data := map[string]string{
		"level": "debug",
		"debug": text,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Debugln(data)
	}
}

func (l *Logger) DebugWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "debug",
		"debug": text,
		"data":  x,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
		if err != nil {
			logrus.Errorln("Error on Send Log to Fluentd: ", err)
		}
	} else {
		logrus.Debugln(data)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
