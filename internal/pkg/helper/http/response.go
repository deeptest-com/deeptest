package httpHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"strings"
)

func IsJsonResp(resp domain.DebugResponse) (ret bool) {
	ret = strings.Contains(
		strings.ToLower(resp.ContentType.String()),
		strings.ToLower(consts.ContentTypeJSON.String()))

	return
}
