package ptlog

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

var Logger *zap.Logger

func Logf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)
	log.Println(msg)

	if Logger != nil {
		Logger.Debug(msg)
	}
}
