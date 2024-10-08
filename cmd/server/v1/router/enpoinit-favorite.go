package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointFavoriteModule struct {
	EndpointFavoriteCtrl *handler.EndpointFavoriteCtrl `inject:""`
}

func (m *EndpointFavoriteModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		public.Post("/favorite", m.EndpointFavoriteCtrl.Favorite).Name = "收藏"
		public.Post("/favorite/list", m.EndpointFavoriteCtrl.Index).Name = "收藏列表"
	}
	return module.NewModule("/endpoints", handler)
}
