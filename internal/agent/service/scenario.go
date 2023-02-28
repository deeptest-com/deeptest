package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var (
	breakMap sync.Map
)

type ScenarioService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *ScenarioService) ExecScenario(req *agentExec.ProcessorExecReq, wsMsg *websocket.Message) (err error) {
	consts.ServerUrl = req.ServerUrl
	consts.ServerToken = req.Token

	scenarioExecReq := s.RemoteService.GetScenarioToExec(req)
	agentExec.Variables = scenarioExecReq.Variables
	agentExec.DatapoolData = scenarioExecReq.Datapools

	s.RestoreEntityFromRawAndSetParent(scenarioExecReq.RootProcessor)

	agentExec.InitExecContext(scenarioExecReq)
	agentExec.InitJsRuntime()

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// execution
	session := agentExec.NewSession(scenarioExecReq, false, wsMsg)
	session.Run()

	// submit result
	s.RemoteService.SubmitScenarioResult(*session.RootProcessor.Result, scenarioExecReq.RootProcessor.ScenarioId, scenarioExecReq.ServerUrl, scenarioExecReq.Token)
	s.sendSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func (s *ScenarioService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	execUtils.SendCancelMsg(wsMsg)
	return
}

func (s *ScenarioService) RestoreEntityFromRawAndSetParent(root *agentExec.Processor) (err error) {
	processors := make([]*agentExec.Processor, 0)

	agentExec.GetProcessorList(root, &processors)

	processorMap := map[uint]*agentExec.Processor{}
	for _, processor := range processors {
		processorMap[processor.ID] = processor

		processor.RestoreEntity()
	}

	for _, obj := range processorMap {
		obj.Parent = processorMap[obj.ParentId]
	}

	return
}

func (s *ScenarioService) sendSubmitResult(rootId uint, wsMsg *websocket.Message) (err error) {
	result := agentDomain.Result{
		ID:       -3,
		ParentId: int(rootId),
		Name:     "提交执行结果成功",
		//Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}
	execUtils.SendExecMsg(result, wsMsg)

	return
}
