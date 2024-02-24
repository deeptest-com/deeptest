package conductorService

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

func RunPerformanceLog(act consts.ExecType, req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	if act == consts.StartPerformanceLog {
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

func sendLogMsg(line string, room string, wsMsg *websocket.Message) {
	websocketHelper.SendExecInstructionToClient(
		line, nil, ptconsts.MsgInstructionTerminal, room, wsMsg)
}
