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
		&InterfaceExtractor{},
		&InterfaceCheckpoint{},

		&Invocation{},
		&Environment{},
		&EnvironmentVar{},
	}
)
