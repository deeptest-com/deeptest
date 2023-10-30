package service

import (
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseOptService struct {
	PostConditionRepo *repo.PostConditionRepo `inject:""`
	DatabaseOptRepo   *repo.DatabaseOptRepo   `inject:""`
	EnvironmentRepo   *repo.EnvironmentRepo   `inject:""`
}

func (s *DatabaseOptService) Get(id uint) (checkpoint model.DebugConditionDatabaseOpt, err error) {
	checkpoint, err = s.DatabaseOptRepo.Get(id)

	return
}

func (s *DatabaseOptService) Create(checkpoint *model.DebugConditionDatabaseOpt) (err error) {
	err = s.DatabaseOptRepo.Save(checkpoint)

	return
}

func (s *DatabaseOptService) Update(checkpoint *model.DebugConditionDatabaseOpt) (err error) {
	err = s.DatabaseOptRepo.Save(checkpoint)

	return
}

func (s *DatabaseOptService) Delete(reqId uint) (err error) {
	err = s.DatabaseOptRepo.Delete(reqId)

	return
}
