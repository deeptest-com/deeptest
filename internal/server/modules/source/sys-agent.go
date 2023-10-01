package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type SysAgentSource struct {
	AgentRepo *repo.SysAgentRepo `inject:""`
}

func (s *SysAgentSource) GetSources() (configs []model.SysAgent, err error) {
	if config.CONFIG.System.SysEnv == "deeptest_demo" {
		configs = []model.SysAgent{
			{
				Name: "本地代理",
				Code: "local",
				Url:  "http://127.0.0.1:8086/api/v1",
				Desc: "客户端默认在本机启动的代理",
			},
			{
				Name: "演示站点代理",
				Code: "test",
				Url:  "http://111.231.16.35:8086/api/v1",
				Desc: "DeepTest演示站点上的代理",
			},
		}
	}

	return
}

func (s *SysAgentSource) Init() (err error) {
	if config.CONFIG.System.SysEnv != "deeptest_demo" {
		return
	}

	s.AgentRepo.DB.Delete(&model.SysAgent{}, "1=1")

	sources, err := s.GetSources()
	if err != nil {
		return err
	}

	for _, source := range sources {
		if err := s.AgentRepo.Save(&source); err != nil {
			return err
		}
	}
	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.SysAgent{}.TableName())
	return nil
}
