package service

import (
	"fmt"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type MessageService struct {
	UserRepo        *repo.UserRepo        `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
	MessageRepo     *repo.MessageRepo     `inject:""`
	BaseRepo        *repo.BaseRepo        `inject:""`
	RoleService     *RoleService          `inject:""`
	RemoteService   *RemoteService        `inject:""`
	Cron            *cron.ServerCron      `inject:""`
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

	host, _ := cache.GetCacheString("thirdPartyHost")

	projectHomePage := fmt.Sprintf("%s/lyapi/%s/workspace", host, project.ShortName)
	mcsData = integrationDomain.ApprovalReq{
		CreatorId:    sender.Username,
		ApproveIds:   userAccount,
		ApprovalType: 1,
		Title:        "乐研API通知-项目权限申请",
		Content:      fmt.Sprintf("您好！%s申请\"%s\"项目的【%s】角色。请审批！\n项目链接：%s", sender.Name, project.Name, applyRole, projectHomePage),
		SourceIds:    []int{0},
		Remark:       auditData.Description,
		NotifyUrl:    fmt.Sprintf("%s/api/v1/message/receiveMcsApprovalData", config.CONFIG.Environment.ServerHost),
	}

	return
}

func (s *MessageService) SendMessageToMcs(message model.Message) (mcsMessageId string, err error) {
	mcsMessageId, err = s.RemoteService.ApprovalAndMsg(message.Content)
	if err != nil {
		_ = s.MessageRepo.UpdateSendStatusById(message.ID, consts.MessageSendFailed)
		return
	}

	message.McsMessageId = mcsMessageId
	if mcsMessageId != "" {
		if message.ServiceType == consts.ServiceTypeInfo {
			err = s.MessageRepo.UpdateCombinedSendStatus(message.MessageSource, message.BusinessId, consts.MessageSendSuccess)
		} else {
			message.SendStatus = consts.MessageSendSuccess
			s.MessageRepo.DB.Save(&message)
		}
	}

	return
}

func (s *MessageService) SendMessageToMcsAsync() (err error) {
	messages, err := s.MessageRepo.ListMsgNeedAsyncToMcs()
	if err != nil {
		return
	}

	for _, message := range messages {
		_, err = s.SendMessageToMcs(message)
	}

	return
}

func (s *MessageService) SendMessageCron() {
	name := "SendMessageSync"

	s.Cron.RemoveTask(name)

	s.Cron.AddCommonTask(name, "*/5 * * * *", func() {
		err := s.SendMessageToMcsAsync()
		if err != nil {
			logUtils.Error("send message 定时导入任务失败，错误原因：" + err.Error())
		}

		logUtils.Info("send message 定时任务结束")
	})
}
