package extractorHelper

import (
	"fmt"
	queryUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

func Extract(extractor *domain.ExtractorBase, resp domain.DebugResponse) (err error) {
	result := ""

	if extractor.Disabled {
		result = ""
		return
	}

	if extractor.Src == consts.Header {
		for _, h := range resp.Headers {
			if h.Name == extractor.Key {
				result = h.Value
				break
			}
		}
	} else {
		if httpHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
			result = queryUtils.JsonQuery(resp.Content, extractor.Expression)

		} else if httpHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
			result = queryUtils.HtmlQuery(resp.Content, extractor.Expression)

		} else if httpHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
			result = queryUtils.XmlQuery(resp.Content, extractor.Expression)

		} else if extractor.Type == consts.Boundary && (extractor.BoundaryStart != "" || extractor.BoundaryEnd != "") {
			result = queryUtils.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
				extractor.BoundaryIndex, extractor.BoundaryIncluded)
		} else if extractor.Type == consts.Regx {
			result = queryUtils.RegxQuery(resp.Content, extractor.Expression)
		}
	}

	extractor.Result = strings.TrimSpace(result)
	extractor.ResultStatus = consts.Pass
	if extractor.Result == "" {
		extractor.ResultStatus = consts.Fail
	}

	_logUtils.Infof(fmt.Sprintf("提取器调试 result:%+v", extractor.Result))

	return
}

func GenDesc(src consts.ExtractorSrc, typ consts.ExtractorType,
	expression, boundaryStart, boundaryEnd string) (ret string) {
	srcDesc := ""
	if src == consts.Header {
		srcDesc = "响应头"
	} else if src == consts.Body {
		srcDesc = "响应体"
	}

	nameDesc := ""
	if typ == consts.Boundary {
		nameDesc = fmt.Sprintf("边界\"%s - %s\"", boundaryStart, boundaryEnd)
	} else if typ == consts.JsonQuery {
		nameDesc = fmt.Sprintf("JSON\"%s\"", expression)
	} else if typ == consts.HtmlQuery {
		nameDesc = fmt.Sprintf("HTML\"%s\"", expression)
	} else if typ == consts.XmlQuery {
		nameDesc = fmt.Sprintf("XML\"%s\"", expression)
	} else if typ == consts.Regx {
		nameDesc = fmt.Sprintf("正则表达式\"%s\"", expression)
	}

	ret = fmt.Sprintf("提取变量 %s%s", srcDesc, nameDesc)

	return
}

func GenResultMsg(po *domain.ExtractorBase) (ret string) {
	desc := GenDesc(po.Src, po.Type, po.Expression, po.BoundaryStart, po.BoundaryEnd)

	po.ResultMsg = fmt.Sprintf("%s， 结果\"%s\"。", desc, po.Result)

	return
}
