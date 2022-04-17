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

func (s *InvocationService) ListByInterface(interfId int) (requests []model.Invocation, err error) {
	requests, err = s.InvocationRepo.List(interfId)

	return
}

func (s *InvocationService) Get(reqId int) (interf model.Invocation, err error) {
	interf, err = s.InvocationRepo.Get(uint(reqId))

	return
}

func (s *InvocationService) LoadHistoryAsInterface(requestId int) (interf model.Interface, err error) {
	request, err := s.InvocationRepo.Get(uint(requestId))

	interfReq := serverDomain.InvocationRequest{}
	interfResp := serverDomain.InvocationResponse{}

	json.Unmarshal([]byte(request.ReqContent), &interfReq)
	json.Unmarshal([]byte(request.RespContent), &interfResp)

	copier.Copy(&interf, interfResp)
	copier.Copy(&interf, interfReq)

	interf.ID = request.InterfaceId

	return
}

func (s *InvocationService) Create(req serverDomain.InvocationRequest,
	resp serverDomain.InvocationResponse, projectId int) (request model.Invocation, err error) {
	request = model.Invocation{
		Name:        time.Now().Format("01-02 15:04:05"),
		InterfaceId: req.Id,
		ProjectId:   uint(projectId),
	}

	bytesReq, _ := json.Marshal(req)
	request.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	request.RespContent = string(bytesReps)

	err = s.InvocationRepo.Save(&request)

	return
}

func (s *InvocationService) Update(req model.Invocation) (err error) {

	return
}

func (s *InvocationService) Delete(reqId uint) (err error) {
	err = s.InvocationRepo.Delete(reqId)

	return
}

func (s *InvocationService) CopyValueFromRequest(interf *model.Invocation, req serverDomain.InvocationRequest) (err error) {
	interf.ID = req.Id

	copier.Copy(interf, req)

	return
}
