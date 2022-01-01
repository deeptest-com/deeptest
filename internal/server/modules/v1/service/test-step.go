package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type TestStepService struct {
	TestStepRepo *repo.TestStepRepo `inject:""`
}

func NewTestStepService() *TestStepService {
	return &TestStepService{}
}

func (s *TestStepService) FindById(id uint) (serverDomain.TestStepResp, error) {
	return s.TestStepRepo.FindById(id)
}

func (s *TestStepService) Create(req serverDomain.TestStepReq) (uint, error) {
	return s.TestStepRepo.Create(req)
}

func (s *TestStepService) Update(id uint, req serverDomain.TestStepReq) error {
	return s.TestStepRepo.Update(id, req)
}

func (s *TestStepService) DeleteById(id uint) error {
	return s.TestStepRepo.DeleteById(id)
}
