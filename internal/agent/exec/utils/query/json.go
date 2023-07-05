package queryUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/antchfx/jsonquery"
	"strings"
)

func JsonQuery(content string, expression string) (result string) {
	doc, err := jsonquery.Parse(strings.NewReader(content))
	_logUtils.Infof(fmt.Sprintf("提取器调试 doc:%+v, err:%+v", doc, err))
	if err != nil {
		result = consts.ContentErr
		return
	}
	elem, err := jsonquery.Query(doc, expression)
	_logUtils.Infof(fmt.Sprintf("提取器调试 elem:%+v, err:%+v", elem, err))

	if err != nil || elem == nil {
		result = consts.ExtractorErr
		return
	}

	obj := elem.Value()
	_logUtils.Infof(fmt.Sprintf("提取器调试 obj:%+v", obj))

	switch obj.(type) {
	case string:
		result = obj.(string)
	default:
		bytes, err := json.Marshal(obj)
		if err != nil {
			result = err.Error()
		} else {
			result = string(bytes)
		}
	}

	return
}
