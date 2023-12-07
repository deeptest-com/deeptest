package commUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

func GetValueInfo(obj interface{}) (value interface{}, valueType consts.ExtractorResultType) {
	value = ""
	valueType = consts.ExtractorResultTypeString

	switch obj.(type) {
	case string:
		value = obj.(string)
		valueType = consts.ExtractorResultTypeString

	case float64:
		value = fmt.Sprintf("%d", obj)
		valueType = consts.ExtractorResultTypeNumber

	case bool:
		value = fmt.Sprintf("%t", obj)
		valueType = consts.ExtractorResultTypeBool

	default:
		var err error
		value = _stringUtils.JsonWithoutHtmlEscaped(obj)
		if err != nil {
			value = err.Error()
		}

		valueType = consts.ExtractorResultTypeObject
	}

	return
}

func GetValueObj(value interface{}, valueType consts.ExtractorResultType) (obj interface{}, err error) {
	switch value.(type) {
	case string:
		if valueType == consts.ExtractorResultTypeObject {
			err = json.Unmarshal([]byte(value.(string)), &obj)
			return
		} else if valueType == consts.ExtractorResultTypeString {
			obj = value.(string)
			return
		} else if valueType == consts.ExtractorResultTypeNumber {
			obj = value.(float64)
			return
		} else if valueType == consts.ExtractorResultTypeBool {
			obj = value.(bool)
			return
		}

	case float64:
		obj = value

	case bool:
		obj = value

	default:
		obj = value
	}

	return
}

func ObjectToStrAsVariValue(obj interface{}) (ret string) {
	switch obj.(type) {

	case string:
		ret = obj.(string)

	case float64:
		ret = fmt.Sprintf("%d", obj)

	case bool:
		ret = fmt.Sprintf("%t", obj)

	default:
		ret = _stringUtils.JsonWithoutHtmlEscaped(obj)
	}

	return
}
