package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
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
		logUtils.Info("****** execUuid is empty")
		return
	}

	/** 1. dealwith performance testing */
	if strings.Include(
		[]string{
			consts.JoinPerformanceTest.String(),
			consts.StartPerformanceTest.String(),
			consts.StopPerformanceTest.String(),
		}, act.String()) {

		go func() {
			defer errDefer(wsMsg)
			RunPerformanceTest(act, req.PerformanceTestExecReq, wsMsg)
		}()

		return
	}

	/** 2. dealwith other type of testing */
	isRunning := agentExec.GetIsRunning(execUuid)

	// stop exec
	if act == consts.ExecStop {
		StopExec(execUuid, wsMsg)

		return
	}

	// already running
	if isRunning && strings.Include([]string{
		consts.ExecScenario.String(),
		consts.ExecPlan.String(),
		consts.ExecCase.String(),
	}, act.String()) {
		execUtils.SendAlreadyRunningMsg(consts.Processor, wsMsg)
		return
	}

	// exec task
	go func() {
		defer errDefer(wsMsg)

		if act == consts.ExecScenario {
			RunScenario(&req.ScenarioExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecPlan {
			RunPlan(&req.PlanExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecCase {
			RunCases(&req.CasesExecReq, req.LocalVarsCache, wsMsg)

		} else if act == consts.ExecMessage {
			RunMessage(&req.MessageReq, wsMsg)
		}

		agentExec.ClearExecContext(execUuid)
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

	} else if req.PerformanceTestExecReq.Room != "" {
		ret = req.PerformanceTestExecReq.Room

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
		ptlog.Logf("err=%v, stack=%s\n", err, s)

		execUtils.SendErrorMsg(err, consts.Processor, wsMsg)
	}
}
