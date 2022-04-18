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

		fluentd, err := fluent.New(fluent.Config{
			FluentPort: portVal,
			FluentHost: host,
		})
		if err != nil {
			panic(err)
		}

		logger := &Logger{
			fluent: fluentd,
			tag:    tag,
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
	data := map[string]string{
		"level": "info",
		"info":  text,
	}
	if getEnv("ENV", "DEV") != "LOCAL" {
		now := time.Now()
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
	if getEnv("ENV", "DEV") != "LOCAL" {
		err := l.fluent.PostWithTime(l.tag, time.Now(), data)
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
