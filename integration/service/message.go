package service

import (
	"fmt"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type MessageService struct {
	UserRepo        *repo.UserRepo        `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
	BaseRepo        *repo.BaseRepo        `inject:""`
	RoleService     *RoleService          `inject:""`
}

func (s *MessageService) GetJoinProjectMcsData(senderId, projectId, auditId uint, roleValue consts.RoleType) (mcsData integrationDomain.ApprovalReq, err error) {
	sender, err := s.UserRepo.GetByUserId(senderId)
	if err != nil {
		return
	}

	project, err := s.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	adminRole, err := s.ProjectRoleRepo.FindByName(s.BaseRepo.GetAdminRoleName())
	if err != nil {
		return
	}

	applyRole, err := s.RoleService.GetRoleNameByValue(string(roleValue))
	if err != nil {
		return
	}

	userAccount, err := s.ProjectRepo.GetUsernamesByProjectAndRole(projectId, adminRole.ID, serverConsts.AdminUserName)
	if err != nil {
		return
	}

	auditData, err := s.ProjectRepo.GetAudit(auditId)
	if err != nil {
		return
	}

	host, _ := cache.GetCacheString("host")

	projectHomePage := fmt.Sprintf("%s/%s/workspace", host, project.ShortName)
	mcsData = integrationDomain.ApprovalReq{
		CreatorId:    sender.Username,
		ApproveIds:   userAccount,
		ApprovalType: 1,
		Title:        "乐研API通知-项目权限申请",
		Content:      fmt.Sprintf("您好！%s申请\"%s\"项目的【%s】角色。请审批！\n项目链接：%s \n查看更多：%s", sender.Name, project.Name, applyRole, projectHomePage, host+"/notification"),
		SourceIds:    []int{0},
		Remark:       auditData.Description,
		NotifyUrl:    fmt.Sprintf("%s/api/v1/message/receiveMcsApprovalData", config.CONFIG.Environment.ServerHost),
	}

	return
}
