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

	req.BaseUrl, req.ShareVariables, req.EnvVars, req.GlobalEnvVars, req.GlobalParamVars =
		s.DebugSceneService.LoadScene(req.EndpointId, req.InterfaceId, req.UsedBy)

	return
}

func (s *DebugService) SubmitResult(req v1.SubmitDebugResultRequest) (err error) {
	processorInterface, _ := s.ProcessorInterfaceRepo.GetDetail(req.Response.Id)

	s.ExtractorService.ExtractInterface(processorInterface.ID, req.Response, consts.ScenarioDebug)
	s.CheckpointService.CheckInterface(processorInterface.ID, req.Response, consts.ScenarioDebug)

	_, err = s.CreateForScenarioInterface(req.Request, req.Response, processorInterface.ProjectId)

	if err != nil {
		return
	}

	return
}

func (s *DebugService) CreateForScenarioInterface(req v1.DebugRequest,
	resp v1.DebugResponse, projectId uint) (invocation model.Debug, err error) {

	invocation = model.Debug{
		InvocationBase: model.InvocationBase{
			Name:        time.Now().Format("01-02 15:04:05"),
			InterfaceId: req.InterfaceId,
			ProjectId:   uint(projectId),
		},
	}

	bytesReq, _ := json.Marshal(req)
	invocation.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	invocation.RespContent = string(bytesReps)

	err = s.DebugRepo.Save(&invocation)

	return
}

func (s *DebugService) ListByInterface(interfId int) (invocations []model.Debug, err error) {
	invocations, err = s.DebugRepo.List(interfId)

	return
}

func (s *DebugService) GetLastResp(interfId uint) (resp v1.DebugResponse, err error) {
	invocation, _ := s.DebugRepo.GetLast(interfId)
	if invocation.ID > 0 {
		json.Unmarshal([]byte(invocation.RespContent), &resp)
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
