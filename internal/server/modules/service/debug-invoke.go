package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"time"
)

type DebugInvokeService struct {
	DebugRepo          *repo.DebugRepo          `inject:""`
	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`
	DebugInvokeRepo    *repo.DebugInvokeRepo    `inject:""`

	ProcessorInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`
	EndpointRepo           *repo.EndpointRepo           `inject:""`
	ScenarioProcessorRepo  *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo           *repo.ScenarioRepo           `inject:""`

	DebugSceneService     *DebugSceneService     `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`

	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
	VariableService   *VariableService   `inject:""`
	DatapoolService   *DatapoolService   `inject:""`
	EndpointService   *EndpointService   `inject:""`
}

func (s *DebugInvokeService) SubmitResult(req v1.SubmitDebugResultRequest) (err error) {
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

	}

	s.ExtractorService.ExtractInterface(req.Request.EndpointInterfaceId, serveId, processorId, scenarioId, req.Response, usedBy)
	s.CheckpointService.CheckInterface(req.Request.EndpointInterfaceId, req.Response, usedBy)

	_, err = s.Create(req.Request, req.Response, serveId, processorId, scenarioId, projectId)

	if err != nil {
		return
	}

	return
}

func (s *DebugInvokeService) Create(debugData v1.DebugData, resp v1.DebugResponse,
	serveId, processorId, scenarioId, projectId uint) (po model.DebugInvoke, err error) {

	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(debugData.EndpointInterfaceId)

	po = model.DebugInvoke{
		ServeId: serveId,

		ProcessorId: processorId,
		ScenarioId:  scenarioId,

		InvocationBase: model.InvocationBase{
			Name:                time.Now().Format("01-02 15:04:05"),
			EndpointInterfaceId: debugData.EndpointInterfaceId,
			DebugInterfaceId:    debugInterfaceId,
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

func (s *DebugInvokeService) ListByInterface(endpointInterfaceId uint) (invocations []model.DebugInvoke, err error) {
	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(endpointInterfaceId)

	invocations, err = s.DebugRepo.List(endpointInterfaceId, debugInterfaceId)

	return
}

func (s *DebugInvokeService) GetLastResp(endpointInterfaceId uint) (resp v1.DebugResponse, err error) {
	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(endpointInterfaceId)

	po, _ := s.DebugRepo.GetLast(endpointInterfaceId, debugInterfaceId)

	if po.ID > 0 {
		json.Unmarshal([]byte(po.RespContent), &resp)
	} else {
		resp = v1.DebugResponse{
			ContentLang: consts.LangHTML,
			Content:     "",
		}
	}

	return
}

func (s *DebugInvokeService) GetAsInterface(id int) (debugData v1.DebugData, interfResp v1.DebugResponse, err error) {
	invocation, err := s.DebugInvokeRepo.Get(uint(id))

	json.Unmarshal([]byte(invocation.ReqContent), &debugData)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	return
}

func (s *DebugInvokeService) Delete(id uint) (err error) {
	err = s.DebugRepo.Delete(id)

	return
}
