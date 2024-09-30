package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type HealthzModule struct {
	HealthzCtrl *handler.HealthzCtrl `inject:""`
}

func NewHealthzModule() *HealthzModule {
	return &HealthzModule{}
}

// Party
func (m *HealthzModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck())
		index.Get("/", m.HealthzCtrl.Get).Name = "健康检查"
	}
	return module.NewModule("/healthz", handler)
}
