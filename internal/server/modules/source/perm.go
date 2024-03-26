package source

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type PermSource struct {
	PermRepo *repo.PermRepo `inject:""`
}

func (s *PermSource) GetSources() []model.SysPerm {
	permRouteLen := len(config.PermRoutes)
	ch := make(chan model.SysPerm, permRouteLen)

	for _, permRoute := range config.PermRoutes {
		p := permRoute
		go func(permRoute map[string]string) {
			perm := model.SysPerm{PermBase: v1.PermBase{
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
		color.Info.Printf("\n[Mysql] --> %s table success to create %d records\n", model.SysPerm{}.TableName(), count)
	}

	return nil
}
