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
	ok := true
	switch value.(type) {
	case string:
		if valueType == consts.ExtractorResultTypeObject {
			err = json.Unmarshal([]byte(value.(string)), &obj)
			if err != nil {
				ok = false
			}
		} else if valueType == consts.ExtractorResultTypeString {
			obj, ok = value.(string)
		} else if valueType == consts.ExtractorResultTypeNumber {
			obj, ok = value.(float64)
		} else if valueType == consts.ExtractorResultTypeBool {
			obj, ok = value.(bool)
		}
		//类型转换失败返回默认数据
		if !ok {
			return value, err
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
