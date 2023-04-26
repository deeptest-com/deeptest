package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/source"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectService struct {
	ProjectRepo  *repo.ProjectRepo    `inject:""`
	ServeRepo    *repo.ServeRepo      `inject:""`
	SampleSource *source.SampleSource `inject:""`
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

func (s *ProjectService) Get(id uint) (model.Project, error) {
	return s.ProjectRepo.Get(id)
}

func (s *ProjectService) Create(req v1.ProjectReq, userId uint) (id uint, err *_domain.BizErr) {
	id, err = s.ProjectRepo.Create(req, userId)
	return
}

func (s *ProjectService) Update(id uint, req v1.ProjectReq) error {
	return s.ProjectRepo.Update(id, req)
}

func (s *ProjectService) DeleteById(id uint) error {
	count, err := s.ServeRepo.GetCountByProject(id)
	if err != nil {
		return err
	}

	if count > 0 {
		err = fmt.Errorf("services under the project, cannot be deleted")
		return err
	}
	return s.ProjectRepo.DeleteById(id)
}

func (s *ProjectService) GetByUser(userId uint) (projects []model.Project, currProject model.Project, recentProjects []model.Project, err error) {
	projects, err = s.ProjectRepo.ListProjectByUser(userId)
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(userId)
	recentProjects, err = s.ProjectRepo.ListProjectsRecentlyVisited(userId)

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

func (s *ProjectService) UpdateMemberRole(req v1.UpdateProjectMemberReq) (err error) {
	return s.ProjectRepo.UpdateUserRole(req)
}

func (s *ProjectService) GetCurrProjectByUser(userId uint) (currProject model.Project, err error) {
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(userId)

	return
}

/*
func (s *ProjectService) createSample(projectId uint) (err error) {
	serve, endpoint, _ := s.SampleSource.GetSources()

	serve.ProjectId = projectId
	endpoint.ProjectId = projectId
	//err = s.ProjectRepo.CreateSample(serve, endpoint)

	return err
}
*/
