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

		&Interface{},
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
		&Snippet{},

		&Invocation{},
		&Auth2Token{},

		&ScenarioCategory{},
		&Scenario{},

		&PlanCategory{},
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
		&InterfaceCookie{},
		&Endpoint{},
		&EndpointPathParam{},
		&InterfaceRequestBody{},
		&InterfaceRequestBodyItem{},
		&InterfaceResponseBodyItem{},
		&InterfaceResponseBodyHeader{},
		&Serve{},
		&ServeServer{},
		&ServeVersion{},
		&EndpointVersion{},
		&ServeEndpointVersion{},
		&SummaryBugs{},
		&SummaryDetails{},
		&SummaryProjectUserRanking{},
		&EnvironmentParam{},
	}
)
