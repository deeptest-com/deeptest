package conductorService

import (
	conductorExec "github.com/aaronchen2k/deeptest/internal/performance/conductor/exec"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/kataras/iris/v12/websocket"
)

func JoinPerformanceTest(room string, wsMsg *websocket.Message) (err error) {
	runningTest := conductorExec.GetRunningTest()

	if runningTest == nil { // no exist room to join
		websocketHelper.SendExecInstructionToClient(
			"", nil, ptconsts.MsgInstructionJoinExist, room, wsMsg)

	} else {
		if room != runningTest.Room { // notify client to join
			websocketHelper.SendExecInstructionToClient(
				runningTest.Room, nil, ptconsts.MsgInstructionJoinExist, room, wsMsg)

			conductorExec.ResumeWsMsg()

		} else { //  client joined successfully
			websocketHelper.SendExecInstructionToClient(
				"performance testing joined", runningTest, ptconsts.MsgInstructionStart, room, wsMsg)
		}
	}

	return
}

func StartPerformanceTest(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	runningTest := conductorExec.GetRunningTest()
	if runningTest != nil { // client should call like this
		return
	}

	ptlog.Init(req.Room)

	websocketHelper.SendExecInstructionToClient(
		"performance testing start", nil, ptconsts.MsgInstructionStart, req.Room, wsMsg)

	performanceTestService := NewPerformanceTestServiceRef(req)

	conductorExec.SetRunningTest(&req)
	PerformanceTestServicesMap.Store(req.Room, performanceTestService)

	go performanceTestService.ExecStart(req, wsMsg)

	return
}

func StopPerformanceTest(room string, wsMsg *websocket.Message) (err error) {
	performanceTestService := GetPerformanceTestServiceRef(room)
	if performanceTestService == nil {
		sendStopMsg("get performanceTestService failed", room, wsMsg)
		return
	}

	err = performanceTestService.ExecStop(wsMsg)
	if err != nil {
		conductorExec.SetRunningTest(nil)
		sendStopMsg("stop failed", room, wsMsg)
		return
	}

	conductorExec.SetRunningTest(nil)
	sendStopMsg("stop successfully", room, wsMsg)

	return
}

func StartPerformanceLog(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	room := req.Room
	performanceTestService := GetPerformanceTestServiceRef(room)
	if performanceTestService == nil {
		sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
		return
	}

	go performanceTestService.StartSendLog(req, wsMsg)

	return
}

func StopPerformanceLog(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	room := req.Room
	performanceTestService := GetPerformanceTestServiceRef(room)
	if performanceTestService == nil {
		sendStopMsg("get performanceTestService failed", req.Room, wsMsg)
		return
	}

	performanceTestService.StopSendLog(req, wsMsg)

	return
}

func sendStopMsg(data interface{}, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		"performance testing stop", data, ptconsts.MsgInstructionTerminal, room, wsMsg)
}
