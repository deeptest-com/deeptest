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

		&Invocation{},
		&Environment{},
		&EnvironmentVar{},

		&Auth2Token{},

		&Scenario{},
		&Processor{},
		//&ProcessorThreadGroup{},
		&ProcessorGroup{},
		&ProcessorLogic{},
		&ProcessorLoop{},
		&ProcessorTimer{},
		&ProcessorPrint{},
		&ProcessorVariable{},
		&ProcessorAssertion{},
		&ProcessorExtractor{},
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
		&Report{},
		&ExecLogProcessor{},
		&ExecLogExtractor{},
		&ExecLogCheckpoint{},
	}
)
