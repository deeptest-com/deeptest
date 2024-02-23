package service

import (
	"github.com/aaronchen2k/deeptest/internal/performance/controller"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

func RunPerformanceLog(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	ptlog.Logf("run performance log act=%s, room=%s", act.String(), req.Room)

	controller.RunPerformanceLog(act, req, wsMsg)

	return
}
