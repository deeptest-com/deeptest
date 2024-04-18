package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/saas/tenant"
	"github.com/gookit/color"
)

type SysAgentSource struct {
	AgentRepo *repo.SysAgentRepo `inject:""`
}

func (s *SysAgentSource) GetSources(tenantId consts.TenantId) (configs []model.SysAgent, err error) {
	if config.CONFIG.System.SysEnv == "deeptest_demo" {
		configs = []model.SysAgent{
			{
				BaseModel: model.BaseModel{ID: 1},
				Name:      "本地代理",
				Url:       "http://127.0.0.1:8086/api/v1",
				Desc:      "客户端默认在本机启动的代理",
			},
			{
				BaseModel: model.BaseModel{ID: 2},
				Name:      "演示站点代理",
				Url:       "http://111.231.16.35:8086/api/v1",
				Desc:      "DeepTest演示站点上的代理",
			},
		}
	}

	isFree := tenant.NewTenant().ForFree(tenantId)

	if config.CONFIG.Saas.Switch && !isFree {
		configs = []model.SysAgent{
			{
				BaseModel: model.BaseModel{ID: 1},
				Name:      "默认代理",
				Url:       "/lya/api/v1",
				Desc:      "付费版本免费代理",
			},
		}
	}

	return
}

func (s *SysAgentSource) Init(tenantId consts.TenantId) (err error) {
	if config.CONFIG.System.SysEnv != "deeptest_demo" {
		return
	}

	//db := s.AgentRepo.GetDB(tenantId)
	//db.Delete(&model.SysAgent{}, "1=1")

	sources, err := s.GetSources(tenantId)
	if err != nil {
		return err
	}

	for _, source := range sources {
		if err := s.AgentRepo.Save(tenantId, &source); err != nil {
			return err
		}
	}
	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.SysAgent{}.TableName())
	return nil
}
