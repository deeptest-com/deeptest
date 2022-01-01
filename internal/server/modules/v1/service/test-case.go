package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type TestCaseService struct {
	TestCaseRepo *repo.TestCaseRepo `inject:""`
}

func NewTestCaseService() *TestCaseService {
	return &TestCaseService{}
}

func (s *TestCaseService) Paginate(req serverDomain.TestCaseReqPaginate) (ret domain.PageData, err error) {
	ret, err = s.TestCaseRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *TestCaseService) FindById(id uint) (serverDomain.TestCaseResp, error) {
	return s.TestCaseRepo.FindById(id)
}

func (s *TestCaseService) Create(req serverDomain.TestCaseReq) (uint, error) {
	return s.TestCaseRepo.Create(req)
}

func (s *TestCaseService) Update(id uint, req serverDomain.TestCaseReq) error {
	return s.TestCaseRepo.Update(id, req)
}

func (s *TestCaseService) DeleteById(id uint) error {
	return s.TestCaseRepo.BatchDelete(id)
}
