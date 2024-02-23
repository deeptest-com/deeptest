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
	runningTest := controllerExec.GetRunningTest()

	if act == consts.JoinPerformanceTest {
		if runningTest == nil { // no exist room to join
			websocketHelper.SendExecInstructionToClient(
				"", err, ptconsts.MsgInstructionJoinExist, req.Room, wsMsg)

		} else {
			if req.Room != runningTest.Room { // notify client to join
				websocketHelper.SendExecInstructionToClient(
					runningTest.Room, err, ptconsts.MsgInstructionJoinExist, req.Room, wsMsg)

				controllerExec.ResumeLog()

			} else { //  client joined successfully
				websocketHelper.SendExecInstructionToClient(
					"performance testing joined", runningTest, ptconsts.MsgInstructionStart, req.Room, wsMsg)
			}
		}

	} else if act == consts.StartPerformanceTest {
		if runningTest != nil { // client should call like this
			return
		}

		ptlog.Init(req.Room)

		websocketHelper.SendExecInstructionToClient(
			"performance testing start", nil, ptconsts.MsgInstructionStart, req.Room, wsMsg)

		performanceTestService := controllerExec.NewPerformanceTestServiceRef(req)

		controllerExec.SetRunningTest(&req)
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
			controllerExec.SetRunningTest(nil)
			sendStopMsg("stop failed", req.Room, wsMsg)
			return
		}

		controllerExec.SetRunningTest(nil)
		sendStopMsg("stop successfully", req.Room, wsMsg)

	}

	return
}

func getPerformanceTestServiceRef(room string) (ret *controllerExec.PerformanceTestService) {
	performanceTestServiceObj, ok := PerformanceTestServicesMap.Load(room)
	if !ok {
		controllerExec.SetRunningTest(nil)
		return
	}

	ret, ok = performanceTestServiceObj.(*controllerExec.PerformanceTestService)
	if !ok {
		controllerExec.SetRunningTest(nil)
		return
	}

	return
}

func sendStopMsg(data interface{}, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		"performance testing stop", data, ptconsts.MsgInstructionTerminal, room, wsMsg)
}
