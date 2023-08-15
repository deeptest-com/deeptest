package cookieHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

func Exec(cookie *domain.CookieBase, resp domain.DebugResponse) (err error) {
	for _, item := range resp.Cookies {
		if item.Name == cookie.CookieName && (cookie.CookieDomain == "" || item.Domain == cookie.CookieDomain) {
			cookie.Result = fmt.Sprintf("%v", item.Value)
			break
		}
	}

	cookie.Result = strings.TrimSpace(cookie.Result)
	cookie.ResultStatus = consts.Pass
	if cookie.Result == "" {
		cookie.ResultStatus = consts.Fail
	}

	_logUtils.Infof(fmt.Sprintf("Cookie调试 result: %+v", cookie.Result))

	return
}

func GenDesc(cookieName, variableName string) (ret string) {
	ret = fmt.Sprintf("<b>获取Cookie&nbsp;%s</b>&nbsp;&nbsp;赋予变量%s", cookieName, variableName)

	return
}

func GenResultMsg(po *domain.CookieBase) (ret string) {
	desc := GenDesc(po.CookieName, po.VariableName)

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
