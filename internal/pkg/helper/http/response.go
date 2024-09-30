package httpHelper

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"strings"
)

func IsJsonResp(resp domain.DebugResponse) (ret bool) {
	ret = strings.Contains(
		strings.ToLower(resp.ContentType.String()),
		strings.ToLower(consts.ContentTypeJSON.String()))

	return
}

func IsJsonRespType(typ consts.HttpContentType) (ret bool) {
	ret = strings.Contains(
		strings.ToLower(typ.String()),
		strings.ToLower(consts.ContentTypeJSON.String()))

	return
}
