package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) Paginate(req v1.ProjectReqPaginate, userId uint) (ret _domain.PageData, err error) {
	ret, err = s.ProjectRepo.Paginate(req, userId)

	if err != nil {
		return
	}

	return
}

func (s *ProjectService) GetById(id uint) (model.Project, error) {
	return s.ProjectRepo.FindById(id)
}

func (s *ProjectService) Create(req v1.ProjectReq, userId uint) (uint, *_domain.BizErr) {
	return s.ProjectRepo.Create(req, userId)
}

func (s *ProjectService) Update(id uint, req v1.ProjectReq) error {
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

func (s *ProjectService) Members(req v1.ProjectReqPaginate, projectId int) (data _domain.PageData, err error) {
	data, err = s.ProjectRepo.Members(req, projectId)

	return
}

func (s *ProjectService) RemoveMember(req v1.ProjectMemberRemoveReq) (err error) {
	err = s.ProjectRepo.RemoveMember(req.UserId, req.ProjectId)

	return
}
