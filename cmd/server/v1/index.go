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

	DebugInterfaceModule    *router.DebugInterfaceModule    `inject:""`
	DebugInvokeModule       *router.DebugInvokeModule       `inject:""`
	DiagnoseInterfaceModule *router.DiagnoseInterfaceModule `inject:""`

	ProjectModule     *router.ProjectModule     `inject:""`
	ProjectPerModule  *router.ProjectPermModule `inject:""`
	ProjectMenuModule *router.ProjectMenuModule `inject:""`
	DatapoolModule    *router.DatapoolModule    `inject:""`
	SnippetModule     *router.SnippetModule     `inject:""`

	MockModule               *router.MockModule               `inject:""`
	EndpointMockScriptModule *router.EndpointMockScriptModule `inject:""`
	EndpointMockExpectModule *router.EndpointMockExpectModule `inject:""`
	MockJsModule             *router.MockJsModule             `inject:""`

	ImportModule      *router.ImportModule      `inject:""`
	AuthModule        *router.AuthModule        `inject:""`
	EnvironmentModule *router.EnvironmentModule `inject:""`
	ShareVarModule    *router.ShareVarModule    `inject:""`
	ParserModule      *router.ParserModule      `inject:""`

	ConditionModule  *router.ConditionModule  `inject:""`
	ExtractorModule  *router.ExtractorModule  `inject:""`
	CheckpointModule *router.CheckpointModule `inject:""`
	ScriptModule     *router.ScriptModule     `inject:""`

	CategoryModule          *router.CategoryModule          `inject:""`
	ScenarioModule          *router.ScenarioModule          `inject:""`
	ScenarioNodeModule      *router.ScenarioNodeModule      `inject:""`
	ScenarioProcessorModule *router.ScenarioProcessorModule `inject:""`
	ScenarioExecModule      *router.ScenarioExecModule      `inject:""`

	ScenarioInterfaceModule *router.ScenarioInterfaceModule `inject:""`
	//ReportModule             *router.ReportModule              `inject:""`
	EndpointModule          *router.EndpointModule          `inject:""`
	EndpointInterfaceModule *router.EndpointInterfaceModule `inject:""`
	EndpointTagModule       *router.EndpointTagModule       `inject:""`

	EndpointCaseModule            *router.EndpointCaseModule            `inject:""`
	EndpointCaseAlternativeModule *router.EndpointCaseAlternativeModule `inject:""`

	ServeModule *router.ServeModule `inject:""`

	PlanModule     *router.PlanModule     `inject:""`
	PlanExecModule *router.PlanExecModule `inject:""`

	PerformanceTestPlanModule *router.PerformanceTestPlanModule `inject:""`
	PerformanceRunnerModule   *router.PerformanceRunnerModule   `inject:""`
	PerformanceExecModule     *router.PerformanceExecModule     `inject:""`

	ScenarioReportModule *router.ScenarioReportModule `inject:""`
	PlanReportModule     *router.PlanReportModule     `inject:""`
	SummaryModule        *router.SummaryModule        `inject:""`
	MessageModule        *router.MessageModule        `inject:""`
	DocumentModule       *router.DocumentModule       `inject:""`
	HealthzModule        *router.HealthzModule        `inject:""`

	ProjectSettingsModule *router.ProjectSettingsModule `inject:""`
	ConfigModule          *router.ConfigModule          `inject:""`
	SysAgentModule        *router.SysAgentModule        `inject:""`
	TestsModule           *router.TestsModule           `inject:""`

	ResponseDefineModule *router.ResponseDefineModule `inject:""`
	JslibModule          *router.JslibModule          `inject:""`

	EndpointCodeModule *router.EndpointCodeModule `inject:""`
	DatabaseConnModule *router.DatabaseConnModule `inject:""`
	DatabaseOptModule  *router.DatabaseOptModule  `inject:""`
	OpenModule         *router.OpenModule         `inject:""`
}

func NewIndexModule() *IndexModule {
	return &IndexModule{}
}

// Party API 模块
func (m *IndexModule) ApiParty() module.WebModule {
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

		m.ProjectModule.Party(),
		m.ProjectPerModule.Party(),
		m.ProjectMenuModule.Party(),
		m.DatapoolModule.Party(),
		m.SnippetModule.Party(),

		m.EndpointCaseAlternativeModule.Party(),
		m.EndpointMockScriptModule.Party(),
		m.EndpointMockExpectModule.Party(),
		m.MockJsModule.Party(),

		m.ImportModule.Party(),
		m.AuthModule.Party(),
		m.EnvironmentModule.Party(),
		m.ShareVarModule.Party(),
		m.ParserModule.Party(),

		m.ConditionModule.Party(),
		m.ExtractorModule.Party(),
		m.CheckpointModule.Party(),
		m.ScriptModule.Party(),

		m.CategoryModule.Party(),
		m.ScenarioModule.Party(),
		m.ScenarioNodeModule.Party(),
		m.ScenarioProcessorModule.Party(),
		m.ScenarioExecModule.Party(),
		m.ScenarioInterfaceModule.Party(),

		m.EndpointModule.Party(),
		m.EndpointInterfaceModule.Party(),
		m.EndpointTagModule.Party(),
		m.EndpointCaseModule.Party(),
		m.EndpointCaseAlternativeModule.Party(),

		m.ServeModule.Party(),

		m.PlanModule.Party(),
		m.PlanExecModule.Party(),
		m.PerformanceTestPlanModule.Party(),
		m.PerformanceRunnerModule.Party(),
		m.PerformanceExecModule.Party(),

		m.ScenarioReportModule.Party(),
		m.PlanReportModule.Party(),
		m.SummaryModule.Party(),

		m.DebugInterfaceModule.Party(),
		m.DebugInvokeModule.Party(),
		m.DiagnoseInterfaceModule.Party(),
		m.MessageModule.Party(),
		m.DocumentModule.Party(),
		m.HealthzModule.Party(),

		m.ProjectSettingsModule.Party(),
		m.ConfigModule.Party(),
		m.SysAgentModule.Party(),
		m.ResponseDefineModule.Party(),

		m.JslibModule.Party(),

		m.EndpointCodeModule.Party(),
		m.DatabaseConnModule.Party(),
		m.DatabaseOptModule.Party(),
		m.OpenModule.Party(),
	}

	return module.NewModule(consts.ApiPathServer, handler, modules...)
}

// Party Mock 模块
func (m *IndexModule) MockParty() module.WebModule {
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
		m.MockModule.Party(),
		m.TestsModule.Party(),
	}
	return module.NewModule(consts.ApiPathMock, handler, modules...)
}
