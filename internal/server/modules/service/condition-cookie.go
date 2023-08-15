package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	cookieHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cookie"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type CookieService struct {
	CookieRepo *repo.CookieRepo `inject:""`

	PostConditionRepo    *repo.PostConditionRepo `inject:""`
	PostConditionService *PostConditionService   `inject:""`
	ShareVarService      *ShareVarService        `inject:""`
}

func (s *CookieService) Get(id uint) (extractor model.DebugConditionCookie, err error) {
	extractor, err = s.CookieRepo.Get(id)

	return
}

func (s *CookieService) Create(extractor *model.DebugConditionCookie) (err error) {
	_, err = s.CookieRepo.Save(extractor)

	return
}
func (s *CookieService) QuickCreate(req serverDomain.CookieConditionQuickCreateReq, usedBy consts.UsedBy) (err error) {
	debugInfo := req.Info
	config := req.Config

	// create post-condition
	condition := model.DebugPostCondition{}
	copier.CopyWithOption(&condition, debugInfo, copier.Option{DeepCopy: true})

	condition.EntityId = 0 // update later
	condition.EntityType = consts.ConditionTypeCookie
	condition.UsedBy = debugInfo.UsedBy
	condition.Desc = cookieHelper.GenDesc(config.CookieName, config.VariableName)

	err = s.PostConditionRepo.Save(&condition)

	// create extractor
	var extractor model.DebugConditionCookie
	copier.CopyWithOption(&extractor, config, copier.Option{DeepCopy: true})
	extractor.ConditionId = condition.ID

	_, err = s.CookieRepo.Save(&extractor)

	s.PostConditionRepo.UpdateEntityId(condition.ID, extractor.ID)

	return
}

func (s *CookieService) Update(extractor *model.DebugConditionCookie) (err error) {
	s.CookieRepo.Update(extractor)

	return
}

func (s *CookieService) Delete(reqId uint) (err error) {
	err = s.CookieRepo.Delete(reqId)

	return
}

func (s *CookieService) ListCookieVariableByInterface(req domain.DebugInfo) (variables []domain.Variable, err error) {
	extractorConditions, err := s.PostConditionRepo.ListExtractor(req.DebugInterfaceId, req.EndpointInterfaceId)

	var conditionIds []uint
	for _, item := range extractorConditions {
		conditionIds = append(conditionIds, item.ID)
	}

	variables, err = s.CookieRepo.ListCookieVariableByInterface(conditionIds)

	return
}
