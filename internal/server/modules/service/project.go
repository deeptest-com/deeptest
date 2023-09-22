package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/source"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectService struct {
	ProjectRepo  *repo.ProjectRepo    `inject:""`
	ServeRepo    *repo.ServeRepo      `inject:""`
	SampleSource *source.SampleSource `inject:""`
	UserRepo     *repo.UserRepo       `inject:""`
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

func (s *ProjectService) Create(req v1.ProjectReq, userId uint) (id uint, err _domain.BizErr) {
	id, err = s.ProjectRepo.Create(req, userId)
	return
}

func (s *ProjectService) Update(req v1.ProjectReq) error {
	return s.ProjectRepo.Update(req)
}

func (s *ProjectService) DeleteById(id uint) error {
	/*
		count, err := s.ServeRepo.GetCountByProject(id)
		if err != nil {
			return err
		}

		if count > 0 {
			err = fmt.Errorf("services under the project, cannot be deleted")
			return err
		}
	*/
	return s.ProjectRepo.DeleteById(id)
}

func (s *ProjectService) GetByUser(userId uint) (projects []model.ProjectMemberRole, currProject model.Project, recentProjects []model.Project, err error) {
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

func (s *ProjectService) Apply(req v1.ApplyProjectReq) (err error) {
	err = s.ProjectRepo.SaveAudit(model.ProjectMemberAudit{ProjectId: req.ProjectId, ApplyUserId: req.ApplyUserId, ProjectRoleName: req.ProjectRoleName, Description: req.Description})
	return
}

func (s *ProjectService) Audit(id, auditUserId uint, status consts.AuditStatus) (err error) {

	var record model.ProjectMemberAudit
	record, err = s.ProjectRepo.GetAudit(id)
	if err != nil {
		return err
	}

	err = s.ProjectRepo.UpdateAuditStatus(id, auditUserId, status)
	if err != nil {
		return err
	}

	if status == consts.Refused {
		return
	}

	var res bool
	res, err = s.ProjectRepo.IfProjectMember(record.ApplyUserId, record.ProjectId)
	if err != nil {
		return
	}

	if res {
		return
	}

	err = s.ProjectRepo.AddProjectMember(record.ProjectId, record.ApplyUserId, record.ProjectRoleName)
	if err != nil {
		return
	}
	return
}

func (s *ProjectService) AuditList(req v1.AuditProjectPaginate) (data _domain.PageData, err error) {
	return s.ProjectRepo.GetAuditList(req)
}

func (s *ProjectService) AuditUsers(projectId uint) (data []model.SysUser, err error) {
	return s.ProjectRepo.GetAuditUsers(projectId)
}

func (s *ProjectService) CheckProjectAndUser(shortName string, userId uint) (project model.Project, userInProject bool, err error) {
	project, err = s.ProjectRepo.GetByShortName(shortName)
	if err != nil {
		return
	}

	isAdminUser, err := s.UserRepo.IsAdminUser(userId)
	if err != nil {
		return
	}
	if isAdminUser {
		return project, true, nil
	}

	userInProject, err = s.ProjectRepo.IfProjectMember(userId, project.ID)
	if err != nil {
		return
	}

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
