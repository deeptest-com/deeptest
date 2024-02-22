package controller

import (
	controllerService "github.com/aaronchen2k/deeptest/internal/performance/controller/service"
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

		performanceTestService := controllerService.NewPerformanceTestServiceRef(req)

		performanceTestService.ExecStart(req, wsMsg)

		PerformanceTestServicesMap.Store(req.Room, performanceTestService)
		runningRoom = req.Room

	} else if act == consts.StopPerformanceTest {
		performanceTestService := getPerformanceTestServiceRef(req.Room)
		if performanceTestService == nil {
			sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
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

	} else if act == consts.StartPerformanceLog {
		room := req.Room
		performanceTestService := getPerformanceTestServiceRef(room)
		if performanceTestService == nil {
			sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
			return
		}

		performanceTestService.SendLogAsync(req, wsMsg)
	}

	return
}

func getPerformanceTestServiceRef(room string) (ret *controllerService.PerformanceTestService) {
	performanceTestServiceObj, ok := PerformanceTestServicesMap.Load(room)
	if !ok {
		runningRoom = ""
		return
	}

	ret, ok = performanceTestServiceObj.(*controllerService.PerformanceTestService)
	if !ok {
		runningRoom = ""
		return
	}

	return
}

func sendStopMsg(data interface{}, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		"performance testing stop", data, ptconsts.MsgInstructionTerminal, room, wsMsg)
}

func sendLogMsg(line string, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		line, nil, ptconsts.MsgInstructionTerminal, room, wsMsg)
}
