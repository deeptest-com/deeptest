package service

import (
	"context"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

func RunMessage(ctx context.Context, req *agentExec.MessageExecReq, wsMsg *websocket.Message) (err error) {
	// start msg
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

	err = sendSubmitResult(req.UserId, wsMsg)
	if err != nil {
		return err
	}

	// end msg
	err = execUtils.SendEndMsg(wsMsg)
	if err != nil {
		return err
	}

	return
}

func sendSubmitResult(userId uint, wsMsg *websocket.Message) (err error) {
	result := agentDomain.MessageExecResult{
		UserId: userId,
		Name:   "提交执行结果成功",
	}
	execUtils.SendExecMsg(result, consts.Processor, wsMsg)

	return
}
