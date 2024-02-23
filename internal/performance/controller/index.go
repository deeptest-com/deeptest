package controller

import (
	controllerExec "github.com/aaronchen2k/deeptest/internal/performance/controller/exec"
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
)

func RunPerformanceTest(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	runningRoom := controllerExec.GetRunningRoom()

	if act == consts.JoinPerformanceTest {
		existRunningRoom := ""
		if runningRoom != "" && req.Room != runningRoom {
			existRunningRoom = runningRoom
			controllerExec.ResumeLog()
		}

		websocketHelper.SendExecInstructionToClient(existRunningRoom, err, ptconsts.MsgInstructionJoinExist, req.Room, wsMsg)

	} else if act == consts.StartPerformanceTest {
		ptlog.Init(req.Room)

		websocketHelper.SendExecInstructionToClient(
			"performance testing start", err, ptconsts.MsgInstructionStart, req.Room, wsMsg)

		performanceTestService := controllerExec.NewPerformanceTestServiceRef(req)

		controllerExec.SetRunningRoom(req.Room)
		PerformanceTestServicesMap.Store(req.Room, performanceTestService)

		performanceTestService.ExecStart(req, wsMsg)

	} else if act == consts.StopPerformanceTest {
		performanceTestService := getPerformanceTestServiceRef(req.Room)
		if performanceTestService == nil {
			sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
			return
		}

		err = performanceTestService.ExecStop(wsMsg)
		if err != nil {
			controllerExec.SetRunningRoom("")
			sendStopMsg("stop failed", req.Room, wsMsg)
			return
		}

		controllerExec.SetRunningRoom("")
		sendStopMsg("stop successfully", req.Room, wsMsg)

	} else if act == consts.StartPerformanceLog {
		room := req.Room
		performanceTestService := getPerformanceTestServiceRef(room)
		if performanceTestService == nil {
			sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
			return
		}

		performanceTestService.StartSendLog(req, wsMsg)

	} else if act == consts.StopPerformanceLog {
		room := req.Room
		performanceTestService := getPerformanceTestServiceRef(room)
		if performanceTestService == nil {
			sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
			return
		}

		performanceTestService.StopSendLog(req, wsMsg)
	}

	return
}

func getPerformanceTestServiceRef(room string) (ret *controllerExec.PerformanceTestService) {
	performanceTestServiceObj, ok := PerformanceTestServicesMap.Load(room)
	if !ok {
		controllerExec.SetRunningRoom("")
		return
	}

	ret, ok = performanceTestServiceObj.(*controllerExec.PerformanceTestService)
	if !ok {
		controllerExec.SetRunningRoom("")
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
