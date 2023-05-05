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

func NewProjectRolePermSource() *ProjectRolePermSource {
	return &ProjectRolePermSource{}
}

func (s *ProjectRolePermSource) GetSources() (res map[consts.RoleType][]uint, err error) {
	return s.ProjectRolePermRepo.GetProjectPermsForRole()
}

func (s *ProjectRolePermSource) Init() (err error) {
	sources, err := s.GetSources()
	if err != nil {
		return
	}

	var successCount int
	var failItems []string
	for roleName, source := range sources {
		successCount, failItems = s.ProjectRolePermRepo.AddPermForProjectRole(roleName, source)
		color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据,角色名：%+v,失败数据：%+v!\n", model.ProjectRolePerm{}.TableName(), successCount, roleName, failItems)
	}

	return
}
