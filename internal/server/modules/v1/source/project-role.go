package source

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/gookit/color"
)

type ProjectRoleSource struct {
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
	PermRepo        *repo.PermRepo        `inject:""`
}

func NewProjectRoleSource() *ProjectRoleSource {
	return &ProjectRoleSource{}
}

func (s *ProjectRoleSource) GetSources() (sources []model.ProjectRole, err error) {
	sources = []model.ProjectRole{
		{
			Name: "项目管理员",
		},
	}
	return
}

func (s *ProjectRoleSource) Init() (err error) {
	sources, err := s.GetSources()
	if err != nil {
		return
	}

	for _, source := range sources {
		err = s.ProjectRoleRepo.Create(source)
		if err != nil { // 遇到错误时回滚事务
			return
		}
	}

	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.ProjectRole{}.TableName())
	return
}
