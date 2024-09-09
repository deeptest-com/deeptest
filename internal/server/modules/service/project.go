package service

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	integrationService "github.com/aaronchen2k/deeptest/integration/service"
	thirdparty "github.com/aaronchen2k/deeptest/integration/thirdparty/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type ProjectService struct {
	ProjectRepo               *repo.ProjectRepo                  `inject:""`
	ServeRepo                 *repo.ServeRepo                    `inject:""`
	UserRepo                  *repo.UserRepo                     `inject:""`
	ProjectRoleRepo           *repo.ProjectRoleRepo              `inject:""`
	MessageRepo               *repo.MessageRepo                  `inject:""`
	BaseRepo                  *repo.BaseRepo                     `inject:""`
	IntegrationRepo           *repo.IntegrationRepo              `inject:""`
	RemoteService             *thirdparty.RemoteService          `inject:""`
	MessageService            *MessageService                    `inject:""`
	UserService               *UserService                       `inject:""`
	IntegrationProjectService *integrationService.ProjectService `inject:""`
}

func (s *ProjectService) Paginate(tenantId consts.TenantId, req v1.ProjectReqPaginate, userId uint) (ret _domain.PageData, err error) {
	ret, err = s.ProjectRepo.Paginate(tenantId, req, userId)

	if err != nil {
		return
	}

	return
}

func (s *ProjectService) Get(tenantId consts.TenantId, id uint) (model.Project, error) {
	return s.ProjectRepo.Get(tenantId, id)
}

func (s *ProjectService) Create(tenantId consts.TenantId, req v1.ProjectReq, userId uint) (id uint, err _domain.BizErr) {
	id, err = s.ProjectRepo.Create(tenantId, req, userId)
	if err.Code != 0 {
		return
	}

	integrationErr := s.IntegrationProjectService.Save(tenantId, req.ProjectReq, id)
	if integrationErr != nil {
		err = _domain.SystemErr
	}

	return
}

func (s *ProjectService) Update(tenantId consts.TenantId, req v1.ProjectReq) (err error) {
	err = s.ProjectRepo.Update(tenantId, req)
	if err != nil {
		return
	}

	err = s.IntegrationProjectService.Save(tenantId, req.ProjectReq, req.Id)

	return
}

func (s *ProjectService) DeleteById(tenantId consts.TenantId, id uint) error {
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
	return s.ProjectRepo.DeleteById(tenantId, id)
}

func (s *ProjectService) GetByUser(tenantId consts.TenantId, userId uint) (projects []model.ProjectMemberRole, currProject model.Project, recentProjects []model.Project, err error) {
	projects, err = s.ProjectRepo.ListProjectByUser(tenantId, userId)
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(tenantId, userId)
	recentProjects, err = s.ProjectRepo.ListProjectsRecentlyVisited(tenantId, userId)

	return
}

func (s *ProjectService) ChangeProject(tenantId consts.TenantId, projectId, userId uint) (err error) {
	err = s.ProjectRepo.ChangeProject(tenantId, projectId, userId)

	return
}

func (s *ProjectService) Members(tenantId consts.TenantId, req v1.ProjectReqPaginate, projectId int) (data _domain.PageData, err error) {
	data, err = s.ProjectRepo.Members(tenantId, req, projectId)

	return
}

func (s *ProjectService) RemoveMember(tenantId consts.TenantId, req v1.ProjectMemberRemoveReq) (err error) {
	err = s.ProjectRepo.RemoveMember(tenantId, req.UserId, req.ProjectId)

	return
}

func (s *ProjectService) UpdateMemberRole(tenantId consts.TenantId, req v1.UpdateProjectMemberReq) (err error) {
	return s.ProjectRepo.UpdateUserRole(tenantId, req)
}

func (s *ProjectService) GetCurrProjectByUser(tenantId consts.TenantId, userId uint) (currProject model.Project, err error) {
	currProject, err = s.ProjectRepo.GetCurrProjectByUser(tenantId, userId)

	return
}

func (s *ProjectService) Apply(tenantId consts.TenantId, req v1.ApplyProjectReq) (err error) {
	//如果已经有审批记录，就不创建新的了
	var b bool
	b, err = s.ProjectRepo.IfProjectMember(tenantId, req.ApplyUserId, req.ProjectId)
	if err != nil || b {
		return
	}
	result, _ := s.ProjectRepo.GetAuditByItem(tenantId, req.ProjectId, req.ApplyUserId, []consts.AuditStatus{consts.Init})
	if result.ID != 0 {
		return
		//return fmt.Errorf("您已提交了申请，请联系审批人审批")
	}
	auditId, err := s.ProjectRepo.SaveAudit(tenantId, model.ProjectMemberAudit{ProjectId: req.ProjectId, ApplyUserId: req.ApplyUserId, ProjectRoleName: req.ProjectRoleName, Description: req.Description})
	if err != nil {
		return
	}

	go func() {
		err = s.IntegrationProjectService.SendApplyMessage(tenantId, req.ProjectId, req.ApplyUserId, auditId, req.ProjectRoleName)
		if err != nil {
			logUtils.Infof("申请加入项目发送消息失败，err:%+v", err)
		}
	}()

	return
}

