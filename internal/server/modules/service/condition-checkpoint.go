package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func (s *CheckpointService) Get(tenantId consts.TenantId, id uint) (checkpoint model.DebugConditionCheckpoint, err error) {
	checkpoint, err = s.CheckpointRepo.Get(tenantId, id)

	return
}

func (s *CheckpointService) Create(tenantId consts.TenantId, checkpoint *model.DebugConditionCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(tenantId, checkpoint)

	return
}

func (s *CheckpointService) Update(tenantId consts.TenantId, checkpoint *model.DebugConditionCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(tenantId, checkpoint)

	return
}
