package source

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	_fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	"github.com/gookit/color"
	"strings"
)

type SampleSource struct {
}

func (s *SampleSource) Init(tenantId consts.TenantId) (err error) {

	/*
		if !config.CONFIG.Saas.Switch {
			return nil
		}
	*/

	var ids []uint
	dao.GetDB(tenantId).Table("biz_endpoint").Pluck("id", &ids)
	if len(ids) > 1 {
		return nil
	}

	sqlStr := _fileUtils.ReadFile("./config/sample/demo.sql")

	sqls := strings.Split(sqlStr, "\n")

	for _, sql := range sqls {
		dao.GetDB(tenantId).Exec(sql)
	}

	color.Info.Printf("\n[Mysql] --> 租户%s初始数据成功!", tenantId)
	return
}
