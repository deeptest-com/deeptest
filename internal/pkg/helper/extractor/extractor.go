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

func Extract(extractor domain.ExtractorBase, resp domain.DebugResponse) (result string, err error) {
	_logUtils.Infof(fmt.Sprintf("提取器调试 extractor:%+v, resp:%+v", extractor, resp))

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

		} else if extractor.Type == consts.Boundary {
			result = queryUtils.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
				extractor.BoundaryIndex, extractor.BoundaryIncluded)
		}
	}

	result = strings.TrimSpace(result)

	return
}
