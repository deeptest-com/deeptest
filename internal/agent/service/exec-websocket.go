package service

import (
	"fmt"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"runtime/debug"
)

func StartExec(req agentDomain.WsReq, wsMsg *websocket.Message) (err error) {
	act := req.Act
	execUuid := getExecUuid(req)
	if execUuid == "" {
		logUtils.Info("****** execUuid is empty")
		logUtils.Infof("%v", req)
		return
	}

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
		defer errDefer(wsMsg)

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
	if req.ScenarioExecReq.ExecUuid != "" {
		ret = req.ScenarioExecReq.ExecUuid

	} else if req.PlanExecReq.ExecUuid != "" {
		ret = req.PlanExecReq.ExecUuid

	} else if req.CasesExecReq.ExecUuid != "" {
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

func errDefer(wsMsg *websocket.Message) {
	err := recover()

	if err != nil {
		s := string(debug.Stack())
		fmt.Printf("err=%v, stack=%s\n", err, s)

		execUtils.SendErrorMsg(err, consts.Processor, wsMsg)
	}
}
