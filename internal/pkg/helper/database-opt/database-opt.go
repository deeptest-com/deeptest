package databaseOptHelpper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
)

func GenDesc(typ consts.DatabaseType, sql string) (ret string) {

	ret = fmt.Sprintf("执行%s查询", _i118Utils.Sprintf(typ.String()))

	return
}

func GenResultMsg(po *domain.DatabaseOptBase) {
	desc := GenDesc(po.Type, po.Sql)

	if po.DatabaseConnIsDisabled {
		po.ResultMsg = fmt.Sprintf("%s，返回\"%s\"。", desc, po.Result)
		return
	}

	po.ResultMsg = fmt.Sprintf("%s，表达式%s提取变量%s为\"%s\"。", desc, po.JsonPath, po.Variable, po.Result)

	return
}
