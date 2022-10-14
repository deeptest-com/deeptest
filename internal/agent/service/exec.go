package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var (
	breakMap sync.Map
)

type ExecService struct {
	ExecContextService  *business.ExecContext  `inject:""`
	ExecComm            *business.ExecComm     `inject:""`
	ExecIteratorService *business.ExecIterator `inject:""`
}

func (s *ExecService) ExecScenario(req *agentExec.ExecReq, wsMsg *websocket.Message) (err error) {
	session := agentExec.NewSession(req, false, wsMsg)
	session.Run()

	return
}

func (s *ExecService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	exec.SendCancelMsg(wsMsg)
	return
}
