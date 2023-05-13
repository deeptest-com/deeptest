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

		&InterfaceParam{},
		&InterfaceBodyFormDataItem{},
		&InterfaceBodyFormUrlEncodedItem{},
		&InterfaceHeader{},
		&InterfaceBasicAuth{},
		&InterfaceBearerToken{},
		&InterfaceOAuth20{},
		&InterfaceApiKey{},
		&InterfaceExtractor{},
		&InterfaceCheckpoint{},

		&DebugInterface{},
		&DebugInterfaceParam{},
		&DebugInterfaceBodyFormDataItem{},
		&DebugInterfaceBodyFormUrlEncodedItem{},
		&DebugInterfaceHeader{},
		&DebugInterfaceBasicAuth{},
		&DebugInterfaceBearerToken{},
		&DebugInterfaceOAuth20{},
		&DebugInterfaceApiKey{},
		&DebugInterfaceExtractor{},
		&DebugInterfaceCheckpoint{},

		&Snippet{},

		&Invocation{},
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

		&ProcessorInterface{},
		&ProcessorInterfaceParam{},
		&ProcessorInterfaceBodyFormDataItem{},
		&ProcessorInterfaceBodyFormUrlEncodedItem{},
		&ProcessorInterfaceHeader{},
		&ProcessorInterfaceBasicAuth{},
		&ProcessorInterfaceBearerToken{},
		&ProcessorInterfaceOAuth20{},
		&ProcessorInterfaceApiKey{},

		&ProcessorInvocation{},
		&ScenarioReport{},
		&ExecLogProcessor{},
		&ExecLogExtractor{},
		&ExecLogCheckpoint{},

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
		&EndpointInterfaceParam{},
		&EndpointInterfaceCookie{},
		&EndpointInterfaceHeader{},

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
	}
)
