package service

import (
	"fmt"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
	"runtime/debug"
)

func StartExec(req agentDomain.WsReq, wsMsg *websocket.Message) (err error) {
	act := req.Act

	execUuid := getExecUuid(req)
	isRunning := agentExec.GetIsRunning(execUuid)

	// stop exec
	if act == consts.ExecStop {
		StopExec(execUuid, wsMsg)

		return
	}

	// already running
	if isRunning && (act == consts.ExecStart) {
		if req.ScenarioExecReq.ScenarioId > 0 {
			execUtils.SendAlreadyRunningMsg(req.ScenarioExecReq.ScenarioId, consts.Processor, wsMsg)
		}
		return
	}

	// exec task
	go func() {
		defer errDefer(req, wsMsg)

		if act == consts.ExecScenario {
			RunScenario(&req.ScenarioExecReq, wsMsg)

		} else if act == consts.ExecPlan {
			RunPlan(&req.PlanExecReq, wsMsg)

		} else if act == consts.ExecCase {
			RunCases(&req.CasesExecReq, wsMsg)

		} else if act == consts.ExecMessage {
			RunMessage(&req.MessageReq, wsMsg)
		}

		agentExec.ClearExec(execUuid)
	}()

	return
}

func getExecUuid(req agentDomain.WsReq) (ret string) {
	if req.ScenarioExecReq.ScenarioId > 0 {
		ret = req.ScenarioExecReq.ExecUuid
	} else if req.PlanExecReq.PlanId > 0 {
		ret = req.PlanExecReq.ExecUuid
	} else if len(req.CasesExecReq.Cases) > 0 {
		ret = req.CasesExecReq.ExecUuid
	}

	return
}

func StopExec(execUuid string, wsMsg *websocket.Message) (err error) {
	agentExec.SetForceStopExec(execUuid, true)

	agentExec.SetIsRunning(execUuid, false)
	execUtils.SendCancelMsg(wsMsg)

	return
}

func SendExecErr(req agentDomain.WsReq, err error, wsMsg *websocket.Message) {
	if req.ScenarioExecReq.ScenarioId > 0 {
		sendScenarioErr(err, wsMsg)
	}
}

func errDefer(req agentDomain.WsReq, wsMsg *websocket.Message) {
	wsMsgErr := recover()

	if wsMsgErr != nil {
		s := string(debug.Stack())
		fmt.Printf("err=%v, stack=%s\n", wsMsgErr, s)

		SendExecErr(req, fmt.Errorf("%+v", wsMsgErr), wsMsg)
	}
}
