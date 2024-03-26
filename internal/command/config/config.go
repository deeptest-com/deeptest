package commandConfig

import (
	commandConsts "github.com/aaronchen2k/deeptest/internal/command/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

func InitConfig() {
	consts.IsRelease = _commUtils.IsRelease()
	consts.ExecDir = commUtils.GetExecDir()

	if !consts.IsRelease {
		log.Println("ExecDir=" + consts.ExecDir)
	}

	commandConsts.DB, _ = NewGormDB()

	return
}

func NewGormDB() (gormDb *gorm.DB, err error) {
	commandConsts.SqliteFile = "file:" + filepath.Join(consts.ExecDir, "deeptest.db")
	gormDb, err = gorm.Open(sqlite.Open(commandConsts.SqliteFile), &gorm.Config{})

	if consts.Verbose {
		gormDb = gormDb.Debug()
	}

	return
}
