package zapLog

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/snowlyg/helper/dir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"runtime"
)

func Init() {
	logPath := getLogPath()

	logUtils.Logger = getLogger(logPath)
}

func getLogger(logPath string) (logger *zap.Logger) {
	logDir := filepath.Dir(logPath)
	if !dir.IsExist(logDir) {
		dir.InsureDir(logDir)
	}

	level := getLogLevel()

	logger, err := CreateLogger(logPath, level)

	if err != nil {
		panic(err)
	}

	return
}

func getLogLevel() (level zapcore.Level) {
	switch config.CONFIG.Zap.Level {
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

	return
}

func getLogPath() (ret string) {
	if runtime.GOOS == "windows" {
		ret = filepath.Join("log", fmt.Sprintf("%s.log", consts.RunFrom))
	} else {
		ret = filepath.Join(consts.WorkDir, "log", fmt.Sprintf("%s.log", consts.RunFrom))
	}

	return
}

func CreateLogger(logPath string, level zapcore.Level) (ret *zap.Logger, err error) {
	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {

		return lev >= level
	})
	writeSyncer, lowClose, err := zap.Open(logPath)
	if err != nil {
		if lowClose != nil {
			lowClose()
		}
		return
	}

	swSugar := zapcore.NewMultiWriteSyncer(
		writeSyncer,
		zapcore.AddSync(os.Stdout),
	)
	infoCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), swSugar, infoPriority)

	ret = zap.New(zapcore.NewTee(infoCore))

	ret = ret.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel))

	return
}
