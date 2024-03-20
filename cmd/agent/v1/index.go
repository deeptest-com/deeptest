package v1

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/router"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/rate"
)

type IndexModule struct {
	SpecModule        *router.SpecModule        `inject:""`
	InvocationModule  *router.ExecModule        `inject:""`
	PerformanceModule *router.PerformanceModule `inject:""`
	HealthzModule     *router.HealthzModule     `inject:""`
}

func NewIndexModule() *IndexModule {
	return &IndexModule{}
}

// Party v1 模块
func (m *IndexModule) Party() module.WebModule {
	handler := func(v1 iris.Party) {
		if !config.CONFIG.Limit.Disable {
			limitV1 := rate.Limit(
				config.CONFIG.Limit.Limit,
				config.CONFIG.Limit.Burst,
				rate.PurgeEvery(time.Minute, 5*time.Minute))
			v1.Use(limitV1)
		}
	}
	modules := []module.WebModule{
		m.SpecModule.Party(),
		m.InvocationModule.Party(),
		m.PerformanceModule.Party(),
		m.HealthzModule.Party(),
	}
	return module.NewModule(consts.ApiPathAgent, handler, modules...)
}
