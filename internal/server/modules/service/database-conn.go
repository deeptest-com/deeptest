package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseConnService struct {
	DatabaseConnRepo *repo.DatabaseConnRepo `inject:""`
}

func (s *DatabaseConnService) List(envId uint) (ret []model.DatabaseConn, err error) {
	ret, err = s.DatabaseConnRepo.List(envId)
	return
}

func (s *DatabaseConnService) Get(id uint) (model.DatabaseConn, error) {
	return s.DatabaseConnRepo.Get(id)
}

func (s *DatabaseConnService) Save(po *model.DatabaseConn) (err error) {
	return s.DatabaseConnRepo.Save(po)
}

func (s *DatabaseConnService) Delete(id uint) (err error) {
	return s.DatabaseConnRepo.Delete(id)
}

func (s *DatabaseConnService) Disable(id uint) (err error) {
	return s.DatabaseConnRepo.Disable(id)
}
