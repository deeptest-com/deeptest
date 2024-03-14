package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointFavoriteModule struct {
	EndpointFavoriteCtrl *handler.EndpointFavoriteCtrl `inject:""`
}

func (m *EndpointFavoriteModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		public.Post("/favorite", m.EndpointFavoriteCtrl.Favorite).Name = "收藏"
		public.Put("/favorite/list", m.EndpointFavoriteCtrl.Index).Name = "更新设计器状态"
	}
	return module.NewModule("/endpoints", handler)
}
