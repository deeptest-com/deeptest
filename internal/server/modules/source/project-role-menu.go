package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ProjectRoleMenuSource struct {
	ProjectRoleMenuRepo *repo.ProjectRoleMenuRepo `inject:""`
}

func (s *ProjectRoleMenuSource) GetSources(tenantId consts.TenantId) (res []model.ProjectRoleMenu, err error) {
	return s.ProjectRoleMenuRepo.GetConfigData(tenantId)
}

func (s *ProjectRoleMenuSource) Init(tenantId consts.TenantId) (err error) {
	sources, err := s.GetSources(tenantId)
	if err != nil {
		return
	}
	s.ProjectRoleMenuRepo.DeleteAllData(tenantId)

	successCount, failItems := s.ProjectRoleMenuRepo.BatchCreate(tenantId, sources)
	color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据,失败数据：%+v!\n", model.ProjectRoleMenu{}.TableName(), successCount, failItems)

	return
}
