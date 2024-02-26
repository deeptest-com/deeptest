package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseOptService struct {
	ConditionRepo   *repo.ConditionRepo   `inject:""`
	DatabaseOptRepo *repo.DatabaseOptRepo `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *DatabaseOptService) Get(tenantId consts.TenantId, id uint) (opt model.DebugConditionDatabaseOpt, err error) {
	opt, err = s.DatabaseOptRepo.Get(tenantId, id)

	return
}

func (s *DatabaseOptService) Update(tenantId consts.TenantId, opt *model.DebugConditionDatabaseOpt) (err error) {
	err = s.DatabaseOptRepo.Save(tenantId, opt)

	return
}
