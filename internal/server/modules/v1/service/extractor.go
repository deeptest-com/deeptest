package service

import (
	"encoding/json"
	_cacheUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/extractor"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
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

func (s *ExtractorService) Create(extractor *model.InterfaceExtractor) (err error) {
	_, err = s.ExtractorRepo.Save(extractor)

	return
}

func (s *ExtractorService) Update(extractor *model.InterfaceExtractor) (err error) {
	_, err = s.ExtractorRepo.Save(extractor)

	return
}

func (s *ExtractorService) Delete(reqId uint) (err error) {
	err = s.ExtractorRepo.Delete(reqId)

	return
}

func (s *ExtractorService) ExtractByInterface(interfaceId uint, resp serverDomain.InvocationResponse, projectId int) (err error) {
	extractors, _ := s.ExtractorRepo.List(interfaceId)

	for _, extractor := range extractors {
		s.Extract(extractor, resp, projectId)
	}

	return
}

func (s *ExtractorService) Extract(extractor model.InterfaceExtractor, resp serverDomain.InvocationResponse,
	projectId int) (err error) {
	if extractor.Disabled {
		extractor.Result = ""
		s.ExtractorRepo.UpdateResult(extractor)
		return
	}

	if extractor.Src == serverConsts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Key {
				extractor.Result = h.Value
				break
			}
		}

		s.ExtractorRepo.UpdateResult(extractor)
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	if requestHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == serverConsts.XmlQuery {
		extractorHelper.ParserXPath(resp.Content, &extractor)

	} else if requestHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == serverConsts.HtmlQuery {
		extractorHelper.ParserCssSelector(resp.Content, &extractor)

	} else if requestHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == serverConsts.JsonQuery {
		extractorHelper.ParserJsonPath(resp.Content, &extractor)

	} else if extractor.Type == serverConsts.Boundary {
		extractorHelper.ParserBoundary(resp.Content, &extractor)

	}

	s.ExtractorRepo.UpdateResult(extractor)

	_cacheUtils.SetCache(strconv.Itoa(projectId), extractor.Variable, extractor.Result)

	val := _cacheUtils.GetCache(strconv.Itoa(projectId), extractor.Variable)
	logUtils.Infof("%s = %v", extractor.Variable, val)

	return
}

func (s *ExtractorService) ListExtractorVariable(interfaceId int) (variables []serverDomain.Variable, err error) {
	variables, err = s.ExtractorRepo.ListExtractorVariable(uint(interfaceId))

	return
}
