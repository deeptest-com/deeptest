package model

import (
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
)

var (
	Models = []interface{}{
		&middleware.Oplog{},

		&SysPerm{},
		&SysRole{},
		&SysUser{},
		&SysUserProfile{},

		&ProjectRole{},
		&Org{},
		&Project{},
		&ProjectMember{},
		&Datapool{},
		&Environment{},
		&EnvironmentVar{},
		&ShareVariable{},

		&DebugInterface{},
		&DebugInterfaceParam{},
		&DebugInterfaceBodyFormDataItem{},
		&DebugInterfaceBodyFormUrlEncodedItem{},
		&DebugInterfaceHeader{},
		&DebugInterfaceBasicAuth{},
		&DebugInterfaceBearerToken{},
		&DebugInterfaceOAuth20{},
		&DebugInterfaceApiKey{},

		&DebugPreCondition{},
		&DebugPostCondition{},
		&DebugConditionExtractor{},
		&DebugConditionCheckpoint{},
		&DebugConditionScript{},

		&DiagnoseInterface{},

		&Snippet{},
		&MockJsExpression{},

		&MockInvocation{},
		&Auth2Token{},

		&Category{},
		&Scenario{},

		&Plan{},
		&RelaPlanScenario{},

		&Processor{},
		//&ProcessorThreadGroup{},
		&ProcessorGroup{},
		&ProcessorLogic{},
		&ProcessorLoop{},
		&ProcessorTimer{},
		&ProcessorPrint{},
		&ProcessorVariable{},
		&ProcessorAssertion{},
		&ProcessorData{},
		&ProcessorCookie{},
		&ProcessorExtractor{},
		&ProcessorCustomCode{},

		&ScenarioReport{},
		&PlanReport{},
		&ExecLogProcessor{},
		&ExecLogExtractor{},
		&ExecLogCheckpoint{},
		&ExecLogScript{},

		&ComponentSchema{},
		&ComponentSchemaSecurity{},

		&Endpoint{},
		&EndpointPathParam{},
		&EndpointInterfaceRequestBody{},
		&EndpointInterfaceRequestBodyItem{},
		&EndpointInterfaceResponseBodyItem{},
		&EndpointInterfaceResponseBodyHeader{},
		&EndpointInterfaceResponseBody{},
		&EndpointInterface{},
		&EndpointCase{},
		&EndpointInterfaceParam{},
		&EndpointInterfaceCookie{},
		&EndpointInterfaceHeader{},
		&EndpointDocument{},
		&EndpointSnapshot{},
		&EndpointTag{},
		&EndpointTagRel{},

		&Serve{},
		&ServeServer{},
		&ServeVersion{},
		&EndpointVersion{},
		&ServeEndpointVersion{},
		&SummaryBugs{},
		&SummaryDetails{},
		&SummaryProjectUserRanking{},
		&EnvironmentParam{},
		&Message{},
		&MessageRead{},
		&DebugInvoke{},
		&ProjectPerm{},
		&ProjectRolePerm{},
		&ProjectRoleMenu{},
		&ProjectMenu{},
		&ProjectRecentlyVisited{},
		&ProjectMemberAudit{},

		&SwaggerSync{},
		&ProjectMockSetting{},

		&EndpointMockScript{},
		&EndpointMockExpect{},
		&EndpointMockExpectRequest{},
		&EndpointMockExpectResponse{},
		&EndpointMockExpectResponseHeader{},

		&DebugConditionResponseDefine{},
		&ExecLogResponseDefine{},

		&SysConfig{},
		&SysJslib{},
		&SysAgent{},
		&ProjectUserServer{},
	}
)
