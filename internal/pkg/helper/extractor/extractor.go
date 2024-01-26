package extractorHelper

import (
	"fmt"
	queryUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

func Extract(extractor *domain.ExtractorBase, resp domain.DebugResponse) (err error) {
	var result interface{}
	resultType := consts.ExtractorResultTypeString

	if extractor.Src == consts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Key {
				result = h.Value
				break
			}
		}
	} else if extractor.Src == consts.Cookie {
		for _, cookie := range resp.Cookies {
			if cookie.Name == extractor.Key {
				result = fmt.Sprintf("%v", cookie.Value)
				break
			}
		}

	} else {
		extractor.Expression = strings.TrimSpace(extractor.Expression)

		if httpHelper.IsJsonContent(resp.ContentType.String()) { // json path
			if extractor.Type == consts.JSONPath {
				result, resultType, err = queryUtils.JsonPath(resp.Content, extractor.Expression)
			} else if extractor.Type == consts.JsonQuery {
				result, resultType = queryUtils.JsonQuery(resp.Content, extractor.Expression)
			}

		} else if httpHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery { // html query
			result = queryUtils.HtmlQuery(resp.Content, extractor.Expression)

		} else if httpHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery { // html query
			result = queryUtils.XmlQuery(resp.Content, extractor.Expression)

		} else if extractor.Type == consts.Boundary && (extractor.BoundaryStart != "" || extractor.BoundaryEnd != "") { // boundary
			result = queryUtils.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
				extractor.BoundaryIndex, extractor.BoundaryIncluded)

		} else if extractor.Type == consts.Regx { // regx
			result = queryUtils.RegxQuery(resp.Content, extractor.Expression)

		}
	}

	extractor.ResultType = resultType
	extractor.Result = valueUtils.InterfaceToStr(result)
	extractor.ResultStatus = consts.Pass
	if extractor.Result == "" {
		extractor.ResultStatus = consts.Fail
	}

	_logUtils.Infof(fmt.Sprintf("提取器调试 result:%+v", extractor.Result))

	return
}

func GenDesc(varName string, src consts.ExtractorSrc, key string, typ consts.ExtractorType,
	expression, boundaryStart, boundaryEnd string) (ret string) {
	srcDesc := ""
	if src == consts.Header {
		srcDesc = "响应头"
	} else if src == consts.Cookie {
		srcDesc = "Cookie"
	} else if src == consts.Body {
		srcDesc = "响应体"
	}

	if src != consts.Body {
		ret = fmt.Sprintf("<b>提取变量&nbsp;%s</b>&nbsp;&nbsp;%s&nbsp;%s", varName, srcDesc, key)
		return
	}

	name := ""
	expr := ""
	if typ == consts.Boundary {
		name = fmt.Sprintf("边界选择器")
		expr = fmt.Sprintf("%s ~ %s", getLimitStr(boundaryStart, 26), getLimitStr(boundaryEnd, 26))

	} else if typ == consts.JSONPath {
		name = fmt.Sprintf("JSONPath")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.JsonQuery {
		name = fmt.Sprintf("JSON查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.HtmlQuery {
		name = fmt.Sprintf("HTML查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.XmlQuery {
		name = fmt.Sprintf("XML查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.Regx {
		name = fmt.Sprintf("正则表达式")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))
	}

	ret = fmt.Sprintf("提取变量%s&nbsp;&nbsp;%s&nbsp;%s（%s）", varName, srcDesc, name, expr)

	return
}

func GenDescForCheckpoint(typ consts.ExtractorType, expression string) (ret string) {
	srcDesc := ""
	srcDesc = "响应体"

	name := ""
	expr := ""
	if typ == consts.JSONPath {
		name = fmt.Sprintf("JSONPath")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.JsonQuery {
		name = fmt.Sprintf("JSON查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.HtmlQuery {
		name = fmt.Sprintf("HTML查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.XmlQuery {
		name = fmt.Sprintf("XML查询")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))

	} else if typ == consts.Regx {
		name = fmt.Sprintf("正则表达式")
		expr = fmt.Sprintf("%s", getLimitStr(expression, 50))
	}

	ret = fmt.Sprintf("%s %s(%s)", srcDesc, name, expr)

	return
}

func GenResultMsg(po *domain.ExtractorBase) (ret string) {
	desc := GenDesc(po.Variable, po.Src, po.Key, po.Type, po.Expression, po.BoundaryStart, po.BoundaryEnd)

	po.ResultMsg = fmt.Sprintf("%s，结果\"%s\"。", desc, po.Result)

	return
}

func getLimitStr(str string, limit int) (ret string) {
	if len(str) <= limit-3 {
		return str
	}

	ret = str[:limit-3] + "..."

	return
}
