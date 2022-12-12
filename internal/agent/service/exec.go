package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var (
	breakMap sync.Map
)

type ExecService struct {
}

func (s *ExecService) ExecScenario(req *agentExec.ExecReq, wsMsg *websocket.Message) (err error) {
	agentExec.InitScopeHierarchy(req.RootProcessor)
	s.SetEntityAndParent(req.RootProcessor)

	// start msg
	exec.SendStartMsg(wsMsg)

	// execution
	session := agentExec.NewSession(req, false, wsMsg)
	session.Run()

	// submit result
	s.SubmitResult(*session.RootProcessor.Result, req.RootProcessor.ScenarioId, req.ServerUrl, req.Token)
	s.sendSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	exec.SendEndMsg(wsMsg)

	return
}

func (s *ExecService) SubmitResult(result domain.Result, scenarioId uint, serverUrl, token string) (err error) {
	bodyBytes, _ := json.Marshal(result)
	req := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("scenarios/exec/submitResult/%d", scenarioId),
		Body:              string(bodyBytes),
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: token,
		},
	}
	resp, err := httpHelper.Post(req)

	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("submit result failed, response %v", resp)
		return
	}

	return
}

func (s *ExecService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	exec.SendCancelMsg(wsMsg)
	return
}

func (s *ExecService) SetEntityAndParent(root *agentExec.Processor) (err error) {
	processors := make([]*agentExec.Processor, 0)
	agentExec.GetProcessorList(root, &processors)

	processorMap := map[uint]*agentExec.Processor{}
	for _, processor := range processors {
		processorMap[processor.ID] = processor

		processor.UnmarshalEntity()
	}

	for _, obj := range processorMap {
		obj.Parent = processorMap[obj.ParentId]
	}

	return
}

func (s *ExecService) sendSubmitResult(rootId uint, wsMsg *websocket.Message) (err error) {
	result := domain.Result{
		ID:       -3,
		ParentId: int(rootId),
		Name:     "提交执行结果成功",
		//Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}
	exec.SendExecMsg(result, wsMsg)

	return
}
