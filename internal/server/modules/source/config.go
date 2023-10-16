package source

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ConfigSource struct {
	ConfigRepo *repo.ConfigRepo `inject:""`
}

func (s *ConfigSource) GetSources() ([]model.SysConfig, error) {
	configs := []model.SysConfig{
		{
			Key:   "agentUrlOpts",
			Value: "[     {         \"label\": \"本地\",         \"value\": \"http://127.0.0.1:8086/api/v1\",         \"desc\": \"客户端默认在本机启动的代理\"     },     {         \"label\": \"演示站点\",         \"value\": \"http://111.231.16.35:8086/api/v1\",         \"desc\": \"DeepTest测试站点代理\"     }  ]",
		},
	}

	return configs, nil
}

func (s *ConfigSource) Init() error {
	s.ConfigRepo.DB.Delete(&model.SysConfig{}, "1=1")

	sources, err := s.GetSources()
	if err != nil {
		return err
	}

	for _, source := range sources {
		if err := s.ConfigRepo.Save(source); err != nil { // 遇到错误时回滚事务
			return err
		}
	}
	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.SysConfig{}.TableName())
	return nil
}
