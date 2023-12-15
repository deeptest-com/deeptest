package commUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

func ConvertValueForPersistence(obj interface{}) (value string, valueType consts.ExtractorResultType) {
	value = ""
	valueType = consts.ExtractorResultTypeString

	switch obj.(type) {
	case string:
		value = obj.(string)
		valueType = consts.ExtractorResultTypeString

	case float64:
		value = fmt.Sprintf("%v", obj)
		valueType = consts.ExtractorResultTypeNumber

	case bool:
		value = fmt.Sprintf("%t", obj)
		valueType = consts.ExtractorResultTypeBool

	default:
		value = _stringUtils.JsonWithoutHtmlEscaped(obj)
		valueType = consts.ExtractorResultTypeObject
	}

	return
}

func ConvertValueForUse(value interface{}, valueType consts.ExtractorResultType) (obj interface{}, err error) {
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
