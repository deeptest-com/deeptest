package databaseOptHelpper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func GenDesc(typ consts.DatabaseType, sql string) (ret string) {

	ret = fmt.Sprintf("%s数据库中执行SQL\"%s\"", typ, sql)

	return
}

func GenResultMsg(po *domain.DatabaseOptBase) {
	po.ResultMsg = GenDesc(po.Type, po.Sql)

	if po.ResultStatus != consts.Pass {
		po.ResultMsg += fmt.Sprintf("，输出\"%s\"。", po.ResultMsg)
	}

	return
}
