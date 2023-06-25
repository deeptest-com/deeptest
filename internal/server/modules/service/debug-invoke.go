package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"time"
)

type DebugInvokeService struct {
	DebugRepo          *repo.DebugRepo          `inject:""`
	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`
	DebugInvokeRepo    *repo.DebugInvokeRepo    `inject:""`

	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestInterfaceRepo     *repo.TestInterfaceRepo     `inject:""`

	DebugSceneService     *DebugSceneService     `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`

	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
	VariableService   *VariableService   `inject:""`
	DatapoolService   *DatapoolService   `inject:""`
	EndpointService   *EndpointService   `inject:""`
}

func (s *DebugInvokeService) SubmitResult(req domain.SubmitDebugResultRequest) (err error) {
	usedBy := req.Request.UsedBy
	var endpointId, serveId, processorId, scenarioId, projectId uint

	if usedBy == consts.InterfaceDebug {
		endpointId, serveId = s.DebugInterfaceService.GetEndpointAndServeIdForEndpointInterface(req.Request.EndpointInterfaceId)

		endpoint, _ := s.EndpointRepo.Get(endpointId)
		serveId = endpoint.ServeId
		projectId = endpoint.ProjectId

	} else if usedBy == consts.ScenarioDebug {
		processorId = req.Request.ScenarioProcessorId
		scenarioId = s.DebugInterfaceService.GetScenarioIdForDebugInterface(req.Request.ScenarioProcessorId)

		scenario, _ := s.ScenarioRepo.Get(scenarioId)
		scenarioId = scenario.ID
		projectId = scenario.ProjectId
	} else if usedBy == consts.TestDebug {
		testInterface, _ := s.TestInterfaceRepo.Get(req.Request.TestInterfaceId)

		serveId = testInterface.ServeId
		projectId = testInterface.ProjectId
	}

	s.ExtractorService.ExtractInterface(req.Request.DebugInterfaceId, req.Request.EndpointInterfaceId, serveId, processorId, scenarioId, req.Response, usedBy)
	s.CheckpointService.CheckInterface(req.Request.DebugInterfaceId, req.Request.EndpointInterfaceId, req.Request.ScenarioProcessorId, req.Response, usedBy)

	_, err = s.Create(req.Request, req.Response, serveId, processorId, scenarioId, projectId)

	if err != nil {
		return
	}

	return
}

func (s *DebugInvokeService) Create(debugData domain.DebugData, resp domain.DebugResponse,
	serveId, scenarioProcessorId, scenarioId, projectId uint) (po model.DebugInvoke, err error) {

	debugInterface, _ := s.DebugInterfaceRepo.Get(debugData.DebugInterfaceId)

	po = model.DebugInvoke{
		ServeId: serveId,

		ScenarioProcessorId: scenarioProcessorId,
		ScenarioId:          scenarioId,

		InvocationBase: model.InvocationBase{
			Name:                time.Now().Format("01-02 15:04:05"),
			EndpointInterfaceId: debugData.EndpointInterfaceId,
			DebugInterfaceId:    debugInterface.ID, // may be 0
			ProjectId:           projectId,
		},
	}

	bytesDebugData, _ := json.Marshal(debugData)
	po.ReqContent = string(bytesDebugData)

	bytesResp, _ := json.Marshal(resp)
	po.RespContent = string(bytesResp)

	err = s.DebugInvokeRepo.Save(&po)

	return
}

func (s *DebugInvokeService) ListByInterface(debugInterfaceId, endpointInterfaceId uint) (invocations []model.DebugInvoke, err error) {
	invocations, err = s.DebugRepo.List(debugInterfaceId, endpointInterfaceId)

	return
}

func (s *DebugInvokeService) GetLastResp(debugInterfaceId, endpointInterfaceId uint) (resp domain.DebugResponse, err error) {
	po, _ := s.DebugRepo.GetLast(debugInterfaceId, endpointInterfaceId)

	if po.ID > 0 {
		json.Unmarshal([]byte(po.RespContent), &resp)
	} else {
		resp = domain.DebugResponse{
			ContentLang: consts.LangHTML,
			Content:     "",
		}
	}

	return
}

func (s *DebugInvokeService) GetAsInterface(id int) (debugData domain.DebugData, interfResp domain.DebugResponse, err error) {
	invocation, err := s.DebugInvokeRepo.Get(uint(id))

	json.Unmarshal([]byte(invocation.ReqContent), &debugData)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	return
}

func (s *DebugInvokeService) Delete(id uint) (err error) {
	err = s.DebugRepo.Delete(id)

	return
}
