package zapLog

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	myZap "github.com/aaronchen2k/deeptest/pkg/core/zap"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/snowlyg/helper/dir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// level 日志级别
var level zapcore.Level

// Init 初始化日志服务
func Init(app string) {
	var logger *zap.Logger

	logDir := "log"
	if app == "server" {
		logDir = config.CONFIG.Zap.Director
	}

	if !dir.IsExist(logDir) {
		dir.InsureDir(logDir)
	}

	switch config.CONFIG.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(myZap.GetEncoderCore(level), zap.AddStacktrace(level))
	} else {
		logger = zap.New(myZap.GetEncoderCore(level))
	}
	if config.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	logUtils.Logger = logger
}
