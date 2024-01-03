package service

import (
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type CheckpointService struct {
	PostConditionRepo *repo.ConditionRepo   `inject:""`
	CheckpointRepo    *repo.CheckpointRepo  `inject:""`
	EnvironmentRepo   *repo.EnvironmentRepo `inject:""`
	ProjectRepo       *repo.ProjectRepo     `inject:""`
	ExtractorRepo     *repo.ExtractorRepo   `inject:""`
	VariableService   *VariableService      `inject:""`
}

func (s *CheckpointService) Get(id uint) (checkpoint model.DebugConditionCheckpoint, err error) {
	checkpoint, err = s.CheckpointRepo.Get(id)

	return
}

func (s *CheckpointService) Create(checkpoint *model.DebugConditionCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}

func (s *CheckpointService) Update(checkpoint *model.DebugConditionCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}
