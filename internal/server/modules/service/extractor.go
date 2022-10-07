package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"strings"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo   `inject:""`
	InterfaceRepo *repo.InterfaceRepo   `inject:""`
	ExecContext   *business.ExecContext `inject:""`
}

func (s *ExtractorService) List(interfaceId int) (extractors []model.InterfaceExtractor, err error) {
	extractors, err = s.ExtractorRepo.List(uint(interfaceId))

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

func (s *ExtractorService) Update(extractor *model.InterfaceExtractor) (err _domain.BizErr) {
	_, err = s.ExtractorRepo.Save(extractor)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ExtractInterface(interf model.Interface, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor) (logExtractors []domain.ExecInterfaceExtractor, err error) {
	extractors, _ := s.ExtractorRepo.List(interf.ID)

	for _, extractor := range extractors {
		logExtractor, err := s.Extract(extractor, resp, interfaceExecLog)

		if err == nil && interfaceExecLog != nil { // gen report for processor
			interfaceExtractor := domain.ExecInterfaceExtractor{}
			copier.CopyWithOption(&interfaceExtractor, logExtractor, copier.Option{DeepCopy: true})

			logExtractors = append(logExtractors, interfaceExtractor)
		}
	}

	return
}

func (s *ExtractorService) Extract(extractor model.InterfaceExtractor, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor) (logExtractor model.ExecLogExtractor, err error) {

	s.ExtractValue(&extractor, resp)

	if interfaceExecLog == nil { // run by interface
		s.ExtractorRepo.UpdateResult(extractor)

	} else { // run by processor
		s.ExecContext.SetVariable(interfaceExecLog.ProcessorId, extractor.Variable, extractor.Result,
			extractor.Scope)

		logExtractor, err = s.ExtractorRepo.UpdateResultToExecLog(extractor, interfaceExecLog)

	}

	return
}

func (s *ExtractorService) ExtractValue(extractor *model.InterfaceExtractor, resp v1.InvocationResponse) (err error) {
	if extractor.Disabled {
		extractor.Result = ""
	} else {
		if extractor.Src == consts.Header {
			for _, h := range resp.Headers {
				if h.Name == extractor.Key {
					extractor.Result = h.Value
					break
				}
			}
		} else {
			if requestHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
				extractor.Result = queryHelper.JsonQuery(resp.Content, extractor.Expression)

			} else if requestHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
				extractor.Result = queryHelper.HtmlQuery(resp.Content, extractor.Expression)

			} else if requestHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
				extractor.Result = queryHelper.XmlQuery(resp.Content, extractor.Expression)

			} else if extractor.Type == consts.Boundary {
				extractor.Result = queryHelper.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
					extractor.BoundaryIndex, extractor.BoundaryIncluded)
			}
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
