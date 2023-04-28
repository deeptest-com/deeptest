package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`
	InterfaceRepo *repo.InterfaceRepo `inject:""`

	ShareVarService *ShareVarService `inject:""`
}

func (s *ExtractorService) List(interfaceId uint, usedBy consts.UsedBy) (extractors []model.DebugInterfaceExtractor, err error) {
	extractors, err = s.ExtractorRepo.List(interfaceId)

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

func (s *ExtractorService) ExtractInterface(interfaceId, serveId, scenarioId uint, resp v1.DebugResponse, usedBy consts.UsedBy) (err error) {
	extractors, _ := s.ExtractorRepo.List(interfaceId)

	for _, extractor := range extractors {
		s.Extract(&extractor, resp, usedBy)
		s.ShareVarService.Save(extractor.Variable, extractor.Result, interfaceId, serveId, scenarioId, extractor.Scope, usedBy)
	}

	return
}

func (s *ExtractorService) Extract(extractor *model.DebugInterfaceExtractor, resp v1.DebugResponse,
	usedBy consts.UsedBy) (err error) {

	extractor.Result, err = extractorHelper.Extract(extractor.ExtractorBase, resp)
	if err != nil {
		return
	}

	s.ExtractorRepo.UpdateResult(*extractor, usedBy)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(interfaceId int) (variables []v1.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariableByInterface(uint(interfaceId))

	return
}

func (s *ExtractorService) ListValidExtractorVarForInterface(interfaceId int, usedBy consts.UsedBy) (variables []v1.Variable, err error) {
	interf, _ := s.InterfaceRepo.Get(uint(interfaceId))

	variables, err = s.ExtractorRepo.ListValidExtractorVariableForInterface(uint(interfaceId), interf.ProjectId, usedBy)

	return
}
