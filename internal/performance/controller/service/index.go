package controllerService

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var (
	PerformanceTestServicesMap sync.Map

	runningRoom string
)

func RunPerformanceTest(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	if act == consts.JoinPerformanceTest {
		if runningRoom != "" && req.Room != runningRoom {
			websocketHelper.SendExecInstructionToClient(
				runningRoom, err, ptconsts.MsgInstructionJoinExist, req.Room, wsMsg)
		}

	} else if act == consts.StartPerformanceTest {
		ptlog.Init(req.Room)

		websocketHelper.SendExecInstructionToClient(
			"performance testing start", err, ptconsts.MsgInstructionStart, req.Room, wsMsg)

		performanceTestService := NewPerformanceTestServiceRef(req)

		go func() {
			performanceTestService.ExecStart(req, wsMsg)
		}()

		PerformanceTestServicesMap.Store(req.Room, performanceTestService)
		runningRoom = req.Room

	} else if act == consts.StopPerformanceTest {
		performanceTestServiceObj, ok := PerformanceTestServicesMap.Load(req.Room)
		if !ok {
			runningRoom = ""
			sendStopMsg("performanceTestService not found", req.Room, wsMsg)
			return
		}

		performanceTestService, ok := performanceTestServiceObj.(*PerformanceTestService)
		if !ok {
			runningRoom = ""
			sendStopMsg("performanceTestService failed convert", req.Room, wsMsg)
			return
		}

		err = performanceTestService.ExecStop(wsMsg)
		if err != nil {
			runningRoom = ""
			sendStopMsg("stop failed", req.Room, wsMsg)
			return
		}

		runningRoom = ""
		sendStopMsg("stop successfully", req.Room, wsMsg)
	}

	return
}

func sendStopMsg(data interface{}, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		"performance testing stop", data, ptconsts.MsgInstructionTerminal, room, wsMsg)
}
