package v1

import (
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/index"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/rate"
)

type IndexModule struct {
	DataModule *index.DataModule `inject:""`
	FileModule *index.FileModule `inject:""`

	AccountModule *index.AccountModule `inject:""`
	UserModule    *index.UserModule    `inject:""`
	RoleModule    *index.RoleModule    `inject:""`
	PermModule    *index.PermModule    `inject:""`

	MockModule *index.MockModule `inject:""`

	ProjectModule     *index.ProjectModule     `inject:""`
	InterfaceModule   *index.InterfaceModule   `inject:""`
	InvocationModule  *index.InvocationModule  `inject:""`
	EnvironmentModule *index.EnvironmentModule `inject:""`
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
		m.EnvironmentModule.Party(),
	}
	return module.NewModule(serverConsts.ApiPath, handler, modules...)
}
