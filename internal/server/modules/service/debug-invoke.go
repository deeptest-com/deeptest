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
	DebugRepo              *repo.DebugRepo              `inject:""`
	DebugInterfaceRepo     *repo.DebugInterfaceRepo     `inject:""`
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
		if req.Request.EndpointInterfaceId > 0 {
			endpointId, serveId = s.DebugInterfaceService.GetEndpointAndServeIdForEndpointInterface(req.Request.EndpointInterfaceId)
		} else if req.Request.DebugInterfaceId > 0 {
			endpointId, serveId = s.DebugInterfaceService.GetEndpointAndServeIdForDebugInterface(req.Request.DebugInterfaceId)
		}

		endpoint, _ := s.EndpointRepo.Get(endpointId)
		serveId = endpoint.ServeId
		projectId = endpoint.ProjectId

	} else if usedBy == consts.ScenarioDebug {
		processorId = req.Request.ProcessorId
		scenarioId = s.DebugInterfaceService.GetScenarioIdForDebugInterface(req.Request.ProcessorId)

		scenario, _ := s.ScenarioRepo.Get(scenarioId)
		scenarioId = scenario.ID
		projectId = scenario.ProjectId

	}

	s.ExtractorService.ExtractInterface(req.Request.EndpointInterfaceId, serveId, processorId, req.Response, usedBy)
	s.CheckpointService.CheckInterface(req.Request.EndpointInterfaceId, req.Response, usedBy)

	_, err = s.Create(req.Request, req.Response, serveId, processorId, scenarioId, projectId)

	if err != nil {
		return
	}

	return
}

func (s *DebugInvokeService) Create(req v1.DebugData, resp v1.DebugResponse,
	serveId, processorId, scenarioId, projectId uint) (po model.DebugInvoke, err error) {

	po = model.DebugInvoke{
		ServeId: serveId,

		ProcessorId: processorId,
		ScenarioId:  scenarioId,

		InvocationBase: model.InvocationBase{
			Name:                time.Now().Format("01-02 15:04:05"),
			EndpointInterfaceId: req.EndpointInterfaceId,
			DebugInterfaceId:    req.DebugInterfaceId,
			ProjectId:           projectId,
		},
	}

	bytesReq, _ := json.Marshal(req)
	po.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	po.RespContent = string(bytesReps)

	err = s.DebugRepo.Save(&po)

	return
}

func (s *DebugInvokeService) ListByInterface(endpointInterfaceId, debugInterfaceId int) (invocations []model.DebugInvoke, err error) {
	invocations, err = s.DebugRepo.List(endpointInterfaceId, debugInterfaceId)

	return
}

func (s *DebugInvokeService) GetLastResp(endpointInterfaceId, debugInterfaceId int) (resp v1.DebugResponse, err error) {
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

//func (s *DebugInvokeService) GetLastReq(interfId uint) (req v1.DebugData, err error) {
//	invocation, _ := s.DebugRepo.GetLast(interfId)
//
//	if invocation.ID > 0 {
//		json.Unmarshal([]byte(invocation.ReqContent), &req)
//	} else {
//		req = v1.DebugData{}
//	}
//
//	return
//}

func (s *DebugInvokeService) GetAsInterface(id int) (interf model.ProcessorInterface, interfResp v1.DebugResponse, err error) {
	//invocation, err := s.DebugRepo.Get(uint(id))
	//
	//interfReq := v1.DebugData{}
	//
	//json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	//json.Unmarshal([]byte(invocation.RespContent), &interfResp)
	//
	//copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})
	//
	//interf.ID = invocation.InterfaceId

	return
}

func (s *DebugInvokeService) Delete(id uint) (err error) {
	err = s.DebugRepo.Delete(id)

	return
}
