package execUtils

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

func SendStartMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(true)
	websocketHelper.SendExecStatus(consts.ProgressInProgress, wsMsg)

	return
}

func SendEndMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(false)
	websocketHelper.SendExecStatus(consts.ProgressEnd, wsMsg)

	return
}

func SendResultMsg(report agentDomain.Report, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendExecResult(report, wsMsg)

	return
}

func SendResult(data interface{}, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendExecResult(data, wsMsg)

	return
}

func SendCancelMsg(wsMsg websocket.Message) (err error) {
	SetRunning(false)

	websocketHelper.SendExecStatus(consts.ProgressEnd, &wsMsg)

	return
}

func SendAlreadyRunningMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("pls_stop_previous")
	websocketHelper.SendExecMsg(msg, agentDomain.ScenarioExecResult{ProgressStatus: consts.InProgress}, &wsMsg)
	_logUtils.Infof(msg)

	return
}

func SendExecMsg(log interface{}, wsMsg *websocket.Message) (err error) {
	SetRunning(true)
	msg := _i118Utils.Sprintf("exec")
	websocketHelper.SendExecMsg(msg, log, wsMsg)

	return
}

func SendErrorMsg(log agentDomain.ScenarioExecResult, wsMsg *websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("exec_fail")
	websocketHelper.SendExecMsg(msg, log, wsMsg)

	return
}
