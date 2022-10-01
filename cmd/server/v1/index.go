package v1

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/router"
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/rate"
)

type IndexModule struct {
	DataModule *router.DataModule `inject:""`
	FileModule *router.FileModule `inject:""`

	AccountModule *router.AccountModule `inject:""`
	UserModule    *router.UserModule    `inject:""`
	RoleModule    *router.RoleModule    `inject:""`
	PermModule    *router.PermModule    `inject:""`

	MockModule *router.MockModule `inject:""`

	ProjectModule     *router.ProjectModule     `inject:""`
	InterfaceModule   *router.InterfaceModule   `inject:""`
	InvocationModule  *router.InvocationModule  `inject:""`
	AuthModule        *router.AuthModule        `inject:""`
	EnvironmentModule *router.EnvironmentModule `inject:""`
	ExtractorModule   *router.ExtractorModule   `inject:""`
	CheckpointModule  *router.CheckpointModule  `inject:""`

	ScenarioModule          *router.ScenarioModule          `inject:""`
	ScenarioNodeModule      *router.ScenarioNodeModule      `inject:""`
	ScenarioProcessorModule *router.ScenarioProcessorModule `inject:""`
	ScenarioExecModule      *router.ScenarioExecModule      `inject:""`
	ReportModule            *router.ReportModule            `inject:""`
}

func NewIndexModule() *IndexModule {
	return &IndexModule{}
}

// Party v1 模块
func (m *IndexModule) Party() module.WebModule {
	handler := func(v1 iris.Party) {
		if !serverConfig.CONFIG.Limit.Disable {
			limitV1 := rate.Limit(
				serverConfig.CONFIG.Limit.Limit,
				serverConfig.CONFIG.Limit.Burst,
				rate.PurgeEvery(time.Minute, 5*time.Minute))
			v1.Use(limitV1)
		}
	}
	modules := []module.WebModule{
		m.DataModule.Party(),
		m.FileModule.Party(),
		m.AccountModule.Party(),
		m.RoleModule.Party(),
		m.PermModule.Party(),
		m.UserModule.Party(),

		m.MockModule.Party(),

		m.ProjectModule.Party(),
		m.InterfaceModule.Party(),
		m.InvocationModule.Party(),
		m.AuthModule.Party(),
		m.EnvironmentModule.Party(),
		m.ExtractorModule.Party(),
		m.CheckpointModule.Party(),
		m.ScenarioModule.Party(),
		m.ScenarioNodeModule.Party(),
		m.ScenarioProcessorModule.Party(),
		m.ScenarioExecModule.Party(),
		m.ReportModule.Party(),
	}
	return module.NewModule(serverConsts.ApiPath, handler, modules...)
}
