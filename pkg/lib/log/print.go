package _logUtils

import (
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Info(str string) {
	Logger.Info(str)
	//log.Println(str)
}
func Infof(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Info(msg)
	//log.Printf(msg+"\n")
}

func Warn(str string) {
	Logger.Warn(str)
	//log.Println(str)
}
func Warnf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Warn(msg)
	//log.Printf(msg+"\n")
}

func Error(str string) {
	Logger.Error(str)
	//log.Println(str)
}
func Errorf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Error(msg)
	//log.Printf(msg+"\n")
}
func Debug(str string) {
	Logger.Debug(str)
	//log.Println(str)
}
func Debugf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Debug(msg)
	//log.Printf(msg+"\n")
}
