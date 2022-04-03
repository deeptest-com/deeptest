package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
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

func (s *TestRequestService) Create(req serverDomain.TestRequestReq) (interf *model.TestRequest, err error) {

	return
}

func (s *TestRequestService) Update(req model.TestRequest) (err error) {

	return
}

func (s *TestRequestService) Delete(reqId uint) (err error) {

	return
}