func (s *ProjectService) Audit(tenantId consts.TenantId, id, auditUserId uint, status consts.AuditStatus) (err error) {

	var record model.ProjectMemberAudit
	record, err = s.ProjectRepo.GetAudit(tenantId, id)
	if err != nil {
		return err
	}

	//防止重复审批
	if record.Status != consts.Init {
		return
	}

	err = s.ProjectRepo.UpdateAuditStatus(tenantId, id, auditUserId, status)
	if err != nil {
		return err
	}

	if status == consts.Refused {
		return
	}

	var res bool
	res, err = s.ProjectRepo.IfProjectMember(tenantId, record.ApplyUserId, record.ProjectId)
	if err != nil {
		return
	}

	if res {
		return
	}

	err = s.ProjectRepo.AddProjectMember(tenantId, record.ProjectId, record.ApplyUserId, record.ProjectRoleName)
	if err != nil {
		return
	}
	return
}

func (s *ProjectService) AuditList(tenantId consts.TenantId, req v1.AuditProjectPaginate) (data _domain.PageData, err error) {
	return s.ProjectRepo.GetAuditList(tenantId, req)
}

func (s *ProjectService) AuditUsers(tenantId consts.TenantId, projectId uint) (data []model.SysUser, err error) {
	return s.ProjectRepo.GetAuditUsers(tenantId, projectId)
}

func (s *ProjectService) CheckProjectAndUser(tenantId consts.TenantId, shortName string, userId uint) (project model.Project, userInProject bool, err error) {
	project, err = s.ProjectRepo.GetByShortName(tenantId, shortName)
	if err != nil {
		return
	}

	//if err != nil {
	//	if err != gorm.ErrRecordNotFound || xToken == "" {
	//		return project, userInProject, err
	//	}
	//
	//	thirdPartyProject, err := s.RemoteService.GetProjectInfo(xToken, shortName)
	//	if err != nil {
	//		return project, userInProject, err
	//	}
	//
	//	_, err = s.CreateProjectForThirdParty(thirdPartyProject)
	//	if err != nil {
	//		return project, userInProject, err
	//	}
	//
	//	project, err = s.ProjectRepo.GetByShortName(shortName)
	//	if err != nil {
	//		return project, userInProject, err
	//	}
	//}
	//
	//if xToken != "" && project.Source != serverConsts.ProjectSourceLY {
	//	err = s.ProjectRepo.UpdateProjectSource(project.ID, serverConsts.ProjectSourceLY)
	//	if err != nil {
	//		return project, userInProject, err
	//	}
	//}

	isAdminUser, err := s.UserRepo.IsAdminUser(tenantId, userId)
	if err != nil {
		return
	}
	if isAdminUser {
		return project, true, nil
	}

	userInProject, err = s.ProjectRepo.IfProjectMember(tenantId, userId, project.ID)
	if err != nil {
		return
	}

	//if !userInProject && xToken != "" {
	//	err = s.ProjectRepo.AddProjectMember(project.ID, userId, consts.User)
	//	if err != nil {
	//		return
	//	}
	//
	//	userInProject = true
	//}

	return
}

func (s *ProjectService) CreateProjectForThirdParty(tenantId consts.TenantId, project integrationDomain.ProjectInfo) (projectId uint, err error) {
	adminName := "admin"
	adminUser, err := s.UserRepo.GetByUserName(tenantId, adminName)
	if err != nil {
		return
	}

	//建项目
	createReq := v1.ProjectReq{
		ProjectBase: v1.ProjectBase{
			Name:      project.Name,
			ShortName: project.NameEngAbbr,
			AdminId:   adminUser.ID,
			AdminName: adminName,
			Source:    serverConsts.ProjectSourceLY,
		},
	}
	projectId, createErr := s.Create(tenantId, createReq, adminUser.ID)
	if projectId == 0 {
		err = errors.New(createErr.Error())
		return
	}

	//创建项目管理员
	for _, spaceAdmin := range project.SpaceAdmins {
		spaceAdminUser, err := s.UserRepo.GetByUserName(tenantId, spaceAdmin.Username)
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		var spaceAdminId uint
		if spaceAdminUser.ID != 0 {
			spaceAdminId = spaceAdminUser.ID
		} else {
			createUserReq := v1.UserReq{
				UserBase: v1.UserBase{
					Username:  spaceAdmin.Username,
					Name:      spaceAdmin.RealName,
					Email:     spaceAdmin.Mail,
					ImAccount: spaceAdmin.WxName,
					Password:  commonUtils.RandStr(8),
				},
			}
			spaceAdminId, err = s.UserService.Create(tenantId, createUserReq)
			if err != nil {
				continue
			}
		}

		err = s.ProjectRepo.AddProjectMember(tenantId, projectId, spaceAdminId, s.BaseRepo.GetAdminRoleName())
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

func (s *ProjectService) AllProjectList(tenantId consts.TenantId, username string) (res []model.Project, err error) {
	return s.ProjectRepo.ListByUsername(tenantId, username)
}

func (s *ProjectService) GetProjectRole(tenantId consts.TenantId, username, projectCode string) (role string, err error) {
	var user model.SysUser
	user, _ = s.UserRepo.GetByUserName(tenantId, username)
	if user.ID == 0 {
		err = fmt.Errorf("用户名不存在")
		return
	}
	var project model.Project
	project, _ = s.ProjectRepo.GetByShortName(tenantId, projectCode)
	if project.ID == 0 {
		err = fmt.Errorf("项目不存在")
		return
	}

	var projectRole model.ProjectRole
	projectRole, _ = s.ProjectRoleRepo.ProjectUserRoleList(tenantId, user.ID, project.ID)
	if projectRole.Name == "" {
		err = fmt.Errorf("用户角色不存在")
		return
	}

	return string(projectRole.Name), err
}
