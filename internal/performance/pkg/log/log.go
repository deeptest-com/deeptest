package ptlog

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/snowlyg/helper/dir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

// level 日志级别
var level zapcore.Level

func Init() {
	logDir := filepath.Join(consts.WorkDir, "log")
	if !dir.IsExist(logDir) {
		dir.InsureDir(logDir)
	}

	Logger, _ = getLogger(logDir)
}

func getLogger(logDir string) (*zap.Logger, error) {
	performancePriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev > zap.DebugLevel && lev < zap.ErrorLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	performanceLogPath := filepath.Join(logDir, "performance.log")
	performanceWriteSyncer, lowClose, err := zap.Open(performanceLogPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	infoLogPath := filepath.Join(logDir, "info.log")
	lowWriteSyncer, lowClose, err := zap.Open(infoLogPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	errorLogPath := filepath.Join(consts.WorkDir, "err.log")
	highWriteSyncer, highClose, err := zap.Open(errorLogPath)
	if err != nil {
		highClose()
		return nil, err
	}

	performanceCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), performanceWriteSyncer, performancePriority)
	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), highWriteSyncer, highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), lowWriteSyncer, lowPriority)

	return zap.New(zapcore.NewTee(performanceCore, highCore, lowCore), zap.AddCaller()), nil
}
