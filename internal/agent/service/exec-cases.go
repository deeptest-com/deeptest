package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(req *agentExec.CasesExecReq, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUUid, req.EnvironmentId)

	// reset exec
	agentExec.ResetStat()
	agentExec.ForceStopExec = false

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	for _, cs := range req.Cases {
		caseInterfaceExecObj := GetCaseToExec(req.BaseCaseId, cs, req.ServerUrl, req.Token)
		if caseInterfaceExecObj == nil {
			execUtils.SendEndMsg(wsMsg)
			return
		}

		// execution
		result, err1 := ExecCase(caseInterfaceExecObj, wsMsg)
		if err1 == nil {
			execUtils.SendErrorMsg(err1, consts.ProgressResult, wsMsg)
			return
		}

		// send result
		execUtils.SendExecMsg(*result, consts.ProgressResult, wsMsg)

		// stop if needed
		if agentExec.ForceStopExec {
			break
		}
	}

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func ExecCase(execObj *agentExec.CaseInterfaceExecObj, wsMsg *websocket.Message) (
	result *agentDomain.CaseExecResult, err error) {

	return
}
