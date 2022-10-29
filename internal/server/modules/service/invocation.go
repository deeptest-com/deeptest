package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model2 "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"time"
)

type InvocationService struct {
	InvocationRepo *repo.InvocationRepo `inject:""`
	InterfaceRepo  *repo.InterfaceRepo  `inject:""`

	InterfaceService  *InterfaceService  `inject:""`
	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
	VariableService   *VariableService   `inject:""`
}

func (s *InvocationService) Invoke(req v1.InvocationRequest, projectId int) (
	resp v1.InvocationResponse, err error) {
	err = s.InterfaceService.UpdateByInvocation(req)
	if err != nil {
		return
	}

	reqNew, err := s.ReplaceEnvironmentExtractorAndExecVariables(req)
	if err != nil {
		return
	}

	resp, err = s.InterfaceService.Test(reqNew)
	if err != nil {
		return
	}

	interf, _ := s.InterfaceRepo.GetDetail(req.Id)
	s.ExtractorService.ExtractInterface(interf, resp, nil)
	s.CheckpointService.CheckInterface(interf, resp, nil)

	_, err = s.Create(req, resp, projectId)
	if err != nil {
		return
	}

	return
}

func (s *InvocationService) ListByInterface(interfId int) (invocations []model2.Invocation, err error) {
	invocations, err = s.InvocationRepo.List(interfId)

	return
}

func (s *InvocationService) GetLastResp(interfId int) (resp v1.InvocationResponse, err error) {
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

func (s *InvocationService) Get(id int) (invocation model2.Invocation, err error) {
	invocation, err = s.InvocationRepo.Get(uint(id))

	return
}

func (s *InvocationService) GetAsInterface(id int) (interf model2.Interface, err error) {
	invocation, err := s.InvocationRepo.Get(uint(id))

	interfReq := v1.InvocationRequest{}
	interfResp := v1.InvocationResponse{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfResp, copier.Option{DeepCopy: true})
	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *InvocationService) Create(req v1.InvocationRequest,
	resp v1.InvocationResponse, projectId int) (invocation model2.Invocation, err error) {
	invocation = model2.Invocation{
		Name:        time.Now().Format("01-02 15:04:05"),
		InterfaceId: req.Id,
		ProjectId:   uint(projectId),
	}

	bytesReq, _ := json.Marshal(req)
	invocation.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	invocation.RespContent = string(bytesReps)

	err = s.InvocationRepo.Save(&invocation)

	return
}

func (s *InvocationService) Update(invocation model2.Invocation) (err error) {

	return
}

func (s *InvocationService) Delete(id uint) (err error) {
	err = s.InvocationRepo.Delete(id)

	return
}

func (s *InvocationService) CopyValueFromRequest(invocation *model2.Invocation, req v1.InvocationRequest) (err error) {
	invocation.ID = req.Id

	copier.CopyWithOption(invocation, req, copier.Option{DeepCopy: true})

	return
}

func (s *InvocationService) ReplaceEnvironmentExtractorAndExecVariables(req v1.InvocationRequest) (
	ret v1.InvocationRequest, err error) {
	variableMap, _ := s.VariableService.GetVariablesByInterface(req.Id)
	agentExec.ReplaceAll(&req.BaseRequest, variableMap)

	ret = req

	return
}
