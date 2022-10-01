package execHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/websocket"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

func SendStartMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(true)
	websocketHelper.SendExecResult(consts.ProgressInProgress, nil, wsMsg)

	return
}

func SendEndMsg(wsMsg *websocket.Message) (err error) {
	SetRunning(false)
	websocketHelper.SendExecResult(consts.ProgressEnd, nil, wsMsg)

	return
}

func SendResultMsg(report domain.ReportSimple, wsMsg *websocket.Message) (err error) {
	websocketHelper.SendExecResult(consts.Result, report, wsMsg)

	return
}

func SendCancelMsg(wsMsg websocket.Message) (err error) {
	SetRunning(false)

	websocketHelper.SendExecResult(consts.ProgressEnd, nil, &wsMsg)

	return
}

func SendAlreadyRunningMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("pls_stop_previous")
	websocketHelper.SendExecMsg(msg, domain.ExecLog{ProgressStatus: consts.InProgress}, &wsMsg)
	_logUtils.Infof(msg)

	return
}

func SendExecMsg(log domain.ExecLog, wsMsg *websocket.Message) (err error) {
	SetRunning(true)
	msg := _i118Utils.Sprintf("start_exec")
	websocketHelper.SendExecMsg(msg, log, wsMsg)
	//_logUtils.Infof("=== " + log.Name)

	return
}

func SendErrorMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	msg := _i118Utils.Sprintf("wrong_req_params", err.Error())
	websocketHelper.SendExecMsg(msg, domain.ExecLog{ProgressStatus: consts.Error}, &wsMsg)
	_logUtils.Infof(msg)

	return
}
