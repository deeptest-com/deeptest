package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ScriptService struct {
	ScriptRepo      *repo.ScriptRepo      `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	VariableService *VariableService      `inject:""`
}

func (s *ScriptService) Get(tenantId consts.TenantId, id uint) (script model.DebugConditionScript, err error) {
	script, err = s.ScriptRepo.Get(tenantId, id)

	return
}

func (s *ScriptService) Create(tenantId consts.TenantId, script *model.DebugConditionScript) (err error) {
	err = s.ScriptRepo.Save(tenantId, script)

	return
}

func (s *ScriptService) Update(tenantId consts.TenantId, script *model.DebugConditionScript) (err error) {
	err = s.ScriptRepo.Save(tenantId, script)

	return
}
