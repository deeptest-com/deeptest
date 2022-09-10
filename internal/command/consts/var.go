package commandConsts

import (
	"gorm.io/gorm"
	"sync"
)

var (
	SqliteFile = ""

	DB     *gorm.DB
	EnvVar sync.Map
)
