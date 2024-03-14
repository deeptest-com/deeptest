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

func SendInitializeMsg(data interface{}, wsMsg *websocket.Message) {
	websocketHelper.SendInitializeMsg(data, wsMsg)
}

func SendEndMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(false)
	websocketHelper.SendExecStatus(consts.ProgressEnd, wsMsg)

	return
}

func SendResultMsg(report agentExecDomain.Report, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendExecResult(report, wsMsg)

	return
}

func SendResult(data interface{}, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendExecResult(data, wsMsg)

	return
}

func SendCancelMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(false)

	websocketHelper.SendExecStatus(consts.ProgressEnd, wsMsg)

	return
}

func SendExecMsg(log interface{}, category consts.WsMsgCategory, wsMsg *websocket.Message) (err error) {
	SetRunning(true)
	msg := _i118Utils.Sprintf("exec")
	websocketHelper.SendExecMsg(msg, log, category, wsMsg)

	return
}

func SendErrorMsg(log interface{}, category consts.WsMsgCategory, wsMsg *websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("exec_fail")
	websocketHelper.SendExecMsg(msg, log, category, wsMsg)

	return
}

func SendAlreadyRunningMsg(wsMsg *websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("pls_stop_previous")

	websocketHelper.SendExecMsg(msg, agentExecDomain.ScenarioExecResult{ProgressStatus: consts.InProgress},
		consts.Processor, wsMsg)

	_logUtils.Infof(msg)

	return
}

func SendStatMsg(data agentExecDomain.InterfaceStat, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendStatInfo(data, wsMsg)

	return
}
