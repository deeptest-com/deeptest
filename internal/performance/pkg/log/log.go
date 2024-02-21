package ptlog

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/snowlyg/helper/dir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

func Init(room string) {
	logDir := filepath.Join(consts.WorkDir, "log", "performance")
	if !dir.IsExist(logDir) {
		dir.InsureDir(logDir)
	}

	Logger, _ = getLogger(logDir, room)
}

func getLogger(logDir, room string) (*zap.Logger, error) {
	performancePriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {

		return lev >= zap.DebugLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	performanceLogPath := filepath.Join(logDir, fmt.Sprintf("%s.log", room))
	_fileUtils.RmDir(performanceLogPath)

	performanceWriteSyncer, lowClose, err := zap.Open(performanceLogPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	performanceCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), performanceWriteSyncer, performancePriority)

	return zap.New(zapcore.NewTee(performanceCore), zap.AddCaller()), nil
}
