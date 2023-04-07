package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

type MessageService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *MessageService) ExecMessage(req *agentExec.MessageExecReq, wsMsg *websocket.Message) (err error) {
	consts.ServerUrl = req.ServerUrl
	consts.ServerToken = req.Token

	// start msg
	err = execUtils.SendStartMsg(wsMsg)
	if err != nil {
		return err
	}

	result := s.RemoteService.GetMessageToExec(req)

	// submit result
	err = execUtils.SendResult(result, wsMsg)
	if err != nil {
		return err
	}

	err = s.sendSubmitResult(req.UserId, wsMsg)
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

func (s *MessageService) sendSubmitResult(userId uint, wsMsg *websocket.Message) (err error) {
	result := agentDomain.MessageExecResult{
		UserId: userId,
		Name:   "提交执行结果成功",
	}
	execUtils.SendExecMsg(result, wsMsg)

	return
}
