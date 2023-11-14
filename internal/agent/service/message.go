package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

func RunMessage(req *agentExec.MessageExecReq, wsMsg *websocket.Message) (err error) {
	// start msg
	agentExec.SetIsRunning(req.ExecUuid, true)
	err = execUtils.SendStartMsg(wsMsg)
	if err != nil {
		return err
	}

	result := GetMessageToExec(req)

	// submit result
	err = execUtils.SendResult(result, wsMsg)
	if err != nil {
		return err
	}

	err = sendSubmitResult(req.ExecUuid, wsMsg)
	if err != nil {
		return err
	}

	// end msg
	agentExec.SetIsRunning(req.ExecUuid, false)
	err = execUtils.SendEndMsg(wsMsg)
	if err != nil {
		return err
	}

	return
}

func sendSubmitResult(execUuid string, wsMsg *websocket.Message) (err error) {
	result := agentDomain.MessageExecResult{
		ExecUuid: execUuid,
		Name:     "提交执行结果成功",
	}
	execUtils.SendExecMsg(result, consts.Processor, wsMsg)

	return
}
