package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/gookit/color"
	"strings"
)

type SampleService struct {
	ProjectRolePermService *ProjectRolePermService `inject:""`
	ProjectRepo            *repo.ProjectRepo       `inject:""`
}

func (s *SampleService) Init(tenantId consts.TenantId) (err error) {

	if config.CONFIG.System.SysEnv == "ly" {
		_, err = s.ProjectRolePermService.GetRoleFromOther(tenantId)
		if err != nil { // 遇到错误时回滚事务
			logUtils.Errorf("[Mysql] --> %s 表初始数据失败!,err:%s", model.ProjectRole{}.TableName(), err.Error())
			return nil
		}
	}

	project := v1.ProjectReq{ProjectBase: v1.ProjectBase{Name: "示例项目", AdminId: 1, Logo: "default_logo1", ShortName: "Demo", Desc: "示例项目包含样例数据，用于展示API管理基本功能和使用"}}

	//创建项目
	s.ProjectRepo.Create(tenantId, project, 1)

	if !config.CONFIG.Saas.Switch {
		return nil
	}

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
