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

type InvocationProcessorService struct {
	ProcessorInvocationRepo *repo.ProcessorInvocationRepo `inject:""`
	ProcessorInterfaceRepo  *repo.ProcessorInterfaceRepo  `inject:""`
	InterfaceRepo           *repo.InterfaceRepo           `inject:""`

	InterfaceService          *InterfaceService          `inject:""`
	ProcessorInterfaceService *ProcessorInterfaceService `inject:""`
	ExtractorService          *ExtractorService          `inject:""`
	CheckpointService         *CheckpointService         `inject:""`
	VariableService           *VariableService           `inject:""`
	DatapoolService           *DatapoolService           `inject:""`
}

func (s *InvocationProcessorService) LoadInterfaceExecData(req v1.InvocationRequest) (ret v1.InvocationRequest, err error) {
	err = s.ProcessorInterfaceService.UpdateByInvocation(req)
	if err != nil {
		return
	}

	ret, err = s.ReplaceEnvironmentAndExtractorVariables(req)

	return
}

func (s *InvocationProcessorService) SubmitInterfaceInvokeResult(req v1.SubmitInvocationResultRequest) (err error) {
	processorInterface, _ := s.ProcessorInterfaceRepo.GetDetail(req.Response.Id)

	s.ExtractorService.ExtractInterface(processorInterface.ID, req.Response, consts.UsedByScenario)
	s.CheckpointService.CheckInterface(processorInterface.ID, req.Response, consts.UsedByScenario)

	_, err = s.CreateForScenarioInterface(req.Request, req.Response, processorInterface.ProjectId)

	if err != nil {
		return
	}

	return
}

func (s *InvocationProcessorService) CreateForScenarioInterface(req v1.InvocationRequest,
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

	err = s.ProcessorInvocationRepo.Save(&invocation)

	return
}

func (s *InvocationProcessorService) ReplaceEnvironmentAndExtractorVariables(req v1.InvocationRequest) (
	ret v1.InvocationRequest, err error) {

	interf, _ := s.ProcessorInterfaceRepo.Get(req.Id)

	variableMap, _ := s.VariableService.GetVariablesByInterface(req.Id, consts.UsedByScenario)
	datapools, _ := s.DatapoolService.ListForExec(interf.ProjectId)

	agentExec.ReplaceAll(&req.BaseRequest, variableMap, datapools)

	ret = req

	return
}

func (s *InvocationProcessorService) ListByInterface(interfId int) (invocations []model.ProcessorInvocation, err error) {
	invocations, err = s.ProcessorInvocationRepo.List(interfId)

	return
}

func (s *InvocationProcessorService) GetLastResp(interfId int) (resp v1.InvocationResponse, err error) {
	invocation, _ := s.ProcessorInvocationRepo.GetLast(interfId)
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

func (s *InvocationProcessorService) GetAsInterface(id int) (interf model.ProcessorInterface, interfResp v1.InvocationResponse, err error) {
	invocation, err := s.ProcessorInvocationRepo.Get(uint(id))

	interfReq := v1.InvocationRequest{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *InvocationProcessorService) Delete(id uint) (err error) {
	err = s.ProcessorInvocationRepo.Delete(id)

	return
}
