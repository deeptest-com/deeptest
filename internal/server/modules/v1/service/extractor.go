package service

import (
	"encoding/json"
	"fmt"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/oliveagle/jsonpath"
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

func (s *ExtractorService) ExtractByInterface(interfaceId uint, resp serverDomain.InvocationResponse) (err error) {
	extractors, _ := s.ExtractorRepo.List(interfaceId)

	for _, extractor := range extractors {
		s.Extract(extractor, resp)
	}

	return
}

func (s *ExtractorService) Extract(extractor model.InterfaceExtractor, resp serverDomain.InvocationResponse) (err error) {
	if extractor.Src == serverConsts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Expression {
				extractor.Result = h.Value
				break
			}
		}

		s.ExtractorRepo.UpdateResult(extractor)
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	if extractor.Type == serverConsts.JsonPath {
		pat, _ := jsonpath.Compile(extractor.Expression)
		result, err := pat.Lookup(jsonData)

		if err == nil && result != nil {
			extractor.Result = fmt.Sprintf("%v", result)
		}
	}

	s.ExtractorRepo.UpdateResult(extractor)

	serverConsts.EnvVar.Store(extractor.Variable, extractor.Result)

	return
}
