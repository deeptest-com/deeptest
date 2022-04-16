package service

import (
	"encoding/json"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
)

type TestRequestService struct {
	TestRequestRepo *repo.TestRequestRepo `inject:""`
}

func NewTestRequestService() *TestRequestService {
	return &TestRequestService{}
}

func (s *TestRequestService) ListByInterface(interfId int) (requests model.TestRequest, err error) {

	return
}

func (s *TestRequestService) Get(reqId int) (interf model.TestRequest, err error) {
	interf, err = s.TestRequestRepo.Get(uint(reqId))

	return
}

func (s *TestRequestService) Create(req serverDomain.TestRequest,
	resp serverDomain.TestResponse, projectId int) (request model.TestRequest, err error) {
	request = model.TestRequest{
		Name:        req.Name,
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

	return
}

func (s *TestRequestService) CopyValueFromRequest(interf *model.TestRequest, req serverDomain.TestRequest) (err error) {
	interf.ID = req.Id

	copier.Copy(interf, req)

	return
}
