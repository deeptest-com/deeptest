package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) Paginate(req serverDomain.ProjectReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ProjectRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *ProjectService) GetById(id uint) (model.Project, error) {
	return s.ProjectRepo.FindById(id)
}

func (s *ProjectService) Create(req serverDomain.ProjectReq, userId uint) (uint, *_domain.BizErr) {
	return s.ProjectRepo.Create(req, userId)
}

func (s *ProjectService) Update(id uint, req serverDomain.ProjectReq) error {
	return s.ProjectRepo.Update(id, req)
}

func (s *ProjectService) DeleteById(id uint) error {
	return s.ProjectRepo.DeleteById(id)
}

func (s *ProjectService) GetByUser(userId uint) (projects []model.Project, currProject model.Project, err error) {
	projects, err = s.ProjectRepo.ListProjectByUser(userId)
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(userId)

	return
}

func (s *ProjectService) ChangeProject(projectId, userId uint) (err error) {
	err = s.ProjectRepo.ChangeProject(projectId, userId)

	return
}
