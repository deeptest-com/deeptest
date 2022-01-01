package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type TestScriptService struct {
	TestScriptRepo *repo.TestScriptRepo `inject:""`
}

func NewTestScriptService() *TestScriptService {
	return &TestScriptService{}
}

func (s *TestScriptService) Paginate(req serverDomain.TestScriptReqPaginate) (ret domain.PageData, err error) {

	ret, err = s.TestScriptRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *TestScriptService) GetDetail(id uint) (serverDomain.TestScriptResp, error) {
	return s.TestScriptRepo.GetDetail(id)
}

func (s *TestScriptService) Create(req serverDomain.TestScriptReq) (uint, error) {
	return s.TestScriptRepo.Create(req)
}

func (s *TestScriptService) Update(id uint, req serverDomain.TestScriptReq) error {
	return s.TestScriptRepo.Update(id, req)
}

func (s *TestScriptService) DeleteById(id uint) error {
	return s.TestScriptRepo.DeleteById(id)
}
