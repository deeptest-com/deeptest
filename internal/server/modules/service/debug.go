package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"time"
)

type DebugService struct {
	DebugRepo              *repo.DebugRepo              `inject:""`
	DebugInterfaceRepo     *repo.InterfaceRepo          `inject:""`
	ProcessorInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`
	EndpointRepo           *repo.EndpointRepo           `inject:""`
	ScenarioProcessorRepo  *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo           *repo.ScenarioRepo           `inject:""`

	DebugSceneService *DebugSceneService `inject:""`

	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
	VariableService   *VariableService   `inject:""`
	DatapoolService   *DatapoolService   `inject:""`
	EndpointService   *EndpointService   `inject:""`
}

func (s *DebugService) LoadData(call v1.DebugCall) (req v1.DebugRequest, err error) {
	isInterfaceHasDebugRecord, err := s.DebugRepo.IsInterfaceHasDebug(call.InterfaceId)

	if isInterfaceHasDebugRecord {
		req, err = s.GetLastReq(call.InterfaceId)
	} else {
		req, err = s.EndpointService.GenerateReq(call.InterfaceId, call.EndpointId)
	}

	req.BaseUrl, req.ShareVars, req.EnvVars, req.GlobalEnvVars, req.GlobalParamVars =
		s.DebugSceneService.LoadScene(req.InterfaceId, req.EndpointId, req.ProcessorId, req.UsedBy)

	return
}

func (s *DebugService) SubmitResult(req v1.SubmitDebugResultRequest) (err error) {
	usedBy := req.Request.UsedBy
	var serveId, processorId, scenarioId, projectId uint

	if usedBy == consts.InterfaceDebug {
		endpoint, _ := s.EndpointRepo.Get(req.Request.EndpointId)
		serveId = endpoint.ServeId
		projectId = endpoint.ProjectId
	} else if usedBy == consts.ScenarioDebug {
		processor, _ := s.ScenarioProcessorRepo.Get(req.Request.ProcessorId)
		scenario, _ := s.ScenarioRepo.Get(processor.ScenarioId)
		processorId = processor.ID
		scenarioId = scenario.ID
		projectId = scenario.ProjectId
	}

	s.ExtractorService.ExtractInterface(req.Request.InterfaceId, serveId, processorId, req.Response, usedBy)
	s.CheckpointService.CheckInterface(req.Request.InterfaceId, req.Response, usedBy)

	_, err = s.Create(req.Request, req.Response, serveId, processorId, scenarioId, projectId)

	if err != nil {
		return
	}

	return
}

func (s *DebugService) Create(req v1.DebugRequest, resp v1.DebugResponse,
	serveId, processorId, scenarioId, projectId uint) (po model.Debug, err error) {

	po = model.Debug{
		ServeId: serveId,

		ProcessorId: processorId,
		ScenarioId:  scenarioId,

		InvocationBase: model.InvocationBase{
			Name:        time.Now().Format("01-02 15:04:05"),
			InterfaceId: req.InterfaceId,
			ProjectId:   projectId,
		},
	}

	bytesReq, _ := json.Marshal(req)
	po.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	po.RespContent = string(bytesReps)

	err = s.DebugRepo.Save(&po)

	return
}

func (s *DebugService) ListByInterface(interfId int) (invocations []model.Debug, err error) {
	invocations, err = s.DebugRepo.List(interfId)

	return
}

func (s *DebugService) GetLastResp(interfaceId uint) (resp v1.DebugResponse, err error) {
	po, _ := s.DebugRepo.GetLast(interfaceId)

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

func (s *DebugService) GetLastReq(interfId uint) (req v1.DebugRequest, err error) {
	invocation, _ := s.DebugRepo.GetLast(interfId)

	if invocation.ID > 0 {
		json.Unmarshal([]byte(invocation.ReqContent), &req)
	} else {
		req = v1.DebugRequest{}
	}

	return
}

func (s *DebugService) GetAsInterface(id int) (interf model.ProcessorInterface, interfResp v1.DebugResponse, err error) {
	invocation, err := s.DebugRepo.Get(uint(id))

	interfReq := v1.DebugRequest{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *DebugService) Delete(id uint) (err error) {
	err = s.DebugRepo.Delete(id)

	return
}
