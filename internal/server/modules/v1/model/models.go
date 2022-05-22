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

		&TestConfig{},
		&TestScenario{},
		&TestSet{},
		&TestProcessor{},
		&ProcessorThreadGroup{},
		&ProcessorSimple{},
		&ProcessorFlow{},
		&ProcessorIterator{},
		&ProcessorTimer{},
		&ProcessorAssertion{},
		&ProcessorExtractor{},
		&ProcessorData{},
		&ProcessorCookie{},
	}
)
