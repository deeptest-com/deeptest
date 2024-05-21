package _logUtils

import (
	"fmt"
	"runtime/debug"
)

func Info(str string) {
	Logger.Info(str)
}
func Infof(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Info(msg)
}

func Warn(str string) {
	Logger.Warn(str)
}
func Warnf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Warn(msg)
}

func Error(str string) {
	Logger.Error(str)
	s := string(debug.Stack())
	fmt.Printf("err=%v, stack=%s\n", str, s)
}
func Errorf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Error(msg)
	s := string(debug.Stack())
	fmt.Printf("err=%v, stack=%s\n", msg, s)
}
func Debug(str string) {
	Logger.Debug(str)
}
func Debugf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	Logger.Debug(msg)
}
