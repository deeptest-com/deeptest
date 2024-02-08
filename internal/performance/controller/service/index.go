package controllerService

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var (
	PerformanceTestServicesMap sync.Map
)

func RunPerformanceTest(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	if act == consts.StartPerformanceTest {
		websocketHelper.SendExecInstructionToClient(
			"exec start", err, ptconsts.MsgInstructionStart, req.Room, wsMsg)

		go func() {
			performanceTestService := NewPerformanceTestService()
			err = NewPerformanceTestService().ExecStart(req, wsMsg)

			PerformanceTestServicesMap.Store(wsMsg.Room, &performanceTestService)
		}()

	} else if act == consts.StopPerformanceTest {
		performanceTestServiceObj, ok := PerformanceTestServicesMap.Load(req.Room)
		if !ok {
			return
		}

		performanceTestService := performanceTestServiceObj.(*PerformanceTestService)

		err = performanceTestService.ExecStop(req, wsMsg)
		if err == nil {
			websocketHelper.SendExecInstructionToClient(
				"exec continue", err, ptconsts.MsgInstructionTerminal, req.Room, wsMsg)
		}
	}

	return
}
