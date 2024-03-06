package source

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type ProjectPermSource struct {
	ProjectPermRepo *repo.ProjectPermRepo `inject:""`
}

func (s *ProjectPermSource) GetSources() []model.ProjectPerm {
	permRouteLen := len(config.PermRoutes)
	ch := make(chan model.ProjectPerm, permRouteLen)

	for _, permRoute := range config.PermRoutes {
		p := permRoute
		go func(permRoute map[string]string) {
			perm := model.ProjectPerm{ProjectPermBase: v1.ProjectPermBase{
				Name:        permRoute["path"],
				DisplayName: permRoute["name"],
				Description: permRoute["name"],
				Act:         permRoute["act"],
			}}
			ch <- perm
		}(p)
	}
	perms := make([]model.ProjectPerm, permRouteLen)
	for i := 0; i < permRouteLen; i++ {
		perms[i] = <-ch
	}
	return perms
}

func (s *ProjectPermSource) Init(tenantId consts.TenantId) error {
	sources := s.GetSources()

	successCount, failItems := s.ProjectPermRepo.CreateIfNotExist(tenantId, sources)
	color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据,失败数据：%+v!\n", model.ProjectPerm{}.TableName(), successCount, failItems)

	return nil
}
