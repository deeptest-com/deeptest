package source

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type SysConfigSource struct {
	ConfigRepo *repo.ConfigRepo `inject:""`
}

func (s *SysConfigSource) GetSources() (configs []model.SysConfig, err error) {
	return
}

func (s *SysConfigSource) Init(consts.TenantId) (err error) {
	return
}
