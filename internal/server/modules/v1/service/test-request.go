package service

import (
	"encoding/json"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"time"
)

type TestRequestService struct {
	TestRequestRepo *repo.TestRequestRepo `inject:""`
}

func NewTestRequestService() *TestRequestService {
	return &TestRequestService{}
}

func (s *TestRequestService) ListByInterface(interfId int) (requests []model.TestRequest, err error) {
	requests, err = s.TestRequestRepo.List(interfId)

	return
}

func (s *TestRequestService) Get(reqId int) (interf model.TestRequest, err error) {
	interf, err = s.TestRequestRepo.Get(uint(reqId))

	return
}

func (s *TestRequestService) LoadHistoryAsInterface(requestId int) (interf model.TestInterface, err error) {
	request, err := s.TestRequestRepo.Get(uint(requestId))

	interfReq := serverDomain.TestRequest{}
	interfResp := serverDomain.TestResponse{}

	json.Unmarshal([]byte(request.ReqContent), &interfReq)
	json.Unmarshal([]byte(request.RespContent), &interfResp)

	copier.Copy(&interf, interfResp)
	copier.Copy(&interf, interfReq)

	interf.ID = request.InterfaceId

	return
}

func (s *TestRequestService) Create(req serverDomain.TestRequest,
	resp serverDomain.TestResponse, projectId int) (request model.TestRequest, err error) {
	request = model.TestRequest{
		Name:        time.Now().Format("01-02 15:04:05"),
		InterfaceId: req.Id,
		ProjectId:   uint(projectId),
	}

	bytesReq, _ := json.Marshal(req)
	request.ReqContent = string(bytesReq)

	bytesReps, _ := json.Marshal(resp)
	request.RespContent = string(bytesReps)

	err = s.TestRequestRepo.Save(&request)

	return
}

func (s *TestRequestService) Update(req model.TestRequest) (err error) {

	return
}

func (s *TestRequestService) Delete(reqId uint) (err error) {
	err = s.TestRequestRepo.Delete(reqId)

	return
}

func (s *TestRequestService) CopyValueFromRequest(interf *model.TestRequest, req serverDomain.TestRequest) (err error) {
	interf.ID = req.Id

	copier.Copy(interf, req)

	return
}
