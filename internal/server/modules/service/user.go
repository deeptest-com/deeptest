package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type UserService struct {
	UserRepo *repo.UserRepo `inject:""`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Paginate(req v1.UserReqPaginate) (_domain.PageData, error) {
	return s.UserRepo.Paginate(req)
}

// getRoles
func (s *UserService) getRoles(users ...*v1.UserResp) {
	s.UserRepo.GetRoles(users...)
}

func (s *UserService) FindByUserName(username string, ids ...uint) (v1.UserResp, error) {
	return s.UserRepo.FindByUserName(username, ids...)
}

func (s *UserService) FindPasswordByUserName(username string, ids ...uint) (v1.LoginResp, error) {
	return s.UserRepo.FindPasswordByUserName(username, ids...)
}

func (s *UserService) Create(req v1.UserReq) (uint, error) {
	return s.UserRepo.Create(req)
}

func (s *UserService) Update(id uint, req v1.UserReq) error {
	return s.UserRepo.Update(id, req)
}

func (s *UserService) IsAdminUser(id uint) (bool, error) {
	return s.UserRepo.IsAdminUser(id)
}

func (s *UserService) FindById(id uint) (v1.UserResp, error) {
	return s.UserRepo.FindById(id)
}

func (s *UserService) DeleteById(id uint) error {
	return s.UserRepo.DeleteById(id)
}

// AddRoleForUser add roles for user
func (s *UserService) AddRoleForUser(user *model.SysUser) error {
	return s.UserRepo.AddRoleForUser(user)
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
