package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type OpenModule struct {
	OpenCtrl *handler.OpenCtrl `inject:""`
}

func (m *OpenModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.OpenCheck())
		index.Get("/allProjectList", m.OpenCtrl.AllProjectList).Name = "所有项目列表"
		index.Get("/getProjectsBySpace", m.OpenCtrl.GetProjectsBySpace).Name = "获取空间关联的项目列表"
		index.Post("/saveSpaceRelatedProjects", m.OpenCtrl.SaveSpaceRelatedProjects).Name = "保存空间和项目的关系"

	}

	return module.NewModule("/openApi", handler)
}
