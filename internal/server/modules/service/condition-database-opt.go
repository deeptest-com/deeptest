package service

import (
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseOptService struct {
	ConditionRepo   *repo.ConditionRepo   `inject:""`
	DatabaseOptRepo *repo.DatabaseOptRepo `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *DatabaseOptService) Get(id uint) (opt model.DebugConditionDatabaseOpt, err error) {
	opt, err = s.DatabaseOptRepo.Get(id)

	return
}

func (s *DatabaseOptService) Create(opt *model.DebugConditionDatabaseOpt) (err error) {
	err = s.DatabaseOptRepo.Save(opt)

	return
}

func (s *DatabaseOptService) Update(opt *model.DebugConditionDatabaseOpt) (err error) {
	err = s.DatabaseOptRepo.Save(opt)

	return
}

func (s *DatabaseOptService) Delete(reqId uint) (err error) {
	err = s.DatabaseOptRepo.Delete(reqId)

	return
}
