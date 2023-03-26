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
	InterfaceRepo          *repo.InterfaceRepo          `inject:""`
	ProcessorInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`

	InterfaceService          *InterfaceService          `inject:""`
	ProcessorInterfaceService *ProcessorInterfaceService `inject:""`
	ExtractorService          *ExtractorService          `inject:""`
	CheckpointService         *CheckpointService         `inject:""`
	VariableService           *VariableService           `inject:""`
	DatapoolService           *DatapoolService           `inject:""`
}

func (s *DebugService) LoadData(req v1.DebugRequest) (ret v1.DebugRequest, err error) {
	err = s.ProcessorInterfaceService.UpdateByInvocation(req)
	if err != nil {
		return
	}

	ret, err = s.ReplaceEnvironmentAndExtractorVariables(req)

	return
}

func (s *DebugService) SubmitResult(req v1.SubmitDebugResultRequest) (err error) {
	processorInterface, _ := s.ProcessorInterfaceRepo.GetDetail(req.Response.Id)

	s.ExtractorService.ExtractInterface(processorInterface.ID, req.Response, consts.UsedByScenario)
	s.CheckpointService.CheckInterface(processorInterface.ID, req.Response, consts.UsedByScenario)

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
			InterfaceId: req.Id,
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

func (s *DebugService) ReplaceEnvironmentAndExtractorVariables(req v1.DebugRequest) (
	ret v1.DebugRequest, err error) {

	interf, _ := s.ProcessorInterfaceRepo.Get(req.Id)

	req.Environment, _ = s.VariableService.GetEnvironmentVariablesByInterface(req.Id, req.UsedBy)
	req.Variables, _ = s.VariableService.GetVariablesByInterface(req.Id, req.UsedBy)
	req.Datapools, _ = s.DatapoolService.ListForExec(interf.ProjectId)

	ret = req

	return
}

func (s *DebugService) ListByInterface(interfId int) (invocations []model.Debug, err error) {
	invocations, err = s.DebugRepo.List(interfId)

	return
}

func (s *DebugService) GetLastResp(interfId int) (resp v1.DebugResponse, err error) {
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
