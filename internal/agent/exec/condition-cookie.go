package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	cookieHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cookie"
)

func ExecCookie(extractor *domain.CookieBase, resp domain.DebugResponse) (err error) {
	err = cookieHelper.Exec(extractor, resp)

	return
}
