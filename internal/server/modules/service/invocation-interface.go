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

type InvocationInterfaceService struct {
	InvocationRepo          *repo.InvocationRepo          `inject:""`
	ProcessorInvocationRepo *repo.ProcessorInvocationRepo `inject:""`
	InterfaceRepo           *repo.InterfaceRepo           `inject:""`
	ScenarioInterfaceRepo   *repo.ProcessorInterfaceRepo  `inject:""`

	InterfaceService         *InterfaceService          `inject:""`
	ScenarioInterfaceService *ProcessorInterfaceService `inject:""`
	ExtractorService         *ExtractorService          `inject:""`
	CheckpointService        *CheckpointService         `inject:""`
	VariableService          *VariableService           `inject:""`
	DatapoolService          *DatapoolService           `inject:""`
}

func (s *InvocationInterfaceService) LoadInterfaceExecData(req v1.DebugRequest) (ret v1.DebugRequest, err error) {
	err = s.InterfaceService.UpdateByInvocation(req)
	if err != nil {
		return
	}

	ret, err = s.ReplaceEnvironmentAndExtractorVariables(req)

	return
}

func (s *InvocationInterfaceService) SubmitInterfaceInvokeResult(req v1.SubmitDebugResultRequest) (err error) {
	interf, _ := s.InterfaceRepo.GetDetail(req.Response.Id)

	s.ExtractorService.ExtractInterface(interf.ID, req.Response, consts.InterfaceDebug)
	s.CheckpointService.CheckInterface(interf.ID, req.Response, consts.InterfaceDebug)

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

func (s *InvocationInterfaceService) GetLastResp(interfId int) (resp v1.DebugResponse, err error) {
	invocation, _ := s.InvocationRepo.GetLast(interfId)
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

func (s *InvocationInterfaceService) Get(id int) (invocation model.Invocation, err error) {
	invocation, err = s.InvocationRepo.Get(uint(id))

	return
}

func (s *InvocationInterfaceService) GetAsInterface(id int) (interf model.Interface, interfResp v1.DebugResponse, err error) {
	invocation, err := s.InvocationRepo.Get(uint(id))

	interfReq := v1.DebugRequest{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *InvocationInterfaceService) CreateForInterface(req v1.DebugRequest,
	resp v1.DebugResponse, projectId uint) (invocation model.Invocation, err error) {
	invocation = model.Invocation{
		InvocationBase: model.InvocationBase{
			Name:        time.Now().Format("01-02 15:04:05"),
			InterfaceId: req.InterfaceId,
			ProjectId:   projectId,
		},
	}

	bytesReq, _ := json.Marshal(req)
	invocation.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	invocation.RespContent = string(bytesReps)

	err = s.InvocationRepo.Save(&invocation)

	return
}

func (s *InvocationInterfaceService) CreateForScenarioInterface(req v1.DebugRequest,
	resp v1.DebugResponse, projectId uint) (invocation model.ProcessorInvocation, err error) {

	invocation = model.ProcessorInvocation{
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

	err = s.ProcessorInvocationRepo.Save(&invocation)

	return
}

func (s *InvocationInterfaceService) Delete(id uint) (err error) {
	err = s.InvocationRepo.Delete(id)

	return
}

func (s *InvocationInterfaceService) CopyValueFromRequest(invocation *model.Invocation, req v1.DebugRequest) (err error) {
	invocation.ID = req.InterfaceId

	copier.CopyWithOption(invocation, req, copier.Option{DeepCopy: true})

	return
}

func (s *InvocationInterfaceService) ReplaceEnvironmentAndExtractorVariables(req v1.DebugRequest) (
	ret v1.DebugRequest, err error) {

	//interf, _ := s.InterfaceRepo.Get(req.InterfaceId)
	//
	//req.Environment, _ = s.VariableService.GetEnvVarsByInterface(req.InterfaceId, consts.InterfaceDebug)
	//req.Variables, _ = s.VariableService.GetShareVarsByInterface(req.InterfaceId, consts.InterfaceDebug)
	//req.Datapools, _ = s.DatapoolService.ListForExec(interf.ProjectId)

	ret = req

	return
}
