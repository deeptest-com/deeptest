package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"time"
)

type InvocationInterfaceService struct {
	InvocationRepo         *repo.InvocationRepo          `inject:""`
	ScenarioInvocationRepo *repo.ProcessorInvocationRepo `inject:""`
	InterfaceRepo          *repo.InterfaceRepo           `inject:""`
	ScenarioInterfaceRepo  *repo.ProcessorInterfaceRepo  `inject:""`

	InterfaceService         *InterfaceService          `inject:""`
	ScenarioInterfaceService *ProcessorInterfaceService `inject:""`
	ExtractorService         *ExtractorService          `inject:""`
	CheckpointService        *CheckpointService         `inject:""`
	VariableService          *VariableService           `inject:""`
	DatapoolService          *DatapoolService           `inject:""`
}

func (s *InvocationInterfaceService) LoadInterfaceExecData(req v1.InvocationRequest) (ret v1.InvocationRequest, err error) {
	err = s.InterfaceService.UpdateByInvocation(req)
	if err != nil {
		return
	}

	ret, err = s.ReplaceEnvironmentAndExtractorVariables(req)

	return
}

func (s *InvocationInterfaceService) SubmitInterfaceInvokeResult(req v1.SubmitInvocationResultRequest) (err error) {
	interf, _ := s.InterfaceRepo.GetDetail(req.Response.Id)

	s.ExtractorService.ExtractInterface(interf.ID, req.Response, consts.UsedByInterface)
	s.CheckpointService.CheckInterface(interf.ID, req.Response, consts.UsedByInterface)

	_, err = s.CreateForInterface(req.Request, req.Response, interf.ProjectId)

	if err != nil {
		return
	}

	return
}

func (s *InvocationInterfaceService) ListByInterface(interfId int) (invocations []model.Invocation, err error) {
	invocations, err = s.InvocationRepo.List(interfId)

	return
}

func (s *InvocationInterfaceService) GetLastResp(interfId int) (resp v1.InvocationResponse, err error) {
	invocation, _ := s.InvocationRepo.GetLast(interfId)
	if invocation.ID > 0 {
		json.Unmarshal([]byte(invocation.RespContent), &resp)
	} else {
		resp = v1.InvocationResponse{
			ContentLang: consts.LangHTML,
			Content:     "",
		}
	}

	return
}

func (s *InvocationInterfaceService) Get(id int) (invocation model.Invocation, err error) {
	invocation, err = s.InvocationRepo.Get(uint(id))

	return
}

func (s *InvocationInterfaceService) GetAsInterface(id int) (interf model.Interface, interfResp v1.InvocationResponse, err error) {
	invocation, err := s.InvocationRepo.Get(uint(id))

	interfReq := v1.InvocationRequest{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *InvocationInterfaceService) CreateForInterface(req v1.InvocationRequest,
	resp v1.InvocationResponse, projectId uint) (invocation model.Invocation, err error) {
	invocation = model.Invocation{
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

	err = s.InvocationRepo.Save(&invocation)

	return
}

func (s *InvocationInterfaceService) CreateForScenarioInterface(req v1.InvocationRequest,
	resp v1.InvocationResponse, projectId uint) (invocation model.ProcessorInvocation, err error) {

	invocation = model.ProcessorInvocation{
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

	err = s.ScenarioInvocationRepo.Save(&invocation)

	return
}

func (s *InvocationInterfaceService) Delete(id uint) (err error) {
	err = s.InvocationRepo.Delete(id)

	return
}

func (s *InvocationInterfaceService) CopyValueFromRequest(invocation *model.Invocation, req v1.InvocationRequest) (err error) {
	invocation.ID = req.Id

	copier.CopyWithOption(invocation, req, copier.Option{DeepCopy: true})

	return
}

func (s *InvocationInterfaceService) ReplaceEnvironmentAndExtractorVariables(req v1.InvocationRequest) (
	ret v1.InvocationRequest, err error) {

	interf, _ := s.InterfaceRepo.Get(req.Id)

	variableMap, _ := s.VariableService.GetVariablesByInterface(req.Id, consts.UsedByInterface)
	datapools, _ := s.DatapoolService.ListForExec(interf.ProjectId)

	agentExec.ReplaceAll(&req.BaseRequest, variableMap, datapools)

	ret = req

	return
}
