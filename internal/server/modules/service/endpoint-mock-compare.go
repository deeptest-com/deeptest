package service

import (
	"fmt"
	queryUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/bitly/go-simplejson"
	"regexp"
	"strconv"
	"strings"
)

type EndpointMockCompareService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
}

func (s *EndpointMockCompareService) CompareBody(expectRequest model.EndpointMockExpectRequest, contentType consts.HttpContentType,
	body string, bodyForm map[string][]string) (ret bool) {

	if httpHelper.IsJsonContent(contentType.String()) { // json
		if expectRequest.SelectType == consts.KeyValue {
			jsn, err := simplejson.NewJson([]byte(body))
			if err != nil {
				return false
			}

			expectValue := expectRequest.Value
			var actualValue interface{}

			actualJson := jsn.Get(expectRequest.Name) // get value of key on first level
			actualValueFloat, err := actualJson.Float64()
			if err == nil {
				actualValue = actualValueFloat
			} else {
				actualValue = actualJson.MustString()
			}

			ret = s.compareObject(actualValue, expectValue, expectRequest.CompareWay)

		} else if expectRequest.SelectType == consts.Xpath { // use xpath
			xpath := expectRequest.Name

			expectValue := expectRequest.Value
			actualValue := queryUtils.JsonQueryWithType(body, xpath)

			ret = s.compareObject(actualValue, expectValue, expectRequest.CompareWay)

		} else if expectRequest.SelectType == consts.FullText || expectRequest.SelectType == "" {
			expectValue := expectRequest.Value
			actualValue := body

			ret = s.compareObject(actualValue, expectValue, expectRequest.CompareWay)
		}

	} else if contentType == consts.ContentTypeFormData {
		if expectRequest.SelectType == consts.KeyValue { // key/value
			expectValue := expectRequest.Value

			items := bodyForm[expectRequest.Name]
			for _, item := range items {
				actualValue := item
				result := s.compareObject(actualValue, expectValue, expectRequest.CompareWay)
				if result {
					return true
				}
			}
		}

	} else if contentType == consts.ContentTypeFormUrlencoded {
		if expectRequest.SelectType == consts.KeyValue { // key/value
			expectValue := expectRequest.Value

			items := bodyForm[expectRequest.Name]
			for _, item := range items {
				actualValue := item
				result := s.compareObject(actualValue, expectValue, expectRequest.CompareWay)
				if result {
					return true
				}
			}
		}

	} else { // xml, html, text etc.
		if expectRequest.SelectType == consts.FullText {
			expectValue := expectRequest.Value
			actualValue := body

			ret = s.compareObject(actualValue, expectValue, expectRequest.CompareWay)
		}

	}

	return
}

func (s *EndpointMockCompareService) CompareString(actualValue interface{}, expectValue string,
	comparator consts.ComparisonOperator) (ret bool) {

	if comparator == consts.Equal {
		ret = fmt.Sprintf("%v", actualValue) == expectValue

	} else if comparator == consts.NotEqual {
		ret = fmt.Sprintf("%v", actualValue) != expectValue

	} else if comparator == consts.Contain {
		ret = strings.Contains(fmt.Sprintf("%v", actualValue), expectValue)

	} else if comparator == consts.NotContain {
		ret = !strings.Contains(fmt.Sprintf("%v", actualValue), expectValue)

	} else if comparator == consts.RegularMatch {
		regx := regexp.MustCompile(expectValue)
		ret = regx.MatchString(fmt.Sprintf("%v", actualValue))

	}

	return
}

func (s *EndpointMockCompareService) compareObject(actualValue interface{}, expectValue string, comparator consts.ComparisonOperator) (ret bool) {
	isFloat := s.isFloat(actualValue)

	if comparator == consts.Equal {
		if isFloat {
			actualFloat, expectFloat, err := s.getFloatValues(actualValue, expectValue)
			ret = err == nil && actualFloat == expectFloat
		} else {
			ret = fmt.Sprintf("%v", actualValue) == expectValue
		}

	} else if comparator == consts.NotEqual {
		if isFloat {
			actualFloat, expectFloat, err := s.getFloatValues(actualValue, expectValue)
			ret = err == nil && actualFloat != expectFloat
		} else {
			ret = fmt.Sprintf("%v", actualValue) != expectValue
		}

	} else if comparator == consts.GreaterThan {
		actualFloat, err2 := s.interfaceToNumber(actualValue)
		expectFloat, err1 := s.strToNumber(expectValue)
		if err1 != nil || err2 != nil {
			return false
		}

		ret = actualFloat > expectFloat

	} else if comparator == consts.GreaterThanOrEqual {
		actualFloat, err2 := s.interfaceToNumber(actualValue)
		expectFloat, err1 := s.strToNumber(expectValue)
		if err1 != nil || err2 != nil {
			return false
		}

		ret = actualFloat >= expectFloat

	} else if comparator == consts.LessThan {
		actualFloat, err2 := s.interfaceToNumber(actualValue)
		expectFloat, err1 := s.strToNumber(expectValue)
		if err1 != nil || err2 != nil {
			return false
		}

		ret = actualFloat < expectFloat

	} else if comparator == consts.LessThanOrEqual {
		actualFloat, err2 := s.interfaceToNumber(actualValue)
		expectFloat, err1 := s.strToNumber(expectValue)
		if err1 != nil || err2 != nil {
			return false
		}

		ret = actualFloat <= expectFloat

	} else if comparator == consts.Contain {
		ret = strings.Contains(fmt.Sprintf("%v", actualValue), expectValue)

	} else if comparator == consts.NotContain {
		ret = !strings.Contains(fmt.Sprintf("%v", actualValue), expectValue)

	} else if comparator == consts.RegularMatch {
		regx := regexp.MustCompile(expectValue)
		ret = regx.MatchString(fmt.Sprintf("%v", actualValue))

	}

	return
}

func (s *EndpointMockCompareService) strToNumber(val string) (ret float64, err error) {
	ret, err = strconv.ParseFloat(val, 64)

	return
}

func (s *EndpointMockCompareService) interfaceToNumber(val interface{}) (ret float64, err error) {
	str := fmt.Sprintf("%v", val)
	ret, err = s.strToNumber(str)

	return
}

func (s *EndpointMockCompareService) getFloatValues(actual interface{}, expect string) (actualFloat, expectFloat float64, err error) {
	actualFloat = actual.(float64)
	expectFloat, err = strconv.ParseFloat(expect, 64)

	return
}

func (s *EndpointMockCompareService) isFloat(val interface{}) (ret bool) {
	switch val.(type) {
	case float64:
		return true
	default:
	}

	return false
}
