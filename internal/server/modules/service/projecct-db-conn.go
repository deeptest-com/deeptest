package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseConnService struct {
	DatabaseConnRepo *repo.DatabaseConnRepo `inject:""`
}

func (s *DatabaseConnService) List(keywords string, projectId int) (ret []model.DatabaseConn, err error) {
	ret, err = s.DatabaseConnRepo.List(keywords, projectId, false)
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

func (s *DatabaseConnService) UpdateName(req v1.DbConnReq) (err error) {
	err = s.DatabaseConnRepo.UpdateName(req)

	return
}
