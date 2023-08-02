package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type CommonService struct {
	BaseRepo *repo.BaseRepo `inject:""`
}

func (s *CommonService) GetEntityByModule(module string) (entity interface{}) {
	moduleEntityMap := map[string]interface{}{
		"endpoint": model.Endpoint{},
		"scenario": model.Scenario{},
		"plan":     model.Plan{},
		"report":   model.PlanReport{},
	}

	if v, ok := moduleEntityMap[module]; ok {
		entity = v
	}

	return
}

func (s *CommonService) BatchUpdateField(req v1.BatchUpdateReq) error {
	entity := s.GetEntityByModule(req.Module)
	return s.BaseRepo.BatchUpdateStatus(entity, req.Ids, req.FieldName, req.Value)
}
