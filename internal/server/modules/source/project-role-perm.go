package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ProjectRolePermSource struct {
	ProjectRolePermRepo *repo2.ProjectRolePermRepo `inject:""`
}

func (s *ProjectRolePermSource) GetSources(tenantId consts.TenantId) (res map[consts.RoleType][]uint, err error) {
	return s.ProjectRolePermRepo.GetProjectPermsForRole(tenantId)
}

func (s *ProjectRolePermSource) Init(tenantId consts.TenantId) (err error) {
	sources, err := s.GetSources(tenantId)
	if err != nil {
		return
	}

	var successCount int
	var failItems []string
	for roleName, source := range sources {
		successCount, failItems = s.ProjectRolePermRepo.AddPermForProjectRole(tenantId, roleName, source)
		color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据,角色名：%+v,失败数据：%+v!\n", model.ProjectRolePerm{}.TableName(), successCount, roleName, failItems)
	}

	return
}
