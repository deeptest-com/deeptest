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
	ConditionRepo *repo.ConditionRepo `inject:""`

	ConditionService *ConditionService `inject:""`
	ShareVarService  *ShareVarService  `inject:""`
}

func (s *ExtractorService) Get(tenantId consts.TenantId, id uint) (extractor model.DebugConditionExtractor, err error) {
	extractor, err = s.ExtractorRepo.Get(tenantId, id)

	return
}

func (s *ExtractorService) Create(tenantId consts.TenantId, extractor *model.DebugConditionExtractor) (err error) {
	_, err = s.ExtractorRepo.Save(tenantId, extractor)

	return
}
func (s *ExtractorService) QuickCreate(tenantId consts.TenantId, req serverDomain.ExtractorConditionQuickCreateReq, usedBy consts.UsedBy) (err error) {
	debugInfo := req.Info
	config := req.Config

	// create post-condition
	condition := model.DebugCondition{
		ConditionSrc: req.ConditionSrc,
	}
	copier.CopyWithOption(&condition, debugInfo, copier.Option{DeepCopy: true})

	condition.EntityId = 0 // update later
	condition.EntityType = consts.ConditionTypeExtractor
	condition.UsedBy = debugInfo.UsedBy
	condition.Desc = extractorHelper.GenDesc(config.Variable, config.Src, config.Key, config.Type, config.Expression, "", "")

	err = s.ConditionRepo.Save(tenantId, &condition)

	// create extractor
	extractor := model.DebugConditionExtractor{
		ExtractorBase: domain.ExtractorBase{
			Src: consts.Body,
		},
	}
	copier.CopyWithOption(&extractor, config, copier.Option{DeepCopy: true})
	extractor.ConditionId = condition.ID

	_, err = s.ExtractorRepo.Save(tenantId, &extractor)

	s.ConditionRepo.UpdateEntityId(tenantId, condition.ID, extractor.ID)

	return
}

func (s *ExtractorService) Update(tenantId consts.TenantId, extractor *model.DebugConditionExtractor) (err error) {
	s.ExtractorRepo.Update(tenantId, extractor)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(tenantId consts.TenantId, req domain.DebugInfo) (variables []domain.Variable, err error) {
	extractorConditions, err := s.ConditionRepo.ListExtractor(tenantId, req)
	var conditionIds1 []uint
	for _, item := range extractorConditions {
		conditionIds1 = append(conditionIds1, item.ID)
	}
	variables1, err := s.ExtractorRepo.ListExtractorVariableByConditions(tenantId, conditionIds1)

	var conditionIds2 []uint
	dbOptConditions, err := s.ConditionRepo.ListDbOpt(tenantId, req)
	for _, item := range dbOptConditions {
		conditionIds2 = append(conditionIds2, item.ID)
	}
	variables2, err := s.ExtractorRepo.ListDbOptVariableByConditions(tenantId, conditionIds2)

	// combine
	mp := map[string]bool{}
	for _, item := range variables1 {
		if item.Name == "" || mp[item.Name] {
			continue
		}

		mp[item.Name] = true
		variables = append(variables, item)
	}
	for _, item := range variables2 {
		if item.Name == "" || mp[item.Name] {
			continue
		}

		mp[item.Name] = true
		variables = append(variables, item)
	}

	return
}
