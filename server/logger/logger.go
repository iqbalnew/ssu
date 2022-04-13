package logger

import (
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
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}

func (l *Logger) ErrorWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "info",
		"error": text,
		"data":  x,
	}
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}

func (l *Logger) Info(text string) {
	data := map[string]string{
		"level": "info",
		"info":  text,
	}
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}

func (l *Logger) InfoWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "info",
		"info":  text,
		"data":  x,
	}
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
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
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}

func (l *Logger) Debug(text string) {
	data := map[string]string{
		"level": "debug",
		"debug": text,
	}
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}

func (l *Logger) DebugWithData(text string, x interface{}) {
	data := map[string]interface{}{
		"level": "debug",
		"debug": text,
		"data":  x,
	}
	err := l.fluent.PostWithTime(l.tag, time.Now(), data)
	if err != nil {
		logrus.Errorln("Error on Send Log to Fluentd: ", err)
	}
}
