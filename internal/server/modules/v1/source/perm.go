package source

import (
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/gookit/color"
)

type PermSource struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermSource() *PermSource {
	return &PermSource{}
}

func (s *PermSource) GetSources() []model.SysPerm {
	permRouteLen := len(serverConfig.PermRoutes)
	ch := make(chan model.SysPerm, permRouteLen)

	for _, permRoute := range serverConfig.PermRoutes {
		p := permRoute
		go func(permRoute map[string]string) {
			perm := model.SysPerm{BasePerm: model.BasePerm{
				Name:        permRoute["path"],
				DisplayName: permRoute["name"],
				Description: permRoute["name"],
				Act:         permRoute["act"],
			}}
			ch <- perm
		}(p)
	}
	perms := make([]model.SysPerm, permRouteLen)
	for i := 0; i < permRouteLen; i++ {
		perms[i] = <-ch
	}
	return perms
}

func (s *PermSource) Init() error {
	sources := s.GetSources()

	count, err := s.PermRepo.CreateIfNotExist(sources)
	if err == nil {
		color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据!\n", model.SysPerm{}.TableName(), count)
	}

	return nil
}
