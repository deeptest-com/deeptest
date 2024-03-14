package service

import (
	"fmt"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/savsgio/gotils/strings"
	"runtime/debug"
)

func StartExec(req agentDomain.WsReq, wsMsg *websocket.Message) (err error) {
	act := req.Act
	execUuid := getExecUuid(req)
	if execUuid == "" {
		logUtils.Info("!!!!!! execUuid is empty")
		return
	}

	// stop exec
	if act == consts.ExecStop {
		StopExec(execUuid, wsMsg)
		return
	}

	// is running
	ctx, _ := agentExec.GetExecCtx(execUuid)
	if ctx != nil && (strings.Include([]string{consts.ExecScenario.String(), consts.ExecPlan.String(), consts.ExecCase.String()}, act.String())) {
		IsRunning(wsMsg)
		return
	}

	// exec task
	go func() {
		defer errDefer(wsMsg)

		agentExec.InitUserExecContext(execUuid)
		ctx, _ := agentExec.GetExecCtx(execUuid)

		if act == consts.ExecScenario {
			RunScenario(ctx, &req.ScenarioExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecPlan {
			RunPlan(ctx, &req.PlanExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecCase {
			RunCases(ctx, &req.CasesExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecMessage {
			RunMessage(ctx, &req.MessageReq, wsMsg)
		}

		agentExec.CloseExecCtx(execUuid)
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
	agentExec.CloseExecCtx(execUuid)
	execUtils.SendCancelMsg(wsMsg)

	return
}

func IsRunning(wsMsg *websocket.Message) (err error) {
	execUtils.SendAlreadyRunningMsg(wsMsg)

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
