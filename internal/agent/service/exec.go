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
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
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

func (s *ExecService) ExecScenario(execReq *agentExec.ProcessorExecReq, wsMsg *websocket.Message) (err error) {
	execObj := s.getScenarioToExec(execReq)
	s.RestoreEntityFromRawAndSetParent(execObj.RootProcessor)

	agentExec.InitScopeHierarchy(execObj.RootProcessor)

	// start msg
	exec.SendStartMsg(wsMsg)

	// execution
	session := agentExec.NewSession(execObj, false, wsMsg)
	session.Run()

	// submit result
	s.SubmitResult(*session.RootProcessor.Result, execObj.RootProcessor.ScenarioId, execObj.ServerUrl, execObj.Token)
	s.sendSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	exec.SendEndMsg(wsMsg)

	return
}

func (s *ExecService) getScenarioToExec(req *agentExec.ProcessorExecReq) (ret *agentExec.ProcessorExecObj) {
	url := "scenarios/exec/loadExecScenario"

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
		Params: []v1.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.ScenarioId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	ret.ServerUrl = req.ServerUrl
	ret.Token = req.Token

	return
}

func (s *ExecService) SubmitResult(result domain.Result, scenarioId uint, serverUrl, token string) (err error) {
	bodyBytes, _ := json.Marshal(result)
	req := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("scenarios/exec/submitResult/%d", scenarioId),
		Body:              string(bodyBytes),
		BodyType:          consts.ContentTypeJSON,
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

	ret := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &ret)

	if ret.Code != 0 {
		logUtils.Infof("submit result failed, response %v", resp.Content)
		return
	}

	return
}

func (s *ExecService) CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	exec.SendCancelMsg(wsMsg)
	return
}

func (s *ExecService) RestoreEntityFromRawAndSetParent(root *agentExec.Processor) (err error) {
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
