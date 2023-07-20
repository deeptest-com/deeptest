package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
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

func (s *ExtractorService) CreateOrUpdateResult(extractor *model.DebugConditionExtractor, usedBy consts.UsedBy) (err error) {
	s.ExtractorRepo.CreateOrUpdateResult(extractor, usedBy)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ExtractInterface(invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, resp domain.DebugResponse, usedBy consts.UsedBy) (err error) {
	extractors, _ := s.ExtractorRepo.List(debugInterfaceId, endpointInterfaceId)

	for _, extractor := range extractors {
		s.Extract(&extractor, resp, invokeId, usedBy)
		s.ShareVarService.Save(extractor.Variable, extractor.Result, invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, extractor.Scope, usedBy)
	}

	return
}

func (s *ExtractorService) Extract(extractor *model.DebugConditionExtractor, resp domain.DebugResponse,
	invokeId uint, usedBy consts.UsedBy) (err error) {

	extractor.Result, err = extractorHelper.Extract(extractor.ExtractorBase, resp)
	if err != nil {
		return
	}

	s.ExtractorRepo.UpdateResult(*extractor, usedBy)
	s.ExtractorRepo.CreateLog(*extractor, invokeId, usedBy)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(req domain.DebugReq) (variables []domain.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariableByInterface(req)

	return
}
