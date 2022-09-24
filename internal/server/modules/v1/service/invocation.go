package service

import (
	"encoding/json"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"time"
)

type InvocationService struct {
	InvocationRepo *repo.InvocationRepo `inject:""`
}

func (s *InvocationService) ListByInterface(interfId int) (invocations []model.Invocation, err error) {
	invocations, err = s.InvocationRepo.List(interfId)

	return
}

func (s *InvocationService) GetLastResp(interfId int) (resp serverDomain.InvocationResponse, err error) {
	invocation, err := s.InvocationRepo.GetLast(interfId)

	json.Unmarshal([]byte(invocation.RespContent), &resp)

	return
}

func (s *InvocationService) Get(id int) (invocation model.Invocation, err error) {
	invocation, err = s.InvocationRepo.Get(uint(id))

	return
}

func (s *InvocationService) GetAsInterface(id int) (interf model.Interface, err error) {
	invocation, err := s.InvocationRepo.Get(uint(id))

	interfReq := serverDomain.InvocationRequest{}
	interfResp := serverDomain.InvocationResponse{}

	json.Unmarshal([]byte(invocation.ReqContent), &interfReq)
	json.Unmarshal([]byte(invocation.RespContent), &interfResp)

	copier.CopyWithOption(&interf, interfResp, copier.Option{DeepCopy: true})
	copier.CopyWithOption(&interf, interfReq, copier.Option{DeepCopy: true})

	interf.ID = invocation.InterfaceId

	return
}

func (s *InvocationService) Create(req serverDomain.InvocationRequest,
	resp serverDomain.InvocationResponse, projectId int) (invocation model.Invocation, err error) {
	invocation = model.Invocation{
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

func (s *InvocationService) Update(invocation model.Invocation) (err error) {

	return
}

func (s *InvocationService) Delete(id uint) (err error) {
	err = s.InvocationRepo.Delete(id)

	return
}

func (s *InvocationService) CopyValueFromRequest(invocation *model.Invocation, req serverDomain.InvocationRequest) (err error) {
	invocation.ID = req.Id

	copier.CopyWithOption(invocation, req, copier.Option{DeepCopy: true})

	return
}
