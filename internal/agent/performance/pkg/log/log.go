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
	logPath := GetLogPath(room)

	logDir := filepath.Dir(logPath)
	if !dir.IsExist(logDir) {
		dir.InsureDir(logDir)
	}

	Logger, _ = getLogger(logPath)
}

func getLogger(logPath string) (*zap.Logger, error) {
	performancePriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {

		return lev >= zap.DebugLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	_fileUtils.RmDir(logPath)

	performanceWriteSyncer, lowClose, err := zap.Open(logPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	performanceCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), performanceWriteSyncer, performancePriority)

	return zap.New(zapcore.NewTee(performanceCore), zap.AddCaller()), nil
}

func GetLogPath(room string) (ret string) {
	ret = filepath.Join(consts.WorkDir, "log", "performance", fmt.Sprintf("%s.log", room))

	return
}
