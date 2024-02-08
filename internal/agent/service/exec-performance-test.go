package service

import (
	controllerService "github.com/aaronchen2k/deeptest/internal/performance/controller/service"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"go.uber.org/zap"
)

func RunPerformanceTest(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run performance test", zap.String("act", act.String()), zap.String("room", req.Room))

	controllerService.RunPerformanceTest(act, req, wsMsg)

	return
}
