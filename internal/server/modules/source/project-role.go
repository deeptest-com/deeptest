package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ProjectRoleSource struct {
	ProjectRoleRepo *repo2.ProjectRoleRepo `inject:""`
	PermRepo        *repo2.PermRepo        `inject:""`
}

func (s *ProjectRoleSource) GetSources() (sources []model.ProjectRole, err error) {
	sources = []model.ProjectRole{
		{
			Name:        consts.Admin,
			DisplayName: "项目空间管理员",
		},
		{
			Name:        consts.User,
			DisplayName: "普通用户",
		},
		{
			Name:        consts.Tester,
			DisplayName: "测试工程师",
		},
		{
			Name:        consts.Developer,
			DisplayName: "开发工程师",
		},
		{
			Name:        consts.ProductManager,
			DisplayName: "产品经理",
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
			color.Info.Printf("\n[Mysql] --> %s 表初始数据失败!,err:%s", model.ProjectRole{}.TableName(), err.Error())
			return
		}
	}

	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.ProjectRole{}.TableName())
	return
}
