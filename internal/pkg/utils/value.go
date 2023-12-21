package commUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"strconv"
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
		} else if valueType == consts.ExtractorResultTypeString {
			obj = value.(string)
		} else if valueType == consts.ExtractorResultTypeNumber {
			obj, err = strconv.ParseFloat(value.(string), 64)
		} else if valueType == consts.ExtractorResultTypeBool {
			obj, err = strconv.ParseBool(value.(string))
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
