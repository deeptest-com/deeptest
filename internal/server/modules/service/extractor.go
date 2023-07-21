package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`

	ShareVarService *ShareVarService `inject:""`
}

func (s *ExtractorService) List(debugInterfaceId, endpointInterfaceId uint) (extractors []model.DebugConditionExtractor, err error) {
	extractors, err = s.ExtractorRepo.List(debugInterfaceId, endpointInterfaceId)

	return
}

func (s *ExtractorService) Get(id uint) (extractor model.DebugConditionExtractor, err error) {
	extractor, err = s.ExtractorRepo.Get(id)

	return
}

func (s *ExtractorService) Create(extractor *model.DebugConditionExtractor) (bizErr _domain.BizErr) {
	_, bizErr = s.ExtractorRepo.Save(extractor)

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

func (s *ExtractorService) ListExtractorVariableByInterface(req domain.DebugReq) (variables []domain.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariableByInterface(req)

	return
}
