package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	"github.com/deeptest-com/deeptest/pkg/domain"
	_i118Utils "github.com/deeptest-com/deeptest/pkg/lib/i118"
	_mailUtils "github.com/deeptest-com/deeptest/pkg/lib/mail"
	"strconv"
)

type UserService struct {
	UserRepo *repo.UserRepo `inject:""`
}

func (s *UserService) Paginate(tenantId consts.TenantId, req v1.UserReqPaginate) (_domain.PageData, error) {
	return s.UserRepo.Paginate(tenantId, req)
}

// getRoles
func (s *UserService) getRoles(tenantId consts.TenantId, users ...*v1.UserResp) {
	s.UserRepo.GetSysRoles(tenantId, users...)
}

func (s *UserService) FindByUserName(tenantId consts.TenantId, username string, ids ...uint) (v1.UserResp, error) {
	return s.UserRepo.FindByUserName(tenantId, username, ids...)
}

func (s *UserService) FindPasswordByUserName(tenantId consts.TenantId, username string, ids ...uint) (v1.LoginResp, error) {
	return s.UserRepo.FindPasswordByUserName(tenantId, username, ids...)
}

func (s *UserService) Create(tenantId consts.TenantId, req v1.UserReq) (uint, error) {
	return s.UserRepo.Create(tenantId, req)
}

func (s *UserService) Update(tenantId consts.TenantId, userId, id uint, req v1.UserReq) error {
	return s.UserRepo.Update(tenantId, userId, id, req)
}

func (s *UserService) IsAdminUser(tenantId consts.TenantId, id uint) (bool, error) {
	return s.UserRepo.IsAdminUser(tenantId, id)
}

func (s *UserService) FindById(tenantId consts.TenantId, id uint) (v1.UserResp, error) {
	return s.UserRepo.FindById(tenantId, id)
}

func (s *UserService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.UserRepo.DeleteById(tenantId, id)
}

// AddRoleForUser add roles for user
func (s *UserService) AddRoleForUser(tenantId consts.TenantId, user *model.SysUser) error {
	return s.UserRepo.AddRoleForUser(tenantId, user)
}

// DelToken 删除token
func (s *UserService) DelToken(token string) error {
	return s.UserRepo.DelToken(token)
}

// CleanToken 清空 token
func (s *UserService) CleanToken(authorityType int, userId string) error {
	return s.UserRepo.CleanToken(authorityType, userId)
}

func (s *UserService) UpdateAvatar(id uint, avatar string) error {
	return s.UserRepo.UpdateAvatar(id, avatar)
}

func (s *UserService) Invite(tenantId consts.TenantId, req v1.InviteUserReq) (user model.SysUser, bizErr *_domain.BizErr) {
	_, err := s.UserRepo.InviteToProject(tenantId, req)

	if err != nil {
		bizErr = &_domain.BizErr{Code: _domain.ErrNoUser.Code}
		return
	}

	vcode, _ := s.UserRepo.GenAndUpdateVcode(tenantId, user.ID)

	url := config.CONFIG.System.Website
	if !consts.IsRelease {
		url = consts.WebsiteDev
	}
	settings := map[string]string{
		"name":  user.Username,
		"url":   url,
		"vcode": vcode,
		"sys":   config.CONFIG.System.Name,
	}
	_mailUtils.Send(user.Email, _i118Utils.Sprintf("invite_user"), "invite-user", settings)

	return
}

func (s *UserService) GetUsersNotExistedInProject(tenantId consts.TenantId, projectId uint) (ret []v1.UserResp, err error) {
	ret, err = s.UserRepo.GetUsersNotExistedInProject(tenantId, projectId)

	return
}

func (s *UserService) UpdateSysRoleForUser(tenantId consts.TenantId, userId uint, roleIds []uint) (err error) {
	strRoleIds := make([]string, 0)
	for _, v := range roleIds {
		strRoleIds = append(strRoleIds, strconv.Itoa(int(v)))
	}
	err = s.UserRepo.UpdateRoleForUser(tenantId, strconv.Itoa(int(userId)), strRoleIds)

	return
}
