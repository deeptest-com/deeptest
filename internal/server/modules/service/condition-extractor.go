package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`

	PostConditionRepo    *repo.PostConditionRepo `inject:""`
	PostConditionService *PostConditionService   `inject:""`
	ShareVarService      *ShareVarService        `inject:""`
}

func (s *ExtractorService) Get(id uint) (extractor model.DebugConditionExtractor, err error) {
	extractor, err = s.ExtractorRepo.Get(id)

	return
}

func (s *ExtractorService) Create(extractor *model.DebugConditionExtractor) (err error) {
	_, err = s.ExtractorRepo.Save(extractor)

	return
}
func (s *ExtractorService) QuickCreate(req serverDomain.ExtractorConditionQuickCreateReq, usedBy consts.UsedBy) (err error) {
	debugInfo := req.Info
	config := req.Config

	// create post-condition
	condition := model.DebugPostCondition{}
	copier.CopyWithOption(&condition, debugInfo, copier.Option{DeepCopy: true})

	condition.EntityId = 0 // update later
	condition.EntityType = consts.ConditionTypeExtractor
	condition.UsedBy = debugInfo.UsedBy
	condition.Desc = extractorHelper.GenDesc(config.Variable, config.Src, config.Key, config.Type, config.Expression, "", "")

	err = s.PostConditionRepo.Save(&condition)

	// create extractor
	var extractor model.DebugConditionExtractor
	copier.CopyWithOption(&extractor, config, copier.Option{DeepCopy: true})
	extractor.ConditionId = condition.ID

	_, err = s.ExtractorRepo.Save(&extractor)

	s.PostConditionRepo.UpdateEntityId(condition.ID, extractor.ID)

	return
}

func (s *ExtractorService) Update(extractor *model.DebugConditionExtractor) (err error) {
	s.ExtractorRepo.Update(extractor)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(req domain.DebugInfo) (variables []domain.Variable, err error) {
	extractorConditions, err := s.PostConditionRepo.ListExtractor(req.DebugInterfaceId, req.EndpointInterfaceId)
	var conditionIds1 []uint
	for _, item := range extractorConditions {
		conditionIds1 = append(conditionIds1, item.ID)
	}
	variables1, err := s.ExtractorRepo.ListExtractorVariableByInterface(conditionIds1)

	var conditionIds2 []uint
	dbOptConditions, err := s.PostConditionRepo.ListDbOpt(req.DebugInterfaceId, req.EndpointInterfaceId)
	for _, item := range dbOptConditions {
		conditionIds2 = append(conditionIds2, item.ID)
	}
	variables2, err := s.ExtractorRepo.ListDbOptVariableByInterface(conditionIds2)

	// combine
	mp := map[string]bool{}
	for _, item := range variables1 {
		mp[item.Name] = true
		variables = append(variables, item)
	}
	for _, item := range variables2 {
		if mp[item.Name] {
			continue
		}
		variables = append(variables, item)
	}

	return
}
