package v1

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/router"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
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

	DebugModule         *router.DebugModule         `inject:""`
	TestInterfaceModule *router.TestInterfaceModule `inject:""`

	MockModule *router.MockModule `inject:""`

	ProjectModule     *router.ProjectModule     `inject:""`
	ProjectPerModule  *router.ProjectPermModule `inject:""`
	ProjectMenuModule *router.ProjectMenuModule `inject:""`
	DatapoolModule    *router.DatapoolModule    `inject:""`
	SnippetModule     *router.SnippetModule     `inject:""`

	ImportModule      *router.ImportModule      `inject:""`
	AuthModule        *router.AuthModule        `inject:""`
	EnvironmentModule *router.EnvironmentModule `inject:""`
	ExtractorModule   *router.ExtractorModule   `inject:""`
	ShareVarModule    *router.ShareVarModule    `inject:""`
	CheckpointModule  *router.CheckpointModule  `inject:""`
	ParserModule      *router.ParserModule      `inject:""`

	CategoryModule          *router.CategoryModule          `inject:""`
	ScenarioModule          *router.ScenarioModule          `inject:""`
	ScenarioNodeModule      *router.ScenarioNodeModule      `inject:""`
	ScenarioProcessorModule *router.ScenarioProcessorModule `inject:""`
	ScenarioExecModule      *router.ScenarioExecModule      `inject:""`

	ScenarioInterfaceModule *router.ScenarioInterfaceModule `inject:""`
	//ReportModule             *router.ReportModule              `inject:""`
	EndpointModule          *router.EndpointModule          `inject:""`
	EndpointInterfaceModule *router.EndpointInterfaceModule `inject:""`
	ServeModule             *router.ServeModule             `inject:""`
	PlanModule              *router.PlanModule              `inject:""`
	PlanExecModule          *router.PlanExecModule          `inject:""`
	ScenarioReportModule    *router.ScenarioReportModule    `inject:""`
	PlanReportModule        *router.PlanReportModule        `inject:""`
	SummaryModule           *router.SummaryModule           `inject:""`
	MessageModule           *router.MessageModule           `inject:""`
	DocumentModule          *router.DocumentModule          `inject:""`

	TestsModule *router.TestsModule `inject:""`
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
		m.DataModule.Party(),
		m.FileModule.Party(),
		m.AccountModule.Party(),
		m.RoleModule.Party(),
		m.PermModule.Party(),
		m.UserModule.Party(),

		m.MockModule.Party(),

		m.ProjectModule.Party(),
		m.ProjectPerModule.Party(),
		m.ProjectMenuModule.Party(),
		m.DatapoolModule.Party(),
		m.SnippetModule.Party(),

		m.ImportModule.Party(),
		m.AuthModule.Party(),
		m.EnvironmentModule.Party(),
		m.ExtractorModule.Party(),
		m.ShareVarModule.Party(),
		m.CheckpointModule.Party(),
		m.ParserModule.Party(),

		m.CategoryModule.Party(),
		m.ScenarioModule.Party(),
		m.ScenarioNodeModule.Party(),
		m.ScenarioProcessorModule.Party(),
		m.ScenarioExecModule.Party(),
		m.ScenarioInterfaceModule.Party(),

		//m.ReportModule.Party(),
		m.EndpointModule.Party(),
		m.EndpointInterfaceModule.Party(),
		m.ServeModule.Party(),

		m.PlanModule.Party(),
		m.PlanExecModule.Party(),

		m.ScenarioReportModule.Party(),
		m.PlanReportModule.Party(),
		m.SummaryModule.Party(),

		m.DebugModule.Party(),
		m.TestInterfaceModule.Party(),
		m.MessageModule.Party(),
		m.DocumentModule.Party(),

		m.TestsModule.Party(),
	}
	return module.NewModule(consts.ApiPath, handler, modules...)
}
