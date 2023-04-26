package source

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ProjectRoleMenuSource struct {
	ProjectRoleMenuRepo *repo2.ProjectRoleMenuRepo `inject:""`
}

func NewProjectRoleMenuSource() *ProjectRoleMenuSource {
	return &ProjectRoleMenuSource{}
}

func (s *ProjectRoleMenuSource) GetSources() (res []model.ProjectRoleMenu, err error) {
	return s.ProjectRoleMenuRepo.GetConfigData()
}

func (s *ProjectRoleMenuSource) Init() (err error) {
	sources, err := s.GetSources()
	if err != nil {
		return
	}

	successCount, failItems := s.ProjectRoleMenuRepo.BatchCreate(sources)
	color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据,失败数据：%+v!\n", model.ProjectRoleMenu{}.TableName(), successCount, failItems)

	return
}
