package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
)

type ExecScenarioService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *ExecScenarioService) ExecScenario(req *agentExec.ScenarioExecReq, wsMsg *websocket.Message) (err error) {
	consts.ServerUrl = req.ServerUrl
	consts.ServerToken = req.Token

	scenarioExecObj := s.RemoteService.GetScenarioToExec(req)

	session, err := s.Exec(scenarioExecObj, wsMsg)

	// submit result
	report, _ := s.RemoteService.SubmitScenarioResult(*session.RootProcessor.Result, scenarioExecObj.RootProcessor.ScenarioId,
		scenarioExecObj.ServerUrl, scenarioExecObj.Token)

	execUtils.SendResultMsg(report, session.WsMsg)
	s.sendSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func (s *ExecScenarioService) Exec(execObj *agentExec.ScenarioExecObj, wsMsg *websocket.Message) (
	session *agentExec.Session, err error) {

	// variables etc.
	agentExec.EnvToVariablesMap = execObj.EnvToVariablesMap
	agentExec.InterfaceToEnvMap = execObj.InterfaceToEnvMap
	agentExec.GlobalVars = execObj.GlobalEnvVars
	agentExec.GlobalParams = execObj.GlobalParamVars
	agentExec.DatapoolData = execObj.Datapools

	s.RestoreEntityFromRawAndSetParent(execObj.RootProcessor)

	agentExec.InitExecContext(execObj)
	agentExec.InitJsRuntime()

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// execution
	session = agentExec.NewSession(execObj, false, wsMsg)
	session.Run()

	return
}

func (s *ExecScenarioService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	execUtils.SendCancelMsg(wsMsg)
	return
}

func (s *ExecScenarioService) RestoreEntityFromRawAndSetParent(root *agentExec.Processor) (err error) {
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

func (s *ExecScenarioService) sendSubmitResult(rootId uint, wsMsg *websocket.Message) (err error) {
	result := agentDomain.ScenarioExecResult{
		ID:       -3,
		ParentId: int(rootId),
		Name:     "提交执行结果成功",
		//Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}

	execUtils.SendExecMsg(result, wsMsg)

	return
}
