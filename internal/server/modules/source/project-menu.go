package source

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectMenuSource struct {
	ProjectMenuRepo *repo2.ProjectMenuRepo `inject:""`
}

func (s *ProjectMenuSource) Init(tenantId consts.TenantId) (err error) {
	defer s.ProjectMenuRepo.BatchInitData(tenantId, "buttonLevel")
	defer s.ProjectMenuRepo.BatchInitData(tenantId, "secondLevel")
	defer s.ProjectMenuRepo.BatchInitData(tenantId, "firstLevel")
	defer s.ProjectMenuRepo.DeleteAllData(tenantId)

	return
}
