package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ProjectService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) Paginate(req serverDomain.ProjectReqPaginate) (ret domain.PageData, err error) {

	ret, err = s.ProjectRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *ProjectService) FindById(id uint) (serverDomain.ProjectResp, error) {
	return s.ProjectRepo.FindById(id)
}

func (s *ProjectService) Create(req serverDomain.ProjectReq) (uint, error) {
	return s.ProjectRepo.Create(req)
}

func (s *ProjectService) Update(id uint, req serverDomain.ProjectReq) error {
	return s.ProjectRepo.Update(id, req)
}

func (s *ProjectService) DeleteById(id uint) error {
	return s.ProjectRepo.BatchDelete(id)
}

func (s *ProjectService) GetByUser(userId uint) (projects []model.Project, currProject model.Project, err error) {
	projects, err = s.ProjectRepo.ListProjectByUser(userId)
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(userId)

	return
}
