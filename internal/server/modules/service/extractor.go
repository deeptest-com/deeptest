package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"strings"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`
	InterfaceRepo *repo.InterfaceRepo `inject:""`
}

func (s *ExtractorService) List(interfaceId int, usedBy consts.UsedBy) (extractors []model.InterfaceExtractor, err error) {
	extractors, err = s.ExtractorRepo.List(uint(interfaceId), usedBy)

	return
}

func (s *ExtractorService) Get(id uint) (extractor model.InterfaceExtractor, err error) {
	extractor, err = s.ExtractorRepo.Get(id)

	return
}

func (s *ExtractorService) Create(extractor *model.InterfaceExtractor) (bizErr _domain.BizErr) {
	_, bizErr = s.ExtractorRepo.Save(extractor)

	return
}

func (s *ExtractorService) Update(extractor *model.InterfaceExtractor) (err error) {
	s.ExtractorRepo.Update(extractor)

	return
}

func (s *ExtractorService) CreateOrUpdateResult(extractor *model.InterfaceExtractor, usedBy consts.UsedBy) (err error) {
	s.ExtractorRepo.CreateOrUpdateResult(extractor, usedBy)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ExtractInterface(interfaceId uint, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor, usedBy consts.UsedBy) (logExtractors []domain.ExecInterfaceExtractor, err error) {

	extractors, _ := s.ExtractorRepo.List(interfaceId, usedBy)

	for _, extractor := range extractors {
		logExtractor, err := s.Extract(extractor, resp, interfaceExecLog, usedBy)

		if err == nil && interfaceExecLog != nil { // gen report for processor
			interfaceExtractor := domain.ExecInterfaceExtractor{}
			copier.CopyWithOption(&interfaceExtractor, logExtractor, copier.Option{DeepCopy: true})

			logExtractors = append(logExtractors, interfaceExtractor)
		}
	}

	return
}

func (s *ExtractorService) Extract(extractor model.InterfaceExtractor, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor, usedBy consts.UsedBy) (logExtractor model.ExecLogExtractor, err error) {

	s.ExtractValue(&extractor, resp)

	s.ExtractorRepo.UpdateResult(extractor, usedBy)

	return
}

func (s *ExtractorService) ExtractValue(extractor *model.InterfaceExtractor, resp v1.InvocationResponse) (err error) {
	if extractor.Disabled {
		extractor.Result = ""
		return
	}

	if extractor.Src == consts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Key {
				extractor.Result = h.Value
				break
			}
		}
	} else {
		if httpHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
			extractor.Result = queryHelper.JsonQuery(resp.Content, extractor.Expression)

		} else if httpHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
			extractor.Result = queryHelper.HtmlQuery(resp.Content, extractor.Expression)

		} else if httpHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
			extractor.Result = queryHelper.XmlQuery(resp.Content, extractor.Expression)

		} else if extractor.Type == consts.Boundary {
			extractor.Result = queryHelper.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
				extractor.BoundaryIndex, extractor.BoundaryIncluded)
		}
	}

	extractor.Result = strings.TrimSpace(extractor.Result)

	return
}

func (s *ExtractorService) ListValidExtractorVariable(interfaceId int) (variables []v1.Variable, err error) {
	interf, _ := s.InterfaceRepo.Get(uint(interfaceId))
	variables, err = s.ExtractorRepo.ListValidExtractorVariable(uint(interfaceId), interf.ProjectId)

	return
}

func (s *ExtractorService) ListExtractorVariableByInterface(interfaceId int) (variables []v1.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariableByInterface(uint(interfaceId))

	return
}
