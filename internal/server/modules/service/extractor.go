package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`

	ShareVarService *ShareVarService `inject:""`
}

func (s *ExtractorService) List(debugInterfaceId, endpointInterfaceId uint) (extractors []model.DebugInterfaceExtractor, err error) {
	extractors, err = s.ExtractorRepo.List(debugInterfaceId, endpointInterfaceId)

	return
}

func (s *ExtractorService) Get(id uint) (extractor model.DebugInterfaceExtractor, err error) {
	extractor, err = s.ExtractorRepo.Get(id)

	return
}

func (s *ExtractorService) Create(extractor *model.DebugInterfaceExtractor) (bizErr _domain.BizErr) {
	_, bizErr = s.ExtractorRepo.Save(extractor)

	return
}

func (s *ExtractorService) Update(extractor *model.DebugInterfaceExtractor) (err error) {
	s.ExtractorRepo.Update(extractor)

	return
}

func (s *ExtractorService) CreateOrUpdateResult(extractor *model.DebugInterfaceExtractor, usedBy consts.UsedBy) (err error) {
	s.ExtractorRepo.CreateOrUpdateResult(extractor, usedBy)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ExtractInterface(debugInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, resp domain.DebugResponse, usedBy consts.UsedBy) (err error) {
	extractors, _ := s.ExtractorRepo.List(debugInterfaceId, endpointInterfaceId)

	for _, extractor := range extractors {
		s.Extract(&extractor, resp, usedBy)
		s.ShareVarService.Save(extractor.Variable, extractor.Result, debugInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, extractor.Scope, usedBy)
	}

	return
}

func (s *ExtractorService) Extract(extractor *model.DebugInterfaceExtractor, resp domain.DebugResponse,
	usedBy consts.UsedBy) (err error) {

	extractor.Result, err = extractorHelper.Extract(extractor.ExtractorBase, resp)
	_logUtils.Infof(fmt.Sprintf("提取器调试 extractor.Result:%+v, err:%+v, usedBy:%+v", extractor.Result, err, usedBy))

	if err != nil {
		return
	}

	s.ExtractorRepo.UpdateResult(*extractor, usedBy)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(req domain.DebugReq) (variables []domain.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariableByInterface(req)

	return
}
