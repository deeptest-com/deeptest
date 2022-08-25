package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/query"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_cacheUtils "github.com/aaronchen2k/deeptest/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strconv"
)

type ExtractorService struct {
	ExtractorRepo *repo.ExtractorRepo `inject:""`
	InterfaceRepo *repo.InterfaceRepo `inject:""`
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

func (s *ExtractorService) ExtractInterface(interf model.Interface, resp serverDomain.InvocationResponse,
	interfaceLog *model.Log) (err error) {
	extractors, _ := s.ExtractorRepo.List(interf.ID)

	for _, extractor := range extractors {
		s.Extract(extractor, resp, interf.ProjectId, interfaceLog)
	}

	return
}

func (s *ExtractorService) Extract(extractor model.InterfaceExtractor, resp serverDomain.InvocationResponse,
	projectId uint, interfaceLog *model.Log) (err error) {
	if extractor.Disabled {
		extractor.Result = ""

		if interfaceLog == nil { // run by interface
			s.ExtractorRepo.UpdateResult(extractor)
		} else { // run by processor

		}

		return
	}

	if extractor.Src == consts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Key {
				extractor.Result = h.Value
				break
			}
		}

		if interfaceLog == nil { // run by interface
			s.ExtractorRepo.UpdateResult(extractor)
		} else { // run by processor

		}

		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	if requestHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
		extractorHelper.JsonQuery(resp.Content, &extractor)

	} else if requestHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
		extractorHelper.HtmlQuery(resp.Content, &extractor)

	} else if requestHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
		extractorHelper.XmlQuery(resp.Content, &extractor)

	} else if extractor.Type == consts.Boundary {
		extractorHelper.BoundaryQuery(resp.Content, &extractor)
	}

	if interfaceLog == nil { // run by interface
		s.ExtractorRepo.UpdateResult(extractor)
	} else { // run by processor

	}

	_cacheUtils.SetCache(strconv.Itoa(int(projectId)), extractor.Variable, extractor.Result)

	val := _cacheUtils.GetCache(strconv.Itoa(int(projectId)), extractor.Variable)
	logUtils.Infof("%s = %v", extractor.Variable, val)

	return
}

func (s *ExtractorService) ListExtractorVariable(interfaceId int) (variables []serverDomain.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariable(uint(interfaceId))

	return
}
